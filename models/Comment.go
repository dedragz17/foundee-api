package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	Id        int
	Answer    Answer
	Question  Question
	Value     string
	Votes     []Vote
	CreatedAt time.Time
	CreatedBy User
}
