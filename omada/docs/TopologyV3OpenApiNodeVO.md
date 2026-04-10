# TopologyV3OpenApiNodeVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbnormalReason** | Pointer to **int32** | Stack Abnormal Reason. 0: Not All Connected. 2: Not All Activated. | [optional] 
**Active** | Pointer to **bool** | Whether The Device Is License Active Or Not | [optional] 
**Capability** | Pointer to **string** | Other Device Capability | [optional] 
**Channel2g** | Pointer to **int32** | channel2g | [optional] 
**Channel5g** | Pointer to **int32** | channel5g | [optional] 
**Channel5g2** | Pointer to **int32** | channel5g2 | [optional] 
**Channel6g** | Pointer to **int32** | channel6g | [optional] 
**ChannelNum** | Pointer to **int32** | ChannelNum Of IPC/NVR | [optional] 
**ClientCount** | Pointer to **int32** | Directly Connected Client Count | [optional] 
**ClientHealth** | Pointer to [**ClientAggHealthDTO**](ClientAggHealthDTO.md) |  | [optional] 
**ClientType** | Pointer to **string** | Type of Third Party Device, if type is other | [optional] 
**ClientVlanId** | Pointer to **int32** | Vlan Id Of Client That Exists In Topology As Other Device | [optional] 
**Compatible** | Pointer to **int32** | The Compatibility Type Of Device Firmware And Controller | [optional] 
**ConnectedChannel** | Pointer to **int32** | Connected Channel Of IPC/NVR | [optional] 
**DadLink** | Pointer to [**[]WiredPortV3DTO**](WiredPortV3DTO.md) | DadLink for mlag node | [optional] 
**Description** | Pointer to **string** | Other Device Description | [optional] 
**DevRxRate** | Pointer to **int64** | Device RxRate | [optional] 
**DevTxRate** | Pointer to **int64** | Device TxRate | [optional] 
**DeviceSeriesType** | Pointer to **int32** | Device Series Type | [optional] 
**Disconnected** | Pointer to **bool** | Whether The Device Is Disconnected Or Not | [optional] 
**EcspFirstVersion** | Pointer to **int32** | Device Protocol First Version | [optional] 
**ExistStpLoop** | Pointer to **bool** | Exist Stp Loop Or Not | [optional] 
**FilterClientCount** | Pointer to [**TopologyFilterClientCountDTO**](TopologyFilterClientCountDTO.md) |  | [optional] 
**FirmwareVersion** | Pointer to **string** | FirmwareVersion Of IPC/NVR | [optional] 
**HealthScore** | Pointer to **int32** | Health Score | [optional] 
**Ip** | Pointer to **string** | Device Ip | [optional] 
**IpcCount** | Pointer to **int32** | Directly Connected IPC Client Count | [optional] 
**Ippt** | Pointer to **bool** | If Gateway Is Ippt Mode | [optional] 
**IsAllClients** | Pointer to **bool** | Whether The Client Group Is All Clients Or Just Other Clients | [optional] 
**LanPorts** | Pointer to [**[]LanPort**](LanPort.md) | Gateway Lan Port List | [optional] 
**Mac** | Pointer to **string** | Device Mac | [optional] 
**MacList** | Pointer to **[]string** | Device Macs | [optional] 
**MlagId** | Pointer to **string** | Mlag group id | [optional] 
**Model** | Pointer to **string** | Device Model | [optional] 
**ModelVersion** | Pointer to **string** | Device ModelVersion | [optional] 
**MultiSwitchDownlinkClient** | Pointer to [**[]ClientsQueryMacAndFilterType**](ClientsQueryMacAndFilterType.md) | Mlag Or Vrrp Group All Members Client Query Parameter | [optional] 
**MultiSwitchList** | Pointer to [**[]TopologyV3OpenApiNodeVO**](TopologyV3OpenApiNodeVO.md) | All members for multiSwitchNode including mlag and vrrp node | [optional] 
**MultiSwitchNum** | Pointer to **int32** | Serial number for mlag and vrrp node members, which is unique | [optional] 
**MultiSwitchRole** | Pointer to **int32** | Role for mlag and vrrp node members | [optional] 
**Name** | Pointer to **string** | Device Name | [optional] 
**OmadaDeviceType** | Pointer to **string** | Type of Omada Device, if type is other | [optional] 
**PeerLink** | Pointer to [**[]WiredPortV3DTO**](WiredPortV3DTO.md) | PeerLink for mlag node | [optional] 
**PortLabels** | Pointer to [**PortLabelDTO**](PortLabelDTO.md) |  | [optional] 
**RdMode2g** | Pointer to **string** | rdMode2g | [optional] 
**RdMode5g** | Pointer to **string** | rdMode5g | [optional] 
**RdMode5g2** | Pointer to **string** | rdMode5g2 | [optional] 
**RdMode6g** | Pointer to **string** | rdMode6g | [optional] 
**Role** | Pointer to **int32** | Identify Device Role In P2P Scenarios | [optional] 
**ShowModel** | Pointer to **string** | Device Show Model | [optional] 
**SpecificType** | Pointer to **int32** | SpecificType for multiSwitchNode, 0 means mlag node and 1 means vrrp node | [optional] 
**StackGroup** | Pointer to **bool** | Whether The Switch Device Is Stack Group Or Not | [optional] 
**StackId** | Pointer to **string** | StackGroup Switch Device Id | [optional] 
**StackStatus** | Pointer to **int32** | Status of the Stack Group. 0: normal; 1: abnormal; 2: stack not ready. | [optional] 
**Status** | Pointer to **int32** | Device Status | [optional] 
**StatusCategory** | Pointer to **int32** | Device Status Category | [optional] 
**Successors** | Pointer to [**[]TopologyV3OpenApiNodeVO**](TopologyV3OpenApiNodeVO.md) | Device Successors | [optional] 
**Support5g2** | Pointer to **bool** | Whether The Device Supports 5g2 Or Not | [optional] 
**Type** | Pointer to **string** | Device Type | [optional] 
**UplinkMacList** | Pointer to **[]string** | Mlag Or Vrrp Group All Members Mac List | [optional] 
**VigiManaged** | Pointer to **bool** | If Vigi Is Managed By Vms | [optional] 
**VigiWirelessUpInfo** | Pointer to [**VigiWirelessUpInfoDTO**](VigiWirelessUpInfoDTO.md) |  | [optional] 
**VlansAsMaster** | Pointer to **string** | All vlans that vrrp member act as master | [optional] 
**VrrpGroupList** | Pointer to [**[]VrrpGroupDTO**](VrrpGroupDTO.md) | Vrrp node member&#39;s vrrp group information | [optional] 
**VrrpLink** | Pointer to [**[]VrrpLinkDTO**](VrrpLinkDTO.md) | Vrrp members internal link | [optional] 
**WanPorts** | Pointer to [**[]WanPort**](WanPort.md) | Gateway Wan Port List | [optional] 
**WarnNvrFirmware** | Pointer to **bool** | Whether IPC/NVR Needs To Upgrade Firmware To Support Lldp Feature | [optional] 
**WiredUpInfo** | Pointer to [**WiredUpInfoDTO**](WiredUpInfoDTO.md) |  | [optional] 
**WiredUpInfos** | Pointer to [**[]WiredUpInfoDTO**](WiredUpInfoDTO.md) | Multi wiredUpInfo for mlag and vrrp member or downlink device | [optional] 
**WirelessUpInfo** | Pointer to [**WirelessUpInfoDTO**](WirelessUpInfoDTO.md) |  | [optional] 
**WirelessUplink** | Pointer to **bool** | Whether The IPC/NVR Is Wireless Connected Or Not | [optional] 

