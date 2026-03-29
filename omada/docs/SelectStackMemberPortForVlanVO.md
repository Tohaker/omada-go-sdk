# SelectStackMemberPortForVlanVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Whether the port is affected. | [optional] 
**AutoSelect** | Pointer to **bool** | Whether the port needs to be automatically selected. | [optional] 
**ConfigStack** | Pointer to **bool** | Whether the port is configured as a stack port. | [optional] 
**DefaultVlan** | Pointer to **int32** | The vlan of default network. | [optional] 
**DownlinkDevices** | Pointer to [**[]DeviceBriefVO**](DeviceBriefVO.md) | The downlink devices of the port | [optional] 
**EditEnable** | Pointer to **bool** | Whether the port is selectable. | [optional] 
**Id** | Pointer to **int32** | port number, for example: 1 | [optional] 
**LagId** | Pointer to **int32** | Lag ID. It indicates the lag id of the port when the port is in a lag. | [optional] 
**MadEnable** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** | Name. | [optional] 
**NativeIsDefault** | Pointer to **bool** | It indicates whether the native network of the port is default. | [optional] 
**NativeNetworkName** | Pointer to **string** | The native network name of the port. | [optional] 
**NativeNetworkVlan** | Pointer to **int32** | The native network vlan of the port. | [optional] 
**NeedConfirm** | Pointer to **bool** | Whether the port needs confirm for binding non-default vlan. | [optional] 
**NeedConfirmCascadePort** | Pointer to **bool** | Whether the port needs confirm for being cascade port. | [optional] 
**NeedConfirmVoiceNetwork** | Pointer to **bool** | When creating single vlan and the port&#39;s voice network is enabled, needConfirmVoiceNetwork is true. | [optional] 
**OswPortNetworkTagsSetting** | Pointer to **int32** | The port network tag setting. 0:allow all, 1:block all, 2:custom | [optional] 
**PortId** | Pointer to **string** | ID, for example: \&quot;1/0/1\&quot; | [optional] 
**Reasons** | Pointer to **[]int32** | Only valid when editEnable is false. It indicates the reason why the port is not selectable.Each Item should be a value as follows: -2: The port has been added to LAG. Lag member port can not be selected; -8: The port is wan port. Wan port can not be selected; -10: The AP’s LAN port cannot be configured; -12: The port&#39;s tags setting is not custom and therefore it can not be selected when creating multi vlan; -13: The port is stack port. Stack port can not be selected; -14: The number of VLANs has reached the limit of the Easy Managed Switch. | [optional] 
**SelectByUser** | Pointer to **bool** | Whether the port is selected by user. | [optional] 
**Type** | Pointer to **int32** | Type. 1: Copper, 2: Combo, 3: SFP | [optional] 
**UnSelectByUser** | Pointer to **bool** | Whether the auto-select port is unselected by user. | [optional] 
**UpperDevice** | Pointer to [**DeviceBriefVO**](DeviceBriefVO.md) |  | [optional] 

## Methods

### NewSelectStackMemberPortForVlanVO

`func NewSelectStackMemberPortForVlanVO() *SelectStackMemberPortForVlanVO`

NewSelectStackMemberPortForVlanVO instantiates a new SelectStackMemberPortForVlanVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectStackMemberPortForVlanVOWithDefaults

`func NewSelectStackMemberPortForVlanVOWithDefaults() *SelectStackMemberPortForVlanVO`

NewSelectStackMemberPortForVlanVOWithDefaults instantiates a new SelectStackMemberPortForVlanVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *SelectStackMemberPortForVlanVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *SelectStackMemberPortForVlanVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *SelectStackMemberPortForVlanVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *SelectStackMemberPortForVlanVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetAutoSelect

`func (o *SelectStackMemberPortForVlanVO) GetAutoSelect() bool`

GetAutoSelect returns the AutoSelect field if non-nil, zero value otherwise.

### GetAutoSelectOk

`func (o *SelectStackMemberPortForVlanVO) GetAutoSelectOk() (*bool, bool)`

GetAutoSelectOk returns a tuple with the AutoSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSelect

`func (o *SelectStackMemberPortForVlanVO) SetAutoSelect(v bool)`

SetAutoSelect sets AutoSelect field to given value.

### HasAutoSelect

`func (o *SelectStackMemberPortForVlanVO) HasAutoSelect() bool`

HasAutoSelect returns a boolean if a field has been set.

### GetConfigStack

`func (o *SelectStackMemberPortForVlanVO) GetConfigStack() bool`

