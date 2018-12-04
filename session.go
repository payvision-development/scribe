package scribe

import (
	"fmt"
	"time"

	"github.com/payvision-development/scribe/freshservice"
	"github.com/payvision-development/scribe/release"
	"github.com/payvision-development/scribe/vss"
)

type state struct {
	ChangeID  int64
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

				id, err := c.Create(event.ReleaseName, event.EnvironmentName, event.DetailedMessageHTML, event.Timestamp)
				if err != nil {
					fmt.Println(err)
					return
				}

				s.ChangeID = id

				if nil != s.LastEvent && "ms.vss-release.deployment-approval-pending-event" == s.LastEvent.EventType {
					var status freshservice.Status

					if "preDeploy" == s.LastEvent.ApprovalType {
						status = freshservice.StatusAwaitingApproval
					} else {
						status = freshservice.StatusPendingReview
					}

					err := c.Update(event.DetailedMessageHTML, string(status))
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

					err := c.Update(event.DetailedMessageHTML, string(status))
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

					err := c.Update(event.DetailedMessageHTML, string(status))
					if err != nil {
						fmt.Println(err)
						return
					}
				}

			case "ms.vss-release.deployment-completed-event":

				if 0 != s.ChangeID {
					err := c.Update(event.DetailedMessageHTML, string(freshservice.StatusClosed))
					if err != nil {
						fmt.Println(err)
						return
					}
				}
			}

			s.LastEvent = event

		case <-time.After(5000 * time.Millisecond):

			if 0 != s.ChangeID {
				err := c.Update("Deployment timeout<br>Status: Failed", string(freshservice.StatusClosed))
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			return
		}
	}
}
