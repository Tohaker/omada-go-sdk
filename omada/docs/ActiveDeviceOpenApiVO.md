# ActiveDeviceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePairList** | [**[]ActivePairOpenApiDTO**](ActivePairOpenApiDTO.md) |  | 
**Category** | **string** | It should be a value as follows: basic; ap; l2Switch; l3Switch; gateway | 

## Methods

### NewActiveDeviceOpenApiVO

`func NewActiveDeviceOpenApiVO(activePairList []ActivePairOpenApiDTO, category string, ) *ActiveDeviceOpenApiVO`

NewActiveDeviceOpenApiVO instantiates a new ActiveDeviceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveDeviceOpenApiVOWithDefaults

`func NewActiveDeviceOpenApiVOWithDefaults() *ActiveDeviceOpenApiVO`

NewActiveDeviceOpenApiVOWithDefaults instantiates a new ActiveDeviceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePairList

`func (o *ActiveDeviceOpenApiVO) GetActivePairList() []ActivePairOpenApiDTO`

GetActivePairList returns the ActivePairList field if non-nil, zero value otherwise.

### GetActivePairListOk

`func (o *ActiveDeviceOpenApiVO) GetActivePairListOk() (*[]ActivePairOpenApiDTO, bool)`

GetActivePairListOk returns a tuple with the ActivePairList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePairList

`func (o *ActiveDeviceOpenApiVO) SetActivePairList(v []ActivePairOpenApiDTO)`

SetActivePairList sets ActivePairList field to given value.


### GetCategory

`func (o *ActiveDeviceOpenApiVO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ActiveDeviceOpenApiVO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ActiveDeviceOpenApiVO) SetCategory(v string)`

SetCategory sets Category field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


