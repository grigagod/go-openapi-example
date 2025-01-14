# BookingPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the payment. This will be a unique identifier for the payment, and is used to reference the payment in other objects. | [optional] [readonly] 
**Amount** | Pointer to **float32** | Amount intended to be collected by this payment. A positive decimal figure describing the amount to be collected. | [optional] 
**Currency** | Pointer to **string** | Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. | [optional] 
**Source** | Pointer to [**BookingPaymentSource**](BookingPaymentSource.md) |  | [optional] 
**Status** | Pointer to **string** | The status of the payment, one of &#x60;pending&#x60;, &#x60;succeeded&#x60;, or &#x60;failed&#x60;. | [optional] [readonly] 

## Methods

### NewBookingPayment

`func NewBookingPayment() *BookingPayment`

NewBookingPayment instantiates a new BookingPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingPaymentWithDefaults

`func NewBookingPaymentWithDefaults() *BookingPayment`

NewBookingPaymentWithDefaults instantiates a new BookingPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BookingPayment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BookingPayment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BookingPayment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BookingPayment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAmount

`func (o *BookingPayment) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BookingPayment) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BookingPayment) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BookingPayment) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *BookingPayment) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BookingPayment) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BookingPayment) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BookingPayment) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSource

`func (o *BookingPayment) GetSource() BookingPaymentSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BookingPayment) GetSourceOk() (*BookingPaymentSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BookingPayment) SetSource(v BookingPaymentSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *BookingPayment) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *BookingPayment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BookingPayment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BookingPayment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BookingPayment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


