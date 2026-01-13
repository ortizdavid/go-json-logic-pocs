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

func applyJsonLogicRules(rulePath string, payload Payload) (map[string]interface{}, error) {
    // Lê regra JsonLogic
    ruleBytes, err := os.ReadFile(rulePath)
    if err != nil {
        return nil, err
    }

    // Decodifica arquivo de regras (array de {field, rule})
    var rules []struct {
        Field string                 `json:"field"`
        Rule  map[string]interface{} `json:"rule"`
    }
    if err := json.Unmarshal(ruleBytes, &rules); err != nil {
        return nil, err
    }

    // Resultado final
    result := make(map[string]interface{})

    for _, r := range rules {
        // Marshal rule individual
        ruleJSON, _ := json.Marshal(r.Rule)
        // Marshal payload
        payloadBytes, _ := json.Marshal(payload)
        var buf bytes.Buffer
        if err := jsonlogic.Apply(bytes.NewReader(ruleJSON), bytes.NewReader(payloadBytes), &buf); err != nil {
            return nil, err
        }
        // Decodifica valor final
        var value interface{}
        if err := json.NewDecoder(&buf).Decode(&value); err != nil {
            return nil, err
        }
        // Salva no resultado final
        result[r.Field] = value
    }

    return result, nil
}


func applyDiscounts(c echo.Context) error {
	var payload Payload
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	res, err := applyJsonLogicRules("./rules/discounts.json", payload)
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

	res, err := applyJsonLogicRules("./rules/restrictions.json", payload)
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
