# DynamicIpv6SettingEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsMode** | **int32** | only for \&quot;dynamic\&quot; mode (0: Get Dynamic DNS; 1: Use the Following DNS Addresses) | 
**PriDns** | Pointer to **string** | primary DNS address, pattern: ipv6 address. This value is required when dnsMode is 1. | [optional] 
**SndDns** | Pointer to **string** | secondary DNS address, pattern: ipv6 address. This value is optional when dnsMode is 1. | [optional] 

## Methods

### NewDynamicIpv6SettingEntity

`func NewDynamicIpv6SettingEntity(dnsMode int32, ) *DynamicIpv6SettingEntity`

NewDynamicIpv6SettingEntity instantiates a new DynamicIpv6SettingEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicIpv6SettingEntityWithDefaults

`func NewDynamicIpv6SettingEntityWithDefaults() *DynamicIpv6SettingEntity`

NewDynamicIpv6SettingEntityWithDefaults instantiates a new DynamicIpv6SettingEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsMode

`func (o *DynamicIpv6SettingEntity) GetDnsMode() int32`

GetDnsMode returns the DnsMode field if non-nil, zero value otherwise.

### GetDnsModeOk

`func (o *DynamicIpv6SettingEntity) GetDnsModeOk() (*int32, bool)`

GetDnsModeOk returns a tuple with the DnsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsMode

`func (o *DynamicIpv6SettingEntity) SetDnsMode(v int32)`

SetDnsMode sets DnsMode field to given value.


### GetPriDns

`func (o *DynamicIpv6SettingEntity) GetPriDns() string`

GetPriDns returns the PriDns field if non-nil, zero value otherwise.

### GetPriDnsOk

`func (o *DynamicIpv6SettingEntity) GetPriDnsOk() (*string, bool)`

GetPriDnsOk returns a tuple with the PriDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriDns

`func (o *DynamicIpv6SettingEntity) SetPriDns(v string)`

SetPriDns sets PriDns field to given value.

### HasPriDns

`func (o *DynamicIpv6SettingEntity) HasPriDns() bool`

HasPriDns returns a boolean if a field has been set.

### GetSndDns

`func (o *DynamicIpv6SettingEntity) GetSndDns() string`

GetSndDns returns the SndDns field if non-nil, zero value otherwise.

### GetSndDnsOk

`func (o *DynamicIpv6SettingEntity) GetSndDnsOk() (*string, bool)`

GetSndDnsOk returns a tuple with the SndDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSndDns

`func (o *DynamicIpv6SettingEntity) SetSndDns(v string)`

SetSndDns sets SndDns field to given value.

### HasSndDns

`func (o *DynamicIpv6SettingEntity) HasSndDns() bool`

HasSndDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


