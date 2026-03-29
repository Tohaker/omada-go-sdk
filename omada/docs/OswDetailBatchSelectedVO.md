# OswDetailBatchSelectedVO

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
**DueTime** | Pointer to **int64** | Expire timestamp of license(cloud base exclusive) | [optional] 
**DueTimeLeft** | Pointer to **int64** | Milliseconds from the current moment to the expiration time(cloud base exclusive) | [optional] 
**EcspFirstVersion** | Pointer to **int32** | Ecsp first version | [optional] 
**Eos** | Pointer to **int64** | End of support time of device(CBC exclusive) | [optional] 
**Eost** | Pointer to **int64** | End of service time of device(CBC exclusive) | [optional] 
**Es** | Pointer to **bool** | Whether the switch is Agile Series Switch | [optional] 
**FirmwareVersion** | Pointer to **string** | Version of firmware,for example:2.5.0 Build 20190118 Rel. 64821 | [optional] 
**ForgetId** | Pointer to **string** | Forget ID of device | [optional] 
**HwVersion** | Pointer to **string** | Version of hardware,for example 1.0 | [optional] 
**InWhitelist** | Pointer to **bool** | Whether the device is in white list | [optional] 
**InitialUnbindingLimit** | Pointer to **int32** | Initial unbind count for license(cloud base exclusive) | [optional] 
**Lags** | Pointer to [**[]OswLagVO**](OswLagVO.md) | Lag List | [optional] 
**LicenseId** | Pointer to **string** | License key on detail page of device(cloud base exclusive) | [optional] 
**LicenseStatus** | Pointer to **int32** | License status(cloud base exclusive).LicenseStatus should be a value as follows: 0:unActive 1:Unbind 2:Expired 3:active | [optional] 
**LicenseUnbindingLimit** | Pointer to **int32** | Remaining unbind count for license on detail Page of device(cloud base exclusive) | [optional] 
**Location** | Pointer to [**DeviceLocationDetailVO**](DeviceLocationDetailVO.md) |  | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**Model** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**MoveSiteId** | Pointer to **string** | Record that the device is in a moveSite operation; if it is null, then it is not in the moveSite operation. | [optional] 
**Name** | Pointer to **string** | Device name,default value is the mac address of device | [optional] 
**OmadacId** | Pointer to **string** | OmadacId of the device | [optional] 
**Ports** | Pointer to [**[]OswPortVO**](OswPortVO.md) | Port List | [optional] 
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
**StackPorts** | Pointer to [**[]OswStackPortGroupVO**](OswStackPortGroupVO.md) | Stack Port List | [optional] 
**Status** | Pointer to **int32** | Status of device,status should be a value as follows: 0:Disconnected;1:Disconnected(Migrating);10:Provisioning;11:Configuring;12:Upgrading;13:Rebooting;14:Connected;15:Connected(Wireless);16:Connected(Migrating);17:Connected(Wireless,Migrating);20:Pending;21:Pending(Wireless);22:Adopting;23:Adopting(Wireless);24:Adopt Failed;25:Adopt Failed(Wireless);26:Managed By Others;27:Managed By Others(Wireless);30:Heartbeat Missed;31:Heartbeat Missed(Wireless);32:Heartbeat Missed(Migrating);33:Heartbeat Missed(Wireless,Migrating);40:Isolated;41:Isolated(Migrating);50:Slice Configuring | [optional] 
**StatusCategory** | Pointer to **int32** | Category of device status,statusCategory should be a value as follows: 0:Disconnected;1:Connected;2:Pending;3:Heartbeat Missed;4:Isolated | [optional] 
**SupportAnomaly** | Pointer to **bool** | Whether the device firmware support intelligent anomaly detection | [optional] 
**SupportLocatePort** | Pointer to **bool** | Whether the device supports locating port | [optional] 
**TemplateId** | Pointer to **string** | ID of the template bound to the device | [optional] 
**TemplateName** | Pointer to **string** | Name of the template bound to the device | [optional] 
**Type** | Pointer to **string** | Device type:ap、gateway、switch、olt | [optional] 
**Unit** | Pointer to **int32** | Unit ID | [optional] 
**Uplink** | Pointer to [**OswUplinkVO**](OswUplinkVO.md) |  | [optional] 
**Version** | Pointer to **string** | Simplified version of firmware,for example:2.5.0 | [optional] 

