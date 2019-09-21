package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Answer struct {
	gorm.Model
	Id       int
	Question Question
	//votes
	Author      User
	CreatedAt   time.Time
	AnswerValue string
}
