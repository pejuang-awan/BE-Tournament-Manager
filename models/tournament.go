package models

type Game struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"unique" json:"name"`
	Description string `json:"description"`
}

type Tournament struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	GameID      uint   `json:"game_id"`
	Description string `json:"description"`
	Location    string `json:"location"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Prize       string `json:"prize"`
	Organizer   string `json:"organizer"`
	Contact     string `json:"contact"`
	MaxTeam     int    `json:"max_team"`
}
