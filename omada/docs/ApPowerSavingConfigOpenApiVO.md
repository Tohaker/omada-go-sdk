# ApPowerSavingConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandEnable** | Pointer to **bool** | Power Saving trigger by band config status. True: enable, false: disable. | [optional] 
**Bands** | Pointer to **[]int32** | Select bands list config of trigger by band;It should be a value as follows: 0: 2.4GHz; 1: 5GHz; 2: 5G2Hz; 3: 6GHz. | [optional] 
**EndTimeH** | Pointer to **int32** | End time of trigger by time(unit: hour); It should be within the range of 0–23. | [optional] 
**EndTimeM** | Pointer to **int32** | End time of trigger by time(unit: minute); It should be within the range of 0–59. | [optional] 
**IdleDuration** | Pointer to **int32** | Idle duration config of trigger by band. | [optional] 
**StartTimeH** | Pointer to **int32** | Start time of trigger by time(unit: hour); It should be within the range of 0–23. | [optional] 
**StartTimeM** | Pointer to **int32** | Start time of trigger by time(unit: minute); It should be within the range of 0–59. | [optional] 
**SupportPowerSaving** | Pointer to **bool** | Indicates whether the device supports power saving. True: support, false: unSupport. | [optional] 
**TimeEnable** | Pointer to **bool** | Power Saving trigger by time config status. True: enable, false: disable. | [optional] 

## Methods

### NewApPowerSavingConfigOpenApiVO

`func NewApPowerSavingConfigOpenApiVO() *ApPowerSavingConfigOpenApiVO`

NewApPowerSavingConfigOpenApiVO instantiates a new ApPowerSavingConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApPowerSavingConfigOpenApiVOWithDefaults

`func NewApPowerSavingConfigOpenApiVOWithDefaults() *ApPowerSavingConfigOpenApiVO`

NewApPowerSavingConfigOpenApiVOWithDefaults instantiates a new ApPowerSavingConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandEnable

`func (o *ApPowerSavingConfigOpenApiVO) GetBandEnable() bool`

GetBandEnable returns the BandEnable field if non-nil, zero value otherwise.

### GetBandEnableOk

`func (o *ApPowerSavingConfigOpenApiVO) GetBandEnableOk() (*bool, bool)`

GetBandEnableOk returns a tuple with the BandEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandEnable

`func (o *ApPowerSavingConfigOpenApiVO) SetBandEnable(v bool)`

SetBandEnable sets BandEnable field to given value.

### HasBandEnable

`func (o *ApPowerSavingConfigOpenApiVO) HasBandEnable() bool`

HasBandEnable returns a boolean if a field has been set.

### GetBands

`func (o *ApPowerSavingConfigOpenApiVO) GetBands() []int32`

GetBands returns the Bands field if non-nil, zero value otherwise.

### GetBandsOk

`func (o *ApPowerSavingConfigOpenApiVO) GetBandsOk() (*[]int32, bool)`

GetBandsOk returns a tuple with the Bands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBands

`func (o *ApPowerSavingConfigOpenApiVO) SetBands(v []int32)`

SetBands sets Bands field to given value.

### HasBands

`func (o *ApPowerSavingConfigOpenApiVO) HasBands() bool`

HasBands returns a boolean if a field has been set.

### GetEndTimeH

`func (o *ApPowerSavingConfigOpenApiVO) GetEndTimeH() int32`

GetEndTimeH returns the EndTimeH field if non-nil, zero value otherwise.

### GetEndTimeHOk

`func (o *ApPowerSavingConfigOpenApiVO) GetEndTimeHOk() (*int32, bool)`

GetEndTimeHOk returns a tuple with the EndTimeH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeH

`func (o *ApPowerSavingConfigOpenApiVO) SetEndTimeH(v int32)`

SetEndTimeH sets EndTimeH field to given value.

### HasEndTimeH

`func (o *ApPowerSavingConfigOpenApiVO) HasEndTimeH() bool`

HasEndTimeH returns a boolean if a field has been set.

### GetEndTimeM

`func (o *ApPowerSavingConfigOpenApiVO) GetEndTimeM() int32`

GetEndTimeM returns the EndTimeM field if non-nil, zero value otherwise.

### GetEndTimeMOk

