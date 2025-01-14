# GetBookings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Booking**](Booking.md) |  | [optional] 
**Links** | Pointer to [**GetStations200ResponseAllOfLinks**](GetStations200ResponseAllOfLinks.md) |  | [optional] 

## Methods

### NewGetBookings200Response

`func NewGetBookings200Response() *GetBookings200Response`

NewGetBookings200Response instantiates a new GetBookings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBookings200ResponseWithDefaults

`func NewGetBookings200ResponseWithDefaults() *GetBookings200Response`

NewGetBookings200ResponseWithDefaults instantiates a new GetBookings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetBookings200Response) GetData() []Booking`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetBookings200Response) GetDataOk() (*[]Booking, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetBookings200Response) SetData(v []Booking)`

SetData sets Data field to given value.

### HasData

`func (o *GetBookings200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *GetBookings200Response) GetLinks() GetStations200ResponseAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetBookings200Response) GetLinksOk() (*GetStations200ResponseAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetBookings200Response) SetLinks(v GetStations200ResponseAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetBookings200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


