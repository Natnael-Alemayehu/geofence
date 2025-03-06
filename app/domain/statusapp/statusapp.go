package statusapp

import (
	"context"
	"net/http"

	"github.com/Natnael-Alemayehu/geofence/foundation/web"
)

type app struct {
}

func newApp() *app {
	return &app{}
}

func (a *app) status(_ context.Context, _ *http.Request) web.Encoder {
	return Status{
		Status: "OK",
	}
}
