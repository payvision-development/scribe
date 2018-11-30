package freshservice

// Status enum
type Status uint8

// Status values
const (
	StatusOpen Status = iota + 1
	StatusPlanning
	StatusAwaitingApproval
	StatusPendingRelease
	StatusPendingReview
	StatusClosed
)

// Priority enum
type Priority uint8

// Priority values
const (
	PriorityLow Priority = iota + 1
	PriorityMedium
	PriorityHigh
	PriorityUrgent
)

// Impact enum
type Impact uint8

// Impact values
const (
	ImpactLow Impact = iota + 1
	ImpactMedium
	ImpactHigh
)

// Type enum
type Type uint8

// Type values
const (
	TypeMinor Type = iota + 1
	TypeStandard
	TypeMajor
	TypeEmergency
)

// Risk enum
type Risk uint8

// Risk values
const (
	RiskLow Risk = iota + 1
	RiskMedium
	RiskHigh
	RiskVeryHigh
)
