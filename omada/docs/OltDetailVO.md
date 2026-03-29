# OltDetailVO

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
**Components** | Pointer to **map[string]string** | Component information of the Omada Controller sent to the OLT | [optional] 
**CompoundModel** | Pointer to **string** | Model complex used in the backend.Ap：model+(country)+modelVersion,  EAP225(EU) v3.0 Ap: specialModel+modelVersion, EAP225-Outdoor-1a20a950b8d950e8 v1.0  Gateway/Switch：model+modelVersion, Osg v3.0 | [optional] 
**CpuUtil** | Pointer to **int32** | Cpu utilization | [optional] 
**CustomId** | Pointer to **string** | Customer ID | [optional] 
**CustomName** | Pointer to **string** | Customer name | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DevCap** | Pointer to [**OltDevCapVO**](OltDevCapVO.md) |  | [optional] 
**DeviceSeriesType** | Pointer to **int32** | Device series type.DeviceSeriesType should be a value as follows: 0:advanced;1:pro | [optional] 
**DeviceTemplateAvailable** | Pointer to **bool** | Whether there is an available device template for the device; it is false if the model is not supported or the site template has not created the corresponding device template. | [optional] 
**DisableHwReset** | Pointer to **bool** |  | [optional] 
**Down** | Pointer to **int64** | Total real-time down-link traffic(byte) | [optional] 
**DownlinkList** | Pointer to [**[]OltDetailDownlinkVO**](OltDetailDownlinkVO.md) | Down-link list | [optional] 
**DueTime** | Pointer to **int64** | Expire timestamp of license(cloud base exclusive) | [optional] 
**DueTimeLeft** | Pointer to **int64** | Milliseconds from the current moment to the expiration time(cloud base exclusive) | [optional] 
**EcspFirstVersion** | Pointer to **int32** | Ecsp first version | [optional] 
**Eos** | Pointer to **int64** | End of support time of device(CBC exclusive) | [optional] 
**Eost** | Pointer to **int64** | End of service time of device(CBC exclusive) | [optional] 
**Es** | Pointer to **bool** | Whether the device is Agile Series Switch | [optional] 
**FirmwareVersion** | Pointer to **string** | Version of firmware,for example:2.5.0 Build 20190118 Rel. 64821 | [optional] 
**ForgetId** | Pointer to **string** | Forget ID of device | [optional] 
**HwVersion** | Pointer to **string** | Version of hardware,for example 1.0 | [optional] 
**InWhitelist** | Pointer to **bool** | Whether the device is in white list | [optional] 
**InitialUnbindingLimit** | Pointer to **int32** | Initial unbind count for license(cloud base exclusive) | [optional] 
**Ip** | Pointer to **string** | Ip address | [optional] 
**Ipv6List** | Pointer to **[]string** | Ipv6 address List | [optional] 
**LastSeen** | Pointer to **int64** | Last online timestamp (ms) | [optional] 
**LicenseId** | Pointer to **string** | License key on detail page of device(cloud base exclusive) | [optional] 
**LicenseStatus** | Pointer to **int32** | License status(cloud base exclusive).LicenseStatus should be a value as follows: 0:unActive 1:Unbind 2:Expired 3:active | [optional] 
**LicenseUnbindingLimit** | Pointer to **int32** | Remaining unbind count for license on detail Page of device(cloud base exclusive) | [optional] 
**Location** | Pointer to [**DeviceLocationDetailVO**](DeviceLocationDetailVO.md) |  | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**MemUtil** | Pointer to **int32** | Memory utilization | [optional] 
**Model** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**MoveSiteId** | Pointer to **string** | Record that the device is in a moveSite operation; if it is null, then it is not in the moveSite operation. | [optional] 
**Name** | Pointer to **string** | Device name,default value is the mac address of device | [optional] 
**OmadacId** | Pointer to **string** | OmadacId of the device | [optional] 
**OnuCount** | Pointer to **int32** | Number of ONUs managed by the device | [optional] 
**PublicIp** | Pointer to **string** | Public ip address | [optional] 
**Remember** | Pointer to **bool** | Whether to remember the device(deprecated) | [optional] 
**RememberDevice** | Pointer to **int32** | Whether to remember the device.RememberDevice should be a value as follows: 0:off, 1:on, 2: follow site | [optional] 
**Resource** | Pointer to **int32** | Data source.Resource should be a value as follows: 0:new created;1:from template;2:override | [optional] 
**ShowModel** | Pointer to **string** | Model complex shown in the front end.Ap：model+(country)+modelVersion,EAP225(EU) v3.0  Gateway/Switch：model+modelVersion,Osg v3.0 | [optional] 
**SiteName** | Pointer to **string** | Site name of the device | [optional] 
**SiteTemplateId** | Pointer to **string** | Template ID bound to the site | [optional] 
**SiteTemplateName** | Pointer to **string** | Template name bound to the site | [optional] 
**Sn** | Pointer to **string** | SN code of device | [optional] 
**SpecialModel** | Pointer to **string** | Special device model,for example:EAP225-Outdoor-1a20a950b8d950e8 | [optional] 
**Speeds** | Pointer to **[]int32** | Supported rate list for all ports. Speeds should be a value as follows: 0:auto; 1:10M; 2:100M; 3:1000M; 4:2.5G; 5:10G; 6:5G; 7:25G; 8:100G; 9:40G; -1:error; no value:all rate supported | [optional] 
**Status** | Pointer to **int32** | Status of device,status should be a value as follows: 0:Disconnected;1:Disconnected(Migrating);10:Provisioning;11:Configuring;12:Upgrading;13:Rebooting;14:Connected;15:Connected(Wireless);16:Connected(Migrating);17:Connected(Wireless,Migrating);20:Pending;21:Pending(Wireless);22:Adopting;23:Adopting(Wireless);24:Adopt Failed;25:Adopt Failed(Wireless);26:Managed By Others;27:Managed By Others(Wireless);30:Heartbeat Missed;31:Heartbeat Missed(Wireless);32:Heartbeat Missed(Migrating);33:Heartbeat Missed(Wireless,Migrating);40:Isolated;41:Isolated(Migrating);50:Slice Configuring | [optional] 
**StatusCategory** | Pointer to **int32** | Category of device status,statusCategory should be a value as follows: 0:Disconnected;1:Connected;2:Pending;3:Heartbeat Missed;4:Isolated | [optional] 
**SupportAnomaly** | Pointer to **bool** | Whether the device firmware support intelligent anomaly detection | [optional] 
**SupportLocatePort** | Pointer to **bool** | Whether the device supports locating port | [optional] 
**TagIds** | Pointer to **[]string** | Bound tag ID | [optional] 
**TemplateId** | Pointer to **string** | ID of the template bound to the device | [optional] 
**TemplateName** | Pointer to **string** | Name of the template bound to the device | [optional] 
**Type** | Pointer to **string** | Device type:ap、gateway、switch、olt | [optional] 
**Up** | Pointer to **int64** | Total real-time up-link traffic(byte) | [optional] 
**Uplink** | Pointer to [**OltUplinkVO**](OltUplinkVO.md) |  | [optional] 
**Uptime** | Pointer to **string** | Runtime duration, for example:0 days 00:44:09 | [optional] 
**UptimeLong** | Pointer to **int64** | Runtime duration, in seconds (s). | [optional] 
**Version** | Pointer to **string** | Simplified version of firmware,for example:2.5.0 | [optional] 

