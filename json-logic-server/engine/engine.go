package engine

import (
	"encoding/json"
	"errors"

	jsonlogic "github.com/diegoholiveira/jsonlogic/v3"
)

type RuleType string

const (
	Validation RuleType = "VALIDATION"
	Workflow   RuleType = "WORKFLOW"
)

type Rule struct {
	ID           string          `json:"id"`
	Type         RuleType        `json:"type"`
	Logic        json.RawMessage `json:"logic"`
	ErrorMessage string          `json:"errorMessage,omitempty"`
}

type Engine struct{}

func New() *Engine {
	return &Engine{}
}

func (e *Engine) Evaluate(rule Rule, ctx map[string]any) (any, error) {
	if len(rule.Logic) == 0 {
		return nil, errors.New("empty rule logic")
	}
	return jsonlogic.ApplyInterface(rule.Logic, ctx)
}
