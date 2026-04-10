# ApMgtSsidVlanCustomSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BridgeVlan** | Pointer to **int32** | bridgeVlan. If support vlan pool, use lanNetworkVlanIds instead.  If customMode&#x3D;1, this filed must be null. | [optional] 
**CustomMode** | **int32** | If mode&#x3D;1, this field must be entered.If a device does not support multiple VLANs, the smallest VLAN you configured will be applied to the SSID. customMode should be a value as follows: 0:by Network; 1:by Vlan. | 
**LanNetworkId** | Pointer to **string** | lanNetworkId. If support vlan pool, use lanNetworkVlanIds instead. If customMode&#x3D;1, this filed must be null. If both lanNetwork and lanNetworkVlanIds parameters exist, the vlanId will actually take effect with the lanNetwork with the smallest vlanId in the lanNetworkVlanIds parameter. | [optional] 
**VlanId** | Pointer to **int32** | vlanId. If support vlanPool, use vlanPoolIds instead. If customMode&#x3D;0, this filed must be null. If both the vlanId and vlanPoolIds parameters exist, the vlanId will actually take effect at the minimum value in the vlanPoolIds parameter. | [optional] 

## Methods

### NewApMgtSsidVlanCustomSettingOpenApiVO

`func NewApMgtSsidVlanCustomSettingOpenApiVO(customMode int32, ) *ApMgtSsidVlanCustomSettingOpenApiVO`

NewApMgtSsidVlanCustomSettingOpenApiVO instantiates a new ApMgtSsidVlanCustomSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApMgtSsidVlanCustomSettingOpenApiVOWithDefaults

`func NewApMgtSsidVlanCustomSettingOpenApiVOWithDefaults() *ApMgtSsidVlanCustomSettingOpenApiVO`

NewApMgtSsidVlanCustomSettingOpenApiVOWithDefaults instantiates a new ApMgtSsidVlanCustomSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBridgeVlan

`func (o *ApMgtSsidVlanCustomSettingOpenApiVO) GetBridgeVlan() int32`

GetBridgeVlan returns the BridgeVlan field if non-nil, zero value otherwise.

### GetBridgeVlanOk

`func (o *ApMgtSsidVlanCustomSettingOpenApiVO) GetBridgeVlanOk() (*int32, bool)`

GetBridgeVlanOk returns a tuple with the BridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeVlan

`func (o *ApMgtSsidVlanCustomSettingOpenApiVO) SetBridgeVlan(v int32)`

SetBridgeVlan sets BridgeVlan field to given value.

### HasBridgeVlan

`func (o *ApMgtSsidVlanCustomSettingOpenApiVO) HasBridgeVlan() bool`

HasBridgeVlan returns a boolean if a field has been set.

### GetCustomMode

`func (o *ApMgtSsidVlanCustomSettingOpenApiVO) GetCustomMode() int32`

GetCustomMode returns the CustomMode field if non-nil, zero value otherwise.

### GetCustomModeOk

`func (o *ApMgtSsidVlanCustomSettingOpenApiVO) GetCustomModeOk() (*int32, bool)`

GetCustomModeOk returns a tuple with the CustomMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMode

`func (o *ApMgtSsidVlanCustomSettingOpenApiVO) SetCustomMode(v int32)`

SetCustomMode sets CustomMode field to given value.


### GetLanNetworkId

`func (o *ApMgtSsidVlanCustomSettingOpenApiVO) GetLanNetworkId() string`

GetLanNetworkId returns the LanNetworkId field if non-nil, zero value otherwise.

### GetLanNetworkIdOk

`func (o *ApMgtSsidVlanCustomSettingOpenApiVO) GetLanNetworkIdOk() (*string, bool)`

GetLanNetworkIdOk returns a tuple with the LanNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetworkId

`func (o *ApMgtSsidVlanCustomSettingOpenApiVO) SetLanNetworkId(v string)`

SetLanNetworkId sets LanNetworkId field to given value.

### HasLanNetworkId

`func (o *ApMgtSsidVlanCustomSettingOpenApiVO) HasLanNetworkId() bool`

HasLanNetworkId returns a boolean if a field has been set.

### GetVlanId

`func (o *ApMgtSsidVlanCustomSettingOpenApiVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *ApMgtSsidVlanCustomSettingOpenApiVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *ApMgtSsidVlanCustomSettingOpenApiVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *ApMgtSsidVlanCustomSettingOpenApiVO) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


