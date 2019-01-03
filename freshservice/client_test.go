package freshservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestCreateChange(t *testing.T) {
	fsURL := "https://foo.freshservice.com"
	fs := NewClient(fsURL, "hulk@outerspace.com", "key")

	resJSON := loadTestData(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, fsURL+"/itil/changes.json",
		httpmock.NewBytesResponder(http.StatusOK, resJSON))

	c := &RequestItilChange{}
	c.ItilChange.Subject = "change for support"
	c.ItilChange.DescriptionHTML = "change description"
	c.ItilChange.Status = StatusOpen
	c.ItilChange.Priority = PriorityLow
	c.ItilChange.ChangeType = TypeMinor
	c.ItilChange.Risk = RiskLow
	c.ItilChange.Impact = ImpactLow
	c.ItilChange.PlannedStartDate = "2018-01-01T00:00:00.00Z"
	c.ItilChange.PlannedEndDate = "2018-01-01T00:00:00.00Z"

	res, err := fs.CreateChange(c)
	if err != nil {
		t.Error(err)
	}

	var resItilChange ResponseItilChange

	err = json.Unmarshal(resJSON, &resItilChange)
	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(res.Item, resItilChange.Item) != true {
		t.Error(err)
	}
}

func TestUpdateChangeStatus(t *testing.T) {
	fsURL := "https://foo.freshservice.com"
	fs := NewClient(fsURL, "hulk@outerspace.com", "key")

	changeID := 1
	changeStatus := StatusOpen
	resJSON := fmt.Sprintf(string(loadTestData(t)), changeStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPut,
		fmt.Sprintf("%s/itil/changes/%d.json", fsURL, changeID),
		httpmock.NewStringResponder(http.StatusOK, resJSON))

	c := &RequestItilChange{}

	c.ItilChange.Subject = "change for support"
	c.ItilChange.DescriptionHTML = "change description"
	c.ItilChange.Status = changeStatus
	c.ItilChange.Priority = PriorityLow
	c.ItilChange.ChangeType = TypeMinor
	c.ItilChange.Risk = RiskLow
	c.ItilChange.Impact = ImpactLow
	c.ItilChange.PlannedStartDate = "2018-01-01T00:00:00.00Z"
	c.ItilChange.PlannedEndDate = "2018-01-01T00:00:00.00Z"

	res, err := fs.UpdateChange(int64(changeID), c)
	if err != nil {
		t.Error(err)
	}

	if int(changeStatus) != res.Item.ItilChange.Status {
		t.Error(err)
	}
}

func TestAddChangeNote(t *testing.T) {
	fsURL := "https://foo.freshservice.com"
	fs := NewClient(fsURL, "hulk@outerspace.com", "key")

	changeID := 1
	changeNote := "Hi Hulk, Still Angry"
	resJSON := fmt.Sprintf(string(loadTestData(t)), changeNote, changeNote)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost,
		fmt.Sprintf("%s/itil/changes/%d/notes.json", fsURL, changeID),
		httpmock.NewStringResponder(http.StatusOK, resJSON))

	n := &RequestNote{}
	n.Note.BodyHTML = changeNote

	res, err := fs.AddChangeNote(int64(changeID), n)
	if err != nil {
		t.Error(err)
	}

	if changeNote != res.Item.Note.Body && "<div>"+changeNote+"<div>" != res.Item.Note.BodyHTML {
		t.Error(err)
	}
}

// loadTestData loads and returns the file contents of "_testdata/<TestName>.data"
func loadTestData(t *testing.T) []byte {
	t.Helper()

	name := fmt.Sprintf("_testdata/%s.data", t.Name())
	bytes, err := ioutil.ReadFile(name)
	if err != nil {
		t.Errorf("couldn't load test data (%q)", name)
		return nil
	}

	return bytes
}
