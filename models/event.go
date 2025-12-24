package models

import (
	"fmt"
	"time"

	"example.com/myapp/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DateTime    time.Time `json:"date_time"`
	Location    string    `json:"location"`
	UserID      int       `json:"user_id"`
}

func (e *Event) Save() error {
	query := "INSERT INTO events(name,description,location,date,user_id) VALUES(?,?,?,?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(e.Name, e.Description, e.Location, time.Now(), e.UserID)
	defer stmt.Close()
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
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
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}		
	return &event, nil
}

func (e *Event) UpdateEvent() error {

	_,err:=db.DB.Exec("update events set name=?,description=?,location=? where id=?",e.Name,e.Description,e.Location,e.ID)
	if err!=nil{
		return err}
	fmt.Printf("event: %v\n", e)

	return nil
	

}

func DeleteEventByID(id int64) error{
    _,err:=db.DB.Exec("delete from events where id=?",id)
	if err!=nil{
		return err}
	return nil
}