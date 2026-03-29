# MultiOswPortSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to **map[string]map[string]interface{}** | Filter conditions in the form of Map.It is effected when [selectAll] is &#39;true&#39;, filter key is the filter field and the value is the filter content.Filter fields include: [connectedStatus], [networkMode], [poeDisplayType], [linkSpeed], [duplex],[switchMac], [switchStatusCategory], [switchSupportPoe], [nativeNetworkId], [networkTagsSetting], [profileId], [operation], [tagIds]. | [optional] 
**NativeBridgeVlan** | Pointer to **int32** | Native Network Bridge Vlan. | [optional] 
**NativeNetworkId** | Pointer to **string** | Native Network ID, Native Network cannot be selected from Tagged Networks or Untagged Networks. | [optional] 
**NetworkTagsSetting** | Pointer to **int32** | Network Tags Setting should be a value as follows: 0: Allow All; 1: Block All; 2: Custom | [optional] 
**ProfileId** | Pointer to **string** | Profile ID | [optional] 
**ProfileOverrideEnable** | Pointer to **bool** | Indicates whether to enable Profile Override | [optional] 
**SearchKey** | Pointer to **string** | The keywords of the searchIt is effected when [selectAll] is &#39;true&#39;. | [optional] 
**SelectAll** | **bool** | Indicates whether select all switch ports.false: include selected switch ports and lags in Parameter [switchList], true: all switch ports and lags but exclude selected switch ports and lags in Parameter [switchList]. | 
**SwitchList** | [**[]OswPortLagListVO**](OswPortLagListVO.md) | Switch List with port and LAG info | 
**TagBridgeVlanMap** | Pointer to **map[string][]int32** | Tag Network Bridge Vlan Map | [optional] 
**TagIds** | Pointer to **[]string** | Tag ID List | [optional] 
**TagNetworkIds** | Pointer to **[]string** | Tag Network IDs | [optional] 
**UntagBridgeVlanMap** | Pointer to **map[string][]int32** | Untag Network Bridge Vlan Map | [optional] 
**UntagNetworkIds** | Pointer to **[]string** | Untag Network IDs | [optional] 
**VoiceBridgeVlan** | Pointer to **int32** | Voice Network Bridge Vlan | [optional] 
**VoiceNetworkEnable** | Pointer to **bool** | Indicates whether voice network is enabled | [optional] 
**VoiceNetworkId** | Pointer to **string** | Voice Network ID | [optional] 

## Methods

### NewMultiOswPortSettingOpenApiVO

`func NewMultiOswPortSettingOpenApiVO(selectAll bool, switchList []OswPortLagListVO, ) *MultiOswPortSettingOpenApiVO`

NewMultiOswPortSettingOpenApiVO instantiates a new MultiOswPortSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiOswPortSettingOpenApiVOWithDefaults

`func NewMultiOswPortSettingOpenApiVOWithDefaults() *MultiOswPortSettingOpenApiVO`

NewMultiOswPortSettingOpenApiVOWithDefaults instantiates a new MultiOswPortSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *MultiOswPortSettingOpenApiVO) GetFilters() map[string]map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *MultiOswPortSettingOpenApiVO) GetFiltersOk() (*map[string]map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *MultiOswPortSettingOpenApiVO) SetFilters(v map[string]map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *MultiOswPortSettingOpenApiVO) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetNativeBridgeVlan

`func (o *MultiOswPortSettingOpenApiVO) GetNativeBridgeVlan() int32`

GetNativeBridgeVlan returns the NativeBridgeVlan field if non-nil, zero value otherwise.

### GetNativeBridgeVlanOk

`func (o *MultiOswPortSettingOpenApiVO) GetNativeBridgeVlanOk() (*int32, bool)`

GetNativeBridgeVlanOk returns a tuple with the NativeBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeBridgeVlan

`func (o *MultiOswPortSettingOpenApiVO) SetNativeBridgeVlan(v int32)`

SetNativeBridgeVlan sets NativeBridgeVlan field to given value.

### HasNativeBridgeVlan

`func (o *MultiOswPortSettingOpenApiVO) HasNativeBridgeVlan() bool`

HasNativeBridgeVlan returns a boolean if a field has been set.

### GetNativeNetworkId

`func (o *MultiOswPortSettingOpenApiVO) GetNativeNetworkId() string`

GetNativeNetworkId returns the NativeNetworkId field if non-nil, zero value otherwise.

### GetNativeNetworkIdOk

`func (o *MultiOswPortSettingOpenApiVO) GetNativeNetworkIdOk() (*string, bool)`

GetNativeNetworkIdOk returns a tuple with the NativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeNetworkId

`func (o *MultiOswPortSettingOpenApiVO) SetNativeNetworkId(v string)`

SetNativeNetworkId sets NativeNetworkId field to given value.

### HasNativeNetworkId

`func (o *MultiOswPortSettingOpenApiVO) HasNativeNetworkId() bool`

