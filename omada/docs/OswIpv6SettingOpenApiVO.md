# OswIpv6SettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsMode** | Pointer to **int32** | DNS mode for dynamic mode, 0: Get Dynamic DNS, 1: Use the Following DNS Addresses. | [optional] 
**Ipv6Addr** | Pointer to **string** | IPv6 Address for static mode, like 2001:4860:4860::8888 | [optional] 
**Mode** | **string** | Ipv6Setting parameter [mode] should be dynamic: 1 or static: 0 | 
**PrefixLen** | Pointer to **int32** | Prefix Length for static mode, which should be within the range of 1–128 | [optional] 
**PriDns** | Pointer to **string** | Primary DNS, like 2001:4860:4860::8888 | [optional] 
**SndDns** | Pointer to **string** | Second DNS, like 2001:4860:4860::8844 | [optional] 

## Methods

### NewOswIpv6SettingOpenApiVO

`func NewOswIpv6SettingOpenApiVO(mode string, ) *OswIpv6SettingOpenApiVO`

NewOswIpv6SettingOpenApiVO instantiates a new OswIpv6SettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswIpv6SettingOpenApiVOWithDefaults

`func NewOswIpv6SettingOpenApiVOWithDefaults() *OswIpv6SettingOpenApiVO`

NewOswIpv6SettingOpenApiVOWithDefaults instantiates a new OswIpv6SettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsMode

`func (o *OswIpv6SettingOpenApiVO) GetDnsMode() int32`

GetDnsMode returns the DnsMode field if non-nil, zero value otherwise.

### GetDnsModeOk

`func (o *OswIpv6SettingOpenApiVO) GetDnsModeOk() (*int32, bool)`

GetDnsModeOk returns a tuple with the DnsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsMode

`func (o *OswIpv6SettingOpenApiVO) SetDnsMode(v int32)`

SetDnsMode sets DnsMode field to given value.

### HasDnsMode

`func (o *OswIpv6SettingOpenApiVO) HasDnsMode() bool`

HasDnsMode returns a boolean if a field has been set.

### GetIpv6Addr

`func (o *OswIpv6SettingOpenApiVO) GetIpv6Addr() string`

GetIpv6Addr returns the Ipv6Addr field if non-nil, zero value otherwise.

### GetIpv6AddrOk

`func (o *OswIpv6SettingOpenApiVO) GetIpv6AddrOk() (*string, bool)`

GetIpv6AddrOk returns a tuple with the Ipv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addr

`func (o *OswIpv6SettingOpenApiVO) SetIpv6Addr(v string)`

SetIpv6Addr sets Ipv6Addr field to given value.

### HasIpv6Addr

`func (o *OswIpv6SettingOpenApiVO) HasIpv6Addr() bool`

HasIpv6Addr returns a boolean if a field has been set.

### GetMode

`func (o *OswIpv6SettingOpenApiVO) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OswIpv6SettingOpenApiVO) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OswIpv6SettingOpenApiVO) SetMode(v string)`

SetMode sets Mode field to given value.


### GetPrefixLen

`func (o *OswIpv6SettingOpenApiVO) GetPrefixLen() int32`

GetPrefixLen returns the PrefixLen field if non-nil, zero value otherwise.

### GetPrefixLenOk

`func (o *OswIpv6SettingOpenApiVO) GetPrefixLenOk() (*int32, bool)`

GetPrefixLenOk returns a tuple with the PrefixLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLen

`func (o *OswIpv6SettingOpenApiVO) SetPrefixLen(v int32)`

SetPrefixLen sets PrefixLen field to given value.

### HasPrefixLen

`func (o *OswIpv6SettingOpenApiVO) HasPrefixLen() bool`

HasPrefixLen returns a boolean if a field has been set.

### GetPriDns

`func (o *OswIpv6SettingOpenApiVO) GetPriDns() string`

GetPriDns returns the PriDns field if non-nil, zero value otherwise.

### GetPriDnsOk

`func (o *OswIpv6SettingOpenApiVO) GetPriDnsOk() (*string, bool)`

GetPriDnsOk returns a tuple with the PriDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriDns

`func (o *OswIpv6SettingOpenApiVO) SetPriDns(v string)`

SetPriDns sets PriDns field to given value.

### HasPriDns

`func (o *OswIpv6SettingOpenApiVO) HasPriDns() bool`

HasPriDns returns a boolean if a field has been set.

### GetSndDns

`func (o *OswIpv6SettingOpenApiVO) GetSndDns() string`

GetSndDns returns the SndDns field if non-nil, zero value otherwise.

### GetSndDnsOk

`func (o *OswIpv6SettingOpenApiVO) GetSndDnsOk() (*string, bool)`

GetSndDnsOk returns a tuple with the SndDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSndDns

`func (o *OswIpv6SettingOpenApiVO) SetSndDns(v string)`

SetSndDns sets SndDns field to given value.

### HasSndDns

`func (o *OswIpv6SettingOpenApiVO) HasSndDns() bool`

HasSndDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


