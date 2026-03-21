# OswDetailForPortsViewVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | Site of the device | [optional] 
**Active** | Pointer to **bool** | whether to active the device(cloud base exclusive) | [optional] 
**AddedInAdvanced** | Pointer to **bool** | Whether the device is added in advanced. | [optional] 
**BoundDeviceTemplate** | Pointer to **bool** | Whether the device is bound to device template | [optional] 
**BoundSiteTemplate** | Pointer to **bool** | Whether the site where the device is located is bound to a site template | [optional] 
**Category** | Pointer to **string** | Category of license | [optional] 
**ClientList** | Pointer to [**[]OswDownLinkClientVO**](OswDownLinkClientVO.md) | Downlink Client device list | [optional] 
**Compatible** | Pointer to **int32** | Device firmware and controller compatibility type.Compatible should be a value as follows: 0:COMPATIBLE;1:HIGH_MAJOR_VER;2:LOW_MAJOR_VER;3:HIGH_MINOR_VER;4:LOW_MINOR_VER;7:HIGH_COMPONENT_VER;10:DEVICE_NOT_COMPATIBLE;11:HIGH_ADOPT_COMMPONENT;12:DEVICE_CATEGORY_NOT_COMPATIBLE;14:DEVICE_NOT_COMPATIBLE_IN_CLUSTER | [optional] 
**CompoundModel** | Pointer to **string** | Model complex used in the backend.Ap：model+(country)+modelVersion,  EAP225(EU) v3.0 Ap: specialModel+modelVersion, EAP225-Outdoor-1a20a950b8d950e8 v1.0  Gateway/Switch：model+modelVersion, Osg v3.0 | [optional] 
**CustomId** | Pointer to **string** | Customer ID | [optional] 
**CustomName** | Pointer to **string** | Customer name | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DeviceSeriesType** | Pointer to **int32** | Device series type.DeviceSeriesType should be a value as follows: 0:advanced;1:pro | [optional] 
**DeviceTemplateAvailable** | Pointer to **bool** | Whether there is an available device template for the device; it is false if the model is not supported or the site template has not created the corresponding device template. | [optional] 
**DisableHwReset** | Pointer to **bool** |  | [optional] 
**DownlinkList** | Pointer to [**[]OswDownlinkVO**](OswDownlinkVO.md) | Downlink Omada device list | [optional] 
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
**LicenseId** | Pointer to **string** | License key on detail page of device(cloud base exclusive) | [optional] 
**LicenseStatus** | Pointer to **int32** | License status(cloud base exclusive).LicenseStatus should be a value as follows: 0:unActive 1:Unbind 2:Expired 3:active | [optional] 
**LicenseUnbindingLimit** | Pointer to **int32** | Remaining unbind count for license on detail Page of device(cloud base exclusive) | [optional] 
**LocatingPorts** | Pointer to **[]int32** | Locating switch ports, effective when parameter[locateEnable] is true and device is not a member of stack | [optional] 
**LocatingStandardPorts** | Pointer to **[]string** | Locating switch standard ports, effective when parameter[locateEnable] is true and device is a member of stack | [optional] 
**Location** | Pointer to [**DeviceLocationDetailVO**](DeviceLocationDetailVO.md) |  | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**Model** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**MoveSiteId** | Pointer to **string** | Record that the device is in a moveSite operation; if it is null, then it is not in the moveSite operation. | [optional] 
**Name** | Pointer to **string** | Device name,default value is the mac address of device | [optional] 
**OmadacId** | Pointer to **string** | OmadacId of the device | [optional] 
**Remember** | Pointer to **bool** | Whether to remember the device(deprecated) | [optional] 
**RememberDevice** | Pointer to **int32** | Whether to remember the device.RememberDevice should be a value as follows: 0:off, 1:on, 2: follow site | [optional] 
**Resource** | Pointer to **int32** | Data source.Resource should be a value as follows: 0:new created;1:from template;2:override | [optional] 
**ShowModel** | Pointer to **string** | Model complex shown in the front end.Ap：model+(country)+modelVersion,EAP225(EU) v3.0  Gateway/Switch：model+modelVersion,Osg v3.0 | [optional] 
**SiteName** | Pointer to **string** | Site name of the device | [optional] 
**SiteTemplateId** | Pointer to **string** | Template ID bound to the site | [optional] 
**SiteTemplateName** | Pointer to **string** | Template name bound to the site | [optional] 
**Sn** | Pointer to **string** | SN code of device | [optional] 
**SpecialModel** | Pointer to **string** | Special device model,for example:EAP225-Outdoor-1a20a950b8d950e8 | [optional] 
**StackMsg** | Pointer to [**StackMsgVO**](StackMsgVO.md) |  | [optional] 
**Status** | Pointer to **int32** | Status of device,status should be a value as follows: 0:Disconnected;1:Disconnected(Migrating);10:Provisioning;11:Configuring;12:Upgrading;13:Rebooting;14:Connected;15:Connected(Wireless);16:Connected(Migrating);17:Connected(Wireless,Migrating);20:Pending;21:Pending(Wireless);22:Adopting;23:Adopting(Wireless);24:Adopt Failed;25:Adopt Failed(Wireless);26:Managed By Others;27:Managed By Others(Wireless);30:Heartbeat Missed;31:Heartbeat Missed(Wireless);32:Heartbeat Missed(Migrating);33:Heartbeat Missed(Wireless,Migrating);40:Isolated;41:Isolated(Migrating);50:Slice Configuring | [optional] 
**StatusCategory** | Pointer to **int32** | Category of device status,statusCategory should be a value as follows: 0:Disconnected;1:Connected;2:Pending;3:Heartbeat Missed;4:Isolated | [optional] 
**StpLinkList** | Pointer to [**[]OswDownlinkVO**](OswDownlinkVO.md) | STP Blocked Link Omada device list | [optional] 
**SupportAnomaly** | Pointer to **bool** | Whether the device firmware support intelligent anomaly detection | [optional] 
**SupportClearCounters** | Pointer to **bool** | Clear Counters Support | [optional] 
**SupportLocatePort** | Pointer to **bool** | Whether the device supports locating port | [optional] 
**TemplateId** | Pointer to **string** | ID of the template bound to the device | [optional] 
**TemplateName** | Pointer to **string** | Name of the template bound to the device | [optional] 
**Type** | Pointer to **string** | Device type:ap、gateway、switch、olt | [optional] 
**Uplink** | Pointer to [**OswUplinkVO**](OswUplinkVO.md) |  | [optional] 
**Version** | Pointer to **string** | Simplified version of firmware,for example:2.5.0 | [optional] 

