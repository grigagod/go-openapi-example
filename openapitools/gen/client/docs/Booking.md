# Booking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the booking | [optional] [readonly] 
**TripId** | Pointer to **string** | Identifier of the booked trip | [optional] 
**PassengerName** | Pointer to **string** | Name of the passenger | [optional] 
**HasBicycle** | Pointer to **bool** | Indicates whether the passenger has a bicycle. | [optional] 
**HasDog** | Pointer to **bool** | Indicates whether the passenger has a dog. | [optional] 

## Methods

### NewBooking

`func NewBooking() *Booking`

NewBooking instantiates a new Booking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingWithDefaults

`func NewBookingWithDefaults() *Booking`

NewBookingWithDefaults instantiates a new Booking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Booking) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Booking) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Booking) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Booking) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTripId

`func (o *Booking) GetTripId() string`

GetTripId returns the TripId field if non-nil, zero value otherwise.

### GetTripIdOk

`func (o *Booking) GetTripIdOk() (*string, bool)`

GetTripIdOk returns a tuple with the TripId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTripId

`func (o *Booking) SetTripId(v string)`

SetTripId sets TripId field to given value.

### HasTripId

`func (o *Booking) HasTripId() bool`

HasTripId returns a boolean if a field has been set.

### GetPassengerName

`func (o *Booking) GetPassengerName() string`

GetPassengerName returns the PassengerName field if non-nil, zero value otherwise.

### GetPassengerNameOk

`func (o *Booking) GetPassengerNameOk() (*string, bool)`

GetPassengerNameOk returns a tuple with the PassengerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassengerName

`func (o *Booking) SetPassengerName(v string)`

SetPassengerName sets PassengerName field to given value.

### HasPassengerName

`func (o *Booking) HasPassengerName() bool`

HasPassengerName returns a boolean if a field has been set.

### GetHasBicycle

`func (o *Booking) GetHasBicycle() bool`

GetHasBicycle returns the HasBicycle field if non-nil, zero value otherwise.

### GetHasBicycleOk

`func (o *Booking) GetHasBicycleOk() (*bool, bool)`

GetHasBicycleOk returns a tuple with the HasBicycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBicycle

`func (o *Booking) SetHasBicycle(v bool)`

SetHasBicycle sets HasBicycle field to given value.

### HasHasBicycle

`func (o *Booking) HasHasBicycle() bool`

HasHasBicycle returns a boolean if a field has been set.

### GetHasDog

`func (o *Booking) GetHasDog() bool`

GetHasDog returns the HasDog field if non-nil, zero value otherwise.

### GetHasDogOk

`func (o *Booking) GetHasDogOk() (*bool, bool)`

GetHasDogOk returns a tuple with the HasDog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDog

`func (o *Booking) SetHasDog(v bool)`

SetHasDog sets HasDog field to given value.

### HasHasDog

`func (o *Booking) HasHasDog() bool`

HasHasDog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