## Methods

### NewOltDetailVO

`func NewOltDetailVO() *OltDetailVO`

NewOltDetailVO instantiates a new OltDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOltDetailVOWithDefaults

`func NewOltDetailVOWithDefaults() *OltDetailVO`

NewOltDetailVOWithDefaults instantiates a new OltDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *OltDetailVO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *OltDetailVO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *OltDetailVO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *OltDetailVO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetActive

`func (o *OltDetailVO) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OltDetailVO) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OltDetailVO) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *OltDetailVO) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddedInAdvanced

`func (o *OltDetailVO) GetAddedInAdvanced() bool`

GetAddedInAdvanced returns the AddedInAdvanced field if non-nil, zero value otherwise.

### GetAddedInAdvancedOk

`func (o *OltDetailVO) GetAddedInAdvancedOk() (*bool, bool)`

GetAddedInAdvancedOk returns a tuple with the AddedInAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedInAdvanced

`func (o *OltDetailVO) SetAddedInAdvanced(v bool)`

SetAddedInAdvanced sets AddedInAdvanced field to given value.

### HasAddedInAdvanced

`func (o *OltDetailVO) HasAddedInAdvanced() bool`

HasAddedInAdvanced returns a boolean if a field has been set.

### GetBoundDeviceTemplate

`func (o *OltDetailVO) GetBoundDeviceTemplate() bool`

GetBoundDeviceTemplate returns the BoundDeviceTemplate field if non-nil, zero value otherwise.

### GetBoundDeviceTemplateOk

`func (o *OltDetailVO) GetBoundDeviceTemplateOk() (*bool, bool)`

GetBoundDeviceTemplateOk returns a tuple with the BoundDeviceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundDeviceTemplate

`func (o *OltDetailVO) SetBoundDeviceTemplate(v bool)`

SetBoundDeviceTemplate sets BoundDeviceTemplate field to given value.

### HasBoundDeviceTemplate

`func (o *OltDetailVO) HasBoundDeviceTemplate() bool`

HasBoundDeviceTemplate returns a boolean if a field has been set.

### GetBoundSiteTemplate

`func (o *OltDetailVO) GetBoundSiteTemplate() bool`

GetBoundSiteTemplate returns the BoundSiteTemplate field if non-nil, zero value otherwise.

### GetBoundSiteTemplateOk

`func (o *OltDetailVO) GetBoundSiteTemplateOk() (*bool, bool)`

GetBoundSiteTemplateOk returns a tuple with the BoundSiteTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSiteTemplate

`func (o *OltDetailVO) SetBoundSiteTemplate(v bool)`

SetBoundSiteTemplate sets BoundSiteTemplate field to given value.

### HasBoundSiteTemplate

`func (o *OltDetailVO) HasBoundSiteTemplate() bool`

HasBoundSiteTemplate returns a boolean if a field has been set.

### GetCategory

`func (o *OltDetailVO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *OltDetailVO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *OltDetailVO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *OltDetailVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCompatible

`func (o *OltDetailVO) GetCompatible() int32`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *OltDetailVO) GetCompatibleOk() (*int32, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *OltDetailVO) SetCompatible(v int32)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *OltDetailVO) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetComponents

`func (o *OltDetailVO) GetComponents() map[string]string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *OltDetailVO) GetComponentsOk() (*map[string]string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *OltDetailVO) SetComponents(v map[string]string)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *OltDetailVO) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetCompoundModel

`func (o *OltDetailVO) GetCompoundModel() string`

GetCompoundModel returns the CompoundModel field if non-nil, zero value otherwise.

### GetCompoundModelOk

`func (o *OltDetailVO) GetCompoundModelOk() (*string, bool)`

GetCompoundModelOk returns a tuple with the CompoundModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompoundModel

`func (o *OltDetailVO) SetCompoundModel(v string)`

SetCompoundModel sets CompoundModel field to given value.

### HasCompoundModel

`func (o *OltDetailVO) HasCompoundModel() bool`

HasCompoundModel returns a boolean if a field has been set.

### GetCpuUtil

`func (o *OltDetailVO) GetCpuUtil() int32`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *OltDetailVO) GetCpuUtilOk() (*int32, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *OltDetailVO) SetCpuUtil(v int32)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *OltDetailVO) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetCustomId

`func (o *OltDetailVO) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *OltDetailVO) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *OltDetailVO) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.

### HasCustomId

`func (o *OltDetailVO) HasCustomId() bool`

HasCustomId returns a boolean if a field has been set.

### GetCustomName

