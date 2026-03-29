# OswStackMemberVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the device is activated | [optional] 
**Category** | Pointer to **string** | Category | [optional] 
**Compatible** | Pointer to **int32** | Compatible Type | [optional] 
**CompoundModel** | Pointer to **string** |  | [optional] 
**CpuUtil** | Pointer to **int32** | Real-time CPU usage | [optional] 
**DeviceSeriesType** | Pointer to **int32** | DeviceSeriesType should be a value as follows: 0: advanced; 1: pro | [optional] 
**DownlinkList** | Pointer to [**[]OswDownlinkVO**](OswDownlinkVO.md) | Downlink Omada device list | [optional] 
**DueTime** | Pointer to **int64** | Expiration time | [optional] 
**DueTimeLeft** | Pointer to **int64** | The number of milliseconds from the current time to the expiration time | [optional] 
**InWhitelist** | Pointer to **bool** | Indicates whether it is in the whitelist | [optional] 
**InitialUnbindingLimit** | Pointer to **int32** | The initial number of times to unbind the license | [optional] 
**Ip** | Pointer to **string** | IP | [optional] 
**LicenseId** | Pointer to **string** | License Id | [optional] 
**LicenseStatus** | Pointer to **int32** | License Status should be a value as follows: 0: UnActive; 1: Unbind; 2: Expired; 3: Active; 4: NotMatch; 5: NearExpired | [optional] 
**LicenseUnbindingLimit** | Pointer to **int32** | The remaining number of times to unbind the license | [optional] 
**LocateEnable** | Pointer to **bool** | Indicates whether the locate function is enabled | [optional] 
**Loop** | Pointer to **string** | Loopback port | [optional] 
**LoopbackNum** | Pointer to **int32** | Number of loops | [optional] 
**Mac** | **string** | Device Mac | 
**MaxStackGroups** | Pointer to **int32** | Maximum number of stacking port aggregation groups supported | [optional] 
**MaxStackUnitNumber** | Pointer to **int32** | The maximum unit number supported by the device, starting from 1 | [optional] 
**MemUitl** | Pointer to **int32** | Real-time memory usage | [optional] 
**MemberStatus** | Pointer to **int32** | Stack Member Status should be a value as follows: 0: master; 1: slaver; 2: notInitial; 3: licenseOverdue; 4: licenseUnactive | [optional] 
**Model** | Pointer to **string** | Member Model | [optional] 
**ModelVersion** | Pointer to **string** | Member Model Version | [optional] 
**Name** | Pointer to **string** | Device Name | [optional] 
**OmadacId** | Pointer to **string** | OmadacId | [optional] 
**OswStackPortGroupStatus** | Pointer to [**[]OswLagVO**](OswLagVO.md) | Including some port configuration information and port real-time status | [optional] 
**PoeRemain** | Pointer to **float64** | PoE Residual Power (W) | [optional] 
**PoeRemainPercent** | Pointer to **float64** | PoE Residual Power Percentage | [optional] 
**PoeTotalPower** | Pointer to **float64** | PoE Total Power (W) | [optional] 
**PortCaps** | Pointer to [**[]OswStackMemberPortCapAndStatusVO**](OswStackMemberPortCapAndStatusVO.md) | Port Information | [optional] 
**Ports** | Pointer to [**[]OswStackMemberPortVO**](OswStackMemberPortVO.md) | Including some port configuration information and port real-time status | [optional] 
**Priority** | **int32** | Priority of the device in the local stacking system | 
**ShowModel** | Pointer to **string** | ShowModel | [optional] 
**StackPortCap** | Pointer to **map[string][]string** |  | [optional] 
**StackPorts** | [**[]OswStackPortGroupVO**](OswStackPortGroupVO.md) | Stack port list | 
**Status** | Pointer to **int32** | Device Status | [optional] 
**StatusCategory** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** | Type | [optional] 
**Unit** | **int32** | Unit number of the local stacking system of the device | 
**Uplink** | Pointer to [**OswUplinkVO**](OswUplinkVO.md) |  | [optional] 
**Uptime** | Pointer to **string** | Uptime | [optional] 
**Version** | Pointer to **string** | Software Version | [optional] 

## Methods

### NewOswStackMemberVO

`func NewOswStackMemberVO(mac string, priority int32, stackPorts []OswStackPortGroupVO, unit int32, ) *OswStackMemberVO`

NewOswStackMemberVO instantiates a new OswStackMemberVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackMemberVOWithDefaults

`func NewOswStackMemberVOWithDefaults() *OswStackMemberVO`

