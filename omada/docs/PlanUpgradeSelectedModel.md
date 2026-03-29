# PlanUpgradeSelectedModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVersion** | **[]string** | Selected model version list, software version, such as \&quot;2.5.0 Build 20190118 Rel. 64821\&quot;. | 
**ModelTypeInfo** | [**ModelTypeInfoOpenApiVO**](ModelTypeInfoOpenApiVO.md) |  | 
**TargetVersion** | **string** | Target version. | 

## Methods

### NewPlanUpgradeSelectedModel

`func NewPlanUpgradeSelectedModel(currentVersion []string, modelTypeInfo ModelTypeInfoOpenApiVO, targetVersion string, ) *PlanUpgradeSelectedModel`

NewPlanUpgradeSelectedModel instantiates a new PlanUpgradeSelectedModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanUpgradeSelectedModelWithDefaults

`func NewPlanUpgradeSelectedModelWithDefaults() *PlanUpgradeSelectedModel`

NewPlanUpgradeSelectedModelWithDefaults instantiates a new PlanUpgradeSelectedModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVersion

`func (o *PlanUpgradeSelectedModel) GetCurrentVersion() []string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *PlanUpgradeSelectedModel) GetCurrentVersionOk() (*[]string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *PlanUpgradeSelectedModel) SetCurrentVersion(v []string)`

SetCurrentVersion sets CurrentVersion field to given value.


### GetModelTypeInfo

`func (o *PlanUpgradeSelectedModel) GetModelTypeInfo() ModelTypeInfoOpenApiVO`

GetModelTypeInfo returns the ModelTypeInfo field if non-nil, zero value otherwise.

### GetModelTypeInfoOk

`func (o *PlanUpgradeSelectedModel) GetModelTypeInfoOk() (*ModelTypeInfoOpenApiVO, bool)`

GetModelTypeInfoOk returns a tuple with the ModelTypeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeInfo

`func (o *PlanUpgradeSelectedModel) SetModelTypeInfo(v ModelTypeInfoOpenApiVO)`

SetModelTypeInfo sets ModelTypeInfo field to given value.


### GetTargetVersion

`func (o *PlanUpgradeSelectedModel) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *PlanUpgradeSelectedModel) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *PlanUpgradeSelectedModel) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


