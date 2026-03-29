# OswVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | \&quot;Site of the device\&quot; | [optional] 
**Active** | Pointer to **bool** | Mark whether the device is activated: When license (specific to cloud base) is false, the status column shows the pre-bound status. When it is true or null, the status column displays as it originally did. | [optional] 
**AddedInAdvanced** | Pointer to **bool** | Whether the device is added in advanced. | [optional] 
**Address** | Pointer to **string** | Address | [optional] 
**AdoptFailType** | Pointer to **int32** | Adopt fail reason should be a value as follows: -1:adopt timeout;-2:user/password error | [optional] 
**Category** | Pointer to **string** | Category of license.When activating in bulk for the pro site, the front end can only select the same type for bulk activation. | [optional] 
**ClientNum** | Pointer to **int32** | Number of clients. | [optional] 
**CombinedGateway** | Pointer to **bool** | Is it an combined gateway? | [optional] 
**Compatible** | Pointer to **int32** | Device firmware and controller compatibility type.Compatible should be a value as follows: 0:COMPATIBLE;1:HIGH_MAJOR_VER;2:LOW_MAJOR_VER;3:HIGH_MINOR_VER;4:LOW_MINOR_VER;7:HIGH_COMPONENT_VER;10:DEVICE_NOT_COMPATIBLE;11:HIGH_ADOPT_COMMPONENT;12:DEVICE_CATEGORY_NOT_COMPATIBLE;14:DEVICE_NOT_COMPATIBLE_IN_CLUSTER | [optional] 
**CompoundModel** | Pointer to **string** | Model complex used in the backend.Ap：model+(country)+modelVersion,  EAP225(EU) v3.0 Ap: specialModel+modelVersion, EAP225-Outdoor-1a20a950b8d950e8 v1.0  Gateway/Switch：model+modelVersion, Osg v3.0 | [optional] 
**ConfigSyncStatus** | Pointer to **int32** | Device configuration synchronization status , including: 0: noConfig (No configuration) 1: success (Success) 2: fail (Failure) 3: info (Information) 4: configuring (Configuring) 5: preConfig (Pre-configuration) | [optional] 
**CpuUtil** | Pointer to **int32** | Cpu utilization | [optional] 
**CustomId** | Pointer to **string** | Customer ID | [optional] 
**CustomName** | Pointer to **string** | Customer name displayed in MSP mode | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DeviceSeriesType** | Pointer to **int32** | Device type should be a value as follows:0:advanced,1:pro | [optional] 
**Download** | Pointer to **int64** | Real-time total downstream traffic (bytes). | [optional] 
**DueTime** | Pointer to **int64** | Expire timestamp of license(cloud base exclusive) | [optional] 
**DueTimeLeft** | Pointer to **int64** | Milliseconds from the current moment to the expiration time(cloud base exclusive) | [optional] 
**Duplex** | Pointer to **int32** | Duplex should be a value as follows:0:Auto,1:Half, 2:Full | [optional] 
**EcspFirstVersion** | Pointer to **int32** | Ecsp first version | [optional] 
**Eos** | Pointer to **int64** | End of support time of device(CBC exclusive) | [optional] 
**Eost** | Pointer to **int64** | End of service time of device(CBC exclusive) | [optional] 
**Es** | Pointer to **bool** | Whether it is Agile Series Switch | [optional] 
**FanStatus** | Pointer to **int32** | FanStatus should be a value as follows: 0: normal; 1: fault; 2: no fan | [optional] 
**FeatureLimit** | Pointer to [**ApFeatureLimitVO**](ApFeatureLimitVO.md) |  | [optional] 
**FirmwareVersion** | Pointer to **string** | Version of firmware,for example:2.5.0 Build 20190118 Rel. 64821 | [optional] 
**FwDownload** | Pointer to **bool** | Whether the device is downloading firmware | [optional] 
**HealthScore** | Pointer to **int32** | 1~3: poor; 4~7: fair; 0: no data; 8~10 good. | [optional] 
**HealthScoreTime** | Pointer to **int64** | The time of healthScore | [optional] 
**HwVersion** | Pointer to **string** | Version of hardware,for example 1.0 | [optional] 
**InWhitelist** | Pointer to **bool** | Whether the device is in white list | [optional] 
**Ip** | Pointer to **string** | Ip address,such as 192.168.0.105 | [optional] 
**Ippt** | Pointer to **bool** | Whether it is LTE Backup | [optional] 
**Ipv6List** | Pointer to **[]string** | Ipv6 address List | [optional] 
**LastSeen** | Pointer to **int64** | Last active time. | [optional] 
**LatestVersion** | Pointer to **string** | Latest firmware version | [optional] 
**LicenseStatus** | Pointer to **int32** | License status should be a value as follows:0: unActive 1: Unbind 2: Expired 3: active If there is a value and it is not 3, display the Active button. license (specific to cloud base). | [optional] 
**LicenseStatusStr** | Pointer to **string** |  | [optional] 
**LicenseUnbindingLimit** | Pointer to **int32** | Remaining unbind count for license on detail Page of device(cloud base exclusive) | [optional] 
**LinkSpeed** | Pointer to **int32** | Link speed | [optional] 
**LocateEnable** | Pointer to **bool** | Whether locate function is enabled | [optional] 
**LocatingPorts** | Pointer to **[]int32** | Locating switch ports, effective when parameter[locateEnable] is true and device is not a member of stack | [optional] 
**LocatingStandardPorts** | Pointer to **[]string** | Locating switch standard ports, effective when parameter[locateEnable] is true and device is a member of stack | [optional] 
**Location** | Pointer to [**LocationVO**](LocationVO.md) |  | [optional] 
**Loop** | Pointer to **string** | Set of loop port | [optional] 
**LoopbackNum** | Pointer to **int32** | Number of loops | [optional] 
**Mac** | Pointer to **string** | Mac address | [optional] 
**MaxStackGroups** | Pointer to **int32** | The maximum number of stacking port aggregation groups that can be configured | [optional] 
**MaxStackUnitNumber** | Pointer to **int32** | The maximum unit number supported by the device | [optional] 
**MemUtil** | Pointer to **int32** | Memory utilization | [optional] 
**MlagMsg** | Pointer to [**MlagMsgVO**](MlagMsgVO.md) |  | [optional] 
**Model** | Pointer to **string** | Model, such as EAP225. | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**MstpInsNum** | Pointer to **int32** | The number of MSTP instances | [optional] 
**Name** | Pointer to **string** | Default uses the MAC address as the name. | [optional] 
**NeedUpgrade** | Pointer to **bool** | Whether the device needs upgrade | [optional] 
**OnlineUpgradeStatus** | Pointer to **int32** | Device online upgrade status should be a value as follows: 0:IDLE,1:DOWNLOADING,2:UPGRADING | [optional] 
**PackageCaptureStatus** | Pointer to **int32** |  | [optional] 
**PoeRemain** | Pointer to **float64** | PoE remaining power | [optional] 
**PoeSupport** | Pointer to **bool** | Indicates whether the switch supports PoE | [optional] 
**PoeTotalPower** | Pointer to **float64** | PoE Total Power (W) | [optional] 
**Ports** | Pointer to [**[]OswStackMemberPortVO**](OswStackMemberPortVO.md) | Ports | [optional] 
**PowerMode** | Pointer to **int32** | AP power mode.The power supply types are as follows: 0: DC power supply 1: 802.3bt power supply 2: 802.3at power supply 3: 802.3af power supply 4: Switch power supply (display restart) | [optional] 
**PowerModeList** | Pointer to **[]int32** | Indicating the power supply modes for AP devices across frequency bands, with array indices corresponding to the frequency bands: 0: 2.4G 1: 5G/5G1 2: 5G2 3: 6G. The corresponding values for power supply modes are: 0: Normal operation 1: Power limited 2: Frequency band disabled | [optional] 
**PreConfigErrorCode** | Pointer to **int32** | When the pre-added devices cannot properly pass through, an error will be returned. | [optional] 
**PreConfigRetryType** | Pointer to **int32** | When retrying the path through, if there is a value, the front end will display a retry button. The types are: 1:Requires password input.2: No password input needed. | [optional] 
**Priority** | Pointer to **int32** | Stack member priority | [optional] 
**Profiles** | Pointer to **map[string]string** | profiles | [optional] 
**PublicIp** | Pointer to **string** | Public ip address | [optional] 
**Resource** | Pointer to **int32** | Data source.Resource should be a value as follows: 0:new created;1:from template;2:override | [optional] 
**Sdm** | Pointer to [**OswSdmTemplateVO**](OswSdmTemplateVO.md) |  | [optional] 
**ShowModel** | Pointer to **string** | Model complex shown in the front end.Ap：model+(country)+modelVersion,EAP225(EU) v3.0  Gateway/Switch：model+modelVersion,Osg v3.0 | [optional] 
**SiteName** | Pointer to **string** | Site name | [optional] 
**Sn** | Pointer to **string** | SN code of device | [optional] 
**SpecialModel** | Pointer to **string** | Special device model,for example:EAP225-Outdoor-1a20a950b8d950e8 | [optional] 
**StackMsg** | Pointer to [**StackMsgVO**](StackMsgVO.md) |  | [optional] 
**StackPortCap** | Pointer to **map[string][]string** |  | [optional] 
**StackPorts** | Pointer to [**[]OswStackPortGroupVO**](OswStackPortGroupVO.md) | Stack ports | [optional] 
**StackSupportPorts** | Pointer to [**[]OswStandPortVO**](OswStandPortVO.md) | Stack support ports | [optional] 
**Status** | Pointer to **int32** | Status of device,status should be a value as follows: 0:Disconnected;1:Disconnected(Migrating);10:Provisioning;11:Configuring;12:Upgrading;13:Rebooting;14:Connected;15:Connected(Wireless);16:Connected(Migrating);17:Connected(Wireless,Migrating);20:Pending;21:Pending(Wireless);22:Adopting;23:Adopting(Wireless);24:Adopt Failed;25:Adopt Failed(Wireless);26:Managed By Others;27:Managed By Others(Wireless);30:Heartbeat Missed;31:Heartbeat Missed(Wireless);32:Heartbeat Missed(Migrating);33:Heartbeat Missed(Wireless,Migrating);40:Isolated;41:Isolated(Migrating);50:Slice Configuring | [optional] 
**StatusCategory** | Pointer to **int32** | Category of device status,statusCategory should be a value as follows: 0:Disconnected;1:Connected;2:Pending;3:Heartbeat Missed;4:Isolated | [optional] 
**StkVer** | Pointer to **string** | The stacking component version number used by the switch is used to determine whether two devices can be stacked | [optional] 
**StkableGroupId** | Pointer to **int32** | The stack support list number where the switch is located, used to determine whether two devices can be stacked | [optional] 
**SubDevType** | Pointer to **int32** | Device subtype, indicating a special device type for wireless routers. This field is not present for ordinary devices, where 1 represents a wireless router. | [optional] 
**SupportAnomaly** | Pointer to **bool** | Whether the device firmware support intelligent anomaly detection | [optional] 
**SupportCableTest** | Pointer to **bool** | Indicates whether the switch supports cable test | [optional] 
**SupportCustomDhcpOption** | Pointer to **bool** | Whether the device supports custom dhcp option. Only valid when type is switch. | [optional] 
**SupportDhcpRange** | Pointer to **bool** | Indicates whether the switch supports DHCP Range Pool | [optional] 
**SupportDhcpReservation** | Pointer to **bool** | Indicates whether the switch supports DHCP Reservation | [optional] 
**SupportDomainPing** | Pointer to **bool** | Indicates whether the switch supports pinging domain | [optional] 
**SupportDomainTraceRoute** | Pointer to **bool** | Indicates whether the switch supports tracerouting domain | [optional] 
**SupportExtendStp** | Pointer to **bool** | Indicates whether the switch supports mstp | [optional] 
**SupportGetOspfNeighborTable** | Pointer to **bool** |  | [optional] 
**SupportIppt** | Pointer to **bool** |  | [optional] 
**SupportMacDelay** | Pointer to **bool** |  | [optional] 
**SupportPowerAlert** | Pointer to **bool** | Indicates whether the switch supports power alert | [optional] 
**SupportRunningConfig** | Pointer to **bool** | Whether the device supports show running config. | [optional] 
**SupportSdm** | Pointer to **bool** | Indicates whether the switch supports SDM template | [optional] 
**SupportStp** | Pointer to **bool** | Indicates whether the switch supports stp | [optional] 
**SupportVrf** | Pointer to **bool** | Indicates whether the switch supports Vrf | [optional] 
**SwitchConsistent** | Pointer to **bool** | Whether the device can be adopted by the site. | [optional] 
**TagName** | Pointer to **string** | Device tag name | [optional] 
**Type** | Pointer to **string** | Device type:ap、gateway、switch、olt | [optional] 
**Unit** | Pointer to **int32** | Unit | [optional] 
**UplinkDeviceMac** | Pointer to **string** | Uplink device mac | [optional] 
**UplinkDeviceName** | Pointer to **string** | Uplink device mac name | [optional] 
**UplinkDevicePort** | Pointer to **string** | Uplink device port | [optional] 
**Upload** | Pointer to **int64** | Real-time total upstream traffic (bytes). | [optional] 
**Uptime** | Pointer to **string** | Device uptime | [optional] 
**UptimeLong** | Pointer to **int64** | Runtime duration, in seconds (s). | [optional] 
**Version** | Pointer to **string** | Software version, such as \&quot;2.5.0,\&quot; extracted from DeviceDO.firmwareVersion - \&quot;2.5.0 Build 20190118 Rel. 64821.\&quot; | [optional] 
**WirelessRouter** | Pointer to **bool** | Indicates whether it is a wireless router; it is true only when this device is a wireless router. | [optional] 

