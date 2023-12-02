package services

import (
	"math"

	"github.com/smugg99/coding-night-2023/models"
	"gorm.io/gorm"
)


func GetNearbyIncidents(db *gorm.DB, position *models.Coords, distance float32) ([]models.Incident, error) {
    var incidents []models.Incident 

    topLeft, bottomRight := getCorners(*position, distance) 
    
    err := db.Where("lat BETWEEN ? AND ? AND lon BETWEEN ? AND ?",
        math.Min(topLeft.Lat, bottomRight.Lat), math.Max(topLeft.Lat, bottomRight.Lat),
        math.Min(topLeft.Lon, bottomRight.Lon), math.Max(topLeft.Lon, bottomRight.Lon),
    ).Find(&incidents).Error
    
    if err != nil {
        return []models.Incident{}, err
    }
    
    return incidents, nil
}

func CreateIncident(db *gorm.DB, incident *models.Incident) (uint, error) {
    if err := db.Create(incident).Error; err != nil {
        return 0, err
    }
    return incident.ID, nil
}

func GetCategoriesByIds(db *gorm.DB, ids []uint) ([]models.IncidentCategory, error) {
    var categories []models.IncidentCategory
    if err := db.Find(&categories, ids).Error; err != nil {
        return []models.IncidentCategory{}, err
    }

    return categories, nil
}

func CreateCategory(db *gorm.DB, category *models.IncidentCategory) (uint, error) {
    if err := db.Create(category).Error; err != nil {
        return 0, err
    }
    return category.ID, nil
}

func getCorners(position models.Coords, distance float32) (topLeft models.Coords, bottomRight models.Coords) {
    topLeft = position
    topLeft.MoveLat(distance)
    topLeft.MoveLon(-distance)

    bottomRight = position
    bottomRight.MoveLat(-distance)
    bottomRight.MoveLon(distance)

    return topLeft, bottomRight
}

