# GetStations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Station**](Station.md) |  | [optional] 
**Links** | Pointer to [**GetStations200ResponseAllOfLinks**](GetStations200ResponseAllOfLinks.md) |  | [optional] 

## Methods

### NewGetStations200Response

`func NewGetStations200Response() *GetStations200Response`

NewGetStations200Response instantiates a new GetStations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStations200ResponseWithDefaults

`func NewGetStations200ResponseWithDefaults() *GetStations200Response`

NewGetStations200ResponseWithDefaults instantiates a new GetStations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetStations200Response) GetData() []Station`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetStations200Response) GetDataOk() (*[]Station, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetStations200Response) SetData(v []Station)`

SetData sets Data field to given value.

### HasData

`func (o *GetStations200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *GetStations200Response) GetLinks() GetStations200ResponseAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetStations200Response) GetLinksOk() (*GetStations200ResponseAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetStations200Response) SetLinks(v GetStations200ResponseAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetStations200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


