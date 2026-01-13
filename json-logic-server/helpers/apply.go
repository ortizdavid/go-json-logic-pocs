package helpers

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/diegoholiveira/jsonlogic/v3"
)

// Função genérica para aplicar JsonLogic
func ApplyJsonLogic(rulePath string, payload interface{}) (interface{}, error) {
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

func ApplyJsonLogicRules(rulePath string, payload interface{}) (map[string]interface{}, error) {
    // Lê regra JsonLogic
    ruleBytes, err := os.ReadFile(rulePath)
    if err != nil {
        return nil, err
    }

    // Decodifica arquivo de regras (array de {field, rule})
    var rules [] LogicRule
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