NewOswStackMemberVOWithDefaults instantiates a new OswStackMemberVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *OswStackMemberVO) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OswStackMemberVO) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OswStackMemberVO) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *OswStackMemberVO) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCategory

`func (o *OswStackMemberVO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *OswStackMemberVO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *OswStackMemberVO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *OswStackMemberVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCompatible

`func (o *OswStackMemberVO) GetCompatible() int32`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *OswStackMemberVO) GetCompatibleOk() (*int32, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *OswStackMemberVO) SetCompatible(v int32)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *OswStackMemberVO) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetCompoundModel

`func (o *OswStackMemberVO) GetCompoundModel() string`

GetCompoundModel returns the CompoundModel field if non-nil, zero value otherwise.

### GetCompoundModelOk

`func (o *OswStackMemberVO) GetCompoundModelOk() (*string, bool)`

GetCompoundModelOk returns a tuple with the CompoundModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompoundModel

`func (o *OswStackMemberVO) SetCompoundModel(v string)`

SetCompoundModel sets CompoundModel field to given value.

### HasCompoundModel

`func (o *OswStackMemberVO) HasCompoundModel() bool`

HasCompoundModel returns a boolean if a field has been set.

### GetCpuUtil

`func (o *OswStackMemberVO) GetCpuUtil() int32`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *OswStackMemberVO) GetCpuUtilOk() (*int32, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *OswStackMemberVO) SetCpuUtil(v int32)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *OswStackMemberVO) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetDeviceSeriesType

`func (o *OswStackMemberVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *OswStackMemberVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *OswStackMemberVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *OswStackMemberVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetDownlinkList

`func (o *OswStackMemberVO) GetDownlinkList() []OswDownlinkVO`

GetDownlinkList returns the DownlinkList field if non-nil, zero value otherwise.

### GetDownlinkListOk

`func (o *OswStackMemberVO) GetDownlinkListOk() (*[]OswDownlinkVO, bool)`

GetDownlinkListOk returns a tuple with the DownlinkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkList

`func (o *OswStackMemberVO) SetDownlinkList(v []OswDownlinkVO)`

SetDownlinkList sets DownlinkList field to given value.

### HasDownlinkList

`func (o *OswStackMemberVO) HasDownlinkList() bool`

HasDownlinkList returns a boolean if a field has been set.

### GetDueTime

`func (o *OswStackMemberVO) GetDueTime() int64`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *OswStackMemberVO) GetDueTimeOk() (*int64, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *OswStackMemberVO) SetDueTime(v int64)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *OswStackMemberVO) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### GetDueTimeLeft

`func (o *OswStackMemberVO) GetDueTimeLeft() int64`

GetDueTimeLeft returns the DueTimeLeft field if non-nil, zero value otherwise.

### GetDueTimeLeftOk

`func (o *OswStackMemberVO) GetDueTimeLeftOk() (*int64, bool)`

GetDueTimeLeftOk returns a tuple with the DueTimeLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTimeLeft

`func (o *OswStackMemberVO) SetDueTimeLeft(v int64)`

SetDueTimeLeft sets DueTimeLeft field to given value.

### HasDueTimeLeft

`func (o *OswStackMemberVO) HasDueTimeLeft() bool`

HasDueTimeLeft returns a boolean if a field has been set.

### GetInWhitelist

`func (o *OswStackMemberVO) GetInWhitelist() bool`

GetInWhitelist returns the InWhitelist field if non-nil, zero value otherwise.

### GetInWhitelistOk

`func (o *OswStackMemberVO) GetInWhitelistOk() (*bool, bool)`

GetInWhitelistOk returns a tuple with the InWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInWhitelist

`func (o *OswStackMemberVO) SetInWhitelist(v bool)`

SetInWhitelist sets InWhitelist field to given value.

### HasInWhitelist

`func (o *OswStackMemberVO) HasInWhitelist() bool`

HasInWhitelist returns a boolean if a field has been set.

### GetInitialUnbindingLimit

`func (o *OswStackMemberVO) GetInitialUnbindingLimit() int32`

GetInitialUnbindingLimit returns the InitialUnbindingLimit field if non-nil, zero value otherwise.

### GetInitialUnbindingLimitOk

`func (o *OswStackMemberVO) GetInitialUnbindingLimitOk() (*int32, bool)`

GetInitialUnbindingLimitOk returns a tuple with the InitialUnbindingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialUnbindingLimit

`func (o *OswStackMemberVO) SetInitialUnbindingLimit(v int32)`

SetInitialUnbindingLimit sets InitialUnbindingLimit field to given value.

### HasInitialUnbindingLimit

