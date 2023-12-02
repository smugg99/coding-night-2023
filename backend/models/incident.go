package models

import (
	"time"
    "math"
	"gorm.io/gorm"
)

const earthRadius float64 = 6371 // In km

type Incident struct {
    gorm.Model
    User User `json:"user"`
    UserID uint `json:"user_id"`
    Name string `json:"name" validate:"required"`
    Coords `json:"coords"` 
    Date time.Time `json:"date" validate:"required"`
    DangerLvl uint `json:"danger_lvl" validate:"max=4"`
    Categories []IncidentCategory `gorm:"many2many:incidents_categories" json:"categories"`
}

type IncidentCategory struct {
    gorm.Model
    Name string `json:"name" validate:"required"`
    Icon string `json:"icon" validation:"required"`
    Incident []Incident `gorm:"many2many:incidents_categories" json:"incident"`
}

type Coords struct {
    Lat float64 `json:"lat" validate:"latitude"`
    Lon float64 `json:"lon" validate:"longitude"`
}

type CreateIncidentDTO struct {
    Name string `json:"name" validate:"required"`
    Coords `json:"coords"` 
    Date time.Time `json:"date" validate:"required"`
    DangerLvl uint `json:"danger_lvl" validate:"max=4"`
    CategoryIds []uint `json:"category_ids"`
}

type CreateIncidentCategoryDTO struct {
    Name string `json:"name" validate:"required"`
    Icon string `json:"icon" validation:"required"`
}


func (c *Coords) MoveLat(km float32) {
	c.Lat = c.Lat + (float64(km)/earthRadius)*(180.0/math.Pi)
}

func (c *Coords) MoveLon(km float32) {
	c.Lon = c.Lon + (float64(km)/earthRadius)*(180/math.Pi)/math.Cos(c.Lat*math.Pi/180)
}