## Methods

### NewOswDetailBatchSelectedVO

`func NewOswDetailBatchSelectedVO() *OswDetailBatchSelectedVO`

NewOswDetailBatchSelectedVO instantiates a new OswDetailBatchSelectedVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswDetailBatchSelectedVOWithDefaults

`func NewOswDetailBatchSelectedVOWithDefaults() *OswDetailBatchSelectedVO`

NewOswDetailBatchSelectedVOWithDefaults instantiates a new OswDetailBatchSelectedVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *OswDetailBatchSelectedVO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *OswDetailBatchSelectedVO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *OswDetailBatchSelectedVO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *OswDetailBatchSelectedVO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetActive

`func (o *OswDetailBatchSelectedVO) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OswDetailBatchSelectedVO) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OswDetailBatchSelectedVO) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *OswDetailBatchSelectedVO) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddedInAdvanced

`func (o *OswDetailBatchSelectedVO) GetAddedInAdvanced() bool`

GetAddedInAdvanced returns the AddedInAdvanced field if non-nil, zero value otherwise.

### GetAddedInAdvancedOk

`func (o *OswDetailBatchSelectedVO) GetAddedInAdvancedOk() (*bool, bool)`

GetAddedInAdvancedOk returns a tuple with the AddedInAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedInAdvanced

`func (o *OswDetailBatchSelectedVO) SetAddedInAdvanced(v bool)`

SetAddedInAdvanced sets AddedInAdvanced field to given value.

### HasAddedInAdvanced

`func (o *OswDetailBatchSelectedVO) HasAddedInAdvanced() bool`

HasAddedInAdvanced returns a boolean if a field has been set.

### GetBoundDeviceTemplate

`func (o *OswDetailBatchSelectedVO) GetBoundDeviceTemplate() bool`

GetBoundDeviceTemplate returns the BoundDeviceTemplate field if non-nil, zero value otherwise.

### GetBoundDeviceTemplateOk

`func (o *OswDetailBatchSelectedVO) GetBoundDeviceTemplateOk() (*bool, bool)`

GetBoundDeviceTemplateOk returns a tuple with the BoundDeviceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundDeviceTemplate

`func (o *OswDetailBatchSelectedVO) SetBoundDeviceTemplate(v bool)`

SetBoundDeviceTemplate sets BoundDeviceTemplate field to given value.

### HasBoundDeviceTemplate

`func (o *OswDetailBatchSelectedVO) HasBoundDeviceTemplate() bool`

HasBoundDeviceTemplate returns a boolean if a field has been set.

### GetBoundSiteTemplate

`func (o *OswDetailBatchSelectedVO) GetBoundSiteTemplate() bool`

GetBoundSiteTemplate returns the BoundSiteTemplate field if non-nil, zero value otherwise.

### GetBoundSiteTemplateOk

`func (o *OswDetailBatchSelectedVO) GetBoundSiteTemplateOk() (*bool, bool)`

GetBoundSiteTemplateOk returns a tuple with the BoundSiteTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSiteTemplate

`func (o *OswDetailBatchSelectedVO) SetBoundSiteTemplate(v bool)`

SetBoundSiteTemplate sets BoundSiteTemplate field to given value.

### HasBoundSiteTemplate

`func (o *OswDetailBatchSelectedVO) HasBoundSiteTemplate() bool`

HasBoundSiteTemplate returns a boolean if a field has been set.

### GetCategory

`func (o *OswDetailBatchSelectedVO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *OswDetailBatchSelectedVO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *OswDetailBatchSelectedVO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *OswDetailBatchSelectedVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCompatible

`func (o *OswDetailBatchSelectedVO) GetCompatible() int32`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *OswDetailBatchSelectedVO) GetCompatibleOk() (*int32, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *OswDetailBatchSelectedVO) SetCompatible(v int32)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *OswDetailBatchSelectedVO) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetCompoundModel

`func (o *OswDetailBatchSelectedVO) GetCompoundModel() string`

GetCompoundModel returns the CompoundModel field if non-nil, zero value otherwise.

### GetCompoundModelOk

`func (o *OswDetailBatchSelectedVO) GetCompoundModelOk() (*string, bool)`

GetCompoundModelOk returns a tuple with the CompoundModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompoundModel

`func (o *OswDetailBatchSelectedVO) SetCompoundModel(v string)`

SetCompoundModel sets CompoundModel field to given value.

### HasCompoundModel

`func (o *OswDetailBatchSelectedVO) HasCompoundModel() bool`

HasCompoundModel returns a boolean if a field has been set.

### GetCustomId

`func (o *OswDetailBatchSelectedVO) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *OswDetailBatchSelectedVO) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *OswDetailBatchSelectedVO) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.

### HasCustomId

`func (o *OswDetailBatchSelectedVO) HasCustomId() bool`

HasCustomId returns a boolean if a field has been set.

### GetCustomName

`func (o *OswDetailBatchSelectedVO) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *OswDetailBatchSelectedVO) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *OswDetailBatchSelectedVO) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *OswDetailBatchSelectedVO) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### GetDescription

`func (o *OswDetailBatchSelectedVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OswDetailBatchSelectedVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OswDetailBatchSelectedVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OswDetailBatchSelectedVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevCap

`func (o *OswDetailBatchSelectedVO) GetDevCap() OswDevCapVO`

GetDevCap returns the DevCap field if non-nil, zero value otherwise.

### GetDevCapOk

`func (o *OswDetailBatchSelectedVO) GetDevCapOk() (*OswDevCapVO, bool)`

GetDevCapOk returns a tuple with the DevCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevCap

`func (o *OswDetailBatchSelectedVO) SetDevCap(v OswDevCapVO)`

SetDevCap sets DevCap field to given value.

### HasDevCap

`func (o *OswDetailBatchSelectedVO) HasDevCap() bool`

HasDevCap returns a boolean if a field has been set.

### GetDeviceMisc

`func (o *OswDetailBatchSelectedVO) GetDeviceMisc() OswDeviceMiscVO`

GetDeviceMisc returns the DeviceMisc field if non-nil, zero value otherwise.

### GetDeviceMiscOk

`func (o *OswDetailBatchSelectedVO) GetDeviceMiscOk() (*OswDeviceMiscVO, bool)`

GetDeviceMiscOk returns a tuple with the DeviceMisc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMisc

`func (o *OswDetailBatchSelectedVO) SetDeviceMisc(v OswDeviceMiscVO)`

SetDeviceMisc sets DeviceMisc field to given value.

### HasDeviceMisc

`func (o *OswDetailBatchSelectedVO) HasDeviceMisc() bool`

HasDeviceMisc returns a boolean if a field has been set.

### GetDeviceSeriesType

`func (o *OswDetailBatchSelectedVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *OswDetailBatchSelectedVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *OswDetailBatchSelectedVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *OswDetailBatchSelectedVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetDeviceTemplateAvailable

`func (o *OswDetailBatchSelectedVO) GetDeviceTemplateAvailable() bool`

GetDeviceTemplateAvailable returns the DeviceTemplateAvailable field if non-nil, zero value otherwise.

### GetDeviceTemplateAvailableOk

`func (o *OswDetailBatchSelectedVO) GetDeviceTemplateAvailableOk() (*bool, bool)`

GetDeviceTemplateAvailableOk returns a tuple with the DeviceTemplateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTemplateAvailable

`func (o *OswDetailBatchSelectedVO) SetDeviceTemplateAvailable(v bool)`

SetDeviceTemplateAvailable sets DeviceTemplateAvailable field to given value.

### HasDeviceTemplateAvailable

`func (o *OswDetailBatchSelectedVO) HasDeviceTemplateAvailable() bool`

HasDeviceTemplateAvailable returns a boolean if a field has been set.

### GetDisableHwReset

`func (o *OswDetailBatchSelectedVO) GetDisableHwReset() bool`

GetDisableHwReset returns the DisableHwReset field if non-nil, zero value otherwise.

### GetDisableHwResetOk

`func (o *OswDetailBatchSelectedVO) GetDisableHwResetOk() (*bool, bool)`

GetDisableHwResetOk returns a tuple with the DisableHwReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHwReset

`func (o *OswDetailBatchSelectedVO) SetDisableHwReset(v bool)`

SetDisableHwReset sets DisableHwReset field to given value.

### HasDisableHwReset

`func (o *OswDetailBatchSelectedVO) HasDisableHwReset() bool`

HasDisableHwReset returns a boolean if a field has been set.

### GetDueTime

`func (o *OswDetailBatchSelectedVO) GetDueTime() int64`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *OswDetailBatchSelectedVO) GetDueTimeOk() (*int64, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *OswDetailBatchSelectedVO) SetDueTime(v int64)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *OswDetailBatchSelectedVO) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### GetDueTimeLeft

