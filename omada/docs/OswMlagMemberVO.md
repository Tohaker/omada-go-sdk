# OswMlagMemberVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | whether to active the device(cloud base exclusive). | [optional] 
**AddedInAdvanced** | Pointer to **bool** |  | [optional] 
**Category** | Pointer to **string** | Category of license. | [optional] 
**CfgCheck** | Pointer to **int32** | MLAG group configuration check status should be a value as follows: 0: All cfg pass; 1: Type2 cfg not pass;2: Type1 and Type2 not pass;3: Type1 not pass | [optional] 
**Compatible** | Pointer to **int32** | Device firmware and controller compatibility type.Compatible should be a value as follows: 0:COMPATIBLE;1:HIGH_MAJOR_VER;2:LOW_MAJOR_VER;3:HIGH_MINOR_VER;4:LOW_MINOR_VER;7:HIGH_COMPONENT_VER;10:DEVICE_NOT_COMPATIBLE;11:HIGH_ADOPT_COMMPONENT;12:DEVICE_CATEGORY_NOT_COMPATIBLE;14:DEVICE_NOT_COMPATIBLE_IN_CLUSTER. | [optional] 
**CompoundModel** | Pointer to **string** | Model complex used in the backend. | [optional] 
**DadEnable** | Pointer to **bool** | Whether the DAD enable. | [optional] 
**DadIp** | Pointer to **string** | DAD IP Address | [optional] 
**DadLinkPorts** | Pointer to **[]string** | DAD Link Ports | [optional] 
**DadLocalIp** | Pointer to **string** | DAD Local IP | [optional] 
**DadLocalIpv6** | Pointer to **string** | DAD Local IPv6 | [optional] 
**DadPeerIp** | Pointer to **string** | DAD Peer IP | [optional] 
**DadPeerIpv6** | Pointer to **string** | DAD Peer IPv6 | [optional] 
**DadStatus** | Pointer to **int32** | DAD status should be a value as follows: 0: PEER_UNDETECTED; 1: NORMAL; 2: DAD_DISABLE. | [optional] 
**DeviceType** | Pointer to **int32** | Device type, 1: Gateway; 2: Switch; 3: Ap. | [optional] 
**DueTime** | Pointer to **int64** | Expire timestamp of license(cloud base exclusive). | [optional] 
**DueTimeLeft** | Pointer to **int64** | Milliseconds from the current moment to the expiration time(cloud base exclusive) | [optional] 
**Eos** | Pointer to **int64** | End of support time of device(CBC exclusive). | [optional] 
**Eost** | Pointer to **int64** | End of service time of device(CBC exclusive). | [optional] 
**FirmwareVersion** | Pointer to **string** | Version of firmware,for example:2.5.0 Build 20190118 Rel. 64821. | [optional] 
**HwVersion** | Pointer to **string** | Version of hardware,for example 1.0. | [optional] 
**InWhiteList** | Pointer to **bool** | Whether the device is in white list. | [optional] 
**Ip** | Pointer to **string** | IP | [optional] 
**Ipv6** | Pointer to **[]string** | Device IPv6 list. | [optional] 
**KaStatus** | Pointer to **int32** | Keep Alive status should be a value as follows: 0: HEARTBEAT_MISSED; 1: HEARTBEAT_NORMAL. | [optional] 
**LicenseStatus** | Pointer to **int32** | License status(cloud base exclusive).LicenseStatus should be a value as follows: 0:unActive 1:Unbind 2:Expired 3:active. | [optional] 
**LicenseUnbindingLimit** | Pointer to **int32** | Remaining unbind count for license on detail Page of device(cloud base exclusive). | [optional] 
**LocateEnable** | Pointer to **bool** | Whether the locate function is enabled. | [optional] 
**Mac** | Pointer to **string** | Device Mac | [optional] 
**MlagGroupId** | Pointer to **int32** | M-LAG group ID. | [optional] 
**MlagVersion** | Pointer to **string** | M-LAG version. | [optional] 
**Model** | Pointer to **string** | Model of device,for example:EAP225. | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**Name** | Pointer to **string** | Device Name | [optional] 
**PeerLinkPorts** | Pointer to **[]string** | Peer Link Ports | [optional] 
**Ports** | Pointer to [**[]OswMlagPortVO**](OswMlagPortVO.md) | M-LAG group device ports. | [optional] 
**Priority** | Pointer to **int32** | Priority of the device in the M-LAG group. | [optional] 
**PublicIp** | Pointer to **string** | Device public IP. | [optional] 
**Role** | Pointer to **int32** | Role | [optional] 
**ShowModel** | Pointer to **string** | Model complex shown in the front end. | [optional] 
**SpecialModel** | Pointer to **string** | Special device model,for example:EAP225-Outdoor-1a20a950b8d950e8. | [optional] 
**Status** | Pointer to **int32** | Device Status | [optional] 
**StatusCategory** | Pointer to **int32** | Category of device status,statusCategory should be a value as follows: 0:Disconnected;1:Connected;2:Pending;3:Heartbeat Missed;4:Isolated | [optional] 
**SupportMlagIpv6** | Pointer to **bool** | Whether the device support the configuration of M-LAG group IPv6. | [optional] 
**SupportPeerLinkPorts** | Pointer to [**[]OswStandPortVO**](OswStandPortVO.md) | Support the configuration of peer link ports for devices in the MLAG group. | [optional] 
**Uptime** | Pointer to **string** | Device uptime. | [optional] 
**Version** | Pointer to **string** | Simplified version of firmware,for example:2.5.0. | [optional] 

