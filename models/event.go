package models

import "time"

type Event struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        time.Time `json:"date"`
	Location    string `json:"location"`
	UserID      int    `json:"user_id"`
}
var events = []Event{}

func (e Event) Save(){
	events=append(events,e)

}

func GetAllEvents() []Event {
	return events
}
