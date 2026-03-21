# ApTrunkSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Whether the device enable trunk setting. | [optional] 
**Mode** | Pointer to **int32** | Trunk setting mode. Mode should be a value as follows: 0：SRC MAC + DST MAC; 1：DST MAC; 2：SRC MAC. | [optional] 
**SupportTrunkSetting** | Pointer to **bool** | Whether the device supports trunk setting. | [optional] 

## Methods

### NewApTrunkSettingOpenApiVO

`func NewApTrunkSettingOpenApiVO() *ApTrunkSettingOpenApiVO`

NewApTrunkSettingOpenApiVO instantiates a new ApTrunkSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApTrunkSettingOpenApiVOWithDefaults

`func NewApTrunkSettingOpenApiVOWithDefaults() *ApTrunkSettingOpenApiVO`

NewApTrunkSettingOpenApiVOWithDefaults instantiates a new ApTrunkSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *ApTrunkSettingOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ApTrunkSettingOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ApTrunkSettingOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ApTrunkSettingOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetMode

`func (o *ApTrunkSettingOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ApTrunkSettingOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ApTrunkSettingOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ApTrunkSettingOpenApiVO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSupportTrunkSetting

`func (o *ApTrunkSettingOpenApiVO) GetSupportTrunkSetting() bool`

GetSupportTrunkSetting returns the SupportTrunkSetting field if non-nil, zero value otherwise.

### GetSupportTrunkSettingOk

`func (o *ApTrunkSettingOpenApiVO) GetSupportTrunkSettingOk() (*bool, bool)`

GetSupportTrunkSettingOk returns a tuple with the SupportTrunkSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTrunkSetting

`func (o *ApTrunkSettingOpenApiVO) SetSupportTrunkSetting(v bool)`

SetSupportTrunkSetting sets SupportTrunkSetting field to given value.

### HasSupportTrunkSetting

`func (o *ApTrunkSettingOpenApiVO) HasSupportTrunkSetting() bool`

HasSupportTrunkSetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