`func (o *OswDetailBatchSelectedVO) GetDueTimeLeft() int64`

GetDueTimeLeft returns the DueTimeLeft field if non-nil, zero value otherwise.

### GetDueTimeLeftOk

`func (o *OswDetailBatchSelectedVO) GetDueTimeLeftOk() (*int64, bool)`

GetDueTimeLeftOk returns a tuple with the DueTimeLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTimeLeft

`func (o *OswDetailBatchSelectedVO) SetDueTimeLeft(v int64)`

SetDueTimeLeft sets DueTimeLeft field to given value.

### HasDueTimeLeft

`func (o *OswDetailBatchSelectedVO) HasDueTimeLeft() bool`

HasDueTimeLeft returns a boolean if a field has been set.

### GetEcspFirstVersion

`func (o *OswDetailBatchSelectedVO) GetEcspFirstVersion() int32`

GetEcspFirstVersion returns the EcspFirstVersion field if non-nil, zero value otherwise.

### GetEcspFirstVersionOk

`func (o *OswDetailBatchSelectedVO) GetEcspFirstVersionOk() (*int32, bool)`

GetEcspFirstVersionOk returns a tuple with the EcspFirstVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcspFirstVersion

`func (o *OswDetailBatchSelectedVO) SetEcspFirstVersion(v int32)`

SetEcspFirstVersion sets EcspFirstVersion field to given value.

### HasEcspFirstVersion

`func (o *OswDetailBatchSelectedVO) HasEcspFirstVersion() bool`

HasEcspFirstVersion returns a boolean if a field has been set.

### GetEos

`func (o *OswDetailBatchSelectedVO) GetEos() int64`

GetEos returns the Eos field if non-nil, zero value otherwise.

### GetEosOk

`func (o *OswDetailBatchSelectedVO) GetEosOk() (*int64, bool)`

GetEosOk returns a tuple with the Eos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEos

`func (o *OswDetailBatchSelectedVO) SetEos(v int64)`

SetEos sets Eos field to given value.

### HasEos

`func (o *OswDetailBatchSelectedVO) HasEos() bool`

HasEos returns a boolean if a field has been set.

### GetEost

`func (o *OswDetailBatchSelectedVO) GetEost() int64`

GetEost returns the Eost field if non-nil, zero value otherwise.

### GetEostOk

`func (o *OswDetailBatchSelectedVO) GetEostOk() (*int64, bool)`

GetEostOk returns a tuple with the Eost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEost

`func (o *OswDetailBatchSelectedVO) SetEost(v int64)`

SetEost sets Eost field to given value.

### HasEost

`func (o *OswDetailBatchSelectedVO) HasEost() bool`

HasEost returns a boolean if a field has been set.

### GetEs

`func (o *OswDetailBatchSelectedVO) GetEs() bool`

GetEs returns the Es field if non-nil, zero value otherwise.

### GetEsOk

`func (o *OswDetailBatchSelectedVO) GetEsOk() (*bool, bool)`

GetEsOk returns a tuple with the Es field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEs

`func (o *OswDetailBatchSelectedVO) SetEs(v bool)`

SetEs sets Es field to given value.

### HasEs

`func (o *OswDetailBatchSelectedVO) HasEs() bool`

