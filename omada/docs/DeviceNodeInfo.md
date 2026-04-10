# DeviceNodeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | Pointer to **bool** |  | [optional] 
**ApInfo** | Pointer to [**APInfo**](APInfo.md) |  | [optional] 
**DevRxRate** | Pointer to **int64** | Device real-time downloadRate | [optional] 
**DevTxRate** | Pointer to **int64** | Device real-time uploadRate | [optional] 
**DeviceType** | Pointer to **int32** | Device type, 0: AP; 1: Switch; 2: Gateway. | [optional] 
**GatewayInfo** | Pointer to [**GatewayInfoEntity**](GatewayInfoEntity.md) |  | [optional] 
**HealthScore** | Pointer to **int32** | 1~3: poor; 4~7: fair; 0: no data; 8~10 good. | [optional] 
**Ip** | Pointer to **string** | Device IP | [optional] 
**Ipv6List** | Pointer to **[]string** | Device IPv6 list | [optional] 
**Mac** | Pointer to **string** | Device MAC | [optional] 
**Model** | Pointer to **string** | Device model name | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**Name** | Pointer to **string** | Device name | [optional] 
**Osg** | Pointer to **bool** |  | [optional] 
**Osw** | Pointer to **bool** |  | [optional] 
**ShowModel** | Pointer to **string** | Device model name with version | [optional] 
**StackGroup** | Pointer to **bool** | Stack Group | [optional] 
**StackId** | Pointer to **string** | Stack Id | [optional] 
**Status** | Pointer to **int32** | Status of device,status should be a value as follows: 0:Disconnected;1:Disconnected(Migrating);10:Provisioning;11:Configuring;12:Upgrading;13:Rebooting;14:Connected;15:Connected(Wireless);16:Connected(Migrating);17:Connected(Wireless,Migrating);20:Pending;21:Pending(Wireless);22:Adopting;23:Adopting(Wireless);24:Adopt Failed;25:Adopt Failed(Wireless);26:Managed By Others;27:Managed By Others(Wireless);30:Heartbeat Missed;31:Heartbeat Missed(Wireless);32:Heartbeat Missed(Migrating);33:Heartbeat Missed(Wireless,Migrating);40:Isolated;41:Isolated(Migrating);50:Slice Configuring | [optional] 
**StatusCategory** | Pointer to **int32** | Device status should be a value as follows: 0: Disconnected; 1: Connected; 2: Pending; 3: Heartbeat Missed; 4: Isolated | [optional] 
**SwitchInfo** | Pointer to [**SwitchInfo**](SwitchInfo.md) |  | [optional] 

## Methods

### NewDeviceNodeInfo

`func NewDeviceNodeInfo() *DeviceNodeInfo`

NewDeviceNodeInfo instantiates a new DeviceNodeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceNodeInfoWithDefaults

`func NewDeviceNodeInfoWithDefaults() *DeviceNodeInfo`

NewDeviceNodeInfoWithDefaults instantiates a new DeviceNodeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *DeviceNodeInfo) GetAp() bool`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *DeviceNodeInfo) GetApOk() (*bool, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *DeviceNodeInfo) SetAp(v bool)`

SetAp sets Ap field to given value.

### HasAp

`func (o *DeviceNodeInfo) HasAp() bool`

HasAp returns a boolean if a field has been set.

### GetApInfo

`func (o *DeviceNodeInfo) GetApInfo() APInfo`

GetApInfo returns the ApInfo field if non-nil, zero value otherwise.

### GetApInfoOk

`func (o *DeviceNodeInfo) GetApInfoOk() (*APInfo, bool)`

GetApInfoOk returns a tuple with the ApInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApInfo

`func (o *DeviceNodeInfo) SetApInfo(v APInfo)`

SetApInfo sets ApInfo field to given value.

### HasApInfo

`func (o *DeviceNodeInfo) HasApInfo() bool`

HasApInfo returns a boolean if a field has been set.

### GetDevRxRate

`func (o *DeviceNodeInfo) GetDevRxRate() int64`

GetDevRxRate returns the DevRxRate field if non-nil, zero value otherwise.

### GetDevRxRateOk

`func (o *DeviceNodeInfo) GetDevRxRateOk() (*int64, bool)`

GetDevRxRateOk returns a tuple with the DevRxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevRxRate

`func (o *DeviceNodeInfo) SetDevRxRate(v int64)`

SetDevRxRate sets DevRxRate field to given value.

### HasDevRxRate

`func (o *DeviceNodeInfo) HasDevRxRate() bool`

HasDevRxRate returns a boolean if a field has been set.

### GetDevTxRate

`func (o *DeviceNodeInfo) GetDevTxRate() int64`

GetDevTxRate returns the DevTxRate field if non-nil, zero value otherwise.

### GetDevTxRateOk

`func (o *DeviceNodeInfo) GetDevTxRateOk() (*int64, bool)`

GetDevTxRateOk returns a tuple with the DevTxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevTxRate

`func (o *DeviceNodeInfo) SetDevTxRate(v int64)`

SetDevTxRate sets DevTxRate field to given value.

### HasDevTxRate

`func (o *DeviceNodeInfo) HasDevTxRate() bool`

HasDevTxRate returns a boolean if a field has been set.

### GetDeviceType

