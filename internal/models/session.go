package models

import "time"

type Session struct {
	UserLogin    string
	LastActivity time.Time
}
