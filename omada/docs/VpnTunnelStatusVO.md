# VpnTunnelStatusVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AhAuthentication** | Pointer to **string** |  | [optional] 
**AuthType** | Pointer to **int32** |  | [optional] 
**ClientMode** | Pointer to **int32** |  | [optional] 
**ConnectedNum** | Pointer to **int64** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**DisconnectedNum** | Pointer to **int64** |  | [optional] 
**Dns** | Pointer to **string** |  | [optional] 
**DownBytes** | Pointer to **int64** |  | [optional] 
**DownPkts** | Pointer to **int64** |  | [optional] 
**EspAuthentication** | Pointer to **string** |  | [optional] 
**EspEncryption** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InterfaceName** | Pointer to **string** |  | [optional] 
**IpPool** | Pointer to **string** |  | [optional] 
**IpPoolEnd** | Pointer to **string** |  | [optional] 
**IpPoolStart** | Pointer to **string** |  | [optional] 
**IpPoolType** | Pointer to **int32** |  | [optional] 
**LocalIp** | Pointer to **string** |  | [optional] 
**LocalPeerIp** | Pointer to **string** |  | [optional] 
**LocalSa** | Pointer to **string** |  | [optional] 
**LoginIp** | Pointer to **string** |  | [optional] 
**LoginTime** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**RemoteIp** | Pointer to **string** |  | [optional] 
**RemotePeerIp** | Pointer to **string** |  | [optional] 
**RemoteSa** | Pointer to **string** |  | [optional] 
**ServerType** | Pointer to **int32** |  | [optional] 
**ServicePort** | Pointer to **int32** |  | [optional] 
**Spi** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**TotalRemoteNum** | Pointer to **int64** |  | [optional] 
**UpBytes** | Pointer to **int64** |  | [optional] 
**UpPkts** | Pointer to **int64** |  | [optional] 
**Uptime** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**VirtualIp** | Pointer to **string** |  | [optional] 
**VpnId** | Pointer to **string** |  | [optional] 
**VpnType** | Pointer to **int32** |  | [optional] 

## Methods

### NewVpnTunnelStatusVO

`func NewVpnTunnelStatusVO() *VpnTunnelStatusVO`

NewVpnTunnelStatusVO instantiates a new VpnTunnelStatusVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnTunnelStatusVOWithDefaults

`func NewVpnTunnelStatusVOWithDefaults() *VpnTunnelStatusVO`

NewVpnTunnelStatusVOWithDefaults instantiates a new VpnTunnelStatusVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAhAuthentication

`func (o *VpnTunnelStatusVO) GetAhAuthentication() string`

GetAhAuthentication returns the AhAuthentication field if non-nil, zero value otherwise.

### GetAhAuthenticationOk

`func (o *VpnTunnelStatusVO) GetAhAuthenticationOk() (*string, bool)`

GetAhAuthenticationOk returns a tuple with the AhAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAhAuthentication

`func (o *VpnTunnelStatusVO) SetAhAuthentication(v string)`

SetAhAuthentication sets AhAuthentication field to given value.

### HasAhAuthentication

`func (o *VpnTunnelStatusVO) HasAhAuthentication() bool`

HasAhAuthentication returns a boolean if a field has been set.

### GetAuthType

`func (o *VpnTunnelStatusVO) GetAuthType() int32`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *VpnTunnelStatusVO) GetAuthTypeOk() (*int32, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *VpnTunnelStatusVO) SetAuthType(v int32)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *VpnTunnelStatusVO) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetClientMode

`func (o *VpnTunnelStatusVO) GetClientMode() int32`

GetClientMode returns the ClientMode field if non-nil, zero value otherwise.

### GetClientModeOk

`func (o *VpnTunnelStatusVO) GetClientModeOk() (*int32, bool)`

GetClientModeOk returns a tuple with the ClientMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMode

`func (o *VpnTunnelStatusVO) SetClientMode(v int32)`

SetClientMode sets ClientMode field to given value.

### HasClientMode

`func (o *VpnTunnelStatusVO) HasClientMode() bool`

HasClientMode returns a boolean if a field has been set.

### GetConnectedNum

`func (o *VpnTunnelStatusVO) GetConnectedNum() int64`

GetConnectedNum returns the ConnectedNum field if non-nil, zero value otherwise.

### GetConnectedNumOk

`func (o *VpnTunnelStatusVO) GetConnectedNumOk() (*int64, bool)`

GetConnectedNumOk returns a tuple with the ConnectedNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedNum

`func (o *VpnTunnelStatusVO) SetConnectedNum(v int64)`

SetConnectedNum sets ConnectedNum field to given value.

### HasConnectedNum

`func (o *VpnTunnelStatusVO) HasConnectedNum() bool`

HasConnectedNum returns a boolean if a field has been set.

### GetDirection

`func (o *VpnTunnelStatusVO) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *VpnTunnelStatusVO) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *VpnTunnelStatusVO) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *VpnTunnelStatusVO) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetDisconnectedNum

`func (o *VpnTunnelStatusVO) GetDisconnectedNum() int64`

GetDisconnectedNum returns the DisconnectedNum field if non-nil, zero value otherwise.

### GetDisconnectedNumOk

`func (o *VpnTunnelStatusVO) GetDisconnectedNumOk() (*int64, bool)`

GetDisconnectedNumOk returns a tuple with the DisconnectedNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNum

`func (o *VpnTunnelStatusVO) SetDisconnectedNum(v int64)`

SetDisconnectedNum sets DisconnectedNum field to given value.

### HasDisconnectedNum

`func (o *VpnTunnelStatusVO) HasDisconnectedNum() bool`

HasDisconnectedNum returns a boolean if a field has been set.

### GetDns

`func (o *VpnTunnelStatusVO) GetDns() string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *VpnTunnelStatusVO) GetDnsOk() (*string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *VpnTunnelStatusVO) SetDns(v string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *VpnTunnelStatusVO) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetDownBytes

`func (o *VpnTunnelStatusVO) GetDownBytes() int64`

GetDownBytes returns the DownBytes field if non-nil, zero value otherwise.

### GetDownBytesOk

`func (o *VpnTunnelStatusVO) GetDownBytesOk() (*int64, bool)`

GetDownBytesOk returns a tuple with the DownBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownBytes

`func (o *VpnTunnelStatusVO) SetDownBytes(v int64)`

SetDownBytes sets DownBytes field to given value.

### HasDownBytes

`func (o *VpnTunnelStatusVO) HasDownBytes() bool`

HasDownBytes returns a boolean if a field has been set.

### GetDownPkts

`func (o *VpnTunnelStatusVO) GetDownPkts() int64`

GetDownPkts returns the DownPkts field if non-nil, zero value otherwise.

### GetDownPktsOk

`func (o *VpnTunnelStatusVO) GetDownPktsOk() (*int64, bool)`

GetDownPktsOk returns a tuple with the DownPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownPkts

`func (o *VpnTunnelStatusVO) SetDownPkts(v int64)`

SetDownPkts sets DownPkts field to given value.

### HasDownPkts

`func (o *VpnTunnelStatusVO) HasDownPkts() bool`

HasDownPkts returns a boolean if a field has been set.

### GetEspAuthentication

`func (o *VpnTunnelStatusVO) GetEspAuthentication() string`

GetEspAuthentication returns the EspAuthentication field if non-nil, zero value otherwise.

### GetEspAuthenticationOk

`func (o *VpnTunnelStatusVO) GetEspAuthenticationOk() (*string, bool)`

GetEspAuthenticationOk returns a tuple with the EspAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEspAuthentication

`func (o *VpnTunnelStatusVO) SetEspAuthentication(v string)`

SetEspAuthentication sets EspAuthentication field to given value.

### HasEspAuthentication

`func (o *VpnTunnelStatusVO) HasEspAuthentication() bool`

HasEspAuthentication returns a boolean if a field has been set.

### GetEspEncryption

`func (o *VpnTunnelStatusVO) GetEspEncryption() string`

GetEspEncryption returns the EspEncryption field if non-nil, zero value otherwise.

### GetEspEncryptionOk

`func (o *VpnTunnelStatusVO) GetEspEncryptionOk() (*string, bool)`

GetEspEncryptionOk returns a tuple with the EspEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEspEncryption

`func (o *VpnTunnelStatusVO) SetEspEncryption(v string)`

SetEspEncryption sets EspEncryption field to given value.

### HasEspEncryption

`func (o *VpnTunnelStatusVO) HasEspEncryption() bool`

HasEspEncryption returns a boolean if a field has been set.

### GetId

`func (o *VpnTunnelStatusVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpnTunnelStatusVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpnTunnelStatusVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VpnTunnelStatusVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceName

`func (o *VpnTunnelStatusVO) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *VpnTunnelStatusVO) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *VpnTunnelStatusVO) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *VpnTunnelStatusVO) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetIpPool

