# OpenApiClientInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Whether the client is online. | [optional] 
**Activity** | Pointer to **int64** | Real-time downlink rate (Byte/s). | [optional] 
**ApMac** | Pointer to **string** | (Wireless)  AP MAC Address. | [optional] 
**ApName** | Pointer to **string** | (Wireless)  AP Name. | [optional] 
**AuthInfo** | Pointer to [**[]AuthInfoOpenApiVO**](AuthInfoOpenApiVO.md) | Client portal authentication information | [optional] 
**AuthStatus** | Pointer to **int32** | Authentication status should be a value as follows: 0: CONNECTED // Access without any authentication method; 1: PENDING // Access to Portal, but authentication failed; 2: AUTHORIZED // Pass through portal, pass other authentication without portal; 3: AUTH-FREE // No portal authentication required. | [optional] 
**BlockDisable** | Pointer to **bool** | Block client disabled, default value: false. | [optional] 
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
**HealthScore** | Pointer to **int32** | 1~3: poor; 4~7: fair; 0: no data; 8~10 good. | [optional] 
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
**PortName** | Pointer to **string** | (Wired) Port name. | [optional] 
**PowerSave** | Pointer to **bool** | (Wireless)  true: Power save mode enabled. | [optional] 
**PpskProfileName** | Pointer to **string** | (Wireless)  SSID PPSK profile name | [optional] 
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
**StackName** | Pointer to **string** | Stack name | [optional] 
**StackableSwitch** | Pointer to **bool** | Whether the connected device is a stackable switch. | [optional] 
**StandardPort** | Pointer to **string** | Standard port. | [optional] 
**Support5g2** | Pointer to **bool** | Whether the client is connecting under 5g2. | [optional] 
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
**UploadActivity** | Pointer to **int64** | Real-time uplink rate (Byte/s). | [optional] 
**Uptime** | Pointer to **int64** | Up time (unit: s). | [optional] 
**Vendor** | Pointer to **string** | Vendor. | [optional] 
**Vid** | Pointer to **int32** | (Wired) vlan. | [optional] 
**WifiMode** | Pointer to **int32** | (Wireless) Wifi mode should be a value as follows: 0: 11a; 1: 11b; 2: 11g; 3: 11na; 4: 11ng; 5: 11ac; 6: 11axa; 7: 11axg; 8: 11beg; 9: 11bea. | [optional] 
**Wireless** | Pointer to **bool** | true: Wireless device (connectDevType&#x3D;ap);  false: Not wireless device(connectDevType&#x3D;switch or gateway). | [optional] 

## Methods

### NewOpenApiClientInfo

`func NewOpenApiClientInfo() *OpenApiClientInfo`

NewOpenApiClientInfo instantiates a new OpenApiClientInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenApiClientInfoWithDefaults

`func NewOpenApiClientInfoWithDefaults() *OpenApiClientInfo`

NewOpenApiClientInfoWithDefaults instantiates a new OpenApiClientInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *OpenApiClientInfo) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OpenApiClientInfo) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OpenApiClientInfo) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *OpenApiClientInfo) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetActivity

`func (o *OpenApiClientInfo) GetActivity() int64`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *OpenApiClientInfo) GetActivityOk() (*int64, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *OpenApiClientInfo) SetActivity(v int64)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *OpenApiClientInfo) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetApMac

`func (o *OpenApiClientInfo) GetApMac() string`

GetApMac returns the ApMac field if non-nil, zero value otherwise.

### GetApMacOk

`func (o *OpenApiClientInfo) GetApMacOk() (*string, bool)`

GetApMacOk returns a tuple with the ApMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMac

`func (o *OpenApiClientInfo) SetApMac(v string)`

SetApMac sets ApMac field to given value.

### HasApMac

`func (o *OpenApiClientInfo) HasApMac() bool`

HasApMac returns a boolean if a field has been set.

### GetApName

`func (o *OpenApiClientInfo) GetApName() string`

GetApName returns the ApName field if non-nil, zero value otherwise.

### GetApNameOk

`func (o *OpenApiClientInfo) GetApNameOk() (*string, bool)`

GetApNameOk returns a tuple with the ApName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApName

`func (o *OpenApiClientInfo) SetApName(v string)`

SetApName sets ApName field to given value.

### HasApName

`func (o *OpenApiClientInfo) HasApName() bool`

HasApName returns a boolean if a field has been set.

### GetAuthInfo

