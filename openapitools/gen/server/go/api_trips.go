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
	"net/http"
	"strings"
	"time"
)

// TripsAPIController binds http requests to an api service and writes the service results to the http response
type TripsAPIController struct {
	service      TripsAPIServicer
	errorHandler ErrorHandler
}

// TripsAPIOption for how the controller is set up.
type TripsAPIOption func(*TripsAPIController)

// WithTripsAPIErrorHandler inject ErrorHandler into controller
func WithTripsAPIErrorHandler(h ErrorHandler) TripsAPIOption {
	return func(c *TripsAPIController) {
		c.errorHandler = h
	}
}

// NewTripsAPIController creates a default api controller
func NewTripsAPIController(s TripsAPIServicer, opts ...TripsAPIOption) *TripsAPIController {
	controller := &TripsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the TripsAPIController
func (c *TripsAPIController) Routes() Routes {
	return Routes{
		"GetTrips": Route{
			strings.ToUpper("Get"),
			"/rest/Train+Travel+API/1.0.0/trips",
			c.GetTrips,
		},
	}
}

// GetTrips - Get available train trips
func (c *TripsAPIController) GetTrips(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var originParam string
	if query.Has("origin") {
		param := query.Get("origin")

		originParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "origin"}, nil)
		return
	}
	var destinationParam string
	if query.Has("destination") {
		param := query.Get("destination")

		destinationParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "destination"}, nil)
		return
	}
	var dateParam time.Time
	if query.Has("date") {
		param, err := parseTime(query.Get("date"))
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Param: "date", Err: err}, nil)
			return
		}

		dateParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{"date"}, nil)
		return
	}
	var pageParam int32
	if query.Has("page") {
		param, err := parseNumericParameter[int32](
			query.Get("page"),
			WithParse[int32](parseInt32),
			WithMinimum[int32](1),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Param: "page", Err: err}, nil)
			return
		}

		pageParam = param
	} else {
		var param int32 = 1
		pageParam = param
	}
	var limitParam int32
	if query.Has("limit") {
		param, err := parseNumericParameter[int32](
			query.Get("limit"),
			WithParse[int32](parseInt32),
			WithMinimum[int32](1),
			WithMaximum[int32](100),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Param: "limit", Err: err}, nil)
			return
		}

		limitParam = param
	} else {
		var param int32 = 10
		limitParam = param
	}
	var bicyclesParam bool
	if query.Has("bicycles") {
		param, err := parseBoolParameter(
			query.Get("bicycles"),
			WithParse[bool](parseBool),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Param: "bicycles", Err: err}, nil)
			return
		}

		bicyclesParam = param
	} else {
		var param bool = false
		bicyclesParam = param
	}
	var dogsParam bool
	if query.Has("dogs") {
		param, err := parseBoolParameter(
			query.Get("dogs"),
			WithParse[bool](parseBool),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Param: "dogs", Err: err}, nil)
			return
		}

		dogsParam = param
	} else {
		var param bool = false
		dogsParam = param
	}
	result, err := c.service.GetTrips(r.Context(), originParam, destinationParam, dateParam, pageParam, limitParam, bicyclesParam, dogsParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
