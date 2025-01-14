# \PaymentsAPI

All URIs are relative to *https://try.microcks.io/rest/Train+Travel+API/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBookingPayment**](PaymentsAPI.md#CreateBookingPayment) | **Post** /bookings/{bookingId}/payment | Pay for a Booking



## CreateBookingPayment

> CreateBookingPayment200Response CreateBookingPayment(ctx, bookingId).BookingPayment(bookingPayment).Execute()

Pay for a Booking



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bookingId := "1725ff48-ab45-4bb5-9d02-88745177dedb" // string | The ID of the booking to pay for.
	bookingPayment := *openapiclient.NewBookingPayment() // BookingPayment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.CreateBookingPayment(context.Background(), bookingId).BookingPayment(bookingPayment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.CreateBookingPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBookingPayment`: CreateBookingPayment200Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.CreateBookingPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookingId** | **string** | The ID of the booking to pay for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBookingPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bookingPayment** | [**BookingPayment**](BookingPayment.md) |  | 

### Return type

[**CreateBookingPayment200Response**](CreateBookingPayment200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json, application/problem+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

