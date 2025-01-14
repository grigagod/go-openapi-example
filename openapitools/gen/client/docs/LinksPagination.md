# LinksPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **string** |  | [optional] 
**Prev** | Pointer to **string** |  | [optional] 

## Methods

### NewLinksPagination

`func NewLinksPagination() *LinksPagination`

NewLinksPagination instantiates a new LinksPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinksPaginationWithDefaults

`func NewLinksPaginationWithDefaults() *LinksPagination`

NewLinksPaginationWithDefaults instantiates a new LinksPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *LinksPagination) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *LinksPagination) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *LinksPagination) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *LinksPagination) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrev

`func (o *LinksPagination) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *LinksPagination) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *LinksPagination) SetPrev(v string)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *LinksPagination) HasPrev() bool`

HasPrev returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


