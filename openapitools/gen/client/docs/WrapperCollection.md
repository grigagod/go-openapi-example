# WrapperCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **[]map[string]interface{}** | The wrapper for a collection is an array of objects. | [optional] 
**Links** | Pointer to **map[string]interface{}** | A set of hypermedia links which serve as controls for the client. | [optional] [readonly] 

## Methods

### NewWrapperCollection

`func NewWrapperCollection() *WrapperCollection`

NewWrapperCollection instantiates a new WrapperCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWrapperCollectionWithDefaults

`func NewWrapperCollectionWithDefaults() *WrapperCollection`

NewWrapperCollectionWithDefaults instantiates a new WrapperCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WrapperCollection) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WrapperCollection) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WrapperCollection) SetData(v []map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *WrapperCollection) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *WrapperCollection) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WrapperCollection) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WrapperCollection) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WrapperCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


