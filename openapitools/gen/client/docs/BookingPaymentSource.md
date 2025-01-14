# BookingPaymentSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Number** | **string** | The account number for the bank account, in string form. Must be a current account. | 
**Cvc** | **string** | Card security code, 3 or 4 digits usually found on the back of the card. | 
**ExpMonth** | **int64** | Two-digit number representing the card&#39;s expiration month. | 
**ExpYear** | **int64** | Four-digit number representing the card&#39;s expiration year. | 
**AddressLine1** | Pointer to **string** |  | [optional] 
**AddressLine2** | Pointer to **string** |  | [optional] 
**AddressCity** | Pointer to **string** |  | [optional] 
**AddressCountry** | **string** |  | 
**AddressPostCode** | Pointer to **string** |  | [optional] 
**SortCode** | Pointer to **string** | The sort code for the bank account, in string form. Must be a six-digit number. | [optional] 
**AccountType** | **string** | The type of entity that holds the account. This can be either &#x60;individual&#x60; or &#x60;company&#x60;. | 
**BankName** | **string** | The name of the bank associated with the routing number. | 
**Country** | **string** | Two-letter country code (ISO 3166-1 alpha-2). | 

## Methods

### NewBookingPaymentSource

`func NewBookingPaymentSource(name string, number string, cvc string, expMonth int64, expYear int64, addressCountry string, accountType string, bankName string, country string, ) *BookingPaymentSource`

NewBookingPaymentSource instantiates a new BookingPaymentSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingPaymentSourceWithDefaults

`func NewBookingPaymentSourceWithDefaults() *BookingPaymentSource`

NewBookingPaymentSourceWithDefaults instantiates a new BookingPaymentSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *BookingPaymentSource) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BookingPaymentSource) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BookingPaymentSource) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *BookingPaymentSource) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetName

`func (o *BookingPaymentSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BookingPaymentSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BookingPaymentSource) SetName(v string)`

SetName sets Name field to given value.


### GetNumber

`func (o *BookingPaymentSource) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *BookingPaymentSource) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *BookingPaymentSource) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetCvc

`func (o *BookingPaymentSource) GetCvc() string`

GetCvc returns the Cvc field if non-nil, zero value otherwise.

### GetCvcOk

`func (o *BookingPaymentSource) GetCvcOk() (*string, bool)`

GetCvcOk returns a tuple with the Cvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvc

`func (o *BookingPaymentSource) SetCvc(v string)`

SetCvc sets Cvc field to given value.


### GetExpMonth

`func (o *BookingPaymentSource) GetExpMonth() int64`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *BookingPaymentSource) GetExpMonthOk() (*int64, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *BookingPaymentSource) SetExpMonth(v int64)`

SetExpMonth sets ExpMonth field to given value.


### GetExpYear

`func (o *BookingPaymentSource) GetExpYear() int64`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *BookingPaymentSource) GetExpYearOk() (*int64, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *BookingPaymentSource) SetExpYear(v int64)`

SetExpYear sets ExpYear field to given value.


### GetAddressLine1

`func (o *BookingPaymentSource) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *BookingPaymentSource) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *BookingPaymentSource) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *BookingPaymentSource) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *BookingPaymentSource) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *BookingPaymentSource) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *BookingPaymentSource) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *BookingPaymentSource) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressCity

`func (o *BookingPaymentSource) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *BookingPaymentSource) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *BookingPaymentSource) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.

### HasAddressCity

`func (o *BookingPaymentSource) HasAddressCity() bool`

HasAddressCity returns a boolean if a field has been set.

### GetAddressCountry

`func (o *BookingPaymentSource) GetAddressCountry() string`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *BookingPaymentSource) GetAddressCountryOk() (*string, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *BookingPaymentSource) SetAddressCountry(v string)`

SetAddressCountry sets AddressCountry field to given value.


### GetAddressPostCode

`func (o *BookingPaymentSource) GetAddressPostCode() string`

GetAddressPostCode returns the AddressPostCode field if non-nil, zero value otherwise.

### GetAddressPostCodeOk

`func (o *BookingPaymentSource) GetAddressPostCodeOk() (*string, bool)`

GetAddressPostCodeOk returns a tuple with the AddressPostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPostCode

`func (o *BookingPaymentSource) SetAddressPostCode(v string)`

SetAddressPostCode sets AddressPostCode field to given value.

### HasAddressPostCode

`func (o *BookingPaymentSource) HasAddressPostCode() bool`

HasAddressPostCode returns a boolean if a field has been set.

### GetSortCode

`func (o *BookingPaymentSource) GetSortCode() string`

GetSortCode returns the SortCode field if non-nil, zero value otherwise.

### GetSortCodeOk

`func (o *BookingPaymentSource) GetSortCodeOk() (*string, bool)`

GetSortCodeOk returns a tuple with the SortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCode

`func (o *BookingPaymentSource) SetSortCode(v string)`

SetSortCode sets SortCode field to given value.

### HasSortCode

`func (o *BookingPaymentSource) HasSortCode() bool`

HasSortCode returns a boolean if a field has been set.

### GetAccountType

`func (o *BookingPaymentSource) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BookingPaymentSource) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BookingPaymentSource) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.


### GetBankName

`func (o *BookingPaymentSource) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *BookingPaymentSource) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *BookingPaymentSource) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetCountry

`func (o *BookingPaymentSource) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BookingPaymentSource) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BookingPaymentSource) SetCountry(v string)`

SetCountry sets Country field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


