package db

import (
	"database/sql"
	"fmt"

	"example.com/rest-api/models"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Error opening database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		location TEXT NOT NULL,
		date DATETIME NOT NULL,
		invitees INTEGER NOT NULL,
		userid INTEGER
	)`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic(err)
	}
}

func InsertEventIntoDB(e *models.Event) error {
	insertEvent := `
	INSERT INTO events (Name, Description, Location, Date, Invitees, Userid)
	VALUES(?, ?, ?, ?, ?, ?);`
	stmt, err := DB.Prepare(insertEvent)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.Date, e.Invitees, e.UserID)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func UpdateEvent(e *models.Event) error {
	updateEvent := `
	UPDATE events 
	SET Name = ?, 
		Description = ?, 
		Location = ?, 
		Date = ?, 
		Invitees = ?, 
		Userid = ?
	WHERE id = ?;`
	stmt, err := DB.Prepare(updateEvent)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.Date, e.Invitees, e.UserID, e.ID)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func DeleteEvent(id int64) error {
	deleteEvent := `
	DELETE FROM events 
	WHERE id = ?;`
	stmt, err := DB.Prepare(deleteEvent)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	id, err = result.LastInsertId()
	return err
}

func GetEvents() []models.Event {
	getEvents := `
	SELECT *
	FROM events;
	`

	var events []models.Event

	rows, err := DB.Query(getEvents)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var event models.Event

		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.Invitees, &event.UserID)

		if err != nil {
			panic(err)
		}

		events = append(events, event)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return events
}

func GetEventByID(id int64) (*models.Event, error) {
	var event models.Event

	getEvent := `
	SELECT *
	FROM events
	WHERE id = ?;
	`

	row := DB.QueryRow(getEvent, id)

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.Invitees, &event.UserID)

	if err != nil {
		fmt.Println("Error Retrieving Event")
		fmt.Println(err)
		return nil, err
	}

	return &event, nil
}
