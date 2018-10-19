package models

import (
	"database/sql"
)

type RoleType int8

const (
	Default RoleType = iota
	Admin
)

type User struct {
	ID       int64    `json:"id"`
	Login    string   `json:"login"`
	Password string   `json:"password,omitempty"`
	Email    string   `json:"email,omitempty"`
	Role     RoleType `json:"role"`
	Active   bool     `json:"active"`
}

func GetAllUsers(db *sql.DB) (users []User, err error) {
	rows, err := db.Query("SELECT user_id, user_login, user_role, user_active FROM users")
	for rows.Next() {
		var user User
		if ok := rows.Scan(&user.ID, &user.Login, &user.Role, &user.Active); ok == nil {
			users = append(users, user)
		}
	}
	return users, err
}
