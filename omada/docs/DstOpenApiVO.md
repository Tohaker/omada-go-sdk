# DstOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | DST config status; If false, other parameters are not required. | [optional] 
**End** | Pointer to [**DSTTime**](DSTTime.md) |  | [optional] 
**EndTime** | Pointer to **int64** | The timeStamp of the DST available end time | [optional] 
**LastEnd** | Pointer to **int64** | The timeStamp of the DST end time of the last year(Unit: ms) | [optional] 
**LastStart** | Pointer to **int64** | The timeStamp of the DST start time of the last year(Unit: ms) | [optional] 
**Mode** | Pointer to **int32** | DST mode should be a value as follows: 0: disable; 1: auto; 2: manual. | [optional] 
**NextEnd** | Pointer to **int64** | The timeStamp of the DST end time of the next year(Unit: ms) | [optional] 
**NextStart** | Pointer to **int64** | The timeStamp of the DST start time of the next year(Unit: ms) | [optional] 
**Offset** | Pointer to **int64** | DST offset config(Unit: ms); It should be a value as follows: [1800000, 3600000, 5400000, 7200000]. | [optional] 
**Start** | Pointer to [**DSTTime**](DSTTime.md) |  | [optional] 
**StartTime** | Pointer to **int64** | The timeStamp of the DST available start time | [optional] 
**Status** | Pointer to **bool** | DST available status | [optional] 
**TimeZone** | Pointer to **string** | Timezone of the site | [optional] 

## Methods

### NewDstOpenApiVO

`func NewDstOpenApiVO() *DstOpenApiVO`

NewDstOpenApiVO instantiates a new DstOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDstOpenApiVOWithDefaults

`func NewDstOpenApiVOWithDefaults() *DstOpenApiVO`

NewDstOpenApiVOWithDefaults instantiates a new DstOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *DstOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *DstOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *DstOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *DstOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnd

`func (o *DstOpenApiVO) GetEnd() DSTTime`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *DstOpenApiVO) GetEndOk() (*DSTTime, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *DstOpenApiVO) SetEnd(v DSTTime)`

SetEnd sets End field to given value.

### HasEnd

`func (o *DstOpenApiVO) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetEndTime

`func (o *DstOpenApiVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DstOpenApiVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DstOpenApiVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *DstOpenApiVO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetLastEnd

`func (o *DstOpenApiVO) GetLastEnd() int64`

GetLastEnd returns the LastEnd field if non-nil, zero value otherwise.

### GetLastEndOk

`func (o *DstOpenApiVO) GetLastEndOk() (*int64, bool)`

GetLastEndOk returns a tuple with the LastEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEnd

`func (o *DstOpenApiVO) SetLastEnd(v int64)`

SetLastEnd sets LastEnd field to given value.

### HasLastEnd

`func (o *DstOpenApiVO) HasLastEnd() bool`

HasLastEnd returns a boolean if a field has been set.

### GetLastStart

`func (o *DstOpenApiVO) GetLastStart() int64`

GetLastStart returns the LastStart field if non-nil, zero value otherwise.

### GetLastStartOk

`func (o *DstOpenApiVO) GetLastStartOk() (*int64, bool)`

GetLastStartOk returns a tuple with the LastStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStart

`func (o *DstOpenApiVO) SetLastStart(v int64)`

SetLastStart sets LastStart field to given value.

### HasLastStart

`func (o *DstOpenApiVO) HasLastStart() bool`

HasLastStart returns a boolean if a field has been set.

### GetMode

`func (o *DstOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DstOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DstOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DstOpenApiVO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNextEnd

`func (o *DstOpenApiVO) GetNextEnd() int64`

GetNextEnd returns the NextEnd field if non-nil, zero value otherwise.

### GetNextEndOk

`func (o *DstOpenApiVO) GetNextEndOk() (*int64, bool)`

GetNextEndOk returns a tuple with the NextEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextEnd

`func (o *DstOpenApiVO) SetNextEnd(v int64)`

SetNextEnd sets NextEnd field to given value.

### HasNextEnd

`func (o *DstOpenApiVO) HasNextEnd() bool`

HasNextEnd returns a boolean if a field has been set.

### GetNextStart

`func (o *DstOpenApiVO) GetNextStart() int64`

GetNextStart returns the NextStart field if non-nil, zero value otherwise.

### GetNextStartOk

`func (o *DstOpenApiVO) GetNextStartOk() (*int64, bool)`

GetNextStartOk returns a tuple with the NextStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStart

`func (o *DstOpenApiVO) SetNextStart(v int64)`

SetNextStart sets NextStart field to given value.

### HasNextStart

`func (o *DstOpenApiVO) HasNextStart() bool`

HasNextStart returns a boolean if a field has been set.

### GetOffset

`func (o *DstOpenApiVO) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *DstOpenApiVO) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *DstOpenApiVO) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *DstOpenApiVO) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetStart

`func (o *DstOpenApiVO) GetStart() DSTTime`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *DstOpenApiVO) GetStartOk() (*DSTTime, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *DstOpenApiVO) SetStart(v DSTTime)`

SetStart sets Start field to given value.

### HasStart

`func (o *DstOpenApiVO) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStartTime

`func (o *DstOpenApiVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DstOpenApiVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DstOpenApiVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DstOpenApiVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *DstOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DstOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DstOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DstOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeZone

`func (o *DstOpenApiVO) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *DstOpenApiVO) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *DstOpenApiVO) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *DstOpenApiVO) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


