# NewMcastRateLimitSettingOpenApiVO

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

### NewNewMcastRateLimitSettingOpenApiVO

`func NewNewMcastRateLimitSettingOpenApiVO() *NewMcastRateLimitSettingOpenApiVO`

NewNewMcastRateLimitSettingOpenApiVO instantiates a new NewMcastRateLimitSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewMcastRateLimitSettingOpenApiVOWithDefaults

`func NewNewMcastRateLimitSettingOpenApiVOWithDefaults() *NewMcastRateLimitSettingOpenApiVO`

NewNewMcastRateLimitSettingOpenApiVOWithDefaults instantiates a new NewMcastRateLimitSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArpEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) GetArpEnable() bool`

GetArpEnable returns the ArpEnable field if non-nil, zero value otherwise.

### GetArpEnableOk

`func (o *NewMcastRateLimitSettingOpenApiVO) GetArpEnableOk() (*bool, bool)`

GetArpEnableOk returns a tuple with the ArpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) SetArpEnable(v bool)`

SetArpEnable sets ArpEnable field to given value.

### HasArpEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) HasArpEnable() bool`

HasArpEnable returns a boolean if a field has been set.

### GetArpPps

`func (o *NewMcastRateLimitSettingOpenApiVO) GetArpPps() int32`

GetArpPps returns the ArpPps field if non-nil, zero value otherwise.

### GetArpPpsOk

`func (o *NewMcastRateLimitSettingOpenApiVO) GetArpPpsOk() (*int32, bool)`

GetArpPpsOk returns a tuple with the ArpPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpPps

`func (o *NewMcastRateLimitSettingOpenApiVO) SetArpPps(v int32)`

SetArpPps sets ArpPps field to given value.

### HasArpPps

`func (o *NewMcastRateLimitSettingOpenApiVO) HasArpPps() bool`

HasArpPps returns a boolean if a field has been set.

### GetDhcpEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) GetDhcpEnable() bool`

GetDhcpEnable returns the DhcpEnable field if non-nil, zero value otherwise.

### GetDhcpEnableOk

`func (o *NewMcastRateLimitSettingOpenApiVO) GetDhcpEnableOk() (*bool, bool)`

GetDhcpEnableOk returns a tuple with the DhcpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) SetDhcpEnable(v bool)`

SetDhcpEnable sets DhcpEnable field to given value.

### HasDhcpEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) HasDhcpEnable() bool`

HasDhcpEnable returns a boolean if a field has been set.

### GetDhcpPps

`func (o *NewMcastRateLimitSettingOpenApiVO) GetDhcpPps() int32`

GetDhcpPps returns the DhcpPps field if non-nil, zero value otherwise.

### GetDhcpPpsOk

`func (o *NewMcastRateLimitSettingOpenApiVO) GetDhcpPpsOk() (*int32, bool)`

GetDhcpPpsOk returns a tuple with the DhcpPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpPps

`func (o *NewMcastRateLimitSettingOpenApiVO) SetDhcpPps(v int32)`

SetDhcpPps sets DhcpPps field to given value.

### HasDhcpPps

`func (o *NewMcastRateLimitSettingOpenApiVO) HasDhcpPps() bool`

HasDhcpPps returns a boolean if a field has been set.

### GetDhcpv6Enable

`func (o *NewMcastRateLimitSettingOpenApiVO) GetDhcpv6Enable() bool`

GetDhcpv6Enable returns the Dhcpv6Enable field if non-nil, zero value otherwise.

### GetDhcpv6EnableOk

`func (o *NewMcastRateLimitSettingOpenApiVO) GetDhcpv6EnableOk() (*bool, bool)`

GetDhcpv6EnableOk returns a tuple with the Dhcpv6Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6Enable

`func (o *NewMcastRateLimitSettingOpenApiVO) SetDhcpv6Enable(v bool)`

SetDhcpv6Enable sets Dhcpv6Enable field to given value.

### HasDhcpv6Enable

`func (o *NewMcastRateLimitSettingOpenApiVO) HasDhcpv6Enable() bool`

