# SsidDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoWanAccess** | Pointer to **bool** | Whether to enable auto wan access. True: enable, false: disable. | [optional] 
**Band** | Pointer to **int32** | SSID band. The lowest bit indicates whether 2.4G is included; the second lowest bit indicates whether 5G is included; the third lowest bit indicates whether 6G is included; 1 means included while 0 means not included. For example, 7(111) means that 2G/5G/6G are enabled; 1(001) means that 2G is enabled. (When 5G is included，it means 5G/5G1/5G2 are enabled.) | [optional] 
**Broadcast** | Pointer to **bool** | SSID broadcast config status. True: enable, false: disable. | [optional] 
**ClientRateLimit** | Pointer to [**RateLimitSettingOpenApiVO**](RateLimitSettingOpenApiVO.md) |  | [optional] 
**DeviceType** | Pointer to **int32** | SSID device type, identify which devices this SSID will take effect. The lowest bit indicates whether [EAP] is included, the second low bit indicates whether [Gateway] is included, 1 means included while 0 means not included; For example, 3(11) means that EAP/Gateway is enabled, 1(01) means that EAP is enabled. | [optional] 
**DhcpOption82** | Pointer to [**SsidDhcpOptionOpenApiVO**](SsidDhcpOptionOpenApiVO.md) |  | [optional] 
**Enable11r** | Pointer to **bool** | SSID 802.11r config status. True: enable, false: disable. | [optional] 
**EntSetting** | Pointer to [**SsidEnterpriseSettingOpenApiVO**](SsidEnterpriseSettingOpenApiVO.md) |  | [optional] 
**GuestNetEnable** | Pointer to **bool** | SSID guest network config status. True: enable, false: disable. | [optional] 
**HidePwd** | Pointer to **bool** | If this field is true, the SSID password will be hidden. | [optional] 
**HotspotV2Setting** | Pointer to [**HotspotV2SettingOpenApiVO**](HotspotV2SettingOpenApiVO.md) |  | [optional] 
**MacFilter** | Pointer to [**SsidMacFilterOpenApiVO**](SsidMacFilterOpenApiVO.md) |  | [optional] 
**MloEnable** | Pointer to **bool** | SSID MLO config status. True: enable, false: disable. | [optional] 
**MultiCast** | Pointer to [**SsidMultiCastOpenApiVO**](SsidMultiCastOpenApiVO.md) |  | [optional] 
**Name** | Pointer to **string** | SSID name. It should contain 1 to 32 UTF-8 characters. | [optional] 
**OweEnable** | Pointer to **bool** | Opportunistic Wireless Encryption, also known as Enhanced Open, is a certification provided by the Wi-Fi Alliance as part of the WPA3 wireless security standard. OWE will enable two wireless VAPs per radio, one for access of OWE-supported stations, and one for access of other stations. An SSID with OWE enabled will be counted as two SSID entries for 2G and 5G. | [optional] 
**PmfMode** | Pointer to **int32** | SSID PMF mode. It should be a value as follows: 1: Mandatory; 2: Capable; 3: Disable. | [optional] 
**PpskSetting** | Pointer to [**SsidPpskSettingOpenApiVO**](SsidPpskSettingOpenApiVO.md) |  | [optional] 
**ProhibitWifiShare** | Pointer to **bool** | SSID prohibitWifiShare config status. True: enable, false: disable. | [optional] 
**PskSetting** | Pointer to [**SsidPskSettingOpenApiVO**](SsidPskSettingOpenApiVO.md) |  | [optional] 
**RateControl** | Pointer to [**SsidRateControlOpenApiVO**](SsidRateControlOpenApiVO.md) |  | [optional] 
**Security** | Pointer to **int32** | SSID security mode; Security should be a value as follows: 0: None; 2: WPA-Enterprise; 3: WPA-Personal; 4: PPSK without RADIUS; 5: PPSK with RADIUS. | [optional] 
**SsidId** | Pointer to **string** | SSID ID | [optional] 
**SsidRateLimit** | Pointer to [**RateLimitSettingOpenApiVO**](RateLimitSettingOpenApiVO.md) |  | [optional] 
**VlanEnable** | Pointer to **bool** | SSID VLAN config status. True: enable, false: disable. | [optional] 
**VlanId** | Pointer to **int32** | SSID VLAN ID. This field is required when Parameter [vlanEnable] is true; It should be within the range of 1–4094. | [optional] 
**VlanSetting** | Pointer to [**SsidVlanSettingOpenApiVO**](SsidVlanSettingOpenApiVO.md) |  | [optional] 
**WlanSchedule** | Pointer to [**SsidWlanScheduleOpenApiVO**](SsidWlanScheduleOpenApiVO.md) |  | [optional] 

