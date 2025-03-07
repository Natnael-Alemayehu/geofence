package geofenceapp

import (
	"context"
	"net/http"

	"github.com/Natnael-Alemayehu/geofence/app/sdk/errs"
	"github.com/Natnael-Alemayehu/geofence/app/sdk/geofence"
	"github.com/Natnael-Alemayehu/geofence/foundation/web"
)

type app struct {
}

func newApp() *app {
	return &app{}
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