HasEs returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *OswDetailBatchSelectedVO) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *OswDetailBatchSelectedVO) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *OswDetailBatchSelectedVO) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *OswDetailBatchSelectedVO) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetForgetId

`func (o *OswDetailBatchSelectedVO) GetForgetId() string`

GetForgetId returns the ForgetId field if non-nil, zero value otherwise.

### GetForgetIdOk

`func (o *OswDetailBatchSelectedVO) GetForgetIdOk() (*string, bool)`

GetForgetIdOk returns a tuple with the ForgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgetId

`func (o *OswDetailBatchSelectedVO) SetForgetId(v string)`

SetForgetId sets ForgetId field to given value.

### HasForgetId

`func (o *OswDetailBatchSelectedVO) HasForgetId() bool`

HasForgetId returns a boolean if a field has been set.

### GetHwVersion

`func (o *OswDetailBatchSelectedVO) GetHwVersion() string`

GetHwVersion returns the HwVersion field if non-nil, zero value otherwise.

### GetHwVersionOk

`func (o *OswDetailBatchSelectedVO) GetHwVersionOk() (*string, bool)`

GetHwVersionOk returns a tuple with the HwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwVersion

`func (o *OswDetailBatchSelectedVO) SetHwVersion(v string)`

SetHwVersion sets HwVersion field to given value.

### HasHwVersion

`func (o *OswDetailBatchSelectedVO) HasHwVersion() bool`

HasHwVersion returns a boolean if a field has been set.

### GetInWhitelist

`func (o *OswDetailBatchSelectedVO) GetInWhitelist() bool`

GetInWhitelist returns the InWhitelist field if non-nil, zero value otherwise.

### GetInWhitelistOk

`func (o *OswDetailBatchSelectedVO) GetInWhitelistOk() (*bool, bool)`

GetInWhitelistOk returns a tuple with the InWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInWhitelist

`func (o *OswDetailBatchSelectedVO) SetInWhitelist(v bool)`

SetInWhitelist sets InWhitelist field to given value.

### HasInWhitelist

`func (o *OswDetailBatchSelectedVO) HasInWhitelist() bool`

HasInWhitelist returns a boolean if a field has been set.

### GetInitialUnbindingLimit

`func (o *OswDetailBatchSelectedVO) GetInitialUnbindingLimit() int32`

GetInitialUnbindingLimit returns the InitialUnbindingLimit field if non-nil, zero value otherwise.

### GetInitialUnbindingLimitOk

`func (o *OswDetailBatchSelectedVO) GetInitialUnbindingLimitOk() (*int32, bool)`

GetInitialUnbindingLimitOk returns a tuple with the InitialUnbindingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialUnbindingLimit

`func (o *OswDetailBatchSelectedVO) SetInitialUnbindingLimit(v int32)`

SetInitialUnbindingLimit sets InitialUnbindingLimit field to given value.

### HasInitialUnbindingLimit

`func (o *OswDetailBatchSelectedVO) HasInitialUnbindingLimit() bool`

HasInitialUnbindingLimit returns a boolean if a field has been set.

### GetLags

`func (o *OswDetailBatchSelectedVO) GetLags() []OswLagVO`

GetLags returns the Lags field if non-nil, zero value otherwise.

### GetLagsOk

`func (o *OswDetailBatchSelectedVO) GetLagsOk() (*[]OswLagVO, bool)`

GetLagsOk returns a tuple with the Lags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLags

`func (o *OswDetailBatchSelectedVO) SetLags(v []OswLagVO)`

SetLags sets Lags field to given value.

### HasLags

`func (o *OswDetailBatchSelectedVO) HasLags() bool`

HasLags returns a boolean if a field has been set.

### GetLicenseId

