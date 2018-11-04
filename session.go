package scribe

import (
	"fmt"
	"log"
	"time"

	"github.com/payvision-development/scribe/freshservice"
	"github.com/payvision-development/scribe/vss"
)

type state struct {
	ChangeID  int
	LastEvent *vss.Event
}

// Session func
func Session(ch chan *vss.Event) {

	s := state{}
	fs := freshservice.NewClient("https://foo.freshservice.com", "key")

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
					log.Fatal(err.Error())
				}

				s.ChangeID = change.Item.ItilChange.DisplayID

			case "ms.vss-release.deployment-approval-pending-event":

			case "ms.vss-release.deployment-approval-completed-event":

			case "ms.vss-release.deployment-completed-event":

			}

			s.LastEvent = event

		case <-time.After(5000 * time.Millisecond):
			fmt.Println("timeout!")
			return
		}
	}
}