`func (o *OpenApiClientInfo) GetAuthInfo() []AuthInfoOpenApiVO`

GetAuthInfo returns the AuthInfo field if non-nil, zero value otherwise.

### GetAuthInfoOk

`func (o *OpenApiClientInfo) GetAuthInfoOk() (*[]AuthInfoOpenApiVO, bool)`

GetAuthInfoOk returns a tuple with the AuthInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthInfo

`func (o *OpenApiClientInfo) SetAuthInfo(v []AuthInfoOpenApiVO)`

SetAuthInfo sets AuthInfo field to given value.

### HasAuthInfo

`func (o *OpenApiClientInfo) HasAuthInfo() bool`

HasAuthInfo returns a boolean if a field has been set.

### GetAuthStatus

`func (o *OpenApiClientInfo) GetAuthStatus() int32`

GetAuthStatus returns the AuthStatus field if non-nil, zero value otherwise.

### GetAuthStatusOk

`func (o *OpenApiClientInfo) GetAuthStatusOk() (*int32, bool)`

GetAuthStatusOk returns a tuple with the AuthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthStatus

`func (o *OpenApiClientInfo) SetAuthStatus(v int32)`

SetAuthStatus sets AuthStatus field to given value.

### HasAuthStatus

`func (o *OpenApiClientInfo) HasAuthStatus() bool`

HasAuthStatus returns a boolean if a field has been set.

### GetBlockDisable

`func (o *OpenApiClientInfo) GetBlockDisable() bool`

GetBlockDisable returns the BlockDisable field if non-nil, zero value otherwise.

### GetBlockDisableOk

`func (o *OpenApiClientInfo) GetBlockDisableOk() (*bool, bool)`

GetBlockDisableOk returns a tuple with the BlockDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDisable

`func (o *OpenApiClientInfo) SetBlockDisable(v bool)`

SetBlockDisable sets BlockDisable field to given value.

### HasBlockDisable

`func (o *OpenApiClientInfo) HasBlockDisable() bool`

HasBlockDisable returns a boolean if a field has been set.

### GetBlocked

`func (o *OpenApiClientInfo) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *OpenApiClientInfo) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *OpenApiClientInfo) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *OpenApiClientInfo) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetCapabilities

`func (o *OpenApiClientInfo) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *OpenApiClientInfo) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *OpenApiClientInfo) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *OpenApiClientInfo) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetChannel

