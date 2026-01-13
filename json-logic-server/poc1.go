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

// Struct do pedido
type Order struct {
	CustomerType    string  `json:"customerType"`
	DiscountPercent float64 `json:"discountPercent"`
	GrossTotal      float64 `json:"grossTotal"`
	DeliveryType    string  `json:"deliveryType"`
	DeliveryFee     float64 `json:"deliveryFee"`
	ManagerApproved bool    `json:"managerApproved"`
}

// Payload esperado: { "order": {...} }
type Payload struct {
	Order Order `json:"order"`
}

// Função genérica para aplicar JsonLogic
func applyJsonLogic(rulePath string, payload Payload) (interface{}, error) {
	// Lê regra JsonLogic do arquivo
	ruleBytes, err := os.ReadFile(rulePath)
	if err != nil {
		return nil, err
	}

	// Converte payload para Reader
	payloadBytes, _ := json.Marshal(payload)
	payloadReader := bytes.NewReader(payloadBytes)

	var result bytes.Buffer
	if err := jsonlogic.Apply(bytes.NewReader(ruleBytes), payloadReader, &result); err != nil {
		return nil, err
	}

	// Decodifica resultado em interface{} (array ou objeto)
	var res interface{}
	if err := json.NewDecoder(&result).Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func applyDiscounts(c echo.Context) error {
	var payload Payload
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	res, err := applyJsonLogic("./rules/discounts.json", payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	//return c.JSON(http.StatusOK, res)
	return c.JSONPretty(http.StatusOK, res, "  ")
}

func applyRestrictions(c echo.Context) error {
	var payload Payload
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	res, err := applyJsonLogic("./rules/restrictions.json", payload)
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
