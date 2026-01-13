package main

import (
	"fmt"
	"json-logic-server/helpers"
	"json-logic-server/user-age/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func applyLogic(c echo.Context) error {
	var payload struct {
		Users []models.User `json:"users"`
	}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	res, err := helpers.ApplyJsonLogicRules("rules/rules.json", payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func main() {
	e := echo.New()
	e.POST("/apply-logic", applyLogic)

	fmt.Println("Server running on :8081")
	e.Logger.Fatal(e.Start(":8081"))
}
