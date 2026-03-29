# SelectLagForVlanVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Whether the port is affected. | [optional] 
**AutoSelect** | Pointer to **bool** | Whether the port needs to be automatically selected. | [optional] 
**DefaultVlan** | Pointer to **int32** | The vlan of default network. | [optional] 
**DownlinkDevices** | Pointer to [**[]DeviceBriefVO**](DeviceBriefVO.md) | The downlink devices of the port | [optional] 
**EditEnable** | Pointer to **bool** | Whether the port is selectable. | [optional] 
**Id** | Pointer to **int32** | Lag ID | [optional] 
**MlagEnable** | Pointer to **bool** | Indicates whether M-LAG port is enabled | [optional] 
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
**StPorts** | Pointer to **[]string** | The ports in the lag. Each item is String, for example: [\&quot;1/0/1\&quot;, \&quot;1/0/2\&quot;]. | [optional] 
**UnSelectByUser** | Pointer to **bool** | Whether the auto-select port is unselected by user. | [optional] 
**UpperDevice** | Pointer to [**DeviceBriefVO**](DeviceBriefVO.md) |  | [optional] 

## Methods

### NewSelectLagForVlanVO

`func NewSelectLagForVlanVO() *SelectLagForVlanVO`

NewSelectLagForVlanVO instantiates a new SelectLagForVlanVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectLagForVlanVOWithDefaults

`func NewSelectLagForVlanVOWithDefaults() *SelectLagForVlanVO`

NewSelectLagForVlanVOWithDefaults instantiates a new SelectLagForVlanVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *SelectLagForVlanVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *SelectLagForVlanVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *SelectLagForVlanVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *SelectLagForVlanVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetAutoSelect

`func (o *SelectLagForVlanVO) GetAutoSelect() bool`

GetAutoSelect returns the AutoSelect field if non-nil, zero value otherwise.

### GetAutoSelectOk

`func (o *SelectLagForVlanVO) GetAutoSelectOk() (*bool, bool)`

GetAutoSelectOk returns a tuple with the AutoSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSelect

`func (o *SelectLagForVlanVO) SetAutoSelect(v bool)`

SetAutoSelect sets AutoSelect field to given value.

### HasAutoSelect

`func (o *SelectLagForVlanVO) HasAutoSelect() bool`

HasAutoSelect returns a boolean if a field has been set.

### GetDefaultVlan

`func (o *SelectLagForVlanVO) GetDefaultVlan() int32`

GetDefaultVlan returns the DefaultVlan field if non-nil, zero value otherwise.

### GetDefaultVlanOk

`func (o *SelectLagForVlanVO) GetDefaultVlanOk() (*int32, bool)`

GetDefaultVlanOk returns a tuple with the DefaultVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlan

`func (o *SelectLagForVlanVO) SetDefaultVlan(v int32)`

SetDefaultVlan sets DefaultVlan field to given value.

### HasDefaultVlan

`func (o *SelectLagForVlanVO) HasDefaultVlan() bool`

HasDefaultVlan returns a boolean if a field has been set.

### GetDownlinkDevices

`func (o *SelectLagForVlanVO) GetDownlinkDevices() []DeviceBriefVO`

GetDownlinkDevices returns the DownlinkDevices field if non-nil, zero value otherwise.

### GetDownlinkDevicesOk

`func (o *SelectLagForVlanVO) GetDownlinkDevicesOk() (*[]DeviceBriefVO, bool)`

GetDownlinkDevicesOk returns a tuple with the DownlinkDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkDevices

`func (o *SelectLagForVlanVO) SetDownlinkDevices(v []DeviceBriefVO)`

SetDownlinkDevices sets DownlinkDevices field to given value.

### HasDownlinkDevices

`func (o *SelectLagForVlanVO) HasDownlinkDevices() bool`

HasDownlinkDevices returns a boolean if a field has been set.

### GetEditEnable

`func (o *SelectLagForVlanVO) GetEditEnable() bool`

GetEditEnable returns the EditEnable field if non-nil, zero value otherwise.

### GetEditEnableOk

`func (o *SelectLagForVlanVO) GetEditEnableOk() (*bool, bool)`

GetEditEnableOk returns a tuple with the EditEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditEnable

`func (o *SelectLagForVlanVO) SetEditEnable(v bool)`

SetEditEnable sets EditEnable field to given value.

### HasEditEnable

`func (o *SelectLagForVlanVO) HasEditEnable() bool`

HasEditEnable returns a boolean if a field has been set.

### GetId