`func (o *VpnTunnelStatusVO) GetIpPool() string`

GetIpPool returns the IpPool field if non-nil, zero value otherwise.

### GetIpPoolOk

`func (o *VpnTunnelStatusVO) GetIpPoolOk() (*string, bool)`

GetIpPoolOk returns a tuple with the IpPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPool

`func (o *VpnTunnelStatusVO) SetIpPool(v string)`

SetIpPool sets IpPool field to given value.

### HasIpPool

`func (o *VpnTunnelStatusVO) HasIpPool() bool`

HasIpPool returns a boolean if a field has been set.

### GetIpPoolEnd

`func (o *VpnTunnelStatusVO) GetIpPoolEnd() string`

GetIpPoolEnd returns the IpPoolEnd field if non-nil, zero value otherwise.

### GetIpPoolEndOk

`func (o *VpnTunnelStatusVO) GetIpPoolEndOk() (*string, bool)`

GetIpPoolEndOk returns a tuple with the IpPoolEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolEnd

`func (o *VpnTunnelStatusVO) SetIpPoolEnd(v string)`

SetIpPoolEnd sets IpPoolEnd field to given value.

### HasIpPoolEnd

`func (o *VpnTunnelStatusVO) HasIpPoolEnd() bool`

HasIpPoolEnd returns a boolean if a field has been set.

### GetIpPoolStart

`func (o *VpnTunnelStatusVO) GetIpPoolStart() string`

GetIpPoolStart returns the IpPoolStart field if non-nil, zero value otherwise.

### GetIpPoolStartOk

`func (o *VpnTunnelStatusVO) GetIpPoolStartOk() (*string, bool)`

GetIpPoolStartOk returns a tuple with the IpPoolStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolStart

`func (o *VpnTunnelStatusVO) SetIpPoolStart(v string)`

SetIpPoolStart sets IpPoolStart field to given value.

### HasIpPoolStart

`func (o *VpnTunnelStatusVO) HasIpPoolStart() bool`

HasIpPoolStart returns a boolean if a field has been set.

### GetIpPoolType

`func (o *VpnTunnelStatusVO) GetIpPoolType() int32`

GetIpPoolType returns the IpPoolType field if non-nil, zero value otherwise.

### GetIpPoolTypeOk

`func (o *VpnTunnelStatusVO) GetIpPoolTypeOk() (*int32, bool)`

GetIpPoolTypeOk returns a tuple with the IpPoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolType

`func (o *VpnTunnelStatusVO) SetIpPoolType(v int32)`

SetIpPoolType sets IpPoolType field to given value.

### HasIpPoolType

`func (o *VpnTunnelStatusVO) HasIpPoolType() bool`

HasIpPoolType returns a boolean if a field has been set.

### GetLocalIp

