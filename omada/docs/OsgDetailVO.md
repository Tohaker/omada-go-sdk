# OsgDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | Site of the device | [optional] 
**Active** | Pointer to **bool** | whether to active the device(cloud base exclusive) | [optional] 
**AddByTemplate** | Pointer to **bool** |  | [optional] 
**AddedInAdvanced** | Pointer to **bool** | Whether the device is added in advanced. | [optional] 
**AdvancedResource** | Pointer to **int32** |  | [optional] 
**BoundDeviceTemplate** | Pointer to **bool** | Whether the device is bound to device template | [optional] 
**BoundSiteTemplate** | Pointer to **bool** | Whether the site where the device is located is bound to a site template | [optional] 
**Category** | Pointer to **string** | Category of license | [optional] 
**ChannelLimitType** | Pointer to **int32** |  | [optional] 
**CombinedGateway** | Pointer to **bool** |  | [optional] 
**CommonAdvancedResource** | Pointer to **int32** |  | [optional] 
**Compatible** | Pointer to **int32** | Device firmware and controller compatibility type.Compatible should be a value as follows: 0:COMPATIBLE;1:HIGH_MAJOR_VER;2:LOW_MAJOR_VER;3:HIGH_MINOR_VER;4:LOW_MINOR_VER;7:HIGH_COMPONENT_VER;10:DEVICE_NOT_COMPATIBLE;11:HIGH_ADOPT_COMMPONENT;12:DEVICE_CATEGORY_NOT_COMPATIBLE;14:DEVICE_NOT_COMPATIBLE_IN_CLUSTER | [optional] 
**CompoundModel** | Pointer to **string** | Model complex used in the backend.Ap：model+(country)+modelVersion,  EAP225(EU) v3.0 Ap: specialModel+modelVersion, EAP225-Outdoor-1a20a950b8d950e8 v1.0  Gateway/Switch：model+modelVersion, Osg v3.0 | [optional] 
**ContainNetType5G** | Pointer to **bool** |  | [optional] 
**ControllerId** | Pointer to **string** |  | [optional] 
**CpuUtil** | Pointer to **int32** |  | [optional] 
**CustomId** | Pointer to **string** | Customer ID | [optional] 
**CustomName** | Pointer to **string** | Customer name | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DevCap** | Pointer to [**OsgCapVO**](OsgCapVO.md) |  | [optional] 
**DeviceMisc** | Pointer to [**WirelessRouterMiscVO**](WirelessRouterMiscVO.md) |  | [optional] 
**DeviceSeriesType** | Pointer to **int32** | Device series type.DeviceSeriesType should be a value as follows: 0:advanced;1:pro | [optional] 
**DeviceTemplateAvailable** | Pointer to **bool** | Whether there is an available device template for the device; it is false if the model is not supported or the site template has not created the corresponding device template. | [optional] 
**DisableHwReset** | Pointer to **bool** |  | [optional] 
**DownlinkList** | Pointer to [**[]OsgDownLinkVO**](OsgDownLinkVO.md) |  | [optional] 
**Download** | Pointer to **int64** |  | [optional] 
**DueTime** | Pointer to **int64** | Expire timestamp of license(cloud base exclusive) | [optional] 
**DueTimeLeft** | Pointer to **int64** | Milliseconds from the current moment to the expiration time(cloud base exclusive) | [optional] 
**EchoServer** | Pointer to **string** |  | [optional] 
**EcspFirstVersion** | Pointer to **int32** | Ecsp first version | [optional] 
**Eos** | Pointer to **int64** | End of support time of device(CBC exclusive) | [optional] 
**Eost** | Pointer to **int64** | End of service time of device(CBC exclusive) | [optional] 
**Es** | Pointer to **bool** | Whether the device is Agile Series Switch | [optional] 
**Fan** | Pointer to [**[]OsgFanStatusVO**](OsgFanStatusVO.md) |  | [optional] 
**FirmwareVersion** | Pointer to **string** | Version of firmware,for example:2.5.0 Build 20190118 Rel. 64821 | [optional] 
**ForgetId** | Pointer to **string** | Forget ID of device | [optional] 
**GuestNum** | Pointer to **int32** |  | [optional] 
**HwOffloadEnable** | Pointer to **bool** |  | [optional] 
**HwVersion** | Pointer to **string** | Version of hardware,for example 1.0 | [optional] 
**InWhitelist** | Pointer to **bool** | Whether the device is in white list | [optional] 
**InitialUnbindingLimit** | Pointer to **int32** | Initial unbind count for license(cloud base exclusive) | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Ippt** | Pointer to **bool** |  | [optional] 
**IpptPreconfig** | Pointer to **bool** |  | [optional] 
**IptvSetting** | Pointer to [**OsgIptvVO**](OsgIptvVO.md) |  | [optional] 
**Ipv6List** | Pointer to **[]string** |  | [optional] 
**JumboOptions** | Pointer to **[]int32** |  | [optional] 
**JumboSize** | Pointer to **int32** |  | [optional] 
**LanClientStats** | Pointer to [**[]OsgLanStatVO**](OsgLanStatVO.md) |  | [optional] 
**LastSeen** | Pointer to **int64** |  | [optional] 
**LbSetting2g** | Pointer to [**ApLoadBalanceVO**](ApLoadBalanceVO.md) |  | [optional] 
**LbSetting5g** | Pointer to [**ApLoadBalanceVO**](ApLoadBalanceVO.md) |  | [optional] 
**LbSetting5g2** | Pointer to [**ApLoadBalanceVO**](ApLoadBalanceVO.md) |  | [optional] 
**LbSetting6g** | Pointer to [**ApLoadBalanceVO**](ApLoadBalanceVO.md) |  | [optional] 
**LedSetting** | Pointer to **int32** |  | [optional] 
**LicenseId** | Pointer to **string** | License key on detail page of device(cloud base exclusive) | [optional] 
**LicenseStatus** | Pointer to **int32** | License status(cloud base exclusive).LicenseStatus should be a value as follows: 0:unActive 1:Unbind 2:Expired 3:active | [optional] 
**LicenseUnbindingLimit** | Pointer to **int32** | Remaining unbind count for license on detail Page of device(cloud base exclusive) | [optional] 
**LldpEnable** | Pointer to **bool** |  | [optional] 
**LldpSetting** | Pointer to **int32** |  | [optional] 
**Location** | Pointer to [**DeviceLocationDetailVO**](DeviceLocationDetailVO.md) |  | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**MemUtil** | Pointer to **int32** |  | [optional] 
**Model** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**MoveSiteId** | Pointer to **string** | Record that the device is in a moveSite operation; if it is null, then it is not in the moveSite operation. | [optional] 
**MultiChipGateway** | Pointer to **bool** |  | [optional] 
**MultiChipInfos** | Pointer to **[][]int32** |  | [optional] 
**Name** | Pointer to **string** | Device name,default value is the mac address of device | [optional] 
**NeedUpgrade** | Pointer to **bool** |  | [optional] 
**NetworkComptent** | Pointer to **int32** |  | [optional] 
**NonPscEnable** | Pointer to **bool** |  | [optional] 
**OfdmaEnable2g** | Pointer to **bool** |  | [optional] 
**OfdmaEnable5g** | Pointer to **bool** |  | [optional] 
**OfdmaEnable5g2** | Pointer to **bool** |  | [optional] 
**OfdmaEnable6g** | Pointer to **bool** |  | [optional] 
**OmadacId** | Pointer to **string** | OmadacId of the device | [optional] 
**OsgCap** | Pointer to [**SiteSettingCapVO**](SiteSettingCapVO.md) |  | [optional] 
**OsgExist** | Pointer to [**ExistSiteSettingVO**](ExistSiteSettingVO.md) |  | [optional] 
**PinSetting** | Pointer to [**OsgLtePinSettingVO**](OsgLtePinSettingVO.md) |  | [optional] 
**PoeLimit** | Pointer to **float64** |  | [optional] 
**PoeRemain** | Pointer to **float64** |  | [optional] 
**PoeRemainPercent** | Pointer to **float64** |  | [optional] 
**PoeSettings** | Pointer to [**[]OsgPortPoeVO**](OsgPortPoeVO.md) |  | [optional] 
**PortConfigs** | Pointer to [**[]OsgPortConfigVO**](OsgPortConfigVO.md) |  | [optional] 
**PortNum** | Pointer to **int32** |  | [optional] 
**PortStats** | Pointer to [**[]OsgPortStatVO**](OsgPortStatVO.md) |  | [optional] 
**PublicIp** | Pointer to **string** |  | [optional] 
**QosSetting2g** | Pointer to [**ApQosVO**](ApQosVO.md) |  | [optional] 
**QosSetting5g** | Pointer to [**ApQosVO**](ApQosVO.md) |  | [optional] 
**QosSetting5g2** | Pointer to [**ApQosVO**](ApQosVO.md) |  | [optional] 
**QosSetting6g** | Pointer to [**ApQosVO**](ApQosVO.md) |  | [optional] 
**RadioSetting2g** | Pointer to [**ApRadioSetting**](ApRadioSetting.md) |  | [optional] 
**RadioSetting5g** | Pointer to [**ApRadioSetting**](ApRadioSetting.md) |  | [optional] 
**RadioSetting5g2** | Pointer to [**ApRadioSetting**](ApRadioSetting.md) |  | [optional] 
**RadioSetting6g** | Pointer to [**ApRadioSetting**](ApRadioSetting.md) |  | [optional] 
**RadioTraffic2g** | Pointer to [**APRadioTrafficEntity**](APRadioTrafficEntity.md) |  | [optional] 
**RadioTraffic5g** | Pointer to [**APRadioTrafficEntity**](APRadioTrafficEntity.md) |  | [optional] 
**RadioTraffic5g2** | Pointer to [**APRadioTrafficEntity**](APRadioTrafficEntity.md) |  | [optional] 
**RadioTraffic6g** | Pointer to [**APRadioTrafficEntity**](APRadioTrafficEntity.md) |  | [optional] 
**RadiosResource** | Pointer to **int32** |  | [optional] 
**Remember** | Pointer to **bool** | Whether to remember the device(deprecated) | [optional] 
**RememberDevice** | Pointer to **int32** | Whether to remember the device.RememberDevice should be a value as follows: 0:off, 1:on, 2: follow site | [optional] 
**Resource** | Pointer to **int32** | Data source.Resource should be a value as follows: 0:new created;1:from template;2:override | [optional] 
**Rps** | Pointer to [**[]OsgRpsStatusVO**](OsgRpsStatusVO.md) |  | [optional] 
**RssiSetting2g** | Pointer to [**ApRssiThresholdVO**](ApRssiThresholdVO.md) |  | [optional] 
**RssiSetting5g** | Pointer to [**ApRssiThresholdVO**](ApRssiThresholdVO.md) |  | [optional] 
**RssiSetting5g2** | Pointer to [**ApRssiThresholdVO**](ApRssiThresholdVO.md) |  | [optional] 
**RssiSetting6g** | Pointer to [**ApRssiThresholdVO**](ApRssiThresholdVO.md) |  | [optional] 
**RxRate** | Pointer to **int64** |  | [optional] 
**ServicesResource** | Pointer to **int32** |  | [optional] 
**ShowModel** | Pointer to **string** | Model complex shown in the front end.Ap：model+(country)+modelVersion,EAP225(EU) v3.0  Gateway/Switch：model+modelVersion,Osg v3.0 | [optional] 
**SiteName** | Pointer to **string** | Site name of the device | [optional] 
**SiteTemplateId** | Pointer to **string** | Template ID bound to the site | [optional] 
**SiteTemplateName** | Pointer to **string** | Template name bound to the site | [optional] 
**Sn** | Pointer to **string** | SN code of device | [optional] 
**SnmpSeting** | Pointer to [**OsgSnmpVO**](OsgSnmpVO.md) |  | [optional] 
**SpecialModel** | Pointer to **string** | Special device model,for example:EAP225-Outdoor-1a20a950b8d950e8 | [optional] 
**Speeds** | Pointer to **[]int32** |  | [optional] 
**SshSetting** | Pointer to [**SshSettingVO**](SshSettingVO.md) |  | [optional] 
**SsidOverrides** | Pointer to [**[]SsidOverrideVO**](SsidOverrideVO.md) |  | [optional] 
**Status** | Pointer to **int32** | Status of device,status should be a value as follows: 0:Disconnected;1:Disconnected(Migrating);10:Provisioning;11:Configuring;12:Upgrading;13:Rebooting;14:Connected;15:Connected(Wireless);16:Connected(Migrating);17:Connected(Wireless,Migrating);20:Pending;21:Pending(Wireless);22:Adopting;23:Adopting(Wireless);24:Adopt Failed;25:Adopt Failed(Wireless);26:Managed By Others;27:Managed By Others(Wireless);30:Heartbeat Missed;31:Heartbeat Missed(Wireless);32:Heartbeat Missed(Migrating);33:Heartbeat Missed(Wireless,Migrating);40:Isolated;41:Isolated(Migrating);50:Slice Configuring | [optional] 
**StatusCategory** | Pointer to **int32** | Category of device status,statusCategory should be a value as follows: 0:Disconnected;1:Connected;2:Pending;3:Heartbeat Missed;4:Isolated | [optional] 
**SupportAnomaly** | Pointer to **bool** | Whether the device firmware support intelligent anomaly detection | [optional] 
**SupportBandScan** | Pointer to **bool** |  | [optional] 
**SupportBandWidthCtrl** | Pointer to **bool** |  | [optional] 
**SupportBatchEdit** | Pointer to **bool** |  | [optional] 
**SupportChannelUtilizationStatus** | Pointer to **bool** |  | [optional] 
**SupportDroppedPacketsStatus** | Pointer to **bool** |  | [optional] 
**SupportDsl** | Pointer to **bool** |  | [optional] 
**SupportDualSim** | Pointer to **int32** |  | [optional] 
**SupportFlowControl** | Pointer to **bool** |  | [optional] 
**SupportHwOffload** | Pointer to **bool** |  | [optional] 
**SupportIpMacBinding** | Pointer to **bool** |  | [optional] 
**SupportIppt** | Pointer to **bool** |  | [optional] 
**SupportJumbo** | Pointer to **bool** |  | [optional] 
**SupportLanClientStats** | Pointer to **bool** |  | [optional] 
**SupportLocatePort** | Pointer to **bool** | Whether the device supports locating port | [optional] 
**SupportLoopbackControl** | Pointer to **bool** |  | [optional] 
**SupportLte** | Pointer to **bool** |  | [optional] 
**SupportMirror** | Pointer to **bool** |  | [optional] 
**SupportNetworkSearch** | Pointer to **bool** |  | [optional] 
**SupportNewIpMacBinding** | Pointer to **bool** |  | [optional] 
**SupportPoe** | Pointer to **bool** |  | [optional] 
**SupportPortControl** | Pointer to **bool** |  | [optional] 
**SupportPortForwardingStatus** | Pointer to **bool** |  | [optional] 
**SupportPortIsolation** | Pointer to **bool** |  | [optional] 
**SupportPvid** | Pointer to **bool** |  | [optional] 
**SupportRFScan** | Pointer to **bool** |  | [optional] 
**SupportRetriedPacketsStatus** | Pointer to **bool** |  | [optional] 
**SupportSessionLimitStatus** | Pointer to **bool** |  | [optional] 
**SupportSnmp** | Pointer to **bool** |  | [optional] 
**SupportSpeedDuplex** | Pointer to **bool** |  | [optional] 
**SupportStormCtrlAction** | Pointer to **bool** |  | [optional] 
**SupportVirtualWan** | Pointer to **bool** |  | [optional] 
**SupportWanSetPvid** | Pointer to **bool** |  | [optional] 
**TagIds** | Pointer to **[]string** |  | [optional] 
**Temp** | Pointer to **int32** |  | [optional] 
**TemplateId** | Pointer to **string** | ID of the template bound to the device | [optional] 
**TemplateName** | Pointer to **string** | Name of the template bound to the device | [optional] 
**TemplateSettings** | Pointer to **[]int32** |  | [optional] 
**TxRate** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** | Device type:ap、gateway、switch、olt | [optional] 
**UnsupportedPorts** | Pointer to **[]int32** |  | [optional] 
**Upload** | Pointer to **int64** |  | [optional] 
**Uptime** | Pointer to **string** |  | [optional] 
**UptimeLong** | Pointer to **int64** |  | [optional] 
**UserNum** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** | Simplified version of firmware,for example:2.5.0 | [optional] 
**VirtualWanStats** | Pointer to [**[]OsgVirtualWanStatVO**](OsgVirtualWanStatVO.md) |  | [optional] 
**WirelessAdvancedResource** | Pointer to **int32** |  | [optional] 
**WirelessHealth** | Pointer to **bool** |  | [optional] 
**WirelessRouter** | Pointer to **bool** |  | [optional] 
**WlanId** | Pointer to **string** |  | [optional] 
**WlansResource** | Pointer to **int32** |  | [optional] 
**Wp2g** | Pointer to [**ApRadioChannel**](ApRadioChannel.md) |  | [optional] 
**Wp5g** | Pointer to [**ApRadioChannel**](ApRadioChannel.md) |  | [optional] 
**Wp5g2** | Pointer to [**ApRadioChannel**](ApRadioChannel.md) |  | [optional] 
**Wp6g** | Pointer to [**ApRadioChannel**](ApRadioChannel.md) |  | [optional] 

