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

type GetBookings200Response struct {
	Data []Booking `json:"data,omitempty"`

	Links GetStations200ResponseAllOfLinks `json:"links,omitempty"`
}

// AssertGetBookings200ResponseRequired checks if the required fields are not zero-ed
func AssertGetBookings200ResponseRequired(obj GetBookings200Response) error {
	for _, el := range obj.Data {
		if err := AssertBookingRequired(el); err != nil {
			return err
		}
	}
	if err := AssertGetStations200ResponseAllOfLinksRequired(obj.Links); err != nil {
		return err
	}
	return nil
}

// AssertGetBookings200ResponseConstraints checks if the values respects the defined constraints
func AssertGetBookings200ResponseConstraints(obj GetBookings200Response) error {
	for _, el := range obj.Data {
		if err := AssertBookingConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertGetStations200ResponseAllOfLinksConstraints(obj.Links); err != nil {
		return err
	}
	return nil
}