## Methods

### NewOswDetailForPortsViewVO

`func NewOswDetailForPortsViewVO() *OswDetailForPortsViewVO`

NewOswDetailForPortsViewVO instantiates a new OswDetailForPortsViewVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswDetailForPortsViewVOWithDefaults

`func NewOswDetailForPortsViewVOWithDefaults() *OswDetailForPortsViewVO`

NewOswDetailForPortsViewVOWithDefaults instantiates a new OswDetailForPortsViewVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *OswDetailForPortsViewVO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *OswDetailForPortsViewVO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *OswDetailForPortsViewVO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *OswDetailForPortsViewVO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetActive

`func (o *OswDetailForPortsViewVO) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OswDetailForPortsViewVO) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OswDetailForPortsViewVO) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *OswDetailForPortsViewVO) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddedInAdvanced

`func (o *OswDetailForPortsViewVO) GetAddedInAdvanced() bool`

GetAddedInAdvanced returns the AddedInAdvanced field if non-nil, zero value otherwise.

### GetAddedInAdvancedOk

`func (o *OswDetailForPortsViewVO) GetAddedInAdvancedOk() (*bool, bool)`

GetAddedInAdvancedOk returns a tuple with the AddedInAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedInAdvanced

`func (o *OswDetailForPortsViewVO) SetAddedInAdvanced(v bool)`

SetAddedInAdvanced sets AddedInAdvanced field to given value.

### HasAddedInAdvanced

`func (o *OswDetailForPortsViewVO) HasAddedInAdvanced() bool`

HasAddedInAdvanced returns a boolean if a field has been set.

### GetBoundDeviceTemplate

`func (o *OswDetailForPortsViewVO) GetBoundDeviceTemplate() bool`

GetBoundDeviceTemplate returns the BoundDeviceTemplate field if non-nil, zero value otherwise.

### GetBoundDeviceTemplateOk

`func (o *OswDetailForPortsViewVO) GetBoundDeviceTemplateOk() (*bool, bool)`

GetBoundDeviceTemplateOk returns a tuple with the BoundDeviceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundDeviceTemplate

`func (o *OswDetailForPortsViewVO) SetBoundDeviceTemplate(v bool)`

SetBoundDeviceTemplate sets BoundDeviceTemplate field to given value.

### HasBoundDeviceTemplate

`func (o *OswDetailForPortsViewVO) HasBoundDeviceTemplate() bool`

HasBoundDeviceTemplate returns a boolean if a field has been set.

### GetBoundSiteTemplate

`func (o *OswDetailForPortsViewVO) GetBoundSiteTemplate() bool`

GetBoundSiteTemplate returns the BoundSiteTemplate field if non-nil, zero value otherwise.

### GetBoundSiteTemplateOk

`func (o *OswDetailForPortsViewVO) GetBoundSiteTemplateOk() (*bool, bool)`

GetBoundSiteTemplateOk returns a tuple with the BoundSiteTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSiteTemplate

`func (o *OswDetailForPortsViewVO) SetBoundSiteTemplate(v bool)`

SetBoundSiteTemplate sets BoundSiteTemplate field to given value.

### HasBoundSiteTemplate

`func (o *OswDetailForPortsViewVO) HasBoundSiteTemplate() bool`

HasBoundSiteTemplate returns a boolean if a field has been set.

### GetCategory

`func (o *OswDetailForPortsViewVO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *OswDetailForPortsViewVO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *OswDetailForPortsViewVO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *OswDetailForPortsViewVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetClientList

`func (o *OswDetailForPortsViewVO) GetClientList() []OswDownLinkClientVO`

GetClientList returns the ClientList field if non-nil, zero value otherwise.

### GetClientListOk

`func (o *OswDetailForPortsViewVO) GetClientListOk() (*[]OswDownLinkClientVO, bool)`

GetClientListOk returns a tuple with the ClientList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientList

`func (o *OswDetailForPortsViewVO) SetClientList(v []OswDownLinkClientVO)`

SetClientList sets ClientList field to given value.

### HasClientList

`func (o *OswDetailForPortsViewVO) HasClientList() bool`

HasClientList returns a boolean if a field has been set.

### GetCompatible

`func (o *OswDetailForPortsViewVO) GetCompatible() int32`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *OswDetailForPortsViewVO) GetCompatibleOk() (*int32, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *OswDetailForPortsViewVO) SetCompatible(v int32)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *OswDetailForPortsViewVO) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetCompoundModel

`func (o *OswDetailForPortsViewVO) GetCompoundModel() string`

GetCompoundModel returns the CompoundModel field if non-nil, zero value otherwise.

### GetCompoundModelOk

`func (o *OswDetailForPortsViewVO) GetCompoundModelOk() (*string, bool)`

GetCompoundModelOk returns a tuple with the CompoundModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompoundModel

`func (o *OswDetailForPortsViewVO) SetCompoundModel(v string)`

SetCompoundModel sets CompoundModel field to given value.

### HasCompoundModel

`func (o *OswDetailForPortsViewVO) HasCompoundModel() bool`

HasCompoundModel returns a boolean if a field has been set.

### GetCustomId

`func (o *OswDetailForPortsViewVO) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *OswDetailForPortsViewVO) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *OswDetailForPortsViewVO) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.

### HasCustomId

`func (o *OswDetailForPortsViewVO) HasCustomId() bool`

HasCustomId returns a boolean if a field has been set.

### GetCustomName

`func (o *OswDetailForPortsViewVO) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *OswDetailForPortsViewVO) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *OswDetailForPortsViewVO) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *OswDetailForPortsViewVO) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### GetDescription

`func (o *OswDetailForPortsViewVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OswDetailForPortsViewVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OswDetailForPortsViewVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OswDetailForPortsViewVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceSeriesType

`func (o *OswDetailForPortsViewVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *OswDetailForPortsViewVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *OswDetailForPortsViewVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *OswDetailForPortsViewVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetDeviceTemplateAvailable

`func (o *OswDetailForPortsViewVO) GetDeviceTemplateAvailable() bool`

GetDeviceTemplateAvailable returns the DeviceTemplateAvailable field if non-nil, zero value otherwise.

### GetDeviceTemplateAvailableOk

`func (o *OswDetailForPortsViewVO) GetDeviceTemplateAvailableOk() (*bool, bool)`

GetDeviceTemplateAvailableOk returns a tuple with the DeviceTemplateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTemplateAvailable

`func (o *OswDetailForPortsViewVO) SetDeviceTemplateAvailable(v bool)`

SetDeviceTemplateAvailable sets DeviceTemplateAvailable field to given value.

### HasDeviceTemplateAvailable

`func (o *OswDetailForPortsViewVO) HasDeviceTemplateAvailable() bool`

HasDeviceTemplateAvailable returns a boolean if a field has been set.

### GetDisableHwReset

`func (o *OswDetailForPortsViewVO) GetDisableHwReset() bool`

GetDisableHwReset returns the DisableHwReset field if non-nil, zero value otherwise.

### GetDisableHwResetOk

`func (o *OswDetailForPortsViewVO) GetDisableHwResetOk() (*bool, bool)`

GetDisableHwResetOk returns a tuple with the DisableHwReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHwReset

`func (o *OswDetailForPortsViewVO) SetDisableHwReset(v bool)`

SetDisableHwReset sets DisableHwReset field to given value.

### HasDisableHwReset

`func (o *OswDetailForPortsViewVO) HasDisableHwReset() bool`

HasDisableHwReset returns a boolean if a field has been set.

### GetDownlinkList

`func (o *OswDetailForPortsViewVO) GetDownlinkList() []OswDownlinkVO`

GetDownlinkList returns the DownlinkList field if non-nil, zero value otherwise.

### GetDownlinkListOk

`func (o *OswDetailForPortsViewVO) GetDownlinkListOk() (*[]OswDownlinkVO, bool)`

GetDownlinkListOk returns a tuple with the DownlinkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkList

`func (o *OswDetailForPortsViewVO) SetDownlinkList(v []OswDownlinkVO)`

SetDownlinkList sets DownlinkList field to given value.

### HasDownlinkList

`func (o *OswDetailForPortsViewVO) HasDownlinkList() bool`

HasDownlinkList returns a boolean if a field has been set.

### GetDueTime

`func (o *OswDetailForPortsViewVO) GetDueTime() int64`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *OswDetailForPortsViewVO) GetDueTimeOk() (*int64, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *OswDetailForPortsViewVO) SetDueTime(v int64)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *OswDetailForPortsViewVO) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### GetDueTimeLeft

