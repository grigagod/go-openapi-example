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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// BookingsAPIController binds http requests to an api service and writes the service results to the http response
type BookingsAPIController struct {
	service      BookingsAPIServicer
	errorHandler ErrorHandler
}

// BookingsAPIOption for how the controller is set up.
type BookingsAPIOption func(*BookingsAPIController)

// WithBookingsAPIErrorHandler inject ErrorHandler into controller
func WithBookingsAPIErrorHandler(h ErrorHandler) BookingsAPIOption {
	return func(c *BookingsAPIController) {
		c.errorHandler = h
	}
}

// NewBookingsAPIController creates a default api controller
func NewBookingsAPIController(s BookingsAPIServicer, opts ...BookingsAPIOption) *BookingsAPIController {
	controller := &BookingsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the BookingsAPIController
func (c *BookingsAPIController) Routes() Routes {
	return Routes{
		"GetBookings": Route{
			strings.ToUpper("Get"),
			"/rest/Train+Travel+API/1.0.0/bookings",
			c.GetBookings,
		},
		"CreateBooking": Route{
			strings.ToUpper("Post"),
			"/rest/Train+Travel+API/1.0.0/bookings",
			c.CreateBooking,
		},
		"GetBooking": Route{
			strings.ToUpper("Get"),
			"/rest/Train+Travel+API/1.0.0/bookings/{bookingId}",
			c.GetBooking,
		},
		"DeleteBooking": Route{
			strings.ToUpper("Delete"),
			"/rest/Train+Travel+API/1.0.0/bookings/{bookingId}",
			c.DeleteBooking,
		},
	}
}

// GetBookings - List existing bookings
func (c *BookingsAPIController) GetBookings(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
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
	result, err := c.service.GetBookings(r.Context(), pageParam, limitParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// CreateBooking - Create a booking
func (c *BookingsAPIController) CreateBooking(w http.ResponseWriter, r *http.Request) {
	bookingParam := Booking{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bookingParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertBookingRequired(bookingParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertBookingConstraints(bookingParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateBooking(r.Context(), bookingParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetBooking - Get a booking
func (c *BookingsAPIController) GetBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookingIdParam := params["bookingId"]
	if bookingIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"bookingId"}, nil)
		return
	}
	result, err := c.service.GetBooking(r.Context(), bookingIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DeleteBooking - Delete a booking
func (c *BookingsAPIController) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookingIdParam := params["bookingId"]
	if bookingIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"bookingId"}, nil)
		return
	}
	result, err := c.service.DeleteBooking(r.Context(), bookingIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
