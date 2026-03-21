# AntSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntGain** | Pointer to **int32** | antenna gain | [optional] 
**AntMode** | Pointer to **int32** | antenna mode | [optional] 
**AntPattern** | Pointer to **int32** | antenna pattern,0 is directional and 1 is omnidirectional. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


