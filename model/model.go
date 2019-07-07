package model

import (
	"github.com/jinzhu/gorm"
)

// Config Exported
type Config struct {
	ElasticSearchUrl string `env:"ELASTICSEARCH_URL" envDefault:"Slomek"`
}

// Ticket data 
Cache struct {
	URL string `env:"REDIS_URL" envDefault:"Slomek"`
}

// Message Exported
type Message struct {
	gorm.Model
	Category string `json:"category"`
	Call []string `json:"call"`
	Text string `json:"text"`
	User User   `json:"user"`
}

// User Exported
type User struct {
	gorm.Model
	Name   string `json:"name"`
	Msisdn string `json:"msisdn"`
}
