# SelectStackLagForVlanVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Whether the port is affected. | [optional] 
**AutoSelect** | Pointer to **bool** | Whether the port needs to be automatically selected. | [optional] 
**DefaultVlan** | Pointer to **int32** | The vlan of default network. | [optional] 
**DownlinkDevices** | Pointer to [**[]DeviceBriefVO**](DeviceBriefVO.md) | The downlink devices of the port | [optional] 
**EditEnable** | Pointer to **bool** | Whether the port is selectable. | [optional] 
**Id** | Pointer to **int32** | Lag ID | [optional] 
**Name** | Pointer to **string** | Name. | [optional] 
**NativeIsDefault** | Pointer to **bool** | It indicates whether the native network of the port is default. | [optional] 
**NativeNetworkName** | Pointer to **string** | The native network name of the port. | [optional] 
**NativeNetworkVlan** | Pointer to **int32** | The native network vlan of the port. | [optional] 
**NeedConfirm** | Pointer to **bool** | Whether the port needs confirm for binding non-default vlan. | [optional] 
**NeedConfirmCascadePort** | Pointer to **bool** | Whether the port needs confirm for being cascade port. | [optional] 
**NeedConfirmVoiceNetwork** | Pointer to **bool** | When creating single vlan and the port&#39;s voice network is enabled, needConfirmVoiceNetwork is true. | [optional] 
**OswPortNetworkTagsSetting** | Pointer to **int32** | The port network tag setting. 0:allow all, 1:block all, 2:custom | [optional] 
**Ports** | Pointer to **[]int32** | The ports in the lag. Each item is Integer, for example: [1, 2]. | [optional] 
**Reasons** | Pointer to **[]int32** | Only valid when editEnable is false. It indicates the reason why the port is not selectable.Each Item should be a value as follows: -2: The port has been added to LAG. Lag member port can not be selected; -8: The port is wan port. Wan port can not be selected; -10: The AP’s LAN port cannot be configured; -12: The port&#39;s tags setting is not custom and therefore it can not be selected when creating multi vlan; -13: The port is stack port. Stack port can not be selected; -14: The number of VLANs has reached the limit of the Easy Managed Switch. | [optional] 
**SelectByUser** | Pointer to **bool** | Whether the port is selected by user. | [optional] 
**StPorts** | Pointer to **[]string** | The ports in the lag. Each item is String, for example: [1/0/1, 1/0/2]. | [optional] 
**UnSelectByUser** | Pointer to **bool** | Whether the auto-select port is unselected by user. | [optional] 
**UpperDevice** | Pointer to [**DeviceBriefVO**](DeviceBriefVO.md) |  | [optional] 

## Methods

### NewSelectStackLagForVlanVO

`func NewSelectStackLagForVlanVO() *SelectStackLagForVlanVO`

NewSelectStackLagForVlanVO instantiates a new SelectStackLagForVlanVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectStackLagForVlanVOWithDefaults

`func NewSelectStackLagForVlanVOWithDefaults() *SelectStackLagForVlanVO`

NewSelectStackLagForVlanVOWithDefaults instantiates a new SelectStackLagForVlanVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *SelectStackLagForVlanVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *SelectStackLagForVlanVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *SelectStackLagForVlanVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *SelectStackLagForVlanVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetAutoSelect

`func (o *SelectStackLagForVlanVO) GetAutoSelect() bool`

GetAutoSelect returns the AutoSelect field if non-nil, zero value otherwise.

### GetAutoSelectOk

`func (o *SelectStackLagForVlanVO) GetAutoSelectOk() (*bool, bool)`

GetAutoSelectOk returns a tuple with the AutoSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSelect

`func (o *SelectStackLagForVlanVO) SetAutoSelect(v bool)`

SetAutoSelect sets AutoSelect field to given value.

### HasAutoSelect

`func (o *SelectStackLagForVlanVO) HasAutoSelect() bool`

HasAutoSelect returns a boolean if a field has been set.

