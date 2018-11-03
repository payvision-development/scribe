package freshservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/payvision-development/scribe/httpclient"
)

// Freshservice struct
type Freshservice struct {
	URL string
}

// NewClient func
func NewClient(url string) *Freshservice {
	fs := new(Freshservice)
	fs.URL = url
	return fs
}

// CreateChange func
func (fs *Freshservice) CreateChange(c *Change) ([]byte, error) {
	url := fmt.Sprintf(fs.URL + "/itil/changes.json")

	change, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(change))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth("X", "X")

	return httpclient.DoRequest(req)
}

// AddChangeNote func
func (fs *Freshservice) AddChangeNote(change int, note string) ([]byte, error) {
	return nil, nil
}

// UpdateChangeStatus func
func (fs *Freshservice) UpdateChangeStatus(change int, status int) ([]byte, error) {
	return nil, nil
}
