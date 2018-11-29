package scribe

import (
	"fmt"
	"time"

	"github.com/payvision-development/scribe/freshservice"
	"github.com/payvision-development/scribe/vss"
)

var open = 1
var planning = 2
var awaitingApproval = 3
var pendingRelease = 4
var pendingReview = 5
var closed = 6

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
					Status:           4, // Pending release
					Priority:         2, // Medium
					ChangeType:       2, // Standard
					Risk:             2, // Medium
					Impact:           2, // Medium
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
						_, err = fs.UpdateChangeStatus(s.ChangeID, awaitingApproval)
					} else {
						_, err = fs.UpdateChangeStatus(s.ChangeID, pendingReview)
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
						_, err = fs.UpdateChangeStatus(s.ChangeID, awaitingApproval)
					} else {
						_, err = fs.UpdateChangeStatus(s.ChangeID, pendingReview)
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
						_, err = fs.UpdateChangeStatus(s.ChangeID, pendingRelease)
					} else {
						_, err = fs.UpdateChangeStatus(s.ChangeID, open)
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

					_, err = fs.UpdateChangeStatus(s.ChangeID, closed)
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

				_, err = fs.UpdateChangeStatus(s.ChangeID, closed)
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			return
		}
	}
}
