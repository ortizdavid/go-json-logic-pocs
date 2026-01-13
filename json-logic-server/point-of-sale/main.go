package main

import (
	"fmt"
	"json-logic-server/helpers"
	"json-logic-server/point-of-sale/models"
	"net/http"
	"github.com/labstack/echo/v4"
)

func applyDiscounts(c echo.Context) error {
	var payload models.Payload
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	res, err := helpers.ApplyJsonLogicRules("rules/discounts.json", payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	//return c.JSON(http.StatusOK, res)
	return c.JSONPretty(http.StatusOK, res, "  ")
}

func applyRestrictions(c echo.Context) error {
	var payload models.Payload
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	res, err := helpers.ApplyJsonLogicRules("rules/restrictions.json", payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	//return c.JSON(http.StatusOK, res)
	return c.JSONPretty(http.StatusOK, res, "  ")
}

func main() {
	e := echo.New()
	e.POST("/apply-discounts", applyDiscounts)
	e.POST("/apply-restrictions", applyRestrictions)

	fmt.Println("Server running on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
