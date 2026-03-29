# BatchSiteBackupVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteIds** | **[]string** | Site ID list to backup, at least one site, max 300 sites. | 

## Methods

### NewBatchSiteBackupVO

`func NewBatchSiteBackupVO(siteIds []string, ) *BatchSiteBackupVO`

NewBatchSiteBackupVO instantiates a new BatchSiteBackupVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchSiteBackupVOWithDefaults

`func NewBatchSiteBackupVOWithDefaults() *BatchSiteBackupVO`

NewBatchSiteBackupVOWithDefaults instantiates a new BatchSiteBackupVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteIds

`func (o *BatchSiteBackupVO) GetSiteIds() []string`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *BatchSiteBackupVO) GetSiteIdsOk() (*[]string, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *BatchSiteBackupVO) SetSiteIds(v []string)`

SetSiteIds sets SiteIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


