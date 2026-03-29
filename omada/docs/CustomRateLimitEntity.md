# CustomRateLimitEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownEnable** | Pointer to **bool** | Down limit enable | [optional] 
**DownLimit** | Pointer to **int32** | Down limit should be within the range of 1–1024. | [optional] 
**DownUnit** | Pointer to **int32** | Down limit unit should be a value as follows: 1: Kbps; 2: Mbps | [optional] 
**UpEnable** | Pointer to **bool** | Up limit enable | [optional] 
**UpLimit** | Pointer to **int32** | Up limit should be within the range of 1–1024. | [optional] 
**UpUnit** | Pointer to **int32** | Up limit unit should be a value as follows: 1: Kbps; 2: Mbps | [optional] 

## Methods

### NewCustomRateLimitEntity

`func NewCustomRateLimitEntity() *CustomRateLimitEntity`

NewCustomRateLimitEntity instantiates a new CustomRateLimitEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomRateLimitEntityWithDefaults

`func NewCustomRateLimitEntityWithDefaults() *CustomRateLimitEntity`

NewCustomRateLimitEntityWithDefaults instantiates a new CustomRateLimitEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownEnable

`func (o *CustomRateLimitEntity) GetDownEnable() bool`

GetDownEnable returns the DownEnable field if non-nil, zero value otherwise.

### GetDownEnableOk

`func (o *CustomRateLimitEntity) GetDownEnableOk() (*bool, bool)`

GetDownEnableOk returns a tuple with the DownEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownEnable

`func (o *CustomRateLimitEntity) SetDownEnable(v bool)`

SetDownEnable sets DownEnable field to given value.

### HasDownEnable

`func (o *CustomRateLimitEntity) HasDownEnable() bool`

HasDownEnable returns a boolean if a field has been set.

### GetDownLimit

`func (o *CustomRateLimitEntity) GetDownLimit() int32`

GetDownLimit returns the DownLimit field if non-nil, zero value otherwise.

### GetDownLimitOk

`func (o *CustomRateLimitEntity) GetDownLimitOk() (*int32, bool)`

GetDownLimitOk returns a tuple with the DownLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimit

`func (o *CustomRateLimitEntity) SetDownLimit(v int32)`

SetDownLimit sets DownLimit field to given value.

### HasDownLimit

`func (o *CustomRateLimitEntity) HasDownLimit() bool`

HasDownLimit returns a boolean if a field has been set.

### GetDownUnit

`func (o *CustomRateLimitEntity) GetDownUnit() int32`

GetDownUnit returns the DownUnit field if non-nil, zero value otherwise.

### GetDownUnitOk

`func (o *CustomRateLimitEntity) GetDownUnitOk() (*int32, bool)`

GetDownUnitOk returns a tuple with the DownUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownUnit

`func (o *CustomRateLimitEntity) SetDownUnit(v int32)`

SetDownUnit sets DownUnit field to given value.

### HasDownUnit

`func (o *CustomRateLimitEntity) HasDownUnit() bool`

HasDownUnit returns a boolean if a field has been set.

### GetUpEnable

`func (o *CustomRateLimitEntity) GetUpEnable() bool`

GetUpEnable returns the UpEnable field if non-nil, zero value otherwise.

### GetUpEnableOk

`func (o *CustomRateLimitEntity) GetUpEnableOk() (*bool, bool)`

GetUpEnableOk returns a tuple with the UpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpEnable

`func (o *CustomRateLimitEntity) SetUpEnable(v bool)`

SetUpEnable sets UpEnable field to given value.

### HasUpEnable

`func (o *CustomRateLimitEntity) HasUpEnable() bool`

HasUpEnable returns a boolean if a field has been set.

### GetUpLimit

`func (o *CustomRateLimitEntity) GetUpLimit() int32`

GetUpLimit returns the UpLimit field if non-nil, zero value otherwise.

### GetUpLimitOk

`func (o *CustomRateLimitEntity) GetUpLimitOk() (*int32, bool)`

GetUpLimitOk returns a tuple with the UpLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimit

`func (o *CustomRateLimitEntity) SetUpLimit(v int32)`

SetUpLimit sets UpLimit field to given value.

### HasUpLimit

`func (o *CustomRateLimitEntity) HasUpLimit() bool`

HasUpLimit returns a boolean if a field has been set.

### GetUpUnit

`func (o *CustomRateLimitEntity) GetUpUnit() int32`

GetUpUnit returns the UpUnit field if non-nil, zero value otherwise.

### GetUpUnitOk

`func (o *CustomRateLimitEntity) GetUpUnitOk() (*int32, bool)`

GetUpUnitOk returns a tuple with the UpUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpUnit

`func (o *CustomRateLimitEntity) SetUpUnit(v int32)`

SetUpUnit sets UpUnit field to given value.

### HasUpUnit

`func (o *CustomRateLimitEntity) HasUpUnit() bool`

HasUpUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