`func (o *OltDetailVO) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *OltDetailVO) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *OltDetailVO) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *OltDetailVO) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### GetDescription

`func (o *OltDetailVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OltDetailVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OltDetailVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OltDetailVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevCap

`func (o *OltDetailVO) GetDevCap() OltDevCapVO`

GetDevCap returns the DevCap field if non-nil, zero value otherwise.

### GetDevCapOk

`func (o *OltDetailVO) GetDevCapOk() (*OltDevCapVO, bool)`

GetDevCapOk returns a tuple with the DevCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevCap

`func (o *OltDetailVO) SetDevCap(v OltDevCapVO)`

SetDevCap sets DevCap field to given value.

### HasDevCap

`func (o *OltDetailVO) HasDevCap() bool`

HasDevCap returns a boolean if a field has been set.

### GetDeviceSeriesType

`func (o *OltDetailVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *OltDetailVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *OltDetailVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *OltDetailVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetDeviceTemplateAvailable

`func (o *OltDetailVO) GetDeviceTemplateAvailable() bool`

GetDeviceTemplateAvailable returns the DeviceTemplateAvailable field if non-nil, zero value otherwise.

### GetDeviceTemplateAvailableOk

`func (o *OltDetailVO) GetDeviceTemplateAvailableOk() (*bool, bool)`

GetDeviceTemplateAvailableOk returns a tuple with the DeviceTemplateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTemplateAvailable

`func (o *OltDetailVO) SetDeviceTemplateAvailable(v bool)`

SetDeviceTemplateAvailable sets DeviceTemplateAvailable field to given value.

### HasDeviceTemplateAvailable

`func (o *OltDetailVO) HasDeviceTemplateAvailable() bool`

HasDeviceTemplateAvailable returns a boolean if a field has been set.

### GetDisableHwReset

`func (o *OltDetailVO) GetDisableHwReset() bool`

GetDisableHwReset returns the DisableHwReset field if non-nil, zero value otherwise.

### GetDisableHwResetOk

`func (o *OltDetailVO) GetDisableHwResetOk() (*bool, bool)`

GetDisableHwResetOk returns a tuple with the DisableHwReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHwReset

`func (o *OltDetailVO) SetDisableHwReset(v bool)`

SetDisableHwReset sets DisableHwReset field to given value.

### HasDisableHwReset

`func (o *OltDetailVO) HasDisableHwReset() bool`

HasDisableHwReset returns a boolean if a field has been set.

### GetDown

`func (o *OltDetailVO) GetDown() int64`

GetDown returns the Down field if non-nil, zero value otherwise.

### GetDownOk

`func (o *OltDetailVO) GetDownOk() (*int64, bool)`

GetDownOk returns a tuple with the Down field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDown

`func (o *OltDetailVO) SetDown(v int64)`

SetDown sets Down field to given value.

### HasDown

`func (o *OltDetailVO) HasDown() bool`

HasDown returns a boolean if a field has been set.

### GetDownlinkList

`func (o *OltDetailVO) GetDownlinkList() []OltDetailDownlinkVO`

GetDownlinkList returns the DownlinkList field if non-nil, zero value otherwise.

### GetDownlinkListOk

`func (o *OltDetailVO) GetDownlinkListOk() (*[]OltDetailDownlinkVO, bool)`

GetDownlinkListOk returns a tuple with the DownlinkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkList

`func (o *OltDetailVO) SetDownlinkList(v []OltDetailDownlinkVO)`

SetDownlinkList sets DownlinkList field to given value.

### HasDownlinkList

`func (o *OltDetailVO) HasDownlinkList() bool`

HasDownlinkList returns a boolean if a field has been set.

### GetDueTime

`func (o *OltDetailVO) GetDueTime() int64`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *OltDetailVO) GetDueTimeOk() (*int64, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *OltDetailVO) SetDueTime(v int64)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *OltDetailVO) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### GetDueTimeLeft

`func (o *OltDetailVO) GetDueTimeLeft() int64`

GetDueTimeLeft returns the DueTimeLeft field if non-nil, zero value otherwise.

### GetDueTimeLeftOk

`func (o *OltDetailVO) GetDueTimeLeftOk() (*int64, bool)`

GetDueTimeLeftOk returns a tuple with the DueTimeLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTimeLeft

`func (o *OltDetailVO) SetDueTimeLeft(v int64)`

SetDueTimeLeft sets DueTimeLeft field to given value.

### HasDueTimeLeft

`func (o *OltDetailVO) HasDueTimeLeft() bool`

HasDueTimeLeft returns a boolean if a field has been set.

### GetEcspFirstVersion

`func (o *OltDetailVO) GetEcspFirstVersion() int32`

GetEcspFirstVersion returns the EcspFirstVersion field if non-nil, zero value otherwise.

### GetEcspFirstVersionOk

`func (o *OltDetailVO) GetEcspFirstVersionOk() (*int32, bool)`

GetEcspFirstVersionOk returns a tuple with the EcspFirstVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcspFirstVersion

`func (o *OltDetailVO) SetEcspFirstVersion(v int32)`

SetEcspFirstVersion sets EcspFirstVersion field to given value.

### HasEcspFirstVersion

`func (o *OltDetailVO) HasEcspFirstVersion() bool`

HasEcspFirstVersion returns a boolean if a field has been set.

### GetEos

`func (o *OltDetailVO) GetEos() int64`

GetEos returns the Eos field if non-nil, zero value otherwise.

### GetEosOk

`func (o *OltDetailVO) GetEosOk() (*int64, bool)`

GetEosOk returns a tuple with the Eos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEos

`func (o *OltDetailVO) SetEos(v int64)`

SetEos sets Eos field to given value.

### HasEos

`func (o *OltDetailVO) HasEos() bool`

HasEos returns a boolean if a field has been set.

### GetEost

`func (o *OltDetailVO) GetEost() int64`

GetEost returns the Eost field if non-nil, zero value otherwise.

### GetEostOk

`func (o *OltDetailVO) GetEostOk() (*int64, bool)`

GetEostOk returns a tuple with the Eost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEost

`func (o *OltDetailVO) SetEost(v int64)`

SetEost sets Eost field to given value.

### HasEost

`func (o *OltDetailVO) HasEost() bool`

HasEost returns a boolean if a field has been set.

### GetEs

`func (o *OltDetailVO) GetEs() bool`

GetEs returns the Es field if non-nil, zero value otherwise.

### GetEsOk

`func (o *OltDetailVO) GetEsOk() (*bool, bool)`

GetEsOk returns a tuple with the Es field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEs

`func (o *OltDetailVO) SetEs(v bool)`

SetEs sets Es field to given value.

### HasEs

`func (o *OltDetailVO) HasEs() bool`

HasEs returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *OltDetailVO) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *OltDetailVO) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *OltDetailVO) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *OltDetailVO) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetForgetId

`func (o *OltDetailVO) GetForgetId() string`

GetForgetId returns the ForgetId field if non-nil, zero value otherwise.

### GetForgetIdOk

`func (o *OltDetailVO) GetForgetIdOk() (*string, bool)`

GetForgetIdOk returns a tuple with the ForgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgetId

`func (o *OltDetailVO) SetForgetId(v string)`

SetForgetId sets ForgetId field to given value.

### HasForgetId

`func (o *OltDetailVO) HasForgetId() bool`

HasForgetId returns a boolean if a field has been set.

### GetHwVersion

`func (o *OltDetailVO) GetHwVersion() string`

GetHwVersion returns the HwVersion field if non-nil, zero value otherwise.

### GetHwVersionOk

`func (o *OltDetailVO) GetHwVersionOk() (*string, bool)`

GetHwVersionOk returns a tuple with the HwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwVersion

`func (o *OltDetailVO) SetHwVersion(v string)`

SetHwVersion sets HwVersion field to given value.

### HasHwVersion

`func (o *OltDetailVO) HasHwVersion() bool`

HasHwVersion returns a boolean if a field has been set.

### GetInWhitelist

`func (o *OltDetailVO) GetInWhitelist() bool`

GetInWhitelist returns the InWhitelist field if non-nil, zero value otherwise.

### GetInWhitelistOk

`func (o *OltDetailVO) GetInWhitelistOk() (*bool, bool)`

GetInWhitelistOk returns a tuple with the InWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInWhitelist

`func (o *OltDetailVO) SetInWhitelist(v bool)`

SetInWhitelist sets InWhitelist field to given value.

### HasInWhitelist

`func (o *OltDetailVO) HasInWhitelist() bool`

HasInWhitelist returns a boolean if a field has been set.

### GetInitialUnbindingLimit

`func (o *OltDetailVO) GetInitialUnbindingLimit() int32`

GetInitialUnbindingLimit returns the InitialUnbindingLimit field if non-nil, zero value otherwise.

### GetInitialUnbindingLimitOk

`func (o *OltDetailVO) GetInitialUnbindingLimitOk() (*int32, bool)`

GetInitialUnbindingLimitOk returns a tuple with the InitialUnbindingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialUnbindingLimit

`func (o *OltDetailVO) SetInitialUnbindingLimit(v int32)`

SetInitialUnbindingLimit sets InitialUnbindingLimit field to given value.

### HasInitialUnbindingLimit

`func (o *OltDetailVO) HasInitialUnbindingLimit() bool`

HasInitialUnbindingLimit returns a boolean if a field has been set.

### GetIp

`func (o *OltDetailVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OltDetailVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OltDetailVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OltDetailVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6List

`func (o *OltDetailVO) GetIpv6List() []string`

GetIpv6List returns the Ipv6List field if non-nil, zero value otherwise.

### GetIpv6ListOk

`func (o *OltDetailVO) GetIpv6ListOk() (*[]string, bool)`

GetIpv6ListOk returns a tuple with the Ipv6List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6List

`func (o *OltDetailVO) SetIpv6List(v []string)`

SetIpv6List sets Ipv6List field to given value.

### HasIpv6List

`func (o *OltDetailVO) HasIpv6List() bool`

HasIpv6List returns a boolean if a field has been set.

### GetLastSeen

`func (o *OltDetailVO) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *OltDetailVO) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *OltDetailVO) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *OltDetailVO) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLicenseId

`func (o *OltDetailVO) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *OltDetailVO) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *OltDetailVO) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *OltDetailVO) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetLicenseStatus

`func (o *OltDetailVO) GetLicenseStatus() int32`

GetLicenseStatus returns the LicenseStatus field if non-nil, zero value otherwise.

### GetLicenseStatusOk

`func (o *OltDetailVO) GetLicenseStatusOk() (*int32, bool)`

GetLicenseStatusOk returns a tuple with the LicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatus

`func (o *OltDetailVO) SetLicenseStatus(v int32)`

SetLicenseStatus sets LicenseStatus field to given value.

### HasLicenseStatus

`func (o *OltDetailVO) HasLicenseStatus() bool`

HasLicenseStatus returns a boolean if a field has been set.

### GetLicenseUnbindingLimit

`func (o *OltDetailVO) GetLicenseUnbindingLimit() int32`

GetLicenseUnbindingLimit returns the LicenseUnbindingLimit field if non-nil, zero value otherwise.

### GetLicenseUnbindingLimitOk

`func (o *OltDetailVO) GetLicenseUnbindingLimitOk() (*int32, bool)`

GetLicenseUnbindingLimitOk returns a tuple with the LicenseUnbindingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUnbindingLimit

`func (o *OltDetailVO) SetLicenseUnbindingLimit(v int32)`

SetLicenseUnbindingLimit sets LicenseUnbindingLimit field to given value.

