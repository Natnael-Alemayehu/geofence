package geofencebus

import (
	"context"
	"fmt"

	"github.com/Natnael-Alemayehu/geofence/app/sdk/errs"
	"github.com/go-redis/redis/v8"
)

func checkCoordinate(ctx context.Context, delivery Delivery, geoloc Geolocation) (Verification, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:9851",
	})

	zone := Zone{
		ID:      geoloc.Location_ID,
		GeoJSON: geoloc.GeoJSON,
	}

	_, err := client.Do(ctx, "SET", "zones", zone.ID, "OBJECT", zone.GeoJSON).Result()
	if err != nil {
		return Verification{}, fmt.Errorf("client do: %v", err)
	}

	result, err := client.Do(ctx, "INTERSECTS", "zones", "POINT", delivery.Latitude, delivery.Longitude).Result()
	if err != nil {
		return Verification{}, errs.New(errs.Aborted, err)
	}

	ver := Verification{
		Longitude: delivery.Longitude,
		Latitude:  delivery.Latitude,
	}

	// Check if result is an array
	if resultArray, ok := result.([]interface{}); ok && len(resultArray) > 1 {
		// The intersecting objects are in resultArray[1]
		if objects, ok := resultArray[1].([]interface{}); ok {
			var idStrings []string
			// Iterate over each object in the objects array
			for _, obj := range objects {
				if objArray, ok := obj.([]interface{}); ok && len(objArray) > 0 {
					// The first element of objArray is the ID
					if id, ok := objArray[0].(string); ok {
						idStrings = append(idStrings, id)
					}
				}
			}
			// Determine if the point is inside based on whether we found any IDs
			if len(idStrings) > 0 {
				ver.Status = Status{"Inside"}
				ver.LocationID = idStrings
				return ver, nil
			} else {
				ver.Status = Status{"Outside"}
				return ver, err
			}
		}
	}
	return Verification{}, nil
}
