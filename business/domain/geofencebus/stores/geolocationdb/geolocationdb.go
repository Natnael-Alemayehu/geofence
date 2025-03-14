package geolocationdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/Natnael-Alemayehu/geofence/business/domain/geofencebus"
	"github.com/Natnael-Alemayehu/geofence/business/sdk/sqldb"
	"github.com/Natnael-Alemayehu/geofence/foundation/logger"
	"github.com/jmoiron/sqlx"
)

// Store manages the set of APIs for user database access.
type Store struct {
	log *logger.Logger
	db  sqlx.ExtContext
}

// NewStore constructs the api for data access.
func NewStore(log *logger.Logger, db *sqlx.DB) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}

// Create inserts a new user into the database.
func (s *Store) Create(ctx context.Context, usr geofencebus.Geolocation) error {
	const q = `
	INSERT INTO geolocation
		(location_id, geojson)
	VALUES
		:location_id,
		ST_GeomFromGeoJSON(:geojson)`

	if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, toDBGeolocation(usr)); err != nil {
		if errors.Is(err, sqldb.ErrDBDuplicatedEntry) {
			return fmt.Errorf("namedexeccontext: %w", geofencebus.ErrUniqueLocation)
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// QueryByID finds the product identified by a given ID.
func (s *Store) QueryByID(ctx context.Context, locationID string) (geofencebus.Geolocation, error) {
	data := struct {
		ID string `db:"location_id"`
	}{
		ID: locationID,
	}

	const q = `
	SELECT
	    location_id, ST_AsGeoJSON(geojson) as geojson
	FROM
		geolocation
	WHERE
		location_id = :location_id`

	var dbPrd geolocation
	if err := sqldb.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbPrd); err != nil {
		if errors.Is(err, sqldb.ErrDBNotFound) {
			return geofencebus.Geolocation{}, fmt.Errorf("db: %w", geofencebus.ErrNotFound)
		}
		return geofencebus.Geolocation{}, fmt.Errorf("db: %w", err)
	}

	return toBusGeolocation(dbPrd)
}
