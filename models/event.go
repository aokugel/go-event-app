package models

import "time"

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string
	Location    string    `binding:"required"`
	Date        time.Time `binding:"required"`
	Invitees    int
	UserID      int64
}

func NewEvent(id int64, name string, description string, location string, date time.Time, invitees int, uid int64) *Event {
	return &Event{
		ID:          id,
		Name:        name,
		Description: description,
		Location:    location,
		Date:        date,
		Invitees:    invitees,
		UserID:      uid,
	}

}
