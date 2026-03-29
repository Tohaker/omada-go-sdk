# BatchSiteSelfRestoreVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteRestoreInfos** | [**[]SelfSiteRestoreVO**](SelfSiteRestoreVO.md) | Site restore info list to restore. Up to 300 entries are allowed for the site restore info list. | 

## Methods

### NewBatchSiteSelfRestoreVO

`func NewBatchSiteSelfRestoreVO(siteRestoreInfos []SelfSiteRestoreVO, ) *BatchSiteSelfRestoreVO`

NewBatchSiteSelfRestoreVO instantiates a new BatchSiteSelfRestoreVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchSiteSelfRestoreVOWithDefaults

`func NewBatchSiteSelfRestoreVOWithDefaults() *BatchSiteSelfRestoreVO`

NewBatchSiteSelfRestoreVOWithDefaults instantiates a new BatchSiteSelfRestoreVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteRestoreInfos

`func (o *BatchSiteSelfRestoreVO) GetSiteRestoreInfos() []SelfSiteRestoreVO`

GetSiteRestoreInfos returns the SiteRestoreInfos field if non-nil, zero value otherwise.

### GetSiteRestoreInfosOk

`func (o *BatchSiteSelfRestoreVO) GetSiteRestoreInfosOk() (*[]SelfSiteRestoreVO, bool)`

GetSiteRestoreInfosOk returns a tuple with the SiteRestoreInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteRestoreInfos

`func (o *BatchSiteSelfRestoreVO) SetSiteRestoreInfos(v []SelfSiteRestoreVO)`

SetSiteRestoreInfos sets SiteRestoreInfos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


