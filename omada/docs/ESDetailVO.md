# ESDetailVO

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
**CustomId** | Pointer to **string** | Customer ID | [optional] 
**CustomName** | Pointer to **string** | Customer name | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DevCap** | Pointer to [**OswDevCapVO**](OswDevCapVO.md) |  | [optional] 
**DeviceMisc** | Pointer to [**OswDeviceMiscVO**](OswDeviceMiscVO.md) |  | [optional] 
**DeviceSeriesType** | Pointer to **int32** | Device series type.DeviceSeriesType should be a value as follows: 0:advanced;1:pro | [optional] 
**DeviceTemplateAvailable** | Pointer to **bool** | Whether there is an available device template for the device; it is false if the model is not supported or the site template has not created the corresponding device template. | [optional] 
**DisableHwReset** | Pointer to **bool** |  | [optional] 
**DownlinkList** | Pointer to [**[]OswDownlinkVO**](OswDownlinkVO.md) |  | [optional] 
**Download** | Pointer to **int64** |  | [optional] 
**DueTime** | Pointer to **int64** | Expire timestamp of license(cloud base exclusive) | [optional] 
**DueTimeLeft** | Pointer to **int64** | Milliseconds from the current moment to the expiration time(cloud base exclusive) | [optional] 
**EcspFirstVersion** | Pointer to **int32** | Ecsp first version | [optional] 
**Eos** | Pointer to **int64** | End of support time of device(CBC exclusive) | [optional] 
**Eost** | Pointer to **int64** | End of service time of device(CBC exclusive) | [optional] 
**Es** | Pointer to **bool** | Whether the device is Agile Series Switch | [optional] 
**FanStatus** | Pointer to **int32** |  | [optional] 
**FirmwareVersion** | Pointer to **string** | Version of firmware,for example:2.5.0 Build 20190118 Rel. 64821 | [optional] 
**ForgetId** | Pointer to **string** | Forget ID of device | [optional] 
**HwVersion** | Pointer to **string** | Version of hardware,for example 1.0 | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InWhitelist** | Pointer to **bool** | Whether the device is in white list | [optional] 
**InitialUnbindingLimit** | Pointer to **int32** | Initial unbind count for license(cloud base exclusive) | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**IpSetting** | Pointer to [**IpSettingVO**](IpSettingVO.md) |  | [optional] 
**JumboEnable** | Pointer to **bool** |  | [optional] 
**Lags** | Pointer to [**[]OswLagVO**](OswLagVO.md) |  | [optional] 
**LastSeen** | Pointer to **int64** |  | [optional] 
**LedSetting** | Pointer to **int32** |  | [optional] 
**LicenseId** | Pointer to **string** | License key on detail page of device(cloud base exclusive) | [optional] 
**LicenseStatus** | Pointer to **int32** | License status(cloud base exclusive).LicenseStatus should be a value as follows: 0:unActive 1:Unbind 2:Expired 3:active | [optional] 
**LicenseUnbindingLimit** | Pointer to **int32** | Remaining unbind count for license on detail Page of device(cloud base exclusive) | [optional] 
**Location** | Pointer to [**DeviceLocationDetailVO**](DeviceLocationDetailVO.md) |  | [optional] 
**Loop** | Pointer to **string** |  | [optional] 
**LoopbackDetectEnable** | Pointer to **bool** |  | [optional] 
**LoopbackNum** | Pointer to **int32** |  | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**ManagerMark** | Pointer to **int32** |  | [optional] 
**Model** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**MoveSiteId** | Pointer to **string** | Record that the device is in a moveSite operation; if it is null, then it is not in the moveSite operation. | [optional] 
**Multicast** | Pointer to [**OswLanMulticastVO**](OswLanMulticastVO.md) |  | [optional] 
**MvlanBridgeVlan** | Pointer to **int32** |  | [optional] 
**MvlanNetworkId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Device name,default value is the mac address of device | [optional] 
**NeedUpgrade** | Pointer to **bool** |  | [optional] 
**NetworkNotify** | Pointer to **bool** |  | [optional] 
**OmadacId** | Pointer to **string** | OmadacId of the device | [optional] 
**PoeRemain** | Pointer to **float64** |  | [optional] 
**PoeRemainPercent** | Pointer to **float64** |  | [optional] 
**PoeTotalPower** | Pointer to **float64** |  | [optional] 
**Ports** | Pointer to [**[]OswPortVO**](OswPortVO.md) |  | [optional] 
**PowerAlertEnable** | Pointer to **bool** |  | [optional] 
**PowerStatusList** | Pointer to [**[]OswPortAlertStatusVO**](OswPortAlertStatusVO.md) |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**PublicIp** | Pointer to **string** |  | [optional] 
**Remember** | Pointer to **bool** | Whether to remember the device(deprecated) | [optional] 
**RememberDevice** | Pointer to **int32** | Whether to remember the device.RememberDevice should be a value as follows: 0:off, 1:on, 2: follow site | [optional] 
**Resource** | Pointer to **int32** | Data source.Resource should be a value as follows: 0:new created;1:from template;2:override | [optional] 
**RxRate** | Pointer to **int64** | Rx Rate | [optional] 
**ShowModel** | Pointer to **string** | Model complex shown in the front end.Ap：model+(country)+modelVersion,EAP225(EU) v3.0  Gateway/Switch：model+modelVersion,Osg v3.0 | [optional] 
**SiteName** | Pointer to **string** | Site name of the device | [optional] 
**SiteTemplateId** | Pointer to **string** | Template ID bound to the site | [optional] 
**SiteTemplateName** | Pointer to **string** | Template name bound to the site | [optional] 
**Sn** | Pointer to **string** |  | [optional] 
**SpecialModel** | Pointer to **string** | Special device model,for example:EAP225-Outdoor-1a20a950b8d950e8 | [optional] 
**Speeds** | Pointer to **[]int32** |  | [optional] 
**Status** | Pointer to **int32** | Status of device,status should be a value as follows: 0:Disconnected;1:Disconnected(Migrating);10:Provisioning;11:Configuring;12:Upgrading;13:Rebooting;14:Connected;15:Connected(Wireless);16:Connected(Migrating);17:Connected(Wireless,Migrating);20:Pending;21:Pending(Wireless);22:Adopting;23:Adopting(Wireless);24:Adopt Failed;25:Adopt Failed(Wireless);26:Managed By Others;27:Managed By Others(Wireless);30:Heartbeat Missed;31:Heartbeat Missed(Wireless);32:Heartbeat Missed(Migrating);33:Heartbeat Missed(Wireless,Migrating);40:Isolated;41:Isolated(Migrating);50:Slice Configuring | [optional] 
**StatusCategory** | Pointer to **int32** | Category of device status,statusCategory should be a value as follows: 0:Disconnected;1:Connected;2:Pending;3:Heartbeat Missed;4:Isolated | [optional] 
**Stp** | Pointer to **int32** |  | [optional] 
**SupportAnomaly** | Pointer to **bool** | Whether the device firmware support intelligent anomaly detection | [optional] 
**SupportCableTest** | Pointer to **bool** |  | [optional] 
**SupportGetOspfNeighborTable** | Pointer to **bool** |  | [optional] 
**SupportHealth** | Pointer to **bool** | Support health | [optional] 
**SupportLocatePort** | Pointer to **bool** | Whether the device supports locating port | [optional] 
**SupportPortAlert** | Pointer to **bool** |  | [optional] 
**SupportPowerAlert** | Pointer to **bool** |  | [optional] 
**SupportStp** | Pointer to **bool** |  | [optional] 
**TagIds** | Pointer to **[]string** |  | [optional] 
**TemplateId** | Pointer to **string** | ID of the template bound to the device | [optional] 
**TemplateName** | Pointer to **string** | Name of the template bound to the device | [optional] 
**TemplateSettings** | Pointer to **[]int32** |  | [optional] 
**TerminalPrefix** | Pointer to **string** | TerminalPrefix represents the device name within the terminal function, designed to prevent terminal command recognition errors when device name contains illegal characters such as &#39;#&#39;. | [optional] 
**TxRate** | Pointer to **int64** | Tx Rate | [optional] 
**Type** | Pointer to **string** | Device type:ap、gateway、switch、olt | [optional] 
**UnknownMulticastRule** | Pointer to **int32** |  | [optional] 
**Uplink** | Pointer to [**OswUplinkVO**](OswUplinkVO.md) |  | [optional] 
**Upload** | Pointer to **int64** |  | [optional] 
**Uptime** | Pointer to **string** |  | [optional] 
**UptimeLong** | Pointer to **int64** |  | [optional] 
**Version** | Pointer to **string** | Simplified version of firmware,for example:2.5.0 | [optional] 

