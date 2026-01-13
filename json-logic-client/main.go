package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
    input := map[string]interface{}{
        "user": map[string]interface{}{"type": "cliente_fidelidade", "age": 30},
        "cart": map[string]interface{}{"total": 250.0, "items": 6},
        "promo": map[string]interface{}{"code": "NATAL10", "discount": 10},
    }

    payload, _ := json.Marshal(input)
    resp, err := http.Post("http://localhost:8080/apply-rules", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Println(err)
	}
    defer resp.Body.Close()

    var result map[string]interface{}
    _ = json.NewDecoder(resp.Body).Decode(&result)

    total := input["cart"].(map[string]interface{})["total"].(float64) * result["multiplier"].(float64)
    fmt.Printf("Total final com desconto: %.2f\n", total)
}
