package geofence

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Natnael-Alemayehu/geofence/app/sdk/errs"
	"github.com/go-redis/redis/v8"
)

func VerifyCoordinate(ctx context.Context, delivery Delivery) (Verification, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:9851",
	})

	zone := Zone{
		ID: "delivery_zone_1",
		GeoJSON: map[string]interface{}{
			"type": "Polygon",
			"coordinates": [][]interface{}{
				{
					[]interface{}{38.74256704424312, 9.033138471223111},
					[]interface{}{38.736467448797924, 9.032775582701646},
					[]interface{}{38.73709210616215, 9.030017618002177},
					[]interface{}{38.738194442688865, 9.025989500072342},
					[]interface{}{38.74322844615821, 9.02439275618923},
					[]interface{}{38.74894222381815, 9.028021709336926},
					[]interface{}{38.747435697232305, 9.032194960307649},
					[]interface{}{38.74256704424312, 9.033138471223111},
				},
			},
		},
	}

	zoneJSON, err := json.Marshal(zone.GeoJSON)
	if err != nil {
		return Verification{}, fmt.Errorf("Marshal Error: ", err)
	}

	_, err = client.Do(ctx, "SET", "zones", zone.ID, "OBJECT", zoneJSON).Result()
	if err != nil {
		return Verification{}, fmt.Errorf("Marshal Error: ", err)
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
