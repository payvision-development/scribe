package vss

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/payvision-development/scribe/httpclient"
)

// TFS struct
type TFS struct {
	ServerURL     string
	CollectionURL string
	APIKey        string
}

// NewClient func
func NewClient(serverURL string, collectionURL string, apikey string) *TFS {
	tfs := new(TFS)
	tfs.ServerURL = serverURL
	tfs.CollectionURL = collectionURL
	tfs.APIKey = apikey
	return tfs
}

// GetRelease func
func (tfs *TFS) GetRelease(projectID string, releaseID int) (*Release, error) {

	url := fmt.Sprintf("%v/%v/_apis/release/releases/%v/?api-version=4.1-preview.6", tfs.CollectionURL, projectID, releaseID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth("", tfs.APIKey)
	req.Header.Set("Content-Type", "application/json")
	res, err := httpclient.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var r Release

	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// GetWorkItems func
func (tfs *TFS) GetWorkItems(projectID string, buildID string) (*WorkItems, error) {

	url := fmt.Sprintf("%v/%v/_apis/build/builds/%v/workitems?api-version=4.1", tfs.CollectionURL, projectID, buildID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth("", tfs.APIKey)
	req.Header.Set("Content-Type", "application/json")
	res, err := httpclient.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var r WorkItems

	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// GetWorkItem func
func (tfs *TFS) GetWorkItem(workItemID string) (*WorkItem, error) {

	url := fmt.Sprintf("%v/_apis/wit/workitems/%v?api-version=4.1", tfs.CollectionURL, workItemID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth("", tfs.APIKey)
	req.Header.Set("Content-Type", "application/json")
	res, err := httpclient.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var r WorkItem

	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
