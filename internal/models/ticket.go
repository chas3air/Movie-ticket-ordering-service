package models

import (
	"fmt"
	"time"
)

type Ticket struct {
	Id          string
	MovieTitle  string
	MovieTime   time.Time
	ViewingArea int
}

func (t Ticket) String() string {
	return fmt.Sprintf("id: %s, movie title: %s, movie time: %v, seat: %d", t.Id, t.MovieTitle, t.MovieTime.Format("01-02-2006 03/04"), t.ViewingArea)
}