HasNativeNetworkId returns a boolean if a field has been set.

### GetNetworkTagsSetting

`func (o *MultiOswPortSettingOpenApiVO) GetNetworkTagsSetting() int32`

GetNetworkTagsSetting returns the NetworkTagsSetting field if non-nil, zero value otherwise.

### GetNetworkTagsSettingOk

`func (o *MultiOswPortSettingOpenApiVO) GetNetworkTagsSettingOk() (*int32, bool)`

GetNetworkTagsSettingOk returns a tuple with the NetworkTagsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTagsSetting

`func (o *MultiOswPortSettingOpenApiVO) SetNetworkTagsSetting(v int32)`

SetNetworkTagsSetting sets NetworkTagsSetting field to given value.

### HasNetworkTagsSetting

`func (o *MultiOswPortSettingOpenApiVO) HasNetworkTagsSetting() bool`

HasNetworkTagsSetting returns a boolean if a field has been set.

### GetProfileId

`func (o *MultiOswPortSettingOpenApiVO) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *MultiOswPortSettingOpenApiVO) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *MultiOswPortSettingOpenApiVO) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *MultiOswPortSettingOpenApiVO) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileOverrideEnable

`func (o *MultiOswPortSettingOpenApiVO) GetProfileOverrideEnable() bool`

GetProfileOverrideEnable returns the ProfileOverrideEnable field if non-nil, zero value otherwise.

### GetProfileOverrideEnableOk

`func (o *MultiOswPortSettingOpenApiVO) GetProfileOverrideEnableOk() (*bool, bool)`

GetProfileOverrideEnableOk returns a tuple with the ProfileOverrideEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileOverrideEnable

`func (o *MultiOswPortSettingOpenApiVO) SetProfileOverrideEnable(v bool)`

SetProfileOverrideEnable sets ProfileOverrideEnable field to given value.

### HasProfileOverrideEnable

`func (o *MultiOswPortSettingOpenApiVO) HasProfileOverrideEnable() bool`

HasProfileOverrideEnable returns a boolean if a field has been set.

### GetSearchKey

`func (o *MultiOswPortSettingOpenApiVO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *MultiOswPortSettingOpenApiVO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *MultiOswPortSettingOpenApiVO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *MultiOswPortSettingOpenApiVO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSelectAll

`func (o *MultiOswPortSettingOpenApiVO) GetSelectAll() bool`

GetSelectAll returns the SelectAll field if non-nil, zero value otherwise.

### GetSelectAllOk

`func (o *MultiOswPortSettingOpenApiVO) GetSelectAllOk() (*bool, bool)`

GetSelectAllOk returns a tuple with the SelectAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectAll

`func (o *MultiOswPortSettingOpenApiVO) SetSelectAll(v bool)`

SetSelectAll sets SelectAll field to given value.


### GetSwitchList

`func (o *MultiOswPortSettingOpenApiVO) GetSwitchList() []OswPortLagListVO`

GetSwitchList returns the SwitchList field if non-nil, zero value otherwise.

### GetSwitchListOk

`func (o *MultiOswPortSettingOpenApiVO) GetSwitchListOk() (*[]OswPortLagListVO, bool)`

GetSwitchListOk returns a tuple with the SwitchList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchList

`func (o *MultiOswPortSettingOpenApiVO) SetSwitchList(v []OswPortLagListVO)`

SetSwitchList sets SwitchList field to given value.


### GetTagBridgeVlanMap

`func (o *MultiOswPortSettingOpenApiVO) GetTagBridgeVlanMap() map[string][]int32`

GetTagBridgeVlanMap returns the TagBridgeVlanMap field if non-nil, zero value otherwise.

### GetTagBridgeVlanMapOk

`func (o *MultiOswPortSettingOpenApiVO) GetTagBridgeVlanMapOk() (*map[string][]int32, bool)`

GetTagBridgeVlanMapOk returns a tuple with the TagBridgeVlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagBridgeVlanMap

`func (o *MultiOswPortSettingOpenApiVO) SetTagBridgeVlanMap(v map[string][]int32)`

SetTagBridgeVlanMap sets TagBridgeVlanMap field to given value.

### HasTagBridgeVlanMap

`func (o *MultiOswPortSettingOpenApiVO) HasTagBridgeVlanMap() bool`

HasTagBridgeVlanMap returns a boolean if a field has been set.

### GetTagIds

`func (o *MultiOswPortSettingOpenApiVO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *MultiOswPortSettingOpenApiVO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *MultiOswPortSettingOpenApiVO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *MultiOswPortSettingOpenApiVO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTagNetworkIds

`func (o *MultiOswPortSettingOpenApiVO) GetTagNetworkIds() []string`

GetTagNetworkIds returns the TagNetworkIds field if non-nil, zero value otherwise.

### GetTagNetworkIdsOk

`func (o *MultiOswPortSettingOpenApiVO) GetTagNetworkIdsOk() (*[]string, bool)`

GetTagNetworkIdsOk returns a tuple with the TagNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNetworkIds

`func (o *MultiOswPortSettingOpenApiVO) SetTagNetworkIds(v []string)`

SetTagNetworkIds sets TagNetworkIds field to given value.

### HasTagNetworkIds

`func (o *MultiOswPortSettingOpenApiVO) HasTagNetworkIds() bool`

HasTagNetworkIds returns a boolean if a field has been set.

### GetUntagBridgeVlanMap

`func (o *MultiOswPortSettingOpenApiVO) GetUntagBridgeVlanMap() map[string][]int32`

GetUntagBridgeVlanMap returns the UntagBridgeVlanMap field if non-nil, zero value otherwise.

### GetUntagBridgeVlanMapOk

`func (o *MultiOswPortSettingOpenApiVO) GetUntagBridgeVlanMapOk() (*map[string][]int32, bool)`

GetUntagBridgeVlanMapOk returns a tuple with the UntagBridgeVlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagBridgeVlanMap

`func (o *MultiOswPortSettingOpenApiVO) SetUntagBridgeVlanMap(v map[string][]int32)`

SetUntagBridgeVlanMap sets UntagBridgeVlanMap field to given value.

### HasUntagBridgeVlanMap

`func (o *MultiOswPortSettingOpenApiVO) HasUntagBridgeVlanMap() bool`

HasUntagBridgeVlanMap returns a boolean if a field has been set.

### GetUntagNetworkIds

`func (o *MultiOswPortSettingOpenApiVO) GetUntagNetworkIds() []string`

GetUntagNetworkIds returns the UntagNetworkIds field if non-nil, zero value otherwise.

### GetUntagNetworkIdsOk

`func (o *MultiOswPortSettingOpenApiVO) GetUntagNetworkIdsOk() (*[]string, bool)`

GetUntagNetworkIdsOk returns a tuple with the UntagNetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagNetworkIds

`func (o *MultiOswPortSettingOpenApiVO) SetUntagNetworkIds(v []string)`

SetUntagNetworkIds sets UntagNetworkIds field to given value.

### HasUntagNetworkIds

`func (o *MultiOswPortSettingOpenApiVO) HasUntagNetworkIds() bool`

HasUntagNetworkIds returns a boolean if a field has been set.

### GetVoiceBridgeVlan

`func (o *MultiOswPortSettingOpenApiVO) GetVoiceBridgeVlan() int32`

GetVoiceBridgeVlan returns the VoiceBridgeVlan field if non-nil, zero value otherwise.

### GetVoiceBridgeVlanOk

`func (o *MultiOswPortSettingOpenApiVO) GetVoiceBridgeVlanOk() (*int32, bool)`

GetVoiceBridgeVlanOk returns a tuple with the VoiceBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceBridgeVlan

`func (o *MultiOswPortSettingOpenApiVO) SetVoiceBridgeVlan(v int32)`

SetVoiceBridgeVlan sets VoiceBridgeVlan field to given value.

### HasVoiceBridgeVlan

`func (o *MultiOswPortSettingOpenApiVO) HasVoiceBridgeVlan() bool`

HasVoiceBridgeVlan returns a boolean if a field has been set.

### GetVoiceNetworkEnable

`func (o *MultiOswPortSettingOpenApiVO) GetVoiceNetworkEnable() bool`

GetVoiceNetworkEnable returns the VoiceNetworkEnable field if non-nil, zero value otherwise.

### GetVoiceNetworkEnableOk

`func (o *MultiOswPortSettingOpenApiVO) GetVoiceNetworkEnableOk() (*bool, bool)`

GetVoiceNetworkEnableOk returns a tuple with the VoiceNetworkEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNetworkEnable

`func (o *MultiOswPortSettingOpenApiVO) SetVoiceNetworkEnable(v bool)`

SetVoiceNetworkEnable sets VoiceNetworkEnable field to given value.

### HasVoiceNetworkEnable

`func (o *MultiOswPortSettingOpenApiVO) HasVoiceNetworkEnable() bool`

HasVoiceNetworkEnable returns a boolean if a field has been set.

### GetVoiceNetworkId

`func (o *MultiOswPortSettingOpenApiVO) GetVoiceNetworkId() string`

GetVoiceNetworkId returns the VoiceNetworkId field if non-nil, zero value otherwise.

### GetVoiceNetworkIdOk

`func (o *MultiOswPortSettingOpenApiVO) GetVoiceNetworkIdOk() (*string, bool)`

GetVoiceNetworkIdOk returns a tuple with the VoiceNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNetworkId

`func (o *MultiOswPortSettingOpenApiVO) SetVoiceNetworkId(v string)`

SetVoiceNetworkId sets VoiceNetworkId field to given value.

### HasVoiceNetworkId

`func (o *MultiOswPortSettingOpenApiVO) HasVoiceNetworkId() bool`

HasVoiceNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


