# SSLVPNLockEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the SSL VPN lock | [optional] 
**Ip** | Pointer to **string** | IP of the SSL VPN lock | [optional] 
**LeftLockTime** | Pointer to **int32** | Left lock time(min) of the SSL VPN lock. | [optional] 
**TotalLockTime** | **int32** | Total lock time of the SSL VPN lock should be within the range of 1–1080(min). | 
**Type** | **int32** | Type of the SSL VPN lock should be a value as follows: 0:username; 1:IP | 
**Username** | Pointer to **string** | Username of the SSL VPN lock | [optional] 

## Methods

### NewSSLVPNLockEntity

`func NewSSLVPNLockEntity(totalLockTime int32, type_ int32, ) *SSLVPNLockEntity`

NewSSLVPNLockEntity instantiates a new SSLVPNLockEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSLVPNLockEntityWithDefaults

`func NewSSLVPNLockEntityWithDefaults() *SSLVPNLockEntity`

NewSSLVPNLockEntityWithDefaults instantiates a new SSLVPNLockEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SSLVPNLockEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SSLVPNLockEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SSLVPNLockEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SSLVPNLockEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *SSLVPNLockEntity) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SSLVPNLockEntity) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SSLVPNLockEntity) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SSLVPNLockEntity) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLeftLockTime

`func (o *SSLVPNLockEntity) GetLeftLockTime() int32`

GetLeftLockTime returns the LeftLockTime field if non-nil, zero value otherwise.

### GetLeftLockTimeOk

`func (o *SSLVPNLockEntity) GetLeftLockTimeOk() (*int32, bool)`

GetLeftLockTimeOk returns a tuple with the LeftLockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftLockTime

`func (o *SSLVPNLockEntity) SetLeftLockTime(v int32)`

SetLeftLockTime sets LeftLockTime field to given value.

### HasLeftLockTime

`func (o *SSLVPNLockEntity) HasLeftLockTime() bool`

HasLeftLockTime returns a boolean if a field has been set.

### GetTotalLockTime

`func (o *SSLVPNLockEntity) GetTotalLockTime() int32`

GetTotalLockTime returns the TotalLockTime field if non-nil, zero value otherwise.

### GetTotalLockTimeOk

`func (o *SSLVPNLockEntity) GetTotalLockTimeOk() (*int32, bool)`

GetTotalLockTimeOk returns a tuple with the TotalLockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLockTime

`func (o *SSLVPNLockEntity) SetTotalLockTime(v int32)`

SetTotalLockTime sets TotalLockTime field to given value.


### GetType

`func (o *SSLVPNLockEntity) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SSLVPNLockEntity) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SSLVPNLockEntity) SetType(v int32)`

SetType sets Type field to given value.


### GetUsername

`func (o *SSLVPNLockEntity) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SSLVPNLockEntity) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SSLVPNLockEntity) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SSLVPNLockEntity) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


