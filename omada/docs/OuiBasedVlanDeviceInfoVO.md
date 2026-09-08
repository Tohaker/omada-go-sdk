# OuiBasedVlanDeviceInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Specifies whether the device is currently in an active state. | [optional] 
**AddedInAdvanced** | Pointer to **bool** | Specifies whether the device was added offline in advance. A value of true indicates that the device was pre-added offline and has not yet been adopted. | [optional] 
**CustomLagIds** | Pointer to **[]int32** | The lag ids that has some vlan not included in port vlan. | [optional] 
**CustomPorts** | Pointer to **[]int32** | The ports that has some vlan not included in port vlan. | [optional] 
**DeviceMisc** | Pointer to [**OswDeviceMiscVO**](OswDeviceMiscVO.md) |  | [optional] 
**DeviceSeriesType** | Pointer to **int32** | DeviceSeriesType should be a value as follows: 0:advanced; 1:pro | [optional] 
**Es** | Pointer to **bool** | Whether it is Agile Series Switch | [optional] 
**Ip** | Pointer to **string** | IP Address | [optional] 
**LagList** | Pointer to **[]int32** | The lag list of the switch device. | [optional] 
**Lags** | Pointer to [**[]LagInfoVO**](LagInfoVO.md) | Device Lag Infos | [optional] 
**LanList** | Pointer to **[]string** | Lan list. | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**MlagMsg** | Pointer to [**MlagMsgVO**](MlagMsgVO.md) |  | [optional] 
**Model** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**Name** | Pointer to **string** | Device name,default value is the mac address of device | [optional] 
**OldFirmwareDevice** | Pointer to **bool** | Whether the device is an old firmware device. | [optional] 
**OldFirmwareUsed** | Pointer to **string** | The old firmware used by the device. | [optional] 
**OuiBasedVlanVersion** | Pointer to **int32** | Specifies the OUI-based VLAN version: 0 indicates the feature is unsupported, while 1, 2, and 3 represent versions v1, v2, and v3 respectively. | [optional] 
**Ports** | Pointer to [**[]OswPortVO**](OswPortVO.md) | The infos of ports. | [optional] 
**SsidIds** | Pointer to **[]string** | The ssid list ap device used. | [optional] 
**Stack** | Pointer to **bool** | Parameter [stack] indicates whether the device supports stacking. | [optional] 
**StackMsg** | Pointer to [**StackMsgVO**](StackMsgVO.md) |  | [optional] 
**StackPorts** | Pointer to [**[]OswStackPortGroupVO**](OswStackPortGroupVO.md) | Stack ports info. | [optional] 
**Status** | Pointer to **int32** | Device status. The value must be one of the following: 0: Disconnected; 1: Disconnected (Migrating); 10: Provisioning; 11: Configuring; 12: Upgrading; 13: Rebooting; 14: Connected; 15: Connected (Wireless); 16: Connected (Migrating); 17: Connected (Wireless, Migrating); 20: Pending; 21: Pending (Wireless); 22: Adopting; 23: Adopting (Wireless); 24: Adoption Failed; 25: Adoption Failed (Wireless); 26: Managed by Others; 27: Managed by Others (Wireless); 30: Heartbeat Missed; 31: Heartbeat Missed (Wireless); 32: Heartbeat Missed (Migrating); 33: Heartbeat Missed (Wireless, Migrating); 40: Isolated; 41: Isolated (Migrating); 50: Slice Configuring | [optional] 
**StatusCategory** | Pointer to **int32** | Device Status Category | [optional] 
**SupportLayout** | Pointer to **bool** | Whether the device supports reporting port layout information. | [optional] 
**Type** | Pointer to **string** | Device type:ap、gateway、switch、olt | [optional] 
**Unit** | Pointer to **int32** | Unit ID. | [optional] 
**Version** | Pointer to **string** | The version. | [optional] 

## Methods

### NewOuiBasedVlanDeviceInfoVO

`func NewOuiBasedVlanDeviceInfoVO() *OuiBasedVlanDeviceInfoVO`

