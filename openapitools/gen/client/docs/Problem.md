# Problem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A URI reference that identifies the problem type | [optional] 
**Title** | Pointer to **string** | A short, human-readable summary of the problem type | [optional] 
**Detail** | Pointer to **string** | A human-readable explanation specific to this occurrence of the problem | [optional] 
**Instance** | Pointer to **string** | A URI reference that identifies the specific occurrence of the problem | [optional] 
**Status** | Pointer to **int32** | The HTTP status code | [optional] 

## Methods

### NewProblem

`func NewProblem() *Problem`

NewProblem instantiates a new Problem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemWithDefaults

`func NewProblemWithDefaults() *Problem`

NewProblemWithDefaults instantiates a new Problem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Problem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Problem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Problem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Problem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *Problem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Problem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Problem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Problem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDetail

`func (o *Problem) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *Problem) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *Problem) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *Problem) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetInstance

`func (o *Problem) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *Problem) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *Problem) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *Problem) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetStatus

`func (o *Problem) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Problem) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Problem) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Problem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


