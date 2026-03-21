# RateLimitSettingOfClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownEnable** | Pointer to **bool** | Down limit enable. | [optional] 
**DownLimit** | Pointer to **int64** | Down limit should be within the range of 1–1024. | [optional] 
**DownUnit** | Pointer to **int32** | Down limit unit should be a value as follows: 1: Kbps; 2: Mbps. | [optional] 
**Enable** | Pointer to **bool** | Rate limit enable. | [optional] 
**RateLimitId** | Pointer to **string** | Rate limit profile ID. Nullable when ratelimit type is custom. | [optional] 
**UpEnable** | Pointer to **bool** | Up limit enable. | [optional] 
**UpLimit** | Pointer to **int64** | Up limit should be within the range of 1–1024. | [optional] 
**UpUnit** | Pointer to **int32** | Up limit unit should be a value as follows: 1: Kbps; 2: Mbps. | [optional] 

## Methods

### NewRateLimitSettingOfClient

`func NewRateLimitSettingOfClient() *RateLimitSettingOfClient`

NewRateLimitSettingOfClient instantiates a new RateLimitSettingOfClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitSettingOfClientWithDefaults

`func NewRateLimitSettingOfClientWithDefaults() *RateLimitSettingOfClient`

NewRateLimitSettingOfClientWithDefaults instantiates a new RateLimitSettingOfClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownEnable

`func (o *RateLimitSettingOfClient) GetDownEnable() bool`

GetDownEnable returns the DownEnable field if non-nil, zero value otherwise.

### GetDownEnableOk

`func (o *RateLimitSettingOfClient) GetDownEnableOk() (*bool, bool)`

GetDownEnableOk returns a tuple with the DownEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownEnable

`func (o *RateLimitSettingOfClient) SetDownEnable(v bool)`

SetDownEnable sets DownEnable field to given value.

### HasDownEnable

`func (o *RateLimitSettingOfClient) HasDownEnable() bool`

HasDownEnable returns a boolean if a field has been set.

### GetDownLimit

`func (o *RateLimitSettingOfClient) GetDownLimit() int64`

GetDownLimit returns the DownLimit field if non-nil, zero value otherwise.

### GetDownLimitOk

`func (o *RateLimitSettingOfClient) GetDownLimitOk() (*int64, bool)`

GetDownLimitOk returns a tuple with the DownLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimit

`func (o *RateLimitSettingOfClient) SetDownLimit(v int64)`

SetDownLimit sets DownLimit field to given value.

### HasDownLimit

`func (o *RateLimitSettingOfClient) HasDownLimit() bool`

HasDownLimit returns a boolean if a field has been set.

### GetDownUnit

`func (o *RateLimitSettingOfClient) GetDownUnit() int32`

GetDownUnit returns the DownUnit field if non-nil, zero value otherwise.

### GetDownUnitOk

`func (o *RateLimitSettingOfClient) GetDownUnitOk() (*int32, bool)`

GetDownUnitOk returns a tuple with the DownUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownUnit

`func (o *RateLimitSettingOfClient) SetDownUnit(v int32)`

SetDownUnit sets DownUnit field to given value.

### HasDownUnit

`func (o *RateLimitSettingOfClient) HasDownUnit() bool`

HasDownUnit returns a boolean if a field has been set.

### GetEnable

`func (o *RateLimitSettingOfClient) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *RateLimitSettingOfClient) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *RateLimitSettingOfClient) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *RateLimitSettingOfClient) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetRateLimitId

`func (o *RateLimitSettingOfClient) GetRateLimitId() string`

GetRateLimitId returns the RateLimitId field if non-nil, zero value otherwise.

### GetRateLimitIdOk

`func (o *RateLimitSettingOfClient) GetRateLimitIdOk() (*string, bool)`

GetRateLimitIdOk returns a tuple with the RateLimitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitId

`func (o *RateLimitSettingOfClient) SetRateLimitId(v string)`

SetRateLimitId sets RateLimitId field to given value.

### HasRateLimitId

`func (o *RateLimitSettingOfClient) HasRateLimitId() bool`

HasRateLimitId returns a boolean if a field has been set.

### GetUpEnable

`func (o *RateLimitSettingOfClient) GetUpEnable() bool`

GetUpEnable returns the UpEnable field if non-nil, zero value otherwise.

### GetUpEnableOk

`func (o *RateLimitSettingOfClient) GetUpEnableOk() (*bool, bool)`

GetUpEnableOk returns a tuple with the UpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpEnable

`func (o *RateLimitSettingOfClient) SetUpEnable(v bool)`

SetUpEnable sets UpEnable field to given value.

### HasUpEnable

`func (o *RateLimitSettingOfClient) HasUpEnable() bool`

HasUpEnable returns a boolean if a field has been set.

### GetUpLimit

`func (o *RateLimitSettingOfClient) GetUpLimit() int64`

GetUpLimit returns the UpLimit field if non-nil, zero value otherwise.

### GetUpLimitOk

`func (o *RateLimitSettingOfClient) GetUpLimitOk() (*int64, bool)`

GetUpLimitOk returns a tuple with the UpLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimit

`func (o *RateLimitSettingOfClient) SetUpLimit(v int64)`

SetUpLimit sets UpLimit field to given value.

### HasUpLimit

`func (o *RateLimitSettingOfClient) HasUpLimit() bool`

HasUpLimit returns a boolean if a field has been set.

### GetUpUnit

`func (o *RateLimitSettingOfClient) GetUpUnit() int32`

GetUpUnit returns the UpUnit field if non-nil, zero value otherwise.

### GetUpUnitOk

`func (o *RateLimitSettingOfClient) GetUpUnitOk() (*int32, bool)`

GetUpUnitOk returns a tuple with the UpUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpUnit

`func (o *RateLimitSettingOfClient) SetUpUnit(v int32)`

SetUpUnit sets UpUnit field to given value.

### HasUpUnit

`func (o *RateLimitSettingOfClient) HasUpUnit() bool`

HasUpUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


