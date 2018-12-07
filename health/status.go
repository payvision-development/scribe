package health

import "time"

// Status type
type Status struct {
	Service     string
	Description string
	Status      string
	Version     string
	Info        Info
}

// Info type
type Info struct {
	Started time.Time
	Events  uint32
}