`func (o *OswDetailForPortsViewVO) GetDueTimeLeft() int64`

GetDueTimeLeft returns the DueTimeLeft field if non-nil, zero value otherwise.

### GetDueTimeLeftOk

`func (o *OswDetailForPortsViewVO) GetDueTimeLeftOk() (*int64, bool)`

GetDueTimeLeftOk returns a tuple with the DueTimeLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTimeLeft

`func (o *OswDetailForPortsViewVO) SetDueTimeLeft(v int64)`

SetDueTimeLeft sets DueTimeLeft field to given value.

### HasDueTimeLeft

`func (o *OswDetailForPortsViewVO) HasDueTimeLeft() bool`

HasDueTimeLeft returns a boolean if a field has been set.

### GetEcspFirstVersion

`func (o *OswDetailForPortsViewVO) GetEcspFirstVersion() int32`

GetEcspFirstVersion returns the EcspFirstVersion field if non-nil, zero value otherwise.

### GetEcspFirstVersionOk

`func (o *OswDetailForPortsViewVO) GetEcspFirstVersionOk() (*int32, bool)`

GetEcspFirstVersionOk returns a tuple with the EcspFirstVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcspFirstVersion

`func (o *OswDetailForPortsViewVO) SetEcspFirstVersion(v int32)`

SetEcspFirstVersion sets EcspFirstVersion field to given value.

### HasEcspFirstVersion

`func (o *OswDetailForPortsViewVO) HasEcspFirstVersion() bool`

HasEcspFirstVersion returns a boolean if a field has been set.

### GetEos

`func (o *OswDetailForPortsViewVO) GetEos() int64`

GetEos returns the Eos field if non-nil, zero value otherwise.

### GetEosOk

`func (o *OswDetailForPortsViewVO) GetEosOk() (*int64, bool)`

GetEosOk returns a tuple with the Eos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEos

`func (o *OswDetailForPortsViewVO) SetEos(v int64)`

SetEos sets Eos field to given value.

### HasEos

`func (o *OswDetailForPortsViewVO) HasEos() bool`

HasEos returns a boolean if a field has been set.

### GetEost

`func (o *OswDetailForPortsViewVO) GetEost() int64`

GetEost returns the Eost field if non-nil, zero value otherwise.

### GetEostOk

`func (o *OswDetailForPortsViewVO) GetEostOk() (*int64, bool)`

GetEostOk returns a tuple with the Eost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEost

`func (o *OswDetailForPortsViewVO) SetEost(v int64)`

SetEost sets Eost field to given value.

### HasEost

`func (o *OswDetailForPortsViewVO) HasEost() bool`

HasEost returns a boolean if a field has been set.

