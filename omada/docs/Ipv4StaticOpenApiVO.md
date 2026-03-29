# Ipv4StaticOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultGateway** | Pointer to **string** |  | [optional] 
**IpAddress** | **string** | IP address | 
**Mtu** | **int32** | 576-1500, default:1500 | 
**PrimaryDns** | Pointer to **string** | Primary DNS | [optional] 
**SecondaryDns** | Pointer to **string** | Secondary DNS | [optional] 
**SubnetMask** | **string** |  | 
**WanMultipleIps** | Pointer to [**[]WanMultipleIpOpenApiVO**](WanMultipleIpOpenApiVO.md) |  | [optional] 

## Methods

### NewIpv4StaticOpenApiVO

`func NewIpv4StaticOpenApiVO(ipAddress string, mtu int32, subnetMask string, ) *Ipv4StaticOpenApiVO`

NewIpv4StaticOpenApiVO instantiates a new Ipv4StaticOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv4StaticOpenApiVOWithDefaults

`func NewIpv4StaticOpenApiVOWithDefaults() *Ipv4StaticOpenApiVO`

NewIpv4StaticOpenApiVOWithDefaults instantiates a new Ipv4StaticOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultGateway

`func (o *Ipv4StaticOpenApiVO) GetDefaultGateway() string`

GetDefaultGateway returns the DefaultGateway field if non-nil, zero value otherwise.

### GetDefaultGatewayOk

`func (o *Ipv4StaticOpenApiVO) GetDefaultGatewayOk() (*string, bool)`

GetDefaultGatewayOk returns a tuple with the DefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateway

`func (o *Ipv4StaticOpenApiVO) SetDefaultGateway(v string)`

SetDefaultGateway sets DefaultGateway field to given value.

### HasDefaultGateway

`func (o *Ipv4StaticOpenApiVO) HasDefaultGateway() bool`

HasDefaultGateway returns a boolean if a field has been set.

### GetIpAddress

`func (o *Ipv4StaticOpenApiVO) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Ipv4StaticOpenApiVO) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Ipv4StaticOpenApiVO) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetMtu

`func (o *Ipv4StaticOpenApiVO) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *Ipv4StaticOpenApiVO) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *Ipv4StaticOpenApiVO) SetMtu(v int32)`

SetMtu sets Mtu field to given value.


### GetPrimaryDns

`func (o *Ipv4StaticOpenApiVO) GetPrimaryDns() string`

GetPrimaryDns returns the PrimaryDns field if non-nil, zero value otherwise.

### GetPrimaryDnsOk

`func (o *Ipv4StaticOpenApiVO) GetPrimaryDnsOk() (*string, bool)`

GetPrimaryDnsOk returns a tuple with the PrimaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDns

`func (o *Ipv4StaticOpenApiVO) SetPrimaryDns(v string)`

SetPrimaryDns sets PrimaryDns field to given value.

### HasPrimaryDns

`func (o *Ipv4StaticOpenApiVO) HasPrimaryDns() bool`

HasPrimaryDns returns a boolean if a field has been set.

### GetSecondaryDns

`func (o *Ipv4StaticOpenApiVO) GetSecondaryDns() string`

GetSecondaryDns returns the SecondaryDns field if non-nil, zero value otherwise.

### GetSecondaryDnsOk

`func (o *Ipv4StaticOpenApiVO) GetSecondaryDnsOk() (*string, bool)`

GetSecondaryDnsOk returns a tuple with the SecondaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDns

`func (o *Ipv4StaticOpenApiVO) SetSecondaryDns(v string)`

SetSecondaryDns sets SecondaryDns field to given value.

### HasSecondaryDns

`func (o *Ipv4StaticOpenApiVO) HasSecondaryDns() bool`

HasSecondaryDns returns a boolean if a field has been set.

### GetSubnetMask

`func (o *Ipv4StaticOpenApiVO) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *Ipv4StaticOpenApiVO) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *Ipv4StaticOpenApiVO) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.


### GetWanMultipleIps

`func (o *Ipv4StaticOpenApiVO) GetWanMultipleIps() []WanMultipleIpOpenApiVO`

GetWanMultipleIps returns the WanMultipleIps field if non-nil, zero value otherwise.

### GetWanMultipleIpsOk

`func (o *Ipv4StaticOpenApiVO) GetWanMultipleIpsOk() (*[]WanMultipleIpOpenApiVO, bool)`

GetWanMultipleIpsOk returns a tuple with the WanMultipleIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanMultipleIps

`func (o *Ipv4StaticOpenApiVO) SetWanMultipleIps(v []WanMultipleIpOpenApiVO)`

SetWanMultipleIps sets WanMultipleIps field to given value.

### HasWanMultipleIps

`func (o *Ipv4StaticOpenApiVO) HasWanMultipleIps() bool`

HasWanMultipleIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


