package scribe

import (
	"fmt"
	"log"
	"time"

	"github.com/payvision-development/scribe/freshservice"
	"github.com/payvision-development/scribe/vss"
	"github.com/valyala/fastjson"
)

type state struct {
	ChangeID  int
	LastEvent *vss.Event
}

// Session func
func Session(ch chan *vss.Event) {

	s := state{}
	fs := freshservice.NewClient("https://foo.freshservice.com")

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

				change, _ := fs.CreateChange(&c)

				var p fastjson.Parser
				v, err := p.Parse(string(change))
				if err != nil {
					log.Fatal(err)
				}

				s.ChangeID = v.Get("item").Get("itil_change").GetInt("display_id")

			case "ms.vss-release.deployment-approval-pending-event":

			case "ms.vss-release.deployment-approval-completed-event":

			case "ms.vss-release.deployment-completed-event":

			}

			s.LastEvent = event

		case <-time.After(5000 * time.Millisecond):
			fmt.Println("timeout puto!")
			return
		}
	}
}