### GetEs

`func (o *OswDetailForPortsViewVO) GetEs() bool`

GetEs returns the Es field if non-nil, zero value otherwise.

### GetEsOk

`func (o *OswDetailForPortsViewVO) GetEsOk() (*bool, bool)`

GetEsOk returns a tuple with the Es field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEs

`func (o *OswDetailForPortsViewVO) SetEs(v bool)`

SetEs sets Es field to given value.

### HasEs

`func (o *OswDetailForPortsViewVO) HasEs() bool`

HasEs returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *OswDetailForPortsViewVO) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *OswDetailForPortsViewVO) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *OswDetailForPortsViewVO) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *OswDetailForPortsViewVO) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetForgetId

`func (o *OswDetailForPortsViewVO) GetForgetId() string`

GetForgetId returns the ForgetId field if non-nil, zero value otherwise.

### GetForgetIdOk

`func (o *OswDetailForPortsViewVO) GetForgetIdOk() (*string, bool)`

GetForgetIdOk returns a tuple with the ForgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgetId

`func (o *OswDetailForPortsViewVO) SetForgetId(v string)`

SetForgetId sets ForgetId field to given value.

### HasForgetId

`func (o *OswDetailForPortsViewVO) HasForgetId() bool`

HasForgetId returns a boolean if a field has been set.

### GetHwVersion

`func (o *OswDetailForPortsViewVO) GetHwVersion() string`

GetHwVersion returns the HwVersion field if non-nil, zero value otherwise.

### GetHwVersionOk

`func (o *OswDetailForPortsViewVO) GetHwVersionOk() (*string, bool)`

GetHwVersionOk returns a tuple with the HwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwVersion

`func (o *OswDetailForPortsViewVO) SetHwVersion(v string)`

SetHwVersion sets HwVersion field to given value.

### HasHwVersion

`func (o *OswDetailForPortsViewVO) HasHwVersion() bool`

HasHwVersion returns a boolean if a field has been set.

### GetInWhitelist

`func (o *OswDetailForPortsViewVO) GetInWhitelist() bool`

GetInWhitelist returns the InWhitelist field if non-nil, zero value otherwise.

### GetInWhitelistOk

`func (o *OswDetailForPortsViewVO) GetInWhitelistOk() (*bool, bool)`

GetInWhitelistOk returns a tuple with the InWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInWhitelist

`func (o *OswDetailForPortsViewVO) SetInWhitelist(v bool)`

SetInWhitelist sets InWhitelist field to given value.

### HasInWhitelist

`func (o *OswDetailForPortsViewVO) HasInWhitelist() bool`

HasInWhitelist returns a boolean if a field has been set.

### GetInitialUnbindingLimit

`func (o *OswDetailForPortsViewVO) GetInitialUnbindingLimit() int32`

GetInitialUnbindingLimit returns the InitialUnbindingLimit field if non-nil, zero value otherwise.

### GetInitialUnbindingLimitOk

`func (o *OswDetailForPortsViewVO) GetInitialUnbindingLimitOk() (*int32, bool)`

GetInitialUnbindingLimitOk returns a tuple with the InitialUnbindingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialUnbindingLimit

`func (o *OswDetailForPortsViewVO) SetInitialUnbindingLimit(v int32)`

SetInitialUnbindingLimit sets InitialUnbindingLimit field to given value.

### HasInitialUnbindingLimit

`func (o *OswDetailForPortsViewVO) HasInitialUnbindingLimit() bool`

HasInitialUnbindingLimit returns a boolean if a field has been set.

### GetLicenseId

