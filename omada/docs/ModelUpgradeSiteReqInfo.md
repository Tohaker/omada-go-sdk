# ModelUpgradeSiteReqInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | **int32** | Start page number. Start from 1. | 
**CurrentPageSize** | **int32** | Number of entries per page. It should be within the range of 1–100. | 
**Models** | [**[]ModelBaseInfo**](ModelBaseInfo.md) | Model base information list, include Model type information and currentVersion list | 

## Methods

### NewModelUpgradeSiteReqInfo

`func NewModelUpgradeSiteReqInfo(currentPage int32, currentPageSize int32, models []ModelBaseInfo, ) *ModelUpgradeSiteReqInfo`

NewModelUpgradeSiteReqInfo instantiates a new ModelUpgradeSiteReqInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelUpgradeSiteReqInfoWithDefaults

`func NewModelUpgradeSiteReqInfoWithDefaults() *ModelUpgradeSiteReqInfo`

NewModelUpgradeSiteReqInfoWithDefaults instantiates a new ModelUpgradeSiteReqInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *ModelUpgradeSiteReqInfo) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ModelUpgradeSiteReqInfo) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ModelUpgradeSiteReqInfo) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.


### GetCurrentPageSize

`func (o *ModelUpgradeSiteReqInfo) GetCurrentPageSize() int32`

GetCurrentPageSize returns the CurrentPageSize field if non-nil, zero value otherwise.

### GetCurrentPageSizeOk

`func (o *ModelUpgradeSiteReqInfo) GetCurrentPageSizeOk() (*int32, bool)`

GetCurrentPageSizeOk returns a tuple with the CurrentPageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPageSize

`func (o *ModelUpgradeSiteReqInfo) SetCurrentPageSize(v int32)`

SetCurrentPageSize sets CurrentPageSize field to given value.


### GetModels

`func (o *ModelUpgradeSiteReqInfo) GetModels() []ModelBaseInfo`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ModelUpgradeSiteReqInfo) GetModelsOk() (*[]ModelBaseInfo, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ModelUpgradeSiteReqInfo) SetModels(v []ModelBaseInfo)`

SetModels sets Models field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


