# OswStackDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | Site of the device | [optional] 
**AbnormalReason** | Pointer to **int32** | When status is 1, show abnormal reason. | [optional] 
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
**Id** | Pointer to **string** | Stack ID | [optional] 
**InWhitelist** | Pointer to **bool** | Whether the device is in white list | [optional] 
**InactiveUnits** | Pointer to **[]int32** | Inactive units | [optional] 
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
**LocateEnable** | Pointer to **bool** | Locate Enable | [optional] 
**Location** | Pointer to [**DeviceLocationDetailVO**](DeviceLocationDetailVO.md) |  | [optional] 
**Loop** | Pointer to **string** | Loop information | [optional] 
**LoopbackDetectEnable** | Pointer to **bool** | Loopback Detect Enable | [optional] 
**LoopbackNum** | Pointer to **int32** | Loopback Num | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**MacDelay** | Pointer to [**MacDelayVO**](MacDelayVO.md) |  | [optional] 
**MasterActive** | Pointer to **bool** | Master Active | [optional] 
**MasterMac** | Pointer to **string** | Master Device Mac | [optional] 
**MaxAge** | Pointer to **int32** | STP maxAge | [optional] 
**MaxHops** | Pointer to **int32** | STP maxHops | [optional] 
**MemUitl** | Pointer to **int32** | Real-time memory usage | [optional] 
**Member** | Pointer to [**[]OswStackMemberVO**](OswStackMemberVO.md) | Member List | [optional] 
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
**SiteId** | Pointer to **string** | Site ID | [optional] 
**SiteName** | Pointer to **string** | Site Name | [optional] 
**SiteTemplateId** | Pointer to **string** | Template ID bound to the site | [optional] 
**SiteTemplateName** | Pointer to **string** | Template name bound to the site | [optional] 
**Sn** | Pointer to **string** | Device serial number | [optional] 
**Snmp** | Pointer to [**OswSnmpVO**](OswSnmpVO.md) |  | [optional] 
**SpecialModel** | Pointer to **string** | Special device model,for example:EAP225-Outdoor-1a20a950b8d950e8 | [optional] 
**Speeds** | Pointer to **[]int32** | Supported rate list for all ports. Speeds should be a value as follows: 0:auto; 1:10M; 2:100M; 3:1000M; 4:2.5G; 5:10G; 6:5G; 7:25G; 8:100G; 9:40G; -1:error; no value:all rate supported | [optional] 
**StackDevice** | Pointer to **bool** | Stack Device | [optional] 
**StackLags** | Pointer to [**[]OswStackMemberLagVO**](OswStackMemberLagVO.md) | Stack Lag List | [optional] 
**StackMsg** | Pointer to [**StackMsgVO**](StackMsgVO.md) |  | [optional] 
**StackPorts** | Pointer to [**[]OswStackPortGroupVO**](OswStackPortGroupVO.md) | Stack Port List | [optional] 
**Status** | Pointer to **int32** | Status is a value as follows: 0: normal; 1:abnormal; 2:stackNotReady | [optional] 
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
**VirtualMac** | Pointer to **string** | Virtual Mac | [optional] 

## Methods

### NewOswStackDetailVO

`func NewOswStackDetailVO() *OswStackDetailVO`

NewOswStackDetailVO instantiates a new OswStackDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackDetailVOWithDefaults

`func NewOswStackDetailVOWithDefaults() *OswStackDetailVO`

NewOswStackDetailVOWithDefaults instantiates a new OswStackDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *OswStackDetailVO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *OswStackDetailVO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *OswStackDetailVO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *OswStackDetailVO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetAbnormalReason

`func (o *OswStackDetailVO) GetAbnormalReason() int32`

GetAbnormalReason returns the AbnormalReason field if non-nil, zero value otherwise.

### GetAbnormalReasonOk

`func (o *OswStackDetailVO) GetAbnormalReasonOk() (*int32, bool)`

GetAbnormalReasonOk returns a tuple with the AbnormalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbnormalReason

`func (o *OswStackDetailVO) SetAbnormalReason(v int32)`

SetAbnormalReason sets AbnormalReason field to given value.

### HasAbnormalReason

`func (o *OswStackDetailVO) HasAbnormalReason() bool`

HasAbnormalReason returns a boolean if a field has been set.

### GetActive

`func (o *OswStackDetailVO) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OswStackDetailVO) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OswStackDetailVO) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *OswStackDetailVO) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddedInAdvanced

`func (o *OswStackDetailVO) GetAddedInAdvanced() bool`

GetAddedInAdvanced returns the AddedInAdvanced field if non-nil, zero value otherwise.

### GetAddedInAdvancedOk

`func (o *OswStackDetailVO) GetAddedInAdvancedOk() (*bool, bool)`

GetAddedInAdvancedOk returns a tuple with the AddedInAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedInAdvanced

`func (o *OswStackDetailVO) SetAddedInAdvanced(v bool)`

SetAddedInAdvanced sets AddedInAdvanced field to given value.

### HasAddedInAdvanced

`func (o *OswStackDetailVO) HasAddedInAdvanced() bool`

HasAddedInAdvanced returns a boolean if a field has been set.

### GetBoundDeviceTemplate

`func (o *OswStackDetailVO) GetBoundDeviceTemplate() bool`

GetBoundDeviceTemplate returns the BoundDeviceTemplate field if non-nil, zero value otherwise.

### GetBoundDeviceTemplateOk

`func (o *OswStackDetailVO) GetBoundDeviceTemplateOk() (*bool, bool)`

GetBoundDeviceTemplateOk returns a tuple with the BoundDeviceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundDeviceTemplate

`func (o *OswStackDetailVO) SetBoundDeviceTemplate(v bool)`

SetBoundDeviceTemplate sets BoundDeviceTemplate field to given value.

### HasBoundDeviceTemplate

`func (o *OswStackDetailVO) HasBoundDeviceTemplate() bool`

HasBoundDeviceTemplate returns a boolean if a field has been set.

### GetBoundSiteTemplate

`func (o *OswStackDetailVO) GetBoundSiteTemplate() bool`

GetBoundSiteTemplate returns the BoundSiteTemplate field if non-nil, zero value otherwise.

### GetBoundSiteTemplateOk

`func (o *OswStackDetailVO) GetBoundSiteTemplateOk() (*bool, bool)`

GetBoundSiteTemplateOk returns a tuple with the BoundSiteTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSiteTemplate

`func (o *OswStackDetailVO) SetBoundSiteTemplate(v bool)`

SetBoundSiteTemplate sets BoundSiteTemplate field to given value.

### HasBoundSiteTemplate

`func (o *OswStackDetailVO) HasBoundSiteTemplate() bool`

HasBoundSiteTemplate returns a boolean if a field has been set.

### GetCategory

`func (o *OswStackDetailVO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *OswStackDetailVO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *OswStackDetailVO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *OswStackDetailVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCompatible

`func (o *OswStackDetailVO) GetCompatible() int32`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *OswStackDetailVO) GetCompatibleOk() (*int32, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *OswStackDetailVO) SetCompatible(v int32)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *OswStackDetailVO) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetCompoundModel

`func (o *OswStackDetailVO) GetCompoundModel() string`

GetCompoundModel returns the CompoundModel field if non-nil, zero value otherwise.

### GetCompoundModelOk

`func (o *OswStackDetailVO) GetCompoundModelOk() (*string, bool)`

GetCompoundModelOk returns a tuple with the CompoundModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompoundModel

`func (o *OswStackDetailVO) SetCompoundModel(v string)`

SetCompoundModel sets CompoundModel field to given value.

### HasCompoundModel

`func (o *OswStackDetailVO) HasCompoundModel() bool`

HasCompoundModel returns a boolean if a field has been set.

### GetCpuUtil

