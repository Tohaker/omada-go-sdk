# AvailableWanPortOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configured** | Pointer to **bool** | qos configured | [optional] 
**Ip** | Pointer to **string** | IPv4 address. | [optional] 
**Ipv6** | Pointer to **bool** | IPv6 enable. | [optional] 
**Ipv6Proto** | Pointer to **string** | IPv6 connection type: 0: static; 1: dynamic; 2: PPPoE; 3: 6to4Tunnel; 4: bridge. | [optional] 
**PortName** | Pointer to **string** | port name | [optional] 
**PortUuid** | Pointer to **string** | port id | [optional] 
**Proto** | Pointer to **string** | WAN IPv4 connection type, it supports Static IP, DHCP, PPPoE, L2TP, PPTP, DS-Lite, and MAP-E. | [optional] 
**RecommendedWan** | Pointer to **bool** | Recommended WAN port. | [optional] 
**SupportIptv** | Pointer to **bool** | support iptv? | [optional] 
**Type** | Pointer to **int32** | wan type | [optional] 
**VirtualWanNum** | Pointer to **int32** | How many virtual wan does this wan port support | [optional] 

## Methods

### NewAvailableWanPortOpenApiVO

`func NewAvailableWanPortOpenApiVO() *AvailableWanPortOpenApiVO`

NewAvailableWanPortOpenApiVO instantiates a new AvailableWanPortOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableWanPortOpenApiVOWithDefaults

`func NewAvailableWanPortOpenApiVOWithDefaults() *AvailableWanPortOpenApiVO`

NewAvailableWanPortOpenApiVOWithDefaults instantiates a new AvailableWanPortOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigured

`func (o *AvailableWanPortOpenApiVO) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *AvailableWanPortOpenApiVO) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *AvailableWanPortOpenApiVO) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *AvailableWanPortOpenApiVO) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetIp

`func (o *AvailableWanPortOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AvailableWanPortOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AvailableWanPortOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AvailableWanPortOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6

`func (o *AvailableWanPortOpenApiVO) GetIpv6() bool`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *AvailableWanPortOpenApiVO) GetIpv6Ok() (*bool, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *AvailableWanPortOpenApiVO) SetIpv6(v bool)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *AvailableWanPortOpenApiVO) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetIpv6Proto

`func (o *AvailableWanPortOpenApiVO) GetIpv6Proto() string`

GetIpv6Proto returns the Ipv6Proto field if non-nil, zero value otherwise.

### GetIpv6ProtoOk

`func (o *AvailableWanPortOpenApiVO) GetIpv6ProtoOk() (*string, bool)`

GetIpv6ProtoOk returns a tuple with the Ipv6Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Proto

`func (o *AvailableWanPortOpenApiVO) SetIpv6Proto(v string)`

SetIpv6Proto sets Ipv6Proto field to given value.

### HasIpv6Proto

`func (o *AvailableWanPortOpenApiVO) HasIpv6Proto() bool`

HasIpv6Proto returns a boolean if a field has been set.

### GetPortName

`func (o *AvailableWanPortOpenApiVO) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *AvailableWanPortOpenApiVO) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *AvailableWanPortOpenApiVO) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *AvailableWanPortOpenApiVO) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetPortUuid

`func (o *AvailableWanPortOpenApiVO) GetPortUuid() string`

GetPortUuid returns the PortUuid field if non-nil, zero value otherwise.

### GetPortUuidOk

`func (o *AvailableWanPortOpenApiVO) GetPortUuidOk() (*string, bool)`

GetPortUuidOk returns a tuple with the PortUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUuid

`func (o *AvailableWanPortOpenApiVO) SetPortUuid(v string)`

SetPortUuid sets PortUuid field to given value.

### HasPortUuid

`func (o *AvailableWanPortOpenApiVO) HasPortUuid() bool`

HasPortUuid returns a boolean if a field has been set.

### GetProto

`func (o *AvailableWanPortOpenApiVO) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *AvailableWanPortOpenApiVO) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *AvailableWanPortOpenApiVO) SetProto(v string)`

SetProto sets Proto field to given value.

### HasProto

`func (o *AvailableWanPortOpenApiVO) HasProto() bool`

HasProto returns a boolean if a field has been set.

### GetRecommendedWan

`func (o *AvailableWanPortOpenApiVO) GetRecommendedWan() bool`

GetRecommendedWan returns the RecommendedWan field if non-nil, zero value otherwise.

### GetRecommendedWanOk

`func (o *AvailableWanPortOpenApiVO) GetRecommendedWanOk() (*bool, bool)`

GetRecommendedWanOk returns a tuple with the RecommendedWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedWan

`func (o *AvailableWanPortOpenApiVO) SetRecommendedWan(v bool)`

SetRecommendedWan sets RecommendedWan field to given value.

### HasRecommendedWan

`func (o *AvailableWanPortOpenApiVO) HasRecommendedWan() bool`

HasRecommendedWan returns a boolean if a field has been set.

### GetSupportIptv

`func (o *AvailableWanPortOpenApiVO) GetSupportIptv() bool`

GetSupportIptv returns the SupportIptv field if non-nil, zero value otherwise.

### GetSupportIptvOk

`func (o *AvailableWanPortOpenApiVO) GetSupportIptvOk() (*bool, bool)`

GetSupportIptvOk returns a tuple with the SupportIptv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIptv

`func (o *AvailableWanPortOpenApiVO) SetSupportIptv(v bool)`

SetSupportIptv sets SupportIptv field to given value.

### HasSupportIptv

`func (o *AvailableWanPortOpenApiVO) HasSupportIptv() bool`

HasSupportIptv returns a boolean if a field has been set.

### GetType

`func (o *AvailableWanPortOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AvailableWanPortOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AvailableWanPortOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *AvailableWanPortOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVirtualWanNum

`func (o *AvailableWanPortOpenApiVO) GetVirtualWanNum() int32`

GetVirtualWanNum returns the VirtualWanNum field if non-nil, zero value otherwise.

### GetVirtualWanNumOk

`func (o *AvailableWanPortOpenApiVO) GetVirtualWanNumOk() (*int32, bool)`

GetVirtualWanNumOk returns a tuple with the VirtualWanNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanNum

`func (o *AvailableWanPortOpenApiVO) SetVirtualWanNum(v int32)`

SetVirtualWanNum sets VirtualWanNum field to given value.

### HasVirtualWanNum

`func (o *AvailableWanPortOpenApiVO) HasVirtualWanNum() bool`

HasVirtualWanNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


