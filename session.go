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
func Session(tc uint32, ch chan *vss.Event, c release.Changer) {

	s := state{}

	for {
		select {
		case event := <-ch:

			switch et := event.EventType; et {

			case vss.DeploymentStartedEvent:

				fmt.Printf("[Release: %v] Event: %v\n", tc, vss.DeploymentStartedEvent)

				id, err := c.Create(event.ReleaseName, event.EnvironmentName, event.DetailedMessageHTML, event.Timestamp)
				if err != nil {
					fmt.Println(err)
				} else {
					s.ChangeID = id

					if nil != s.LastEvent && vss.DeploymentApprovalPendingEvent == s.LastEvent.EventType {
						var status int

						if "preDeploy" == s.LastEvent.ApprovalType {
							status = freshservice.StatusAwaitingApproval
						} else {
							status = freshservice.StatusPendingReview
						}

						err := c.Update(event.DetailedMessageHTML, string(status))
						if err != nil {
							fmt.Println(err)
						}
					}
				}

			case vss.DeploymentApprovalPendingEvent:

				fmt.Printf("[Release: %v] Event: %v\n", tc, vss.DeploymentApprovalPendingEvent)

				if 0 != s.ChangeID {
					var status int

					if "preDeploy" == event.ApprovalType {
						status = freshservice.StatusAwaitingApproval
					} else {
						status = freshservice.StatusPendingReview
					}

					err := c.Update(event.DetailedMessageHTML, string(status))
					if err != nil {
						fmt.Println(err)
					}
				}

			case vss.DeploymentApprovalCompletedEvent:

				fmt.Printf("[Release: %v] Event: %v\n", tc, vss.DeploymentApprovalCompletedEvent)

				if 0 != s.ChangeID {
					var status int

					if "preDeploy" == event.ApprovalType {
						status = freshservice.StatusPendingRelease
					} else {
						status = freshservice.StatusOpen
					}

					err := c.Update(event.DetailedMessageHTML, string(status))
					if err != nil {
						fmt.Println(err)
					}
				}

			case vss.DeploymentCompletedEvent:

				fmt.Printf("[Release: %v] Event: %v\n", tc, vss.DeploymentCompletedEvent)

				if 0 != s.ChangeID {
					err := c.Update(event.DetailedMessageHTML, string(freshservice.StatusClosed))
					if err != nil {
						fmt.Println(err)
					}
				}
			}

			s.LastEvent = event

		case <-time.After(10000 * time.Millisecond):

			fmt.Printf("[Release: %v] Event: %v\n", tc, "Deployment timeout")

			if 0 != s.ChangeID {
				err := c.Update("Deployment timeout<br>Status: Failed", string(freshservice.StatusClosed))
				if err != nil {
					fmt.Println(err)
				}
			}

			return
		}
	}
}