`func (o *OpenApiClientInfo) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *OpenApiClientInfo) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *OpenApiClientInfo) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *OpenApiClientInfo) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetClientLockToApSetting

`func (o *OpenApiClientInfo) GetClientLockToApSetting() ClientLockToApDetailSetting`

GetClientLockToApSetting returns the ClientLockToApSetting field if non-nil, zero value otherwise.

### GetClientLockToApSettingOk

`func (o *OpenApiClientInfo) GetClientLockToApSettingOk() (*ClientLockToApDetailSetting, bool)`

GetClientLockToApSettingOk returns a tuple with the ClientLockToApSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientLockToApSetting

`func (o *OpenApiClientInfo) SetClientLockToApSetting(v ClientLockToApDetailSetting)`

SetClientLockToApSetting sets ClientLockToApSetting field to given value.

### HasClientLockToApSetting

`func (o *OpenApiClientInfo) HasClientLockToApSetting() bool`

HasClientLockToApSetting returns a boolean if a field has been set.

### GetConnectDevType

`func (o *OpenApiClientInfo) GetConnectDevType() string`

GetConnectDevType returns the ConnectDevType field if non-nil, zero value otherwise.

### GetConnectDevTypeOk

`func (o *OpenApiClientInfo) GetConnectDevTypeOk() (*string, bool)`

GetConnectDevTypeOk returns a tuple with the ConnectDevType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectDevType

`func (o *OpenApiClientInfo) SetConnectDevType(v string)`

SetConnectDevType sets ConnectDevType field to given value.

### HasConnectDevType

`func (o *OpenApiClientInfo) HasConnectDevType() bool`

HasConnectDevType returns a boolean if a field has been set.

### GetConnectType

`func (o *OpenApiClientInfo) GetConnectType() int32`

GetConnectType returns the ConnectType field if non-nil, zero value otherwise.

### GetConnectTypeOk

`func (o *OpenApiClientInfo) GetConnectTypeOk() (*int32, bool)`

GetConnectTypeOk returns a tuple with the ConnectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectType

`func (o *OpenApiClientInfo) SetConnectType(v int32)`

SetConnectType sets ConnectType field to given value.

### HasConnectType

`func (o *OpenApiClientInfo) HasConnectType() bool`

HasConnectType returns a boolean if a field has been set.

### GetConnectedToWirelessRouter

`func (o *OpenApiClientInfo) GetConnectedToWirelessRouter() bool`

GetConnectedToWirelessRouter returns the ConnectedToWirelessRouter field if non-nil, zero value otherwise.

### GetConnectedToWirelessRouterOk

`func (o *OpenApiClientInfo) GetConnectedToWirelessRouterOk() (*bool, bool)`

GetConnectedToWirelessRouterOk returns a tuple with the ConnectedToWirelessRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedToWirelessRouter

`func (o *OpenApiClientInfo) SetConnectedToWirelessRouter(v bool)`

SetConnectedToWirelessRouter sets ConnectedToWirelessRouter field to given value.

### HasConnectedToWirelessRouter

`func (o *OpenApiClientInfo) HasConnectedToWirelessRouter() bool`

HasConnectedToWirelessRouter returns a boolean if a field has been set.

### GetDescription

`func (o *OpenApiClientInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpenApiClientInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpenApiClientInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpenApiClientInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceCategory

`func (o *OpenApiClientInfo) GetDeviceCategory() string`

GetDeviceCategory returns the DeviceCategory field if non-nil, zero value otherwise.

### GetDeviceCategoryOk

`func (o *OpenApiClientInfo) GetDeviceCategoryOk() (*string, bool)`

GetDeviceCategoryOk returns a tuple with the DeviceCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCategory

`func (o *OpenApiClientInfo) SetDeviceCategory(v string)`

SetDeviceCategory sets DeviceCategory field to given value.

### HasDeviceCategory

`func (o *OpenApiClientInfo) HasDeviceCategory() bool`

HasDeviceCategory returns a boolean if a field has been set.

### GetDeviceType

`func (o *OpenApiClientInfo) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *OpenApiClientInfo) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *OpenApiClientInfo) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *OpenApiClientInfo) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDhcpLeaseTime

`func (o *OpenApiClientInfo) GetDhcpLeaseTime() int64`

GetDhcpLeaseTime returns the DhcpLeaseTime field if non-nil, zero value otherwise.

### GetDhcpLeaseTimeOk

`func (o *OpenApiClientInfo) GetDhcpLeaseTimeOk() (*int64, bool)`

GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpLeaseTime

`func (o *OpenApiClientInfo) SetDhcpLeaseTime(v int64)`

SetDhcpLeaseTime sets DhcpLeaseTime field to given value.

### HasDhcpLeaseTime

`func (o *OpenApiClientInfo) HasDhcpLeaseTime() bool`

HasDhcpLeaseTime returns a boolean if a field has been set.

### GetDot1xIdentity

`func (o *OpenApiClientInfo) GetDot1xIdentity() string`

GetDot1xIdentity returns the Dot1xIdentity field if non-nil, zero value otherwise.

### GetDot1xIdentityOk

`func (o *OpenApiClientInfo) GetDot1xIdentityOk() (*string, bool)`

GetDot1xIdentityOk returns a tuple with the Dot1xIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1xIdentity

`func (o *OpenApiClientInfo) SetDot1xIdentity(v string)`

SetDot1xIdentity sets Dot1xIdentity field to given value.

### HasDot1xIdentity

`func (o *OpenApiClientInfo) HasDot1xIdentity() bool`

HasDot1xIdentity returns a boolean if a field has been set.

### GetDot1xVlan

`func (o *OpenApiClientInfo) GetDot1xVlan() int32`

GetDot1xVlan returns the Dot1xVlan field if non-nil, zero value otherwise.

### GetDot1xVlanOk

`func (o *OpenApiClientInfo) GetDot1xVlanOk() (*int32, bool)`

GetDot1xVlanOk returns a tuple with the Dot1xVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1xVlan

`func (o *OpenApiClientInfo) SetDot1xVlan(v int32)`

SetDot1xVlan sets Dot1xVlan field to given value.

### HasDot1xVlan

`func (o *OpenApiClientInfo) HasDot1xVlan() bool`

HasDot1xVlan returns a boolean if a field has been set.

### GetDownPacket

`func (o *OpenApiClientInfo) GetDownPacket() int64`

GetDownPacket returns the DownPacket field if non-nil, zero value otherwise.

### GetDownPacketOk

`func (o *OpenApiClientInfo) GetDownPacketOk() (*int64, bool)`

GetDownPacketOk returns a tuple with the DownPacket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownPacket

`func (o *OpenApiClientInfo) SetDownPacket(v int64)`

SetDownPacket sets DownPacket field to given value.

### HasDownPacket

`func (o *OpenApiClientInfo) HasDownPacket() bool`

HasDownPacket returns a boolean if a field has been set.

### GetGatewayMac

`func (o *OpenApiClientInfo) GetGatewayMac() string`

GetGatewayMac returns the GatewayMac field if non-nil, zero value otherwise.

### GetGatewayMacOk

`func (o *OpenApiClientInfo) GetGatewayMacOk() (*string, bool)`

GetGatewayMacOk returns a tuple with the GatewayMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayMac

`func (o *OpenApiClientInfo) SetGatewayMac(v string)`

SetGatewayMac sets GatewayMac field to given value.

### HasGatewayMac

`func (o *OpenApiClientInfo) HasGatewayMac() bool`

HasGatewayMac returns a boolean if a field has been set.

### GetGatewayName

`func (o *OpenApiClientInfo) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *OpenApiClientInfo) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *OpenApiClientInfo) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.

