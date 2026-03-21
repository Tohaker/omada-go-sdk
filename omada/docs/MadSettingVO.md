# MadSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Enable | [optional] 
**Mode** | Pointer to **int32** | Mode should be 0 | [optional] 
**SelectPorts** | Pointer to **[]string** | Select Ports | [optional] 

## Methods

### NewMadSettingVO

`func NewMadSettingVO() *MadSettingVO`

NewMadSettingVO instantiates a new MadSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMadSettingVOWithDefaults

`func NewMadSettingVOWithDefaults() *MadSettingVO`

NewMadSettingVOWithDefaults instantiates a new MadSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *MadSettingVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *MadSettingVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *MadSettingVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *MadSettingVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetMode

`func (o *MadSettingVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *MadSettingVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *MadSettingVO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *MadSettingVO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSelectPorts

`func (o *MadSettingVO) GetSelectPorts() []string`

GetSelectPorts returns the SelectPorts field if non-nil, zero value otherwise.

### GetSelectPortsOk

`func (o *MadSettingVO) GetSelectPortsOk() (*[]string, bool)`

GetSelectPortsOk returns a tuple with the SelectPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectPorts

`func (o *MadSettingVO) SetSelectPorts(v []string)`

SetSelectPorts sets SelectPorts field to given value.

### HasSelectPorts

`func (o *MadSettingVO) HasSelectPorts() bool`

HasSelectPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