## Methods

### NewTopologyV3OpenApiNodeVO

`func NewTopologyV3OpenApiNodeVO() *TopologyV3OpenApiNodeVO`

NewTopologyV3OpenApiNodeVO instantiates a new TopologyV3OpenApiNodeVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyV3OpenApiNodeVOWithDefaults

`func NewTopologyV3OpenApiNodeVOWithDefaults() *TopologyV3OpenApiNodeVO`

NewTopologyV3OpenApiNodeVOWithDefaults instantiates a new TopologyV3OpenApiNodeVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbnormalReason

`func (o *TopologyV3OpenApiNodeVO) GetAbnormalReason() int32`

GetAbnormalReason returns the AbnormalReason field if non-nil, zero value otherwise.

### GetAbnormalReasonOk

`func (o *TopologyV3OpenApiNodeVO) GetAbnormalReasonOk() (*int32, bool)`

GetAbnormalReasonOk returns a tuple with the AbnormalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbnormalReason

`func (o *TopologyV3OpenApiNodeVO) SetAbnormalReason(v int32)`

SetAbnormalReason sets AbnormalReason field to given value.

### HasAbnormalReason

`func (o *TopologyV3OpenApiNodeVO) HasAbnormalReason() bool`

HasAbnormalReason returns a boolean if a field has been set.

### GetActive

`func (o *TopologyV3OpenApiNodeVO) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *TopologyV3OpenApiNodeVO) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *TopologyV3OpenApiNodeVO) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *TopologyV3OpenApiNodeVO) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCapability

`func (o *TopologyV3OpenApiNodeVO) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *TopologyV3OpenApiNodeVO) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *TopologyV3OpenApiNodeVO) SetCapability(v string)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *TopologyV3OpenApiNodeVO) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### GetChannel2g

`func (o *TopologyV3OpenApiNodeVO) GetChannel2g() int32`

GetChannel2g returns the Channel2g field if non-nil, zero value otherwise.

### GetChannel2gOk

`func (o *TopologyV3OpenApiNodeVO) GetChannel2gOk() (*int32, bool)`

GetChannel2gOk returns a tuple with the Channel2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel2g

`func (o *TopologyV3OpenApiNodeVO) SetChannel2g(v int32)`

SetChannel2g sets Channel2g field to given value.

### HasChannel2g

`func (o *TopologyV3OpenApiNodeVO) HasChannel2g() bool`

HasChannel2g returns a boolean if a field has been set.

### GetChannel5g

`func (o *TopologyV3OpenApiNodeVO) GetChannel5g() int32`

GetChannel5g returns the Channel5g field if non-nil, zero value otherwise.

### GetChannel5gOk

`func (o *TopologyV3OpenApiNodeVO) GetChannel5gOk() (*int32, bool)`

GetChannel5gOk returns a tuple with the Channel5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel5g

`func (o *TopologyV3OpenApiNodeVO) SetChannel5g(v int32)`

SetChannel5g sets Channel5g field to given value.

### HasChannel5g

`func (o *TopologyV3OpenApiNodeVO) HasChannel5g() bool`

HasChannel5g returns a boolean if a field has been set.

### GetChannel5g2

`func (o *TopologyV3OpenApiNodeVO) GetChannel5g2() int32`

GetChannel5g2 returns the Channel5g2 field if non-nil, zero value otherwise.

### GetChannel5g2Ok

`func (o *TopologyV3OpenApiNodeVO) GetChannel5g2Ok() (*int32, bool)`

GetChannel5g2Ok returns a tuple with the Channel5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel5g2

`func (o *TopologyV3OpenApiNodeVO) SetChannel5g2(v int32)`

SetChannel5g2 sets Channel5g2 field to given value.

### HasChannel5g2

`func (o *TopologyV3OpenApiNodeVO) HasChannel5g2() bool`

HasChannel5g2 returns a boolean if a field has been set.

### GetChannel6g

`func (o *TopologyV3OpenApiNodeVO) GetChannel6g() int32`

GetChannel6g returns the Channel6g field if non-nil, zero value otherwise.

### GetChannel6gOk

`func (o *TopologyV3OpenApiNodeVO) GetChannel6gOk() (*int32, bool)`

GetChannel6gOk returns a tuple with the Channel6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel6g

`func (o *TopologyV3OpenApiNodeVO) SetChannel6g(v int32)`

SetChannel6g sets Channel6g field to given value.

### HasChannel6g

`func (o *TopologyV3OpenApiNodeVO) HasChannel6g() bool`

HasChannel6g returns a boolean if a field has been set.

### GetChannelNum

`func (o *TopologyV3OpenApiNodeVO) GetChannelNum() int32`

GetChannelNum returns the ChannelNum field if non-nil, zero value otherwise.

### GetChannelNumOk

`func (o *TopologyV3OpenApiNodeVO) GetChannelNumOk() (*int32, bool)`

GetChannelNumOk returns a tuple with the ChannelNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelNum

`func (o *TopologyV3OpenApiNodeVO) SetChannelNum(v int32)`

SetChannelNum sets ChannelNum field to given value.

### HasChannelNum

`func (o *TopologyV3OpenApiNodeVO) HasChannelNum() bool`

