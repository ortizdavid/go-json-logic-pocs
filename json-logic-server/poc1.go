package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	jsonlogic "github.com/diegoholiveira/jsonlogic/v3"
)

type Rule struct {
	ID    string      `json:"id"`
	Logic interface{} `json:"logic"`
}

func main() {
	http.HandleFunc("/evaluate", handler)

	fmt.Println("JSON Logic Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var input map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	multiplier, err := applyDiscounts(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	allowed, err := applyRestrictions(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"multiplier": multiplier,
		"allowed":    allowed,
	})
}

func applyDiscounts(input map[string]interface{}) (float64, error) {
	rules, err := loadRules("rules/discounts.json")
	if err != nil {
		return 1, err
	}

	multiplier := 1.0

	for _, r := range rules {
		var result float64

		err := applyRule(r.Logic, input, &result)
		if err != nil {
			return 1, err
		}

		multiplier *= result
	}

	return multiplier, nil
}

func applyRestrictions(input map[string]interface{}) (bool, error) {
	rules, err := loadRules("rules/restrictions.json")
	if err != nil {
		return false, err
	}

	for _, r := range rules {
		var allowed bool

		err := applyRule(r.Logic, input, &allowed)
		if err != nil {
			return false, err
		}

		if !allowed {
			return false, nil
		}
	}

	return true, nil
}

func loadRules(path string) ([]Rule, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var rules []Rule
	if err := json.Unmarshal(data, &rules); err != nil {
		return nil, err
	}

	return rules, nil
}

func applyRule(logic interface{}, data interface{}, result interface{}) error {
	logicBytes, _ := json.Marshal(logic)
	dataBytes, _ := json.Marshal(data)

	return jsonlogic.Apply(
		bytes.NewReader(logicBytes),
		bytes.NewReader(dataBytes),
		result,
	)
}
