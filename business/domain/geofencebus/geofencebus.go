package geofencebus

import (
	"context"
	"errors"
	"fmt"

	"github.com/Natnael-Alemayehu/geofence/foundation/logger"
)

var (
	ErrNotFound       = errors.New("product not found")
	ErrUniqueLocation = errors.New("location not unique")
)

type Storer interface {
	Create(ctx context.Context, prd Geolocation) error
	QueryByID(ctx context.Context, productID string) (Geolocation, error)
}

type Business struct {
	log    *logger.Logger
	storer Storer
}

// NewBusiness constructs a product business API for use.
func NewBusiness(log *logger.Logger, storer Storer) *Business {
	return &Business{
		log:    log,
		storer: storer,
	}
}

// Create adds a new geolocation to the system.
func (b *Business) Create(ctx context.Context, nu Geolocation) (Geolocation, error) {

	if err := b.storer.Create(ctx, nu); err != nil {
		return Geolocation{}, fmt.Errorf("create: %w", err)
	}

	return nu, nil
}

// QueryByID finds the product by the specified ID.
func (b *Business) QueryByID(ctx context.Context, locationID string) (Geolocation, error) {

	prd, err := b.storer.QueryByID(ctx, locationID)
	if err != nil {
		return Geolocation{}, fmt.Errorf("query: productID[%s]: %w", locationID, err)
	}

	return prd, nil
}