### GetDefaultVlan

`func (o *SelectStackLagForVlanVO) GetDefaultVlan() int32`

GetDefaultVlan returns the DefaultVlan field if non-nil, zero value otherwise.

### GetDefaultVlanOk

`func (o *SelectStackLagForVlanVO) GetDefaultVlanOk() (*int32, bool)`

GetDefaultVlanOk returns a tuple with the DefaultVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlan

`func (o *SelectStackLagForVlanVO) SetDefaultVlan(v int32)`

SetDefaultVlan sets DefaultVlan field to given value.

### HasDefaultVlan

`func (o *SelectStackLagForVlanVO) HasDefaultVlan() bool`

HasDefaultVlan returns a boolean if a field has been set.

### GetDownlinkDevices

`func (o *SelectStackLagForVlanVO) GetDownlinkDevices() []DeviceBriefVO`

GetDownlinkDevices returns the DownlinkDevices field if non-nil, zero value otherwise.

### GetDownlinkDevicesOk

`func (o *SelectStackLagForVlanVO) GetDownlinkDevicesOk() (*[]DeviceBriefVO, bool)`

GetDownlinkDevicesOk returns a tuple with the DownlinkDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkDevices

`func (o *SelectStackLagForVlanVO) SetDownlinkDevices(v []DeviceBriefVO)`

SetDownlinkDevices sets DownlinkDevices field to given value.

### HasDownlinkDevices

`func (o *SelectStackLagForVlanVO) HasDownlinkDevices() bool`

HasDownlinkDevices returns a boolean if a field has been set.

### GetEditEnable

`func (o *SelectStackLagForVlanVO) GetEditEnable() bool`

GetEditEnable returns the EditEnable field if non-nil, zero value otherwise.

### GetEditEnableOk

`func (o *SelectStackLagForVlanVO) GetEditEnableOk() (*bool, bool)`

GetEditEnableOk returns a tuple with the EditEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditEnable

`func (o *SelectStackLagForVlanVO) SetEditEnable(v bool)`

SetEditEnable sets EditEnable field to given value.

### HasEditEnable

`func (o *SelectStackLagForVlanVO) HasEditEnable() bool`

HasEditEnable returns a boolean if a field has been set.

### GetId

