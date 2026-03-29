# ClientRateLimitSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomRateLimit** | Pointer to [**CustomRateLimitEntity**](CustomRateLimitEntity.md) |  | [optional] 
**Mode** | **int32** | Rate limit mode should be a value as follows: &lt;br/&gt;0: Custom mode. Apply the given rate limit value to the client; &lt;br/&gt;1: Rate limit profile mode. Find the corresponding rate limit file with rate limit ID and apply it to the client. | 
**RateLimitProfileId** | Pointer to **string** | Rate limit profile ID. Required when ratelimit mode is 1 | [optional] 

## Methods

### NewClientRateLimitSetting

`func NewClientRateLimitSetting(mode int32, ) *ClientRateLimitSetting`

NewClientRateLimitSetting instantiates a new ClientRateLimitSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientRateLimitSettingWithDefaults

`func NewClientRateLimitSettingWithDefaults() *ClientRateLimitSetting`

NewClientRateLimitSettingWithDefaults instantiates a new ClientRateLimitSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomRateLimit

`func (o *ClientRateLimitSetting) GetCustomRateLimit() CustomRateLimitEntity`

GetCustomRateLimit returns the CustomRateLimit field if non-nil, zero value otherwise.

### GetCustomRateLimitOk

`func (o *ClientRateLimitSetting) GetCustomRateLimitOk() (*CustomRateLimitEntity, bool)`

GetCustomRateLimitOk returns a tuple with the CustomRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRateLimit

`func (o *ClientRateLimitSetting) SetCustomRateLimit(v CustomRateLimitEntity)`

SetCustomRateLimit sets CustomRateLimit field to given value.

### HasCustomRateLimit

`func (o *ClientRateLimitSetting) HasCustomRateLimit() bool`

HasCustomRateLimit returns a boolean if a field has been set.

### GetMode

`func (o *ClientRateLimitSetting) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ClientRateLimitSetting) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ClientRateLimitSetting) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetRateLimitProfileId

`func (o *ClientRateLimitSetting) GetRateLimitProfileId() string`

GetRateLimitProfileId returns the RateLimitProfileId field if non-nil, zero value otherwise.

### GetRateLimitProfileIdOk

`func (o *ClientRateLimitSetting) GetRateLimitProfileIdOk() (*string, bool)`

GetRateLimitProfileIdOk returns a tuple with the RateLimitProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitProfileId

`func (o *ClientRateLimitSetting) SetRateLimitProfileId(v string)`

SetRateLimitProfileId sets RateLimitProfileId field to given value.

### HasRateLimitProfileId

`func (o *ClientRateLimitSetting) HasRateLimitProfileId() bool`

HasRateLimitProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


