# UpdateApPowerSavingConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandEnable** | **bool** | Power Saving trigger by band config status. True: enable, false: disable. | 
**Bands** | Pointer to **[]int32** | Select bands list config of trigger by band;It should be a value as follows: 0: 2.4GHz; 1: 5GHz; 2: 5G2Hz; 3: 6GHz;This field is required when Parameter [bandEnable] is true;Please note that the band filled in must be supported by the device, otherwise the modified configuration will not take effect properly | [optional] 
**EndTimeH** | Pointer to **int32** | End time of trigger by time(unit: hour); It should be within the range of 0–23; This field is required when Parameter [timeEnable] is true. | [optional] 
**EndTimeM** | Pointer to **int32** | End time of trigger by time(unit: minute); It should be within the range of 0–59; This field is required when Parameter [timeEnable] is true. | [optional] 
**IdleDuration** | Pointer to **int32** | Idle duration config of trigger by band(unit: minute); It should be within the range of 60–1440; This field is required when Parameter [bandEnable] is true. | [optional] 
**StartTimeH** | Pointer to **int32** | Start time of trigger by time(unit: hour); It should be within the range of 0–23; This field is required when Parameter [timeEnable] is true. | [optional] 
**StartTimeM** | Pointer to **int32** | Start time of trigger by time(unit: minute); It should be within the range of 0–59; This field is required when Parameter [timeEnable] is true. | [optional] 
**TimeEnable** | **bool** | Power Saving trigger by time config status. True: enable, false: disable. | 

## Methods

### NewUpdateApPowerSavingConfigOpenApiVO

`func NewUpdateApPowerSavingConfigOpenApiVO(bandEnable bool, timeEnable bool, ) *UpdateApPowerSavingConfigOpenApiVO`

NewUpdateApPowerSavingConfigOpenApiVO instantiates a new UpdateApPowerSavingConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApPowerSavingConfigOpenApiVOWithDefaults

`func NewUpdateApPowerSavingConfigOpenApiVOWithDefaults() *UpdateApPowerSavingConfigOpenApiVO`

NewUpdateApPowerSavingConfigOpenApiVOWithDefaults instantiates a new UpdateApPowerSavingConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandEnable

`func (o *UpdateApPowerSavingConfigOpenApiVO) GetBandEnable() bool`

GetBandEnable returns the BandEnable field if non-nil, zero value otherwise.

### GetBandEnableOk

`func (o *UpdateApPowerSavingConfigOpenApiVO) GetBandEnableOk() (*bool, bool)`

GetBandEnableOk returns a tuple with the BandEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandEnable

`func (o *UpdateApPowerSavingConfigOpenApiVO) SetBandEnable(v bool)`

SetBandEnable sets BandEnable field to given value.


### GetBands

`func (o *UpdateApPowerSavingConfigOpenApiVO) GetBands() []int32`

GetBands returns the Bands field if non-nil, zero value otherwise.

### GetBandsOk

`func (o *UpdateApPowerSavingConfigOpenApiVO) GetBandsOk() (*[]int32, bool)`

GetBandsOk returns a tuple with the Bands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBands

`func (o *UpdateApPowerSavingConfigOpenApiVO) SetBands(v []int32)`

SetBands sets Bands field to given value.

### HasBands

`func (o *UpdateApPowerSavingConfigOpenApiVO) HasBands() bool`

HasBands returns a boolean if a field has been set.

### GetEndTimeH

`func (o *UpdateApPowerSavingConfigOpenApiVO) GetEndTimeH() int32`

GetEndTimeH returns the EndTimeH field if non-nil, zero value otherwise.

### GetEndTimeHOk

`func (o *UpdateApPowerSavingConfigOpenApiVO) GetEndTimeHOk() (*int32, bool)`

GetEndTimeHOk returns a tuple with the EndTimeH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeH

`func (o *UpdateApPowerSavingConfigOpenApiVO) SetEndTimeH(v int32)`