## Methods

### NewOsgDetailVO

`func NewOsgDetailVO() *OsgDetailVO`

NewOsgDetailVO instantiates a new OsgDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgDetailVOWithDefaults

`func NewOsgDetailVOWithDefaults() *OsgDetailVO`

NewOsgDetailVOWithDefaults instantiates a new OsgDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *OsgDetailVO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *OsgDetailVO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *OsgDetailVO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *OsgDetailVO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetActive

`func (o *OsgDetailVO) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OsgDetailVO) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OsgDetailVO) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *OsgDetailVO) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddByTemplate

`func (o *OsgDetailVO) GetAddByTemplate() bool`

GetAddByTemplate returns the AddByTemplate field if non-nil, zero value otherwise.

### GetAddByTemplateOk

`func (o *OsgDetailVO) GetAddByTemplateOk() (*bool, bool)`

GetAddByTemplateOk returns a tuple with the AddByTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddByTemplate

`func (o *OsgDetailVO) SetAddByTemplate(v bool)`

SetAddByTemplate sets AddByTemplate field to given value.

### HasAddByTemplate

`func (o *OsgDetailVO) HasAddByTemplate() bool`

HasAddByTemplate returns a boolean if a field has been set.

### GetAddedInAdvanced

`func (o *OsgDetailVO) GetAddedInAdvanced() bool`

GetAddedInAdvanced returns the AddedInAdvanced field if non-nil, zero value otherwise.

### GetAddedInAdvancedOk

`func (o *OsgDetailVO) GetAddedInAdvancedOk() (*bool, bool)`

GetAddedInAdvancedOk returns a tuple with the AddedInAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedInAdvanced

`func (o *OsgDetailVO) SetAddedInAdvanced(v bool)`

SetAddedInAdvanced sets AddedInAdvanced field to given value.

### HasAddedInAdvanced

`func (o *OsgDetailVO) HasAddedInAdvanced() bool`

HasAddedInAdvanced returns a boolean if a field has been set.

### GetAdvancedResource

`func (o *OsgDetailVO) GetAdvancedResource() int32`

GetAdvancedResource returns the AdvancedResource field if non-nil, zero value otherwise.

### GetAdvancedResourceOk

`func (o *OsgDetailVO) GetAdvancedResourceOk() (*int32, bool)`

GetAdvancedResourceOk returns a tuple with the AdvancedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedResource

`func (o *OsgDetailVO) SetAdvancedResource(v int32)`

SetAdvancedResource sets AdvancedResource field to given value.

### HasAdvancedResource

`func (o *OsgDetailVO) HasAdvancedResource() bool`

HasAdvancedResource returns a boolean if a field has been set.

### GetBoundDeviceTemplate

`func (o *OsgDetailVO) GetBoundDeviceTemplate() bool`

GetBoundDeviceTemplate returns the BoundDeviceTemplate field if non-nil, zero value otherwise.

### GetBoundDeviceTemplateOk

`func (o *OsgDetailVO) GetBoundDeviceTemplateOk() (*bool, bool)`

GetBoundDeviceTemplateOk returns a tuple with the BoundDeviceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundDeviceTemplate

`func (o *OsgDetailVO) SetBoundDeviceTemplate(v bool)`

SetBoundDeviceTemplate sets BoundDeviceTemplate field to given value.

### HasBoundDeviceTemplate

`func (o *OsgDetailVO) HasBoundDeviceTemplate() bool`

HasBoundDeviceTemplate returns a boolean if a field has been set.

### GetBoundSiteTemplate

`func (o *OsgDetailVO) GetBoundSiteTemplate() bool`

GetBoundSiteTemplate returns the BoundSiteTemplate field if non-nil, zero value otherwise.

### GetBoundSiteTemplateOk

`func (o *OsgDetailVO) GetBoundSiteTemplateOk() (*bool, bool)`

GetBoundSiteTemplateOk returns a tuple with the BoundSiteTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSiteTemplate

`func (o *OsgDetailVO) SetBoundSiteTemplate(v bool)`

SetBoundSiteTemplate sets BoundSiteTemplate field to given value.

### HasBoundSiteTemplate

`func (o *OsgDetailVO) HasBoundSiteTemplate() bool`

HasBoundSiteTemplate returns a boolean if a field has been set.

### GetCategory

`func (o *OsgDetailVO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *OsgDetailVO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *OsgDetailVO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *OsgDetailVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetChannelLimitType

`func (o *OsgDetailVO) GetChannelLimitType() int32`

GetChannelLimitType returns the ChannelLimitType field if non-nil, zero value otherwise.

### GetChannelLimitTypeOk

`func (o *OsgDetailVO) GetChannelLimitTypeOk() (*int32, bool)`

GetChannelLimitTypeOk returns a tuple with the ChannelLimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLimitType

`func (o *OsgDetailVO) SetChannelLimitType(v int32)`

SetChannelLimitType sets ChannelLimitType field to given value.

### HasChannelLimitType

`func (o *OsgDetailVO) HasChannelLimitType() bool`

HasChannelLimitType returns a boolean if a field has been set.

### GetCombinedGateway

`func (o *OsgDetailVO) GetCombinedGateway() bool`

GetCombinedGateway returns the CombinedGateway field if non-nil, zero value otherwise.

### GetCombinedGatewayOk

`func (o *OsgDetailVO) GetCombinedGatewayOk() (*bool, bool)`

GetCombinedGatewayOk returns a tuple with the CombinedGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedGateway

`func (o *OsgDetailVO) SetCombinedGateway(v bool)`

SetCombinedGateway sets CombinedGateway field to given value.

### HasCombinedGateway

`func (o *OsgDetailVO) HasCombinedGateway() bool`

HasCombinedGateway returns a boolean if a field has been set.

### GetCommonAdvancedResource

`func (o *OsgDetailVO) GetCommonAdvancedResource() int32`

GetCommonAdvancedResource returns the CommonAdvancedResource field if non-nil, zero value otherwise.

### GetCommonAdvancedResourceOk

`func (o *OsgDetailVO) GetCommonAdvancedResourceOk() (*int32, bool)`

GetCommonAdvancedResourceOk returns a tuple with the CommonAdvancedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonAdvancedResource

`func (o *OsgDetailVO) SetCommonAdvancedResource(v int32)`

SetCommonAdvancedResource sets CommonAdvancedResource field to given value.

### HasCommonAdvancedResource

`func (o *OsgDetailVO) HasCommonAdvancedResource() bool`

HasCommonAdvancedResource returns a boolean if a field has been set.

### GetCompatible

`func (o *OsgDetailVO) GetCompatible() int32`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *OsgDetailVO) GetCompatibleOk() (*int32, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *OsgDetailVO) SetCompatible(v int32)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *OsgDetailVO) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetCompoundModel

`func (o *OsgDetailVO) GetCompoundModel() string`

GetCompoundModel returns the CompoundModel field if non-nil, zero value otherwise.

### GetCompoundModelOk

`func (o *OsgDetailVO) GetCompoundModelOk() (*string, bool)`

GetCompoundModelOk returns a tuple with the CompoundModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompoundModel

`func (o *OsgDetailVO) SetCompoundModel(v string)`

SetCompoundModel sets CompoundModel field to given value.

### HasCompoundModel

`func (o *OsgDetailVO) HasCompoundModel() bool`

HasCompoundModel returns a boolean if a field has been set.

### GetContainNetType5G

`func (o *OsgDetailVO) GetContainNetType5G() bool`

GetContainNetType5G returns the ContainNetType5G field if non-nil, zero value otherwise.

### GetContainNetType5GOk

`func (o *OsgDetailVO) GetContainNetType5GOk() (*bool, bool)`

GetContainNetType5GOk returns a tuple with the ContainNetType5G field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainNetType5G

`func (o *OsgDetailVO) SetContainNetType5G(v bool)`

SetContainNetType5G sets ContainNetType5G field to given value.

### HasContainNetType5G

`func (o *OsgDetailVO) HasContainNetType5G() bool`

HasContainNetType5G returns a boolean if a field has been set.

### GetControllerId

`func (o *OsgDetailVO) GetControllerId() string`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *OsgDetailVO) GetControllerIdOk() (*string, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *OsgDetailVO) SetControllerId(v string)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *OsgDetailVO) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### GetCpuUtil

`func (o *OsgDetailVO) GetCpuUtil() int32`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *OsgDetailVO) GetCpuUtilOk() (*int32, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *OsgDetailVO) SetCpuUtil(v int32)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *OsgDetailVO) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetCustomId

`func (o *OsgDetailVO) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *OsgDetailVO) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *OsgDetailVO) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.

### HasCustomId

`func (o *OsgDetailVO) HasCustomId() bool`

HasCustomId returns a boolean if a field has been set.

### GetCustomName

`func (o *OsgDetailVO) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *OsgDetailVO) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *OsgDetailVO) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *OsgDetailVO) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### GetDescription

`func (o *OsgDetailVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OsgDetailVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OsgDetailVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OsgDetailVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevCap

`func (o *OsgDetailVO) GetDevCap() OsgCapVO`

GetDevCap returns the DevCap field if non-nil, zero value otherwise.

### GetDevCapOk

`func (o *OsgDetailVO) GetDevCapOk() (*OsgCapVO, bool)`

GetDevCapOk returns a tuple with the DevCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevCap

`func (o *OsgDetailVO) SetDevCap(v OsgCapVO)`

SetDevCap sets DevCap field to given value.

### HasDevCap

`func (o *OsgDetailVO) HasDevCap() bool`

HasDevCap returns a boolean if a field has been set.

### GetDeviceMisc

`func (o *OsgDetailVO) GetDeviceMisc() WirelessRouterMiscVO`

GetDeviceMisc returns the DeviceMisc field if non-nil, zero value otherwise.

### GetDeviceMiscOk

`func (o *OsgDetailVO) GetDeviceMiscOk() (*WirelessRouterMiscVO, bool)`

GetDeviceMiscOk returns a tuple with the DeviceMisc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMisc

`func (o *OsgDetailVO) SetDeviceMisc(v WirelessRouterMiscVO)`

SetDeviceMisc sets DeviceMisc field to given value.

### HasDeviceMisc

`func (o *OsgDetailVO) HasDeviceMisc() bool`

HasDeviceMisc returns a boolean if a field has been set.

### GetDeviceSeriesType

`func (o *OsgDetailVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *OsgDetailVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *OsgDetailVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *OsgDetailVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetDeviceTemplateAvailable

`func (o *OsgDetailVO) GetDeviceTemplateAvailable() bool`

GetDeviceTemplateAvailable returns the DeviceTemplateAvailable field if non-nil, zero value otherwise.

### GetDeviceTemplateAvailableOk

`func (o *OsgDetailVO) GetDeviceTemplateAvailableOk() (*bool, bool)`

GetDeviceTemplateAvailableOk returns a tuple with the DeviceTemplateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTemplateAvailable

`func (o *OsgDetailVO) SetDeviceTemplateAvailable(v bool)`

SetDeviceTemplateAvailable sets DeviceTemplateAvailable field to given value.

### HasDeviceTemplateAvailable

`func (o *OsgDetailVO) HasDeviceTemplateAvailable() bool`

HasDeviceTemplateAvailable returns a boolean if a field has been set.

### GetDisableHwReset

`func (o *OsgDetailVO) GetDisableHwReset() bool`

GetDisableHwReset returns the DisableHwReset field if non-nil, zero value otherwise.

### GetDisableHwResetOk

`func (o *OsgDetailVO) GetDisableHwResetOk() (*bool, bool)`

GetDisableHwResetOk returns a tuple with the DisableHwReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHwReset

`func (o *OsgDetailVO) SetDisableHwReset(v bool)`

SetDisableHwReset sets DisableHwReset field to given value.

### HasDisableHwReset

`func (o *OsgDetailVO) HasDisableHwReset() bool`

HasDisableHwReset returns a boolean if a field has been set.

### GetDownlinkList

`func (o *OsgDetailVO) GetDownlinkList() []OsgDownLinkVO`

GetDownlinkList returns the DownlinkList field if non-nil, zero value otherwise.

### GetDownlinkListOk

`func (o *OsgDetailVO) GetDownlinkListOk() (*[]OsgDownLinkVO, bool)`

GetDownlinkListOk returns a tuple with the DownlinkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkList

`func (o *OsgDetailVO) SetDownlinkList(v []OsgDownLinkVO)`

SetDownlinkList sets DownlinkList field to given value.

### HasDownlinkList

`func (o *OsgDetailVO) HasDownlinkList() bool`

HasDownlinkList returns a boolean if a field has been set.

### GetDownload

`func (o *OsgDetailVO) GetDownload() int64`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *OsgDetailVO) GetDownloadOk() (*int64, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *OsgDetailVO) SetDownload(v int64)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *OsgDetailVO) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetDueTime

`func (o *OsgDetailVO) GetDueTime() int64`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *OsgDetailVO) GetDueTimeOk() (*int64, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *OsgDetailVO) SetDueTime(v int64)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *OsgDetailVO) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### GetDueTimeLeft

`func (o *OsgDetailVO) GetDueTimeLeft() int64`

GetDueTimeLeft returns the DueTimeLeft field if non-nil, zero value otherwise.

### GetDueTimeLeftOk

`func (o *OsgDetailVO) GetDueTimeLeftOk() (*int64, bool)`

GetDueTimeLeftOk returns a tuple with the DueTimeLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTimeLeft

`func (o *OsgDetailVO) SetDueTimeLeft(v int64)`

SetDueTimeLeft sets DueTimeLeft field to given value.

### HasDueTimeLeft

`func (o *OsgDetailVO) HasDueTimeLeft() bool`

HasDueTimeLeft returns a boolean if a field has been set.

### GetEchoServer

`func (o *OsgDetailVO) GetEchoServer() string`

GetEchoServer returns the EchoServer field if non-nil, zero value otherwise.

### GetEchoServerOk

`func (o *OsgDetailVO) GetEchoServerOk() (*string, bool)`

GetEchoServerOk returns a tuple with the EchoServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEchoServer

`func (o *OsgDetailVO) SetEchoServer(v string)`

SetEchoServer sets EchoServer field to given value.

### HasEchoServer

`func (o *OsgDetailVO) HasEchoServer() bool`

HasEchoServer returns a boolean if a field has been set.

### GetEcspFirstVersion

`func (o *OsgDetailVO) GetEcspFirstVersion() int32`

GetEcspFirstVersion returns the EcspFirstVersion field if non-nil, zero value otherwise.

### GetEcspFirstVersionOk

`func (o *OsgDetailVO) GetEcspFirstVersionOk() (*int32, bool)`

GetEcspFirstVersionOk returns a tuple with the EcspFirstVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcspFirstVersion

`func (o *OsgDetailVO) SetEcspFirstVersion(v int32)`

SetEcspFirstVersion sets EcspFirstVersion field to given value.

### HasEcspFirstVersion

`func (o *OsgDetailVO) HasEcspFirstVersion() bool`

HasEcspFirstVersion returns a boolean if a field has been set.

### GetEos

`func (o *OsgDetailVO) GetEos() int64`

GetEos returns the Eos field if non-nil, zero value otherwise.

### GetEosOk

`func (o *OsgDetailVO) GetEosOk() (*int64, bool)`

GetEosOk returns a tuple with the Eos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEos

`func (o *OsgDetailVO) SetEos(v int64)`

SetEos sets Eos field to given value.

### HasEos

`func (o *OsgDetailVO) HasEos() bool`

HasEos returns a boolean if a field has been set.

### GetEost

`func (o *OsgDetailVO) GetEost() int64`

GetEost returns the Eost field if non-nil, zero value otherwise.

### GetEostOk

`func (o *OsgDetailVO) GetEostOk() (*int64, bool)`

GetEostOk returns a tuple with the Eost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEost

`func (o *OsgDetailVO) SetEost(v int64)`

SetEost sets Eost field to given value.

### HasEost

`func (o *OsgDetailVO) HasEost() bool`

HasEost returns a boolean if a field has been set.

### GetEs

`func (o *OsgDetailVO) GetEs() bool`

GetEs returns the Es field if non-nil, zero value otherwise.

### GetEsOk

`func (o *OsgDetailVO) GetEsOk() (*bool, bool)`

GetEsOk returns a tuple with the Es field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEs

`func (o *OsgDetailVO) SetEs(v bool)`

SetEs sets Es field to given value.

### HasEs

`func (o *OsgDetailVO) HasEs() bool`

HasEs returns a boolean if a field has been set.

### GetFan

`func (o *OsgDetailVO) GetFan() []OsgFanStatusVO`

GetFan returns the Fan field if non-nil, zero value otherwise.

### GetFanOk

`func (o *OsgDetailVO) GetFanOk() (*[]OsgFanStatusVO, bool)`

GetFanOk returns a tuple with the Fan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFan

`func (o *OsgDetailVO) SetFan(v []OsgFanStatusVO)`

SetFan sets Fan field to given value.

### HasFan

`func (o *OsgDetailVO) HasFan() bool`

HasFan returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *OsgDetailVO) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *OsgDetailVO) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *OsgDetailVO) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *OsgDetailVO) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetForgetId

`func (o *OsgDetailVO) GetForgetId() string`

GetForgetId returns the ForgetId field if non-nil, zero value otherwise.

### GetForgetIdOk

`func (o *OsgDetailVO) GetForgetIdOk() (*string, bool)`

GetForgetIdOk returns a tuple with the ForgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgetId

`func (o *OsgDetailVO) SetForgetId(v string)`

SetForgetId sets ForgetId field to given value.

### HasForgetId

`func (o *OsgDetailVO) HasForgetId() bool`

HasForgetId returns a boolean if a field has been set.

### GetGuestNum

`func (o *OsgDetailVO) GetGuestNum() int32`

GetGuestNum returns the GuestNum field if non-nil, zero value otherwise.

### GetGuestNumOk

`func (o *OsgDetailVO) GetGuestNumOk() (*int32, bool)`

GetGuestNumOk returns a tuple with the GuestNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestNum

`func (o *OsgDetailVO) SetGuestNum(v int32)`

SetGuestNum sets GuestNum field to given value.

### HasGuestNum

`func (o *OsgDetailVO) HasGuestNum() bool`

HasGuestNum returns a boolean if a field has been set.

### GetHwOffloadEnable

`func (o *OsgDetailVO) GetHwOffloadEnable() bool`

GetHwOffloadEnable returns the HwOffloadEnable field if non-nil, zero value otherwise.

### GetHwOffloadEnableOk

`func (o *OsgDetailVO) GetHwOffloadEnableOk() (*bool, bool)`

GetHwOffloadEnableOk returns a tuple with the HwOffloadEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwOffloadEnable

`func (o *OsgDetailVO) SetHwOffloadEnable(v bool)`

SetHwOffloadEnable sets HwOffloadEnable field to given value.

### HasHwOffloadEnable

`func (o *OsgDetailVO) HasHwOffloadEnable() bool`

HasHwOffloadEnable returns a boolean if a field has been set.

### GetHwVersion

`func (o *OsgDetailVO) GetHwVersion() string`

GetHwVersion returns the HwVersion field if non-nil, zero value otherwise.

### GetHwVersionOk

`func (o *OsgDetailVO) GetHwVersionOk() (*string, bool)`

GetHwVersionOk returns a tuple with the HwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwVersion

`func (o *OsgDetailVO) SetHwVersion(v string)`

SetHwVersion sets HwVersion field to given value.

### HasHwVersion

`func (o *OsgDetailVO) HasHwVersion() bool`

HasHwVersion returns a boolean if a field has been set.

### GetInWhitelist

`func (o *OsgDetailVO) GetInWhitelist() bool`

GetInWhitelist returns the InWhitelist field if non-nil, zero value otherwise.

### GetInWhitelistOk

`func (o *OsgDetailVO) GetInWhitelistOk() (*bool, bool)`

GetInWhitelistOk returns a tuple with the InWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInWhitelist

`func (o *OsgDetailVO) SetInWhitelist(v bool)`

SetInWhitelist sets InWhitelist field to given value.

### HasInWhitelist

`func (o *OsgDetailVO) HasInWhitelist() bool`

HasInWhitelist returns a boolean if a field has been set.

### GetInitialUnbindingLimit

`func (o *OsgDetailVO) GetInitialUnbindingLimit() int32`

GetInitialUnbindingLimit returns the InitialUnbindingLimit field if non-nil, zero value otherwise.

### GetInitialUnbindingLimitOk

`func (o *OsgDetailVO) GetInitialUnbindingLimitOk() (*int32, bool)`

GetInitialUnbindingLimitOk returns a tuple with the InitialUnbindingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialUnbindingLimit

`func (o *OsgDetailVO) SetInitialUnbindingLimit(v int32)`

SetInitialUnbindingLimit sets InitialUnbindingLimit field to given value.

### HasInitialUnbindingLimit

`func (o *OsgDetailVO) HasInitialUnbindingLimit() bool`

HasInitialUnbindingLimit returns a boolean if a field has been set.

### GetIp

`func (o *OsgDetailVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OsgDetailVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OsgDetailVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OsgDetailVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIppt

`func (o *OsgDetailVO) GetIppt() bool`

GetIppt returns the Ippt field if non-nil, zero value otherwise.

### GetIpptOk

`func (o *OsgDetailVO) GetIpptOk() (*bool, bool)`

GetIpptOk returns a tuple with the Ippt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIppt

`func (o *OsgDetailVO) SetIppt(v bool)`

SetIppt sets Ippt field to given value.

### HasIppt

`func (o *OsgDetailVO) HasIppt() bool`

HasIppt returns a boolean if a field has been set.

### GetIpptPreconfig

`func (o *OsgDetailVO) GetIpptPreconfig() bool`

GetIpptPreconfig returns the IpptPreconfig field if non-nil, zero value otherwise.

### GetIpptPreconfigOk

`func (o *OsgDetailVO) GetIpptPreconfigOk() (*bool, bool)`

GetIpptPreconfigOk returns a tuple with the IpptPreconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpptPreconfig

`func (o *OsgDetailVO) SetIpptPreconfig(v bool)`

SetIpptPreconfig sets IpptPreconfig field to given value.

### HasIpptPreconfig

`func (o *OsgDetailVO) HasIpptPreconfig() bool`

HasIpptPreconfig returns a boolean if a field has been set.

### GetIptvSetting

`func (o *OsgDetailVO) GetIptvSetting() OsgIptvVO`

GetIptvSetting returns the IptvSetting field if non-nil, zero value otherwise.

### GetIptvSettingOk

`func (o *OsgDetailVO) GetIptvSettingOk() (*OsgIptvVO, bool)`

GetIptvSettingOk returns a tuple with the IptvSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIptvSetting

`func (o *OsgDetailVO) SetIptvSetting(v OsgIptvVO)`

SetIptvSetting sets IptvSetting field to given value.

### HasIptvSetting

`func (o *OsgDetailVO) HasIptvSetting() bool`

HasIptvSetting returns a boolean if a field has been set.

### GetIpv6List

`func (o *OsgDetailVO) GetIpv6List() []string`

GetIpv6List returns the Ipv6List field if non-nil, zero value otherwise.

### GetIpv6ListOk

`func (o *OsgDetailVO) GetIpv6ListOk() (*[]string, bool)`

GetIpv6ListOk returns a tuple with the Ipv6List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6List

`func (o *OsgDetailVO) SetIpv6List(v []string)`

SetIpv6List sets Ipv6List field to given value.

### HasIpv6List

`func (o *OsgDetailVO) HasIpv6List() bool`

HasIpv6List returns a boolean if a field has been set.

### GetJumboOptions

`func (o *OsgDetailVO) GetJumboOptions() []int32`

GetJumboOptions returns the JumboOptions field if non-nil, zero value otherwise.

### GetJumboOptionsOk

`func (o *OsgDetailVO) GetJumboOptionsOk() (*[]int32, bool)`

GetJumboOptionsOk returns a tuple with the JumboOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumboOptions

`func (o *OsgDetailVO) SetJumboOptions(v []int32)`

SetJumboOptions sets JumboOptions field to given value.

### HasJumboOptions

`func (o *OsgDetailVO) HasJumboOptions() bool`

HasJumboOptions returns a boolean if a field has been set.

### GetJumboSize

`func (o *OsgDetailVO) GetJumboSize() int32`

GetJumboSize returns the JumboSize field if non-nil, zero value otherwise.

### GetJumboSizeOk

`func (o *OsgDetailVO) GetJumboSizeOk() (*int32, bool)`

GetJumboSizeOk returns a tuple with the JumboSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumboSize

`func (o *OsgDetailVO) SetJumboSize(v int32)`

SetJumboSize sets JumboSize field to given value.

### HasJumboSize

`func (o *OsgDetailVO) HasJumboSize() bool`

HasJumboSize returns a boolean if a field has been set.

### GetLanClientStats

`func (o *OsgDetailVO) GetLanClientStats() []OsgLanStatVO`

GetLanClientStats returns the LanClientStats field if non-nil, zero value otherwise.

### GetLanClientStatsOk

`func (o *OsgDetailVO) GetLanClientStatsOk() (*[]OsgLanStatVO, bool)`

GetLanClientStatsOk returns a tuple with the LanClientStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanClientStats

`func (o *OsgDetailVO) SetLanClientStats(v []OsgLanStatVO)`

SetLanClientStats sets LanClientStats field to given value.

### HasLanClientStats

`func (o *OsgDetailVO) HasLanClientStats() bool`

HasLanClientStats returns a boolean if a field has been set.

### GetLastSeen

`func (o *OsgDetailVO) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *OsgDetailVO) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *OsgDetailVO) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *OsgDetailVO) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLbSetting2g

`func (o *OsgDetailVO) GetLbSetting2g() ApLoadBalanceVO`

GetLbSetting2g returns the LbSetting2g field if non-nil, zero value otherwise.

### GetLbSetting2gOk

`func (o *OsgDetailVO) GetLbSetting2gOk() (*ApLoadBalanceVO, bool)`

GetLbSetting2gOk returns a tuple with the LbSetting2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbSetting2g

`func (o *OsgDetailVO) SetLbSetting2g(v ApLoadBalanceVO)`

SetLbSetting2g sets LbSetting2g field to given value.

### HasLbSetting2g

`func (o *OsgDetailVO) HasLbSetting2g() bool`

HasLbSetting2g returns a boolean if a field has been set.

### GetLbSetting5g

`func (o *OsgDetailVO) GetLbSetting5g() ApLoadBalanceVO`

GetLbSetting5g returns the LbSetting5g field if non-nil, zero value otherwise.

### GetLbSetting5gOk

`func (o *OsgDetailVO) GetLbSetting5gOk() (*ApLoadBalanceVO, bool)`

GetLbSetting5gOk returns a tuple with the LbSetting5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbSetting5g

`func (o *OsgDetailVO) SetLbSetting5g(v ApLoadBalanceVO)`

SetLbSetting5g sets LbSetting5g field to given value.

### HasLbSetting5g

`func (o *OsgDetailVO) HasLbSetting5g() bool`

HasLbSetting5g returns a boolean if a field has been set.

### GetLbSetting5g2

`func (o *OsgDetailVO) GetLbSetting5g2() ApLoadBalanceVO`

GetLbSetting5g2 returns the LbSetting5g2 field if non-nil, zero value otherwise.

### GetLbSetting5g2Ok

`func (o *OsgDetailVO) GetLbSetting5g2Ok() (*ApLoadBalanceVO, bool)`

GetLbSetting5g2Ok returns a tuple with the LbSetting5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbSetting5g2

`func (o *OsgDetailVO) SetLbSetting5g2(v ApLoadBalanceVO)`

SetLbSetting5g2 sets LbSetting5g2 field to given value.

### HasLbSetting5g2

