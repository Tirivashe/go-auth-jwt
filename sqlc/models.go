// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"database/sql"
)

type User struct {
	ID        int32
	Name      string
	Email     string
	Password  sql.NullString
	Createdat sql.NullTime
	Updatedat sql.NullTime
}