## Methods

### NewOswMlagMemberVO

`func NewOswMlagMemberVO() *OswMlagMemberVO`

NewOswMlagMemberVO instantiates a new OswMlagMemberVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswMlagMemberVOWithDefaults

`func NewOswMlagMemberVOWithDefaults() *OswMlagMemberVO`

NewOswMlagMemberVOWithDefaults instantiates a new OswMlagMemberVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *OswMlagMemberVO) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OswMlagMemberVO) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OswMlagMemberVO) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *OswMlagMemberVO) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddedInAdvanced

`func (o *OswMlagMemberVO) GetAddedInAdvanced() bool`

GetAddedInAdvanced returns the AddedInAdvanced field if non-nil, zero value otherwise.

### GetAddedInAdvancedOk

`func (o *OswMlagMemberVO) GetAddedInAdvancedOk() (*bool, bool)`

GetAddedInAdvancedOk returns a tuple with the AddedInAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedInAdvanced

`func (o *OswMlagMemberVO) SetAddedInAdvanced(v bool)`

SetAddedInAdvanced sets AddedInAdvanced field to given value.

### HasAddedInAdvanced

`func (o *OswMlagMemberVO) HasAddedInAdvanced() bool`

HasAddedInAdvanced returns a boolean if a field has been set.

### GetCategory

