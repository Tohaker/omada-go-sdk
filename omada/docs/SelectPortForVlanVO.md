# SelectPortForVlanVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Whether the port is affected. | [optional] 
**AutoSelect** | Pointer to **bool** | Whether the port needs to be automatically selected. | [optional] 
**ConfigMlagDad** | Pointer to **bool** | Whether the port is configured as a mlag dad port. | [optional] 
**ConfigMlagPeerLink** | Pointer to **bool** | Whether the port is configured as a mlag peer-link port. | [optional] 
**DefaultVlan** | Pointer to **int32** | The vlan of default network. | [optional] 
**DownlinkDevices** | Pointer to [**[]DeviceBriefVO**](DeviceBriefVO.md) | The downlink devices of the port | [optional] 
**EditEnable** | Pointer to **bool** | Whether the port is selectable. | [optional] 
**Id** | Pointer to **int32** | port number, for example: 1 | [optional] 
**LagId** | Pointer to **int32** | Lag ID. It indicates the lag id of the port when the port is in a lag. | [optional] 
**Mode** | Pointer to **int32** | Mode, 0:WAN,1:LAN; | [optional] 
**Name** | Pointer to **string** | Name. | [optional] 
**NativeIsDefault** | Pointer to **bool** | It indicates whether the native network of the port is default. | [optional] 
**NativeNetworkName** | Pointer to **string** | The native network name of the port. | [optional] 
**NativeNetworkVlan** | Pointer to **int32** | The native network vlan of the port. | [optional] 
**NeedConfirm** | Pointer to **bool** | Whether the port needs confirm for binding non-default vlan. | [optional] 
**NeedConfirmCascadePort** | Pointer to **bool** | Whether the port needs confirm for being cascade port. | [optional] 
**NeedConfirmVoiceNetwork** | Pointer to **bool** | When creating single vlan and the port&#39;s voice network is enabled, needConfirmVoiceNetwork is true. | [optional] 
**OswPortNetworkTagsSetting** | Pointer to **int32** | The port network tag setting. 0:allow all, 1:block all, 2:custom | [optional] 
**PortId** | Pointer to **string** | ID, for example: \&quot;1/0/1\&quot; (for switch port), \&quot;3_e3f176e1c95a49849477636c43a49d3f\&quot; (for gateway port). | [optional] 
**Reasons** | Pointer to **[]int32** | Only valid when editEnable is false. It indicates the reason why the port is not selectable.Each Item should be a value as follows: -2: The port has been added to LAG. Lag member port can not be selected; -8: The port is wan port. Wan port can not be selected; -10: The AP’s LAN port cannot be configured; -12: The port&#39;s tags setting is not custom and therefore it can not be selected when creating multi vlan; -13: The port is stack port. Stack port can not be selected; -14: The number of VLANs has reached the limit of the Easy Managed Switch. | [optional] 
**SelectByUser** | Pointer to **bool** | Whether the port is selected by user. | [optional] 
**Type** | Pointer to **int32** | Type. 1: Copper, 2: Combo, 3: SFP | [optional] 
**UnSelectByUser** | Pointer to **bool** | Whether the auto-select port is unselected by user. | [optional] 
**UpperDevice** | Pointer to [**DeviceBriefVO**](DeviceBriefVO.md) |  | [optional] 

## Methods

### NewSelectPortForVlanVO

`func NewSelectPortForVlanVO() *SelectPortForVlanVO`

NewSelectPortForVlanVO instantiates a new SelectPortForVlanVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectPortForVlanVOWithDefaults

`func NewSelectPortForVlanVOWithDefaults() *SelectPortForVlanVO`

NewSelectPortForVlanVOWithDefaults instantiates a new SelectPortForVlanVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *SelectPortForVlanVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *SelectPortForVlanVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *SelectPortForVlanVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *SelectPortForVlanVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetAutoSelect

`func (o *SelectPortForVlanVO) GetAutoSelect() bool`

GetAutoSelect returns the AutoSelect field if non-nil, zero value otherwise.

### GetAutoSelectOk

`func (o *SelectPortForVlanVO) GetAutoSelectOk() (*bool, bool)`

GetAutoSelectOk returns a tuple with the AutoSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSelect

`func (o *SelectPortForVlanVO) SetAutoSelect(v bool)`

SetAutoSelect sets AutoSelect field to given value.

### HasAutoSelect

