# OswDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | Site of the device | [optional] 
**Active** | Pointer to **bool** | whether to active the device(cloud base exclusive) | [optional] 
**AddedInAdvanced** | Pointer to **bool** | Whether the device is added in advanced. | [optional] 
**BoundDeviceTemplate** | Pointer to **bool** | Whether the device is bound to device template | [optional] 
**BoundSiteTemplate** | Pointer to **bool** | Whether the site where the device is located is bound to a site template | [optional] 
**Category** | Pointer to **string** | Category of license | [optional] 
**Compatible** | Pointer to **int32** | Device firmware and controller compatibility type.Compatible should be a value as follows: 0:COMPATIBLE;1:HIGH_MAJOR_VER;2:LOW_MAJOR_VER;3:HIGH_MINOR_VER;4:LOW_MINOR_VER;7:HIGH_COMPONENT_VER;10:DEVICE_NOT_COMPATIBLE;11:HIGH_ADOPT_COMMPONENT;12:DEVICE_CATEGORY_NOT_COMPATIBLE;14:DEVICE_NOT_COMPATIBLE_IN_CLUSTER | [optional] 
**CompoundModel** | Pointer to **string** | Model complex used in the backend.Ap：model+(country)+modelVersion,  EAP225(EU) v3.0 Ap: specialModel+modelVersion, EAP225-Outdoor-1a20a950b8d950e8 v1.0  Gateway/Switch：model+modelVersion, Osg v3.0 | [optional] 
**CpuUtil** | Pointer to **int32** | Real-time CPU usage | [optional] 
**CustomId** | Pointer to **string** | Customer ID | [optional] 
**CustomName** | Pointer to **string** | Customer name | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DevCap** | Pointer to [**OswDevCapVO**](OswDevCapVO.md) |  | [optional] 
**DeviceMisc** | Pointer to [**OswDeviceMiscVO**](OswDeviceMiscVO.md) |  | [optional] 
**DeviceSeriesType** | Pointer to **int32** | Device series type.DeviceSeriesType should be a value as follows: 0:advanced;1:pro | [optional] 
**DeviceTemplateAvailable** | Pointer to **bool** | Whether there is an available device template for the device; it is false if the model is not supported or the site template has not created the corresponding device template. | [optional] 
**DisableHwReset** | Pointer to **bool** |  | [optional] 
**DownlinkList** | Pointer to [**[]OswDownlinkVO**](OswDownlinkVO.md) | Downlink Omada device list | [optional] 
**Download** | Pointer to **int64** | Total Download (Byte) | [optional] 
**DueTime** | Pointer to **int64** | Expire timestamp of license(cloud base exclusive) | [optional] 
**DueTimeLeft** | Pointer to **int64** | Milliseconds from the current moment to the expiration time(cloud base exclusive) | [optional] 
**EcspFirstVersion** | Pointer to **int32** | Ecsp first version | [optional] 
**Eos** | Pointer to **int64** | End of support time of device(CBC exclusive) | [optional] 
**Eost** | Pointer to **int64** | End of service time of device(CBC exclusive) | [optional] 
**Es** | Pointer to **bool** | Whether the switch is Agile Series Switch | [optional] 
**FanStatus** | Pointer to **int32** | Fan Status should be a value as follows: 0:normal; 1:fault; 2:no fan | [optional] 
**FirmwareVersion** | Pointer to **string** | Version of firmware,for example:2.5.0 Build 20190118 Rel. 64821 | [optional] 
**ForgetId** | Pointer to **string** | Forget ID of device | [optional] 
**ForwardDelay** | Pointer to **int32** | STP forwardDelay | [optional] 
**HelloTime** | Pointer to **int32** | STP helloTime | [optional] 
**HwVersion** | Pointer to **string** | Version of hardware,for example 1.0 | [optional] 
**InWhitelist** | Pointer to **bool** | Whether the device is in white list | [optional] 
**InitialUnbindingLimit** | Pointer to **int32** | Initial unbind count for license(cloud base exclusive) | [optional] 
**Ip** | Pointer to **string** | Switch IP | [optional] 
**IpSetting** | Pointer to [**IpSettingVO**](IpSettingVO.md) |  | [optional] 
**Ipv6List** | Pointer to **[]string** | IPV6 List | [optional] 
**Jumbo** | Pointer to **int32** | Jumbo should be within the range of 1518-9216. | [optional] 
**LagHashAlg** | Pointer to **int32** | It should be a value as follows: 0: SRC MAC; 1: DST MAC; 2: SRC MAC + DST MAC; 3: SRC IP; 4: DST IP; 5: SRC IP + DST IP | [optional] 
**Lags** | Pointer to [**[]OswLagVO**](OswLagVO.md) | Lag List | [optional] 
**LastSeen** | Pointer to **int64** | Last Seen | [optional] 
**LedSetting** | Pointer to **int32** | LedSetting should be a value as follows: 0: Off; 1: On; 2:follow site | [optional] 
**LicenseId** | Pointer to **string** | License key on detail page of device(cloud base exclusive) | [optional] 
**LicenseStatus** | Pointer to **int32** | License status(cloud base exclusive).LicenseStatus should be a value as follows: 0:unActive 1:Unbind 2:Expired 3:active | [optional] 
**LicenseUnbindingLimit** | Pointer to **int32** | Remaining unbind count for license on detail Page of device(cloud base exclusive) | [optional] 
**Location** | Pointer to [**DeviceLocationDetailVO**](DeviceLocationDetailVO.md) |  | [optional] 
**Loop** | Pointer to **string** | Loop information | [optional] 
**LoopbackDetectEnable** | Pointer to **bool** | Loopback Detect Enable | [optional] 
**LoopbackNum** | Pointer to **int32** | Loopback Num | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**MaxAge** | Pointer to **int32** | STP maxAge | [optional] 
**MaxHops** | Pointer to **int32** | STP maxHops | [optional] 
**MemUitl** | Pointer to **int32** | Real-time memory usage | [optional] 
**MlagMsg** | Pointer to [**MlagMsgVO**](MlagMsgVO.md) |  | [optional] 
**MlagPeerInfo** | Pointer to [**OswMlagPeerInfoVO**](OswMlagPeerInfoVO.md) |  | [optional] 
**Model** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**MoveSiteId** | Pointer to **string** | Record that the device is in a moveSite operation; if it is null, then it is not in the moveSite operation. | [optional] 
**Mstp** | Pointer to [**OswStpMstpVO**](OswStpMstpVO.md) |  | [optional] 
**Multicast** | Pointer to [**OswLanMulticastVO**](OswLanMulticastVO.md) |  | [optional] 
**MvlanBridgeVlan** | Pointer to **int32** | Only valid when mvlanNetworkId is bridge vlan | [optional] 
**MvlanNetworkId** | Pointer to **string** | Management VLAN network ID | [optional] 
**Name** | Pointer to **string** | Device name,default value is the mac address of device | [optional] 
**NeedUpgrade** | Pointer to **bool** | Need Upgrade | [optional] 
**OmadacId** | Pointer to **string** | OmadacId of the device | [optional] 
**PoeRemain** | Pointer to **float64** | PoE Residual Power (W) | [optional] 
**PoeRemainPercent** | Pointer to **float64** | PoE Residual Power Percentage | [optional] 
**PoeTotalPower** | Pointer to **float64** | PoE Total Power (W) | [optional] 
**Ports** | Pointer to [**[]OswPortVO**](OswPortVO.md) | Port List | [optional] 
**Priority** | Pointer to **int32** | STP priority | [optional] 
**PublicIp** | Pointer to **string** | Public IP | [optional] 
**QosConfig** | Pointer to [**OswQosConfigVO**](OswQosConfigVO.md) |  | [optional] 
**Remember** | Pointer to **bool** | Whether to remember the device(deprecated) | [optional] 
**RememberDevice** | Pointer to **int32** | Whether to remember the device.RememberDevice should be a value as follows: 0:off, 1:on, 2: follow site | [optional] 
**Resource** | Pointer to **int32** | Data source.Resource should be a value as follows: 0:new created;1:from template;2:override | [optional] 
**RxRate** | Pointer to **int64** | Rx Rate | [optional] 
**Sdm** | Pointer to [**OswSdmTemplateVO**](OswSdmTemplateVO.md) |  | [optional] 
**ShowModel** | Pointer to **string** | Model complex shown in the front end.Ap：model+(country)+modelVersion,EAP225(EU) v3.0  Gateway/Switch：model+modelVersion,Osg v3.0 | [optional] 
**SiteName** | Pointer to **string** | Site name of the device | [optional] 
**SiteTemplateId** | Pointer to **string** | Template ID bound to the site | [optional] 
**SiteTemplateName** | Pointer to **string** | Template name bound to the site | [optional] 
**Sn** | Pointer to **string** | Device serial number | [optional] 
**Snmp** | Pointer to [**OswSnmpVO**](OswSnmpVO.md) |  | [optional] 
**SpecialModel** | Pointer to **string** | Special device model,for example:EAP225-Outdoor-1a20a950b8d950e8 | [optional] 
**Speeds** | Pointer to **[]int32** | Supported rate list for all ports. Speeds should be a value as follows: 0:auto; 1:10M; 2:100M; 3:1000M; 4:2.5G; 5:10G; 6:5G; 7:25G; 8:100G; 9:40G; -1:error; no value:all rate supported | [optional] 
**StackDevice** | Pointer to **bool** | Stack Device | [optional] 
**StackMsg** | Pointer to [**StackMsgVO**](StackMsgVO.md) |  | [optional] 
**StackPorts** | Pointer to [**[]OswStackPortGroupVO**](OswStackPortGroupVO.md) | Stack Port List | [optional] 
**Status** | Pointer to **int32** | Status of device,status should be a value as follows: 0:Disconnected;1:Disconnected(Migrating);10:Provisioning;11:Configuring;12:Upgrading;13:Rebooting;14:Connected;15:Connected(Wireless);16:Connected(Migrating);17:Connected(Wireless,Migrating);20:Pending;21:Pending(Wireless);22:Adopting;23:Adopting(Wireless);24:Adopt Failed;25:Adopt Failed(Wireless);26:Managed By Others;27:Managed By Others(Wireless);30:Heartbeat Missed;31:Heartbeat Missed(Wireless);32:Heartbeat Missed(Migrating);33:Heartbeat Missed(Wireless,Migrating);40:Isolated;41:Isolated(Migrating);50:Slice Configuring | [optional] 
**StatusCategory** | Pointer to **int32** | Category of device status,statusCategory should be a value as follows: 0:Disconnected;1:Connected;2:Pending;3:Heartbeat Missed;4:Isolated | [optional] 
**Stp** | Pointer to **int32** | Spanning Tree Protocol should be a value as follows: 1: STP; 2: RSTP; 3: MSTP; 0: OFF | [optional] 
**StpLinkList** | Pointer to [**[]OswDownlinkVO**](OswDownlinkVO.md) | STP Blocked Link Omada device list | [optional] 
**SupportAnomaly** | Pointer to **bool** | Whether the device firmware support intelligent anomaly detection | [optional] 
**SupportHealth** | Pointer to **bool** | Support health | [optional] 
**SupportIpv6Acl** | Pointer to **bool** | Support Ipv6 Acl | [optional] 
**SupportLocatePort** | Pointer to **bool** | Whether the device supports locating port | [optional] 
**SupportVlanIf** | Pointer to **bool** | Support Vlan Interface | [optional] 
**SupportVrf** | Pointer to **bool** | Support Vrf | [optional] 
**TagIds** | Pointer to **[]string** | The bound TAG ID List | [optional] 
**TemplateId** | Pointer to **string** | ID of the template bound to the device | [optional] 
**TemplateName** | Pointer to **string** | Name of the template bound to the device | [optional] 
**TemplateSettings** | Pointer to **[]int32** | Template Setting List | [optional] 
**TerminalPrefix** | Pointer to **string** | TerminalPrefix represents the device name within the terminal function, designed to prevent terminal command recognition errors when device name contains illegal characters such as &#39;#&#39;. | [optional] 
**TxHoldCount** | Pointer to **int32** | STP txHoldCount | [optional] 
**TxRate** | Pointer to **int64** | Tx Rate | [optional] 
**Type** | Pointer to **string** | Device type:ap、gateway、switch、olt | [optional] 
**Unit** | Pointer to **int32** | Unit ID | [optional] 
**Uplink** | Pointer to [**OswUplinkVO**](OswUplinkVO.md) |  | [optional] 
**UplinkPort** | Pointer to **int32** | Uplink port | [optional] 
**UplinkStPort** | Pointer to **string** | Uplink standard port,unit/slot/port | [optional] 
**Upload** | Pointer to **int64** | Total Upload (Byte) | [optional] 
**Uptime** | Pointer to **string** | Uptime | [optional] 
**UptimeLong** | Pointer to **int64** | UptimeLong, Running duration, Units: seconds | [optional] 
**Version** | Pointer to **string** | Simplified version of firmware,for example:2.5.0 | [optional] 

