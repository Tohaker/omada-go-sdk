# SitesSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllSiteIds** | Pointer to **[]string** | all sites id. | [optional] 
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]Site**](Site.md) |  | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewSitesSite

`func NewSitesSite() *SitesSite`

NewSitesSite instantiates a new SitesSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSitesSiteWithDefaults

`func NewSitesSiteWithDefaults() *SitesSite`

NewSitesSiteWithDefaults instantiates a new SitesSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllSiteIds

`func (o *SitesSite) GetAllSiteIds() []string`

GetAllSiteIds returns the AllSiteIds field if non-nil, zero value otherwise.

### GetAllSiteIdsOk

`func (o *SitesSite) GetAllSiteIdsOk() (*[]string, bool)`

GetAllSiteIdsOk returns a tuple with the AllSiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSiteIds

`func (o *SitesSite) SetAllSiteIds(v []string)`

SetAllSiteIds sets AllSiteIds field to given value.

### HasAllSiteIds

`func (o *SitesSite) HasAllSiteIds() bool`

HasAllSiteIds returns a boolean if a field has been set.

### GetCurrentPage

`func (o *SitesSite) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *SitesSite) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *SitesSite) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *SitesSite) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *SitesSite) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *SitesSite) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *SitesSite) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *SitesSite) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *SitesSite) GetData() []Site`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SitesSite) GetDataOk() (*[]Site, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SitesSite) SetData(v []Site)`

SetData sets Data field to given value.

### HasData

`func (o *SitesSite) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalRows

`func (o *SitesSite) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *SitesSite) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *SitesSite) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *SitesSite) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


