// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Train Travel API
 *
 * API for finding and booking train trips across Europe.  ## Run in Postman  Experiment with this API in Postman, using our Postman Collection.  [![Run In Postman](https://run.pstmn.io/button.svg =128pxx32px)](https://app.getpostman.com/run-collection/9265903-7a75a0d0-b108-4436-ba54-c6139698dc08?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D9265903-7a75a0d0-b108-4436-ba54-c6139698dc08%26entityType%3Dcollection%26workspaceId%3Df507f69d-9564-419c-89a2-cb8e4c8c7b8f)
 *
 * API version: 1.2.0
 * Contact: support@example.com
 */

package oas

import (
	"time"
)

type GetTrips200ResponseAllOfDataInner struct {

	// Unique identifier for the trip
	Id string `json:"id,omitempty"`

	// The starting station of the trip
	Origin string `json:"origin,omitempty"`

	// The destination station of the trip
	Destination string `json:"destination,omitempty"`

	// The date and time when the trip departs
	DepartureTime time.Time `json:"departure_time,omitempty"`

	// The date and time when the trip arrives
	ArrivalTime time.Time `json:"arrival_time,omitempty"`

	// The name of the operator of the trip
	Operator string `json:"operator,omitempty"`

	// The cost of the trip
	Price float32 `json:"price,omitempty"`

	// Indicates whether bicycles are allowed on the trip
	BicyclesAllowed bool `json:"bicycles_allowed,omitempty"`

	// Indicates whether dogs are allowed on the trip
	DogsAllowed bool `json:"dogs_allowed,omitempty"`

	Self string `json:"self,omitempty"`
}

// AssertGetTrips200ResponseAllOfDataInnerRequired checks if the required fields are not zero-ed
func AssertGetTrips200ResponseAllOfDataInnerRequired(obj GetTrips200ResponseAllOfDataInner) error {
	return nil
}

// AssertGetTrips200ResponseAllOfDataInnerConstraints checks if the values respects the defined constraints
func AssertGetTrips200ResponseAllOfDataInnerConstraints(obj GetTrips200ResponseAllOfDataInner) error {
	return nil
}
