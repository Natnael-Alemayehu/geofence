package statusapp

import (
	"net/http"

	"github.com/Natnael-Alemayehu/geofence/foundation/logger"
	"github.com/Natnael-Alemayehu/geofence/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(app *web.App, log *logger.Logger) {
	const version = "v1"

	api := newApp()

	app.HandlerFunc(http.MethodGet, version, "/status", api.status)
}