## Methods

### NewSsidDetailOpenApiVO

`func NewSsidDetailOpenApiVO() *SsidDetailOpenApiVO`

NewSsidDetailOpenApiVO instantiates a new SsidDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsidDetailOpenApiVOWithDefaults

`func NewSsidDetailOpenApiVOWithDefaults() *SsidDetailOpenApiVO`

NewSsidDetailOpenApiVOWithDefaults instantiates a new SsidDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoWanAccess

`func (o *SsidDetailOpenApiVO) GetAutoWanAccess() bool`

GetAutoWanAccess returns the AutoWanAccess field if non-nil, zero value otherwise.

### GetAutoWanAccessOk

`func (o *SsidDetailOpenApiVO) GetAutoWanAccessOk() (*bool, bool)`

GetAutoWanAccessOk returns a tuple with the AutoWanAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoWanAccess

`func (o *SsidDetailOpenApiVO) SetAutoWanAccess(v bool)`

SetAutoWanAccess sets AutoWanAccess field to given value.

### HasAutoWanAccess

`func (o *SsidDetailOpenApiVO) HasAutoWanAccess() bool`

HasAutoWanAccess returns a boolean if a field has been set.

### GetBand

`func (o *SsidDetailOpenApiVO) GetBand() int32`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *SsidDetailOpenApiVO) GetBandOk() (*int32, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *SsidDetailOpenApiVO) SetBand(v int32)`

SetBand sets Band field to given value.

### HasBand

`func (o *SsidDetailOpenApiVO) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetBroadcast

`func (o *SsidDetailOpenApiVO) GetBroadcast() bool`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *SsidDetailOpenApiVO) GetBroadcastOk() (*bool, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *SsidDetailOpenApiVO) SetBroadcast(v bool)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *SsidDetailOpenApiVO) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.

### GetClientRateLimit

`func (o *SsidDetailOpenApiVO) GetClientRateLimit() RateLimitSettingOpenApiVO`

GetClientRateLimit returns the ClientRateLimit field if non-nil, zero value otherwise.

### GetClientRateLimitOk

`func (o *SsidDetailOpenApiVO) GetClientRateLimitOk() (*RateLimitSettingOpenApiVO, bool)`

GetClientRateLimitOk returns a tuple with the ClientRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRateLimit

`func (o *SsidDetailOpenApiVO) SetClientRateLimit(v RateLimitSettingOpenApiVO)`

SetClientRateLimit sets ClientRateLimit field to given value.

### HasClientRateLimit

`func (o *SsidDetailOpenApiVO) HasClientRateLimit() bool`

HasClientRateLimit returns a boolean if a field has been set.

### GetDeviceType

