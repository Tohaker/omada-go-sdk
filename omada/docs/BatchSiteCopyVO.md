# BatchSiteCopyVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteNamePrefix** | **string** | Site name prefix, new sites name will use this prefix. Parameter [siteNamePrefix] should be 1 - 128 ASCII characters. | 
**SourceSiteId** | **string** | Source site ID to be copied. | 
**TargetSiteNum** | **int32** | Site num to be created should between 1 and 300. | 

## Methods

### NewBatchSiteCopyVO

`func NewBatchSiteCopyVO(siteNamePrefix string, sourceSiteId string, targetSiteNum int32, ) *BatchSiteCopyVO`

NewBatchSiteCopyVO instantiates a new BatchSiteCopyVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchSiteCopyVOWithDefaults

`func NewBatchSiteCopyVOWithDefaults() *BatchSiteCopyVO`

NewBatchSiteCopyVOWithDefaults instantiates a new BatchSiteCopyVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteNamePrefix

`func (o *BatchSiteCopyVO) GetSiteNamePrefix() string`

GetSiteNamePrefix returns the SiteNamePrefix field if non-nil, zero value otherwise.

### GetSiteNamePrefixOk

`func (o *BatchSiteCopyVO) GetSiteNamePrefixOk() (*string, bool)`

GetSiteNamePrefixOk returns a tuple with the SiteNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteNamePrefix

`func (o *BatchSiteCopyVO) SetSiteNamePrefix(v string)`

SetSiteNamePrefix sets SiteNamePrefix field to given value.


### GetSourceSiteId

`func (o *BatchSiteCopyVO) GetSourceSiteId() string`

GetSourceSiteId returns the SourceSiteId field if non-nil, zero value otherwise.

### GetSourceSiteIdOk

`func (o *BatchSiteCopyVO) GetSourceSiteIdOk() (*string, bool)`

GetSourceSiteIdOk returns a tuple with the SourceSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSiteId

`func (o *BatchSiteCopyVO) SetSourceSiteId(v string)`

SetSourceSiteId sets SourceSiteId field to given value.


### GetTargetSiteNum

`func (o *BatchSiteCopyVO) GetTargetSiteNum() int32`

GetTargetSiteNum returns the TargetSiteNum field if non-nil, zero value otherwise.

### GetTargetSiteNumOk

`func (o *BatchSiteCopyVO) GetTargetSiteNumOk() (*int32, bool)`

GetTargetSiteNumOk returns a tuple with the TargetSiteNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSiteNum

`func (o *BatchSiteCopyVO) SetTargetSiteNum(v int32)`

SetTargetSiteNum sets TargetSiteNum field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