`func (o *VpnTunnelStatusVO) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *VpnTunnelStatusVO) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *VpnTunnelStatusVO) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.

### HasLocalIp

`func (o *VpnTunnelStatusVO) HasLocalIp() bool`

HasLocalIp returns a boolean if a field has been set.

### GetLocalPeerIp

`func (o *VpnTunnelStatusVO) GetLocalPeerIp() string`

GetLocalPeerIp returns the LocalPeerIp field if non-nil, zero value otherwise.

### GetLocalPeerIpOk

`func (o *VpnTunnelStatusVO) GetLocalPeerIpOk() (*string, bool)`

GetLocalPeerIpOk returns a tuple with the LocalPeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPeerIp

`func (o *VpnTunnelStatusVO) SetLocalPeerIp(v string)`

SetLocalPeerIp sets LocalPeerIp field to given value.

### HasLocalPeerIp

`func (o *VpnTunnelStatusVO) HasLocalPeerIp() bool`

HasLocalPeerIp returns a boolean if a field has been set.

### GetLocalSa

`func (o *VpnTunnelStatusVO) GetLocalSa() string`

GetLocalSa returns the LocalSa field if non-nil, zero value otherwise.

### GetLocalSaOk

`func (o *VpnTunnelStatusVO) GetLocalSaOk() (*string, bool)`

GetLocalSaOk returns a tuple with the LocalSa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSa

`func (o *VpnTunnelStatusVO) SetLocalSa(v string)`

SetLocalSa sets LocalSa field to given value.

### HasLocalSa

`func (o *VpnTunnelStatusVO) HasLocalSa() bool`

HasLocalSa returns a boolean if a field has been set.

### GetLoginIp

`func (o *VpnTunnelStatusVO) GetLoginIp() string`

GetLoginIp returns the LoginIp field if non-nil, zero value otherwise.

### GetLoginIpOk

`func (o *VpnTunnelStatusVO) GetLoginIpOk() (*string, bool)`

GetLoginIpOk returns a tuple with the LoginIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginIp

`func (o *VpnTunnelStatusVO) SetLoginIp(v string)`

SetLoginIp sets LoginIp field to given value.

### HasLoginIp

`func (o *VpnTunnelStatusVO) HasLoginIp() bool`

HasLoginIp returns a boolean if a field has been set.

### GetLoginTime

`func (o *VpnTunnelStatusVO) GetLoginTime() int64`

GetLoginTime returns the LoginTime field if non-nil, zero value otherwise.

### GetLoginTimeOk

`func (o *VpnTunnelStatusVO) GetLoginTimeOk() (*int64, bool)`

GetLoginTimeOk returns a tuple with the LoginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginTime

`func (o *VpnTunnelStatusVO) SetLoginTime(v int64)`

SetLoginTime sets LoginTime field to given value.

### HasLoginTime

`func (o *VpnTunnelStatusVO) HasLoginTime() bool`

HasLoginTime returns a boolean if a field has been set.

### GetName

`func (o *VpnTunnelStatusVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnTunnelStatusVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnTunnelStatusVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpnTunnelStatusVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *VpnTunnelStatusVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VpnTunnelStatusVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VpnTunnelStatusVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *VpnTunnelStatusVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *VpnTunnelStatusVO) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *VpnTunnelStatusVO) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *VpnTunnelStatusVO) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *VpnTunnelStatusVO) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRemoteIp

`func (o *VpnTunnelStatusVO) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *VpnTunnelStatusVO) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *VpnTunnelStatusVO) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *VpnTunnelStatusVO) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetRemotePeerIp

`func (o *VpnTunnelStatusVO) GetRemotePeerIp() string`

GetRemotePeerIp returns the RemotePeerIp field if non-nil, zero value otherwise.

### GetRemotePeerIpOk

`func (o *VpnTunnelStatusVO) GetRemotePeerIpOk() (*string, bool)`

GetRemotePeerIpOk returns a tuple with the RemotePeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePeerIp

`func (o *VpnTunnelStatusVO) SetRemotePeerIp(v string)`

SetRemotePeerIp sets RemotePeerIp field to given value.

### HasRemotePeerIp

`func (o *VpnTunnelStatusVO) HasRemotePeerIp() bool`

HasRemotePeerIp returns a boolean if a field has been set.

### GetRemoteSa

`func (o *VpnTunnelStatusVO) GetRemoteSa() string`

GetRemoteSa returns the RemoteSa field if non-nil, zero value otherwise.

### GetRemoteSaOk

`func (o *VpnTunnelStatusVO) GetRemoteSaOk() (*string, bool)`

GetRemoteSaOk returns a tuple with the RemoteSa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSa

`func (o *VpnTunnelStatusVO) SetRemoteSa(v string)`

SetRemoteSa sets RemoteSa field to given value.

### HasRemoteSa

`func (o *VpnTunnelStatusVO) HasRemoteSa() bool`

HasRemoteSa returns a boolean if a field has been set.

### GetServerType

`func (o *VpnTunnelStatusVO) GetServerType() int32`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *VpnTunnelStatusVO) GetServerTypeOk() (*int32, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *VpnTunnelStatusVO) SetServerType(v int32)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *VpnTunnelStatusVO) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetServicePort

`func (o *VpnTunnelStatusVO) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *VpnTunnelStatusVO) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *VpnTunnelStatusVO) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *VpnTunnelStatusVO) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetSpi

`func (o *VpnTunnelStatusVO) GetSpi() int64`

GetSpi returns the Spi field if non-nil, zero value otherwise.

### GetSpiOk

`func (o *VpnTunnelStatusVO) GetSpiOk() (*int64, bool)`

GetSpiOk returns a tuple with the Spi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpi

`func (o *VpnTunnelStatusVO) SetSpi(v int64)`

SetSpi sets Spi field to given value.

### HasSpi

`func (o *VpnTunnelStatusVO) HasSpi() bool`

HasSpi returns a boolean if a field has been set.

### GetStatus

