# SelfSiteRestoreVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | **string** | Site backup file name. Parameter [fileName] should be 1 - 128 ASCII characters. | 
**SiteId** | **string** | Site ID to restore | 

## Methods

### NewSelfSiteRestoreVO

`func NewSelfSiteRestoreVO(fileName string, siteId string, ) *SelfSiteRestoreVO`

NewSelfSiteRestoreVO instantiates a new SelfSiteRestoreVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfSiteRestoreVOWithDefaults

`func NewSelfSiteRestoreVOWithDefaults() *SelfSiteRestoreVO`

NewSelfSiteRestoreVOWithDefaults instantiates a new SelfSiteRestoreVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *SelfSiteRestoreVO) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *SelfSiteRestoreVO) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *SelfSiteRestoreVO) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetSiteId

`func (o *SelfSiteRestoreVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SelfSiteRestoreVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SelfSiteRestoreVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


