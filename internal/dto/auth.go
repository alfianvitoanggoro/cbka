package dto

import "time"

type RequestRegistration struct {
	Username string `gorm:"size:100;not null" json:"username"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Role     string `gorm:"size:20;not null" json:"role"`
}

type ResponseRegistration struct {
	UserID    int        `gorm:"primary_key;auto_increment" json:"user_id"`
	Username  string     `gorm:"size:100;not null" json:"username"`
	Email     string     `gorm:"size:100;not null;unique" json:"email"`
	Role      string     `gorm:"size:20;not null" json:"role"`
	IsActive  bool       `gorm:"not null" json:"is_active"`
	LastLogin *time.Time `json:"last_login,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseLogin struct {
	Token string `json:"token"`
}
