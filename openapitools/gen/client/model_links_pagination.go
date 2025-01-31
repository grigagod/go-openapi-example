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

// checks if the LinksPagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksPagination{}

// LinksPagination struct for LinksPagination
type LinksPagination struct {
	Next *string `json:"next,omitempty"`
	Prev *string `json:"prev,omitempty"`
}

// NewLinksPagination instantiates a new LinksPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksPagination() *LinksPagination {
	this := LinksPagination{}
	return &this
}

// NewLinksPaginationWithDefaults instantiates a new LinksPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksPaginationWithDefaults() *LinksPagination {
	this := LinksPagination{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *LinksPagination) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksPagination) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *LinksPagination) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *LinksPagination) SetNext(v string) {
	o.Next = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *LinksPagination) GetPrev() string {
	if o == nil || IsNil(o.Prev) {
		var ret string
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksPagination) GetPrevOk() (*string, bool) {
	if o == nil || IsNil(o.Prev) {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *LinksPagination) HasPrev() bool {
	if o != nil && !IsNil(o.Prev) {
		return true
	}

	return false
}

// SetPrev gets a reference to the given string and assigns it to the Prev field.
func (o *LinksPagination) SetPrev(v string) {
	o.Prev = &v
}

func (o LinksPagination) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksPagination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Prev) {
		toSerialize["prev"] = o.Prev
	}
	return toSerialize, nil
}

type NullableLinksPagination struct {
	value *LinksPagination
	isSet bool
}

func (v NullableLinksPagination) Get() *LinksPagination {
	return v.value
}

func (v *NullableLinksPagination) Set(val *LinksPagination) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksPagination) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksPagination(val *LinksPagination) *NullableLinksPagination {
	return &NullableLinksPagination{value: val, isSet: true}
}

func (v NullableLinksPagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
