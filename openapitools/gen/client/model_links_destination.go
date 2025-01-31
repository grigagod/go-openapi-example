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
)

// checks if the LinksDestination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksDestination{}

// LinksDestination The link to the destination station resource.
type LinksDestination struct {
	Self *string `json:"self,omitempty"`
}

// NewLinksDestination instantiates a new LinksDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksDestination() *LinksDestination {
	this := LinksDestination{}
	return &this
}

// NewLinksDestinationWithDefaults instantiates a new LinksDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksDestinationWithDefaults() *LinksDestination {
	this := LinksDestination{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *LinksDestination) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksDestination) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *LinksDestination) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *LinksDestination) SetSelf(v string) {
	o.Self = &v
}

func (o LinksDestination) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksDestination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullableLinksDestination struct {
	value *LinksDestination
	isSet bool
}

func (v NullableLinksDestination) Get() *LinksDestination {
	return v.value
}

func (v *NullableLinksDestination) Set(val *LinksDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksDestination(val *LinksDestination) *NullableLinksDestination {
	return &NullableLinksDestination{value: val, isSet: true}
}

func (v NullableLinksDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