`func (o *OswStackDetailVO) GetCpuUtil() int32`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *OswStackDetailVO) GetCpuUtilOk() (*int32, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *OswStackDetailVO) SetCpuUtil(v int32)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *OswStackDetailVO) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetCustomId

`func (o *OswStackDetailVO) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *OswStackDetailVO) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *OswStackDetailVO) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.

### HasCustomId

`func (o *OswStackDetailVO) HasCustomId() bool`

HasCustomId returns a boolean if a field has been set.

### GetCustomName

`func (o *OswStackDetailVO) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *OswStackDetailVO) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *OswStackDetailVO) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *OswStackDetailVO) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### GetDescription

`func (o *OswStackDetailVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OswStackDetailVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OswStackDetailVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OswStackDetailVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevCap

`func (o *OswStackDetailVO) GetDevCap() OswDevCapVO`

GetDevCap returns the DevCap field if non-nil, zero value otherwise.

### GetDevCapOk

`func (o *OswStackDetailVO) GetDevCapOk() (*OswDevCapVO, bool)`

GetDevCapOk returns a tuple with the DevCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevCap

`func (o *OswStackDetailVO) SetDevCap(v OswDevCapVO)`

SetDevCap sets DevCap field to given value.

### HasDevCap

`func (o *OswStackDetailVO) HasDevCap() bool`

HasDevCap returns a boolean if a field has been set.

### GetDeviceMisc

`func (o *OswStackDetailVO) GetDeviceMisc() OswDeviceMiscVO`

GetDeviceMisc returns the DeviceMisc field if non-nil, zero value otherwise.

### GetDeviceMiscOk

`func (o *OswStackDetailVO) GetDeviceMiscOk() (*OswDeviceMiscVO, bool)`

GetDeviceMiscOk returns a tuple with the DeviceMisc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMisc

`func (o *OswStackDetailVO) SetDeviceMisc(v OswDeviceMiscVO)`

SetDeviceMisc sets DeviceMisc field to given value.

### HasDeviceMisc

`func (o *OswStackDetailVO) HasDeviceMisc() bool`

HasDeviceMisc returns a boolean if a field has been set.

### GetDeviceSeriesType

`func (o *OswStackDetailVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *OswStackDetailVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *OswStackDetailVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *OswStackDetailVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetDeviceTemplateAvailable

`func (o *OswStackDetailVO) GetDeviceTemplateAvailable() bool`

GetDeviceTemplateAvailable returns the DeviceTemplateAvailable field if non-nil, zero value otherwise.

### GetDeviceTemplateAvailableOk

`func (o *OswStackDetailVO) GetDeviceTemplateAvailableOk() (*bool, bool)`

GetDeviceTemplateAvailableOk returns a tuple with the DeviceTemplateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTemplateAvailable

`func (o *OswStackDetailVO) SetDeviceTemplateAvailable(v bool)`

SetDeviceTemplateAvailable sets DeviceTemplateAvailable field to given value.

### HasDeviceTemplateAvailable

`func (o *OswStackDetailVO) HasDeviceTemplateAvailable() bool`

HasDeviceTemplateAvailable returns a boolean if a field has been set.

### GetDisableHwReset

`func (o *OswStackDetailVO) GetDisableHwReset() bool`

GetDisableHwReset returns the DisableHwReset field if non-nil, zero value otherwise.

### GetDisableHwResetOk

`func (o *OswStackDetailVO) GetDisableHwResetOk() (*bool, bool)`

GetDisableHwResetOk returns a tuple with the DisableHwReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHwReset

`func (o *OswStackDetailVO) SetDisableHwReset(v bool)`

SetDisableHwReset sets DisableHwReset field to given value.

### HasDisableHwReset

`func (o *OswStackDetailVO) HasDisableHwReset() bool`

HasDisableHwReset returns a boolean if a field has been set.

### GetDownlinkList

`func (o *OswStackDetailVO) GetDownlinkList() []OswDownlinkVO`

GetDownlinkList returns the DownlinkList field if non-nil, zero value otherwise.

### GetDownlinkListOk

`func (o *OswStackDetailVO) GetDownlinkListOk() (*[]OswDownlinkVO, bool)`

GetDownlinkListOk returns a tuple with the DownlinkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkList

`func (o *OswStackDetailVO) SetDownlinkList(v []OswDownlinkVO)`

SetDownlinkList sets DownlinkList field to given value.

### HasDownlinkList

`func (o *OswStackDetailVO) HasDownlinkList() bool`

HasDownlinkList returns a boolean if a field has been set.

### GetDownload

`func (o *OswStackDetailVO) GetDownload() int64`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *OswStackDetailVO) GetDownloadOk() (*int64, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *OswStackDetailVO) SetDownload(v int64)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *OswStackDetailVO) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetDueTime

`func (o *OswStackDetailVO) GetDueTime() int64`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *OswStackDetailVO) GetDueTimeOk() (*int64, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *OswStackDetailVO) SetDueTime(v int64)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *OswStackDetailVO) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### GetDueTimeLeft

`func (o *OswStackDetailVO) GetDueTimeLeft() int64`

GetDueTimeLeft returns the DueTimeLeft field if non-nil, zero value otherwise.

### GetDueTimeLeftOk

`func (o *OswStackDetailVO) GetDueTimeLeftOk() (*int64, bool)`

GetDueTimeLeftOk returns a tuple with the DueTimeLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTimeLeft

`func (o *OswStackDetailVO) SetDueTimeLeft(v int64)`

SetDueTimeLeft sets DueTimeLeft field to given value.

### HasDueTimeLeft

`func (o *OswStackDetailVO) HasDueTimeLeft() bool`

HasDueTimeLeft returns a boolean if a field has been set.

### GetEcspFirstVersion

`func (o *OswStackDetailVO) GetEcspFirstVersion() int32`

GetEcspFirstVersion returns the EcspFirstVersion field if non-nil, zero value otherwise.

### GetEcspFirstVersionOk

`func (o *OswStackDetailVO) GetEcspFirstVersionOk() (*int32, bool)`

GetEcspFirstVersionOk returns a tuple with the EcspFirstVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcspFirstVersion

`func (o *OswStackDetailVO) SetEcspFirstVersion(v int32)`

SetEcspFirstVersion sets EcspFirstVersion field to given value.

### HasEcspFirstVersion

`func (o *OswStackDetailVO) HasEcspFirstVersion() bool`

HasEcspFirstVersion returns a boolean if a field has been set.

### GetEos

`func (o *OswStackDetailVO) GetEos() int64`

GetEos returns the Eos field if non-nil, zero value otherwise.

### GetEosOk

`func (o *OswStackDetailVO) GetEosOk() (*int64, bool)`

GetEosOk returns a tuple with the Eos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEos

`func (o *OswStackDetailVO) SetEos(v int64)`

SetEos sets Eos field to given value.

### HasEos

`func (o *OswStackDetailVO) HasEos() bool`

HasEos returns a boolean if a field has been set.

### GetEost

`func (o *OswStackDetailVO) GetEost() int64`

GetEost returns the Eost field if non-nil, zero value otherwise.

### GetEostOk

`func (o *OswStackDetailVO) GetEostOk() (*int64, bool)`

GetEostOk returns a tuple with the Eost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEost

`func (o *OswStackDetailVO) SetEost(v int64)`

SetEost sets Eost field to given value.

### HasEost

`func (o *OswStackDetailVO) HasEost() bool`

HasEost returns a boolean if a field has been set.

### GetEs

`func (o *OswStackDetailVO) GetEs() bool`

GetEs returns the Es field if non-nil, zero value otherwise.

### GetEsOk

`func (o *OswStackDetailVO) GetEsOk() (*bool, bool)`

GetEsOk returns a tuple with the Es field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEs

`func (o *OswStackDetailVO) SetEs(v bool)`

SetEs sets Es field to given value.

### HasEs

`func (o *OswStackDetailVO) HasEs() bool`

HasEs returns a boolean if a field has been set.

