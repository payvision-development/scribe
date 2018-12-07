package vss

// Event type
type Event struct {
	ReleaseTrackingCode uint32
	ID                  string
	EventType           string
	EnvironmentID       int
	EnvironmentName     string
	ReleaseID           int
	ReleaseName         string
	ReleaseURL          string
	ProjectID           string
	ProjectName         string
	ProjectURL          string
	ServerURL           string
	CollectionURL       string
	Status              string
	ApprovalType        string
	SubscriptionID      string
	NotificationID      int
	PublisherID         string
	Message             string
	DetailedMessageHTML string
	ResourceVersion     string
	Timestamp           string
}
