# GetBookings200Response1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Booking**](Booking.md) |  | [optional] 
**Links** | Pointer to [**GetStations200ResponseAllOfLinks**](GetStations200ResponseAllOfLinks.md) |  | [optional] 

## Methods

### NewGetBookings200Response1

`func NewGetBookings200Response1() *GetBookings200Response1`

NewGetBookings200Response1 instantiates a new GetBookings200Response1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBookings200Response1WithDefaults

`func NewGetBookings200Response1WithDefaults() *GetBookings200Response1`

NewGetBookings200Response1WithDefaults instantiates a new GetBookings200Response1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetBookings200Response1) GetData() []Booking`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetBookings200Response1) GetDataOk() (*[]Booking, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetBookings200Response1) SetData(v []Booking)`

SetData sets Data field to given value.

### HasData

`func (o *GetBookings200Response1) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *GetBookings200Response1) GetLinks() GetStations200ResponseAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetBookings200Response1) GetLinksOk() (*GetStations200ResponseAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetBookings200Response1) SetLinks(v GetStations200ResponseAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetBookings200Response1) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


