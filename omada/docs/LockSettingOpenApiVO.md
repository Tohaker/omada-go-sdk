# LockSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | Lock duration should be within the range of 1–1080(min). | [optional] 
**Status** | **bool** | Status of the SSL VPN server name lock or IP lock. | 
**Times** | Pointer to **int32** | The number of login failures that trigger the lock. It is required when parameter [status] is true, and it should be within the range of 1–10 | [optional] 

## Methods

### NewLockSettingOpenApiVO

`func NewLockSettingOpenApiVO(status bool, ) *LockSettingOpenApiVO`

NewLockSettingOpenApiVO instantiates a new LockSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockSettingOpenApiVOWithDefaults

`func NewLockSettingOpenApiVOWithDefaults() *LockSettingOpenApiVO`

NewLockSettingOpenApiVOWithDefaults instantiates a new LockSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *LockSettingOpenApiVO) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *LockSettingOpenApiVO) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *LockSettingOpenApiVO) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *LockSettingOpenApiVO) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetStatus

`func (o *LockSettingOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LockSettingOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LockSettingOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetTimes

`func (o *LockSettingOpenApiVO) GetTimes() int32`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *LockSettingOpenApiVO) GetTimesOk() (*int32, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *LockSettingOpenApiVO) SetTimes(v int32)`

SetTimes sets Times field to given value.

### HasTimes

`func (o *LockSettingOpenApiVO) HasTimes() bool`

HasTimes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


