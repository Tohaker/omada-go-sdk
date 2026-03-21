# UpdateApTrunkSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Whether the device enable trunk setting. | [optional] 
**Mode** | Pointer to **int32** | Trunk setting mode. Mode should be a value as follows: 0：SRC MAC + DST MAC; 1：DST MAC; 2：SRC MAC. | [optional] 

## Methods

### NewUpdateApTrunkSettingOpenApiVO

`func NewUpdateApTrunkSettingOpenApiVO() *UpdateApTrunkSettingOpenApiVO`

NewUpdateApTrunkSettingOpenApiVO instantiates a new UpdateApTrunkSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApTrunkSettingOpenApiVOWithDefaults

`func NewUpdateApTrunkSettingOpenApiVOWithDefaults() *UpdateApTrunkSettingOpenApiVO`

NewUpdateApTrunkSettingOpenApiVOWithDefaults instantiates a new UpdateApTrunkSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *UpdateApTrunkSettingOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *UpdateApTrunkSettingOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *UpdateApTrunkSettingOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *UpdateApTrunkSettingOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetMode

`func (o *UpdateApTrunkSettingOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *UpdateApTrunkSettingOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *UpdateApTrunkSettingOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *UpdateApTrunkSettingOpenApiVO) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


