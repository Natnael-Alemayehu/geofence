package geofenceapp

import (
	"encoding/json"
	"fmt"

	"github.com/Natnael-Alemayehu/geofence/app/sdk/geofence"
	"github.com/Natnael-Alemayehu/geofence/business/domain/geofencebus"
)

type Zone struct {
	ID      string                 `json:"id"`
	GeoJSON map[string]interface{} `json:"geojson"`
}

func (a *Zone) Decode(data []byte) error {
	return json.Unmarshal(data, a)
}

// Encode implements the encoder interface.
func (app Zone) Encode() ([]byte, string, error) {
	data, err := json.Marshal(app)
	return data, "application/json", err
}

// ====================================================================================

type Delivery struct {
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

// Decode implements the decoder interface.
func (a *Delivery) Decode(data []byte) error {
	fmt.Printf("\n\n %v \n", string(data))
	return json.Unmarshal(data, a)
}

// ====================================================================================

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

func toAppGeolocation(loc geofencebus.Geolocation) Zone {
	var geoJSONMap map[string]interface{}
	json.Unmarshal([]byte(loc.GeoJSON), &geoJSONMap)
	return Zone{
		ID:      loc.Location_ID,
		GeoJSON: geoJSONMap,
	}
}

func toBusGeolocation(loc Zone) geofencebus.Geolocation {
	geoJSONString, err := json.Marshal(loc.GeoJSON)
	if err != nil {
		return geofencebus.Geolocation{}
	}

	busGeo := geofencebus.Geolocation{
		Location_ID: loc.ID,
		GeoJSON:     string(geoJSONString),
	}
	return busGeo
}

// ==================================================================
type message struct {
	Message string `json:"message"`
}

// Encode implements the encoder interface.
func (app message) Encode() ([]byte, string, error) {
	data, err := json.Marshal(app)
	return data, "application/json", err
}
