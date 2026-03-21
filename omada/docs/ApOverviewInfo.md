# ApOverviewInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuUtil** | Pointer to **int32** | Device cpu util(like 1 means 1% cpu util) | [optional] 
**FirmwareVersion** | Pointer to **string** | Device firmware version | [optional] 
**Ip** | Pointer to **string** | Device IP | [optional] 
**Ipv6List** | Pointer to **[]string** | Device IPv6 | [optional] 
**Mac** | Pointer to **string** | Device MAC, e.g. 00-00-FF-FF-0C-E9 | [optional] 
**MemoryUtil** | Pointer to **int32** | Device memory util(like 50 means 50% memory util) | [optional] 
**Model** | Pointer to **string** | Device model | [optional] 
**Name** | Pointer to **string** | Device name | [optional] 
**Type** | Pointer to **string** | Device type | [optional] 
**UptimeLong** | Pointer to **int64** | Device uptime(unit:second) | [optional] 
**WirelessUplinkInfo** | Pointer to [**ApWirelessUplink**](ApWirelessUplink.md) |  | [optional] 
**WlanGroupId** | Pointer to **string** | WLAN group ID | [optional] 

## Methods

### NewApOverviewInfo

`func NewApOverviewInfo() *ApOverviewInfo`

NewApOverviewInfo instantiates a new ApOverviewInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApOverviewInfoWithDefaults

`func NewApOverviewInfoWithDefaults() *ApOverviewInfo`

NewApOverviewInfoWithDefaults instantiates a new ApOverviewInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuUtil

`func (o *ApOverviewInfo) GetCpuUtil() int32`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *ApOverviewInfo) GetCpuUtilOk() (*int32, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *ApOverviewInfo) SetCpuUtil(v int32)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *ApOverviewInfo) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *ApOverviewInfo) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *ApOverviewInfo) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *ApOverviewInfo) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *ApOverviewInfo) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetIp

`func (o *ApOverviewInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ApOverviewInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ApOverviewInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ApOverviewInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6List

`func (o *ApOverviewInfo) GetIpv6List() []string`

GetIpv6List returns the Ipv6List field if non-nil, zero value otherwise.

### GetIpv6ListOk

`func (o *ApOverviewInfo) GetIpv6ListOk() (*[]string, bool)`

GetIpv6ListOk returns a tuple with the Ipv6List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6List

`func (o *ApOverviewInfo) SetIpv6List(v []string)`

SetIpv6List sets Ipv6List field to given value.

### HasIpv6List

`func (o *ApOverviewInfo) HasIpv6List() bool`

HasIpv6List returns a boolean if a field has been set.

### GetMac

`func (o *ApOverviewInfo) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ApOverviewInfo) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ApOverviewInfo) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ApOverviewInfo) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMemoryUtil

`func (o *ApOverviewInfo) GetMemoryUtil() int32`

GetMemoryUtil returns the MemoryUtil field if non-nil, zero value otherwise.

### GetMemoryUtilOk

`func (o *ApOverviewInfo) GetMemoryUtilOk() (*int32, bool)`

GetMemoryUtilOk returns a tuple with the MemoryUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUtil

`func (o *ApOverviewInfo) SetMemoryUtil(v int32)`

SetMemoryUtil sets MemoryUtil field to given value.

### HasMemoryUtil

`func (o *ApOverviewInfo) HasMemoryUtil() bool`

HasMemoryUtil returns a boolean if a field has been set.

### GetModel

`func (o *ApOverviewInfo) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ApOverviewInfo) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ApOverviewInfo) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ApOverviewInfo) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *ApOverviewInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApOverviewInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApOverviewInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApOverviewInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ApOverviewInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApOverviewInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApOverviewInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApOverviewInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUptimeLong

`func (o *ApOverviewInfo) GetUptimeLong() int64`

GetUptimeLong returns the UptimeLong field if non-nil, zero value otherwise.

### GetUptimeLongOk

`func (o *ApOverviewInfo) GetUptimeLongOk() (*int64, bool)`

GetUptimeLongOk returns a tuple with the UptimeLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeLong

`func (o *ApOverviewInfo) SetUptimeLong(v int64)`

SetUptimeLong sets UptimeLong field to given value.

### HasUptimeLong

`func (o *ApOverviewInfo) HasUptimeLong() bool`

HasUptimeLong returns a boolean if a field has been set.

### GetWirelessUplinkInfo

`func (o *ApOverviewInfo) GetWirelessUplinkInfo() ApWirelessUplink`

GetWirelessUplinkInfo returns the WirelessUplinkInfo field if non-nil, zero value otherwise.

### GetWirelessUplinkInfoOk

`func (o *ApOverviewInfo) GetWirelessUplinkInfoOk() (*ApWirelessUplink, bool)`

GetWirelessUplinkInfoOk returns a tuple with the WirelessUplinkInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessUplinkInfo

`func (o *ApOverviewInfo) SetWirelessUplinkInfo(v ApWirelessUplink)`

SetWirelessUplinkInfo sets WirelessUplinkInfo field to given value.

### HasWirelessUplinkInfo

`func (o *ApOverviewInfo) HasWirelessUplinkInfo() bool`

HasWirelessUplinkInfo returns a boolean if a field has been set.

### GetWlanGroupId

`func (o *ApOverviewInfo) GetWlanGroupId() string`

GetWlanGroupId returns the WlanGroupId field if non-nil, zero value otherwise.

### GetWlanGroupIdOk

`func (o *ApOverviewInfo) GetWlanGroupIdOk() (*string, bool)`

GetWlanGroupIdOk returns a tuple with the WlanGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanGroupId

`func (o *ApOverviewInfo) SetWlanGroupId(v string)`

SetWlanGroupId sets WlanGroupId field to given value.

### HasWlanGroupId

`func (o *ApOverviewInfo) HasWlanGroupId() bool`

HasWlanGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


