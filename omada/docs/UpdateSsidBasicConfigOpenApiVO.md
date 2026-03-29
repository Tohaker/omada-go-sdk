# UpdateSsidBasicConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoWanAccess** | Pointer to **bool** | Whether to enable auto wan access. True: enable, false: disable. | [optional] 
**Band** | **int32** | SSID band. The lowest bit indicates whether 2.4G is included; the second lowest bit indicates whether 5G is included; the third lowest bit indicates whether 6G is included; 1 means included while 0 means not included. For example, 7(111) means that 2G/5G/6G are enabled; 1(001) means that 2G is enabled. (When 5G is included，it means 5G/5G1/5G2 are enabled.) | 
**Broadcast** | **bool** | SSID broadcast config status. True: enable, false: disable. | 
**Enable11r** | **bool** | SSID 802.11r config status. True: enable, false: disable. | 
**EntSetting** | Pointer to [**SsidEnterpriseSettingOpenApiVO**](SsidEnterpriseSettingOpenApiVO.md) |  | [optional] 
**GreEnable** | Pointer to **bool** | SSID EoGre Tunnel config status. True: enable, false: disable; This configuration can be enabled only when the [VPN - EoGre Tunnel] global config is enabled. | [optional] 
**GuestNetEnable** | **bool** | SSID guest network config status. True: enable, false: disable. | 
**HidePwd** | Pointer to **bool** | If this field is true, the SSID password will be hidden. | [optional] 
**MloEnable** | **bool** | SSID MLO config status. True: enable, false: disable. | 
**Name** | **string** | SSID name. It should contain 1 to 32 UTF-8 characters. | 
**OweEnable** | Pointer to **bool** | Opportunistic Wireless Encryption, also known as Enhanced Open, is a certification provided by the Wi-Fi Alliance as part of the WPA3 wireless security standard. OWE will enable two wireless VAPs per radio, one for access of OWE-supported stations, and one for access of other stations. An SSID with OWE enabled will be counted as two SSID entries for 2G and 5G. Only for security is None and band contains 2.4G or 5G. | [optional] 
**PmfMode** | **int32** | SSID PMF mode. It should be a value as follows: 1: Mandatory; 2: Capable; 3: Disable. | 
**PpskSetting** | Pointer to [**SsidPpskSettingOpenApiVO**](SsidPpskSettingOpenApiVO.md) |  | [optional] 
**ProhibitWifiShare** | Pointer to **bool** | SSID prohibitWifiShare config status. True: enable, false: disable. | [optional] 
**PskSetting** | Pointer to [**SsidPskSettingOpenApiVO**](SsidPskSettingOpenApiVO.md) |  | [optional] 
**Security** | **int32** | SSID security mode; Security should be a value as follows: 0: None; 2: WPA-Enterprise; 3: WPA-Personal; 4: PPSK without RADIUS; 5: PPSK with RADIUS. | 
**VlanEnable** | **bool** | SSID VLAN config status. True: enable, false: disable. | 
**VlanId** | Pointer to **int32** | SSID VLAN ID. This field is required when Parameter [vlanEnable] is true; It should be within the range of 1–4094.  | [optional] 
**VlanSetting** | Pointer to [**SsidVlanSettingOpenApiVO**](SsidVlanSettingOpenApiVO.md) |  | [optional] 

## Methods

### NewUpdateSsidBasicConfigOpenApiVO

`func NewUpdateSsidBasicConfigOpenApiVO(band int32, broadcast bool, enable11r bool, guestNetEnable bool, mloEnable bool, name string, pmfMode int32, security int32, vlanEnable bool, ) *UpdateSsidBasicConfigOpenApiVO`

NewUpdateSsidBasicConfigOpenApiVO instantiates a new UpdateSsidBasicConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSsidBasicConfigOpenApiVOWithDefaults

`func NewUpdateSsidBasicConfigOpenApiVOWithDefaults() *UpdateSsidBasicConfigOpenApiVO`

