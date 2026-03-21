# Ipv4DhcpOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpOptions** | Pointer to [**[]WanDhcpOptionOpenApiVO**](WanDhcpOptionOpenApiVO.md) |  | [optional] 
**Hostname** | Pointer to **string** | Host name | [optional] 
**Mtu** | **int32** | 576-1500, default:1500. | 
**PrimaryDns** | Pointer to **string** | Primary DNS | [optional] 
**SecondaryDns** | Pointer to **string** | Secondary DNS | [optional] 
**UnicastDhcp** | Pointer to **bool** | Unicast DHCP | [optional] 
**WanMultipleIps** | Pointer to [**[]WanMultipleIpOpenApiVO**](WanMultipleIpOpenApiVO.md) |  | [optional] 

## Methods

### NewIpv4DhcpOpenApiVO

`func NewIpv4DhcpOpenApiVO(mtu int32, ) *Ipv4DhcpOpenApiVO`

NewIpv4DhcpOpenApiVO instantiates a new Ipv4DhcpOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv4DhcpOpenApiVOWithDefaults

`func NewIpv4DhcpOpenApiVOWithDefaults() *Ipv4DhcpOpenApiVO`

NewIpv4DhcpOpenApiVOWithDefaults instantiates a new Ipv4DhcpOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpOptions

`func (o *Ipv4DhcpOpenApiVO) GetDhcpOptions() []WanDhcpOptionOpenApiVO`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *Ipv4DhcpOpenApiVO) GetDhcpOptionsOk() (*[]WanDhcpOptionOpenApiVO, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *Ipv4DhcpOpenApiVO) SetDhcpOptions(v []WanDhcpOptionOpenApiVO)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *Ipv4DhcpOpenApiVO) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetHostname

`func (o *Ipv4DhcpOpenApiVO) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Ipv4DhcpOpenApiVO) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Ipv4DhcpOpenApiVO) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *Ipv4DhcpOpenApiVO) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetMtu

`func (o *Ipv4DhcpOpenApiVO) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *Ipv4DhcpOpenApiVO) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *Ipv4DhcpOpenApiVO) SetMtu(v int32)`

SetMtu sets Mtu field to given value.


### GetPrimaryDns

`func (o *Ipv4DhcpOpenApiVO) GetPrimaryDns() string`

GetPrimaryDns returns the PrimaryDns field if non-nil, zero value otherwise.

### GetPrimaryDnsOk

`func (o *Ipv4DhcpOpenApiVO) GetPrimaryDnsOk() (*string, bool)`

GetPrimaryDnsOk returns a tuple with the PrimaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDns

`func (o *Ipv4DhcpOpenApiVO) SetPrimaryDns(v string)`

SetPrimaryDns sets PrimaryDns field to given value.

### HasPrimaryDns

`func (o *Ipv4DhcpOpenApiVO) HasPrimaryDns() bool`

HasPrimaryDns returns a boolean if a field has been set.

### GetSecondaryDns

`func (o *Ipv4DhcpOpenApiVO) GetSecondaryDns() string`

GetSecondaryDns returns the SecondaryDns field if non-nil, zero value otherwise.

### GetSecondaryDnsOk

`func (o *Ipv4DhcpOpenApiVO) GetSecondaryDnsOk() (*string, bool)`

GetSecondaryDnsOk returns a tuple with the SecondaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDns

`func (o *Ipv4DhcpOpenApiVO) SetSecondaryDns(v string)`

SetSecondaryDns sets SecondaryDns field to given value.

### HasSecondaryDns

`func (o *Ipv4DhcpOpenApiVO) HasSecondaryDns() bool`

HasSecondaryDns returns a boolean if a field has been set.

### GetUnicastDhcp

`func (o *Ipv4DhcpOpenApiVO) GetUnicastDhcp() bool`

GetUnicastDhcp returns the UnicastDhcp field if non-nil, zero value otherwise.

### GetUnicastDhcpOk

`func (o *Ipv4DhcpOpenApiVO) GetUnicastDhcpOk() (*bool, bool)`

GetUnicastDhcpOk returns a tuple with the UnicastDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnicastDhcp

`func (o *Ipv4DhcpOpenApiVO) SetUnicastDhcp(v bool)`

SetUnicastDhcp sets UnicastDhcp field to given value.

### HasUnicastDhcp

`func (o *Ipv4DhcpOpenApiVO) HasUnicastDhcp() bool`

HasUnicastDhcp returns a boolean if a field has been set.

### GetWanMultipleIps

`func (o *Ipv4DhcpOpenApiVO) GetWanMultipleIps() []WanMultipleIpOpenApiVO`

GetWanMultipleIps returns the WanMultipleIps field if non-nil, zero value otherwise.

### GetWanMultipleIpsOk

`func (o *Ipv4DhcpOpenApiVO) GetWanMultipleIpsOk() (*[]WanMultipleIpOpenApiVO, bool)`

GetWanMultipleIpsOk returns a tuple with the WanMultipleIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanMultipleIps

`func (o *Ipv4DhcpOpenApiVO) SetWanMultipleIps(v []WanMultipleIpOpenApiVO)`

SetWanMultipleIps sets WanMultipleIps field to given value.

### HasWanMultipleIps

`func (o *Ipv4DhcpOpenApiVO) HasWanMultipleIps() bool`

HasWanMultipleIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


