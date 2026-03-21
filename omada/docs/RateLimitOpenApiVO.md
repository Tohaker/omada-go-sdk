# RateLimitOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomRateLimit** | Pointer to [**CustomRateLimitOpenApiVO**](CustomRateLimitOpenApiVO.md) |  | [optional] 
**Mode** | **int32** | Mode of configure rate limit should be a value as follows: 0: customRateLimit; 1: rateLimitProfileId. | 
**RateLimitProfileId** | Pointer to **string** | This field represents Rate limit profile ID. Rate limit profile can be created using &#39;Create rate limit profile&#39; interface, and Rate limit profile ID can be obtained from &#39;Get rate limit profile list&#39; interface | [optional] 

## Methods

### NewRateLimitOpenApiVO

`func NewRateLimitOpenApiVO(mode int32, ) *RateLimitOpenApiVO`

NewRateLimitOpenApiVO instantiates a new RateLimitOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitOpenApiVOWithDefaults

`func NewRateLimitOpenApiVOWithDefaults() *RateLimitOpenApiVO`

NewRateLimitOpenApiVOWithDefaults instantiates a new RateLimitOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomRateLimit

`func (o *RateLimitOpenApiVO) GetCustomRateLimit() CustomRateLimitOpenApiVO`

GetCustomRateLimit returns the CustomRateLimit field if non-nil, zero value otherwise.

### GetCustomRateLimitOk

`func (o *RateLimitOpenApiVO) GetCustomRateLimitOk() (*CustomRateLimitOpenApiVO, bool)`

GetCustomRateLimitOk returns a tuple with the CustomRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRateLimit

`func (o *RateLimitOpenApiVO) SetCustomRateLimit(v CustomRateLimitOpenApiVO)`

SetCustomRateLimit sets CustomRateLimit field to given value.

### HasCustomRateLimit

`func (o *RateLimitOpenApiVO) HasCustomRateLimit() bool`

HasCustomRateLimit returns a boolean if a field has been set.

### GetMode

`func (o *RateLimitOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RateLimitOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RateLimitOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetRateLimitProfileId

`func (o *RateLimitOpenApiVO) GetRateLimitProfileId() string`

GetRateLimitProfileId returns the RateLimitProfileId field if non-nil, zero value otherwise.

### GetRateLimitProfileIdOk

`func (o *RateLimitOpenApiVO) GetRateLimitProfileIdOk() (*string, bool)`

GetRateLimitProfileIdOk returns a tuple with the RateLimitProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitProfileId

`func (o *RateLimitOpenApiVO) SetRateLimitProfileId(v string)`

SetRateLimitProfileId sets RateLimitProfileId field to given value.

### HasRateLimitProfileId

`func (o *RateLimitOpenApiVO) HasRateLimitProfileId() bool`

HasRateLimitProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


