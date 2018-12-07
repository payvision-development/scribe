package vss

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/payvision-development/scribe/httpclient"
)

// VSTS struct
type VSTS struct {
	URL    string
	APIKey string
}

// NewClient func
func NewClient(url string, apikey string) *VSTS {
	fs := new(VSTS)
	fs.URL = url
	fs.APIKey = apikey
	return fs
}

// Release func
func (fs *VSTS) Release(ReleaseURL string) (*Release, error) {

	req, err := http.NewRequest("GET", ReleaseURL, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth("", fs.APIKey)
	req.Header.Set("Content-Type", "application/json")
	res, err := httpclient.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var resRelease Release

	err = json.Unmarshal(res, &resRelease)
	if err != nil {
		return nil, err
	}

	return &resRelease, nil
}

// WorkItems func
func (fs *VSTS) WorkItems(r *Release) (*WorkItems, error) {

	// TODO: Check in case several artifacts were associated to the build
	url := fmt.Sprintf(fs.URL + r.Artifacts[0].DefinitionReference.Project.Name + "/_apis/build/builds/" + r.Artifacts[0].DefinitionReference.Version.ID + "/workitems?api-version=2.0")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth("", fs.APIKey)
	req.Header.Set("Content-Type", "application/json")
	res, err := httpclient.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var resWorkItems WorkItems

	err = json.Unmarshal(res, &resWorkItems)
	if err != nil {
		return nil, err
	}

	return &resWorkItems, nil
}

// WorkItem func
func (fs *VSTS) WorkItem(URL string) (*WorkItem, error) {

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth("", fs.APIKey)
	req.Header.Set("Content-Type", "application/json")
	res, err := httpclient.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var resWorkItem WorkItem

	err = json.Unmarshal(res, &resWorkItem)
	if err != nil {
		return nil, err
	}

	return &resWorkItem, nil
}
