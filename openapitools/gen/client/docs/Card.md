# Card

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**Name** | **string** | Cardholder&#39;s full name as it appears on the card. | 
**Number** | **string** | The card number, as a string without any separators. On read all but the last four digits will be masked for security. | 
**Cvc** | **string** | Card security code, 3 or 4 digits usually found on the back of the card. | 
**ExpMonth** | **int64** | Two-digit number representing the card&#39;s expiration month. | 
**ExpYear** | **int64** | Four-digit number representing the card&#39;s expiration year. | 
**AddressLine1** | Pointer to **string** |  | [optional] 
**AddressLine2** | Pointer to **string** |  | [optional] 
**AddressCity** | Pointer to **string** |  | [optional] 
**AddressCountry** | **string** |  | 
**AddressPostCode** | Pointer to **string** |  | [optional] 

## Methods

### NewCard

`func NewCard(name string, number string, cvc string, expMonth int64, expYear int64, addressCountry string, ) *Card`

NewCard instantiates a new Card object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardWithDefaults

`func NewCardWithDefaults() *Card`

NewCardWithDefaults instantiates a new Card object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Card) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Card) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Card) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Card) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetName

`func (o *Card) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Card) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Card) SetName(v string)`

SetName sets Name field to given value.


### GetNumber

`func (o *Card) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Card) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Card) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetCvc

`func (o *Card) GetCvc() string`

GetCvc returns the Cvc field if non-nil, zero value otherwise.

### GetCvcOk

`func (o *Card) GetCvcOk() (*string, bool)`

GetCvcOk returns a tuple with the Cvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvc

`func (o *Card) SetCvc(v string)`

SetCvc sets Cvc field to given value.


### GetExpMonth

`func (o *Card) GetExpMonth() int64`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *Card) GetExpMonthOk() (*int64, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *Card) SetExpMonth(v int64)`

SetExpMonth sets ExpMonth field to given value.


### GetExpYear

`func (o *Card) GetExpYear() int64`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *Card) GetExpYearOk() (*int64, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *Card) SetExpYear(v int64)`

SetExpYear sets ExpYear field to given value.


### GetAddressLine1

`func (o *Card) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *Card) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *Card) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *Card) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *Card) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *Card) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *Card) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *Card) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressCity

`func (o *Card) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *Card) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *Card) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.

### HasAddressCity

`func (o *Card) HasAddressCity() bool`

HasAddressCity returns a boolean if a field has been set.

### GetAddressCountry

`func (o *Card) GetAddressCountry() string`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *Card) GetAddressCountryOk() (*string, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *Card) SetAddressCountry(v string)`

SetAddressCountry sets AddressCountry field to given value.


### GetAddressPostCode

`func (o *Card) GetAddressPostCode() string`

GetAddressPostCode returns the AddressPostCode field if non-nil, zero value otherwise.

### GetAddressPostCodeOk

`func (o *Card) GetAddressPostCodeOk() (*string, bool)`

GetAddressPostCodeOk returns a tuple with the AddressPostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPostCode

`func (o *Card) SetAddressPostCode(v string)`

SetAddressPostCode sets AddressPostCode field to given value.

### HasAddressPostCode

`func (o *Card) HasAddressPostCode() bool`

HasAddressPostCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