### HasLicenseUnbindingLimit

`func (o *OltDetailVO) HasLicenseUnbindingLimit() bool`

HasLicenseUnbindingLimit returns a boolean if a field has been set.

### GetLocation

`func (o *OltDetailVO) GetLocation() DeviceLocationDetailVO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OltDetailVO) GetLocationOk() (*DeviceLocationDetailVO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OltDetailVO) SetLocation(v DeviceLocationDetailVO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OltDetailVO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMac

`func (o *OltDetailVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OltDetailVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OltDetailVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OltDetailVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMemUtil

`func (o *OltDetailVO) GetMemUtil() int32`

GetMemUtil returns the MemUtil field if non-nil, zero value otherwise.

### GetMemUtilOk

`func (o *OltDetailVO) GetMemUtilOk() (*int32, bool)`

GetMemUtilOk returns a tuple with the MemUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUtil

`func (o *OltDetailVO) SetMemUtil(v int32)`

SetMemUtil sets MemUtil field to given value.

### HasMemUtil

`func (o *OltDetailVO) HasMemUtil() bool`

HasMemUtil returns a boolean if a field has been set.

### GetModel

`func (o *OltDetailVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OltDetailVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OltDetailVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OltDetailVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OltDetailVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OltDetailVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OltDetailVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OltDetailVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetMoveSiteId

`func (o *OltDetailVO) GetMoveSiteId() string`

GetMoveSiteId returns the MoveSiteId field if non-nil, zero value otherwise.

### GetMoveSiteIdOk

`func (o *OltDetailVO) GetMoveSiteIdOk() (*string, bool)`

GetMoveSiteIdOk returns a tuple with the MoveSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveSiteId

`func (o *OltDetailVO) SetMoveSiteId(v string)`

SetMoveSiteId sets MoveSiteId field to given value.

### HasMoveSiteId

`func (o *OltDetailVO) HasMoveSiteId() bool`

HasMoveSiteId returns a boolean if a field has been set.

### GetName

`func (o *OltDetailVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OltDetailVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OltDetailVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OltDetailVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOmadacId

`func (o *OltDetailVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *OltDetailVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *OltDetailVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *OltDetailVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetOnuCount

`func (o *OltDetailVO) GetOnuCount() int32`

GetOnuCount returns the OnuCount field if non-nil, zero value otherwise.

### GetOnuCountOk

`func (o *OltDetailVO) GetOnuCountOk() (*int32, bool)`

GetOnuCountOk returns a tuple with the OnuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuCount

`func (o *OltDetailVO) SetOnuCount(v int32)`

SetOnuCount sets OnuCount field to given value.

### HasOnuCount

`func (o *OltDetailVO) HasOnuCount() bool`

HasOnuCount returns a boolean if a field has been set.

### GetPublicIp

`func (o *OltDetailVO) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *OltDetailVO) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *OltDetailVO) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *OltDetailVO) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRemember

`func (o *OltDetailVO) GetRemember() bool`

GetRemember returns the Remember field if non-nil, zero value otherwise.

### GetRememberOk

`func (o *OltDetailVO) GetRememberOk() (*bool, bool)`

GetRememberOk returns a tuple with the Remember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemember

`func (o *OltDetailVO) SetRemember(v bool)`

SetRemember sets Remember field to given value.

### HasRemember

`func (o *OltDetailVO) HasRemember() bool`

HasRemember returns a boolean if a field has been set.

### GetRememberDevice

`func (o *OltDetailVO) GetRememberDevice() int32`

GetRememberDevice returns the RememberDevice field if non-nil, zero value otherwise.

### GetRememberDeviceOk

`func (o *OltDetailVO) GetRememberDeviceOk() (*int32, bool)`

GetRememberDeviceOk returns a tuple with the RememberDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberDevice

`func (o *OltDetailVO) SetRememberDevice(v int32)`

SetRememberDevice sets RememberDevice field to given value.

### HasRememberDevice

`func (o *OltDetailVO) HasRememberDevice() bool`

HasRememberDevice returns a boolean if a field has been set.

### GetResource

`func (o *OltDetailVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *OltDetailVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *OltDetailVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *OltDetailVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetShowModel

`func (o *OltDetailVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *OltDetailVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *OltDetailVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *OltDetailVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetSiteName

`func (o *OltDetailVO) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *OltDetailVO) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *OltDetailVO) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *OltDetailVO) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSiteTemplateId

`func (o *OltDetailVO) GetSiteTemplateId() string`

GetSiteTemplateId returns the SiteTemplateId field if non-nil, zero value otherwise.

### GetSiteTemplateIdOk

`func (o *OltDetailVO) GetSiteTemplateIdOk() (*string, bool)`

GetSiteTemplateIdOk returns a tuple with the SiteTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTemplateId

`func (o *OltDetailVO) SetSiteTemplateId(v string)`

SetSiteTemplateId sets SiteTemplateId field to given value.

### HasSiteTemplateId

`func (o *OltDetailVO) HasSiteTemplateId() bool`

HasSiteTemplateId returns a boolean if a field has been set.

### GetSiteTemplateName

`func (o *OltDetailVO) GetSiteTemplateName() string`

GetSiteTemplateName returns the SiteTemplateName field if non-nil, zero value otherwise.

### GetSiteTemplateNameOk

`func (o *OltDetailVO) GetSiteTemplateNameOk() (*string, bool)`

GetSiteTemplateNameOk returns a tuple with the SiteTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTemplateName

`func (o *OltDetailVO) SetSiteTemplateName(v string)`

SetSiteTemplateName sets SiteTemplateName field to given value.

### HasSiteTemplateName

`func (o *OltDetailVO) HasSiteTemplateName() bool`

HasSiteTemplateName returns a boolean if a field has been set.

### GetSn

`func (o *OltDetailVO) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *OltDetailVO) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *OltDetailVO) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *OltDetailVO) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetSpecialModel

`func (o *OltDetailVO) GetSpecialModel() string`

GetSpecialModel returns the SpecialModel field if non-nil, zero value otherwise.

### GetSpecialModelOk

`func (o *OltDetailVO) GetSpecialModelOk() (*string, bool)`

GetSpecialModelOk returns a tuple with the SpecialModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialModel

`func (o *OltDetailVO) SetSpecialModel(v string)`

SetSpecialModel sets SpecialModel field to given value.

### HasSpecialModel

`func (o *OltDetailVO) HasSpecialModel() bool`

HasSpecialModel returns a boolean if a field has been set.

### GetSpeeds

`func (o *OltDetailVO) GetSpeeds() []int32`

GetSpeeds returns the Speeds field if non-nil, zero value otherwise.

### GetSpeedsOk

`func (o *OltDetailVO) GetSpeedsOk() (*[]int32, bool)`

GetSpeedsOk returns a tuple with the Speeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeeds

`func (o *OltDetailVO) SetSpeeds(v []int32)`

SetSpeeds sets Speeds field to given value.

### HasSpeeds

`func (o *OltDetailVO) HasSpeeds() bool`

HasSpeeds returns a boolean if a field has been set.

### GetStatus

`func (o *OltDetailVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OltDetailVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OltDetailVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OltDetailVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *OltDetailVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *OltDetailVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *OltDetailVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *OltDetailVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetSupportAnomaly

`func (o *OltDetailVO) GetSupportAnomaly() bool`

GetSupportAnomaly returns the SupportAnomaly field if non-nil, zero value otherwise.

### GetSupportAnomalyOk

`func (o *OltDetailVO) GetSupportAnomalyOk() (*bool, bool)`

GetSupportAnomalyOk returns a tuple with the SupportAnomaly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAnomaly

`func (o *OltDetailVO) SetSupportAnomaly(v bool)`

SetSupportAnomaly sets SupportAnomaly field to given value.

### HasSupportAnomaly

`func (o *OltDetailVO) HasSupportAnomaly() bool`

HasSupportAnomaly returns a boolean if a field has been set.

### GetSupportLocatePort

`func (o *OltDetailVO) GetSupportLocatePort() bool`

GetSupportLocatePort returns the SupportLocatePort field if non-nil, zero value otherwise.

### GetSupportLocatePortOk

`func (o *OltDetailVO) GetSupportLocatePortOk() (*bool, bool)`

GetSupportLocatePortOk returns a tuple with the SupportLocatePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLocatePort

`func (o *OltDetailVO) SetSupportLocatePort(v bool)`

SetSupportLocatePort sets SupportLocatePort field to given value.

### HasSupportLocatePort

`func (o *OltDetailVO) HasSupportLocatePort() bool`

HasSupportLocatePort returns a boolean if a field has been set.

### GetTagIds

`func (o *OltDetailVO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *OltDetailVO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *OltDetailVO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *OltDetailVO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTemplateId

`func (o *OltDetailVO) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *OltDetailVO) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *OltDetailVO) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *OltDetailVO) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateName

`func (o *OltDetailVO) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *OltDetailVO) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *OltDetailVO) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *OltDetailVO) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetType

`func (o *OltDetailVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OltDetailVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OltDetailVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OltDetailVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUp

`func (o *OltDetailVO) GetUp() int64`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *OltDetailVO) GetUpOk() (*int64, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *OltDetailVO) SetUp(v int64)`

SetUp sets Up field to given value.

### HasUp

`func (o *OltDetailVO) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetUplink

`func (o *OltDetailVO) GetUplink() OltUplinkVO`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *OltDetailVO) GetUplinkOk() (*OltUplinkVO, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *OltDetailVO) SetUplink(v OltUplinkVO)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *OltDetailVO) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetUptime

`func (o *OltDetailVO) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *OltDetailVO) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *OltDetailVO) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *OltDetailVO) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetUptimeLong

`func (o *OltDetailVO) GetUptimeLong() int64`

GetUptimeLong returns the UptimeLong field if non-nil, zero value otherwise.

### GetUptimeLongOk

`func (o *OltDetailVO) GetUptimeLongOk() (*int64, bool)`

GetUptimeLongOk returns a tuple with the UptimeLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeLong

`func (o *OltDetailVO) SetUptimeLong(v int64)`

SetUptimeLong sets UptimeLong field to given value.

### HasUptimeLong

`func (o *OltDetailVO) HasUptimeLong() bool`

HasUptimeLong returns a boolean if a field has been set.

### GetVersion

`func (o *OltDetailVO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OltDetailVO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OltDetailVO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OltDetailVO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