### GetFanStatus

`func (o *OswStackDetailVO) GetFanStatus() int32`

GetFanStatus returns the FanStatus field if non-nil, zero value otherwise.

### GetFanStatusOk

`func (o *OswStackDetailVO) GetFanStatusOk() (*int32, bool)`

GetFanStatusOk returns a tuple with the FanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanStatus

`func (o *OswStackDetailVO) SetFanStatus(v int32)`

SetFanStatus sets FanStatus field to given value.

### HasFanStatus

`func (o *OswStackDetailVO) HasFanStatus() bool`

HasFanStatus returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *OswStackDetailVO) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *OswStackDetailVO) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *OswStackDetailVO) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *OswStackDetailVO) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetForgetId

`func (o *OswStackDetailVO) GetForgetId() string`

GetForgetId returns the ForgetId field if non-nil, zero value otherwise.

### GetForgetIdOk

`func (o *OswStackDetailVO) GetForgetIdOk() (*string, bool)`

GetForgetIdOk returns a tuple with the ForgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgetId

`func (o *OswStackDetailVO) SetForgetId(v string)`

SetForgetId sets ForgetId field to given value.

### HasForgetId

`func (o *OswStackDetailVO) HasForgetId() bool`

HasForgetId returns a boolean if a field has been set.

### GetForwardDelay

`func (o *OswStackDetailVO) GetForwardDelay() int32`

GetForwardDelay returns the ForwardDelay field if non-nil, zero value otherwise.

### GetForwardDelayOk

`func (o *OswStackDetailVO) GetForwardDelayOk() (*int32, bool)`

GetForwardDelayOk returns a tuple with the ForwardDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardDelay

`func (o *OswStackDetailVO) SetForwardDelay(v int32)`

SetForwardDelay sets ForwardDelay field to given value.

### HasForwardDelay

`func (o *OswStackDetailVO) HasForwardDelay() bool`

HasForwardDelay returns a boolean if a field has been set.

### GetHelloTime

`func (o *OswStackDetailVO) GetHelloTime() int32`

GetHelloTime returns the HelloTime field if non-nil, zero value otherwise.

### GetHelloTimeOk

`func (o *OswStackDetailVO) GetHelloTimeOk() (*int32, bool)`

GetHelloTimeOk returns a tuple with the HelloTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloTime

`func (o *OswStackDetailVO) SetHelloTime(v int32)`

SetHelloTime sets HelloTime field to given value.

### HasHelloTime

`func (o *OswStackDetailVO) HasHelloTime() bool`

HasHelloTime returns a boolean if a field has been set.

### GetHwVersion

`func (o *OswStackDetailVO) GetHwVersion() string`

GetHwVersion returns the HwVersion field if non-nil, zero value otherwise.

### GetHwVersionOk

`func (o *OswStackDetailVO) GetHwVersionOk() (*string, bool)`

GetHwVersionOk returns a tuple with the HwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwVersion

`func (o *OswStackDetailVO) SetHwVersion(v string)`

SetHwVersion sets HwVersion field to given value.

### HasHwVersion

`func (o *OswStackDetailVO) HasHwVersion() bool`

HasHwVersion returns a boolean if a field has been set.

### GetId

`func (o *OswStackDetailVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OswStackDetailVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OswStackDetailVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OswStackDetailVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInWhitelist

`func (o *OswStackDetailVO) GetInWhitelist() bool`

GetInWhitelist returns the InWhitelist field if non-nil, zero value otherwise.

### GetInWhitelistOk

`func (o *OswStackDetailVO) GetInWhitelistOk() (*bool, bool)`

GetInWhitelistOk returns a tuple with the InWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInWhitelist

`func (o *OswStackDetailVO) SetInWhitelist(v bool)`

SetInWhitelist sets InWhitelist field to given value.

### HasInWhitelist

`func (o *OswStackDetailVO) HasInWhitelist() bool`

HasInWhitelist returns a boolean if a field has been set.

### GetInactiveUnits

`func (o *OswStackDetailVO) GetInactiveUnits() []int32`

GetInactiveUnits returns the InactiveUnits field if non-nil, zero value otherwise.

### GetInactiveUnitsOk

`func (o *OswStackDetailVO) GetInactiveUnitsOk() (*[]int32, bool)`

GetInactiveUnitsOk returns a tuple with the InactiveUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveUnits

`func (o *OswStackDetailVO) SetInactiveUnits(v []int32)`

SetInactiveUnits sets InactiveUnits field to given value.

### HasInactiveUnits

`func (o *OswStackDetailVO) HasInactiveUnits() bool`

HasInactiveUnits returns a boolean if a field has been set.

### GetInitialUnbindingLimit

`func (o *OswStackDetailVO) GetInitialUnbindingLimit() int32`

GetInitialUnbindingLimit returns the InitialUnbindingLimit field if non-nil, zero value otherwise.

### GetInitialUnbindingLimitOk

`func (o *OswStackDetailVO) GetInitialUnbindingLimitOk() (*int32, bool)`

GetInitialUnbindingLimitOk returns a tuple with the InitialUnbindingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialUnbindingLimit

`func (o *OswStackDetailVO) SetInitialUnbindingLimit(v int32)`

SetInitialUnbindingLimit sets InitialUnbindingLimit field to given value.

### HasInitialUnbindingLimit

`func (o *OswStackDetailVO) HasInitialUnbindingLimit() bool`

HasInitialUnbindingLimit returns a boolean if a field has been set.

### GetIp

`func (o *OswStackDetailVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OswStackDetailVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OswStackDetailVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OswStackDetailVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpSetting

`func (o *OswStackDetailVO) GetIpSetting() IpSettingVO`

GetIpSetting returns the IpSetting field if non-nil, zero value otherwise.

### GetIpSettingOk

`func (o *OswStackDetailVO) GetIpSettingOk() (*IpSettingVO, bool)`

GetIpSettingOk returns a tuple with the IpSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSetting

`func (o *OswStackDetailVO) SetIpSetting(v IpSettingVO)`

SetIpSetting sets IpSetting field to given value.

### HasIpSetting

`func (o *OswStackDetailVO) HasIpSetting() bool`

HasIpSetting returns a boolean if a field has been set.

### GetIpv6List

`func (o *OswStackDetailVO) GetIpv6List() []string`

GetIpv6List returns the Ipv6List field if non-nil, zero value otherwise.

### GetIpv6ListOk

`func (o *OswStackDetailVO) GetIpv6ListOk() (*[]string, bool)`

GetIpv6ListOk returns a tuple with the Ipv6List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6List

`func (o *OswStackDetailVO) SetIpv6List(v []string)`

SetIpv6List sets Ipv6List field to given value.

### HasIpv6List

`func (o *OswStackDetailVO) HasIpv6List() bool`

HasIpv6List returns a boolean if a field has been set.

### GetJumbo

`func (o *OswStackDetailVO) GetJumbo() int32`

GetJumbo returns the Jumbo field if non-nil, zero value otherwise.

### GetJumboOk

`func (o *OswStackDetailVO) GetJumboOk() (*int32, bool)`

GetJumboOk returns a tuple with the Jumbo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumbo

`func (o *OswStackDetailVO) SetJumbo(v int32)`

SetJumbo sets Jumbo field to given value.

### HasJumbo

`func (o *OswStackDetailVO) HasJumbo() bool`

HasJumbo returns a boolean if a field has been set.

### GetLagHashAlg

`func (o *OswStackDetailVO) GetLagHashAlg() int32`

GetLagHashAlg returns the LagHashAlg field if non-nil, zero value otherwise.

### GetLagHashAlgOk

`func (o *OswStackDetailVO) GetLagHashAlgOk() (*int32, bool)`

GetLagHashAlgOk returns a tuple with the LagHashAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagHashAlg

`func (o *OswStackDetailVO) SetLagHashAlg(v int32)`

SetLagHashAlg sets LagHashAlg field to given value.

### HasLagHashAlg

`func (o *OswStackDetailVO) HasLagHashAlg() bool`

HasLagHashAlg returns a boolean if a field has been set.

### GetLags

`func (o *OswStackDetailVO) GetLags() []OswLagVO`

GetLags returns the Lags field if non-nil, zero value otherwise.

### GetLagsOk

`func (o *OswStackDetailVO) GetLagsOk() (*[]OswLagVO, bool)`

GetLagsOk returns a tuple with the Lags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLags

`func (o *OswStackDetailVO) SetLags(v []OswLagVO)`

SetLags sets Lags field to given value.

### HasLags

`func (o *OswStackDetailVO) HasLags() bool`

HasLags returns a boolean if a field has been set.

### GetLastSeen

`func (o *OswStackDetailVO) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *OswStackDetailVO) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *OswStackDetailVO) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *OswStackDetailVO) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLedSetting

`func (o *OswStackDetailVO) GetLedSetting() int32`

GetLedSetting returns the LedSetting field if non-nil, zero value otherwise.

### GetLedSettingOk

`func (o *OswStackDetailVO) GetLedSettingOk() (*int32, bool)`

GetLedSettingOk returns a tuple with the LedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedSetting

`func (o *OswStackDetailVO) SetLedSetting(v int32)`

SetLedSetting sets LedSetting field to given value.

### HasLedSetting

`func (o *OswStackDetailVO) HasLedSetting() bool`

HasLedSetting returns a boolean if a field has been set.

### GetLicenseId

`func (o *OswStackDetailVO) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *OswStackDetailVO) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *OswStackDetailVO) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *OswStackDetailVO) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetLicenseStatus

