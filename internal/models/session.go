package models

import (
	"fmt"
	"time"
)

type Session struct {
	UserLogin    string    `json:"userLogin"`
	LastActivity time.Time `json:"lastActivity"`
}

func (s Session) String() string {
	return fmt.Sprintf("userlogin: %s, last_time_activity: %v", s.UserLogin, s.LastActivity.Format("01-02-2006/03-04-05"))
}