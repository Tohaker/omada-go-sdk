# SslVpnUserEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available of the SSL VPN user | [optional] 
**ConcurrentNumber** | **int32** | Concurrent number of the SSL VPN user | 
**GroupId** | Pointer to **string** | Group ID of the SSL VPN user. User group can be created using &#39;Create SSL VPN user group&#39; interface, and User Group ID can be obtained from &#39;Get user group list for SSL VPN server&#39; interface. | [optional] 
**GroupName** | Pointer to **string** | Group name of the SSL VPN user | [optional] 
**Id** | Pointer to **string** | ID of the SSL VPN user | [optional] 
**Name** | **string** | Name of the SSL VPN user | 
**Password** | **string** | Password of the SSL VPN user | 
**Status** | **bool** | Status of the SSL VPN user | 
**Validity** | **string** | Validity of the SSL VPN user | 

## Methods

### NewSslVpnUserEntity

`func NewSslVpnUserEntity(concurrentNumber int32, name string, password string, status bool, validity string, ) *SslVpnUserEntity`

NewSslVpnUserEntity instantiates a new SslVpnUserEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslVpnUserEntityWithDefaults

`func NewSslVpnUserEntityWithDefaults() *SslVpnUserEntity`

NewSslVpnUserEntityWithDefaults instantiates a new SslVpnUserEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *SslVpnUserEntity) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *SslVpnUserEntity) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *SslVpnUserEntity) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *SslVpnUserEntity) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetConcurrentNumber

`func (o *SslVpnUserEntity) GetConcurrentNumber() int32`

GetConcurrentNumber returns the ConcurrentNumber field if non-nil, zero value otherwise.

### GetConcurrentNumberOk

`func (o *SslVpnUserEntity) GetConcurrentNumberOk() (*int32, bool)`

GetConcurrentNumberOk returns a tuple with the ConcurrentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentNumber

`func (o *SslVpnUserEntity) SetConcurrentNumber(v int32)`

SetConcurrentNumber sets ConcurrentNumber field to given value.


### GetGroupId

`func (o *SslVpnUserEntity) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SslVpnUserEntity) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SslVpnUserEntity) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *SslVpnUserEntity) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *SslVpnUserEntity) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *SslVpnUserEntity) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *SslVpnUserEntity) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *SslVpnUserEntity) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetId

`func (o *SslVpnUserEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SslVpnUserEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SslVpnUserEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SslVpnUserEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SslVpnUserEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SslVpnUserEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SslVpnUserEntity) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *SslVpnUserEntity) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SslVpnUserEntity) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SslVpnUserEntity) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetStatus

`func (o *SslVpnUserEntity) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SslVpnUserEntity) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SslVpnUserEntity) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetValidity

`func (o *SslVpnUserEntity) GetValidity() string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *SslVpnUserEntity) GetValidityOk() (*string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *SslVpnUserEntity) SetValidity(v string)`

SetValidity sets Validity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


