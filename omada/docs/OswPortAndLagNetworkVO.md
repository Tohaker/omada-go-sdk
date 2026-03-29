# OswPortAndLagNetworkVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipaddr** | Pointer to **string** | IP address. | [optional] 
**Mode** | Pointer to **int32** | Ip address mode, it should be a value as follows: 0:static 1:dynamic | [optional] 
**NativeLags** | Pointer to **[]int32** | Native Vlan Lag List, for example [1, 2] | [optional] 
**NativePorts** | Pointer to **[]int32** | Native Vlan Port List, for example [1, 2] | [optional] 
**NativeStPorts** | Pointer to **[]string** | Native Vlan Standard Port List, for example [\&quot;1/0/1\&quot;, \&quot;1/0/2\&quot;] | [optional] 
**TagLags** | Pointer to **[]int32** | Tag Vlan Lag List, for example [1, 2] | [optional] 
**TagPorts** | Pointer to **[]int32** | Tag Vlan Port List, for example [1, 2] | [optional] 
**TagStPorts** | Pointer to **[]string** | Tag Vlan Standard Port List, for example [\&quot;1/0/1\&quot;, \&quot;1/0/2\&quot;] | [optional] 
**UntagLags** | Pointer to **[]int32** | Untag Vlan Lag List, for example [1, 2] | [optional] 
**UntagPorts** | Pointer to **[]int32** | Untag Vlan Port List, for example [1, 2] | [optional] 
**UntagStPorts** | Pointer to **[]string** | Untag Vlan Standard Port List, for example [\&quot;1/0/1\&quot;, \&quot;1/0/2\&quot;] | [optional] 
**Vlan** | Pointer to **int32** | Vlan ID | [optional] 

## Methods

### NewOswPortAndLagNetworkVO

`func NewOswPortAndLagNetworkVO() *OswPortAndLagNetworkVO`

NewOswPortAndLagNetworkVO instantiates a new OswPortAndLagNetworkVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswPortAndLagNetworkVOWithDefaults

`func NewOswPortAndLagNetworkVOWithDefaults() *OswPortAndLagNetworkVO`

NewOswPortAndLagNetworkVOWithDefaults instantiates a new OswPortAndLagNetworkVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpaddr

`func (o *OswPortAndLagNetworkVO) GetIpaddr() string`

GetIpaddr returns the Ipaddr field if non-nil, zero value otherwise.

### GetIpaddrOk

`func (o *OswPortAndLagNetworkVO) GetIpaddrOk() (*string, bool)`

GetIpaddrOk returns a tuple with the Ipaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddr

`func (o *OswPortAndLagNetworkVO) SetIpaddr(v string)`

SetIpaddr sets Ipaddr field to given value.

### HasIpaddr

`func (o *OswPortAndLagNetworkVO) HasIpaddr() bool`

HasIpaddr returns a boolean if a field has been set.

### GetMode

`func (o *OswPortAndLagNetworkVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OswPortAndLagNetworkVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OswPortAndLagNetworkVO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *OswPortAndLagNetworkVO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNativeLags

`func (o *OswPortAndLagNetworkVO) GetNativeLags() []int32`

GetNativeLags returns the NativeLags field if non-nil, zero value otherwise.

### GetNativeLagsOk

`func (o *OswPortAndLagNetworkVO) GetNativeLagsOk() (*[]int32, bool)`

GetNativeLagsOk returns a tuple with the NativeLags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeLags

`func (o *OswPortAndLagNetworkVO) SetNativeLags(v []int32)`

SetNativeLags sets NativeLags field to given value.

### HasNativeLags

`func (o *OswPortAndLagNetworkVO) HasNativeLags() bool`

HasNativeLags returns a boolean if a field has been set.

### GetNativePorts

`func (o *OswPortAndLagNetworkVO) GetNativePorts() []int32`

GetNativePorts returns the NativePorts field if non-nil, zero value otherwise.

### GetNativePortsOk

`func (o *OswPortAndLagNetworkVO) GetNativePortsOk() (*[]int32, bool)`

GetNativePortsOk returns a tuple with the NativePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativePorts

`func (o *OswPortAndLagNetworkVO) SetNativePorts(v []int32)`

SetNativePorts sets NativePorts field to given value.

### HasNativePorts

`func (o *OswPortAndLagNetworkVO) HasNativePorts() bool`

HasNativePorts returns a boolean if a field has been set.

### GetNativeStPorts

`func (o *OswPortAndLagNetworkVO) GetNativeStPorts() []string`

GetNativeStPorts returns the NativeStPorts field if non-nil, zero value otherwise.

### GetNativeStPortsOk

`func (o *OswPortAndLagNetworkVO) GetNativeStPortsOk() (*[]string, bool)`

GetNativeStPortsOk returns a tuple with the NativeStPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeStPorts

`func (o *OswPortAndLagNetworkVO) SetNativeStPorts(v []string)`

SetNativeStPorts sets NativeStPorts field to given value.

### HasNativeStPorts

`func (o *OswPortAndLagNetworkVO) HasNativeStPorts() bool`

HasNativeStPorts returns a boolean if a field has been set.

### GetTagLags

`func (o *OswPortAndLagNetworkVO) GetTagLags() []int32`

GetTagLags returns the TagLags field if non-nil, zero value otherwise.

### GetTagLagsOk

`func (o *OswPortAndLagNetworkVO) GetTagLagsOk() (*[]int32, bool)`

GetTagLagsOk returns a tuple with the TagLags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagLags

`func (o *OswPortAndLagNetworkVO) SetTagLags(v []int32)`

SetTagLags sets TagLags field to given value.

### HasTagLags

`func (o *OswPortAndLagNetworkVO) HasTagLags() bool`

HasTagLags returns a boolean if a field has been set.

### GetTagPorts

`func (o *OswPortAndLagNetworkVO) GetTagPorts() []int32`

GetTagPorts returns the TagPorts field if non-nil, zero value otherwise.

### GetTagPortsOk

`func (o *OswPortAndLagNetworkVO) GetTagPortsOk() (*[]int32, bool)`

GetTagPortsOk returns a tuple with the TagPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagPorts

`func (o *OswPortAndLagNetworkVO) SetTagPorts(v []int32)`

SetTagPorts sets TagPorts field to given value.

### HasTagPorts

`func (o *OswPortAndLagNetworkVO) HasTagPorts() bool`

HasTagPorts returns a boolean if a field has been set.

### GetTagStPorts

`func (o *OswPortAndLagNetworkVO) GetTagStPorts() []string`

GetTagStPorts returns the TagStPorts field if non-nil, zero value otherwise.

### GetTagStPortsOk

`func (o *OswPortAndLagNetworkVO) GetTagStPortsOk() (*[]string, bool)`

GetTagStPortsOk returns a tuple with the TagStPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagStPorts

`func (o *OswPortAndLagNetworkVO) SetTagStPorts(v []string)`

SetTagStPorts sets TagStPorts field to given value.

### HasTagStPorts

`func (o *OswPortAndLagNetworkVO) HasTagStPorts() bool`

HasTagStPorts returns a boolean if a field has been set.

### GetUntagLags

`func (o *OswPortAndLagNetworkVO) GetUntagLags() []int32`

GetUntagLags returns the UntagLags field if non-nil, zero value otherwise.

### GetUntagLagsOk

`func (o *OswPortAndLagNetworkVO) GetUntagLagsOk() (*[]int32, bool)`

GetUntagLagsOk returns a tuple with the UntagLags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagLags

`func (o *OswPortAndLagNetworkVO) SetUntagLags(v []int32)`

SetUntagLags sets UntagLags field to given value.

### HasUntagLags

`func (o *OswPortAndLagNetworkVO) HasUntagLags() bool`

HasUntagLags returns a boolean if a field has been set.

### GetUntagPorts

`func (o *OswPortAndLagNetworkVO) GetUntagPorts() []int32`

GetUntagPorts returns the UntagPorts field if non-nil, zero value otherwise.

### GetUntagPortsOk

`func (o *OswPortAndLagNetworkVO) GetUntagPortsOk() (*[]int32, bool)`

GetUntagPortsOk returns a tuple with the UntagPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagPorts

`func (o *OswPortAndLagNetworkVO) SetUntagPorts(v []int32)`

SetUntagPorts sets UntagPorts field to given value.

### HasUntagPorts

`func (o *OswPortAndLagNetworkVO) HasUntagPorts() bool`

HasUntagPorts returns a boolean if a field has been set.

### GetUntagStPorts

`func (o *OswPortAndLagNetworkVO) GetUntagStPorts() []string`

GetUntagStPorts returns the UntagStPorts field if non-nil, zero value otherwise.

### GetUntagStPortsOk

`func (o *OswPortAndLagNetworkVO) GetUntagStPortsOk() (*[]string, bool)`

GetUntagStPortsOk returns a tuple with the UntagStPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagStPorts

`func (o *OswPortAndLagNetworkVO) SetUntagStPorts(v []string)`

SetUntagStPorts sets UntagStPorts field to given value.

### HasUntagStPorts

`func (o *OswPortAndLagNetworkVO) HasUntagStPorts() bool`

HasUntagStPorts returns a boolean if a field has been set.

### GetVlan

`func (o *OswPortAndLagNetworkVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *OswPortAndLagNetworkVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *OswPortAndLagNetworkVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *OswPortAndLagNetworkVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


