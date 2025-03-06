// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"context"
	"net/http"

	"github.com/Natnael-Alemayehu/geofence/app/domain/geofenceapp"
	"github.com/Natnael-Alemayehu/geofence/app/sdk/mid"
	"github.com/Natnael-Alemayehu/geofence/foundation/logger"
	"github.com/Natnael-Alemayehu/geofence/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Build string
	Log   *logger.Logger
}

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(cfg Config) http.Handler {
	logger := func(ctx context.Context, msg string, args ...any) {
		cfg.Log.Info(ctx, msg, args...)
	}

	app := web.NewApp(
		logger,
		mid.Logger(cfg.Log),
		mid.Errors(cfg.Log),
		mid.Panics(),
	)

	geofenceapp.Routes(app, cfg.Log)

	return app
}