HasChannelNum returns a boolean if a field has been set.

### GetClientCount

`func (o *TopologyV3OpenApiNodeVO) GetClientCount() int32`

GetClientCount returns the ClientCount field if non-nil, zero value otherwise.

### GetClientCountOk

`func (o *TopologyV3OpenApiNodeVO) GetClientCountOk() (*int32, bool)`

GetClientCountOk returns a tuple with the ClientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCount

`func (o *TopologyV3OpenApiNodeVO) SetClientCount(v int32)`

SetClientCount sets ClientCount field to given value.

### HasClientCount

`func (o *TopologyV3OpenApiNodeVO) HasClientCount() bool`

HasClientCount returns a boolean if a field has been set.

### GetClientHealth

`func (o *TopologyV3OpenApiNodeVO) GetClientHealth() ClientAggHealthDTO`

GetClientHealth returns the ClientHealth field if non-nil, zero value otherwise.

### GetClientHealthOk

`func (o *TopologyV3OpenApiNodeVO) GetClientHealthOk() (*ClientAggHealthDTO, bool)`

GetClientHealthOk returns a tuple with the ClientHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHealth

`func (o *TopologyV3OpenApiNodeVO) SetClientHealth(v ClientAggHealthDTO)`

SetClientHealth sets ClientHealth field to given value.

### HasClientHealth

`func (o *TopologyV3OpenApiNodeVO) HasClientHealth() bool`

HasClientHealth returns a boolean if a field has been set.

### GetClientType

