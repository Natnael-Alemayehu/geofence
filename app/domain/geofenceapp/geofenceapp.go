package geofenceapp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Natnael-Alemayehu/geofence/app/sdk/errs"
	"github.com/Natnael-Alemayehu/geofence/foundation/web"
	"github.com/go-redis/redis/v8"
)

type app struct {
}

func newApp() *app {
	return &app{}
}

func (a *app) VerifyLocation(ctx context.Context, _ *http.Request) web.Encoder {

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
		log.Fatal("Marshal Error: ", err)
	}

	_, err = client.Do(ctx, "SET", "zones", zone.ID, "OBJECT", zoneJSON).Result()
	if err != nil {
		log.Fatal("Error setting geofence:", err)
	}

	delivery := Delivery{
		Latitude:  9.02921925586169,
		Longitude: 38.741409590890214,
	}

	result, err := client.Do(ctx, "INTERSECTS", "zones", "POINT", delivery.Latitude, delivery.Longitude).Result()
	if err != nil {
		errs.Newf(errs.Aborted, "Error executing INTERSECTS:", err)
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
				fmt.Printf("Delivery at (%.3f, %.3f) is INSIDE the delivery zone.\n", delivery.Latitude, delivery.Longitude)
				fmt.Println("Matching Zone IDs:", idStrings)
			} else {
				fmt.Printf("Delivery at (%.3f, %.3f) is OUTSIDE the delivery zone.\n", delivery.Latitude, delivery.Longitude)
			}
		} else {
			fmt.Println("No objects found in result.")
		}
	} else {
		fmt.Println("Unexpected result format:", result)
	}

	return web.NewNoResponse()
}
