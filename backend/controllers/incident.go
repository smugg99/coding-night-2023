package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/smugg99/coding-night-2023/models"
	"github.com/smugg99/coding-night-2023/services"
)

func (base *Controller) Incidents(c *fiber.Ctx) error {
    lat := c.QueryFloat("lat")
    lon := c.QueryFloat("lon")

    distance := c.QueryFloat("distance")

    incidents, err := services.GetNearbyIncidents(base.Db, &models.Coords{Lat: lat, Lon: lon}, float32(distance))
    if err != nil {
        return c.SendStatus(fiber.StatusInternalServerError)
    }

    return c.JSON(incidents)
}

func (base *Controller) CreateIncident(c *fiber.Ctx) error {
    var newIncident models.CreateIncidentDTO
    if err := c.BodyParser(&newIncident); err != nil {
        return c.SendStatus(fiber.StatusBadRequest)
    }

    user := c.Locals("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    userID := uint(claims["id"].(float64))
    
    categories, err := services.GetCategoriesByIds(base.Db, newIncident.CategoryIds)
    if err != nil {
        c.SendStatus(fiber.StatusInternalServerError)
    }

    incident := models.Incident{
        UserID: userID,
        Name: newIncident.Name,
        Coords: newIncident.Coords,
        DangerLvl: newIncident.DangerLvl,
        Date: newIncident.Date,
        Categories: categories,
    }

    id, err := services.CreateIncident(base.Db, &incident)
    if err != nil {
        c.SendStatus(fiber.StatusInternalServerError)
    }

    return c.JSON(id)
}

func (base *Controller) CreateCategory(c *fiber.Ctx) error {
    var newCategory models.CreateIncidentCategoryDTO
    if err := c.BodyParser(&newCategory); err != nil {
        return c.SendStatus(fiber.StatusInternalServerError)
    }

    category := models.IncidentCategory{
        Name: newCategory.Name,
        Icon: newCategory.Icon,
    }

    id, err := services.CreateCategory(base.Db, &category)
    if err != nil {
        return c.SendStatus(fiber.StatusInternalServerError)
    }

    return c.JSON(id)
}
