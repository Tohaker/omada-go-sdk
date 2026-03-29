# Retry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetryRate** | Pointer to **float64** | AP re-contracting rate within one hour | [optional] 
**RetryTimes** | Pointer to **int32** | Number of new contracts within one hour, such as 60 : 60% | [optional] 
**Time** | Pointer to **int64** | Sampling time second | [optional] 

## Methods

### NewRetry

`func NewRetry() *Retry`

NewRetry instantiates a new Retry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryWithDefaults

`func NewRetryWithDefaults() *Retry`

NewRetryWithDefaults instantiates a new Retry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetryRate

`func (o *Retry) GetRetryRate() float64`

GetRetryRate returns the RetryRate field if non-nil, zero value otherwise.

### GetRetryRateOk

`func (o *Retry) GetRetryRateOk() (*float64, bool)`

GetRetryRateOk returns a tuple with the RetryRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryRate

`func (o *Retry) SetRetryRate(v float64)`

SetRetryRate sets RetryRate field to given value.

### HasRetryRate

`func (o *Retry) HasRetryRate() bool`

HasRetryRate returns a boolean if a field has been set.

### GetRetryTimes

`func (o *Retry) GetRetryTimes() int32`

GetRetryTimes returns the RetryTimes field if non-nil, zero value otherwise.

### GetRetryTimesOk

`func (o *Retry) GetRetryTimesOk() (*int32, bool)`

GetRetryTimesOk returns a tuple with the RetryTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryTimes

`func (o *Retry) SetRetryTimes(v int32)`

SetRetryTimes sets RetryTimes field to given value.

### HasRetryTimes

`func (o *Retry) HasRetryTimes() bool`

HasRetryTimes returns a boolean if a field has been set.

### GetTime

`func (o *Retry) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Retry) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Retry) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *Retry) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


