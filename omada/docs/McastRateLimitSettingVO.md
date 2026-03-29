# McastRateLimitSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArpEnable** | Pointer to **bool** | Whether to enable ARP. | [optional] 
**ArpPps** | Pointer to **int32** | Parameter [arpPps] indicates the number of ARP packets allowed to pass per second. | [optional] 
**DhcpEnable** | Pointer to **bool** | Whether to enable DHCP. | [optional] 
**DhcpPps** | Pointer to **int32** | Parameter [dhcpPps] indicates the number of DHCP packets allowed to pass per second. | [optional] 
**Dhcpv6Enable** | Pointer to **bool** | Whether to enable DHCP v6. | [optional] 
**Dhcpv6Pps** | Pointer to **int32** | Parameter [dhcpv6Pps] indicates the number of DHCP v6 packets allowed to pass per second. | [optional] 
**Enable** | Pointer to **bool** | Parameter [enable] indicates whether to enable the multicast rate limiting function. True: on, false: off | [optional] 
**IgmpEnable** | Pointer to **bool** | Whether to enable IGMP. | [optional] 
**IgmpPps** | Pointer to **int32** | Parameter [igmpPps] indicates the number of IGMP packets allowed to pass per second. | [optional] 
**MdnsEnable** | Pointer to **bool** | Whether to enable MDNS. | [optional] 
**MdnsPps** | Pointer to **int32** | Parameter [mdnsPps] indicates the number of MDNS packets allowed to pass per second. | [optional] 
**NdEnable** | Pointer to **bool** | Whether to enable ND. | [optional] 
**NdPps** | Pointer to **int32** | Parameter [ndPps] indicates the number of ND packets allowed to pass per second. | [optional] 
**OtherBcastEnable** | Pointer to **bool** | Whether to enable OTHER BROADCAST. | [optional] 
**OtherBcastPps** | Pointer to **int32** | Parameter [otherBcastPps] indicates the number of OTHER BROADCAST packets allowed to pass per second. | [optional] 
**OtherMcastEnable** | Pointer to **bool** | Whether to enable OTHER MULTICAST. | [optional] 
**OtherMcastPps** | Pointer to **int32** | Parameter [otherMcastPps] indicates the number of OTHER MULTICAST packets allowed to pass per second. | [optional] 

## Methods

### NewMcastRateLimitSettingVO

`func NewMcastRateLimitSettingVO() *McastRateLimitSettingVO`

NewMcastRateLimitSettingVO instantiates a new McastRateLimitSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcastRateLimitSettingVOWithDefaults

`func NewMcastRateLimitSettingVOWithDefaults() *McastRateLimitSettingVO`

NewMcastRateLimitSettingVOWithDefaults instantiates a new McastRateLimitSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArpEnable

`func (o *McastRateLimitSettingVO) GetArpEnable() bool`

GetArpEnable returns the ArpEnable field if non-nil, zero value otherwise.

### GetArpEnableOk

`func (o *McastRateLimitSettingVO) GetArpEnableOk() (*bool, bool)`

GetArpEnableOk returns a tuple with the ArpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpEnable

`func (o *McastRateLimitSettingVO) SetArpEnable(v bool)`

SetArpEnable sets ArpEnable field to given value.

### HasArpEnable

`func (o *McastRateLimitSettingVO) HasArpEnable() bool`

HasArpEnable returns a boolean if a field has been set.

### GetArpPps

`func (o *McastRateLimitSettingVO) GetArpPps() int32`

GetArpPps returns the ArpPps field if non-nil, zero value otherwise.

### GetArpPpsOk

`func (o *McastRateLimitSettingVO) GetArpPpsOk() (*int32, bool)`

GetArpPpsOk returns a tuple with the ArpPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpPps

`func (o *McastRateLimitSettingVO) SetArpPps(v int32)`

SetArpPps sets ArpPps field to given value.

### HasArpPps

`func (o *McastRateLimitSettingVO) HasArpPps() bool`

HasArpPps returns a boolean if a field has been set.

### GetDhcpEnable

`func (o *McastRateLimitSettingVO) GetDhcpEnable() bool`

GetDhcpEnable returns the DhcpEnable field if non-nil, zero value otherwise.

### GetDhcpEnableOk

`func (o *McastRateLimitSettingVO) GetDhcpEnableOk() (*bool, bool)`

GetDhcpEnableOk returns a tuple with the DhcpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpEnable

`func (o *McastRateLimitSettingVO) SetDhcpEnable(v bool)`

SetDhcpEnable sets DhcpEnable field to given value.

### HasDhcpEnable

`func (o *McastRateLimitSettingVO) HasDhcpEnable() bool`

HasDhcpEnable returns a boolean if a field has been set.

### GetDhcpPps