`func (o *VpnTunnelStatusVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnTunnelStatusVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnTunnelStatusVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VpnTunnelStatusVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalRemoteNum

`func (o *VpnTunnelStatusVO) GetTotalRemoteNum() int64`

GetTotalRemoteNum returns the TotalRemoteNum field if non-nil, zero value otherwise.

### GetTotalRemoteNumOk

`func (o *VpnTunnelStatusVO) GetTotalRemoteNumOk() (*int64, bool)`

GetTotalRemoteNumOk returns a tuple with the TotalRemoteNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRemoteNum

`func (o *VpnTunnelStatusVO) SetTotalRemoteNum(v int64)`

SetTotalRemoteNum sets TotalRemoteNum field to given value.

### HasTotalRemoteNum

`func (o *VpnTunnelStatusVO) HasTotalRemoteNum() bool`

HasTotalRemoteNum returns a boolean if a field has been set.

### GetUpBytes

`func (o *VpnTunnelStatusVO) GetUpBytes() int64`

GetUpBytes returns the UpBytes field if non-nil, zero value otherwise.

### GetUpBytesOk

`func (o *VpnTunnelStatusVO) GetUpBytesOk() (*int64, bool)`

GetUpBytesOk returns a tuple with the UpBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpBytes

`func (o *VpnTunnelStatusVO) SetUpBytes(v int64)`

SetUpBytes sets UpBytes field to given value.

### HasUpBytes

`func (o *VpnTunnelStatusVO) HasUpBytes() bool`

HasUpBytes returns a boolean if a field has been set.

### GetUpPkts

`func (o *VpnTunnelStatusVO) GetUpPkts() int64`

GetUpPkts returns the UpPkts field if non-nil, zero value otherwise.

### GetUpPktsOk

`func (o *VpnTunnelStatusVO) GetUpPktsOk() (*int64, bool)`

GetUpPktsOk returns a tuple with the UpPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPkts

`func (o *VpnTunnelStatusVO) SetUpPkts(v int64)`

SetUpPkts sets UpPkts field to given value.

### HasUpPkts

`func (o *VpnTunnelStatusVO) HasUpPkts() bool`

HasUpPkts returns a boolean if a field has been set.

### GetUptime

`func (o *VpnTunnelStatusVO) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *VpnTunnelStatusVO) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *VpnTunnelStatusVO) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *VpnTunnelStatusVO) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetUserId

`func (o *VpnTunnelStatusVO) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *VpnTunnelStatusVO) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *VpnTunnelStatusVO) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *VpnTunnelStatusVO) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *VpnTunnelStatusVO) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *VpnTunnelStatusVO) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *VpnTunnelStatusVO) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *VpnTunnelStatusVO) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetVirtualIp

`func (o *VpnTunnelStatusVO) GetVirtualIp() string`

GetVirtualIp returns the VirtualIp field if non-nil, zero value otherwise.

### GetVirtualIpOk

`func (o *VpnTunnelStatusVO) GetVirtualIpOk() (*string, bool)`

GetVirtualIpOk returns a tuple with the VirtualIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIp

`func (o *VpnTunnelStatusVO) SetVirtualIp(v string)`

SetVirtualIp sets VirtualIp field to given value.

### HasVirtualIp

`func (o *VpnTunnelStatusVO) HasVirtualIp() bool`

HasVirtualIp returns a boolean if a field has been set.

### GetVpnId

`func (o *VpnTunnelStatusVO) GetVpnId() string`

GetVpnId returns the VpnId field if non-nil, zero value otherwise.

### GetVpnIdOk

`func (o *VpnTunnelStatusVO) GetVpnIdOk() (*string, bool)`

GetVpnIdOk returns a tuple with the VpnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnId

`func (o *VpnTunnelStatusVO) SetVpnId(v string)`

SetVpnId sets VpnId field to given value.

### HasVpnId

`func (o *VpnTunnelStatusVO) HasVpnId() bool`

HasVpnId returns a boolean if a field has been set.

### GetVpnType

`func (o *VpnTunnelStatusVO) GetVpnType() int32`

GetVpnType returns the VpnType field if non-nil, zero value otherwise.

### GetVpnTypeOk

`func (o *VpnTunnelStatusVO) GetVpnTypeOk() (*int32, bool)`

GetVpnTypeOk returns a tuple with the VpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnType

`func (o *VpnTunnelStatusVO) SetVpnType(v int32)`

SetVpnType sets VpnType field to given value.

### HasVpnType

`func (o *VpnTunnelStatusVO) HasVpnType() bool`

HasVpnType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


