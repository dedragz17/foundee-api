package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Id        int
	Name      string
	Questions []Question
}