`func (o *SsidDetailOpenApiVO) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *SsidDetailOpenApiVO) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *SsidDetailOpenApiVO) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *SsidDetailOpenApiVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDhcpOption82

`func (o *SsidDetailOpenApiVO) GetDhcpOption82() SsidDhcpOptionOpenApiVO`

GetDhcpOption82 returns the DhcpOption82 field if non-nil, zero value otherwise.

### GetDhcpOption82Ok

`func (o *SsidDetailOpenApiVO) GetDhcpOption82Ok() (*SsidDhcpOptionOpenApiVO, bool)`

GetDhcpOption82Ok returns a tuple with the DhcpOption82 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOption82

`func (o *SsidDetailOpenApiVO) SetDhcpOption82(v SsidDhcpOptionOpenApiVO)`

SetDhcpOption82 sets DhcpOption82 field to given value.

### HasDhcpOption82

`func (o *SsidDetailOpenApiVO) HasDhcpOption82() bool`

HasDhcpOption82 returns a boolean if a field has been set.

### GetEnable11r

`func (o *SsidDetailOpenApiVO) GetEnable11r() bool`

GetEnable11r returns the Enable11r field if non-nil, zero value otherwise.

### GetEnable11rOk

`func (o *SsidDetailOpenApiVO) GetEnable11rOk() (*bool, bool)`

GetEnable11rOk returns a tuple with the Enable11r field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable11r

`func (o *SsidDetailOpenApiVO) SetEnable11r(v bool)`

SetEnable11r sets Enable11r field to given value.

### HasEnable11r

`func (o *SsidDetailOpenApiVO) HasEnable11r() bool`

HasEnable11r returns a boolean if a field has been set.

### GetEntSetting

`func (o *SsidDetailOpenApiVO) GetEntSetting() SsidEnterpriseSettingOpenApiVO`

GetEntSetting returns the EntSetting field if non-nil, zero value otherwise.

### GetEntSettingOk

`func (o *SsidDetailOpenApiVO) GetEntSettingOk() (*SsidEnterpriseSettingOpenApiVO, bool)`

GetEntSettingOk returns a tuple with the EntSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntSetting

`func (o *SsidDetailOpenApiVO) SetEntSetting(v SsidEnterpriseSettingOpenApiVO)`

SetEntSetting sets EntSetting field to given value.

### HasEntSetting

`func (o *SsidDetailOpenApiVO) HasEntSetting() bool`

HasEntSetting returns a boolean if a field has been set.

### GetGuestNetEnable

`func (o *SsidDetailOpenApiVO) GetGuestNetEnable() bool`

GetGuestNetEnable returns the GuestNetEnable field if non-nil, zero value otherwise.

### GetGuestNetEnableOk

`func (o *SsidDetailOpenApiVO) GetGuestNetEnableOk() (*bool, bool)`

GetGuestNetEnableOk returns a tuple with the GuestNetEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestNetEnable

`func (o *SsidDetailOpenApiVO) SetGuestNetEnable(v bool)`

SetGuestNetEnable sets GuestNetEnable field to given value.

### HasGuestNetEnable

`func (o *SsidDetailOpenApiVO) HasGuestNetEnable() bool`

HasGuestNetEnable returns a boolean if a field has been set.

### GetHidePwd

`func (o *SsidDetailOpenApiVO) GetHidePwd() bool`

GetHidePwd returns the HidePwd field if non-nil, zero value otherwise.

### GetHidePwdOk

`func (o *SsidDetailOpenApiVO) GetHidePwdOk() (*bool, bool)`

GetHidePwdOk returns a tuple with the HidePwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePwd

`func (o *SsidDetailOpenApiVO) SetHidePwd(v bool)`

SetHidePwd sets HidePwd field to given value.

### HasHidePwd

`func (o *SsidDetailOpenApiVO) HasHidePwd() bool`

HasHidePwd returns a boolean if a field has been set.

### GetHotspotV2Setting

`func (o *SsidDetailOpenApiVO) GetHotspotV2Setting() HotspotV2SettingOpenApiVO`

GetHotspotV2Setting returns the HotspotV2Setting field if non-nil, zero value otherwise.

### GetHotspotV2SettingOk

`func (o *SsidDetailOpenApiVO) GetHotspotV2SettingOk() (*HotspotV2SettingOpenApiVO, bool)`

GetHotspotV2SettingOk returns a tuple with the HotspotV2Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotspotV2Setting

`func (o *SsidDetailOpenApiVO) SetHotspotV2Setting(v HotspotV2SettingOpenApiVO)`

SetHotspotV2Setting sets HotspotV2Setting field to given value.

### HasHotspotV2Setting

`func (o *SsidDetailOpenApiVO) HasHotspotV2Setting() bool`

HasHotspotV2Setting returns a boolean if a field has been set.

### GetMacFilter

`func (o *SsidDetailOpenApiVO) GetMacFilter() SsidMacFilterOpenApiVO`

GetMacFilter returns the MacFilter field if non-nil, zero value otherwise.

### GetMacFilterOk

`func (o *SsidDetailOpenApiVO) GetMacFilterOk() (*SsidMacFilterOpenApiVO, bool)`

GetMacFilterOk returns a tuple with the MacFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacFilter

`func (o *SsidDetailOpenApiVO) SetMacFilter(v SsidMacFilterOpenApiVO)`

SetMacFilter sets MacFilter field to given value.

### HasMacFilter

`func (o *SsidDetailOpenApiVO) HasMacFilter() bool`

HasMacFilter returns a boolean if a field has been set.

### GetMloEnable

`func (o *SsidDetailOpenApiVO) GetMloEnable() bool`

GetMloEnable returns the MloEnable field if non-nil, zero value otherwise.

### GetMloEnableOk

`func (o *SsidDetailOpenApiVO) GetMloEnableOk() (*bool, bool)`

GetMloEnableOk returns a tuple with the MloEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMloEnable

`func (o *SsidDetailOpenApiVO) SetMloEnable(v bool)`

SetMloEnable sets MloEnable field to given value.

### HasMloEnable

`func (o *SsidDetailOpenApiVO) HasMloEnable() bool`

HasMloEnable returns a boolean if a field has been set.

### GetMultiCast

`func (o *SsidDetailOpenApiVO) GetMultiCast() SsidMultiCastOpenApiVO`

GetMultiCast returns the MultiCast field if non-nil, zero value otherwise.

### GetMultiCastOk

`func (o *SsidDetailOpenApiVO) GetMultiCastOk() (*SsidMultiCastOpenApiVO, bool)`

GetMultiCastOk returns a tuple with the MultiCast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiCast

`func (o *SsidDetailOpenApiVO) SetMultiCast(v SsidMultiCastOpenApiVO)`

SetMultiCast sets MultiCast field to given value.

### HasMultiCast

`func (o *SsidDetailOpenApiVO) HasMultiCast() bool`

HasMultiCast returns a boolean if a field has been set.

### GetName

`func (o *SsidDetailOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SsidDetailOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SsidDetailOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SsidDetailOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOweEnable

`func (o *SsidDetailOpenApiVO) GetOweEnable() bool`

GetOweEnable returns the OweEnable field if non-nil, zero value otherwise.

### GetOweEnableOk

`func (o *SsidDetailOpenApiVO) GetOweEnableOk() (*bool, bool)`

GetOweEnableOk returns a tuple with the OweEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOweEnable

`func (o *SsidDetailOpenApiVO) SetOweEnable(v bool)`

SetOweEnable sets OweEnable field to given value.

### HasOweEnable

`func (o *SsidDetailOpenApiVO) HasOweEnable() bool`

HasOweEnable returns a boolean if a field has been set.

### GetPmfMode

`func (o *SsidDetailOpenApiVO) GetPmfMode() int32`

GetPmfMode returns the PmfMode field if non-nil, zero value otherwise.

### GetPmfModeOk

`func (o *SsidDetailOpenApiVO) GetPmfModeOk() (*int32, bool)`

GetPmfModeOk returns a tuple with the PmfMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmfMode

`func (o *SsidDetailOpenApiVO) SetPmfMode(v int32)`

SetPmfMode sets PmfMode field to given value.

### HasPmfMode

`func (o *SsidDetailOpenApiVO) HasPmfMode() bool`

HasPmfMode returns a boolean if a field has been set.

### GetPpskSetting

`func (o *SsidDetailOpenApiVO) GetPpskSetting() SsidPpskSettingOpenApiVO`

GetPpskSetting returns the PpskSetting field if non-nil, zero value otherwise.

### GetPpskSettingOk

`func (o *SsidDetailOpenApiVO) GetPpskSettingOk() (*SsidPpskSettingOpenApiVO, bool)`

GetPpskSettingOk returns a tuple with the PpskSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpskSetting

`func (o *SsidDetailOpenApiVO) SetPpskSetting(v SsidPpskSettingOpenApiVO)`

SetPpskSetting sets PpskSetting field to given value.

### HasPpskSetting

`func (o *SsidDetailOpenApiVO) HasPpskSetting() bool`

HasPpskSetting returns a boolean if a field has been set.

### GetProhibitWifiShare

`func (o *SsidDetailOpenApiVO) GetProhibitWifiShare() bool`

GetProhibitWifiShare returns the ProhibitWifiShare field if non-nil, zero value otherwise.

### GetProhibitWifiShareOk

`func (o *SsidDetailOpenApiVO) GetProhibitWifiShareOk() (*bool, bool)`

GetProhibitWifiShareOk returns a tuple with the ProhibitWifiShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProhibitWifiShare

`func (o *SsidDetailOpenApiVO) SetProhibitWifiShare(v bool)`

SetProhibitWifiShare sets ProhibitWifiShare field to given value.

### HasProhibitWifiShare

`func (o *SsidDetailOpenApiVO) HasProhibitWifiShare() bool`

HasProhibitWifiShare returns a boolean if a field has been set.

### GetPskSetting

`func (o *SsidDetailOpenApiVO) GetPskSetting() SsidPskSettingOpenApiVO`

GetPskSetting returns the PskSetting field if non-nil, zero value otherwise.

### GetPskSettingOk

`func (o *SsidDetailOpenApiVO) GetPskSettingOk() (*SsidPskSettingOpenApiVO, bool)`

GetPskSettingOk returns a tuple with the PskSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPskSetting

`func (o *SsidDetailOpenApiVO) SetPskSetting(v SsidPskSettingOpenApiVO)`

SetPskSetting sets PskSetting field to given value.

### HasPskSetting

`func (o *SsidDetailOpenApiVO) HasPskSetting() bool`

HasPskSetting returns a boolean if a field has been set.

### GetRateControl

`func (o *SsidDetailOpenApiVO) GetRateControl() SsidRateControlOpenApiVO`

GetRateControl returns the RateControl field if non-nil, zero value otherwise.

### GetRateControlOk

`func (o *SsidDetailOpenApiVO) GetRateControlOk() (*SsidRateControlOpenApiVO, bool)`

GetRateControlOk returns a tuple with the RateControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateControl

`func (o *SsidDetailOpenApiVO) SetRateControl(v SsidRateControlOpenApiVO)`

SetRateControl sets RateControl field to given value.

### HasRateControl

`func (o *SsidDetailOpenApiVO) HasRateControl() bool`

HasRateControl returns a boolean if a field has been set.

### GetSecurity

`func (o *SsidDetailOpenApiVO) GetSecurity() int32`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *SsidDetailOpenApiVO) GetSecurityOk() (*int32, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *SsidDetailOpenApiVO) SetSecurity(v int32)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *SsidDetailOpenApiVO) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetSsidId

`func (o *SsidDetailOpenApiVO) GetSsidId() string`

GetSsidId returns the SsidId field if non-nil, zero value otherwise.

### GetSsidIdOk

`func (o *SsidDetailOpenApiVO) GetSsidIdOk() (*string, bool)`

GetSsidIdOk returns a tuple with the SsidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidId

`func (o *SsidDetailOpenApiVO) SetSsidId(v string)`

SetSsidId sets SsidId field to given value.

### HasSsidId

`func (o *SsidDetailOpenApiVO) HasSsidId() bool`

HasSsidId returns a boolean if a field has been set.

### GetSsidRateLimit

`func (o *SsidDetailOpenApiVO) GetSsidRateLimit() RateLimitSettingOpenApiVO`

GetSsidRateLimit returns the SsidRateLimit field if non-nil, zero value otherwise.

### GetSsidRateLimitOk

`func (o *SsidDetailOpenApiVO) GetSsidRateLimitOk() (*RateLimitSettingOpenApiVO, bool)`

GetSsidRateLimitOk returns a tuple with the SsidRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidRateLimit

`func (o *SsidDetailOpenApiVO) SetSsidRateLimit(v RateLimitSettingOpenApiVO)`

SetSsidRateLimit sets SsidRateLimit field to given value.

### HasSsidRateLimit

