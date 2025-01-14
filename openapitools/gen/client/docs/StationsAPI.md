# \StationsAPI

All URIs are relative to *https://try.microcks.io/rest/Train+Travel+API/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStations**](StationsAPI.md#GetStations) | **Get** /stations | Get a list of train stations



## GetStations

> GetStations200Response GetStations(ctx).Page(page).Limit(limit).Coordinates(coordinates).Search(search).Country(country).Execute()

Get a list of train stations



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
	coordinates := "52.5200,13.4050" // string | The latitude and longitude of the user's location, to narrow down the search results to sites within a proximity of this location.  (optional)
	search := "search_example" // string | A search term to filter the list of stations by name or address.  (optional)
	country := "DE" // string | Filter stations by country code (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StationsAPI.GetStations(context.Background()).Page(page).Limit(limit).Coordinates(coordinates).Search(search).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StationsAPI.GetStations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStations`: GetStations200Response
	fmt.Fprintf(os.Stdout, "Response from `StationsAPI.GetStations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page number to return | [default to 1]
 **limit** | **int32** | The number of items to return per page | [default to 10]
 **coordinates** | **string** | The latitude and longitude of the user&#39;s location, to narrow down the search results to sites within a proximity of this location.  | 
 **search** | **string** | A search term to filter the list of stations by name or address.  | 
 **country** | **string** | Filter stations by country code | 

### Return type

[**GetStations200Response**](GetStations200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/problem+json, application/problem+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

