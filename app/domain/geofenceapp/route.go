package geofenceapp

import (
	"net/http"

	"github.com/Natnael-Alemayehu/geofence/business/domain/geofencebus"
	"github.com/Natnael-Alemayehu/geofence/foundation/logger"
	"github.com/Natnael-Alemayehu/geofence/foundation/web"
)

type Config struct {
	Log         *logger.Logger
	GeofenceBus *geofencebus.Business
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	api := newApp(cfg.GeofenceBus)

	app.HandlerFunc(http.MethodPost, version, "/verify_location", api.VerifyLocation)
	app.HandlerFunc(http.MethodGet, version, "/location/id/{location_id}", api.SearchLocationbyID)
	app.HandlerFunc(http.MethodGet, version, "/location/name/{location_name}", api.SearchLocationbyName)
	app.HandlerFunc(http.MethodPost, version, "/location", api.CreateGeoLocation)
	app.HandlerFunc(http.MethodGet, version, "/location/delete/{location_id}", api.DeleteGeoLocation)
}
