# ClientDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Whether the client is online. | [optional] 
**Activity** | Pointer to **int64** | Real-time downlink rate (Byte/s). | [optional] 
**ApMac** | Pointer to **string** | (Wireless)  AP MAC Address. | [optional] 
**ApName** | Pointer to **string** | (Wireless)  AP Name. | [optional] 
**AuthInfo** | Pointer to [**[]AuthInfoOpenApiVO**](AuthInfoOpenApiVO.md) | Client portal authentication information | [optional] 
**AuthStatus** | Pointer to **int32** | Authentication status should be a value as follows: 0: CONNECTED // Access without any authentication method; 1: PENDING // Access to Portal, but authentication failed; 2: AUTHORIZED // Pass through portal, pass other authentication without portal; 3: AUTH-FREE // No portal authentication required. | [optional] 
**BlockDisable** | Pointer to **bool** | Block client disabled, default value: false | [optional] 
**Blocked** | Pointer to **bool** | Whether the client is blocked. | [optional] 
**Capabilities** | Pointer to **[]string** | One or more of the following values: Station、DOCSIS cable device、Telephone、Router、WLAN access point、Bridge、Repeater、other. | [optional] 
**Channel** | Pointer to **int32** | (Wireless)  Actual channel. | [optional] 
**ClientLockToApSetting** | Pointer to [**ClientLockToApDetailSetting**](ClientLockToApDetailSetting.md) |  | [optional] 
**ConnectDevType** | Pointer to **string** | connect device type should be a value as follows: ap, switch, gateway. | [optional] 
**ConnectType** | Pointer to **int32** | Connect type should be a value as follows: 0: wireless guest; 1: wireless user; 2: wired user. | [optional] 
**ConnectedToWirelessRouter** | Pointer to **bool** | true: Client is connecting to a wireless router. | [optional] 
**Description** | Pointer to **string** | Device description. | [optional] 
**DeviceCategory** | Pointer to **string** | Device Category: loT, TV, computer, phone... | [optional] 
**DeviceType** | Pointer to **string** | Device Type: iphone, ipod, android, pc, printer, tv... | [optional] 
**DhcpLeaseTime** | Pointer to **int64** | DHCP lease time, unit seconds | [optional] 
**Dot1xIdentity** | Pointer to **string** | (Wired) 802.1x authentication identity. | [optional] 
**Dot1xVlan** | Pointer to **int32** | (Wired) Network name corresponding to the VLAN obtained by 802.1x D-VLAN. | [optional] 
**DownPacket** | Pointer to **int64** | Number of downstream packets. | [optional] 
**GatewayMac** | Pointer to **string** | (Wired, connectDevType&#x3D;gateway)  Gateway MAC Address. | [optional] 
**GatewayName** | Pointer to **string** | (Wired, connectDevType&#x3D;gateway)  Gateway name. | [optional] 
**Guest** | Pointer to **bool** | (Wireless) Whether it is Guest (used to display the wireless Guest client icon). | [optional] 
**HostName** | Pointer to **string** | Host name, device name. | [optional] 
**Id** | Pointer to **string** | Client ID. | [optional] 
**Ip** | Pointer to **string** | IP Address. | [optional] 
**IpSetting** | Pointer to [**ClientIpSetting**](ClientIpSetting.md) |  | [optional] 
**Ipv6List** | Pointer to **[]string** | IPv6 Address. | [optional] 
**LagId** | Pointer to **int32** | (Wired) LAG ID. Exists only when the client is connected to the LAG. | [optional] 
**LastSeen** | Pointer to **int64** | Last found time, timestamp (ms). | [optional] 
**LocateEnable** | Pointer to **bool** | Whether locate function is enabled | [optional] 
**Mac** | Pointer to **string** | Client MAC Address. | [optional] 
**Manager** | Pointer to **bool** | Whether it is the client currently being managed. | [optional] 
**Model** | Pointer to **string** | Model of client device. | [optional] 
**MultiLink** | Pointer to [**[]ClientMultifrequencyInfo**](ClientMultifrequencyInfo.md) | (Wireless) Client multifrequency info list. | [optional] 
**Name** | Pointer to **string** | Client Name, alias. | [optional] 
**NetworkName** | Pointer to **string** | (Wired) Network name. | [optional] 
**OsName** | Pointer to **string** | Device system version. | [optional] 
**Port** | Pointer to **int32** | (Wired) Port ID. | [optional] 
**PowerSave** | Pointer to **bool** | (Wireless)  true: Power save mode enabled. | [optional] 
**RadioId** | Pointer to **int32** | (Wireless) Radio ID should be a value as follows: 0: 2.4GHz; 1: 5GHz-1; 2:5GHz-2; 3: 6GHz. | [optional] 
**RateLimit** | Pointer to [**RateLimitSettingOfClient**](RateLimitSettingOfClient.md) |  | [optional] 
**Rssi** | Pointer to **int32** | (Wireless) Signal strength, unit: dBm. | [optional] 
**RxRate** | Pointer to **int64** | (Wireless) Uplink negotiation rate (Kbit/s). | [optional] 
**SignalLevel** | Pointer to **int32** | (Wireless) Signal strength percentage should be within the range of 0-100. | [optional] 
**SignalRank** | Pointer to **int32** | (Wireless) Signal strength level should be within the range of 0-5. | [optional] 
**Snr** | Pointer to **int32** | (Wireless) Signal Noise Ratio. | [optional] 
**Ssid** | Pointer to **string** | (Wireless)  SSID name. | [optional] 
**StPortsInLag** | Pointer to **[]string** | Standard switch ports in Lag. Exists only when lagId exists and stackId exists | [optional] 
**StackId** | Pointer to **string** | Stack Id | [optional] 
**StandardPort** | Pointer to **string** | Standard port. | [optional] 
**SupportLocate** | Pointer to **bool** | Whether the client supports locate.True when the client is wired connected to a switch or stack that supports locate port. | [optional] 
**SwitchMac** | Pointer to **string** | (Wired, connectDevType&#x3D;switch)  Switch MAC address. | [optional] 
**SwitchName** | Pointer to **string** | (Wired, connectDevType&#x3D;switch)  Switch name. | [optional] 
**SwitchPortsInLag** | Pointer to **[]int32** | Switch ports in Lag. Exists only when lagId exists and stackId is null | [optional] 
**SystemName** | Pointer to **string** | Device system name. | [optional] 
**TrafficDown** | Pointer to **int64** | Downstream traffic (Byte). | [optional] 
**TrafficUp** | Pointer to **int64** | Upstream traffic (Byte). | [optional] 
**TxRate** | Pointer to **int64** | (Wireless) Downlink negotiation rate (Kbit/s). | [optional] 
**Unit** | Pointer to **int32** | Unit ID. | [optional] 
**UpPacket** | Pointer to **int64** | Number of upstream packets. | [optional] 
**UploadActivity** | Pointer to **int64** | Real-time downlink rate (Byte/s). | [optional] 
**Uptime** | Pointer to **int64** | Up time (unit: s). | [optional] 
**Vendor** | Pointer to **string** | Vendor. | [optional] 
**Vid** | Pointer to **int32** | (Wired) vlan. | [optional] 
**WifiMode** | Pointer to **int32** | (Wireless) Wifi mode should be a value as follows: 0: 11a; 1: 11b; 2: 11g; 3: 11na; 4: 11ng; 5: 11ac; 6: 11axa; 7: 11axg; 8: 11beg; 9: 11bea. | [optional] 
**Wireless** | Pointer to **bool** | true: Wireless device (connectDevType&#x3D;ap);  false: Not wireless device(connectDevType&#x3D;switch or gateway) | [optional] 

## Methods

### NewClientDetail

`func NewClientDetail() *ClientDetail`

NewClientDetail instantiates a new ClientDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientDetailWithDefaults

`func NewClientDetailWithDefaults() *ClientDetail`

NewClientDetailWithDefaults instantiates a new ClientDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ClientDetail) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ClientDetail) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ClientDetail) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ClientDetail) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetActivity

`func (o *ClientDetail) GetActivity() int64`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *ClientDetail) GetActivityOk() (*int64, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *ClientDetail) SetActivity(v int64)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *ClientDetail) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetApMac

`func (o *ClientDetail) GetApMac() string`

GetApMac returns the ApMac field if non-nil, zero value otherwise.

### GetApMacOk

`func (o *ClientDetail) GetApMacOk() (*string, bool)`

GetApMacOk returns a tuple with the ApMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMac

`func (o *ClientDetail) SetApMac(v string)`

SetApMac sets ApMac field to given value.

### HasApMac

`func (o *ClientDetail) HasApMac() bool`

HasApMac returns a boolean if a field has been set.

### GetApName

`func (o *ClientDetail) GetApName() string`

GetApName returns the ApName field if non-nil, zero value otherwise.

### GetApNameOk

`func (o *ClientDetail) GetApNameOk() (*string, bool)`

GetApNameOk returns a tuple with the ApName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApName

`func (o *ClientDetail) SetApName(v string)`

SetApName sets ApName field to given value.

### HasApName

`func (o *ClientDetail) HasApName() bool`

HasApName returns a boolean if a field has been set.

### GetAuthInfo

`func (o *ClientDetail) GetAuthInfo() []AuthInfoOpenApiVO`

GetAuthInfo returns the AuthInfo field if non-nil, zero value otherwise.

### GetAuthInfoOk

`func (o *ClientDetail) GetAuthInfoOk() (*[]AuthInfoOpenApiVO, bool)`