`func (o *SelectLagForVlanVO) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SelectLagForVlanVO) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SelectLagForVlanVO) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SelectLagForVlanVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMlagEnable

`func (o *SelectLagForVlanVO) GetMlagEnable() bool`

GetMlagEnable returns the MlagEnable field if non-nil, zero value otherwise.

### GetMlagEnableOk

`func (o *SelectLagForVlanVO) GetMlagEnableOk() (*bool, bool)`

GetMlagEnableOk returns a tuple with the MlagEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagEnable

`func (o *SelectLagForVlanVO) SetMlagEnable(v bool)`

SetMlagEnable sets MlagEnable field to given value.

### HasMlagEnable

`func (o *SelectLagForVlanVO) HasMlagEnable() bool`

HasMlagEnable returns a boolean if a field has been set.

### GetName

`func (o *SelectLagForVlanVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SelectLagForVlanVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SelectLagForVlanVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SelectLagForVlanVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNativeIsDefault

`func (o *SelectLagForVlanVO) GetNativeIsDefault() bool`

GetNativeIsDefault returns the NativeIsDefault field if non-nil, zero value otherwise.

### GetNativeIsDefaultOk

`func (o *SelectLagForVlanVO) GetNativeIsDefaultOk() (*bool, bool)`

GetNativeIsDefaultOk returns a tuple with the NativeIsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeIsDefault

`func (o *SelectLagForVlanVO) SetNativeIsDefault(v bool)`

SetNativeIsDefault sets NativeIsDefault field to given value.

### HasNativeIsDefault

`func (o *SelectLagForVlanVO) HasNativeIsDefault() bool`

HasNativeIsDefault returns a boolean if a field has been set.

### GetNativeNetworkName

`func (o *SelectLagForVlanVO) GetNativeNetworkName() string`

GetNativeNetworkName returns the NativeNetworkName field if non-nil, zero value otherwise.

### GetNativeNetworkNameOk

`func (o *SelectLagForVlanVO) GetNativeNetworkNameOk() (*string, bool)`

GetNativeNetworkNameOk returns a tuple with the NativeNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeNetworkName

`func (o *SelectLagForVlanVO) SetNativeNetworkName(v string)`

SetNativeNetworkName sets NativeNetworkName field to given value.

### HasNativeNetworkName

`func (o *SelectLagForVlanVO) HasNativeNetworkName() bool`

HasNativeNetworkName returns a boolean if a field has been set.

### GetNativeNetworkVlan

`func (o *SelectLagForVlanVO) GetNativeNetworkVlan() int32`

GetNativeNetworkVlan returns the NativeNetworkVlan field if non-nil, zero value otherwise.

### GetNativeNetworkVlanOk

`func (o *SelectLagForVlanVO) GetNativeNetworkVlanOk() (*int32, bool)`

GetNativeNetworkVlanOk returns a tuple with the NativeNetworkVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeNetworkVlan

`func (o *SelectLagForVlanVO) SetNativeNetworkVlan(v int32)`

SetNativeNetworkVlan sets NativeNetworkVlan field to given value.

### HasNativeNetworkVlan

`func (o *SelectLagForVlanVO) HasNativeNetworkVlan() bool`

HasNativeNetworkVlan returns a boolean if a field has been set.

### GetNeedConfirm

`func (o *SelectLagForVlanVO) GetNeedConfirm() bool`

GetNeedConfirm returns the NeedConfirm field if non-nil, zero value otherwise.

### GetNeedConfirmOk

`func (o *SelectLagForVlanVO) GetNeedConfirmOk() (*bool, bool)`

GetNeedConfirmOk returns a tuple with the NeedConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedConfirm

`func (o *SelectLagForVlanVO) SetNeedConfirm(v bool)`

SetNeedConfirm sets NeedConfirm field to given value.

### HasNeedConfirm

`func (o *SelectLagForVlanVO) HasNeedConfirm() bool`

HasNeedConfirm returns a boolean if a field has been set.

### GetNeedConfirmCascadePort

`func (o *SelectLagForVlanVO) GetNeedConfirmCascadePort() bool`

GetNeedConfirmCascadePort returns the NeedConfirmCascadePort field if non-nil, zero value otherwise.

### GetNeedConfirmCascadePortOk

`func (o *SelectLagForVlanVO) GetNeedConfirmCascadePortOk() (*bool, bool)`

GetNeedConfirmCascadePortOk returns a tuple with the NeedConfirmCascadePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedConfirmCascadePort

`func (o *SelectLagForVlanVO) SetNeedConfirmCascadePort(v bool)`

SetNeedConfirmCascadePort sets NeedConfirmCascadePort field to given value.

### HasNeedConfirmCascadePort

`func (o *SelectLagForVlanVO) HasNeedConfirmCascadePort() bool`

HasNeedConfirmCascadePort returns a boolean if a field has been set.

### GetNeedConfirmVoiceNetwork

`func (o *SelectLagForVlanVO) GetNeedConfirmVoiceNetwork() bool`

GetNeedConfirmVoiceNetwork returns the NeedConfirmVoiceNetwork field if non-nil, zero value otherwise.

### GetNeedConfirmVoiceNetworkOk

`func (o *SelectLagForVlanVO) GetNeedConfirmVoiceNetworkOk() (*bool, bool)`

GetNeedConfirmVoiceNetworkOk returns a tuple with the NeedConfirmVoiceNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedConfirmVoiceNetwork

`func (o *SelectLagForVlanVO) SetNeedConfirmVoiceNetwork(v bool)`

SetNeedConfirmVoiceNetwork sets NeedConfirmVoiceNetwork field to given value.

### HasNeedConfirmVoiceNetwork

`func (o *SelectLagForVlanVO) HasNeedConfirmVoiceNetwork() bool`

HasNeedConfirmVoiceNetwork returns a boolean if a field has been set.

### GetOswPortNetworkTagsSetting

`func (o *SelectLagForVlanVO) GetOswPortNetworkTagsSetting() int32`

GetOswPortNetworkTagsSetting returns the OswPortNetworkTagsSetting field if non-nil, zero value otherwise.

### GetOswPortNetworkTagsSettingOk

`func (o *SelectLagForVlanVO) GetOswPortNetworkTagsSettingOk() (*int32, bool)`

GetOswPortNetworkTagsSettingOk returns a tuple with the OswPortNetworkTagsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOswPortNetworkTagsSetting

`func (o *SelectLagForVlanVO) SetOswPortNetworkTagsSetting(v int32)`

SetOswPortNetworkTagsSetting sets OswPortNetworkTagsSetting field to given value.

### HasOswPortNetworkTagsSetting

`func (o *SelectLagForVlanVO) HasOswPortNetworkTagsSetting() bool`

HasOswPortNetworkTagsSetting returns a boolean if a field has been set.

### GetPorts

`func (o *SelectLagForVlanVO) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *SelectLagForVlanVO) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *SelectLagForVlanVO) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *SelectLagForVlanVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetReasons

`func (o *SelectLagForVlanVO) GetReasons() []int32`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *SelectLagForVlanVO) GetReasonsOk() (*[]int32, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *SelectLagForVlanVO) SetReasons(v []int32)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *SelectLagForVlanVO) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetSelectByUser

`func (o *SelectLagForVlanVO) GetSelectByUser() bool`

GetSelectByUser returns the SelectByUser field if non-nil, zero value otherwise.

### GetSelectByUserOk

`func (o *SelectLagForVlanVO) GetSelectByUserOk() (*bool, bool)`

GetSelectByUserOk returns a tuple with the SelectByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectByUser

`func (o *SelectLagForVlanVO) SetSelectByUser(v bool)`

SetSelectByUser sets SelectByUser field to given value.

### HasSelectByUser

`func (o *SelectLagForVlanVO) HasSelectByUser() bool`

HasSelectByUser returns a boolean if a field has been set.

### GetStPorts

`func (o *SelectLagForVlanVO) GetStPorts() []string`

GetStPorts returns the StPorts field if non-nil, zero value otherwise.

### GetStPortsOk

`func (o *SelectLagForVlanVO) GetStPortsOk() (*[]string, bool)`

GetStPortsOk returns a tuple with the StPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStPorts

`func (o *SelectLagForVlanVO) SetStPorts(v []string)`

SetStPorts sets StPorts field to given value.

### HasStPorts

`func (o *SelectLagForVlanVO) HasStPorts() bool`

HasStPorts returns a boolean if a field has been set.

### GetUnSelectByUser

`func (o *SelectLagForVlanVO) GetUnSelectByUser() bool`

GetUnSelectByUser returns the UnSelectByUser field if non-nil, zero value otherwise.

### GetUnSelectByUserOk

`func (o *SelectLagForVlanVO) GetUnSelectByUserOk() (*bool, bool)`

GetUnSelectByUserOk returns a tuple with the UnSelectByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnSelectByUser

`func (o *SelectLagForVlanVO) SetUnSelectByUser(v bool)`

SetUnSelectByUser sets UnSelectByUser field to given value.

### HasUnSelectByUser

`func (o *SelectLagForVlanVO) HasUnSelectByUser() bool`

HasUnSelectByUser returns a boolean if a field has been set.

### GetUpperDevice

`func (o *SelectLagForVlanVO) GetUpperDevice() DeviceBriefVO`

GetUpperDevice returns the UpperDevice field if non-nil, zero value otherwise.

### GetUpperDeviceOk

`func (o *SelectLagForVlanVO) GetUpperDeviceOk() (*DeviceBriefVO, bool)`

GetUpperDeviceOk returns a tuple with the UpperDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperDevice

`func (o *SelectLagForVlanVO) SetUpperDevice(v DeviceBriefVO)`

SetUpperDevice sets UpperDevice field to given value.

### HasUpperDevice

`func (o *SelectLagForVlanVO) HasUpperDevice() bool`

HasUpperDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


