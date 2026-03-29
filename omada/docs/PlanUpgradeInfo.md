# PlanUpgradeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVersion** | Pointer to **[]string** | Current version number. | [optional] 
**FromRollback** | Pointer to **bool** | Is it from rollback? | [optional] 
**Id** | Pointer to **string** | ID. | [optional] 
**ModelTypeInfo** | Pointer to [**ModelTypeInfoOpenApiVO**](ModelTypeInfoOpenApiVO.md) |  | [optional] 
**Operator** | Pointer to **string** | operator. | [optional] 
**ScheduledUpgradeTime** | Pointer to **string** | Plan upgrade time. | [optional] 
**SiteNames** | Pointer to **[]string** | Site names where the model is located. | [optional] 
**SiteNum** | Pointer to **int32** | Number of sites where the model is located. | [optional] 
**TargetVersion** | Pointer to **string** | Target upgrade version. | [optional] 

## Methods

### NewPlanUpgradeInfo

`func NewPlanUpgradeInfo() *PlanUpgradeInfo`

NewPlanUpgradeInfo instantiates a new PlanUpgradeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanUpgradeInfoWithDefaults

`func NewPlanUpgradeInfoWithDefaults() *PlanUpgradeInfo`

NewPlanUpgradeInfoWithDefaults instantiates a new PlanUpgradeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVersion

`func (o *PlanUpgradeInfo) GetCurrentVersion() []string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *PlanUpgradeInfo) GetCurrentVersionOk() (*[]string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *PlanUpgradeInfo) SetCurrentVersion(v []string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *PlanUpgradeInfo) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetFromRollback

`func (o *PlanUpgradeInfo) GetFromRollback() bool`

GetFromRollback returns the FromRollback field if non-nil, zero value otherwise.

### GetFromRollbackOk

`func (o *PlanUpgradeInfo) GetFromRollbackOk() (*bool, bool)`

GetFromRollbackOk returns a tuple with the FromRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromRollback

`func (o *PlanUpgradeInfo) SetFromRollback(v bool)`

SetFromRollback sets FromRollback field to given value.

### HasFromRollback

`func (o *PlanUpgradeInfo) HasFromRollback() bool`

HasFromRollback returns a boolean if a field has been set.

### GetId

`func (o *PlanUpgradeInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanUpgradeInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanUpgradeInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlanUpgradeInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModelTypeInfo

`func (o *PlanUpgradeInfo) GetModelTypeInfo() ModelTypeInfoOpenApiVO`

GetModelTypeInfo returns the ModelTypeInfo field if non-nil, zero value otherwise.

### GetModelTypeInfoOk

`func (o *PlanUpgradeInfo) GetModelTypeInfoOk() (*ModelTypeInfoOpenApiVO, bool)`

GetModelTypeInfoOk returns a tuple with the ModelTypeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeInfo

`func (o *PlanUpgradeInfo) SetModelTypeInfo(v ModelTypeInfoOpenApiVO)`

SetModelTypeInfo sets ModelTypeInfo field to given value.

### HasModelTypeInfo

`func (o *PlanUpgradeInfo) HasModelTypeInfo() bool`

HasModelTypeInfo returns a boolean if a field has been set.

### GetOperator

`func (o *PlanUpgradeInfo) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PlanUpgradeInfo) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PlanUpgradeInfo) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *PlanUpgradeInfo) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetScheduledUpgradeTime

`func (o *PlanUpgradeInfo) GetScheduledUpgradeTime() string`

GetScheduledUpgradeTime returns the ScheduledUpgradeTime field if non-nil, zero value otherwise.

### GetScheduledUpgradeTimeOk

`func (o *PlanUpgradeInfo) GetScheduledUpgradeTimeOk() (*string, bool)`

GetScheduledUpgradeTimeOk returns a tuple with the ScheduledUpgradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledUpgradeTime

`func (o *PlanUpgradeInfo) SetScheduledUpgradeTime(v string)`

SetScheduledUpgradeTime sets ScheduledUpgradeTime field to given value.

### HasScheduledUpgradeTime

`func (o *PlanUpgradeInfo) HasScheduledUpgradeTime() bool`

HasScheduledUpgradeTime returns a boolean if a field has been set.

### GetSiteNames

`func (o *PlanUpgradeInfo) GetSiteNames() []string`

GetSiteNames returns the SiteNames field if non-nil, zero value otherwise.

### GetSiteNamesOk

`func (o *PlanUpgradeInfo) GetSiteNamesOk() (*[]string, bool)`

GetSiteNamesOk returns a tuple with the SiteNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteNames

`func (o *PlanUpgradeInfo) SetSiteNames(v []string)`

SetSiteNames sets SiteNames field to given value.

### HasSiteNames

`func (o *PlanUpgradeInfo) HasSiteNames() bool`

HasSiteNames returns a boolean if a field has been set.

### GetSiteNum

`func (o *PlanUpgradeInfo) GetSiteNum() int32`

GetSiteNum returns the SiteNum field if non-nil, zero value otherwise.

### GetSiteNumOk

`func (o *PlanUpgradeInfo) GetSiteNumOk() (*int32, bool)`

GetSiteNumOk returns a tuple with the SiteNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteNum

`func (o *PlanUpgradeInfo) SetSiteNum(v int32)`

SetSiteNum sets SiteNum field to given value.

### HasSiteNum

`func (o *PlanUpgradeInfo) HasSiteNum() bool`

HasSiteNum returns a boolean if a field has been set.

### GetTargetVersion

`func (o *PlanUpgradeInfo) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *PlanUpgradeInfo) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *PlanUpgradeInfo) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *PlanUpgradeInfo) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


