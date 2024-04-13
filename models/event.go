package models

import "time"

var Events = []Event{
	{ID: 1, Name: "Wedding", Location: "Downtown Abbey", Date: time.Now(), Invitees: 400},
	{ID: 2, Name: "Funeral", Location: "Some church", Date: time.Now(), Invitees: 25},
	{ID: 3, Name: "Deposition", Location: "Courthouse", Date: time.Now(), Invitees: 100},
}

type Event struct {
	ID       int
	Name     string    `binding:"required"`
	Location string    `binding:"required"`
	Date     time.Time `binding:"required"`
	Invitees int       `binding:"required"`
}

func NewEvent(id int, name string, location string, date time.Time, invitees int) *Event {
	return &Event{
		ID:       id,
		Name:     name,
		Location: location,
		Date:     date,
		Invitees: invitees,
	}

}
