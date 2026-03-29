# SslVpnLockCreateAndModifyOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | IP of the SSL VPN lock. It is required when parameter [type] is 1. | [optional] 
**TotalLockTime** | **int32** | Total lock time of the SSL VPN lock should be within the range of 1–1080(min). | 
**Type** | **int32** | Type of the SSL VPN lock should be a value as follows: 0:username; 1:IP | 
**Username** | Pointer to **string** | Username of the SSL VPN lock. It is required when parameter [type] is 0. | [optional] 

## Methods

### NewSslVpnLockCreateAndModifyOpenApiVO

`func NewSslVpnLockCreateAndModifyOpenApiVO(totalLockTime int32, type_ int32, ) *SslVpnLockCreateAndModifyOpenApiVO`

NewSslVpnLockCreateAndModifyOpenApiVO instantiates a new SslVpnLockCreateAndModifyOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslVpnLockCreateAndModifyOpenApiVOWithDefaults

`func NewSslVpnLockCreateAndModifyOpenApiVOWithDefaults() *SslVpnLockCreateAndModifyOpenApiVO`

NewSslVpnLockCreateAndModifyOpenApiVOWithDefaults instantiates a new SslVpnLockCreateAndModifyOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *SslVpnLockCreateAndModifyOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SslVpnLockCreateAndModifyOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SslVpnLockCreateAndModifyOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SslVpnLockCreateAndModifyOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetTotalLockTime

`func (o *SslVpnLockCreateAndModifyOpenApiVO) GetTotalLockTime() int32`

GetTotalLockTime returns the TotalLockTime field if non-nil, zero value otherwise.

### GetTotalLockTimeOk

`func (o *SslVpnLockCreateAndModifyOpenApiVO) GetTotalLockTimeOk() (*int32, bool)`

GetTotalLockTimeOk returns a tuple with the TotalLockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLockTime

`func (o *SslVpnLockCreateAndModifyOpenApiVO) SetTotalLockTime(v int32)`

SetTotalLockTime sets TotalLockTime field to given value.


### GetType

`func (o *SslVpnLockCreateAndModifyOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SslVpnLockCreateAndModifyOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SslVpnLockCreateAndModifyOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.


### GetUsername

`func (o *SslVpnLockCreateAndModifyOpenApiVO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SslVpnLockCreateAndModifyOpenApiVO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SslVpnLockCreateAndModifyOpenApiVO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SslVpnLockCreateAndModifyOpenApiVO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