`func (o *OswMlagMemberVO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *OswMlagMemberVO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *OswMlagMemberVO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *OswMlagMemberVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCfgCheck

`func (o *OswMlagMemberVO) GetCfgCheck() int32`

GetCfgCheck returns the CfgCheck field if non-nil, zero value otherwise.

### GetCfgCheckOk

`func (o *OswMlagMemberVO) GetCfgCheckOk() (*int32, bool)`

GetCfgCheckOk returns a tuple with the CfgCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfgCheck

`func (o *OswMlagMemberVO) SetCfgCheck(v int32)`

SetCfgCheck sets CfgCheck field to given value.

### HasCfgCheck

`func (o *OswMlagMemberVO) HasCfgCheck() bool`

HasCfgCheck returns a boolean if a field has been set.

### GetCompatible

`func (o *OswMlagMemberVO) GetCompatible() int32`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *OswMlagMemberVO) GetCompatibleOk() (*int32, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *OswMlagMemberVO) SetCompatible(v int32)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *OswMlagMemberVO) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetCompoundModel

`func (o *OswMlagMemberVO) GetCompoundModel() string`

GetCompoundModel returns the CompoundModel field if non-nil, zero value otherwise.

### GetCompoundModelOk

`func (o *OswMlagMemberVO) GetCompoundModelOk() (*string, bool)`

GetCompoundModelOk returns a tuple with the CompoundModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompoundModel

`func (o *OswMlagMemberVO) SetCompoundModel(v string)`

SetCompoundModel sets CompoundModel field to given value.

### HasCompoundModel

`func (o *OswMlagMemberVO) HasCompoundModel() bool`

HasCompoundModel returns a boolean if a field has been set.

### GetDadEnable

`func (o *OswMlagMemberVO) GetDadEnable() bool`

GetDadEnable returns the DadEnable field if non-nil, zero value otherwise.

### GetDadEnableOk

`func (o *OswMlagMemberVO) GetDadEnableOk() (*bool, bool)`

GetDadEnableOk returns a tuple with the DadEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDadEnable

`func (o *OswMlagMemberVO) SetDadEnable(v bool)`

SetDadEnable sets DadEnable field to given value.

### HasDadEnable

`func (o *OswMlagMemberVO) HasDadEnable() bool`

HasDadEnable returns a boolean if a field has been set.

### GetDadIp

`func (o *OswMlagMemberVO) GetDadIp() string`

GetDadIp returns the DadIp field if non-nil, zero value otherwise.

### GetDadIpOk

`func (o *OswMlagMemberVO) GetDadIpOk() (*string, bool)`

GetDadIpOk returns a tuple with the DadIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDadIp

`func (o *OswMlagMemberVO) SetDadIp(v string)`

SetDadIp sets DadIp field to given value.

### HasDadIp

`func (o *OswMlagMemberVO) HasDadIp() bool`

HasDadIp returns a boolean if a field has been set.

### GetDadLinkPorts

`func (o *OswMlagMemberVO) GetDadLinkPorts() []string`

GetDadLinkPorts returns the DadLinkPorts field if non-nil, zero value otherwise.

### GetDadLinkPortsOk

`func (o *OswMlagMemberVO) GetDadLinkPortsOk() (*[]string, bool)`

GetDadLinkPortsOk returns a tuple with the DadLinkPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDadLinkPorts

`func (o *OswMlagMemberVO) SetDadLinkPorts(v []string)`

SetDadLinkPorts sets DadLinkPorts field to given value.

### HasDadLinkPorts

`func (o *OswMlagMemberVO) HasDadLinkPorts() bool`

HasDadLinkPorts returns a boolean if a field has been set.

### GetDadLocalIp

`func (o *OswMlagMemberVO) GetDadLocalIp() string`

GetDadLocalIp returns the DadLocalIp field if non-nil, zero value otherwise.

### GetDadLocalIpOk

`func (o *OswMlagMemberVO) GetDadLocalIpOk() (*string, bool)`

GetDadLocalIpOk returns a tuple with the DadLocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDadLocalIp

`func (o *OswMlagMemberVO) SetDadLocalIp(v string)`

SetDadLocalIp sets DadLocalIp field to given value.

### HasDadLocalIp

`func (o *OswMlagMemberVO) HasDadLocalIp() bool`

HasDadLocalIp returns a boolean if a field has been set.

### GetDadLocalIpv6

`func (o *OswMlagMemberVO) GetDadLocalIpv6() string`

GetDadLocalIpv6 returns the DadLocalIpv6 field if non-nil, zero value otherwise.

### GetDadLocalIpv6Ok

`func (o *OswMlagMemberVO) GetDadLocalIpv6Ok() (*string, bool)`

GetDadLocalIpv6Ok returns a tuple with the DadLocalIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDadLocalIpv6

`func (o *OswMlagMemberVO) SetDadLocalIpv6(v string)`

SetDadLocalIpv6 sets DadLocalIpv6 field to given value.

### HasDadLocalIpv6

`func (o *OswMlagMemberVO) HasDadLocalIpv6() bool`

HasDadLocalIpv6 returns a boolean if a field has been set.

### GetDadPeerIp

`func (o *OswMlagMemberVO) GetDadPeerIp() string`

GetDadPeerIp returns the DadPeerIp field if non-nil, zero value otherwise.

### GetDadPeerIpOk

`func (o *OswMlagMemberVO) GetDadPeerIpOk() (*string, bool)`

GetDadPeerIpOk returns a tuple with the DadPeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDadPeerIp

`func (o *OswMlagMemberVO) SetDadPeerIp(v string)`

SetDadPeerIp sets DadPeerIp field to given value.

### HasDadPeerIp

`func (o *OswMlagMemberVO) HasDadPeerIp() bool`

HasDadPeerIp returns a boolean if a field has been set.

### GetDadPeerIpv6

`func (o *OswMlagMemberVO) GetDadPeerIpv6() string`

GetDadPeerIpv6 returns the DadPeerIpv6 field if non-nil, zero value otherwise.

### GetDadPeerIpv6Ok

`func (o *OswMlagMemberVO) GetDadPeerIpv6Ok() (*string, bool)`

GetDadPeerIpv6Ok returns a tuple with the DadPeerIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDadPeerIpv6

`func (o *OswMlagMemberVO) SetDadPeerIpv6(v string)`

SetDadPeerIpv6 sets DadPeerIpv6 field to given value.

### HasDadPeerIpv6

`func (o *OswMlagMemberVO) HasDadPeerIpv6() bool`

HasDadPeerIpv6 returns a boolean if a field has been set.

### GetDadStatus

`func (o *OswMlagMemberVO) GetDadStatus() int32`

GetDadStatus returns the DadStatus field if non-nil, zero value otherwise.

### GetDadStatusOk

`func (o *OswMlagMemberVO) GetDadStatusOk() (*int32, bool)`

GetDadStatusOk returns a tuple with the DadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDadStatus

`func (o *OswMlagMemberVO) SetDadStatus(v int32)`

SetDadStatus sets DadStatus field to given value.

### HasDadStatus

`func (o *OswMlagMemberVO) HasDadStatus() bool`

HasDadStatus returns a boolean if a field has been set.

### GetDeviceType

`func (o *OswMlagMemberVO) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *OswMlagMemberVO) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *OswMlagMemberVO) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *OswMlagMemberVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDueTime

`func (o *OswMlagMemberVO) GetDueTime() int64`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *OswMlagMemberVO) GetDueTimeOk() (*int64, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *OswMlagMemberVO) SetDueTime(v int64)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *OswMlagMemberVO) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### GetDueTimeLeft

`func (o *OswMlagMemberVO) GetDueTimeLeft() int64`

GetDueTimeLeft returns the DueTimeLeft field if non-nil, zero value otherwise.

### GetDueTimeLeftOk

`func (o *OswMlagMemberVO) GetDueTimeLeftOk() (*int64, bool)`

GetDueTimeLeftOk returns a tuple with the DueTimeLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTimeLeft

`func (o *OswMlagMemberVO) SetDueTimeLeft(v int64)`

SetDueTimeLeft sets DueTimeLeft field to given value.

### HasDueTimeLeft

`func (o *OswMlagMemberVO) HasDueTimeLeft() bool`

HasDueTimeLeft returns a boolean if a field has been set.

### GetEos

`func (o *OswMlagMemberVO) GetEos() int64`

GetEos returns the Eos field if non-nil, zero value otherwise.

### GetEosOk

`func (o *OswMlagMemberVO) GetEosOk() (*int64, bool)`

GetEosOk returns a tuple with the Eos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEos

`func (o *OswMlagMemberVO) SetEos(v int64)`

