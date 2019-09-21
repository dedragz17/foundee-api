package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Vote struct {
	gorm.Model
	Id        int
	Answer    Answer
	Question  Question
	VotedAt   time.Time
	VotedBy   User
	VoteValue string
}
