# Dot1xGuestVlanSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BridgeVlan** | Pointer to **int32** | Actual VLAN ID to apply when the selected LAN network is a Multi-VLAN network. Valid range: 1-4094. | [optional] 
**LanNetworkId** | Pointer to **string** | LAN network ID used when [mode] &#x3D; 1 (By Network). Specifies the target LAN network to which the Guest VLAN is mapped. The LAN network ID can be obtained from the &#39;Get LAN network list V3&#39; interface. | [optional] 
**Mode** | **int32** | Guest VLAN application mode. Valid values: 0 &#x3D; None (Disable Guest Vlan), 1 &#x3D; By Network (select a LAN network), 2 &#x3D; By VLAN ID (specify a VLAN ID directly). | 
**VlanId** | Pointer to **int32** | Guest VLAN ID used when [mode] &#x3D; 2 (By VLAN ID). Valid range: 1-4094. | [optional] 

## Methods

### NewDot1xGuestVlanSettingOpenApiVO

`func NewDot1xGuestVlanSettingOpenApiVO(mode int32, ) *Dot1xGuestVlanSettingOpenApiVO`

NewDot1xGuestVlanSettingOpenApiVO instantiates a new Dot1xGuestVlanSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDot1xGuestVlanSettingOpenApiVOWithDefaults

`func NewDot1xGuestVlanSettingOpenApiVOWithDefaults() *Dot1xGuestVlanSettingOpenApiVO`

NewDot1xGuestVlanSettingOpenApiVOWithDefaults instantiates a new Dot1xGuestVlanSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBridgeVlan

`func (o *Dot1xGuestVlanSettingOpenApiVO) GetBridgeVlan() int32`

GetBridgeVlan returns the BridgeVlan field if non-nil, zero value otherwise.

### GetBridgeVlanOk

`func (o *Dot1xGuestVlanSettingOpenApiVO) GetBridgeVlanOk() (*int32, bool)`

GetBridgeVlanOk returns a tuple with the BridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeVlan

`func (o *Dot1xGuestVlanSettingOpenApiVO) SetBridgeVlan(v int32)`

SetBridgeVlan sets BridgeVlan field to given value.

### HasBridgeVlan

`func (o *Dot1xGuestVlanSettingOpenApiVO) HasBridgeVlan() bool`

HasBridgeVlan returns a boolean if a field has been set.

### GetLanNetworkId

`func (o *Dot1xGuestVlanSettingOpenApiVO) GetLanNetworkId() string`

GetLanNetworkId returns the LanNetworkId field if non-nil, zero value otherwise.

### GetLanNetworkIdOk

`func (o *Dot1xGuestVlanSettingOpenApiVO) GetLanNetworkIdOk() (*string, bool)`

GetLanNetworkIdOk returns a tuple with the LanNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetworkId

`func (o *Dot1xGuestVlanSettingOpenApiVO) SetLanNetworkId(v string)`

SetLanNetworkId sets LanNetworkId field to given value.

### HasLanNetworkId

`func (o *Dot1xGuestVlanSettingOpenApiVO) HasLanNetworkId() bool`

HasLanNetworkId returns a boolean if a field has been set.

### GetMode

`func (o *Dot1xGuestVlanSettingOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Dot1xGuestVlanSettingOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Dot1xGuestVlanSettingOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetVlanId

`func (o *Dot1xGuestVlanSettingOpenApiVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *Dot1xGuestVlanSettingOpenApiVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *Dot1xGuestVlanSettingOpenApiVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *Dot1xGuestVlanSettingOpenApiVO) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