### HasGatewayName

`func (o *OpenApiClientInfo) HasGatewayName() bool`

HasGatewayName returns a boolean if a field has been set.

### GetGuest

`func (o *OpenApiClientInfo) GetGuest() bool`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *OpenApiClientInfo) GetGuestOk() (*bool, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *OpenApiClientInfo) SetGuest(v bool)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *OpenApiClientInfo) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetHealthScore

`func (o *OpenApiClientInfo) GetHealthScore() int32`

GetHealthScore returns the HealthScore field if non-nil, zero value otherwise.

### GetHealthScoreOk

`func (o *OpenApiClientInfo) GetHealthScoreOk() (*int32, bool)`

GetHealthScoreOk returns a tuple with the HealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthScore

`func (o *OpenApiClientInfo) SetHealthScore(v int32)`

SetHealthScore sets HealthScore field to given value.

### HasHealthScore

`func (o *OpenApiClientInfo) HasHealthScore() bool`

HasHealthScore returns a boolean if a field has been set.

### GetHostName

`func (o *OpenApiClientInfo) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *OpenApiClientInfo) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *OpenApiClientInfo) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *OpenApiClientInfo) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetId

`func (o *OpenApiClientInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenApiClientInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenApiClientInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OpenApiClientInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *OpenApiClientInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OpenApiClientInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OpenApiClientInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OpenApiClientInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpSetting

`func (o *OpenApiClientInfo) GetIpSetting() ClientIpSetting`

GetIpSetting returns the IpSetting field if non-nil, zero value otherwise.

### GetIpSettingOk

`func (o *OpenApiClientInfo) GetIpSettingOk() (*ClientIpSetting, bool)`

GetIpSettingOk returns a tuple with the IpSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSetting

`func (o *OpenApiClientInfo) SetIpSetting(v ClientIpSetting)`

SetIpSetting sets IpSetting field to given value.

### HasIpSetting

`func (o *OpenApiClientInfo) HasIpSetting() bool`

HasIpSetting returns a boolean if a field has been set.

### GetIpv6List

`func (o *OpenApiClientInfo) GetIpv6List() []string`

GetIpv6List returns the Ipv6List field if non-nil, zero value otherwise.

### GetIpv6ListOk

`func (o *OpenApiClientInfo) GetIpv6ListOk() (*[]string, bool)`

GetIpv6ListOk returns a tuple with the Ipv6List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6List

`func (o *OpenApiClientInfo) SetIpv6List(v []string)`

SetIpv6List sets Ipv6List field to given value.

### HasIpv6List

`func (o *OpenApiClientInfo) HasIpv6List() bool`

HasIpv6List returns a boolean if a field has been set.

### GetLagId

`func (o *OpenApiClientInfo) GetLagId() int32`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *OpenApiClientInfo) GetLagIdOk() (*int32, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *OpenApiClientInfo) SetLagId(v int32)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *OpenApiClientInfo) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetLastSeen

`func (o *OpenApiClientInfo) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *OpenApiClientInfo) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *OpenApiClientInfo) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *OpenApiClientInfo) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLocateEnable

`func (o *OpenApiClientInfo) GetLocateEnable() bool`

GetLocateEnable returns the LocateEnable field if non-nil, zero value otherwise.

### GetLocateEnableOk

`func (o *OpenApiClientInfo) GetLocateEnableOk() (*bool, bool)`

GetLocateEnableOk returns a tuple with the LocateEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocateEnable

`func (o *OpenApiClientInfo) SetLocateEnable(v bool)`

SetLocateEnable sets LocateEnable field to given value.

### HasLocateEnable

`func (o *OpenApiClientInfo) HasLocateEnable() bool`

HasLocateEnable returns a boolean if a field has been set.

### GetMac

`func (o *OpenApiClientInfo) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OpenApiClientInfo) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OpenApiClientInfo) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OpenApiClientInfo) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetManager

`func (o *OpenApiClientInfo) GetManager() bool`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *OpenApiClientInfo) GetManagerOk() (*bool, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *OpenApiClientInfo) SetManager(v bool)`

SetManager sets Manager field to given value.

### HasManager

`func (o *OpenApiClientInfo) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetModel

`func (o *OpenApiClientInfo) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OpenApiClientInfo) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OpenApiClientInfo) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OpenApiClientInfo) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetMultiLink

`func (o *OpenApiClientInfo) GetMultiLink() []ClientMultifrequencyInfo`

GetMultiLink returns the MultiLink field if non-nil, zero value otherwise.

### GetMultiLinkOk

`func (o *OpenApiClientInfo) GetMultiLinkOk() (*[]ClientMultifrequencyInfo, bool)`

GetMultiLinkOk returns a tuple with the MultiLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiLink

`func (o *OpenApiClientInfo) SetMultiLink(v []ClientMultifrequencyInfo)`

SetMultiLink sets MultiLink field to given value.

### HasMultiLink

`func (o *OpenApiClientInfo) HasMultiLink() bool`

HasMultiLink returns a boolean if a field has been set.

### GetName

`func (o *OpenApiClientInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenApiClientInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenApiClientInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenApiClientInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkName

`func (o *OpenApiClientInfo) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *OpenApiClientInfo) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *OpenApiClientInfo) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *OpenApiClientInfo) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetOsName

`func (o *OpenApiClientInfo) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *OpenApiClientInfo) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *OpenApiClientInfo) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *OpenApiClientInfo) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetPort

`func (o *OpenApiClientInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OpenApiClientInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OpenApiClientInfo) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OpenApiClientInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortName

`func (o *OpenApiClientInfo) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *OpenApiClientInfo) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *OpenApiClientInfo) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *OpenApiClientInfo) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetPowerSave

`func (o *OpenApiClientInfo) GetPowerSave() bool`

GetPowerSave returns the PowerSave field if non-nil, zero value otherwise.

### GetPowerSaveOk

`func (o *OpenApiClientInfo) GetPowerSaveOk() (*bool, bool)`

GetPowerSaveOk returns a tuple with the PowerSave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSave

`func (o *OpenApiClientInfo) SetPowerSave(v bool)`

SetPowerSave sets PowerSave field to given value.

### HasPowerSave

`func (o *OpenApiClientInfo) HasPowerSave() bool`

HasPowerSave returns a boolean if a field has been set.

### GetPpskProfileName

`func (o *OpenApiClientInfo) GetPpskProfileName() string`

GetPpskProfileName returns the PpskProfileName field if non-nil, zero value otherwise.

### GetPpskProfileNameOk

`func (o *OpenApiClientInfo) GetPpskProfileNameOk() (*string, bool)`

GetPpskProfileNameOk returns a tuple with the PpskProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpskProfileName

`func (o *OpenApiClientInfo) SetPpskProfileName(v string)`

SetPpskProfileName sets PpskProfileName field to given value.

### HasPpskProfileName

`func (o *OpenApiClientInfo) HasPpskProfileName() bool`

HasPpskProfileName returns a boolean if a field has been set.

### GetRadioId

`func (o *OpenApiClientInfo) GetRadioId() int32`

GetRadioId returns the RadioId field if non-nil, zero value otherwise.

### GetRadioIdOk

`func (o *OpenApiClientInfo) GetRadioIdOk() (*int32, bool)`

GetRadioIdOk returns a tuple with the RadioId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioId

`func (o *OpenApiClientInfo) SetRadioId(v int32)`