## Methods

### NewOswDetailVO

`func NewOswDetailVO() *OswDetailVO`

NewOswDetailVO instantiates a new OswDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswDetailVOWithDefaults

`func NewOswDetailVOWithDefaults() *OswDetailVO`

NewOswDetailVOWithDefaults instantiates a new OswDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *OswDetailVO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *OswDetailVO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *OswDetailVO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *OswDetailVO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetActive

`func (o *OswDetailVO) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OswDetailVO) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OswDetailVO) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *OswDetailVO) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddedInAdvanced

`func (o *OswDetailVO) GetAddedInAdvanced() bool`

GetAddedInAdvanced returns the AddedInAdvanced field if non-nil, zero value otherwise.

### GetAddedInAdvancedOk

`func (o *OswDetailVO) GetAddedInAdvancedOk() (*bool, bool)`

GetAddedInAdvancedOk returns a tuple with the AddedInAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedInAdvanced

`func (o *OswDetailVO) SetAddedInAdvanced(v bool)`

SetAddedInAdvanced sets AddedInAdvanced field to given value.

### HasAddedInAdvanced

`func (o *OswDetailVO) HasAddedInAdvanced() bool`

HasAddedInAdvanced returns a boolean if a field has been set.

### GetBoundDeviceTemplate

`func (o *OswDetailVO) GetBoundDeviceTemplate() bool`

GetBoundDeviceTemplate returns the BoundDeviceTemplate field if non-nil, zero value otherwise.

### GetBoundDeviceTemplateOk

`func (o *OswDetailVO) GetBoundDeviceTemplateOk() (*bool, bool)`

GetBoundDeviceTemplateOk returns a tuple with the BoundDeviceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundDeviceTemplate

`func (o *OswDetailVO) SetBoundDeviceTemplate(v bool)`

SetBoundDeviceTemplate sets BoundDeviceTemplate field to given value.

### HasBoundDeviceTemplate

`func (o *OswDetailVO) HasBoundDeviceTemplate() bool`

HasBoundDeviceTemplate returns a boolean if a field has been set.

### GetBoundSiteTemplate

`func (o *OswDetailVO) GetBoundSiteTemplate() bool`

GetBoundSiteTemplate returns the BoundSiteTemplate field if non-nil, zero value otherwise.

### GetBoundSiteTemplateOk

`func (o *OswDetailVO) GetBoundSiteTemplateOk() (*bool, bool)`

GetBoundSiteTemplateOk returns a tuple with the BoundSiteTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSiteTemplate

`func (o *OswDetailVO) SetBoundSiteTemplate(v bool)`

SetBoundSiteTemplate sets BoundSiteTemplate field to given value.

### HasBoundSiteTemplate

`func (o *OswDetailVO) HasBoundSiteTemplate() bool`

HasBoundSiteTemplate returns a boolean if a field has been set.

### GetCategory

`func (o *OswDetailVO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *OswDetailVO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *OswDetailVO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *OswDetailVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCompatible

`func (o *OswDetailVO) GetCompatible() int32`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *OswDetailVO) GetCompatibleOk() (*int32, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *OswDetailVO) SetCompatible(v int32)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *OswDetailVO) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetCompoundModel

`func (o *OswDetailVO) GetCompoundModel() string`

GetCompoundModel returns the CompoundModel field if non-nil, zero value otherwise.

### GetCompoundModelOk

`func (o *OswDetailVO) GetCompoundModelOk() (*string, bool)`

GetCompoundModelOk returns a tuple with the CompoundModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompoundModel

`func (o *OswDetailVO) SetCompoundModel(v string)`

SetCompoundModel sets CompoundModel field to given value.

### HasCompoundModel

`func (o *OswDetailVO) HasCompoundModel() bool`

HasCompoundModel returns a boolean if a field has been set.

### GetCpuUtil

`func (o *OswDetailVO) GetCpuUtil() int32`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *OswDetailVO) GetCpuUtilOk() (*int32, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *OswDetailVO) SetCpuUtil(v int32)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *OswDetailVO) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetCustomId

`func (o *OswDetailVO) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *OswDetailVO) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *OswDetailVO) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.

### HasCustomId

`func (o *OswDetailVO) HasCustomId() bool`

HasCustomId returns a boolean if a field has been set.

### GetCustomName

`func (o *OswDetailVO) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *OswDetailVO) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *OswDetailVO) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *OswDetailVO) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### GetDescription

`func (o *OswDetailVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OswDetailVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OswDetailVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OswDetailVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevCap

`func (o *OswDetailVO) GetDevCap() OswDevCapVO`

GetDevCap returns the DevCap field if non-nil, zero value otherwise.

### GetDevCapOk

`func (o *OswDetailVO) GetDevCapOk() (*OswDevCapVO, bool)`

GetDevCapOk returns a tuple with the DevCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevCap

`func (o *OswDetailVO) SetDevCap(v OswDevCapVO)`

SetDevCap sets DevCap field to given value.

### HasDevCap

`func (o *OswDetailVO) HasDevCap() bool`

HasDevCap returns a boolean if a field has been set.

### GetDeviceMisc

`func (o *OswDetailVO) GetDeviceMisc() OswDeviceMiscVO`

GetDeviceMisc returns the DeviceMisc field if non-nil, zero value otherwise.

### GetDeviceMiscOk

`func (o *OswDetailVO) GetDeviceMiscOk() (*OswDeviceMiscVO, bool)`

GetDeviceMiscOk returns a tuple with the DeviceMisc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMisc

`func (o *OswDetailVO) SetDeviceMisc(v OswDeviceMiscVO)`

SetDeviceMisc sets DeviceMisc field to given value.

### HasDeviceMisc

`func (o *OswDetailVO) HasDeviceMisc() bool`

HasDeviceMisc returns a boolean if a field has been set.

### GetDeviceSeriesType

`func (o *OswDetailVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *OswDetailVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *OswDetailVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *OswDetailVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetDeviceTemplateAvailable

`func (o *OswDetailVO) GetDeviceTemplateAvailable() bool`

GetDeviceTemplateAvailable returns the DeviceTemplateAvailable field if non-nil, zero value otherwise.

### GetDeviceTemplateAvailableOk

`func (o *OswDetailVO) GetDeviceTemplateAvailableOk() (*bool, bool)`

GetDeviceTemplateAvailableOk returns a tuple with the DeviceTemplateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTemplateAvailable

`func (o *OswDetailVO) SetDeviceTemplateAvailable(v bool)`

SetDeviceTemplateAvailable sets DeviceTemplateAvailable field to given value.

### HasDeviceTemplateAvailable

`func (o *OswDetailVO) HasDeviceTemplateAvailable() bool`

HasDeviceTemplateAvailable returns a boolean if a field has been set.

### GetDisableHwReset

`func (o *OswDetailVO) GetDisableHwReset() bool`

GetDisableHwReset returns the DisableHwReset field if non-nil, zero value otherwise.

### GetDisableHwResetOk

`func (o *OswDetailVO) GetDisableHwResetOk() (*bool, bool)`

GetDisableHwResetOk returns a tuple with the DisableHwReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHwReset

`func (o *OswDetailVO) SetDisableHwReset(v bool)`

SetDisableHwReset sets DisableHwReset field to given value.

### HasDisableHwReset

`func (o *OswDetailVO) HasDisableHwReset() bool`

HasDisableHwReset returns a boolean if a field has been set.

### GetDownlinkList

`func (o *OswDetailVO) GetDownlinkList() []OswDownlinkVO`

GetDownlinkList returns the DownlinkList field if non-nil, zero value otherwise.

### GetDownlinkListOk

`func (o *OswDetailVO) GetDownlinkListOk() (*[]OswDownlinkVO, bool)`

GetDownlinkListOk returns a tuple with the DownlinkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkList

`func (o *OswDetailVO) SetDownlinkList(v []OswDownlinkVO)`

SetDownlinkList sets DownlinkList field to given value.

### HasDownlinkList

`func (o *OswDetailVO) HasDownlinkList() bool`

HasDownlinkList returns a boolean if a field has been set.

### GetDownload

`func (o *OswDetailVO) GetDownload() int64`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *OswDetailVO) GetDownloadOk() (*int64, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *OswDetailVO) SetDownload(v int64)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *OswDetailVO) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetDueTime

`func (o *OswDetailVO) GetDueTime() int64`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *OswDetailVO) GetDueTimeOk() (*int64, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *OswDetailVO) SetDueTime(v int64)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *OswDetailVO) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### GetDueTimeLeft

`func (o *OswDetailVO) GetDueTimeLeft() int64`

GetDueTimeLeft returns the DueTimeLeft field if non-nil, zero value otherwise.

### GetDueTimeLeftOk

`func (o *OswDetailVO) GetDueTimeLeftOk() (*int64, bool)`

GetDueTimeLeftOk returns a tuple with the DueTimeLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTimeLeft

`func (o *OswDetailVO) SetDueTimeLeft(v int64)`

SetDueTimeLeft sets DueTimeLeft field to given value.

### HasDueTimeLeft

`func (o *OswDetailVO) HasDueTimeLeft() bool`

HasDueTimeLeft returns a boolean if a field has been set.

### GetEcspFirstVersion

`func (o *OswDetailVO) GetEcspFirstVersion() int32`

GetEcspFirstVersion returns the EcspFirstVersion field if non-nil, zero value otherwise.

### GetEcspFirstVersionOk

`func (o *OswDetailVO) GetEcspFirstVersionOk() (*int32, bool)`

GetEcspFirstVersionOk returns a tuple with the EcspFirstVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcspFirstVersion

`func (o *OswDetailVO) SetEcspFirstVersion(v int32)`

SetEcspFirstVersion sets EcspFirstVersion field to given value.

### HasEcspFirstVersion

`func (o *OswDetailVO) HasEcspFirstVersion() bool`

HasEcspFirstVersion returns a boolean if a field has been set.

### GetEos

`func (o *OswDetailVO) GetEos() int64`

GetEos returns the Eos field if non-nil, zero value otherwise.

### GetEosOk

`func (o *OswDetailVO) GetEosOk() (*int64, bool)`

GetEosOk returns a tuple with the Eos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEos

`func (o *OswDetailVO) SetEos(v int64)`

SetEos sets Eos field to given value.

### HasEos

`func (o *OswDetailVO) HasEos() bool`

HasEos returns a boolean if a field has been set.

### GetEost

`func (o *OswDetailVO) GetEost() int64`

GetEost returns the Eost field if non-nil, zero value otherwise.

### GetEostOk

`func (o *OswDetailVO) GetEostOk() (*int64, bool)`

GetEostOk returns a tuple with the Eost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEost

`func (o *OswDetailVO) SetEost(v int64)`

SetEost sets Eost field to given value.

### HasEost

`func (o *OswDetailVO) HasEost() bool`

HasEost returns a boolean if a field has been set.

### GetEs

`func (o *OswDetailVO) GetEs() bool`

GetEs returns the Es field if non-nil, zero value otherwise.

### GetEsOk

`func (o *OswDetailVO) GetEsOk() (*bool, bool)`

GetEsOk returns a tuple with the Es field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEs

`func (o *OswDetailVO) SetEs(v bool)`