`func (o *OswStackMemberVO) HasInitialUnbindingLimit() bool`

HasInitialUnbindingLimit returns a boolean if a field has been set.

### GetIp

`func (o *OswStackMemberVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OswStackMemberVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OswStackMemberVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OswStackMemberVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLicenseId

`func (o *OswStackMemberVO) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *OswStackMemberVO) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *OswStackMemberVO) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *OswStackMemberVO) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetLicenseStatus

`func (o *OswStackMemberVO) GetLicenseStatus() int32`

GetLicenseStatus returns the LicenseStatus field if non-nil, zero value otherwise.

### GetLicenseStatusOk

`func (o *OswStackMemberVO) GetLicenseStatusOk() (*int32, bool)`

GetLicenseStatusOk returns a tuple with the LicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatus

`func (o *OswStackMemberVO) SetLicenseStatus(v int32)`

SetLicenseStatus sets LicenseStatus field to given value.

### HasLicenseStatus

`func (o *OswStackMemberVO) HasLicenseStatus() bool`

HasLicenseStatus returns a boolean if a field has been set.

### GetLicenseUnbindingLimit

`func (o *OswStackMemberVO) GetLicenseUnbindingLimit() int32`

GetLicenseUnbindingLimit returns the LicenseUnbindingLimit field if non-nil, zero value otherwise.

### GetLicenseUnbindingLimitOk

`func (o *OswStackMemberVO) GetLicenseUnbindingLimitOk() (*int32, bool)`

GetLicenseUnbindingLimitOk returns a tuple with the LicenseUnbindingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUnbindingLimit

`func (o *OswStackMemberVO) SetLicenseUnbindingLimit(v int32)`

SetLicenseUnbindingLimit sets LicenseUnbindingLimit field to given value.

### HasLicenseUnbindingLimit

`func (o *OswStackMemberVO) HasLicenseUnbindingLimit() bool`

HasLicenseUnbindingLimit returns a boolean if a field has been set.

### GetLocateEnable

`func (o *OswStackMemberVO) GetLocateEnable() bool`

GetLocateEnable returns the LocateEnable field if non-nil, zero value otherwise.

### GetLocateEnableOk

`func (o *OswStackMemberVO) GetLocateEnableOk() (*bool, bool)`

GetLocateEnableOk returns a tuple with the LocateEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocateEnable

`func (o *OswStackMemberVO) SetLocateEnable(v bool)`

SetLocateEnable sets LocateEnable field to given value.

### HasLocateEnable

`func (o *OswStackMemberVO) HasLocateEnable() bool`

HasLocateEnable returns a boolean if a field has been set.

### GetLoop

`func (o *OswStackMemberVO) GetLoop() string`

GetLoop returns the Loop field if non-nil, zero value otherwise.

### GetLoopOk

`func (o *OswStackMemberVO) GetLoopOk() (*string, bool)`

GetLoopOk returns a tuple with the Loop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoop

`func (o *OswStackMemberVO) SetLoop(v string)`

SetLoop sets Loop field to given value.

### HasLoop

`func (o *OswStackMemberVO) HasLoop() bool`

HasLoop returns a boolean if a field has been set.

### GetLoopbackNum

`func (o *OswStackMemberVO) GetLoopbackNum() int32`

GetLoopbackNum returns the LoopbackNum field if non-nil, zero value otherwise.

### GetLoopbackNumOk

`func (o *OswStackMemberVO) GetLoopbackNumOk() (*int32, bool)`

GetLoopbackNumOk returns a tuple with the LoopbackNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackNum

`func (o *OswStackMemberVO) SetLoopbackNum(v int32)`

SetLoopbackNum sets LoopbackNum field to given value.

### HasLoopbackNum

`func (o *OswStackMemberVO) HasLoopbackNum() bool`

HasLoopbackNum returns a boolean if a field has been set.

### GetMac

`func (o *OswStackMemberVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswStackMemberVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswStackMemberVO) SetMac(v string)`

SetMac sets Mac field to given value.


### GetMaxStackGroups

`func (o *OswStackMemberVO) GetMaxStackGroups() int32`

GetMaxStackGroups returns the MaxStackGroups field if non-nil, zero value otherwise.

### GetMaxStackGroupsOk

`func (o *OswStackMemberVO) GetMaxStackGroupsOk() (*int32, bool)`

GetMaxStackGroupsOk returns a tuple with the MaxStackGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStackGroups

`func (o *OswStackMemberVO) SetMaxStackGroups(v int32)`

SetMaxStackGroups sets MaxStackGroups field to given value.

### HasMaxStackGroups

