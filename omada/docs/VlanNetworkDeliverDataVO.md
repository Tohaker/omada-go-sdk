# VlanNetworkDeliverDataVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedDeviceList** | Pointer to [**[]SelectDeviceForVlanVO**](SelectDeviceForVlanVO.md) | Affected device list | [optional] 
**AffectedStackList** | Pointer to [**[]SelectStackForVlanVO**](SelectStackForVlanVO.md) | Affected stack list | [optional] 
**DeliveringDevices** | Pointer to **[]string** | Delivering devices | [optional] 
**DeliveringStacks** | Pointer to **[]string** | Delivering stacks | [optional] 
**DeviceConfig** | Pointer to [**SelectPortBindingBriefVO**](SelectPortBindingBriefVO.md) |  | [optional] 
**DoneDevicesNum** | Pointer to **int32** | Delivered device num | [optional] 
**FailDevices** | Pointer to **[]string** | Failed devices | [optional] 
**FailStacks** | Pointer to **[]string** | Failed stacks | [optional] 
**LanNetwork** | Pointer to [**LanNetworkVO**](LanNetworkVO.md) |  | [optional] 
**State** | Pointer to **int32** | Delivering state. 0: free ( it means the network has been already delivered)  1: delivering 2. deliver done. | [optional] 
**SuccessDevices** | Pointer to **[]string** | Successfully delivered devices | [optional] 
**SuccessStacks** | Pointer to **[]string** | Successfully delivered stacks | [optional] 
**TotalDevicesNum** | Pointer to **int32** | Total device num | [optional] 

## Methods

### NewVlanNetworkDeliverDataVO

`func NewVlanNetworkDeliverDataVO() *VlanNetworkDeliverDataVO`

NewVlanNetworkDeliverDataVO instantiates a new VlanNetworkDeliverDataVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanNetworkDeliverDataVOWithDefaults

`func NewVlanNetworkDeliverDataVOWithDefaults() *VlanNetworkDeliverDataVO`

NewVlanNetworkDeliverDataVOWithDefaults instantiates a new VlanNetworkDeliverDataVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedDeviceList

`func (o *VlanNetworkDeliverDataVO) GetAffectedDeviceList() []SelectDeviceForVlanVO`

GetAffectedDeviceList returns the AffectedDeviceList field if non-nil, zero value otherwise.

### GetAffectedDeviceListOk

`func (o *VlanNetworkDeliverDataVO) GetAffectedDeviceListOk() (*[]SelectDeviceForVlanVO, bool)`

GetAffectedDeviceListOk returns a tuple with the AffectedDeviceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedDeviceList

`func (o *VlanNetworkDeliverDataVO) SetAffectedDeviceList(v []SelectDeviceForVlanVO)`

SetAffectedDeviceList sets AffectedDeviceList field to given value.

### HasAffectedDeviceList

`func (o *VlanNetworkDeliverDataVO) HasAffectedDeviceList() bool`

HasAffectedDeviceList returns a boolean if a field has been set.

### GetAffectedStackList

`func (o *VlanNetworkDeliverDataVO) GetAffectedStackList() []SelectStackForVlanVO`

GetAffectedStackList returns the AffectedStackList field if non-nil, zero value otherwise.

### GetAffectedStackListOk

`func (o *VlanNetworkDeliverDataVO) GetAffectedStackListOk() (*[]SelectStackForVlanVO, bool)`

GetAffectedStackListOk returns a tuple with the AffectedStackList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedStackList

`func (o *VlanNetworkDeliverDataVO) SetAffectedStackList(v []SelectStackForVlanVO)`

SetAffectedStackList sets AffectedStackList field to given value.

### HasAffectedStackList

`func (o *VlanNetworkDeliverDataVO) HasAffectedStackList() bool`

HasAffectedStackList returns a boolean if a field has been set.

### GetDeliveringDevices

`func (o *VlanNetworkDeliverDataVO) GetDeliveringDevices() []string`