HasDhcpv6Enable returns a boolean if a field has been set.

### GetDhcpv6Pps

`func (o *NewMcastRateLimitSettingOpenApiVO) GetDhcpv6Pps() int32`

GetDhcpv6Pps returns the Dhcpv6Pps field if non-nil, zero value otherwise.

### GetDhcpv6PpsOk

`func (o *NewMcastRateLimitSettingOpenApiVO) GetDhcpv6PpsOk() (*int32, bool)`

GetDhcpv6PpsOk returns a tuple with the Dhcpv6Pps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6Pps

`func (o *NewMcastRateLimitSettingOpenApiVO) SetDhcpv6Pps(v int32)`

SetDhcpv6Pps sets Dhcpv6Pps field to given value.

### HasDhcpv6Pps

`func (o *NewMcastRateLimitSettingOpenApiVO) HasDhcpv6Pps() bool`

HasDhcpv6Pps returns a boolean if a field has been set.

### GetEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *NewMcastRateLimitSettingOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetIgmpEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) GetIgmpEnable() bool`

GetIgmpEnable returns the IgmpEnable field if non-nil, zero value otherwise.

### GetIgmpEnableOk

`func (o *NewMcastRateLimitSettingOpenApiVO) GetIgmpEnableOk() (*bool, bool)`

GetIgmpEnableOk returns a tuple with the IgmpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) SetIgmpEnable(v bool)`

SetIgmpEnable sets IgmpEnable field to given value.

### HasIgmpEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) HasIgmpEnable() bool`

HasIgmpEnable returns a boolean if a field has been set.

### GetIgmpPps

`func (o *NewMcastRateLimitSettingOpenApiVO) GetIgmpPps() int32`

GetIgmpPps returns the IgmpPps field if non-nil, zero value otherwise.

### GetIgmpPpsOk

`func (o *NewMcastRateLimitSettingOpenApiVO) GetIgmpPpsOk() (*int32, bool)`

GetIgmpPpsOk returns a tuple with the IgmpPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpPps

`func (o *NewMcastRateLimitSettingOpenApiVO) SetIgmpPps(v int32)`

SetIgmpPps sets IgmpPps field to given value.

### HasIgmpPps

`func (o *NewMcastRateLimitSettingOpenApiVO) HasIgmpPps() bool`

HasIgmpPps returns a boolean if a field has been set.

### GetMdnsEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) GetMdnsEnable() bool`

GetMdnsEnable returns the MdnsEnable field if non-nil, zero value otherwise.

### GetMdnsEnableOk

`func (o *NewMcastRateLimitSettingOpenApiVO) GetMdnsEnableOk() (*bool, bool)`

GetMdnsEnableOk returns a tuple with the MdnsEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdnsEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) SetMdnsEnable(v bool)`

SetMdnsEnable sets MdnsEnable field to given value.

### HasMdnsEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) HasMdnsEnable() bool`

HasMdnsEnable returns a boolean if a field has been set.

### GetMdnsPps

`func (o *NewMcastRateLimitSettingOpenApiVO) GetMdnsPps() int32`

GetMdnsPps returns the MdnsPps field if non-nil, zero value otherwise.

### GetMdnsPpsOk

`func (o *NewMcastRateLimitSettingOpenApiVO) GetMdnsPpsOk() (*int32, bool)`

GetMdnsPpsOk returns a tuple with the MdnsPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdnsPps

`func (o *NewMcastRateLimitSettingOpenApiVO) SetMdnsPps(v int32)`

SetMdnsPps sets MdnsPps field to given value.

### HasMdnsPps

`func (o *NewMcastRateLimitSettingOpenApiVO) HasMdnsPps() bool`

HasMdnsPps returns a boolean if a field has been set.

### GetNdEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) GetNdEnable() bool`

GetNdEnable returns the NdEnable field if non-nil, zero value otherwise.

### GetNdEnableOk

`func (o *NewMcastRateLimitSettingOpenApiVO) GetNdEnableOk() (*bool, bool)`