SetEndTimeH sets EndTimeH field to given value.

### HasEndTimeH

`func (o *UpdateApPowerSavingConfigOpenApiVO) HasEndTimeH() bool`

HasEndTimeH returns a boolean if a field has been set.

### GetEndTimeM

`func (o *UpdateApPowerSavingConfigOpenApiVO) GetEndTimeM() int32`

GetEndTimeM returns the EndTimeM field if non-nil, zero value otherwise.

### GetEndTimeMOk

`func (o *UpdateApPowerSavingConfigOpenApiVO) GetEndTimeMOk() (*int32, bool)`

GetEndTimeMOk returns a tuple with the EndTimeM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeM

`func (o *UpdateApPowerSavingConfigOpenApiVO) SetEndTimeM(v int32)`

SetEndTimeM sets EndTimeM field to given value.

### HasEndTimeM

`func (o *UpdateApPowerSavingConfigOpenApiVO) HasEndTimeM() bool`

HasEndTimeM returns a boolean if a field has been set.

### GetIdleDuration

`func (o *UpdateApPowerSavingConfigOpenApiVO) GetIdleDuration() int32`

GetIdleDuration returns the IdleDuration field if non-nil, zero value otherwise.

### GetIdleDurationOk

`func (o *UpdateApPowerSavingConfigOpenApiVO) GetIdleDurationOk() (*int32, bool)`

GetIdleDurationOk returns a tuple with the IdleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleDuration

`func (o *UpdateApPowerSavingConfigOpenApiVO) SetIdleDuration(v int32)`

SetIdleDuration sets IdleDuration field to given value.

### HasIdleDuration

`func (o *UpdateApPowerSavingConfigOpenApiVO) HasIdleDuration() bool`

HasIdleDuration returns a boolean if a field has been set.

### GetStartTimeH

`func (o *UpdateApPowerSavingConfigOpenApiVO) GetStartTimeH() int32`

GetStartTimeH returns the StartTimeH field if non-nil, zero value otherwise.

### GetStartTimeHOk

`func (o *UpdateApPowerSavingConfigOpenApiVO) GetStartTimeHOk() (*int32, bool)`

GetStartTimeHOk returns a tuple with the StartTimeH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeH

`func (o *UpdateApPowerSavingConfigOpenApiVO) SetStartTimeH(v int32)`

SetStartTimeH sets StartTimeH field to given value.

### HasStartTimeH

`func (o *UpdateApPowerSavingConfigOpenApiVO) HasStartTimeH() bool`

HasStartTimeH returns a boolean if a field has been set.

### GetStartTimeM

`func (o *UpdateApPowerSavingConfigOpenApiVO) GetStartTimeM() int32`

GetStartTimeM returns the StartTimeM field if non-nil, zero value otherwise.

### GetStartTimeMOk

`func (o *UpdateApPowerSavingConfigOpenApiVO) GetStartTimeMOk() (*int32, bool)`

GetStartTimeMOk returns a tuple with the StartTimeM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeM

`func (o *UpdateApPowerSavingConfigOpenApiVO) SetStartTimeM(v int32)`

SetStartTimeM sets StartTimeM field to given value.

### HasStartTimeM

`func (o *UpdateApPowerSavingConfigOpenApiVO) HasStartTimeM() bool`

HasStartTimeM returns a boolean if a field has been set.

### GetTimeEnable

`func (o *UpdateApPowerSavingConfigOpenApiVO) GetTimeEnable() bool`

GetTimeEnable returns the TimeEnable field if non-nil, zero value otherwise.

### GetTimeEnableOk

`func (o *UpdateApPowerSavingConfigOpenApiVO) GetTimeEnableOk() (*bool, bool)`

GetTimeEnableOk returns a tuple with the TimeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEnable

`func (o *UpdateApPowerSavingConfigOpenApiVO) SetTimeEnable(v bool)`

SetTimeEnable sets TimeEnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