`func (o *OswDetailForPortsViewVO) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *OswDetailForPortsViewVO) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *OswDetailForPortsViewVO) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *OswDetailForPortsViewVO) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetLicenseStatus

`func (o *OswDetailForPortsViewVO) GetLicenseStatus() int32`

GetLicenseStatus returns the LicenseStatus field if non-nil, zero value otherwise.

### GetLicenseStatusOk

`func (o *OswDetailForPortsViewVO) GetLicenseStatusOk() (*int32, bool)`

GetLicenseStatusOk returns a tuple with the LicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatus

`func (o *OswDetailForPortsViewVO) SetLicenseStatus(v int32)`

SetLicenseStatus sets LicenseStatus field to given value.

### HasLicenseStatus

`func (o *OswDetailForPortsViewVO) HasLicenseStatus() bool`

HasLicenseStatus returns a boolean if a field has been set.

### GetLicenseUnbindingLimit

`func (o *OswDetailForPortsViewVO) GetLicenseUnbindingLimit() int32`

GetLicenseUnbindingLimit returns the LicenseUnbindingLimit field if non-nil, zero value otherwise.

### GetLicenseUnbindingLimitOk

`func (o *OswDetailForPortsViewVO) GetLicenseUnbindingLimitOk() (*int32, bool)`

GetLicenseUnbindingLimitOk returns a tuple with the LicenseUnbindingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUnbindingLimit

`func (o *OswDetailForPortsViewVO) SetLicenseUnbindingLimit(v int32)`

SetLicenseUnbindingLimit sets LicenseUnbindingLimit field to given value.

### HasLicenseUnbindingLimit

`func (o *OswDetailForPortsViewVO) HasLicenseUnbindingLimit() bool`

HasLicenseUnbindingLimit returns a boolean if a field has been set.

### GetLocatingPorts

`func (o *OswDetailForPortsViewVO) GetLocatingPorts() []int32`

GetLocatingPorts returns the LocatingPorts field if non-nil, zero value otherwise.

### GetLocatingPortsOk

`func (o *OswDetailForPortsViewVO) GetLocatingPortsOk() (*[]int32, bool)`

GetLocatingPortsOk returns a tuple with the LocatingPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatingPorts

`func (o *OswDetailForPortsViewVO) SetLocatingPorts(v []int32)`

SetLocatingPorts sets LocatingPorts field to given value.

### HasLocatingPorts

`func (o *OswDetailForPortsViewVO) HasLocatingPorts() bool`

HasLocatingPorts returns a boolean if a field has been set.

### GetLocatingStandardPorts

`func (o *OswDetailForPortsViewVO) GetLocatingStandardPorts() []string`

GetLocatingStandardPorts returns the LocatingStandardPorts field if non-nil, zero value otherwise.

### GetLocatingStandardPortsOk

`func (o *OswDetailForPortsViewVO) GetLocatingStandardPortsOk() (*[]string, bool)`

GetLocatingStandardPortsOk returns a tuple with the LocatingStandardPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatingStandardPorts

`func (o *OswDetailForPortsViewVO) SetLocatingStandardPorts(v []string)`

SetLocatingStandardPorts sets LocatingStandardPorts field to given value.

### HasLocatingStandardPorts

`func (o *OswDetailForPortsViewVO) HasLocatingStandardPorts() bool`

HasLocatingStandardPorts returns a boolean if a field has been set.

### GetLocation

`func (o *OswDetailForPortsViewVO) GetLocation() DeviceLocationDetailVO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OswDetailForPortsViewVO) GetLocationOk() (*DeviceLocationDetailVO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OswDetailForPortsViewVO) SetLocation(v DeviceLocationDetailVO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OswDetailForPortsViewVO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMac

`func (o *OswDetailForPortsViewVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswDetailForPortsViewVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswDetailForPortsViewVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswDetailForPortsViewVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *OswDetailForPortsViewVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OswDetailForPortsViewVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OswDetailForPortsViewVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OswDetailForPortsViewVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OswDetailForPortsViewVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OswDetailForPortsViewVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OswDetailForPortsViewVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OswDetailForPortsViewVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetMoveSiteId

`func (o *OswDetailForPortsViewVO) GetMoveSiteId() string`

GetMoveSiteId returns the MoveSiteId field if non-nil, zero value otherwise.

### GetMoveSiteIdOk

`func (o *OswDetailForPortsViewVO) GetMoveSiteIdOk() (*string, bool)`

GetMoveSiteIdOk returns a tuple with the MoveSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveSiteId

`func (o *OswDetailForPortsViewVO) SetMoveSiteId(v string)`

SetMoveSiteId sets MoveSiteId field to given value.

### HasMoveSiteId

`func (o *OswDetailForPortsViewVO) HasMoveSiteId() bool`

HasMoveSiteId returns a boolean if a field has been set.

### GetName

`func (o *OswDetailForPortsViewVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswDetailForPortsViewVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswDetailForPortsViewVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswDetailForPortsViewVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOmadacId

`func (o *OswDetailForPortsViewVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *OswDetailForPortsViewVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *OswDetailForPortsViewVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *OswDetailForPortsViewVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetRemember

`func (o *OswDetailForPortsViewVO) GetRemember() bool`

GetRemember returns the Remember field if non-nil, zero value otherwise.

### GetRememberOk

`func (o *OswDetailForPortsViewVO) GetRememberOk() (*bool, bool)`

GetRememberOk returns a tuple with the Remember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemember

`func (o *OswDetailForPortsViewVO) SetRemember(v bool)`

SetRemember sets Remember field to given value.

### HasRemember

`func (o *OswDetailForPortsViewVO) HasRemember() bool`

HasRemember returns a boolean if a field has been set.

### GetRememberDevice

`func (o *OswDetailForPortsViewVO) GetRememberDevice() int32`

GetRememberDevice returns the RememberDevice field if non-nil, zero value otherwise.

### GetRememberDeviceOk

`func (o *OswDetailForPortsViewVO) GetRememberDeviceOk() (*int32, bool)`

GetRememberDeviceOk returns a tuple with the RememberDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberDevice

`func (o *OswDetailForPortsViewVO) SetRememberDevice(v int32)`

SetRememberDevice sets RememberDevice field to given value.

### HasRememberDevice

`func (o *OswDetailForPortsViewVO) HasRememberDevice() bool`

HasRememberDevice returns a boolean if a field has been set.

### GetResource

`func (o *OswDetailForPortsViewVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *OswDetailForPortsViewVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *OswDetailForPortsViewVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *OswDetailForPortsViewVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetShowModel

`func (o *OswDetailForPortsViewVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *OswDetailForPortsViewVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *OswDetailForPortsViewVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *OswDetailForPortsViewVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetSiteName

`func (o *OswDetailForPortsViewVO) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *OswDetailForPortsViewVO) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *OswDetailForPortsViewVO) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *OswDetailForPortsViewVO) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSiteTemplateId

`func (o *OswDetailForPortsViewVO) GetSiteTemplateId() string`

GetSiteTemplateId returns the SiteTemplateId field if non-nil, zero value otherwise.

### GetSiteTemplateIdOk

`func (o *OswDetailForPortsViewVO) GetSiteTemplateIdOk() (*string, bool)`

GetSiteTemplateIdOk returns a tuple with the SiteTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTemplateId

`func (o *OswDetailForPortsViewVO) SetSiteTemplateId(v string)`

SetSiteTemplateId sets SiteTemplateId field to given value.

### HasSiteTemplateId

`func (o *OswDetailForPortsViewVO) HasSiteTemplateId() bool`

HasSiteTemplateId returns a boolean if a field has been set.

### GetSiteTemplateName

`func (o *OswDetailForPortsViewVO) GetSiteTemplateName() string`

GetSiteTemplateName returns the SiteTemplateName field if non-nil, zero value otherwise.

### GetSiteTemplateNameOk

`func (o *OswDetailForPortsViewVO) GetSiteTemplateNameOk() (*string, bool)`

GetSiteTemplateNameOk returns a tuple with the SiteTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTemplateName

`func (o *OswDetailForPortsViewVO) SetSiteTemplateName(v string)`

SetSiteTemplateName sets SiteTemplateName field to given value.

### HasSiteTemplateName

`func (o *OswDetailForPortsViewVO) HasSiteTemplateName() bool`

HasSiteTemplateName returns a boolean if a field has been set.

### GetSn

