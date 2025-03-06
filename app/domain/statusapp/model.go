package statusapp

import "encoding/json"

type Status struct {
	Status string `json:"status"`
}

// Encode implements the encoder interface.
func (app Status) Encode() ([]byte, string, error) {
	data, err := json.Marshal(app)
	return data, "application/json", err
}
