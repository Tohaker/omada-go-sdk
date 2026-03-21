# AntSwitchRadioConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntGain** | Pointer to **int32** | antenna Gain | [optional] 
**AntMode** | Pointer to **int32** | antenna mode in radio setting. 0:Auto, 1: Built-In, 2: Omni | [optional] 

## Methods

### NewAntSwitchRadioConfig

`func NewAntSwitchRadioConfig() *AntSwitchRadioConfig`

NewAntSwitchRadioConfig instantiates a new AntSwitchRadioConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntSwitchRadioConfigWithDefaults

`func NewAntSwitchRadioConfigWithDefaults() *AntSwitchRadioConfig`

NewAntSwitchRadioConfigWithDefaults instantiates a new AntSwitchRadioConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntGain

`func (o *AntSwitchRadioConfig) GetAntGain() int32`

GetAntGain returns the AntGain field if non-nil, zero value otherwise.

### GetAntGainOk

`func (o *AntSwitchRadioConfig) GetAntGainOk() (*int32, bool)`

GetAntGainOk returns a tuple with the AntGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntGain

`func (o *AntSwitchRadioConfig) SetAntGain(v int32)`

SetAntGain sets AntGain field to given value.

### HasAntGain

`func (o *AntSwitchRadioConfig) HasAntGain() bool`

HasAntGain returns a boolean if a field has been set.

### GetAntMode

`func (o *AntSwitchRadioConfig) GetAntMode() int32`

GetAntMode returns the AntMode field if non-nil, zero value otherwise.

### GetAntModeOk

`func (o *AntSwitchRadioConfig) GetAntModeOk() (*int32, bool)`

GetAntModeOk returns a tuple with the AntMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntMode

`func (o *AntSwitchRadioConfig) SetAntMode(v int32)`

SetAntMode sets AntMode field to given value.

### HasAntMode

`func (o *AntSwitchRadioConfig) HasAntMode() bool`

HasAntMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