GetNdEnableOk returns a tuple with the NdEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) SetNdEnable(v bool)`

SetNdEnable sets NdEnable field to given value.

### HasNdEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) HasNdEnable() bool`

HasNdEnable returns a boolean if a field has been set.

### GetNdPps

`func (o *NewMcastRateLimitSettingOpenApiVO) GetNdPps() int32`

GetNdPps returns the NdPps field if non-nil, zero value otherwise.

### GetNdPpsOk

`func (o *NewMcastRateLimitSettingOpenApiVO) GetNdPpsOk() (*int32, bool)`

GetNdPpsOk returns a tuple with the NdPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdPps

`func (o *NewMcastRateLimitSettingOpenApiVO) SetNdPps(v int32)`

SetNdPps sets NdPps field to given value.

### HasNdPps

`func (o *NewMcastRateLimitSettingOpenApiVO) HasNdPps() bool`

HasNdPps returns a boolean if a field has been set.

### GetOtherBcastEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) GetOtherBcastEnable() bool`

GetOtherBcastEnable returns the OtherBcastEnable field if non-nil, zero value otherwise.

### GetOtherBcastEnableOk

`func (o *NewMcastRateLimitSettingOpenApiVO) GetOtherBcastEnableOk() (*bool, bool)`

GetOtherBcastEnableOk returns a tuple with the OtherBcastEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherBcastEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) SetOtherBcastEnable(v bool)`

SetOtherBcastEnable sets OtherBcastEnable field to given value.

### HasOtherBcastEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) HasOtherBcastEnable() bool`

HasOtherBcastEnable returns a boolean if a field has been set.

### GetOtherBcastPps

`func (o *NewMcastRateLimitSettingOpenApiVO) GetOtherBcastPps() int32`

GetOtherBcastPps returns the OtherBcastPps field if non-nil, zero value otherwise.

### GetOtherBcastPpsOk

`func (o *NewMcastRateLimitSettingOpenApiVO) GetOtherBcastPpsOk() (*int32, bool)`

GetOtherBcastPpsOk returns a tuple with the OtherBcastPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherBcastPps

`func (o *NewMcastRateLimitSettingOpenApiVO) SetOtherBcastPps(v int32)`

SetOtherBcastPps sets OtherBcastPps field to given value.

### HasOtherBcastPps

`func (o *NewMcastRateLimitSettingOpenApiVO) HasOtherBcastPps() bool`

HasOtherBcastPps returns a boolean if a field has been set.

### GetOtherMcastEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) GetOtherMcastEnable() bool`

GetOtherMcastEnable returns the OtherMcastEnable field if non-nil, zero value otherwise.

### GetOtherMcastEnableOk

`func (o *NewMcastRateLimitSettingOpenApiVO) GetOtherMcastEnableOk() (*bool, bool)`

GetOtherMcastEnableOk returns a tuple with the OtherMcastEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherMcastEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) SetOtherMcastEnable(v bool)`

SetOtherMcastEnable sets OtherMcastEnable field to given value.

### HasOtherMcastEnable

`func (o *NewMcastRateLimitSettingOpenApiVO) HasOtherMcastEnable() bool`

HasOtherMcastEnable returns a boolean if a field has been set.

### GetOtherMcastPps

`func (o *NewMcastRateLimitSettingOpenApiVO) GetOtherMcastPps() int32`

GetOtherMcastPps returns the OtherMcastPps field if non-nil, zero value otherwise.

### GetOtherMcastPpsOk

`func (o *NewMcastRateLimitSettingOpenApiVO) GetOtherMcastPpsOk() (*int32, bool)`

GetOtherMcastPpsOk returns a tuple with the OtherMcastPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherMcastPps

`func (o *NewMcastRateLimitSettingOpenApiVO) SetOtherMcastPps(v int32)`

SetOtherMcastPps sets OtherMcastPps field to given value.

### HasOtherMcastPps

`func (o *NewMcastRateLimitSettingOpenApiVO) HasOtherMcastPps() bool`

HasOtherMcastPps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