SetEs sets Es field to given value.

### HasEs

`func (o *OswDetailVO) HasEs() bool`

HasEs returns a boolean if a field has been set.

### GetFanStatus

`func (o *OswDetailVO) GetFanStatus() int32`

GetFanStatus returns the FanStatus field if non-nil, zero value otherwise.

### GetFanStatusOk

`func (o *OswDetailVO) GetFanStatusOk() (*int32, bool)`

GetFanStatusOk returns a tuple with the FanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanStatus

`func (o *OswDetailVO) SetFanStatus(v int32)`

SetFanStatus sets FanStatus field to given value.

### HasFanStatus

`func (o *OswDetailVO) HasFanStatus() bool`

HasFanStatus returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *OswDetailVO) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *OswDetailVO) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *OswDetailVO) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *OswDetailVO) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetForgetId

`func (o *OswDetailVO) GetForgetId() string`

GetForgetId returns the ForgetId field if non-nil, zero value otherwise.

### GetForgetIdOk

`func (o *OswDetailVO) GetForgetIdOk() (*string, bool)`

GetForgetIdOk returns a tuple with the ForgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgetId

`func (o *OswDetailVO) SetForgetId(v string)`

SetForgetId sets ForgetId field to given value.

### HasForgetId

`func (o *OswDetailVO) HasForgetId() bool`

HasForgetId returns a boolean if a field has been set.

### GetForwardDelay

`func (o *OswDetailVO) GetForwardDelay() int32`

GetForwardDelay returns the ForwardDelay field if non-nil, zero value otherwise.

### GetForwardDelayOk

`func (o *OswDetailVO) GetForwardDelayOk() (*int32, bool)`

GetForwardDelayOk returns a tuple with the ForwardDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardDelay

`func (o *OswDetailVO) SetForwardDelay(v int32)`

SetForwardDelay sets ForwardDelay field to given value.

### HasForwardDelay

`func (o *OswDetailVO) HasForwardDelay() bool`

HasForwardDelay returns a boolean if a field has been set.

### GetHelloTime

`func (o *OswDetailVO) GetHelloTime() int32`

GetHelloTime returns the HelloTime field if non-nil, zero value otherwise.

### GetHelloTimeOk

`func (o *OswDetailVO) GetHelloTimeOk() (*int32, bool)`

GetHelloTimeOk returns a tuple with the HelloTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloTime

`func (o *OswDetailVO) SetHelloTime(v int32)`

SetHelloTime sets HelloTime field to given value.

### HasHelloTime

`func (o *OswDetailVO) HasHelloTime() bool`

HasHelloTime returns a boolean if a field has been set.

### GetHwVersion

`func (o *OswDetailVO) GetHwVersion() string`

GetHwVersion returns the HwVersion field if non-nil, zero value otherwise.

### GetHwVersionOk

`func (o *OswDetailVO) GetHwVersionOk() (*string, bool)`

GetHwVersionOk returns a tuple with the HwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwVersion

`func (o *OswDetailVO) SetHwVersion(v string)`

SetHwVersion sets HwVersion field to given value.

### HasHwVersion

`func (o *OswDetailVO) HasHwVersion() bool`

HasHwVersion returns a boolean if a field has been set.

### GetInWhitelist

`func (o *OswDetailVO) GetInWhitelist() bool`

GetInWhitelist returns the InWhitelist field if non-nil, zero value otherwise.

### GetInWhitelistOk

`func (o *OswDetailVO) GetInWhitelistOk() (*bool, bool)`

GetInWhitelistOk returns a tuple with the InWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInWhitelist

`func (o *OswDetailVO) SetInWhitelist(v bool)`

SetInWhitelist sets InWhitelist field to given value.

### HasInWhitelist

`func (o *OswDetailVO) HasInWhitelist() bool`

HasInWhitelist returns a boolean if a field has been set.

### GetInitialUnbindingLimit

`func (o *OswDetailVO) GetInitialUnbindingLimit() int32`

GetInitialUnbindingLimit returns the InitialUnbindingLimit field if non-nil, zero value otherwise.

### GetInitialUnbindingLimitOk

`func (o *OswDetailVO) GetInitialUnbindingLimitOk() (*int32, bool)`

GetInitialUnbindingLimitOk returns a tuple with the InitialUnbindingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialUnbindingLimit

`func (o *OswDetailVO) SetInitialUnbindingLimit(v int32)`

SetInitialUnbindingLimit sets InitialUnbindingLimit field to given value.

### HasInitialUnbindingLimit

`func (o *OswDetailVO) HasInitialUnbindingLimit() bool`

HasInitialUnbindingLimit returns a boolean if a field has been set.

### GetIp

`func (o *OswDetailVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OswDetailVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OswDetailVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OswDetailVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpSetting

`func (o *OswDetailVO) GetIpSetting() IpSettingVO`

GetIpSetting returns the IpSetting field if non-nil, zero value otherwise.

### GetIpSettingOk

`func (o *OswDetailVO) GetIpSettingOk() (*IpSettingVO, bool)`

GetIpSettingOk returns a tuple with the IpSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSetting

`func (o *OswDetailVO) SetIpSetting(v IpSettingVO)`

SetIpSetting sets IpSetting field to given value.

### HasIpSetting

`func (o *OswDetailVO) HasIpSetting() bool`

HasIpSetting returns a boolean if a field has been set.

### GetIpv6List

`func (o *OswDetailVO) GetIpv6List() []string`

GetIpv6List returns the Ipv6List field if non-nil, zero value otherwise.

### GetIpv6ListOk

`func (o *OswDetailVO) GetIpv6ListOk() (*[]string, bool)`

GetIpv6ListOk returns a tuple with the Ipv6List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6List

`func (o *OswDetailVO) SetIpv6List(v []string)`

SetIpv6List sets Ipv6List field to given value.

### HasIpv6List

`func (o *OswDetailVO) HasIpv6List() bool`

HasIpv6List returns a boolean if a field has been set.

### GetJumbo

`func (o *OswDetailVO) GetJumbo() int32`

GetJumbo returns the Jumbo field if non-nil, zero value otherwise.

### GetJumboOk

`func (o *OswDetailVO) GetJumboOk() (*int32, bool)`

GetJumboOk returns a tuple with the Jumbo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumbo

`func (o *OswDetailVO) SetJumbo(v int32)`

SetJumbo sets Jumbo field to given value.

### HasJumbo

`func (o *OswDetailVO) HasJumbo() bool`

HasJumbo returns a boolean if a field has been set.

### GetLagHashAlg

`func (o *OswDetailVO) GetLagHashAlg() int32`

GetLagHashAlg returns the LagHashAlg field if non-nil, zero value otherwise.

### GetLagHashAlgOk

`func (o *OswDetailVO) GetLagHashAlgOk() (*int32, bool)`

GetLagHashAlgOk returns a tuple with the LagHashAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagHashAlg

`func (o *OswDetailVO) SetLagHashAlg(v int32)`

SetLagHashAlg sets LagHashAlg field to given value.

### HasLagHashAlg

`func (o *OswDetailVO) HasLagHashAlg() bool`

HasLagHashAlg returns a boolean if a field has been set.

### GetLags

`func (o *OswDetailVO) GetLags() []OswLagVO`

GetLags returns the Lags field if non-nil, zero value otherwise.

### GetLagsOk

`func (o *OswDetailVO) GetLagsOk() (*[]OswLagVO, bool)`

GetLagsOk returns a tuple with the Lags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLags

`func (o *OswDetailVO) SetLags(v []OswLagVO)`

SetLags sets Lags field to given value.

### HasLags

`func (o *OswDetailVO) HasLags() bool`

HasLags returns a boolean if a field has been set.

### GetLastSeen

`func (o *OswDetailVO) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *OswDetailVO) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *OswDetailVO) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *OswDetailVO) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLedSetting

`func (o *OswDetailVO) GetLedSetting() int32`

GetLedSetting returns the LedSetting field if non-nil, zero value otherwise.

### GetLedSettingOk

`func (o *OswDetailVO) GetLedSettingOk() (*int32, bool)`

