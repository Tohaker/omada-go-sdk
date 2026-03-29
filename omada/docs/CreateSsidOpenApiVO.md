# CreateSsidOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Band** | **int32** | SSID band. The lowest bit indicates whether 2.4G is included; the second lowest bit indicates whether 5G is included; the third lowest bit indicates whether 6G is included; 1 means included while 0 means not included. For example, 7(111) means that 2G/5G/6G are enabled; 1(001) means that 2G is enabled. (When 5G is included，it means 5G/5G1/5G2 are enabled.) | 
**Broadcast** | **bool** | SSID broadcast config status. True: enable, false: disable. | 
**DeviceType** | **int32** | SSID device type, identify which devices this SSID will take effect; The lowest bit indicates whether [EAP] is included, the second low bit indicates whether [Gateway] is included, 1 means included while 0 means not included; For example, 3(11) means that EAP/Gateway is enabled, 1(01) means that EAP is enabled. | 
**Enable11r** | **bool** | SSID 802.11r config status. True: enable, false: disable. | 
**EntSetting** | Pointer to [**SsidEnterpriseSettingOpenApiVO**](SsidEnterpriseSettingOpenApiVO.md) |  | [optional] 
**GreEnable** | Pointer to **bool** | SSID EoGre Tunnel config status. True: enable, false: disable. This configuration can be enabled only when the [VPN - EoGre Tunnel] global config is enabled;(This configuration applies to the Pro Site of the Omada Pro Controller only). | [optional] 
**GuestNetEnable** | **bool** | SSID guest network config status. True: enable, false: disable. | 
**HidePwd** | **bool** | If this field is true, the SSID password will be hidden. | 
**MloEnable** | **bool** | SSID MLO config status. True: enable, false: disable. | 
**Name** | **string** | SSID name. It should contain 1 to 32 UTF-8 characters. | 
**PmfMode** | **int32** | SSID PMF mode. It should be a value as follows: 1: Mandatory; 2: Capable; 3: Disable. | 
**PpskSetting** | Pointer to [**SsidPpskSettingOpenApiVO**](SsidPpskSettingOpenApiVO.md) |  | [optional] 
**ProhibitWifiShare** | Pointer to **bool** | SSID prohibitWifiShare config status. True: enable, false: disable. | [optional] 
**PskSetting** | Pointer to [**SsidPskSettingOpenApiVO**](SsidPskSettingOpenApiVO.md) |  | [optional] 
**Security** | **int32** | SSID security mode; Security should be a value as follows: 0: None; 2: WPA-Enterprise; 3: WPA-Personal; 4: PPSK without RADIUS; 5: PPSK with RADIUS. | 
**VlanEnable** | **bool** | SSID VLAN config status. True: enable, false: disable. | 
**VlanId** | Pointer to **int32** | SSID VLAN ID. This field is required when Parameter [vlanEnable] is true; It should be within the range of 1–4094. If the field vlanSetting is entered, this field must be null. | [optional] 
**VlanSetting** | Pointer to [**SsidVlanSettingOpenApiVO**](SsidVlanSettingOpenApiVO.md) |  | [optional] 

## Methods

### NewCreateSsidOpenApiVO

`func NewCreateSsidOpenApiVO(band int32, broadcast bool, deviceType int32, enable11r bool, guestNetEnable bool, hidePwd bool, mloEnable bool, name string, pmfMode int32, security int32, vlanEnable bool, ) *CreateSsidOpenApiVO`

NewCreateSsidOpenApiVO instantiates a new CreateSsidOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSsidOpenApiVOWithDefaults

`func NewCreateSsidOpenApiVOWithDefaults() *CreateSsidOpenApiVO`

NewCreateSsidOpenApiVOWithDefaults instantiates a new CreateSsidOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBand

`func (o *CreateSsidOpenApiVO) GetBand() int32`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *CreateSsidOpenApiVO) GetBandOk() (*int32, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *CreateSsidOpenApiVO) SetBand(v int32)`

SetBand sets Band field to given value.


### GetBroadcast

`func (o *CreateSsidOpenApiVO) GetBroadcast() bool`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *CreateSsidOpenApiVO) GetBroadcastOk() (*bool, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *CreateSsidOpenApiVO) SetBroadcast(v bool)`