## Methods

### NewESDetailVO

`func NewESDetailVO() *ESDetailVO`

NewESDetailVO instantiates a new ESDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewESDetailVOWithDefaults

`func NewESDetailVOWithDefaults() *ESDetailVO`

NewESDetailVOWithDefaults instantiates a new ESDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *ESDetailVO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ESDetailVO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ESDetailVO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *ESDetailVO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetActive

`func (o *ESDetailVO) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ESDetailVO) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ESDetailVO) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ESDetailVO) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddedInAdvanced

`func (o *ESDetailVO) GetAddedInAdvanced() bool`

GetAddedInAdvanced returns the AddedInAdvanced field if non-nil, zero value otherwise.

### GetAddedInAdvancedOk

`func (o *ESDetailVO) GetAddedInAdvancedOk() (*bool, bool)`

GetAddedInAdvancedOk returns a tuple with the AddedInAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedInAdvanced

`func (o *ESDetailVO) SetAddedInAdvanced(v bool)`

SetAddedInAdvanced sets AddedInAdvanced field to given value.

### HasAddedInAdvanced

`func (o *ESDetailVO) HasAddedInAdvanced() bool`

HasAddedInAdvanced returns a boolean if a field has been set.

### GetBoundDeviceTemplate

`func (o *ESDetailVO) GetBoundDeviceTemplate() bool`

GetBoundDeviceTemplate returns the BoundDeviceTemplate field if non-nil, zero value otherwise.

### GetBoundDeviceTemplateOk

`func (o *ESDetailVO) GetBoundDeviceTemplateOk() (*bool, bool)`

GetBoundDeviceTemplateOk returns a tuple with the BoundDeviceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundDeviceTemplate

`func (o *ESDetailVO) SetBoundDeviceTemplate(v bool)`

SetBoundDeviceTemplate sets BoundDeviceTemplate field to given value.

### HasBoundDeviceTemplate

`func (o *ESDetailVO) HasBoundDeviceTemplate() bool`

HasBoundDeviceTemplate returns a boolean if a field has been set.

### GetBoundSiteTemplate

`func (o *ESDetailVO) GetBoundSiteTemplate() bool`

GetBoundSiteTemplate returns the BoundSiteTemplate field if non-nil, zero value otherwise.

### GetBoundSiteTemplateOk

`func (o *ESDetailVO) GetBoundSiteTemplateOk() (*bool, bool)`

GetBoundSiteTemplateOk returns a tuple with the BoundSiteTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSiteTemplate

`func (o *ESDetailVO) SetBoundSiteTemplate(v bool)`

SetBoundSiteTemplate sets BoundSiteTemplate field to given value.

### HasBoundSiteTemplate

`func (o *ESDetailVO) HasBoundSiteTemplate() bool`

HasBoundSiteTemplate returns a boolean if a field has been set.

### GetCategory

`func (o *ESDetailVO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ESDetailVO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ESDetailVO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ESDetailVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCompatible

`func (o *ESDetailVO) GetCompatible() int32`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *ESDetailVO) GetCompatibleOk() (*int32, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *ESDetailVO) SetCompatible(v int32)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *ESDetailVO) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetCompoundModel

`func (o *ESDetailVO) GetCompoundModel() string`

GetCompoundModel returns the CompoundModel field if non-nil, zero value otherwise.

### GetCompoundModelOk

`func (o *ESDetailVO) GetCompoundModelOk() (*string, bool)`

GetCompoundModelOk returns a tuple with the CompoundModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompoundModel

`func (o *ESDetailVO) SetCompoundModel(v string)`

SetCompoundModel sets CompoundModel field to given value.

### HasCompoundModel

`func (o *ESDetailVO) HasCompoundModel() bool`

HasCompoundModel returns a boolean if a field has been set.

### GetCustomId

`func (o *ESDetailVO) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *ESDetailVO) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *ESDetailVO) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.

### HasCustomId

`func (o *ESDetailVO) HasCustomId() bool`

HasCustomId returns a boolean if a field has been set.

### GetCustomName

`func (o *ESDetailVO) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *ESDetailVO) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *ESDetailVO) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *ESDetailVO) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### GetDescription

