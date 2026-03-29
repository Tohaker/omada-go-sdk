# CreateRadiusUserOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownLimit** | Pointer to **int64** | Build-in Radius profile user downlink traffic limit, unit: MB. DownLimit should be within the range of 1-10485760 | [optional] 
**DownRateLimit** | Pointer to **int64** | Build-in Radius profile user downlink rate limit, unit: Kbps. DownRateLimit should be within the range of 1-10485760 | [optional] 
**MacAddress** | Pointer to **string** | MAC address, required when parameter [type] is 1. Should be a valid MAC address format | [optional] 
**Timeout** | Pointer to **int32** | Session timeout, unit: second | [optional] 
**Type** | **int32** | Type of Build-in Radius profile user, 0: user auth; 1: MAC auth | 
**UpLimit** | Pointer to **int64** | Build-in Radius profile user uplink traffic limit, unit: MB. UpLimit should be within the range of 1-10485760 | [optional] 
**UpRateLimit** | Pointer to **int64** | Build-in Radius profile user uplink rate limit, unit: Kbps. UpRateLimit should be within the range of 1-10485760 | [optional] 
**UserInfo** | Pointer to [**RadiusUserInfoOpenApiVO**](RadiusUserInfoOpenApiVO.md) |  | [optional] 
**VlanId** | Pointer to **int32** | VLAN ID. VlanId should be within the range of 1-4096 | [optional] 

## Methods

### NewCreateRadiusUserOpenApiVO

`func NewCreateRadiusUserOpenApiVO(type_ int32, ) *CreateRadiusUserOpenApiVO`

NewCreateRadiusUserOpenApiVO instantiates a new CreateRadiusUserOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRadiusUserOpenApiVOWithDefaults

`func NewCreateRadiusUserOpenApiVOWithDefaults() *CreateRadiusUserOpenApiVO`

NewCreateRadiusUserOpenApiVOWithDefaults instantiates a new CreateRadiusUserOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownLimit

`func (o *CreateRadiusUserOpenApiVO) GetDownLimit() int64`

GetDownLimit returns the DownLimit field if non-nil, zero value otherwise.

### GetDownLimitOk

`func (o *CreateRadiusUserOpenApiVO) GetDownLimitOk() (*int64, bool)`

GetDownLimitOk returns a tuple with the DownLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimit

`func (o *CreateRadiusUserOpenApiVO) SetDownLimit(v int64)`

SetDownLimit sets DownLimit field to given value.

### HasDownLimit

`func (o *CreateRadiusUserOpenApiVO) HasDownLimit() bool`

HasDownLimit returns a boolean if a field has been set.

### GetDownRateLimit

`func (o *CreateRadiusUserOpenApiVO) GetDownRateLimit() int64`

GetDownRateLimit returns the DownRateLimit field if non-nil, zero value otherwise.

### GetDownRateLimitOk

`func (o *CreateRadiusUserOpenApiVO) GetDownRateLimitOk() (*int64, bool)`

GetDownRateLimitOk returns a tuple with the DownRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownRateLimit

`func (o *CreateRadiusUserOpenApiVO) SetDownRateLimit(v int64)`

SetDownRateLimit sets DownRateLimit field to given value.

### HasDownRateLimit

`func (o *CreateRadiusUserOpenApiVO) HasDownRateLimit() bool`

HasDownRateLimit returns a boolean if a field has been set.

### GetMacAddress

`func (o *CreateRadiusUserOpenApiVO) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *CreateRadiusUserOpenApiVO) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *CreateRadiusUserOpenApiVO) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *CreateRadiusUserOpenApiVO) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetTimeout

`func (o *CreateRadiusUserOpenApiVO) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CreateRadiusUserOpenApiVO) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CreateRadiusUserOpenApiVO) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CreateRadiusUserOpenApiVO) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetType

`func (o *CreateRadiusUserOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateRadiusUserOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateRadiusUserOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.


### GetUpLimit

`func (o *CreateRadiusUserOpenApiVO) GetUpLimit() int64`

GetUpLimit returns the UpLimit field if non-nil, zero value otherwise.

### GetUpLimitOk

`func (o *CreateRadiusUserOpenApiVO) GetUpLimitOk() (*int64, bool)`

GetUpLimitOk returns a tuple with the UpLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimit

`func (o *CreateRadiusUserOpenApiVO) SetUpLimit(v int64)`

SetUpLimit sets UpLimit field to given value.

### HasUpLimit

`func (o *CreateRadiusUserOpenApiVO) HasUpLimit() bool`

HasUpLimit returns a boolean if a field has been set.

### GetUpRateLimit

`func (o *CreateRadiusUserOpenApiVO) GetUpRateLimit() int64`

GetUpRateLimit returns the UpRateLimit field if non-nil, zero value otherwise.

### GetUpRateLimitOk

`func (o *CreateRadiusUserOpenApiVO) GetUpRateLimitOk() (*int64, bool)`

GetUpRateLimitOk returns a tuple with the UpRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpRateLimit

`func (o *CreateRadiusUserOpenApiVO) SetUpRateLimit(v int64)`

SetUpRateLimit sets UpRateLimit field to given value.

### HasUpRateLimit

`func (o *CreateRadiusUserOpenApiVO) HasUpRateLimit() bool`

HasUpRateLimit returns a boolean if a field has been set.

### GetUserInfo

`func (o *CreateRadiusUserOpenApiVO) GetUserInfo() RadiusUserInfoOpenApiVO`

GetUserInfo returns the UserInfo field if non-nil, zero value otherwise.

### GetUserInfoOk

`func (o *CreateRadiusUserOpenApiVO) GetUserInfoOk() (*RadiusUserInfoOpenApiVO, bool)`

GetUserInfoOk returns a tuple with the UserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfo

`func (o *CreateRadiusUserOpenApiVO) SetUserInfo(v RadiusUserInfoOpenApiVO)`

SetUserInfo sets UserInfo field to given value.

### HasUserInfo

`func (o *CreateRadiusUserOpenApiVO) HasUserInfo() bool`

HasUserInfo returns a boolean if a field has been set.

### GetVlanId

`func (o *CreateRadiusUserOpenApiVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *CreateRadiusUserOpenApiVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *CreateRadiusUserOpenApiVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *CreateRadiusUserOpenApiVO) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


