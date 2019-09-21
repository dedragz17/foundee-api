package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Question struct {
	gorm.Model
	Id          int
	Author      User
	Title       string
	Tags        string
	Description string
	Answers     []Answer
	Comments    []Comment
	Votes       []Vote
	Category    Category
	CreatedAt   time.Time
}
