package models

import (
	"errors"

	"example.com/myapp/db"
	"example.com/myapp/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error{
	query := "INSERT INTO users(email,password) VALUES(?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err}
	res, err := stmt.Exec(u.Email, hashedPassword)
	defer stmt.Close()
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	u.ID = id
	return err
}

func (u *User) ValidateCredentials() (error,string) {
	query := "SELECT id,password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var storedHashedPassword ,jwtToken string
	var userId int64;
	err := row.Scan(&userId, &storedHashedPassword)
	if err != nil {
		return err,""
	}
	passwordIsValid:= utils.ComparePassword(storedHashedPassword, u.Password)
	if !passwordIsValid {
		return  errors.New("invalid credentials"),""
	}
	jwtToken, err = utils.GenerateJWT(u.Email, u.ID)
	if err != nil {
		return err, ""
	}
	return nil, jwtToken

}