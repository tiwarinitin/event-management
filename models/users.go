package models

import (
	"errors"

	"example.com/event-mgmt/db"
	"example.com/event-mgmt/utils"
)

type Users struct {
	ID       int64
	Name     string
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

	hashedPassword, err := utils.HashedPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(
		u.Name,
		u.Email,
		hashedPassword,
		u.Role,
	)

	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	//u.ID = userId

	return err
}

func (u *Users) ValidateUser() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
