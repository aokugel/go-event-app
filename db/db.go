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
		location TEXT NOT NULL,
		date DATETIME NOT NULL,
		invitees INTEGER NOT NULL
	)`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic(err)
	}
}

func InsertEventIntoDB(e *models.Event) error {
	insertEvent := `
	INSERT INTO events (Name, Location, Date, Invitees)
	VALUES(?, ?, ?, ?);`
	stmt, err := DB.Prepare(insertEvent)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	fmt.Println(stmt, e.Name, e.Location, e.Date, e.Invitees)
	result, err := stmt.Exec(e.Name, e.Location, e.Date, e.Invitees)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	e.ID = id
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

		err = rows.Scan(&event.ID, &event.Name, &event.Location, &event.Date, &event.Invitees)
		fmt.Println(event)

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
