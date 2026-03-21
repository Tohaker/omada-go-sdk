# PPSKRateLimitSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownLimit** | Pointer to **int32** | Downlink speed limit value. When the value of Parameter [downLimitType] is 0(Kbps), downLimit should be within the range of 1–10485760; when the value of Parameter [downLimitType] is 1(Mbps), downLimit should be within the range of 1-10240. | [optional] 
**DownLimitEnable** | Pointer to **bool** | Whether to limit downlink speed; This field is required when select custom setting. True: enable, false: disable. | [optional] 
**DownLimitType** | Pointer to **int32** | Downlink speed limit unit config; DownLimitType should be a value as follows: 0: Kbps; 1: Mbps. | [optional] 
**RateLimitProfileId** | Pointer to **string** | This field represents Rate limit profile ID. Rate limit profile can be created using &#39;Create rate limit profile&#39; interface, and Rate limit profile ID can be obtained from &#39;Get rate limit profile list&#39; interface | [optional] 
**UpLimit** | Pointer to **int32** | Uplink speed limit value. When the value of Parameter [upLimitType] is 0(Kbps), upLimit should be within the range of 1–10485760; when the value of Parameter [upLimitType] is 1(Mbps), upLimit should be within the range of 1-10240. | [optional] 
**UpLimitEnable** | Pointer to **bool** | Whether to limit uplink speed; This field is required when select custom setting. True: enable, false: disable. | [optional] 
**UpLimitType** | Pointer to **int32** | Uplink speed limit unit config; UpLimitType should be a value as follows: 0: Kbps; 1: Mbps. | [optional] 

## Methods

### NewPPSKRateLimitSettingVO

`func NewPPSKRateLimitSettingVO() *PPSKRateLimitSettingVO`

NewPPSKRateLimitSettingVO instantiates a new PPSKRateLimitSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPPSKRateLimitSettingVOWithDefaults

`func NewPPSKRateLimitSettingVOWithDefaults() *PPSKRateLimitSettingVO`

NewPPSKRateLimitSettingVOWithDefaults instantiates a new PPSKRateLimitSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownLimit

`func (o *PPSKRateLimitSettingVO) GetDownLimit() int32`

GetDownLimit returns the DownLimit field if non-nil, zero value otherwise.

### GetDownLimitOk

`func (o *PPSKRateLimitSettingVO) GetDownLimitOk() (*int32, bool)`

GetDownLimitOk returns a tuple with the DownLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimit

`func (o *PPSKRateLimitSettingVO) SetDownLimit(v int32)`

SetDownLimit sets DownLimit field to given value.

### HasDownLimit

`func (o *PPSKRateLimitSettingVO) HasDownLimit() bool`

HasDownLimit returns a boolean if a field has been set.

### GetDownLimitEnable

`func (o *PPSKRateLimitSettingVO) GetDownLimitEnable() bool`

GetDownLimitEnable returns the DownLimitEnable field if non-nil, zero value otherwise.

### GetDownLimitEnableOk

`func (o *PPSKRateLimitSettingVO) GetDownLimitEnableOk() (*bool, bool)`

GetDownLimitEnableOk returns a tuple with the DownLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimitEnable

`func (o *PPSKRateLimitSettingVO) SetDownLimitEnable(v bool)`

SetDownLimitEnable sets DownLimitEnable field to given value.

### HasDownLimitEnable

`func (o *PPSKRateLimitSettingVO) HasDownLimitEnable() bool`

HasDownLimitEnable returns a boolean if a field has been set.

### GetDownLimitType

`func (o *PPSKRateLimitSettingVO) GetDownLimitType() int32`

GetDownLimitType returns the DownLimitType field if non-nil, zero value otherwise.

### GetDownLimitTypeOk

`func (o *PPSKRateLimitSettingVO) GetDownLimitTypeOk() (*int32, bool)`

GetDownLimitTypeOk returns a tuple with the DownLimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimitType

`func (o *PPSKRateLimitSettingVO) SetDownLimitType(v int32)`

SetDownLimitType sets DownLimitType field to given value.

### HasDownLimitType

`func (o *PPSKRateLimitSettingVO) HasDownLimitType() bool`

HasDownLimitType returns a boolean if a field has been set.

### GetRateLimitProfileId

`func (o *PPSKRateLimitSettingVO) GetRateLimitProfileId() string`

GetRateLimitProfileId returns the RateLimitProfileId field if non-nil, zero value otherwise.

### GetRateLimitProfileIdOk

`func (o *PPSKRateLimitSettingVO) GetRateLimitProfileIdOk() (*string, bool)`

GetRateLimitProfileIdOk returns a tuple with the RateLimitProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitProfileId

`func (o *PPSKRateLimitSettingVO) SetRateLimitProfileId(v string)`

SetRateLimitProfileId sets RateLimitProfileId field to given value.

### HasRateLimitProfileId

`func (o *PPSKRateLimitSettingVO) HasRateLimitProfileId() bool`

HasRateLimitProfileId returns a boolean if a field has been set.

### GetUpLimit

`func (o *PPSKRateLimitSettingVO) GetUpLimit() int32`

GetUpLimit returns the UpLimit field if non-nil, zero value otherwise.

### GetUpLimitOk

`func (o *PPSKRateLimitSettingVO) GetUpLimitOk() (*int32, bool)`

GetUpLimitOk returns a tuple with the UpLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimit

`func (o *PPSKRateLimitSettingVO) SetUpLimit(v int32)`

SetUpLimit sets UpLimit field to given value.

### HasUpLimit

`func (o *PPSKRateLimitSettingVO) HasUpLimit() bool`

HasUpLimit returns a boolean if a field has been set.

### GetUpLimitEnable

`func (o *PPSKRateLimitSettingVO) GetUpLimitEnable() bool`

GetUpLimitEnable returns the UpLimitEnable field if non-nil, zero value otherwise.

### GetUpLimitEnableOk

`func (o *PPSKRateLimitSettingVO) GetUpLimitEnableOk() (*bool, bool)`

GetUpLimitEnableOk returns a tuple with the UpLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimitEnable

`func (o *PPSKRateLimitSettingVO) SetUpLimitEnable(v bool)`

SetUpLimitEnable sets UpLimitEnable field to given value.

### HasUpLimitEnable

`func (o *PPSKRateLimitSettingVO) HasUpLimitEnable() bool`

HasUpLimitEnable returns a boolean if a field has been set.

### GetUpLimitType

`func (o *PPSKRateLimitSettingVO) GetUpLimitType() int32`

GetUpLimitType returns the UpLimitType field if non-nil, zero value otherwise.

### GetUpLimitTypeOk

`func (o *PPSKRateLimitSettingVO) GetUpLimitTypeOk() (*int32, bool)`

GetUpLimitTypeOk returns a tuple with the UpLimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimitType

`func (o *PPSKRateLimitSettingVO) SetUpLimitType(v int32)`

SetUpLimitType sets UpLimitType field to given value.

### HasUpLimitType

`func (o *PPSKRateLimitSettingVO) HasUpLimitType() bool`

HasUpLimitType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


