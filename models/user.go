package models

import (
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