`func (o *OswStackDetailVO) GetLicenseStatus() int32`

GetLicenseStatus returns the LicenseStatus field if non-nil, zero value otherwise.

### GetLicenseStatusOk

`func (o *OswStackDetailVO) GetLicenseStatusOk() (*int32, bool)`

GetLicenseStatusOk returns a tuple with the LicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatus

`func (o *OswStackDetailVO) SetLicenseStatus(v int32)`

SetLicenseStatus sets LicenseStatus field to given value.

### HasLicenseStatus

`func (o *OswStackDetailVO) HasLicenseStatus() bool`

HasLicenseStatus returns a boolean if a field has been set.

### GetLicenseUnbindingLimit

`func (o *OswStackDetailVO) GetLicenseUnbindingLimit() int32`

GetLicenseUnbindingLimit returns the LicenseUnbindingLimit field if non-nil, zero value otherwise.

### GetLicenseUnbindingLimitOk

`func (o *OswStackDetailVO) GetLicenseUnbindingLimitOk() (*int32, bool)`

GetLicenseUnbindingLimitOk returns a tuple with the LicenseUnbindingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUnbindingLimit

`func (o *OswStackDetailVO) SetLicenseUnbindingLimit(v int32)`

SetLicenseUnbindingLimit sets LicenseUnbindingLimit field to given value.

### HasLicenseUnbindingLimit

`func (o *OswStackDetailVO) HasLicenseUnbindingLimit() bool`

HasLicenseUnbindingLimit returns a boolean if a field has been set.

### GetLocateEnable

`func (o *OswStackDetailVO) GetLocateEnable() bool`

GetLocateEnable returns the LocateEnable field if non-nil, zero value otherwise.

### GetLocateEnableOk

`func (o *OswStackDetailVO) GetLocateEnableOk() (*bool, bool)`

GetLocateEnableOk returns a tuple with the LocateEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocateEnable

`func (o *OswStackDetailVO) SetLocateEnable(v bool)`

SetLocateEnable sets LocateEnable field to given value.

### HasLocateEnable

`func (o *OswStackDetailVO) HasLocateEnable() bool`

HasLocateEnable returns a boolean if a field has been set.

### GetLocation