`func (o *OswDetailBatchSelectedVO) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *OswDetailBatchSelectedVO) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *OswDetailBatchSelectedVO) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *OswDetailBatchSelectedVO) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetLicenseStatus

`func (o *OswDetailBatchSelectedVO) GetLicenseStatus() int32`

GetLicenseStatus returns the LicenseStatus field if non-nil, zero value otherwise.

### GetLicenseStatusOk

`func (o *OswDetailBatchSelectedVO) GetLicenseStatusOk() (*int32, bool)`

GetLicenseStatusOk returns a tuple with the LicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatus

`func (o *OswDetailBatchSelectedVO) SetLicenseStatus(v int32)`

SetLicenseStatus sets LicenseStatus field to given value.

### HasLicenseStatus

`func (o *OswDetailBatchSelectedVO) HasLicenseStatus() bool`

HasLicenseStatus returns a boolean if a field has been set.

### GetLicenseUnbindingLimit

`func (o *OswDetailBatchSelectedVO) GetLicenseUnbindingLimit() int32`

GetLicenseUnbindingLimit returns the LicenseUnbindingLimit field if non-nil, zero value otherwise.

### GetLicenseUnbindingLimitOk

`func (o *OswDetailBatchSelectedVO) GetLicenseUnbindingLimitOk() (*int32, bool)`

GetLicenseUnbindingLimitOk returns a tuple with the LicenseUnbindingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUnbindingLimit

`func (o *OswDetailBatchSelectedVO) SetLicenseUnbindingLimit(v int32)`

SetLicenseUnbindingLimit sets LicenseUnbindingLimit field to given value.

### HasLicenseUnbindingLimit

`func (o *OswDetailBatchSelectedVO) HasLicenseUnbindingLimit() bool`

HasLicenseUnbindingLimit returns a boolean if a field has been set.

### GetLocation

`func (o *OswDetailBatchSelectedVO) GetLocation() DeviceLocationDetailVO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OswDetailBatchSelectedVO) GetLocationOk() (*DeviceLocationDetailVO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OswDetailBatchSelectedVO) SetLocation(v DeviceLocationDetailVO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OswDetailBatchSelectedVO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMac

`func (o *OswDetailBatchSelectedVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswDetailBatchSelectedVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswDetailBatchSelectedVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswDetailBatchSelectedVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *OswDetailBatchSelectedVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OswDetailBatchSelectedVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OswDetailBatchSelectedVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OswDetailBatchSelectedVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OswDetailBatchSelectedVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OswDetailBatchSelectedVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OswDetailBatchSelectedVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OswDetailBatchSelectedVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetMoveSiteId

`func (o *OswDetailBatchSelectedVO) GetMoveSiteId() string`

GetMoveSiteId returns the MoveSiteId field if non-nil, zero value otherwise.

### GetMoveSiteIdOk

`func (o *OswDetailBatchSelectedVO) GetMoveSiteIdOk() (*string, bool)`

GetMoveSiteIdOk returns a tuple with the MoveSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveSiteId

`func (o *OswDetailBatchSelectedVO) SetMoveSiteId(v string)`

SetMoveSiteId sets MoveSiteId field to given value.

### HasMoveSiteId

`func (o *OswDetailBatchSelectedVO) HasMoveSiteId() bool`

HasMoveSiteId returns a boolean if a field has been set.

### GetName

`func (o *OswDetailBatchSelectedVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswDetailBatchSelectedVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswDetailBatchSelectedVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswDetailBatchSelectedVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOmadacId

`func (o *OswDetailBatchSelectedVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *OswDetailBatchSelectedVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *OswDetailBatchSelectedVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *OswDetailBatchSelectedVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetPorts

`func (o *OswDetailBatchSelectedVO) GetPorts() []OswPortVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OswDetailBatchSelectedVO) GetPortsOk() (*[]OswPortVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OswDetailBatchSelectedVO) SetPorts(v []OswPortVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OswDetailBatchSelectedVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetRemember

`func (o *OswDetailBatchSelectedVO) GetRemember() bool`

GetRemember returns the Remember field if non-nil, zero value otherwise.

### GetRememberOk

`func (o *OswDetailBatchSelectedVO) GetRememberOk() (*bool, bool)`

GetRememberOk returns a tuple with the Remember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemember

`func (o *OswDetailBatchSelectedVO) SetRemember(v bool)`

SetRemember sets Remember field to given value.

### HasRemember

`func (o *OswDetailBatchSelectedVO) HasRemember() bool`

HasRemember returns a boolean if a field has been set.

### GetRememberDevice

`func (o *OswDetailBatchSelectedVO) GetRememberDevice() int32`

GetRememberDevice returns the RememberDevice field if non-nil, zero value otherwise.

### GetRememberDeviceOk

`func (o *OswDetailBatchSelectedVO) GetRememberDeviceOk() (*int32, bool)`

GetRememberDeviceOk returns a tuple with the RememberDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberDevice

`func (o *OswDetailBatchSelectedVO) SetRememberDevice(v int32)`

SetRememberDevice sets RememberDevice field to given value.

### HasRememberDevice

`func (o *OswDetailBatchSelectedVO) HasRememberDevice() bool`

HasRememberDevice returns a boolean if a field has been set.

### GetResource

`func (o *OswDetailBatchSelectedVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *OswDetailBatchSelectedVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *OswDetailBatchSelectedVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *OswDetailBatchSelectedVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetShowModel

`func (o *OswDetailBatchSelectedVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *OswDetailBatchSelectedVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *OswDetailBatchSelectedVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *OswDetailBatchSelectedVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetSiteName

`func (o *OswDetailBatchSelectedVO) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *OswDetailBatchSelectedVO) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *OswDetailBatchSelectedVO) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *OswDetailBatchSelectedVO) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSiteTemplateId

`func (o *OswDetailBatchSelectedVO) GetSiteTemplateId() string`

GetSiteTemplateId returns the SiteTemplateId field if non-nil, zero value otherwise.

### GetSiteTemplateIdOk

`func (o *OswDetailBatchSelectedVO) GetSiteTemplateIdOk() (*string, bool)`

GetSiteTemplateIdOk returns a tuple with the SiteTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTemplateId

`func (o *OswDetailBatchSelectedVO) SetSiteTemplateId(v string)`

SetSiteTemplateId sets SiteTemplateId field to given value.

### HasSiteTemplateId

`func (o *OswDetailBatchSelectedVO) HasSiteTemplateId() bool`

HasSiteTemplateId returns a boolean if a field has been set.

### GetSiteTemplateName

`func (o *OswDetailBatchSelectedVO) GetSiteTemplateName() string`

GetSiteTemplateName returns the SiteTemplateName field if non-nil, zero value otherwise.

### GetSiteTemplateNameOk

`func (o *OswDetailBatchSelectedVO) GetSiteTemplateNameOk() (*string, bool)`

GetSiteTemplateNameOk returns a tuple with the SiteTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTemplateName

`func (o *OswDetailBatchSelectedVO) SetSiteTemplateName(v string)`

SetSiteTemplateName sets SiteTemplateName field to given value.

### HasSiteTemplateName

`func (o *OswDetailBatchSelectedVO) HasSiteTemplateName() bool`

HasSiteTemplateName returns a boolean if a field has been set.

### GetSn

`func (o *OswDetailBatchSelectedVO) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *OswDetailBatchSelectedVO) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *OswDetailBatchSelectedVO) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *OswDetailBatchSelectedVO) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetSpecialModel

`func (o *OswDetailBatchSelectedVO) GetSpecialModel() string`

GetSpecialModel returns the SpecialModel field if non-nil, zero value otherwise.

### GetSpecialModelOk

`func (o *OswDetailBatchSelectedVO) GetSpecialModelOk() (*string, bool)`

GetSpecialModelOk returns a tuple with the SpecialModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialModel

`func (o *OswDetailBatchSelectedVO) SetSpecialModel(v string)`

SetSpecialModel sets SpecialModel field to given value.

### HasSpecialModel

`func (o *OswDetailBatchSelectedVO) HasSpecialModel() bool`

HasSpecialModel returns a boolean if a field has been set.

### GetStackMsg

`func (o *OswDetailBatchSelectedVO) GetStackMsg() StackMsgVO`

GetStackMsg returns the StackMsg field if non-nil, zero value otherwise.

### GetStackMsgOk

`func (o *OswDetailBatchSelectedVO) GetStackMsgOk() (*StackMsgVO, bool)`

GetStackMsgOk returns a tuple with the StackMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackMsg

`func (o *OswDetailBatchSelectedVO) SetStackMsg(v StackMsgVO)`

SetStackMsg sets StackMsg field to given value.

### HasStackMsg

`func (o *OswDetailBatchSelectedVO) HasStackMsg() bool`

