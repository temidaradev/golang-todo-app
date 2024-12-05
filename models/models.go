package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Todo string `json:"todo" gorm:"text;not null;default;null`
}