SetRadioId sets RadioId field to given value.

### HasRadioId

`func (o *OpenApiClientInfo) HasRadioId() bool`

HasRadioId returns a boolean if a field has been set.

### GetRateLimit

`func (o *OpenApiClientInfo) GetRateLimit() RateLimitSettingOfClient`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *OpenApiClientInfo) GetRateLimitOk() (*RateLimitSettingOfClient, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *OpenApiClientInfo) SetRateLimit(v RateLimitSettingOfClient)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *OpenApiClientInfo) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetRssi

`func (o *OpenApiClientInfo) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *OpenApiClientInfo) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *OpenApiClientInfo) SetRssi(v int32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *OpenApiClientInfo) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetRxRate

`func (o *OpenApiClientInfo) GetRxRate() int64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *OpenApiClientInfo) GetRxRateOk() (*int64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *OpenApiClientInfo) SetRxRate(v int64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *OpenApiClientInfo) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetSignalLevel

`func (o *OpenApiClientInfo) GetSignalLevel() int32`

GetSignalLevel returns the SignalLevel field if non-nil, zero value otherwise.

### GetSignalLevelOk

`func (o *OpenApiClientInfo) GetSignalLevelOk() (*int32, bool)`

GetSignalLevelOk returns a tuple with the SignalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalLevel

`func (o *OpenApiClientInfo) SetSignalLevel(v int32)`

SetSignalLevel sets SignalLevel field to given value.

### HasSignalLevel

`func (o *OpenApiClientInfo) HasSignalLevel() bool`

HasSignalLevel returns a boolean if a field has been set.

### GetSignalRank

`func (o *OpenApiClientInfo) GetSignalRank() int32`

GetSignalRank returns the SignalRank field if non-nil, zero value otherwise.

### GetSignalRankOk

`func (o *OpenApiClientInfo) GetSignalRankOk() (*int32, bool)`

GetSignalRankOk returns a tuple with the SignalRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalRank

`func (o *OpenApiClientInfo) SetSignalRank(v int32)`

SetSignalRank sets SignalRank field to given value.

### HasSignalRank

`func (o *OpenApiClientInfo) HasSignalRank() bool`

HasSignalRank returns a boolean if a field has been set.

### GetSnr

`func (o *OpenApiClientInfo) GetSnr() int32`

GetSnr returns the Snr field if non-nil, zero value otherwise.

### GetSnrOk

`func (o *OpenApiClientInfo) GetSnrOk() (*int32, bool)`

GetSnrOk returns a tuple with the Snr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnr

`func (o *OpenApiClientInfo) SetSnr(v int32)`

SetSnr sets Snr field to given value.

### HasSnr

`func (o *OpenApiClientInfo) HasSnr() bool`

HasSnr returns a boolean if a field has been set.

### GetSsid

`func (o *OpenApiClientInfo) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *OpenApiClientInfo) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *OpenApiClientInfo) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *OpenApiClientInfo) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetStPortsInLag

`func (o *OpenApiClientInfo) GetStPortsInLag() []string`

GetStPortsInLag returns the StPortsInLag field if non-nil, zero value otherwise.

### GetStPortsInLagOk

`func (o *OpenApiClientInfo) GetStPortsInLagOk() (*[]string, bool)`

GetStPortsInLagOk returns a tuple with the StPortsInLag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStPortsInLag

`func (o *OpenApiClientInfo) SetStPortsInLag(v []string)`

SetStPortsInLag sets StPortsInLag field to given value.

### HasStPortsInLag

`func (o *OpenApiClientInfo) HasStPortsInLag() bool`

HasStPortsInLag returns a boolean if a field has been set.

### GetStackId

`func (o *OpenApiClientInfo) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OpenApiClientInfo) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OpenApiClientInfo) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OpenApiClientInfo) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *OpenApiClientInfo) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *OpenApiClientInfo) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *OpenApiClientInfo) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *OpenApiClientInfo) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetStackableSwitch

`func (o *OpenApiClientInfo) GetStackableSwitch() bool`

GetStackableSwitch returns the StackableSwitch field if non-nil, zero value otherwise.

### GetStackableSwitchOk

`func (o *OpenApiClientInfo) GetStackableSwitchOk() (*bool, bool)`

GetStackableSwitchOk returns a tuple with the StackableSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackableSwitch

`func (o *OpenApiClientInfo) SetStackableSwitch(v bool)`

