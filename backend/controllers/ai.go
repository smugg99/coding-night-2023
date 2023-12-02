package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/smugg99/coding-night-2023/config"
	"github.com/smugg99/coding-night-2023/models"
	"github.com/smugg99/coding-night-2023/services"
)

const system = "Your main goal is to provide information and advice on how to stay safe from specified incidents in the nearby area and how to get help in case something bad happens"

func (base *Controller) GetSuggestions(c *fiber.Ctx) error {
    lat := c.QueryFloat("lat")
    lon := c.QueryFloat("lon")

    distance := c.QueryFloat("distance")

    coords := models.Coords{Lat: lat, Lon: lon}
    incidents, err := services.GetNearbyIncidents(base.Db, &coords, float32(distance))
    if err != nil {
        return c.SendStatus(fiber.StatusInternalServerError)
    }

    myPosition, err := getMyPosition(&coords)
    if err != nil {
        return c.SendStatus(fiber.StatusServiceUnavailable)
    }

    user := fmt.Sprintf("Give me istructions how to stay safe from specified incidents: %v. I'm based in %s.", incidents, myPosition)

    payload := fmt.Sprintf("{\"model\": \"orca-mini:3b\",\"system\": \"%s\",\"prompt\": \"%s\",\"stream\": false}", system, user)

	body := strings.NewReader(payload)



    req, err := http.NewRequest("POST",fmt.Sprintf("http://%s:11434/api/generate", config.Conf.SrvConfig.AiHost), body)

	if err != nil {

		return c.SendStatus(fiber.StatusInternalServerError)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {

        return c.SendStatus(fiber.StatusServiceUnavailable)
	}

	defer resp.Body.Close()

    result, err := io.ReadAll(resp.Body)
    if err != nil {

        return c.SendStatus(fiber.StatusInternalServerError)
    }


    var mapResult map[string]interface{}

    err = json.Unmarshal(result, &mapResult)
    if err != nil {

        return c.SendStatus(fiber.StatusServiceUnavailable)
    }

    return c.JSON(mapResult["response"])

}


func getMyPosition(position *models.Coords) (string, error) {
    res, err := http.Get(fmt.Sprintf("https://geocode.maps.co/reverse?lat=%f&lon=%f", position.Lat, position.Lon))
    if err != nil {
        return "", err
    }
    defer res.Body.Close()
    jsonBody, err := io.ReadAll(res.Body)
    if err != nil {
        return "", err
    }

    var jsonMap map[string]interface{}

    json.Unmarshal(jsonBody, &jsonMap)

    return jsonMap["display_name"].(string), nil
}

