package freshservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/payvision-development/scribe/httpclient"
)

// Freshservice struct
type Freshservice struct {
	URL    string
	Email  string
	APIKey string
}

// NewClient func
func NewClient(url string, email string, apikey string) *Freshservice {
	fs := new(Freshservice)
	fs.URL = url
	fs.Email = email
	fs.APIKey = apikey
	return fs
}

// CreateChange func
func (fs *Freshservice) CreateChange(c *RequestItilChange) (*ResponseItilChange, error) {
	url := fmt.Sprintf(fs.URL + "/itil/changes.json")

	c.ItilChange.Email = fs.Email

	reqItilChange, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqItilChange))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(fs.APIKey, "X")
	req.Header.Set("Content-Type", "application/json")
	res, err := httpclient.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var resItilChange ResponseItilChange

	err = json.Unmarshal(res, &resItilChange)
	if err != nil {
		return nil, err
	}

	return &resItilChange, nil
}

// UpdateChange func
func (fs *Freshservice) UpdateChange(change int64, c *RequestItilChange) (*ResponseItilChange, error) {
	url := fmt.Sprintf(fs.URL + "/itil/changes/" + strconv.FormatInt(change, 10) + ".json")

	reqItilChange, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(reqItilChange))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(fs.APIKey, "X")
	req.Header.Set("Content-Type", "application/json")
	res, err := httpclient.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var resItilChange ResponseItilChange

	err = json.Unmarshal(res, &resItilChange)
	if err != nil {
		return nil, err
	}

	return &resItilChange, nil
}

// AddChangeNote func
func (fs *Freshservice) AddChangeNote(change int64, n *RequestNote) (*ResponseNote, error) {
	url := fmt.Sprintf(fs.URL + "/itil/changes/" + strconv.FormatInt(change, 10) + "/notes.json")

	reqNote, err := json.Marshal(n)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqNote))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(fs.APIKey, "X")
	req.Header.Set("Content-Type", "application/json")
	res, err := httpclient.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var resNote ResponseNote

	err = json.Unmarshal(res, &resNote)
	if err != nil {
		return nil, err
	}

	return &resNote, nil
}

// CheckEndpoint func
func (fs *Freshservice) CheckEndpoint() error {
	url := fmt.Sprintf(fs.URL + "/itil/changes.json")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(fs.APIKey, "X")
	req.Header.Set("Content-Type", "application/json")
	_, err = httpclient.DoRequest(req)
	if err != nil {
		return err
	}

	return nil
}
