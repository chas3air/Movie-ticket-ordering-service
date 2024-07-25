package models

import (
	"fmt"
	"time"
)

type Ticket struct {
	Id          string    `json:"id"`
	MovieTitle  string    `json:"movieTitle"`
	MovieTime   time.Time `json:"movieTime"`
	ViewingArea int       `json:"viewingArea"`
	Email       string    `json:"email"`
}

func (t Ticket) String() string {
	return fmt.Sprintf("id: %s, movie title: %s, movie time: %v, seat: %d, email: %s", t.Id, t.MovieTitle, t.MovieTime.Format("01-02-2006 03/04"), t.ViewingArea, t.Email)
}
