# CustomRateLimitSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownLimit** | Pointer to **int32** | Downlink speed limit value. When the value of Parameter [downLimitType] is 0(Kbps), downLimit should be within the range of 1–10485760; when the value of Parameter [downLimitType] is 1(Mbps), downLimit should be within the range of 1-10240. | [optional] 
**DownLimitEnable** | **bool** | Whether to limit downlink speed; This field is required when select custom setting. True: enable, false: disable. | 
**DownLimitType** | Pointer to **int32** | Downlink speed limit unit config; DownLimitType should be a value as follows: 0: Kbps; 1: Mbps. | [optional] 
**UpLimit** | Pointer to **int32** | Uplink speed limit value. When the value of Parameter [upLimitType] is 0(Kbps), upLimit should be within the range of 1–10485760; when the value of Parameter [upLimitType] is 1(Mbps), upLimit should be within the range of 1-10240. | [optional] 
**UpLimitEnable** | **bool** | Whether to limit uplink speed; This field is required when select custom setting. True: enable, false: disable. | 
**UpLimitType** | Pointer to **int32** | Uplink speed limit unit config; UpLimitType should be a value as follows: 0: Kbps; 1: Mbps. | [optional] 

## Methods

### NewCustomRateLimitSettingOpenApiVO

`func NewCustomRateLimitSettingOpenApiVO(downLimitEnable bool, upLimitEnable bool, ) *CustomRateLimitSettingOpenApiVO`

NewCustomRateLimitSettingOpenApiVO instantiates a new CustomRateLimitSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomRateLimitSettingOpenApiVOWithDefaults

`func NewCustomRateLimitSettingOpenApiVOWithDefaults() *CustomRateLimitSettingOpenApiVO`

NewCustomRateLimitSettingOpenApiVOWithDefaults instantiates a new CustomRateLimitSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownLimit

`func (o *CustomRateLimitSettingOpenApiVO) GetDownLimit() int32`

GetDownLimit returns the DownLimit field if non-nil, zero value otherwise.

### GetDownLimitOk

`func (o *CustomRateLimitSettingOpenApiVO) GetDownLimitOk() (*int32, bool)`

GetDownLimitOk returns a tuple with the DownLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimit

`func (o *CustomRateLimitSettingOpenApiVO) SetDownLimit(v int32)`

SetDownLimit sets DownLimit field to given value.

### HasDownLimit

`func (o *CustomRateLimitSettingOpenApiVO) HasDownLimit() bool`

HasDownLimit returns a boolean if a field has been set.

### GetDownLimitEnable

`func (o *CustomRateLimitSettingOpenApiVO) GetDownLimitEnable() bool`

GetDownLimitEnable returns the DownLimitEnable field if non-nil, zero value otherwise.

### GetDownLimitEnableOk

`func (o *CustomRateLimitSettingOpenApiVO) GetDownLimitEnableOk() (*bool, bool)`

GetDownLimitEnableOk returns a tuple with the DownLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimitEnable

`func (o *CustomRateLimitSettingOpenApiVO) SetDownLimitEnable(v bool)`

SetDownLimitEnable sets DownLimitEnable field to given value.


### GetDownLimitType

`func (o *CustomRateLimitSettingOpenApiVO) GetDownLimitType() int32`

GetDownLimitType returns the DownLimitType field if non-nil, zero value otherwise.

### GetDownLimitTypeOk

`func (o *CustomRateLimitSettingOpenApiVO) GetDownLimitTypeOk() (*int32, bool)`

GetDownLimitTypeOk returns a tuple with the DownLimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimitType

`func (o *CustomRateLimitSettingOpenApiVO) SetDownLimitType(v int32)`

SetDownLimitType sets DownLimitType field to given value.

### HasDownLimitType

`func (o *CustomRateLimitSettingOpenApiVO) HasDownLimitType() bool`

HasDownLimitType returns a boolean if a field has been set.

### GetUpLimit

`func (o *CustomRateLimitSettingOpenApiVO) GetUpLimit() int32`

GetUpLimit returns the UpLimit field if non-nil, zero value otherwise.

### GetUpLimitOk

`func (o *CustomRateLimitSettingOpenApiVO) GetUpLimitOk() (*int32, bool)`

GetUpLimitOk returns a tuple with the UpLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimit

`func (o *CustomRateLimitSettingOpenApiVO) SetUpLimit(v int32)`

SetUpLimit sets UpLimit field to given value.

### HasUpLimit

`func (o *CustomRateLimitSettingOpenApiVO) HasUpLimit() bool`

HasUpLimit returns a boolean if a field has been set.

### GetUpLimitEnable

`func (o *CustomRateLimitSettingOpenApiVO) GetUpLimitEnable() bool`

GetUpLimitEnable returns the UpLimitEnable field if non-nil, zero value otherwise.

### GetUpLimitEnableOk

`func (o *CustomRateLimitSettingOpenApiVO) GetUpLimitEnableOk() (*bool, bool)`

GetUpLimitEnableOk returns a tuple with the UpLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimitEnable

`func (o *CustomRateLimitSettingOpenApiVO) SetUpLimitEnable(v bool)`

SetUpLimitEnable sets UpLimitEnable field to given value.


### GetUpLimitType

`func (o *CustomRateLimitSettingOpenApiVO) GetUpLimitType() int32`

GetUpLimitType returns the UpLimitType field if non-nil, zero value otherwise.

### GetUpLimitTypeOk

`func (o *CustomRateLimitSettingOpenApiVO) GetUpLimitTypeOk() (*int32, bool)`

GetUpLimitTypeOk returns a tuple with the UpLimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimitType

`func (o *CustomRateLimitSettingOpenApiVO) SetUpLimitType(v int32)`

SetUpLimitType sets UpLimitType field to given value.

### HasUpLimitType

`func (o *CustomRateLimitSettingOpenApiVO) HasUpLimitType() bool`

HasUpLimitType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


