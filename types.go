package main

type Plant struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Seed struct {
	ID              uint   `gorm:"primaryKey"`
	PlantID         uint   `json:"plant_id"`
	Characteristics string `json:"characteristics"`
}

type Species struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
