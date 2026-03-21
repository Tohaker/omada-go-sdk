# ApMvlanSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BridgeVlan** | Pointer to **int32** | When mvlanNetworkId is bridge vlan, mvlanBridgeVlan has a value. | [optional] 
**LanNetworkId** | Pointer to **string** | This field indicates the currently effective lanNetworkId. | [optional] 
**Mode** | **int32** | Mode for the AP management VLAN configuration to take effect.[0:Default,1:Custom] | 

## Methods

### NewApMvlanSettingOpenApiVO

`func NewApMvlanSettingOpenApiVO(mode int32, ) *ApMvlanSettingOpenApiVO`

NewApMvlanSettingOpenApiVO instantiates a new ApMvlanSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApMvlanSettingOpenApiVOWithDefaults

`func NewApMvlanSettingOpenApiVOWithDefaults() *ApMvlanSettingOpenApiVO`

NewApMvlanSettingOpenApiVOWithDefaults instantiates a new ApMvlanSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBridgeVlan

`func (o *ApMvlanSettingOpenApiVO) GetBridgeVlan() int32`

GetBridgeVlan returns the BridgeVlan field if non-nil, zero value otherwise.

### GetBridgeVlanOk

`func (o *ApMvlanSettingOpenApiVO) GetBridgeVlanOk() (*int32, bool)`

GetBridgeVlanOk returns a tuple with the BridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeVlan

`func (o *ApMvlanSettingOpenApiVO) SetBridgeVlan(v int32)`

SetBridgeVlan sets BridgeVlan field to given value.

### HasBridgeVlan

`func (o *ApMvlanSettingOpenApiVO) HasBridgeVlan() bool`

HasBridgeVlan returns a boolean if a field has been set.

### GetLanNetworkId

`func (o *ApMvlanSettingOpenApiVO) GetLanNetworkId() string`

GetLanNetworkId returns the LanNetworkId field if non-nil, zero value otherwise.

### GetLanNetworkIdOk

`func (o *ApMvlanSettingOpenApiVO) GetLanNetworkIdOk() (*string, bool)`

GetLanNetworkIdOk returns a tuple with the LanNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetworkId

`func (o *ApMvlanSettingOpenApiVO) SetLanNetworkId(v string)`

SetLanNetworkId sets LanNetworkId field to given value.

### HasLanNetworkId

`func (o *ApMvlanSettingOpenApiVO) HasLanNetworkId() bool`

HasLanNetworkId returns a boolean if a field has been set.

### GetMode

`func (o *ApMvlanSettingOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ApMvlanSettingOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ApMvlanSettingOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