`func (o *OswDetailForPortsViewVO) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *OswDetailForPortsViewVO) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *OswDetailForPortsViewVO) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *OswDetailForPortsViewVO) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetSpecialModel

`func (o *OswDetailForPortsViewVO) GetSpecialModel() string`

GetSpecialModel returns the SpecialModel field if non-nil, zero value otherwise.

### GetSpecialModelOk

`func (o *OswDetailForPortsViewVO) GetSpecialModelOk() (*string, bool)`

GetSpecialModelOk returns a tuple with the SpecialModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialModel

`func (o *OswDetailForPortsViewVO) SetSpecialModel(v string)`

SetSpecialModel sets SpecialModel field to given value.

### HasSpecialModel

`func (o *OswDetailForPortsViewVO) HasSpecialModel() bool`

HasSpecialModel returns a boolean if a field has been set.

### GetStackMsg

`func (o *OswDetailForPortsViewVO) GetStackMsg() StackMsgVO`

GetStackMsg returns the StackMsg field if non-nil, zero value otherwise.

### GetStackMsgOk

`func (o *OswDetailForPortsViewVO) GetStackMsgOk() (*StackMsgVO, bool)`

GetStackMsgOk returns a tuple with the StackMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackMsg

`func (o *OswDetailForPortsViewVO) SetStackMsg(v StackMsgVO)`

SetStackMsg sets StackMsg field to given value.

### HasStackMsg

`func (o *OswDetailForPortsViewVO) HasStackMsg() bool`

HasStackMsg returns a boolean if a field has been set.

### GetStatus

`func (o *OswDetailForPortsViewVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswDetailForPortsViewVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswDetailForPortsViewVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OswDetailForPortsViewVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *OswDetailForPortsViewVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *OswDetailForPortsViewVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *OswDetailForPortsViewVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *OswDetailForPortsViewVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetStpLinkList

`func (o *OswDetailForPortsViewVO) GetStpLinkList() []OswDownlinkVO`

GetStpLinkList returns the StpLinkList field if non-nil, zero value otherwise.

### GetStpLinkListOk

`func (o *OswDetailForPortsViewVO) GetStpLinkListOk() (*[]OswDownlinkVO, bool)`

GetStpLinkListOk returns a tuple with the StpLinkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpLinkList

`func (o *OswDetailForPortsViewVO) SetStpLinkList(v []OswDownlinkVO)`

SetStpLinkList sets StpLinkList field to given value.

### HasStpLinkList

`func (o *OswDetailForPortsViewVO) HasStpLinkList() bool`

HasStpLinkList returns a boolean if a field has been set.

### GetSupportAnomaly

`func (o *OswDetailForPortsViewVO) GetSupportAnomaly() bool`

GetSupportAnomaly returns the SupportAnomaly field if non-nil, zero value otherwise.

### GetSupportAnomalyOk

`func (o *OswDetailForPortsViewVO) GetSupportAnomalyOk() (*bool, bool)`

GetSupportAnomalyOk returns a tuple with the SupportAnomaly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAnomaly

`func (o *OswDetailForPortsViewVO) SetSupportAnomaly(v bool)`

SetSupportAnomaly sets SupportAnomaly field to given value.

### HasSupportAnomaly

`func (o *OswDetailForPortsViewVO) HasSupportAnomaly() bool`

HasSupportAnomaly returns a boolean if a field has been set.

### GetSupportClearCounters

`func (o *OswDetailForPortsViewVO) GetSupportClearCounters() bool`

GetSupportClearCounters returns the SupportClearCounters field if non-nil, zero value otherwise.

### GetSupportClearCountersOk

`func (o *OswDetailForPortsViewVO) GetSupportClearCountersOk() (*bool, bool)`

GetSupportClearCountersOk returns a tuple with the SupportClearCounters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportClearCounters

`func (o *OswDetailForPortsViewVO) SetSupportClearCounters(v bool)`

SetSupportClearCounters sets SupportClearCounters field to given value.

### HasSupportClearCounters

`func (o *OswDetailForPortsViewVO) HasSupportClearCounters() bool`

HasSupportClearCounters returns a boolean if a field has been set.

### GetSupportLocatePort

`func (o *OswDetailForPortsViewVO) GetSupportLocatePort() bool`

GetSupportLocatePort returns the SupportLocatePort field if non-nil, zero value otherwise.

### GetSupportLocatePortOk

`func (o *OswDetailForPortsViewVO) GetSupportLocatePortOk() (*bool, bool)`

GetSupportLocatePortOk returns a tuple with the SupportLocatePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLocatePort

`func (o *OswDetailForPortsViewVO) SetSupportLocatePort(v bool)`

SetSupportLocatePort sets SupportLocatePort field to given value.

### HasSupportLocatePort

`func (o *OswDetailForPortsViewVO) HasSupportLocatePort() bool`

HasSupportLocatePort returns a boolean if a field has been set.

### GetTemplateId

`func (o *OswDetailForPortsViewVO) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *OswDetailForPortsViewVO) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *OswDetailForPortsViewVO) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *OswDetailForPortsViewVO) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateName

`func (o *OswDetailForPortsViewVO) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *OswDetailForPortsViewVO) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *OswDetailForPortsViewVO) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *OswDetailForPortsViewVO) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetType

`func (o *OswDetailForPortsViewVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswDetailForPortsViewVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswDetailForPortsViewVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OswDetailForPortsViewVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUplink

`func (o *OswDetailForPortsViewVO) GetUplink() OswUplinkVO`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *OswDetailForPortsViewVO) GetUplinkOk() (*OswUplinkVO, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *OswDetailForPortsViewVO) SetUplink(v OswUplinkVO)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *OswDetailForPortsViewVO) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetVersion

`func (o *OswDetailForPortsViewVO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OswDetailForPortsViewVO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OswDetailForPortsViewVO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OswDetailForPortsViewVO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


