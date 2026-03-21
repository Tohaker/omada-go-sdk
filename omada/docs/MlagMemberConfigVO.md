# MlagMemberConfigVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DadEnable** | Pointer to **bool** | Whether the DAD enable. | [optional] 
**DadLinkPorts** | Pointer to **[]string** | DAD Link Ports | [optional] 
**DadLocalIp** | Pointer to **string** | DAD Local IP | [optional] 
**DadLocalIpv6** | Pointer to **string** | DAD Local IPv6 | [optional] 
**DadPeerIp** | Pointer to **string** | DAD Peer IP | [optional] 
**DadPeerIpv6** | Pointer to **string** | DAD Peer IPv6 | [optional] 
**Mac** | **string** | Device Mac. | 
**PeerLinkPorts** | Pointer to **[]string** | Peer Link Ports | [optional] 
**Priority** | **int32** | Priority of the device in the M-LAG group. | 

## Methods

### NewMlagMemberConfigVO

`func NewMlagMemberConfigVO(mac string, priority int32, ) *MlagMemberConfigVO`

NewMlagMemberConfigVO instantiates a new MlagMemberConfigVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlagMemberConfigVOWithDefaults

`func NewMlagMemberConfigVOWithDefaults() *MlagMemberConfigVO`

NewMlagMemberConfigVOWithDefaults instantiates a new MlagMemberConfigVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDadEnable

`func (o *MlagMemberConfigVO) GetDadEnable() bool`

GetDadEnable returns the DadEnable field if non-nil, zero value otherwise.

### GetDadEnableOk

`func (o *MlagMemberConfigVO) GetDadEnableOk() (*bool, bool)`

GetDadEnableOk returns a tuple with the DadEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDadEnable

`func (o *MlagMemberConfigVO) SetDadEnable(v bool)`

SetDadEnable sets DadEnable field to given value.

### HasDadEnable

`func (o *MlagMemberConfigVO) HasDadEnable() bool`

HasDadEnable returns a boolean if a field has been set.

### GetDadLinkPorts

`func (o *MlagMemberConfigVO) GetDadLinkPorts() []string`

GetDadLinkPorts returns the DadLinkPorts field if non-nil, zero value otherwise.

### GetDadLinkPortsOk

`func (o *MlagMemberConfigVO) GetDadLinkPortsOk() (*[]string, bool)`

GetDadLinkPortsOk returns a tuple with the DadLinkPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDadLinkPorts

`func (o *MlagMemberConfigVO) SetDadLinkPorts(v []string)`

SetDadLinkPorts sets DadLinkPorts field to given value.

### HasDadLinkPorts

`func (o *MlagMemberConfigVO) HasDadLinkPorts() bool`

HasDadLinkPorts returns a boolean if a field has been set.

### GetDadLocalIp

`func (o *MlagMemberConfigVO) GetDadLocalIp() string`

GetDadLocalIp returns the DadLocalIp field if non-nil, zero value otherwise.

### GetDadLocalIpOk

`func (o *MlagMemberConfigVO) GetDadLocalIpOk() (*string, bool)`

GetDadLocalIpOk returns a tuple with the DadLocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDadLocalIp

`func (o *MlagMemberConfigVO) SetDadLocalIp(v string)`

SetDadLocalIp sets DadLocalIp field to given value.

### HasDadLocalIp

`func (o *MlagMemberConfigVO) HasDadLocalIp() bool`

HasDadLocalIp returns a boolean if a field has been set.

### GetDadLocalIpv6

`func (o *MlagMemberConfigVO) GetDadLocalIpv6() string`

GetDadLocalIpv6 returns the DadLocalIpv6 field if non-nil, zero value otherwise.

### GetDadLocalIpv6Ok

`func (o *MlagMemberConfigVO) GetDadLocalIpv6Ok() (*string, bool)`

GetDadLocalIpv6Ok returns a tuple with the DadLocalIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDadLocalIpv6

`func (o *MlagMemberConfigVO) SetDadLocalIpv6(v string)`

SetDadLocalIpv6 sets DadLocalIpv6 field to given value.

### HasDadLocalIpv6

`func (o *MlagMemberConfigVO) HasDadLocalIpv6() bool`

HasDadLocalIpv6 returns a boolean if a field has been set.

### GetDadPeerIp

`func (o *MlagMemberConfigVO) GetDadPeerIp() string`

GetDadPeerIp returns the DadPeerIp field if non-nil, zero value otherwise.

### GetDadPeerIpOk

`func (o *MlagMemberConfigVO) GetDadPeerIpOk() (*string, bool)`

GetDadPeerIpOk returns a tuple with the DadPeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDadPeerIp

`func (o *MlagMemberConfigVO) SetDadPeerIp(v string)`

SetDadPeerIp sets DadPeerIp field to given value.

### HasDadPeerIp

`func (o *MlagMemberConfigVO) HasDadPeerIp() bool`

HasDadPeerIp returns a boolean if a field has been set.

### GetDadPeerIpv6

`func (o *MlagMemberConfigVO) GetDadPeerIpv6() string`

GetDadPeerIpv6 returns the DadPeerIpv6 field if non-nil, zero value otherwise.

### GetDadPeerIpv6Ok

`func (o *MlagMemberConfigVO) GetDadPeerIpv6Ok() (*string, bool)`

GetDadPeerIpv6Ok returns a tuple with the DadPeerIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDadPeerIpv6

`func (o *MlagMemberConfigVO) SetDadPeerIpv6(v string)`

SetDadPeerIpv6 sets DadPeerIpv6 field to given value.

### HasDadPeerIpv6

`func (o *MlagMemberConfigVO) HasDadPeerIpv6() bool`

HasDadPeerIpv6 returns a boolean if a field has been set.

### GetMac

`func (o *MlagMemberConfigVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *MlagMemberConfigVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *MlagMemberConfigVO) SetMac(v string)`

SetMac sets Mac field to given value.


### GetPeerLinkPorts

`func (o *MlagMemberConfigVO) GetPeerLinkPorts() []string`

GetPeerLinkPorts returns the PeerLinkPorts field if non-nil, zero value otherwise.

### GetPeerLinkPortsOk

`func (o *MlagMemberConfigVO) GetPeerLinkPortsOk() (*[]string, bool)`

GetPeerLinkPortsOk returns a tuple with the PeerLinkPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerLinkPorts

`func (o *MlagMemberConfigVO) SetPeerLinkPorts(v []string)`

SetPeerLinkPorts sets PeerLinkPorts field to given value.

### HasPeerLinkPorts

`func (o *MlagMemberConfigVO) HasPeerLinkPorts() bool`

HasPeerLinkPorts returns a boolean if a field has been set.

### GetPriority

`func (o *MlagMemberConfigVO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MlagMemberConfigVO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MlagMemberConfigVO) SetPriority(v int32)`

SetPriority sets Priority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


