# Ipv6TunnelOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsType** | **int32** | DNS Address Type should be a value as follows: 0: Get from ISP Dynamically; 1: Static. | 
**PrimaryDns** | Pointer to **string** | It takes effect when [dnsType] is 1. | [optional] 

## Methods

### NewIpv6TunnelOpenApiVO

`func NewIpv6TunnelOpenApiVO(dnsType int32, ) *Ipv6TunnelOpenApiVO`

NewIpv6TunnelOpenApiVO instantiates a new Ipv6TunnelOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6TunnelOpenApiVOWithDefaults

`func NewIpv6TunnelOpenApiVOWithDefaults() *Ipv6TunnelOpenApiVO`

NewIpv6TunnelOpenApiVOWithDefaults instantiates a new Ipv6TunnelOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsType

`func (o *Ipv6TunnelOpenApiVO) GetDnsType() int32`

GetDnsType returns the DnsType field if non-nil, zero value otherwise.

### GetDnsTypeOk

`func (o *Ipv6TunnelOpenApiVO) GetDnsTypeOk() (*int32, bool)`

GetDnsTypeOk returns a tuple with the DnsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsType

`func (o *Ipv6TunnelOpenApiVO) SetDnsType(v int32)`

SetDnsType sets DnsType field to given value.


### GetPrimaryDns

`func (o *Ipv6TunnelOpenApiVO) GetPrimaryDns() string`

GetPrimaryDns returns the PrimaryDns field if non-nil, zero value otherwise.

### GetPrimaryDnsOk

`func (o *Ipv6TunnelOpenApiVO) GetPrimaryDnsOk() (*string, bool)`

GetPrimaryDnsOk returns a tuple with the PrimaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDns

`func (o *Ipv6TunnelOpenApiVO) SetPrimaryDns(v string)`

SetPrimaryDns sets PrimaryDns field to given value.

### HasPrimaryDns

`func (o *Ipv6TunnelOpenApiVO) HasPrimaryDns() bool`

HasPrimaryDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