GetConfigStack returns the ConfigStack field if non-nil, zero value otherwise.

### GetConfigStackOk

`func (o *SelectStackMemberPortForVlanVO) GetConfigStackOk() (*bool, bool)`

GetConfigStackOk returns a tuple with the ConfigStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStack

`func (o *SelectStackMemberPortForVlanVO) SetConfigStack(v bool)`

SetConfigStack sets ConfigStack field to given value.

### HasConfigStack

`func (o *SelectStackMemberPortForVlanVO) HasConfigStack() bool`

HasConfigStack returns a boolean if a field has been set.

### GetDefaultVlan

`func (o *SelectStackMemberPortForVlanVO) GetDefaultVlan() int32`

GetDefaultVlan returns the DefaultVlan field if non-nil, zero value otherwise.

### GetDefaultVlanOk

`func (o *SelectStackMemberPortForVlanVO) GetDefaultVlanOk() (*int32, bool)`

GetDefaultVlanOk returns a tuple with the DefaultVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlan

`func (o *SelectStackMemberPortForVlanVO) SetDefaultVlan(v int32)`

SetDefaultVlan sets DefaultVlan field to given value.

### HasDefaultVlan

`func (o *SelectStackMemberPortForVlanVO) HasDefaultVlan() bool`

HasDefaultVlan returns a boolean if a field has been set.

### GetDownlinkDevices

`func (o *SelectStackMemberPortForVlanVO) GetDownlinkDevices() []DeviceBriefVO`

GetDownlinkDevices returns the DownlinkDevices field if non-nil, zero value otherwise.

### GetDownlinkDevicesOk

`func (o *SelectStackMemberPortForVlanVO) GetDownlinkDevicesOk() (*[]DeviceBriefVO, bool)`

GetDownlinkDevicesOk returns a tuple with the DownlinkDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkDevices

`func (o *SelectStackMemberPortForVlanVO) SetDownlinkDevices(v []DeviceBriefVO)`

SetDownlinkDevices sets DownlinkDevices field to given value.

### HasDownlinkDevices

`func (o *SelectStackMemberPortForVlanVO) HasDownlinkDevices() bool`

HasDownlinkDevices returns a boolean if a field has been set.

### GetEditEnable

`func (o *SelectStackMemberPortForVlanVO) GetEditEnable() bool`

GetEditEnable returns the EditEnable field if non-nil, zero value otherwise.

### GetEditEnableOk

`func (o *SelectStackMemberPortForVlanVO) GetEditEnableOk() (*bool, bool)`

GetEditEnableOk returns a tuple with the EditEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditEnable

`func (o *SelectStackMemberPortForVlanVO) SetEditEnable(v bool)`

SetEditEnable sets EditEnable field to given value.

### HasEditEnable

`func (o *SelectStackMemberPortForVlanVO) HasEditEnable() bool`

HasEditEnable returns a boolean if a field has been set.

### GetId

