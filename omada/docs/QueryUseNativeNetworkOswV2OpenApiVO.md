# QueryUseNativeNetworkOswV2OpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | List of VLAN and Interface network IDs. The valid network IDs can be obtained from \&quot;Get LAN network list\&quot;. | [optional] 
**Page** | **int32** | Start from 1. | 
**PageSize** | **int32** | It should be within the range of 1–1000. | 
**SelectType** | Pointer to **string** | selectType should be a value as follows: all: select all; include: include ids; exclude: exclude ids. | [optional] 

## Methods

### NewQueryUseNativeNetworkOswV2OpenApiVO

`func NewQueryUseNativeNetworkOswV2OpenApiVO(page int32, pageSize int32, ) *QueryUseNativeNetworkOswV2OpenApiVO`

NewQueryUseNativeNetworkOswV2OpenApiVO instantiates a new QueryUseNativeNetworkOswV2OpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryUseNativeNetworkOswV2OpenApiVOWithDefaults

`func NewQueryUseNativeNetworkOswV2OpenApiVOWithDefaults() *QueryUseNativeNetworkOswV2OpenApiVO`

NewQueryUseNativeNetworkOswV2OpenApiVOWithDefaults instantiates a new QueryUseNativeNetworkOswV2OpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *QueryUseNativeNetworkOswV2OpenApiVO) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *QueryUseNativeNetworkOswV2OpenApiVO) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *QueryUseNativeNetworkOswV2OpenApiVO) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *QueryUseNativeNetworkOswV2OpenApiVO) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetPage

`func (o *QueryUseNativeNetworkOswV2OpenApiVO) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *QueryUseNativeNetworkOswV2OpenApiVO) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *QueryUseNativeNetworkOswV2OpenApiVO) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *QueryUseNativeNetworkOswV2OpenApiVO) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *QueryUseNativeNetworkOswV2OpenApiVO) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *QueryUseNativeNetworkOswV2OpenApiVO) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetSelectType

`func (o *QueryUseNativeNetworkOswV2OpenApiVO) GetSelectType() string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *QueryUseNativeNetworkOswV2OpenApiVO) GetSelectTypeOk() (*string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *QueryUseNativeNetworkOswV2OpenApiVO) SetSelectType(v string)`

SetSelectType sets SelectType field to given value.

### HasSelectType

`func (o *QueryUseNativeNetworkOswV2OpenApiVO) HasSelectType() bool`

HasSelectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


