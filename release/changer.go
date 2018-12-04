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
	Client *freshservice.Freshservice
	Change *freshservice.ItilChange
}

// Create func
func (f FreshserviceChanger) Create(name string, environment string, msg string, date string) (int64, error) {

	c := freshservice.Change{
		Email:            "hulk@outerspace.com",
		Subject:          "[Release Management] Deployment of release " + name + " to environment " + environment,
		DescriptionHTML:  msg,
		Status:           freshservice.StatusPendingRelease,
		Priority:         freshservice.PriorityMedium,
		ChangeType:       freshservice.TypeStandard,
		Risk:             freshservice.RiskMedium,
		Impact:           freshservice.ImpactMedium,
		PlannedStartDate: date,
		PlannedEndDate:   date,
	}

	change, err := f.Client.CreateChange(&c)
	if err != nil {
		return -1, err
	}

	f.Change = change

	return f.Change.Item.ItilChange.ID, nil
}

// Update func
func (f FreshserviceChanger) Update(msg string, status string) error {
	_, err := f.Client.AddChangeNote(f.Change.Item.ItilChange.ID, msg)
	if err != nil {
		return err
	}

	s, err := strconv.ParseInt(status, 10, 8)
	if err != nil {
		return err
	}

	_, err = f.Client.UpdateChangeStatus(f.Change.Item.ItilChange.ID, freshservice.Status(s))
	if err != nil {
		return err
	}

	return nil
}