`func (o *OswStackMemberVO) HasMaxStackGroups() bool`

HasMaxStackGroups returns a boolean if a field has been set.

### GetMaxStackUnitNumber

`func (o *OswStackMemberVO) GetMaxStackUnitNumber() int32`

GetMaxStackUnitNumber returns the MaxStackUnitNumber field if non-nil, zero value otherwise.

### GetMaxStackUnitNumberOk

`func (o *OswStackMemberVO) GetMaxStackUnitNumberOk() (*int32, bool)`

GetMaxStackUnitNumberOk returns a tuple with the MaxStackUnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStackUnitNumber

`func (o *OswStackMemberVO) SetMaxStackUnitNumber(v int32)`

SetMaxStackUnitNumber sets MaxStackUnitNumber field to given value.

### HasMaxStackUnitNumber

`func (o *OswStackMemberVO) HasMaxStackUnitNumber() bool`

HasMaxStackUnitNumber returns a boolean if a field has been set.

### GetMemUitl

`func (o *OswStackMemberVO) GetMemUitl() int32`

GetMemUitl returns the MemUitl field if non-nil, zero value otherwise.

### GetMemUitlOk

`func (o *OswStackMemberVO) GetMemUitlOk() (*int32, bool)`

GetMemUitlOk returns a tuple with the MemUitl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUitl

`func (o *OswStackMemberVO) SetMemUitl(v int32)`

SetMemUitl sets MemUitl field to given value.

### HasMemUitl

`func (o *OswStackMemberVO) HasMemUitl() bool`

HasMemUitl returns a boolean if a field has been set.

### GetMemberStatus

`func (o *OswStackMemberVO) GetMemberStatus() int32`

GetMemberStatus returns the MemberStatus field if non-nil, zero value otherwise.

### GetMemberStatusOk

`func (o *OswStackMemberVO) GetMemberStatusOk() (*int32, bool)`

GetMemberStatusOk returns a tuple with the MemberStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberStatus

`func (o *OswStackMemberVO) SetMemberStatus(v int32)`

SetMemberStatus sets MemberStatus field to given value.

### HasMemberStatus

`func (o *OswStackMemberVO) HasMemberStatus() bool`

HasMemberStatus returns a boolean if a field has been set.

### GetModel