GetLedSettingOk returns a tuple with the LedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedSetting

`func (o *OswDetailVO) SetLedSetting(v int32)`

SetLedSetting sets LedSetting field to given value.

### HasLedSetting

`func (o *OswDetailVO) HasLedSetting() bool`

HasLedSetting returns a boolean if a field has been set.

### GetLicenseId

`func (o *OswDetailVO) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *OswDetailVO) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *OswDetailVO) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *OswDetailVO) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetLicenseStatus

`func (o *OswDetailVO) GetLicenseStatus() int32`

GetLicenseStatus returns the LicenseStatus field if non-nil, zero value otherwise.

### GetLicenseStatusOk

`func (o *OswDetailVO) GetLicenseStatusOk() (*int32, bool)`

GetLicenseStatusOk returns a tuple with the LicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatus

`func (o *OswDetailVO) SetLicenseStatus(v int32)`

SetLicenseStatus sets LicenseStatus field to given value.

### HasLicenseStatus

`func (o *OswDetailVO) HasLicenseStatus() bool`

HasLicenseStatus returns a boolean if a field has been set.

### GetLicenseUnbindingLimit

`func (o *OswDetailVO) GetLicenseUnbindingLimit() int32`

GetLicenseUnbindingLimit returns the LicenseUnbindingLimit field if non-nil, zero value otherwise.

### GetLicenseUnbindingLimitOk

`func (o *OswDetailVO) GetLicenseUnbindingLimitOk() (*int32, bool)`

GetLicenseUnbindingLimitOk returns a tuple with the LicenseUnbindingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUnbindingLimit

`func (o *OswDetailVO) SetLicenseUnbindingLimit(v int32)`

SetLicenseUnbindingLimit sets LicenseUnbindingLimit field to given value.

### HasLicenseUnbindingLimit

`func (o *OswDetailVO) HasLicenseUnbindingLimit() bool`

HasLicenseUnbindingLimit returns a boolean if a field has been set.

### GetLocation

`func (o *OswDetailVO) GetLocation() DeviceLocationDetailVO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OswDetailVO) GetLocationOk() (*DeviceLocationDetailVO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OswDetailVO) SetLocation(v DeviceLocationDetailVO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OswDetailVO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLoop

`func (o *OswDetailVO) GetLoop() string`

GetLoop returns the Loop field if non-nil, zero value otherwise.

### GetLoopOk

`func (o *OswDetailVO) GetLoopOk() (*string, bool)`

GetLoopOk returns a tuple with the Loop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoop

`func (o *OswDetailVO) SetLoop(v string)`

SetLoop sets Loop field to given value.

### HasLoop

`func (o *OswDetailVO) HasLoop() bool`

HasLoop returns a boolean if a field has been set.

### GetLoopbackDetectEnable

`func (o *OswDetailVO) GetLoopbackDetectEnable() bool`

GetLoopbackDetectEnable returns the LoopbackDetectEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectEnableOk

`func (o *OswDetailVO) GetLoopbackDetectEnableOk() (*bool, bool)`

GetLoopbackDetectEnableOk returns a tuple with the LoopbackDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectEnable

`func (o *OswDetailVO) SetLoopbackDetectEnable(v bool)`

SetLoopbackDetectEnable sets LoopbackDetectEnable field to given value.

### HasLoopbackDetectEnable

`func (o *OswDetailVO) HasLoopbackDetectEnable() bool`

HasLoopbackDetectEnable returns a boolean if a field has been set.

### GetLoopbackNum

`func (o *OswDetailVO) GetLoopbackNum() int32`

GetLoopbackNum returns the LoopbackNum field if non-nil, zero value otherwise.

### GetLoopbackNumOk

`func (o *OswDetailVO) GetLoopbackNumOk() (*int32, bool)`

GetLoopbackNumOk returns a tuple with the LoopbackNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackNum

`func (o *OswDetailVO) SetLoopbackNum(v int32)`

SetLoopbackNum sets LoopbackNum field to given value.

### HasLoopbackNum

`func (o *OswDetailVO) HasLoopbackNum() bool`

HasLoopbackNum returns a boolean if a field has been set.

### GetMac

`func (o *OswDetailVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswDetailVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswDetailVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswDetailVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMaxAge

`func (o *OswDetailVO) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *OswDetailVO) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *OswDetailVO) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *OswDetailVO) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetMaxHops

`func (o *OswDetailVO) GetMaxHops() int32`

GetMaxHops returns the MaxHops field if non-nil, zero value otherwise.

### GetMaxHopsOk

`func (o *OswDetailVO) GetMaxHopsOk() (*int32, bool)`

GetMaxHopsOk returns a tuple with the MaxHops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHops

`func (o *OswDetailVO) SetMaxHops(v int32)`

SetMaxHops sets MaxHops field to given value.

### HasMaxHops

`func (o *OswDetailVO) HasMaxHops() bool`

HasMaxHops returns a boolean if a field has been set.

### GetMemUitl

`func (o *OswDetailVO) GetMemUitl() int32`

GetMemUitl returns the MemUitl field if non-nil, zero value otherwise.

### GetMemUitlOk

`func (o *OswDetailVO) GetMemUitlOk() (*int32, bool)`

GetMemUitlOk returns a tuple with the MemUitl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUitl

`func (o *OswDetailVO) SetMemUitl(v int32)`

SetMemUitl sets MemUitl field to given value.

### HasMemUitl

`func (o *OswDetailVO) HasMemUitl() bool`

HasMemUitl returns a boolean if a field has been set.

### GetMlagMsg

`func (o *OswDetailVO) GetMlagMsg() MlagMsgVO`

GetMlagMsg returns the MlagMsg field if non-nil, zero value otherwise.

### GetMlagMsgOk

`func (o *OswDetailVO) GetMlagMsgOk() (*MlagMsgVO, bool)`

GetMlagMsgOk returns a tuple with the MlagMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagMsg

`func (o *OswDetailVO) SetMlagMsg(v MlagMsgVO)`

SetMlagMsg sets MlagMsg field to given value.

### HasMlagMsg

`func (o *OswDetailVO) HasMlagMsg() bool`

HasMlagMsg returns a boolean if a field has been set.

### GetMlagPeerInfo

`func (o *OswDetailVO) GetMlagPeerInfo() OswMlagPeerInfoVO`

GetMlagPeerInfo returns the MlagPeerInfo field if non-nil, zero value otherwise.

### GetMlagPeerInfoOk

`func (o *OswDetailVO) GetMlagPeerInfoOk() (*OswMlagPeerInfoVO, bool)`

GetMlagPeerInfoOk returns a tuple with the MlagPeerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPeerInfo

`func (o *OswDetailVO) SetMlagPeerInfo(v OswMlagPeerInfoVO)`

SetMlagPeerInfo sets MlagPeerInfo field to given value.

### HasMlagPeerInfo

`func (o *OswDetailVO) HasMlagPeerInfo() bool`

HasMlagPeerInfo returns a boolean if a field has been set.

### GetModel

