# ModelUpgradeSiteInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllSiteIds** | Pointer to **[]string** | All siteIDs in the query result. | [optional] 
**CurrentPage** | Pointer to **int32** | Start page number. Start from 1. | [optional] 
**CurrentPageSize** | Pointer to **int32** | Number of entries per page. It should be within the range of 1–100. | [optional] 
**Data** | Pointer to [**[]SiteBasicInfo**](SiteBasicInfo.md) | Site basic information on one page. | [optional] 
**TotalRows** | Pointer to **int32** | Total rows. | [optional] 

## Methods

### NewModelUpgradeSiteInfo

`func NewModelUpgradeSiteInfo() *ModelUpgradeSiteInfo`

NewModelUpgradeSiteInfo instantiates a new ModelUpgradeSiteInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelUpgradeSiteInfoWithDefaults

`func NewModelUpgradeSiteInfoWithDefaults() *ModelUpgradeSiteInfo`

NewModelUpgradeSiteInfoWithDefaults instantiates a new ModelUpgradeSiteInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllSiteIds

`func (o *ModelUpgradeSiteInfo) GetAllSiteIds() []string`

GetAllSiteIds returns the AllSiteIds field if non-nil, zero value otherwise.

### GetAllSiteIdsOk

`func (o *ModelUpgradeSiteInfo) GetAllSiteIdsOk() (*[]string, bool)`

GetAllSiteIdsOk returns a tuple with the AllSiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSiteIds

`func (o *ModelUpgradeSiteInfo) SetAllSiteIds(v []string)`

SetAllSiteIds sets AllSiteIds field to given value.

### HasAllSiteIds

`func (o *ModelUpgradeSiteInfo) HasAllSiteIds() bool`

HasAllSiteIds returns a boolean if a field has been set.

### GetCurrentPage

`func (o *ModelUpgradeSiteInfo) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ModelUpgradeSiteInfo) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ModelUpgradeSiteInfo) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *ModelUpgradeSiteInfo) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentPageSize

`func (o *ModelUpgradeSiteInfo) GetCurrentPageSize() int32`

GetCurrentPageSize returns the CurrentPageSize field if non-nil, zero value otherwise.

### GetCurrentPageSizeOk

`func (o *ModelUpgradeSiteInfo) GetCurrentPageSizeOk() (*int32, bool)`

GetCurrentPageSizeOk returns a tuple with the CurrentPageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPageSize

`func (o *ModelUpgradeSiteInfo) SetCurrentPageSize(v int32)`

SetCurrentPageSize sets CurrentPageSize field to given value.

### HasCurrentPageSize

`func (o *ModelUpgradeSiteInfo) HasCurrentPageSize() bool`

HasCurrentPageSize returns a boolean if a field has been set.

### GetData

`func (o *ModelUpgradeSiteInfo) GetData() []SiteBasicInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModelUpgradeSiteInfo) GetDataOk() (*[]SiteBasicInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModelUpgradeSiteInfo) SetData(v []SiteBasicInfo)`

SetData sets Data field to given value.

### HasData

`func (o *ModelUpgradeSiteInfo) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalRows

`func (o *ModelUpgradeSiteInfo) GetTotalRows() int32`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *ModelUpgradeSiteInfo) GetTotalRowsOk() (*int32, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *ModelUpgradeSiteInfo) SetTotalRows(v int32)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *ModelUpgradeSiteInfo) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


