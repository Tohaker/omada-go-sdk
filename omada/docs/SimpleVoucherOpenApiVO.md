# SimpleVoucherOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Voucher code | [optional] 
**DownLimit** | Pointer to **int64** | Downlink speed limit in Kbps. The value of limit should be within the range of 0–10485760. | [optional] 
**Id** | Pointer to **string** | Voucher ID | [optional] 
**StartTime** | Pointer to **int64** | The expiration date of the voucher, unit: MS | [optional] 
**Status** | Pointer to **int32** | Voucher status. It should be a value as follows: 0: unused, 1: in use, 2: expired | [optional] 
**TimeLeftSec** | Pointer to **int64** | Left duration of voucher, unit: Second | [optional] 
**TimeUsedSec** | Pointer to **int64** | Used duration of voucher, unit: Second | [optional] 
**TimingByClientUsage** | Pointer to **bool** | Whether the voucher is timing by usage and duration type is client duration. When this parameter is true, will not display parameter [timeUsedSec] and [timeLeftSec] | [optional] 
**TrafficLimit** | Pointer to **int64** | Traffic limit in MB. It should be within the range of 1–10485760 | [optional] 
**TrafficLimitFrequency** | Pointer to **int32** | Frequency of traffic limit should be a value as follows: 0: total; 1: daily; 2: weekly; 3: monthly. | [optional] 
**TrafficUnused** | Pointer to **int64** | Unused traffic of the voucher, unit: Byte | [optional] 
**TrafficUsed** | Pointer to **int64** | Used traffic of the voucher, unit: Byte | [optional] 
**UpLimit** | Pointer to **int64** | Uplink speed limit in Kbps. The value of limit should be within the range of 0–10485760. | [optional] 

## Methods

### NewSimpleVoucherOpenApiVO

`func NewSimpleVoucherOpenApiVO() *SimpleVoucherOpenApiVO`

NewSimpleVoucherOpenApiVO instantiates a new SimpleVoucherOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleVoucherOpenApiVOWithDefaults

`func NewSimpleVoucherOpenApiVOWithDefaults() *SimpleVoucherOpenApiVO`

NewSimpleVoucherOpenApiVOWithDefaults instantiates a new SimpleVoucherOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *SimpleVoucherOpenApiVO) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SimpleVoucherOpenApiVO) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SimpleVoucherOpenApiVO) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SimpleVoucherOpenApiVO) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDownLimit

`func (o *SimpleVoucherOpenApiVO) GetDownLimit() int64`

GetDownLimit returns the DownLimit field if non-nil, zero value otherwise.

### GetDownLimitOk

`func (o *SimpleVoucherOpenApiVO) GetDownLimitOk() (*int64, bool)`

GetDownLimitOk returns a tuple with the DownLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimit

`func (o *SimpleVoucherOpenApiVO) SetDownLimit(v int64)`

SetDownLimit sets DownLimit field to given value.

### HasDownLimit

`func (o *SimpleVoucherOpenApiVO) HasDownLimit() bool`

HasDownLimit returns a boolean if a field has been set.

### GetId