`func (o *SelectPortForVlanVO) HasAutoSelect() bool`

HasAutoSelect returns a boolean if a field has been set.

### GetConfigMlagDad

`func (o *SelectPortForVlanVO) GetConfigMlagDad() bool`

GetConfigMlagDad returns the ConfigMlagDad field if non-nil, zero value otherwise.

### GetConfigMlagDadOk

`func (o *SelectPortForVlanVO) GetConfigMlagDadOk() (*bool, bool)`

GetConfigMlagDadOk returns a tuple with the ConfigMlagDad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMlagDad

`func (o *SelectPortForVlanVO) SetConfigMlagDad(v bool)`

SetConfigMlagDad sets ConfigMlagDad field to given value.

### HasConfigMlagDad

`func (o *SelectPortForVlanVO) HasConfigMlagDad() bool`

HasConfigMlagDad returns a boolean if a field has been set.

### GetConfigMlagPeerLink

`func (o *SelectPortForVlanVO) GetConfigMlagPeerLink() bool`

GetConfigMlagPeerLink returns the ConfigMlagPeerLink field if non-nil, zero value otherwise.

### GetConfigMlagPeerLinkOk

`func (o *SelectPortForVlanVO) GetConfigMlagPeerLinkOk() (*bool, bool)`

GetConfigMlagPeerLinkOk returns a tuple with the ConfigMlagPeerLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMlagPeerLink

`func (o *SelectPortForVlanVO) SetConfigMlagPeerLink(v bool)`

SetConfigMlagPeerLink sets ConfigMlagPeerLink field to given value.

### HasConfigMlagPeerLink

`func (o *SelectPortForVlanVO) HasConfigMlagPeerLink() bool`

HasConfigMlagPeerLink returns a boolean if a field has been set.

### GetDefaultVlan

`func (o *SelectPortForVlanVO) GetDefaultVlan() int32`

GetDefaultVlan returns the DefaultVlan field if non-nil, zero value otherwise.

### GetDefaultVlanOk

`func (o *SelectPortForVlanVO) GetDefaultVlanOk() (*int32, bool)`

GetDefaultVlanOk returns a tuple with the DefaultVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlan

`func (o *SelectPortForVlanVO) SetDefaultVlan(v int32)`

SetDefaultVlan sets DefaultVlan field to given value.

### HasDefaultVlan

`func (o *SelectPortForVlanVO) HasDefaultVlan() bool`

HasDefaultVlan returns a boolean if a field has been set.

### GetDownlinkDevices

`func (o *SelectPortForVlanVO) GetDownlinkDevices() []DeviceBriefVO`

GetDownlinkDevices returns the DownlinkDevices field if non-nil, zero value otherwise.

### GetDownlinkDevicesOk

`func (o *SelectPortForVlanVO) GetDownlinkDevicesOk() (*[]DeviceBriefVO, bool)`

GetDownlinkDevicesOk returns a tuple with the DownlinkDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkDevices

`func (o *SelectPortForVlanVO) SetDownlinkDevices(v []DeviceBriefVO)`

SetDownlinkDevices sets DownlinkDevices field to given value.

### HasDownlinkDevices

`func (o *SelectPortForVlanVO) HasDownlinkDevices() bool`

HasDownlinkDevices returns a boolean if a field has been set.

### GetEditEnable

`func (o *SelectPortForVlanVO) GetEditEnable() bool`

GetEditEnable returns the EditEnable field if non-nil, zero value otherwise.

### GetEditEnableOk

`func (o *SelectPortForVlanVO) GetEditEnableOk() (*bool, bool)`

GetEditEnableOk returns a tuple with the EditEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditEnable

`func (o *SelectPortForVlanVO) SetEditEnable(v bool)`

SetEditEnable sets EditEnable field to given value.

### HasEditEnable

`func (o *SelectPortForVlanVO) HasEditEnable() bool`

HasEditEnable returns a boolean if a field has been set.

### GetId

