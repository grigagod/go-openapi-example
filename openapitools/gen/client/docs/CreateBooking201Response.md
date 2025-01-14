# CreateBooking201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the booking | [optional] [readonly] 
**TripId** | Pointer to **string** | Identifier of the booked trip | [optional] 
**PassengerName** | Pointer to **string** | Name of the passenger | [optional] 
**HasBicycle** | Pointer to **bool** | Indicates whether the passenger has a bicycle. | [optional] 
**HasDog** | Pointer to **bool** | Indicates whether the passenger has a dog. | [optional] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewCreateBooking201Response

`func NewCreateBooking201Response() *CreateBooking201Response`

NewCreateBooking201Response instantiates a new CreateBooking201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBooking201ResponseWithDefaults

`func NewCreateBooking201ResponseWithDefaults() *CreateBooking201Response`

NewCreateBooking201ResponseWithDefaults instantiates a new CreateBooking201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateBooking201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateBooking201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateBooking201Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateBooking201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTripId

`func (o *CreateBooking201Response) GetTripId() string`

GetTripId returns the TripId field if non-nil, zero value otherwise.

### GetTripIdOk

`func (o *CreateBooking201Response) GetTripIdOk() (*string, bool)`

GetTripIdOk returns a tuple with the TripId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTripId

`func (o *CreateBooking201Response) SetTripId(v string)`

SetTripId sets TripId field to given value.

### HasTripId

`func (o *CreateBooking201Response) HasTripId() bool`

HasTripId returns a boolean if a field has been set.

### GetPassengerName

`func (o *CreateBooking201Response) GetPassengerName() string`

GetPassengerName returns the PassengerName field if non-nil, zero value otherwise.

### GetPassengerNameOk

`func (o *CreateBooking201Response) GetPassengerNameOk() (*string, bool)`

GetPassengerNameOk returns a tuple with the PassengerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassengerName

`func (o *CreateBooking201Response) SetPassengerName(v string)`

SetPassengerName sets PassengerName field to given value.

### HasPassengerName

`func (o *CreateBooking201Response) HasPassengerName() bool`

HasPassengerName returns a boolean if a field has been set.

### GetHasBicycle

`func (o *CreateBooking201Response) GetHasBicycle() bool`

GetHasBicycle returns the HasBicycle field if non-nil, zero value otherwise.

### GetHasBicycleOk

`func (o *CreateBooking201Response) GetHasBicycleOk() (*bool, bool)`

GetHasBicycleOk returns a tuple with the HasBicycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBicycle

`func (o *CreateBooking201Response) SetHasBicycle(v bool)`

SetHasBicycle sets HasBicycle field to given value.

### HasHasBicycle

`func (o *CreateBooking201Response) HasHasBicycle() bool`

HasHasBicycle returns a boolean if a field has been set.

### GetHasDog

`func (o *CreateBooking201Response) GetHasDog() bool`

GetHasDog returns the HasDog field if non-nil, zero value otherwise.

### GetHasDogOk

`func (o *CreateBooking201Response) GetHasDogOk() (*bool, bool)`

GetHasDogOk returns a tuple with the HasDog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDog

`func (o *CreateBooking201Response) SetHasDog(v bool)`

SetHasDog sets HasDog field to given value.

### HasHasDog

`func (o *CreateBooking201Response) HasHasDog() bool`

HasHasDog returns a boolean if a field has been set.

### GetLinks

`func (o *CreateBooking201Response) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateBooking201Response) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateBooking201Response) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateBooking201Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


