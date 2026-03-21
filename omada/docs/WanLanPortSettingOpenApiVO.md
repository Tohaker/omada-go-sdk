# WanLanPortSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Closable** | Pointer to **bool** | can this port be closed or not, true: support | [optional] 
**LanNetworkNames** | Pointer to **[]string** | names of lan networks | [optional] 
**Mode** | Pointer to **int32** | 0: WAN, 1: LAN | [optional] 
**PortName** | Pointer to **string** | port name | [optional] 
**PortUuid** | Pointer to **string** | port uuid | [optional] 
**Status** | Pointer to **int32** | whether to open this port | [optional] 
**SupportInternetVlan** | Pointer to **bool** | does this port support vlan or not, true: support | [optional] 
**SupportIptv** | Pointer to **bool** | support iptv or not, true: support | [optional] 
**SupportQosTagEnable** | Pointer to **bool** | does this port support qos tage enable | [optional] 
**SupportVpn** | Pointer to **bool** | support vpn or not, true: support | [optional] 
**Type** | Pointer to **int32** | 0: WAN, 1: WAN/LAN, 2: LAN | [optional] 
**WanIpv6Component** | Pointer to **int32** | component version of gateway wan ipv6 | [optional] 

## Methods

### NewWanLanPortSettingOpenApiVO

`func NewWanLanPortSettingOpenApiVO() *WanLanPortSettingOpenApiVO`

NewWanLanPortSettingOpenApiVO instantiates a new WanLanPortSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWanLanPortSettingOpenApiVOWithDefaults

`func NewWanLanPortSettingOpenApiVOWithDefaults() *WanLanPortSettingOpenApiVO`

NewWanLanPortSettingOpenApiVOWithDefaults instantiates a new WanLanPortSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClosable

`func (o *WanLanPortSettingOpenApiVO) GetClosable() bool`

GetClosable returns the Closable field if non-nil, zero value otherwise.

### GetClosableOk

`func (o *WanLanPortSettingOpenApiVO) GetClosableOk() (*bool, bool)`

GetClosableOk returns a tuple with the Closable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosable

`func (o *WanLanPortSettingOpenApiVO) SetClosable(v bool)`

SetClosable sets Closable field to given value.

### HasClosable

`func (o *WanLanPortSettingOpenApiVO) HasClosable() bool`

HasClosable returns a boolean if a field has been set.

### GetLanNetworkNames

`func (o *WanLanPortSettingOpenApiVO) GetLanNetworkNames() []string`

GetLanNetworkNames returns the LanNetworkNames field if non-nil, zero value otherwise.

### GetLanNetworkNamesOk

`func (o *WanLanPortSettingOpenApiVO) GetLanNetworkNamesOk() (*[]string, bool)`

GetLanNetworkNamesOk returns a tuple with the LanNetworkNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetworkNames

`func (o *WanLanPortSettingOpenApiVO) SetLanNetworkNames(v []string)`

SetLanNetworkNames sets LanNetworkNames field to given value.

### HasLanNetworkNames

`func (o *WanLanPortSettingOpenApiVO) HasLanNetworkNames() bool`

HasLanNetworkNames returns a boolean if a field has been set.

### GetMode

`func (o *WanLanPortSettingOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *WanLanPortSettingOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *WanLanPortSettingOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *WanLanPortSettingOpenApiVO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPortName

`func (o *WanLanPortSettingOpenApiVO) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *WanLanPortSettingOpenApiVO) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *WanLanPortSettingOpenApiVO) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *WanLanPortSettingOpenApiVO) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetPortUuid

`func (o *WanLanPortSettingOpenApiVO) GetPortUuid() string`

GetPortUuid returns the PortUuid field if non-nil, zero value otherwise.

### GetPortUuidOk

`func (o *WanLanPortSettingOpenApiVO) GetPortUuidOk() (*string, bool)`

GetPortUuidOk returns a tuple with the PortUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUuid

`func (o *WanLanPortSettingOpenApiVO) SetPortUuid(v string)`

SetPortUuid sets PortUuid field to given value.

### HasPortUuid

`func (o *WanLanPortSettingOpenApiVO) HasPortUuid() bool`

HasPortUuid returns a boolean if a field has been set.