`func (o *OswStackDetailVO) GetLocation() DeviceLocationDetailVO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OswStackDetailVO) GetLocationOk() (*DeviceLocationDetailVO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OswStackDetailVO) SetLocation(v DeviceLocationDetailVO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OswStackDetailVO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLoop

`func (o *OswStackDetailVO) GetLoop() string`

GetLoop returns the Loop field if non-nil, zero value otherwise.

### GetLoopOk

`func (o *OswStackDetailVO) GetLoopOk() (*string, bool)`

GetLoopOk returns a tuple with the Loop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoop

`func (o *OswStackDetailVO) SetLoop(v string)`

SetLoop sets Loop field to given value.

### HasLoop

`func (o *OswStackDetailVO) HasLoop() bool`

HasLoop returns a boolean if a field has been set.

### GetLoopbackDetectEnable

`func (o *OswStackDetailVO) GetLoopbackDetectEnable() bool`

GetLoopbackDetectEnable returns the LoopbackDetectEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectEnableOk

`func (o *OswStackDetailVO) GetLoopbackDetectEnableOk() (*bool, bool)`

GetLoopbackDetectEnableOk returns a tuple with the LoopbackDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectEnable

`func (o *OswStackDetailVO) SetLoopbackDetectEnable(v bool)`

SetLoopbackDetectEnable sets LoopbackDetectEnable field to given value.

### HasLoopbackDetectEnable

`func (o *OswStackDetailVO) HasLoopbackDetectEnable() bool`

HasLoopbackDetectEnable returns a boolean if a field has been set.

### GetLoopbackNum

`func (o *OswStackDetailVO) GetLoopbackNum() int32`

GetLoopbackNum returns the LoopbackNum field if non-nil, zero value otherwise.

### GetLoopbackNumOk

`func (o *OswStackDetailVO) GetLoopbackNumOk() (*int32, bool)`

GetLoopbackNumOk returns a tuple with the LoopbackNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackNum

`func (o *OswStackDetailVO) SetLoopbackNum(v int32)`

SetLoopbackNum sets LoopbackNum field to given value.

### HasLoopbackNum

`func (o *OswStackDetailVO) HasLoopbackNum() bool`

HasLoopbackNum returns a boolean if a field has been set.

### GetMac

`func (o *OswStackDetailVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswStackDetailVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswStackDetailVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswStackDetailVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMacDelay

`func (o *OswStackDetailVO) GetMacDelay() MacDelayVO`

GetMacDelay returns the MacDelay field if non-nil, zero value otherwise.

### GetMacDelayOk

`func (o *OswStackDetailVO) GetMacDelayOk() (*MacDelayVO, bool)`

GetMacDelayOk returns a tuple with the MacDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacDelay

`func (o *OswStackDetailVO) SetMacDelay(v MacDelayVO)`

SetMacDelay sets MacDelay field to given value.

### HasMacDelay

`func (o *OswStackDetailVO) HasMacDelay() bool`

HasMacDelay returns a boolean if a field has been set.

### GetMasterActive

`func (o *OswStackDetailVO) GetMasterActive() bool`

GetMasterActive returns the MasterActive field if non-nil, zero value otherwise.

### GetMasterActiveOk

`func (o *OswStackDetailVO) GetMasterActiveOk() (*bool, bool)`

GetMasterActiveOk returns a tuple with the MasterActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterActive

`func (o *OswStackDetailVO) SetMasterActive(v bool)`

SetMasterActive sets MasterActive field to given value.

### HasMasterActive

`func (o *OswStackDetailVO) HasMasterActive() bool`

HasMasterActive returns a boolean if a field has been set.

### GetMasterMac

`func (o *OswStackDetailVO) GetMasterMac() string`

GetMasterMac returns the MasterMac field if non-nil, zero value otherwise.

### GetMasterMacOk

`func (o *OswStackDetailVO) GetMasterMacOk() (*string, bool)`

GetMasterMacOk returns a tuple with the MasterMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterMac

`func (o *OswStackDetailVO) SetMasterMac(v string)`

SetMasterMac sets MasterMac field to given value.

### HasMasterMac

`func (o *OswStackDetailVO) HasMasterMac() bool`

HasMasterMac returns a boolean if a field has been set.

### GetMaxAge

`func (o *OswStackDetailVO) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *OswStackDetailVO) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *OswStackDetailVO) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *OswStackDetailVO) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetMaxHops

`func (o *OswStackDetailVO) GetMaxHops() int32`

GetMaxHops returns the MaxHops field if non-nil, zero value otherwise.

### GetMaxHopsOk

`func (o *OswStackDetailVO) GetMaxHopsOk() (*int32, bool)`

GetMaxHopsOk returns a tuple with the MaxHops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHops

`func (o *OswStackDetailVO) SetMaxHops(v int32)`

SetMaxHops sets MaxHops field to given value.

### HasMaxHops

`func (o *OswStackDetailVO) HasMaxHops() bool`

HasMaxHops returns a boolean if a field has been set.

### GetMemUitl

`func (o *OswStackDetailVO) GetMemUitl() int32`

GetMemUitl returns the MemUitl field if non-nil, zero value otherwise.

### GetMemUitlOk

`func (o *OswStackDetailVO) GetMemUitlOk() (*int32, bool)`

GetMemUitlOk returns a tuple with the MemUitl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUitl

`func (o *OswStackDetailVO) SetMemUitl(v int32)`

SetMemUitl sets MemUitl field to given value.

### HasMemUitl

`func (o *OswStackDetailVO) HasMemUitl() bool`

HasMemUitl returns a boolean if a field has been set.

### GetMember

`func (o *OswStackDetailVO) GetMember() []OswStackMemberVO`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *OswStackDetailVO) GetMemberOk() (*[]OswStackMemberVO, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *OswStackDetailVO) SetMember(v []OswStackMemberVO)`

SetMember sets Member field to given value.

### HasMember

`func (o *OswStackDetailVO) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetMlagMsg

`func (o *OswStackDetailVO) GetMlagMsg() MlagMsgVO`

GetMlagMsg returns the MlagMsg field if non-nil, zero value otherwise.

### GetMlagMsgOk

`func (o *OswStackDetailVO) GetMlagMsgOk() (*MlagMsgVO, bool)`

GetMlagMsgOk returns a tuple with the MlagMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagMsg

`func (o *OswStackDetailVO) SetMlagMsg(v MlagMsgVO)`

SetMlagMsg sets MlagMsg field to given value.

### HasMlagMsg

`func (o *OswStackDetailVO) HasMlagMsg() bool`

HasMlagMsg returns a boolean if a field has been set.

### GetMlagPeerInfo

`func (o *OswStackDetailVO) GetMlagPeerInfo() OswMlagPeerInfoVO`

GetMlagPeerInfo returns the MlagPeerInfo field if non-nil, zero value otherwise.

### GetMlagPeerInfoOk

`func (o *OswStackDetailVO) GetMlagPeerInfoOk() (*OswMlagPeerInfoVO, bool)`

GetMlagPeerInfoOk returns a tuple with the MlagPeerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPeerInfo

`func (o *OswStackDetailVO) SetMlagPeerInfo(v OswMlagPeerInfoVO)`

SetMlagPeerInfo sets MlagPeerInfo field to given value.

### HasMlagPeerInfo

`func (o *OswStackDetailVO) HasMlagPeerInfo() bool`

HasMlagPeerInfo returns a boolean if a field has been set.

### GetModel

`func (o *OswStackDetailVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OswStackDetailVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OswStackDetailVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OswStackDetailVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OswStackDetailVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OswStackDetailVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OswStackDetailVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OswStackDetailVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetMoveSiteId

`func (o *OswStackDetailVO) GetMoveSiteId() string`

GetMoveSiteId returns the MoveSiteId field if non-nil, zero value otherwise.

### GetMoveSiteIdOk

`func (o *OswStackDetailVO) GetMoveSiteIdOk() (*string, bool)`

GetMoveSiteIdOk returns a tuple with the MoveSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveSiteId

`func (o *OswStackDetailVO) SetMoveSiteId(v string)`

SetMoveSiteId sets MoveSiteId field to given value.

### HasMoveSiteId

`func (o *OswStackDetailVO) HasMoveSiteId() bool`

HasMoveSiteId returns a boolean if a field has been set.

### GetMstp

`func (o *OswStackDetailVO) GetMstp() OswStpMstpVO`

GetMstp returns the Mstp field if non-nil, zero value otherwise.

### GetMstpOk

`func (o *OswStackDetailVO) GetMstpOk() (*OswStpMstpVO, bool)`

GetMstpOk returns a tuple with the Mstp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMstp

`func (o *OswStackDetailVO) SetMstp(v OswStpMstpVO)`

SetMstp sets Mstp field to given value.

### HasMstp

`func (o *OswStackDetailVO) HasMstp() bool`

HasMstp returns a boolean if a field has been set.

### GetMulticast

`func (o *OswStackDetailVO) GetMulticast() OswLanMulticastVO`

GetMulticast returns the Multicast field if non-nil, zero value otherwise.

### GetMulticastOk

`func (o *OswStackDetailVO) GetMulticastOk() (*OswLanMulticastVO, bool)`

GetMulticastOk returns a tuple with the Multicast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticast

`func (o *OswStackDetailVO) SetMulticast(v OswLanMulticastVO)`

SetMulticast sets Multicast field to given value.

### HasMulticast

`func (o *OswStackDetailVO) HasMulticast() bool`

HasMulticast returns a boolean if a field has been set.

### GetMvlanBridgeVlan

`func (o *OswStackDetailVO) GetMvlanBridgeVlan() int32`

GetMvlanBridgeVlan returns the MvlanBridgeVlan field if non-nil, zero value otherwise.

### GetMvlanBridgeVlanOk

`func (o *OswStackDetailVO) GetMvlanBridgeVlanOk() (*int32, bool)`

GetMvlanBridgeVlanOk returns a tuple with the MvlanBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMvlanBridgeVlan

`func (o *OswStackDetailVO) SetMvlanBridgeVlan(v int32)`

SetMvlanBridgeVlan sets MvlanBridgeVlan field to given value.

### HasMvlanBridgeVlan

`func (o *OswStackDetailVO) HasMvlanBridgeVlan() bool`

HasMvlanBridgeVlan returns a boolean if a field has been set.

### GetMvlanNetworkId

`func (o *OswStackDetailVO) GetMvlanNetworkId() string`

GetMvlanNetworkId returns the MvlanNetworkId field if non-nil, zero value otherwise.

### GetMvlanNetworkIdOk

`func (o *OswStackDetailVO) GetMvlanNetworkIdOk() (*string, bool)`

GetMvlanNetworkIdOk returns a tuple with the MvlanNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMvlanNetworkId

`func (o *OswStackDetailVO) SetMvlanNetworkId(v string)`

SetMvlanNetworkId sets MvlanNetworkId field to given value.

### HasMvlanNetworkId

`func (o *OswStackDetailVO) HasMvlanNetworkId() bool`

HasMvlanNetworkId returns a boolean if a field has been set.

### GetName

`func (o *OswStackDetailVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswStackDetailVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswStackDetailVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswStackDetailVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedUpgrade

`func (o *OswStackDetailVO) GetNeedUpgrade() bool`

GetNeedUpgrade returns the NeedUpgrade field if non-nil, zero value otherwise.

### GetNeedUpgradeOk

`func (o *OswStackDetailVO) GetNeedUpgradeOk() (*bool, bool)`

GetNeedUpgradeOk returns a tuple with the NeedUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedUpgrade

`func (o *OswStackDetailVO) SetNeedUpgrade(v bool)`

SetNeedUpgrade sets NeedUpgrade field to given value.

### HasNeedUpgrade

`func (o *OswStackDetailVO) HasNeedUpgrade() bool`

HasNeedUpgrade returns a boolean if a field has been set.

### GetOmadacId

`func (o *OswStackDetailVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *OswStackDetailVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *OswStackDetailVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *OswStackDetailVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetPoeRemain

`func (o *OswStackDetailVO) GetPoeRemain() float64`

GetPoeRemain returns the PoeRemain field if non-nil, zero value otherwise.

### GetPoeRemainOk

`func (o *OswStackDetailVO) GetPoeRemainOk() (*float64, bool)`

GetPoeRemainOk returns a tuple with the PoeRemain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeRemain

`func (o *OswStackDetailVO) SetPoeRemain(v float64)`

SetPoeRemain sets PoeRemain field to given value.

### HasPoeRemain

`func (o *OswStackDetailVO) HasPoeRemain() bool`

HasPoeRemain returns a boolean if a field has been set.

### GetPoeRemainPercent

`func (o *OswStackDetailVO) GetPoeRemainPercent() float64`

GetPoeRemainPercent returns the PoeRemainPercent field if non-nil, zero value otherwise.

### GetPoeRemainPercentOk

`func (o *OswStackDetailVO) GetPoeRemainPercentOk() (*float64, bool)`

GetPoeRemainPercentOk returns a tuple with the PoeRemainPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeRemainPercent

`func (o *OswStackDetailVO) SetPoeRemainPercent(v float64)`

SetPoeRemainPercent sets PoeRemainPercent field to given value.

### HasPoeRemainPercent

`func (o *OswStackDetailVO) HasPoeRemainPercent() bool`

HasPoeRemainPercent returns a boolean if a field has been set.

### GetPoeTotalPower

`func (o *OswStackDetailVO) GetPoeTotalPower() float64`

GetPoeTotalPower returns the PoeTotalPower field if non-nil, zero value otherwise.

### GetPoeTotalPowerOk

`func (o *OswStackDetailVO) GetPoeTotalPowerOk() (*float64, bool)`

GetPoeTotalPowerOk returns a tuple with the PoeTotalPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeTotalPower

`func (o *OswStackDetailVO) SetPoeTotalPower(v float64)`

SetPoeTotalPower sets PoeTotalPower field to given value.

### HasPoeTotalPower

`func (o *OswStackDetailVO) HasPoeTotalPower() bool`

HasPoeTotalPower returns a boolean if a field has been set.

### GetPorts

`func (o *OswStackDetailVO) GetPorts() []OswPortVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OswStackDetailVO) GetPortsOk() (*[]OswPortVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OswStackDetailVO) SetPorts(v []OswPortVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OswStackDetailVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPriority

`func (o *OswStackDetailVO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *OswStackDetailVO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *OswStackDetailVO) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *OswStackDetailVO) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPublicIp

`func (o *OswStackDetailVO) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *OswStackDetailVO) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *OswStackDetailVO) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *OswStackDetailVO) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetQosConfig

`func (o *OswStackDetailVO) GetQosConfig() OswQosConfigVO`

GetQosConfig returns the QosConfig field if non-nil, zero value otherwise.

### GetQosConfigOk

`func (o *OswStackDetailVO) GetQosConfigOk() (*OswQosConfigVO, bool)`

GetQosConfigOk returns a tuple with the QosConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosConfig

`func (o *OswStackDetailVO) SetQosConfig(v OswQosConfigVO)`

SetQosConfig sets QosConfig field to given value.

### HasQosConfig

`func (o *OswStackDetailVO) HasQosConfig() bool`

HasQosConfig returns a boolean if a field has been set.

### GetRemember

`func (o *OswStackDetailVO) GetRemember() bool`

GetRemember returns the Remember field if non-nil, zero value otherwise.

### GetRememberOk

`func (o *OswStackDetailVO) GetRememberOk() (*bool, bool)`

GetRememberOk returns a tuple with the Remember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemember

`func (o *OswStackDetailVO) SetRemember(v bool)`

SetRemember sets Remember field to given value.

### HasRemember

`func (o *OswStackDetailVO) HasRemember() bool`

HasRemember returns a boolean if a field has been set.

### GetRememberDevice

`func (o *OswStackDetailVO) GetRememberDevice() int32`

GetRememberDevice returns the RememberDevice field if non-nil, zero value otherwise.

### GetRememberDeviceOk

`func (o *OswStackDetailVO) GetRememberDeviceOk() (*int32, bool)`

GetRememberDeviceOk returns a tuple with the RememberDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberDevice

`func (o *OswStackDetailVO) SetRememberDevice(v int32)`

SetRememberDevice sets RememberDevice field to given value.

### HasRememberDevice

`func (o *OswStackDetailVO) HasRememberDevice() bool`

HasRememberDevice returns a boolean if a field has been set.

### GetResource

`func (o *OswStackDetailVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *OswStackDetailVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *OswStackDetailVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *OswStackDetailVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetRxRate

`func (o *OswStackDetailVO) GetRxRate() int64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *OswStackDetailVO) GetRxRateOk() (*int64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *OswStackDetailVO) SetRxRate(v int64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *OswStackDetailVO) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetSdm

`func (o *OswStackDetailVO) GetSdm() OswSdmTemplateVO`

GetSdm returns the Sdm field if non-nil, zero value otherwise.

### GetSdmOk

`func (o *OswStackDetailVO) GetSdmOk() (*OswSdmTemplateVO, bool)`

GetSdmOk returns a tuple with the Sdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdm

`func (o *OswStackDetailVO) SetSdm(v OswSdmTemplateVO)`

SetSdm sets Sdm field to given value.

### HasSdm

`func (o *OswStackDetailVO) HasSdm() bool`

HasSdm returns a boolean if a field has been set.

### GetShowModel

`func (o *OswStackDetailVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *OswStackDetailVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *OswStackDetailVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *OswStackDetailVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetSiteId

`func (o *OswStackDetailVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *OswStackDetailVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *OswStackDetailVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *OswStackDetailVO) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteName

`func (o *OswStackDetailVO) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *OswStackDetailVO) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *OswStackDetailVO) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *OswStackDetailVO) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSiteTemplateId

`func (o *OswStackDetailVO) GetSiteTemplateId() string`

GetSiteTemplateId returns the SiteTemplateId field if non-nil, zero value otherwise.

### GetSiteTemplateIdOk

`func (o *OswStackDetailVO) GetSiteTemplateIdOk() (*string, bool)`

GetSiteTemplateIdOk returns a tuple with the SiteTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTemplateId

`func (o *OswStackDetailVO) SetSiteTemplateId(v string)`

SetSiteTemplateId sets SiteTemplateId field to given value.

### HasSiteTemplateId

`func (o *OswStackDetailVO) HasSiteTemplateId() bool`

HasSiteTemplateId returns a boolean if a field has been set.

### GetSiteTemplateName

`func (o *OswStackDetailVO) GetSiteTemplateName() string`

GetSiteTemplateName returns the SiteTemplateName field if non-nil, zero value otherwise.

### GetSiteTemplateNameOk

`func (o *OswStackDetailVO) GetSiteTemplateNameOk() (*string, bool)`

GetSiteTemplateNameOk returns a tuple with the SiteTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTemplateName

`func (o *OswStackDetailVO) SetSiteTemplateName(v string)`

SetSiteTemplateName sets SiteTemplateName field to given value.

### HasSiteTemplateName

`func (o *OswStackDetailVO) HasSiteTemplateName() bool`

HasSiteTemplateName returns a boolean if a field has been set.

### GetSn

`func (o *OswStackDetailVO) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *OswStackDetailVO) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *OswStackDetailVO) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *OswStackDetailVO) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetSnmp

`func (o *OswStackDetailVO) GetSnmp() OswSnmpVO`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *OswStackDetailVO) GetSnmpOk() (*OswSnmpVO, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *OswStackDetailVO) SetSnmp(v OswSnmpVO)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *OswStackDetailVO) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetSpecialModel

`func (o *OswStackDetailVO) GetSpecialModel() string`

GetSpecialModel returns the SpecialModel field if non-nil, zero value otherwise.

### GetSpecialModelOk

`func (o *OswStackDetailVO) GetSpecialModelOk() (*string, bool)`

GetSpecialModelOk returns a tuple with the SpecialModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialModel

`func (o *OswStackDetailVO) SetSpecialModel(v string)`

SetSpecialModel sets SpecialModel field to given value.

### HasSpecialModel

`func (o *OswStackDetailVO) HasSpecialModel() bool`

HasSpecialModel returns a boolean if a field has been set.

### GetSpeeds

`func (o *OswStackDetailVO) GetSpeeds() []int32`

GetSpeeds returns the Speeds field if non-nil, zero value otherwise.

### GetSpeedsOk

`func (o *OswStackDetailVO) GetSpeedsOk() (*[]int32, bool)`

GetSpeedsOk returns a tuple with the Speeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeeds

`func (o *OswStackDetailVO) SetSpeeds(v []int32)`

SetSpeeds sets Speeds field to given value.

### HasSpeeds

`func (o *OswStackDetailVO) HasSpeeds() bool`

HasSpeeds returns a boolean if a field has been set.

### GetStackDevice

`func (o *OswStackDetailVO) GetStackDevice() bool`

GetStackDevice returns the StackDevice field if non-nil, zero value otherwise.

### GetStackDeviceOk

`func (o *OswStackDetailVO) GetStackDeviceOk() (*bool, bool)`

GetStackDeviceOk returns a tuple with the StackDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackDevice

`func (o *OswStackDetailVO) SetStackDevice(v bool)`

SetStackDevice sets StackDevice field to given value.

### HasStackDevice

`func (o *OswStackDetailVO) HasStackDevice() bool`

HasStackDevice returns a boolean if a field has been set.

### GetStackLags

`func (o *OswStackDetailVO) GetStackLags() []OswStackMemberLagVO`

GetStackLags returns the StackLags field if non-nil, zero value otherwise.

### GetStackLagsOk

`func (o *OswStackDetailVO) GetStackLagsOk() (*[]OswStackMemberLagVO, bool)`

GetStackLagsOk returns a tuple with the StackLags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackLags

`func (o *OswStackDetailVO) SetStackLags(v []OswStackMemberLagVO)`

SetStackLags sets StackLags field to given value.

### HasStackLags

`func (o *OswStackDetailVO) HasStackLags() bool`

HasStackLags returns a boolean if a field has been set.

### GetStackMsg

`func (o *OswStackDetailVO) GetStackMsg() StackMsgVO`

GetStackMsg returns the StackMsg field if non-nil, zero value otherwise.

### GetStackMsgOk

`func (o *OswStackDetailVO) GetStackMsgOk() (*StackMsgVO, bool)`

GetStackMsgOk returns a tuple with the StackMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackMsg

`func (o *OswStackDetailVO) SetStackMsg(v StackMsgVO)`

SetStackMsg sets StackMsg field to given value.

### HasStackMsg

`func (o *OswStackDetailVO) HasStackMsg() bool`

HasStackMsg returns a boolean if a field has been set.

### GetStackPorts

`func (o *OswStackDetailVO) GetStackPorts() []OswStackPortGroupVO`

GetStackPorts returns the StackPorts field if non-nil, zero value otherwise.

### GetStackPortsOk

`func (o *OswStackDetailVO) GetStackPortsOk() (*[]OswStackPortGroupVO, bool)`

GetStackPortsOk returns a tuple with the StackPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPorts

`func (o *OswStackDetailVO) SetStackPorts(v []OswStackPortGroupVO)`

SetStackPorts sets StackPorts field to given value.

### HasStackPorts

`func (o *OswStackDetailVO) HasStackPorts() bool`

HasStackPorts returns a boolean if a field has been set.

### GetStatus

`func (o *OswStackDetailVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswStackDetailVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswStackDetailVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OswStackDetailVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *OswStackDetailVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *OswStackDetailVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *OswStackDetailVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *OswStackDetailVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetStp

`func (o *OswStackDetailVO) GetStp() int32`

GetStp returns the Stp field if non-nil, zero value otherwise.

### GetStpOk

`func (o *OswStackDetailVO) GetStpOk() (*int32, bool)`

GetStpOk returns a tuple with the Stp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStp

`func (o *OswStackDetailVO) SetStp(v int32)`

SetStp sets Stp field to given value.

### HasStp

`func (o *OswStackDetailVO) HasStp() bool`

HasStp returns a boolean if a field has been set.

### GetStpLinkList

`func (o *OswStackDetailVO) GetStpLinkList() []OswDownlinkVO`

GetStpLinkList returns the StpLinkList field if non-nil, zero value otherwise.

### GetStpLinkListOk

`func (o *OswStackDetailVO) GetStpLinkListOk() (*[]OswDownlinkVO, bool)`

GetStpLinkListOk returns a tuple with the StpLinkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpLinkList

`func (o *OswStackDetailVO) SetStpLinkList(v []OswDownlinkVO)`

SetStpLinkList sets StpLinkList field to given value.

### HasStpLinkList

`func (o *OswStackDetailVO) HasStpLinkList() bool`

HasStpLinkList returns a boolean if a field has been set.

### GetSupportAnomaly

`func (o *OswStackDetailVO) GetSupportAnomaly() bool`

GetSupportAnomaly returns the SupportAnomaly field if non-nil, zero value otherwise.

### GetSupportAnomalyOk

`func (o *OswStackDetailVO) GetSupportAnomalyOk() (*bool, bool)`

GetSupportAnomalyOk returns a tuple with the SupportAnomaly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAnomaly

`func (o *OswStackDetailVO) SetSupportAnomaly(v bool)`

SetSupportAnomaly sets SupportAnomaly field to given value.

### HasSupportAnomaly

`func (o *OswStackDetailVO) HasSupportAnomaly() bool`

HasSupportAnomaly returns a boolean if a field has been set.

### GetSupportHealth

`func (o *OswStackDetailVO) GetSupportHealth() bool`

GetSupportHealth returns the SupportHealth field if non-nil, zero value otherwise.

### GetSupportHealthOk

`func (o *OswStackDetailVO) GetSupportHealthOk() (*bool, bool)`

GetSupportHealthOk returns a tuple with the SupportHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportHealth

`func (o *OswStackDetailVO) SetSupportHealth(v bool)`

SetSupportHealth sets SupportHealth field to given value.

### HasSupportHealth

`func (o *OswStackDetailVO) HasSupportHealth() bool`

HasSupportHealth returns a boolean if a field has been set.

### GetSupportIpv6Acl

`func (o *OswStackDetailVO) GetSupportIpv6Acl() bool`

GetSupportIpv6Acl returns the SupportIpv6Acl field if non-nil, zero value otherwise.

### GetSupportIpv6AclOk

`func (o *OswStackDetailVO) GetSupportIpv6AclOk() (*bool, bool)`

GetSupportIpv6AclOk returns a tuple with the SupportIpv6Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIpv6Acl

`func (o *OswStackDetailVO) SetSupportIpv6Acl(v bool)`

SetSupportIpv6Acl sets SupportIpv6Acl field to given value.

### HasSupportIpv6Acl

`func (o *OswStackDetailVO) HasSupportIpv6Acl() bool`

HasSupportIpv6Acl returns a boolean if a field has been set.

### GetSupportLocatePort

`func (o *OswStackDetailVO) GetSupportLocatePort() bool`

GetSupportLocatePort returns the SupportLocatePort field if non-nil, zero value otherwise.

### GetSupportLocatePortOk

`func (o *OswStackDetailVO) GetSupportLocatePortOk() (*bool, bool)`

GetSupportLocatePortOk returns a tuple with the SupportLocatePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLocatePort

`func (o *OswStackDetailVO) SetSupportLocatePort(v bool)`

SetSupportLocatePort sets SupportLocatePort field to given value.

### HasSupportLocatePort

`func (o *OswStackDetailVO) HasSupportLocatePort() bool`

HasSupportLocatePort returns a boolean if a field has been set.

### GetSupportVlanIf

`func (o *OswStackDetailVO) GetSupportVlanIf() bool`

GetSupportVlanIf returns the SupportVlanIf field if non-nil, zero value otherwise.

### GetSupportVlanIfOk

`func (o *OswStackDetailVO) GetSupportVlanIfOk() (*bool, bool)`

GetSupportVlanIfOk returns a tuple with the SupportVlanIf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVlanIf

`func (o *OswStackDetailVO) SetSupportVlanIf(v bool)`

SetSupportVlanIf sets SupportVlanIf field to given value.

### HasSupportVlanIf

`func (o *OswStackDetailVO) HasSupportVlanIf() bool`

HasSupportVlanIf returns a boolean if a field has been set.

### GetSupportVrf

`func (o *OswStackDetailVO) GetSupportVrf() bool`

GetSupportVrf returns the SupportVrf field if non-nil, zero value otherwise.

### GetSupportVrfOk

`func (o *OswStackDetailVO) GetSupportVrfOk() (*bool, bool)`

GetSupportVrfOk returns a tuple with the SupportVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVrf

`func (o *OswStackDetailVO) SetSupportVrf(v bool)`

SetSupportVrf sets SupportVrf field to given value.

### HasSupportVrf

`func (o *OswStackDetailVO) HasSupportVrf() bool`

HasSupportVrf returns a boolean if a field has been set.

### GetTagIds

`func (o *OswStackDetailVO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *OswStackDetailVO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *OswStackDetailVO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *OswStackDetailVO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTemplateId

`func (o *OswStackDetailVO) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *OswStackDetailVO) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *OswStackDetailVO) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *OswStackDetailVO) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateName

