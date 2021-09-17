package models

import "database/sql"

// UserModel Structure of the users table
type User struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Password      string         `json:"password" gorm:"->:false;<-:update,create"`
	Status        int            `json:"status"`
	LastLoginAt   sql.NullString `json:"last_login_at"`
	RememberToken string         `json:"remember_token" gorm:"->:false;<-:create"`
	LastLoginIP   sql.NullString `json:"last_login_ip"`
	CreatedAt     sql.NullString `json:"created_at"`
	UpdatedAt     sql.NullString `json:"updated_at"`
	ForceChange   bool           `json:"forcechange" pg:"forceChange" gorm:"column:forceChange"`
	ObjectGUID    string         `json:"objectguid" pg:"objectguid" gorm:"column:objectguid"`
	AuthType      string         `json:"auth_type"`
	tableName     struct{}       `pg:"users"`
}
