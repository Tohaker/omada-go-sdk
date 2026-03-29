# ONUPortRateLimitVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rate** | Pointer to **int32** | Rate should be within the range 1 - 10240 | [optional] 
**Status** | Pointer to **bool** | Enable the switch. | [optional] 

## Methods

### NewONUPortRateLimitVO

`func NewONUPortRateLimitVO() *ONUPortRateLimitVO`

NewONUPortRateLimitVO instantiates a new ONUPortRateLimitVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewONUPortRateLimitVOWithDefaults

`func NewONUPortRateLimitVOWithDefaults() *ONUPortRateLimitVO`

NewONUPortRateLimitVOWithDefaults instantiates a new ONUPortRateLimitVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRate

`func (o *ONUPortRateLimitVO) GetRate() int32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ONUPortRateLimitVO) GetRateOk() (*int32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ONUPortRateLimitVO) SetRate(v int32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ONUPortRateLimitVO) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetStatus

`func (o *ONUPortRateLimitVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ONUPortRateLimitVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ONUPortRateLimitVO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ONUPortRateLimitVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


