# VpnDefaultValueRespVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int32** | Default service port. | [optional] 
**TunnelIp** | Pointer to **string** | Default tunnel IP for Wire Guard. | [optional] 
**VpnName** | Pointer to **map[string]string** | Default VPN name for each VPN type. | [optional] 

## Methods

### NewVpnDefaultValueRespVO

`func NewVpnDefaultValueRespVO() *VpnDefaultValueRespVO`

NewVpnDefaultValueRespVO instantiates a new VpnDefaultValueRespVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnDefaultValueRespVOWithDefaults

`func NewVpnDefaultValueRespVOWithDefaults() *VpnDefaultValueRespVO`

NewVpnDefaultValueRespVOWithDefaults instantiates a new VpnDefaultValueRespVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *VpnDefaultValueRespVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VpnDefaultValueRespVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VpnDefaultValueRespVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *VpnDefaultValueRespVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetTunnelIp

`func (o *VpnDefaultValueRespVO) GetTunnelIp() string`

GetTunnelIp returns the TunnelIp field if non-nil, zero value otherwise.

### GetTunnelIpOk

`func (o *VpnDefaultValueRespVO) GetTunnelIpOk() (*string, bool)`

GetTunnelIpOk returns a tuple with the TunnelIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelIp

`func (o *VpnDefaultValueRespVO) SetTunnelIp(v string)`

SetTunnelIp sets TunnelIp field to given value.

### HasTunnelIp

`func (o *VpnDefaultValueRespVO) HasTunnelIp() bool`

HasTunnelIp returns a boolean if a field has been set.

### GetVpnName

`func (o *VpnDefaultValueRespVO) GetVpnName() map[string]string`

GetVpnName returns the VpnName field if non-nil, zero value otherwise.

### GetVpnNameOk

`func (o *VpnDefaultValueRespVO) GetVpnNameOk() (*map[string]string, bool)`

GetVpnNameOk returns a tuple with the VpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnName

`func (o *VpnDefaultValueRespVO) SetVpnName(v map[string]string)`

SetVpnName sets VpnName field to given value.

### HasVpnName

`func (o *VpnDefaultValueRespVO) HasVpnName() bool`

HasVpnName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