SetStackableSwitch sets StackableSwitch field to given value.

### HasStackableSwitch

`func (o *OpenApiClientInfo) HasStackableSwitch() bool`

HasStackableSwitch returns a boolean if a field has been set.

### GetStandardPort

`func (o *OpenApiClientInfo) GetStandardPort() string`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *OpenApiClientInfo) GetStandardPortOk() (*string, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *OpenApiClientInfo) SetStandardPort(v string)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *OpenApiClientInfo) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.

### GetSupport5g2

`func (o *OpenApiClientInfo) GetSupport5g2() bool`

GetSupport5g2 returns the Support5g2 field if non-nil, zero value otherwise.

### GetSupport5g2Ok

`func (o *OpenApiClientInfo) GetSupport5g2Ok() (*bool, bool)`

GetSupport5g2Ok returns a tuple with the Support5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport5g2

`func (o *OpenApiClientInfo) SetSupport5g2(v bool)`

SetSupport5g2 sets Support5g2 field to given value.

### HasSupport5g2

`func (o *OpenApiClientInfo) HasSupport5g2() bool`

HasSupport5g2 returns a boolean if a field has been set.

### GetSupportLocate

`func (o *OpenApiClientInfo) GetSupportLocate() bool`

GetSupportLocate returns the SupportLocate field if non-nil, zero value otherwise.

### GetSupportLocateOk

`func (o *OpenApiClientInfo) GetSupportLocateOk() (*bool, bool)`

GetSupportLocateOk returns a tuple with the SupportLocate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLocate

`func (o *OpenApiClientInfo) SetSupportLocate(v bool)`

SetSupportLocate sets SupportLocate field to given value.

### HasSupportLocate

`func (o *OpenApiClientInfo) HasSupportLocate() bool`

HasSupportLocate returns a boolean if a field has been set.

### GetSwitchMac

`func (o *OpenApiClientInfo) GetSwitchMac() string`

GetSwitchMac returns the SwitchMac field if non-nil, zero value otherwise.

### GetSwitchMacOk

`func (o *OpenApiClientInfo) GetSwitchMacOk() (*string, bool)`

GetSwitchMacOk returns a tuple with the SwitchMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMac

`func (o *OpenApiClientInfo) SetSwitchMac(v string)`

SetSwitchMac sets SwitchMac field to given value.

### HasSwitchMac

`func (o *OpenApiClientInfo) HasSwitchMac() bool`

HasSwitchMac returns a boolean if a field has been set.

### GetSwitchName

`func (o *OpenApiClientInfo) GetSwitchName() string`

GetSwitchName returns the SwitchName field if non-nil, zero value otherwise.

### GetSwitchNameOk

`func (o *OpenApiClientInfo) GetSwitchNameOk() (*string, bool)`

GetSwitchNameOk returns a tuple with the SwitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchName

`func (o *OpenApiClientInfo) SetSwitchName(v string)`

SetSwitchName sets SwitchName field to given value.

### HasSwitchName

`func (o *OpenApiClientInfo) HasSwitchName() bool`

HasSwitchName returns a boolean if a field has been set.

### GetSwitchPortsInLag

`func (o *OpenApiClientInfo) GetSwitchPortsInLag() []int32`

GetSwitchPortsInLag returns the SwitchPortsInLag field if non-nil, zero value otherwise.

### GetSwitchPortsInLagOk

`func (o *OpenApiClientInfo) GetSwitchPortsInLagOk() (*[]int32, bool)`

GetSwitchPortsInLagOk returns a tuple with the SwitchPortsInLag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortsInLag

`func (o *OpenApiClientInfo) SetSwitchPortsInLag(v []int32)`

SetSwitchPortsInLag sets SwitchPortsInLag field to given value.

### HasSwitchPortsInLag

`func (o *OpenApiClientInfo) HasSwitchPortsInLag() bool`

HasSwitchPortsInLag returns a boolean if a field has been set.

### GetSystemName

