package geolocationdb

import "github.com/Natnael-Alemayehu/geofence/business/domain/geofencebus"

type geolocation struct {
	Location_ID string
	GeoJSON     string
}

func toDBGeolocation(bus geofencebus.Geolocation) geolocation {
	return geolocation{
		Location_ID: bus.Location_ID,
		GeoJSON:     bus.GeoJSON,
	}
}

func toBusGeolocation(gl geolocation) (geofencebus.Geolocation, error) {
	geoloc := geofencebus.Geolocation{
		Location_ID: gl.Location_ID,
		GeoJSON:     gl.GeoJSON,
	}

	return geoloc, nil
}
