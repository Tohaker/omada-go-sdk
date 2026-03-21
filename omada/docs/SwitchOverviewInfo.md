# SwitchOverviewInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuUtil** | Pointer to **int32** | CpuUtil | [optional] 
**FirmwareVersion** | Pointer to **string** | Firmware Version e.g:2.5.0 Build 20190118 Rel. 64821 | [optional] 
**HwVersion** | Pointer to **string** | Hardware Version | [optional] 
**Ip** | Pointer to **string** | Switch IP address | [optional] 
**Ipv6List** | Pointer to **[]string** | Switch IPv6 list | [optional] 
**Mac** | Pointer to **string** | Switch MAC address | [optional] 
**MemUtil** | Pointer to **int32** | MemUtil | [optional] 
**Model** | Pointer to **string** | Model | [optional] 
**PortList** | Pointer to [**[]PortInfo**](PortInfo.md) | Port List | [optional] 
**Uptime** | Pointer to **string** | Uptime | [optional] 
**Version** | Pointer to **string** | Firmware Version e.g:2.5.0 | [optional] 

## Methods

### NewSwitchOverviewInfo

`func NewSwitchOverviewInfo() *SwitchOverviewInfo`

NewSwitchOverviewInfo instantiates a new SwitchOverviewInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchOverviewInfoWithDefaults

`func NewSwitchOverviewInfoWithDefaults() *SwitchOverviewInfo`

NewSwitchOverviewInfoWithDefaults instantiates a new SwitchOverviewInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuUtil

`func (o *SwitchOverviewInfo) GetCpuUtil() int32`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *SwitchOverviewInfo) GetCpuUtilOk() (*int32, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *SwitchOverviewInfo) SetCpuUtil(v int32)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *SwitchOverviewInfo) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *SwitchOverviewInfo) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *SwitchOverviewInfo) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *SwitchOverviewInfo) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *SwitchOverviewInfo) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetHwVersion

`func (o *SwitchOverviewInfo) GetHwVersion() string`

GetHwVersion returns the HwVersion field if non-nil, zero value otherwise.

### GetHwVersionOk

`func (o *SwitchOverviewInfo) GetHwVersionOk() (*string, bool)`

GetHwVersionOk returns a tuple with the HwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwVersion

`func (o *SwitchOverviewInfo) SetHwVersion(v string)`

SetHwVersion sets HwVersion field to given value.

### HasHwVersion

`func (o *SwitchOverviewInfo) HasHwVersion() bool`

HasHwVersion returns a boolean if a field has been set.

### GetIp

`func (o *SwitchOverviewInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SwitchOverviewInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SwitchOverviewInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SwitchOverviewInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6List

`func (o *SwitchOverviewInfo) GetIpv6List() []string`

GetIpv6List returns the Ipv6List field if non-nil, zero value otherwise.

### GetIpv6ListOk

`func (o *SwitchOverviewInfo) GetIpv6ListOk() (*[]string, bool)`

GetIpv6ListOk returns a tuple with the Ipv6List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6List

`func (o *SwitchOverviewInfo) SetIpv6List(v []string)`

SetIpv6List sets Ipv6List field to given value.

### HasIpv6List

`func (o *SwitchOverviewInfo) HasIpv6List() bool`

HasIpv6List returns a boolean if a field has been set.

### GetMac

`func (o *SwitchOverviewInfo) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SwitchOverviewInfo) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SwitchOverviewInfo) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *SwitchOverviewInfo) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMemUtil

`func (o *SwitchOverviewInfo) GetMemUtil() int32`

GetMemUtil returns the MemUtil field if non-nil, zero value otherwise.

### GetMemUtilOk

`func (o *SwitchOverviewInfo) GetMemUtilOk() (*int32, bool)`

GetMemUtilOk returns a tuple with the MemUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUtil

`func (o *SwitchOverviewInfo) SetMemUtil(v int32)`

SetMemUtil sets MemUtil field to given value.

### HasMemUtil

`func (o *SwitchOverviewInfo) HasMemUtil() bool`

HasMemUtil returns a boolean if a field has been set.

### GetModel

`func (o *SwitchOverviewInfo) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SwitchOverviewInfo) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SwitchOverviewInfo) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SwitchOverviewInfo) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPortList

`func (o *SwitchOverviewInfo) GetPortList() []PortInfo`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *SwitchOverviewInfo) GetPortListOk() (*[]PortInfo, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *SwitchOverviewInfo) SetPortList(v []PortInfo)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *SwitchOverviewInfo) HasPortList() bool`

HasPortList returns a boolean if a field has been set.

### GetUptime

`func (o *SwitchOverviewInfo) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *SwitchOverviewInfo) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *SwitchOverviewInfo) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *SwitchOverviewInfo) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVersion

`func (o *SwitchOverviewInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SwitchOverviewInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SwitchOverviewInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SwitchOverviewInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


