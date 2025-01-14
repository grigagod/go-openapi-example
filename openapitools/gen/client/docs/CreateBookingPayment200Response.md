# CreateBookingPayment200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the payment. This will be a unique identifier for the payment, and is used to reference the payment in other objects. | [optional] [readonly] 
**Amount** | Pointer to **float32** | Amount intended to be collected by this payment. A positive decimal figure describing the amount to be collected. | [optional] 
**Currency** | Pointer to **string** | Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. | [optional] 
**Source** | Pointer to [**BookingPaymentSource**](BookingPaymentSource.md) |  | [optional] 
**Status** | Pointer to **string** | The status of the payment, one of &#x60;pending&#x60;, &#x60;succeeded&#x60;, or &#x60;failed&#x60;. | [optional] [readonly] 
**Links** | Pointer to [**LinksBooking**](LinksBooking.md) |  | [optional] 

## Methods

### NewCreateBookingPayment200Response

`func NewCreateBookingPayment200Response() *CreateBookingPayment200Response`

NewCreateBookingPayment200Response instantiates a new CreateBookingPayment200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBookingPayment200ResponseWithDefaults

`func NewCreateBookingPayment200ResponseWithDefaults() *CreateBookingPayment200Response`

NewCreateBookingPayment200ResponseWithDefaults instantiates a new CreateBookingPayment200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateBookingPayment200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateBookingPayment200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateBookingPayment200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateBookingPayment200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAmount

`func (o *CreateBookingPayment200Response) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateBookingPayment200Response) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateBookingPayment200Response) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CreateBookingPayment200Response) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *CreateBookingPayment200Response) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateBookingPayment200Response) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateBookingPayment200Response) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateBookingPayment200Response) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSource

`func (o *CreateBookingPayment200Response) GetSource() BookingPaymentSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateBookingPayment200Response) GetSourceOk() (*BookingPaymentSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateBookingPayment200Response) SetSource(v BookingPaymentSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *CreateBookingPayment200Response) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *CreateBookingPayment200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateBookingPayment200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateBookingPayment200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateBookingPayment200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *CreateBookingPayment200Response) GetLinks() LinksBooking`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateBookingPayment200Response) GetLinksOk() (*LinksBooking, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateBookingPayment200Response) SetLinks(v LinksBooking)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateBookingPayment200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


