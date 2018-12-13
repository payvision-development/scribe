package freshservice

// Status values
const (
	StatusOpen int = iota + 1
	StatusPlanning
	StatusAwaitingApproval
	StatusPendingRelease
	StatusPendingReview
	StatusClosed
)

// Priority values
const (
	PriorityLow int = iota + 1
	PriorityMedium
	PriorityHigh
	PriorityUrgent
)

// Impact values
const (
	ImpactLow int = iota + 1
	ImpactMedium
	ImpactHigh
)

// Type values
const (
	TypeMinor int = iota + 1
	TypeStandard
	TypeMajor
	TypeEmergency
)

// Risk values
const (
	RiskLow int = iota + 1
	RiskMedium
	RiskHigh
	RiskVeryHigh
)
