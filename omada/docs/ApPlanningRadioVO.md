# ApPlanningRadioVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrCode** | Pointer to **int32** | Error code returned by the algorithm. 0: OK. -34820: The configured bandwidth has no optional channel. -34821: The lowest power supported by this device is out of the configured Power Range. -34822: The highest power supported by this device is out of the configured Power Range. -34815: Failed to optimize device because of no scan result. -34816: Failed to apply deploy config because of this device is not connected. -34823: Cannot optimize the channels because you have excluded them. | [optional] 
**OrigBand** | Pointer to **int32** | The original switch status of band. 0: off, 1: on. | [optional] 
**OrigChan** | Pointer to **string** | The original channel. | [optional] 
**OrigChanWidth** | Pointer to **string** | The original channel width. | [optional] 
**OrigPower** | Pointer to **int32** | The original power value. | [optional] 
**RcmdBand** | Pointer to **int32** | Band switch status recommended by the algorithm. 0: off, 1: on. | [optional] 
**RcmdChan** | Pointer to **string** | Channel recommended by the algorithm. | [optional] 
**RcmdChanWidth** | Pointer to **string** | Channel width recommended by the algorithm. | [optional] 
**RcmdPower** | Pointer to **int32** | Power value recommended by the algorithm. | [optional] 

## Methods

### NewApPlanningRadioVO

`func NewApPlanningRadioVO() *ApPlanningRadioVO`

NewApPlanningRadioVO instantiates a new ApPlanningRadioVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApPlanningRadioVOWithDefaults

`func NewApPlanningRadioVOWithDefaults() *ApPlanningRadioVO`

NewApPlanningRadioVOWithDefaults instantiates a new ApPlanningRadioVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrCode

`func (o *ApPlanningRadioVO) GetErrCode() int32`

GetErrCode returns the ErrCode field if non-nil, zero value otherwise.

### GetErrCodeOk

`func (o *ApPlanningRadioVO) GetErrCodeOk() (*int32, bool)`

GetErrCodeOk returns a tuple with the ErrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrCode

`func (o *ApPlanningRadioVO) SetErrCode(v int32)`

SetErrCode sets ErrCode field to given value.

### HasErrCode

`func (o *ApPlanningRadioVO) HasErrCode() bool`

HasErrCode returns a boolean if a field has been set.

### GetOrigBand

`func (o *ApPlanningRadioVO) GetOrigBand() int32`

GetOrigBand returns the OrigBand field if non-nil, zero value otherwise.

### GetOrigBandOk

`func (o *ApPlanningRadioVO) GetOrigBandOk() (*int32, bool)`

GetOrigBandOk returns a tuple with the OrigBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigBand

`func (o *ApPlanningRadioVO) SetOrigBand(v int32)`

SetOrigBand sets OrigBand field to given value.

### HasOrigBand

`func (o *ApPlanningRadioVO) HasOrigBand() bool`

HasOrigBand returns a boolean if a field has been set.

### GetOrigChan

`func (o *ApPlanningRadioVO) GetOrigChan() string`

GetOrigChan returns the OrigChan field if non-nil, zero value otherwise.

### GetOrigChanOk

`func (o *ApPlanningRadioVO) GetOrigChanOk() (*string, bool)`

GetOrigChanOk returns a tuple with the OrigChan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigChan

`func (o *ApPlanningRadioVO) SetOrigChan(v string)`

SetOrigChan sets OrigChan field to given value.

### HasOrigChan

`func (o *ApPlanningRadioVO) HasOrigChan() bool`

HasOrigChan returns a boolean if a field has been set.

### GetOrigChanWidth

`func (o *ApPlanningRadioVO) GetOrigChanWidth() string`

GetOrigChanWidth returns the OrigChanWidth field if non-nil, zero value otherwise.

### GetOrigChanWidthOk

`func (o *ApPlanningRadioVO) GetOrigChanWidthOk() (*string, bool)`