`func (o *TopologyV3OpenApiNodeVO) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *TopologyV3OpenApiNodeVO) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *TopologyV3OpenApiNodeVO) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *TopologyV3OpenApiNodeVO) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetClientVlanId

`func (o *TopologyV3OpenApiNodeVO) GetClientVlanId() int32`

GetClientVlanId returns the ClientVlanId field if non-nil, zero value otherwise.

### GetClientVlanIdOk

`func (o *TopologyV3OpenApiNodeVO) GetClientVlanIdOk() (*int32, bool)`

GetClientVlanIdOk returns a tuple with the ClientVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVlanId

`func (o *TopologyV3OpenApiNodeVO) SetClientVlanId(v int32)`

SetClientVlanId sets ClientVlanId field to given value.

### HasClientVlanId

`func (o *TopologyV3OpenApiNodeVO) HasClientVlanId() bool`

HasClientVlanId returns a boolean if a field has been set.

### GetCompatible

`func (o *TopologyV3OpenApiNodeVO) GetCompatible() int32`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *TopologyV3OpenApiNodeVO) GetCompatibleOk() (*int32, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *TopologyV3OpenApiNodeVO) SetCompatible(v int32)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *TopologyV3OpenApiNodeVO) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetConnectedChannel

`func (o *TopologyV3OpenApiNodeVO) GetConnectedChannel() int32`

GetConnectedChannel returns the ConnectedChannel field if non-nil, zero value otherwise.

### GetConnectedChannelOk

`func (o *TopologyV3OpenApiNodeVO) GetConnectedChannelOk() (*int32, bool)`

GetConnectedChannelOk returns a tuple with the ConnectedChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedChannel

`func (o *TopologyV3OpenApiNodeVO) SetConnectedChannel(v int32)`

SetConnectedChannel sets ConnectedChannel field to given value.

### HasConnectedChannel

`func (o *TopologyV3OpenApiNodeVO) HasConnectedChannel() bool`

HasConnectedChannel returns a boolean if a field has been set.

### GetDadLink

`func (o *TopologyV3OpenApiNodeVO) GetDadLink() []WiredPortV3DTO`

GetDadLink returns the DadLink field if non-nil, zero value otherwise.

### GetDadLinkOk

`func (o *TopologyV3OpenApiNodeVO) GetDadLinkOk() (*[]WiredPortV3DTO, bool)`

GetDadLinkOk returns a tuple with the DadLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDadLink

`func (o *TopologyV3OpenApiNodeVO) SetDadLink(v []WiredPortV3DTO)`

SetDadLink sets DadLink field to given value.

### HasDadLink

`func (o *TopologyV3OpenApiNodeVO) HasDadLink() bool`

HasDadLink returns a boolean if a field has been set.

### GetDescription

`func (o *TopologyV3OpenApiNodeVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TopologyV3OpenApiNodeVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TopologyV3OpenApiNodeVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TopologyV3OpenApiNodeVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevRxRate

`func (o *TopologyV3OpenApiNodeVO) GetDevRxRate() int64`

GetDevRxRate returns the DevRxRate field if non-nil, zero value otherwise.

### GetDevRxRateOk

`func (o *TopologyV3OpenApiNodeVO) GetDevRxRateOk() (*int64, bool)`

GetDevRxRateOk returns a tuple with the DevRxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevRxRate

`func (o *TopologyV3OpenApiNodeVO) SetDevRxRate(v int64)`

SetDevRxRate sets DevRxRate field to given value.

### HasDevRxRate

`func (o *TopologyV3OpenApiNodeVO) HasDevRxRate() bool`

HasDevRxRate returns a boolean if a field has been set.

### GetDevTxRate

`func (o *TopologyV3OpenApiNodeVO) GetDevTxRate() int64`

GetDevTxRate returns the DevTxRate field if non-nil, zero value otherwise.

### GetDevTxRateOk

`func (o *TopologyV3OpenApiNodeVO) GetDevTxRateOk() (*int64, bool)`

GetDevTxRateOk returns a tuple with the DevTxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevTxRate

`func (o *TopologyV3OpenApiNodeVO) SetDevTxRate(v int64)`

SetDevTxRate sets DevTxRate field to given value.

### HasDevTxRate

`func (o *TopologyV3OpenApiNodeVO) HasDevTxRate() bool`

HasDevTxRate returns a boolean if a field has been set.

### GetDeviceSeriesType

`func (o *TopologyV3OpenApiNodeVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *TopologyV3OpenApiNodeVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *TopologyV3OpenApiNodeVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *TopologyV3OpenApiNodeVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetDisconnected

`func (o *TopologyV3OpenApiNodeVO) GetDisconnected() bool`

GetDisconnected returns the Disconnected field if non-nil, zero value otherwise.

### GetDisconnectedOk

`func (o *TopologyV3OpenApiNodeVO) GetDisconnectedOk() (*bool, bool)`

GetDisconnectedOk returns a tuple with the Disconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnected

`func (o *TopologyV3OpenApiNodeVO) SetDisconnected(v bool)`

SetDisconnected sets Disconnected field to given value.

### HasDisconnected

`func (o *TopologyV3OpenApiNodeVO) HasDisconnected() bool`

HasDisconnected returns a boolean if a field has been set.

### GetEcspFirstVersion

`func (o *TopologyV3OpenApiNodeVO) GetEcspFirstVersion() int32`

GetEcspFirstVersion returns the EcspFirstVersion field if non-nil, zero value otherwise.

### GetEcspFirstVersionOk

`func (o *TopologyV3OpenApiNodeVO) GetEcspFirstVersionOk() (*int32, bool)`

GetEcspFirstVersionOk returns a tuple with the EcspFirstVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcspFirstVersion

`func (o *TopologyV3OpenApiNodeVO) SetEcspFirstVersion(v int32)`

SetEcspFirstVersion sets EcspFirstVersion field to given value.

### HasEcspFirstVersion

`func (o *TopologyV3OpenApiNodeVO) HasEcspFirstVersion() bool`

HasEcspFirstVersion returns a boolean if a field has been set.

### GetExistStpLoop

`func (o *TopologyV3OpenApiNodeVO) GetExistStpLoop() bool`

GetExistStpLoop returns the ExistStpLoop field if non-nil, zero value otherwise.

### GetExistStpLoopOk

`func (o *TopologyV3OpenApiNodeVO) GetExistStpLoopOk() (*bool, bool)`

GetExistStpLoopOk returns a tuple with the ExistStpLoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistStpLoop

`func (o *TopologyV3OpenApiNodeVO) SetExistStpLoop(v bool)`

SetExistStpLoop sets ExistStpLoop field to given value.

### HasExistStpLoop

`func (o *TopologyV3OpenApiNodeVO) HasExistStpLoop() bool`

HasExistStpLoop returns a boolean if a field has been set.

### GetFilterClientCount

`func (o *TopologyV3OpenApiNodeVO) GetFilterClientCount() TopologyFilterClientCountDTO`

GetFilterClientCount returns the FilterClientCount field if non-nil, zero value otherwise.

### GetFilterClientCountOk

`func (o *TopologyV3OpenApiNodeVO) GetFilterClientCountOk() (*TopologyFilterClientCountDTO, bool)`

GetFilterClientCountOk returns a tuple with the FilterClientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterClientCount

`func (o *TopologyV3OpenApiNodeVO) SetFilterClientCount(v TopologyFilterClientCountDTO)`

SetFilterClientCount sets FilterClientCount field to given value.

### HasFilterClientCount

`func (o *TopologyV3OpenApiNodeVO) HasFilterClientCount() bool`

HasFilterClientCount returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *TopologyV3OpenApiNodeVO) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *TopologyV3OpenApiNodeVO) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *TopologyV3OpenApiNodeVO) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *TopologyV3OpenApiNodeVO) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetHealthScore

`func (o *TopologyV3OpenApiNodeVO) GetHealthScore() int32`

GetHealthScore returns the HealthScore field if non-nil, zero value otherwise.

### GetHealthScoreOk

`func (o *TopologyV3OpenApiNodeVO) GetHealthScoreOk() (*int32, bool)`

GetHealthScoreOk returns a tuple with the HealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthScore

`func (o *TopologyV3OpenApiNodeVO) SetHealthScore(v int32)`

SetHealthScore sets HealthScore field to given value.

### HasHealthScore

`func (o *TopologyV3OpenApiNodeVO) HasHealthScore() bool`

HasHealthScore returns a boolean if a field has been set.

### GetIp

`func (o *TopologyV3OpenApiNodeVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *TopologyV3OpenApiNodeVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *TopologyV3OpenApiNodeVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *TopologyV3OpenApiNodeVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpcCount

`func (o *TopologyV3OpenApiNodeVO) GetIpcCount() int32`

GetIpcCount returns the IpcCount field if non-nil, zero value otherwise.

### GetIpcCountOk

`func (o *TopologyV3OpenApiNodeVO) GetIpcCountOk() (*int32, bool)`

GetIpcCountOk returns a tuple with the IpcCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpcCount

`func (o *TopologyV3OpenApiNodeVO) SetIpcCount(v int32)`

SetIpcCount sets IpcCount field to given value.

### HasIpcCount

`func (o *TopologyV3OpenApiNodeVO) HasIpcCount() bool`

HasIpcCount returns a boolean if a field has been set.

### GetIppt

`func (o *TopologyV3OpenApiNodeVO) GetIppt() bool`

GetIppt returns the Ippt field if non-nil, zero value otherwise.

### GetIpptOk

`func (o *TopologyV3OpenApiNodeVO) GetIpptOk() (*bool, bool)`

GetIpptOk returns a tuple with the Ippt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIppt

`func (o *TopologyV3OpenApiNodeVO) SetIppt(v bool)`

SetIppt sets Ippt field to given value.

### HasIppt

`func (o *TopologyV3OpenApiNodeVO) HasIppt() bool`

HasIppt returns a boolean if a field has been set.

### GetIsAllClients

`func (o *TopologyV3OpenApiNodeVO) GetIsAllClients() bool`

GetIsAllClients returns the IsAllClients field if non-nil, zero value otherwise.

### GetIsAllClientsOk

`func (o *TopologyV3OpenApiNodeVO) GetIsAllClientsOk() (*bool, bool)`

GetIsAllClientsOk returns a tuple with the IsAllClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllClients

`func (o *TopologyV3OpenApiNodeVO) SetIsAllClients(v bool)`

SetIsAllClients sets IsAllClients field to given value.

### HasIsAllClients

`func (o *TopologyV3OpenApiNodeVO) HasIsAllClients() bool`

HasIsAllClients returns a boolean if a field has been set.

### GetLanPorts

`func (o *TopologyV3OpenApiNodeVO) GetLanPorts() []LanPort`

GetLanPorts returns the LanPorts field if non-nil, zero value otherwise.

### GetLanPortsOk

`func (o *TopologyV3OpenApiNodeVO) GetLanPortsOk() (*[]LanPort, bool)`

GetLanPortsOk returns a tuple with the LanPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanPorts

`func (o *TopologyV3OpenApiNodeVO) SetLanPorts(v []LanPort)`

SetLanPorts sets LanPorts field to given value.

### HasLanPorts

`func (o *TopologyV3OpenApiNodeVO) HasLanPorts() bool`

HasLanPorts returns a boolean if a field has been set.

### GetMac

`func (o *TopologyV3OpenApiNodeVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *TopologyV3OpenApiNodeVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *TopologyV3OpenApiNodeVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *TopologyV3OpenApiNodeVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMacList

`func (o *TopologyV3OpenApiNodeVO) GetMacList() []string`

GetMacList returns the MacList field if non-nil, zero value otherwise.

### GetMacListOk

`func (o *TopologyV3OpenApiNodeVO) GetMacListOk() (*[]string, bool)`

GetMacListOk returns a tuple with the MacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacList

`func (o *TopologyV3OpenApiNodeVO) SetMacList(v []string)`

SetMacList sets MacList field to given value.

### HasMacList

`func (o *TopologyV3OpenApiNodeVO) HasMacList() bool`

HasMacList returns a boolean if a field has been set.

### GetMlagId

`func (o *TopologyV3OpenApiNodeVO) GetMlagId() string`

GetMlagId returns the MlagId field if non-nil, zero value otherwise.

### GetMlagIdOk

`func (o *TopologyV3OpenApiNodeVO) GetMlagIdOk() (*string, bool)`

GetMlagIdOk returns a tuple with the MlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagId

`func (o *TopologyV3OpenApiNodeVO) SetMlagId(v string)`

SetMlagId sets MlagId field to given value.

### HasMlagId

`func (o *TopologyV3OpenApiNodeVO) HasMlagId() bool`

HasMlagId returns a boolean if a field has been set.

### GetModel

`func (o *TopologyV3OpenApiNodeVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TopologyV3OpenApiNodeVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TopologyV3OpenApiNodeVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TopologyV3OpenApiNodeVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *TopologyV3OpenApiNodeVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *TopologyV3OpenApiNodeVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *TopologyV3OpenApiNodeVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *TopologyV3OpenApiNodeVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetMultiSwitchDownlinkClient

`func (o *TopologyV3OpenApiNodeVO) GetMultiSwitchDownlinkClient() []ClientsQueryMacAndFilterType`

GetMultiSwitchDownlinkClient returns the MultiSwitchDownlinkClient field if non-nil, zero value otherwise.

### GetMultiSwitchDownlinkClientOk

`func (o *TopologyV3OpenApiNodeVO) GetMultiSwitchDownlinkClientOk() (*[]ClientsQueryMacAndFilterType, bool)`

GetMultiSwitchDownlinkClientOk returns a tuple with the MultiSwitchDownlinkClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSwitchDownlinkClient

`func (o *TopologyV3OpenApiNodeVO) SetMultiSwitchDownlinkClient(v []ClientsQueryMacAndFilterType)`

SetMultiSwitchDownlinkClient sets MultiSwitchDownlinkClient field to given value.

### HasMultiSwitchDownlinkClient

`func (o *TopologyV3OpenApiNodeVO) HasMultiSwitchDownlinkClient() bool`

HasMultiSwitchDownlinkClient returns a boolean if a field has been set.

### GetMultiSwitchList

`func (o *TopologyV3OpenApiNodeVO) GetMultiSwitchList() []TopologyV3OpenApiNodeVO`

GetMultiSwitchList returns the MultiSwitchList field if non-nil, zero value otherwise.

### GetMultiSwitchListOk

`func (o *TopologyV3OpenApiNodeVO) GetMultiSwitchListOk() (*[]TopologyV3OpenApiNodeVO, bool)`

GetMultiSwitchListOk returns a tuple with the MultiSwitchList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSwitchList

`func (o *TopologyV3OpenApiNodeVO) SetMultiSwitchList(v []TopologyV3OpenApiNodeVO)`

SetMultiSwitchList sets MultiSwitchList field to given value.

### HasMultiSwitchList

`func (o *TopologyV3OpenApiNodeVO) HasMultiSwitchList() bool`

HasMultiSwitchList returns a boolean if a field has been set.

### GetMultiSwitchNum

`func (o *TopologyV3OpenApiNodeVO) GetMultiSwitchNum() int32`

GetMultiSwitchNum returns the MultiSwitchNum field if non-nil, zero value otherwise.

### GetMultiSwitchNumOk

`func (o *TopologyV3OpenApiNodeVO) GetMultiSwitchNumOk() (*int32, bool)`

GetMultiSwitchNumOk returns a tuple with the MultiSwitchNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSwitchNum

`func (o *TopologyV3OpenApiNodeVO) SetMultiSwitchNum(v int32)`

SetMultiSwitchNum sets MultiSwitchNum field to given value.

### HasMultiSwitchNum

`func (o *TopologyV3OpenApiNodeVO) HasMultiSwitchNum() bool`

HasMultiSwitchNum returns a boolean if a field has been set.

### GetMultiSwitchRole

`func (o *TopologyV3OpenApiNodeVO) GetMultiSwitchRole() int32`

GetMultiSwitchRole returns the MultiSwitchRole field if non-nil, zero value otherwise.

### GetMultiSwitchRoleOk

`func (o *TopologyV3OpenApiNodeVO) GetMultiSwitchRoleOk() (*int32, bool)`

GetMultiSwitchRoleOk returns a tuple with the MultiSwitchRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSwitchRole

`func (o *TopologyV3OpenApiNodeVO) SetMultiSwitchRole(v int32)`

SetMultiSwitchRole sets MultiSwitchRole field to given value.

### HasMultiSwitchRole

`func (o *TopologyV3OpenApiNodeVO) HasMultiSwitchRole() bool`

HasMultiSwitchRole returns a boolean if a field has been set.

### GetName

`func (o *TopologyV3OpenApiNodeVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopologyV3OpenApiNodeVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopologyV3OpenApiNodeVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopologyV3OpenApiNodeVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOmadaDeviceType

`func (o *TopologyV3OpenApiNodeVO) GetOmadaDeviceType() string`

GetOmadaDeviceType returns the OmadaDeviceType field if non-nil, zero value otherwise.

### GetOmadaDeviceTypeOk

`func (o *TopologyV3OpenApiNodeVO) GetOmadaDeviceTypeOk() (*string, bool)`

GetOmadaDeviceTypeOk returns a tuple with the OmadaDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadaDeviceType

`func (o *TopologyV3OpenApiNodeVO) SetOmadaDeviceType(v string)`

SetOmadaDeviceType sets OmadaDeviceType field to given value.

### HasOmadaDeviceType

`func (o *TopologyV3OpenApiNodeVO) HasOmadaDeviceType() bool`

HasOmadaDeviceType returns a boolean if a field has been set.

### GetPeerLink

`func (o *TopologyV3OpenApiNodeVO) GetPeerLink() []WiredPortV3DTO`

GetPeerLink returns the PeerLink field if non-nil, zero value otherwise.

### GetPeerLinkOk

`func (o *TopologyV3OpenApiNodeVO) GetPeerLinkOk() (*[]WiredPortV3DTO, bool)`

GetPeerLinkOk returns a tuple with the PeerLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerLink

`func (o *TopologyV3OpenApiNodeVO) SetPeerLink(v []WiredPortV3DTO)`

SetPeerLink sets PeerLink field to given value.

### HasPeerLink

`func (o *TopologyV3OpenApiNodeVO) HasPeerLink() bool`

HasPeerLink returns a boolean if a field has been set.

### GetPortLabels

`func (o *TopologyV3OpenApiNodeVO) GetPortLabels() PortLabelDTO`

GetPortLabels returns the PortLabels field if non-nil, zero value otherwise.

### GetPortLabelsOk

`func (o *TopologyV3OpenApiNodeVO) GetPortLabelsOk() (*PortLabelDTO, bool)`

GetPortLabelsOk returns a tuple with the PortLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortLabels

`func (o *TopologyV3OpenApiNodeVO) SetPortLabels(v PortLabelDTO)`

SetPortLabels sets PortLabels field to given value.

### HasPortLabels

`func (o *TopologyV3OpenApiNodeVO) HasPortLabels() bool`

HasPortLabels returns a boolean if a field has been set.

### GetRdMode2g

`func (o *TopologyV3OpenApiNodeVO) GetRdMode2g() string`

GetRdMode2g returns the RdMode2g field if non-nil, zero value otherwise.

### GetRdMode2gOk

`func (o *TopologyV3OpenApiNodeVO) GetRdMode2gOk() (*string, bool)`

GetRdMode2gOk returns a tuple with the RdMode2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdMode2g

`func (o *TopologyV3OpenApiNodeVO) SetRdMode2g(v string)`

SetRdMode2g sets RdMode2g field to given value.

### HasRdMode2g

`func (o *TopologyV3OpenApiNodeVO) HasRdMode2g() bool`

HasRdMode2g returns a boolean if a field has been set.

### GetRdMode5g

`func (o *TopologyV3OpenApiNodeVO) GetRdMode5g() string`

GetRdMode5g returns the RdMode5g field if non-nil, zero value otherwise.

### GetRdMode5gOk

`func (o *TopologyV3OpenApiNodeVO) GetRdMode5gOk() (*string, bool)`

GetRdMode5gOk returns a tuple with the RdMode5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdMode5g

`func (o *TopologyV3OpenApiNodeVO) SetRdMode5g(v string)`

SetRdMode5g sets RdMode5g field to given value.

### HasRdMode5g

`func (o *TopologyV3OpenApiNodeVO) HasRdMode5g() bool`

HasRdMode5g returns a boolean if a field has been set.

### GetRdMode5g2

`func (o *TopologyV3OpenApiNodeVO) GetRdMode5g2() string`

GetRdMode5g2 returns the RdMode5g2 field if non-nil, zero value otherwise.

### GetRdMode5g2Ok

`func (o *TopologyV3OpenApiNodeVO) GetRdMode5g2Ok() (*string, bool)`

GetRdMode5g2Ok returns a tuple with the RdMode5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdMode5g2

`func (o *TopologyV3OpenApiNodeVO) SetRdMode5g2(v string)`

SetRdMode5g2 sets RdMode5g2 field to given value.

### HasRdMode5g2

`func (o *TopologyV3OpenApiNodeVO) HasRdMode5g2() bool`

HasRdMode5g2 returns a boolean if a field has been set.

### GetRdMode6g

`func (o *TopologyV3OpenApiNodeVO) GetRdMode6g() string`

GetRdMode6g returns the RdMode6g field if non-nil, zero value otherwise.

### GetRdMode6gOk

`func (o *TopologyV3OpenApiNodeVO) GetRdMode6gOk() (*string, bool)`

GetRdMode6gOk returns a tuple with the RdMode6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdMode6g

`func (o *TopologyV3OpenApiNodeVO) SetRdMode6g(v string)`

SetRdMode6g sets RdMode6g field to given value.

### HasRdMode6g

`func (o *TopologyV3OpenApiNodeVO) HasRdMode6g() bool`

HasRdMode6g returns a boolean if a field has been set.

### GetRole

`func (o *TopologyV3OpenApiNodeVO) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *TopologyV3OpenApiNodeVO) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *TopologyV3OpenApiNodeVO) SetRole(v int32)`

SetRole sets Role field to given value.

### HasRole

`func (o *TopologyV3OpenApiNodeVO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetShowModel

`func (o *TopologyV3OpenApiNodeVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *TopologyV3OpenApiNodeVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *TopologyV3OpenApiNodeVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *TopologyV3OpenApiNodeVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetSpecificType

`func (o *TopologyV3OpenApiNodeVO) GetSpecificType() int32`

GetSpecificType returns the SpecificType field if non-nil, zero value otherwise.

### GetSpecificTypeOk

`func (o *TopologyV3OpenApiNodeVO) GetSpecificTypeOk() (*int32, bool)`

GetSpecificTypeOk returns a tuple with the SpecificType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificType

`func (o *TopologyV3OpenApiNodeVO) SetSpecificType(v int32)`

SetSpecificType sets SpecificType field to given value.

### HasSpecificType

`func (o *TopologyV3OpenApiNodeVO) HasSpecificType() bool`

HasSpecificType returns a boolean if a field has been set.

### GetStackGroup

`func (o *TopologyV3OpenApiNodeVO) GetStackGroup() bool`

GetStackGroup returns the StackGroup field if non-nil, zero value otherwise.

### GetStackGroupOk

`func (o *TopologyV3OpenApiNodeVO) GetStackGroupOk() (*bool, bool)`

GetStackGroupOk returns a tuple with the StackGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackGroup

`func (o *TopologyV3OpenApiNodeVO) SetStackGroup(v bool)`

SetStackGroup sets StackGroup field to given value.

### HasStackGroup

`func (o *TopologyV3OpenApiNodeVO) HasStackGroup() bool`

HasStackGroup returns a boolean if a field has been set.

### GetStackId

`func (o *TopologyV3OpenApiNodeVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *TopologyV3OpenApiNodeVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *TopologyV3OpenApiNodeVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *TopologyV3OpenApiNodeVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackStatus

`func (o *TopologyV3OpenApiNodeVO) GetStackStatus() int32`

GetStackStatus returns the StackStatus field if non-nil, zero value otherwise.

### GetStackStatusOk

`func (o *TopologyV3OpenApiNodeVO) GetStackStatusOk() (*int32, bool)`

GetStackStatusOk returns a tuple with the StackStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackStatus

`func (o *TopologyV3OpenApiNodeVO) SetStackStatus(v int32)`

SetStackStatus sets StackStatus field to given value.

### HasStackStatus

`func (o *TopologyV3OpenApiNodeVO) HasStackStatus() bool`

HasStackStatus returns a boolean if a field has been set.

### GetStatus

`func (o *TopologyV3OpenApiNodeVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TopologyV3OpenApiNodeVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TopologyV3OpenApiNodeVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TopologyV3OpenApiNodeVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *TopologyV3OpenApiNodeVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *TopologyV3OpenApiNodeVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *TopologyV3OpenApiNodeVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *TopologyV3OpenApiNodeVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetSuccessors

`func (o *TopologyV3OpenApiNodeVO) GetSuccessors() []TopologyV3OpenApiNodeVO`

GetSuccessors returns the Successors field if non-nil, zero value otherwise.

### GetSuccessorsOk

`func (o *TopologyV3OpenApiNodeVO) GetSuccessorsOk() (*[]TopologyV3OpenApiNodeVO, bool)`

GetSuccessorsOk returns a tuple with the Successors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessors

`func (o *TopologyV3OpenApiNodeVO) SetSuccessors(v []TopologyV3OpenApiNodeVO)`

SetSuccessors sets Successors field to given value.

### HasSuccessors

`func (o *TopologyV3OpenApiNodeVO) HasSuccessors() bool`

HasSuccessors returns a boolean if a field has been set.

### GetSupport5g2

`func (o *TopologyV3OpenApiNodeVO) GetSupport5g2() bool`

GetSupport5g2 returns the Support5g2 field if non-nil, zero value otherwise.

### GetSupport5g2Ok

`func (o *TopologyV3OpenApiNodeVO) GetSupport5g2Ok() (*bool, bool)`

GetSupport5g2Ok returns a tuple with the Support5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport5g2

`func (o *TopologyV3OpenApiNodeVO) SetSupport5g2(v bool)`

SetSupport5g2 sets Support5g2 field to given value.

### HasSupport5g2

`func (o *TopologyV3OpenApiNodeVO) HasSupport5g2() bool`

HasSupport5g2 returns a boolean if a field has been set.

### GetType

`func (o *TopologyV3OpenApiNodeVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TopologyV3OpenApiNodeVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TopologyV3OpenApiNodeVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TopologyV3OpenApiNodeVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUplinkMacList

`func (o *TopologyV3OpenApiNodeVO) GetUplinkMacList() []string`

GetUplinkMacList returns the UplinkMacList field if non-nil, zero value otherwise.

### GetUplinkMacListOk

`func (o *TopologyV3OpenApiNodeVO) GetUplinkMacListOk() (*[]string, bool)`

GetUplinkMacListOk returns a tuple with the UplinkMacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkMacList

`func (o *TopologyV3OpenApiNodeVO) SetUplinkMacList(v []string)`

SetUplinkMacList sets UplinkMacList field to given value.

### HasUplinkMacList

`func (o *TopologyV3OpenApiNodeVO) HasUplinkMacList() bool`

HasUplinkMacList returns a boolean if a field has been set.

### GetVigiManaged

`func (o *TopologyV3OpenApiNodeVO) GetVigiManaged() bool`

GetVigiManaged returns the VigiManaged field if non-nil, zero value otherwise.

### GetVigiManagedOk

`func (o *TopologyV3OpenApiNodeVO) GetVigiManagedOk() (*bool, bool)`

GetVigiManagedOk returns a tuple with the VigiManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVigiManaged

`func (o *TopologyV3OpenApiNodeVO) SetVigiManaged(v bool)`

SetVigiManaged sets VigiManaged field to given value.

### HasVigiManaged

`func (o *TopologyV3OpenApiNodeVO) HasVigiManaged() bool`

HasVigiManaged returns a boolean if a field has been set.

### GetVigiWirelessUpInfo

`func (o *TopologyV3OpenApiNodeVO) GetVigiWirelessUpInfo() VigiWirelessUpInfoDTO`

GetVigiWirelessUpInfo returns the VigiWirelessUpInfo field if non-nil, zero value otherwise.

### GetVigiWirelessUpInfoOk

`func (o *TopologyV3OpenApiNodeVO) GetVigiWirelessUpInfoOk() (*VigiWirelessUpInfoDTO, bool)`

GetVigiWirelessUpInfoOk returns a tuple with the VigiWirelessUpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVigiWirelessUpInfo

`func (o *TopologyV3OpenApiNodeVO) SetVigiWirelessUpInfo(v VigiWirelessUpInfoDTO)`

SetVigiWirelessUpInfo sets VigiWirelessUpInfo field to given value.

### HasVigiWirelessUpInfo

`func (o *TopologyV3OpenApiNodeVO) HasVigiWirelessUpInfo() bool`

HasVigiWirelessUpInfo returns a boolean if a field has been set.

### GetVlansAsMaster

`func (o *TopologyV3OpenApiNodeVO) GetVlansAsMaster() string`

GetVlansAsMaster returns the VlansAsMaster field if non-nil, zero value otherwise.

### GetVlansAsMasterOk

`func (o *TopologyV3OpenApiNodeVO) GetVlansAsMasterOk() (*string, bool)`

GetVlansAsMasterOk returns a tuple with the VlansAsMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlansAsMaster

`func (o *TopologyV3OpenApiNodeVO) SetVlansAsMaster(v string)`

SetVlansAsMaster sets VlansAsMaster field to given value.

### HasVlansAsMaster

`func (o *TopologyV3OpenApiNodeVO) HasVlansAsMaster() bool`

HasVlansAsMaster returns a boolean if a field has been set.

### GetVrrpGroupList

`func (o *TopologyV3OpenApiNodeVO) GetVrrpGroupList() []VrrpGroupDTO`

GetVrrpGroupList returns the VrrpGroupList field if non-nil, zero value otherwise.

### GetVrrpGroupListOk

`func (o *TopologyV3OpenApiNodeVO) GetVrrpGroupListOk() (*[]VrrpGroupDTO, bool)`

GetVrrpGroupListOk returns a tuple with the VrrpGroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrrpGroupList

`func (o *TopologyV3OpenApiNodeVO) SetVrrpGroupList(v []VrrpGroupDTO)`

SetVrrpGroupList sets VrrpGroupList field to given value.

### HasVrrpGroupList

`func (o *TopologyV3OpenApiNodeVO) HasVrrpGroupList() bool`

HasVrrpGroupList returns a boolean if a field has been set.

### GetVrrpLink

`func (o *TopologyV3OpenApiNodeVO) GetVrrpLink() []VrrpLinkDTO`

GetVrrpLink returns the VrrpLink field if non-nil, zero value otherwise.

### GetVrrpLinkOk

`func (o *TopologyV3OpenApiNodeVO) GetVrrpLinkOk() (*[]VrrpLinkDTO, bool)`

GetVrrpLinkOk returns a tuple with the VrrpLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrrpLink

`func (o *TopologyV3OpenApiNodeVO) SetVrrpLink(v []VrrpLinkDTO)`

SetVrrpLink sets VrrpLink field to given value.

### HasVrrpLink

`func (o *TopologyV3OpenApiNodeVO) HasVrrpLink() bool`

HasVrrpLink returns a boolean if a field has been set.

### GetWanPorts

`func (o *TopologyV3OpenApiNodeVO) GetWanPorts() []WanPort`

GetWanPorts returns the WanPorts field if non-nil, zero value otherwise.

### GetWanPortsOk

`func (o *TopologyV3OpenApiNodeVO) GetWanPortsOk() (*[]WanPort, bool)`

GetWanPortsOk returns a tuple with the WanPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPorts

`func (o *TopologyV3OpenApiNodeVO) SetWanPorts(v []WanPort)`

SetWanPorts sets WanPorts field to given value.

### HasWanPorts

`func (o *TopologyV3OpenApiNodeVO) HasWanPorts() bool`

HasWanPorts returns a boolean if a field has been set.

### GetWarnNvrFirmware

`func (o *TopologyV3OpenApiNodeVO) GetWarnNvrFirmware() bool`

GetWarnNvrFirmware returns the WarnNvrFirmware field if non-nil, zero value otherwise.

### GetWarnNvrFirmwareOk

`func (o *TopologyV3OpenApiNodeVO) GetWarnNvrFirmwareOk() (*bool, bool)`

GetWarnNvrFirmwareOk returns a tuple with the WarnNvrFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnNvrFirmware

`func (o *TopologyV3OpenApiNodeVO) SetWarnNvrFirmware(v bool)`

SetWarnNvrFirmware sets WarnNvrFirmware field to given value.

### HasWarnNvrFirmware

`func (o *TopologyV3OpenApiNodeVO) HasWarnNvrFirmware() bool`

HasWarnNvrFirmware returns a boolean if a field has been set.

### GetWiredUpInfo

`func (o *TopologyV3OpenApiNodeVO) GetWiredUpInfo() WiredUpInfoDTO`

GetWiredUpInfo returns the WiredUpInfo field if non-nil, zero value otherwise.

### GetWiredUpInfoOk

`func (o *TopologyV3OpenApiNodeVO) GetWiredUpInfoOk() (*WiredUpInfoDTO, bool)`

GetWiredUpInfoOk returns a tuple with the WiredUpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredUpInfo

`func (o *TopologyV3OpenApiNodeVO) SetWiredUpInfo(v WiredUpInfoDTO)`

SetWiredUpInfo sets WiredUpInfo field to given value.

### HasWiredUpInfo

`func (o *TopologyV3OpenApiNodeVO) HasWiredUpInfo() bool`

HasWiredUpInfo returns a boolean if a field has been set.

### GetWiredUpInfos

`func (o *TopologyV3OpenApiNodeVO) GetWiredUpInfos() []WiredUpInfoDTO`

GetWiredUpInfos returns the WiredUpInfos field if non-nil, zero value otherwise.

### GetWiredUpInfosOk

`func (o *TopologyV3OpenApiNodeVO) GetWiredUpInfosOk() (*[]WiredUpInfoDTO, bool)`

GetWiredUpInfosOk returns a tuple with the WiredUpInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredUpInfos

`func (o *TopologyV3OpenApiNodeVO) SetWiredUpInfos(v []WiredUpInfoDTO)`

SetWiredUpInfos sets WiredUpInfos field to given value.

### HasWiredUpInfos

`func (o *TopologyV3OpenApiNodeVO) HasWiredUpInfos() bool`

HasWiredUpInfos returns a boolean if a field has been set.

### GetWirelessUpInfo

`func (o *TopologyV3OpenApiNodeVO) GetWirelessUpInfo() WirelessUpInfoDTO`

GetWirelessUpInfo returns the WirelessUpInfo field if non-nil, zero value otherwise.

### GetWirelessUpInfoOk

`func (o *TopologyV3OpenApiNodeVO) GetWirelessUpInfoOk() (*WirelessUpInfoDTO, bool)`

GetWirelessUpInfoOk returns a tuple with the WirelessUpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessUpInfo

`func (o *TopologyV3OpenApiNodeVO) SetWirelessUpInfo(v WirelessUpInfoDTO)`

SetWirelessUpInfo sets WirelessUpInfo field to given value.

### HasWirelessUpInfo

`func (o *TopologyV3OpenApiNodeVO) HasWirelessUpInfo() bool`

HasWirelessUpInfo returns a boolean if a field has been set.

### GetWirelessUplink

`func (o *TopologyV3OpenApiNodeVO) GetWirelessUplink() bool`

GetWirelessUplink returns the WirelessUplink field if non-nil, zero value otherwise.

### GetWirelessUplinkOk

`func (o *TopologyV3OpenApiNodeVO) GetWirelessUplinkOk() (*bool, bool)`

GetWirelessUplinkOk returns a tuple with the WirelessUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessUplink

`func (o *TopologyV3OpenApiNodeVO) SetWirelessUplink(v bool)`

SetWirelessUplink sets WirelessUplink field to given value.

### HasWirelessUplink

`func (o *TopologyV3OpenApiNodeVO) HasWirelessUplink() bool`

HasWirelessUplink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


