package geofenceapp

import (
	"encoding/json"

	"github.com/Natnael-Alemayehu/geofence/app/sdk/geofence"
)

type Zone struct {
	ID      string      `json:"id"`
	GeoJSON interface{} `json:"geojson"`
}

type Delivery struct {
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

// Decode implements the decoder interface.
func (a *Delivery) Decode(data []byte) error {
	return json.Unmarshal(data, a)
}

type Verification struct {
	Latitude   float64  `json:"latitude" `
	Longitude  float64  `json:"longitude"`
	Status     string   `json:"status"`
	LocationID []string `json:"location_id,omitempty"`
}

// Encode implements the encoder interface.
func (app Verification) Encode() ([]byte, string, error) {
	data, err := json.Marshal(app)
	return data, "application/json", err
}

func tosdkDelivery(delapp Delivery) geofence.Delivery {
	del := geofence.Delivery{
		Latitude:  delapp.Latitude,
		Longitude: delapp.Longitude,
	}
	return del
}

func toSDKVerification(geover geofence.Verification) Verification {
	st_geover := geover.Status.ToString(geover.Status)
	return Verification{
		Latitude:   geover.Latitude,
		Longitude:  geover.Longitude,
		Status:     st_geover,
		LocationID: geover.LocationID,
	}
}
