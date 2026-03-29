# ModelBaseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVersion** | Pointer to **[]string** | Model version list, software version, such as \&quot;2.5.0 Build 20190118 Rel. 64821\&quot;. | [optional] 
**ModelTypeInfo** | [**ModelTypeInfoOpenApiVO**](ModelTypeInfoOpenApiVO.md) |  | 

## Methods

### NewModelBaseInfo

`func NewModelBaseInfo(modelTypeInfo ModelTypeInfoOpenApiVO, ) *ModelBaseInfo`

NewModelBaseInfo instantiates a new ModelBaseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelBaseInfoWithDefaults

`func NewModelBaseInfoWithDefaults() *ModelBaseInfo`

NewModelBaseInfoWithDefaults instantiates a new ModelBaseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVersion

`func (o *ModelBaseInfo) GetCurrentVersion() []string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *ModelBaseInfo) GetCurrentVersionOk() (*[]string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *ModelBaseInfo) SetCurrentVersion(v []string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *ModelBaseInfo) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetModelTypeInfo

`func (o *ModelBaseInfo) GetModelTypeInfo() ModelTypeInfoOpenApiVO`

GetModelTypeInfo returns the ModelTypeInfo field if non-nil, zero value otherwise.

### GetModelTypeInfoOk

`func (o *ModelBaseInfo) GetModelTypeInfoOk() (*ModelTypeInfoOpenApiVO, bool)`

GetModelTypeInfoOk returns a tuple with the ModelTypeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeInfo

`func (o *ModelBaseInfo) SetModelTypeInfo(v ModelTypeInfoOpenApiVO)`

SetModelTypeInfo sets ModelTypeInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