NewUpdateSsidBasicConfigOpenApiVOWithDefaults instantiates a new UpdateSsidBasicConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoWanAccess

`func (o *UpdateSsidBasicConfigOpenApiVO) GetAutoWanAccess() bool`

GetAutoWanAccess returns the AutoWanAccess field if non-nil, zero value otherwise.

### GetAutoWanAccessOk

`func (o *UpdateSsidBasicConfigOpenApiVO) GetAutoWanAccessOk() (*bool, bool)`

GetAutoWanAccessOk returns a tuple with the AutoWanAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoWanAccess

`func (o *UpdateSsidBasicConfigOpenApiVO) SetAutoWanAccess(v bool)`

SetAutoWanAccess sets AutoWanAccess field to given value.

### HasAutoWanAccess

`func (o *UpdateSsidBasicConfigOpenApiVO) HasAutoWanAccess() bool`

HasAutoWanAccess returns a boolean if a field has been set.

### GetBand

`func (o *UpdateSsidBasicConfigOpenApiVO) GetBand() int32`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *UpdateSsidBasicConfigOpenApiVO) GetBandOk() (*int32, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *UpdateSsidBasicConfigOpenApiVO) SetBand(v int32)`

SetBand sets Band field to given value.


### GetBroadcast

`func (o *UpdateSsidBasicConfigOpenApiVO) GetBroadcast() bool`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *UpdateSsidBasicConfigOpenApiVO) GetBroadcastOk() (*bool, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *UpdateSsidBasicConfigOpenApiVO) SetBroadcast(v bool)`

SetBroadcast sets Broadcast field to given value.


### GetEnable11r

`func (o *UpdateSsidBasicConfigOpenApiVO) GetEnable11r() bool`

GetEnable11r returns the Enable11r field if non-nil, zero value otherwise.

### GetEnable11rOk

`func (o *UpdateSsidBasicConfigOpenApiVO) GetEnable11rOk() (*bool, bool)`

GetEnable11rOk returns a tuple with the Enable11r field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable11r

`func (o *UpdateSsidBasicConfigOpenApiVO) SetEnable11r(v bool)`

SetEnable11r sets Enable11r field to given value.


### GetEntSetting

`func (o *UpdateSsidBasicConfigOpenApiVO) GetEntSetting() SsidEnterpriseSettingOpenApiVO`

GetEntSetting returns the EntSetting field if non-nil, zero value otherwise.

### GetEntSettingOk

`func (o *UpdateSsidBasicConfigOpenApiVO) GetEntSettingOk() (*SsidEnterpriseSettingOpenApiVO, bool)`

GetEntSettingOk returns a tuple with the EntSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntSetting

`func (o *UpdateSsidBasicConfigOpenApiVO) SetEntSetting(v SsidEnterpriseSettingOpenApiVO)`

SetEntSetting sets EntSetting field to given value.

### HasEntSetting

`func (o *UpdateSsidBasicConfigOpenApiVO) HasEntSetting() bool`

HasEntSetting returns a boolean if a field has been set.

### GetGreEnable

`func (o *UpdateSsidBasicConfigOpenApiVO) GetGreEnable() bool`

GetGreEnable returns the GreEnable field if non-nil, zero value otherwise.

### GetGreEnableOk

`func (o *UpdateSsidBasicConfigOpenApiVO) GetGreEnableOk() (*bool, bool)`

GetGreEnableOk returns a tuple with the GreEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreEnable

`func (o *UpdateSsidBasicConfigOpenApiVO) SetGreEnable(v bool)`

SetGreEnable sets GreEnable field to given value.

### HasGreEnable

`func (o *UpdateSsidBasicConfigOpenApiVO) HasGreEnable() bool`

HasGreEnable returns a boolean if a field has been set.

### GetGuestNetEnable

`func (o *UpdateSsidBasicConfigOpenApiVO) GetGuestNetEnable() bool`