`func (o *ESDetailVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ESDetailVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ESDetailVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ESDetailVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevCap

`func (o *ESDetailVO) GetDevCap() OswDevCapVO`

GetDevCap returns the DevCap field if non-nil, zero value otherwise.

### GetDevCapOk

`func (o *ESDetailVO) GetDevCapOk() (*OswDevCapVO, bool)`

GetDevCapOk returns a tuple with the DevCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevCap

`func (o *ESDetailVO) SetDevCap(v OswDevCapVO)`

SetDevCap sets DevCap field to given value.

### HasDevCap

`func (o *ESDetailVO) HasDevCap() bool`

HasDevCap returns a boolean if a field has been set.

### GetDeviceMisc

`func (o *ESDetailVO) GetDeviceMisc() OswDeviceMiscVO`

GetDeviceMisc returns the DeviceMisc field if non-nil, zero value otherwise.

### GetDeviceMiscOk

`func (o *ESDetailVO) GetDeviceMiscOk() (*OswDeviceMiscVO, bool)`

GetDeviceMiscOk returns a tuple with the DeviceMisc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMisc

`func (o *ESDetailVO) SetDeviceMisc(v OswDeviceMiscVO)`

SetDeviceMisc sets DeviceMisc field to given value.

### HasDeviceMisc

`func (o *ESDetailVO) HasDeviceMisc() bool`

HasDeviceMisc returns a boolean if a field has been set.

### GetDeviceSeriesType

`func (o *ESDetailVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *ESDetailVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *ESDetailVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *ESDetailVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetDeviceTemplateAvailable

`func (o *ESDetailVO) GetDeviceTemplateAvailable() bool`

GetDeviceTemplateAvailable returns the DeviceTemplateAvailable field if non-nil, zero value otherwise.

### GetDeviceTemplateAvailableOk

`func (o *ESDetailVO) GetDeviceTemplateAvailableOk() (*bool, bool)`

GetDeviceTemplateAvailableOk returns a tuple with the DeviceTemplateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTemplateAvailable

`func (o *ESDetailVO) SetDeviceTemplateAvailable(v bool)`

SetDeviceTemplateAvailable sets DeviceTemplateAvailable field to given value.

### HasDeviceTemplateAvailable

`func (o *ESDetailVO) HasDeviceTemplateAvailable() bool`

HasDeviceTemplateAvailable returns a boolean if a field has been set.

### GetDisableHwReset

`func (o *ESDetailVO) GetDisableHwReset() bool`

GetDisableHwReset returns the DisableHwReset field if non-nil, zero value otherwise.

### GetDisableHwResetOk

`func (o *ESDetailVO) GetDisableHwResetOk() (*bool, bool)`

GetDisableHwResetOk returns a tuple with the DisableHwReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHwReset

`func (o *ESDetailVO) SetDisableHwReset(v bool)`

SetDisableHwReset sets DisableHwReset field to given value.

### HasDisableHwReset

`func (o *ESDetailVO) HasDisableHwReset() bool`

HasDisableHwReset returns a boolean if a field has been set.

### GetDownlinkList

`func (o *ESDetailVO) GetDownlinkList() []OswDownlinkVO`

GetDownlinkList returns the DownlinkList field if non-nil, zero value otherwise.

### GetDownlinkListOk

`func (o *ESDetailVO) GetDownlinkListOk() (*[]OswDownlinkVO, bool)`

GetDownlinkListOk returns a tuple with the DownlinkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkList

`func (o *ESDetailVO) SetDownlinkList(v []OswDownlinkVO)`

SetDownlinkList sets DownlinkList field to given value.

### HasDownlinkList

`func (o *ESDetailVO) HasDownlinkList() bool`

HasDownlinkList returns a boolean if a field has been set.

### GetDownload

`func (o *ESDetailVO) GetDownload() int64`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ESDetailVO) GetDownloadOk() (*int64, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ESDetailVO) SetDownload(v int64)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *ESDetailVO) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetDueTime

`func (o *ESDetailVO) GetDueTime() int64`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *ESDetailVO) GetDueTimeOk() (*int64, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *ESDetailVO) SetDueTime(v int64)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *ESDetailVO) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### GetDueTimeLeft

`func (o *ESDetailVO) GetDueTimeLeft() int64`

GetDueTimeLeft returns the DueTimeLeft field if non-nil, zero value otherwise.

### GetDueTimeLeftOk

`func (o *ESDetailVO) GetDueTimeLeftOk() (*int64, bool)`

GetDueTimeLeftOk returns a tuple with the DueTimeLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTimeLeft

`func (o *ESDetailVO) SetDueTimeLeft(v int64)`

SetDueTimeLeft sets DueTimeLeft field to given value.

### HasDueTimeLeft

`func (o *ESDetailVO) HasDueTimeLeft() bool`

HasDueTimeLeft returns a boolean if a field has been set.

### GetEcspFirstVersion

`func (o *ESDetailVO) GetEcspFirstVersion() int32`

GetEcspFirstVersion returns the EcspFirstVersion field if non-nil, zero value otherwise.

### GetEcspFirstVersionOk

`func (o *ESDetailVO) GetEcspFirstVersionOk() (*int32, bool)`

GetEcspFirstVersionOk returns a tuple with the EcspFirstVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcspFirstVersion

`func (o *ESDetailVO) SetEcspFirstVersion(v int32)`

SetEcspFirstVersion sets EcspFirstVersion field to given value.

### HasEcspFirstVersion

`func (o *ESDetailVO) HasEcspFirstVersion() bool`

HasEcspFirstVersion returns a boolean if a field has been set.

### GetEos

`func (o *ESDetailVO) GetEos() int64`

GetEos returns the Eos field if non-nil, zero value otherwise.

### GetEosOk

`func (o *ESDetailVO) GetEosOk() (*int64, bool)`

GetEosOk returns a tuple with the Eos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEos

`func (o *ESDetailVO) SetEos(v int64)`

SetEos sets Eos field to given value.

### HasEos

`func (o *ESDetailVO) HasEos() bool`

HasEos returns a boolean if a field has been set.

### GetEost

`func (o *ESDetailVO) GetEost() int64`

GetEost returns the Eost field if non-nil, zero value otherwise.

### GetEostOk

`func (o *ESDetailVO) GetEostOk() (*int64, bool)`

GetEostOk returns a tuple with the Eost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEost

`func (o *ESDetailVO) SetEost(v int64)`

SetEost sets Eost field to given value.

### HasEost

`func (o *ESDetailVO) HasEost() bool`

HasEost returns a boolean if a field has been set.

### GetEs

`func (o *ESDetailVO) GetEs() bool`

GetEs returns the Es field if non-nil, zero value otherwise.

### GetEsOk

`func (o *ESDetailVO) GetEsOk() (*bool, bool)`

GetEsOk returns a tuple with the Es field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEs

`func (o *ESDetailVO) SetEs(v bool)`

SetEs sets Es field to given value.

### HasEs

`func (o *ESDetailVO) HasEs() bool`

HasEs returns a boolean if a field has been set.

### GetFanStatus

`func (o *ESDetailVO) GetFanStatus() int32`

GetFanStatus returns the FanStatus field if non-nil, zero value otherwise.

### GetFanStatusOk