`func (o *OswStackMemberVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OswStackMemberVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OswStackMemberVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OswStackMemberVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OswStackMemberVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OswStackMemberVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OswStackMemberVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OswStackMemberVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *OswStackMemberVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswStackMemberVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswStackMemberVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswStackMemberVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOmadacId

`func (o *OswStackMemberVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *OswStackMemberVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *OswStackMemberVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *OswStackMemberVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetOswStackPortGroupStatus

`func (o *OswStackMemberVO) GetOswStackPortGroupStatus() []OswLagVO`

GetOswStackPortGroupStatus returns the OswStackPortGroupStatus field if non-nil, zero value otherwise.

### GetOswStackPortGroupStatusOk

`func (o *OswStackMemberVO) GetOswStackPortGroupStatusOk() (*[]OswLagVO, bool)`

GetOswStackPortGroupStatusOk returns a tuple with the OswStackPortGroupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOswStackPortGroupStatus

`func (o *OswStackMemberVO) SetOswStackPortGroupStatus(v []OswLagVO)`

SetOswStackPortGroupStatus sets OswStackPortGroupStatus field to given value.

### HasOswStackPortGroupStatus

`func (o *OswStackMemberVO) HasOswStackPortGroupStatus() bool`

HasOswStackPortGroupStatus returns a boolean if a field has been set.

### GetPoeRemain

`func (o *OswStackMemberVO) GetPoeRemain() float64`

GetPoeRemain returns the PoeRemain field if non-nil, zero value otherwise.

### GetPoeRemainOk

`func (o *OswStackMemberVO) GetPoeRemainOk() (*float64, bool)`

GetPoeRemainOk returns a tuple with the PoeRemain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeRemain

`func (o *OswStackMemberVO) SetPoeRemain(v float64)`

SetPoeRemain sets PoeRemain field to given value.

### HasPoeRemain

`func (o *OswStackMemberVO) HasPoeRemain() bool`

HasPoeRemain returns a boolean if a field has been set.

### GetPoeRemainPercent

`func (o *OswStackMemberVO) GetPoeRemainPercent() float64`

GetPoeRemainPercent returns the PoeRemainPercent field if non-nil, zero value otherwise.

### GetPoeRemainPercentOk

`func (o *OswStackMemberVO) GetPoeRemainPercentOk() (*float64, bool)`

GetPoeRemainPercentOk returns a tuple with the PoeRemainPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeRemainPercent

`func (o *OswStackMemberVO) SetPoeRemainPercent(v float64)`

SetPoeRemainPercent sets PoeRemainPercent field to given value.

### HasPoeRemainPercent

`func (o *OswStackMemberVO) HasPoeRemainPercent() bool`

HasPoeRemainPercent returns a boolean if a field has been set.

### GetPoeTotalPower

`func (o *OswStackMemberVO) GetPoeTotalPower() float64`

GetPoeTotalPower returns the PoeTotalPower field if non-nil, zero value otherwise.

### GetPoeTotalPowerOk

`func (o *OswStackMemberVO) GetPoeTotalPowerOk() (*float64, bool)`

GetPoeTotalPowerOk returns a tuple with the PoeTotalPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeTotalPower

`func (o *OswStackMemberVO) SetPoeTotalPower(v float64)`

SetPoeTotalPower sets PoeTotalPower field to given value.

### HasPoeTotalPower

`func (o *OswStackMemberVO) HasPoeTotalPower() bool`

HasPoeTotalPower returns a boolean if a field has been set.

### GetPortCaps

`func (o *OswStackMemberVO) GetPortCaps() []OswStackMemberPortCapAndStatusVO`

GetPortCaps returns the PortCaps field if non-nil, zero value otherwise.

### GetPortCapsOk

`func (o *OswStackMemberVO) GetPortCapsOk() (*[]OswStackMemberPortCapAndStatusVO, bool)`

GetPortCapsOk returns a tuple with the PortCaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCaps

`func (o *OswStackMemberVO) SetPortCaps(v []OswStackMemberPortCapAndStatusVO)`

SetPortCaps sets PortCaps field to given value.

### HasPortCaps

`func (o *OswStackMemberVO) HasPortCaps() bool`

HasPortCaps returns a boolean if a field has been set.

### GetPorts

`func (o *OswStackMemberVO) GetPorts() []OswStackMemberPortVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OswStackMemberVO) GetPortsOk() (*[]OswStackMemberPortVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OswStackMemberVO) SetPorts(v []OswStackMemberPortVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OswStackMemberVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPriority

`func (o *OswStackMemberVO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *OswStackMemberVO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *OswStackMemberVO) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetShowModel

`func (o *OswStackMemberVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *OswStackMemberVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *OswStackMemberVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *OswStackMemberVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetStackPortCap

`func (o *OswStackMemberVO) GetStackPortCap() map[string][]string`

GetStackPortCap returns the StackPortCap field if non-nil, zero value otherwise.

### GetStackPortCapOk

`func (o *OswStackMemberVO) GetStackPortCapOk() (*map[string][]string, bool)`

GetStackPortCapOk returns a tuple with the StackPortCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPortCap

`func (o *OswStackMemberVO) SetStackPortCap(v map[string][]string)`

SetStackPortCap sets StackPortCap field to given value.

### HasStackPortCap

`func (o *OswStackMemberVO) HasStackPortCap() bool`

HasStackPortCap returns a boolean if a field has been set.

### GetStackPorts

`func (o *OswStackMemberVO) GetStackPorts() []OswStackPortGroupVO`

GetStackPorts returns the StackPorts field if non-nil, zero value otherwise.

### GetStackPortsOk

`func (o *OswStackMemberVO) GetStackPortsOk() (*[]OswStackPortGroupVO, bool)`

GetStackPortsOk returns a tuple with the StackPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPorts

`func (o *OswStackMemberVO) SetStackPorts(v []OswStackPortGroupVO)`

SetStackPorts sets StackPorts field to given value.


### GetStatus

`func (o *OswStackMemberVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswStackMemberVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswStackMemberVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OswStackMemberVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *OswStackMemberVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *OswStackMemberVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *OswStackMemberVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *OswStackMemberVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetType

`func (o *OswStackMemberVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswStackMemberVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswStackMemberVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OswStackMemberVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *OswStackMemberVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OswStackMemberVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OswStackMemberVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.


### GetUplink

`func (o *OswStackMemberVO) GetUplink() OswUplinkVO`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *OswStackMemberVO) GetUplinkOk() (*OswUplinkVO, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *OswStackMemberVO) SetUplink(v OswUplinkVO)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *OswStackMemberVO) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetUptime

`func (o *OswStackMemberVO) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *OswStackMemberVO) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *OswStackMemberVO) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *OswStackMemberVO) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVersion

`func (o *OswStackMemberVO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OswStackMemberVO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OswStackMemberVO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OswStackMemberVO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


