package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Note struct {
	gorm.Model
	Id        int
	User      User
	CreatedAt time.Time
	CreatedBy User
	Note      string
}
