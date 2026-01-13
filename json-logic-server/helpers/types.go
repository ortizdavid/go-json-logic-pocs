package helpers

  type LogicRule struct {
        Field string                 `json:"field"`
        Rule  map[string]interface{} `json:"rule"`
    }