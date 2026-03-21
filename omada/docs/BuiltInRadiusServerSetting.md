# BuiltInRadiusServerSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthServerPort** | **int32** | Server port. | 
**CustomAddress** | Pointer to **string** | Customize server address, required when parameter [serverAddressType] is 1. | [optional] 
**Enable** | **bool** | Whether the Built-In RADIUS server is enabled. | 
**RadiusServerCertificate** | Pointer to [**RADIUSServerCertificateSetting**](RADIUSServerCertificateSetting.md) |  | [optional] 
**Secret** | **string** | Server secrect. | 
**ServerAddressType** | **int32** | Built In RADIUS server address type, 0: Auto generated, 1: Custom | 
**TunnelReplyEnable** | **bool** | Whether the Tunneled Reply is enabled. | 

## Methods

### NewBuiltInRadiusServerSetting

`func NewBuiltInRadiusServerSetting(authServerPort int32, enable bool, secret string, serverAddressType int32, tunnelReplyEnable bool, ) *BuiltInRadiusServerSetting`

NewBuiltInRadiusServerSetting instantiates a new BuiltInRadiusServerSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuiltInRadiusServerSettingWithDefaults

`func NewBuiltInRadiusServerSettingWithDefaults() *BuiltInRadiusServerSetting`

NewBuiltInRadiusServerSettingWithDefaults instantiates a new BuiltInRadiusServerSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthServerPort

`func (o *BuiltInRadiusServerSetting) GetAuthServerPort() int32`

GetAuthServerPort returns the AuthServerPort field if non-nil, zero value otherwise.

### GetAuthServerPortOk

`func (o *BuiltInRadiusServerSetting) GetAuthServerPortOk() (*int32, bool)`

GetAuthServerPortOk returns a tuple with the AuthServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServerPort

`func (o *BuiltInRadiusServerSetting) SetAuthServerPort(v int32)`

SetAuthServerPort sets AuthServerPort field to given value.


### GetCustomAddress

`func (o *BuiltInRadiusServerSetting) GetCustomAddress() string`

GetCustomAddress returns the CustomAddress field if non-nil, zero value otherwise.

### GetCustomAddressOk

`func (o *BuiltInRadiusServerSetting) GetCustomAddressOk() (*string, bool)`

GetCustomAddressOk returns a tuple with the CustomAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAddress

`func (o *BuiltInRadiusServerSetting) SetCustomAddress(v string)`

SetCustomAddress sets CustomAddress field to given value.

### HasCustomAddress

`func (o *BuiltInRadiusServerSetting) HasCustomAddress() bool`

HasCustomAddress returns a boolean if a field has been set.

### GetEnable

`func (o *BuiltInRadiusServerSetting) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *BuiltInRadiusServerSetting) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *BuiltInRadiusServerSetting) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetRadiusServerCertificate

`func (o *BuiltInRadiusServerSetting) GetRadiusServerCertificate() RADIUSServerCertificateSetting`

GetRadiusServerCertificate returns the RadiusServerCertificate field if non-nil, zero value otherwise.

### GetRadiusServerCertificateOk

`func (o *BuiltInRadiusServerSetting) GetRadiusServerCertificateOk() (*RADIUSServerCertificateSetting, bool)`

GetRadiusServerCertificateOk returns a tuple with the RadiusServerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServerCertificate

`func (o *BuiltInRadiusServerSetting) SetRadiusServerCertificate(v RADIUSServerCertificateSetting)`

SetRadiusServerCertificate sets RadiusServerCertificate field to given value.

### HasRadiusServerCertificate

`func (o *BuiltInRadiusServerSetting) HasRadiusServerCertificate() bool`

HasRadiusServerCertificate returns a boolean if a field has been set.

### GetSecret

`func (o *BuiltInRadiusServerSetting) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *BuiltInRadiusServerSetting) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *BuiltInRadiusServerSetting) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetServerAddressType

`func (o *BuiltInRadiusServerSetting) GetServerAddressType() int32`

GetServerAddressType returns the ServerAddressType field if non-nil, zero value otherwise.

### GetServerAddressTypeOk

`func (o *BuiltInRadiusServerSetting) GetServerAddressTypeOk() (*int32, bool)`

GetServerAddressTypeOk returns a tuple with the ServerAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddressType

`func (o *BuiltInRadiusServerSetting) SetServerAddressType(v int32)`

SetServerAddressType sets ServerAddressType field to given value.


### GetTunnelReplyEnable

`func (o *BuiltInRadiusServerSetting) GetTunnelReplyEnable() bool`

GetTunnelReplyEnable returns the TunnelReplyEnable field if non-nil, zero value otherwise.

### GetTunnelReplyEnableOk

`func (o *BuiltInRadiusServerSetting) GetTunnelReplyEnableOk() (*bool, bool)`

GetTunnelReplyEnableOk returns a tuple with the TunnelReplyEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelReplyEnable

`func (o *BuiltInRadiusServerSetting) SetTunnelReplyEnable(v bool)`

SetTunnelReplyEnable sets TunnelReplyEnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


