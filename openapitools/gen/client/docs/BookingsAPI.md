# \BookingsAPI

All URIs are relative to *https://try.microcks.io/rest/Train+Travel+API/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBooking**](BookingsAPI.md#CreateBooking) | **Post** /bookings | Create a booking
[**DeleteBooking**](BookingsAPI.md#DeleteBooking) | **Delete** /bookings/{bookingId} | Delete a booking
[**GetBooking**](BookingsAPI.md#GetBooking) | **Get** /bookings/{bookingId} | Get a booking
[**GetBookings**](BookingsAPI.md#GetBookings) | **Get** /bookings | List existing bookings



## CreateBooking

> CreateBooking201Response CreateBooking(ctx).Booking(booking).Execute()

Create a booking



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
	booking := *openapiclient.NewBooking() // Booking | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BookingsAPI.CreateBooking(context.Background()).Booking(booking).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookingsAPI.CreateBooking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBooking`: CreateBooking201Response
	fmt.Fprintf(os.Stdout, "Response from `BookingsAPI.CreateBooking`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBookingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **booking** | [**Booking**](Booking.md) |  | 

### Return type

[**CreateBooking201Response**](CreateBooking201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml, application/problem+json, application/problem+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBooking

> DeleteBooking(ctx, bookingId).Execute()

Delete a booking



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
	bookingId := "1725ff48-ab45-4bb5-9d02-88745177dedb" // string | The ID of the booking to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BookingsAPI.DeleteBooking(context.Background(), bookingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookingsAPI.DeleteBooking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookingId** | **string** | The ID of the booking to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBookingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json, application/problem+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBooking

> CreateBooking201Response GetBooking(ctx, bookingId).Execute()

Get a booking



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
	bookingId := "1725ff48-ab45-4bb5-9d02-88745177dedb" // string | The ID of the booking to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BookingsAPI.GetBooking(context.Background(), bookingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookingsAPI.GetBooking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBooking`: CreateBooking201Response
	fmt.Fprintf(os.Stdout, "Response from `BookingsAPI.GetBooking`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookingId** | **string** | The ID of the booking to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateBooking201Response**](CreateBooking201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/problem+json, application/problem+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBookings

> GetBookings200Response GetBookings(ctx).Page(page).Limit(limit).Execute()

List existing bookings



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
	page := int32(1) // int32 | The page number to return (optional) (default to 1)
	limit := int32(10) // int32 | The number of items to return per page (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BookingsAPI.GetBookings(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookingsAPI.GetBookings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBookings`: GetBookings200Response
	fmt.Fprintf(os.Stdout, "Response from `BookingsAPI.GetBookings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBookingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page number to return | [default to 1]
 **limit** | **int32** | The number of items to return per page | [default to 10]

### Return type

[**GetBookings200Response**](GetBookings200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/problem+json, application/problem+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

