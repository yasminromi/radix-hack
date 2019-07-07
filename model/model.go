package model

import (
	"github.com/jinzhu/gorm"
)

// Config Exported
type Config struct {
	ElasticSearchUrl string `env:"BONSAI_URL" envDefault:"Slomek"`
}

// Ticket data
type Cache struct {
	URL string `env:"REDISCLOUD_URL" envDefault:"Slomek"`
}

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
