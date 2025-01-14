/*
Train Travel API

API for finding and booking train trips across Europe.  ## Run in Postman  Experiment with this API in Postman, using our Postman Collection.  [![Run In Postman](https://run.pstmn.io/button.svg =128pxx32px)](https://app.getpostman.com/run-collection/9265903-7a75a0d0-b108-4436-ba54-c6139698dc08?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D9265903-7a75a0d0-b108-4436-ba54-c6139698dc08%26entityType%3Dcollection%26workspaceId%3Df507f69d-9564-419c-89a2-cb8e4c8c7b8f)

API version: 1.2.0
Contact: support@example.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oas

import (
	"encoding/json"
	"time"
)

// checks if the GetTrips200ResponseAllOfDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTrips200ResponseAllOfDataInner{}

// GetTrips200ResponseAllOfDataInner struct for GetTrips200ResponseAllOfDataInner
type GetTrips200ResponseAllOfDataInner struct {
	// Unique identifier for the trip
	Id *string `json:"id,omitempty"`
	// The starting station of the trip
	Origin *string `json:"origin,omitempty"`
	// The destination station of the trip
	Destination *string `json:"destination,omitempty"`
	// The date and time when the trip departs
	DepartureTime *time.Time `json:"departure_time,omitempty"`
	// The date and time when the trip arrives
	ArrivalTime *time.Time `json:"arrival_time,omitempty"`
	// The name of the operator of the trip
	Operator *string `json:"operator,omitempty"`
	// The cost of the trip
	Price *float32 `json:"price,omitempty"`
	// Indicates whether bicycles are allowed on the trip
	BicyclesAllowed *bool `json:"bicycles_allowed,omitempty"`
	// Indicates whether dogs are allowed on the trip
	DogsAllowed *bool   `json:"dogs_allowed,omitempty"`
	Self        *string `json:"self,omitempty"`
}

// NewGetTrips200ResponseAllOfDataInner instantiates a new GetTrips200ResponseAllOfDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTrips200ResponseAllOfDataInner() *GetTrips200ResponseAllOfDataInner {
	this := GetTrips200ResponseAllOfDataInner{}
	return &this
}

// NewGetTrips200ResponseAllOfDataInnerWithDefaults instantiates a new GetTrips200ResponseAllOfDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTrips200ResponseAllOfDataInnerWithDefaults() *GetTrips200ResponseAllOfDataInner {
	this := GetTrips200ResponseAllOfDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetTrips200ResponseAllOfDataInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrips200ResponseAllOfDataInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetTrips200ResponseAllOfDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetTrips200ResponseAllOfDataInner) SetId(v string) {
	o.Id = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *GetTrips200ResponseAllOfDataInner) GetOrigin() string {
	if o == nil || IsNil(o.Origin) {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrips200ResponseAllOfDataInner) GetOriginOk() (*string, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *GetTrips200ResponseAllOfDataInner) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *GetTrips200ResponseAllOfDataInner) SetOrigin(v string) {
	o.Origin = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *GetTrips200ResponseAllOfDataInner) GetDestination() string {
	if o == nil || IsNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrips200ResponseAllOfDataInner) GetDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *GetTrips200ResponseAllOfDataInner) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *GetTrips200ResponseAllOfDataInner) SetDestination(v string) {
	o.Destination = &v
}

// GetDepartureTime returns the DepartureTime field value if set, zero value otherwise.
func (o *GetTrips200ResponseAllOfDataInner) GetDepartureTime() time.Time {
	if o == nil || IsNil(o.DepartureTime) {
		var ret time.Time
		return ret
	}
	return *o.DepartureTime
}

// GetDepartureTimeOk returns a tuple with the DepartureTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrips200ResponseAllOfDataInner) GetDepartureTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DepartureTime) {
		return nil, false
	}
	return o.DepartureTime, true
}

// HasDepartureTime returns a boolean if a field has been set.
func (o *GetTrips200ResponseAllOfDataInner) HasDepartureTime() bool {
	if o != nil && !IsNil(o.DepartureTime) {
		return true
	}

	return false
}

// SetDepartureTime gets a reference to the given time.Time and assigns it to the DepartureTime field.
func (o *GetTrips200ResponseAllOfDataInner) SetDepartureTime(v time.Time) {
	o.DepartureTime = &v
}

// GetArrivalTime returns the ArrivalTime field value if set, zero value otherwise.
func (o *GetTrips200ResponseAllOfDataInner) GetArrivalTime() time.Time {
	if o == nil || IsNil(o.ArrivalTime) {
		var ret time.Time
		return ret
	}
	return *o.ArrivalTime
}

// GetArrivalTimeOk returns a tuple with the ArrivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrips200ResponseAllOfDataInner) GetArrivalTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ArrivalTime) {
		return nil, false
	}
	return o.ArrivalTime, true
}

// HasArrivalTime returns a boolean if a field has been set.
func (o *GetTrips200ResponseAllOfDataInner) HasArrivalTime() bool {
	if o != nil && !IsNil(o.ArrivalTime) {
		return true
	}

	return false
}

// SetArrivalTime gets a reference to the given time.Time and assigns it to the ArrivalTime field.
func (o *GetTrips200ResponseAllOfDataInner) SetArrivalTime(v time.Time) {
	o.ArrivalTime = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *GetTrips200ResponseAllOfDataInner) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrips200ResponseAllOfDataInner) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *GetTrips200ResponseAllOfDataInner) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *GetTrips200ResponseAllOfDataInner) SetOperator(v string) {
	o.Operator = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GetTrips200ResponseAllOfDataInner) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrips200ResponseAllOfDataInner) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *GetTrips200ResponseAllOfDataInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *GetTrips200ResponseAllOfDataInner) SetPrice(v float32) {
	o.Price = &v
}

// GetBicyclesAllowed returns the BicyclesAllowed field value if set, zero value otherwise.
func (o *GetTrips200ResponseAllOfDataInner) GetBicyclesAllowed() bool {
	if o == nil || IsNil(o.BicyclesAllowed) {
		var ret bool
		return ret
	}
	return *o.BicyclesAllowed
}

// GetBicyclesAllowedOk returns a tuple with the BicyclesAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrips200ResponseAllOfDataInner) GetBicyclesAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.BicyclesAllowed) {
		return nil, false
	}
	return o.BicyclesAllowed, true
}

// HasBicyclesAllowed returns a boolean if a field has been set.
func (o *GetTrips200ResponseAllOfDataInner) HasBicyclesAllowed() bool {
	if o != nil && !IsNil(o.BicyclesAllowed) {
		return true
	}

	return false
}

// SetBicyclesAllowed gets a reference to the given bool and assigns it to the BicyclesAllowed field.
func (o *GetTrips200ResponseAllOfDataInner) SetBicyclesAllowed(v bool) {
	o.BicyclesAllowed = &v
}

// GetDogsAllowed returns the DogsAllowed field value if set, zero value otherwise.
func (o *GetTrips200ResponseAllOfDataInner) GetDogsAllowed() bool {
	if o == nil || IsNil(o.DogsAllowed) {
		var ret bool
		return ret
	}
	return *o.DogsAllowed
}

// GetDogsAllowedOk returns a tuple with the DogsAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrips200ResponseAllOfDataInner) GetDogsAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.DogsAllowed) {
		return nil, false
	}
	return o.DogsAllowed, true
}

// HasDogsAllowed returns a boolean if a field has been set.
func (o *GetTrips200ResponseAllOfDataInner) HasDogsAllowed() bool {
	if o != nil && !IsNil(o.DogsAllowed) {
		return true
	}

	return false
}

// SetDogsAllowed gets a reference to the given bool and assigns it to the DogsAllowed field.
func (o *GetTrips200ResponseAllOfDataInner) SetDogsAllowed(v bool) {
	o.DogsAllowed = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *GetTrips200ResponseAllOfDataInner) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTrips200ResponseAllOfDataInner) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *GetTrips200ResponseAllOfDataInner) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *GetTrips200ResponseAllOfDataInner) SetSelf(v string) {
	o.Self = &v
}

func (o GetTrips200ResponseAllOfDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTrips200ResponseAllOfDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.DepartureTime) {
		toSerialize["departure_time"] = o.DepartureTime
	}
	if !IsNil(o.ArrivalTime) {
		toSerialize["arrival_time"] = o.ArrivalTime
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.BicyclesAllowed) {
		toSerialize["bicycles_allowed"] = o.BicyclesAllowed
	}
	if !IsNil(o.DogsAllowed) {
		toSerialize["dogs_allowed"] = o.DogsAllowed
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullableGetTrips200ResponseAllOfDataInner struct {
	value *GetTrips200ResponseAllOfDataInner
	isSet bool
}

func (v NullableGetTrips200ResponseAllOfDataInner) Get() *GetTrips200ResponseAllOfDataInner {
	return v.value
}

func (v *NullableGetTrips200ResponseAllOfDataInner) Set(val *GetTrips200ResponseAllOfDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTrips200ResponseAllOfDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTrips200ResponseAllOfDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTrips200ResponseAllOfDataInner(val *GetTrips200ResponseAllOfDataInner) *NullableGetTrips200ResponseAllOfDataInner {
	return &NullableGetTrips200ResponseAllOfDataInner{value: val, isSet: true}
}

func (v NullableGetTrips200ResponseAllOfDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTrips200ResponseAllOfDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