SetEos sets Eos field to given value.

### HasEos

`func (o *OswMlagMemberVO) HasEos() bool`

HasEos returns a boolean if a field has been set.

### GetEost

`func (o *OswMlagMemberVO) GetEost() int64`

GetEost returns the Eost field if non-nil, zero value otherwise.

### GetEostOk

`func (o *OswMlagMemberVO) GetEostOk() (*int64, bool)`

GetEostOk returns a tuple with the Eost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEost

`func (o *OswMlagMemberVO) SetEost(v int64)`

SetEost sets Eost field to given value.

### HasEost

`func (o *OswMlagMemberVO) HasEost() bool`

HasEost returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *OswMlagMemberVO) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *OswMlagMemberVO) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *OswMlagMemberVO) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *OswMlagMemberVO) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetHwVersion

`func (o *OswMlagMemberVO) GetHwVersion() string`

GetHwVersion returns the HwVersion field if non-nil, zero value otherwise.

### GetHwVersionOk

`func (o *OswMlagMemberVO) GetHwVersionOk() (*string, bool)`

GetHwVersionOk returns a tuple with the HwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwVersion

`func (o *OswMlagMemberVO) SetHwVersion(v string)`

SetHwVersion sets HwVersion field to given value.

### HasHwVersion

`func (o *OswMlagMemberVO) HasHwVersion() bool`

HasHwVersion returns a boolean if a field has been set.

### GetInWhiteList

`func (o *OswMlagMemberVO) GetInWhiteList() bool`

GetInWhiteList returns the InWhiteList field if non-nil, zero value otherwise.

### GetInWhiteListOk

`func (o *OswMlagMemberVO) GetInWhiteListOk() (*bool, bool)`

GetInWhiteListOk returns a tuple with the InWhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInWhiteList

`func (o *OswMlagMemberVO) SetInWhiteList(v bool)`

SetInWhiteList sets InWhiteList field to given value.

### HasInWhiteList

`func (o *OswMlagMemberVO) HasInWhiteList() bool`

HasInWhiteList returns a boolean if a field has been set.

### GetIp

`func (o *OswMlagMemberVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OswMlagMemberVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OswMlagMemberVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OswMlagMemberVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6

`func (o *OswMlagMemberVO) GetIpv6() []string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *OswMlagMemberVO) GetIpv6Ok() (*[]string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *OswMlagMemberVO) SetIpv6(v []string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *OswMlagMemberVO) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetKaStatus

`func (o *OswMlagMemberVO) GetKaStatus() int32`

GetKaStatus returns the KaStatus field if non-nil, zero value otherwise.

### GetKaStatusOk

`func (o *OswMlagMemberVO) GetKaStatusOk() (*int32, bool)`

GetKaStatusOk returns a tuple with the KaStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKaStatus

`func (o *OswMlagMemberVO) SetKaStatus(v int32)`

SetKaStatus sets KaStatus field to given value.

### HasKaStatus

`func (o *OswMlagMemberVO) HasKaStatus() bool`

HasKaStatus returns a boolean if a field has been set.

### GetLicenseStatus

`func (o *OswMlagMemberVO) GetLicenseStatus() int32`

GetLicenseStatus returns the LicenseStatus field if non-nil, zero value otherwise.

### GetLicenseStatusOk

`func (o *OswMlagMemberVO) GetLicenseStatusOk() (*int32, bool)`

GetLicenseStatusOk returns a tuple with the LicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatus

`func (o *OswMlagMemberVO) SetLicenseStatus(v int32)`

SetLicenseStatus sets LicenseStatus field to given value.

### HasLicenseStatus

`func (o *OswMlagMemberVO) HasLicenseStatus() bool`

HasLicenseStatus returns a boolean if a field has been set.

### GetLicenseUnbindingLimit

`func (o *OswMlagMemberVO) GetLicenseUnbindingLimit() int32`

