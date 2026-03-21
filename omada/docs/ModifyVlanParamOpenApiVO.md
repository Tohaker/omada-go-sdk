# ModifyVlanParamOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceConfig** | [**SelectPortBindingBriefVO**](SelectPortBindingBriefVO.md) |  | 
**LanNetwork** | [**LanNetworkOpenApiV3VO**](LanNetworkOpenApiV3VO.md) |  | 

## Methods

### NewModifyVlanParamOpenApiVO

`func NewModifyVlanParamOpenApiVO(deviceConfig SelectPortBindingBriefVO, lanNetwork LanNetworkOpenApiV3VO, ) *ModifyVlanParamOpenApiVO`

NewModifyVlanParamOpenApiVO instantiates a new ModifyVlanParamOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyVlanParamOpenApiVOWithDefaults

`func NewModifyVlanParamOpenApiVOWithDefaults() *ModifyVlanParamOpenApiVO`

NewModifyVlanParamOpenApiVOWithDefaults instantiates a new ModifyVlanParamOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceConfig

`func (o *ModifyVlanParamOpenApiVO) GetDeviceConfig() SelectPortBindingBriefVO`

GetDeviceConfig returns the DeviceConfig field if non-nil, zero value otherwise.

### GetDeviceConfigOk

`func (o *ModifyVlanParamOpenApiVO) GetDeviceConfigOk() (*SelectPortBindingBriefVO, bool)`

GetDeviceConfigOk returns a tuple with the DeviceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceConfig

`func (o *ModifyVlanParamOpenApiVO) SetDeviceConfig(v SelectPortBindingBriefVO)`

SetDeviceConfig sets DeviceConfig field to given value.


### GetLanNetwork

`func (o *ModifyVlanParamOpenApiVO) GetLanNetwork() LanNetworkOpenApiV3VO`

GetLanNetwork returns the LanNetwork field if non-nil, zero value otherwise.

### GetLanNetworkOk

`func (o *ModifyVlanParamOpenApiVO) GetLanNetworkOk() (*LanNetworkOpenApiV3VO, bool)`

GetLanNetworkOk returns a tuple with the LanNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetwork

`func (o *ModifyVlanParamOpenApiVO) SetLanNetwork(v LanNetworkOpenApiV3VO)`

SetLanNetwork sets LanNetwork field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


