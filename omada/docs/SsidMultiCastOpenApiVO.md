# SsidMultiCastOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArpCastEnable** | Pointer to **bool** | Whether to enable ARP cast to unicast, which is enabled by default. True: enable, false: disable. | [optional] 
**ChannelUtil** | Pointer to **int32** | This item indicates that when the channel utilization reaches the threshold, multicast will no longer be converted to unicast, the default threshold is 100, and the value should be within the range of 0-100. | [optional] 
**FilterEnable** | Pointer to **bool** | Whether to enable the multicast filter switch, which is disabled by default. True: enable, false: disable. | [optional] 
**FilterMode** | Pointer to **int32** | This item indicates the status of the filtering protocol. The lowest bit indicates whether IGMP is enabled; the second lowest bit indicates whether MDNS is enabled; and the third lowest bit indicates whether Others is enabled. 1 means enable while 0 means disable. For example, 7(111) means that all are enabled; 1(001) means that only IGMP is enabled.  | [optional] 
**Ipv6CastEnable** | Pointer to **bool** | Whether to enable IPv6 multicast to unicast, which is enabled by default. True: enable, false: disable. | [optional] 
**MacGroupId** | Pointer to **string** | This field represents MAC Group Profile ID. MAC Group Profile can be created using Create a new group profile interface, and MAC Group Profile ID can be obtained from Get group profile list by type interface. | [optional] 
**MultiCastEnable** | Pointer to **bool** | Whether to enable multicast to unicast, which is enabled by default. True: enable, false: disable. | [optional] 

## Methods

### NewSsidMultiCastOpenApiVO

`func NewSsidMultiCastOpenApiVO() *SsidMultiCastOpenApiVO`

NewSsidMultiCastOpenApiVO instantiates a new SsidMultiCastOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsidMultiCastOpenApiVOWithDefaults

`func NewSsidMultiCastOpenApiVOWithDefaults() *SsidMultiCastOpenApiVO`

NewSsidMultiCastOpenApiVOWithDefaults instantiates a new SsidMultiCastOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArpCastEnable

`func (o *SsidMultiCastOpenApiVO) GetArpCastEnable() bool`

GetArpCastEnable returns the ArpCastEnable field if non-nil, zero value otherwise.

### GetArpCastEnableOk

`func (o *SsidMultiCastOpenApiVO) GetArpCastEnableOk() (*bool, bool)`

GetArpCastEnableOk returns a tuple with the ArpCastEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpCastEnable

`func (o *SsidMultiCastOpenApiVO) SetArpCastEnable(v bool)`

SetArpCastEnable sets ArpCastEnable field to given value.

### HasArpCastEnable

`func (o *SsidMultiCastOpenApiVO) HasArpCastEnable() bool`

HasArpCastEnable returns a boolean if a field has been set.

### GetChannelUtil

`func (o *SsidMultiCastOpenApiVO) GetChannelUtil() int32`

GetChannelUtil returns the ChannelUtil field if non-nil, zero value otherwise.

### GetChannelUtilOk

`func (o *SsidMultiCastOpenApiVO) GetChannelUtilOk() (*int32, bool)`

GetChannelUtilOk returns a tuple with the ChannelUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelUtil

`func (o *SsidMultiCastOpenApiVO) SetChannelUtil(v int32)`

SetChannelUtil sets ChannelUtil field to given value.

### HasChannelUtil

`func (o *SsidMultiCastOpenApiVO) HasChannelUtil() bool`

HasChannelUtil returns a boolean if a field has been set.

### GetFilterEnable

`func (o *SsidMultiCastOpenApiVO) GetFilterEnable() bool`

GetFilterEnable returns the FilterEnable field if non-nil, zero value otherwise.

### GetFilterEnableOk

`func (o *SsidMultiCastOpenApiVO) GetFilterEnableOk() (*bool, bool)`

GetFilterEnableOk returns a tuple with the FilterEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterEnable

`func (o *SsidMultiCastOpenApiVO) SetFilterEnable(v bool)`

SetFilterEnable sets FilterEnable field to given value.

### HasFilterEnable

`func (o *SsidMultiCastOpenApiVO) HasFilterEnable() bool`

HasFilterEnable returns a boolean if a field has been set.

### GetFilterMode

`func (o *SsidMultiCastOpenApiVO) GetFilterMode() int32`

GetFilterMode returns the FilterMode field if non-nil, zero value otherwise.

### GetFilterModeOk

`func (o *SsidMultiCastOpenApiVO) GetFilterModeOk() (*int32, bool)`

GetFilterModeOk returns a tuple with the FilterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterMode

`func (o *SsidMultiCastOpenApiVO) SetFilterMode(v int32)`

SetFilterMode sets FilterMode field to given value.

### HasFilterMode

`func (o *SsidMultiCastOpenApiVO) HasFilterMode() bool`

HasFilterMode returns a boolean if a field has been set.

### GetIpv6CastEnable

`func (o *SsidMultiCastOpenApiVO) GetIpv6CastEnable() bool`

GetIpv6CastEnable returns the Ipv6CastEnable field if non-nil, zero value otherwise.

### GetIpv6CastEnableOk

`func (o *SsidMultiCastOpenApiVO) GetIpv6CastEnableOk() (*bool, bool)`

GetIpv6CastEnableOk returns a tuple with the Ipv6CastEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6CastEnable

`func (o *SsidMultiCastOpenApiVO) SetIpv6CastEnable(v bool)`

SetIpv6CastEnable sets Ipv6CastEnable field to given value.

### HasIpv6CastEnable

`func (o *SsidMultiCastOpenApiVO) HasIpv6CastEnable() bool`

HasIpv6CastEnable returns a boolean if a field has been set.

### GetMacGroupId

`func (o *SsidMultiCastOpenApiVO) GetMacGroupId() string`

GetMacGroupId returns the MacGroupId field if non-nil, zero value otherwise.

### GetMacGroupIdOk

`func (o *SsidMultiCastOpenApiVO) GetMacGroupIdOk() (*string, bool)`

GetMacGroupIdOk returns a tuple with the MacGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacGroupId

`func (o *SsidMultiCastOpenApiVO) SetMacGroupId(v string)`

SetMacGroupId sets MacGroupId field to given value.

### HasMacGroupId

`func (o *SsidMultiCastOpenApiVO) HasMacGroupId() bool`

HasMacGroupId returns a boolean if a field has been set.

### GetMultiCastEnable

`func (o *SsidMultiCastOpenApiVO) GetMultiCastEnable() bool`

GetMultiCastEnable returns the MultiCastEnable field if non-nil, zero value otherwise.

### GetMultiCastEnableOk

`func (o *SsidMultiCastOpenApiVO) GetMultiCastEnableOk() (*bool, bool)`

GetMultiCastEnableOk returns a tuple with the MultiCastEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiCastEnable

`func (o *SsidMultiCastOpenApiVO) SetMultiCastEnable(v bool)`

SetMultiCastEnable sets MultiCastEnable field to given value.

### HasMultiCastEnable

`func (o *SsidMultiCastOpenApiVO) HasMultiCastEnable() bool`

HasMultiCastEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


