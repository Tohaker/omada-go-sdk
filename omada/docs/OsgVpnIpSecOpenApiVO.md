# OsgVpnIpSecOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AhAuthentication** | Pointer to **string** | AH Authentication Algorithm. | [optional] 
**Direction** | Pointer to **string** | SA direction, in/out. | [optional] 
**EspAuthentication** | Pointer to **string** | ESP Authentication Algorithm. | [optional] 
**EspEncryption** | Pointer to **string** | ESP Encryption Algorithm. | [optional] 
**LocalPeerIp** | Pointer to **string** | IP address of the local peer. | [optional] 
**LocalSa** | Pointer to **string** | Local network segment of SA cover. | [optional] 
**Name** | Pointer to **string** | IPsec name. | [optional] 
**Protocol** | Pointer to **string** | SA Authentication protocol and Encryption protocol. | [optional] 
**RemotePeerIp** | Pointer to **string** | IP address of the remote peer. | [optional] 
**RemoteSa** | Pointer to **string** | Remote network segment of SA cover. | [optional] 
**Spi** | Pointer to **int64** | Security Parameter Index of SA. | [optional] 
**VpnId** | Pointer to **int32** | VPN Item Id. | [optional] 

## Methods

### NewOsgVpnIpSecOpenApiVO

`func NewOsgVpnIpSecOpenApiVO() *OsgVpnIpSecOpenApiVO`

NewOsgVpnIpSecOpenApiVO instantiates a new OsgVpnIpSecOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgVpnIpSecOpenApiVOWithDefaults

`func NewOsgVpnIpSecOpenApiVOWithDefaults() *OsgVpnIpSecOpenApiVO`

NewOsgVpnIpSecOpenApiVOWithDefaults instantiates a new OsgVpnIpSecOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAhAuthentication

`func (o *OsgVpnIpSecOpenApiVO) GetAhAuthentication() string`

GetAhAuthentication returns the AhAuthentication field if non-nil, zero value otherwise.

### GetAhAuthenticationOk

`func (o *OsgVpnIpSecOpenApiVO) GetAhAuthenticationOk() (*string, bool)`

GetAhAuthenticationOk returns a tuple with the AhAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAhAuthentication

`func (o *OsgVpnIpSecOpenApiVO) SetAhAuthentication(v string)`

SetAhAuthentication sets AhAuthentication field to given value.

### HasAhAuthentication

`func (o *OsgVpnIpSecOpenApiVO) HasAhAuthentication() bool`

HasAhAuthentication returns a boolean if a field has been set.

### GetDirection

