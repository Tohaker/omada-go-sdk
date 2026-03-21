# VlanPreConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedDeviceList** | Pointer to [**[]SelectDeviceForVlanVO**](SelectDeviceForVlanVO.md) | Affected device list | [optional] 
**AffectedStackList** | Pointer to [**[]SelectStackForVlanVO**](SelectStackForVlanVO.md) | Affected stack list | [optional] 
**DeviceConfig** | Pointer to [**SelectPortBindingBriefVO**](SelectPortBindingBriefVO.md) |  | [optional] 
**LanNetwork** | Pointer to [**LanNetworkQueryOpenApiV3VO**](LanNetworkQueryOpenApiV3VO.md) |  | [optional] 
**NetworkSegmentChanged** | Pointer to **bool** | Indicate whether the network segment change when modifying network | [optional] 
**TotalAffectedDeviceNum** | Pointer to **int32** | Indicate total affected devices num, including devices with ports that have Network Tags Setting set to Allow all | [optional] 

## Methods

### NewVlanPreConfigOpenApiVO

`func NewVlanPreConfigOpenApiVO() *VlanPreConfigOpenApiVO`

NewVlanPreConfigOpenApiVO instantiates a new VlanPreConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanPreConfigOpenApiVOWithDefaults

`func NewVlanPreConfigOpenApiVOWithDefaults() *VlanPreConfigOpenApiVO`

NewVlanPreConfigOpenApiVOWithDefaults instantiates a new VlanPreConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedDeviceList

`func (o *VlanPreConfigOpenApiVO) GetAffectedDeviceList() []SelectDeviceForVlanVO`

GetAffectedDeviceList returns the AffectedDeviceList field if non-nil, zero value otherwise.

### GetAffectedDeviceListOk

`func (o *VlanPreConfigOpenApiVO) GetAffectedDeviceListOk() (*[]SelectDeviceForVlanVO, bool)`

GetAffectedDeviceListOk returns a tuple with the AffectedDeviceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedDeviceList

`func (o *VlanPreConfigOpenApiVO) SetAffectedDeviceList(v []SelectDeviceForVlanVO)`

SetAffectedDeviceList sets AffectedDeviceList field to given value.

### HasAffectedDeviceList

`func (o *VlanPreConfigOpenApiVO) HasAffectedDeviceList() bool`

HasAffectedDeviceList returns a boolean if a field has been set.

### GetAffectedStackList

`func (o *VlanPreConfigOpenApiVO) GetAffectedStackList() []SelectStackForVlanVO`

GetAffectedStackList returns the AffectedStackList field if non-nil, zero value otherwise.

### GetAffectedStackListOk

`func (o *VlanPreConfigOpenApiVO) GetAffectedStackListOk() (*[]SelectStackForVlanVO, bool)`

GetAffectedStackListOk returns a tuple with the AffectedStackList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedStackList

`func (o *VlanPreConfigOpenApiVO) SetAffectedStackList(v []SelectStackForVlanVO)`

SetAffectedStackList sets AffectedStackList field to given value.

### HasAffectedStackList

`func (o *VlanPreConfigOpenApiVO) HasAffectedStackList() bool`

HasAffectedStackList returns a boolean if a field has been set.

### GetDeviceConfig

`func (o *VlanPreConfigOpenApiVO) GetDeviceConfig() SelectPortBindingBriefVO`

GetDeviceConfig returns the DeviceConfig field if non-nil, zero value otherwise.

### GetDeviceConfigOk

`func (o *VlanPreConfigOpenApiVO) GetDeviceConfigOk() (*SelectPortBindingBriefVO, bool)`

GetDeviceConfigOk returns a tuple with the DeviceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceConfig

`func (o *VlanPreConfigOpenApiVO) SetDeviceConfig(v SelectPortBindingBriefVO)`

SetDeviceConfig sets DeviceConfig field to given value.

### HasDeviceConfig

`func (o *VlanPreConfigOpenApiVO) HasDeviceConfig() bool`

HasDeviceConfig returns a boolean if a field has been set.

### GetLanNetwork

`func (o *VlanPreConfigOpenApiVO) GetLanNetwork() LanNetworkQueryOpenApiV3VO`

GetLanNetwork returns the LanNetwork field if non-nil, zero value otherwise.

### GetLanNetworkOk

`func (o *VlanPreConfigOpenApiVO) GetLanNetworkOk() (*LanNetworkQueryOpenApiV3VO, bool)`

GetLanNetworkOk returns a tuple with the LanNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetwork

`func (o *VlanPreConfigOpenApiVO) SetLanNetwork(v LanNetworkQueryOpenApiV3VO)`

SetLanNetwork sets LanNetwork field to given value.

### HasLanNetwork

`func (o *VlanPreConfigOpenApiVO) HasLanNetwork() bool`

HasLanNetwork returns a boolean if a field has been set.

### GetNetworkSegmentChanged

`func (o *VlanPreConfigOpenApiVO) GetNetworkSegmentChanged() bool`

GetNetworkSegmentChanged returns the NetworkSegmentChanged field if non-nil, zero value otherwise.

### GetNetworkSegmentChangedOk

`func (o *VlanPreConfigOpenApiVO) GetNetworkSegmentChangedOk() (*bool, bool)`

GetNetworkSegmentChangedOk returns a tuple with the NetworkSegmentChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSegmentChanged

`func (o *VlanPreConfigOpenApiVO) SetNetworkSegmentChanged(v bool)`

SetNetworkSegmentChanged sets NetworkSegmentChanged field to given value.

### HasNetworkSegmentChanged

`func (o *VlanPreConfigOpenApiVO) HasNetworkSegmentChanged() bool`

HasNetworkSegmentChanged returns a boolean if a field has been set.

### GetTotalAffectedDeviceNum

`func (o *VlanPreConfigOpenApiVO) GetTotalAffectedDeviceNum() int32`

GetTotalAffectedDeviceNum returns the TotalAffectedDeviceNum field if non-nil, zero value otherwise.

### GetTotalAffectedDeviceNumOk

`func (o *VlanPreConfigOpenApiVO) GetTotalAffectedDeviceNumOk() (*int32, bool)`

GetTotalAffectedDeviceNumOk returns a tuple with the TotalAffectedDeviceNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAffectedDeviceNum

`func (o *VlanPreConfigOpenApiVO) SetTotalAffectedDeviceNum(v int32)`

SetTotalAffectedDeviceNum sets TotalAffectedDeviceNum field to given value.

### HasTotalAffectedDeviceNum

`func (o *VlanPreConfigOpenApiVO) HasTotalAffectedDeviceNum() bool`

HasTotalAffectedDeviceNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