GetDeliveringDevices returns the DeliveringDevices field if non-nil, zero value otherwise.

### GetDeliveringDevicesOk

`func (o *VlanNetworkDeliverDataVO) GetDeliveringDevicesOk() (*[]string, bool)`

GetDeliveringDevicesOk returns a tuple with the DeliveringDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveringDevices

`func (o *VlanNetworkDeliverDataVO) SetDeliveringDevices(v []string)`

SetDeliveringDevices sets DeliveringDevices field to given value.

### HasDeliveringDevices

`func (o *VlanNetworkDeliverDataVO) HasDeliveringDevices() bool`

HasDeliveringDevices returns a boolean if a field has been set.

### GetDeliveringStacks

`func (o *VlanNetworkDeliverDataVO) GetDeliveringStacks() []string`

GetDeliveringStacks returns the DeliveringStacks field if non-nil, zero value otherwise.

### GetDeliveringStacksOk

`func (o *VlanNetworkDeliverDataVO) GetDeliveringStacksOk() (*[]string, bool)`

GetDeliveringStacksOk returns a tuple with the DeliveringStacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveringStacks

`func (o *VlanNetworkDeliverDataVO) SetDeliveringStacks(v []string)`

SetDeliveringStacks sets DeliveringStacks field to given value.

### HasDeliveringStacks

`func (o *VlanNetworkDeliverDataVO) HasDeliveringStacks() bool`

HasDeliveringStacks returns a boolean if a field has been set.

### GetDeviceConfig

`func (o *VlanNetworkDeliverDataVO) GetDeviceConfig() SelectPortBindingBriefVO`

GetDeviceConfig returns the DeviceConfig field if non-nil, zero value otherwise.

### GetDeviceConfigOk

`func (o *VlanNetworkDeliverDataVO) GetDeviceConfigOk() (*SelectPortBindingBriefVO, bool)`

GetDeviceConfigOk returns a tuple with the DeviceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceConfig

`func (o *VlanNetworkDeliverDataVO) SetDeviceConfig(v SelectPortBindingBriefVO)`

SetDeviceConfig sets DeviceConfig field to given value.

### HasDeviceConfig

`func (o *VlanNetworkDeliverDataVO) HasDeviceConfig() bool`

HasDeviceConfig returns a boolean if a field has been set.

### GetDoneDevicesNum

`func (o *VlanNetworkDeliverDataVO) GetDoneDevicesNum() int32`

GetDoneDevicesNum returns the DoneDevicesNum field if non-nil, zero value otherwise.

### GetDoneDevicesNumOk

`func (o *VlanNetworkDeliverDataVO) GetDoneDevicesNumOk() (*int32, bool)`

GetDoneDevicesNumOk returns a tuple with the DoneDevicesNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoneDevicesNum

`func (o *VlanNetworkDeliverDataVO) SetDoneDevicesNum(v int32)`

SetDoneDevicesNum sets DoneDevicesNum field to given value.

### HasDoneDevicesNum

`func (o *VlanNetworkDeliverDataVO) HasDoneDevicesNum() bool`

HasDoneDevicesNum returns a boolean if a field has been set.

### GetFailDevices

`func (o *VlanNetworkDeliverDataVO) GetFailDevices() []string`

GetFailDevices returns the FailDevices field if non-nil, zero value otherwise.

### GetFailDevicesOk

`func (o *VlanNetworkDeliverDataVO) GetFailDevicesOk() (*[]string, bool)`

GetFailDevicesOk returns a tuple with the FailDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailDevices

`func (o *VlanNetworkDeliverDataVO) SetFailDevices(v []string)`

SetFailDevices sets FailDevices field to given value.

### HasFailDevices

`func (o *VlanNetworkDeliverDataVO) HasFailDevices() bool`

HasFailDevices returns a boolean if a field has been set.

### GetFailStacks

