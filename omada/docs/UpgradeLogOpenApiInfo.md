# UpgradeLogOpenApiInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVersion** | Pointer to **[]string** | Model firmware version list before upgrade ,such as \&quot;[2.5.0 Build 20190118 Rel. 64821, 2.4.8 Build 20190118 Rel. 64821]\&quot;. This field is the same as \&quot;PREVIOUS VERSION\&quot; in Controller Upgrade Logs | [optional] 
**Id** | Pointer to **string** | Upgrade log ID. | [optional] 
**ModelTypeInfo** | Pointer to [**ModelTypeInfoOpenApiVO**](ModelTypeInfoOpenApiVO.md) |  | [optional] 
**Operator** | Pointer to **string** | Operator. | [optional] 
**Rollbacked** | Pointer to **bool** | Whether the log is rolled back or not. The logs that are rolled back cannot be rolled back repeatedly. | [optional] 
**SiteNames** | Pointer to **[]string** | The siteName lists where the model belongs. | [optional] 
**SiteNum** | Pointer to **int32** | Number of sites where the model belongs. | [optional] 
**TargetVersion** | Pointer to **string** | Model firmware version after device upgrade, such as \&quot;2.5.1 Build 20190118 Rel. 64821\&quot;. This field is the same as \&quot;CURRENT VERSION\&quot; in Controller Upgrade Logs | [optional] 
**UpgradeTime** | Pointer to **string** | The time when the upgrade is complete. | [optional] 

## Methods

### NewUpgradeLogOpenApiInfo

`func NewUpgradeLogOpenApiInfo() *UpgradeLogOpenApiInfo`

NewUpgradeLogOpenApiInfo instantiates a new UpgradeLogOpenApiInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeLogOpenApiInfoWithDefaults

`func NewUpgradeLogOpenApiInfoWithDefaults() *UpgradeLogOpenApiInfo`

NewUpgradeLogOpenApiInfoWithDefaults instantiates a new UpgradeLogOpenApiInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVersion

`func (o *UpgradeLogOpenApiInfo) GetCurrentVersion() []string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *UpgradeLogOpenApiInfo) GetCurrentVersionOk() (*[]string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *UpgradeLogOpenApiInfo) SetCurrentVersion(v []string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *UpgradeLogOpenApiInfo) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetId

`func (o *UpgradeLogOpenApiInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpgradeLogOpenApiInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpgradeLogOpenApiInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpgradeLogOpenApiInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModelTypeInfo

`func (o *UpgradeLogOpenApiInfo) GetModelTypeInfo() ModelTypeInfoOpenApiVO`

GetModelTypeInfo returns the ModelTypeInfo field if non-nil, zero value otherwise.

### GetModelTypeInfoOk

`func (o *UpgradeLogOpenApiInfo) GetModelTypeInfoOk() (*ModelTypeInfoOpenApiVO, bool)`

GetModelTypeInfoOk returns a tuple with the ModelTypeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeInfo

`func (o *UpgradeLogOpenApiInfo) SetModelTypeInfo(v ModelTypeInfoOpenApiVO)`

SetModelTypeInfo sets ModelTypeInfo field to given value.

### HasModelTypeInfo

`func (o *UpgradeLogOpenApiInfo) HasModelTypeInfo() bool`

HasModelTypeInfo returns a boolean if a field has been set.

### GetOperator

`func (o *UpgradeLogOpenApiInfo) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *UpgradeLogOpenApiInfo) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *UpgradeLogOpenApiInfo) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *UpgradeLogOpenApiInfo) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetRollbacked

`func (o *UpgradeLogOpenApiInfo) GetRollbacked() bool`

GetRollbacked returns the Rollbacked field if non-nil, zero value otherwise.

### GetRollbackedOk

`func (o *UpgradeLogOpenApiInfo) GetRollbackedOk() (*bool, bool)`

GetRollbackedOk returns a tuple with the Rollbacked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbacked

`func (o *UpgradeLogOpenApiInfo) SetRollbacked(v bool)`

SetRollbacked sets Rollbacked field to given value.

### HasRollbacked

`func (o *UpgradeLogOpenApiInfo) HasRollbacked() bool`

HasRollbacked returns a boolean if a field has been set.

### GetSiteNames

`func (o *UpgradeLogOpenApiInfo) GetSiteNames() []string`

GetSiteNames returns the SiteNames field if non-nil, zero value otherwise.

### GetSiteNamesOk

`func (o *UpgradeLogOpenApiInfo) GetSiteNamesOk() (*[]string, bool)`

GetSiteNamesOk returns a tuple with the SiteNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteNames

`func (o *UpgradeLogOpenApiInfo) SetSiteNames(v []string)`

SetSiteNames sets SiteNames field to given value.

### HasSiteNames

`func (o *UpgradeLogOpenApiInfo) HasSiteNames() bool`

HasSiteNames returns a boolean if a field has been set.

### GetSiteNum

`func (o *UpgradeLogOpenApiInfo) GetSiteNum() int32`

GetSiteNum returns the SiteNum field if non-nil, zero value otherwise.

### GetSiteNumOk

`func (o *UpgradeLogOpenApiInfo) GetSiteNumOk() (*int32, bool)`

GetSiteNumOk returns a tuple with the SiteNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteNum

`func (o *UpgradeLogOpenApiInfo) SetSiteNum(v int32)`

SetSiteNum sets SiteNum field to given value.

### HasSiteNum

`func (o *UpgradeLogOpenApiInfo) HasSiteNum() bool`

HasSiteNum returns a boolean if a field has been set.

### GetTargetVersion

`func (o *UpgradeLogOpenApiInfo) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *UpgradeLogOpenApiInfo) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *UpgradeLogOpenApiInfo) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *UpgradeLogOpenApiInfo) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.

### GetUpgradeTime

`func (o *UpgradeLogOpenApiInfo) GetUpgradeTime() string`

GetUpgradeTime returns the UpgradeTime field if non-nil, zero value otherwise.

### GetUpgradeTimeOk

`func (o *UpgradeLogOpenApiInfo) GetUpgradeTimeOk() (*string, bool)`

GetUpgradeTimeOk returns a tuple with the UpgradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeTime

`func (o *UpgradeLogOpenApiInfo) SetUpgradeTime(v string)`

SetUpgradeTime sets UpgradeTime field to given value.

### HasUpgradeTime

`func (o *UpgradeLogOpenApiInfo) HasUpgradeTime() bool`

HasUpgradeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