## Methods

### NewOswVO

`func NewOswVO() *OswVO`

NewOswVO instantiates a new OswVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswVOWithDefaults

`func NewOswVOWithDefaults() *OswVO`

NewOswVOWithDefaults instantiates a new OswVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *OswVO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *OswVO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *OswVO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *OswVO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetActive

`func (o *OswVO) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OswVO) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OswVO) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *OswVO) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddedInAdvanced

`func (o *OswVO) GetAddedInAdvanced() bool`

GetAddedInAdvanced returns the AddedInAdvanced field if non-nil, zero value otherwise.

### GetAddedInAdvancedOk

`func (o *OswVO) GetAddedInAdvancedOk() (*bool, bool)`

GetAddedInAdvancedOk returns a tuple with the AddedInAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedInAdvanced

`func (o *OswVO) SetAddedInAdvanced(v bool)`

SetAddedInAdvanced sets AddedInAdvanced field to given value.

### HasAddedInAdvanced

`func (o *OswVO) HasAddedInAdvanced() bool`

HasAddedInAdvanced returns a boolean if a field has been set.

### GetAddress

`func (o *OswVO) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OswVO) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OswVO) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *OswVO) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAdoptFailType

`func (o *OswVO) GetAdoptFailType() int32`

GetAdoptFailType returns the AdoptFailType field if non-nil, zero value otherwise.

### GetAdoptFailTypeOk

`func (o *OswVO) GetAdoptFailTypeOk() (*int32, bool)`

GetAdoptFailTypeOk returns a tuple with the AdoptFailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdoptFailType

`func (o *OswVO) SetAdoptFailType(v int32)`

SetAdoptFailType sets AdoptFailType field to given value.

### HasAdoptFailType

`func (o *OswVO) HasAdoptFailType() bool`

HasAdoptFailType returns a boolean if a field has been set.

### GetCategory

