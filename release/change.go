package release

import "github.com/sergiovhe/scribe/vss"

// Changer interface
type Changer interface {
	Create(name string, environment string, msg string, date string) error
	Update(msg string, status string) error
}

// FreshserviceChanger struct
type FreshserviceChanger struct {
	Client freshservice.Freshservice
	ID     int
	Change feshservice.ItilChange
}

// Create func
func (f Changer) Create(name string, environment string, msg string, date string) error {
	c := f.Change{
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

	f.Change, err := f.Client.CreateChange(&c)
	if err != nil {
		return nil, err
	}
	
	return nil
}

// Update func
func (f Changer) Update(msg string, status string) error {
	_, err := fs.AddChangeNote(id, msg)
	if err != nil {
		return err
	}

	_, err = fs.UpdateChangeStatus(id, freshservice.Status(status))
	if err != nil {
		return err
	}

	return nil
}
