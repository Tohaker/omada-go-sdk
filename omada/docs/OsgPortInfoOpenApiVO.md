# OsgPortInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreOsgModel** | Pointer to **int32** | preconfigured gateway model | [optional] 
**RealOsgModel** | Pointer to **string** | real gateway model | [optional] 
**TargetModels** | Pointer to **[]int32** | target gateway models when gateway change | [optional] 
**WanLanPortSettings** | Pointer to [**[]WanLanPortSettingOpenApiVO**](WanLanPortSettingOpenApiVO.md) |  | [optional] 
**WanPortNum** | Pointer to **int32** | custom wan port num | [optional] 

## Methods

### NewOsgPortInfoOpenApiVO

`func NewOsgPortInfoOpenApiVO() *OsgPortInfoOpenApiVO`

NewOsgPortInfoOpenApiVO instantiates a new OsgPortInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgPortInfoOpenApiVOWithDefaults

`func NewOsgPortInfoOpenApiVOWithDefaults() *OsgPortInfoOpenApiVO`

NewOsgPortInfoOpenApiVOWithDefaults instantiates a new OsgPortInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreOsgModel

`func (o *OsgPortInfoOpenApiVO) GetPreOsgModel() int32`

GetPreOsgModel returns the PreOsgModel field if non-nil, zero value otherwise.

### GetPreOsgModelOk

`func (o *OsgPortInfoOpenApiVO) GetPreOsgModelOk() (*int32, bool)`

GetPreOsgModelOk returns a tuple with the PreOsgModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreOsgModel

`func (o *OsgPortInfoOpenApiVO) SetPreOsgModel(v int32)`

SetPreOsgModel sets PreOsgModel field to given value.

### HasPreOsgModel

`func (o *OsgPortInfoOpenApiVO) HasPreOsgModel() bool`

HasPreOsgModel returns a boolean if a field has been set.

### GetRealOsgModel

`func (o *OsgPortInfoOpenApiVO) GetRealOsgModel() string`

GetRealOsgModel returns the RealOsgModel field if non-nil, zero value otherwise.

### GetRealOsgModelOk

`func (o *OsgPortInfoOpenApiVO) GetRealOsgModelOk() (*string, bool)`

GetRealOsgModelOk returns a tuple with the RealOsgModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealOsgModel

`func (o *OsgPortInfoOpenApiVO) SetRealOsgModel(v string)`

SetRealOsgModel sets RealOsgModel field to given value.

### HasRealOsgModel

`func (o *OsgPortInfoOpenApiVO) HasRealOsgModel() bool`

HasRealOsgModel returns a boolean if a field has been set.

### GetTargetModels

`func (o *OsgPortInfoOpenApiVO) GetTargetModels() []int32`

GetTargetModels returns the TargetModels field if non-nil, zero value otherwise.

### GetTargetModelsOk

`func (o *OsgPortInfoOpenApiVO) GetTargetModelsOk() (*[]int32, bool)`

GetTargetModelsOk returns a tuple with the TargetModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetModels

`func (o *OsgPortInfoOpenApiVO) SetTargetModels(v []int32)`

SetTargetModels sets TargetModels field to given value.

### HasTargetModels

`func (o *OsgPortInfoOpenApiVO) HasTargetModels() bool`

HasTargetModels returns a boolean if a field has been set.

### GetWanLanPortSettings

`func (o *OsgPortInfoOpenApiVO) GetWanLanPortSettings() []WanLanPortSettingOpenApiVO`

GetWanLanPortSettings returns the WanLanPortSettings field if non-nil, zero value otherwise.

### GetWanLanPortSettingsOk

`func (o *OsgPortInfoOpenApiVO) GetWanLanPortSettingsOk() (*[]WanLanPortSettingOpenApiVO, bool)`

GetWanLanPortSettingsOk returns a tuple with the WanLanPortSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanLanPortSettings

`func (o *OsgPortInfoOpenApiVO) SetWanLanPortSettings(v []WanLanPortSettingOpenApiVO)`

SetWanLanPortSettings sets WanLanPortSettings field to given value.

### HasWanLanPortSettings

`func (o *OsgPortInfoOpenApiVO) HasWanLanPortSettings() bool`

HasWanLanPortSettings returns a boolean if a field has been set.

### GetWanPortNum

`func (o *OsgPortInfoOpenApiVO) GetWanPortNum() int32`

GetWanPortNum returns the WanPortNum field if non-nil, zero value otherwise.

### GetWanPortNumOk

`func (o *OsgPortInfoOpenApiVO) GetWanPortNumOk() (*int32, bool)`

GetWanPortNumOk returns a tuple with the WanPortNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortNum

`func (o *OsgPortInfoOpenApiVO) SetWanPortNum(v int32)`

SetWanPortNum sets WanPortNum field to given value.

### HasWanPortNum

`func (o *OsgPortInfoOpenApiVO) HasWanPortNum() bool`

HasWanPortNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


