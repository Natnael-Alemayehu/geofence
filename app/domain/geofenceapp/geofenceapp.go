package geofenceapp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Natnael-Alemayehu/geofence/app/sdk/errs"
	"github.com/Natnael-Alemayehu/geofence/app/sdk/geofence"
	"github.com/Natnael-Alemayehu/geofence/business/domain/geofencebus"
	"github.com/Natnael-Alemayehu/geofence/foundation/web"
)

type app struct {
	geolocBus *geofencebus.Business
}

func newApp(geolocBus *geofencebus.Business) *app {
	return &app{
		geolocBus: geolocBus,
	}
}

func (a *app) VerifyLocation(ctx context.Context, r *http.Request) web.Encoder {

	var delivery Delivery
	if err := web.Decode(r, &delivery); err != nil {
		return errs.New(errs.InvalidArgument, err)
	}

	delsdk := tosdkDelivery(delivery)

	ver, err := geofence.VerifyCoordinate(ctx, delsdk)
	if err != nil {
		return errs.Newf(errs.Aborted, "Verification failed: %v", err)
	}

	verif := toSDKVerification(ver)

	return verif
}

func (a *app) SearchLocation(ctx context.Context, r *http.Request) web.Encoder {

	id := web.Param(r, "location_id")

	if id == "" {
		return errs.Newf(errs.Aborted, "location id formatting: %v", fmt.Errorf("not found"))
	}

	geoloc, err := a.geolocBus.QueryByID(ctx, id)
	if err != nil {
		return errs.Newf(errs.Aborted, "Query by ID: %v", err)
	}

	return toAppGeolocation(geoloc)
}

func (a *app) CreateGeolocation(ctx context.Context, r *http.Request) web.Encoder {
	var newloc Zone
	if err := web.Decode(r, &newloc); err != nil {
		return errs.Newf(errs.InvalidArgument, "Decode: %v", err)
	}
	nl := toBusGeolocation(newloc)

	loc, err := a.geolocBus.Create(ctx, nl)
	if err != nil {
		return errs.Newf(errs.Aborted, "create: %v", err)
	}

	return toAppGeolocation(loc)
}
