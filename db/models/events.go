package models

import (
	"time"
)

type event struct {
	ID           string    `json:"id" required:"true"`
	Name         string    `json:"name" required:"true"`
	Description  string    `json:"description" required:"true"`
	Destination  string    `json:"destination"`
	Datetime     time.Time `json:"datetime"`
	HostedBy     string    `json:"hosted_by" required:"true"`
	Participants string    `json:"participants"`
	Pictures     []string  `json:"pictures"`
}

type events []event
