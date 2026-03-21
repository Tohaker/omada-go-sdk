# OswDhcpServerOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | Pointer to **string** | Gateway IP, like 192.168.0.1 | [optional] 
**Ip** | **string** | DHCP Server IP, like 192.168.0.1. | 
**IpRangePool** | Pointer to [**[]OswDhcpServerRangeOpenApiVO**](OswDhcpServerRangeOpenApiVO.md) | The list of DHCP Range | [optional] 
**Leasetime** | **int32** | Lease time should be within the range of 2–2880 | 
**Netmask** | **string** | Parameter [netmask] should not within the range of 1-30 | 
**Option138** | Pointer to **string** | option138 ip, like 192.168.0.1 | [optional] 
**Options** | Pointer to [**[]SwitchCustomDHCPOptions**](SwitchCustomDHCPOptions.md) | Custom DHCP options. | [optional] 
**PriDns** | Pointer to **string** | Primary DNS, like 192.0.0.1 | [optional] 
**SndDns** | Pointer to **string** | Second DNS, like 8.8.8.8 | [optional] 
**VrfId** | Pointer to **string** | VRF ID | [optional] 

## Methods

### NewOswDhcpServerOpenApiVO

`func NewOswDhcpServerOpenApiVO(ip string, leasetime int32, netmask string, ) *OswDhcpServerOpenApiVO`

NewOswDhcpServerOpenApiVO instantiates a new OswDhcpServerOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswDhcpServerOpenApiVOWithDefaults

`func NewOswDhcpServerOpenApiVOWithDefaults() *OswDhcpServerOpenApiVO`

NewOswDhcpServerOpenApiVOWithDefaults instantiates a new OswDhcpServerOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *OswDhcpServerOpenApiVO) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *OswDhcpServerOpenApiVO) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *OswDhcpServerOpenApiVO) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *OswDhcpServerOpenApiVO) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIp

`func (o *OswDhcpServerOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OswDhcpServerOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OswDhcpServerOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.


### GetIpRangePool

`func (o *OswDhcpServerOpenApiVO) GetIpRangePool() []OswDhcpServerRangeOpenApiVO`

GetIpRangePool returns the IpRangePool field if non-nil, zero value otherwise.

### GetIpRangePoolOk

`func (o *OswDhcpServerOpenApiVO) GetIpRangePoolOk() (*[]OswDhcpServerRangeOpenApiVO, bool)`

GetIpRangePoolOk returns a tuple with the IpRangePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRangePool

`func (o *OswDhcpServerOpenApiVO) SetIpRangePool(v []OswDhcpServerRangeOpenApiVO)`

SetIpRangePool sets IpRangePool field to given value.

### HasIpRangePool

`func (o *OswDhcpServerOpenApiVO) HasIpRangePool() bool`

HasIpRangePool returns a boolean if a field has been set.

### GetLeasetime

`func (o *OswDhcpServerOpenApiVO) GetLeasetime() int32`

GetLeasetime returns the Leasetime field if non-nil, zero value otherwise.

### GetLeasetimeOk

`func (o *OswDhcpServerOpenApiVO) GetLeasetimeOk() (*int32, bool)`

GetLeasetimeOk returns a tuple with the Leasetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasetime

`func (o *OswDhcpServerOpenApiVO) SetLeasetime(v int32)`

SetLeasetime sets Leasetime field to given value.


### GetNetmask

`func (o *OswDhcpServerOpenApiVO) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *OswDhcpServerOpenApiVO) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *OswDhcpServerOpenApiVO) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.


### GetOption138

`func (o *OswDhcpServerOpenApiVO) GetOption138() string`

GetOption138 returns the Option138 field if non-nil, zero value otherwise.

### GetOption138Ok

`func (o *OswDhcpServerOpenApiVO) GetOption138Ok() (*string, bool)`

GetOption138Ok returns a tuple with the Option138 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption138

`func (o *OswDhcpServerOpenApiVO) SetOption138(v string)`

SetOption138 sets Option138 field to given value.

### HasOption138

`func (o *OswDhcpServerOpenApiVO) HasOption138() bool`

HasOption138 returns a boolean if a field has been set.

### GetOptions

`func (o *OswDhcpServerOpenApiVO) GetOptions() []SwitchCustomDHCPOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *OswDhcpServerOpenApiVO) GetOptionsOk() (*[]SwitchCustomDHCPOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *OswDhcpServerOpenApiVO) SetOptions(v []SwitchCustomDHCPOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *OswDhcpServerOpenApiVO) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPriDns

`func (o *OswDhcpServerOpenApiVO) GetPriDns() string`

GetPriDns returns the PriDns field if non-nil, zero value otherwise.

### GetPriDnsOk

`func (o *OswDhcpServerOpenApiVO) GetPriDnsOk() (*string, bool)`

GetPriDnsOk returns a tuple with the PriDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriDns

`func (o *OswDhcpServerOpenApiVO) SetPriDns(v string)`

SetPriDns sets PriDns field to given value.

### HasPriDns

`func (o *OswDhcpServerOpenApiVO) HasPriDns() bool`

HasPriDns returns a boolean if a field has been set.

### GetSndDns

`func (o *OswDhcpServerOpenApiVO) GetSndDns() string`

GetSndDns returns the SndDns field if non-nil, zero value otherwise.

### GetSndDnsOk

`func (o *OswDhcpServerOpenApiVO) GetSndDnsOk() (*string, bool)`

GetSndDnsOk returns a tuple with the SndDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSndDns

`func (o *OswDhcpServerOpenApiVO) SetSndDns(v string)`

SetSndDns sets SndDns field to given value.

### HasSndDns

`func (o *OswDhcpServerOpenApiVO) HasSndDns() bool`

HasSndDns returns a boolean if a field has been set.

### GetVrfId

`func (o *OswDhcpServerOpenApiVO) GetVrfId() string`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *OswDhcpServerOpenApiVO) GetVrfIdOk() (*string, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *OswDhcpServerOpenApiVO) SetVrfId(v string)`

SetVrfId sets VrfId field to given value.

### HasVrfId

`func (o *OswDhcpServerOpenApiVO) HasVrfId() bool`

HasVrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


