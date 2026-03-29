# PlanUpgradeModelList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVersion** | Pointer to **[]string** | Model version list, software version, such as \&quot;2.5.0 Build 20190118 Rel. 64821\&quot; | [optional] 
**ModelTypeInfo** | Pointer to [**ModelTypeInfoOpenApiVO**](ModelTypeInfoOpenApiVO.md) |  | [optional] 
**Status** | Pointer to **int32** | Model critical status should be a value as follows: 0: good; 1: critical | [optional] 
**UpgradeStatus** | Pointer to **int32** | Model upgrade status should be a value as follows: 0: up-to-date; 1: new Version Available | [optional] 

## Methods

### NewPlanUpgradeModelList

`func NewPlanUpgradeModelList() *PlanUpgradeModelList`

NewPlanUpgradeModelList instantiates a new PlanUpgradeModelList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanUpgradeModelListWithDefaults

`func NewPlanUpgradeModelListWithDefaults() *PlanUpgradeModelList`

NewPlanUpgradeModelListWithDefaults instantiates a new PlanUpgradeModelList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVersion

`func (o *PlanUpgradeModelList) GetCurrentVersion() []string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *PlanUpgradeModelList) GetCurrentVersionOk() (*[]string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *PlanUpgradeModelList) SetCurrentVersion(v []string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *PlanUpgradeModelList) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetModelTypeInfo

`func (o *PlanUpgradeModelList) GetModelTypeInfo() ModelTypeInfoOpenApiVO`

GetModelTypeInfo returns the ModelTypeInfo field if non-nil, zero value otherwise.

### GetModelTypeInfoOk

`func (o *PlanUpgradeModelList) GetModelTypeInfoOk() (*ModelTypeInfoOpenApiVO, bool)`

GetModelTypeInfoOk returns a tuple with the ModelTypeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeInfo

`func (o *PlanUpgradeModelList) SetModelTypeInfo(v ModelTypeInfoOpenApiVO)`

SetModelTypeInfo sets ModelTypeInfo field to given value.

### HasModelTypeInfo

`func (o *PlanUpgradeModelList) HasModelTypeInfo() bool`

HasModelTypeInfo returns a boolean if a field has been set.

### GetStatus

`func (o *PlanUpgradeModelList) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlanUpgradeModelList) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlanUpgradeModelList) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlanUpgradeModelList) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpgradeStatus

`func (o *PlanUpgradeModelList) GetUpgradeStatus() int32`

GetUpgradeStatus returns the UpgradeStatus field if non-nil, zero value otherwise.

### GetUpgradeStatusOk

`func (o *PlanUpgradeModelList) GetUpgradeStatusOk() (*int32, bool)`

GetUpgradeStatusOk returns a tuple with the UpgradeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStatus

`func (o *PlanUpgradeModelList) SetUpgradeStatus(v int32)`

SetUpgradeStatus sets UpgradeStatus field to given value.

### HasUpgradeStatus

`func (o *PlanUpgradeModelList) HasUpgradeStatus() bool`

HasUpgradeStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