`func (o *OsgDetailVO) HasLbSetting5g2() bool`

HasLbSetting5g2 returns a boolean if a field has been set.

### GetLbSetting6g

`func (o *OsgDetailVO) GetLbSetting6g() ApLoadBalanceVO`

GetLbSetting6g returns the LbSetting6g field if non-nil, zero value otherwise.

### GetLbSetting6gOk

`func (o *OsgDetailVO) GetLbSetting6gOk() (*ApLoadBalanceVO, bool)`

GetLbSetting6gOk returns a tuple with the LbSetting6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbSetting6g

`func (o *OsgDetailVO) SetLbSetting6g(v ApLoadBalanceVO)`

SetLbSetting6g sets LbSetting6g field to given value.

### HasLbSetting6g

`func (o *OsgDetailVO) HasLbSetting6g() bool`

HasLbSetting6g returns a boolean if a field has been set.

### GetLedSetting

`func (o *OsgDetailVO) GetLedSetting() int32`

GetLedSetting returns the LedSetting field if non-nil, zero value otherwise.

### GetLedSettingOk

`func (o *OsgDetailVO) GetLedSettingOk() (*int32, bool)`

GetLedSettingOk returns a tuple with the LedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedSetting

`func (o *OsgDetailVO) SetLedSetting(v int32)`

SetLedSetting sets LedSetting field to given value.

### HasLedSetting

`func (o *OsgDetailVO) HasLedSetting() bool`

HasLedSetting returns a boolean if a field has been set.

### GetLicenseId

