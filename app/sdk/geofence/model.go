package geofence

import (
	"encoding/json"
	"fmt"

	"github.com/Natnael-Alemayehu/geofence/app/sdk/errs"
)

type Zone struct {
	ID      string      `json:"id"`
	GeoJSON interface{} `json:"geojson"`
}

type Delivery struct {
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

// Validate checks the data in the model is considered clean.
func (app Delivery) Validate() error {
	if err := errs.Check(app); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	return nil
}

// Encode implements the encoder interface.
func (app Delivery) Encode() ([]byte, string, error) {
	data, err := json.Marshal(app)
	return data, "application/json", err
}

// Decode implements the decoder interface.
func (a *Delivery) Decode(data []byte) error {
	return json.Unmarshal(data, a)
}

type Status struct {
	VrfStatus string
}

func (Status) ToString(st Status) string {
	return st.VrfStatus
}

type Verification struct {
	Latitude   float64
	Longitude  float64
	Status     Status
	LocationID []string
}
