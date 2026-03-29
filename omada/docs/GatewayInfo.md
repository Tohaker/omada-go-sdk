# GatewayInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuUtil** | Pointer to **int32** | Gateway cpu utilization rate | [optional] 
**FirmwareVersion** | Pointer to **string** | Gateway software version | [optional] 
**Ip** | Pointer to **string** | Gateway IPv4 | [optional] 
**Ipv6List** | Pointer to **[]string** | IPv6 address List | [optional] 
**LastSeen** | Pointer to **int64** | Last seen time | [optional] 
**Mac** | Pointer to **string** | Gateway MAC address | [optional] 
**MemUtil** | Pointer to **int32** | Gateway memory utilization rate | [optional] 
**MultiChipGateway** | Pointer to **bool** | Whether this gateway is a multi-chip gateway | [optional] 
**MultiChipInfos** | Pointer to **[][]int32** | Only ports within the same group can be mirrored | [optional] 
**PortConfigs** | Pointer to [**[]GatewayPortConfig**](GatewayPortConfig.md) | Gateway port configs | [optional] 
**ShowModel** | Pointer to **string** | Gateway model description | [optional] 
**Temp** | Pointer to **int32** | Gateway temperature | [optional] 
**Uptime** | Pointer to **string** | Gateway uptime | [optional] 

## Methods

### NewGatewayInfo

`func NewGatewayInfo() *GatewayInfo`

NewGatewayInfo instantiates a new GatewayInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayInfoWithDefaults

`func NewGatewayInfoWithDefaults() *GatewayInfo`

NewGatewayInfoWithDefaults instantiates a new GatewayInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuUtil

`func (o *GatewayInfo) GetCpuUtil() int32`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *GatewayInfo) GetCpuUtilOk() (*int32, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *GatewayInfo) SetCpuUtil(v int32)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *GatewayInfo) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *GatewayInfo) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *GatewayInfo) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *GatewayInfo) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *GatewayInfo) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetIp

`func (o *GatewayInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GatewayInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GatewayInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GatewayInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6List

`func (o *GatewayInfo) GetIpv6List() []string`

GetIpv6List returns the Ipv6List field if non-nil, zero value otherwise.

### GetIpv6ListOk

`func (o *GatewayInfo) GetIpv6ListOk() (*[]string, bool)`

GetIpv6ListOk returns a tuple with the Ipv6List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6List

`func (o *GatewayInfo) SetIpv6List(v []string)`

SetIpv6List sets Ipv6List field to given value.

### HasIpv6List

`func (o *GatewayInfo) HasIpv6List() bool`

HasIpv6List returns a boolean if a field has been set.

### GetLastSeen

`func (o *GatewayInfo) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *GatewayInfo) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *GatewayInfo) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *GatewayInfo) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMac

`func (o *GatewayInfo) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GatewayInfo) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GatewayInfo) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GatewayInfo) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMemUtil

`func (o *GatewayInfo) GetMemUtil() int32`

GetMemUtil returns the MemUtil field if non-nil, zero value otherwise.

### GetMemUtilOk

`func (o *GatewayInfo) GetMemUtilOk() (*int32, bool)`

GetMemUtilOk returns a tuple with the MemUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUtil

`func (o *GatewayInfo) SetMemUtil(v int32)`

SetMemUtil sets MemUtil field to given value.

### HasMemUtil

`func (o *GatewayInfo) HasMemUtil() bool`

HasMemUtil returns a boolean if a field has been set.

### GetMultiChipGateway

`func (o *GatewayInfo) GetMultiChipGateway() bool`

GetMultiChipGateway returns the MultiChipGateway field if non-nil, zero value otherwise.

### GetMultiChipGatewayOk

`func (o *GatewayInfo) GetMultiChipGatewayOk() (*bool, bool)`

GetMultiChipGatewayOk returns a tuple with the MultiChipGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiChipGateway

`func (o *GatewayInfo) SetMultiChipGateway(v bool)`

SetMultiChipGateway sets MultiChipGateway field to given value.

### HasMultiChipGateway

`func (o *GatewayInfo) HasMultiChipGateway() bool`

HasMultiChipGateway returns a boolean if a field has been set.

### GetMultiChipInfos

`func (o *GatewayInfo) GetMultiChipInfos() [][]int32`

GetMultiChipInfos returns the MultiChipInfos field if non-nil, zero value otherwise.

### GetMultiChipInfosOk

`func (o *GatewayInfo) GetMultiChipInfosOk() (*[][]int32, bool)`

GetMultiChipInfosOk returns a tuple with the MultiChipInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiChipInfos

`func (o *GatewayInfo) SetMultiChipInfos(v [][]int32)`

SetMultiChipInfos sets MultiChipInfos field to given value.

### HasMultiChipInfos

`func (o *GatewayInfo) HasMultiChipInfos() bool`

HasMultiChipInfos returns a boolean if a field has been set.

### GetPortConfigs

`func (o *GatewayInfo) GetPortConfigs() []GatewayPortConfig`

GetPortConfigs returns the PortConfigs field if non-nil, zero value otherwise.

### GetPortConfigsOk

`func (o *GatewayInfo) GetPortConfigsOk() (*[]GatewayPortConfig, bool)`

GetPortConfigsOk returns a tuple with the PortConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfigs

`func (o *GatewayInfo) SetPortConfigs(v []GatewayPortConfig)`

SetPortConfigs sets PortConfigs field to given value.

### HasPortConfigs

`func (o *GatewayInfo) HasPortConfigs() bool`

HasPortConfigs returns a boolean if a field has been set.

### GetShowModel

`func (o *GatewayInfo) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *GatewayInfo) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *GatewayInfo) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *GatewayInfo) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetTemp

`func (o *GatewayInfo) GetTemp() int32`

GetTemp returns the Temp field if non-nil, zero value otherwise.

### GetTempOk

`func (o *GatewayInfo) GetTempOk() (*int32, bool)`

GetTempOk returns a tuple with the Temp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemp

`func (o *GatewayInfo) SetTemp(v int32)`

SetTemp sets Temp field to given value.

### HasTemp

`func (o *GatewayInfo) HasTemp() bool`

HasTemp returns a boolean if a field has been set.

### GetUptime

`func (o *GatewayInfo) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *GatewayInfo) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *GatewayInfo) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *GatewayInfo) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


