# DeviceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the device is activated. | [optional] 
**Compatible** | Pointer to **int32** | The compatible type of device. | [optional] 
**CpuUtil** | Pointer to **int32** | Device cpuUtil | [optional] 
**DetailStatus** | Pointer to **int32** | Status of device,status should be a value as follows: 0:Disconnected;1:Disconnected(Migrating);10:Provisioning;11:Configuring;12:Upgrading;13:Rebooting;14:Connected;15:Connected(Wireless);16:Connected(Migrating);17:Connected(Wireless,Migrating);20:Pending;21:Pending(Wireless);22:Adopting;23:Adopting(Wireless);24:Adopt Failed;25:Adopt Failed(Wireless);26:Managed By Others;27:Managed By Others(Wireless);30:Heartbeat Missed;31:Heartbeat Missed(Wireless);32:Heartbeat Missed(Migrating);33:Heartbeat Missed(Wireless,Migrating);40:Isolated;41:Isolated(Migrating);50:Slice Configuring | [optional] 
**DeviceSeriesType** | Pointer to **int32** | Device series type. 0 means basic, 1 means pro. | [optional] 
**Duplex** | Pointer to **int32** | Device uplink port duplex mode, duplex should be a value as follows: 0: Auto; 1: Half; 2: Full. | [optional] 
**FirmwareVersion** | Pointer to **string** | The device firmware version. | [optional] 
**InWhiteList** | Pointer to **bool** | Whether the device is in white list. | [optional] 
**Ip** | Pointer to **string** | Device IP | [optional] 
**Ipv6** | Pointer to **[]string** | Device IPv6 list | [optional] 
**LastSeen** | Pointer to **int64** | Device lastSeen | [optional] 
**LicenseStatus** | Pointer to **int32** | Device license status (Only for cloud base) should be a value as follows: 0: unActive; 1: Unbind; 2: Expired; 3: active | [optional] 
**LinkSpeed** | Pointer to **int32** | Device uplink port linkSpeed, linkSpeed should be a value as follows: 0: Auto; 1: 10M; 2: 100M; 3: 1000M; 4: 2500M; 5: 10G; 6: 5G; 7: 25G, 8: 100G. | [optional] 
**Mac** | Pointer to **string** | Device MAC | [optional] 
**MemUtil** | Pointer to **int32** | Device memUtil | [optional] 
**Model** | Pointer to **string** | Device model name with version | [optional] 
**ModelName** | Pointer to **string** | Device model name | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**Name** | Pointer to **string** | Device name | [optional] 
**PublicIp** | Pointer to **string** | Device public IP | [optional] 
**Sn** | Pointer to **string** | Device serial number | [optional] 
**Status** | Pointer to **int32** | Device status should be a value as follows: 0: Disconnected; 1: Connected; 2: Pending; 3: Heartbeat Missed; 4: Isolated | [optional] 
**Subtype** | Pointer to **string** | Switch subtype should be a value as follows: smart: Non-Agile Series Switch; es: Agile Series Switch. | [optional] 
**SwitchConsistent** | Pointer to **bool** | Whether the device can be adopted by the site. | [optional] 
**TagName** | Pointer to **string** | Device tag name | [optional] 
**Type** | Pointer to **string** | Device type | [optional] 
**UplinkDeviceMac** | Pointer to **string** | Uplink device mac | [optional] 
**UplinkDeviceName** | Pointer to **string** | Uplink device name | [optional] 
**UplinkDevicePort** | Pointer to **string** | Uplink device port | [optional] 
**Uptime** | Pointer to **string** | Device uptime | [optional] 

## Methods

### NewDeviceInfo

`func NewDeviceInfo() *DeviceInfo`

NewDeviceInfo instantiates a new DeviceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInfoWithDefaults

`func NewDeviceInfoWithDefaults() *DeviceInfo`

NewDeviceInfoWithDefaults instantiates a new DeviceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *DeviceInfo) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DeviceInfo) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DeviceInfo) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *DeviceInfo) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCompatible

`func (o *DeviceInfo) GetCompatible() int32`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *DeviceInfo) GetCompatibleOk() (*int32, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *DeviceInfo) SetCompatible(v int32)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *DeviceInfo) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetCpuUtil

`func (o *DeviceInfo) GetCpuUtil() int32`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *DeviceInfo) GetCpuUtilOk() (*int32, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *DeviceInfo) SetCpuUtil(v int32)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *DeviceInfo) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetDetailStatus

`func (o *DeviceInfo) GetDetailStatus() int32`

GetDetailStatus returns the DetailStatus field if non-nil, zero value otherwise.

### GetDetailStatusOk

`func (o *DeviceInfo) GetDetailStatusOk() (*int32, bool)`

GetDetailStatusOk returns a tuple with the DetailStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailStatus

`func (o *DeviceInfo) SetDetailStatus(v int32)`

SetDetailStatus sets DetailStatus field to given value.

### HasDetailStatus

`func (o *DeviceInfo) HasDetailStatus() bool`

HasDetailStatus returns a boolean if a field has been set.

### GetDeviceSeriesType

`func (o *DeviceInfo) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *DeviceInfo) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *DeviceInfo) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *DeviceInfo) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetDuplex

`func (o *DeviceInfo) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *DeviceInfo) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *DeviceInfo) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *DeviceInfo) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *DeviceInfo) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *DeviceInfo) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *DeviceInfo) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *DeviceInfo) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetInWhiteList

`func (o *DeviceInfo) GetInWhiteList() bool`

GetInWhiteList returns the InWhiteList field if non-nil, zero value otherwise.

### GetInWhiteListOk

`func (o *DeviceInfo) GetInWhiteListOk() (*bool, bool)`

GetInWhiteListOk returns a tuple with the InWhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInWhiteList

`func (o *DeviceInfo) SetInWhiteList(v bool)`

SetInWhiteList sets InWhiteList field to given value.

### HasInWhiteList

`func (o *DeviceInfo) HasInWhiteList() bool`

HasInWhiteList returns a boolean if a field has been set.

### GetIp

`func (o *DeviceInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DeviceInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DeviceInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *DeviceInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6

`func (o *DeviceInfo) GetIpv6() []string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *DeviceInfo) GetIpv6Ok() (*[]string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *DeviceInfo) SetIpv6(v []string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *DeviceInfo) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetLastSeen

`func (o *DeviceInfo) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *DeviceInfo) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *DeviceInfo) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *DeviceInfo) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLicenseStatus

`func (o *DeviceInfo) GetLicenseStatus() int32`

GetLicenseStatus returns the LicenseStatus field if non-nil, zero value otherwise.

### GetLicenseStatusOk

`func (o *DeviceInfo) GetLicenseStatusOk() (*int32, bool)`

GetLicenseStatusOk returns a tuple with the LicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatus

`func (o *DeviceInfo) SetLicenseStatus(v int32)`

SetLicenseStatus sets LicenseStatus field to given value.

### HasLicenseStatus

`func (o *DeviceInfo) HasLicenseStatus() bool`

HasLicenseStatus returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *DeviceInfo) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *DeviceInfo) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *DeviceInfo) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *DeviceInfo) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetMac

`func (o *DeviceInfo) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DeviceInfo) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DeviceInfo) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *DeviceInfo) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMemUtil

`func (o *DeviceInfo) GetMemUtil() int32`

GetMemUtil returns the MemUtil field if non-nil, zero value otherwise.

### GetMemUtilOk

`func (o *DeviceInfo) GetMemUtilOk() (*int32, bool)`

GetMemUtilOk returns a tuple with the MemUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUtil

`func (o *DeviceInfo) SetMemUtil(v int32)`

SetMemUtil sets MemUtil field to given value.

### HasMemUtil

`func (o *DeviceInfo) HasMemUtil() bool`

HasMemUtil returns a boolean if a field has been set.

### GetModel

