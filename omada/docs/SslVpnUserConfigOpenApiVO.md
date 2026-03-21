# SslVpnUserConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConcurrentNumber** | **int32** | Concurrent number of the SSL VPN user should be within the range of 1–100 | 
**GroupId** | Pointer to **string** | Group ID of the SSL VPN user. User group can be created using &#39;Create SSL VPN user group&#39; interface, and User Group ID can be obtained from &#39;Get user group list for SSL VPN server&#39; interface. | [optional] 
**Name** | **string** | Name of the SSL VPN user should contain 1 to 20 characters. | 
**Password** | **string** | Password of the SSL VPN user should contain 1 to 64 characters. | 
**Status** | **bool** | Status of the SSL VPN user | 
**Validity** | **string** | Validity of the SSL VPN user. The format is Month/Day/Year, for example 08/20/2022 | 

## Methods

### NewSslVpnUserConfigOpenApiVO

`func NewSslVpnUserConfigOpenApiVO(concurrentNumber int32, name string, password string, status bool, validity string, ) *SslVpnUserConfigOpenApiVO`

NewSslVpnUserConfigOpenApiVO instantiates a new SslVpnUserConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslVpnUserConfigOpenApiVOWithDefaults

`func NewSslVpnUserConfigOpenApiVOWithDefaults() *SslVpnUserConfigOpenApiVO`

NewSslVpnUserConfigOpenApiVOWithDefaults instantiates a new SslVpnUserConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConcurrentNumber

`func (o *SslVpnUserConfigOpenApiVO) GetConcurrentNumber() int32`

GetConcurrentNumber returns the ConcurrentNumber field if non-nil, zero value otherwise.

### GetConcurrentNumberOk

`func (o *SslVpnUserConfigOpenApiVO) GetConcurrentNumberOk() (*int32, bool)`

GetConcurrentNumberOk returns a tuple with the ConcurrentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentNumber

`func (o *SslVpnUserConfigOpenApiVO) SetConcurrentNumber(v int32)`

SetConcurrentNumber sets ConcurrentNumber field to given value.


### GetGroupId

`func (o *SslVpnUserConfigOpenApiVO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SslVpnUserConfigOpenApiVO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SslVpnUserConfigOpenApiVO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *SslVpnUserConfigOpenApiVO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *SslVpnUserConfigOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SslVpnUserConfigOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SslVpnUserConfigOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *SslVpnUserConfigOpenApiVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SslVpnUserConfigOpenApiVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SslVpnUserConfigOpenApiVO) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetStatus

`func (o *SslVpnUserConfigOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SslVpnUserConfigOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SslVpnUserConfigOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetValidity

`func (o *SslVpnUserConfigOpenApiVO) GetValidity() string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *SslVpnUserConfigOpenApiVO) GetValidityOk() (*string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *SslVpnUserConfigOpenApiVO) SetValidity(v string)`

SetValidity sets Validity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


