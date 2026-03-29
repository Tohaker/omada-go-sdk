# Ipv6DynamicOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsType** | **int32** | DNS Address Type, 0: dynamic, 1: static. | 
**GetIpv6Type** | **int32** | It should be a value as follows: 0: SLAAC; 1: DHCPv6; 2: specified; 3: auto; 4:non-address. | 
**PrefixEnable** | **bool** |  | 
**PrefixSize** | Pointer to **int32** | It is required when [prefixEnable] is true, which ranges from 48 ~ 64. | [optional] 
**PrimaryDns** | Pointer to **string** | It is required when [dnsType] is 1. | [optional] 
**SecondaryDns** | Pointer to **string** | It takes effect when [dnsType] is 1 | [optional] 

## Methods

### NewIpv6DynamicOpenApiVO

`func NewIpv6DynamicOpenApiVO(dnsType int32, getIpv6Type int32, prefixEnable bool, ) *Ipv6DynamicOpenApiVO`

NewIpv6DynamicOpenApiVO instantiates a new Ipv6DynamicOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6DynamicOpenApiVOWithDefaults

`func NewIpv6DynamicOpenApiVOWithDefaults() *Ipv6DynamicOpenApiVO`

NewIpv6DynamicOpenApiVOWithDefaults instantiates a new Ipv6DynamicOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsType

`func (o *Ipv6DynamicOpenApiVO) GetDnsType() int32`

GetDnsType returns the DnsType field if non-nil, zero value otherwise.

### GetDnsTypeOk

`func (o *Ipv6DynamicOpenApiVO) GetDnsTypeOk() (*int32, bool)`

GetDnsTypeOk returns a tuple with the DnsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsType

`func (o *Ipv6DynamicOpenApiVO) SetDnsType(v int32)`

SetDnsType sets DnsType field to given value.


### GetGetIpv6Type

`func (o *Ipv6DynamicOpenApiVO) GetGetIpv6Type() int32`

GetGetIpv6Type returns the GetIpv6Type field if non-nil, zero value otherwise.

### GetGetIpv6TypeOk

`func (o *Ipv6DynamicOpenApiVO) GetGetIpv6TypeOk() (*int32, bool)`

GetGetIpv6TypeOk returns a tuple with the GetIpv6Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetIpv6Type

`func (o *Ipv6DynamicOpenApiVO) SetGetIpv6Type(v int32)`

SetGetIpv6Type sets GetIpv6Type field to given value.


### GetPrefixEnable

`func (o *Ipv6DynamicOpenApiVO) GetPrefixEnable() bool`

GetPrefixEnable returns the PrefixEnable field if non-nil, zero value otherwise.

### GetPrefixEnableOk

`func (o *Ipv6DynamicOpenApiVO) GetPrefixEnableOk() (*bool, bool)`

GetPrefixEnableOk returns a tuple with the PrefixEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixEnable

`func (o *Ipv6DynamicOpenApiVO) SetPrefixEnable(v bool)`

SetPrefixEnable sets PrefixEnable field to given value.


### GetPrefixSize

`func (o *Ipv6DynamicOpenApiVO) GetPrefixSize() int32`

GetPrefixSize returns the PrefixSize field if non-nil, zero value otherwise.

### GetPrefixSizeOk

`func (o *Ipv6DynamicOpenApiVO) GetPrefixSizeOk() (*int32, bool)`

GetPrefixSizeOk returns a tuple with the PrefixSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSize

`func (o *Ipv6DynamicOpenApiVO) SetPrefixSize(v int32)`

SetPrefixSize sets PrefixSize field to given value.

### HasPrefixSize

`func (o *Ipv6DynamicOpenApiVO) HasPrefixSize() bool`

HasPrefixSize returns a boolean if a field has been set.

### GetPrimaryDns

`func (o *Ipv6DynamicOpenApiVO) GetPrimaryDns() string`

GetPrimaryDns returns the PrimaryDns field if non-nil, zero value otherwise.

### GetPrimaryDnsOk

`func (o *Ipv6DynamicOpenApiVO) GetPrimaryDnsOk() (*string, bool)`

GetPrimaryDnsOk returns a tuple with the PrimaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDns

`func (o *Ipv6DynamicOpenApiVO) SetPrimaryDns(v string)`

SetPrimaryDns sets PrimaryDns field to given value.

### HasPrimaryDns

`func (o *Ipv6DynamicOpenApiVO) HasPrimaryDns() bool`

HasPrimaryDns returns a boolean if a field has been set.

### GetSecondaryDns

`func (o *Ipv6DynamicOpenApiVO) GetSecondaryDns() string`

GetSecondaryDns returns the SecondaryDns field if non-nil, zero value otherwise.

### GetSecondaryDnsOk

`func (o *Ipv6DynamicOpenApiVO) GetSecondaryDnsOk() (*string, bool)`

GetSecondaryDnsOk returns a tuple with the SecondaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDns

`func (o *Ipv6DynamicOpenApiVO) SetSecondaryDns(v string)`

SetSecondaryDns sets SecondaryDns field to given value.

### HasSecondaryDns

`func (o *Ipv6DynamicOpenApiVO) HasSecondaryDns() bool`

HasSecondaryDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


