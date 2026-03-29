# CustomRateLimitOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownLimit** | Pointer to **int64** | Downlink speed limit in Kbps. The value of limit should be within the range of 0–10485760(Kbps). | [optional] 
**DownLimitEnable** | **bool** | Whether to enable downlink speed limit. | 
**UpLimit** | Pointer to **int64** | Uplink speed limit in Kbps. The value of limit should be within the range of 0–10485760(Kbps). | [optional] 
**UpLimitEnable** | **bool** | Whether to enable uplink speed limit. | 

## Methods

### NewCustomRateLimitOpenApiVO

`func NewCustomRateLimitOpenApiVO(downLimitEnable bool, upLimitEnable bool, ) *CustomRateLimitOpenApiVO`

NewCustomRateLimitOpenApiVO instantiates a new CustomRateLimitOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomRateLimitOpenApiVOWithDefaults

`func NewCustomRateLimitOpenApiVOWithDefaults() *CustomRateLimitOpenApiVO`

NewCustomRateLimitOpenApiVOWithDefaults instantiates a new CustomRateLimitOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownLimit

`func (o *CustomRateLimitOpenApiVO) GetDownLimit() int64`

GetDownLimit returns the DownLimit field if non-nil, zero value otherwise.

### GetDownLimitOk

`func (o *CustomRateLimitOpenApiVO) GetDownLimitOk() (*int64, bool)`

GetDownLimitOk returns a tuple with the DownLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimit

`func (o *CustomRateLimitOpenApiVO) SetDownLimit(v int64)`

SetDownLimit sets DownLimit field to given value.

### HasDownLimit

`func (o *CustomRateLimitOpenApiVO) HasDownLimit() bool`

HasDownLimit returns a boolean if a field has been set.

### GetDownLimitEnable

`func (o *CustomRateLimitOpenApiVO) GetDownLimitEnable() bool`

GetDownLimitEnable returns the DownLimitEnable field if non-nil, zero value otherwise.

### GetDownLimitEnableOk

`func (o *CustomRateLimitOpenApiVO) GetDownLimitEnableOk() (*bool, bool)`

GetDownLimitEnableOk returns a tuple with the DownLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimitEnable

`func (o *CustomRateLimitOpenApiVO) SetDownLimitEnable(v bool)`

SetDownLimitEnable sets DownLimitEnable field to given value.


### GetUpLimit

`func (o *CustomRateLimitOpenApiVO) GetUpLimit() int64`

GetUpLimit returns the UpLimit field if non-nil, zero value otherwise.

### GetUpLimitOk

`func (o *CustomRateLimitOpenApiVO) GetUpLimitOk() (*int64, bool)`

GetUpLimitOk returns a tuple with the UpLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimit

`func (o *CustomRateLimitOpenApiVO) SetUpLimit(v int64)`

SetUpLimit sets UpLimit field to given value.

### HasUpLimit

`func (o *CustomRateLimitOpenApiVO) HasUpLimit() bool`

HasUpLimit returns a boolean if a field has been set.

### GetUpLimitEnable

`func (o *CustomRateLimitOpenApiVO) GetUpLimitEnable() bool`

GetUpLimitEnable returns the UpLimitEnable field if non-nil, zero value otherwise.

### GetUpLimitEnableOk

`func (o *CustomRateLimitOpenApiVO) GetUpLimitEnableOk() (*bool, bool)`

GetUpLimitEnableOk returns a tuple with the UpLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimitEnable

`func (o *CustomRateLimitOpenApiVO) SetUpLimitEnable(v bool)`

SetUpLimitEnable sets UpLimitEnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