`func (o *SelectStackLagForVlanVO) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SelectStackLagForVlanVO) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SelectStackLagForVlanVO) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SelectStackLagForVlanVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SelectStackLagForVlanVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SelectStackLagForVlanVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SelectStackLagForVlanVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SelectStackLagForVlanVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNativeIsDefault

`func (o *SelectStackLagForVlanVO) GetNativeIsDefault() bool`

GetNativeIsDefault returns the NativeIsDefault field if non-nil, zero value otherwise.

### GetNativeIsDefaultOk

`func (o *SelectStackLagForVlanVO) GetNativeIsDefaultOk() (*bool, bool)`

GetNativeIsDefaultOk returns a tuple with the NativeIsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeIsDefault

`func (o *SelectStackLagForVlanVO) SetNativeIsDefault(v bool)`

SetNativeIsDefault sets NativeIsDefault field to given value.

### HasNativeIsDefault

`func (o *SelectStackLagForVlanVO) HasNativeIsDefault() bool`

HasNativeIsDefault returns a boolean if a field has been set.

### GetNativeNetworkName

`func (o *SelectStackLagForVlanVO) GetNativeNetworkName() string`

GetNativeNetworkName returns the NativeNetworkName field if non-nil, zero value otherwise.

### GetNativeNetworkNameOk

`func (o *SelectStackLagForVlanVO) GetNativeNetworkNameOk() (*string, bool)`

GetNativeNetworkNameOk returns a tuple with the NativeNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeNetworkName

`func (o *SelectStackLagForVlanVO) SetNativeNetworkName(v string)`

SetNativeNetworkName sets NativeNetworkName field to given value.

### HasNativeNetworkName

`func (o *SelectStackLagForVlanVO) HasNativeNetworkName() bool`

HasNativeNetworkName returns a boolean if a field has been set.

### GetNativeNetworkVlan

`func (o *SelectStackLagForVlanVO) GetNativeNetworkVlan() int32`

GetNativeNetworkVlan returns the NativeNetworkVlan field if non-nil, zero value otherwise.

### GetNativeNetworkVlanOk

`func (o *SelectStackLagForVlanVO) GetNativeNetworkVlanOk() (*int32, bool)`

GetNativeNetworkVlanOk returns a tuple with the NativeNetworkVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeNetworkVlan

`func (o *SelectStackLagForVlanVO) SetNativeNetworkVlan(v int32)`

SetNativeNetworkVlan sets NativeNetworkVlan field to given value.

### HasNativeNetworkVlan

`func (o *SelectStackLagForVlanVO) HasNativeNetworkVlan() bool`

HasNativeNetworkVlan returns a boolean if a field has been set.

### GetNeedConfirm

`func (o *SelectStackLagForVlanVO) GetNeedConfirm() bool`

GetNeedConfirm returns the NeedConfirm field if non-nil, zero value otherwise.

### GetNeedConfirmOk

`func (o *SelectStackLagForVlanVO) GetNeedConfirmOk() (*bool, bool)`

GetNeedConfirmOk returns a tuple with the NeedConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedConfirm

`func (o *SelectStackLagForVlanVO) SetNeedConfirm(v bool)`

SetNeedConfirm sets NeedConfirm field to given value.

### HasNeedConfirm

`func (o *SelectStackLagForVlanVO) HasNeedConfirm() bool`

HasNeedConfirm returns a boolean if a field has been set.

### GetNeedConfirmCascadePort

`func (o *SelectStackLagForVlanVO) GetNeedConfirmCascadePort() bool`

GetNeedConfirmCascadePort returns the NeedConfirmCascadePort field if non-nil, zero value otherwise.

### GetNeedConfirmCascadePortOk

`func (o *SelectStackLagForVlanVO) GetNeedConfirmCascadePortOk() (*bool, bool)`

GetNeedConfirmCascadePortOk returns a tuple with the NeedConfirmCascadePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedConfirmCascadePort

`func (o *SelectStackLagForVlanVO) SetNeedConfirmCascadePort(v bool)`

SetNeedConfirmCascadePort sets NeedConfirmCascadePort field to given value.

### HasNeedConfirmCascadePort

`func (o *SelectStackLagForVlanVO) HasNeedConfirmCascadePort() bool`

HasNeedConfirmCascadePort returns a boolean if a field has been set.

### GetNeedConfirmVoiceNetwork

`func (o *SelectStackLagForVlanVO) GetNeedConfirmVoiceNetwork() bool`

GetNeedConfirmVoiceNetwork returns the NeedConfirmVoiceNetwork field if non-nil, zero value otherwise.

### GetNeedConfirmVoiceNetworkOk

`func (o *SelectStackLagForVlanVO) GetNeedConfirmVoiceNetworkOk() (*bool, bool)`

GetNeedConfirmVoiceNetworkOk returns a tuple with the NeedConfirmVoiceNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedConfirmVoiceNetwork

`func (o *SelectStackLagForVlanVO) SetNeedConfirmVoiceNetwork(v bool)`

SetNeedConfirmVoiceNetwork sets NeedConfirmVoiceNetwork field to given value.

### HasNeedConfirmVoiceNetwork

`func (o *SelectStackLagForVlanVO) HasNeedConfirmVoiceNetwork() bool`

HasNeedConfirmVoiceNetwork returns a boolean if a field has been set.

### GetOswPortNetworkTagsSetting

`func (o *SelectStackLagForVlanVO) GetOswPortNetworkTagsSetting() int32`

GetOswPortNetworkTagsSetting returns the OswPortNetworkTagsSetting field if non-nil, zero value otherwise.

### GetOswPortNetworkTagsSettingOk

`func (o *SelectStackLagForVlanVO) GetOswPortNetworkTagsSettingOk() (*int32, bool)`

GetOswPortNetworkTagsSettingOk returns a tuple with the OswPortNetworkTagsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOswPortNetworkTagsSetting

`func (o *SelectStackLagForVlanVO) SetOswPortNetworkTagsSetting(v int32)`

SetOswPortNetworkTagsSetting sets OswPortNetworkTagsSetting field to given value.

### HasOswPortNetworkTagsSetting

`func (o *SelectStackLagForVlanVO) HasOswPortNetworkTagsSetting() bool`

HasOswPortNetworkTagsSetting returns a boolean if a field has been set.

### GetPorts

`func (o *SelectStackLagForVlanVO) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *SelectStackLagForVlanVO) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *SelectStackLagForVlanVO) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *SelectStackLagForVlanVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetReasons

`func (o *SelectStackLagForVlanVO) GetReasons() []int32`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *SelectStackLagForVlanVO) GetReasonsOk() (*[]int32, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *SelectStackLagForVlanVO) SetReasons(v []int32)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *SelectStackLagForVlanVO) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetSelectByUser

`func (o *SelectStackLagForVlanVO) GetSelectByUser() bool`

GetSelectByUser returns the SelectByUser field if non-nil, zero value otherwise.

### GetSelectByUserOk

`func (o *SelectStackLagForVlanVO) GetSelectByUserOk() (*bool, bool)`

GetSelectByUserOk returns a tuple with the SelectByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectByUser

`func (o *SelectStackLagForVlanVO) SetSelectByUser(v bool)`

SetSelectByUser sets SelectByUser field to given value.

### HasSelectByUser

`func (o *SelectStackLagForVlanVO) HasSelectByUser() bool`

HasSelectByUser returns a boolean if a field has been set.

### GetStPorts

`func (o *SelectStackLagForVlanVO) GetStPorts() []string`

GetStPorts returns the StPorts field if non-nil, zero value otherwise.

### GetStPortsOk

`func (o *SelectStackLagForVlanVO) GetStPortsOk() (*[]string, bool)`

GetStPortsOk returns a tuple with the StPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStPorts

`func (o *SelectStackLagForVlanVO) SetStPorts(v []string)`

SetStPorts sets StPorts field to given value.

### HasStPorts

`func (o *SelectStackLagForVlanVO) HasStPorts() bool`

HasStPorts returns a boolean if a field has been set.

### GetUnSelectByUser

`func (o *SelectStackLagForVlanVO) GetUnSelectByUser() bool`

GetUnSelectByUser returns the UnSelectByUser field if non-nil, zero value otherwise.

### GetUnSelectByUserOk

`func (o *SelectStackLagForVlanVO) GetUnSelectByUserOk() (*bool, bool)`

GetUnSelectByUserOk returns a tuple with the UnSelectByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnSelectByUser

`func (o *SelectStackLagForVlanVO) SetUnSelectByUser(v bool)`

SetUnSelectByUser sets UnSelectByUser field to given value.

### HasUnSelectByUser

`func (o *SelectStackLagForVlanVO) HasUnSelectByUser() bool`

HasUnSelectByUser returns a boolean if a field has been set.

### GetUpperDevice

`func (o *SelectStackLagForVlanVO) GetUpperDevice() DeviceBriefVO`

GetUpperDevice returns the UpperDevice field if non-nil, zero value otherwise.

### GetUpperDeviceOk

`func (o *SelectStackLagForVlanVO) GetUpperDeviceOk() (*DeviceBriefVO, bool)`

GetUpperDeviceOk returns a tuple with the UpperDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperDevice

`func (o *SelectStackLagForVlanVO) SetUpperDevice(v DeviceBriefVO)`

SetUpperDevice sets UpperDevice field to given value.

### HasUpperDevice

`func (o *SelectStackLagForVlanVO) HasUpperDevice() bool`

HasUpperDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