`func (o *OswVO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *OswVO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *OswVO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *OswVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetClientNum

`func (o *OswVO) GetClientNum() int32`

GetClientNum returns the ClientNum field if non-nil, zero value otherwise.

### GetClientNumOk

`func (o *OswVO) GetClientNumOk() (*int32, bool)`

GetClientNumOk returns a tuple with the ClientNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNum

`func (o *OswVO) SetClientNum(v int32)`

SetClientNum sets ClientNum field to given value.

### HasClientNum

`func (o *OswVO) HasClientNum() bool`

HasClientNum returns a boolean if a field has been set.

### GetCombinedGateway

`func (o *OswVO) GetCombinedGateway() bool`

GetCombinedGateway returns the CombinedGateway field if non-nil, zero value otherwise.

### GetCombinedGatewayOk

`func (o *OswVO) GetCombinedGatewayOk() (*bool, bool)`

GetCombinedGatewayOk returns a tuple with the CombinedGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedGateway

`func (o *OswVO) SetCombinedGateway(v bool)`

SetCombinedGateway sets CombinedGateway field to given value.

### HasCombinedGateway

`func (o *OswVO) HasCombinedGateway() bool`

HasCombinedGateway returns a boolean if a field has been set.

### GetCompatible

`func (o *OswVO) GetCompatible() int32`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *OswVO) GetCompatibleOk() (*int32, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *OswVO) SetCompatible(v int32)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *OswVO) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetCompoundModel

`func (o *OswVO) GetCompoundModel() string`

GetCompoundModel returns the CompoundModel field if non-nil, zero value otherwise.

### GetCompoundModelOk

`func (o *OswVO) GetCompoundModelOk() (*string, bool)`

GetCompoundModelOk returns a tuple with the CompoundModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompoundModel

`func (o *OswVO) SetCompoundModel(v string)`

SetCompoundModel sets CompoundModel field to given value.

### HasCompoundModel

`func (o *OswVO) HasCompoundModel() bool`

HasCompoundModel returns a boolean if a field has been set.

### GetConfigSyncStatus

`func (o *OswVO) GetConfigSyncStatus() int32`

GetConfigSyncStatus returns the ConfigSyncStatus field if non-nil, zero value otherwise.

### GetConfigSyncStatusOk

`func (o *OswVO) GetConfigSyncStatusOk() (*int32, bool)`

GetConfigSyncStatusOk returns a tuple with the ConfigSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSyncStatus

`func (o *OswVO) SetConfigSyncStatus(v int32)`

SetConfigSyncStatus sets ConfigSyncStatus field to given value.

### HasConfigSyncStatus

`func (o *OswVO) HasConfigSyncStatus() bool`

HasConfigSyncStatus returns a boolean if a field has been set.

### GetCpuUtil

`func (o *OswVO) GetCpuUtil() int32`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *OswVO) GetCpuUtilOk() (*int32, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *OswVO) SetCpuUtil(v int32)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *OswVO) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetCustomId

`func (o *OswVO) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *OswVO) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *OswVO) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.

### HasCustomId

`func (o *OswVO) HasCustomId() bool`

HasCustomId returns a boolean if a field has been set.

### GetCustomName

`func (o *OswVO) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *OswVO) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *OswVO) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *OswVO) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### GetDescription

`func (o *OswVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OswVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OswVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OswVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceSeriesType

`func (o *OswVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *OswVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *OswVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *OswVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetDownload

`func (o *OswVO) GetDownload() int64`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *OswVO) GetDownloadOk() (*int64, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *OswVO) SetDownload(v int64)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *OswVO) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetDueTime

`func (o *OswVO) GetDueTime() int64`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *OswVO) GetDueTimeOk() (*int64, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *OswVO) SetDueTime(v int64)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *OswVO) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### GetDueTimeLeft

`func (o *OswVO) GetDueTimeLeft() int64`

GetDueTimeLeft returns the DueTimeLeft field if non-nil, zero value otherwise.

### GetDueTimeLeftOk

`func (o *OswVO) GetDueTimeLeftOk() (*int64, bool)`

GetDueTimeLeftOk returns a tuple with the DueTimeLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTimeLeft

`func (o *OswVO) SetDueTimeLeft(v int64)`

SetDueTimeLeft sets DueTimeLeft field to given value.

### HasDueTimeLeft

`func (o *OswVO) HasDueTimeLeft() bool`

HasDueTimeLeft returns a boolean if a field has been set.

### GetDuplex

`func (o *OswVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *OswVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *OswVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *OswVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetEcspFirstVersion

`func (o *OswVO) GetEcspFirstVersion() int32`

GetEcspFirstVersion returns the EcspFirstVersion field if non-nil, zero value otherwise.

### GetEcspFirstVersionOk

`func (o *OswVO) GetEcspFirstVersionOk() (*int32, bool)`

GetEcspFirstVersionOk returns a tuple with the EcspFirstVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcspFirstVersion

`func (o *OswVO) SetEcspFirstVersion(v int32)`

SetEcspFirstVersion sets EcspFirstVersion field to given value.

### HasEcspFirstVersion

`func (o *OswVO) HasEcspFirstVersion() bool`

HasEcspFirstVersion returns a boolean if a field has been set.

### GetEos

`func (o *OswVO) GetEos() int64`

GetEos returns the Eos field if non-nil, zero value otherwise.

### GetEosOk

`func (o *OswVO) GetEosOk() (*int64, bool)`

GetEosOk returns a tuple with the Eos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEos

`func (o *OswVO) SetEos(v int64)`

SetEos sets Eos field to given value.

### HasEos

`func (o *OswVO) HasEos() bool`

HasEos returns a boolean if a field has been set.

### GetEost

`func (o *OswVO) GetEost() int64`

GetEost returns the Eost field if non-nil, zero value otherwise.

### GetEostOk

`func (o *OswVO) GetEostOk() (*int64, bool)`

GetEostOk returns a tuple with the Eost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEost

`func (o *OswVO) SetEost(v int64)`

SetEost sets Eost field to given value.

### HasEost

`func (o *OswVO) HasEost() bool`

HasEost returns a boolean if a field has been set.

### GetEs

`func (o *OswVO) GetEs() bool`

GetEs returns the Es field if non-nil, zero value otherwise.

### GetEsOk

`func (o *OswVO) GetEsOk() (*bool, bool)`

GetEsOk returns a tuple with the Es field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEs

`func (o *OswVO) SetEs(v bool)`

SetEs sets Es field to given value.

### HasEs

`func (o *OswVO) HasEs() bool`

HasEs returns a boolean if a field has been set.

### GetFanStatus

`func (o *OswVO) GetFanStatus() int32`

GetFanStatus returns the FanStatus field if non-nil, zero value otherwise.

### GetFanStatusOk

`func (o *OswVO) GetFanStatusOk() (*int32, bool)`

GetFanStatusOk returns a tuple with the FanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanStatus

`func (o *OswVO) SetFanStatus(v int32)`

SetFanStatus sets FanStatus field to given value.

### HasFanStatus

`func (o *OswVO) HasFanStatus() bool`

HasFanStatus returns a boolean if a field has been set.

### GetFeatureLimit

`func (o *OswVO) GetFeatureLimit() ApFeatureLimitVO`

GetFeatureLimit returns the FeatureLimit field if non-nil, zero value otherwise.

### GetFeatureLimitOk

`func (o *OswVO) GetFeatureLimitOk() (*ApFeatureLimitVO, bool)`

GetFeatureLimitOk returns a tuple with the FeatureLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureLimit

`func (o *OswVO) SetFeatureLimit(v ApFeatureLimitVO)`

SetFeatureLimit sets FeatureLimit field to given value.

### HasFeatureLimit

`func (o *OswVO) HasFeatureLimit() bool`

HasFeatureLimit returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *OswVO) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *OswVO) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *OswVO) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *OswVO) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetFwDownload

`func (o *OswVO) GetFwDownload() bool`

GetFwDownload returns the FwDownload field if non-nil, zero value otherwise.

### GetFwDownloadOk

`func (o *OswVO) GetFwDownloadOk() (*bool, bool)`

GetFwDownloadOk returns a tuple with the FwDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwDownload

`func (o *OswVO) SetFwDownload(v bool)`

SetFwDownload sets FwDownload field to given value.

### HasFwDownload

`func (o *OswVO) HasFwDownload() bool`

HasFwDownload returns a boolean if a field has been set.

### GetHealthScore

`func (o *OswVO) GetHealthScore() int32`

GetHealthScore returns the HealthScore field if non-nil, zero value otherwise.

### GetHealthScoreOk

`func (o *OswVO) GetHealthScoreOk() (*int32, bool)`

GetHealthScoreOk returns a tuple with the HealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthScore

`func (o *OswVO) SetHealthScore(v int32)`

SetHealthScore sets HealthScore field to given value.

