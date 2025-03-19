package geolocationdb

import (
	"github.com/Natnael-Alemayehu/geofence/business/domain/geofencebus"
	"github.com/google/uuid"
)

type geolocation struct {
	Location_ID   uuid.UUID
	Location_Name string
	GeoJSON       string
}

func toDBGeolocation(bus geofencebus.Geolocation) geolocation {
	return geolocation{
		Location_ID:   bus.Location_ID,
		Location_Name: bus.Location_Name,
		GeoJSON:       bus.GeoJSON,
	}
}

func toBusGeolocation(gl geolocation) (geofencebus.Geolocation, error) {
	geoloc := geofencebus.Geolocation{
		Location_Name: gl.Location_Name,
		Location_ID:   gl.Location_ID,
		GeoJSON:       gl.GeoJSON,
	}

	return geoloc, nil
}
