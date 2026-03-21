# LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnsv6** | Pointer to **int32** | DHCP Name Server, should be a value as follows: 0: \&quot;auto\&quot;; 1: \&quot;manual\&quot; | [optional] 
**PortUuid** | Pointer to **string** | The port UUID of WAN | [optional] 
**PreId** | Pointer to **int32** | Prefix ID should be within the range of 0-127 | [optional] 
**PreType** | Pointer to **int32** | Prefix type should be a value as follows:  0: \&quot;manual\&quot;; 1: \&quot;get from PD\&quot; | [optional] 
**Prefix** | Pointer to **string** | Address prefix | [optional] 
**PriDns** | Pointer to **string** | Primary DHCP Name Server, only effective for DNSv6 \&quot;manual\&quot; | [optional] 
**SndDns** | Pointer to **string** | Secondary DHCP Name Server, only effective for dnsv6 \&quot;manual\&quot; | [optional] 

## Methods

### NewLanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode

`func NewLanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode() *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode`

NewLanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode instantiates a new LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSModeWithDefaults

`func NewLanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSModeWithDefaults() *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode`

NewLanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSModeWithDefaults instantiates a new LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsv6

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) GetDnsv6() int32`

GetDnsv6 returns the Dnsv6 field if non-nil, zero value otherwise.

### GetDnsv6Ok

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) GetDnsv6Ok() (*int32, bool)`

GetDnsv6Ok returns a tuple with the Dnsv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsv6

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) SetDnsv6(v int32)`

SetDnsv6 sets Dnsv6 field to given value.

### HasDnsv6

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) HasDnsv6() bool`

HasDnsv6 returns a boolean if a field has been set.

### GetPortUuid

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) GetPortUuid() string`

GetPortUuid returns the PortUuid field if non-nil, zero value otherwise.

### GetPortUuidOk

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) GetPortUuidOk() (*string, bool)`

GetPortUuidOk returns a tuple with the PortUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUuid

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) SetPortUuid(v string)`

SetPortUuid sets PortUuid field to given value.

### HasPortUuid

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) HasPortUuid() bool`

HasPortUuid returns a boolean if a field has been set.

### GetPreId

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) GetPreId() int32`

GetPreId returns the PreId field if non-nil, zero value otherwise.

### GetPreIdOk

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) GetPreIdOk() (*int32, bool)`

GetPreIdOk returns a tuple with the PreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreId

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) SetPreId(v int32)`

SetPreId sets PreId field to given value.

### HasPreId

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) HasPreId() bool`

HasPreId returns a boolean if a field has been set.

### GetPreType

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) GetPreType() int32`

GetPreType returns the PreType field if non-nil, zero value otherwise.

### GetPreTypeOk

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) GetPreTypeOk() (*int32, bool)`

GetPreTypeOk returns a tuple with the PreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreType

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) SetPreType(v int32)`

SetPreType sets PreType field to given value.

### HasPreType

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) HasPreType() bool`

HasPreType returns a boolean if a field has been set.

### GetPrefix

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPriDns

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) GetPriDns() string`

GetPriDns returns the PriDns field if non-nil, zero value otherwise.

### GetPriDnsOk

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) GetPriDnsOk() (*string, bool)`

GetPriDnsOk returns a tuple with the PriDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriDns

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) SetPriDns(v string)`

SetPriDns sets PriDns field to given value.

### HasPriDns

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) HasPriDns() bool`

HasPriDns returns a boolean if a field has been set.

### GetSndDns

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) GetSndDns() string`

GetSndDns returns the SndDns field if non-nil, zero value otherwise.

### GetSndDnsOk

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) GetSndDnsOk() (*string, bool)`

GetSndDnsOk returns a tuple with the SndDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSndDns

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) SetSndDns(v string)`

SetSndDns sets SndDns field to given value.

### HasSndDns

`func (o *LanNetworkProtoSLAACStatelessDHCPModeOrSLAACRDNSSMode) HasSndDns() bool`

HasSndDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