`func (o *SelectStackMemberPortForVlanVO) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SelectStackMemberPortForVlanVO) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SelectStackMemberPortForVlanVO) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SelectStackMemberPortForVlanVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLagId

`func (o *SelectStackMemberPortForVlanVO) GetLagId() int32`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *SelectStackMemberPortForVlanVO) GetLagIdOk() (*int32, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *SelectStackMemberPortForVlanVO) SetLagId(v int32)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *SelectStackMemberPortForVlanVO) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetMadEnable

`func (o *SelectStackMemberPortForVlanVO) GetMadEnable() bool`

GetMadEnable returns the MadEnable field if non-nil, zero value otherwise.

### GetMadEnableOk

`func (o *SelectStackMemberPortForVlanVO) GetMadEnableOk() (*bool, bool)`

GetMadEnableOk returns a tuple with the MadEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMadEnable

`func (o *SelectStackMemberPortForVlanVO) SetMadEnable(v bool)`

SetMadEnable sets MadEnable field to given value.

### HasMadEnable

`func (o *SelectStackMemberPortForVlanVO) HasMadEnable() bool`

HasMadEnable returns a boolean if a field has been set.

### GetName

`func (o *SelectStackMemberPortForVlanVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SelectStackMemberPortForVlanVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SelectStackMemberPortForVlanVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SelectStackMemberPortForVlanVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNativeIsDefault

`func (o *SelectStackMemberPortForVlanVO) GetNativeIsDefault() bool`

GetNativeIsDefault returns the NativeIsDefault field if non-nil, zero value otherwise.

### GetNativeIsDefaultOk

`func (o *SelectStackMemberPortForVlanVO) GetNativeIsDefaultOk() (*bool, bool)`

GetNativeIsDefaultOk returns a tuple with the NativeIsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeIsDefault

`func (o *SelectStackMemberPortForVlanVO) SetNativeIsDefault(v bool)`

SetNativeIsDefault sets NativeIsDefault field to given value.

### HasNativeIsDefault

`func (o *SelectStackMemberPortForVlanVO) HasNativeIsDefault() bool`

HasNativeIsDefault returns a boolean if a field has been set.

### GetNativeNetworkName

`func (o *SelectStackMemberPortForVlanVO) GetNativeNetworkName() string`

GetNativeNetworkName returns the NativeNetworkName field if non-nil, zero value otherwise.

### GetNativeNetworkNameOk

`func (o *SelectStackMemberPortForVlanVO) GetNativeNetworkNameOk() (*string, bool)`

GetNativeNetworkNameOk returns a tuple with the NativeNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeNetworkName

`func (o *SelectStackMemberPortForVlanVO) SetNativeNetworkName(v string)`

SetNativeNetworkName sets NativeNetworkName field to given value.

### HasNativeNetworkName

`func (o *SelectStackMemberPortForVlanVO) HasNativeNetworkName() bool`

HasNativeNetworkName returns a boolean if a field has been set.

### GetNativeNetworkVlan

`func (o *SelectStackMemberPortForVlanVO) GetNativeNetworkVlan() int32`

GetNativeNetworkVlan returns the NativeNetworkVlan field if non-nil, zero value otherwise.

### GetNativeNetworkVlanOk

`func (o *SelectStackMemberPortForVlanVO) GetNativeNetworkVlanOk() (*int32, bool)`

GetNativeNetworkVlanOk returns a tuple with the NativeNetworkVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeNetworkVlan

`func (o *SelectStackMemberPortForVlanVO) SetNativeNetworkVlan(v int32)`

SetNativeNetworkVlan sets NativeNetworkVlan field to given value.

### HasNativeNetworkVlan

`func (o *SelectStackMemberPortForVlanVO) HasNativeNetworkVlan() bool`

HasNativeNetworkVlan returns a boolean if a field has been set.

### GetNeedConfirm

`func (o *SelectStackMemberPortForVlanVO) GetNeedConfirm() bool`

GetNeedConfirm returns the NeedConfirm field if non-nil, zero value otherwise.

### GetNeedConfirmOk

`func (o *SelectStackMemberPortForVlanVO) GetNeedConfirmOk() (*bool, bool)`

GetNeedConfirmOk returns a tuple with the NeedConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedConfirm

`func (o *SelectStackMemberPortForVlanVO) SetNeedConfirm(v bool)`

SetNeedConfirm sets NeedConfirm field to given value.

### HasNeedConfirm

`func (o *SelectStackMemberPortForVlanVO) HasNeedConfirm() bool`

HasNeedConfirm returns a boolean if a field has been set.

### GetNeedConfirmCascadePort

`func (o *SelectStackMemberPortForVlanVO) GetNeedConfirmCascadePort() bool`

GetNeedConfirmCascadePort returns the NeedConfirmCascadePort field if non-nil, zero value otherwise.

### GetNeedConfirmCascadePortOk

`func (o *SelectStackMemberPortForVlanVO) GetNeedConfirmCascadePortOk() (*bool, bool)`

GetNeedConfirmCascadePortOk returns a tuple with the NeedConfirmCascadePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedConfirmCascadePort

`func (o *SelectStackMemberPortForVlanVO) SetNeedConfirmCascadePort(v bool)`

SetNeedConfirmCascadePort sets NeedConfirmCascadePort field to given value.

### HasNeedConfirmCascadePort

`func (o *SelectStackMemberPortForVlanVO) HasNeedConfirmCascadePort() bool`

HasNeedConfirmCascadePort returns a boolean if a field has been set.

### GetNeedConfirmVoiceNetwork

`func (o *SelectStackMemberPortForVlanVO) GetNeedConfirmVoiceNetwork() bool`

GetNeedConfirmVoiceNetwork returns the NeedConfirmVoiceNetwork field if non-nil, zero value otherwise.

### GetNeedConfirmVoiceNetworkOk

`func (o *SelectStackMemberPortForVlanVO) GetNeedConfirmVoiceNetworkOk() (*bool, bool)`

GetNeedConfirmVoiceNetworkOk returns a tuple with the NeedConfirmVoiceNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedConfirmVoiceNetwork

`func (o *SelectStackMemberPortForVlanVO) SetNeedConfirmVoiceNetwork(v bool)`

SetNeedConfirmVoiceNetwork sets NeedConfirmVoiceNetwork field to given value.

### HasNeedConfirmVoiceNetwork

`func (o *SelectStackMemberPortForVlanVO) HasNeedConfirmVoiceNetwork() bool`

HasNeedConfirmVoiceNetwork returns a boolean if a field has been set.

### GetOswPortNetworkTagsSetting

`func (o *SelectStackMemberPortForVlanVO) GetOswPortNetworkTagsSetting() int32`

GetOswPortNetworkTagsSetting returns the OswPortNetworkTagsSetting field if non-nil, zero value otherwise.

### GetOswPortNetworkTagsSettingOk

`func (o *SelectStackMemberPortForVlanVO) GetOswPortNetworkTagsSettingOk() (*int32, bool)`

GetOswPortNetworkTagsSettingOk returns a tuple with the OswPortNetworkTagsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOswPortNetworkTagsSetting

`func (o *SelectStackMemberPortForVlanVO) SetOswPortNetworkTagsSetting(v int32)`

SetOswPortNetworkTagsSetting sets OswPortNetworkTagsSetting field to given value.

### HasOswPortNetworkTagsSetting

`func (o *SelectStackMemberPortForVlanVO) HasOswPortNetworkTagsSetting() bool`

HasOswPortNetworkTagsSetting returns a boolean if a field has been set.

### GetPortId

`func (o *SelectStackMemberPortForVlanVO) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *SelectStackMemberPortForVlanVO) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *SelectStackMemberPortForVlanVO) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *SelectStackMemberPortForVlanVO) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetReasons

`func (o *SelectStackMemberPortForVlanVO) GetReasons() []int32`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *SelectStackMemberPortForVlanVO) GetReasonsOk() (*[]int32, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *SelectStackMemberPortForVlanVO) SetReasons(v []int32)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *SelectStackMemberPortForVlanVO) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetSelectByUser

`func (o *SelectStackMemberPortForVlanVO) GetSelectByUser() bool`

GetSelectByUser returns the SelectByUser field if non-nil, zero value otherwise.

### GetSelectByUserOk

`func (o *SelectStackMemberPortForVlanVO) GetSelectByUserOk() (*bool, bool)`

GetSelectByUserOk returns a tuple with the SelectByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectByUser

`func (o *SelectStackMemberPortForVlanVO) SetSelectByUser(v bool)`

SetSelectByUser sets SelectByUser field to given value.

### HasSelectByUser

`func (o *SelectStackMemberPortForVlanVO) HasSelectByUser() bool`

HasSelectByUser returns a boolean if a field has been set.

### GetType

`func (o *SelectStackMemberPortForVlanVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SelectStackMemberPortForVlanVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SelectStackMemberPortForVlanVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SelectStackMemberPortForVlanVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnSelectByUser

`func (o *SelectStackMemberPortForVlanVO) GetUnSelectByUser() bool`

GetUnSelectByUser returns the UnSelectByUser field if non-nil, zero value otherwise.

### GetUnSelectByUserOk

`func (o *SelectStackMemberPortForVlanVO) GetUnSelectByUserOk() (*bool, bool)`

GetUnSelectByUserOk returns a tuple with the UnSelectByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnSelectByUser

`func (o *SelectStackMemberPortForVlanVO) SetUnSelectByUser(v bool)`

SetUnSelectByUser sets UnSelectByUser field to given value.

### HasUnSelectByUser

`func (o *SelectStackMemberPortForVlanVO) HasUnSelectByUser() bool`

HasUnSelectByUser returns a boolean if a field has been set.

### GetUpperDevice

`func (o *SelectStackMemberPortForVlanVO) GetUpperDevice() DeviceBriefVO`

GetUpperDevice returns the UpperDevice field if non-nil, zero value otherwise.

### GetUpperDeviceOk

`func (o *SelectStackMemberPortForVlanVO) GetUpperDeviceOk() (*DeviceBriefVO, bool)`

GetUpperDeviceOk returns a tuple with the UpperDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperDevice

`func (o *SelectStackMemberPortForVlanVO) SetUpperDevice(v DeviceBriefVO)`

SetUpperDevice sets UpperDevice field to given value.

### HasUpperDevice

`func (o *SelectStackMemberPortForVlanVO) HasUpperDevice() bool`

HasUpperDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


