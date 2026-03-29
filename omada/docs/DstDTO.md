# DstDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | DST config status; If false, other parameters are not required. | [optional] 
**End** | Pointer to [**DstTimeDTO**](DstTimeDTO.md) |  | [optional] 
**EndTime** | Pointer to **int64** | The timeStamp of the DST available end time | [optional] 
**LastEnd** | Pointer to **int64** | The timeStamp of the DST end time of the last year(Unit: ms) | [optional] 
**LastStart** | Pointer to **int64** | The timeStamp of the DST start time of the last year(Unit: ms) | [optional] 
**Mode** | Pointer to **int32** | DST config mode; If disable, other parameters are not required. 0: disable, 1: auto, 2: manually | [optional] 
**NextEnd** | Pointer to **int64** | The timeStamp of the DST end time of the next year(Unit: ms) | [optional] 
**NextStart** | Pointer to **int64** | The timeStamp of the DST start time of the next year(Unit: ms) | [optional] 
**Offset** | Pointer to **int64** | DST offset config(Unit: ms); It should be a value as follows: [1800000, 3600000, 5400000, 7200000]. | [optional] 
**Start** | Pointer to [**DstTimeDTO**](DstTimeDTO.md) |  | [optional] 
**StartTime** | Pointer to **int64** | The timeStamp of the DST available start time | [optional] 
**Status** | Pointer to **bool** | DST available status | [optional] 
**TimeZone** | Pointer to **string** | Timezone of the site | [optional] 

## Methods

### NewDstDTO

`func NewDstDTO() *DstDTO`

NewDstDTO instantiates a new DstDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDstDTOWithDefaults

`func NewDstDTOWithDefaults() *DstDTO`

NewDstDTOWithDefaults instantiates a new DstDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *DstDTO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *DstDTO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *DstDTO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *DstDTO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnd

`func (o *DstDTO) GetEnd() DstTimeDTO`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *DstDTO) GetEndOk() (*DstTimeDTO, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *DstDTO) SetEnd(v DstTimeDTO)`

SetEnd sets End field to given value.

### HasEnd

`func (o *DstDTO) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetEndTime

`func (o *DstDTO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DstDTO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DstDTO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *DstDTO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetLastEnd

`func (o *DstDTO) GetLastEnd() int64`

GetLastEnd returns the LastEnd field if non-nil, zero value otherwise.

### GetLastEndOk

`func (o *DstDTO) GetLastEndOk() (*int64, bool)`

GetLastEndOk returns a tuple with the LastEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEnd

`func (o *DstDTO) SetLastEnd(v int64)`

SetLastEnd sets LastEnd field to given value.

### HasLastEnd

`func (o *DstDTO) HasLastEnd() bool`

HasLastEnd returns a boolean if a field has been set.

### GetLastStart

`func (o *DstDTO) GetLastStart() int64`

GetLastStart returns the LastStart field if non-nil, zero value otherwise.

### GetLastStartOk

`func (o *DstDTO) GetLastStartOk() (*int64, bool)`

GetLastStartOk returns a tuple with the LastStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStart

`func (o *DstDTO) SetLastStart(v int64)`

SetLastStart sets LastStart field to given value.

### HasLastStart

`func (o *DstDTO) HasLastStart() bool`

HasLastStart returns a boolean if a field has been set.

### GetMode

`func (o *DstDTO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DstDTO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DstDTO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DstDTO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNextEnd

`func (o *DstDTO) GetNextEnd() int64`

GetNextEnd returns the NextEnd field if non-nil, zero value otherwise.

### GetNextEndOk

`func (o *DstDTO) GetNextEndOk() (*int64, bool)`

GetNextEndOk returns a tuple with the NextEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextEnd

`func (o *DstDTO) SetNextEnd(v int64)`

SetNextEnd sets NextEnd field to given value.

### HasNextEnd

`func (o *DstDTO) HasNextEnd() bool`

HasNextEnd returns a boolean if a field has been set.

### GetNextStart

`func (o *DstDTO) GetNextStart() int64`

GetNextStart returns the NextStart field if non-nil, zero value otherwise.

### GetNextStartOk

`func (o *DstDTO) GetNextStartOk() (*int64, bool)`

GetNextStartOk returns a tuple with the NextStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStart

`func (o *DstDTO) SetNextStart(v int64)`

SetNextStart sets NextStart field to given value.

### HasNextStart

`func (o *DstDTO) HasNextStart() bool`

HasNextStart returns a boolean if a field has been set.

### GetOffset

`func (o *DstDTO) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *DstDTO) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *DstDTO) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *DstDTO) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetStart

`func (o *DstDTO) GetStart() DstTimeDTO`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *DstDTO) GetStartOk() (*DstTimeDTO, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *DstDTO) SetStart(v DstTimeDTO)`

SetStart sets Start field to given value.

### HasStart

`func (o *DstDTO) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStartTime

`func (o *DstDTO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DstDTO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DstDTO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DstDTO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *DstDTO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DstDTO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DstDTO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DstDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeZone

`func (o *DstDTO) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *DstDTO) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *DstDTO) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *DstDTO) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


