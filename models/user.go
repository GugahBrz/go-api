package models

import "time"

// User Model
type User struct {
	ID        uint64     `gorm:"primary_key;AUTO_INCREMENT" form:"id" json:"id"`
	Name      string     `gorm:"not null" form:"name" json:"name"`
	CreatedAt time.Time  `form:"created_at" json:"created_at"`
	UpdatedAt time.Time  `form:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `form:"deleted_at" json:"-"`
}
