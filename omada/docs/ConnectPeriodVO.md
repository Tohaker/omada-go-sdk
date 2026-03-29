# ConnectPeriodVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int64** | Duration of a connect interval | [optional] 
**End** | Pointer to **int64** | End timestamp of a connect interval. | [optional] 
**Start** | Pointer to **int64** | Start timestamp of a connect interval. | [optional] 

## Methods

### NewConnectPeriodVO

`func NewConnectPeriodVO() *ConnectPeriodVO`

NewConnectPeriodVO instantiates a new ConnectPeriodVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectPeriodVOWithDefaults

`func NewConnectPeriodVOWithDefaults() *ConnectPeriodVO`

NewConnectPeriodVOWithDefaults instantiates a new ConnectPeriodVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *ConnectPeriodVO) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ConnectPeriodVO) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ConnectPeriodVO) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ConnectPeriodVO) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEnd

`func (o *ConnectPeriodVO) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ConnectPeriodVO) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ConnectPeriodVO) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ConnectPeriodVO) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStart

`func (o *ConnectPeriodVO) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ConnectPeriodVO) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ConnectPeriodVO) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *ConnectPeriodVO) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


