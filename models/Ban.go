package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Ban struct {
	gorm.Model
	Id        int
	User      User
	StartDate time.Time
	EndDate   time.Time
	CreatedBy User
	Reason    string
}