`func (o *OsgDetailVO) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *OsgDetailVO) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *OsgDetailVO) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *OsgDetailVO) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetLicenseStatus

`func (o *OsgDetailVO) GetLicenseStatus() int32`

GetLicenseStatus returns the LicenseStatus field if non-nil, zero value otherwise.

### GetLicenseStatusOk

`func (o *OsgDetailVO) GetLicenseStatusOk() (*int32, bool)`

GetLicenseStatusOk returns a tuple with the LicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatus

`func (o *OsgDetailVO) SetLicenseStatus(v int32)`

SetLicenseStatus sets LicenseStatus field to given value.

### HasLicenseStatus

`func (o *OsgDetailVO) HasLicenseStatus() bool`

HasLicenseStatus returns a boolean if a field has been set.

### GetLicenseUnbindingLimit

`func (o *OsgDetailVO) GetLicenseUnbindingLimit() int32`

GetLicenseUnbindingLimit returns the LicenseUnbindingLimit field if non-nil, zero value otherwise.

### GetLicenseUnbindingLimitOk

`func (o *OsgDetailVO) GetLicenseUnbindingLimitOk() (*int32, bool)`

GetLicenseUnbindingLimitOk returns a tuple with the LicenseUnbindingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUnbindingLimit

`func (o *OsgDetailVO) SetLicenseUnbindingLimit(v int32)`

SetLicenseUnbindingLimit sets LicenseUnbindingLimit field to given value.

### HasLicenseUnbindingLimit

`func (o *OsgDetailVO) HasLicenseUnbindingLimit() bool`

HasLicenseUnbindingLimit returns a boolean if a field has been set.

### GetLldpEnable

`func (o *OsgDetailVO) GetLldpEnable() bool`

GetLldpEnable returns the LldpEnable field if non-nil, zero value otherwise.

### GetLldpEnableOk

`func (o *OsgDetailVO) GetLldpEnableOk() (*bool, bool)`

GetLldpEnableOk returns a tuple with the LldpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpEnable

`func (o *OsgDetailVO) SetLldpEnable(v bool)`

SetLldpEnable sets LldpEnable field to given value.

### HasLldpEnable

`func (o *OsgDetailVO) HasLldpEnable() bool`

HasLldpEnable returns a boolean if a field has been set.

### GetLldpSetting

`func (o *OsgDetailVO) GetLldpSetting() int32`

GetLldpSetting returns the LldpSetting field if non-nil, zero value otherwise.

### GetLldpSettingOk

`func (o *OsgDetailVO) GetLldpSettingOk() (*int32, bool)`

GetLldpSettingOk returns a tuple with the LldpSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpSetting

`func (o *OsgDetailVO) SetLldpSetting(v int32)`

SetLldpSetting sets LldpSetting field to given value.

### HasLldpSetting

`func (o *OsgDetailVO) HasLldpSetting() bool`

HasLldpSetting returns a boolean if a field has been set.

### GetLocation

`func (o *OsgDetailVO) GetLocation() DeviceLocationDetailVO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OsgDetailVO) GetLocationOk() (*DeviceLocationDetailVO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OsgDetailVO) SetLocation(v DeviceLocationDetailVO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OsgDetailVO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMac

`func (o *OsgDetailVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OsgDetailVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OsgDetailVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OsgDetailVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMemUtil

`func (o *OsgDetailVO) GetMemUtil() int32`

GetMemUtil returns the MemUtil field if non-nil, zero value otherwise.

### GetMemUtilOk

`func (o *OsgDetailVO) GetMemUtilOk() (*int32, bool)`

GetMemUtilOk returns a tuple with the MemUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUtil

`func (o *OsgDetailVO) SetMemUtil(v int32)`

SetMemUtil sets MemUtil field to given value.

### HasMemUtil

`func (o *OsgDetailVO) HasMemUtil() bool`

HasMemUtil returns a boolean if a field has been set.

### GetModel

`func (o *OsgDetailVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OsgDetailVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OsgDetailVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OsgDetailVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OsgDetailVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OsgDetailVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OsgDetailVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OsgDetailVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetMoveSiteId

`func (o *OsgDetailVO) GetMoveSiteId() string`

GetMoveSiteId returns the MoveSiteId field if non-nil, zero value otherwise.

### GetMoveSiteIdOk

`func (o *OsgDetailVO) GetMoveSiteIdOk() (*string, bool)`

GetMoveSiteIdOk returns a tuple with the MoveSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveSiteId

`func (o *OsgDetailVO) SetMoveSiteId(v string)`

SetMoveSiteId sets MoveSiteId field to given value.

### HasMoveSiteId

`func (o *OsgDetailVO) HasMoveSiteId() bool`

HasMoveSiteId returns a boolean if a field has been set.

### GetMultiChipGateway

`func (o *OsgDetailVO) GetMultiChipGateway() bool`

GetMultiChipGateway returns the MultiChipGateway field if non-nil, zero value otherwise.

### GetMultiChipGatewayOk

`func (o *OsgDetailVO) GetMultiChipGatewayOk() (*bool, bool)`

GetMultiChipGatewayOk returns a tuple with the MultiChipGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiChipGateway

`func (o *OsgDetailVO) SetMultiChipGateway(v bool)`

SetMultiChipGateway sets MultiChipGateway field to given value.

### HasMultiChipGateway

`func (o *OsgDetailVO) HasMultiChipGateway() bool`

HasMultiChipGateway returns a boolean if a field has been set.

### GetMultiChipInfos

`func (o *OsgDetailVO) GetMultiChipInfos() [][]int32`

GetMultiChipInfos returns the MultiChipInfos field if non-nil, zero value otherwise.

### GetMultiChipInfosOk

`func (o *OsgDetailVO) GetMultiChipInfosOk() (*[][]int32, bool)`

GetMultiChipInfosOk returns a tuple with the MultiChipInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiChipInfos

`func (o *OsgDetailVO) SetMultiChipInfos(v [][]int32)`

SetMultiChipInfos sets MultiChipInfos field to given value.

### HasMultiChipInfos

`func (o *OsgDetailVO) HasMultiChipInfos() bool`

HasMultiChipInfos returns a boolean if a field has been set.

### GetName

`func (o *OsgDetailVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsgDetailVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsgDetailVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsgDetailVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedUpgrade

`func (o *OsgDetailVO) GetNeedUpgrade() bool`

GetNeedUpgrade returns the NeedUpgrade field if non-nil, zero value otherwise.

### GetNeedUpgradeOk

`func (o *OsgDetailVO) GetNeedUpgradeOk() (*bool, bool)`

GetNeedUpgradeOk returns a tuple with the NeedUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedUpgrade

`func (o *OsgDetailVO) SetNeedUpgrade(v bool)`

SetNeedUpgrade sets NeedUpgrade field to given value.

### HasNeedUpgrade

`func (o *OsgDetailVO) HasNeedUpgrade() bool`

HasNeedUpgrade returns a boolean if a field has been set.

### GetNetworkComptent

`func (o *OsgDetailVO) GetNetworkComptent() int32`

GetNetworkComptent returns the NetworkComptent field if non-nil, zero value otherwise.

### GetNetworkComptentOk

`func (o *OsgDetailVO) GetNetworkComptentOk() (*int32, bool)`

GetNetworkComptentOk returns a tuple with the NetworkComptent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComptent

`func (o *OsgDetailVO) SetNetworkComptent(v int32)`

SetNetworkComptent sets NetworkComptent field to given value.

### HasNetworkComptent

`func (o *OsgDetailVO) HasNetworkComptent() bool`

HasNetworkComptent returns a boolean if a field has been set.

### GetNonPscEnable

`func (o *OsgDetailVO) GetNonPscEnable() bool`

GetNonPscEnable returns the NonPscEnable field if non-nil, zero value otherwise.

### GetNonPscEnableOk

`func (o *OsgDetailVO) GetNonPscEnableOk() (*bool, bool)`

GetNonPscEnableOk returns a tuple with the NonPscEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonPscEnable

`func (o *OsgDetailVO) SetNonPscEnable(v bool)`

SetNonPscEnable sets NonPscEnable field to given value.

### HasNonPscEnable

`func (o *OsgDetailVO) HasNonPscEnable() bool`

HasNonPscEnable returns a boolean if a field has been set.

### GetOfdmaEnable2g

`func (o *OsgDetailVO) GetOfdmaEnable2g() bool`

GetOfdmaEnable2g returns the OfdmaEnable2g field if non-nil, zero value otherwise.

### GetOfdmaEnable2gOk

`func (o *OsgDetailVO) GetOfdmaEnable2gOk() (*bool, bool)`

GetOfdmaEnable2gOk returns a tuple with the OfdmaEnable2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfdmaEnable2g

`func (o *OsgDetailVO) SetOfdmaEnable2g(v bool)`

SetOfdmaEnable2g sets OfdmaEnable2g field to given value.

### HasOfdmaEnable2g

`func (o *OsgDetailVO) HasOfdmaEnable2g() bool`

HasOfdmaEnable2g returns a boolean if a field has been set.

### GetOfdmaEnable5g

`func (o *OsgDetailVO) GetOfdmaEnable5g() bool`

GetOfdmaEnable5g returns the OfdmaEnable5g field if non-nil, zero value otherwise.

### GetOfdmaEnable5gOk

`func (o *OsgDetailVO) GetOfdmaEnable5gOk() (*bool, bool)`

GetOfdmaEnable5gOk returns a tuple with the OfdmaEnable5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfdmaEnable5g

`func (o *OsgDetailVO) SetOfdmaEnable5g(v bool)`

SetOfdmaEnable5g sets OfdmaEnable5g field to given value.

### HasOfdmaEnable5g

`func (o *OsgDetailVO) HasOfdmaEnable5g() bool`

HasOfdmaEnable5g returns a boolean if a field has been set.

### GetOfdmaEnable5g2

`func (o *OsgDetailVO) GetOfdmaEnable5g2() bool`

GetOfdmaEnable5g2 returns the OfdmaEnable5g2 field if non-nil, zero value otherwise.

### GetOfdmaEnable5g2Ok

`func (o *OsgDetailVO) GetOfdmaEnable5g2Ok() (*bool, bool)`

GetOfdmaEnable5g2Ok returns a tuple with the OfdmaEnable5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfdmaEnable5g2

`func (o *OsgDetailVO) SetOfdmaEnable5g2(v bool)`

SetOfdmaEnable5g2 sets OfdmaEnable5g2 field to given value.

### HasOfdmaEnable5g2

`func (o *OsgDetailVO) HasOfdmaEnable5g2() bool`

HasOfdmaEnable5g2 returns a boolean if a field has been set.

### GetOfdmaEnable6g

`func (o *OsgDetailVO) GetOfdmaEnable6g() bool`

GetOfdmaEnable6g returns the OfdmaEnable6g field if non-nil, zero value otherwise.

### GetOfdmaEnable6gOk

`func (o *OsgDetailVO) GetOfdmaEnable6gOk() (*bool, bool)`

GetOfdmaEnable6gOk returns a tuple with the OfdmaEnable6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfdmaEnable6g

`func (o *OsgDetailVO) SetOfdmaEnable6g(v bool)`

SetOfdmaEnable6g sets OfdmaEnable6g field to given value.

### HasOfdmaEnable6g

`func (o *OsgDetailVO) HasOfdmaEnable6g() bool`

HasOfdmaEnable6g returns a boolean if a field has been set.

### GetOmadacId

`func (o *OsgDetailVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *OsgDetailVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *OsgDetailVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *OsgDetailVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetOsgCap

`func (o *OsgDetailVO) GetOsgCap() SiteSettingCapVO`

GetOsgCap returns the OsgCap field if non-nil, zero value otherwise.

### GetOsgCapOk

`func (o *OsgDetailVO) GetOsgCapOk() (*SiteSettingCapVO, bool)`

GetOsgCapOk returns a tuple with the OsgCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsgCap

`func (o *OsgDetailVO) SetOsgCap(v SiteSettingCapVO)`

SetOsgCap sets OsgCap field to given value.

### HasOsgCap

`func (o *OsgDetailVO) HasOsgCap() bool`

HasOsgCap returns a boolean if a field has been set.

### GetOsgExist

`func (o *OsgDetailVO) GetOsgExist() ExistSiteSettingVO`

GetOsgExist returns the OsgExist field if non-nil, zero value otherwise.

### GetOsgExistOk

`func (o *OsgDetailVO) GetOsgExistOk() (*ExistSiteSettingVO, bool)`

GetOsgExistOk returns a tuple with the OsgExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsgExist

`func (o *OsgDetailVO) SetOsgExist(v ExistSiteSettingVO)`

SetOsgExist sets OsgExist field to given value.

### HasOsgExist

`func (o *OsgDetailVO) HasOsgExist() bool`

HasOsgExist returns a boolean if a field has been set.

### GetPinSetting

`func (o *OsgDetailVO) GetPinSetting() OsgLtePinSettingVO`

GetPinSetting returns the PinSetting field if non-nil, zero value otherwise.

### GetPinSettingOk

`func (o *OsgDetailVO) GetPinSettingOk() (*OsgLtePinSettingVO, bool)`

GetPinSettingOk returns a tuple with the PinSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinSetting

`func (o *OsgDetailVO) SetPinSetting(v OsgLtePinSettingVO)`

SetPinSetting sets PinSetting field to given value.

### HasPinSetting

`func (o *OsgDetailVO) HasPinSetting() bool`

HasPinSetting returns a boolean if a field has been set.

### GetPoeLimit

`func (o *OsgDetailVO) GetPoeLimit() float64`

GetPoeLimit returns the PoeLimit field if non-nil, zero value otherwise.

### GetPoeLimitOk

`func (o *OsgDetailVO) GetPoeLimitOk() (*float64, bool)`

GetPoeLimitOk returns a tuple with the PoeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeLimit

`func (o *OsgDetailVO) SetPoeLimit(v float64)`

SetPoeLimit sets PoeLimit field to given value.

### HasPoeLimit

`func (o *OsgDetailVO) HasPoeLimit() bool`

HasPoeLimit returns a boolean if a field has been set.

### GetPoeRemain

`func (o *OsgDetailVO) GetPoeRemain() float64`

GetPoeRemain returns the PoeRemain field if non-nil, zero value otherwise.

### GetPoeRemainOk

`func (o *OsgDetailVO) GetPoeRemainOk() (*float64, bool)`

GetPoeRemainOk returns a tuple with the PoeRemain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeRemain

`func (o *OsgDetailVO) SetPoeRemain(v float64)`

SetPoeRemain sets PoeRemain field to given value.

### HasPoeRemain

`func (o *OsgDetailVO) HasPoeRemain() bool`

HasPoeRemain returns a boolean if a field has been set.

### GetPoeRemainPercent

`func (o *OsgDetailVO) GetPoeRemainPercent() float64`

GetPoeRemainPercent returns the PoeRemainPercent field if non-nil, zero value otherwise.

### GetPoeRemainPercentOk

`func (o *OsgDetailVO) GetPoeRemainPercentOk() (*float64, bool)`

GetPoeRemainPercentOk returns a tuple with the PoeRemainPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeRemainPercent

`func (o *OsgDetailVO) SetPoeRemainPercent(v float64)`

SetPoeRemainPercent sets PoeRemainPercent field to given value.

### HasPoeRemainPercent

`func (o *OsgDetailVO) HasPoeRemainPercent() bool`

HasPoeRemainPercent returns a boolean if a field has been set.

### GetPoeSettings

`func (o *OsgDetailVO) GetPoeSettings() []OsgPortPoeVO`

GetPoeSettings returns the PoeSettings field if non-nil, zero value otherwise.

### GetPoeSettingsOk

`func (o *OsgDetailVO) GetPoeSettingsOk() (*[]OsgPortPoeVO, bool)`

GetPoeSettingsOk returns a tuple with the PoeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeSettings

`func (o *OsgDetailVO) SetPoeSettings(v []OsgPortPoeVO)`

SetPoeSettings sets PoeSettings field to given value.

### HasPoeSettings

`func (o *OsgDetailVO) HasPoeSettings() bool`

HasPoeSettings returns a boolean if a field has been set.

### GetPortConfigs

`func (o *OsgDetailVO) GetPortConfigs() []OsgPortConfigVO`

GetPortConfigs returns the PortConfigs field if non-nil, zero value otherwise.

### GetPortConfigsOk

`func (o *OsgDetailVO) GetPortConfigsOk() (*[]OsgPortConfigVO, bool)`

GetPortConfigsOk returns a tuple with the PortConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfigs

`func (o *OsgDetailVO) SetPortConfigs(v []OsgPortConfigVO)`

SetPortConfigs sets PortConfigs field to given value.

### HasPortConfigs

`func (o *OsgDetailVO) HasPortConfigs() bool`

HasPortConfigs returns a boolean if a field has been set.

### GetPortNum

`func (o *OsgDetailVO) GetPortNum() int32`

GetPortNum returns the PortNum field if non-nil, zero value otherwise.

### GetPortNumOk

`func (o *OsgDetailVO) GetPortNumOk() (*int32, bool)`

GetPortNumOk returns a tuple with the PortNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNum

`func (o *OsgDetailVO) SetPortNum(v int32)`

SetPortNum sets PortNum field to given value.

### HasPortNum

`func (o *OsgDetailVO) HasPortNum() bool`

HasPortNum returns a boolean if a field has been set.

### GetPortStats

`func (o *OsgDetailVO) GetPortStats() []OsgPortStatVO`

GetPortStats returns the PortStats field if non-nil, zero value otherwise.

### GetPortStatsOk

`func (o *OsgDetailVO) GetPortStatsOk() (*[]OsgPortStatVO, bool)`

GetPortStatsOk returns a tuple with the PortStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStats

`func (o *OsgDetailVO) SetPortStats(v []OsgPortStatVO)`

SetPortStats sets PortStats field to given value.

### HasPortStats

`func (o *OsgDetailVO) HasPortStats() bool`

HasPortStats returns a boolean if a field has been set.

### GetPublicIp

`func (o *OsgDetailVO) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *OsgDetailVO) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *OsgDetailVO) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *OsgDetailVO) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetQosSetting2g

`func (o *OsgDetailVO) GetQosSetting2g() ApQosVO`

GetQosSetting2g returns the QosSetting2g field if non-nil, zero value otherwise.

### GetQosSetting2gOk

`func (o *OsgDetailVO) GetQosSetting2gOk() (*ApQosVO, bool)`

GetQosSetting2gOk returns a tuple with the QosSetting2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosSetting2g

`func (o *OsgDetailVO) SetQosSetting2g(v ApQosVO)`

SetQosSetting2g sets QosSetting2g field to given value.

### HasQosSetting2g

`func (o *OsgDetailVO) HasQosSetting2g() bool`

HasQosSetting2g returns a boolean if a field has been set.

### GetQosSetting5g

`func (o *OsgDetailVO) GetQosSetting5g() ApQosVO`

GetQosSetting5g returns the QosSetting5g field if non-nil, zero value otherwise.

### GetQosSetting5gOk

`func (o *OsgDetailVO) GetQosSetting5gOk() (*ApQosVO, bool)`

GetQosSetting5gOk returns a tuple with the QosSetting5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosSetting5g

`func (o *OsgDetailVO) SetQosSetting5g(v ApQosVO)`

SetQosSetting5g sets QosSetting5g field to given value.

### HasQosSetting5g

`func (o *OsgDetailVO) HasQosSetting5g() bool`

HasQosSetting5g returns a boolean if a field has been set.

### GetQosSetting5g2

`func (o *OsgDetailVO) GetQosSetting5g2() ApQosVO`

GetQosSetting5g2 returns the QosSetting5g2 field if non-nil, zero value otherwise.

### GetQosSetting5g2Ok

`func (o *OsgDetailVO) GetQosSetting5g2Ok() (*ApQosVO, bool)`

GetQosSetting5g2Ok returns a tuple with the QosSetting5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosSetting5g2

`func (o *OsgDetailVO) SetQosSetting5g2(v ApQosVO)`

SetQosSetting5g2 sets QosSetting5g2 field to given value.

### HasQosSetting5g2

`func (o *OsgDetailVO) HasQosSetting5g2() bool`

HasQosSetting5g2 returns a boolean if a field has been set.

### GetQosSetting6g

`func (o *OsgDetailVO) GetQosSetting6g() ApQosVO`

GetQosSetting6g returns the QosSetting6g field if non-nil, zero value otherwise.

### GetQosSetting6gOk

`func (o *OsgDetailVO) GetQosSetting6gOk() (*ApQosVO, bool)`

GetQosSetting6gOk returns a tuple with the QosSetting6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosSetting6g

`func (o *OsgDetailVO) SetQosSetting6g(v ApQosVO)`

SetQosSetting6g sets QosSetting6g field to given value.

### HasQosSetting6g

`func (o *OsgDetailVO) HasQosSetting6g() bool`

HasQosSetting6g returns a boolean if a field has been set.

### GetRadioSetting2g

`func (o *OsgDetailVO) GetRadioSetting2g() ApRadioSetting`

GetRadioSetting2g returns the RadioSetting2g field if non-nil, zero value otherwise.

### GetRadioSetting2gOk

`func (o *OsgDetailVO) GetRadioSetting2gOk() (*ApRadioSetting, bool)`

GetRadioSetting2gOk returns a tuple with the RadioSetting2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioSetting2g

`func (o *OsgDetailVO) SetRadioSetting2g(v ApRadioSetting)`

SetRadioSetting2g sets RadioSetting2g field to given value.

### HasRadioSetting2g

`func (o *OsgDetailVO) HasRadioSetting2g() bool`

HasRadioSetting2g returns a boolean if a field has been set.

### GetRadioSetting5g

`func (o *OsgDetailVO) GetRadioSetting5g() ApRadioSetting`

GetRadioSetting5g returns the RadioSetting5g field if non-nil, zero value otherwise.

### GetRadioSetting5gOk

`func (o *OsgDetailVO) GetRadioSetting5gOk() (*ApRadioSetting, bool)`

GetRadioSetting5gOk returns a tuple with the RadioSetting5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioSetting5g

`func (o *OsgDetailVO) SetRadioSetting5g(v ApRadioSetting)`

SetRadioSetting5g sets RadioSetting5g field to given value.

### HasRadioSetting5g

`func (o *OsgDetailVO) HasRadioSetting5g() bool`

HasRadioSetting5g returns a boolean if a field has been set.

### GetRadioSetting5g2

`func (o *OsgDetailVO) GetRadioSetting5g2() ApRadioSetting`

GetRadioSetting5g2 returns the RadioSetting5g2 field if non-nil, zero value otherwise.

### GetRadioSetting5g2Ok

`func (o *OsgDetailVO) GetRadioSetting5g2Ok() (*ApRadioSetting, bool)`

GetRadioSetting5g2Ok returns a tuple with the RadioSetting5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioSetting5g2

`func (o *OsgDetailVO) SetRadioSetting5g2(v ApRadioSetting)`

SetRadioSetting5g2 sets RadioSetting5g2 field to given value.

### HasRadioSetting5g2

`func (o *OsgDetailVO) HasRadioSetting5g2() bool`

HasRadioSetting5g2 returns a boolean if a field has been set.

### GetRadioSetting6g

`func (o *OsgDetailVO) GetRadioSetting6g() ApRadioSetting`

GetRadioSetting6g returns the RadioSetting6g field if non-nil, zero value otherwise.

### GetRadioSetting6gOk

`func (o *OsgDetailVO) GetRadioSetting6gOk() (*ApRadioSetting, bool)`

GetRadioSetting6gOk returns a tuple with the RadioSetting6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioSetting6g

`func (o *OsgDetailVO) SetRadioSetting6g(v ApRadioSetting)`

SetRadioSetting6g sets RadioSetting6g field to given value.

### HasRadioSetting6g

`func (o *OsgDetailVO) HasRadioSetting6g() bool`

HasRadioSetting6g returns a boolean if a field has been set.

### GetRadioTraffic2g

`func (o *OsgDetailVO) GetRadioTraffic2g() APRadioTrafficEntity`

GetRadioTraffic2g returns the RadioTraffic2g field if non-nil, zero value otherwise.

### GetRadioTraffic2gOk

`func (o *OsgDetailVO) GetRadioTraffic2gOk() (*APRadioTrafficEntity, bool)`

GetRadioTraffic2gOk returns a tuple with the RadioTraffic2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioTraffic2g

`func (o *OsgDetailVO) SetRadioTraffic2g(v APRadioTrafficEntity)`

SetRadioTraffic2g sets RadioTraffic2g field to given value.

### HasRadioTraffic2g

`func (o *OsgDetailVO) HasRadioTraffic2g() bool`

HasRadioTraffic2g returns a boolean if a field has been set.

### GetRadioTraffic5g

`func (o *OsgDetailVO) GetRadioTraffic5g() APRadioTrafficEntity`

GetRadioTraffic5g returns the RadioTraffic5g field if non-nil, zero value otherwise.

### GetRadioTraffic5gOk

`func (o *OsgDetailVO) GetRadioTraffic5gOk() (*APRadioTrafficEntity, bool)`

GetRadioTraffic5gOk returns a tuple with the RadioTraffic5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioTraffic5g

`func (o *OsgDetailVO) SetRadioTraffic5g(v APRadioTrafficEntity)`

SetRadioTraffic5g sets RadioTraffic5g field to given value.

### HasRadioTraffic5g

`func (o *OsgDetailVO) HasRadioTraffic5g() bool`

HasRadioTraffic5g returns a boolean if a field has been set.

### GetRadioTraffic5g2

`func (o *OsgDetailVO) GetRadioTraffic5g2() APRadioTrafficEntity`

GetRadioTraffic5g2 returns the RadioTraffic5g2 field if non-nil, zero value otherwise.

### GetRadioTraffic5g2Ok

`func (o *OsgDetailVO) GetRadioTraffic5g2Ok() (*APRadioTrafficEntity, bool)`

GetRadioTraffic5g2Ok returns a tuple with the RadioTraffic5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioTraffic5g2

`func (o *OsgDetailVO) SetRadioTraffic5g2(v APRadioTrafficEntity)`

SetRadioTraffic5g2 sets RadioTraffic5g2 field to given value.

### HasRadioTraffic5g2

`func (o *OsgDetailVO) HasRadioTraffic5g2() bool`

HasRadioTraffic5g2 returns a boolean if a field has been set.

### GetRadioTraffic6g

`func (o *OsgDetailVO) GetRadioTraffic6g() APRadioTrafficEntity`

GetRadioTraffic6g returns the RadioTraffic6g field if non-nil, zero value otherwise.

### GetRadioTraffic6gOk

`func (o *OsgDetailVO) GetRadioTraffic6gOk() (*APRadioTrafficEntity, bool)`

GetRadioTraffic6gOk returns a tuple with the RadioTraffic6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioTraffic6g

`func (o *OsgDetailVO) SetRadioTraffic6g(v APRadioTrafficEntity)`

SetRadioTraffic6g sets RadioTraffic6g field to given value.

### HasRadioTraffic6g

`func (o *OsgDetailVO) HasRadioTraffic6g() bool`

HasRadioTraffic6g returns a boolean if a field has been set.

### GetRadiosResource

`func (o *OsgDetailVO) GetRadiosResource() int32`

GetRadiosResource returns the RadiosResource field if non-nil, zero value otherwise.

### GetRadiosResourceOk

`func (o *OsgDetailVO) GetRadiosResourceOk() (*int32, bool)`

GetRadiosResourceOk returns a tuple with the RadiosResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiosResource

`func (o *OsgDetailVO) SetRadiosResource(v int32)`

SetRadiosResource sets RadiosResource field to given value.

### HasRadiosResource

`func (o *OsgDetailVO) HasRadiosResource() bool`

HasRadiosResource returns a boolean if a field has been set.

### GetRemember

`func (o *OsgDetailVO) GetRemember() bool`

GetRemember returns the Remember field if non-nil, zero value otherwise.

### GetRememberOk

`func (o *OsgDetailVO) GetRememberOk() (*bool, bool)`

GetRememberOk returns a tuple with the Remember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemember

`func (o *OsgDetailVO) SetRemember(v bool)`

SetRemember sets Remember field to given value.

### HasRemember

`func (o *OsgDetailVO) HasRemember() bool`

HasRemember returns a boolean if a field has been set.

### GetRememberDevice

`func (o *OsgDetailVO) GetRememberDevice() int32`

GetRememberDevice returns the RememberDevice field if non-nil, zero value otherwise.

### GetRememberDeviceOk

`func (o *OsgDetailVO) GetRememberDeviceOk() (*int32, bool)`

GetRememberDeviceOk returns a tuple with the RememberDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberDevice

`func (o *OsgDetailVO) SetRememberDevice(v int32)`

SetRememberDevice sets RememberDevice field to given value.

### HasRememberDevice

`func (o *OsgDetailVO) HasRememberDevice() bool`

HasRememberDevice returns a boolean if a field has been set.

### GetResource

`func (o *OsgDetailVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *OsgDetailVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *OsgDetailVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *OsgDetailVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetRps

`func (o *OsgDetailVO) GetRps() []OsgRpsStatusVO`

GetRps returns the Rps field if non-nil, zero value otherwise.

### GetRpsOk

`func (o *OsgDetailVO) GetRpsOk() (*[]OsgRpsStatusVO, bool)`

GetRpsOk returns a tuple with the Rps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRps

`func (o *OsgDetailVO) SetRps(v []OsgRpsStatusVO)`

SetRps sets Rps field to given value.

### HasRps

`func (o *OsgDetailVO) HasRps() bool`

HasRps returns a boolean if a field has been set.

### GetRssiSetting2g

`func (o *OsgDetailVO) GetRssiSetting2g() ApRssiThresholdVO`

GetRssiSetting2g returns the RssiSetting2g field if non-nil, zero value otherwise.

### GetRssiSetting2gOk

`func (o *OsgDetailVO) GetRssiSetting2gOk() (*ApRssiThresholdVO, bool)`

GetRssiSetting2gOk returns a tuple with the RssiSetting2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiSetting2g

`func (o *OsgDetailVO) SetRssiSetting2g(v ApRssiThresholdVO)`

SetRssiSetting2g sets RssiSetting2g field to given value.

### HasRssiSetting2g

`func (o *OsgDetailVO) HasRssiSetting2g() bool`

HasRssiSetting2g returns a boolean if a field has been set.

### GetRssiSetting5g

`func (o *OsgDetailVO) GetRssiSetting5g() ApRssiThresholdVO`

GetRssiSetting5g returns the RssiSetting5g field if non-nil, zero value otherwise.

### GetRssiSetting5gOk

`func (o *OsgDetailVO) GetRssiSetting5gOk() (*ApRssiThresholdVO, bool)`

GetRssiSetting5gOk returns a tuple with the RssiSetting5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiSetting5g

`func (o *OsgDetailVO) SetRssiSetting5g(v ApRssiThresholdVO)`

SetRssiSetting5g sets RssiSetting5g field to given value.

### HasRssiSetting5g

`func (o *OsgDetailVO) HasRssiSetting5g() bool`

HasRssiSetting5g returns a boolean if a field has been set.

### GetRssiSetting5g2

`func (o *OsgDetailVO) GetRssiSetting5g2() ApRssiThresholdVO`

GetRssiSetting5g2 returns the RssiSetting5g2 field if non-nil, zero value otherwise.

### GetRssiSetting5g2Ok

`func (o *OsgDetailVO) GetRssiSetting5g2Ok() (*ApRssiThresholdVO, bool)`

GetRssiSetting5g2Ok returns a tuple with the RssiSetting5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiSetting5g2

`func (o *OsgDetailVO) SetRssiSetting5g2(v ApRssiThresholdVO)`

SetRssiSetting5g2 sets RssiSetting5g2 field to given value.

### HasRssiSetting5g2

`func (o *OsgDetailVO) HasRssiSetting5g2() bool`

HasRssiSetting5g2 returns a boolean if a field has been set.

### GetRssiSetting6g

`func (o *OsgDetailVO) GetRssiSetting6g() ApRssiThresholdVO`

GetRssiSetting6g returns the RssiSetting6g field if non-nil, zero value otherwise.

### GetRssiSetting6gOk

`func (o *OsgDetailVO) GetRssiSetting6gOk() (*ApRssiThresholdVO, bool)`

GetRssiSetting6gOk returns a tuple with the RssiSetting6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiSetting6g

`func (o *OsgDetailVO) SetRssiSetting6g(v ApRssiThresholdVO)`

SetRssiSetting6g sets RssiSetting6g field to given value.

### HasRssiSetting6g

`func (o *OsgDetailVO) HasRssiSetting6g() bool`

HasRssiSetting6g returns a boolean if a field has been set.

### GetRxRate

`func (o *OsgDetailVO) GetRxRate() int64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *OsgDetailVO) GetRxRateOk() (*int64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *OsgDetailVO) SetRxRate(v int64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *OsgDetailVO) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetServicesResource

`func (o *OsgDetailVO) GetServicesResource() int32`

GetServicesResource returns the ServicesResource field if non-nil, zero value otherwise.

### GetServicesResourceOk

`func (o *OsgDetailVO) GetServicesResourceOk() (*int32, bool)`

GetServicesResourceOk returns a tuple with the ServicesResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesResource

`func (o *OsgDetailVO) SetServicesResource(v int32)`

SetServicesResource sets ServicesResource field to given value.

### HasServicesResource

`func (o *OsgDetailVO) HasServicesResource() bool`

HasServicesResource returns a boolean if a field has been set.

### GetShowModel

`func (o *OsgDetailVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *OsgDetailVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *OsgDetailVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *OsgDetailVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetSiteName

`func (o *OsgDetailVO) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *OsgDetailVO) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *OsgDetailVO) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *OsgDetailVO) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSiteTemplateId

`func (o *OsgDetailVO) GetSiteTemplateId() string`

GetSiteTemplateId returns the SiteTemplateId field if non-nil, zero value otherwise.

### GetSiteTemplateIdOk

`func (o *OsgDetailVO) GetSiteTemplateIdOk() (*string, bool)`

GetSiteTemplateIdOk returns a tuple with the SiteTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTemplateId

`func (o *OsgDetailVO) SetSiteTemplateId(v string)`

SetSiteTemplateId sets SiteTemplateId field to given value.

### HasSiteTemplateId

`func (o *OsgDetailVO) HasSiteTemplateId() bool`

HasSiteTemplateId returns a boolean if a field has been set.

### GetSiteTemplateName

`func (o *OsgDetailVO) GetSiteTemplateName() string`

GetSiteTemplateName returns the SiteTemplateName field if non-nil, zero value otherwise.

### GetSiteTemplateNameOk

`func (o *OsgDetailVO) GetSiteTemplateNameOk() (*string, bool)`

GetSiteTemplateNameOk returns a tuple with the SiteTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTemplateName

`func (o *OsgDetailVO) SetSiteTemplateName(v string)`

SetSiteTemplateName sets SiteTemplateName field to given value.

### HasSiteTemplateName

`func (o *OsgDetailVO) HasSiteTemplateName() bool`

HasSiteTemplateName returns a boolean if a field has been set.

### GetSn

`func (o *OsgDetailVO) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *OsgDetailVO) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *OsgDetailVO) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *OsgDetailVO) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetSnmpSeting

`func (o *OsgDetailVO) GetSnmpSeting() OsgSnmpVO`

GetSnmpSeting returns the SnmpSeting field if non-nil, zero value otherwise.

### GetSnmpSetingOk

`func (o *OsgDetailVO) GetSnmpSetingOk() (*OsgSnmpVO, bool)`

GetSnmpSetingOk returns a tuple with the SnmpSeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpSeting

`func (o *OsgDetailVO) SetSnmpSeting(v OsgSnmpVO)`

SetSnmpSeting sets SnmpSeting field to given value.

### HasSnmpSeting

`func (o *OsgDetailVO) HasSnmpSeting() bool`

HasSnmpSeting returns a boolean if a field has been set.

### GetSpecialModel

`func (o *OsgDetailVO) GetSpecialModel() string`

GetSpecialModel returns the SpecialModel field if non-nil, zero value otherwise.

### GetSpecialModelOk

`func (o *OsgDetailVO) GetSpecialModelOk() (*string, bool)`

GetSpecialModelOk returns a tuple with the SpecialModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialModel

`func (o *OsgDetailVO) SetSpecialModel(v string)`

SetSpecialModel sets SpecialModel field to given value.

### HasSpecialModel

`func (o *OsgDetailVO) HasSpecialModel() bool`

HasSpecialModel returns a boolean if a field has been set.

### GetSpeeds

`func (o *OsgDetailVO) GetSpeeds() []int32`

GetSpeeds returns the Speeds field if non-nil, zero value otherwise.

### GetSpeedsOk

`func (o *OsgDetailVO) GetSpeedsOk() (*[]int32, bool)`

GetSpeedsOk returns a tuple with the Speeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeeds

`func (o *OsgDetailVO) SetSpeeds(v []int32)`

SetSpeeds sets Speeds field to given value.

### HasSpeeds

`func (o *OsgDetailVO) HasSpeeds() bool`

HasSpeeds returns a boolean if a field has been set.

### GetSshSetting

`func (o *OsgDetailVO) GetSshSetting() SshSettingVO`

GetSshSetting returns the SshSetting field if non-nil, zero value otherwise.

### GetSshSettingOk

`func (o *OsgDetailVO) GetSshSettingOk() (*SshSettingVO, bool)`

GetSshSettingOk returns a tuple with the SshSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshSetting

`func (o *OsgDetailVO) SetSshSetting(v SshSettingVO)`

SetSshSetting sets SshSetting field to given value.

### HasSshSetting

`func (o *OsgDetailVO) HasSshSetting() bool`

HasSshSetting returns a boolean if a field has been set.

### GetSsidOverrides

`func (o *OsgDetailVO) GetSsidOverrides() []SsidOverrideVO`

GetSsidOverrides returns the SsidOverrides field if non-nil, zero value otherwise.

### GetSsidOverridesOk

`func (o *OsgDetailVO) GetSsidOverridesOk() (*[]SsidOverrideVO, bool)`

GetSsidOverridesOk returns a tuple with the SsidOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidOverrides

`func (o *OsgDetailVO) SetSsidOverrides(v []SsidOverrideVO)`

SetSsidOverrides sets SsidOverrides field to given value.

### HasSsidOverrides

`func (o *OsgDetailVO) HasSsidOverrides() bool`

HasSsidOverrides returns a boolean if a field has been set.

### GetStatus

`func (o *OsgDetailVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OsgDetailVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OsgDetailVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OsgDetailVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *OsgDetailVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *OsgDetailVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *OsgDetailVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *OsgDetailVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetSupportAnomaly

`func (o *OsgDetailVO) GetSupportAnomaly() bool`

GetSupportAnomaly returns the SupportAnomaly field if non-nil, zero value otherwise.

### GetSupportAnomalyOk

`func (o *OsgDetailVO) GetSupportAnomalyOk() (*bool, bool)`

GetSupportAnomalyOk returns a tuple with the SupportAnomaly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAnomaly

`func (o *OsgDetailVO) SetSupportAnomaly(v bool)`

SetSupportAnomaly sets SupportAnomaly field to given value.

### HasSupportAnomaly

`func (o *OsgDetailVO) HasSupportAnomaly() bool`

HasSupportAnomaly returns a boolean if a field has been set.

### GetSupportBandScan

`func (o *OsgDetailVO) GetSupportBandScan() bool`

GetSupportBandScan returns the SupportBandScan field if non-nil, zero value otherwise.

### GetSupportBandScanOk

`func (o *OsgDetailVO) GetSupportBandScanOk() (*bool, bool)`

GetSupportBandScanOk returns a tuple with the SupportBandScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportBandScan

`func (o *OsgDetailVO) SetSupportBandScan(v bool)`

SetSupportBandScan sets SupportBandScan field to given value.

### HasSupportBandScan

`func (o *OsgDetailVO) HasSupportBandScan() bool`

HasSupportBandScan returns a boolean if a field has been set.

### GetSupportBandWidthCtrl

`func (o *OsgDetailVO) GetSupportBandWidthCtrl() bool`

GetSupportBandWidthCtrl returns the SupportBandWidthCtrl field if non-nil, zero value otherwise.

### GetSupportBandWidthCtrlOk

`func (o *OsgDetailVO) GetSupportBandWidthCtrlOk() (*bool, bool)`

GetSupportBandWidthCtrlOk returns a tuple with the SupportBandWidthCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportBandWidthCtrl

`func (o *OsgDetailVO) SetSupportBandWidthCtrl(v bool)`

SetSupportBandWidthCtrl sets SupportBandWidthCtrl field to given value.

### HasSupportBandWidthCtrl

`func (o *OsgDetailVO) HasSupportBandWidthCtrl() bool`

HasSupportBandWidthCtrl returns a boolean if a field has been set.

### GetSupportBatchEdit

`func (o *OsgDetailVO) GetSupportBatchEdit() bool`

GetSupportBatchEdit returns the SupportBatchEdit field if non-nil, zero value otherwise.

### GetSupportBatchEditOk

`func (o *OsgDetailVO) GetSupportBatchEditOk() (*bool, bool)`

GetSupportBatchEditOk returns a tuple with the SupportBatchEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportBatchEdit

`func (o *OsgDetailVO) SetSupportBatchEdit(v bool)`

SetSupportBatchEdit sets SupportBatchEdit field to given value.

### HasSupportBatchEdit

`func (o *OsgDetailVO) HasSupportBatchEdit() bool`

HasSupportBatchEdit returns a boolean if a field has been set.

### GetSupportChannelUtilizationStatus

`func (o *OsgDetailVO) GetSupportChannelUtilizationStatus() bool`

GetSupportChannelUtilizationStatus returns the SupportChannelUtilizationStatus field if non-nil, zero value otherwise.

### GetSupportChannelUtilizationStatusOk

`func (o *OsgDetailVO) GetSupportChannelUtilizationStatusOk() (*bool, bool)`

GetSupportChannelUtilizationStatusOk returns a tuple with the SupportChannelUtilizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportChannelUtilizationStatus

`func (o *OsgDetailVO) SetSupportChannelUtilizationStatus(v bool)`

SetSupportChannelUtilizationStatus sets SupportChannelUtilizationStatus field to given value.

### HasSupportChannelUtilizationStatus

`func (o *OsgDetailVO) HasSupportChannelUtilizationStatus() bool`

HasSupportChannelUtilizationStatus returns a boolean if a field has been set.

### GetSupportDroppedPacketsStatus

`func (o *OsgDetailVO) GetSupportDroppedPacketsStatus() bool`

GetSupportDroppedPacketsStatus returns the SupportDroppedPacketsStatus field if non-nil, zero value otherwise.

### GetSupportDroppedPacketsStatusOk

`func (o *OsgDetailVO) GetSupportDroppedPacketsStatusOk() (*bool, bool)`

GetSupportDroppedPacketsStatusOk returns a tuple with the SupportDroppedPacketsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDroppedPacketsStatus

`func (o *OsgDetailVO) SetSupportDroppedPacketsStatus(v bool)`

SetSupportDroppedPacketsStatus sets SupportDroppedPacketsStatus field to given value.

### HasSupportDroppedPacketsStatus

`func (o *OsgDetailVO) HasSupportDroppedPacketsStatus() bool`

HasSupportDroppedPacketsStatus returns a boolean if a field has been set.

### GetSupportDsl

`func (o *OsgDetailVO) GetSupportDsl() bool`

GetSupportDsl returns the SupportDsl field if non-nil, zero value otherwise.

### GetSupportDslOk

`func (o *OsgDetailVO) GetSupportDslOk() (*bool, bool)`

GetSupportDslOk returns a tuple with the SupportDsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDsl

`func (o *OsgDetailVO) SetSupportDsl(v bool)`

SetSupportDsl sets SupportDsl field to given value.

### HasSupportDsl

`func (o *OsgDetailVO) HasSupportDsl() bool`

HasSupportDsl returns a boolean if a field has been set.

### GetSupportDualSim

`func (o *OsgDetailVO) GetSupportDualSim() int32`

GetSupportDualSim returns the SupportDualSim field if non-nil, zero value otherwise.

### GetSupportDualSimOk

`func (o *OsgDetailVO) GetSupportDualSimOk() (*int32, bool)`

GetSupportDualSimOk returns a tuple with the SupportDualSim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDualSim

`func (o *OsgDetailVO) SetSupportDualSim(v int32)`

SetSupportDualSim sets SupportDualSim field to given value.

### HasSupportDualSim

`func (o *OsgDetailVO) HasSupportDualSim() bool`

HasSupportDualSim returns a boolean if a field has been set.

### GetSupportFlowControl

`func (o *OsgDetailVO) GetSupportFlowControl() bool`

GetSupportFlowControl returns the SupportFlowControl field if non-nil, zero value otherwise.

### GetSupportFlowControlOk

`func (o *OsgDetailVO) GetSupportFlowControlOk() (*bool, bool)`

GetSupportFlowControlOk returns a tuple with the SupportFlowControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportFlowControl

`func (o *OsgDetailVO) SetSupportFlowControl(v bool)`

SetSupportFlowControl sets SupportFlowControl field to given value.

### HasSupportFlowControl

`func (o *OsgDetailVO) HasSupportFlowControl() bool`

HasSupportFlowControl returns a boolean if a field has been set.

### GetSupportHwOffload

`func (o *OsgDetailVO) GetSupportHwOffload() bool`

GetSupportHwOffload returns the SupportHwOffload field if non-nil, zero value otherwise.

### GetSupportHwOffloadOk

`func (o *OsgDetailVO) GetSupportHwOffloadOk() (*bool, bool)`

GetSupportHwOffloadOk returns a tuple with the SupportHwOffload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportHwOffload

`func (o *OsgDetailVO) SetSupportHwOffload(v bool)`

SetSupportHwOffload sets SupportHwOffload field to given value.

### HasSupportHwOffload

`func (o *OsgDetailVO) HasSupportHwOffload() bool`

HasSupportHwOffload returns a boolean if a field has been set.

### GetSupportIpMacBinding

`func (o *OsgDetailVO) GetSupportIpMacBinding() bool`

GetSupportIpMacBinding returns the SupportIpMacBinding field if non-nil, zero value otherwise.

### GetSupportIpMacBindingOk

`func (o *OsgDetailVO) GetSupportIpMacBindingOk() (*bool, bool)`

GetSupportIpMacBindingOk returns a tuple with the SupportIpMacBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIpMacBinding

`func (o *OsgDetailVO) SetSupportIpMacBinding(v bool)`

SetSupportIpMacBinding sets SupportIpMacBinding field to given value.

### HasSupportIpMacBinding

`func (o *OsgDetailVO) HasSupportIpMacBinding() bool`

HasSupportIpMacBinding returns a boolean if a field has been set.

### GetSupportIppt

`func (o *OsgDetailVO) GetSupportIppt() bool`

GetSupportIppt returns the SupportIppt field if non-nil, zero value otherwise.

### GetSupportIpptOk

`func (o *OsgDetailVO) GetSupportIpptOk() (*bool, bool)`

GetSupportIpptOk returns a tuple with the SupportIppt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIppt

`func (o *OsgDetailVO) SetSupportIppt(v bool)`

SetSupportIppt sets SupportIppt field to given value.

### HasSupportIppt

`func (o *OsgDetailVO) HasSupportIppt() bool`

HasSupportIppt returns a boolean if a field has been set.

### GetSupportJumbo

`func (o *OsgDetailVO) GetSupportJumbo() bool`

GetSupportJumbo returns the SupportJumbo field if non-nil, zero value otherwise.

### GetSupportJumboOk

`func (o *OsgDetailVO) GetSupportJumboOk() (*bool, bool)`

GetSupportJumboOk returns a tuple with the SupportJumbo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportJumbo

`func (o *OsgDetailVO) SetSupportJumbo(v bool)`

SetSupportJumbo sets SupportJumbo field to given value.

### HasSupportJumbo

`func (o *OsgDetailVO) HasSupportJumbo() bool`

HasSupportJumbo returns a boolean if a field has been set.

### GetSupportLanClientStats

`func (o *OsgDetailVO) GetSupportLanClientStats() bool`

GetSupportLanClientStats returns the SupportLanClientStats field if non-nil, zero value otherwise.

### GetSupportLanClientStatsOk

`func (o *OsgDetailVO) GetSupportLanClientStatsOk() (*bool, bool)`

GetSupportLanClientStatsOk returns a tuple with the SupportLanClientStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLanClientStats

`func (o *OsgDetailVO) SetSupportLanClientStats(v bool)`

SetSupportLanClientStats sets SupportLanClientStats field to given value.

### HasSupportLanClientStats

`func (o *OsgDetailVO) HasSupportLanClientStats() bool`

HasSupportLanClientStats returns a boolean if a field has been set.

### GetSupportLocatePort

`func (o *OsgDetailVO) GetSupportLocatePort() bool`

GetSupportLocatePort returns the SupportLocatePort field if non-nil, zero value otherwise.

### GetSupportLocatePortOk

`func (o *OsgDetailVO) GetSupportLocatePortOk() (*bool, bool)`

GetSupportLocatePortOk returns a tuple with the SupportLocatePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLocatePort

`func (o *OsgDetailVO) SetSupportLocatePort(v bool)`

SetSupportLocatePort sets SupportLocatePort field to given value.

### HasSupportLocatePort

`func (o *OsgDetailVO) HasSupportLocatePort() bool`

HasSupportLocatePort returns a boolean if a field has been set.

### GetSupportLoopbackControl

`func (o *OsgDetailVO) GetSupportLoopbackControl() bool`

GetSupportLoopbackControl returns the SupportLoopbackControl field if non-nil, zero value otherwise.

### GetSupportLoopbackControlOk

`func (o *OsgDetailVO) GetSupportLoopbackControlOk() (*bool, bool)`

GetSupportLoopbackControlOk returns a tuple with the SupportLoopbackControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLoopbackControl

`func (o *OsgDetailVO) SetSupportLoopbackControl(v bool)`

SetSupportLoopbackControl sets SupportLoopbackControl field to given value.

### HasSupportLoopbackControl

`func (o *OsgDetailVO) HasSupportLoopbackControl() bool`

HasSupportLoopbackControl returns a boolean if a field has been set.

### GetSupportLte

`func (o *OsgDetailVO) GetSupportLte() bool`

GetSupportLte returns the SupportLte field if non-nil, zero value otherwise.

### GetSupportLteOk

`func (o *OsgDetailVO) GetSupportLteOk() (*bool, bool)`

GetSupportLteOk returns a tuple with the SupportLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLte

`func (o *OsgDetailVO) SetSupportLte(v bool)`

SetSupportLte sets SupportLte field to given value.

### HasSupportLte

`func (o *OsgDetailVO) HasSupportLte() bool`

HasSupportLte returns a boolean if a field has been set.

### GetSupportMirror

`func (o *OsgDetailVO) GetSupportMirror() bool`

GetSupportMirror returns the SupportMirror field if non-nil, zero value otherwise.

### GetSupportMirrorOk

`func (o *OsgDetailVO) GetSupportMirrorOk() (*bool, bool)`

GetSupportMirrorOk returns a tuple with the SupportMirror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMirror

`func (o *OsgDetailVO) SetSupportMirror(v bool)`

SetSupportMirror sets SupportMirror field to given value.

### HasSupportMirror

`func (o *OsgDetailVO) HasSupportMirror() bool`

HasSupportMirror returns a boolean if a field has been set.

### GetSupportNetworkSearch

`func (o *OsgDetailVO) GetSupportNetworkSearch() bool`

GetSupportNetworkSearch returns the SupportNetworkSearch field if non-nil, zero value otherwise.

### GetSupportNetworkSearchOk

`func (o *OsgDetailVO) GetSupportNetworkSearchOk() (*bool, bool)`

GetSupportNetworkSearchOk returns a tuple with the SupportNetworkSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportNetworkSearch

`func (o *OsgDetailVO) SetSupportNetworkSearch(v bool)`

SetSupportNetworkSearch sets SupportNetworkSearch field to given value.

### HasSupportNetworkSearch

`func (o *OsgDetailVO) HasSupportNetworkSearch() bool`

HasSupportNetworkSearch returns a boolean if a field has been set.

### GetSupportNewIpMacBinding

`func (o *OsgDetailVO) GetSupportNewIpMacBinding() bool`

GetSupportNewIpMacBinding returns the SupportNewIpMacBinding field if non-nil, zero value otherwise.

### GetSupportNewIpMacBindingOk

`func (o *OsgDetailVO) GetSupportNewIpMacBindingOk() (*bool, bool)`

GetSupportNewIpMacBindingOk returns a tuple with the SupportNewIpMacBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportNewIpMacBinding

`func (o *OsgDetailVO) SetSupportNewIpMacBinding(v bool)`

SetSupportNewIpMacBinding sets SupportNewIpMacBinding field to given value.

### HasSupportNewIpMacBinding

`func (o *OsgDetailVO) HasSupportNewIpMacBinding() bool`

HasSupportNewIpMacBinding returns a boolean if a field has been set.

### GetSupportPoe

`func (o *OsgDetailVO) GetSupportPoe() bool`

GetSupportPoe returns the SupportPoe field if non-nil, zero value otherwise.

### GetSupportPoeOk

`func (o *OsgDetailVO) GetSupportPoeOk() (*bool, bool)`

GetSupportPoeOk returns a tuple with the SupportPoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPoe

`func (o *OsgDetailVO) SetSupportPoe(v bool)`

SetSupportPoe sets SupportPoe field to given value.

### HasSupportPoe

`func (o *OsgDetailVO) HasSupportPoe() bool`

HasSupportPoe returns a boolean if a field has been set.

### GetSupportPortControl

`func (o *OsgDetailVO) GetSupportPortControl() bool`

GetSupportPortControl returns the SupportPortControl field if non-nil, zero value otherwise.

### GetSupportPortControlOk

`func (o *OsgDetailVO) GetSupportPortControlOk() (*bool, bool)`

GetSupportPortControlOk returns a tuple with the SupportPortControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPortControl

`func (o *OsgDetailVO) SetSupportPortControl(v bool)`

SetSupportPortControl sets SupportPortControl field to given value.

### HasSupportPortControl

`func (o *OsgDetailVO) HasSupportPortControl() bool`

HasSupportPortControl returns a boolean if a field has been set.

### GetSupportPortForwardingStatus

`func (o *OsgDetailVO) GetSupportPortForwardingStatus() bool`

GetSupportPortForwardingStatus returns the SupportPortForwardingStatus field if non-nil, zero value otherwise.

### GetSupportPortForwardingStatusOk

`func (o *OsgDetailVO) GetSupportPortForwardingStatusOk() (*bool, bool)`

GetSupportPortForwardingStatusOk returns a tuple with the SupportPortForwardingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPortForwardingStatus

`func (o *OsgDetailVO) SetSupportPortForwardingStatus(v bool)`

SetSupportPortForwardingStatus sets SupportPortForwardingStatus field to given value.

### HasSupportPortForwardingStatus

`func (o *OsgDetailVO) HasSupportPortForwardingStatus() bool`

HasSupportPortForwardingStatus returns a boolean if a field has been set.

### GetSupportPortIsolation

`func (o *OsgDetailVO) GetSupportPortIsolation() bool`

GetSupportPortIsolation returns the SupportPortIsolation field if non-nil, zero value otherwise.

### GetSupportPortIsolationOk

`func (o *OsgDetailVO) GetSupportPortIsolationOk() (*bool, bool)`

GetSupportPortIsolationOk returns a tuple with the SupportPortIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPortIsolation

`func (o *OsgDetailVO) SetSupportPortIsolation(v bool)`

SetSupportPortIsolation sets SupportPortIsolation field to given value.

### HasSupportPortIsolation

`func (o *OsgDetailVO) HasSupportPortIsolation() bool`

HasSupportPortIsolation returns a boolean if a field has been set.

### GetSupportPvid

`func (o *OsgDetailVO) GetSupportPvid() bool`

GetSupportPvid returns the SupportPvid field if non-nil, zero value otherwise.

### GetSupportPvidOk

`func (o *OsgDetailVO) GetSupportPvidOk() (*bool, bool)`

GetSupportPvidOk returns a tuple with the SupportPvid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPvid

`func (o *OsgDetailVO) SetSupportPvid(v bool)`

SetSupportPvid sets SupportPvid field to given value.

### HasSupportPvid

`func (o *OsgDetailVO) HasSupportPvid() bool`

HasSupportPvid returns a boolean if a field has been set.

### GetSupportRFScan

`func (o *OsgDetailVO) GetSupportRFScan() bool`

GetSupportRFScan returns the SupportRFScan field if non-nil, zero value otherwise.

### GetSupportRFScanOk

`func (o *OsgDetailVO) GetSupportRFScanOk() (*bool, bool)`

GetSupportRFScanOk returns a tuple with the SupportRFScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportRFScan

`func (o *OsgDetailVO) SetSupportRFScan(v bool)`

SetSupportRFScan sets SupportRFScan field to given value.

### HasSupportRFScan

`func (o *OsgDetailVO) HasSupportRFScan() bool`

HasSupportRFScan returns a boolean if a field has been set.

### GetSupportRetriedPacketsStatus

`func (o *OsgDetailVO) GetSupportRetriedPacketsStatus() bool`

GetSupportRetriedPacketsStatus returns the SupportRetriedPacketsStatus field if non-nil, zero value otherwise.

### GetSupportRetriedPacketsStatusOk

`func (o *OsgDetailVO) GetSupportRetriedPacketsStatusOk() (*bool, bool)`

GetSupportRetriedPacketsStatusOk returns a tuple with the SupportRetriedPacketsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportRetriedPacketsStatus

`func (o *OsgDetailVO) SetSupportRetriedPacketsStatus(v bool)`

SetSupportRetriedPacketsStatus sets SupportRetriedPacketsStatus field to given value.

### HasSupportRetriedPacketsStatus

`func (o *OsgDetailVO) HasSupportRetriedPacketsStatus() bool`

HasSupportRetriedPacketsStatus returns a boolean if a field has been set.

### GetSupportSessionLimitStatus

`func (o *OsgDetailVO) GetSupportSessionLimitStatus() bool`

GetSupportSessionLimitStatus returns the SupportSessionLimitStatus field if non-nil, zero value otherwise.

### GetSupportSessionLimitStatusOk

`func (o *OsgDetailVO) GetSupportSessionLimitStatusOk() (*bool, bool)`

GetSupportSessionLimitStatusOk returns a tuple with the SupportSessionLimitStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSessionLimitStatus

`func (o *OsgDetailVO) SetSupportSessionLimitStatus(v bool)`

SetSupportSessionLimitStatus sets SupportSessionLimitStatus field to given value.

### HasSupportSessionLimitStatus

`func (o *OsgDetailVO) HasSupportSessionLimitStatus() bool`

HasSupportSessionLimitStatus returns a boolean if a field has been set.

### GetSupportSnmp

`func (o *OsgDetailVO) GetSupportSnmp() bool`

GetSupportSnmp returns the SupportSnmp field if non-nil, zero value otherwise.

### GetSupportSnmpOk

`func (o *OsgDetailVO) GetSupportSnmpOk() (*bool, bool)`

GetSupportSnmpOk returns a tuple with the SupportSnmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSnmp

`func (o *OsgDetailVO) SetSupportSnmp(v bool)`

SetSupportSnmp sets SupportSnmp field to given value.

### HasSupportSnmp

`func (o *OsgDetailVO) HasSupportSnmp() bool`

HasSupportSnmp returns a boolean if a field has been set.

### GetSupportSpeedDuplex

`func (o *OsgDetailVO) GetSupportSpeedDuplex() bool`

GetSupportSpeedDuplex returns the SupportSpeedDuplex field if non-nil, zero value otherwise.

### GetSupportSpeedDuplexOk

`func (o *OsgDetailVO) GetSupportSpeedDuplexOk() (*bool, bool)`

GetSupportSpeedDuplexOk returns a tuple with the SupportSpeedDuplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSpeedDuplex

`func (o *OsgDetailVO) SetSupportSpeedDuplex(v bool)`

SetSupportSpeedDuplex sets SupportSpeedDuplex field to given value.

### HasSupportSpeedDuplex

`func (o *OsgDetailVO) HasSupportSpeedDuplex() bool`

HasSupportSpeedDuplex returns a boolean if a field has been set.

### GetSupportStormCtrlAction

`func (o *OsgDetailVO) GetSupportStormCtrlAction() bool`

GetSupportStormCtrlAction returns the SupportStormCtrlAction field if non-nil, zero value otherwise.

### GetSupportStormCtrlActionOk

`func (o *OsgDetailVO) GetSupportStormCtrlActionOk() (*bool, bool)`

GetSupportStormCtrlActionOk returns a tuple with the SupportStormCtrlAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportStormCtrlAction

`func (o *OsgDetailVO) SetSupportStormCtrlAction(v bool)`

SetSupportStormCtrlAction sets SupportStormCtrlAction field to given value.

### HasSupportStormCtrlAction

`func (o *OsgDetailVO) HasSupportStormCtrlAction() bool`

HasSupportStormCtrlAction returns a boolean if a field has been set.

### GetSupportVirtualWan

`func (o *OsgDetailVO) GetSupportVirtualWan() bool`

GetSupportVirtualWan returns the SupportVirtualWan field if non-nil, zero value otherwise.

### GetSupportVirtualWanOk

`func (o *OsgDetailVO) GetSupportVirtualWanOk() (*bool, bool)`

GetSupportVirtualWanOk returns a tuple with the SupportVirtualWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVirtualWan

`func (o *OsgDetailVO) SetSupportVirtualWan(v bool)`

SetSupportVirtualWan sets SupportVirtualWan field to given value.

### HasSupportVirtualWan

`func (o *OsgDetailVO) HasSupportVirtualWan() bool`

HasSupportVirtualWan returns a boolean if a field has been set.

### GetSupportWanSetPvid

`func (o *OsgDetailVO) GetSupportWanSetPvid() bool`

GetSupportWanSetPvid returns the SupportWanSetPvid field if non-nil, zero value otherwise.

### GetSupportWanSetPvidOk

`func (o *OsgDetailVO) GetSupportWanSetPvidOk() (*bool, bool)`

GetSupportWanSetPvidOk returns a tuple with the SupportWanSetPvid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportWanSetPvid

`func (o *OsgDetailVO) SetSupportWanSetPvid(v bool)`

SetSupportWanSetPvid sets SupportWanSetPvid field to given value.

### HasSupportWanSetPvid

`func (o *OsgDetailVO) HasSupportWanSetPvid() bool`

HasSupportWanSetPvid returns a boolean if a field has been set.

### GetTagIds

`func (o *OsgDetailVO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *OsgDetailVO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *OsgDetailVO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *OsgDetailVO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTemp

`func (o *OsgDetailVO) GetTemp() int32`

GetTemp returns the Temp field if non-nil, zero value otherwise.

### GetTempOk

`func (o *OsgDetailVO) GetTempOk() (*int32, bool)`

GetTempOk returns a tuple with the Temp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemp

`func (o *OsgDetailVO) SetTemp(v int32)`

SetTemp sets Temp field to given value.

### HasTemp

`func (o *OsgDetailVO) HasTemp() bool`

HasTemp returns a boolean if a field has been set.

### GetTemplateId

`func (o *OsgDetailVO) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *OsgDetailVO) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *OsgDetailVO) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *OsgDetailVO) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateName

`func (o *OsgDetailVO) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *OsgDetailVO) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *OsgDetailVO) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *OsgDetailVO) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetTemplateSettings

`func (o *OsgDetailVO) GetTemplateSettings() []int32`

GetTemplateSettings returns the TemplateSettings field if non-nil, zero value otherwise.

### GetTemplateSettingsOk

`func (o *OsgDetailVO) GetTemplateSettingsOk() (*[]int32, bool)`

GetTemplateSettingsOk returns a tuple with the TemplateSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSettings

`func (o *OsgDetailVO) SetTemplateSettings(v []int32)`

SetTemplateSettings sets TemplateSettings field to given value.

### HasTemplateSettings

`func (o *OsgDetailVO) HasTemplateSettings() bool`

HasTemplateSettings returns a boolean if a field has been set.

### GetTxRate

`func (o *OsgDetailVO) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *OsgDetailVO) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *OsgDetailVO) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *OsgDetailVO) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.

### GetType

`func (o *OsgDetailVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OsgDetailVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OsgDetailVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OsgDetailVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnsupportedPorts

`func (o *OsgDetailVO) GetUnsupportedPorts() []int32`

GetUnsupportedPorts returns the UnsupportedPorts field if non-nil, zero value otherwise.

### GetUnsupportedPortsOk

`func (o *OsgDetailVO) GetUnsupportedPortsOk() (*[]int32, bool)`

GetUnsupportedPortsOk returns a tuple with the UnsupportedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsupportedPorts

`func (o *OsgDetailVO) SetUnsupportedPorts(v []int32)`

SetUnsupportedPorts sets UnsupportedPorts field to given value.

### HasUnsupportedPorts

`func (o *OsgDetailVO) HasUnsupportedPorts() bool`

HasUnsupportedPorts returns a boolean if a field has been set.

### GetUpload

`func (o *OsgDetailVO) GetUpload() int64`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *OsgDetailVO) GetUploadOk() (*int64, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *OsgDetailVO) SetUpload(v int64)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *OsgDetailVO) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetUptime

`func (o *OsgDetailVO) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *OsgDetailVO) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *OsgDetailVO) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *OsgDetailVO) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetUptimeLong

`func (o *OsgDetailVO) GetUptimeLong() int64`

GetUptimeLong returns the UptimeLong field if non-nil, zero value otherwise.

### GetUptimeLongOk

`func (o *OsgDetailVO) GetUptimeLongOk() (*int64, bool)`

GetUptimeLongOk returns a tuple with the UptimeLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeLong

`func (o *OsgDetailVO) SetUptimeLong(v int64)`

SetUptimeLong sets UptimeLong field to given value.

### HasUptimeLong

`func (o *OsgDetailVO) HasUptimeLong() bool`

HasUptimeLong returns a boolean if a field has been set.

### GetUserNum

`func (o *OsgDetailVO) GetUserNum() int32`

GetUserNum returns the UserNum field if non-nil, zero value otherwise.

### GetUserNumOk

`func (o *OsgDetailVO) GetUserNumOk() (*int32, bool)`

GetUserNumOk returns a tuple with the UserNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNum

`func (o *OsgDetailVO) SetUserNum(v int32)`

SetUserNum sets UserNum field to given value.

### HasUserNum

`func (o *OsgDetailVO) HasUserNum() bool`

HasUserNum returns a boolean if a field has been set.

### GetVersion

`func (o *OsgDetailVO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OsgDetailVO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OsgDetailVO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OsgDetailVO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVirtualWanStats

`func (o *OsgDetailVO) GetVirtualWanStats() []OsgVirtualWanStatVO`

GetVirtualWanStats returns the VirtualWanStats field if non-nil, zero value otherwise.

### GetVirtualWanStatsOk

`func (o *OsgDetailVO) GetVirtualWanStatsOk() (*[]OsgVirtualWanStatVO, bool)`

GetVirtualWanStatsOk returns a tuple with the VirtualWanStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanStats

`func (o *OsgDetailVO) SetVirtualWanStats(v []OsgVirtualWanStatVO)`

SetVirtualWanStats sets VirtualWanStats field to given value.

### HasVirtualWanStats

`func (o *OsgDetailVO) HasVirtualWanStats() bool`

HasVirtualWanStats returns a boolean if a field has been set.

### GetWirelessAdvancedResource

`func (o *OsgDetailVO) GetWirelessAdvancedResource() int32`

GetWirelessAdvancedResource returns the WirelessAdvancedResource field if non-nil, zero value otherwise.

### GetWirelessAdvancedResourceOk

`func (o *OsgDetailVO) GetWirelessAdvancedResourceOk() (*int32, bool)`

GetWirelessAdvancedResourceOk returns a tuple with the WirelessAdvancedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessAdvancedResource

`func (o *OsgDetailVO) SetWirelessAdvancedResource(v int32)`

SetWirelessAdvancedResource sets WirelessAdvancedResource field to given value.

### HasWirelessAdvancedResource

`func (o *OsgDetailVO) HasWirelessAdvancedResource() bool`

HasWirelessAdvancedResource returns a boolean if a field has been set.

### GetWirelessHealth

`func (o *OsgDetailVO) GetWirelessHealth() bool`

GetWirelessHealth returns the WirelessHealth field if non-nil, zero value otherwise.

### GetWirelessHealthOk

`func (o *OsgDetailVO) GetWirelessHealthOk() (*bool, bool)`

GetWirelessHealthOk returns a tuple with the WirelessHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessHealth

`func (o *OsgDetailVO) SetWirelessHealth(v bool)`

SetWirelessHealth sets WirelessHealth field to given value.

### HasWirelessHealth

`func (o *OsgDetailVO) HasWirelessHealth() bool`

HasWirelessHealth returns a boolean if a field has been set.

### GetWirelessRouter

`func (o *OsgDetailVO) GetWirelessRouter() bool`

GetWirelessRouter returns the WirelessRouter field if non-nil, zero value otherwise.

### GetWirelessRouterOk

`func (o *OsgDetailVO) GetWirelessRouterOk() (*bool, bool)`

GetWirelessRouterOk returns a tuple with the WirelessRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessRouter

`func (o *OsgDetailVO) SetWirelessRouter(v bool)`

SetWirelessRouter sets WirelessRouter field to given value.

### HasWirelessRouter

`func (o *OsgDetailVO) HasWirelessRouter() bool`

HasWirelessRouter returns a boolean if a field has been set.

### GetWlanId

`func (o *OsgDetailVO) GetWlanId() string`

GetWlanId returns the WlanId field if non-nil, zero value otherwise.

### GetWlanIdOk

`func (o *OsgDetailVO) GetWlanIdOk() (*string, bool)`

GetWlanIdOk returns a tuple with the WlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanId

`func (o *OsgDetailVO) SetWlanId(v string)`

SetWlanId sets WlanId field to given value.

### HasWlanId

`func (o *OsgDetailVO) HasWlanId() bool`

HasWlanId returns a boolean if a field has been set.

### GetWlansResource

`func (o *OsgDetailVO) GetWlansResource() int32`

GetWlansResource returns the WlansResource field if non-nil, zero value otherwise.

### GetWlansResourceOk

`func (o *OsgDetailVO) GetWlansResourceOk() (*int32, bool)`

GetWlansResourceOk returns a tuple with the WlansResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlansResource

`func (o *OsgDetailVO) SetWlansResource(v int32)`

SetWlansResource sets WlansResource field to given value.

### HasWlansResource

`func (o *OsgDetailVO) HasWlansResource() bool`

HasWlansResource returns a boolean if a field has been set.

### GetWp2g

`func (o *OsgDetailVO) GetWp2g() ApRadioChannel`

GetWp2g returns the Wp2g field if non-nil, zero value otherwise.

### GetWp2gOk

`func (o *OsgDetailVO) GetWp2gOk() (*ApRadioChannel, bool)`

GetWp2gOk returns a tuple with the Wp2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWp2g

`func (o *OsgDetailVO) SetWp2g(v ApRadioChannel)`

SetWp2g sets Wp2g field to given value.

### HasWp2g

`func (o *OsgDetailVO) HasWp2g() bool`

HasWp2g returns a boolean if a field has been set.

### GetWp5g

`func (o *OsgDetailVO) GetWp5g() ApRadioChannel`

GetWp5g returns the Wp5g field if non-nil, zero value otherwise.

### GetWp5gOk

`func (o *OsgDetailVO) GetWp5gOk() (*ApRadioChannel, bool)`

GetWp5gOk returns a tuple with the Wp5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWp5g

`func (o *OsgDetailVO) SetWp5g(v ApRadioChannel)`

SetWp5g sets Wp5g field to given value.

### HasWp5g

`func (o *OsgDetailVO) HasWp5g() bool`

HasWp5g returns a boolean if a field has been set.

### GetWp5g2

`func (o *OsgDetailVO) GetWp5g2() ApRadioChannel`

GetWp5g2 returns the Wp5g2 field if non-nil, zero value otherwise.

### GetWp5g2Ok

`func (o *OsgDetailVO) GetWp5g2Ok() (*ApRadioChannel, bool)`

GetWp5g2Ok returns a tuple with the Wp5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWp5g2

`func (o *OsgDetailVO) SetWp5g2(v ApRadioChannel)`

SetWp5g2 sets Wp5g2 field to given value.

### HasWp5g2

`func (o *OsgDetailVO) HasWp5g2() bool`

HasWp5g2 returns a boolean if a field has been set.

### GetWp6g

`func (o *OsgDetailVO) GetWp6g() ApRadioChannel`

GetWp6g returns the Wp6g field if non-nil, zero value otherwise.

### GetWp6gOk

`func (o *OsgDetailVO) GetWp6gOk() (*ApRadioChannel, bool)`

GetWp6gOk returns a tuple with the Wp6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWp6g

`func (o *OsgDetailVO) SetWp6g(v ApRadioChannel)`

SetWp6g sets Wp6g field to given value.

### HasWp6g

`func (o *OsgDetailVO) HasWp6g() bool`

HasWp6g returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


