# VirtualWanIpv4IpoaOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dns1** | Pointer to **string** | Primary DNS server. | [optional] 
**Dns2** | Pointer to **string** | Secondary DNS server. | [optional] 
**Gateway** | **string** | Gateway IP. | 
**Ipaddr** | Pointer to **string** | IP address. | [optional] 
**Mtu** | Pointer to **int32** | Parameter [mtu] should be a value between 576 and 1500. | [optional] 
**Netmask** | Pointer to **string** | Subnet mask. | [optional] 
**WanMultipleIps** | Pointer to [**[]VirtualWanMultipleIpVO**](VirtualWanMultipleIpVO.md) |  | [optional] 

## Methods

### NewVirtualWanIpv4IpoaOpenApiVO

`func NewVirtualWanIpv4IpoaOpenApiVO(gateway string, ) *VirtualWanIpv4IpoaOpenApiVO`

NewVirtualWanIpv4IpoaOpenApiVO instantiates a new VirtualWanIpv4IpoaOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualWanIpv4IpoaOpenApiVOWithDefaults

`func NewVirtualWanIpv4IpoaOpenApiVOWithDefaults() *VirtualWanIpv4IpoaOpenApiVO`

NewVirtualWanIpv4IpoaOpenApiVOWithDefaults instantiates a new VirtualWanIpv4IpoaOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDns1

`func (o *VirtualWanIpv4IpoaOpenApiVO) GetDns1() string`

GetDns1 returns the Dns1 field if non-nil, zero value otherwise.

### GetDns1Ok

`func (o *VirtualWanIpv4IpoaOpenApiVO) GetDns1Ok() (*string, bool)`

GetDns1Ok returns a tuple with the Dns1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns1

`func (o *VirtualWanIpv4IpoaOpenApiVO) SetDns1(v string)`

SetDns1 sets Dns1 field to given value.

### HasDns1

`func (o *VirtualWanIpv4IpoaOpenApiVO) HasDns1() bool`

HasDns1 returns a boolean if a field has been set.

### GetDns2

`func (o *VirtualWanIpv4IpoaOpenApiVO) GetDns2() string`

GetDns2 returns the Dns2 field if non-nil, zero value otherwise.

### GetDns2Ok

`func (o *VirtualWanIpv4IpoaOpenApiVO) GetDns2Ok() (*string, bool)`

GetDns2Ok returns a tuple with the Dns2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns2

`func (o *VirtualWanIpv4IpoaOpenApiVO) SetDns2(v string)`

SetDns2 sets Dns2 field to given value.

### HasDns2

`func (o *VirtualWanIpv4IpoaOpenApiVO) HasDns2() bool`

HasDns2 returns a boolean if a field has been set.

### GetGateway

`func (o *VirtualWanIpv4IpoaOpenApiVO) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *VirtualWanIpv4IpoaOpenApiVO) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *VirtualWanIpv4IpoaOpenApiVO) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetIpaddr

`func (o *VirtualWanIpv4IpoaOpenApiVO) GetIpaddr() string`

GetIpaddr returns the Ipaddr field if non-nil, zero value otherwise.

### GetIpaddrOk

`func (o *VirtualWanIpv4IpoaOpenApiVO) GetIpaddrOk() (*string, bool)`

GetIpaddrOk returns a tuple with the Ipaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddr

`func (o *VirtualWanIpv4IpoaOpenApiVO) SetIpaddr(v string)`

SetIpaddr sets Ipaddr field to given value.

### HasIpaddr

`func (o *VirtualWanIpv4IpoaOpenApiVO) HasIpaddr() bool`

HasIpaddr returns a boolean if a field has been set.

### GetMtu

`func (o *VirtualWanIpv4IpoaOpenApiVO) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VirtualWanIpv4IpoaOpenApiVO) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VirtualWanIpv4IpoaOpenApiVO) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VirtualWanIpv4IpoaOpenApiVO) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetNetmask

`func (o *VirtualWanIpv4IpoaOpenApiVO) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *VirtualWanIpv4IpoaOpenApiVO) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *VirtualWanIpv4IpoaOpenApiVO) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *VirtualWanIpv4IpoaOpenApiVO) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetWanMultipleIps

`func (o *VirtualWanIpv4IpoaOpenApiVO) GetWanMultipleIps() []VirtualWanMultipleIpVO`

GetWanMultipleIps returns the WanMultipleIps field if non-nil, zero value otherwise.

### GetWanMultipleIpsOk

`func (o *VirtualWanIpv4IpoaOpenApiVO) GetWanMultipleIpsOk() (*[]VirtualWanMultipleIpVO, bool)`

GetWanMultipleIpsOk returns a tuple with the WanMultipleIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanMultipleIps

`func (o *VirtualWanIpv4IpoaOpenApiVO) SetWanMultipleIps(v []VirtualWanMultipleIpVO)`

SetWanMultipleIps sets WanMultipleIps field to given value.

### HasWanMultipleIps

`func (o *VirtualWanIpv4IpoaOpenApiVO) HasWanMultipleIps() bool`

HasWanMultipleIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


