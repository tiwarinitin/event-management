package models

import "example.com/event-mgmt/db"

type Users struct {
	ID       int64
	Name     string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
	Role     string
}

func (u Users) Save() error {
	query := `
	INSERT INTO users(name,email,password,role)
	VALUES (?,?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		u.Name,
		u.Email,
		u.Password,
		u.Role,
	)

	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	//u.ID = userId

	return err
}