`func (o *OswStackDetailVO) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *OswStackDetailVO) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *OswStackDetailVO) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *OswStackDetailVO) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetTemplateSettings

`func (o *OswStackDetailVO) GetTemplateSettings() []int32`

GetTemplateSettings returns the TemplateSettings field if non-nil, zero value otherwise.

### GetTemplateSettingsOk

`func (o *OswStackDetailVO) GetTemplateSettingsOk() (*[]int32, bool)`

GetTemplateSettingsOk returns a tuple with the TemplateSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSettings

`func (o *OswStackDetailVO) SetTemplateSettings(v []int32)`

SetTemplateSettings sets TemplateSettings field to given value.

### HasTemplateSettings

`func (o *OswStackDetailVO) HasTemplateSettings() bool`

HasTemplateSettings returns a boolean if a field has been set.

### GetTerminalPrefix

`func (o *OswStackDetailVO) GetTerminalPrefix() string`

GetTerminalPrefix returns the TerminalPrefix field if non-nil, zero value otherwise.

### GetTerminalPrefixOk

`func (o *OswStackDetailVO) GetTerminalPrefixOk() (*string, bool)`

GetTerminalPrefixOk returns a tuple with the TerminalPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalPrefix

`func (o *OswStackDetailVO) SetTerminalPrefix(v string)`

SetTerminalPrefix sets TerminalPrefix field to given value.

### HasTerminalPrefix

`func (o *OswStackDetailVO) HasTerminalPrefix() bool`

HasTerminalPrefix returns a boolean if a field has been set.

### GetTxHoldCount

`func (o *OswStackDetailVO) GetTxHoldCount() int32`

GetTxHoldCount returns the TxHoldCount field if non-nil, zero value otherwise.

### GetTxHoldCountOk

`func (o *OswStackDetailVO) GetTxHoldCountOk() (*int32, bool)`

GetTxHoldCountOk returns a tuple with the TxHoldCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHoldCount

`func (o *OswStackDetailVO) SetTxHoldCount(v int32)`

SetTxHoldCount sets TxHoldCount field to given value.

### HasTxHoldCount

`func (o *OswStackDetailVO) HasTxHoldCount() bool`

HasTxHoldCount returns a boolean if a field has been set.

### GetTxRate

`func (o *OswStackDetailVO) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *OswStackDetailVO) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *OswStackDetailVO) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *OswStackDetailVO) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.