`func (o *McastRateLimitSettingVO) GetDhcpPps() int32`

GetDhcpPps returns the DhcpPps field if non-nil, zero value otherwise.

### GetDhcpPpsOk

`func (o *McastRateLimitSettingVO) GetDhcpPpsOk() (*int32, bool)`

GetDhcpPpsOk returns a tuple with the DhcpPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpPps

`func (o *McastRateLimitSettingVO) SetDhcpPps(v int32)`

SetDhcpPps sets DhcpPps field to given value.

### HasDhcpPps

`func (o *McastRateLimitSettingVO) HasDhcpPps() bool`

HasDhcpPps returns a boolean if a field has been set.

### GetDhcpv6Enable

`func (o *McastRateLimitSettingVO) GetDhcpv6Enable() bool`

GetDhcpv6Enable returns the Dhcpv6Enable field if non-nil, zero value otherwise.

### GetDhcpv6EnableOk

`func (o *McastRateLimitSettingVO) GetDhcpv6EnableOk() (*bool, bool)`

GetDhcpv6EnableOk returns a tuple with the Dhcpv6Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6Enable

`func (o *McastRateLimitSettingVO) SetDhcpv6Enable(v bool)`

SetDhcpv6Enable sets Dhcpv6Enable field to given value.

### HasDhcpv6Enable

`func (o *McastRateLimitSettingVO) HasDhcpv6Enable() bool`

HasDhcpv6Enable returns a boolean if a field has been set.

### GetDhcpv6Pps

`func (o *McastRateLimitSettingVO) GetDhcpv6Pps() int32`

GetDhcpv6Pps returns the Dhcpv6Pps field if non-nil, zero value otherwise.

### GetDhcpv6PpsOk

`func (o *McastRateLimitSettingVO) GetDhcpv6PpsOk() (*int32, bool)`

GetDhcpv6PpsOk returns a tuple with the Dhcpv6Pps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6Pps

`func (o *McastRateLimitSettingVO) SetDhcpv6Pps(v int32)`

SetDhcpv6Pps sets Dhcpv6Pps field to given value.

### HasDhcpv6Pps

`func (o *McastRateLimitSettingVO) HasDhcpv6Pps() bool`

HasDhcpv6Pps returns a boolean if a field has been set.

### GetEnable