`func (o *OswDetailVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OswDetailVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OswDetailVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OswDetailVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OswDetailVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OswDetailVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OswDetailVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OswDetailVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetMoveSiteId

`func (o *OswDetailVO) GetMoveSiteId() string`

GetMoveSiteId returns the MoveSiteId field if non-nil, zero value otherwise.

### GetMoveSiteIdOk

`func (o *OswDetailVO) GetMoveSiteIdOk() (*string, bool)`

GetMoveSiteIdOk returns a tuple with the MoveSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveSiteId

`func (o *OswDetailVO) SetMoveSiteId(v string)`

SetMoveSiteId sets MoveSiteId field to given value.

### HasMoveSiteId

`func (o *OswDetailVO) HasMoveSiteId() bool`

HasMoveSiteId returns a boolean if a field has been set.

### GetMstp

`func (o *OswDetailVO) GetMstp() OswStpMstpVO`

GetMstp returns the Mstp field if non-nil, zero value otherwise.

### GetMstpOk

`func (o *OswDetailVO) GetMstpOk() (*OswStpMstpVO, bool)`

GetMstpOk returns a tuple with the Mstp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMstp

`func (o *OswDetailVO) SetMstp(v OswStpMstpVO)`

SetMstp sets Mstp field to given value.

### HasMstp

`func (o *OswDetailVO) HasMstp() bool`

HasMstp returns a boolean if a field has been set.

### GetMulticast

`func (o *OswDetailVO) GetMulticast() OswLanMulticastVO`

GetMulticast returns the Multicast field if non-nil, zero value otherwise.

### GetMulticastOk

`func (o *OswDetailVO) GetMulticastOk() (*OswLanMulticastVO, bool)`

GetMulticastOk returns a tuple with the Multicast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticast

`func (o *OswDetailVO) SetMulticast(v OswLanMulticastVO)`

SetMulticast sets Multicast field to given value.

### HasMulticast

`func (o *OswDetailVO) HasMulticast() bool`

HasMulticast returns a boolean if a field has been set.

### GetMvlanBridgeVlan

`func (o *OswDetailVO) GetMvlanBridgeVlan() int32`

GetMvlanBridgeVlan returns the MvlanBridgeVlan field if non-nil, zero value otherwise.

### GetMvlanBridgeVlanOk

`func (o *OswDetailVO) GetMvlanBridgeVlanOk() (*int32, bool)`

GetMvlanBridgeVlanOk returns a tuple with the MvlanBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMvlanBridgeVlan

`func (o *OswDetailVO) SetMvlanBridgeVlan(v int32)`

SetMvlanBridgeVlan sets MvlanBridgeVlan field to given value.

### HasMvlanBridgeVlan

`func (o *OswDetailVO) HasMvlanBridgeVlan() bool`

HasMvlanBridgeVlan returns a boolean if a field has been set.

### GetMvlanNetworkId

`func (o *OswDetailVO) GetMvlanNetworkId() string`

GetMvlanNetworkId returns the MvlanNetworkId field if non-nil, zero value otherwise.

### GetMvlanNetworkIdOk

`func (o *OswDetailVO) GetMvlanNetworkIdOk() (*string, bool)`

GetMvlanNetworkIdOk returns a tuple with the MvlanNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMvlanNetworkId

`func (o *OswDetailVO) SetMvlanNetworkId(v string)`

SetMvlanNetworkId sets MvlanNetworkId field to given value.

### HasMvlanNetworkId

`func (o *OswDetailVO) HasMvlanNetworkId() bool`

HasMvlanNetworkId returns a boolean if a field has been set.

### GetName

`func (o *OswDetailVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswDetailVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswDetailVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswDetailVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedUpgrade

`func (o *OswDetailVO) GetNeedUpgrade() bool`

GetNeedUpgrade returns the NeedUpgrade field if non-nil, zero value otherwise.

### GetNeedUpgradeOk

`func (o *OswDetailVO) GetNeedUpgradeOk() (*bool, bool)`

GetNeedUpgradeOk returns a tuple with the NeedUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedUpgrade

`func (o *OswDetailVO) SetNeedUpgrade(v bool)`

SetNeedUpgrade sets NeedUpgrade field to given value.

### HasNeedUpgrade

`func (o *OswDetailVO) HasNeedUpgrade() bool`

HasNeedUpgrade returns a boolean if a field has been set.

### GetOmadacId

`func (o *OswDetailVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *OswDetailVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *OswDetailVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *OswDetailVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetPoeRemain

`func (o *OswDetailVO) GetPoeRemain() float64`

GetPoeRemain returns the PoeRemain field if non-nil, zero value otherwise.

### GetPoeRemainOk

`func (o *OswDetailVO) GetPoeRemainOk() (*float64, bool)`

GetPoeRemainOk returns a tuple with the PoeRemain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeRemain

`func (o *OswDetailVO) SetPoeRemain(v float64)`

SetPoeRemain sets PoeRemain field to given value.

### HasPoeRemain

`func (o *OswDetailVO) HasPoeRemain() bool`

HasPoeRemain returns a boolean if a field has been set.

### GetPoeRemainPercent

`func (o *OswDetailVO) GetPoeRemainPercent() float64`

GetPoeRemainPercent returns the PoeRemainPercent field if non-nil, zero value otherwise.

### GetPoeRemainPercentOk

`func (o *OswDetailVO) GetPoeRemainPercentOk() (*float64, bool)`

GetPoeRemainPercentOk returns a tuple with the PoeRemainPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeRemainPercent

`func (o *OswDetailVO) SetPoeRemainPercent(v float64)`

SetPoeRemainPercent sets PoeRemainPercent field to given value.

### HasPoeRemainPercent

`func (o *OswDetailVO) HasPoeRemainPercent() bool`

HasPoeRemainPercent returns a boolean if a field has been set.

### GetPoeTotalPower

`func (o *OswDetailVO) GetPoeTotalPower() float64`

GetPoeTotalPower returns the PoeTotalPower field if non-nil, zero value otherwise.

### GetPoeTotalPowerOk

`func (o *OswDetailVO) GetPoeTotalPowerOk() (*float64, bool)`

GetPoeTotalPowerOk returns a tuple with the PoeTotalPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeTotalPower

`func (o *OswDetailVO) SetPoeTotalPower(v float64)`

SetPoeTotalPower sets PoeTotalPower field to given value.

### HasPoeTotalPower

`func (o *OswDetailVO) HasPoeTotalPower() bool`

HasPoeTotalPower returns a boolean if a field has been set.

### GetPorts

`func (o *OswDetailVO) GetPorts() []OswPortVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OswDetailVO) GetPortsOk() (*[]OswPortVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OswDetailVO) SetPorts(v []OswPortVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OswDetailVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPriority

`func (o *OswDetailVO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *OswDetailVO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *OswDetailVO) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *OswDetailVO) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPublicIp

`func (o *OswDetailVO) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *OswDetailVO) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *OswDetailVO) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *OswDetailVO) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetQosConfig

`func (o *OswDetailVO) GetQosConfig() OswQosConfigVO`

GetQosConfig returns the QosConfig field if non-nil, zero value otherwise.

### GetQosConfigOk

`func (o *OswDetailVO) GetQosConfigOk() (*OswQosConfigVO, bool)`

GetQosConfigOk returns a tuple with the QosConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosConfig

`func (o *OswDetailVO) SetQosConfig(v OswQosConfigVO)`

SetQosConfig sets QosConfig field to given value.

### HasQosConfig

`func (o *OswDetailVO) HasQosConfig() bool`

HasQosConfig returns a boolean if a field has been set.

### GetRemember

`func (o *OswDetailVO) GetRemember() bool`

GetRemember returns the Remember field if non-nil, zero value otherwise.

### GetRememberOk

`func (o *OswDetailVO) GetRememberOk() (*bool, bool)`

GetRememberOk returns a tuple with the Remember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemember

`func (o *OswDetailVO) SetRemember(v bool)`

SetRemember sets Remember field to given value.

### HasRemember

`func (o *OswDetailVO) HasRemember() bool`

HasRemember returns a boolean if a field has been set.

### GetRememberDevice

`func (o *OswDetailVO) GetRememberDevice() int32`

GetRememberDevice returns the RememberDevice field if non-nil, zero value otherwise.

### GetRememberDeviceOk

`func (o *OswDetailVO) GetRememberDeviceOk() (*int32, bool)`

GetRememberDeviceOk returns a tuple with the RememberDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberDevice

`func (o *OswDetailVO) SetRememberDevice(v int32)`

SetRememberDevice sets RememberDevice field to given value.

### HasRememberDevice

`func (o *OswDetailVO) HasRememberDevice() bool`

HasRememberDevice returns a boolean if a field has been set.

### GetResource

`func (o *OswDetailVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *OswDetailVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *OswDetailVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *OswDetailVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetRxRate

`func (o *OswDetailVO) GetRxRate() int64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *OswDetailVO) GetRxRateOk() (*int64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *OswDetailVO) SetRxRate(v int64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *OswDetailVO) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetSdm

`func (o *OswDetailVO) GetSdm() OswSdmTemplateVO`

GetSdm returns the Sdm field if non-nil, zero value otherwise.

### GetSdmOk