### GetType

`func (o *OswStackDetailVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswStackDetailVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswStackDetailVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OswStackDetailVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *OswStackDetailVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OswStackDetailVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OswStackDetailVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *OswStackDetailVO) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUplink

`func (o *OswStackDetailVO) GetUplink() OswUplinkVO`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *OswStackDetailVO) GetUplinkOk() (*OswUplinkVO, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *OswStackDetailVO) SetUplink(v OswUplinkVO)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *OswStackDetailVO) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetUplinkPort

`func (o *OswStackDetailVO) GetUplinkPort() int32`

GetUplinkPort returns the UplinkPort field if non-nil, zero value otherwise.

### GetUplinkPortOk

`func (o *OswStackDetailVO) GetUplinkPortOk() (*int32, bool)`

GetUplinkPortOk returns a tuple with the UplinkPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkPort

`func (o *OswStackDetailVO) SetUplinkPort(v int32)`

SetUplinkPort sets UplinkPort field to given value.

### HasUplinkPort

`func (o *OswStackDetailVO) HasUplinkPort() bool`

HasUplinkPort returns a boolean if a field has been set.

### GetUplinkStPort

`func (o *OswStackDetailVO) GetUplinkStPort() string`

GetUplinkStPort returns the UplinkStPort field if non-nil, zero value otherwise.

### GetUplinkStPortOk

`func (o *OswStackDetailVO) GetUplinkStPortOk() (*string, bool)`

GetUplinkStPortOk returns a tuple with the UplinkStPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkStPort

`func (o *OswStackDetailVO) SetUplinkStPort(v string)`

SetUplinkStPort sets UplinkStPort field to given value.

### HasUplinkStPort

`func (o *OswStackDetailVO) HasUplinkStPort() bool`

HasUplinkStPort returns a boolean if a field has been set.

### GetUpload

`func (o *OswStackDetailVO) GetUpload() int64`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *OswStackDetailVO) GetUploadOk() (*int64, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *OswStackDetailVO) SetUpload(v int64)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *OswStackDetailVO) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetUptime

`func (o *OswStackDetailVO) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *OswStackDetailVO) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *OswStackDetailVO) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *OswStackDetailVO) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetUptimeLong

`func (o *OswStackDetailVO) GetUptimeLong() int64`

GetUptimeLong returns the UptimeLong field if non-nil, zero value otherwise.

### GetUptimeLongOk

`func (o *OswStackDetailVO) GetUptimeLongOk() (*int64, bool)`

GetUptimeLongOk returns a tuple with the UptimeLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeLong

`func (o *OswStackDetailVO) SetUptimeLong(v int64)`

SetUptimeLong sets UptimeLong field to given value.

### HasUptimeLong

`func (o *OswStackDetailVO) HasUptimeLong() bool`

HasUptimeLong returns a boolean if a field has been set.

### GetVersion

`func (o *OswStackDetailVO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OswStackDetailVO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OswStackDetailVO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OswStackDetailVO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVirtualMac

`func (o *OswStackDetailVO) GetVirtualMac() string`

GetVirtualMac returns the VirtualMac field if non-nil, zero value otherwise.

### GetVirtualMacOk

`func (o *OswStackDetailVO) GetVirtualMacOk() (*string, bool)`

GetVirtualMacOk returns a tuple with the VirtualMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMac

`func (o *OswStackDetailVO) SetVirtualMac(v string)`

SetVirtualMac sets VirtualMac field to given value.

### HasVirtualMac

`func (o *OswStackDetailVO) HasVirtualMac() bool`

HasVirtualMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