`func (o *SsidDetailOpenApiVO) HasSsidRateLimit() bool`

HasSsidRateLimit returns a boolean if a field has been set.

### GetVlanEnable

`func (o *SsidDetailOpenApiVO) GetVlanEnable() bool`

GetVlanEnable returns the VlanEnable field if non-nil, zero value otherwise.

### GetVlanEnableOk

`func (o *SsidDetailOpenApiVO) GetVlanEnableOk() (*bool, bool)`

GetVlanEnableOk returns a tuple with the VlanEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanEnable

`func (o *SsidDetailOpenApiVO) SetVlanEnable(v bool)`

SetVlanEnable sets VlanEnable field to given value.

### HasVlanEnable

`func (o *SsidDetailOpenApiVO) HasVlanEnable() bool`

HasVlanEnable returns a boolean if a field has been set.

### GetVlanId

`func (o *SsidDetailOpenApiVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *SsidDetailOpenApiVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *SsidDetailOpenApiVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *SsidDetailOpenApiVO) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlanSetting

`func (o *SsidDetailOpenApiVO) GetVlanSetting() SsidVlanSettingOpenApiVO`

GetVlanSetting returns the VlanSetting field if non-nil, zero value otherwise.

### GetVlanSettingOk

`func (o *SsidDetailOpenApiVO) GetVlanSettingOk() (*SsidVlanSettingOpenApiVO, bool)`

GetVlanSettingOk returns a tuple with the VlanSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanSetting

`func (o *SsidDetailOpenApiVO) SetVlanSetting(v SsidVlanSettingOpenApiVO)`

SetVlanSetting sets VlanSetting field to given value.

### HasVlanSetting

`func (o *SsidDetailOpenApiVO) HasVlanSetting() bool`

HasVlanSetting returns a boolean if a field has been set.

### GetWlanSchedule

`func (o *SsidDetailOpenApiVO) GetWlanSchedule() SsidWlanScheduleOpenApiVO`

GetWlanSchedule returns the WlanSchedule field if non-nil, zero value otherwise.

### GetWlanScheduleOk

`func (o *SsidDetailOpenApiVO) GetWlanScheduleOk() (*SsidWlanScheduleOpenApiVO, bool)`

GetWlanScheduleOk returns a tuple with the WlanSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanSchedule

`func (o *SsidDetailOpenApiVO) SetWlanSchedule(v SsidWlanScheduleOpenApiVO)`

SetWlanSchedule sets WlanSchedule field to given value.

### HasWlanSchedule

`func (o *SsidDetailOpenApiVO) HasWlanSchedule() bool`

HasWlanSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