GetGuestNetEnable returns the GuestNetEnable field if non-nil, zero value otherwise.

### GetGuestNetEnableOk

`func (o *UpdateSsidBasicConfigOpenApiVO) GetGuestNetEnableOk() (*bool, bool)`

GetGuestNetEnableOk returns a tuple with the GuestNetEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestNetEnable

`func (o *UpdateSsidBasicConfigOpenApiVO) SetGuestNetEnable(v bool)`

SetGuestNetEnable sets GuestNetEnable field to given value.


### GetHidePwd

`func (o *UpdateSsidBasicConfigOpenApiVO) GetHidePwd() bool`

GetHidePwd returns the HidePwd field if non-nil, zero value otherwise.

### GetHidePwdOk

`func (o *UpdateSsidBasicConfigOpenApiVO) GetHidePwdOk() (*bool, bool)`

GetHidePwdOk returns a tuple with the HidePwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePwd

`func (o *UpdateSsidBasicConfigOpenApiVO) SetHidePwd(v bool)`

SetHidePwd sets HidePwd field to given value.

### HasHidePwd

`func (o *UpdateSsidBasicConfigOpenApiVO) HasHidePwd() bool`

HasHidePwd returns a boolean if a field has been set.

### GetMloEnable

`func (o *UpdateSsidBasicConfigOpenApiVO) GetMloEnable() bool`

GetMloEnable returns the MloEnable field if non-nil, zero value otherwise.

### GetMloEnableOk

`func (o *UpdateSsidBasicConfigOpenApiVO) GetMloEnableOk() (*bool, bool)`

GetMloEnableOk returns a tuple with the MloEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMloEnable

`func (o *UpdateSsidBasicConfigOpenApiVO) SetMloEnable(v bool)`

SetMloEnable sets MloEnable field to given value.


### GetName

`func (o *UpdateSsidBasicConfigOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSsidBasicConfigOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSsidBasicConfigOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetOweEnable

`func (o *UpdateSsidBasicConfigOpenApiVO) GetOweEnable() bool`

GetOweEnable returns the OweEnable field if non-nil, zero value otherwise.

### GetOweEnableOk

`func (o *UpdateSsidBasicConfigOpenApiVO) GetOweEnableOk() (*bool, bool)`

GetOweEnableOk returns a tuple with the OweEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOweEnable

`func (o *UpdateSsidBasicConfigOpenApiVO) SetOweEnable(v bool)`

SetOweEnable sets OweEnable field to given value.

### HasOweEnable

`func (o *UpdateSsidBasicConfigOpenApiVO) HasOweEnable() bool`

HasOweEnable returns a boolean if a field has been set.

### GetPmfMode

`func (o *UpdateSsidBasicConfigOpenApiVO) GetPmfMode() int32`

GetPmfMode returns the PmfMode field if non-nil, zero value otherwise.

### GetPmfModeOk

`func (o *UpdateSsidBasicConfigOpenApiVO) GetPmfModeOk() (*int32, bool)`

GetPmfModeOk returns a tuple with the PmfMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmfMode

`func (o *UpdateSsidBasicConfigOpenApiVO) SetPmfMode(v int32)`

SetPmfMode sets PmfMode field to given value.


### GetPpskSetting

`func (o *UpdateSsidBasicConfigOpenApiVO) GetPpskSetting() SsidPpskSettingOpenApiVO`

GetPpskSetting returns the PpskSetting field if non-nil, zero value otherwise.

### GetPpskSettingOk

`func (o *UpdateSsidBasicConfigOpenApiVO) GetPpskSettingOk() (*SsidPpskSettingOpenApiVO, bool)`

GetPpskSettingOk returns a tuple with the PpskSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpskSetting

`func (o *UpdateSsidBasicConfigOpenApiVO) SetPpskSetting(v SsidPpskSettingOpenApiVO)`

SetPpskSetting sets PpskSetting field to given value.

### HasPpskSetting

`func (o *UpdateSsidBasicConfigOpenApiVO) HasPpskSetting() bool`

HasPpskSetting returns a boolean if a field has been set.

### GetProhibitWifiShare

`func (o *UpdateSsidBasicConfigOpenApiVO) GetProhibitWifiShare() bool`

GetProhibitWifiShare returns the ProhibitWifiShare field if non-nil, zero value otherwise.

### GetProhibitWifiShareOk

`func (o *UpdateSsidBasicConfigOpenApiVO) GetProhibitWifiShareOk() (*bool, bool)`

GetProhibitWifiShareOk returns a tuple with the ProhibitWifiShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProhibitWifiShare

`func (o *UpdateSsidBasicConfigOpenApiVO) SetProhibitWifiShare(v bool)`

SetProhibitWifiShare sets ProhibitWifiShare field to given value.

### HasProhibitWifiShare

`func (o *UpdateSsidBasicConfigOpenApiVO) HasProhibitWifiShare() bool`

HasProhibitWifiShare returns a boolean if a field has been set.

### GetPskSetting

`func (o *UpdateSsidBasicConfigOpenApiVO) GetPskSetting() SsidPskSettingOpenApiVO`

GetPskSetting returns the PskSetting field if non-nil, zero value otherwise.

### GetPskSettingOk

`func (o *UpdateSsidBasicConfigOpenApiVO) GetPskSettingOk() (*SsidPskSettingOpenApiVO, bool)`

GetPskSettingOk returns a tuple with the PskSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPskSetting

`func (o *UpdateSsidBasicConfigOpenApiVO) SetPskSetting(v SsidPskSettingOpenApiVO)`

SetPskSetting sets PskSetting field to given value.

### HasPskSetting

`func (o *UpdateSsidBasicConfigOpenApiVO) HasPskSetting() bool`

HasPskSetting returns a boolean if a field has been set.

### GetSecurity

`func (o *UpdateSsidBasicConfigOpenApiVO) GetSecurity() int32`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *UpdateSsidBasicConfigOpenApiVO) GetSecurityOk() (*int32, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *UpdateSsidBasicConfigOpenApiVO) SetSecurity(v int32)`

SetSecurity sets Security field to given value.


### GetVlanEnable

`func (o *UpdateSsidBasicConfigOpenApiVO) GetVlanEnable() bool`

GetVlanEnable returns the VlanEnable field if non-nil, zero value otherwise.

### GetVlanEnableOk

`func (o *UpdateSsidBasicConfigOpenApiVO) GetVlanEnableOk() (*bool, bool)`

GetVlanEnableOk returns a tuple with the VlanEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanEnable

`func (o *UpdateSsidBasicConfigOpenApiVO) SetVlanEnable(v bool)`

SetVlanEnable sets VlanEnable field to given value.


### GetVlanId

`func (o *UpdateSsidBasicConfigOpenApiVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *UpdateSsidBasicConfigOpenApiVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *UpdateSsidBasicConfigOpenApiVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *UpdateSsidBasicConfigOpenApiVO) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlanSetting

`func (o *UpdateSsidBasicConfigOpenApiVO) GetVlanSetting() SsidVlanSettingOpenApiVO`

GetVlanSetting returns the VlanSetting field if non-nil, zero value otherwise.

### GetVlanSettingOk

`func (o *UpdateSsidBasicConfigOpenApiVO) GetVlanSettingOk() (*SsidVlanSettingOpenApiVO, bool)`

GetVlanSettingOk returns a tuple with the VlanSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanSetting

`func (o *UpdateSsidBasicConfigOpenApiVO) SetVlanSetting(v SsidVlanSettingOpenApiVO)`

SetVlanSetting sets VlanSetting field to given value.

### HasVlanSetting

`func (o *UpdateSsidBasicConfigOpenApiVO) HasVlanSetting() bool`

HasVlanSetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


