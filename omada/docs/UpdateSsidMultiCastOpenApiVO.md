# UpdateSsidMultiCastOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArpCastEnable** | **bool** | Whether to enable ARP cast to unicast, which is disabled by default; true: enable, false: disable | 
**ChannelUtil** | **int32** | This item indicates that when the channel utilization reaches the threshold, multicast will no longer be converted to unicast, the default threshold is 100, and the value should be within the range of 0-100. | 
**FilterEnable** | **bool** | Whether to enable the multicast filter switch, which is disabled by default; true: enable, false: disable | 
**FilterMode** | Pointer to **int32** | This item indicates the status of the filtering protocol, the lowest low bit indicates whether IGMP is enabled, the second low bit indicates whether MDNS is enabled, and the third low bit indicates whether Others is enabled; 1 means enable while 0 means disable. For example, 7(111) means that all are enabled, 1(001) means that only IGMP is enabled.  | [optional] 
**Ipv6CastEnable** | **bool** | Whether to enable IPv6 multicast to unicast, which is enabled by default; true: enable, false: disable | 
**MacGroupId** | Pointer to **string** | This field represents MAC Group Profile ID. MAC Group Profile can be created using Create a new group profile interface, and MAC Group Profile ID can be obtained from Get group profile list by type interface. | [optional] 
**MultiCastEnable** | **bool** | Whether to enable multicast to unicast, which is disabled by default; true: enable, false: disable | 

## Methods

### NewUpdateSsidMultiCastOpenApiVO

`func NewUpdateSsidMultiCastOpenApiVO(arpCastEnable bool, channelUtil int32, filterEnable bool, ipv6CastEnable bool, multiCastEnable bool, ) *UpdateSsidMultiCastOpenApiVO`

NewUpdateSsidMultiCastOpenApiVO instantiates a new UpdateSsidMultiCastOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSsidMultiCastOpenApiVOWithDefaults

`func NewUpdateSsidMultiCastOpenApiVOWithDefaults() *UpdateSsidMultiCastOpenApiVO`

NewUpdateSsidMultiCastOpenApiVOWithDefaults instantiates a new UpdateSsidMultiCastOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArpCastEnable

`func (o *UpdateSsidMultiCastOpenApiVO) GetArpCastEnable() bool`

GetArpCastEnable returns the ArpCastEnable field if non-nil, zero value otherwise.

### GetArpCastEnableOk

`func (o *UpdateSsidMultiCastOpenApiVO) GetArpCastEnableOk() (*bool, bool)`

GetArpCastEnableOk returns a tuple with the ArpCastEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpCastEnable

`func (o *UpdateSsidMultiCastOpenApiVO) SetArpCastEnable(v bool)`

SetArpCastEnable sets ArpCastEnable field to given value.


### GetChannelUtil

`func (o *UpdateSsidMultiCastOpenApiVO) GetChannelUtil() int32`

GetChannelUtil returns the ChannelUtil field if non-nil, zero value otherwise.

### GetChannelUtilOk

`func (o *UpdateSsidMultiCastOpenApiVO) GetChannelUtilOk() (*int32, bool)`

GetChannelUtilOk returns a tuple with the ChannelUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelUtil

`func (o *UpdateSsidMultiCastOpenApiVO) SetChannelUtil(v int32)`

SetChannelUtil sets ChannelUtil field to given value.


### GetFilterEnable

`func (o *UpdateSsidMultiCastOpenApiVO) GetFilterEnable() bool`

GetFilterEnable returns the FilterEnable field if non-nil, zero value otherwise.

### GetFilterEnableOk

`func (o *UpdateSsidMultiCastOpenApiVO) GetFilterEnableOk() (*bool, bool)`

GetFilterEnableOk returns a tuple with the FilterEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterEnable

`func (o *UpdateSsidMultiCastOpenApiVO) SetFilterEnable(v bool)`

SetFilterEnable sets FilterEnable field to given value.


### GetFilterMode

`func (o *UpdateSsidMultiCastOpenApiVO) GetFilterMode() int32`

GetFilterMode returns the FilterMode field if non-nil, zero value otherwise.

### GetFilterModeOk

`func (o *UpdateSsidMultiCastOpenApiVO) GetFilterModeOk() (*int32, bool)`

GetFilterModeOk returns a tuple with the FilterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterMode

`func (o *UpdateSsidMultiCastOpenApiVO) SetFilterMode(v int32)`

SetFilterMode sets FilterMode field to given value.

### HasFilterMode

`func (o *UpdateSsidMultiCastOpenApiVO) HasFilterMode() bool`

HasFilterMode returns a boolean if a field has been set.

### GetIpv6CastEnable

`func (o *UpdateSsidMultiCastOpenApiVO) GetIpv6CastEnable() bool`

GetIpv6CastEnable returns the Ipv6CastEnable field if non-nil, zero value otherwise.

### GetIpv6CastEnableOk

`func (o *UpdateSsidMultiCastOpenApiVO) GetIpv6CastEnableOk() (*bool, bool)`

GetIpv6CastEnableOk returns a tuple with the Ipv6CastEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6CastEnable

`func (o *UpdateSsidMultiCastOpenApiVO) SetIpv6CastEnable(v bool)`

SetIpv6CastEnable sets Ipv6CastEnable field to given value.


### GetMacGroupId

`func (o *UpdateSsidMultiCastOpenApiVO) GetMacGroupId() string`

GetMacGroupId returns the MacGroupId field if non-nil, zero value otherwise.

### GetMacGroupIdOk

`func (o *UpdateSsidMultiCastOpenApiVO) GetMacGroupIdOk() (*string, bool)`

GetMacGroupIdOk returns a tuple with the MacGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacGroupId

`func (o *UpdateSsidMultiCastOpenApiVO) SetMacGroupId(v string)`

SetMacGroupId sets MacGroupId field to given value.

### HasMacGroupId

`func (o *UpdateSsidMultiCastOpenApiVO) HasMacGroupId() bool`

HasMacGroupId returns a boolean if a field has been set.

### GetMultiCastEnable

`func (o *UpdateSsidMultiCastOpenApiVO) GetMultiCastEnable() bool`

GetMultiCastEnable returns the MultiCastEnable field if non-nil, zero value otherwise.

### GetMultiCastEnableOk

`func (o *UpdateSsidMultiCastOpenApiVO) GetMultiCastEnableOk() (*bool, bool)`

GetMultiCastEnableOk returns a tuple with the MultiCastEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiCastEnable

`func (o *UpdateSsidMultiCastOpenApiVO) SetMultiCastEnable(v bool)`

SetMultiCastEnable sets MultiCastEnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