NewOuiBasedVlanDeviceInfoVO instantiates a new OuiBasedVlanDeviceInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOuiBasedVlanDeviceInfoVOWithDefaults

`func NewOuiBasedVlanDeviceInfoVOWithDefaults() *OuiBasedVlanDeviceInfoVO`

NewOuiBasedVlanDeviceInfoVOWithDefaults instantiates a new OuiBasedVlanDeviceInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *OuiBasedVlanDeviceInfoVO) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OuiBasedVlanDeviceInfoVO) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OuiBasedVlanDeviceInfoVO) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *OuiBasedVlanDeviceInfoVO) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddedInAdvanced

`func (o *OuiBasedVlanDeviceInfoVO) GetAddedInAdvanced() bool`

GetAddedInAdvanced returns the AddedInAdvanced field if non-nil, zero value otherwise.

### GetAddedInAdvancedOk

`func (o *OuiBasedVlanDeviceInfoVO) GetAddedInAdvancedOk() (*bool, bool)`

GetAddedInAdvancedOk returns a tuple with the AddedInAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedInAdvanced

`func (o *OuiBasedVlanDeviceInfoVO) SetAddedInAdvanced(v bool)`

SetAddedInAdvanced sets AddedInAdvanced field to given value.

### HasAddedInAdvanced

`func (o *OuiBasedVlanDeviceInfoVO) HasAddedInAdvanced() bool`

HasAddedInAdvanced returns a boolean if a field has been set.

### GetCustomLagIds

`func (o *OuiBasedVlanDeviceInfoVO) GetCustomLagIds() []int32`

GetCustomLagIds returns the CustomLagIds field if non-nil, zero value otherwise.

### GetCustomLagIdsOk

`func (o *OuiBasedVlanDeviceInfoVO) GetCustomLagIdsOk() (*[]int32, bool)`

GetCustomLagIdsOk returns a tuple with the CustomLagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLagIds

`func (o *OuiBasedVlanDeviceInfoVO) SetCustomLagIds(v []int32)`

SetCustomLagIds sets CustomLagIds field to given value.

### HasCustomLagIds

`func (o *OuiBasedVlanDeviceInfoVO) HasCustomLagIds() bool`

HasCustomLagIds returns a boolean if a field has been set.

### GetCustomPorts

`func (o *OuiBasedVlanDeviceInfoVO) GetCustomPorts() []int32`

GetCustomPorts returns the CustomPorts field if non-nil, zero value otherwise.

### GetCustomPortsOk

`func (o *OuiBasedVlanDeviceInfoVO) GetCustomPortsOk() (*[]int32, bool)`

GetCustomPortsOk returns a tuple with the CustomPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPorts

`func (o *OuiBasedVlanDeviceInfoVO) SetCustomPorts(v []int32)`

SetCustomPorts sets CustomPorts field to given value.

### HasCustomPorts

`func (o *OuiBasedVlanDeviceInfoVO) HasCustomPorts() bool`

HasCustomPorts returns a boolean if a field has been set.

### GetDeviceMisc

`func (o *OuiBasedVlanDeviceInfoVO) GetDeviceMisc() OswDeviceMiscVO`

GetDeviceMisc returns the DeviceMisc field if non-nil, zero value otherwise.

### GetDeviceMiscOk

`func (o *OuiBasedVlanDeviceInfoVO) GetDeviceMiscOk() (*OswDeviceMiscVO, bool)`

GetDeviceMiscOk returns a tuple with the DeviceMisc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMisc

`func (o *OuiBasedVlanDeviceInfoVO) SetDeviceMisc(v OswDeviceMiscVO)`

SetDeviceMisc sets DeviceMisc field to given value.

### HasDeviceMisc

`func (o *OuiBasedVlanDeviceInfoVO) HasDeviceMisc() bool`

HasDeviceMisc returns a boolean if a field has been set.

### GetDeviceSeriesType

