package model

import "gorm.io/gorm"

type RaceType struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null"`
}