GetOrigChanWidthOk returns a tuple with the OrigChanWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigChanWidth

`func (o *ApPlanningRadioVO) SetOrigChanWidth(v string)`

SetOrigChanWidth sets OrigChanWidth field to given value.

### HasOrigChanWidth

`func (o *ApPlanningRadioVO) HasOrigChanWidth() bool`

HasOrigChanWidth returns a boolean if a field has been set.

### GetOrigPower

`func (o *ApPlanningRadioVO) GetOrigPower() int32`

GetOrigPower returns the OrigPower field if non-nil, zero value otherwise.

### GetOrigPowerOk

`func (o *ApPlanningRadioVO) GetOrigPowerOk() (*int32, bool)`

GetOrigPowerOk returns a tuple with the OrigPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigPower

`func (o *ApPlanningRadioVO) SetOrigPower(v int32)`

SetOrigPower sets OrigPower field to given value.

### HasOrigPower

`func (o *ApPlanningRadioVO) HasOrigPower() bool`

HasOrigPower returns a boolean if a field has been set.

### GetRcmdBand

`func (o *ApPlanningRadioVO) GetRcmdBand() int32`

GetRcmdBand returns the RcmdBand field if non-nil, zero value otherwise.

### GetRcmdBandOk

`func (o *ApPlanningRadioVO) GetRcmdBandOk() (*int32, bool)`

GetRcmdBandOk returns a tuple with the RcmdBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcmdBand

`func (o *ApPlanningRadioVO) SetRcmdBand(v int32)`

SetRcmdBand sets RcmdBand field to given value.

### HasRcmdBand

`func (o *ApPlanningRadioVO) HasRcmdBand() bool`

HasRcmdBand returns a boolean if a field has been set.

### GetRcmdChan

`func (o *ApPlanningRadioVO) GetRcmdChan() string`

GetRcmdChan returns the RcmdChan field if non-nil, zero value otherwise.

### GetRcmdChanOk

`func (o *ApPlanningRadioVO) GetRcmdChanOk() (*string, bool)`

GetRcmdChanOk returns a tuple with the RcmdChan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcmdChan

`func (o *ApPlanningRadioVO) SetRcmdChan(v string)`

SetRcmdChan sets RcmdChan field to given value.

### HasRcmdChan

`func (o *ApPlanningRadioVO) HasRcmdChan() bool`

HasRcmdChan returns a boolean if a field has been set.

### GetRcmdChanWidth

`func (o *ApPlanningRadioVO) GetRcmdChanWidth() string`

GetRcmdChanWidth returns the RcmdChanWidth field if non-nil, zero value otherwise.

### GetRcmdChanWidthOk

`func (o *ApPlanningRadioVO) GetRcmdChanWidthOk() (*string, bool)`

GetRcmdChanWidthOk returns a tuple with the RcmdChanWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcmdChanWidth

`func (o *ApPlanningRadioVO) SetRcmdChanWidth(v string)`

SetRcmdChanWidth sets RcmdChanWidth field to given value.

### HasRcmdChanWidth

`func (o *ApPlanningRadioVO) HasRcmdChanWidth() bool`

HasRcmdChanWidth returns a boolean if a field has been set.

### GetRcmdPower

`func (o *ApPlanningRadioVO) GetRcmdPower() int32`

GetRcmdPower returns the RcmdPower field if non-nil, zero value otherwise.

### GetRcmdPowerOk

`func (o *ApPlanningRadioVO) GetRcmdPowerOk() (*int32, bool)`

GetRcmdPowerOk returns a tuple with the RcmdPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcmdPower

`func (o *ApPlanningRadioVO) SetRcmdPower(v int32)`

SetRcmdPower sets RcmdPower field to given value.

### HasRcmdPower

`func (o *ApPlanningRadioVO) HasRcmdPower() bool`

HasRcmdPower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


