package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	Name        string `gorm:"unique" json:"name"`
	Description string `json:"description"`
}

type Tournament struct {
	gorm.Model
	Name        string `json:"name"`
	GameID      uint   `json:"game"`
	Description string `json:"description"`
	Location    string `json:"location"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Prize       int    `json:"prize"`
	Organizer   string `json:"organizer"`
	Contact     string `json:"contact"`
	MaxTeam     int    `json:"maxTeam"`
}
