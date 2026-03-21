# OsgVpnTunnelOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientMode** | Pointer to **string** | Client Mode: Client or NEM(Network Extension Mode). | [optional] 
**Dns** | Pointer to **string** | DNS for the peer VPN. | [optional] 
**DownBytes** | Pointer to **int64** | Downlink traffic in bytes. | [optional] 
**DownPkts** | Pointer to **int64** | Number of packets in the downlink. | [optional] 
**InterfaceName** | Pointer to **string** | Interface Name. | [optional] 
**LocalIp** | Pointer to **string** | The local Server IP of the tunnel. | [optional] 
**RemoteIp** | Pointer to **string** | The IP address of the peer VPN User. | [optional] 
**ServerType** | Pointer to **int32** | Client-To-Site VPN Type First-level cascade selection 0:VPN Server 1:VPN Client. | [optional] 
**UpBytes** | Pointer to **int64** | Uplink traffic in bytes. | [optional] 
**UpPkts** | Pointer to **int64** | The amount of uplink packets. | [optional] 
**Uptime** | Pointer to **string** | VPN tunnel effective time, it starts from the connection calculation, accurate to the minute. | [optional] 
**UserId** | Pointer to **int32** | User Id. | [optional] 
**UserName** | Pointer to **string** | Peer Username. | [optional] 
**VpnId** | Pointer to **int32** | VPN Item Id. | [optional] 
**VpnName** | Pointer to **string** | VPN Item Name. | [optional] 
**VpnType** | Pointer to **int32** | Client-To-Site VPN Type Second-level cascade selection 0：L2TP 1: PPTP 2: IPSec 3: OpenVPN. | [optional] 

## Methods

### NewOsgVpnTunnelOpenApiVO

`func NewOsgVpnTunnelOpenApiVO() *OsgVpnTunnelOpenApiVO`

NewOsgVpnTunnelOpenApiVO instantiates a new OsgVpnTunnelOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgVpnTunnelOpenApiVOWithDefaults

`func NewOsgVpnTunnelOpenApiVOWithDefaults() *OsgVpnTunnelOpenApiVO`

NewOsgVpnTunnelOpenApiVOWithDefaults instantiates a new OsgVpnTunnelOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientMode

`func (o *OsgVpnTunnelOpenApiVO) GetClientMode() string`

GetClientMode returns the ClientMode field if non-nil, zero value otherwise.

### GetClientModeOk

`func (o *OsgVpnTunnelOpenApiVO) GetClientModeOk() (*string, bool)`

GetClientModeOk returns a tuple with the ClientMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMode

`func (o *OsgVpnTunnelOpenApiVO) SetClientMode(v string)`

SetClientMode sets ClientMode field to given value.

### HasClientMode

`func (o *OsgVpnTunnelOpenApiVO) HasClientMode() bool`

HasClientMode returns a boolean if a field has been set.

### GetDns

`func (o *OsgVpnTunnelOpenApiVO) GetDns() string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *OsgVpnTunnelOpenApiVO) GetDnsOk() (*string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *OsgVpnTunnelOpenApiVO) SetDns(v string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *OsgVpnTunnelOpenApiVO) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetDownBytes

`func (o *OsgVpnTunnelOpenApiVO) GetDownBytes() int64`

GetDownBytes returns the DownBytes field if non-nil, zero value otherwise.

### GetDownBytesOk

`func (o *OsgVpnTunnelOpenApiVO) GetDownBytesOk() (*int64, bool)`

GetDownBytesOk returns a tuple with the DownBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownBytes

`func (o *OsgVpnTunnelOpenApiVO) SetDownBytes(v int64)`

SetDownBytes sets DownBytes field to given value.

### HasDownBytes

`func (o *OsgVpnTunnelOpenApiVO) HasDownBytes() bool`

HasDownBytes returns a boolean if a field has been set.

### GetDownPkts

`func (o *OsgVpnTunnelOpenApiVO) GetDownPkts() int64`

GetDownPkts returns the DownPkts field if non-nil, zero value otherwise.

### GetDownPktsOk

`func (o *OsgVpnTunnelOpenApiVO) GetDownPktsOk() (*int64, bool)`

GetDownPktsOk returns a tuple with the DownPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownPkts

`func (o *OsgVpnTunnelOpenApiVO) SetDownPkts(v int64)`

SetDownPkts sets DownPkts field to given value.

### HasDownPkts

`func (o *OsgVpnTunnelOpenApiVO) HasDownPkts() bool`

HasDownPkts returns a boolean if a field has been set.

### GetInterfaceName

`func (o *OsgVpnTunnelOpenApiVO) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *OsgVpnTunnelOpenApiVO) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *OsgVpnTunnelOpenApiVO) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *OsgVpnTunnelOpenApiVO) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetLocalIp

`func (o *OsgVpnTunnelOpenApiVO) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *OsgVpnTunnelOpenApiVO) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *OsgVpnTunnelOpenApiVO) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.

### HasLocalIp

`func (o *OsgVpnTunnelOpenApiVO) HasLocalIp() bool`

HasLocalIp returns a boolean if a field has been set.

### GetRemoteIp

`func (o *OsgVpnTunnelOpenApiVO) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *OsgVpnTunnelOpenApiVO) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *OsgVpnTunnelOpenApiVO) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *OsgVpnTunnelOpenApiVO) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetServerType

`func (o *OsgVpnTunnelOpenApiVO) GetServerType() int32`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *OsgVpnTunnelOpenApiVO) GetServerTypeOk() (*int32, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *OsgVpnTunnelOpenApiVO) SetServerType(v int32)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *OsgVpnTunnelOpenApiVO) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetUpBytes

`func (o *OsgVpnTunnelOpenApiVO) GetUpBytes() int64`

GetUpBytes returns the UpBytes field if non-nil, zero value otherwise.

### GetUpBytesOk

`func (o *OsgVpnTunnelOpenApiVO) GetUpBytesOk() (*int64, bool)`

GetUpBytesOk returns a tuple with the UpBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpBytes

`func (o *OsgVpnTunnelOpenApiVO) SetUpBytes(v int64)`

SetUpBytes sets UpBytes field to given value.

### HasUpBytes

`func (o *OsgVpnTunnelOpenApiVO) HasUpBytes() bool`

HasUpBytes returns a boolean if a field has been set.

### GetUpPkts

`func (o *OsgVpnTunnelOpenApiVO) GetUpPkts() int64`

GetUpPkts returns the UpPkts field if non-nil, zero value otherwise.

### GetUpPktsOk

`func (o *OsgVpnTunnelOpenApiVO) GetUpPktsOk() (*int64, bool)`

GetUpPktsOk returns a tuple with the UpPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPkts

`func (o *OsgVpnTunnelOpenApiVO) SetUpPkts(v int64)`

SetUpPkts sets UpPkts field to given value.

### HasUpPkts

`func (o *OsgVpnTunnelOpenApiVO) HasUpPkts() bool`

HasUpPkts returns a boolean if a field has been set.

### GetUptime

`func (o *OsgVpnTunnelOpenApiVO) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *OsgVpnTunnelOpenApiVO) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *OsgVpnTunnelOpenApiVO) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *OsgVpnTunnelOpenApiVO) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetUserId

`func (o *OsgVpnTunnelOpenApiVO) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OsgVpnTunnelOpenApiVO) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OsgVpnTunnelOpenApiVO) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *OsgVpnTunnelOpenApiVO) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *OsgVpnTunnelOpenApiVO) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *OsgVpnTunnelOpenApiVO) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *OsgVpnTunnelOpenApiVO) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *OsgVpnTunnelOpenApiVO) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetVpnId

`func (o *OsgVpnTunnelOpenApiVO) GetVpnId() int32`

GetVpnId returns the VpnId field if non-nil, zero value otherwise.

### GetVpnIdOk

`func (o *OsgVpnTunnelOpenApiVO) GetVpnIdOk() (*int32, bool)`

GetVpnIdOk returns a tuple with the VpnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnId

`func (o *OsgVpnTunnelOpenApiVO) SetVpnId(v int32)`

SetVpnId sets VpnId field to given value.

### HasVpnId

`func (o *OsgVpnTunnelOpenApiVO) HasVpnId() bool`

HasVpnId returns a boolean if a field has been set.

### GetVpnName

`func (o *OsgVpnTunnelOpenApiVO) GetVpnName() string`

GetVpnName returns the VpnName field if non-nil, zero value otherwise.

### GetVpnNameOk

`func (o *OsgVpnTunnelOpenApiVO) GetVpnNameOk() (*string, bool)`

GetVpnNameOk returns a tuple with the VpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnName

`func (o *OsgVpnTunnelOpenApiVO) SetVpnName(v string)`

SetVpnName sets VpnName field to given value.

### HasVpnName

`func (o *OsgVpnTunnelOpenApiVO) HasVpnName() bool`

HasVpnName returns a boolean if a field has been set.

### GetVpnType

`func (o *OsgVpnTunnelOpenApiVO) GetVpnType() int32`

GetVpnType returns the VpnType field if non-nil, zero value otherwise.

### GetVpnTypeOk

`func (o *OsgVpnTunnelOpenApiVO) GetVpnTypeOk() (*int32, bool)`

GetVpnTypeOk returns a tuple with the VpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnType

`func (o *OsgVpnTunnelOpenApiVO) SetVpnType(v int32)`

SetVpnType sets VpnType field to given value.

### HasVpnType

`func (o *OsgVpnTunnelOpenApiVO) HasVpnType() bool`

HasVpnType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