`func (o *OuiBasedVlanDeviceInfoVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *OuiBasedVlanDeviceInfoVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *OuiBasedVlanDeviceInfoVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *OuiBasedVlanDeviceInfoVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetEs

`func (o *OuiBasedVlanDeviceInfoVO) GetEs() bool`

GetEs returns the Es field if non-nil, zero value otherwise.

### GetEsOk

`func (o *OuiBasedVlanDeviceInfoVO) GetEsOk() (*bool, bool)`

GetEsOk returns a tuple with the Es field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEs

`func (o *OuiBasedVlanDeviceInfoVO) SetEs(v bool)`

SetEs sets Es field to given value.

### HasEs

`func (o *OuiBasedVlanDeviceInfoVO) HasEs() bool`

HasEs returns a boolean if a field has been set.

### GetIp

`func (o *OuiBasedVlanDeviceInfoVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OuiBasedVlanDeviceInfoVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OuiBasedVlanDeviceInfoVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OuiBasedVlanDeviceInfoVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLagList

`func (o *OuiBasedVlanDeviceInfoVO) GetLagList() []int32`

GetLagList returns the LagList field if non-nil, zero value otherwise.

### GetLagListOk

`func (o *OuiBasedVlanDeviceInfoVO) GetLagListOk() (*[]int32, bool)`

GetLagListOk returns a tuple with the LagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagList

`func (o *OuiBasedVlanDeviceInfoVO) SetLagList(v []int32)`

SetLagList sets LagList field to given value.

### HasLagList

`func (o *OuiBasedVlanDeviceInfoVO) HasLagList() bool`

HasLagList returns a boolean if a field has been set.

### GetLags

`func (o *OuiBasedVlanDeviceInfoVO) GetLags() []LagInfoVO`

GetLags returns the Lags field if non-nil, zero value otherwise.

### GetLagsOk

`func (o *OuiBasedVlanDeviceInfoVO) GetLagsOk() (*[]LagInfoVO, bool)`

GetLagsOk returns a tuple with the Lags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLags

`func (o *OuiBasedVlanDeviceInfoVO) SetLags(v []LagInfoVO)`

SetLags sets Lags field to given value.

### HasLags

`func (o *OuiBasedVlanDeviceInfoVO) HasLags() bool`

HasLags returns a boolean if a field has been set.

### GetLanList

`func (o *OuiBasedVlanDeviceInfoVO) GetLanList() []string`

GetLanList returns the LanList field if non-nil, zero value otherwise.

### GetLanListOk

`func (o *OuiBasedVlanDeviceInfoVO) GetLanListOk() (*[]string, bool)`

GetLanListOk returns a tuple with the LanList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanList

`func (o *OuiBasedVlanDeviceInfoVO) SetLanList(v []string)`

SetLanList sets LanList field to given value.

### HasLanList

`func (o *OuiBasedVlanDeviceInfoVO) HasLanList() bool`

HasLanList returns a boolean if a field has been set.

### GetMac

`func (o *OuiBasedVlanDeviceInfoVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OuiBasedVlanDeviceInfoVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OuiBasedVlanDeviceInfoVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OuiBasedVlanDeviceInfoVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMlagMsg

`func (o *OuiBasedVlanDeviceInfoVO) GetMlagMsg() MlagMsgVO`

GetMlagMsg returns the MlagMsg field if non-nil, zero value otherwise.

### GetMlagMsgOk

`func (o *OuiBasedVlanDeviceInfoVO) GetMlagMsgOk() (*MlagMsgVO, bool)`

GetMlagMsgOk returns a tuple with the MlagMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagMsg

`func (o *OuiBasedVlanDeviceInfoVO) SetMlagMsg(v MlagMsgVO)`

SetMlagMsg sets MlagMsg field to given value.

### HasMlagMsg

`func (o *OuiBasedVlanDeviceInfoVO) HasMlagMsg() bool`

HasMlagMsg returns a boolean if a field has been set.

### GetModel

`func (o *OuiBasedVlanDeviceInfoVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OuiBasedVlanDeviceInfoVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OuiBasedVlanDeviceInfoVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OuiBasedVlanDeviceInfoVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OuiBasedVlanDeviceInfoVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OuiBasedVlanDeviceInfoVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OuiBasedVlanDeviceInfoVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OuiBasedVlanDeviceInfoVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *OuiBasedVlanDeviceInfoVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OuiBasedVlanDeviceInfoVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OuiBasedVlanDeviceInfoVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OuiBasedVlanDeviceInfoVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOldFirmwareDevice

`func (o *OuiBasedVlanDeviceInfoVO) GetOldFirmwareDevice() bool`

GetOldFirmwareDevice returns the OldFirmwareDevice field if non-nil, zero value otherwise.

### GetOldFirmwareDeviceOk

`func (o *OuiBasedVlanDeviceInfoVO) GetOldFirmwareDeviceOk() (*bool, bool)`

GetOldFirmwareDeviceOk returns a tuple with the OldFirmwareDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldFirmwareDevice

`func (o *OuiBasedVlanDeviceInfoVO) SetOldFirmwareDevice(v bool)`

SetOldFirmwareDevice sets OldFirmwareDevice field to given value.

### HasOldFirmwareDevice

`func (o *OuiBasedVlanDeviceInfoVO) HasOldFirmwareDevice() bool`

HasOldFirmwareDevice returns a boolean if a field has been set.

### GetOldFirmwareUsed

`func (o *OuiBasedVlanDeviceInfoVO) GetOldFirmwareUsed() string`

GetOldFirmwareUsed returns the OldFirmwareUsed field if non-nil, zero value otherwise.

### GetOldFirmwareUsedOk

`func (o *OuiBasedVlanDeviceInfoVO) GetOldFirmwareUsedOk() (*string, bool)`

GetOldFirmwareUsedOk returns a tuple with the OldFirmwareUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldFirmwareUsed

`func (o *OuiBasedVlanDeviceInfoVO) SetOldFirmwareUsed(v string)`

SetOldFirmwareUsed sets OldFirmwareUsed field to given value.

### HasOldFirmwareUsed

`func (o *OuiBasedVlanDeviceInfoVO) HasOldFirmwareUsed() bool`

HasOldFirmwareUsed returns a boolean if a field has been set.

### GetOuiBasedVlanVersion

`func (o *OuiBasedVlanDeviceInfoVO) GetOuiBasedVlanVersion() int32`

GetOuiBasedVlanVersion returns the OuiBasedVlanVersion field if non-nil, zero value otherwise.

### GetOuiBasedVlanVersionOk

`func (o *OuiBasedVlanDeviceInfoVO) GetOuiBasedVlanVersionOk() (*int32, bool)`

GetOuiBasedVlanVersionOk returns a tuple with the OuiBasedVlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuiBasedVlanVersion

`func (o *OuiBasedVlanDeviceInfoVO) SetOuiBasedVlanVersion(v int32)`

SetOuiBasedVlanVersion sets OuiBasedVlanVersion field to given value.

### HasOuiBasedVlanVersion

`func (o *OuiBasedVlanDeviceInfoVO) HasOuiBasedVlanVersion() bool`

HasOuiBasedVlanVersion returns a boolean if a field has been set.

### GetPorts

`func (o *OuiBasedVlanDeviceInfoVO) GetPorts() []OswPortVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OuiBasedVlanDeviceInfoVO) GetPortsOk() (*[]OswPortVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OuiBasedVlanDeviceInfoVO) SetPorts(v []OswPortVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OuiBasedVlanDeviceInfoVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetSsidIds

`func (o *OuiBasedVlanDeviceInfoVO) GetSsidIds() []string`

GetSsidIds returns the SsidIds field if non-nil, zero value otherwise.

### GetSsidIdsOk

`func (o *OuiBasedVlanDeviceInfoVO) GetSsidIdsOk() (*[]string, bool)`

GetSsidIdsOk returns a tuple with the SsidIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidIds

`func (o *OuiBasedVlanDeviceInfoVO) SetSsidIds(v []string)`

SetSsidIds sets SsidIds field to given value.

### HasSsidIds

`func (o *OuiBasedVlanDeviceInfoVO) HasSsidIds() bool`

HasSsidIds returns a boolean if a field has been set.

### GetStack

`func (o *OuiBasedVlanDeviceInfoVO) GetStack() bool`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *OuiBasedVlanDeviceInfoVO) GetStackOk() (*bool, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *OuiBasedVlanDeviceInfoVO) SetStack(v bool)`

SetStack sets Stack field to given value.

### HasStack

`func (o *OuiBasedVlanDeviceInfoVO) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetStackMsg

`func (o *OuiBasedVlanDeviceInfoVO) GetStackMsg() StackMsgVO`

GetStackMsg returns the StackMsg field if non-nil, zero value otherwise.

### GetStackMsgOk

`func (o *OuiBasedVlanDeviceInfoVO) GetStackMsgOk() (*StackMsgVO, bool)`

GetStackMsgOk returns a tuple with the StackMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackMsg

`func (o *OuiBasedVlanDeviceInfoVO) SetStackMsg(v StackMsgVO)`

SetStackMsg sets StackMsg field to given value.

### HasStackMsg

`func (o *OuiBasedVlanDeviceInfoVO) HasStackMsg() bool`

HasStackMsg returns a boolean if a field has been set.

### GetStackPorts

`func (o *OuiBasedVlanDeviceInfoVO) GetStackPorts() []OswStackPortGroupVO`

GetStackPorts returns the StackPorts field if non-nil, zero value otherwise.

### GetStackPortsOk

`func (o *OuiBasedVlanDeviceInfoVO) GetStackPortsOk() (*[]OswStackPortGroupVO, bool)`

GetStackPortsOk returns a tuple with the StackPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPorts

`func (o *OuiBasedVlanDeviceInfoVO) SetStackPorts(v []OswStackPortGroupVO)`

SetStackPorts sets StackPorts field to given value.

### HasStackPorts

`func (o *OuiBasedVlanDeviceInfoVO) HasStackPorts() bool`

HasStackPorts returns a boolean if a field has been set.

### GetStatus

`func (o *OuiBasedVlanDeviceInfoVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OuiBasedVlanDeviceInfoVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OuiBasedVlanDeviceInfoVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OuiBasedVlanDeviceInfoVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *OuiBasedVlanDeviceInfoVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *OuiBasedVlanDeviceInfoVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *OuiBasedVlanDeviceInfoVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *OuiBasedVlanDeviceInfoVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetSupportLayout

`func (o *OuiBasedVlanDeviceInfoVO) GetSupportLayout() bool`

GetSupportLayout returns the SupportLayout field if non-nil, zero value otherwise.

### GetSupportLayoutOk

`func (o *OuiBasedVlanDeviceInfoVO) GetSupportLayoutOk() (*bool, bool)`

GetSupportLayoutOk returns a tuple with the SupportLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLayout

`func (o *OuiBasedVlanDeviceInfoVO) SetSupportLayout(v bool)`

SetSupportLayout sets SupportLayout field to given value.

### HasSupportLayout

`func (o *OuiBasedVlanDeviceInfoVO) HasSupportLayout() bool`

HasSupportLayout returns a boolean if a field has been set.

### GetType

`func (o *OuiBasedVlanDeviceInfoVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OuiBasedVlanDeviceInfoVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OuiBasedVlanDeviceInfoVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OuiBasedVlanDeviceInfoVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *OuiBasedVlanDeviceInfoVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OuiBasedVlanDeviceInfoVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OuiBasedVlanDeviceInfoVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *OuiBasedVlanDeviceInfoVO) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetVersion

`func (o *OuiBasedVlanDeviceInfoVO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OuiBasedVlanDeviceInfoVO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OuiBasedVlanDeviceInfoVO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OuiBasedVlanDeviceInfoVO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


