package models

import (
	"database/sql"
	"time"
)

// UserModel Structure of the users table
type User struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Password      string         `json:"password" gorm:"->:false;<-:update,create"`
	Status        int            `json:"status"`
	LastLoginAt   time.Time      `json:"last_login_at"`
	RememberToken string         `json:"remember_token" gorm:"->:false;<-:create"`
	LastLoginIP   sql.NullString `json:"last_login_ip"`
	CreatedAt     sql.NullString `json:"created_at"`
	UpdatedAt     sql.NullString `json:"updated_at"`
	ForceChange   bool           `json:"forcechange" pg:"forceChange" gorm:"column:forceChange"`
	ObjectGUID    string         `json:"objectguid" pg:"objectguid" gorm:"column:objectguid"`
	AuthType      string         `json:"auth_type"`
	SessionTime   int            `json:"session_time" pg:"session_time"`
	tableName     struct{}       `pg:"users"`
}
