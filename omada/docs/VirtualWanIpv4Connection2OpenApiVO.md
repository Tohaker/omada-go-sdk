# VirtualWanIpv4Connection2OpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dns1** | Pointer to **string** | Primary DNS server. | [optional] 
**Dns2** | Pointer to **string** | Secondary DNS server. | [optional] 
**Gateway** | Pointer to **string** | Gateway IP. | [optional] 
**Ipaddr** | Pointer to **string** | IP address. | [optional] 
**MainProto** | Pointer to **string** | The main virtual WAN IPv4 proto type, use static, dhcp, pppoe. | [optional] 
**Netmask** | Pointer to **string** | Subnet mask. | [optional] 
**Proto** | Pointer to **string** | The second virtual WAN IPv4 proto type, use static, dhcp, pppoe. | [optional] 
**Server** | Pointer to **string** | VPN server IP/domain. | [optional] 

## Methods

### NewVirtualWanIpv4Connection2OpenApiVO

`func NewVirtualWanIpv4Connection2OpenApiVO() *VirtualWanIpv4Connection2OpenApiVO`

NewVirtualWanIpv4Connection2OpenApiVO instantiates a new VirtualWanIpv4Connection2OpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualWanIpv4Connection2OpenApiVOWithDefaults

`func NewVirtualWanIpv4Connection2OpenApiVOWithDefaults() *VirtualWanIpv4Connection2OpenApiVO`

NewVirtualWanIpv4Connection2OpenApiVOWithDefaults instantiates a new VirtualWanIpv4Connection2OpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDns1

`func (o *VirtualWanIpv4Connection2OpenApiVO) GetDns1() string`

GetDns1 returns the Dns1 field if non-nil, zero value otherwise.

### GetDns1Ok

`func (o *VirtualWanIpv4Connection2OpenApiVO) GetDns1Ok() (*string, bool)`

GetDns1Ok returns a tuple with the Dns1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns1

`func (o *VirtualWanIpv4Connection2OpenApiVO) SetDns1(v string)`

SetDns1 sets Dns1 field to given value.

### HasDns1

`func (o *VirtualWanIpv4Connection2OpenApiVO) HasDns1() bool`

HasDns1 returns a boolean if a field has been set.

### GetDns2

`func (o *VirtualWanIpv4Connection2OpenApiVO) GetDns2() string`

GetDns2 returns the Dns2 field if non-nil, zero value otherwise.

### GetDns2Ok

`func (o *VirtualWanIpv4Connection2OpenApiVO) GetDns2Ok() (*string, bool)`

GetDns2Ok returns a tuple with the Dns2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns2

`func (o *VirtualWanIpv4Connection2OpenApiVO) SetDns2(v string)`

SetDns2 sets Dns2 field to given value.

### HasDns2

`func (o *VirtualWanIpv4Connection2OpenApiVO) HasDns2() bool`

HasDns2 returns a boolean if a field has been set.

### GetGateway

`func (o *VirtualWanIpv4Connection2OpenApiVO) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *VirtualWanIpv4Connection2OpenApiVO) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *VirtualWanIpv4Connection2OpenApiVO) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *VirtualWanIpv4Connection2OpenApiVO) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIpaddr

`func (o *VirtualWanIpv4Connection2OpenApiVO) GetIpaddr() string`

GetIpaddr returns the Ipaddr field if non-nil, zero value otherwise.

### GetIpaddrOk

`func (o *VirtualWanIpv4Connection2OpenApiVO) GetIpaddrOk() (*string, bool)`

GetIpaddrOk returns a tuple with the Ipaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddr

`func (o *VirtualWanIpv4Connection2OpenApiVO) SetIpaddr(v string)`

SetIpaddr sets Ipaddr field to given value.

### HasIpaddr

`func (o *VirtualWanIpv4Connection2OpenApiVO) HasIpaddr() bool`

HasIpaddr returns a boolean if a field has been set.

### GetMainProto

`func (o *VirtualWanIpv4Connection2OpenApiVO) GetMainProto() string`

GetMainProto returns the MainProto field if non-nil, zero value otherwise.

### GetMainProtoOk

`func (o *VirtualWanIpv4Connection2OpenApiVO) GetMainProtoOk() (*string, bool)`

GetMainProtoOk returns a tuple with the MainProto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainProto

`func (o *VirtualWanIpv4Connection2OpenApiVO) SetMainProto(v string)`

SetMainProto sets MainProto field to given value.

### HasMainProto

`func (o *VirtualWanIpv4Connection2OpenApiVO) HasMainProto() bool`

HasMainProto returns a boolean if a field has been set.

### GetNetmask

`func (o *VirtualWanIpv4Connection2OpenApiVO) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *VirtualWanIpv4Connection2OpenApiVO) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *VirtualWanIpv4Connection2OpenApiVO) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *VirtualWanIpv4Connection2OpenApiVO) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetProto

`func (o *VirtualWanIpv4Connection2OpenApiVO) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *VirtualWanIpv4Connection2OpenApiVO) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *VirtualWanIpv4Connection2OpenApiVO) SetProto(v string)`

SetProto sets Proto field to given value.

### HasProto

`func (o *VirtualWanIpv4Connection2OpenApiVO) HasProto() bool`

HasProto returns a boolean if a field has been set.

### GetServer

`func (o *VirtualWanIpv4Connection2OpenApiVO) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *VirtualWanIpv4Connection2OpenApiVO) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *VirtualWanIpv4Connection2OpenApiVO) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *VirtualWanIpv4Connection2OpenApiVO) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


