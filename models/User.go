package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Id       int    `db:"id, primarykey"`
	Nickname string `db:"nickname"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Bans     []Ban  `db:"bans"`
}
