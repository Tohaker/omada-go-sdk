# MultiOswPortSelectVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to **map[string]map[string]interface{}** | Filter conditions in the form of Map.It is effected when [selectAll] is &#39;true&#39;, filter key is the filter field and the value is the filter content.Filter fields include: [connectedStatus], [networkMode], [poeDisplayType], [linkSpeed], [duplex],[switchMac], [switchStatusCategory], [switchSupportPoe], [nativeNetworkId], [networkTagsSetting], [profileId], [operation], [tagIds]. | [optional] 
**SearchKey** | Pointer to **string** | The keywords of the searchIt is effected when [selectAll] is &#39;true&#39;. | [optional] 
**SelectAll** | **bool** | Indicates whether select all switch ports.false: include selected switch ports and lags in Parameter [switchList], true: all switch ports and lags but exclude selected switch ports and lags in Parameter [switchList]. | 
**SwitchList** | [**[]OswPortLagListVO**](OswPortLagListVO.md) | Switches Port And Lag List | 

## Methods

### NewMultiOswPortSelectVO

`func NewMultiOswPortSelectVO(selectAll bool, switchList []OswPortLagListVO, ) *MultiOswPortSelectVO`

NewMultiOswPortSelectVO instantiates a new MultiOswPortSelectVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiOswPortSelectVOWithDefaults

`func NewMultiOswPortSelectVOWithDefaults() *MultiOswPortSelectVO`

NewMultiOswPortSelectVOWithDefaults instantiates a new MultiOswPortSelectVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *MultiOswPortSelectVO) GetFilters() map[string]map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *MultiOswPortSelectVO) GetFiltersOk() (*map[string]map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *MultiOswPortSelectVO) SetFilters(v map[string]map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *MultiOswPortSelectVO) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSearchKey

`func (o *MultiOswPortSelectVO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *MultiOswPortSelectVO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *MultiOswPortSelectVO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *MultiOswPortSelectVO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSelectAll

`func (o *MultiOswPortSelectVO) GetSelectAll() bool`

GetSelectAll returns the SelectAll field if non-nil, zero value otherwise.

### GetSelectAllOk

`func (o *MultiOswPortSelectVO) GetSelectAllOk() (*bool, bool)`

GetSelectAllOk returns a tuple with the SelectAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectAll

`func (o *MultiOswPortSelectVO) SetSelectAll(v bool)`

SetSelectAll sets SelectAll field to given value.


### GetSwitchList

`func (o *MultiOswPortSelectVO) GetSwitchList() []OswPortLagListVO`

GetSwitchList returns the SwitchList field if non-nil, zero value otherwise.

### GetSwitchListOk

`func (o *MultiOswPortSelectVO) GetSwitchListOk() (*[]OswPortLagListVO, bool)`

GetSwitchListOk returns a tuple with the SwitchList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchList

`func (o *MultiOswPortSelectVO) SetSwitchList(v []OswPortLagListVO)`

SetSwitchList sets SwitchList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


