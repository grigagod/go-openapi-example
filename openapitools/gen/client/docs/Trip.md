# Trip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the trip | [optional] 
**Origin** | Pointer to **string** | The starting station of the trip | [optional] 
**Destination** | Pointer to **string** | The destination station of the trip | [optional] 
**DepartureTime** | Pointer to **time.Time** | The date and time when the trip departs | [optional] 
**ArrivalTime** | Pointer to **time.Time** | The date and time when the trip arrives | [optional] 
**Operator** | Pointer to **string** | The name of the operator of the trip | [optional] 
**Price** | Pointer to **float32** | The cost of the trip | [optional] 
**BicyclesAllowed** | Pointer to **bool** | Indicates whether bicycles are allowed on the trip | [optional] 
**DogsAllowed** | Pointer to **bool** | Indicates whether dogs are allowed on the trip | [optional] 

## Methods

### NewTrip

`func NewTrip() *Trip`

NewTrip instantiates a new Trip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTripWithDefaults

`func NewTripWithDefaults() *Trip`

NewTripWithDefaults instantiates a new Trip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Trip) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Trip) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Trip) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Trip) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrigin

`func (o *Trip) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Trip) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Trip) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *Trip) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetDestination

`func (o *Trip) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Trip) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Trip) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *Trip) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDepartureTime

`func (o *Trip) GetDepartureTime() time.Time`

GetDepartureTime returns the DepartureTime field if non-nil, zero value otherwise.

### GetDepartureTimeOk

`func (o *Trip) GetDepartureTimeOk() (*time.Time, bool)`

GetDepartureTimeOk returns a tuple with the DepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureTime

`func (o *Trip) SetDepartureTime(v time.Time)`

SetDepartureTime sets DepartureTime field to given value.

### HasDepartureTime

`func (o *Trip) HasDepartureTime() bool`

HasDepartureTime returns a boolean if a field has been set.

### GetArrivalTime

`func (o *Trip) GetArrivalTime() time.Time`

GetArrivalTime returns the ArrivalTime field if non-nil, zero value otherwise.

### GetArrivalTimeOk

`func (o *Trip) GetArrivalTimeOk() (*time.Time, bool)`

GetArrivalTimeOk returns a tuple with the ArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalTime

`func (o *Trip) SetArrivalTime(v time.Time)`

SetArrivalTime sets ArrivalTime field to given value.

### HasArrivalTime

`func (o *Trip) HasArrivalTime() bool`

HasArrivalTime returns a boolean if a field has been set.

### GetOperator

`func (o *Trip) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Trip) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Trip) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *Trip) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetPrice

`func (o *Trip) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Trip) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Trip) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Trip) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetBicyclesAllowed

`func (o *Trip) GetBicyclesAllowed() bool`

GetBicyclesAllowed returns the BicyclesAllowed field if non-nil, zero value otherwise.

### GetBicyclesAllowedOk

`func (o *Trip) GetBicyclesAllowedOk() (*bool, bool)`

GetBicyclesAllowedOk returns a tuple with the BicyclesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBicyclesAllowed

`func (o *Trip) SetBicyclesAllowed(v bool)`

SetBicyclesAllowed sets BicyclesAllowed field to given value.

### HasBicyclesAllowed

`func (o *Trip) HasBicyclesAllowed() bool`

HasBicyclesAllowed returns a boolean if a field has been set.

### GetDogsAllowed

`func (o *Trip) GetDogsAllowed() bool`

GetDogsAllowed returns the DogsAllowed field if non-nil, zero value otherwise.

### GetDogsAllowedOk

`func (o *Trip) GetDogsAllowedOk() (*bool, bool)`

GetDogsAllowedOk returns a tuple with the DogsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDogsAllowed

`func (o *Trip) SetDogsAllowed(v bool)`

SetDogsAllowed sets DogsAllowed field to given value.

### HasDogsAllowed

`func (o *Trip) HasDogsAllowed() bool`

HasDogsAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


