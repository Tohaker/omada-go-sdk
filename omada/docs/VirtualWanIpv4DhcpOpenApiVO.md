# VirtualWanIpv4DhcpOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpOptions** | Pointer to [**[]CustomDHCPOptions**](CustomDHCPOptions.md) | Virtual WAN custom DHCP options. | [optional] 
**Dns1** | Pointer to **string** | Primary DNS server. | [optional] 
**Dns2** | Pointer to **string** | Secondary DNS server. | [optional] 
**Hostname** | Pointer to **string** | Host name. Parameter [hostname] should be up to 63 characters long and can only use numbers, letters, and underscores. | [optional] 
**Mtu** | Pointer to **int32** | Parameter [mtu] should be a value between 576 and 1500. | [optional] 
**Unicast** | Pointer to **string** | Subnet mask of virtual WAN. | [optional] 

## Methods

### NewVirtualWanIpv4DhcpOpenApiVO

`func NewVirtualWanIpv4DhcpOpenApiVO() *VirtualWanIpv4DhcpOpenApiVO`

NewVirtualWanIpv4DhcpOpenApiVO instantiates a new VirtualWanIpv4DhcpOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualWanIpv4DhcpOpenApiVOWithDefaults

`func NewVirtualWanIpv4DhcpOpenApiVOWithDefaults() *VirtualWanIpv4DhcpOpenApiVO`

NewVirtualWanIpv4DhcpOpenApiVOWithDefaults instantiates a new VirtualWanIpv4DhcpOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpOptions

`func (o *VirtualWanIpv4DhcpOpenApiVO) GetDhcpOptions() []CustomDHCPOptions`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *VirtualWanIpv4DhcpOpenApiVO) GetDhcpOptionsOk() (*[]CustomDHCPOptions, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *VirtualWanIpv4DhcpOpenApiVO) SetDhcpOptions(v []CustomDHCPOptions)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *VirtualWanIpv4DhcpOpenApiVO) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetDns1

`func (o *VirtualWanIpv4DhcpOpenApiVO) GetDns1() string`

GetDns1 returns the Dns1 field if non-nil, zero value otherwise.

### GetDns1Ok

`func (o *VirtualWanIpv4DhcpOpenApiVO) GetDns1Ok() (*string, bool)`

GetDns1Ok returns a tuple with the Dns1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns1

`func (o *VirtualWanIpv4DhcpOpenApiVO) SetDns1(v string)`

SetDns1 sets Dns1 field to given value.

### HasDns1

`func (o *VirtualWanIpv4DhcpOpenApiVO) HasDns1() bool`

HasDns1 returns a boolean if a field has been set.

### GetDns2

`func (o *VirtualWanIpv4DhcpOpenApiVO) GetDns2() string`

GetDns2 returns the Dns2 field if non-nil, zero value otherwise.

### GetDns2Ok

`func (o *VirtualWanIpv4DhcpOpenApiVO) GetDns2Ok() (*string, bool)`

GetDns2Ok returns a tuple with the Dns2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns2

`func (o *VirtualWanIpv4DhcpOpenApiVO) SetDns2(v string)`

SetDns2 sets Dns2 field to given value.

### HasDns2

`func (o *VirtualWanIpv4DhcpOpenApiVO) HasDns2() bool`

HasDns2 returns a boolean if a field has been set.

### GetHostname

`func (o *VirtualWanIpv4DhcpOpenApiVO) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VirtualWanIpv4DhcpOpenApiVO) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VirtualWanIpv4DhcpOpenApiVO) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *VirtualWanIpv4DhcpOpenApiVO) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetMtu

`func (o *VirtualWanIpv4DhcpOpenApiVO) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VirtualWanIpv4DhcpOpenApiVO) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VirtualWanIpv4DhcpOpenApiVO) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VirtualWanIpv4DhcpOpenApiVO) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetUnicast

`func (o *VirtualWanIpv4DhcpOpenApiVO) GetUnicast() string`

GetUnicast returns the Unicast field if non-nil, zero value otherwise.

### GetUnicastOk

`func (o *VirtualWanIpv4DhcpOpenApiVO) GetUnicastOk() (*string, bool)`

GetUnicastOk returns a tuple with the Unicast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnicast

`func (o *VirtualWanIpv4DhcpOpenApiVO) SetUnicast(v string)`

SetUnicast sets Unicast field to given value.

### HasUnicast

`func (o *VirtualWanIpv4DhcpOpenApiVO) HasUnicast() bool`

HasUnicast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


