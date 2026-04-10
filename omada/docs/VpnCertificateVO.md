# VpnCertificateVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedServerAddress** | Pointer to **[]string** |  | [optional] 
**Dns1** | Pointer to **string** |  | [optional] 
**Dns2** | Pointer to **string** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**KeepAlive** | Pointer to **int32** |  | [optional] 
**LocalIp** | Pointer to **string** |  | [optional] 
**OpenVpnMode** | Pointer to **int32** |  | [optional] 
**PreSharedKey** | Pointer to **string** |  | [optional] 
**PrivateKey** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**RemoteIp** | Pointer to **string** |  | [optional] 
**ServerPublicKey** | Pointer to **string** |  | [optional] 
**ServicePort** | Pointer to **int32** |  | [optional] 

## Methods

### NewVpnCertificateVO

`func NewVpnCertificateVO() *VpnCertificateVO`

NewVpnCertificateVO instantiates a new VpnCertificateVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnCertificateVOWithDefaults

`func NewVpnCertificateVOWithDefaults() *VpnCertificateVO`

NewVpnCertificateVOWithDefaults instantiates a new VpnCertificateVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedServerAddress

`func (o *VpnCertificateVO) GetAllowedServerAddress() []string`

GetAllowedServerAddress returns the AllowedServerAddress field if non-nil, zero value otherwise.

### GetAllowedServerAddressOk

`func (o *VpnCertificateVO) GetAllowedServerAddressOk() (*[]string, bool)`

GetAllowedServerAddressOk returns a tuple with the AllowedServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedServerAddress

`func (o *VpnCertificateVO) SetAllowedServerAddress(v []string)`

SetAllowedServerAddress sets AllowedServerAddress field to given value.

### HasAllowedServerAddress

`func (o *VpnCertificateVO) HasAllowedServerAddress() bool`

HasAllowedServerAddress returns a boolean if a field has been set.

### GetDns1

`func (o *VpnCertificateVO) GetDns1() string`

GetDns1 returns the Dns1 field if non-nil, zero value otherwise.

### GetDns1Ok

`func (o *VpnCertificateVO) GetDns1Ok() (*string, bool)`

GetDns1Ok returns a tuple with the Dns1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns1

`func (o *VpnCertificateVO) SetDns1(v string)`

SetDns1 sets Dns1 field to given value.

### HasDns1

`func (o *VpnCertificateVO) HasDns1() bool`

HasDns1 returns a boolean if a field has been set.

### GetDns2

`func (o *VpnCertificateVO) GetDns2() string`

GetDns2 returns the Dns2 field if non-nil, zero value otherwise.

### GetDns2Ok

`func (o *VpnCertificateVO) GetDns2Ok() (*string, bool)`

GetDns2Ok returns a tuple with the Dns2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns2

`func (o *VpnCertificateVO) SetDns2(v string)`

SetDns2 sets Dns2 field to given value.

### HasDns2

`func (o *VpnCertificateVO) HasDns2() bool`

HasDns2 returns a boolean if a field has been set.

### GetFileName

`func (o *VpnCertificateVO) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *VpnCertificateVO) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *VpnCertificateVO) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *VpnCertificateVO) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetId

`func (o *VpnCertificateVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpnCertificateVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpnCertificateVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VpnCertificateVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeepAlive

`func (o *VpnCertificateVO) GetKeepAlive() int32`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *VpnCertificateVO) GetKeepAliveOk() (*int32, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *VpnCertificateVO) SetKeepAlive(v int32)`

SetKeepAlive sets KeepAlive field to given value.

### HasKeepAlive

`func (o *VpnCertificateVO) HasKeepAlive() bool`

HasKeepAlive returns a boolean if a field has been set.

### GetLocalIp

`func (o *VpnCertificateVO) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *VpnCertificateVO) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *VpnCertificateVO) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.

### HasLocalIp

`func (o *VpnCertificateVO) HasLocalIp() bool`

HasLocalIp returns a boolean if a field has been set.

### GetOpenVpnMode

`func (o *VpnCertificateVO) GetOpenVpnMode() int32`

GetOpenVpnMode returns the OpenVpnMode field if non-nil, zero value otherwise.

### GetOpenVpnModeOk

`func (o *VpnCertificateVO) GetOpenVpnModeOk() (*int32, bool)`

GetOpenVpnModeOk returns a tuple with the OpenVpnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenVpnMode

`func (o *VpnCertificateVO) SetOpenVpnMode(v int32)`

SetOpenVpnMode sets OpenVpnMode field to given value.

### HasOpenVpnMode

`func (o *VpnCertificateVO) HasOpenVpnMode() bool`

HasOpenVpnMode returns a boolean if a field has been set.

### GetPreSharedKey

`func (o *VpnCertificateVO) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *VpnCertificateVO) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *VpnCertificateVO) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.

### HasPreSharedKey

`func (o *VpnCertificateVO) HasPreSharedKey() bool`

HasPreSharedKey returns a boolean if a field has been set.

### GetPrivateKey

`func (o *VpnCertificateVO) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *VpnCertificateVO) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *VpnCertificateVO) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *VpnCertificateVO) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPublicKey

`func (o *VpnCertificateVO) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *VpnCertificateVO) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *VpnCertificateVO) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *VpnCertificateVO) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetRemoteIp

`func (o *VpnCertificateVO) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *VpnCertificateVO) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *VpnCertificateVO) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *VpnCertificateVO) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetServerPublicKey

`func (o *VpnCertificateVO) GetServerPublicKey() string`

GetServerPublicKey returns the ServerPublicKey field if non-nil, zero value otherwise.

### GetServerPublicKeyOk

`func (o *VpnCertificateVO) GetServerPublicKeyOk() (*string, bool)`

GetServerPublicKeyOk returns a tuple with the ServerPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPublicKey

`func (o *VpnCertificateVO) SetServerPublicKey(v string)`

SetServerPublicKey sets ServerPublicKey field to given value.

### HasServerPublicKey

`func (o *VpnCertificateVO) HasServerPublicKey() bool`

HasServerPublicKey returns a boolean if a field has been set.

### GetServicePort

`func (o *VpnCertificateVO) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *VpnCertificateVO) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *VpnCertificateVO) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *VpnCertificateVO) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


