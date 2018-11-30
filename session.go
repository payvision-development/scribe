package scribe

import (
	"fmt"
	"time"

	"github.com/payvision-development/scribe/freshservice"
	"github.com/payvision-development/scribe/vss"
	"github.com/payvision-development/scribe/release"
)

type state struct {
	ChangeID  int
	LastEvent *vss.Event
}

// Session func
func Session(ch chan *vss.Event, c release.Changer) {

	s := state{}

	for {
		select {
		case event := <-ch:

			switch et := event.EventType; et {
			case "ms.vss-release.deployment-started-event":

				c.C

				change, err := createChange(fs, event)
				if err != nil {
					fmt.Println(err)
					return
				}

				s.ChangeID = change.Item.ItilChange.DisplayID

				if nil != s.LastEvent && "ms.vss-release.deployment-approval-pending-event" == s.LastEvent.EventType {
					var status freshservice.Status

					if "preDeploy" == s.LastEvent.ApprovalType {
						status = freshservice.StatusAwaitingApproval
					} else {
						status = freshservice.StatusPendingReview
					}

					err := updateChange(fs, s.ChangeID, event.DetailedMessageHTML, status)
					if err != nil {
						fmt.Println(err)
						return
					}
				}

			case "ms.vss-release.deployment-approval-pending-event":

				if 0 != s.ChangeID {
					var status freshservice.Status

					if "preDeploy" == event.ApprovalType {
						status = freshservice.StatusAwaitingApproval
					} else {
						status = freshservice.StatusPendingReview
					}

					err := updateChange(fs, s.ChangeID, event.DetailedMessageHTML, status)
					if err != nil {
						fmt.Println(err)
						return
					}
				}

			case "ms.vss-release.deployment-approval-completed-event":

				if 0 != s.ChangeID {
					var status freshservice.Status

					if "preDeploy" == event.ApprovalType {
						status = freshservice.StatusPendingRelease
					} else {
						status = freshservice.StatusOpen
					}

					err := updateChange(fs, s.ChangeID, event.DetailedMessageHTML, status)
					if err != nil {
						fmt.Println(err)
						return
					}
				}

			case "ms.vss-release.deployment-completed-event":

				if 0 != s.ChangeID {
					err := updateChange(fs, s.ChangeID, event.DetailedMessageHTML, freshservice.StatusClosed)
					if err != nil {
						fmt.Println(err)
						return
					}
				}
			}

			s.LastEvent = event

		case <-time.After(5000 * time.Millisecond):

			if 0 != s.ChangeID {
				err := updateChange(fs, s.ChangeID, "Deployment timeout<br>Status: Failed", freshservice.StatusClosed)
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			return
		}
	}
}

func createChange(fs *freshservice.Freshservice, event *vss.Event) (*freshservice.ItilChange, error) {

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
		return nil, err
	}

	return change, nil
}

func updateChange(fs *freshservice.Freshservice, id int, msg string, status freshservice.Status) error {

	_, err := fs.AddChangeNote(id, msg)
	if err != nil {
		return err
	}

	_, err = fs.UpdateChangeStatus(id, status)
	if err != nil {
		return err
	}

	return nil
}
