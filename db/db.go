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

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		userid INTEGER PRIMARY KEY AUTOINCREMENT,
		firstName TEXT,
		lastName TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT
	);`

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		eventid INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		location TEXT NOT NULL,
		date DATETIME NOT NULL,
		invitees INTEGER NOT NULL,
		userid INTEGER,
		FOREIGN KEY(userid) REFERENCES users(userid)
	);`

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		registrationsid INTEGER PRIMARY KEY AUTOINCREMENT,
		userid INTEGER,
		eventid INTEGER,
		FOREIGN KEY(eventid) REFERENCES EVENTS(eventid),
		FOREIGN KEY(userid) REFERENCES users(userid)
	);`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(err)

	}

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic(err)

	}

	_, err = DB.Exec(createRegistrationsTable)
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
	WHERE eventid = ?;`
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
	WHERE eventid = ?;`
	stmt, err := DB.Prepare(deleteEvent)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	//id, err = result.LastInsertId()
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
	WHERE eventid = ?;
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

func GetUsers() []models.User {
	getUsers := `
	SELECT *
	FROM users;
	`

	var users []models.User

	rows, err := DB.Query(getUsers)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return users
}

func InsertUserIntoDB(u *models.User) error {
	insertEvent := `
	INSERT INTO users (firstName, lastName, email, password)
	VALUES(?, ?, ?, ?);`
	stmt, err := DB.Prepare(insertEvent)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.Password)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE email = ?`
	row := DB.QueryRow(query, email)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)

	if err != nil {
		fmt.Println("Error Retrieving User")
		return models.User{}, err
	}
	return user, nil
}

func RegisterUserForEvent(userid int64, eventid int64) {
	insertRegistration := `
	INSERT INTO registrations (userid, eventid)
	VALUES(?, ?);`
	stmt, err := DB.Prepare(insertRegistration)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userid, eventid)
	if err != nil {
		panic(err)
	}
}

func UnegisterUserFromEvent(userid int64, eventid int64) {
	deleteRegistration := `
	DELETE FROM registrations
	WHERE userid = ?
	AND eventid = ?;`
	stmt, err := DB.Prepare(deleteRegistration)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userid, eventid)
	if err != nil {
		panic(err)
	}

}

// its going to wind up looking something like this.
func GetEventsByRegisteredUser(userid int64) (events []models.Event) {
	registeredEvents := `
	SELECT name, description, location, date, invitees, events.eventid, events.userid
	FROM events 
	INNER JOIN registrations
	ON registrations.registrationsid = events.eventid
	WHERE registrations.userid = ?;
	`
	stmt, err := DB.Prepare(registeredEvents)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := DB.Query(registeredEvents, userid)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var event models.Event

		err = rows.Scan(&event.Name, &event.Description, &event.Location, &event.Date, &event.Invitees, &event.ID, &event.UserID)

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
