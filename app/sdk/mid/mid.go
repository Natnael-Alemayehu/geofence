// Package mid provides app level middleware support.
package mid

import (
	"context"
	"errors"

	"github.com/Natnael-Alemayehu/geofence/business/domain/geofencebus"
	"github.com/Natnael-Alemayehu/geofence/foundation/web"
)

type ctxKey int

const (
	geolocationID ctxKey = iota + 1
)

// isError tests if the Encoder has an error inside of it.
func isError(e web.Encoder) error {
	err, isError := e.(error)
	if isError {
		return err
	}
	return nil
}

// =============================================================================
func setLocation(ctx context.Context, locationID string) context.Context {
	return context.WithValue(ctx, geolocationID, locationID)
}

func GetLocation(ctx context.Context) (geofencebus.Geolocation, error) {
	v, ok := ctx.Value(geolocationID).(geofencebus.Geolocation)
	if !ok {
		return geofencebus.Geolocation{}, errors.New("product not found in context")
	}

	return v, nil
}
