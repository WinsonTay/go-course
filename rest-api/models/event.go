package models

import (
	"time"

	"example.com/eventbooking-rest-api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `json:"dateTime"`
	UserID      int64     `json:"userID"`
}

func (e *Event) Save() error {

	//Later Add it to a database
	query := `INSERT INTO events (name, description, location, dateTime, user_id) 
			  VALUES (?, ?, ?, ?, ?)
			 `
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err

}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
func GetEvent(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	var event Event

	// GetEvent retrieves a single event from the database by its id
	// Returns the event and an error if it doesn't exist or there was a problem retrieving it
	err := db.DB.QueryRow(query, id)

	/*
		QueryRow executes a query that is expected to return at most one row.
		QueryRow always returns a non-nil value.
		If the query doesn't return any rows, the Scan will return sql.ErrNoRows.
		To check for sql.ErrNoRows, see the example code.
	*/

	// Scan copies the columns from the matched row into the values pointed at by dest
	// If there is no row, or if the query fails, it will return an error
	if err := err.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID); err != nil {
		return nil, err
	}

	return &event, nil
}
func (event *Event) Update() error {
	query := `
	UPDATE events 
	SET name =?, description = ?, location = ?, dateTime = ?, user_id = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID, event.ID)
	return err

}

func (event *Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	return err

}

func (event *Event) Register(userId int64) error {
	query := "INSERT INTO registrations (event_id, user_id) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID, userId)
	if err != nil {
		return err
	}
	return nil
}
