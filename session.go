package scribe

import (
	"fmt"
	"strings"
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
func Session(tc uint32, ch chan *vss.Event, c release.Changer, v *vss.VSTS) {

	s := state{}

	for {
		select {
		case event := <-ch:
			switch et := event.EventType; et {

			case vss.DeploymentStartedEvent:

				fmt.Printf("[Release: %v] Event: %v\n", tc, vss.DeploymentStartedEvent)

				var descriptionHTML strings.Builder

				if v != nil {
					descriptionHTML.WriteString("<br><b>Work Items to deploy</b><br>")

					release, err := v.Release(event.ReleaseURL)
					if err != nil {
						fmt.Println(err)
					} else {
						workItems, err := v.WorkItems(release)
						if err != nil {
							fmt.Println(err)
						} else {
							for _, workItem := range workItems.Value {
								w, err := v.WorkItem(workItem.URL)
								if err != nil {
									fmt.Println(err)
								} else {
									descriptionHTML.WriteString("<br>[" + w.Fields.SystemWorkItemType + "] <a href='" + w.Links.HTML.Href + "'>" + w.Fields.SystemTitle + "</a>")
								}
							}
						}
					}
				} else {
					descriptionHTML.WriteString("<br><b>No Work Items associated to this deploy</b><br>")
				}

				id, err := c.Create(event.ReleaseName, event.EnvironmentName, descriptionHTML.String(), event.Timestamp)
				if err != nil {
					fmt.Println(err)
				} else {
					s.ChangeID = id

					err := c.Update(event.DetailedMessageHTML, freshservice.StatusOpen)
					if err != nil {
						fmt.Println(err)
					} else {
						if nil != s.LastEvent && vss.DeploymentApprovalPendingEvent == s.LastEvent.EventType {
							var status int

							if "preDeploy" == s.LastEvent.ApprovalType {
								status = freshservice.StatusAwaitingApproval
							} else {
								status = freshservice.StatusPendingReview
							}

							err := c.Update(s.LastEvent.DetailedMessageHTML, status)
							if err != nil {
								fmt.Println(err)
							}
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

					err := c.Update(event.DetailedMessageHTML, status)
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

					err := c.Update(event.DetailedMessageHTML, status)
					if err != nil {
						fmt.Println(err)
					}
				}

			case vss.DeploymentCompletedEvent:

				fmt.Printf("[Release: %v] Event: %v\n", tc, vss.DeploymentCompletedEvent)

				if 0 != s.ChangeID {
					err := c.Update(event.DetailedMessageHTML, freshservice.StatusClosed)
					if err != nil {
						fmt.Println(err)
					}
				}
			}

			s.LastEvent = event

		case <-time.After(30 * time.Minute):

			if 0 != s.ChangeID && s.LastEvent.EventType != vss.DeploymentCompletedEvent {

				fmt.Printf("[Release: %v] Event: %v\n", tc, "Deployment timeout")

				err := c.Update("Deployment timeout<br>Status: Failed", freshservice.StatusClosed)
				if err != nil {
					fmt.Println(err)
				}
			}

			return
		}
	}
}