SetBroadcast sets Broadcast field to given value.


### GetDeviceType

`func (o *CreateSsidOpenApiVO) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *CreateSsidOpenApiVO) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *CreateSsidOpenApiVO) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.


### GetEnable11r

`func (o *CreateSsidOpenApiVO) GetEnable11r() bool`

GetEnable11r returns the Enable11r field if non-nil, zero value otherwise.

### GetEnable11rOk

`func (o *CreateSsidOpenApiVO) GetEnable11rOk() (*bool, bool)`

GetEnable11rOk returns a tuple with the Enable11r field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable11r

`func (o *CreateSsidOpenApiVO) SetEnable11r(v bool)`

SetEnable11r sets Enable11r field to given value.


### GetEntSetting

`func (o *CreateSsidOpenApiVO) GetEntSetting() SsidEnterpriseSettingOpenApiVO`

GetEntSetting returns the EntSetting field if non-nil, zero value otherwise.

### GetEntSettingOk

`func (o *CreateSsidOpenApiVO) GetEntSettingOk() (*SsidEnterpriseSettingOpenApiVO, bool)`

GetEntSettingOk returns a tuple with the EntSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntSetting

`func (o *CreateSsidOpenApiVO) SetEntSetting(v SsidEnterpriseSettingOpenApiVO)`

SetEntSetting sets EntSetting field to given value.

### HasEntSetting

`func (o *CreateSsidOpenApiVO) HasEntSetting() bool`

HasEntSetting returns a boolean if a field has been set.

### GetGreEnable

`func (o *CreateSsidOpenApiVO) GetGreEnable() bool`

GetGreEnable returns the GreEnable field if non-nil, zero value otherwise.

### GetGreEnableOk

`func (o *CreateSsidOpenApiVO) GetGreEnableOk() (*bool, bool)`

GetGreEnableOk returns a tuple with the GreEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreEnable

`func (o *CreateSsidOpenApiVO) SetGreEnable(v bool)`

SetGreEnable sets GreEnable field to given value.

### HasGreEnable

`func (o *CreateSsidOpenApiVO) HasGreEnable() bool`

HasGreEnable returns a boolean if a field has been set.

### GetGuestNetEnable

`func (o *CreateSsidOpenApiVO) GetGuestNetEnable() bool`

GetGuestNetEnable returns the GuestNetEnable field if non-nil, zero value otherwise.

### GetGuestNetEnableOk

`func (o *CreateSsidOpenApiVO) GetGuestNetEnableOk() (*bool, bool)`

GetGuestNetEnableOk returns a tuple with the GuestNetEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestNetEnable

`func (o *CreateSsidOpenApiVO) SetGuestNetEnable(v bool)`

SetGuestNetEnable sets GuestNetEnable field to given value.


### GetHidePwd

`func (o *CreateSsidOpenApiVO) GetHidePwd() bool`

GetHidePwd returns the HidePwd field if non-nil, zero value otherwise.

### GetHidePwdOk

`func (o *CreateSsidOpenApiVO) GetHidePwdOk() (*bool, bool)`

GetHidePwdOk returns a tuple with the HidePwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePwd

`func (o *CreateSsidOpenApiVO) SetHidePwd(v bool)`

SetHidePwd sets HidePwd field to given value.


### GetMloEnable

`func (o *CreateSsidOpenApiVO) GetMloEnable() bool`

GetMloEnable returns the MloEnable field if non-nil, zero value otherwise.

### GetMloEnableOk

`func (o *CreateSsidOpenApiVO) GetMloEnableOk() (*bool, bool)`

GetMloEnableOk returns a tuple with the MloEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMloEnable

`func (o *CreateSsidOpenApiVO) SetMloEnable(v bool)`

SetMloEnable sets MloEnable field to given value.


### GetName

`func (o *CreateSsidOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSsidOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSsidOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetPmfMode

`func (o *CreateSsidOpenApiVO) GetPmfMode() int32`

GetPmfMode returns the PmfMode field if non-nil, zero value otherwise.

### GetPmfModeOk

`func (o *CreateSsidOpenApiVO) GetPmfModeOk() (*int32, bool)`

GetPmfModeOk returns a tuple with the PmfMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmfMode

`func (o *CreateSsidOpenApiVO) SetPmfMode(v int32)`

SetPmfMode sets PmfMode field to given value.


### GetPpskSetting

`func (o *CreateSsidOpenApiVO) GetPpskSetting() SsidPpskSettingOpenApiVO`

GetPpskSetting returns the PpskSetting field if non-nil, zero value otherwise.

### GetPpskSettingOk

`func (o *CreateSsidOpenApiVO) GetPpskSettingOk() (*SsidPpskSettingOpenApiVO, bool)`

GetPpskSettingOk returns a tuple with the PpskSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpskSetting

`func (o *CreateSsidOpenApiVO) SetPpskSetting(v SsidPpskSettingOpenApiVO)`

SetPpskSetting sets PpskSetting field to given value.

### HasPpskSetting

`func (o *CreateSsidOpenApiVO) HasPpskSetting() bool`

HasPpskSetting returns a boolean if a field has been set.

### GetProhibitWifiShare

`func (o *CreateSsidOpenApiVO) GetProhibitWifiShare() bool`

GetProhibitWifiShare returns the ProhibitWifiShare field if non-nil, zero value otherwise.

### GetProhibitWifiShareOk

`func (o *CreateSsidOpenApiVO) GetProhibitWifiShareOk() (*bool, bool)`

GetProhibitWifiShareOk returns a tuple with the ProhibitWifiShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProhibitWifiShare

`func (o *CreateSsidOpenApiVO) SetProhibitWifiShare(v bool)`

SetProhibitWifiShare sets ProhibitWifiShare field to given value.

### HasProhibitWifiShare

`func (o *CreateSsidOpenApiVO) HasProhibitWifiShare() bool`

HasProhibitWifiShare returns a boolean if a field has been set.

### GetPskSetting

`func (o *CreateSsidOpenApiVO) GetPskSetting() SsidPskSettingOpenApiVO`

GetPskSetting returns the PskSetting field if non-nil, zero value otherwise.

### GetPskSettingOk

`func (o *CreateSsidOpenApiVO) GetPskSettingOk() (*SsidPskSettingOpenApiVO, bool)`

GetPskSettingOk returns a tuple with the PskSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPskSetting

`func (o *CreateSsidOpenApiVO) SetPskSetting(v SsidPskSettingOpenApiVO)`

SetPskSetting sets PskSetting field to given value.

### HasPskSetting

`func (o *CreateSsidOpenApiVO) HasPskSetting() bool`

HasPskSetting returns a boolean if a field has been set.

### GetSecurity

`func (o *CreateSsidOpenApiVO) GetSecurity() int32`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *CreateSsidOpenApiVO) GetSecurityOk() (*int32, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *CreateSsidOpenApiVO) SetSecurity(v int32)`

SetSecurity sets Security field to given value.


### GetVlanEnable

`func (o *CreateSsidOpenApiVO) GetVlanEnable() bool`

GetVlanEnable returns the VlanEnable field if non-nil, zero value otherwise.

### GetVlanEnableOk

`func (o *CreateSsidOpenApiVO) GetVlanEnableOk() (*bool, bool)`

GetVlanEnableOk returns a tuple with the VlanEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanEnable

`func (o *CreateSsidOpenApiVO) SetVlanEnable(v bool)`

SetVlanEnable sets VlanEnable field to given value.


### GetVlanId

`func (o *CreateSsidOpenApiVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *CreateSsidOpenApiVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *CreateSsidOpenApiVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *CreateSsidOpenApiVO) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlanSetting

`func (o *CreateSsidOpenApiVO) GetVlanSetting() SsidVlanSettingOpenApiVO`

GetVlanSetting returns the VlanSetting field if non-nil, zero value otherwise.

### GetVlanSettingOk

`func (o *CreateSsidOpenApiVO) GetVlanSettingOk() (*SsidVlanSettingOpenApiVO, bool)`

GetVlanSettingOk returns a tuple with the VlanSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanSetting

`func (o *CreateSsidOpenApiVO) SetVlanSetting(v SsidVlanSettingOpenApiVO)`

SetVlanSetting sets VlanSetting field to given value.

### HasVlanSetting

`func (o *CreateSsidOpenApiVO) HasVlanSetting() bool`

HasVlanSetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