`func (o *ESDetailVO) GetFanStatusOk() (*int32, bool)`

GetFanStatusOk returns a tuple with the FanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanStatus

`func (o *ESDetailVO) SetFanStatus(v int32)`

SetFanStatus sets FanStatus field to given value.

### HasFanStatus

`func (o *ESDetailVO) HasFanStatus() bool`

HasFanStatus returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *ESDetailVO) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *ESDetailVO) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *ESDetailVO) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *ESDetailVO) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetForgetId

`func (o *ESDetailVO) GetForgetId() string`

GetForgetId returns the ForgetId field if non-nil, zero value otherwise.

### GetForgetIdOk

`func (o *ESDetailVO) GetForgetIdOk() (*string, bool)`

GetForgetIdOk returns a tuple with the ForgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgetId

`func (o *ESDetailVO) SetForgetId(v string)`

SetForgetId sets ForgetId field to given value.

### HasForgetId

`func (o *ESDetailVO) HasForgetId() bool`

HasForgetId returns a boolean if a field has been set.

### GetHwVersion

`func (o *ESDetailVO) GetHwVersion() string`

GetHwVersion returns the HwVersion field if non-nil, zero value otherwise.

### GetHwVersionOk

`func (o *ESDetailVO) GetHwVersionOk() (*string, bool)`

GetHwVersionOk returns a tuple with the HwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwVersion

`func (o *ESDetailVO) SetHwVersion(v string)`

SetHwVersion sets HwVersion field to given value.

### HasHwVersion

`func (o *ESDetailVO) HasHwVersion() bool`

HasHwVersion returns a boolean if a field has been set.

### GetId

`func (o *ESDetailVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ESDetailVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ESDetailVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ESDetailVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInWhitelist

`func (o *ESDetailVO) GetInWhitelist() bool`

GetInWhitelist returns the InWhitelist field if non-nil, zero value otherwise.

### GetInWhitelistOk

`func (o *ESDetailVO) GetInWhitelistOk() (*bool, bool)`

GetInWhitelistOk returns a tuple with the InWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInWhitelist

`func (o *ESDetailVO) SetInWhitelist(v bool)`

SetInWhitelist sets InWhitelist field to given value.

### HasInWhitelist

`func (o *ESDetailVO) HasInWhitelist() bool`

HasInWhitelist returns a boolean if a field has been set.

### GetInitialUnbindingLimit

`func (o *ESDetailVO) GetInitialUnbindingLimit() int32`

GetInitialUnbindingLimit returns the InitialUnbindingLimit field if non-nil, zero value otherwise.

### GetInitialUnbindingLimitOk

`func (o *ESDetailVO) GetInitialUnbindingLimitOk() (*int32, bool)`

GetInitialUnbindingLimitOk returns a tuple with the InitialUnbindingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialUnbindingLimit

`func (o *ESDetailVO) SetInitialUnbindingLimit(v int32)`

SetInitialUnbindingLimit sets InitialUnbindingLimit field to given value.

### HasInitialUnbindingLimit

`func (o *ESDetailVO) HasInitialUnbindingLimit() bool`

HasInitialUnbindingLimit returns a boolean if a field has been set.

### GetIp

`func (o *ESDetailVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ESDetailVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ESDetailVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ESDetailVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpSetting

`func (o *ESDetailVO) GetIpSetting() IpSettingVO`

GetIpSetting returns the IpSetting field if non-nil, zero value otherwise.

### GetIpSettingOk

`func (o *ESDetailVO) GetIpSettingOk() (*IpSettingVO, bool)`

GetIpSettingOk returns a tuple with the IpSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSetting

`func (o *ESDetailVO) SetIpSetting(v IpSettingVO)`

SetIpSetting sets IpSetting field to given value.

### HasIpSetting

`func (o *ESDetailVO) HasIpSetting() bool`

HasIpSetting returns a boolean if a field has been set.

### GetJumboEnable

`func (o *ESDetailVO) GetJumboEnable() bool`

GetJumboEnable returns the JumboEnable field if non-nil, zero value otherwise.

### GetJumboEnableOk

`func (o *ESDetailVO) GetJumboEnableOk() (*bool, bool)`

GetJumboEnableOk returns a tuple with the JumboEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumboEnable

`func (o *ESDetailVO) SetJumboEnable(v bool)`

SetJumboEnable sets JumboEnable field to given value.

### HasJumboEnable

`func (o *ESDetailVO) HasJumboEnable() bool`

HasJumboEnable returns a boolean if a field has been set.

### GetLags

`func (o *ESDetailVO) GetLags() []OswLagVO`

GetLags returns the Lags field if non-nil, zero value otherwise.

### GetLagsOk

`func (o *ESDetailVO) GetLagsOk() (*[]OswLagVO, bool)`

GetLagsOk returns a tuple with the Lags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLags

`func (o *ESDetailVO) SetLags(v []OswLagVO)`

SetLags sets Lags field to given value.

### HasLags

`func (o *ESDetailVO) HasLags() bool`

HasLags returns a boolean if a field has been set.

### GetLastSeen

`func (o *ESDetailVO) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ESDetailVO) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ESDetailVO) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ESDetailVO) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLedSetting

`func (o *ESDetailVO) GetLedSetting() int32`

GetLedSetting returns the LedSetting field if non-nil, zero value otherwise.

### GetLedSettingOk

`func (o *ESDetailVO) GetLedSettingOk() (*int32, bool)`

GetLedSettingOk returns a tuple with the LedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedSetting

`func (o *ESDetailVO) SetLedSetting(v int32)`

SetLedSetting sets LedSetting field to given value.

### HasLedSetting

`func (o *ESDetailVO) HasLedSetting() bool`

HasLedSetting returns a boolean if a field has been set.

### GetLicenseId

`func (o *ESDetailVO) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *ESDetailVO) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *ESDetailVO) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *ESDetailVO) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetLicenseStatus

`func (o *ESDetailVO) GetLicenseStatus() int32`

GetLicenseStatus returns the LicenseStatus field if non-nil, zero value otherwise.

### GetLicenseStatusOk

`func (o *ESDetailVO) GetLicenseStatusOk() (*int32, bool)`

GetLicenseStatusOk returns a tuple with the LicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatus

`func (o *ESDetailVO) SetLicenseStatus(v int32)`

SetLicenseStatus sets LicenseStatus field to given value.

### HasLicenseStatus

`func (o *ESDetailVO) HasLicenseStatus() bool`

HasLicenseStatus returns a boolean if a field has been set.

### GetLicenseUnbindingLimit

`func (o *ESDetailVO) GetLicenseUnbindingLimit() int32`

GetLicenseUnbindingLimit returns the LicenseUnbindingLimit field if non-nil, zero value otherwise.

