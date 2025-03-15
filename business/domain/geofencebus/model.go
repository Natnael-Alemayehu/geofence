package geofencebus

import (
	"encoding/json"
	"fmt"

	"github.com/Natnael-Alemayehu/geofence/app/sdk/errs"
)

type Geolocation struct {
	Location_ID string
	GeoJSON     string
}

// ==========================================================
// Data from geofence SDK

type Zone struct {
	ID      string `json:"id"`
	GeoJSON string `json:"geojson"`
}

// type Geometry struct {
// 	GeoJSON map[string]interface{}
// }

// // Value implements the driver.Valuer interface.
// func (g Geometry) Value() (driver.Value, error) {
// 	if g.GeoJSON == nil {
// 		return nil, nil
// 	}
// 	return json.Marshal(g.GeoJSON)
// }

// // Scan implements the sql.Scanner interface.
// func (g *Geometry) Scan(value interface{}) error {
// 	if value == nil {
// 		g.GeoJSON = nil
// 		return nil
// 	}

// 	// Convert the value to a string
// 	var geoJSONString string
// 	switch v := value.(type) {
// 	case []byte:
// 		geoJSONString = string(v)
// 	case string:
// 		geoJSONString = v
// 	default:
// 		return fmt.Errorf("unsupported type for GeoJSON: %T", value)
// 	}

// 	// Unmarshal the string into the GeoJSON map
// 	return json.Unmarshal([]byte(geoJSONString), &g.GeoJSON)
// }

type Delivery struct {
	LocationID string  `json:"location_id" validate:"required"`
	Latitude   float64 `json:"latitude" validate:"required"`
	Longitude  float64 `json:"longitude" validate:"required"`
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
