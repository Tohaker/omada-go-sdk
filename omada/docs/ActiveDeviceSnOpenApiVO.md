# ActiveDeviceSnOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePairList** | [**[]ActivePairSnOpenApiDTO**](ActivePairSnOpenApiDTO.md) |  | 
**Category** | **string** | It should be a value as follows: basic; ap; l2Switch; l3Switch; gateway | 

## Methods

### NewActiveDeviceSnOpenApiVO

`func NewActiveDeviceSnOpenApiVO(activePairList []ActivePairSnOpenApiDTO, category string, ) *ActiveDeviceSnOpenApiVO`

NewActiveDeviceSnOpenApiVO instantiates a new ActiveDeviceSnOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveDeviceSnOpenApiVOWithDefaults

`func NewActiveDeviceSnOpenApiVOWithDefaults() *ActiveDeviceSnOpenApiVO`

NewActiveDeviceSnOpenApiVOWithDefaults instantiates a new ActiveDeviceSnOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePairList

`func (o *ActiveDeviceSnOpenApiVO) GetActivePairList() []ActivePairSnOpenApiDTO`

GetActivePairList returns the ActivePairList field if non-nil, zero value otherwise.

### GetActivePairListOk

`func (o *ActiveDeviceSnOpenApiVO) GetActivePairListOk() (*[]ActivePairSnOpenApiDTO, bool)`

GetActivePairListOk returns a tuple with the ActivePairList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePairList

`func (o *ActiveDeviceSnOpenApiVO) SetActivePairList(v []ActivePairSnOpenApiDTO)`

SetActivePairList sets ActivePairList field to given value.


### GetCategory

`func (o *ActiveDeviceSnOpenApiVO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ActiveDeviceSnOpenApiVO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ActiveDeviceSnOpenApiVO) SetCategory(v string)`

SetCategory sets Category field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