### GetLicenseUnbindingLimitOk

`func (o *ESDetailVO) GetLicenseUnbindingLimitOk() (*int32, bool)`

GetLicenseUnbindingLimitOk returns a tuple with the LicenseUnbindingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUnbindingLimit

`func (o *ESDetailVO) SetLicenseUnbindingLimit(v int32)`

SetLicenseUnbindingLimit sets LicenseUnbindingLimit field to given value.

### HasLicenseUnbindingLimit

`func (o *ESDetailVO) HasLicenseUnbindingLimit() bool`

HasLicenseUnbindingLimit returns a boolean if a field has been set.

### GetLocation

`func (o *ESDetailVO) GetLocation() DeviceLocationDetailVO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ESDetailVO) GetLocationOk() (*DeviceLocationDetailVO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ESDetailVO) SetLocation(v DeviceLocationDetailVO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ESDetailVO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLoop

`func (o *ESDetailVO) GetLoop() string`

GetLoop returns the Loop field if non-nil, zero value otherwise.

### GetLoopOk

`func (o *ESDetailVO) GetLoopOk() (*string, bool)`

GetLoopOk returns a tuple with the Loop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoop

`func (o *ESDetailVO) SetLoop(v string)`

SetLoop sets Loop field to given value.

### HasLoop

`func (o *ESDetailVO) HasLoop() bool`

HasLoop returns a boolean if a field has been set.

### GetLoopbackDetectEnable

`func (o *ESDetailVO) GetLoopbackDetectEnable() bool`

GetLoopbackDetectEnable returns the LoopbackDetectEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectEnableOk

`func (o *ESDetailVO) GetLoopbackDetectEnableOk() (*bool, bool)`

GetLoopbackDetectEnableOk returns a tuple with the LoopbackDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectEnable

`func (o *ESDetailVO) SetLoopbackDetectEnable(v bool)`

SetLoopbackDetectEnable sets LoopbackDetectEnable field to given value.

### HasLoopbackDetectEnable

`func (o *ESDetailVO) HasLoopbackDetectEnable() bool`

HasLoopbackDetectEnable returns a boolean if a field has been set.

### GetLoopbackNum

`func (o *ESDetailVO) GetLoopbackNum() int32`

GetLoopbackNum returns the LoopbackNum field if non-nil, zero value otherwise.

### GetLoopbackNumOk

`func (o *ESDetailVO) GetLoopbackNumOk() (*int32, bool)`

GetLoopbackNumOk returns a tuple with the LoopbackNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackNum

`func (o *ESDetailVO) SetLoopbackNum(v int32)`

SetLoopbackNum sets LoopbackNum field to given value.

### HasLoopbackNum

`func (o *ESDetailVO) HasLoopbackNum() bool`

HasLoopbackNum returns a boolean if a field has been set.

### GetMac

`func (o *ESDetailVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ESDetailVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ESDetailVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ESDetailVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetManagerMark

`func (o *ESDetailVO) GetManagerMark() int32`

GetManagerMark returns the ManagerMark field if non-nil, zero value otherwise.

### GetManagerMarkOk

`func (o *ESDetailVO) GetManagerMarkOk() (*int32, bool)`

GetManagerMarkOk returns a tuple with the ManagerMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerMark

`func (o *ESDetailVO) SetManagerMark(v int32)`

SetManagerMark sets ManagerMark field to given value.

### HasManagerMark

`func (o *ESDetailVO) HasManagerMark() bool`

HasManagerMark returns a boolean if a field has been set.

### GetModel

`func (o *ESDetailVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ESDetailVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ESDetailVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ESDetailVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *ESDetailVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ESDetailVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ESDetailVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *ESDetailVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetMoveSiteId

`func (o *ESDetailVO) GetMoveSiteId() string`

GetMoveSiteId returns the MoveSiteId field if non-nil, zero value otherwise.

### GetMoveSiteIdOk

`func (o *ESDetailVO) GetMoveSiteIdOk() (*string, bool)`

GetMoveSiteIdOk returns a tuple with the MoveSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveSiteId

`func (o *ESDetailVO) SetMoveSiteId(v string)`

SetMoveSiteId sets MoveSiteId field to given value.

### HasMoveSiteId

`func (o *ESDetailVO) HasMoveSiteId() bool`

HasMoveSiteId returns a boolean if a field has been set.

### GetMulticast

`func (o *ESDetailVO) GetMulticast() OswLanMulticastVO`

GetMulticast returns the Multicast field if non-nil, zero value otherwise.

### GetMulticastOk

`func (o *ESDetailVO) GetMulticastOk() (*OswLanMulticastVO, bool)`

GetMulticastOk returns a tuple with the Multicast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticast

`func (o *ESDetailVO) SetMulticast(v OswLanMulticastVO)`

SetMulticast sets Multicast field to given value.

### HasMulticast

`func (o *ESDetailVO) HasMulticast() bool`

HasMulticast returns a boolean if a field has been set.

### GetMvlanBridgeVlan

`func (o *ESDetailVO) GetMvlanBridgeVlan() int32`

GetMvlanBridgeVlan returns the MvlanBridgeVlan field if non-nil, zero value otherwise.

### GetMvlanBridgeVlanOk

`func (o *ESDetailVO) GetMvlanBridgeVlanOk() (*int32, bool)`

GetMvlanBridgeVlanOk returns a tuple with the MvlanBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMvlanBridgeVlan

`func (o *ESDetailVO) SetMvlanBridgeVlan(v int32)`

SetMvlanBridgeVlan sets MvlanBridgeVlan field to given value.

### HasMvlanBridgeVlan

`func (o *ESDetailVO) HasMvlanBridgeVlan() bool`

HasMvlanBridgeVlan returns a boolean if a field has been set.

### GetMvlanNetworkId

`func (o *ESDetailVO) GetMvlanNetworkId() string`

GetMvlanNetworkId returns the MvlanNetworkId field if non-nil, zero value otherwise.

### GetMvlanNetworkIdOk

`func (o *ESDetailVO) GetMvlanNetworkIdOk() (*string, bool)`

GetMvlanNetworkIdOk returns a tuple with the MvlanNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMvlanNetworkId

`func (o *ESDetailVO) SetMvlanNetworkId(v string)`

SetMvlanNetworkId sets MvlanNetworkId field to given value.

### HasMvlanNetworkId

`func (o *ESDetailVO) HasMvlanNetworkId() bool`

HasMvlanNetworkId returns a boolean if a field has been set.

### GetName

`func (o *ESDetailVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ESDetailVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ESDetailVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ESDetailVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedUpgrade

`func (o *ESDetailVO) GetNeedUpgrade() bool`

GetNeedUpgrade returns the NeedUpgrade field if non-nil, zero value otherwise.

### GetNeedUpgradeOk

`func (o *ESDetailVO) GetNeedUpgradeOk() (*bool, bool)`

GetNeedUpgradeOk returns a tuple with the NeedUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedUpgrade

`func (o *ESDetailVO) SetNeedUpgrade(v bool)`

SetNeedUpgrade sets NeedUpgrade field to given value.

### HasNeedUpgrade

`func (o *ESDetailVO) HasNeedUpgrade() bool`

HasNeedUpgrade returns a boolean if a field has been set.

### GetNetworkNotify

`func (o *ESDetailVO) GetNetworkNotify() bool`

GetNetworkNotify returns the NetworkNotify field if non-nil, zero value otherwise.

### GetNetworkNotifyOk

`func (o *ESDetailVO) GetNetworkNotifyOk() (*bool, bool)`

GetNetworkNotifyOk returns a tuple with the NetworkNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkNotify

`func (o *ESDetailVO) SetNetworkNotify(v bool)`

SetNetworkNotify sets NetworkNotify field to given value.

### HasNetworkNotify

`func (o *ESDetailVO) HasNetworkNotify() bool`

HasNetworkNotify returns a boolean if a field has been set.

### GetOmadacId

`func (o *ESDetailVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *ESDetailVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *ESDetailVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *ESDetailVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetPoeRemain

`func (o *ESDetailVO) GetPoeRemain() float64`

GetPoeRemain returns the PoeRemain field if non-nil, zero value otherwise.

### GetPoeRemainOk

`func (o *ESDetailVO) GetPoeRemainOk() (*float64, bool)`

GetPoeRemainOk returns a tuple with the PoeRemain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeRemain

`func (o *ESDetailVO) SetPoeRemain(v float64)`

SetPoeRemain sets PoeRemain field to given value.

### HasPoeRemain

`func (o *ESDetailVO) HasPoeRemain() bool`

HasPoeRemain returns a boolean if a field has been set.

### GetPoeRemainPercent

`func (o *ESDetailVO) GetPoeRemainPercent() float64`

GetPoeRemainPercent returns the PoeRemainPercent field if non-nil, zero value otherwise.

### GetPoeRemainPercentOk

`func (o *ESDetailVO) GetPoeRemainPercentOk() (*float64, bool)`

GetPoeRemainPercentOk returns a tuple with the PoeRemainPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeRemainPercent

`func (o *ESDetailVO) SetPoeRemainPercent(v float64)`

SetPoeRemainPercent sets PoeRemainPercent field to given value.

### HasPoeRemainPercent

`func (o *ESDetailVO) HasPoeRemainPercent() bool`

HasPoeRemainPercent returns a boolean if a field has been set.

### GetPoeTotalPower

`func (o *ESDetailVO) GetPoeTotalPower() float64`

GetPoeTotalPower returns the PoeTotalPower field if non-nil, zero value otherwise.

### GetPoeTotalPowerOk

`func (o *ESDetailVO) GetPoeTotalPowerOk() (*float64, bool)`

GetPoeTotalPowerOk returns a tuple with the PoeTotalPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeTotalPower

`func (o *ESDetailVO) SetPoeTotalPower(v float64)`

SetPoeTotalPower sets PoeTotalPower field to given value.

### HasPoeTotalPower

`func (o *ESDetailVO) HasPoeTotalPower() bool`

HasPoeTotalPower returns a boolean if a field has been set.

### GetPorts

`func (o *ESDetailVO) GetPorts() []OswPortVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ESDetailVO) GetPortsOk() (*[]OswPortVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ESDetailVO) SetPorts(v []OswPortVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ESDetailVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPowerAlertEnable

`func (o *ESDetailVO) GetPowerAlertEnable() bool`

GetPowerAlertEnable returns the PowerAlertEnable field if non-nil, zero value otherwise.

### GetPowerAlertEnableOk

`func (o *ESDetailVO) GetPowerAlertEnableOk() (*bool, bool)`

GetPowerAlertEnableOk returns a tuple with the PowerAlertEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerAlertEnable

`func (o *ESDetailVO) SetPowerAlertEnable(v bool)`

SetPowerAlertEnable sets PowerAlertEnable field to given value.

### HasPowerAlertEnable

`func (o *ESDetailVO) HasPowerAlertEnable() bool`

HasPowerAlertEnable returns a boolean if a field has been set.

### GetPowerStatusList

`func (o *ESDetailVO) GetPowerStatusList() []OswPortAlertStatusVO`

GetPowerStatusList returns the PowerStatusList field if non-nil, zero value otherwise.

### GetPowerStatusListOk

`func (o *ESDetailVO) GetPowerStatusListOk() (*[]OswPortAlertStatusVO, bool)`

GetPowerStatusListOk returns a tuple with the PowerStatusList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerStatusList

`func (o *ESDetailVO) SetPowerStatusList(v []OswPortAlertStatusVO)`

SetPowerStatusList sets PowerStatusList field to given value.

### HasPowerStatusList

`func (o *ESDetailVO) HasPowerStatusList() bool`

HasPowerStatusList returns a boolean if a field has been set.

### GetPriority

`func (o *ESDetailVO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ESDetailVO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ESDetailVO) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ESDetailVO) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPublicIp

`func (o *ESDetailVO) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *ESDetailVO) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *ESDetailVO) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *ESDetailVO) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRemember

`func (o *ESDetailVO) GetRemember() bool`

GetRemember returns the Remember field if non-nil, zero value otherwise.

### GetRememberOk

`func (o *ESDetailVO) GetRememberOk() (*bool, bool)`

GetRememberOk returns a tuple with the Remember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemember

`func (o *ESDetailVO) SetRemember(v bool)`

SetRemember sets Remember field to given value.

### HasRemember

`func (o *ESDetailVO) HasRemember() bool`

HasRemember returns a boolean if a field has been set.

### GetRememberDevice

`func (o *ESDetailVO) GetRememberDevice() int32`

GetRememberDevice returns the RememberDevice field if non-nil, zero value otherwise.

### GetRememberDeviceOk

`func (o *ESDetailVO) GetRememberDeviceOk() (*int32, bool)`

GetRememberDeviceOk returns a tuple with the RememberDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberDevice

`func (o *ESDetailVO) SetRememberDevice(v int32)`

SetRememberDevice sets RememberDevice field to given value.

### HasRememberDevice

`func (o *ESDetailVO) HasRememberDevice() bool`

HasRememberDevice returns a boolean if a field has been set.

### GetResource

`func (o *ESDetailVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ESDetailVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ESDetailVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ESDetailVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetRxRate

`func (o *ESDetailVO) GetRxRate() int64`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *ESDetailVO) GetRxRateOk() (*int64, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *ESDetailVO) SetRxRate(v int64)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *ESDetailVO) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetShowModel

`func (o *ESDetailVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *ESDetailVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *ESDetailVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *ESDetailVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetSiteName

`func (o *ESDetailVO) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *ESDetailVO) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *ESDetailVO) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *ESDetailVO) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSiteTemplateId

`func (o *ESDetailVO) GetSiteTemplateId() string`

GetSiteTemplateId returns the SiteTemplateId field if non-nil, zero value otherwise.

### GetSiteTemplateIdOk

`func (o *ESDetailVO) GetSiteTemplateIdOk() (*string, bool)`

GetSiteTemplateIdOk returns a tuple with the SiteTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTemplateId

`func (o *ESDetailVO) SetSiteTemplateId(v string)`

SetSiteTemplateId sets SiteTemplateId field to given value.

### HasSiteTemplateId

`func (o *ESDetailVO) HasSiteTemplateId() bool`

HasSiteTemplateId returns a boolean if a field has been set.

### GetSiteTemplateName

`func (o *ESDetailVO) GetSiteTemplateName() string`

GetSiteTemplateName returns the SiteTemplateName field if non-nil, zero value otherwise.

### GetSiteTemplateNameOk

`func (o *ESDetailVO) GetSiteTemplateNameOk() (*string, bool)`

GetSiteTemplateNameOk returns a tuple with the SiteTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTemplateName

`func (o *ESDetailVO) SetSiteTemplateName(v string)`

SetSiteTemplateName sets SiteTemplateName field to given value.

### HasSiteTemplateName

`func (o *ESDetailVO) HasSiteTemplateName() bool`

HasSiteTemplateName returns a boolean if a field has been set.

### GetSn

`func (o *ESDetailVO) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *ESDetailVO) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *ESDetailVO) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *ESDetailVO) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetSpecialModel

`func (o *ESDetailVO) GetSpecialModel() string`

GetSpecialModel returns the SpecialModel field if non-nil, zero value otherwise.

### GetSpecialModelOk

`func (o *ESDetailVO) GetSpecialModelOk() (*string, bool)`

GetSpecialModelOk returns a tuple with the SpecialModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialModel

`func (o *ESDetailVO) SetSpecialModel(v string)`

SetSpecialModel sets SpecialModel field to given value.

### HasSpecialModel

`func (o *ESDetailVO) HasSpecialModel() bool`

HasSpecialModel returns a boolean if a field has been set.

### GetSpeeds

`func (o *ESDetailVO) GetSpeeds() []int32`

GetSpeeds returns the Speeds field if non-nil, zero value otherwise.

### GetSpeedsOk

`func (o *ESDetailVO) GetSpeedsOk() (*[]int32, bool)`

GetSpeedsOk returns a tuple with the Speeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeeds

`func (o *ESDetailVO) SetSpeeds(v []int32)`

SetSpeeds sets Speeds field to given value.

### HasSpeeds

`func (o *ESDetailVO) HasSpeeds() bool`

HasSpeeds returns a boolean if a field has been set.

### GetStatus

`func (o *ESDetailVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ESDetailVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ESDetailVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ESDetailVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *ESDetailVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *ESDetailVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *ESDetailVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *ESDetailVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetStp

`func (o *ESDetailVO) GetStp() int32`

GetStp returns the Stp field if non-nil, zero value otherwise.

### GetStpOk

`func (o *ESDetailVO) GetStpOk() (*int32, bool)`

GetStpOk returns a tuple with the Stp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStp

`func (o *ESDetailVO) SetStp(v int32)`

SetStp sets Stp field to given value.

### HasStp

`func (o *ESDetailVO) HasStp() bool`

HasStp returns a boolean if a field has been set.

### GetSupportAnomaly

`func (o *ESDetailVO) GetSupportAnomaly() bool`

GetSupportAnomaly returns the SupportAnomaly field if non-nil, zero value otherwise.

### GetSupportAnomalyOk

`func (o *ESDetailVO) GetSupportAnomalyOk() (*bool, bool)`

GetSupportAnomalyOk returns a tuple with the SupportAnomaly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAnomaly

`func (o *ESDetailVO) SetSupportAnomaly(v bool)`

SetSupportAnomaly sets SupportAnomaly field to given value.

### HasSupportAnomaly

`func (o *ESDetailVO) HasSupportAnomaly() bool`

HasSupportAnomaly returns a boolean if a field has been set.

### GetSupportCableTest

`func (o *ESDetailVO) GetSupportCableTest() bool`

GetSupportCableTest returns the SupportCableTest field if non-nil, zero value otherwise.

### GetSupportCableTestOk

`func (o *ESDetailVO) GetSupportCableTestOk() (*bool, bool)`

GetSupportCableTestOk returns a tuple with the SupportCableTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCableTest

`func (o *ESDetailVO) SetSupportCableTest(v bool)`

SetSupportCableTest sets SupportCableTest field to given value.

### HasSupportCableTest

`func (o *ESDetailVO) HasSupportCableTest() bool`

HasSupportCableTest returns a boolean if a field has been set.

### GetSupportGetOspfNeighborTable

`func (o *ESDetailVO) GetSupportGetOspfNeighborTable() bool`

GetSupportGetOspfNeighborTable returns the SupportGetOspfNeighborTable field if non-nil, zero value otherwise.

### GetSupportGetOspfNeighborTableOk

`func (o *ESDetailVO) GetSupportGetOspfNeighborTableOk() (*bool, bool)`

GetSupportGetOspfNeighborTableOk returns a tuple with the SupportGetOspfNeighborTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportGetOspfNeighborTable

`func (o *ESDetailVO) SetSupportGetOspfNeighborTable(v bool)`

SetSupportGetOspfNeighborTable sets SupportGetOspfNeighborTable field to given value.

### HasSupportGetOspfNeighborTable

`func (o *ESDetailVO) HasSupportGetOspfNeighborTable() bool`

HasSupportGetOspfNeighborTable returns a boolean if a field has been set.

### GetSupportHealth

`func (o *ESDetailVO) GetSupportHealth() bool`

GetSupportHealth returns the SupportHealth field if non-nil, zero value otherwise.

### GetSupportHealthOk

`func (o *ESDetailVO) GetSupportHealthOk() (*bool, bool)`

GetSupportHealthOk returns a tuple with the SupportHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportHealth

`func (o *ESDetailVO) SetSupportHealth(v bool)`

SetSupportHealth sets SupportHealth field to given value.

### HasSupportHealth

`func (o *ESDetailVO) HasSupportHealth() bool`

HasSupportHealth returns a boolean if a field has been set.

### GetSupportLocatePort

`func (o *ESDetailVO) GetSupportLocatePort() bool`

GetSupportLocatePort returns the SupportLocatePort field if non-nil, zero value otherwise.

### GetSupportLocatePortOk

`func (o *ESDetailVO) GetSupportLocatePortOk() (*bool, bool)`

GetSupportLocatePortOk returns a tuple with the SupportLocatePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLocatePort

`func (o *ESDetailVO) SetSupportLocatePort(v bool)`

SetSupportLocatePort sets SupportLocatePort field to given value.

### HasSupportLocatePort

`func (o *ESDetailVO) HasSupportLocatePort() bool`

HasSupportLocatePort returns a boolean if a field has been set.

### GetSupportPortAlert

`func (o *ESDetailVO) GetSupportPortAlert() bool`

GetSupportPortAlert returns the SupportPortAlert field if non-nil, zero value otherwise.

### GetSupportPortAlertOk

`func (o *ESDetailVO) GetSupportPortAlertOk() (*bool, bool)`

GetSupportPortAlertOk returns a tuple with the SupportPortAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPortAlert

`func (o *ESDetailVO) SetSupportPortAlert(v bool)`

SetSupportPortAlert sets SupportPortAlert field to given value.

### HasSupportPortAlert

`func (o *ESDetailVO) HasSupportPortAlert() bool`

HasSupportPortAlert returns a boolean if a field has been set.

### GetSupportPowerAlert

`func (o *ESDetailVO) GetSupportPowerAlert() bool`

GetSupportPowerAlert returns the SupportPowerAlert field if non-nil, zero value otherwise.

### GetSupportPowerAlertOk

`func (o *ESDetailVO) GetSupportPowerAlertOk() (*bool, bool)`

GetSupportPowerAlertOk returns a tuple with the SupportPowerAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPowerAlert

`func (o *ESDetailVO) SetSupportPowerAlert(v bool)`

SetSupportPowerAlert sets SupportPowerAlert field to given value.

### HasSupportPowerAlert

`func (o *ESDetailVO) HasSupportPowerAlert() bool`

HasSupportPowerAlert returns a boolean if a field has been set.

### GetSupportStp

`func (o *ESDetailVO) GetSupportStp() bool`

GetSupportStp returns the SupportStp field if non-nil, zero value otherwise.

### GetSupportStpOk

`func (o *ESDetailVO) GetSupportStpOk() (*bool, bool)`

GetSupportStpOk returns a tuple with the SupportStp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportStp

`func (o *ESDetailVO) SetSupportStp(v bool)`

SetSupportStp sets SupportStp field to given value.

### HasSupportStp

`func (o *ESDetailVO) HasSupportStp() bool`

HasSupportStp returns a boolean if a field has been set.

### GetTagIds

`func (o *ESDetailVO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ESDetailVO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ESDetailVO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ESDetailVO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTemplateId

`func (o *ESDetailVO) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ESDetailVO) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ESDetailVO) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ESDetailVO) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateName

`func (o *ESDetailVO) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *ESDetailVO) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *ESDetailVO) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *ESDetailVO) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetTemplateSettings

`func (o *ESDetailVO) GetTemplateSettings() []int32`

GetTemplateSettings returns the TemplateSettings field if non-nil, zero value otherwise.

### GetTemplateSettingsOk

`func (o *ESDetailVO) GetTemplateSettingsOk() (*[]int32, bool)`

GetTemplateSettingsOk returns a tuple with the TemplateSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSettings

`func (o *ESDetailVO) SetTemplateSettings(v []int32)`

SetTemplateSettings sets TemplateSettings field to given value.

### HasTemplateSettings

`func (o *ESDetailVO) HasTemplateSettings() bool`

HasTemplateSettings returns a boolean if a field has been set.

### GetTerminalPrefix

`func (o *ESDetailVO) GetTerminalPrefix() string`

GetTerminalPrefix returns the TerminalPrefix field if non-nil, zero value otherwise.

### GetTerminalPrefixOk

`func (o *ESDetailVO) GetTerminalPrefixOk() (*string, bool)`

GetTerminalPrefixOk returns a tuple with the TerminalPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalPrefix

`func (o *ESDetailVO) SetTerminalPrefix(v string)`

SetTerminalPrefix sets TerminalPrefix field to given value.

### HasTerminalPrefix

`func (o *ESDetailVO) HasTerminalPrefix() bool`

HasTerminalPrefix returns a boolean if a field has been set.

### GetTxRate

`func (o *ESDetailVO) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *ESDetailVO) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *ESDetailVO) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *ESDetailVO) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.

### GetType

`func (o *ESDetailVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ESDetailVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ESDetailVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ESDetailVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnknownMulticastRule

`func (o *ESDetailVO) GetUnknownMulticastRule() int32`

GetUnknownMulticastRule returns the UnknownMulticastRule field if non-nil, zero value otherwise.

### GetUnknownMulticastRuleOk

`func (o *ESDetailVO) GetUnknownMulticastRuleOk() (*int32, bool)`

GetUnknownMulticastRuleOk returns a tuple with the UnknownMulticastRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownMulticastRule

`func (o *ESDetailVO) SetUnknownMulticastRule(v int32)`

SetUnknownMulticastRule sets UnknownMulticastRule field to given value.

### HasUnknownMulticastRule

`func (o *ESDetailVO) HasUnknownMulticastRule() bool`

HasUnknownMulticastRule returns a boolean if a field has been set.

### GetUplink

`func (o *ESDetailVO) GetUplink() OswUplinkVO`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *ESDetailVO) GetUplinkOk() (*OswUplinkVO, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *ESDetailVO) SetUplink(v OswUplinkVO)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *ESDetailVO) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetUpload

`func (o *ESDetailVO) GetUpload() int64`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *ESDetailVO) GetUploadOk() (*int64, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *ESDetailVO) SetUpload(v int64)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *ESDetailVO) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetUptime

`func (o *ESDetailVO) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *ESDetailVO) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *ESDetailVO) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *ESDetailVO) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetUptimeLong

`func (o *ESDetailVO) GetUptimeLong() int64`

GetUptimeLong returns the UptimeLong field if non-nil, zero value otherwise.

### GetUptimeLongOk

`func (o *ESDetailVO) GetUptimeLongOk() (*int64, bool)`

GetUptimeLongOk returns a tuple with the UptimeLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeLong

`func (o *ESDetailVO) SetUptimeLong(v int64)`

SetUptimeLong sets UptimeLong field to given value.

### HasUptimeLong

`func (o *ESDetailVO) HasUptimeLong() bool`

HasUptimeLong returns a boolean if a field has been set.

### GetVersion

`func (o *ESDetailVO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ESDetailVO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ESDetailVO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ESDetailVO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


