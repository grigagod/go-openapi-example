# BankAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Number** | **string** | The account number for the bank account, in string form. Must be a current account. | 
**SortCode** | Pointer to **string** | The sort code for the bank account, in string form. Must be a six-digit number. | [optional] 
**AccountType** | **string** | The type of entity that holds the account. This can be either &#x60;individual&#x60; or &#x60;company&#x60;. | 
**BankName** | **string** | The name of the bank associated with the routing number. | 
**Country** | **string** | Two-letter country code (ISO 3166-1 alpha-2). | 

## Methods

### NewBankAccount

`func NewBankAccount(name string, number string, accountType string, bankName string, country string, ) *BankAccount`

NewBankAccount instantiates a new BankAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountWithDefaults

`func NewBankAccountWithDefaults() *BankAccount`

NewBankAccountWithDefaults instantiates a new BankAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *BankAccount) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BankAccount) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BankAccount) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *BankAccount) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetName

`func (o *BankAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BankAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BankAccount) SetName(v string)`

SetName sets Name field to given value.


### GetNumber

`func (o *BankAccount) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *BankAccount) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *BankAccount) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetSortCode

`func (o *BankAccount) GetSortCode() string`

GetSortCode returns the SortCode field if non-nil, zero value otherwise.

### GetSortCodeOk

`func (o *BankAccount) GetSortCodeOk() (*string, bool)`

GetSortCodeOk returns a tuple with the SortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCode

`func (o *BankAccount) SetSortCode(v string)`

SetSortCode sets SortCode field to given value.

### HasSortCode

`func (o *BankAccount) HasSortCode() bool`

HasSortCode returns a boolean if a field has been set.

### GetAccountType

`func (o *BankAccount) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BankAccount) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BankAccount) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.


### GetBankName

`func (o *BankAccount) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *BankAccount) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *BankAccount) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetCountry

`func (o *BankAccount) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BankAccount) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BankAccount) SetCountry(v string)`

SetCountry sets Country field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


