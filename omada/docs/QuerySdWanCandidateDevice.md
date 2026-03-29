# QuerySdWanCandidateDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeDeviceMacs** | Pointer to **[]string** | A list of exclude Device MAC for the candidate. | [optional] 
**Filters** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**GroupId** | Pointer to **string** | The ID of the SD-WAN group. | [optional] 
**Page** | **int32** | Start from 1. | 
**PageSize** | **int32** | It should be within the range of 1–1000. | 
**Role** | **int32** | The role in SD-WAN Group. 0:Hub,1:Spoke | 
**SearchField** | Pointer to **string** |  | [optional] 
**SearchKey** | Pointer to **string** | Look for a specific piece of data. | [optional] 
**Sorts** | Pointer to **map[string]string** |  | [optional] 
**TunnelLimit** | Pointer to **int32** | The maximum number of VPN tunnels that can be created. | [optional] 

## Methods

### NewQuerySdWanCandidateDevice

`func NewQuerySdWanCandidateDevice(page int32, pageSize int32, role int32, ) *QuerySdWanCandidateDevice`

NewQuerySdWanCandidateDevice instantiates a new QuerySdWanCandidateDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySdWanCandidateDeviceWithDefaults

`func NewQuerySdWanCandidateDeviceWithDefaults() *QuerySdWanCandidateDevice`

NewQuerySdWanCandidateDeviceWithDefaults instantiates a new QuerySdWanCandidateDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeDeviceMacs

`func (o *QuerySdWanCandidateDevice) GetExcludeDeviceMacs() []string`

GetExcludeDeviceMacs returns the ExcludeDeviceMacs field if non-nil, zero value otherwise.

### GetExcludeDeviceMacsOk

`func (o *QuerySdWanCandidateDevice) GetExcludeDeviceMacsOk() (*[]string, bool)`

GetExcludeDeviceMacsOk returns a tuple with the ExcludeDeviceMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeDeviceMacs

`func (o *QuerySdWanCandidateDevice) SetExcludeDeviceMacs(v []string)`

SetExcludeDeviceMacs sets ExcludeDeviceMacs field to given value.

### HasExcludeDeviceMacs

`func (o *QuerySdWanCandidateDevice) HasExcludeDeviceMacs() bool`

HasExcludeDeviceMacs returns a boolean if a field has been set.

### GetFilters

`func (o *QuerySdWanCandidateDevice) GetFilters() map[string]map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *QuerySdWanCandidateDevice) GetFiltersOk() (*map[string]map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *QuerySdWanCandidateDevice) SetFilters(v map[string]map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *QuerySdWanCandidateDevice) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetGroupId

`func (o *QuerySdWanCandidateDevice) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *QuerySdWanCandidateDevice) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *QuerySdWanCandidateDevice) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *QuerySdWanCandidateDevice) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetPage

`func (o *QuerySdWanCandidateDevice) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *QuerySdWanCandidateDevice) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *QuerySdWanCandidateDevice) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *QuerySdWanCandidateDevice) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *QuerySdWanCandidateDevice) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *QuerySdWanCandidateDevice) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetRole

`func (o *QuerySdWanCandidateDevice) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *QuerySdWanCandidateDevice) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *QuerySdWanCandidateDevice) SetRole(v int32)`

SetRole sets Role field to given value.


### GetSearchField

`func (o *QuerySdWanCandidateDevice) GetSearchField() string`

GetSearchField returns the SearchField field if non-nil, zero value otherwise.

### GetSearchFieldOk

`func (o *QuerySdWanCandidateDevice) GetSearchFieldOk() (*string, bool)`

GetSearchFieldOk returns a tuple with the SearchField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchField

`func (o *QuerySdWanCandidateDevice) SetSearchField(v string)`

SetSearchField sets SearchField field to given value.

### HasSearchField

`func (o *QuerySdWanCandidateDevice) HasSearchField() bool`

HasSearchField returns a boolean if a field has been set.

### GetSearchKey

`func (o *QuerySdWanCandidateDevice) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *QuerySdWanCandidateDevice) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *QuerySdWanCandidateDevice) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *QuerySdWanCandidateDevice) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSorts

`func (o *QuerySdWanCandidateDevice) GetSorts() map[string]string`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *QuerySdWanCandidateDevice) GetSortsOk() (*map[string]string, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *QuerySdWanCandidateDevice) SetSorts(v map[string]string)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *QuerySdWanCandidateDevice) HasSorts() bool`

HasSorts returns a boolean if a field has been set.

### GetTunnelLimit

`func (o *QuerySdWanCandidateDevice) GetTunnelLimit() int32`

GetTunnelLimit returns the TunnelLimit field if non-nil, zero value otherwise.

### GetTunnelLimitOk

`func (o *QuerySdWanCandidateDevice) GetTunnelLimitOk() (*int32, bool)`

GetTunnelLimitOk returns a tuple with the TunnelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelLimit

`func (o *QuerySdWanCandidateDevice) SetTunnelLimit(v int32)`

SetTunnelLimit sets TunnelLimit field to given value.

### HasTunnelLimit

`func (o *QuerySdWanCandidateDevice) HasTunnelLimit() bool`

HasTunnelLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