`func (o *SimpleVoucherOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleVoucherOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleVoucherOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SimpleVoucherOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartTime

`func (o *SimpleVoucherOpenApiVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SimpleVoucherOpenApiVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SimpleVoucherOpenApiVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *SimpleVoucherOpenApiVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *SimpleVoucherOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SimpleVoucherOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SimpleVoucherOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SimpleVoucherOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeLeftSec

`func (o *SimpleVoucherOpenApiVO) GetTimeLeftSec() int64`

GetTimeLeftSec returns the TimeLeftSec field if non-nil, zero value otherwise.

### GetTimeLeftSecOk

`func (o *SimpleVoucherOpenApiVO) GetTimeLeftSecOk() (*int64, bool)`

GetTimeLeftSecOk returns a tuple with the TimeLeftSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLeftSec

`func (o *SimpleVoucherOpenApiVO) SetTimeLeftSec(v int64)`

SetTimeLeftSec sets TimeLeftSec field to given value.

### HasTimeLeftSec

`func (o *SimpleVoucherOpenApiVO) HasTimeLeftSec() bool`

HasTimeLeftSec returns a boolean if a field has been set.

### GetTimeUsedSec

`func (o *SimpleVoucherOpenApiVO) GetTimeUsedSec() int64`

GetTimeUsedSec returns the TimeUsedSec field if non-nil, zero value otherwise.

### GetTimeUsedSecOk

`func (o *SimpleVoucherOpenApiVO) GetTimeUsedSecOk() (*int64, bool)`

GetTimeUsedSecOk returns a tuple with the TimeUsedSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUsedSec

`func (o *SimpleVoucherOpenApiVO) SetTimeUsedSec(v int64)`

SetTimeUsedSec sets TimeUsedSec field to given value.

### HasTimeUsedSec

`func (o *SimpleVoucherOpenApiVO) HasTimeUsedSec() bool`

HasTimeUsedSec returns a boolean if a field has been set.

### GetTimingByClientUsage

`func (o *SimpleVoucherOpenApiVO) GetTimingByClientUsage() bool`

GetTimingByClientUsage returns the TimingByClientUsage field if non-nil, zero value otherwise.

### GetTimingByClientUsageOk

`func (o *SimpleVoucherOpenApiVO) GetTimingByClientUsageOk() (*bool, bool)`

GetTimingByClientUsageOk returns a tuple with the TimingByClientUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimingByClientUsage

`func (o *SimpleVoucherOpenApiVO) SetTimingByClientUsage(v bool)`

SetTimingByClientUsage sets TimingByClientUsage field to given value.

### HasTimingByClientUsage

`func (o *SimpleVoucherOpenApiVO) HasTimingByClientUsage() bool`

HasTimingByClientUsage returns a boolean if a field has been set.

### GetTrafficLimit

`func (o *SimpleVoucherOpenApiVO) GetTrafficLimit() int64`

GetTrafficLimit returns the TrafficLimit field if non-nil, zero value otherwise.

### GetTrafficLimitOk

`func (o *SimpleVoucherOpenApiVO) GetTrafficLimitOk() (*int64, bool)`

GetTrafficLimitOk returns a tuple with the TrafficLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimit

`func (o *SimpleVoucherOpenApiVO) SetTrafficLimit(v int64)`

SetTrafficLimit sets TrafficLimit field to given value.

### HasTrafficLimit

`func (o *SimpleVoucherOpenApiVO) HasTrafficLimit() bool`

HasTrafficLimit returns a boolean if a field has been set.

### GetTrafficLimitFrequency

`func (o *SimpleVoucherOpenApiVO) GetTrafficLimitFrequency() int32`

GetTrafficLimitFrequency returns the TrafficLimitFrequency field if non-nil, zero value otherwise.

### GetTrafficLimitFrequencyOk

`func (o *SimpleVoucherOpenApiVO) GetTrafficLimitFrequencyOk() (*int32, bool)`

GetTrafficLimitFrequencyOk returns a tuple with the TrafficLimitFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimitFrequency

`func (o *SimpleVoucherOpenApiVO) SetTrafficLimitFrequency(v int32)`

SetTrafficLimitFrequency sets TrafficLimitFrequency field to given value.

### HasTrafficLimitFrequency

`func (o *SimpleVoucherOpenApiVO) HasTrafficLimitFrequency() bool`

HasTrafficLimitFrequency returns a boolean if a field has been set.

### GetTrafficUnused

`func (o *SimpleVoucherOpenApiVO) GetTrafficUnused() int64`

GetTrafficUnused returns the TrafficUnused field if non-nil, zero value otherwise.

### GetTrafficUnusedOk

`func (o *SimpleVoucherOpenApiVO) GetTrafficUnusedOk() (*int64, bool)`

GetTrafficUnusedOk returns a tuple with the TrafficUnused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficUnused

`func (o *SimpleVoucherOpenApiVO) SetTrafficUnused(v int64)`

SetTrafficUnused sets TrafficUnused field to given value.

### HasTrafficUnused

`func (o *SimpleVoucherOpenApiVO) HasTrafficUnused() bool`

HasTrafficUnused returns a boolean if a field has been set.

### GetTrafficUsed

`func (o *SimpleVoucherOpenApiVO) GetTrafficUsed() int64`

GetTrafficUsed returns the TrafficUsed field if non-nil, zero value otherwise.

### GetTrafficUsedOk

`func (o *SimpleVoucherOpenApiVO) GetTrafficUsedOk() (*int64, bool)`

GetTrafficUsedOk returns a tuple with the TrafficUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficUsed

`func (o *SimpleVoucherOpenApiVO) SetTrafficUsed(v int64)`

SetTrafficUsed sets TrafficUsed field to given value.

### HasTrafficUsed

`func (o *SimpleVoucherOpenApiVO) HasTrafficUsed() bool`

HasTrafficUsed returns a boolean if a field has been set.

### GetUpLimit

`func (o *SimpleVoucherOpenApiVO) GetUpLimit() int64`

GetUpLimit returns the UpLimit field if non-nil, zero value otherwise.

### GetUpLimitOk

`func (o *SimpleVoucherOpenApiVO) GetUpLimitOk() (*int64, bool)`

GetUpLimitOk returns a tuple with the UpLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimit

`func (o *SimpleVoucherOpenApiVO) SetUpLimit(v int64)`

SetUpLimit sets UpLimit field to given value.

### HasUpLimit

`func (o *SimpleVoucherOpenApiVO) HasUpLimit() bool`

HasUpLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


