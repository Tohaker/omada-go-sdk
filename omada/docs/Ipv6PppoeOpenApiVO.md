# Ipv6PppoeOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsType** | **int32** | DNS Address Type should be a value as follows: 0: dynamic, 1: static. | 
**GetIpv6Type** | **int32** | It should be a value as follows: 0: SLAAC, 1: DHCPv6, 2: specified, 3: auto, 4:non_address, which ranges from 48 ~ 64. | 
**Password** | **string** |  | 
**PppShare** | **bool** | Share the same PPPoE session with IPv4 | 
**PrefixEnable** | **bool** |  | 
**PrefixSize** | Pointer to **int32** | It is required when [prefixEnable] is true. | [optional] 
**PrimaryDns** | Pointer to **string** | It is required when [dnsType] is 1. | [optional] 
**SecondaryDns** | Pointer to **string** | It takes effect when [dnsType] is 1. | [optional] 
**SpecificIp** | Pointer to **string** | It is required when [getIpv6Type] is 2. | [optional] 
**UserName** | **string** |  | 

## Methods

### NewIpv6PppoeOpenApiVO

`func NewIpv6PppoeOpenApiVO(dnsType int32, getIpv6Type int32, password string, pppShare bool, prefixEnable bool, userName string, ) *Ipv6PppoeOpenApiVO`

NewIpv6PppoeOpenApiVO instantiates a new Ipv6PppoeOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6PppoeOpenApiVOWithDefaults

`func NewIpv6PppoeOpenApiVOWithDefaults() *Ipv6PppoeOpenApiVO`

NewIpv6PppoeOpenApiVOWithDefaults instantiates a new Ipv6PppoeOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsType

`func (o *Ipv6PppoeOpenApiVO) GetDnsType() int32`

GetDnsType returns the DnsType field if non-nil, zero value otherwise.

### GetDnsTypeOk

`func (o *Ipv6PppoeOpenApiVO) GetDnsTypeOk() (*int32, bool)`

GetDnsTypeOk returns a tuple with the DnsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsType

`func (o *Ipv6PppoeOpenApiVO) SetDnsType(v int32)`

SetDnsType sets DnsType field to given value.


### GetGetIpv6Type

`func (o *Ipv6PppoeOpenApiVO) GetGetIpv6Type() int32`

GetGetIpv6Type returns the GetIpv6Type field if non-nil, zero value otherwise.

### GetGetIpv6TypeOk

`func (o *Ipv6PppoeOpenApiVO) GetGetIpv6TypeOk() (*int32, bool)`

GetGetIpv6TypeOk returns a tuple with the GetIpv6Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetIpv6Type

`func (o *Ipv6PppoeOpenApiVO) SetGetIpv6Type(v int32)`

SetGetIpv6Type sets GetIpv6Type field to given value.


### GetPassword

`func (o *Ipv6PppoeOpenApiVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Ipv6PppoeOpenApiVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Ipv6PppoeOpenApiVO) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPppShare

`func (o *Ipv6PppoeOpenApiVO) GetPppShare() bool`

GetPppShare returns the PppShare field if non-nil, zero value otherwise.

### GetPppShareOk

`func (o *Ipv6PppoeOpenApiVO) GetPppShareOk() (*bool, bool)`

GetPppShareOk returns a tuple with the PppShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPppShare

`func (o *Ipv6PppoeOpenApiVO) SetPppShare(v bool)`

SetPppShare sets PppShare field to given value.


### GetPrefixEnable

`func (o *Ipv6PppoeOpenApiVO) GetPrefixEnable() bool`

GetPrefixEnable returns the PrefixEnable field if non-nil, zero value otherwise.

### GetPrefixEnableOk

`func (o *Ipv6PppoeOpenApiVO) GetPrefixEnableOk() (*bool, bool)`

GetPrefixEnableOk returns a tuple with the PrefixEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixEnable

`func (o *Ipv6PppoeOpenApiVO) SetPrefixEnable(v bool)`

SetPrefixEnable sets PrefixEnable field to given value.


### GetPrefixSize

`func (o *Ipv6PppoeOpenApiVO) GetPrefixSize() int32`

GetPrefixSize returns the PrefixSize field if non-nil, zero value otherwise.

### GetPrefixSizeOk

`func (o *Ipv6PppoeOpenApiVO) GetPrefixSizeOk() (*int32, bool)`

GetPrefixSizeOk returns a tuple with the PrefixSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSize

`func (o *Ipv6PppoeOpenApiVO) SetPrefixSize(v int32)`

SetPrefixSize sets PrefixSize field to given value.

### HasPrefixSize

`func (o *Ipv6PppoeOpenApiVO) HasPrefixSize() bool`

HasPrefixSize returns a boolean if a field has been set.

### GetPrimaryDns

`func (o *Ipv6PppoeOpenApiVO) GetPrimaryDns() string`

GetPrimaryDns returns the PrimaryDns field if non-nil, zero value otherwise.

### GetPrimaryDnsOk

`func (o *Ipv6PppoeOpenApiVO) GetPrimaryDnsOk() (*string, bool)`

GetPrimaryDnsOk returns a tuple with the PrimaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDns

`func (o *Ipv6PppoeOpenApiVO) SetPrimaryDns(v string)`

SetPrimaryDns sets PrimaryDns field to given value.

### HasPrimaryDns

`func (o *Ipv6PppoeOpenApiVO) HasPrimaryDns() bool`

HasPrimaryDns returns a boolean if a field has been set.

### GetSecondaryDns

`func (o *Ipv6PppoeOpenApiVO) GetSecondaryDns() string`

GetSecondaryDns returns the SecondaryDns field if non-nil, zero value otherwise.

### GetSecondaryDnsOk

`func (o *Ipv6PppoeOpenApiVO) GetSecondaryDnsOk() (*string, bool)`

GetSecondaryDnsOk returns a tuple with the SecondaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDns

`func (o *Ipv6PppoeOpenApiVO) SetSecondaryDns(v string)`

SetSecondaryDns sets SecondaryDns field to given value.

### HasSecondaryDns

`func (o *Ipv6PppoeOpenApiVO) HasSecondaryDns() bool`

HasSecondaryDns returns a boolean if a field has been set.

### GetSpecificIp

`func (o *Ipv6PppoeOpenApiVO) GetSpecificIp() string`

GetSpecificIp returns the SpecificIp field if non-nil, zero value otherwise.

### GetSpecificIpOk

`func (o *Ipv6PppoeOpenApiVO) GetSpecificIpOk() (*string, bool)`

GetSpecificIpOk returns a tuple with the SpecificIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificIp

`func (o *Ipv6PppoeOpenApiVO) SetSpecificIp(v string)`

SetSpecificIp sets SpecificIp field to given value.

### HasSpecificIp

`func (o *Ipv6PppoeOpenApiVO) HasSpecificIp() bool`

HasSpecificIp returns a boolean if a field has been set.

### GetUserName

`func (o *Ipv6PppoeOpenApiVO) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *Ipv6PppoeOpenApiVO) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *Ipv6PppoeOpenApiVO) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


