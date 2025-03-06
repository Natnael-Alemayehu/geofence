package geofenceapp

import "encoding/json"

type Delivery struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Zone struct {
	ID      string      `json:"id"`
	GeoJSON interface{} `json:"geojson"`
}

// Encode implements the encoder interface.
func (app Delivery) Encode() ([]byte, string, error) {
	data, err := json.Marshal(app)
	return data, "application/json", err
}
