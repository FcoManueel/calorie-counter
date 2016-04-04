package models

import "time"

type Intake struct {
	ID         string     `json:"id"`
	UserID     string     `json:"userId"`
	Name       string     `json:"name"`
	Calories   int        `json:"calories"`
	CreatedAt  time.Time  `json:"createdAt"`
	ConsumedAt time.Time  `json:"consumedAt"`
	DisabledAt *time.Time `json:"disabledAt,omitempty"`
}

type Intakes []*Intake

// NewRecord implements factory method needed for DB queries
func (intakes *Intakes) NewRecord() interface{} {
	intake := &Intake{}
	*intakes = append(*intakes, intake)
	return intake
}

// IntakesData is a namespaced wrapper for intakes
type IntakesData struct {
	Data Intakes `json:"intakes"`
}
