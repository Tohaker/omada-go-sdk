# OswStackVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbnormalReason** | Pointer to **int32** | Abnormal Reason | [optional] 
**Clients** | Pointer to **int32** | Clients | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DetectedMember** | Pointer to [**[]OswStackMemberVO**](OswStackMemberVO.md) | List of members that have been detected by the stack | [optional] 
**DevicesNum** | Pointer to **int32** | Number of devices | [optional] 
**DueTime** | Pointer to **int64** | Expiration time | [optional] 
**DueTimeLeft** | Pointer to **int64** | The number of milliseconds from the current time to the expiration time | [optional] 
**FwDownload** | Pointer to **bool** | Indicates whether the upgrade status is downloading | [optional] 
**Id** | Pointer to **string** | Stack ID | [optional] 
**Ip** | Pointer to **string** | IP | [optional] 
**LatestVersion** | Pointer to **string** | Latest Version | [optional] 
**LicenseStatus** | Pointer to **int32** | License Status should be a value as follows: 0: Activated; 1: Unactivated; 2: Not All Activated, 3: Expired | [optional] 
**LocateEnable** | Pointer to **bool** | Indicates whether the locate function is enabled | [optional] 
**Loop** | Pointer to **string** | Loopback port | [optional] 
**LoopbackNum** | Pointer to **int32** | Loopback Num | [optional] 
**MacDelay** | Pointer to [**MacDelayVO**](MacDelayVO.md) |  | [optional] 
**MadSetting** | Pointer to [**MadSettingVO**](MadSettingVO.md) |  | [optional] 
**MasterDeviceActive** | Pointer to **bool** | Marks whether the master device of the stack system is activated | [optional] 
**MasterMac** | Pointer to **string** | Master device mac | [optional] 
**Member** | Pointer to [**[]OswStackMemberVO**](OswStackMemberVO.md) | Stack member list | [optional] 
**Name** | **string** | Stack Name | 
**NeedUpgrade** | Pointer to **bool** | Indicates whether an upgrade is required | [optional] 
**QosConfig** | Pointer to [**OswQosConfigVO**](OswQosConfigVO.md) |  | [optional] 
**SiteId** | Pointer to **string** | Site ID | [optional] 
**SiteName** | Pointer to **string** | Site Name | [optional] 
**Status** | Pointer to **int32** | Stack Status should be a value as follows: 0: normal; 1: abnormal; 2: stack not ready | [optional] 
**SupportMacDelay** | Pointer to **bool** |  | [optional] 
**Tag** | Pointer to **string** | Tag | [optional] 
**TrafficDown** | Pointer to **int64** | Traffic Down | [optional] 
**TrafficUp** | Pointer to **int64** | Traffic Up | [optional] 
**Version** | Pointer to **string** | Version | [optional] 
**VirtualMac** | Pointer to **string** | Virtual Mac | [optional] 

## Methods

### NewOswStackVO

`func NewOswStackVO(name string, ) *OswStackVO`

NewOswStackVO instantiates a new OswStackVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackVOWithDefaults

`func NewOswStackVOWithDefaults() *OswStackVO`

NewOswStackVOWithDefaults instantiates a new OswStackVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbnormalReason

`func (o *OswStackVO) GetAbnormalReason() int32`

GetAbnormalReason returns the AbnormalReason field if non-nil, zero value otherwise.

### GetAbnormalReasonOk

`func (o *OswStackVO) GetAbnormalReasonOk() (*int32, bool)`

GetAbnormalReasonOk returns a tuple with the AbnormalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbnormalReason

`func (o *OswStackVO) SetAbnormalReason(v int32)`

SetAbnormalReason sets AbnormalReason field to given value.

### HasAbnormalReason

`func (o *OswStackVO) HasAbnormalReason() bool`

HasAbnormalReason returns a boolean if a field has been set.

### GetClients

`func (o *OswStackVO) GetClients() int32`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *OswStackVO) GetClientsOk() (*int32, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *OswStackVO) SetClients(v int32)`

SetClients sets Clients field to given value.

### HasClients

`func (o *OswStackVO) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetDescription

`func (o *OswStackVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OswStackVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OswStackVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OswStackVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetectedMember

`func (o *OswStackVO) GetDetectedMember() []OswStackMemberVO`

GetDetectedMember returns the DetectedMember field if non-nil, zero value otherwise.

### GetDetectedMemberOk

`func (o *OswStackVO) GetDetectedMemberOk() (*[]OswStackMemberVO, bool)`

GetDetectedMemberOk returns a tuple with the DetectedMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedMember

`func (o *OswStackVO) SetDetectedMember(v []OswStackMemberVO)`

SetDetectedMember sets DetectedMember field to given value.

### HasDetectedMember

`func (o *OswStackVO) HasDetectedMember() bool`

HasDetectedMember returns a boolean if a field has been set.

### GetDevicesNum

`func (o *OswStackVO) GetDevicesNum() int32`

GetDevicesNum returns the DevicesNum field if non-nil, zero value otherwise.

### GetDevicesNumOk

`func (o *OswStackVO) GetDevicesNumOk() (*int32, bool)`

GetDevicesNumOk returns a tuple with the DevicesNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesNum

`func (o *OswStackVO) SetDevicesNum(v int32)`

SetDevicesNum sets DevicesNum field to given value.

### HasDevicesNum

`func (o *OswStackVO) HasDevicesNum() bool`

HasDevicesNum returns a boolean if a field has been set.

### GetDueTime

`func (o *OswStackVO) GetDueTime() int64`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *OswStackVO) GetDueTimeOk() (*int64, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *OswStackVO) SetDueTime(v int64)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *OswStackVO) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### GetDueTimeLeft

`func (o *OswStackVO) GetDueTimeLeft() int64`

GetDueTimeLeft returns the DueTimeLeft field if non-nil, zero value otherwise.

### GetDueTimeLeftOk

`func (o *OswStackVO) GetDueTimeLeftOk() (*int64, bool)`

GetDueTimeLeftOk returns a tuple with the DueTimeLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTimeLeft

`func (o *OswStackVO) SetDueTimeLeft(v int64)`

SetDueTimeLeft sets DueTimeLeft field to given value.

### HasDueTimeLeft

`func (o *OswStackVO) HasDueTimeLeft() bool`

HasDueTimeLeft returns a boolean if a field has been set.

### GetFwDownload

`func (o *OswStackVO) GetFwDownload() bool`

GetFwDownload returns the FwDownload field if non-nil, zero value otherwise.

### GetFwDownloadOk

`func (o *OswStackVO) GetFwDownloadOk() (*bool, bool)`

GetFwDownloadOk returns a tuple with the FwDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwDownload

`func (o *OswStackVO) SetFwDownload(v bool)`

SetFwDownload sets FwDownload field to given value.

### HasFwDownload

`func (o *OswStackVO) HasFwDownload() bool`

HasFwDownload returns a boolean if a field has been set.

### GetId

`func (o *OswStackVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OswStackVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OswStackVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OswStackVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *OswStackVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OswStackVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OswStackVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OswStackVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLatestVersion

`func (o *OswStackVO) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *OswStackVO) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *OswStackVO) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *OswStackVO) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetLicenseStatus

`func (o *OswStackVO) GetLicenseStatus() int32`

GetLicenseStatus returns the LicenseStatus field if non-nil, zero value otherwise.

### GetLicenseStatusOk

`func (o *OswStackVO) GetLicenseStatusOk() (*int32, bool)`

GetLicenseStatusOk returns a tuple with the LicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatus

`func (o *OswStackVO) SetLicenseStatus(v int32)`

SetLicenseStatus sets LicenseStatus field to given value.

### HasLicenseStatus

`func (o *OswStackVO) HasLicenseStatus() bool`

HasLicenseStatus returns a boolean if a field has been set.

### GetLocateEnable

`func (o *OswStackVO) GetLocateEnable() bool`

GetLocateEnable returns the LocateEnable field if non-nil, zero value otherwise.

### GetLocateEnableOk

`func (o *OswStackVO) GetLocateEnableOk() (*bool, bool)`

GetLocateEnableOk returns a tuple with the LocateEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocateEnable

`func (o *OswStackVO) SetLocateEnable(v bool)`

SetLocateEnable sets LocateEnable field to given value.

### HasLocateEnable

`func (o *OswStackVO) HasLocateEnable() bool`

HasLocateEnable returns a boolean if a field has been set.

### GetLoop

`func (o *OswStackVO) GetLoop() string`

GetLoop returns the Loop field if non-nil, zero value otherwise.

### GetLoopOk

`func (o *OswStackVO) GetLoopOk() (*string, bool)`

GetLoopOk returns a tuple with the Loop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoop

`func (o *OswStackVO) SetLoop(v string)`

SetLoop sets Loop field to given value.

### HasLoop

`func (o *OswStackVO) HasLoop() bool`

HasLoop returns a boolean if a field has been set.

### GetLoopbackNum

`func (o *OswStackVO) GetLoopbackNum() int32`

GetLoopbackNum returns the LoopbackNum field if non-nil, zero value otherwise.

### GetLoopbackNumOk

`func (o *OswStackVO) GetLoopbackNumOk() (*int32, bool)`

GetLoopbackNumOk returns a tuple with the LoopbackNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackNum

`func (o *OswStackVO) SetLoopbackNum(v int32)`

SetLoopbackNum sets LoopbackNum field to given value.

### HasLoopbackNum

`func (o *OswStackVO) HasLoopbackNum() bool`

HasLoopbackNum returns a boolean if a field has been set.

### GetMacDelay

`func (o *OswStackVO) GetMacDelay() MacDelayVO`

GetMacDelay returns the MacDelay field if non-nil, zero value otherwise.

### GetMacDelayOk

`func (o *OswStackVO) GetMacDelayOk() (*MacDelayVO, bool)`

GetMacDelayOk returns a tuple with the MacDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacDelay

`func (o *OswStackVO) SetMacDelay(v MacDelayVO)`

SetMacDelay sets MacDelay field to given value.

### HasMacDelay

`func (o *OswStackVO) HasMacDelay() bool`

HasMacDelay returns a boolean if a field has been set.

### GetMadSetting

`func (o *OswStackVO) GetMadSetting() MadSettingVO`

GetMadSetting returns the MadSetting field if non-nil, zero value otherwise.

### GetMadSettingOk

`func (o *OswStackVO) GetMadSettingOk() (*MadSettingVO, bool)`

GetMadSettingOk returns a tuple with the MadSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMadSetting

`func (o *OswStackVO) SetMadSetting(v MadSettingVO)`

SetMadSetting sets MadSetting field to given value.

### HasMadSetting

`func (o *OswStackVO) HasMadSetting() bool`

HasMadSetting returns a boolean if a field has been set.

### GetMasterDeviceActive

`func (o *OswStackVO) GetMasterDeviceActive() bool`

GetMasterDeviceActive returns the MasterDeviceActive field if non-nil, zero value otherwise.

### GetMasterDeviceActiveOk

`func (o *OswStackVO) GetMasterDeviceActiveOk() (*bool, bool)`

GetMasterDeviceActiveOk returns a tuple with the MasterDeviceActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterDeviceActive

`func (o *OswStackVO) SetMasterDeviceActive(v bool)`

SetMasterDeviceActive sets MasterDeviceActive field to given value.

### HasMasterDeviceActive

`func (o *OswStackVO) HasMasterDeviceActive() bool`

HasMasterDeviceActive returns a boolean if a field has been set.

### GetMasterMac

`func (o *OswStackVO) GetMasterMac() string`

GetMasterMac returns the MasterMac field if non-nil, zero value otherwise.

### GetMasterMacOk

`func (o *OswStackVO) GetMasterMacOk() (*string, bool)`

GetMasterMacOk returns a tuple with the MasterMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterMac

`func (o *OswStackVO) SetMasterMac(v string)`

SetMasterMac sets MasterMac field to given value.

### HasMasterMac

`func (o *OswStackVO) HasMasterMac() bool`

HasMasterMac returns a boolean if a field has been set.

### GetMember

`func (o *OswStackVO) GetMember() []OswStackMemberVO`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *OswStackVO) GetMemberOk() (*[]OswStackMemberVO, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *OswStackVO) SetMember(v []OswStackMemberVO)`

SetMember sets Member field to given value.

### HasMember

`func (o *OswStackVO) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetName

`func (o *OswStackVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswStackVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswStackVO) SetName(v string)`

SetName sets Name field to given value.


### GetNeedUpgrade

`func (o *OswStackVO) GetNeedUpgrade() bool`

GetNeedUpgrade returns the NeedUpgrade field if non-nil, zero value otherwise.

### GetNeedUpgradeOk

`func (o *OswStackVO) GetNeedUpgradeOk() (*bool, bool)`

GetNeedUpgradeOk returns a tuple with the NeedUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedUpgrade

`func (o *OswStackVO) SetNeedUpgrade(v bool)`

SetNeedUpgrade sets NeedUpgrade field to given value.

### HasNeedUpgrade

`func (o *OswStackVO) HasNeedUpgrade() bool`

HasNeedUpgrade returns a boolean if a field has been set.

### GetQosConfig

`func (o *OswStackVO) GetQosConfig() OswQosConfigVO`

GetQosConfig returns the QosConfig field if non-nil, zero value otherwise.

### GetQosConfigOk

`func (o *OswStackVO) GetQosConfigOk() (*OswQosConfigVO, bool)`

GetQosConfigOk returns a tuple with the QosConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosConfig

`func (o *OswStackVO) SetQosConfig(v OswQosConfigVO)`

SetQosConfig sets QosConfig field to given value.

### HasQosConfig

`func (o *OswStackVO) HasQosConfig() bool`

HasQosConfig returns a boolean if a field has been set.

### GetSiteId

`func (o *OswStackVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *OswStackVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *OswStackVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *OswStackVO) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteName

`func (o *OswStackVO) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *OswStackVO) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *OswStackVO) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *OswStackVO) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetStatus

`func (o *OswStackVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswStackVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswStackVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OswStackVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportMacDelay

`func (o *OswStackVO) GetSupportMacDelay() bool`

GetSupportMacDelay returns the SupportMacDelay field if non-nil, zero value otherwise.

### GetSupportMacDelayOk

`func (o *OswStackVO) GetSupportMacDelayOk() (*bool, bool)`

GetSupportMacDelayOk returns a tuple with the SupportMacDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMacDelay

`func (o *OswStackVO) SetSupportMacDelay(v bool)`

SetSupportMacDelay sets SupportMacDelay field to given value.

### HasSupportMacDelay

`func (o *OswStackVO) HasSupportMacDelay() bool`

HasSupportMacDelay returns a boolean if a field has been set.

### GetTag

`func (o *OswStackVO) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *OswStackVO) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *OswStackVO) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *OswStackVO) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTrafficDown

`func (o *OswStackVO) GetTrafficDown() int64`

GetTrafficDown returns the TrafficDown field if non-nil, zero value otherwise.

### GetTrafficDownOk

`func (o *OswStackVO) GetTrafficDownOk() (*int64, bool)`

GetTrafficDownOk returns a tuple with the TrafficDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficDown

`func (o *OswStackVO) SetTrafficDown(v int64)`

SetTrafficDown sets TrafficDown field to given value.

### HasTrafficDown

`func (o *OswStackVO) HasTrafficDown() bool`

HasTrafficDown returns a boolean if a field has been set.

### GetTrafficUp

`func (o *OswStackVO) GetTrafficUp() int64`

GetTrafficUp returns the TrafficUp field if non-nil, zero value otherwise.

### GetTrafficUpOk

`func (o *OswStackVO) GetTrafficUpOk() (*int64, bool)`

GetTrafficUpOk returns a tuple with the TrafficUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficUp

`func (o *OswStackVO) SetTrafficUp(v int64)`

SetTrafficUp sets TrafficUp field to given value.

### HasTrafficUp

`func (o *OswStackVO) HasTrafficUp() bool`

HasTrafficUp returns a boolean if a field has been set.

### GetVersion

`func (o *OswStackVO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OswStackVO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OswStackVO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OswStackVO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVirtualMac

`func (o *OswStackVO) GetVirtualMac() string`

GetVirtualMac returns the VirtualMac field if non-nil, zero value otherwise.

### GetVirtualMacOk

`func (o *OswStackVO) GetVirtualMacOk() (*string, bool)`

GetVirtualMacOk returns a tuple with the VirtualMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMac

`func (o *OswStackVO) SetVirtualMac(v string)`

SetVirtualMac sets VirtualMac field to given value.

### HasVirtualMac

`func (o *OswStackVO) HasVirtualMac() bool`

HasVirtualMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