`func (o *DeviceInfo) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceInfo) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceInfo) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DeviceInfo) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelName

`func (o *DeviceInfo) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *DeviceInfo) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *DeviceInfo) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *DeviceInfo) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetModelVersion

`func (o *DeviceInfo) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *DeviceInfo) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *DeviceInfo) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *DeviceInfo) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *DeviceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicIp

`func (o *DeviceInfo) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *DeviceInfo) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *DeviceInfo) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *DeviceInfo) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetSn

`func (o *DeviceInfo) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *DeviceInfo) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *DeviceInfo) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *DeviceInfo) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceInfo) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceInfo) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceInfo) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubtype

`func (o *DeviceInfo) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *DeviceInfo) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *DeviceInfo) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *DeviceInfo) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetSwitchConsistent

`func (o *DeviceInfo) GetSwitchConsistent() bool`

GetSwitchConsistent returns the SwitchConsistent field if non-nil, zero value otherwise.

### GetSwitchConsistentOk

`func (o *DeviceInfo) GetSwitchConsistentOk() (*bool, bool)`

GetSwitchConsistentOk returns a tuple with the SwitchConsistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchConsistent

`func (o *DeviceInfo) SetSwitchConsistent(v bool)`

SetSwitchConsistent sets SwitchConsistent field to given value.

### HasSwitchConsistent

`func (o *DeviceInfo) HasSwitchConsistent() bool`

HasSwitchConsistent returns a boolean if a field has been set.

### GetTagName

`func (o *DeviceInfo) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *DeviceInfo) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *DeviceInfo) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *DeviceInfo) HasTagName() bool`

HasTagName returns a boolean if a field has been set.

### GetType

`func (o *DeviceInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeviceInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUplinkDeviceMac

`func (o *DeviceInfo) GetUplinkDeviceMac() string`

GetUplinkDeviceMac returns the UplinkDeviceMac field if non-nil, zero value otherwise.

### GetUplinkDeviceMacOk

`func (o *DeviceInfo) GetUplinkDeviceMacOk() (*string, bool)`

GetUplinkDeviceMacOk returns a tuple with the UplinkDeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkDeviceMac

`func (o *DeviceInfo) SetUplinkDeviceMac(v string)`

SetUplinkDeviceMac sets UplinkDeviceMac field to given value.

### HasUplinkDeviceMac

`func (o *DeviceInfo) HasUplinkDeviceMac() bool`

HasUplinkDeviceMac returns a boolean if a field has been set.

### GetUplinkDeviceName

`func (o *DeviceInfo) GetUplinkDeviceName() string`

GetUplinkDeviceName returns the UplinkDeviceName field if non-nil, zero value otherwise.

### GetUplinkDeviceNameOk

`func (o *DeviceInfo) GetUplinkDeviceNameOk() (*string, bool)`

GetUplinkDeviceNameOk returns a tuple with the UplinkDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkDeviceName

`func (o *DeviceInfo) SetUplinkDeviceName(v string)`

SetUplinkDeviceName sets UplinkDeviceName field to given value.

### HasUplinkDeviceName

`func (o *DeviceInfo) HasUplinkDeviceName() bool`

HasUplinkDeviceName returns a boolean if a field has been set.

### GetUplinkDevicePort

`func (o *DeviceInfo) GetUplinkDevicePort() string`

GetUplinkDevicePort returns the UplinkDevicePort field if non-nil, zero value otherwise.

### GetUplinkDevicePortOk

`func (o *DeviceInfo) GetUplinkDevicePortOk() (*string, bool)`

GetUplinkDevicePortOk returns a tuple with the UplinkDevicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkDevicePort

`func (o *DeviceInfo) SetUplinkDevicePort(v string)`

SetUplinkDevicePort sets UplinkDevicePort field to given value.

### HasUplinkDevicePort

`func (o *DeviceInfo) HasUplinkDevicePort() bool`

HasUplinkDevicePort returns a boolean if a field has been set.

### GetUptime

`func (o *DeviceInfo) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *DeviceInfo) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *DeviceInfo) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *DeviceInfo) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