`func (o *OpenApiClientInfo) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *OpenApiClientInfo) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *OpenApiClientInfo) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *OpenApiClientInfo) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetTrafficDown

`func (o *OpenApiClientInfo) GetTrafficDown() int64`

GetTrafficDown returns the TrafficDown field if non-nil, zero value otherwise.

### GetTrafficDownOk

`func (o *OpenApiClientInfo) GetTrafficDownOk() (*int64, bool)`

GetTrafficDownOk returns a tuple with the TrafficDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficDown

`func (o *OpenApiClientInfo) SetTrafficDown(v int64)`

SetTrafficDown sets TrafficDown field to given value.

### HasTrafficDown

`func (o *OpenApiClientInfo) HasTrafficDown() bool`

HasTrafficDown returns a boolean if a field has been set.

### GetTrafficUp

`func (o *OpenApiClientInfo) GetTrafficUp() int64`

GetTrafficUp returns the TrafficUp field if non-nil, zero value otherwise.

### GetTrafficUpOk

`func (o *OpenApiClientInfo) GetTrafficUpOk() (*int64, bool)`

GetTrafficUpOk returns a tuple with the TrafficUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficUp

`func (o *OpenApiClientInfo) SetTrafficUp(v int64)`

SetTrafficUp sets TrafficUp field to given value.

### HasTrafficUp

`func (o *OpenApiClientInfo) HasTrafficUp() bool`

HasTrafficUp returns a boolean if a field has been set.

### GetTxRate

`func (o *OpenApiClientInfo) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *OpenApiClientInfo) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *OpenApiClientInfo) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *OpenApiClientInfo) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.

### GetUnit

`func (o *OpenApiClientInfo) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OpenApiClientInfo) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OpenApiClientInfo) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *OpenApiClientInfo) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUpPacket

`func (o *OpenApiClientInfo) GetUpPacket() int64`

GetUpPacket returns the UpPacket field if non-nil, zero value otherwise.

### GetUpPacketOk

`func (o *OpenApiClientInfo) GetUpPacketOk() (*int64, bool)`

GetUpPacketOk returns a tuple with the UpPacket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPacket

`func (o *OpenApiClientInfo) SetUpPacket(v int64)`

SetUpPacket sets UpPacket field to given value.

### HasUpPacket

`func (o *OpenApiClientInfo) HasUpPacket() bool`

HasUpPacket returns a boolean if a field has been set.

### GetUploadActivity

`func (o *OpenApiClientInfo) GetUploadActivity() int64`

GetUploadActivity returns the UploadActivity field if non-nil, zero value otherwise.

### GetUploadActivityOk

`func (o *OpenApiClientInfo) GetUploadActivityOk() (*int64, bool)`

GetUploadActivityOk returns a tuple with the UploadActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadActivity

`func (o *OpenApiClientInfo) SetUploadActivity(v int64)`

SetUploadActivity sets UploadActivity field to given value.

### HasUploadActivity

`func (o *OpenApiClientInfo) HasUploadActivity() bool`

HasUploadActivity returns a boolean if a field has been set.

### GetUptime

`func (o *OpenApiClientInfo) GetUptime() int64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *OpenApiClientInfo) GetUptimeOk() (*int64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *OpenApiClientInfo) SetUptime(v int64)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *OpenApiClientInfo) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVendor

`func (o *OpenApiClientInfo) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *OpenApiClientInfo) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *OpenApiClientInfo) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *OpenApiClientInfo) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVid

`func (o *OpenApiClientInfo) GetVid() int32`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *OpenApiClientInfo) GetVidOk() (*int32, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *OpenApiClientInfo) SetVid(v int32)`

SetVid sets Vid field to given value.

### HasVid

`func (o *OpenApiClientInfo) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetWifiMode

`func (o *OpenApiClientInfo) GetWifiMode() int32`

GetWifiMode returns the WifiMode field if non-nil, zero value otherwise.

### GetWifiModeOk

`func (o *OpenApiClientInfo) GetWifiModeOk() (*int32, bool)`

GetWifiModeOk returns a tuple with the WifiMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMode

`func (o *OpenApiClientInfo) SetWifiMode(v int32)`

SetWifiMode sets WifiMode field to given value.

### HasWifiMode

`func (o *OpenApiClientInfo) HasWifiMode() bool`

HasWifiMode returns a boolean if a field has been set.

### GetWireless

`func (o *OpenApiClientInfo) GetWireless() bool`

GetWireless returns the Wireless field if non-nil, zero value otherwise.

### GetWirelessOk

`func (o *OpenApiClientInfo) GetWirelessOk() (*bool, bool)`

GetWirelessOk returns a tuple with the Wireless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireless

`func (o *OpenApiClientInfo) SetWireless(v bool)`

SetWireless sets Wireless field to given value.

### HasWireless

`func (o *OpenApiClientInfo) HasWireless() bool`

HasWireless returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


