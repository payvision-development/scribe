package release

import (
	"strconv"

	"github.com/payvision-development/scribe/freshservice"
)

// Changer interface
type Changer interface {
	Create(name string, environment string, msg string, date string) (int64, error)
	Update(msg string, status string) error
}

// FreshserviceChanger struct
type FreshserviceChanger struct {
	Client             *freshservice.Freshservice
	RequestItilChange  *freshservice.RequestItilChange
	ResponseItilChange *freshservice.ResponseItilChange
}

// Create func
func (f FreshserviceChanger) Create(name string, environment string, msg string, date string) (int64, error) {

	c := &freshservice.RequestItilChange{}

	c.ItilChange.Email = "hulk@outerspace.com"
	c.ItilChange.Subject = "[Release Management] Deployment of release " + name + " to environment " + environment
	c.ItilChange.DescriptionHTML = msg
	c.ItilChange.Status = freshservice.StatusPendingRelease
	c.ItilChange.Priority = freshservice.PriorityMedium
	c.ItilChange.ChangeType = freshservice.TypeStandard
	c.ItilChange.Risk = freshservice.RiskMedium
	c.ItilChange.Impact = freshservice.ImpactMedium
	c.ItilChange.PlannedStartDate = date
	c.ItilChange.PlannedEndDate = date

	f.RequestItilChange = c

	resItilChange, err := f.Client.CreateChange(c)
	if err != nil {
		return -1, err
	}

	f.ResponseItilChange = resItilChange

	return f.ResponseItilChange.Item.ItilChange.ID, nil
}

// Update func
func (f FreshserviceChanger) Update(msg string, status string) error {

	n := &freshservice.RequestNote{}
	n.Note.Body = msg

	_, err := f.Client.AddChangeNote(f.ResponseItilChange.Item.ItilChange.DisplayID, n)
	if err != nil {
		return err
	}

	i, err := strconv.Atoi(status)
	if err != nil {
		return err
	}

	f.RequestItilChange.ItilChange.Status = i

	_, err = f.Client.UpdateChange(f.ResponseItilChange.Item.ItilChange.DisplayID, f.RequestItilChange)
	if err != nil {
		return err
	}

	return nil
}