GetAuthInfoOk returns a tuple with the AuthInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthInfo

`func (o *ClientDetail) SetAuthInfo(v []AuthInfoOpenApiVO)`

SetAuthInfo sets AuthInfo field to given value.

### HasAuthInfo

`func (o *ClientDetail) HasAuthInfo() bool`

HasAuthInfo returns a boolean if a field has been set.

### GetAuthStatus

`func (o *ClientDetail) GetAuthStatus() int32`

GetAuthStatus returns the AuthStatus field if non-nil, zero value otherwise.

### GetAuthStatusOk

`func (o *ClientDetail) GetAuthStatusOk() (*int32, bool)`

GetAuthStatusOk returns a tuple with the AuthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthStatus

`func (o *ClientDetail) SetAuthStatus(v int32)`

SetAuthStatus sets AuthStatus field to given value.

### HasAuthStatus

`func (o *ClientDetail) HasAuthStatus() bool`

HasAuthStatus returns a boolean if a field has been set.

### GetBlockDisable

`func (o *ClientDetail) GetBlockDisable() bool`

GetBlockDisable returns the BlockDisable field if non-nil, zero value otherwise.

### GetBlockDisableOk

`func (o *ClientDetail) GetBlockDisableOk() (*bool, bool)`

GetBlockDisableOk returns a tuple with the BlockDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDisable

`func (o *ClientDetail) SetBlockDisable(v bool)`

SetBlockDisable sets BlockDisable field to given value.

### HasBlockDisable

`func (o *ClientDetail) HasBlockDisable() bool`

HasBlockDisable returns a boolean if a field has been set.

### GetBlocked

`func (o *ClientDetail) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *ClientDetail) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *ClientDetail) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *ClientDetail) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetCapabilities

`func (o *ClientDetail) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ClientDetail) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ClientDetail) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ClientDetail) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetChannel

`func (o *ClientDetail) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ClientDetail) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ClientDetail) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ClientDetail) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetClientLockToApSetting

`func (o *ClientDetail) GetClientLockToApSetting() ClientLockToApDetailSetting`

GetClientLockToApSetting returns the ClientLockToApSetting field if non-nil, zero value otherwise.

### GetClientLockToApSettingOk

`func (o *ClientDetail) GetClientLockToApSettingOk() (*ClientLockToApDetailSetting, bool)`

GetClientLockToApSettingOk returns a tuple with the ClientLockToApSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientLockToApSetting

`func (o *ClientDetail) SetClientLockToApSetting(v ClientLockToApDetailSetting)`

SetClientLockToApSetting sets ClientLockToApSetting field to given value.

### HasClientLockToApSetting

`func (o *ClientDetail) HasClientLockToApSetting() bool`

HasClientLockToApSetting returns a boolean if a field has been set.

### GetConnectDevType

`func (o *ClientDetail) GetConnectDevType() string`

GetConnectDevType returns the ConnectDevType field if non-nil, zero value otherwise.

### GetConnectDevTypeOk

`func (o *ClientDetail) GetConnectDevTypeOk() (*string, bool)`

GetConnectDevTypeOk returns a tuple with the ConnectDevType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectDevType

`func (o *ClientDetail) SetConnectDevType(v string)`

SetConnectDevType sets ConnectDevType field to given value.

### HasConnectDevType

`func (o *ClientDetail) HasConnectDevType() bool`

HasConnectDevType returns a boolean if a field has been set.

### GetConnectType

`func (o *ClientDetail) GetConnectType() int32`

GetConnectType returns the ConnectType field if non-nil, zero value otherwise.

### GetConnectTypeOk

`func (o *ClientDetail) GetConnectTypeOk() (*int32, bool)`

GetConnectTypeOk returns a tuple with the ConnectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectType

`func (o *ClientDetail) SetConnectType(v int32)`

SetConnectType sets ConnectType field to given value.

### HasConnectType

`func (o *ClientDetail) HasConnectType() bool`

HasConnectType returns a boolean if a field has been set.

### GetConnectedToWirelessRouter

`func (o *ClientDetail) GetConnectedToWirelessRouter() bool`

GetConnectedToWirelessRouter returns the ConnectedToWirelessRouter field if non-nil, zero value otherwise.

### GetConnectedToWirelessRouterOk

`func (o *ClientDetail) GetConnectedToWirelessRouterOk() (*bool, bool)`

GetConnectedToWirelessRouterOk returns a tuple with the ConnectedToWirelessRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedToWirelessRouter

`func (o *ClientDetail) SetConnectedToWirelessRouter(v bool)`

SetConnectedToWirelessRouter sets ConnectedToWirelessRouter field to given value.

### HasConnectedToWirelessRouter

`func (o *ClientDetail) HasConnectedToWirelessRouter() bool`

HasConnectedToWirelessRouter returns a boolean if a field has been set.

### GetDescription

`func (o *ClientDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClientDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClientDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClientDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceCategory

`func (o *ClientDetail) GetDeviceCategory() string`

GetDeviceCategory returns the DeviceCategory field if non-nil, zero value otherwise.

### GetDeviceCategoryOk

`func (o *ClientDetail) GetDeviceCategoryOk() (*string, bool)`

GetDeviceCategoryOk returns a tuple with the DeviceCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCategory

`func (o *ClientDetail) SetDeviceCategory(v string)`

SetDeviceCategory sets DeviceCategory field to given value.

### HasDeviceCategory

`func (o *ClientDetail) HasDeviceCategory() bool`

HasDeviceCategory returns a boolean if a field has been set.

### GetDeviceType

`func (o *ClientDetail) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ClientDetail) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ClientDetail) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ClientDetail) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDhcpLeaseTime

`func (o *ClientDetail) GetDhcpLeaseTime() int64`

GetDhcpLeaseTime returns the DhcpLeaseTime field if non-nil, zero value otherwise.

### GetDhcpLeaseTimeOk

`func (o *ClientDetail) GetDhcpLeaseTimeOk() (*int64, bool)`

GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpLeaseTime

`func (o *ClientDetail) SetDhcpLeaseTime(v int64)`

SetDhcpLeaseTime sets DhcpLeaseTime field to given value.

### HasDhcpLeaseTime

`func (o *ClientDetail) HasDhcpLeaseTime() bool`

HasDhcpLeaseTime returns a boolean if a field has been set.

### GetDot1xIdentity

`func (o *ClientDetail) GetDot1xIdentity() string`

GetDot1xIdentity returns the Dot1xIdentity field if non-nil, zero value otherwise.

### GetDot1xIdentityOk

`func (o *ClientDetail) GetDot1xIdentityOk() (*string, bool)`

GetDot1xIdentityOk returns a tuple with the Dot1xIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1xIdentity

`func (o *ClientDetail) SetDot1xIdentity(v string)`

SetDot1xIdentity sets Dot1xIdentity field to given value.

### HasDot1xIdentity

`func (o *ClientDetail) HasDot1xIdentity() bool`

HasDot1xIdentity returns a boolean if a field has been set.

### GetDot1xVlan

`func (o *ClientDetail) GetDot1xVlan() int32`

GetDot1xVlan returns the Dot1xVlan field if non-nil, zero value otherwise.

### GetDot1xVlanOk

`func (o *ClientDetail) GetDot1xVlanOk() (*int32, bool)`

GetDot1xVlanOk returns a tuple with the Dot1xVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1xVlan

`func (o *ClientDetail) SetDot1xVlan(v int32)`

SetDot1xVlan sets Dot1xVlan field to given value.

### HasDot1xVlan

`func (o *ClientDetail) HasDot1xVlan() bool`

HasDot1xVlan returns a boolean if a field has been set.

### GetDownPacket

`func (o *ClientDetail) GetDownPacket() int64`

GetDownPacket returns the DownPacket field if non-nil, zero value otherwise.

### GetDownPacketOk

`func (o *ClientDetail) GetDownPacketOk() (*int64, bool)`

GetDownPacketOk returns a tuple with the DownPacket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownPacket

`func (o *ClientDetail) SetDownPacket(v int64)`

SetDownPacket sets DownPacket field to given value.

### HasDownPacket

`func (o *ClientDetail) HasDownPacket() bool`

HasDownPacket returns a boolean if a field has been set.

### GetGatewayMac

`func (o *ClientDetail) GetGatewayMac() string`

GetGatewayMac returns the GatewayMac field if non-nil, zero value otherwise.

### GetGatewayMacOk

`func (o *ClientDetail) GetGatewayMacOk() (*string, bool)`

GetGatewayMacOk returns a tuple with the GatewayMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayMac

`func (o *ClientDetail) SetGatewayMac(v string)`

SetGatewayMac sets GatewayMac field to given value.

### HasGatewayMac

`func (o *ClientDetail) HasGatewayMac() bool`

HasGatewayMac returns a boolean if a field has been set.

### GetGatewayName

`func (o *ClientDetail) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *ClientDetail) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *ClientDetail) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.