### HasHealthScore

`func (o *OswVO) HasHealthScore() bool`

HasHealthScore returns a boolean if a field has been set.

### GetHealthScoreTime

`func (o *OswVO) GetHealthScoreTime() int64`

GetHealthScoreTime returns the HealthScoreTime field if non-nil, zero value otherwise.

### GetHealthScoreTimeOk

`func (o *OswVO) GetHealthScoreTimeOk() (*int64, bool)`

GetHealthScoreTimeOk returns a tuple with the HealthScoreTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthScoreTime

`func (o *OswVO) SetHealthScoreTime(v int64)`

SetHealthScoreTime sets HealthScoreTime field to given value.

### HasHealthScoreTime

`func (o *OswVO) HasHealthScoreTime() bool`

HasHealthScoreTime returns a boolean if a field has been set.

### GetHwVersion

`func (o *OswVO) GetHwVersion() string`

GetHwVersion returns the HwVersion field if non-nil, zero value otherwise.

### GetHwVersionOk

`func (o *OswVO) GetHwVersionOk() (*string, bool)`

GetHwVersionOk returns a tuple with the HwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwVersion

`func (o *OswVO) SetHwVersion(v string)`

SetHwVersion sets HwVersion field to given value.

### HasHwVersion

`func (o *OswVO) HasHwVersion() bool`

HasHwVersion returns a boolean if a field has been set.

### GetInWhitelist

`func (o *OswVO) GetInWhitelist() bool`

GetInWhitelist returns the InWhitelist field if non-nil, zero value otherwise.

### GetInWhitelistOk

`func (o *OswVO) GetInWhitelistOk() (*bool, bool)`

GetInWhitelistOk returns a tuple with the InWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInWhitelist

`func (o *OswVO) SetInWhitelist(v bool)`

SetInWhitelist sets InWhitelist field to given value.

### HasInWhitelist

`func (o *OswVO) HasInWhitelist() bool`

HasInWhitelist returns a boolean if a field has been set.

### GetIp

`func (o *OswVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OswVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OswVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OswVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIppt

`func (o *OswVO) GetIppt() bool`

GetIppt returns the Ippt field if non-nil, zero value otherwise.

### GetIpptOk

`func (o *OswVO) GetIpptOk() (*bool, bool)`

GetIpptOk returns a tuple with the Ippt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIppt

`func (o *OswVO) SetIppt(v bool)`

SetIppt sets Ippt field to given value.

### HasIppt

`func (o *OswVO) HasIppt() bool`

HasIppt returns a boolean if a field has been set.

### GetIpv6List

`func (o *OswVO) GetIpv6List() []string`

GetIpv6List returns the Ipv6List field if non-nil, zero value otherwise.

### GetIpv6ListOk

`func (o *OswVO) GetIpv6ListOk() (*[]string, bool)`

GetIpv6ListOk returns a tuple with the Ipv6List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6List

`func (o *OswVO) SetIpv6List(v []string)`

SetIpv6List sets Ipv6List field to given value.

### HasIpv6List

`func (o *OswVO) HasIpv6List() bool`

HasIpv6List returns a boolean if a field has been set.

### GetLastSeen

`func (o *OswVO) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *OswVO) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *OswVO) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *OswVO) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLatestVersion

`func (o *OswVO) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *OswVO) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *OswVO) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *OswVO) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetLicenseStatus

`func (o *OswVO) GetLicenseStatus() int32`

GetLicenseStatus returns the LicenseStatus field if non-nil, zero value otherwise.

### GetLicenseStatusOk

`func (o *OswVO) GetLicenseStatusOk() (*int32, bool)`

GetLicenseStatusOk returns a tuple with the LicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatus

`func (o *OswVO) SetLicenseStatus(v int32)`

SetLicenseStatus sets LicenseStatus field to given value.

### HasLicenseStatus

`func (o *OswVO) HasLicenseStatus() bool`

HasLicenseStatus returns a boolean if a field has been set.

### GetLicenseStatusStr

`func (o *OswVO) GetLicenseStatusStr() string`

GetLicenseStatusStr returns the LicenseStatusStr field if non-nil, zero value otherwise.

### GetLicenseStatusStrOk

`func (o *OswVO) GetLicenseStatusStrOk() (*string, bool)`

GetLicenseStatusStrOk returns a tuple with the LicenseStatusStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatusStr

`func (o *OswVO) SetLicenseStatusStr(v string)`

SetLicenseStatusStr sets LicenseStatusStr field to given value.

### HasLicenseStatusStr

`func (o *OswVO) HasLicenseStatusStr() bool`

HasLicenseStatusStr returns a boolean if a field has been set.

### GetLicenseUnbindingLimit

`func (o *OswVO) GetLicenseUnbindingLimit() int32`

GetLicenseUnbindingLimit returns the LicenseUnbindingLimit field if non-nil, zero value otherwise.

### GetLicenseUnbindingLimitOk

`func (o *OswVO) GetLicenseUnbindingLimitOk() (*int32, bool)`

GetLicenseUnbindingLimitOk returns a tuple with the LicenseUnbindingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUnbindingLimit

`func (o *OswVO) SetLicenseUnbindingLimit(v int32)`

SetLicenseUnbindingLimit sets LicenseUnbindingLimit field to given value.

### HasLicenseUnbindingLimit

`func (o *OswVO) HasLicenseUnbindingLimit() bool`

HasLicenseUnbindingLimit returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *OswVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *OswVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *OswVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *OswVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLocateEnable

`func (o *OswVO) GetLocateEnable() bool`

GetLocateEnable returns the LocateEnable field if non-nil, zero value otherwise.

### GetLocateEnableOk

`func (o *OswVO) GetLocateEnableOk() (*bool, bool)`

GetLocateEnableOk returns a tuple with the LocateEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocateEnable

`func (o *OswVO) SetLocateEnable(v bool)`

SetLocateEnable sets LocateEnable field to given value.

### HasLocateEnable

`func (o *OswVO) HasLocateEnable() bool`

HasLocateEnable returns a boolean if a field has been set.

### GetLocatingPorts

`func (o *OswVO) GetLocatingPorts() []int32`

GetLocatingPorts returns the LocatingPorts field if non-nil, zero value otherwise.

### GetLocatingPortsOk

`func (o *OswVO) GetLocatingPortsOk() (*[]int32, bool)`

GetLocatingPortsOk returns a tuple with the LocatingPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatingPorts

`func (o *OswVO) SetLocatingPorts(v []int32)`

SetLocatingPorts sets LocatingPorts field to given value.

### HasLocatingPorts

`func (o *OswVO) HasLocatingPorts() bool`

HasLocatingPorts returns a boolean if a field has been set.

### GetLocatingStandardPorts

`func (o *OswVO) GetLocatingStandardPorts() []string`

GetLocatingStandardPorts returns the LocatingStandardPorts field if non-nil, zero value otherwise.

### GetLocatingStandardPortsOk

`func (o *OswVO) GetLocatingStandardPortsOk() (*[]string, bool)`

GetLocatingStandardPortsOk returns a tuple with the LocatingStandardPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatingStandardPorts

`func (o *OswVO) SetLocatingStandardPorts(v []string)`

SetLocatingStandardPorts sets LocatingStandardPorts field to given value.

### HasLocatingStandardPorts

`func (o *OswVO) HasLocatingStandardPorts() bool`

HasLocatingStandardPorts returns a boolean if a field has been set.

### GetLocation

`func (o *OswVO) GetLocation() LocationVO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OswVO) GetLocationOk() (*LocationVO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OswVO) SetLocation(v LocationVO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OswVO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLoop

`func (o *OswVO) GetLoop() string`

GetLoop returns the Loop field if non-nil, zero value otherwise.

### GetLoopOk

`func (o *OswVO) GetLoopOk() (*string, bool)`

GetLoopOk returns a tuple with the Loop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoop

`func (o *OswVO) SetLoop(v string)`

SetLoop sets Loop field to given value.

### HasLoop

`func (o *OswVO) HasLoop() bool`

HasLoop returns a boolean if a field has been set.

### GetLoopbackNum

`func (o *OswVO) GetLoopbackNum() int32`

GetLoopbackNum returns the LoopbackNum field if non-nil, zero value otherwise.

### GetLoopbackNumOk

`func (o *OswVO) GetLoopbackNumOk() (*int32, bool)`

GetLoopbackNumOk returns a tuple with the LoopbackNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackNum

`func (o *OswVO) SetLoopbackNum(v int32)`

SetLoopbackNum sets LoopbackNum field to given value.

### HasLoopbackNum

`func (o *OswVO) HasLoopbackNum() bool`

HasLoopbackNum returns a boolean if a field has been set.

### GetMac

`func (o *OswVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMaxStackGroups

`func (o *OswVO) GetMaxStackGroups() int32`

GetMaxStackGroups returns the MaxStackGroups field if non-nil, zero value otherwise.

### GetMaxStackGroupsOk

`func (o *OswVO) GetMaxStackGroupsOk() (*int32, bool)`

GetMaxStackGroupsOk returns a tuple with the MaxStackGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStackGroups

`func (o *OswVO) SetMaxStackGroups(v int32)`

SetMaxStackGroups sets MaxStackGroups field to given value.

### HasMaxStackGroups

`func (o *OswVO) HasMaxStackGroups() bool`

HasMaxStackGroups returns a boolean if a field has been set.

### GetMaxStackUnitNumber

`func (o *OswVO) GetMaxStackUnitNumber() int32`

GetMaxStackUnitNumber returns the MaxStackUnitNumber field if non-nil, zero value otherwise.

### GetMaxStackUnitNumberOk

`func (o *OswVO) GetMaxStackUnitNumberOk() (*int32, bool)`

GetMaxStackUnitNumberOk returns a tuple with the MaxStackUnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStackUnitNumber

`func (o *OswVO) SetMaxStackUnitNumber(v int32)`

SetMaxStackUnitNumber sets MaxStackUnitNumber field to given value.

### HasMaxStackUnitNumber

`func (o *OswVO) HasMaxStackUnitNumber() bool`

HasMaxStackUnitNumber returns a boolean if a field has been set.

### GetMemUtil

`func (o *OswVO) GetMemUtil() int32`

GetMemUtil returns the MemUtil field if non-nil, zero value otherwise.

### GetMemUtilOk

`func (o *OswVO) GetMemUtilOk() (*int32, bool)`

GetMemUtilOk returns a tuple with the MemUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUtil

`func (o *OswVO) SetMemUtil(v int32)`

SetMemUtil sets MemUtil field to given value.

### HasMemUtil

`func (o *OswVO) HasMemUtil() bool`

HasMemUtil returns a boolean if a field has been set.

### GetMlagMsg

`func (o *OswVO) GetMlagMsg() MlagMsgVO`

GetMlagMsg returns the MlagMsg field if non-nil, zero value otherwise.

### GetMlagMsgOk

`func (o *OswVO) GetMlagMsgOk() (*MlagMsgVO, bool)`

GetMlagMsgOk returns a tuple with the MlagMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagMsg

`func (o *OswVO) SetMlagMsg(v MlagMsgVO)`

SetMlagMsg sets MlagMsg field to given value.

### HasMlagMsg

`func (o *OswVO) HasMlagMsg() bool`

HasMlagMsg returns a boolean if a field has been set.

### GetModel

`func (o *OswVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OswVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OswVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OswVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OswVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OswVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OswVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OswVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetMstpInsNum

`func (o *OswVO) GetMstpInsNum() int32`

GetMstpInsNum returns the MstpInsNum field if non-nil, zero value otherwise.

### GetMstpInsNumOk

`func (o *OswVO) GetMstpInsNumOk() (*int32, bool)`

GetMstpInsNumOk returns a tuple with the MstpInsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMstpInsNum

`func (o *OswVO) SetMstpInsNum(v int32)`

SetMstpInsNum sets MstpInsNum field to given value.

### HasMstpInsNum

`func (o *OswVO) HasMstpInsNum() bool`

HasMstpInsNum returns a boolean if a field has been set.

### GetName

`func (o *OswVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedUpgrade

`func (o *OswVO) GetNeedUpgrade() bool`

GetNeedUpgrade returns the NeedUpgrade field if non-nil, zero value otherwise.

### GetNeedUpgradeOk

`func (o *OswVO) GetNeedUpgradeOk() (*bool, bool)`

GetNeedUpgradeOk returns a tuple with the NeedUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedUpgrade

`func (o *OswVO) SetNeedUpgrade(v bool)`

SetNeedUpgrade sets NeedUpgrade field to given value.

### HasNeedUpgrade

`func (o *OswVO) HasNeedUpgrade() bool`

HasNeedUpgrade returns a boolean if a field has been set.

### GetOnlineUpgradeStatus

`func (o *OswVO) GetOnlineUpgradeStatus() int32`

GetOnlineUpgradeStatus returns the OnlineUpgradeStatus field if non-nil, zero value otherwise.

### GetOnlineUpgradeStatusOk

`func (o *OswVO) GetOnlineUpgradeStatusOk() (*int32, bool)`

GetOnlineUpgradeStatusOk returns a tuple with the OnlineUpgradeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineUpgradeStatus

`func (o *OswVO) SetOnlineUpgradeStatus(v int32)`

SetOnlineUpgradeStatus sets OnlineUpgradeStatus field to given value.

### HasOnlineUpgradeStatus

`func (o *OswVO) HasOnlineUpgradeStatus() bool`

HasOnlineUpgradeStatus returns a boolean if a field has been set.

### GetPackageCaptureStatus

`func (o *OswVO) GetPackageCaptureStatus() int32`

GetPackageCaptureStatus returns the PackageCaptureStatus field if non-nil, zero value otherwise.

### GetPackageCaptureStatusOk

`func (o *OswVO) GetPackageCaptureStatusOk() (*int32, bool)`

GetPackageCaptureStatusOk returns a tuple with the PackageCaptureStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCaptureStatus

`func (o *OswVO) SetPackageCaptureStatus(v int32)`

SetPackageCaptureStatus sets PackageCaptureStatus field to given value.

### HasPackageCaptureStatus

`func (o *OswVO) HasPackageCaptureStatus() bool`

HasPackageCaptureStatus returns a boolean if a field has been set.

### GetPoeRemain

`func (o *OswVO) GetPoeRemain() float64`

GetPoeRemain returns the PoeRemain field if non-nil, zero value otherwise.

### GetPoeRemainOk

`func (o *OswVO) GetPoeRemainOk() (*float64, bool)`

GetPoeRemainOk returns a tuple with the PoeRemain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeRemain

`func (o *OswVO) SetPoeRemain(v float64)`

SetPoeRemain sets PoeRemain field to given value.

### HasPoeRemain

`func (o *OswVO) HasPoeRemain() bool`

HasPoeRemain returns a boolean if a field has been set.

### GetPoeSupport

`func (o *OswVO) GetPoeSupport() bool`

GetPoeSupport returns the PoeSupport field if non-nil, zero value otherwise.

### GetPoeSupportOk

`func (o *OswVO) GetPoeSupportOk() (*bool, bool)`

GetPoeSupportOk returns a tuple with the PoeSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeSupport

`func (o *OswVO) SetPoeSupport(v bool)`

SetPoeSupport sets PoeSupport field to given value.

### HasPoeSupport

`func (o *OswVO) HasPoeSupport() bool`

HasPoeSupport returns a boolean if a field has been set.

### GetPoeTotalPower

`func (o *OswVO) GetPoeTotalPower() float64`

GetPoeTotalPower returns the PoeTotalPower field if non-nil, zero value otherwise.

### GetPoeTotalPowerOk

`func (o *OswVO) GetPoeTotalPowerOk() (*float64, bool)`

GetPoeTotalPowerOk returns a tuple with the PoeTotalPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeTotalPower

`func (o *OswVO) SetPoeTotalPower(v float64)`

SetPoeTotalPower sets PoeTotalPower field to given value.

### HasPoeTotalPower

`func (o *OswVO) HasPoeTotalPower() bool`

HasPoeTotalPower returns a boolean if a field has been set.

### GetPorts

`func (o *OswVO) GetPorts() []OswStackMemberPortVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OswVO) GetPortsOk() (*[]OswStackMemberPortVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OswVO) SetPorts(v []OswStackMemberPortVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OswVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPowerMode

`func (o *OswVO) GetPowerMode() int32`

GetPowerMode returns the PowerMode field if non-nil, zero value otherwise.

### GetPowerModeOk

`func (o *OswVO) GetPowerModeOk() (*int32, bool)`

GetPowerModeOk returns a tuple with the PowerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerMode

`func (o *OswVO) SetPowerMode(v int32)`

SetPowerMode sets PowerMode field to given value.

### HasPowerMode

`func (o *OswVO) HasPowerMode() bool`

HasPowerMode returns a boolean if a field has been set.

### GetPowerModeList

`func (o *OswVO) GetPowerModeList() []int32`

GetPowerModeList returns the PowerModeList field if non-nil, zero value otherwise.

### GetPowerModeListOk

`func (o *OswVO) GetPowerModeListOk() (*[]int32, bool)`

GetPowerModeListOk returns a tuple with the PowerModeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerModeList

`func (o *OswVO) SetPowerModeList(v []int32)`

SetPowerModeList sets PowerModeList field to given value.

### HasPowerModeList

`func (o *OswVO) HasPowerModeList() bool`

HasPowerModeList returns a boolean if a field has been set.

### GetPreConfigErrorCode

`func (o *OswVO) GetPreConfigErrorCode() int32`

GetPreConfigErrorCode returns the PreConfigErrorCode field if non-nil, zero value otherwise.

### GetPreConfigErrorCodeOk

`func (o *OswVO) GetPreConfigErrorCodeOk() (*int32, bool)`

GetPreConfigErrorCodeOk returns a tuple with the PreConfigErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreConfigErrorCode

`func (o *OswVO) SetPreConfigErrorCode(v int32)`

SetPreConfigErrorCode sets PreConfigErrorCode field to given value.

### HasPreConfigErrorCode

`func (o *OswVO) HasPreConfigErrorCode() bool`

HasPreConfigErrorCode returns a boolean if a field has been set.

### GetPreConfigRetryType

`func (o *OswVO) GetPreConfigRetryType() int32`

GetPreConfigRetryType returns the PreConfigRetryType field if non-nil, zero value otherwise.

### GetPreConfigRetryTypeOk

`func (o *OswVO) GetPreConfigRetryTypeOk() (*int32, bool)`

GetPreConfigRetryTypeOk returns a tuple with the PreConfigRetryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreConfigRetryType

`func (o *OswVO) SetPreConfigRetryType(v int32)`

SetPreConfigRetryType sets PreConfigRetryType field to given value.

### HasPreConfigRetryType

`func (o *OswVO) HasPreConfigRetryType() bool`

HasPreConfigRetryType returns a boolean if a field has been set.

### GetPriority

`func (o *OswVO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *OswVO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *OswVO) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *OswVO) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProfiles

`func (o *OswVO) GetProfiles() map[string]string`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *OswVO) GetProfilesOk() (*map[string]string, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *OswVO) SetProfiles(v map[string]string)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *OswVO) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetPublicIp

`func (o *OswVO) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *OswVO) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *OswVO) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *OswVO) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetResource

`func (o *OswVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *OswVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *OswVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *OswVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSdm

`func (o *OswVO) GetSdm() OswSdmTemplateVO`

GetSdm returns the Sdm field if non-nil, zero value otherwise.

### GetSdmOk

`func (o *OswVO) GetSdmOk() (*OswSdmTemplateVO, bool)`

GetSdmOk returns a tuple with the Sdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdm

`func (o *OswVO) SetSdm(v OswSdmTemplateVO)`

SetSdm sets Sdm field to given value.

### HasSdm

`func (o *OswVO) HasSdm() bool`

HasSdm returns a boolean if a field has been set.

### GetShowModel

`func (o *OswVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *OswVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *OswVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *OswVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetSiteName

`func (o *OswVO) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *OswVO) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *OswVO) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *OswVO) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSn

`func (o *OswVO) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *OswVO) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *OswVO) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *OswVO) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetSpecialModel

`func (o *OswVO) GetSpecialModel() string`

GetSpecialModel returns the SpecialModel field if non-nil, zero value otherwise.

### GetSpecialModelOk

`func (o *OswVO) GetSpecialModelOk() (*string, bool)`

GetSpecialModelOk returns a tuple with the SpecialModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialModel

`func (o *OswVO) SetSpecialModel(v string)`

SetSpecialModel sets SpecialModel field to given value.

### HasSpecialModel

`func (o *OswVO) HasSpecialModel() bool`

HasSpecialModel returns a boolean if a field has been set.

### GetStackMsg

`func (o *OswVO) GetStackMsg() StackMsgVO`

GetStackMsg returns the StackMsg field if non-nil, zero value otherwise.

### GetStackMsgOk

`func (o *OswVO) GetStackMsgOk() (*StackMsgVO, bool)`

GetStackMsgOk returns a tuple with the StackMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackMsg

`func (o *OswVO) SetStackMsg(v StackMsgVO)`

SetStackMsg sets StackMsg field to given value.

### HasStackMsg

`func (o *OswVO) HasStackMsg() bool`

HasStackMsg returns a boolean if a field has been set.

### GetStackPortCap

`func (o *OswVO) GetStackPortCap() map[string][]string`

GetStackPortCap returns the StackPortCap field if non-nil, zero value otherwise.

### GetStackPortCapOk

`func (o *OswVO) GetStackPortCapOk() (*map[string][]string, bool)`

GetStackPortCapOk returns a tuple with the StackPortCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPortCap

`func (o *OswVO) SetStackPortCap(v map[string][]string)`

SetStackPortCap sets StackPortCap field to given value.

### HasStackPortCap

`func (o *OswVO) HasStackPortCap() bool`

HasStackPortCap returns a boolean if a field has been set.

### GetStackPorts

`func (o *OswVO) GetStackPorts() []OswStackPortGroupVO`

GetStackPorts returns the StackPorts field if non-nil, zero value otherwise.

### GetStackPortsOk

`func (o *OswVO) GetStackPortsOk() (*[]OswStackPortGroupVO, bool)`

GetStackPortsOk returns a tuple with the StackPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPorts

`func (o *OswVO) SetStackPorts(v []OswStackPortGroupVO)`

SetStackPorts sets StackPorts field to given value.

### HasStackPorts

`func (o *OswVO) HasStackPorts() bool`

HasStackPorts returns a boolean if a field has been set.

### GetStackSupportPorts

`func (o *OswVO) GetStackSupportPorts() []OswStandPortVO`

GetStackSupportPorts returns the StackSupportPorts field if non-nil, zero value otherwise.

### GetStackSupportPortsOk

`func (o *OswVO) GetStackSupportPortsOk() (*[]OswStandPortVO, bool)`

GetStackSupportPortsOk returns a tuple with the StackSupportPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackSupportPorts

`func (o *OswVO) SetStackSupportPorts(v []OswStandPortVO)`

SetStackSupportPorts sets StackSupportPorts field to given value.

### HasStackSupportPorts

`func (o *OswVO) HasStackSupportPorts() bool`

HasStackSupportPorts returns a boolean if a field has been set.

### GetStatus

