package model

import (
	"github.com/jinzhu/gorm"
)

// Config Exported
type Config struct {
	ElasticSearchUrl string `env:"ELASTICSEARCH_URL" envDefault:"Slomek"`
}

// Message Exported
type Message struct {
	gorm.Model
	Text string `json:"text"`
	User User   `json:"user"`
}

// User Exported
type User struct {
	gorm.Model
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