### HasGatewayName

`func (o *ClientDetail) HasGatewayName() bool`

HasGatewayName returns a boolean if a field has been set.

### GetGuest

`func (o *ClientDetail) GetGuest() bool`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *ClientDetail) GetGuestOk() (*bool, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *ClientDetail) SetGuest(v bool)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *ClientDetail) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetHostName

`func (o *ClientDetail) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ClientDetail) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ClientDetail) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *ClientDetail) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetId

`func (o *ClientDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *ClientDetail) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ClientDetail) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ClientDetail) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ClientDetail) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpSetting

`func (o *ClientDetail) GetIpSetting() ClientIpSetting`

GetIpSetting returns the IpSetting field if non-nil, zero value otherwise.

### GetIpSettingOk

`func (o *ClientDetail) GetIpSettingOk() (*ClientIpSetting, bool)`

GetIpSettingOk returns a tuple with the IpSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSetting

`func (o *ClientDetail) SetIpSetting(v ClientIpSetting)`

SetIpSetting sets IpSetting field to given value.

### HasIpSetting

`func (o *ClientDetail) HasIpSetting() bool`

HasIpSetting returns a boolean if a field has been set.

### GetIpv6List

`func (o *ClientDetail) GetIpv6List() []string`

GetIpv6List returns the Ipv6List field if non-nil, zero value otherwise.

### GetIpv6ListOk

`func (o *ClientDetail) GetIpv6ListOk() (*[]string, bool)`

GetIpv6ListOk returns a tuple with the Ipv6List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6List

`func (o *ClientDetail) SetIpv6List(v []string)`

SetIpv6List sets Ipv6List field to given value.

### HasIpv6List

`func (o *ClientDetail) HasIpv6List() bool`

HasIpv6List returns a boolean if a field has been set.

### GetLagId

`func (o *ClientDetail) GetLagId() int32`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *ClientDetail) GetLagIdOk() (*int32, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *ClientDetail) SetLagId(v int32)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *ClientDetail) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetLastSeen

`func (o *ClientDetail) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ClientDetail) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ClientDetail) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ClientDetail) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLocateEnable

`func (o *ClientDetail) GetLocateEnable() bool`

GetLocateEnable returns the LocateEnable field if non-nil, zero value otherwise.

### GetLocateEnableOk

`func (o *ClientDetail) GetLocateEnableOk() (*bool, bool)`

GetLocateEnableOk returns a tuple with the LocateEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocateEnable

`func (o *ClientDetail) SetLocateEnable(v bool)`

SetLocateEnable sets LocateEnable field to given value.

### HasLocateEnable

`func (o *ClientDetail) HasLocateEnable() bool`

HasLocateEnable returns a boolean if a field has been set.

### GetMac

`func (o *ClientDetail) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ClientDetail) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ClientDetail) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ClientDetail) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetManager

`func (o *ClientDetail) GetManager() bool`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *ClientDetail) GetManagerOk() (*bool, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *ClientDetail) SetManager(v bool)`

SetManager sets Manager field to given value.

### HasManager

`func (o *ClientDetail) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetModel

`func (o *ClientDetail) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ClientDetail) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ClientDetail) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ClientDetail) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetMultiLink

`func (o *ClientDetail) GetMultiLink() []ClientMultifrequencyInfo`

GetMultiLink returns the MultiLink field if non-nil, zero value otherwise.

### GetMultiLinkOk

`func (o *ClientDetail) GetMultiLinkOk() (*[]ClientMultifrequencyInfo, bool)`

GetMultiLinkOk returns a tuple with the MultiLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiLink

`func (o *ClientDetail) SetMultiLink(v []ClientMultifrequencyInfo)`

SetMultiLink sets MultiLink field to given value.

### HasMultiLink

`func (o *ClientDetail) HasMultiLink() bool`

HasMultiLink returns a boolean if a field has been set.

### GetName

`func (o *ClientDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkName

`func (o *ClientDetail) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *ClientDetail) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *ClientDetail) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *ClientDetail) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetOsName

`func (o *ClientDetail) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *ClientDetail) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *ClientDetail) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *ClientDetail) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetPort

`func (o *ClientDetail) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ClientDetail) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ClientDetail) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ClientDetail) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPowerSave

`func (o *ClientDetail) GetPowerSave() bool`

GetPowerSave returns the PowerSave field if non-nil, zero value otherwise.

### GetPowerSaveOk

`func (o *ClientDetail) GetPowerSaveOk() (*bool, bool)`

GetPowerSaveOk returns a tuple with the PowerSave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSave

`func (o *ClientDetail) SetPowerSave(v bool)`

SetPowerSave sets PowerSave field to given value.

### HasPowerSave

`func (o *ClientDetail) HasPowerSave() bool`

HasPowerSave returns a boolean if a field has been set.

### GetRadioId

`func (o *ClientDetail) GetRadioId() int32`

GetRadioId returns the RadioId field if non-nil, zero value otherwise.

### GetRadioIdOk

`func (o *ClientDetail) GetRadioIdOk() (*int32, bool)`

GetRadioIdOk returns a tuple with the RadioId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioId

`func (o *ClientDetail) SetRadioId(v int32)`

SetRadioId sets RadioId field to given value.

### HasRadioId

`func (o *ClientDetail) HasRadioId() bool`

HasRadioId returns a boolean if a field has been set.

### GetRateLimit

`func (o *ClientDetail) GetRateLimit() RateLimitSettingOfClient`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *ClientDetail) GetRateLimitOk() (*RateLimitSettingOfClient, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *ClientDetail) SetRateLimit(v RateLimitSettingOfClient)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *ClientDetail) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetRssi

`func (o *ClientDetail) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *ClientDetail) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *ClientDetail) SetRssi(v int32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *ClientDetail) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetRxRate

`func (o *ClientDetail) GetRxRate() int64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *ClientDetail) GetRxRateOk() (*int64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *ClientDetail) SetRxRate(v int64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *ClientDetail) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetSignalLevel

`func (o *ClientDetail) GetSignalLevel() int32`

GetSignalLevel returns the SignalLevel field if non-nil, zero value otherwise.

### GetSignalLevelOk

`func (o *ClientDetail) GetSignalLevelOk() (*int32, bool)`

GetSignalLevelOk returns a tuple with the SignalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalLevel

`func (o *ClientDetail) SetSignalLevel(v int32)`

SetSignalLevel sets SignalLevel field to given value.

### HasSignalLevel

`func (o *ClientDetail) HasSignalLevel() bool`

HasSignalLevel returns a boolean if a field has been set.

### GetSignalRank

`func (o *ClientDetail) GetSignalRank() int32`

GetSignalRank returns the SignalRank field if non-nil, zero value otherwise.

### GetSignalRankOk

`func (o *ClientDetail) GetSignalRankOk() (*int32, bool)`

GetSignalRankOk returns a tuple with the SignalRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalRank

`func (o *ClientDetail) SetSignalRank(v int32)`

SetSignalRank sets SignalRank field to given value.

### HasSignalRank

`func (o *ClientDetail) HasSignalRank() bool`

HasSignalRank returns a boolean if a field has been set.

### GetSnr

`func (o *ClientDetail) GetSnr() int32`

GetSnr returns the Snr field if non-nil, zero value otherwise.

### GetSnrOk

`func (o *ClientDetail) GetSnrOk() (*int32, bool)`

GetSnrOk returns a tuple with the Snr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnr

`func (o *ClientDetail) SetSnr(v int32)`

SetSnr sets Snr field to given value.

### HasSnr

`func (o *ClientDetail) HasSnr() bool`

HasSnr returns a boolean if a field has been set.

### GetSsid

`func (o *ClientDetail) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *ClientDetail) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *ClientDetail) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *ClientDetail) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetStPortsInLag

`func (o *ClientDetail) GetStPortsInLag() []string`

GetStPortsInLag returns the StPortsInLag field if non-nil, zero value otherwise.

### GetStPortsInLagOk

`func (o *ClientDetail) GetStPortsInLagOk() (*[]string, bool)`

GetStPortsInLagOk returns a tuple with the StPortsInLag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStPortsInLag

`func (o *ClientDetail) SetStPortsInLag(v []string)`

SetStPortsInLag sets StPortsInLag field to given value.

### HasStPortsInLag

`func (o *ClientDetail) HasStPortsInLag() bool`

HasStPortsInLag returns a boolean if a field has been set.

### GetStackId

`func (o *ClientDetail) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *ClientDetail) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *ClientDetail) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *ClientDetail) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStandardPort

`func (o *ClientDetail) GetStandardPort() string`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *ClientDetail) GetStandardPortOk() (*string, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *ClientDetail) SetStandardPort(v string)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *ClientDetail) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.

### GetSupportLocate

`func (o *ClientDetail) GetSupportLocate() bool`

GetSupportLocate returns the SupportLocate field if non-nil, zero value otherwise.

### GetSupportLocateOk

`func (o *ClientDetail) GetSupportLocateOk() (*bool, bool)`

GetSupportLocateOk returns a tuple with the SupportLocate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLocate

`func (o *ClientDetail) SetSupportLocate(v bool)`

SetSupportLocate sets SupportLocate field to given value.

### HasSupportLocate

`func (o *ClientDetail) HasSupportLocate() bool`

HasSupportLocate returns a boolean if a field has been set.

### GetSwitchMac

`func (o *ClientDetail) GetSwitchMac() string`

GetSwitchMac returns the SwitchMac field if non-nil, zero value otherwise.

### GetSwitchMacOk

`func (o *ClientDetail) GetSwitchMacOk() (*string, bool)`

GetSwitchMacOk returns a tuple with the SwitchMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMac

`func (o *ClientDetail) SetSwitchMac(v string)`

SetSwitchMac sets SwitchMac field to given value.

### HasSwitchMac

`func (o *ClientDetail) HasSwitchMac() bool`

HasSwitchMac returns a boolean if a field has been set.

### GetSwitchName

`func (o *ClientDetail) GetSwitchName() string`

GetSwitchName returns the SwitchName field if non-nil, zero value otherwise.

### GetSwitchNameOk

`func (o *ClientDetail) GetSwitchNameOk() (*string, bool)`

GetSwitchNameOk returns a tuple with the SwitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchName

`func (o *ClientDetail) SetSwitchName(v string)`

SetSwitchName sets SwitchName field to given value.

### HasSwitchName

`func (o *ClientDetail) HasSwitchName() bool`

HasSwitchName returns a boolean if a field has been set.

### GetSwitchPortsInLag

`func (o *ClientDetail) GetSwitchPortsInLag() []int32`

GetSwitchPortsInLag returns the SwitchPortsInLag field if non-nil, zero value otherwise.

### GetSwitchPortsInLagOk

`func (o *ClientDetail) GetSwitchPortsInLagOk() (*[]int32, bool)`

GetSwitchPortsInLagOk returns a tuple with the SwitchPortsInLag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortsInLag

`func (o *ClientDetail) SetSwitchPortsInLag(v []int32)`

SetSwitchPortsInLag sets SwitchPortsInLag field to given value.

### HasSwitchPortsInLag

`func (o *ClientDetail) HasSwitchPortsInLag() bool`

HasSwitchPortsInLag returns a boolean if a field has been set.

### GetSystemName

`func (o *ClientDetail) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *ClientDetail) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *ClientDetail) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *ClientDetail) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetTrafficDown

`func (o *ClientDetail) GetTrafficDown() int64`

GetTrafficDown returns the TrafficDown field if non-nil, zero value otherwise.

### GetTrafficDownOk

`func (o *ClientDetail) GetTrafficDownOk() (*int64, bool)`

GetTrafficDownOk returns a tuple with the TrafficDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficDown

`func (o *ClientDetail) SetTrafficDown(v int64)`

SetTrafficDown sets TrafficDown field to given value.

### HasTrafficDown

`func (o *ClientDetail) HasTrafficDown() bool`

HasTrafficDown returns a boolean if a field has been set.

### GetTrafficUp

`func (o *ClientDetail) GetTrafficUp() int64`

GetTrafficUp returns the TrafficUp field if non-nil, zero value otherwise.

### GetTrafficUpOk

`func (o *ClientDetail) GetTrafficUpOk() (*int64, bool)`

GetTrafficUpOk returns a tuple with the TrafficUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficUp

`func (o *ClientDetail) SetTrafficUp(v int64)`

SetTrafficUp sets TrafficUp field to given value.

### HasTrafficUp

`func (o *ClientDetail) HasTrafficUp() bool`

HasTrafficUp returns a boolean if a field has been set.

### GetTxRate

`func (o *ClientDetail) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *ClientDetail) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *ClientDetail) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *ClientDetail) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.

### GetUnit

`func (o *ClientDetail) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ClientDetail) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ClientDetail) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ClientDetail) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUpPacket

`func (o *ClientDetail) GetUpPacket() int64`

GetUpPacket returns the UpPacket field if non-nil, zero value otherwise.

### GetUpPacketOk

`func (o *ClientDetail) GetUpPacketOk() (*int64, bool)`

GetUpPacketOk returns a tuple with the UpPacket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPacket

`func (o *ClientDetail) SetUpPacket(v int64)`

SetUpPacket sets UpPacket field to given value.

### HasUpPacket

`func (o *ClientDetail) HasUpPacket() bool`

HasUpPacket returns a boolean if a field has been set.

### GetUploadActivity

`func (o *ClientDetail) GetUploadActivity() int64`

GetUploadActivity returns the UploadActivity field if non-nil, zero value otherwise.

### GetUploadActivityOk

`func (o *ClientDetail) GetUploadActivityOk() (*int64, bool)`

GetUploadActivityOk returns a tuple with the UploadActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadActivity

`func (o *ClientDetail) SetUploadActivity(v int64)`

SetUploadActivity sets UploadActivity field to given value.

### HasUploadActivity

`func (o *ClientDetail) HasUploadActivity() bool`

HasUploadActivity returns a boolean if a field has been set.

### GetUptime

`func (o *ClientDetail) GetUptime() int64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *ClientDetail) GetUptimeOk() (*int64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *ClientDetail) SetUptime(v int64)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *ClientDetail) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVendor

`func (o *ClientDetail) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ClientDetail) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ClientDetail) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ClientDetail) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVid

`func (o *ClientDetail) GetVid() int32`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *ClientDetail) GetVidOk() (*int32, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *ClientDetail) SetVid(v int32)`

SetVid sets Vid field to given value.

### HasVid

`func (o *ClientDetail) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetWifiMode

`func (o *ClientDetail) GetWifiMode() int32`

GetWifiMode returns the WifiMode field if non-nil, zero value otherwise.

### GetWifiModeOk

`func (o *ClientDetail) GetWifiModeOk() (*int32, bool)`

GetWifiModeOk returns a tuple with the WifiMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMode

`func (o *ClientDetail) SetWifiMode(v int32)`

SetWifiMode sets WifiMode field to given value.

### HasWifiMode

`func (o *ClientDetail) HasWifiMode() bool`

HasWifiMode returns a boolean if a field has been set.

### GetWireless

`func (o *ClientDetail) GetWireless() bool`

GetWireless returns the Wireless field if non-nil, zero value otherwise.

### GetWirelessOk

`func (o *ClientDetail) GetWirelessOk() (*bool, bool)`

GetWirelessOk returns a tuple with the Wireless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireless

`func (o *ClientDetail) SetWireless(v bool)`

SetWireless sets Wireless field to given value.

### HasWireless

`func (o *ClientDetail) HasWireless() bool`

HasWireless returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


