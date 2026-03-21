# QueryUseNativeNetworkOswOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | List of VLAN and Interface network IDs. The valid network IDs can be obtained from \&quot;Get LAN network list\&quot;. When selectAll is true, it means select all except specified IDs; When selectAll is false, it means select specified IDs. | [optional] 
**Page** | **int32** | Start from 1. | 
**PageSize** | **int32** | It should be within the range of 1–1000. | 
**SelectAll** | Pointer to **bool** | Select all VLAN networks | [optional] 

## Methods

### NewQueryUseNativeNetworkOswOpenApiVO

`func NewQueryUseNativeNetworkOswOpenApiVO(page int32, pageSize int32, ) *QueryUseNativeNetworkOswOpenApiVO`

NewQueryUseNativeNetworkOswOpenApiVO instantiates a new QueryUseNativeNetworkOswOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryUseNativeNetworkOswOpenApiVOWithDefaults

`func NewQueryUseNativeNetworkOswOpenApiVOWithDefaults() *QueryUseNativeNetworkOswOpenApiVO`

NewQueryUseNativeNetworkOswOpenApiVOWithDefaults instantiates a new QueryUseNativeNetworkOswOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *QueryUseNativeNetworkOswOpenApiVO) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *QueryUseNativeNetworkOswOpenApiVO) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *QueryUseNativeNetworkOswOpenApiVO) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *QueryUseNativeNetworkOswOpenApiVO) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetPage

`func (o *QueryUseNativeNetworkOswOpenApiVO) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *QueryUseNativeNetworkOswOpenApiVO) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *QueryUseNativeNetworkOswOpenApiVO) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *QueryUseNativeNetworkOswOpenApiVO) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *QueryUseNativeNetworkOswOpenApiVO) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *QueryUseNativeNetworkOswOpenApiVO) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetSelectAll

`func (o *QueryUseNativeNetworkOswOpenApiVO) GetSelectAll() bool`

GetSelectAll returns the SelectAll field if non-nil, zero value otherwise.

### GetSelectAllOk

`func (o *QueryUseNativeNetworkOswOpenApiVO) GetSelectAllOk() (*bool, bool)`

GetSelectAllOk returns a tuple with the SelectAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectAll

`func (o *QueryUseNativeNetworkOswOpenApiVO) SetSelectAll(v bool)`

SetSelectAll sets SelectAll field to given value.

### HasSelectAll

`func (o *QueryUseNativeNetworkOswOpenApiVO) HasSelectAll() bool`

HasSelectAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