`func (o *ApPowerSavingConfigOpenApiVO) GetEndTimeMOk() (*int32, bool)`

GetEndTimeMOk returns a tuple with the EndTimeM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeM

`func (o *ApPowerSavingConfigOpenApiVO) SetEndTimeM(v int32)`

SetEndTimeM sets EndTimeM field to given value.

### HasEndTimeM

`func (o *ApPowerSavingConfigOpenApiVO) HasEndTimeM() bool`

HasEndTimeM returns a boolean if a field has been set.

### GetIdleDuration

`func (o *ApPowerSavingConfigOpenApiVO) GetIdleDuration() int32`

GetIdleDuration returns the IdleDuration field if non-nil, zero value otherwise.

### GetIdleDurationOk

`func (o *ApPowerSavingConfigOpenApiVO) GetIdleDurationOk() (*int32, bool)`

GetIdleDurationOk returns a tuple with the IdleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleDuration

`func (o *ApPowerSavingConfigOpenApiVO) SetIdleDuration(v int32)`

SetIdleDuration sets IdleDuration field to given value.

### HasIdleDuration

`func (o *ApPowerSavingConfigOpenApiVO) HasIdleDuration() bool`

HasIdleDuration returns a boolean if a field has been set.

### GetStartTimeH

`func (o *ApPowerSavingConfigOpenApiVO) GetStartTimeH() int32`

GetStartTimeH returns the StartTimeH field if non-nil, zero value otherwise.

### GetStartTimeHOk

`func (o *ApPowerSavingConfigOpenApiVO) GetStartTimeHOk() (*int32, bool)`

GetStartTimeHOk returns a tuple with the StartTimeH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeH

`func (o *ApPowerSavingConfigOpenApiVO) SetStartTimeH(v int32)`

SetStartTimeH sets StartTimeH field to given value.

### HasStartTimeH

`func (o *ApPowerSavingConfigOpenApiVO) HasStartTimeH() bool`

HasStartTimeH returns a boolean if a field has been set.

### GetStartTimeM

`func (o *ApPowerSavingConfigOpenApiVO) GetStartTimeM() int32`

GetStartTimeM returns the StartTimeM field if non-nil, zero value otherwise.

### GetStartTimeMOk

`func (o *ApPowerSavingConfigOpenApiVO) GetStartTimeMOk() (*int32, bool)`

GetStartTimeMOk returns a tuple with the StartTimeM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeM

`func (o *ApPowerSavingConfigOpenApiVO) SetStartTimeM(v int32)`

SetStartTimeM sets StartTimeM field to given value.

### HasStartTimeM

`func (o *ApPowerSavingConfigOpenApiVO) HasStartTimeM() bool`

HasStartTimeM returns a boolean if a field has been set.

### GetSupportPowerSaving

`func (o *ApPowerSavingConfigOpenApiVO) GetSupportPowerSaving() bool`

GetSupportPowerSaving returns the SupportPowerSaving field if non-nil, zero value otherwise.

### GetSupportPowerSavingOk

`func (o *ApPowerSavingConfigOpenApiVO) GetSupportPowerSavingOk() (*bool, bool)`

GetSupportPowerSavingOk returns a tuple with the SupportPowerSaving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPowerSaving

`func (o *ApPowerSavingConfigOpenApiVO) SetSupportPowerSaving(v bool)`

SetSupportPowerSaving sets SupportPowerSaving field to given value.

### HasSupportPowerSaving

`func (o *ApPowerSavingConfigOpenApiVO) HasSupportPowerSaving() bool`

HasSupportPowerSaving returns a boolean if a field has been set.

### GetTimeEnable

`func (o *ApPowerSavingConfigOpenApiVO) GetTimeEnable() bool`

GetTimeEnable returns the TimeEnable field if non-nil, zero value otherwise.

### GetTimeEnableOk

`func (o *ApPowerSavingConfigOpenApiVO) GetTimeEnableOk() (*bool, bool)`

GetTimeEnableOk returns a tuple with the TimeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEnable

`func (o *ApPowerSavingConfigOpenApiVO) SetTimeEnable(v bool)`

SetTimeEnable sets TimeEnable field to given value.

### HasTimeEnable

`func (o *ApPowerSavingConfigOpenApiVO) HasTimeEnable() bool`

HasTimeEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


