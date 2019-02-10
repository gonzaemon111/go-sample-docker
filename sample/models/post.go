package models

import (
	"time"
)

// Post Model
type Post struct {
	ID     int            `gorm:"AUTO_INCREMENT;primary_key"`
	Header string         `gorm:"not null;size:255"`
	Body   string         `gorm:"not null;size:13000"`
	Author string         `gorm:"not null;size:30"`
	CreatedAt time.Time   `gorm:"not null"`
}