`func (o *SelectPortForVlanVO) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SelectPortForVlanVO) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SelectPortForVlanVO) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SelectPortForVlanVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLagId

`func (o *SelectPortForVlanVO) GetLagId() int32`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *SelectPortForVlanVO) GetLagIdOk() (*int32, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *SelectPortForVlanVO) SetLagId(v int32)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *SelectPortForVlanVO) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetMode

`func (o *SelectPortForVlanVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SelectPortForVlanVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SelectPortForVlanVO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SelectPortForVlanVO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *SelectPortForVlanVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SelectPortForVlanVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SelectPortForVlanVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SelectPortForVlanVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNativeIsDefault

`func (o *SelectPortForVlanVO) GetNativeIsDefault() bool`

GetNativeIsDefault returns the NativeIsDefault field if non-nil, zero value otherwise.

### GetNativeIsDefaultOk

`func (o *SelectPortForVlanVO) GetNativeIsDefaultOk() (*bool, bool)`

GetNativeIsDefaultOk returns a tuple with the NativeIsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeIsDefault

`func (o *SelectPortForVlanVO) SetNativeIsDefault(v bool)`

SetNativeIsDefault sets NativeIsDefault field to given value.

### HasNativeIsDefault

`func (o *SelectPortForVlanVO) HasNativeIsDefault() bool`

HasNativeIsDefault returns a boolean if a field has been set.

### GetNativeNetworkName

`func (o *SelectPortForVlanVO) GetNativeNetworkName() string`

GetNativeNetworkName returns the NativeNetworkName field if non-nil, zero value otherwise.

### GetNativeNetworkNameOk

`func (o *SelectPortForVlanVO) GetNativeNetworkNameOk() (*string, bool)`

GetNativeNetworkNameOk returns a tuple with the NativeNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeNetworkName

`func (o *SelectPortForVlanVO) SetNativeNetworkName(v string)`

SetNativeNetworkName sets NativeNetworkName field to given value.

### HasNativeNetworkName

`func (o *SelectPortForVlanVO) HasNativeNetworkName() bool`

HasNativeNetworkName returns a boolean if a field has been set.

### GetNativeNetworkVlan

`func (o *SelectPortForVlanVO) GetNativeNetworkVlan() int32`

GetNativeNetworkVlan returns the NativeNetworkVlan field if non-nil, zero value otherwise.

### GetNativeNetworkVlanOk

`func (o *SelectPortForVlanVO) GetNativeNetworkVlanOk() (*int32, bool)`

GetNativeNetworkVlanOk returns a tuple with the NativeNetworkVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeNetworkVlan

`func (o *SelectPortForVlanVO) SetNativeNetworkVlan(v int32)`

SetNativeNetworkVlan sets NativeNetworkVlan field to given value.

### HasNativeNetworkVlan

`func (o *SelectPortForVlanVO) HasNativeNetworkVlan() bool`

HasNativeNetworkVlan returns a boolean if a field has been set.

### GetNeedConfirm

`func (o *SelectPortForVlanVO) GetNeedConfirm() bool`

GetNeedConfirm returns the NeedConfirm field if non-nil, zero value otherwise.

### GetNeedConfirmOk

`func (o *SelectPortForVlanVO) GetNeedConfirmOk() (*bool, bool)`

GetNeedConfirmOk returns a tuple with the NeedConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedConfirm

`func (o *SelectPortForVlanVO) SetNeedConfirm(v bool)`

SetNeedConfirm sets NeedConfirm field to given value.

### HasNeedConfirm

`func (o *SelectPortForVlanVO) HasNeedConfirm() bool`

HasNeedConfirm returns a boolean if a field has been set.

### GetNeedConfirmCascadePort

`func (o *SelectPortForVlanVO) GetNeedConfirmCascadePort() bool`

GetNeedConfirmCascadePort returns the NeedConfirmCascadePort field if non-nil, zero value otherwise.

### GetNeedConfirmCascadePortOk

`func (o *SelectPortForVlanVO) GetNeedConfirmCascadePortOk() (*bool, bool)`

GetNeedConfirmCascadePortOk returns a tuple with the NeedConfirmCascadePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedConfirmCascadePort

`func (o *SelectPortForVlanVO) SetNeedConfirmCascadePort(v bool)`

SetNeedConfirmCascadePort sets NeedConfirmCascadePort field to given value.

### HasNeedConfirmCascadePort

`func (o *SelectPortForVlanVO) HasNeedConfirmCascadePort() bool`

HasNeedConfirmCascadePort returns a boolean if a field has been set.

### GetNeedConfirmVoiceNetwork

`func (o *SelectPortForVlanVO) GetNeedConfirmVoiceNetwork() bool`

GetNeedConfirmVoiceNetwork returns the NeedConfirmVoiceNetwork field if non-nil, zero value otherwise.

### GetNeedConfirmVoiceNetworkOk

`func (o *SelectPortForVlanVO) GetNeedConfirmVoiceNetworkOk() (*bool, bool)`

GetNeedConfirmVoiceNetworkOk returns a tuple with the NeedConfirmVoiceNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedConfirmVoiceNetwork

`func (o *SelectPortForVlanVO) SetNeedConfirmVoiceNetwork(v bool)`

SetNeedConfirmVoiceNetwork sets NeedConfirmVoiceNetwork field to given value.

### HasNeedConfirmVoiceNetwork

`func (o *SelectPortForVlanVO) HasNeedConfirmVoiceNetwork() bool`

HasNeedConfirmVoiceNetwork returns a boolean if a field has been set.

### GetOswPortNetworkTagsSetting

`func (o *SelectPortForVlanVO) GetOswPortNetworkTagsSetting() int32`

GetOswPortNetworkTagsSetting returns the OswPortNetworkTagsSetting field if non-nil, zero value otherwise.

### GetOswPortNetworkTagsSettingOk

`func (o *SelectPortForVlanVO) GetOswPortNetworkTagsSettingOk() (*int32, bool)`

GetOswPortNetworkTagsSettingOk returns a tuple with the OswPortNetworkTagsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOswPortNetworkTagsSetting

`func (o *SelectPortForVlanVO) SetOswPortNetworkTagsSetting(v int32)`

SetOswPortNetworkTagsSetting sets OswPortNetworkTagsSetting field to given value.

### HasOswPortNetworkTagsSetting

`func (o *SelectPortForVlanVO) HasOswPortNetworkTagsSetting() bool`

HasOswPortNetworkTagsSetting returns a boolean if a field has been set.

### GetPortId

`func (o *SelectPortForVlanVO) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *SelectPortForVlanVO) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *SelectPortForVlanVO) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *SelectPortForVlanVO) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetReasons

`func (o *SelectPortForVlanVO) GetReasons() []int32`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *SelectPortForVlanVO) GetReasonsOk() (*[]int32, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *SelectPortForVlanVO) SetReasons(v []int32)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *SelectPortForVlanVO) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetSelectByUser

`func (o *SelectPortForVlanVO) GetSelectByUser() bool`

GetSelectByUser returns the SelectByUser field if non-nil, zero value otherwise.

### GetSelectByUserOk

`func (o *SelectPortForVlanVO) GetSelectByUserOk() (*bool, bool)`

GetSelectByUserOk returns a tuple with the SelectByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectByUser

`func (o *SelectPortForVlanVO) SetSelectByUser(v bool)`

SetSelectByUser sets SelectByUser field to given value.

### HasSelectByUser

`func (o *SelectPortForVlanVO) HasSelectByUser() bool`

HasSelectByUser returns a boolean if a field has been set.

### GetType

`func (o *SelectPortForVlanVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SelectPortForVlanVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SelectPortForVlanVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SelectPortForVlanVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnSelectByUser

`func (o *SelectPortForVlanVO) GetUnSelectByUser() bool`

GetUnSelectByUser returns the UnSelectByUser field if non-nil, zero value otherwise.

### GetUnSelectByUserOk

`func (o *SelectPortForVlanVO) GetUnSelectByUserOk() (*bool, bool)`

GetUnSelectByUserOk returns a tuple with the UnSelectByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnSelectByUser

`func (o *SelectPortForVlanVO) SetUnSelectByUser(v bool)`

SetUnSelectByUser sets UnSelectByUser field to given value.

### HasUnSelectByUser

`func (o *SelectPortForVlanVO) HasUnSelectByUser() bool`

HasUnSelectByUser returns a boolean if a field has been set.

### GetUpperDevice

`func (o *SelectPortForVlanVO) GetUpperDevice() DeviceBriefVO`

GetUpperDevice returns the UpperDevice field if non-nil, zero value otherwise.

### GetUpperDeviceOk

`func (o *SelectPortForVlanVO) GetUpperDeviceOk() (*DeviceBriefVO, bool)`

GetUpperDeviceOk returns a tuple with the UpperDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperDevice

`func (o *SelectPortForVlanVO) SetUpperDevice(v DeviceBriefVO)`

SetUpperDevice sets UpperDevice field to given value.

### HasUpperDevice

`func (o *SelectPortForVlanVO) HasUpperDevice() bool`

HasUpperDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


