# ApMgtSsidVlanSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomConfig** | Pointer to [**ApMgtSsidVlanCustomSettingOpenApiVO**](ApMgtSsidVlanCustomSettingOpenApiVO.md) |  | [optional] 
**Mode** | **int32** | should be a value as follows: 0:Default; 1:Custom. | 

## Methods

### NewApMgtSsidVlanSettingOpenApiVO

`func NewApMgtSsidVlanSettingOpenApiVO(mode int32, ) *ApMgtSsidVlanSettingOpenApiVO`

NewApMgtSsidVlanSettingOpenApiVO instantiates a new ApMgtSsidVlanSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApMgtSsidVlanSettingOpenApiVOWithDefaults

`func NewApMgtSsidVlanSettingOpenApiVOWithDefaults() *ApMgtSsidVlanSettingOpenApiVO`

NewApMgtSsidVlanSettingOpenApiVOWithDefaults instantiates a new ApMgtSsidVlanSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomConfig

`func (o *ApMgtSsidVlanSettingOpenApiVO) GetCustomConfig() ApMgtSsidVlanCustomSettingOpenApiVO`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *ApMgtSsidVlanSettingOpenApiVO) GetCustomConfigOk() (*ApMgtSsidVlanCustomSettingOpenApiVO, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *ApMgtSsidVlanSettingOpenApiVO) SetCustomConfig(v ApMgtSsidVlanCustomSettingOpenApiVO)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *ApMgtSsidVlanSettingOpenApiVO) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### GetMode

`func (o *ApMgtSsidVlanSettingOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ApMgtSsidVlanSettingOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ApMgtSsidVlanSettingOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


