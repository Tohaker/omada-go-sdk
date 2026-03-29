# PPSKExpirationVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayEndHour** | Pointer to **int32** | The hour of end time, should be within the range of 0-23 When parameter [type] is 3, parameter [dayEndHour] is required | [optional] 
**DayEndMin** | Pointer to **int32** | The minute of end time, should be within the range of 0-59 When parameter [type] is 3, parameter [dayEndMin] is required | [optional] 
**DayStartHour** | Pointer to **int32** | The hour of start time, should be within the range of 0-23. When parameter [type] is 3, parameter [dayStartHour] is required | [optional] 
**DayStartMin** | Pointer to **int32** | The minute of start time, should be within the range of 0-59 When parameter [type] is 3, parameter [dayStartMin] is required | [optional] 
**Duration** | Pointer to **int32** | Duration of one use. When parameter [type] is 2, parameter [duration] is required | [optional] 
**DurationUnit** | Pointer to **int32** | Unit of  parameter [duration], should be a value as follows: 0: hour, 1: day, 2: week. When parameter [type] is 2, parameter [duration] is required | [optional] 
**EffectiveTime** | Pointer to **int64** | The timestamp when the PPSK takes effect, unit: millisecond. When parameter [type] is 1, parameter [effectiveTime] is required | [optional] 
**ExpirationTime** | Pointer to **int64** | The timestamp of the expiration of PPSK, unit: millisecond. When parameter [type] is 1, parameter [expirationTime] is required | [optional] 
**Type** | Pointer to **int32** | Expiration type, should be a value as follows: 0: Never expire. 1: PPSK can be used between the effective time and expiration time. 2: PPSK can be used after a period of time of creation. 3: PPSK can be used during designated time periods every day | [optional] 

## Methods

### NewPPSKExpirationVO

`func NewPPSKExpirationVO() *PPSKExpirationVO`

NewPPSKExpirationVO instantiates a new PPSKExpirationVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPPSKExpirationVOWithDefaults

`func NewPPSKExpirationVOWithDefaults() *PPSKExpirationVO`

NewPPSKExpirationVOWithDefaults instantiates a new PPSKExpirationVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayEndHour

`func (o *PPSKExpirationVO) GetDayEndHour() int32`

GetDayEndHour returns the DayEndHour field if non-nil, zero value otherwise.

### GetDayEndHourOk

`func (o *PPSKExpirationVO) GetDayEndHourOk() (*int32, bool)`

GetDayEndHourOk returns a tuple with the DayEndHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayEndHour

`func (o *PPSKExpirationVO) SetDayEndHour(v int32)`

SetDayEndHour sets DayEndHour field to given value.

### HasDayEndHour

`func (o *PPSKExpirationVO) HasDayEndHour() bool`

HasDayEndHour returns a boolean if a field has been set.

### GetDayEndMin

`func (o *PPSKExpirationVO) GetDayEndMin() int32`

GetDayEndMin returns the DayEndMin field if non-nil, zero value otherwise.

### GetDayEndMinOk

`func (o *PPSKExpirationVO) GetDayEndMinOk() (*int32, bool)`

GetDayEndMinOk returns a tuple with the DayEndMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayEndMin

`func (o *PPSKExpirationVO) SetDayEndMin(v int32)`

SetDayEndMin sets DayEndMin field to given value.

### HasDayEndMin

`func (o *PPSKExpirationVO) HasDayEndMin() bool`

HasDayEndMin returns a boolean if a field has been set.

### GetDayStartHour

`func (o *PPSKExpirationVO) GetDayStartHour() int32`

GetDayStartHour returns the DayStartHour field if non-nil, zero value otherwise.

### GetDayStartHourOk

`func (o *PPSKExpirationVO) GetDayStartHourOk() (*int32, bool)`

GetDayStartHourOk returns a tuple with the DayStartHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayStartHour

`func (o *PPSKExpirationVO) SetDayStartHour(v int32)`

SetDayStartHour sets DayStartHour field to given value.

### HasDayStartHour

`func (o *PPSKExpirationVO) HasDayStartHour() bool`

HasDayStartHour returns a boolean if a field has been set.

### GetDayStartMin

`func (o *PPSKExpirationVO) GetDayStartMin() int32`

GetDayStartMin returns the DayStartMin field if non-nil, zero value otherwise.

### GetDayStartMinOk

`func (o *PPSKExpirationVO) GetDayStartMinOk() (*int32, bool)`

GetDayStartMinOk returns a tuple with the DayStartMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayStartMin

`func (o *PPSKExpirationVO) SetDayStartMin(v int32)`

SetDayStartMin sets DayStartMin field to given value.

### HasDayStartMin

`func (o *PPSKExpirationVO) HasDayStartMin() bool`

HasDayStartMin returns a boolean if a field has been set.

### GetDuration

`func (o *PPSKExpirationVO) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PPSKExpirationVO) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PPSKExpirationVO) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *PPSKExpirationVO) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDurationUnit

`func (o *PPSKExpirationVO) GetDurationUnit() int32`

GetDurationUnit returns the DurationUnit field if non-nil, zero value otherwise.

### GetDurationUnitOk

`func (o *PPSKExpirationVO) GetDurationUnitOk() (*int32, bool)`

GetDurationUnitOk returns a tuple with the DurationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationUnit

`func (o *PPSKExpirationVO) SetDurationUnit(v int32)`

SetDurationUnit sets DurationUnit field to given value.

### HasDurationUnit

`func (o *PPSKExpirationVO) HasDurationUnit() bool`

HasDurationUnit returns a boolean if a field has been set.

### GetEffectiveTime

`func (o *PPSKExpirationVO) GetEffectiveTime() int64`

GetEffectiveTime returns the EffectiveTime field if non-nil, zero value otherwise.

### GetEffectiveTimeOk

`func (o *PPSKExpirationVO) GetEffectiveTimeOk() (*int64, bool)`

GetEffectiveTimeOk returns a tuple with the EffectiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTime

`func (o *PPSKExpirationVO) SetEffectiveTime(v int64)`

SetEffectiveTime sets EffectiveTime field to given value.

### HasEffectiveTime

`func (o *PPSKExpirationVO) HasEffectiveTime() bool`

HasEffectiveTime returns a boolean if a field has been set.

### GetExpirationTime

`func (o *PPSKExpirationVO) GetExpirationTime() int64`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *PPSKExpirationVO) GetExpirationTimeOk() (*int64, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *PPSKExpirationVO) SetExpirationTime(v int64)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *PPSKExpirationVO) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetType

`func (o *PPSKExpirationVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PPSKExpirationVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PPSKExpirationVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *PPSKExpirationVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


