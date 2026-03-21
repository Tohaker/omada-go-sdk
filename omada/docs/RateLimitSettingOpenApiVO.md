# RateLimitSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomSetting** | Pointer to [**CustomRateLimitSettingOpenApiVO**](CustomRateLimitSettingOpenApiVO.md) |  | [optional] 
**ProfileId** | Pointer to **string** | This field represents RateLimit Profile ID. RateLimit Profile can be created using Create rate limit profile interface, and RateLimit Profile ID can be obtained from Get rate limit profile list interface.(The validity priority is higher than the custom setting) | [optional] 

## Methods

### NewRateLimitSettingOpenApiVO

`func NewRateLimitSettingOpenApiVO() *RateLimitSettingOpenApiVO`

NewRateLimitSettingOpenApiVO instantiates a new RateLimitSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitSettingOpenApiVOWithDefaults

`func NewRateLimitSettingOpenApiVOWithDefaults() *RateLimitSettingOpenApiVO`

NewRateLimitSettingOpenApiVOWithDefaults instantiates a new RateLimitSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomSetting

`func (o *RateLimitSettingOpenApiVO) GetCustomSetting() CustomRateLimitSettingOpenApiVO`

GetCustomSetting returns the CustomSetting field if non-nil, zero value otherwise.

### GetCustomSettingOk

`func (o *RateLimitSettingOpenApiVO) GetCustomSettingOk() (*CustomRateLimitSettingOpenApiVO, bool)`

GetCustomSettingOk returns a tuple with the CustomSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSetting

`func (o *RateLimitSettingOpenApiVO) SetCustomSetting(v CustomRateLimitSettingOpenApiVO)`

SetCustomSetting sets CustomSetting field to given value.

### HasCustomSetting

`func (o *RateLimitSettingOpenApiVO) HasCustomSetting() bool`

HasCustomSetting returns a boolean if a field has been set.

### GetProfileId

`func (o *RateLimitSettingOpenApiVO) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *RateLimitSettingOpenApiVO) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *RateLimitSettingOpenApiVO) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *RateLimitSettingOpenApiVO) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


