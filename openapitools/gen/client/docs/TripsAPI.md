# \TripsAPI

All URIs are relative to *https://try.microcks.io/rest/Train+Travel+API/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTrips**](TripsAPI.md#GetTrips) | **Get** /trips | Get available train trips



## GetTrips

> GetTrips200Response GetTrips(ctx).Origin(origin).Destination(destination).Date(date).Page(page).Limit(limit).Bicycles(bicycles).Dogs(dogs).Execute()

Get available train trips



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	origin := "efdbb9d1-02c2-4bc3-afb7-6788d8782b1e" // string | The ID of the origin station
	destination := "b2e783e1-c824-4d63-b37a-d8d698862f1d" // string | The ID of the destination station
	date := time.Now() // time.Time | The date and time of the trip in ISO 8601 format in origin station's timezone.
	page := int32(1) // int32 | The page number to return (optional) (default to 1)
	limit := int32(10) // int32 | The number of items to return per page (optional) (default to 10)
	bicycles := true // bool | Only return trips where bicycles are known to be allowed (optional) (default to false)
	dogs := true // bool | Only return trips where dogs are known to be allowed (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TripsAPI.GetTrips(context.Background()).Origin(origin).Destination(destination).Date(date).Page(page).Limit(limit).Bicycles(bicycles).Dogs(dogs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TripsAPI.GetTrips``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrips`: GetTrips200Response
	fmt.Fprintf(os.Stdout, "Response from `TripsAPI.GetTrips`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTripsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **origin** | **string** | The ID of the origin station | 
 **destination** | **string** | The ID of the destination station | 
 **date** | **time.Time** | The date and time of the trip in ISO 8601 format in origin station&#39;s timezone. | 
 **page** | **int32** | The page number to return | [default to 1]
 **limit** | **int32** | The number of items to return per page | [default to 10]
 **bicycles** | **bool** | Only return trips where bicycles are known to be allowed | [default to false]
 **dogs** | **bool** | Only return trips where dogs are known to be allowed | [default to false]

### Return type

[**GetTrips200Response**](GetTrips200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/problem+json, application/problem+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

