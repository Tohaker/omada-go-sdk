# StaticIpv6SettingEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultGate** | Pointer to **string** | default gateway, only for static mode | [optional] 
**Ipv6Addr** | **string** | ipv6 address, only required for static mode | 
**PrefixLen** | **int32** | prefixLen of the ipv6 address, only required for static mode | 
**PriDns** | Pointer to **string** | primary DNS address, pattern: ipv6 address. | [optional] 
**SndDns** | Pointer to **string** | secondary DNS address, pattern: ipv6 address. | [optional] 

## Methods

### NewStaticIpv6SettingEntity

`func NewStaticIpv6SettingEntity(ipv6Addr string, prefixLen int32, ) *StaticIpv6SettingEntity`

NewStaticIpv6SettingEntity instantiates a new StaticIpv6SettingEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticIpv6SettingEntityWithDefaults

`func NewStaticIpv6SettingEntityWithDefaults() *StaticIpv6SettingEntity`

NewStaticIpv6SettingEntityWithDefaults instantiates a new StaticIpv6SettingEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultGate

`func (o *StaticIpv6SettingEntity) GetDefaultGate() string`

GetDefaultGate returns the DefaultGate field if non-nil, zero value otherwise.

### GetDefaultGateOk

`func (o *StaticIpv6SettingEntity) GetDefaultGateOk() (*string, bool)`

GetDefaultGateOk returns a tuple with the DefaultGate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGate

`func (o *StaticIpv6SettingEntity) SetDefaultGate(v string)`

SetDefaultGate sets DefaultGate field to given value.

### HasDefaultGate

`func (o *StaticIpv6SettingEntity) HasDefaultGate() bool`

HasDefaultGate returns a boolean if a field has been set.

### GetIpv6Addr

`func (o *StaticIpv6SettingEntity) GetIpv6Addr() string`

GetIpv6Addr returns the Ipv6Addr field if non-nil, zero value otherwise.

### GetIpv6AddrOk

`func (o *StaticIpv6SettingEntity) GetIpv6AddrOk() (*string, bool)`

GetIpv6AddrOk returns a tuple with the Ipv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addr

`func (o *StaticIpv6SettingEntity) SetIpv6Addr(v string)`

SetIpv6Addr sets Ipv6Addr field to given value.


### GetPrefixLen

`func (o *StaticIpv6SettingEntity) GetPrefixLen() int32`

GetPrefixLen returns the PrefixLen field if non-nil, zero value otherwise.

### GetPrefixLenOk

`func (o *StaticIpv6SettingEntity) GetPrefixLenOk() (*int32, bool)`

GetPrefixLenOk returns a tuple with the PrefixLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLen

`func (o *StaticIpv6SettingEntity) SetPrefixLen(v int32)`

SetPrefixLen sets PrefixLen field to given value.


### GetPriDns

`func (o *StaticIpv6SettingEntity) GetPriDns() string`

GetPriDns returns the PriDns field if non-nil, zero value otherwise.

### GetPriDnsOk

`func (o *StaticIpv6SettingEntity) GetPriDnsOk() (*string, bool)`

GetPriDnsOk returns a tuple with the PriDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriDns

`func (o *StaticIpv6SettingEntity) SetPriDns(v string)`

SetPriDns sets PriDns field to given value.

### HasPriDns

`func (o *StaticIpv6SettingEntity) HasPriDns() bool`

HasPriDns returns a boolean if a field has been set.

### GetSndDns

`func (o *StaticIpv6SettingEntity) GetSndDns() string`

GetSndDns returns the SndDns field if non-nil, zero value otherwise.

### GetSndDnsOk

`func (o *StaticIpv6SettingEntity) GetSndDnsOk() (*string, bool)`

GetSndDnsOk returns a tuple with the SndDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSndDns

`func (o *StaticIpv6SettingEntity) SetSndDns(v string)`

SetSndDns sets SndDns field to given value.

### HasSndDns

`func (o *StaticIpv6SettingEntity) HasSndDns() bool`

HasSndDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


