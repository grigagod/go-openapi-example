# GetTrips200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetTrips200ResponseAllOfDataInner**](GetTrips200ResponseAllOfDataInner.md) |  | [optional] 
**Links** | Pointer to [**GetStations200ResponseAllOfLinks**](GetStations200ResponseAllOfLinks.md) |  | [optional] 

## Methods

### NewGetTrips200Response

`func NewGetTrips200Response() *GetTrips200Response`

NewGetTrips200Response instantiates a new GetTrips200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTrips200ResponseWithDefaults

`func NewGetTrips200ResponseWithDefaults() *GetTrips200Response`

NewGetTrips200ResponseWithDefaults instantiates a new GetTrips200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTrips200Response) GetData() []GetTrips200ResponseAllOfDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTrips200Response) GetDataOk() (*[]GetTrips200ResponseAllOfDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTrips200Response) SetData(v []GetTrips200ResponseAllOfDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetTrips200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *GetTrips200Response) GetLinks() GetStations200ResponseAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetTrips200Response) GetLinksOk() (*GetStations200ResponseAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetTrips200Response) SetLinks(v GetStations200ResponseAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetTrips200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


