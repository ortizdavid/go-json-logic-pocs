package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/diegoholiveira/jsonlogic/v3"
	"github.com/labstack/echo/v4"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Location string `json:"location"`
}

func applyLogic(c echo.Context) error {
	// Lê dados enviados pelo cliente
	var payload struct {
		Users []User `json:"users"`
	}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Lê regra JsonLogic do arquivo
	dataBytes, err := os.ReadFile("poc2.rule.json")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Converte payload para Reader
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	payloadReader := bytes.NewReader(payloadBytes)

	// Aplica regra JsonLogic
	var result bytes.Buffer
	if err := jsonlogic.Apply(bytes.NewReader(dataBytes), payloadReader, &result); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Decodifica resultado direto no struct
	var filteredUsers []User
	if err := json.NewDecoder(&result).Decode(&filteredUsers); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, filteredUsers)
}

func main() {
	e := echo.New()
	e.POST("/apply-logic", applyLogic)

	fmt.Println("Server running on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
