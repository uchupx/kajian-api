package model

import "database/sql"

type User struct {
	BaseModel
	Id        sql.NullInt64  `db:"id"`
	Username  sql.NullString `db:"username"`
	Password  sql.NullString `db:"password"`
	Email     sql.NullString `db:"email"`
	CreatedAt sql.NullTime   `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
	Deletedat sql.NullTime   `db:"deletedat"`
}

func (User) TableName() string {
	return "users"
}