HasStackMsg returns a boolean if a field has been set.

### GetStackPorts

`func (o *OswDetailBatchSelectedVO) GetStackPorts() []OswStackPortGroupVO`

GetStackPorts returns the StackPorts field if non-nil, zero value otherwise.

### GetStackPortsOk

`func (o *OswDetailBatchSelectedVO) GetStackPortsOk() (*[]OswStackPortGroupVO, bool)`

GetStackPortsOk returns a tuple with the StackPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPorts

`func (o *OswDetailBatchSelectedVO) SetStackPorts(v []OswStackPortGroupVO)`

SetStackPorts sets StackPorts field to given value.

### HasStackPorts

`func (o *OswDetailBatchSelectedVO) HasStackPorts() bool`

HasStackPorts returns a boolean if a field has been set.

### GetStatus

`func (o *OswDetailBatchSelectedVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswDetailBatchSelectedVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswDetailBatchSelectedVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OswDetailBatchSelectedVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *OswDetailBatchSelectedVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *OswDetailBatchSelectedVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *OswDetailBatchSelectedVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *OswDetailBatchSelectedVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetSupportAnomaly

`func (o *OswDetailBatchSelectedVO) GetSupportAnomaly() bool`

GetSupportAnomaly returns the SupportAnomaly field if non-nil, zero value otherwise.

### GetSupportAnomalyOk

`func (o *OswDetailBatchSelectedVO) GetSupportAnomalyOk() (*bool, bool)`

GetSupportAnomalyOk returns a tuple with the SupportAnomaly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAnomaly

`func (o *OswDetailBatchSelectedVO) SetSupportAnomaly(v bool)`

SetSupportAnomaly sets SupportAnomaly field to given value.

### HasSupportAnomaly

`func (o *OswDetailBatchSelectedVO) HasSupportAnomaly() bool`

HasSupportAnomaly returns a boolean if a field has been set.

### GetSupportLocatePort

`func (o *OswDetailBatchSelectedVO) GetSupportLocatePort() bool`

GetSupportLocatePort returns the SupportLocatePort field if non-nil, zero value otherwise.

### GetSupportLocatePortOk

`func (o *OswDetailBatchSelectedVO) GetSupportLocatePortOk() (*bool, bool)`

GetSupportLocatePortOk returns a tuple with the SupportLocatePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLocatePort

`func (o *OswDetailBatchSelectedVO) SetSupportLocatePort(v bool)`

SetSupportLocatePort sets SupportLocatePort field to given value.

### HasSupportLocatePort

`func (o *OswDetailBatchSelectedVO) HasSupportLocatePort() bool`

HasSupportLocatePort returns a boolean if a field has been set.

### GetTemplateId

`func (o *OswDetailBatchSelectedVO) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *OswDetailBatchSelectedVO) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *OswDetailBatchSelectedVO) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *OswDetailBatchSelectedVO) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateName

`func (o *OswDetailBatchSelectedVO) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *OswDetailBatchSelectedVO) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *OswDetailBatchSelectedVO) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *OswDetailBatchSelectedVO) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetType

`func (o *OswDetailBatchSelectedVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswDetailBatchSelectedVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswDetailBatchSelectedVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OswDetailBatchSelectedVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *OswDetailBatchSelectedVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OswDetailBatchSelectedVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OswDetailBatchSelectedVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *OswDetailBatchSelectedVO) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUplink

`func (o *OswDetailBatchSelectedVO) GetUplink() OswUplinkVO`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *OswDetailBatchSelectedVO) GetUplinkOk() (*OswUplinkVO, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *OswDetailBatchSelectedVO) SetUplink(v OswUplinkVO)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *OswDetailBatchSelectedVO) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetVersion

`func (o *OswDetailBatchSelectedVO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OswDetailBatchSelectedVO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OswDetailBatchSelectedVO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OswDetailBatchSelectedVO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


