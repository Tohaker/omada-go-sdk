# AntSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntGain** | Pointer to **int32** | antenna gain | [optional] 
**AntMode** | Pointer to **int32** | antenna mode | [optional] 
**AntPattern** | Pointer to **int32** | antenna pattern,0 is built-In and 1 is External. | [optional] 
**CustomGainLimit** | Pointer to **int32** | custom antMode max antenna gain. | [optional] 
**DefaultOmniGain** | Pointer to **int32** | default Omni Antenna Mode gain. | [optional] 

## Methods

### NewAntSettingVO

`func NewAntSettingVO() *AntSettingVO`

NewAntSettingVO instantiates a new AntSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntSettingVOWithDefaults

`func NewAntSettingVOWithDefaults() *AntSettingVO`

NewAntSettingVOWithDefaults instantiates a new AntSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntGain

`func (o *AntSettingVO) GetAntGain() int32`

GetAntGain returns the AntGain field if non-nil, zero value otherwise.

### GetAntGainOk

`func (o *AntSettingVO) GetAntGainOk() (*int32, bool)`

GetAntGainOk returns a tuple with the AntGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntGain

`func (o *AntSettingVO) SetAntGain(v int32)`

SetAntGain sets AntGain field to given value.

### HasAntGain

`func (o *AntSettingVO) HasAntGain() bool`

HasAntGain returns a boolean if a field has been set.

### GetAntMode

`func (o *AntSettingVO) GetAntMode() int32`

GetAntMode returns the AntMode field if non-nil, zero value otherwise.

### GetAntModeOk

`func (o *AntSettingVO) GetAntModeOk() (*int32, bool)`

GetAntModeOk returns a tuple with the AntMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntMode

`func (o *AntSettingVO) SetAntMode(v int32)`

SetAntMode sets AntMode field to given value.

### HasAntMode

`func (o *AntSettingVO) HasAntMode() bool`

HasAntMode returns a boolean if a field has been set.

### GetAntPattern

`func (o *AntSettingVO) GetAntPattern() int32`

GetAntPattern returns the AntPattern field if non-nil, zero value otherwise.

### GetAntPatternOk

`func (o *AntSettingVO) GetAntPatternOk() (*int32, bool)`

GetAntPatternOk returns a tuple with the AntPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntPattern

`func (o *AntSettingVO) SetAntPattern(v int32)`

SetAntPattern sets AntPattern field to given value.

### HasAntPattern

`func (o *AntSettingVO) HasAntPattern() bool`

HasAntPattern returns a boolean if a field has been set.

### GetCustomGainLimit

`func (o *AntSettingVO) GetCustomGainLimit() int32`

GetCustomGainLimit returns the CustomGainLimit field if non-nil, zero value otherwise.

### GetCustomGainLimitOk

`func (o *AntSettingVO) GetCustomGainLimitOk() (*int32, bool)`

GetCustomGainLimitOk returns a tuple with the CustomGainLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomGainLimit

`func (o *AntSettingVO) SetCustomGainLimit(v int32)`

SetCustomGainLimit sets CustomGainLimit field to given value.

### HasCustomGainLimit

`func (o *AntSettingVO) HasCustomGainLimit() bool`

HasCustomGainLimit returns a boolean if a field has been set.

### GetDefaultOmniGain

`func (o *AntSettingVO) GetDefaultOmniGain() int32`

GetDefaultOmniGain returns the DefaultOmniGain field if non-nil, zero value otherwise.

### GetDefaultOmniGainOk

`func (o *AntSettingVO) GetDefaultOmniGainOk() (*int32, bool)`

GetDefaultOmniGainOk returns a tuple with the DefaultOmniGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOmniGain

`func (o *AntSettingVO) SetDefaultOmniGain(v int32)`

SetDefaultOmniGain sets DefaultOmniGain field to given value.

### HasDefaultOmniGain

`func (o *AntSettingVO) HasDefaultOmniGain() bool`

HasDefaultOmniGain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