`func (o *OswDetailVO) GetSdmOk() (*OswSdmTemplateVO, bool)`

GetSdmOk returns a tuple with the Sdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdm

`func (o *OswDetailVO) SetSdm(v OswSdmTemplateVO)`

SetSdm sets Sdm field to given value.

### HasSdm

`func (o *OswDetailVO) HasSdm() bool`

HasSdm returns a boolean if a field has been set.

### GetShowModel

`func (o *OswDetailVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *OswDetailVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *OswDetailVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *OswDetailVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetSiteName

`func (o *OswDetailVO) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *OswDetailVO) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *OswDetailVO) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *OswDetailVO) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSiteTemplateId

`func (o *OswDetailVO) GetSiteTemplateId() string`

GetSiteTemplateId returns the SiteTemplateId field if non-nil, zero value otherwise.

### GetSiteTemplateIdOk

`func (o *OswDetailVO) GetSiteTemplateIdOk() (*string, bool)`

GetSiteTemplateIdOk returns a tuple with the SiteTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTemplateId

`func (o *OswDetailVO) SetSiteTemplateId(v string)`

SetSiteTemplateId sets SiteTemplateId field to given value.

### HasSiteTemplateId

`func (o *OswDetailVO) HasSiteTemplateId() bool`

HasSiteTemplateId returns a boolean if a field has been set.

### GetSiteTemplateName

`func (o *OswDetailVO) GetSiteTemplateName() string`

GetSiteTemplateName returns the SiteTemplateName field if non-nil, zero value otherwise.

### GetSiteTemplateNameOk

`func (o *OswDetailVO) GetSiteTemplateNameOk() (*string, bool)`

GetSiteTemplateNameOk returns a tuple with the SiteTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTemplateName

`func (o *OswDetailVO) SetSiteTemplateName(v string)`

SetSiteTemplateName sets SiteTemplateName field to given value.

### HasSiteTemplateName

`func (o *OswDetailVO) HasSiteTemplateName() bool`

HasSiteTemplateName returns a boolean if a field has been set.

### GetSn

`func (o *OswDetailVO) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *OswDetailVO) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *OswDetailVO) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *OswDetailVO) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetSnmp

`func (o *OswDetailVO) GetSnmp() OswSnmpVO`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *OswDetailVO) GetSnmpOk() (*OswSnmpVO, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *OswDetailVO) SetSnmp(v OswSnmpVO)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *OswDetailVO) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetSpecialModel

`func (o *OswDetailVO) GetSpecialModel() string`

GetSpecialModel returns the SpecialModel field if non-nil, zero value otherwise.

### GetSpecialModelOk

`func (o *OswDetailVO) GetSpecialModelOk() (*string, bool)`

GetSpecialModelOk returns a tuple with the SpecialModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialModel

`func (o *OswDetailVO) SetSpecialModel(v string)`

SetSpecialModel sets SpecialModel field to given value.

### HasSpecialModel

`func (o *OswDetailVO) HasSpecialModel() bool`

HasSpecialModel returns a boolean if a field has been set.

### GetSpeeds

`func (o *OswDetailVO) GetSpeeds() []int32`

GetSpeeds returns the Speeds field if non-nil, zero value otherwise.

### GetSpeedsOk

`func (o *OswDetailVO) GetSpeedsOk() (*[]int32, bool)`

GetSpeedsOk returns a tuple with the Speeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeeds

`func (o *OswDetailVO) SetSpeeds(v []int32)`

SetSpeeds sets Speeds field to given value.

### HasSpeeds

`func (o *OswDetailVO) HasSpeeds() bool`

HasSpeeds returns a boolean if a field has been set.

### GetStackDevice

`func (o *OswDetailVO) GetStackDevice() bool`

GetStackDevice returns the StackDevice field if non-nil, zero value otherwise.

### GetStackDeviceOk

`func (o *OswDetailVO) GetStackDeviceOk() (*bool, bool)`

GetStackDeviceOk returns a tuple with the StackDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackDevice

`func (o *OswDetailVO) SetStackDevice(v bool)`

SetStackDevice sets StackDevice field to given value.

### HasStackDevice

`func (o *OswDetailVO) HasStackDevice() bool`

HasStackDevice returns a boolean if a field has been set.

### GetStackMsg

`func (o *OswDetailVO) GetStackMsg() StackMsgVO`

GetStackMsg returns the StackMsg field if non-nil, zero value otherwise.

### GetStackMsgOk

`func (o *OswDetailVO) GetStackMsgOk() (*StackMsgVO, bool)`

GetStackMsgOk returns a tuple with the StackMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackMsg

`func (o *OswDetailVO) SetStackMsg(v StackMsgVO)`

SetStackMsg sets StackMsg field to given value.

### HasStackMsg

`func (o *OswDetailVO) HasStackMsg() bool`

HasStackMsg returns a boolean if a field has been set.

### GetStackPorts

`func (o *OswDetailVO) GetStackPorts() []OswStackPortGroupVO`

GetStackPorts returns the StackPorts field if non-nil, zero value otherwise.

### GetStackPortsOk

`func (o *OswDetailVO) GetStackPortsOk() (*[]OswStackPortGroupVO, bool)`

GetStackPortsOk returns a tuple with the StackPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPorts

`func (o *OswDetailVO) SetStackPorts(v []OswStackPortGroupVO)`

SetStackPorts sets StackPorts field to given value.

### HasStackPorts

`func (o *OswDetailVO) HasStackPorts() bool`

HasStackPorts returns a boolean if a field has been set.

### GetStatus

`func (o *OswDetailVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswDetailVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswDetailVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OswDetailVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *OswDetailVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *OswDetailVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *OswDetailVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *OswDetailVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetStp

`func (o *OswDetailVO) GetStp() int32`

GetStp returns the Stp field if non-nil, zero value otherwise.

### GetStpOk

`func (o *OswDetailVO) GetStpOk() (*int32, bool)`

GetStpOk returns a tuple with the Stp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStp

`func (o *OswDetailVO) SetStp(v int32)`

SetStp sets Stp field to given value.

### HasStp

`func (o *OswDetailVO) HasStp() bool`

HasStp returns a boolean if a field has been set.

### GetStpLinkList

`func (o *OswDetailVO) GetStpLinkList() []OswDownlinkVO`

GetStpLinkList returns the StpLinkList field if non-nil, zero value otherwise.

### GetStpLinkListOk

`func (o *OswDetailVO) GetStpLinkListOk() (*[]OswDownlinkVO, bool)`

GetStpLinkListOk returns a tuple with the StpLinkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpLinkList

`func (o *OswDetailVO) SetStpLinkList(v []OswDownlinkVO)`

SetStpLinkList sets StpLinkList field to given value.

### HasStpLinkList

`func (o *OswDetailVO) HasStpLinkList() bool`

HasStpLinkList returns a boolean if a field has been set.

### GetSupportAnomaly

`func (o *OswDetailVO) GetSupportAnomaly() bool`

GetSupportAnomaly returns the SupportAnomaly field if non-nil, zero value otherwise.

### GetSupportAnomalyOk

`func (o *OswDetailVO) GetSupportAnomalyOk() (*bool, bool)`

GetSupportAnomalyOk returns a tuple with the SupportAnomaly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAnomaly

`func (o *OswDetailVO) SetSupportAnomaly(v bool)`

SetSupportAnomaly sets SupportAnomaly field to given value.

### HasSupportAnomaly

`func (o *OswDetailVO) HasSupportAnomaly() bool`

HasSupportAnomaly returns a boolean if a field has been set.

### GetSupportHealth

`func (o *OswDetailVO) GetSupportHealth() bool`

GetSupportHealth returns the SupportHealth field if non-nil, zero value otherwise.

### GetSupportHealthOk

`func (o *OswDetailVO) GetSupportHealthOk() (*bool, bool)`

GetSupportHealthOk returns a tuple with the SupportHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportHealth

`func (o *OswDetailVO) SetSupportHealth(v bool)`

SetSupportHealth sets SupportHealth field to given value.

### HasSupportHealth

`func (o *OswDetailVO) HasSupportHealth() bool`

HasSupportHealth returns a boolean if a field has been set.

### GetSupportIpv6Acl

`func (o *OswDetailVO) GetSupportIpv6Acl() bool`

GetSupportIpv6Acl returns the SupportIpv6Acl field if non-nil, zero value otherwise.

### GetSupportIpv6AclOk

`func (o *OswDetailVO) GetSupportIpv6AclOk() (*bool, bool)`

GetSupportIpv6AclOk returns a tuple with the SupportIpv6Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIpv6Acl

`func (o *OswDetailVO) SetSupportIpv6Acl(v bool)`

SetSupportIpv6Acl sets SupportIpv6Acl field to given value.

### HasSupportIpv6Acl

`func (o *OswDetailVO) HasSupportIpv6Acl() bool`

HasSupportIpv6Acl returns a boolean if a field has been set.

### GetSupportLocatePort

`func (o *OswDetailVO) GetSupportLocatePort() bool`

GetSupportLocatePort returns the SupportLocatePort field if non-nil, zero value otherwise.

### GetSupportLocatePortOk

`func (o *OswDetailVO) GetSupportLocatePortOk() (*bool, bool)`

GetSupportLocatePortOk returns a tuple with the SupportLocatePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLocatePort

`func (o *OswDetailVO) SetSupportLocatePort(v bool)`

SetSupportLocatePort sets SupportLocatePort field to given value.

### HasSupportLocatePort

`func (o *OswDetailVO) HasSupportLocatePort() bool`

HasSupportLocatePort returns a boolean if a field has been set.

### GetSupportVlanIf

`func (o *OswDetailVO) GetSupportVlanIf() bool`

GetSupportVlanIf returns the SupportVlanIf field if non-nil, zero value otherwise.

### GetSupportVlanIfOk

`func (o *OswDetailVO) GetSupportVlanIfOk() (*bool, bool)`

GetSupportVlanIfOk returns a tuple with the SupportVlanIf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVlanIf

`func (o *OswDetailVO) SetSupportVlanIf(v bool)`

SetSupportVlanIf sets SupportVlanIf field to given value.

### HasSupportVlanIf

`func (o *OswDetailVO) HasSupportVlanIf() bool`

HasSupportVlanIf returns a boolean if a field has been set.

### GetSupportVrf

`func (o *OswDetailVO) GetSupportVrf() bool`

GetSupportVrf returns the SupportVrf field if non-nil, zero value otherwise.

### GetSupportVrfOk

`func (o *OswDetailVO) GetSupportVrfOk() (*bool, bool)`

GetSupportVrfOk returns a tuple with the SupportVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVrf

`func (o *OswDetailVO) SetSupportVrf(v bool)`

SetSupportVrf sets SupportVrf field to given value.

### HasSupportVrf

`func (o *OswDetailVO) HasSupportVrf() bool`

HasSupportVrf returns a boolean if a field has been set.

### GetTagIds

`func (o *OswDetailVO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *OswDetailVO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *OswDetailVO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *OswDetailVO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTemplateId

`func (o *OswDetailVO) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *OswDetailVO) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *OswDetailVO) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *OswDetailVO) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateName

`func (o *OswDetailVO) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *OswDetailVO) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *OswDetailVO) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *OswDetailVO) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetTemplateSettings

`func (o *OswDetailVO) GetTemplateSettings() []int32`

GetTemplateSettings returns the TemplateSettings field if non-nil, zero value otherwise.

### GetTemplateSettingsOk

`func (o *OswDetailVO) GetTemplateSettingsOk() (*[]int32, bool)`

GetTemplateSettingsOk returns a tuple with the TemplateSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSettings

`func (o *OswDetailVO) SetTemplateSettings(v []int32)`

SetTemplateSettings sets TemplateSettings field to given value.

### HasTemplateSettings

`func (o *OswDetailVO) HasTemplateSettings() bool`

HasTemplateSettings returns a boolean if a field has been set.

### GetTerminalPrefix

`func (o *OswDetailVO) GetTerminalPrefix() string`

GetTerminalPrefix returns the TerminalPrefix field if non-nil, zero value otherwise.

### GetTerminalPrefixOk

`func (o *OswDetailVO) GetTerminalPrefixOk() (*string, bool)`

GetTerminalPrefixOk returns a tuple with the TerminalPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalPrefix

`func (o *OswDetailVO) SetTerminalPrefix(v string)`

SetTerminalPrefix sets TerminalPrefix field to given value.

### HasTerminalPrefix

`func (o *OswDetailVO) HasTerminalPrefix() bool`

HasTerminalPrefix returns a boolean if a field has been set.

### GetTxHoldCount

`func (o *OswDetailVO) GetTxHoldCount() int32`

GetTxHoldCount returns the TxHoldCount field if non-nil, zero value otherwise.

### GetTxHoldCountOk

`func (o *OswDetailVO) GetTxHoldCountOk() (*int32, bool)`

GetTxHoldCountOk returns a tuple with the TxHoldCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHoldCount

`func (o *OswDetailVO) SetTxHoldCount(v int32)`

SetTxHoldCount sets TxHoldCount field to given value.

### HasTxHoldCount

`func (o *OswDetailVO) HasTxHoldCount() bool`

HasTxHoldCount returns a boolean if a field has been set.

### GetTxRate

`func (o *OswDetailVO) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *OswDetailVO) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *OswDetailVO) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *OswDetailVO) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.

### GetType

`func (o *OswDetailVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswDetailVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswDetailVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OswDetailVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *OswDetailVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OswDetailVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OswDetailVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *OswDetailVO) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUplink

`func (o *OswDetailVO) GetUplink() OswUplinkVO`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *OswDetailVO) GetUplinkOk() (*OswUplinkVO, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *OswDetailVO) SetUplink(v OswUplinkVO)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *OswDetailVO) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetUplinkPort

`func (o *OswDetailVO) GetUplinkPort() int32`

GetUplinkPort returns the UplinkPort field if non-nil, zero value otherwise.

### GetUplinkPortOk

`func (o *OswDetailVO) GetUplinkPortOk() (*int32, bool)`

GetUplinkPortOk returns a tuple with the UplinkPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkPort

`func (o *OswDetailVO) SetUplinkPort(v int32)`

SetUplinkPort sets UplinkPort field to given value.

### HasUplinkPort

`func (o *OswDetailVO) HasUplinkPort() bool`

HasUplinkPort returns a boolean if a field has been set.

### GetUplinkStPort

`func (o *OswDetailVO) GetUplinkStPort() string`

GetUplinkStPort returns the UplinkStPort field if non-nil, zero value otherwise.

### GetUplinkStPortOk

`func (o *OswDetailVO) GetUplinkStPortOk() (*string, bool)`

GetUplinkStPortOk returns a tuple with the UplinkStPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkStPort

`func (o *OswDetailVO) SetUplinkStPort(v string)`

SetUplinkStPort sets UplinkStPort field to given value.

### HasUplinkStPort

`func (o *OswDetailVO) HasUplinkStPort() bool`

HasUplinkStPort returns a boolean if a field has been set.

### GetUpload

`func (o *OswDetailVO) GetUpload() int64`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *OswDetailVO) GetUploadOk() (*int64, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *OswDetailVO) SetUpload(v int64)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *OswDetailVO) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetUptime

`func (o *OswDetailVO) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *OswDetailVO) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *OswDetailVO) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *OswDetailVO) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetUptimeLong

`func (o *OswDetailVO) GetUptimeLong() int64`

GetUptimeLong returns the UptimeLong field if non-nil, zero value otherwise.

### GetUptimeLongOk

`func (o *OswDetailVO) GetUptimeLongOk() (*int64, bool)`

GetUptimeLongOk returns a tuple with the UptimeLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeLong

`func (o *OswDetailVO) SetUptimeLong(v int64)`

SetUptimeLong sets UptimeLong field to given value.

### HasUptimeLong

`func (o *OswDetailVO) HasUptimeLong() bool`

HasUptimeLong returns a boolean if a field has been set.

### GetVersion

`func (o *OswDetailVO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OswDetailVO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OswDetailVO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OswDetailVO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