`func (o *OsgVpnIpSecOpenApiVO) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *OsgVpnIpSecOpenApiVO) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *OsgVpnIpSecOpenApiVO) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *OsgVpnIpSecOpenApiVO) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetEspAuthentication

`func (o *OsgVpnIpSecOpenApiVO) GetEspAuthentication() string`

GetEspAuthentication returns the EspAuthentication field if non-nil, zero value otherwise.

### GetEspAuthenticationOk

`func (o *OsgVpnIpSecOpenApiVO) GetEspAuthenticationOk() (*string, bool)`

GetEspAuthenticationOk returns a tuple with the EspAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEspAuthentication

`func (o *OsgVpnIpSecOpenApiVO) SetEspAuthentication(v string)`

SetEspAuthentication sets EspAuthentication field to given value.

### HasEspAuthentication

`func (o *OsgVpnIpSecOpenApiVO) HasEspAuthentication() bool`

HasEspAuthentication returns a boolean if a field has been set.

### GetEspEncryption

`func (o *OsgVpnIpSecOpenApiVO) GetEspEncryption() string`

GetEspEncryption returns the EspEncryption field if non-nil, zero value otherwise.

### GetEspEncryptionOk

`func (o *OsgVpnIpSecOpenApiVO) GetEspEncryptionOk() (*string, bool)`

GetEspEncryptionOk returns a tuple with the EspEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEspEncryption

`func (o *OsgVpnIpSecOpenApiVO) SetEspEncryption(v string)`

SetEspEncryption sets EspEncryption field to given value.

### HasEspEncryption

`func (o *OsgVpnIpSecOpenApiVO) HasEspEncryption() bool`

HasEspEncryption returns a boolean if a field has been set.

### GetLocalPeerIp

`func (o *OsgVpnIpSecOpenApiVO) GetLocalPeerIp() string`

GetLocalPeerIp returns the LocalPeerIp field if non-nil, zero value otherwise.

### GetLocalPeerIpOk

`func (o *OsgVpnIpSecOpenApiVO) GetLocalPeerIpOk() (*string, bool)`

GetLocalPeerIpOk returns a tuple with the LocalPeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPeerIp

`func (o *OsgVpnIpSecOpenApiVO) SetLocalPeerIp(v string)`

SetLocalPeerIp sets LocalPeerIp field to given value.

### HasLocalPeerIp

`func (o *OsgVpnIpSecOpenApiVO) HasLocalPeerIp() bool`

HasLocalPeerIp returns a boolean if a field has been set.

### GetLocalSa

`func (o *OsgVpnIpSecOpenApiVO) GetLocalSa() string`

GetLocalSa returns the LocalSa field if non-nil, zero value otherwise.

### GetLocalSaOk

`func (o *OsgVpnIpSecOpenApiVO) GetLocalSaOk() (*string, bool)`

GetLocalSaOk returns a tuple with the LocalSa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSa

`func (o *OsgVpnIpSecOpenApiVO) SetLocalSa(v string)`

SetLocalSa sets LocalSa field to given value.

### HasLocalSa

`func (o *OsgVpnIpSecOpenApiVO) HasLocalSa() bool`

HasLocalSa returns a boolean if a field has been set.

### GetName

`func (o *OsgVpnIpSecOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsgVpnIpSecOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsgVpnIpSecOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsgVpnIpSecOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProtocol

`func (o *OsgVpnIpSecOpenApiVO) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *OsgVpnIpSecOpenApiVO) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *OsgVpnIpSecOpenApiVO) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *OsgVpnIpSecOpenApiVO) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRemotePeerIp

`func (o *OsgVpnIpSecOpenApiVO) GetRemotePeerIp() string`

GetRemotePeerIp returns the RemotePeerIp field if non-nil, zero value otherwise.

### GetRemotePeerIpOk

`func (o *OsgVpnIpSecOpenApiVO) GetRemotePeerIpOk() (*string, bool)`

GetRemotePeerIpOk returns a tuple with the RemotePeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePeerIp

`func (o *OsgVpnIpSecOpenApiVO) SetRemotePeerIp(v string)`

SetRemotePeerIp sets RemotePeerIp field to given value.

### HasRemotePeerIp

`func (o *OsgVpnIpSecOpenApiVO) HasRemotePeerIp() bool`

HasRemotePeerIp returns a boolean if a field has been set.

### GetRemoteSa

`func (o *OsgVpnIpSecOpenApiVO) GetRemoteSa() string`

GetRemoteSa returns the RemoteSa field if non-nil, zero value otherwise.

### GetRemoteSaOk

`func (o *OsgVpnIpSecOpenApiVO) GetRemoteSaOk() (*string, bool)`

GetRemoteSaOk returns a tuple with the RemoteSa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSa

`func (o *OsgVpnIpSecOpenApiVO) SetRemoteSa(v string)`

SetRemoteSa sets RemoteSa field to given value.

### HasRemoteSa

`func (o *OsgVpnIpSecOpenApiVO) HasRemoteSa() bool`

HasRemoteSa returns a boolean if a field has been set.

### GetSpi

`func (o *OsgVpnIpSecOpenApiVO) GetSpi() int64`

GetSpi returns the Spi field if non-nil, zero value otherwise.

### GetSpiOk

`func (o *OsgVpnIpSecOpenApiVO) GetSpiOk() (*int64, bool)`

GetSpiOk returns a tuple with the Spi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpi

`func (o *OsgVpnIpSecOpenApiVO) SetSpi(v int64)`

SetSpi sets Spi field to given value.

### HasSpi

`func (o *OsgVpnIpSecOpenApiVO) HasSpi() bool`

HasSpi returns a boolean if a field has been set.

### GetVpnId

`func (o *OsgVpnIpSecOpenApiVO) GetVpnId() int32`

GetVpnId returns the VpnId field if non-nil, zero value otherwise.

### GetVpnIdOk

`func (o *OsgVpnIpSecOpenApiVO) GetVpnIdOk() (*int32, bool)`

GetVpnIdOk returns a tuple with the VpnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnId

`func (o *OsgVpnIpSecOpenApiVO) SetVpnId(v int32)`

SetVpnId sets VpnId field to given value.

### HasVpnId

`func (o *OsgVpnIpSecOpenApiVO) HasVpnId() bool`

HasVpnId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