### GetStatus

`func (o *WanLanPortSettingOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WanLanPortSettingOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WanLanPortSettingOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WanLanPortSettingOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportInternetVlan

`func (o *WanLanPortSettingOpenApiVO) GetSupportInternetVlan() bool`

GetSupportInternetVlan returns the SupportInternetVlan field if non-nil, zero value otherwise.

### GetSupportInternetVlanOk

`func (o *WanLanPortSettingOpenApiVO) GetSupportInternetVlanOk() (*bool, bool)`

GetSupportInternetVlanOk returns a tuple with the SupportInternetVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportInternetVlan

`func (o *WanLanPortSettingOpenApiVO) SetSupportInternetVlan(v bool)`

SetSupportInternetVlan sets SupportInternetVlan field to given value.

### HasSupportInternetVlan

`func (o *WanLanPortSettingOpenApiVO) HasSupportInternetVlan() bool`

HasSupportInternetVlan returns a boolean if a field has been set.

### GetSupportIptv

`func (o *WanLanPortSettingOpenApiVO) GetSupportIptv() bool`

GetSupportIptv returns the SupportIptv field if non-nil, zero value otherwise.

### GetSupportIptvOk

`func (o *WanLanPortSettingOpenApiVO) GetSupportIptvOk() (*bool, bool)`

GetSupportIptvOk returns a tuple with the SupportIptv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIptv

`func (o *WanLanPortSettingOpenApiVO) SetSupportIptv(v bool)`

SetSupportIptv sets SupportIptv field to given value.

### HasSupportIptv

`func (o *WanLanPortSettingOpenApiVO) HasSupportIptv() bool`

HasSupportIptv returns a boolean if a field has been set.

### GetSupportQosTagEnable

`func (o *WanLanPortSettingOpenApiVO) GetSupportQosTagEnable() bool`

GetSupportQosTagEnable returns the SupportQosTagEnable field if non-nil, zero value otherwise.

### GetSupportQosTagEnableOk

`func (o *WanLanPortSettingOpenApiVO) GetSupportQosTagEnableOk() (*bool, bool)`

GetSupportQosTagEnableOk returns a tuple with the SupportQosTagEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportQosTagEnable

`func (o *WanLanPortSettingOpenApiVO) SetSupportQosTagEnable(v bool)`

SetSupportQosTagEnable sets SupportQosTagEnable field to given value.

### HasSupportQosTagEnable

`func (o *WanLanPortSettingOpenApiVO) HasSupportQosTagEnable() bool`

HasSupportQosTagEnable returns a boolean if a field has been set.

### GetSupportVpn

`func (o *WanLanPortSettingOpenApiVO) GetSupportVpn() bool`

GetSupportVpn returns the SupportVpn field if non-nil, zero value otherwise.

### GetSupportVpnOk

`func (o *WanLanPortSettingOpenApiVO) GetSupportVpnOk() (*bool, bool)`

GetSupportVpnOk returns a tuple with the SupportVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVpn

`func (o *WanLanPortSettingOpenApiVO) SetSupportVpn(v bool)`

SetSupportVpn sets SupportVpn field to given value.

### HasSupportVpn

`func (o *WanLanPortSettingOpenApiVO) HasSupportVpn() bool`

HasSupportVpn returns a boolean if a field has been set.

### GetType

`func (o *WanLanPortSettingOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WanLanPortSettingOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WanLanPortSettingOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *WanLanPortSettingOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWanIpv6Component

`func (o *WanLanPortSettingOpenApiVO) GetWanIpv6Component() int32`

GetWanIpv6Component returns the WanIpv6Component field if non-nil, zero value otherwise.

### GetWanIpv6ComponentOk

`func (o *WanLanPortSettingOpenApiVO) GetWanIpv6ComponentOk() (*int32, bool)`

GetWanIpv6ComponentOk returns a tuple with the WanIpv6Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanIpv6Component

`func (o *WanLanPortSettingOpenApiVO) SetWanIpv6Component(v int32)`

SetWanIpv6Component sets WanIpv6Component field to given value.

### HasWanIpv6Component

`func (o *WanLanPortSettingOpenApiVO) HasWanIpv6Component() bool`

HasWanIpv6Component returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


