package model

import (
	"github.com/jinzhu/gorm"
)

// Message Exported
type Message struct {
	gorm.Model
	Category string   `json:"category"`
	Call     []string `json:"call"`
	Text     string   `json:"text"`
	User     User     `json:"user"`
}

// User Exported
type User struct {
	gorm.Model
	Name   string `json:"name"`
	Msisdn string `json:"msisdn"`
}