GetLicenseUnbindingLimit returns the LicenseUnbindingLimit field if non-nil, zero value otherwise.

### GetLicenseUnbindingLimitOk

`func (o *OswMlagMemberVO) GetLicenseUnbindingLimitOk() (*int32, bool)`

GetLicenseUnbindingLimitOk returns a tuple with the LicenseUnbindingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUnbindingLimit

`func (o *OswMlagMemberVO) SetLicenseUnbindingLimit(v int32)`

SetLicenseUnbindingLimit sets LicenseUnbindingLimit field to given value.

### HasLicenseUnbindingLimit

`func (o *OswMlagMemberVO) HasLicenseUnbindingLimit() bool`

HasLicenseUnbindingLimit returns a boolean if a field has been set.

### GetLocateEnable

`func (o *OswMlagMemberVO) GetLocateEnable() bool`

GetLocateEnable returns the LocateEnable field if non-nil, zero value otherwise.

### GetLocateEnableOk

`func (o *OswMlagMemberVO) GetLocateEnableOk() (*bool, bool)`

GetLocateEnableOk returns a tuple with the LocateEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocateEnable

`func (o *OswMlagMemberVO) SetLocateEnable(v bool)`

SetLocateEnable sets LocateEnable field to given value.

### HasLocateEnable

`func (o *OswMlagMemberVO) HasLocateEnable() bool`

HasLocateEnable returns a boolean if a field has been set.

### GetMac

`func (o *OswMlagMemberVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswMlagMemberVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswMlagMemberVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswMlagMemberVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMlagGroupId

`func (o *OswMlagMemberVO) GetMlagGroupId() int32`

GetMlagGroupId returns the MlagGroupId field if non-nil, zero value otherwise.

### GetMlagGroupIdOk

`func (o *OswMlagMemberVO) GetMlagGroupIdOk() (*int32, bool)`

GetMlagGroupIdOk returns a tuple with the MlagGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagGroupId

`func (o *OswMlagMemberVO) SetMlagGroupId(v int32)`

SetMlagGroupId sets MlagGroupId field to given value.

### HasMlagGroupId

`func (o *OswMlagMemberVO) HasMlagGroupId() bool`

HasMlagGroupId returns a boolean if a field has been set.

### GetMlagVersion

`func (o *OswMlagMemberVO) GetMlagVersion() string`

GetMlagVersion returns the MlagVersion field if non-nil, zero value otherwise.

### GetMlagVersionOk

`func (o *OswMlagMemberVO) GetMlagVersionOk() (*string, bool)`

GetMlagVersionOk returns a tuple with the MlagVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagVersion

`func (o *OswMlagMemberVO) SetMlagVersion(v string)`

SetMlagVersion sets MlagVersion field to given value.

### HasMlagVersion

`func (o *OswMlagMemberVO) HasMlagVersion() bool`

HasMlagVersion returns a boolean if a field has been set.

### GetModel

`func (o *OswMlagMemberVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OswMlagMemberVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OswMlagMemberVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OswMlagMemberVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OswMlagMemberVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OswMlagMemberVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OswMlagMemberVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OswMlagMemberVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *OswMlagMemberVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswMlagMemberVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswMlagMemberVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswMlagMemberVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPeerLinkPorts

`func (o *OswMlagMemberVO) GetPeerLinkPorts() []string`

GetPeerLinkPorts returns the PeerLinkPorts field if non-nil, zero value otherwise.

### GetPeerLinkPortsOk

`func (o *OswMlagMemberVO) GetPeerLinkPortsOk() (*[]string, bool)`

GetPeerLinkPortsOk returns a tuple with the PeerLinkPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerLinkPorts

`func (o *OswMlagMemberVO) SetPeerLinkPorts(v []string)`

SetPeerLinkPorts sets PeerLinkPorts field to given value.

### HasPeerLinkPorts

`func (o *OswMlagMemberVO) HasPeerLinkPorts() bool`

HasPeerLinkPorts returns a boolean if a field has been set.

### GetPorts

`func (o *OswMlagMemberVO) GetPorts() []OswMlagPortVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OswMlagMemberVO) GetPortsOk() (*[]OswMlagPortVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OswMlagMemberVO) SetPorts(v []OswMlagPortVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OswMlagMemberVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPriority

`func (o *OswMlagMemberVO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *OswMlagMemberVO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *OswMlagMemberVO) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *OswMlagMemberVO) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPublicIp

`func (o *OswMlagMemberVO) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *OswMlagMemberVO) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *OswMlagMemberVO) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *OswMlagMemberVO) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRole

`func (o *OswMlagMemberVO) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OswMlagMemberVO) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OswMlagMemberVO) SetRole(v int32)`

SetRole sets Role field to given value.

### HasRole

`func (o *OswMlagMemberVO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetShowModel

`func (o *OswMlagMemberVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *OswMlagMemberVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *OswMlagMemberVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *OswMlagMemberVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetSpecialModel

`func (o *OswMlagMemberVO) GetSpecialModel() string`

GetSpecialModel returns the SpecialModel field if non-nil, zero value otherwise.

### GetSpecialModelOk

`func (o *OswMlagMemberVO) GetSpecialModelOk() (*string, bool)`

GetSpecialModelOk returns a tuple with the SpecialModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialModel

`func (o *OswMlagMemberVO) SetSpecialModel(v string)`

SetSpecialModel sets SpecialModel field to given value.

### HasSpecialModel

`func (o *OswMlagMemberVO) HasSpecialModel() bool`

HasSpecialModel returns a boolean if a field has been set.

### GetStatus

`func (o *OswMlagMemberVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswMlagMemberVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswMlagMemberVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OswMlagMemberVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *OswMlagMemberVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *OswMlagMemberVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *OswMlagMemberVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *OswMlagMemberVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetSupportMlagIpv6

`func (o *OswMlagMemberVO) GetSupportMlagIpv6() bool`

GetSupportMlagIpv6 returns the SupportMlagIpv6 field if non-nil, zero value otherwise.

### GetSupportMlagIpv6Ok

`func (o *OswMlagMemberVO) GetSupportMlagIpv6Ok() (*bool, bool)`

GetSupportMlagIpv6Ok returns a tuple with the SupportMlagIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMlagIpv6

`func (o *OswMlagMemberVO) SetSupportMlagIpv6(v bool)`

SetSupportMlagIpv6 sets SupportMlagIpv6 field to given value.

### HasSupportMlagIpv6

`func (o *OswMlagMemberVO) HasSupportMlagIpv6() bool`

HasSupportMlagIpv6 returns a boolean if a field has been set.

### GetSupportPeerLinkPorts

`func (o *OswMlagMemberVO) GetSupportPeerLinkPorts() []OswStandPortVO`

GetSupportPeerLinkPorts returns the SupportPeerLinkPorts field if non-nil, zero value otherwise.

### GetSupportPeerLinkPortsOk

`func (o *OswMlagMemberVO) GetSupportPeerLinkPortsOk() (*[]OswStandPortVO, bool)`

GetSupportPeerLinkPortsOk returns a tuple with the SupportPeerLinkPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPeerLinkPorts

`func (o *OswMlagMemberVO) SetSupportPeerLinkPorts(v []OswStandPortVO)`

SetSupportPeerLinkPorts sets SupportPeerLinkPorts field to given value.

### HasSupportPeerLinkPorts

`func (o *OswMlagMemberVO) HasSupportPeerLinkPorts() bool`

HasSupportPeerLinkPorts returns a boolean if a field has been set.

### GetUptime

`func (o *OswMlagMemberVO) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *OswMlagMemberVO) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *OswMlagMemberVO) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *OswMlagMemberVO) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVersion

`func (o *OswMlagMemberVO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OswMlagMemberVO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OswMlagMemberVO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OswMlagMemberVO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


