# RadiusUserOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownLimit** | Pointer to **int64** | Build-in Radius profile user downlink traffic limit, unit: MB | [optional] 
**DownRateLimit** | Pointer to **int64** | Build-in Radius profile user downlink rate limit, unit: Kbps | [optional] 
**Password** | Pointer to **string** | Build-in Radius profile user password, when parameter [type] is 1, [password] is the MAC address | [optional] 
**Timeout** | Pointer to **int32** | Session timeout | [optional] 
**Type** | Pointer to **int32** | Type of Build-in Radius profile user, 0: user auth; 1: MAC auth | [optional] 
**UpLimit** | Pointer to **int64** | Build-in Radius profile user uplink traffic limit, unit: MB | [optional] 
**UpRateLimit** | Pointer to **int64** | Build-in Radius profile user uplink rate limit, unit: Kbps | [optional] 
**UserId** | Pointer to **string** | Build-in Radius profile user ID | [optional] 
**Username** | Pointer to **string** | Build-in Radius profile user name, when parameter [type] is 1, [username] is the MAC address | [optional] 
**VlanId** | Pointer to **int32** | VLAN ID, from 1 to 4096 | [optional] 

## Methods

### NewRadiusUserOpenApiVO

`func NewRadiusUserOpenApiVO() *RadiusUserOpenApiVO`

NewRadiusUserOpenApiVO instantiates a new RadiusUserOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusUserOpenApiVOWithDefaults

`func NewRadiusUserOpenApiVOWithDefaults() *RadiusUserOpenApiVO`

NewRadiusUserOpenApiVOWithDefaults instantiates a new RadiusUserOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownLimit

`func (o *RadiusUserOpenApiVO) GetDownLimit() int64`

GetDownLimit returns the DownLimit field if non-nil, zero value otherwise.

### GetDownLimitOk

`func (o *RadiusUserOpenApiVO) GetDownLimitOk() (*int64, bool)`

GetDownLimitOk returns a tuple with the DownLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimit

`func (o *RadiusUserOpenApiVO) SetDownLimit(v int64)`

SetDownLimit sets DownLimit field to given value.

### HasDownLimit

`func (o *RadiusUserOpenApiVO) HasDownLimit() bool`

HasDownLimit returns a boolean if a field has been set.

### GetDownRateLimit

`func (o *RadiusUserOpenApiVO) GetDownRateLimit() int64`

GetDownRateLimit returns the DownRateLimit field if non-nil, zero value otherwise.

### GetDownRateLimitOk

`func (o *RadiusUserOpenApiVO) GetDownRateLimitOk() (*int64, bool)`

GetDownRateLimitOk returns a tuple with the DownRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownRateLimit

`func (o *RadiusUserOpenApiVO) SetDownRateLimit(v int64)`

SetDownRateLimit sets DownRateLimit field to given value.

### HasDownRateLimit

`func (o *RadiusUserOpenApiVO) HasDownRateLimit() bool`

HasDownRateLimit returns a boolean if a field has been set.

### GetPassword

`func (o *RadiusUserOpenApiVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RadiusUserOpenApiVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RadiusUserOpenApiVO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RadiusUserOpenApiVO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTimeout

`func (o *RadiusUserOpenApiVO) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *RadiusUserOpenApiVO) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *RadiusUserOpenApiVO) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *RadiusUserOpenApiVO) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetType

`func (o *RadiusUserOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RadiusUserOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RadiusUserOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *RadiusUserOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpLimit

`func (o *RadiusUserOpenApiVO) GetUpLimit() int64`

GetUpLimit returns the UpLimit field if non-nil, zero value otherwise.

### GetUpLimitOk

`func (o *RadiusUserOpenApiVO) GetUpLimitOk() (*int64, bool)`

GetUpLimitOk returns a tuple with the UpLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimit

`func (o *RadiusUserOpenApiVO) SetUpLimit(v int64)`

SetUpLimit sets UpLimit field to given value.

### HasUpLimit

`func (o *RadiusUserOpenApiVO) HasUpLimit() bool`

HasUpLimit returns a boolean if a field has been set.

### GetUpRateLimit

`func (o *RadiusUserOpenApiVO) GetUpRateLimit() int64`

GetUpRateLimit returns the UpRateLimit field if non-nil, zero value otherwise.

### GetUpRateLimitOk

`func (o *RadiusUserOpenApiVO) GetUpRateLimitOk() (*int64, bool)`

GetUpRateLimitOk returns a tuple with the UpRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpRateLimit

`func (o *RadiusUserOpenApiVO) SetUpRateLimit(v int64)`

SetUpRateLimit sets UpRateLimit field to given value.

### HasUpRateLimit

`func (o *RadiusUserOpenApiVO) HasUpRateLimit() bool`

HasUpRateLimit returns a boolean if a field has been set.

### GetUserId

`func (o *RadiusUserOpenApiVO) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RadiusUserOpenApiVO) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RadiusUserOpenApiVO) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *RadiusUserOpenApiVO) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *RadiusUserOpenApiVO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RadiusUserOpenApiVO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RadiusUserOpenApiVO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RadiusUserOpenApiVO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVlanId

`func (o *RadiusUserOpenApiVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *RadiusUserOpenApiVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *RadiusUserOpenApiVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *RadiusUserOpenApiVO) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