`func (o *McastRateLimitSettingVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *McastRateLimitSettingVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *McastRateLimitSettingVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *McastRateLimitSettingVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetIgmpEnable

`func (o *McastRateLimitSettingVO) GetIgmpEnable() bool`

GetIgmpEnable returns the IgmpEnable field if non-nil, zero value otherwise.

### GetIgmpEnableOk

`func (o *McastRateLimitSettingVO) GetIgmpEnableOk() (*bool, bool)`

GetIgmpEnableOk returns a tuple with the IgmpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpEnable

`func (o *McastRateLimitSettingVO) SetIgmpEnable(v bool)`

SetIgmpEnable sets IgmpEnable field to given value.

### HasIgmpEnable

`func (o *McastRateLimitSettingVO) HasIgmpEnable() bool`

HasIgmpEnable returns a boolean if a field has been set.

### GetIgmpPps

`func (o *McastRateLimitSettingVO) GetIgmpPps() int32`

GetIgmpPps returns the IgmpPps field if non-nil, zero value otherwise.

### GetIgmpPpsOk

`func (o *McastRateLimitSettingVO) GetIgmpPpsOk() (*int32, bool)`

GetIgmpPpsOk returns a tuple with the IgmpPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpPps

`func (o *McastRateLimitSettingVO) SetIgmpPps(v int32)`

SetIgmpPps sets IgmpPps field to given value.

### HasIgmpPps

`func (o *McastRateLimitSettingVO) HasIgmpPps() bool`

HasIgmpPps returns a boolean if a field has been set.

### GetMdnsEnable

`func (o *McastRateLimitSettingVO) GetMdnsEnable() bool`

GetMdnsEnable returns the MdnsEnable field if non-nil, zero value otherwise.

### GetMdnsEnableOk

`func (o *McastRateLimitSettingVO) GetMdnsEnableOk() (*bool, bool)`

GetMdnsEnableOk returns a tuple with the MdnsEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdnsEnable

`func (o *McastRateLimitSettingVO) SetMdnsEnable(v bool)`

SetMdnsEnable sets MdnsEnable field to given value.

### HasMdnsEnable

`func (o *McastRateLimitSettingVO) HasMdnsEnable() bool`

HasMdnsEnable returns a boolean if a field has been set.

### GetMdnsPps

`func (o *McastRateLimitSettingVO) GetMdnsPps() int32`

GetMdnsPps returns the MdnsPps field if non-nil, zero value otherwise.

### GetMdnsPpsOk

`func (o *McastRateLimitSettingVO) GetMdnsPpsOk() (*int32, bool)`

GetMdnsPpsOk returns a tuple with the MdnsPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdnsPps

`func (o *McastRateLimitSettingVO) SetMdnsPps(v int32)`

SetMdnsPps sets MdnsPps field to given value.

### HasMdnsPps

`func (o *McastRateLimitSettingVO) HasMdnsPps() bool`

HasMdnsPps returns a boolean if a field has been set.

### GetNdEnable

`func (o *McastRateLimitSettingVO) GetNdEnable() bool`

GetNdEnable returns the NdEnable field if non-nil, zero value otherwise.

### GetNdEnableOk

`func (o *McastRateLimitSettingVO) GetNdEnableOk() (*bool, bool)`

GetNdEnableOk returns a tuple with the NdEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdEnable

`func (o *McastRateLimitSettingVO) SetNdEnable(v bool)`

SetNdEnable sets NdEnable field to given value.

### HasNdEnable

`func (o *McastRateLimitSettingVO) HasNdEnable() bool`

HasNdEnable returns a boolean if a field has been set.

### GetNdPps

`func (o *McastRateLimitSettingVO) GetNdPps() int32`

GetNdPps returns the NdPps field if non-nil, zero value otherwise.

### GetNdPpsOk

`func (o *McastRateLimitSettingVO) GetNdPpsOk() (*int32, bool)`

GetNdPpsOk returns a tuple with the NdPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdPps

`func (o *McastRateLimitSettingVO) SetNdPps(v int32)`

SetNdPps sets NdPps field to given value.

### HasNdPps

`func (o *McastRateLimitSettingVO) HasNdPps() bool`

HasNdPps returns a boolean if a field has been set.

### GetOtherBcastEnable

`func (o *McastRateLimitSettingVO) GetOtherBcastEnable() bool`

GetOtherBcastEnable returns the OtherBcastEnable field if non-nil, zero value otherwise.

### GetOtherBcastEnableOk

`func (o *McastRateLimitSettingVO) GetOtherBcastEnableOk() (*bool, bool)`

GetOtherBcastEnableOk returns a tuple with the OtherBcastEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherBcastEnable

`func (o *McastRateLimitSettingVO) SetOtherBcastEnable(v bool)`

SetOtherBcastEnable sets OtherBcastEnable field to given value.

### HasOtherBcastEnable

`func (o *McastRateLimitSettingVO) HasOtherBcastEnable() bool`

HasOtherBcastEnable returns a boolean if a field has been set.

### GetOtherBcastPps

`func (o *McastRateLimitSettingVO) GetOtherBcastPps() int32`

GetOtherBcastPps returns the OtherBcastPps field if non-nil, zero value otherwise.

### GetOtherBcastPpsOk

`func (o *McastRateLimitSettingVO) GetOtherBcastPpsOk() (*int32, bool)`

GetOtherBcastPpsOk returns a tuple with the OtherBcastPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherBcastPps

`func (o *McastRateLimitSettingVO) SetOtherBcastPps(v int32)`

SetOtherBcastPps sets OtherBcastPps field to given value.

### HasOtherBcastPps

`func (o *McastRateLimitSettingVO) HasOtherBcastPps() bool`

HasOtherBcastPps returns a boolean if a field has been set.

### GetOtherMcastEnable

`func (o *McastRateLimitSettingVO) GetOtherMcastEnable() bool`

GetOtherMcastEnable returns the OtherMcastEnable field if non-nil, zero value otherwise.

### GetOtherMcastEnableOk

`func (o *McastRateLimitSettingVO) GetOtherMcastEnableOk() (*bool, bool)`

GetOtherMcastEnableOk returns a tuple with the OtherMcastEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherMcastEnable

`func (o *McastRateLimitSettingVO) SetOtherMcastEnable(v bool)`

SetOtherMcastEnable sets OtherMcastEnable field to given value.

### HasOtherMcastEnable

`func (o *McastRateLimitSettingVO) HasOtherMcastEnable() bool`

HasOtherMcastEnable returns a boolean if a field has been set.

### GetOtherMcastPps

`func (o *McastRateLimitSettingVO) GetOtherMcastPps() int32`

GetOtherMcastPps returns the OtherMcastPps field if non-nil, zero value otherwise.

### GetOtherMcastPpsOk

`func (o *McastRateLimitSettingVO) GetOtherMcastPpsOk() (*int32, bool)`

GetOtherMcastPpsOk returns a tuple with the OtherMcastPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherMcastPps

`func (o *McastRateLimitSettingVO) SetOtherMcastPps(v int32)`

SetOtherMcastPps sets OtherMcastPps field to given value.

### HasOtherMcastPps

`func (o *McastRateLimitSettingVO) HasOtherMcastPps() bool`

HasOtherMcastPps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


