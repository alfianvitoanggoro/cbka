package model

import "time"

type User struct {
	UserID    int        `gorm:"primary_key;auto_increment" json:"user_id"`
	Username  string     `gorm:"size:100;not null" json:"username"`
	Email     string     `gorm:"size:100;not null;unique" json:"email"`
	Password  string     `gorm:"not null" json:""`
	Role      string     `gorm:"size:20;not null" json:"role"`
	IsActive  bool       `gorm:"not null;default:true" json:"is_active"`
	LastLogin *time.Time `json:"last_login,omitempty"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