`func (o *VlanNetworkDeliverDataVO) GetFailStacks() []string`

GetFailStacks returns the FailStacks field if non-nil, zero value otherwise.

### GetFailStacksOk

`func (o *VlanNetworkDeliverDataVO) GetFailStacksOk() (*[]string, bool)`

GetFailStacksOk returns a tuple with the FailStacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailStacks

`func (o *VlanNetworkDeliverDataVO) SetFailStacks(v []string)`

SetFailStacks sets FailStacks field to given value.

### HasFailStacks

`func (o *VlanNetworkDeliverDataVO) HasFailStacks() bool`

HasFailStacks returns a boolean if a field has been set.

### GetLanNetwork

`func (o *VlanNetworkDeliverDataVO) GetLanNetwork() LanNetworkVO`

GetLanNetwork returns the LanNetwork field if non-nil, zero value otherwise.

### GetLanNetworkOk

`func (o *VlanNetworkDeliverDataVO) GetLanNetworkOk() (*LanNetworkVO, bool)`

GetLanNetworkOk returns a tuple with the LanNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetwork

`func (o *VlanNetworkDeliverDataVO) SetLanNetwork(v LanNetworkVO)`

SetLanNetwork sets LanNetwork field to given value.

### HasLanNetwork

`func (o *VlanNetworkDeliverDataVO) HasLanNetwork() bool`

HasLanNetwork returns a boolean if a field has been set.

### GetState

`func (o *VlanNetworkDeliverDataVO) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VlanNetworkDeliverDataVO) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VlanNetworkDeliverDataVO) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *VlanNetworkDeliverDataVO) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSuccessDevices

`func (o *VlanNetworkDeliverDataVO) GetSuccessDevices() []string`

GetSuccessDevices returns the SuccessDevices field if non-nil, zero value otherwise.

### GetSuccessDevicesOk

`func (o *VlanNetworkDeliverDataVO) GetSuccessDevicesOk() (*[]string, bool)`

GetSuccessDevicesOk returns a tuple with the SuccessDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessDevices

`func (o *VlanNetworkDeliverDataVO) SetSuccessDevices(v []string)`

SetSuccessDevices sets SuccessDevices field to given value.

### HasSuccessDevices

`func (o *VlanNetworkDeliverDataVO) HasSuccessDevices() bool`

HasSuccessDevices returns a boolean if a field has been set.

### GetSuccessStacks

`func (o *VlanNetworkDeliverDataVO) GetSuccessStacks() []string`

GetSuccessStacks returns the SuccessStacks field if non-nil, zero value otherwise.

### GetSuccessStacksOk

`func (o *VlanNetworkDeliverDataVO) GetSuccessStacksOk() (*[]string, bool)`

GetSuccessStacksOk returns a tuple with the SuccessStacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessStacks

`func (o *VlanNetworkDeliverDataVO) SetSuccessStacks(v []string)`

SetSuccessStacks sets SuccessStacks field to given value.

### HasSuccessStacks

`func (o *VlanNetworkDeliverDataVO) HasSuccessStacks() bool`

HasSuccessStacks returns a boolean if a field has been set.

### GetTotalDevicesNum

`func (o *VlanNetworkDeliverDataVO) GetTotalDevicesNum() int32`

GetTotalDevicesNum returns the TotalDevicesNum field if non-nil, zero value otherwise.

### GetTotalDevicesNumOk

`func (o *VlanNetworkDeliverDataVO) GetTotalDevicesNumOk() (*int32, bool)`

GetTotalDevicesNumOk returns a tuple with the TotalDevicesNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDevicesNum

`func (o *VlanNetworkDeliverDataVO) SetTotalDevicesNum(v int32)`

SetTotalDevicesNum sets TotalDevicesNum field to given value.

### HasTotalDevicesNum

`func (o *VlanNetworkDeliverDataVO) HasTotalDevicesNum() bool`

HasTotalDevicesNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


