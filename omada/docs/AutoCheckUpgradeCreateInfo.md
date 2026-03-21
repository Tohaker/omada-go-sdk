# AutoCheckUpgradeCreateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **int32** | Channel should be a value as follows: 0: stable; 1: Release Candidate(RC); 2: Beta, and it should not be null | 
**ModelTypeInfos** | [**[]ModelTypeInfoOpenApiVO**](ModelTypeInfoOpenApiVO.md) | List of model type selected by the user, and it should not be null | 
**Occurrence** | [**UpgradeBaseScheduleTimeOpenApiVO**](UpgradeBaseScheduleTimeOpenApiVO.md) |  | 
**SiteIds** | **[]string** | List of sites selected by the user, and it should not be null | 

## Methods

### NewAutoCheckUpgradeCreateInfo

`func NewAutoCheckUpgradeCreateInfo(channel int32, modelTypeInfos []ModelTypeInfoOpenApiVO, occurrence UpgradeBaseScheduleTimeOpenApiVO, siteIds []string, ) *AutoCheckUpgradeCreateInfo`

NewAutoCheckUpgradeCreateInfo instantiates a new AutoCheckUpgradeCreateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoCheckUpgradeCreateInfoWithDefaults

`func NewAutoCheckUpgradeCreateInfoWithDefaults() *AutoCheckUpgradeCreateInfo`

NewAutoCheckUpgradeCreateInfoWithDefaults instantiates a new AutoCheckUpgradeCreateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *AutoCheckUpgradeCreateInfo) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *AutoCheckUpgradeCreateInfo) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *AutoCheckUpgradeCreateInfo) SetChannel(v int32)`

SetChannel sets Channel field to given value.


### GetModelTypeInfos

`func (o *AutoCheckUpgradeCreateInfo) GetModelTypeInfos() []ModelTypeInfoOpenApiVO`

GetModelTypeInfos returns the ModelTypeInfos field if non-nil, zero value otherwise.

### GetModelTypeInfosOk

`func (o *AutoCheckUpgradeCreateInfo) GetModelTypeInfosOk() (*[]ModelTypeInfoOpenApiVO, bool)`

GetModelTypeInfosOk returns a tuple with the ModelTypeInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeInfos

`func (o *AutoCheckUpgradeCreateInfo) SetModelTypeInfos(v []ModelTypeInfoOpenApiVO)`

SetModelTypeInfos sets ModelTypeInfos field to given value.


### GetOccurrence

`func (o *AutoCheckUpgradeCreateInfo) GetOccurrence() UpgradeBaseScheduleTimeOpenApiVO`

GetOccurrence returns the Occurrence field if non-nil, zero value otherwise.

### GetOccurrenceOk

`func (o *AutoCheckUpgradeCreateInfo) GetOccurrenceOk() (*UpgradeBaseScheduleTimeOpenApiVO, bool)`

GetOccurrenceOk returns a tuple with the Occurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrence

`func (o *AutoCheckUpgradeCreateInfo) SetOccurrence(v UpgradeBaseScheduleTimeOpenApiVO)`

SetOccurrence sets Occurrence field to given value.


### GetSiteIds

`func (o *AutoCheckUpgradeCreateInfo) GetSiteIds() []string`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *AutoCheckUpgradeCreateInfo) GetSiteIdsOk() (*[]string, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *AutoCheckUpgradeCreateInfo) SetSiteIds(v []string)`

SetSiteIds sets SiteIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


