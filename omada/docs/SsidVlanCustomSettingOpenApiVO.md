# SsidVlanCustomSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BridgeVlan** | Pointer to **int32** | bridgeVlan. If support vlan pool, use lanNetworkVlanIds instead.  If customMode&#x3D;1, this filed must be null. | [optional] 
**CustomMode** | **int32** | If mode&#x3D;1, this field must be entered.If a device does not support multiple VLANs, the smallest VLAN you configured will be applied to the SSID. customMode should be a value as follows: 0:by Network; 1:by Vlan. | 
**LanNetworkId** | Pointer to **string** | lanNetworkId. If support vlan pool, use lanNetworkVlanIds instead. If customMode&#x3D;1, this filed must be null. If both lanNetwork and lanNetworkVlanIds parameters exist, the vlanId will actually take effect with the lanNetwork with the smallest vlanId in the lanNetworkVlanIds parameter. | [optional] 
**LanNetworkVlanIds** | Pointer to **map[string][]int32** | Indicates the mapping between the lanNetworkId and the vlanId, and if the lanNetwork corresponds to a bridgeVlan, multiple vlanIds may correspond.  If customMode&#x3D;1, this filed must be null. Cbc Pro does not support this filed. | [optional] 
**VlanId** | Pointer to **int32** | vlanId. If support vlanPool, use vlanPoolIds instead. If customMode&#x3D;0, this filed must be null. If both the vlanId and vlanPoolIds parameters exist, the vlanId will actually take effect at the minimum value in the vlanPoolIds parameter. | [optional] 
**VlanPoolIds** | Pointer to **string** | When customMode&#x3D;1 needs to have a value.  If customMode&#x3D;0, this filed must be null. Cbc Pro does not support this filed. | [optional] 

## Methods

### NewSsidVlanCustomSettingOpenApiVO

`func NewSsidVlanCustomSettingOpenApiVO(customMode int32, ) *SsidVlanCustomSettingOpenApiVO`

NewSsidVlanCustomSettingOpenApiVO instantiates a new SsidVlanCustomSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsidVlanCustomSettingOpenApiVOWithDefaults

`func NewSsidVlanCustomSettingOpenApiVOWithDefaults() *SsidVlanCustomSettingOpenApiVO`

NewSsidVlanCustomSettingOpenApiVOWithDefaults instantiates a new SsidVlanCustomSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBridgeVlan

`func (o *SsidVlanCustomSettingOpenApiVO) GetBridgeVlan() int32`

GetBridgeVlan returns the BridgeVlan field if non-nil, zero value otherwise.

### GetBridgeVlanOk

`func (o *SsidVlanCustomSettingOpenApiVO) GetBridgeVlanOk() (*int32, bool)`

GetBridgeVlanOk returns a tuple with the BridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeVlan

`func (o *SsidVlanCustomSettingOpenApiVO) SetBridgeVlan(v int32)`

SetBridgeVlan sets BridgeVlan field to given value.

### HasBridgeVlan

`func (o *SsidVlanCustomSettingOpenApiVO) HasBridgeVlan() bool`

HasBridgeVlan returns a boolean if a field has been set.

### GetCustomMode

`func (o *SsidVlanCustomSettingOpenApiVO) GetCustomMode() int32`

GetCustomMode returns the CustomMode field if non-nil, zero value otherwise.

### GetCustomModeOk

`func (o *SsidVlanCustomSettingOpenApiVO) GetCustomModeOk() (*int32, bool)`

GetCustomModeOk returns a tuple with the CustomMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMode

`func (o *SsidVlanCustomSettingOpenApiVO) SetCustomMode(v int32)`

SetCustomMode sets CustomMode field to given value.


### GetLanNetworkId

`func (o *SsidVlanCustomSettingOpenApiVO) GetLanNetworkId() string`

GetLanNetworkId returns the LanNetworkId field if non-nil, zero value otherwise.

### GetLanNetworkIdOk

`func (o *SsidVlanCustomSettingOpenApiVO) GetLanNetworkIdOk() (*string, bool)`

GetLanNetworkIdOk returns a tuple with the LanNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetworkId

`func (o *SsidVlanCustomSettingOpenApiVO) SetLanNetworkId(v string)`

SetLanNetworkId sets LanNetworkId field to given value.

### HasLanNetworkId

`func (o *SsidVlanCustomSettingOpenApiVO) HasLanNetworkId() bool`

HasLanNetworkId returns a boolean if a field has been set.

### GetLanNetworkVlanIds

`func (o *SsidVlanCustomSettingOpenApiVO) GetLanNetworkVlanIds() map[string][]int32`

GetLanNetworkVlanIds returns the LanNetworkVlanIds field if non-nil, zero value otherwise.

### GetLanNetworkVlanIdsOk

`func (o *SsidVlanCustomSettingOpenApiVO) GetLanNetworkVlanIdsOk() (*map[string][]int32, bool)`

GetLanNetworkVlanIdsOk returns a tuple with the LanNetworkVlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetworkVlanIds

`func (o *SsidVlanCustomSettingOpenApiVO) SetLanNetworkVlanIds(v map[string][]int32)`

SetLanNetworkVlanIds sets LanNetworkVlanIds field to given value.

### HasLanNetworkVlanIds

`func (o *SsidVlanCustomSettingOpenApiVO) HasLanNetworkVlanIds() bool`

HasLanNetworkVlanIds returns a boolean if a field has been set.

### GetVlanId

`func (o *SsidVlanCustomSettingOpenApiVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *SsidVlanCustomSettingOpenApiVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *SsidVlanCustomSettingOpenApiVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *SsidVlanCustomSettingOpenApiVO) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlanPoolIds

`func (o *SsidVlanCustomSettingOpenApiVO) GetVlanPoolIds() string`

GetVlanPoolIds returns the VlanPoolIds field if non-nil, zero value otherwise.

### GetVlanPoolIdsOk

`func (o *SsidVlanCustomSettingOpenApiVO) GetVlanPoolIdsOk() (*string, bool)`

GetVlanPoolIdsOk returns a tuple with the VlanPoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPoolIds

`func (o *SsidVlanCustomSettingOpenApiVO) SetVlanPoolIds(v string)`

SetVlanPoolIds sets VlanPoolIds field to given value.

### HasVlanPoolIds

`func (o *SsidVlanCustomSettingOpenApiVO) HasVlanPoolIds() bool`

HasVlanPoolIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


