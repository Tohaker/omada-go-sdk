# OswDhcpServerVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | Pointer to **string** |  | [optional] 
**Ip** | **string** |  | 
**Leasetime** | **int32** |  | 
**Netmask** | **string** |  | 
**Option138** | Pointer to **string** |  | [optional] 
**Options** | Pointer to [**[]DhcpOptionVO**](DhcpOptionVO.md) |  | [optional] 
**PriDns** | **string** |  | 
**Range** | Pointer to [**[]OswDhcpServerRangeVO**](OswDhcpServerRangeVO.md) |  | [optional] 
**SndDns** | Pointer to **string** |  | [optional] 
**VrfId** | Pointer to **string** |  | [optional] 

## Methods

### NewOswDhcpServerVO

`func NewOswDhcpServerVO(ip string, leasetime int32, netmask string, priDns string, ) *OswDhcpServerVO`

NewOswDhcpServerVO instantiates a new OswDhcpServerVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswDhcpServerVOWithDefaults

`func NewOswDhcpServerVOWithDefaults() *OswDhcpServerVO`

NewOswDhcpServerVOWithDefaults instantiates a new OswDhcpServerVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *OswDhcpServerVO) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *OswDhcpServerVO) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *OswDhcpServerVO) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *OswDhcpServerVO) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIp

`func (o *OswDhcpServerVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OswDhcpServerVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OswDhcpServerVO) SetIp(v string)`

SetIp sets Ip field to given value.


### GetLeasetime

`func (o *OswDhcpServerVO) GetLeasetime() int32`

GetLeasetime returns the Leasetime field if non-nil, zero value otherwise.

### GetLeasetimeOk

`func (o *OswDhcpServerVO) GetLeasetimeOk() (*int32, bool)`

GetLeasetimeOk returns a tuple with the Leasetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasetime

`func (o *OswDhcpServerVO) SetLeasetime(v int32)`

SetLeasetime sets Leasetime field to given value.


### GetNetmask

`func (o *OswDhcpServerVO) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *OswDhcpServerVO) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *OswDhcpServerVO) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.


### GetOption138

`func (o *OswDhcpServerVO) GetOption138() string`

GetOption138 returns the Option138 field if non-nil, zero value otherwise.

### GetOption138Ok

`func (o *OswDhcpServerVO) GetOption138Ok() (*string, bool)`

GetOption138Ok returns a tuple with the Option138 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption138

`func (o *OswDhcpServerVO) SetOption138(v string)`

SetOption138 sets Option138 field to given value.

### HasOption138

`func (o *OswDhcpServerVO) HasOption138() bool`

HasOption138 returns a boolean if a field has been set.

### GetOptions

`func (o *OswDhcpServerVO) GetOptions() []DhcpOptionVO`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *OswDhcpServerVO) GetOptionsOk() (*[]DhcpOptionVO, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *OswDhcpServerVO) SetOptions(v []DhcpOptionVO)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *OswDhcpServerVO) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPriDns

`func (o *OswDhcpServerVO) GetPriDns() string`

GetPriDns returns the PriDns field if non-nil, zero value otherwise.

### GetPriDnsOk

`func (o *OswDhcpServerVO) GetPriDnsOk() (*string, bool)`

GetPriDnsOk returns a tuple with the PriDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriDns

`func (o *OswDhcpServerVO) SetPriDns(v string)`

SetPriDns sets PriDns field to given value.


### GetRange

`func (o *OswDhcpServerVO) GetRange() []OswDhcpServerRangeVO`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *OswDhcpServerVO) GetRangeOk() (*[]OswDhcpServerRangeVO, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *OswDhcpServerVO) SetRange(v []OswDhcpServerRangeVO)`

SetRange sets Range field to given value.

### HasRange

`func (o *OswDhcpServerVO) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetSndDns

`func (o *OswDhcpServerVO) GetSndDns() string`

GetSndDns returns the SndDns field if non-nil, zero value otherwise.

### GetSndDnsOk

`func (o *OswDhcpServerVO) GetSndDnsOk() (*string, bool)`

GetSndDnsOk returns a tuple with the SndDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSndDns

`func (o *OswDhcpServerVO) SetSndDns(v string)`

SetSndDns sets SndDns field to given value.

### HasSndDns

`func (o *OswDhcpServerVO) HasSndDns() bool`

HasSndDns returns a boolean if a field has been set.

### GetVrfId

`func (o *OswDhcpServerVO) GetVrfId() string`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *OswDhcpServerVO) GetVrfIdOk() (*string, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *OswDhcpServerVO) SetVrfId(v string)`

SetVrfId sets VrfId field to given value.

### HasVrfId

`func (o *OswDhcpServerVO) HasVrfId() bool`

HasVrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