`func (o *OswVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OswVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *OswVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *OswVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *OswVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *OswVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetStkVer

`func (o *OswVO) GetStkVer() string`

GetStkVer returns the StkVer field if non-nil, zero value otherwise.

### GetStkVerOk

`func (o *OswVO) GetStkVerOk() (*string, bool)`

GetStkVerOk returns a tuple with the StkVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStkVer

`func (o *OswVO) SetStkVer(v string)`

SetStkVer sets StkVer field to given value.

### HasStkVer

`func (o *OswVO) HasStkVer() bool`

HasStkVer returns a boolean if a field has been set.

### GetStkableGroupId

`func (o *OswVO) GetStkableGroupId() int32`

GetStkableGroupId returns the StkableGroupId field if non-nil, zero value otherwise.

### GetStkableGroupIdOk

`func (o *OswVO) GetStkableGroupIdOk() (*int32, bool)`

GetStkableGroupIdOk returns a tuple with the StkableGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStkableGroupId

`func (o *OswVO) SetStkableGroupId(v int32)`

SetStkableGroupId sets StkableGroupId field to given value.

### HasStkableGroupId

`func (o *OswVO) HasStkableGroupId() bool`

HasStkableGroupId returns a boolean if a field has been set.

### GetSubDevType

`func (o *OswVO) GetSubDevType() int32`

GetSubDevType returns the SubDevType field if non-nil, zero value otherwise.

### GetSubDevTypeOk

`func (o *OswVO) GetSubDevTypeOk() (*int32, bool)`

GetSubDevTypeOk returns a tuple with the SubDevType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubDevType

`func (o *OswVO) SetSubDevType(v int32)`

SetSubDevType sets SubDevType field to given value.

### HasSubDevType

`func (o *OswVO) HasSubDevType() bool`

HasSubDevType returns a boolean if a field has been set.

### GetSupportAnomaly

`func (o *OswVO) GetSupportAnomaly() bool`

GetSupportAnomaly returns the SupportAnomaly field if non-nil, zero value otherwise.

### GetSupportAnomalyOk

`func (o *OswVO) GetSupportAnomalyOk() (*bool, bool)`

GetSupportAnomalyOk returns a tuple with the SupportAnomaly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAnomaly

`func (o *OswVO) SetSupportAnomaly(v bool)`

SetSupportAnomaly sets SupportAnomaly field to given value.

### HasSupportAnomaly

`func (o *OswVO) HasSupportAnomaly() bool`

HasSupportAnomaly returns a boolean if a field has been set.

### GetSupportCableTest

`func (o *OswVO) GetSupportCableTest() bool`

GetSupportCableTest returns the SupportCableTest field if non-nil, zero value otherwise.

### GetSupportCableTestOk

`func (o *OswVO) GetSupportCableTestOk() (*bool, bool)`

GetSupportCableTestOk returns a tuple with the SupportCableTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCableTest

`func (o *OswVO) SetSupportCableTest(v bool)`

SetSupportCableTest sets SupportCableTest field to given value.

### HasSupportCableTest

`func (o *OswVO) HasSupportCableTest() bool`

HasSupportCableTest returns a boolean if a field has been set.

### GetSupportCustomDhcpOption

`func (o *OswVO) GetSupportCustomDhcpOption() bool`

GetSupportCustomDhcpOption returns the SupportCustomDhcpOption field if non-nil, zero value otherwise.

### GetSupportCustomDhcpOptionOk

`func (o *OswVO) GetSupportCustomDhcpOptionOk() (*bool, bool)`

GetSupportCustomDhcpOptionOk returns a tuple with the SupportCustomDhcpOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCustomDhcpOption

`func (o *OswVO) SetSupportCustomDhcpOption(v bool)`

SetSupportCustomDhcpOption sets SupportCustomDhcpOption field to given value.

### HasSupportCustomDhcpOption

`func (o *OswVO) HasSupportCustomDhcpOption() bool`

HasSupportCustomDhcpOption returns a boolean if a field has been set.

### GetSupportDhcpRange

`func (o *OswVO) GetSupportDhcpRange() bool`

GetSupportDhcpRange returns the SupportDhcpRange field if non-nil, zero value otherwise.

### GetSupportDhcpRangeOk

`func (o *OswVO) GetSupportDhcpRangeOk() (*bool, bool)`

GetSupportDhcpRangeOk returns a tuple with the SupportDhcpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDhcpRange

`func (o *OswVO) SetSupportDhcpRange(v bool)`

SetSupportDhcpRange sets SupportDhcpRange field to given value.

### HasSupportDhcpRange

`func (o *OswVO) HasSupportDhcpRange() bool`

HasSupportDhcpRange returns a boolean if a field has been set.

### GetSupportDhcpReservation

`func (o *OswVO) GetSupportDhcpReservation() bool`

GetSupportDhcpReservation returns the SupportDhcpReservation field if non-nil, zero value otherwise.

### GetSupportDhcpReservationOk

`func (o *OswVO) GetSupportDhcpReservationOk() (*bool, bool)`

GetSupportDhcpReservationOk returns a tuple with the SupportDhcpReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDhcpReservation

`func (o *OswVO) SetSupportDhcpReservation(v bool)`

SetSupportDhcpReservation sets SupportDhcpReservation field to given value.

### HasSupportDhcpReservation

`func (o *OswVO) HasSupportDhcpReservation() bool`

HasSupportDhcpReservation returns a boolean if a field has been set.

### GetSupportDomainPing

`func (o *OswVO) GetSupportDomainPing() bool`

GetSupportDomainPing returns the SupportDomainPing field if non-nil, zero value otherwise.

### GetSupportDomainPingOk

`func (o *OswVO) GetSupportDomainPingOk() (*bool, bool)`

GetSupportDomainPingOk returns a tuple with the SupportDomainPing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDomainPing

`func (o *OswVO) SetSupportDomainPing(v bool)`

SetSupportDomainPing sets SupportDomainPing field to given value.

### HasSupportDomainPing

`func (o *OswVO) HasSupportDomainPing() bool`

HasSupportDomainPing returns a boolean if a field has been set.

### GetSupportDomainTraceRoute

`func (o *OswVO) GetSupportDomainTraceRoute() bool`

GetSupportDomainTraceRoute returns the SupportDomainTraceRoute field if non-nil, zero value otherwise.

### GetSupportDomainTraceRouteOk

`func (o *OswVO) GetSupportDomainTraceRouteOk() (*bool, bool)`

GetSupportDomainTraceRouteOk returns a tuple with the SupportDomainTraceRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDomainTraceRoute

`func (o *OswVO) SetSupportDomainTraceRoute(v bool)`

SetSupportDomainTraceRoute sets SupportDomainTraceRoute field to given value.

### HasSupportDomainTraceRoute

`func (o *OswVO) HasSupportDomainTraceRoute() bool`

HasSupportDomainTraceRoute returns a boolean if a field has been set.

### GetSupportExtendStp

`func (o *OswVO) GetSupportExtendStp() bool`

GetSupportExtendStp returns the SupportExtendStp field if non-nil, zero value otherwise.

### GetSupportExtendStpOk

`func (o *OswVO) GetSupportExtendStpOk() (*bool, bool)`

GetSupportExtendStpOk returns a tuple with the SupportExtendStp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportExtendStp

`func (o *OswVO) SetSupportExtendStp(v bool)`

SetSupportExtendStp sets SupportExtendStp field to given value.

### HasSupportExtendStp

`func (o *OswVO) HasSupportExtendStp() bool`

HasSupportExtendStp returns a boolean if a field has been set.

### GetSupportGetOspfNeighborTable

`func (o *OswVO) GetSupportGetOspfNeighborTable() bool`

GetSupportGetOspfNeighborTable returns the SupportGetOspfNeighborTable field if non-nil, zero value otherwise.

### GetSupportGetOspfNeighborTableOk

`func (o *OswVO) GetSupportGetOspfNeighborTableOk() (*bool, bool)`

GetSupportGetOspfNeighborTableOk returns a tuple with the SupportGetOspfNeighborTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportGetOspfNeighborTable

`func (o *OswVO) SetSupportGetOspfNeighborTable(v bool)`

SetSupportGetOspfNeighborTable sets SupportGetOspfNeighborTable field to given value.

### HasSupportGetOspfNeighborTable

`func (o *OswVO) HasSupportGetOspfNeighborTable() bool`

HasSupportGetOspfNeighborTable returns a boolean if a field has been set.

### GetSupportIppt

`func (o *OswVO) GetSupportIppt() bool`

GetSupportIppt returns the SupportIppt field if non-nil, zero value otherwise.

### GetSupportIpptOk

`func (o *OswVO) GetSupportIpptOk() (*bool, bool)`

GetSupportIpptOk returns a tuple with the SupportIppt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIppt

`func (o *OswVO) SetSupportIppt(v bool)`

SetSupportIppt sets SupportIppt field to given value.

### HasSupportIppt

`func (o *OswVO) HasSupportIppt() bool`

HasSupportIppt returns a boolean if a field has been set.

### GetSupportMacDelay

`func (o *OswVO) GetSupportMacDelay() bool`

GetSupportMacDelay returns the SupportMacDelay field if non-nil, zero value otherwise.

### GetSupportMacDelayOk

`func (o *OswVO) GetSupportMacDelayOk() (*bool, bool)`

GetSupportMacDelayOk returns a tuple with the SupportMacDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMacDelay

`func (o *OswVO) SetSupportMacDelay(v bool)`

SetSupportMacDelay sets SupportMacDelay field to given value.

### HasSupportMacDelay

`func (o *OswVO) HasSupportMacDelay() bool`

HasSupportMacDelay returns a boolean if a field has been set.

### GetSupportPowerAlert

`func (o *OswVO) GetSupportPowerAlert() bool`

GetSupportPowerAlert returns the SupportPowerAlert field if non-nil, zero value otherwise.

### GetSupportPowerAlertOk

`func (o *OswVO) GetSupportPowerAlertOk() (*bool, bool)`

GetSupportPowerAlertOk returns a tuple with the SupportPowerAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPowerAlert

`func (o *OswVO) SetSupportPowerAlert(v bool)`

SetSupportPowerAlert sets SupportPowerAlert field to given value.

### HasSupportPowerAlert

`func (o *OswVO) HasSupportPowerAlert() bool`

HasSupportPowerAlert returns a boolean if a field has been set.

### GetSupportRunningConfig

`func (o *OswVO) GetSupportRunningConfig() bool`

GetSupportRunningConfig returns the SupportRunningConfig field if non-nil, zero value otherwise.

### GetSupportRunningConfigOk

`func (o *OswVO) GetSupportRunningConfigOk() (*bool, bool)`

GetSupportRunningConfigOk returns a tuple with the SupportRunningConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportRunningConfig

`func (o *OswVO) SetSupportRunningConfig(v bool)`

SetSupportRunningConfig sets SupportRunningConfig field to given value.

### HasSupportRunningConfig

`func (o *OswVO) HasSupportRunningConfig() bool`

HasSupportRunningConfig returns a boolean if a field has been set.

### GetSupportSdm

`func (o *OswVO) GetSupportSdm() bool`

GetSupportSdm returns the SupportSdm field if non-nil, zero value otherwise.

### GetSupportSdmOk

`func (o *OswVO) GetSupportSdmOk() (*bool, bool)`

GetSupportSdmOk returns a tuple with the SupportSdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSdm

`func (o *OswVO) SetSupportSdm(v bool)`

SetSupportSdm sets SupportSdm field to given value.

### HasSupportSdm

`func (o *OswVO) HasSupportSdm() bool`

HasSupportSdm returns a boolean if a field has been set.

### GetSupportStp

`func (o *OswVO) GetSupportStp() bool`

GetSupportStp returns the SupportStp field if non-nil, zero value otherwise.

### GetSupportStpOk

`func (o *OswVO) GetSupportStpOk() (*bool, bool)`

GetSupportStpOk returns a tuple with the SupportStp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportStp

`func (o *OswVO) SetSupportStp(v bool)`

SetSupportStp sets SupportStp field to given value.

### HasSupportStp

`func (o *OswVO) HasSupportStp() bool`

HasSupportStp returns a boolean if a field has been set.

### GetSupportVrf

`func (o *OswVO) GetSupportVrf() bool`

GetSupportVrf returns the SupportVrf field if non-nil, zero value otherwise.

### GetSupportVrfOk

`func (o *OswVO) GetSupportVrfOk() (*bool, bool)`

GetSupportVrfOk returns a tuple with the SupportVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVrf

`func (o *OswVO) SetSupportVrf(v bool)`

SetSupportVrf sets SupportVrf field to given value.

### HasSupportVrf

`func (o *OswVO) HasSupportVrf() bool`

HasSupportVrf returns a boolean if a field has been set.

### GetSwitchConsistent

`func (o *OswVO) GetSwitchConsistent() bool`

GetSwitchConsistent returns the SwitchConsistent field if non-nil, zero value otherwise.

### GetSwitchConsistentOk

`func (o *OswVO) GetSwitchConsistentOk() (*bool, bool)`

GetSwitchConsistentOk returns a tuple with the SwitchConsistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchConsistent

`func (o *OswVO) SetSwitchConsistent(v bool)`

SetSwitchConsistent sets SwitchConsistent field to given value.

### HasSwitchConsistent

`func (o *OswVO) HasSwitchConsistent() bool`

HasSwitchConsistent returns a boolean if a field has been set.

### GetTagName

`func (o *OswVO) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *OswVO) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *OswVO) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *OswVO) HasTagName() bool`

HasTagName returns a boolean if a field has been set.

### GetType

`func (o *OswVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OswVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *OswVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OswVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OswVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *OswVO) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUplinkDeviceMac

`func (o *OswVO) GetUplinkDeviceMac() string`

GetUplinkDeviceMac returns the UplinkDeviceMac field if non-nil, zero value otherwise.

### GetUplinkDeviceMacOk

`func (o *OswVO) GetUplinkDeviceMacOk() (*string, bool)`

GetUplinkDeviceMacOk returns a tuple with the UplinkDeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkDeviceMac

`func (o *OswVO) SetUplinkDeviceMac(v string)`

SetUplinkDeviceMac sets UplinkDeviceMac field to given value.

### HasUplinkDeviceMac

`func (o *OswVO) HasUplinkDeviceMac() bool`

HasUplinkDeviceMac returns a boolean if a field has been set.

### GetUplinkDeviceName

`func (o *OswVO) GetUplinkDeviceName() string`

GetUplinkDeviceName returns the UplinkDeviceName field if non-nil, zero value otherwise.

### GetUplinkDeviceNameOk

`func (o *OswVO) GetUplinkDeviceNameOk() (*string, bool)`

GetUplinkDeviceNameOk returns a tuple with the UplinkDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkDeviceName

`func (o *OswVO) SetUplinkDeviceName(v string)`

SetUplinkDeviceName sets UplinkDeviceName field to given value.

### HasUplinkDeviceName

`func (o *OswVO) HasUplinkDeviceName() bool`

HasUplinkDeviceName returns a boolean if a field has been set.

### GetUplinkDevicePort

`func (o *OswVO) GetUplinkDevicePort() string`

GetUplinkDevicePort returns the UplinkDevicePort field if non-nil, zero value otherwise.

### GetUplinkDevicePortOk

`func (o *OswVO) GetUplinkDevicePortOk() (*string, bool)`

GetUplinkDevicePortOk returns a tuple with the UplinkDevicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkDevicePort

`func (o *OswVO) SetUplinkDevicePort(v string)`

SetUplinkDevicePort sets UplinkDevicePort field to given value.

### HasUplinkDevicePort

`func (o *OswVO) HasUplinkDevicePort() bool`

HasUplinkDevicePort returns a boolean if a field has been set.

### GetUpload

`func (o *OswVO) GetUpload() int64`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *OswVO) GetUploadOk() (*int64, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *OswVO) SetUpload(v int64)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *OswVO) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetUptime

`func (o *OswVO) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *OswVO) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *OswVO) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *OswVO) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetUptimeLong

`func (o *OswVO) GetUptimeLong() int64`

GetUptimeLong returns the UptimeLong field if non-nil, zero value otherwise.

### GetUptimeLongOk

`func (o *OswVO) GetUptimeLongOk() (*int64, bool)`

GetUptimeLongOk returns a tuple with the UptimeLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeLong

`func (o *OswVO) SetUptimeLong(v int64)`

SetUptimeLong sets UptimeLong field to given value.

### HasUptimeLong

`func (o *OswVO) HasUptimeLong() bool`

HasUptimeLong returns a boolean if a field has been set.

### GetVersion

`func (o *OswVO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OswVO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OswVO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OswVO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWirelessRouter

`func (o *OswVO) GetWirelessRouter() bool`

GetWirelessRouter returns the WirelessRouter field if non-nil, zero value otherwise.

### GetWirelessRouterOk

`func (o *OswVO) GetWirelessRouterOk() (*bool, bool)`

GetWirelessRouterOk returns a tuple with the WirelessRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessRouter

`func (o *OswVO) SetWirelessRouter(v bool)`

SetWirelessRouter sets WirelessRouter field to given value.

### HasWirelessRouter

`func (o *OswVO) HasWirelessRouter() bool`

HasWirelessRouter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


