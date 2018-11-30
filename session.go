package scribe

import (
	"fmt"
	"time"

	"github.com/payvision-development/scribe/freshservice"
	"github.com/payvision-development/scribe/vss"
)

type state struct {
	ChangeID  int
	LastEvent *vss.Event
}

// Session func
func Session(ch chan *vss.Event, url string, apikey string) {

	s := state{}
	fs := freshservice.NewClient(url, apikey)

	for {
		select {
		case event := <-ch:

			switch et := event.EventType; et {
			case "ms.vss-release.deployment-started-event":

				c := freshservice.Change{
					Email:            "hulk@outerspace.com",
					Subject:          "[Release Management] Deployment of release " + event.ReleaseName + " to environment " + event.EnvironmentName,
					DescriptionHTML:  event.DetailedMessageHTML,
					Status:           freshservice.StatusPendingRelease,
					Priority:         freshservice.PriorityMedium,
					ChangeType:       freshservice.TypeStandard,
					Risk:             freshservice.RiskMedium,
					Impact:           freshservice.ImpactMedium,
					PlannedStartDate: event.Timestamp,
					PlannedEndDate:   event.Timestamp,
				}

				change, err := fs.CreateChange(&c)
				if err != nil {
					fmt.Println(err)
					return
				}

				s.ChangeID = change.Item.ItilChange.DisplayID

				if nil != s.LastEvent && "ms.vss-release.deployment-approval-pending-event" == s.LastEvent.EventType {
					_, err := fs.AddChangeNote(s.ChangeID, s.LastEvent.DetailedMessageHTML)
					if err != nil {
						fmt.Println(err)
						return
					}

					if "preDeploy" == s.LastEvent.ApprovalType {
						_, err = fs.UpdateChangeStatus(s.ChangeID, freshservice.StatusAwaitingApproval)
					} else {
						_, err = fs.UpdateChangeStatus(s.ChangeID, freshservice.StatusPendingReview)
					}

					if err != nil {
						fmt.Println(err)
						return
					}
				}

			case "ms.vss-release.deployment-approval-pending-event":

				if 0 != s.ChangeID {
					_, err := fs.AddChangeNote(s.ChangeID, event.DetailedMessageHTML)
					if err != nil {
						fmt.Println(err)
						return
					}

					if "preDeploy" == event.ApprovalType {
						_, err = fs.UpdateChangeStatus(s.ChangeID, freshservice.StatusAwaitingApproval)
					} else {
						_, err = fs.UpdateChangeStatus(s.ChangeID, freshservice.StatusPendingReview)
					}

					if err != nil {
						fmt.Println(err)
						return
					}
				}

			case "ms.vss-release.deployment-approval-completed-event":

				if 0 != s.ChangeID {
					_, err := fs.AddChangeNote(s.ChangeID, event.DetailedMessageHTML)
					if err != nil {
						fmt.Println(err)
						return
					}

					if "preDeploy" == event.ApprovalType {
						_, err = fs.UpdateChangeStatus(s.ChangeID, freshservice.StatusPendingRelease)
					} else {
						_, err = fs.UpdateChangeStatus(s.ChangeID, freshservice.StatusOpen)
					}

					if err != nil {
						fmt.Println(err)
						return
					}
				}

			case "ms.vss-release.deployment-completed-event":

				if 0 != s.ChangeID {
					_, err := fs.AddChangeNote(s.ChangeID, event.DetailedMessageHTML)
					if err != nil {
						fmt.Println(err)
						return
					}

					_, err = fs.UpdateChangeStatus(s.ChangeID, freshservice.StatusClosed)
					if err != nil {
						fmt.Println(err)
						return
					}
				}
			}

			s.LastEvent = event

		case <-time.After(5000 * time.Millisecond):
			if 0 != s.ChangeID {
				_, err := fs.AddChangeNote(s.ChangeID, "Deployment timeout<br>Status: Failed")
				if err != nil {
					fmt.Println(err)
					return
				}

				_, err = fs.UpdateChangeStatus(s.ChangeID, freshservice.StatusClosed)
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			return
		}
	}
}
