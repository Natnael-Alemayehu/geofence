// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"context"
	"net/http"

	"github.com/Natnael-Alemayehu/geofence/app/domain/geofenceapp"
	"github.com/Natnael-Alemayehu/geofence/app/domain/statusapp"
	"github.com/Natnael-Alemayehu/geofence/app/sdk/mid"
	"github.com/Natnael-Alemayehu/geofence/business/domain/geofencebus"
	"github.com/Natnael-Alemayehu/geofence/business/domain/geofencebus/stores/geolocationdb"
	"github.com/Natnael-Alemayehu/geofence/foundation/logger"
	"github.com/Natnael-Alemayehu/geofence/foundation/web"
	"github.com/jmoiron/sqlx"
)

// Options represent optional parameters.
type Options struct {
	corsOrigin []string
}

// WithCORS provides configuration options for CORS.
func WithCORS(origins []string) func(opts *Options) {
	return func(opts *Options) {
		opts.corsOrigin = origins
	}
}

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Build string
	Log   *logger.Logger
	DB    *sqlx.DB
}

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(cfg Config, options ...func(opts *Options)) http.Handler {
	logger := func(ctx context.Context, msg string, args ...any) {
		cfg.Log.Info(ctx, msg, args...)
	}

	app := web.NewApp(
		logger,
		mid.Logger(cfg.Log),
		mid.Errors(cfg.Log),
		mid.Panics(),
	)

	var opts Options
	for _, option := range options {
		option(&opts)
	}

	if len(opts.corsOrigin) > 0 {
		app.EnableCORS(opts.corsOrigin)
	}

	geofenceBus := geofencebus.NewBusiness(cfg.Log, geolocationdb.NewStore(cfg.Log, cfg.DB))

	geofenceapp.Routes(app, geofenceapp.Config{
		Log:         cfg.Log,
		GeofenceBus: geofenceBus,
	})
	statusapp.Routes(app, cfg.Log)

	return app
}
