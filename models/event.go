package models

import (
	"time"

	"example.com/event-mgmt/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

//var events = []Event{}

func (e Event) Save() error {
	query := `
	INSERT INTO events(name,description,dateTime,user_id)
	VALUES (?,?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		e.Name,
		e.Description,
		e.DateTime,
		e.UserID,
	)

	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	//e.ID = id

	return err
}

func GetEventByID(id int64) (*Event, error) {
	query := `
	SELECT * FROM events WHERE id = ?
	`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func GetAllEvents() ([]Event, error) {
	query := `
	SELECT * FROM events
	`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name=?, description=?,dateTime=?, user_id=?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		event.Name,
		event.Description,
		event.DateTime,
		event.UserID,
		event.ID,
	)
	return err

}

func (event Event) Delete() error {
	query := `
	DELETE 
	FROM events
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		event.ID,
	)

	return err
}

func (event Event) DeleteAll() error {
	query := `
	DELETE 
	FROM events
	`
	_, err := db.DB.Exec(query)

	return err
}
