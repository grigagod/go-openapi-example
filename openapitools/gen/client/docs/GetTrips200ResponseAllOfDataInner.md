# GetTrips200ResponseAllOfDataInner

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
**Self** | Pointer to **string** |  | [optional] 

## Methods

### NewGetTrips200ResponseAllOfDataInner

`func NewGetTrips200ResponseAllOfDataInner() *GetTrips200ResponseAllOfDataInner`

NewGetTrips200ResponseAllOfDataInner instantiates a new GetTrips200ResponseAllOfDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTrips200ResponseAllOfDataInnerWithDefaults

`func NewGetTrips200ResponseAllOfDataInnerWithDefaults() *GetTrips200ResponseAllOfDataInner`

NewGetTrips200ResponseAllOfDataInnerWithDefaults instantiates a new GetTrips200ResponseAllOfDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetTrips200ResponseAllOfDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTrips200ResponseAllOfDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTrips200ResponseAllOfDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetTrips200ResponseAllOfDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrigin

`func (o *GetTrips200ResponseAllOfDataInner) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *GetTrips200ResponseAllOfDataInner) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *GetTrips200ResponseAllOfDataInner) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *GetTrips200ResponseAllOfDataInner) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetDestination

`func (o *GetTrips200ResponseAllOfDataInner) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *GetTrips200ResponseAllOfDataInner) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *GetTrips200ResponseAllOfDataInner) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *GetTrips200ResponseAllOfDataInner) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDepartureTime

`func (o *GetTrips200ResponseAllOfDataInner) GetDepartureTime() time.Time`

GetDepartureTime returns the DepartureTime field if non-nil, zero value otherwise.

### GetDepartureTimeOk

`func (o *GetTrips200ResponseAllOfDataInner) GetDepartureTimeOk() (*time.Time, bool)`

GetDepartureTimeOk returns a tuple with the DepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureTime

`func (o *GetTrips200ResponseAllOfDataInner) SetDepartureTime(v time.Time)`

SetDepartureTime sets DepartureTime field to given value.

### HasDepartureTime

`func (o *GetTrips200ResponseAllOfDataInner) HasDepartureTime() bool`

HasDepartureTime returns a boolean if a field has been set.

### GetArrivalTime

`func (o *GetTrips200ResponseAllOfDataInner) GetArrivalTime() time.Time`

GetArrivalTime returns the ArrivalTime field if non-nil, zero value otherwise.

### GetArrivalTimeOk

`func (o *GetTrips200ResponseAllOfDataInner) GetArrivalTimeOk() (*time.Time, bool)`

GetArrivalTimeOk returns a tuple with the ArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalTime

`func (o *GetTrips200ResponseAllOfDataInner) SetArrivalTime(v time.Time)`

SetArrivalTime sets ArrivalTime field to given value.

### HasArrivalTime

`func (o *GetTrips200ResponseAllOfDataInner) HasArrivalTime() bool`

HasArrivalTime returns a boolean if a field has been set.

### GetOperator

`func (o *GetTrips200ResponseAllOfDataInner) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *GetTrips200ResponseAllOfDataInner) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *GetTrips200ResponseAllOfDataInner) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *GetTrips200ResponseAllOfDataInner) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetPrice

`func (o *GetTrips200ResponseAllOfDataInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetTrips200ResponseAllOfDataInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetTrips200ResponseAllOfDataInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetTrips200ResponseAllOfDataInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetBicyclesAllowed

`func (o *GetTrips200ResponseAllOfDataInner) GetBicyclesAllowed() bool`

GetBicyclesAllowed returns the BicyclesAllowed field if non-nil, zero value otherwise.

### GetBicyclesAllowedOk

`func (o *GetTrips200ResponseAllOfDataInner) GetBicyclesAllowedOk() (*bool, bool)`

GetBicyclesAllowedOk returns a tuple with the BicyclesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBicyclesAllowed

`func (o *GetTrips200ResponseAllOfDataInner) SetBicyclesAllowed(v bool)`

SetBicyclesAllowed sets BicyclesAllowed field to given value.

### HasBicyclesAllowed

`func (o *GetTrips200ResponseAllOfDataInner) HasBicyclesAllowed() bool`

HasBicyclesAllowed returns a boolean if a field has been set.

### GetDogsAllowed

`func (o *GetTrips200ResponseAllOfDataInner) GetDogsAllowed() bool`

GetDogsAllowed returns the DogsAllowed field if non-nil, zero value otherwise.

### GetDogsAllowedOk

`func (o *GetTrips200ResponseAllOfDataInner) GetDogsAllowedOk() (*bool, bool)`

GetDogsAllowedOk returns a tuple with the DogsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDogsAllowed

`func (o *GetTrips200ResponseAllOfDataInner) SetDogsAllowed(v bool)`

SetDogsAllowed sets DogsAllowed field to given value.

### HasDogsAllowed

`func (o *GetTrips200ResponseAllOfDataInner) HasDogsAllowed() bool`

HasDogsAllowed returns a boolean if a field has been set.

### GetSelf

`func (o *GetTrips200ResponseAllOfDataInner) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *GetTrips200ResponseAllOfDataInner) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *GetTrips200ResponseAllOfDataInner) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *GetTrips200ResponseAllOfDataInner) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