`func (o *DeviceNodeInfo) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DeviceNodeInfo) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DeviceNodeInfo) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *DeviceNodeInfo) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetGatewayInfo

`func (o *DeviceNodeInfo) GetGatewayInfo() GatewayInfoEntity`

GetGatewayInfo returns the GatewayInfo field if non-nil, zero value otherwise.

### GetGatewayInfoOk

`func (o *DeviceNodeInfo) GetGatewayInfoOk() (*GatewayInfoEntity, bool)`

GetGatewayInfoOk returns a tuple with the GatewayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayInfo

`func (o *DeviceNodeInfo) SetGatewayInfo(v GatewayInfoEntity)`

SetGatewayInfo sets GatewayInfo field to given value.

### HasGatewayInfo

`func (o *DeviceNodeInfo) HasGatewayInfo() bool`

HasGatewayInfo returns a boolean if a field has been set.

### GetHealthScore

`func (o *DeviceNodeInfo) GetHealthScore() int32`

GetHealthScore returns the HealthScore field if non-nil, zero value otherwise.

### GetHealthScoreOk

`func (o *DeviceNodeInfo) GetHealthScoreOk() (*int32, bool)`

GetHealthScoreOk returns a tuple with the HealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthScore

`func (o *DeviceNodeInfo) SetHealthScore(v int32)`

SetHealthScore sets HealthScore field to given value.

### HasHealthScore

`func (o *DeviceNodeInfo) HasHealthScore() bool`

HasHealthScore returns a boolean if a field has been set.

### GetIp

`func (o *DeviceNodeInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DeviceNodeInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DeviceNodeInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *DeviceNodeInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6List

`func (o *DeviceNodeInfo) GetIpv6List() []string`

GetIpv6List returns the Ipv6List field if non-nil, zero value otherwise.

### GetIpv6ListOk

`func (o *DeviceNodeInfo) GetIpv6ListOk() (*[]string, bool)`

GetIpv6ListOk returns a tuple with the Ipv6List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6List

`func (o *DeviceNodeInfo) SetIpv6List(v []string)`

SetIpv6List sets Ipv6List field to given value.

### HasIpv6List

`func (o *DeviceNodeInfo) HasIpv6List() bool`

HasIpv6List returns a boolean if a field has been set.

### GetMac

`func (o *DeviceNodeInfo) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DeviceNodeInfo) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DeviceNodeInfo) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *DeviceNodeInfo) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *DeviceNodeInfo) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceNodeInfo) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceNodeInfo) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DeviceNodeInfo) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *DeviceNodeInfo) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *DeviceNodeInfo) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *DeviceNodeInfo) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *DeviceNodeInfo) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *DeviceNodeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceNodeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceNodeInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceNodeInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsg

`func (o *DeviceNodeInfo) GetOsg() bool`

GetOsg returns the Osg field if non-nil, zero value otherwise.

### GetOsgOk

`func (o *DeviceNodeInfo) GetOsgOk() (*bool, bool)`

GetOsgOk returns a tuple with the Osg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsg

`func (o *DeviceNodeInfo) SetOsg(v bool)`

SetOsg sets Osg field to given value.

### HasOsg

`func (o *DeviceNodeInfo) HasOsg() bool`

HasOsg returns a boolean if a field has been set.

### GetOsw

`func (o *DeviceNodeInfo) GetOsw() bool`

GetOsw returns the Osw field if non-nil, zero value otherwise.

### GetOswOk

`func (o *DeviceNodeInfo) GetOswOk() (*bool, bool)`

GetOswOk returns a tuple with the Osw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsw

`func (o *DeviceNodeInfo) SetOsw(v bool)`

SetOsw sets Osw field to given value.

### HasOsw

`func (o *DeviceNodeInfo) HasOsw() bool`

HasOsw returns a boolean if a field has been set.

### GetShowModel

`func (o *DeviceNodeInfo) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *DeviceNodeInfo) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *DeviceNodeInfo) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *DeviceNodeInfo) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetStackGroup

`func (o *DeviceNodeInfo) GetStackGroup() bool`

GetStackGroup returns the StackGroup field if non-nil, zero value otherwise.

### GetStackGroupOk

`func (o *DeviceNodeInfo) GetStackGroupOk() (*bool, bool)`

GetStackGroupOk returns a tuple with the StackGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackGroup

`func (o *DeviceNodeInfo) SetStackGroup(v bool)`

SetStackGroup sets StackGroup field to given value.

### HasStackGroup

`func (o *DeviceNodeInfo) HasStackGroup() bool`

HasStackGroup returns a boolean if a field has been set.

### GetStackId

`func (o *DeviceNodeInfo) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *DeviceNodeInfo) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *DeviceNodeInfo) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *DeviceNodeInfo) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceNodeInfo) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceNodeInfo) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceNodeInfo) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceNodeInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *DeviceNodeInfo) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *DeviceNodeInfo) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *DeviceNodeInfo) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *DeviceNodeInfo) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetSwitchInfo

`func (o *DeviceNodeInfo) GetSwitchInfo() SwitchInfo`

GetSwitchInfo returns the SwitchInfo field if non-nil, zero value otherwise.

### GetSwitchInfoOk

`func (o *DeviceNodeInfo) GetSwitchInfoOk() (*SwitchInfo, bool)`

GetSwitchInfoOk returns a tuple with the SwitchInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchInfo

`func (o *DeviceNodeInfo) SetSwitchInfo(v SwitchInfo)`

SetSwitchInfo sets SwitchInfo field to given value.

### HasSwitchInfo

`func (o *DeviceNodeInfo) HasSwitchInfo() bool`

HasSwitchInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


