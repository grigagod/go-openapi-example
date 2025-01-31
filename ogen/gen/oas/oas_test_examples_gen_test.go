// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"fmt"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"

	std "encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBooking_EncodeDecode(t *testing.T) {
	var typ Booking
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Booking
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestBookingPayment_EncodeDecode(t *testing.T) {
	var typ BookingPayment
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 BookingPayment
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}

func TestBookingPayment_Examples(t *testing.T) {

	for i, tc := range []struct {
		Input string
	}{
		{Input: "{\"amount\":100.5,\"currency\":\"gbp\",\"source\":{\"account_type\":\"individual\",\"bank_name\":\"Starling Bank\",\"country\":\"gb\",\"name\":\"J. Doe\",\"number\":\"00012345\",\"object\":\"bank_account\",\"sort_code\":\"000123\"}}"},
		{Input: "{\"amount\":49.99,\"currency\":\"gbp\",\"source\":{\"address_city\":\"London\",\"address_country\":\"gb\",\"address_line1\":\"123 Fake Street\",\"address_line2\":\"4th Floor\",\"address_post_code\":\"N12 9XX\",\"cvc\":\"123\",\"exp_month\":12,\"exp_year\":2025,\"name\":\"J. Doe\",\"number\":\"4242424242424242\",\"object\":\"card\"}}"},
	} {
		tc := tc
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			var typ BookingPayment

			if err := typ.Decode(jx.DecodeStr(tc.Input)); err != nil {
				if validateErr, ok := errors.Into[*validate.Error](err); ok {
					t.Skipf("Validation error: %v", validateErr)
					return
				}
				require.NoErrorf(t, err, "Input: %s", tc.Input)
			}

			e := jx.Encoder{}
			typ.Encode(&e)
			require.True(t, std.Valid(e.Bytes()), "Encoded: %s", e.Bytes())

			var typ2 BookingPayment
			require.NoError(t, typ2.Decode(jx.DecodeBytes(e.Bytes())))
		})
	}
}
func TestBookingPaymentCurrency_EncodeDecode(t *testing.T) {
	var typ BookingPaymentCurrency
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 BookingPaymentCurrency
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestBookingPaymentSource_EncodeDecode(t *testing.T) {
	var typ BookingPaymentSource
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 BookingPaymentSource
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestBookingPaymentSource0_EncodeDecode(t *testing.T) {
	var typ BookingPaymentSource0
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 BookingPaymentSource0
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestBookingPaymentSource1_EncodeDecode(t *testing.T) {
	var typ BookingPaymentSource1
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 BookingPaymentSource1
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestBookingPaymentSource1AccountType_EncodeDecode(t *testing.T) {
	var typ BookingPaymentSource1AccountType
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 BookingPaymentSource1AccountType
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestBookingPaymentStatus_EncodeDecode(t *testing.T) {
	var typ BookingPaymentStatus
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 BookingPaymentStatus
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCreateBookingCreated_EncodeDecode(t *testing.T) {
	var typ CreateBookingCreated
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CreateBookingCreated
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCreateBookingPaymentOK_EncodeDecode(t *testing.T) {
	var typ CreateBookingPaymentOK
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CreateBookingPaymentOK
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCreateBookingPaymentOKCurrency_EncodeDecode(t *testing.T) {
	var typ CreateBookingPaymentOKCurrency
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CreateBookingPaymentOKCurrency
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCreateBookingPaymentOKSource_EncodeDecode(t *testing.T) {
	var typ CreateBookingPaymentOKSource
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CreateBookingPaymentOKSource
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCreateBookingPaymentOKSource0_EncodeDecode(t *testing.T) {
	var typ CreateBookingPaymentOKSource0
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CreateBookingPaymentOKSource0
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCreateBookingPaymentOKSource1_EncodeDecode(t *testing.T) {
	var typ CreateBookingPaymentOKSource1
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CreateBookingPaymentOKSource1
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCreateBookingPaymentOKSource1AccountType_EncodeDecode(t *testing.T) {
	var typ CreateBookingPaymentOKSource1AccountType
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CreateBookingPaymentOKSource1AccountType
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCreateBookingPaymentOKStatus_EncodeDecode(t *testing.T) {
	var typ CreateBookingPaymentOKStatus
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CreateBookingPaymentOKStatus
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestGetBookingOK_EncodeDecode(t *testing.T) {
	var typ GetBookingOK
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 GetBookingOK
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestGetBookingsOK_EncodeDecode(t *testing.T) {
	var typ GetBookingsOK
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 GetBookingsOK
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestGetBookingsOKDataItem_EncodeDecode(t *testing.T) {
	var typ GetBookingsOKDataItem
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 GetBookingsOKDataItem
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestGetBookingsOKLinks_EncodeDecode(t *testing.T) {
	var typ GetBookingsOKLinks
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 GetBookingsOKLinks
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestGetStationsOK_EncodeDecode(t *testing.T) {
	var typ GetStationsOK
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 GetStationsOK
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestGetStationsOKDataItem_EncodeDecode(t *testing.T) {
	var typ GetStationsOKDataItem
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 GetStationsOKDataItem
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestGetStationsOKLinks_EncodeDecode(t *testing.T) {
	var typ GetStationsOKLinks
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 GetStationsOKLinks
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestGetTripsOK_EncodeDecode(t *testing.T) {
	var typ GetTripsOK
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 GetTripsOK
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestGetTripsOKDataItem_EncodeDecode(t *testing.T) {
	var typ GetTripsOKDataItem
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 GetTripsOKDataItem
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestGetTripsOKLinks_EncodeDecode(t *testing.T) {
	var typ GetTripsOKLinks
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 GetTripsOKLinks
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestLinksBooking_EncodeDecode(t *testing.T) {
	var typ LinksBooking
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 LinksBooking
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestLinksSelf_EncodeDecode(t *testing.T) {
	var typ LinksSelf
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 LinksSelf
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestNewBookingReq_EncodeDecode(t *testing.T) {
	var typ NewBookingReq
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 NewBookingReq
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestNewBookingReqLinks_EncodeDecode(t *testing.T) {
	var typ NewBookingReqLinks
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 NewBookingReqLinks
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestProblem_EncodeDecode(t *testing.T) {
	var typ Problem
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Problem
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}

func TestProblem_Examples(t *testing.T) {

	for i, tc := range []struct {
		Input string
	}{
		{Input: "{\"detail\":\"Access is forbidden with the provided credentials.\",\"status\":403,\"title\":\"Forbidden\",\"type\":\"https://example.com/errors/forbidden\"}"},
		{Input: "{\"detail\":\"An unexpected error occurred.\",\"status\":500,\"title\":\"Internal Server Error\",\"type\":\"https://example.com/errors/internal-server-error\"}"},
		{Input: "{\"detail\":\"The request is invalid or missing required parameters.\",\"status\":400,\"title\":\"Bad Request\",\"type\":\"https://example.com/errors/bad-request\"}"},
		{Input: "{\"detail\":\"The requested resource was not found.\",\"status\":404,\"title\":\"Not Found\",\"type\":\"https://example.com/errors/not-found\"}"},
		{Input: "{\"detail\":\"There is a conflict with an existing resource.\",\"status\":409,\"title\":\"Conflict\",\"type\":\"https://example.com/errors/conflict\"}"},
		{Input: "{\"detail\":\"You do not have the necessary permissions.\",\"status\":401,\"title\":\"Unauthorized\",\"type\":\"https://example.com/errors/unauthorized\"}"},
		{Input: "{\"detail\":\"You have exceeded the rate limit.\",\"status\":429,\"title\":\"Too Many Requests\",\"type\":\"https://example.com/errors/too-many-requests\"}"},
	} {
		tc := tc
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			var typ Problem

			if err := typ.Decode(jx.DecodeStr(tc.Input)); err != nil {
				if validateErr, ok := errors.Into[*validate.Error](err); ok {
					t.Skipf("Validation error: %v", validateErr)
					return
				}
				require.NoErrorf(t, err, "Input: %s", tc.Input)
			}

			e := jx.Encoder{}
			typ.Encode(&e)
			require.True(t, std.Valid(e.Bytes()), "Encoded: %s", e.Bytes())

			var typ2 Problem
			require.NoError(t, typ2.Decode(jx.DecodeBytes(e.Bytes())))
		})
	}
}
