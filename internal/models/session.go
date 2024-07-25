package models

import "time"

type Session struct {
	UserLogin    string    `json:"userLogin"`
	LastActivity time.Time `json:"lastActivity"`
}
