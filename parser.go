package scribe

import (
	"hash/fnv"
	"strconv"
	"strings"

	"github.com/payvision-development/scribe/vss"
	"github.com/valyala/fastjson"
)

// Parser func
func Parser(b []byte) (*vss.Event, error) {

	var p fastjson.Parser
	var re vss.Event

	v, err := p.Parse(string(b))
	if err != nil {
		return nil, err
	}

	re.SubscriptionID = string(v.GetStringBytes("subscriptionId"))
	re.NotificationID = v.GetInt("notificationId")
	re.ID = string(v.GetStringBytes("id"))
	re.EventType = string(v.GetStringBytes("eventType"))
	re.PublisherID = string(v.GetStringBytes("publisherId"))
	re.Message = string(v.Get("message").GetStringBytes("text"))
	re.DetailedMessageHTML = string(v.Get("detailedMessage").GetStringBytes("html"))
	re.ResourceVersion = string(v.GetStringBytes("resourceVersion"))
	re.Timestamp = string(v.GetStringBytes("createdDate"))
	re.ProjectID = string(v.Get("resource").Get("project").GetStringBytes("id"))
	re.ProjectName = string(v.Get("resource").Get("project").GetStringBytes("name"))
	re.ProjectURL = string(v.Get("resourceContainers").Get("project").GetStringBytes("baseUrl"))
	re.ServerURL = string(v.Get("resourceContainers").Get("server").GetStringBytes("baseUrl"))
	re.CollectionURL = string(v.Get("resourceContainers").Get("collection").GetStringBytes("baseUrl"))

	if strings.Compare(re.ProjectID, "00000000-0000-0000-0000-000000000000") == 0 {

		re.EnvironmentID = v.Get("resource").Get("environment").GetInt("id")
		re.EnvironmentName = string(v.Get("resource").Get("environment").GetStringBytes("name"))
		re.ReleaseID = v.Get("resource").Get("environment").Get("release").GetInt("id")
		re.ReleaseName = string(v.Get("resource").Get("environment").Get("release").GetStringBytes("name"))
		re.Status = string(v.Get("resource").Get("environment").GetStringBytes("status"))

	} else {

		switch et := re.EventType; et {
		case "ms.vss-release.deployment-started-event", "ms.vss-release.deployment-completed-event":

			re.EnvironmentID = v.Get("resource").Get("environment").GetInt("id")
			re.EnvironmentName = string(v.Get("resource").Get("environment").GetStringBytes("name"))
			re.ReleaseID = v.Get("resource").Get("environment").Get("release").GetInt("id")
			re.ReleaseName = string(v.Get("resource").Get("environment").Get("release").GetStringBytes("name"))
			re.ReleaseURL = string(v.Get("resource").Get("environment").Get("release").GetStringBytes("url"))
			re.Status = string(v.Get("resource").Get("environment").GetStringBytes("status"))

		case "ms.vss-release.deployment-approval-pending-event", "ms.vss-release.deployment-approval-completed-event":

			for _, k := range v.Get("resource").Get("release").GetArray("environments") {
				if strings.Compare(string(k.GetStringBytes("status")), "inProgress") == 0 {
					re.EnvironmentID = k.GetInt("id")
					re.EnvironmentName = string(k.GetStringBytes("name"))
					re.ReleaseID = k.Get("release").GetInt("id")
					re.ReleaseName = string(k.Get("release").GetStringBytes("name"))
					re.ReleaseURL = string(k.Get("release").GetStringBytes("url"))
					re.Status = string(k.GetStringBytes("status"))
					re.ApprovalType = string(v.Get("resource").Get("approval").GetStringBytes("approvalType"))

					break
				}
			}

		default:
		}
	}

	re.ReleaseTrackingCode = hash(strings.Join([]string{strconv.Itoa(re.EnvironmentID), re.ProjectID}, "-"))

	return &re, nil
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
