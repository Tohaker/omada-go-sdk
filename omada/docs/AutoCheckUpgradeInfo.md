# AutoCheckUpgradeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoCheckTime** | Pointer to **string** | Next execution time | [optional] 
**Channel** | Pointer to **int32** | Channel should be a value as follows: 0: stable; 1: Release Candidate(RC); 2: Beta | [optional] 
**Id** | Pointer to **string** | ID | [optional] 
**ModelTypeInfos** | Pointer to [**[]ModelTypeInfoOpenApiVO**](ModelTypeInfoOpenApiVO.md) | List of model type selected by the user | [optional] 
**Occurrence** | Pointer to [**UpgradeBaseScheduleTimeOpenApiVO**](UpgradeBaseScheduleTimeOpenApiVO.md) |  | [optional] 
**SiteNames** | Pointer to **[]string** | The siteNames lists selected by the user | [optional] 
**SiteNum** | Pointer to **int32** | Number of sites selected by the user | [optional] 

## Methods

### NewAutoCheckUpgradeInfo

`func NewAutoCheckUpgradeInfo() *AutoCheckUpgradeInfo`

NewAutoCheckUpgradeInfo instantiates a new AutoCheckUpgradeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoCheckUpgradeInfoWithDefaults

`func NewAutoCheckUpgradeInfoWithDefaults() *AutoCheckUpgradeInfo`

NewAutoCheckUpgradeInfoWithDefaults instantiates a new AutoCheckUpgradeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoCheckTime

`func (o *AutoCheckUpgradeInfo) GetAutoCheckTime() string`

GetAutoCheckTime returns the AutoCheckTime field if non-nil, zero value otherwise.

### GetAutoCheckTimeOk

`func (o *AutoCheckUpgradeInfo) GetAutoCheckTimeOk() (*string, bool)`

GetAutoCheckTimeOk returns a tuple with the AutoCheckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCheckTime

`func (o *AutoCheckUpgradeInfo) SetAutoCheckTime(v string)`

SetAutoCheckTime sets AutoCheckTime field to given value.

### HasAutoCheckTime

`func (o *AutoCheckUpgradeInfo) HasAutoCheckTime() bool`

HasAutoCheckTime returns a boolean if a field has been set.

### GetChannel

`func (o *AutoCheckUpgradeInfo) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *AutoCheckUpgradeInfo) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *AutoCheckUpgradeInfo) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *AutoCheckUpgradeInfo) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetId

`func (o *AutoCheckUpgradeInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoCheckUpgradeInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoCheckUpgradeInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutoCheckUpgradeInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModelTypeInfos

`func (o *AutoCheckUpgradeInfo) GetModelTypeInfos() []ModelTypeInfoOpenApiVO`

GetModelTypeInfos returns the ModelTypeInfos field if non-nil, zero value otherwise.

### GetModelTypeInfosOk

`func (o *AutoCheckUpgradeInfo) GetModelTypeInfosOk() (*[]ModelTypeInfoOpenApiVO, bool)`

GetModelTypeInfosOk returns a tuple with the ModelTypeInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeInfos

`func (o *AutoCheckUpgradeInfo) SetModelTypeInfos(v []ModelTypeInfoOpenApiVO)`

SetModelTypeInfos sets ModelTypeInfos field to given value.

### HasModelTypeInfos

`func (o *AutoCheckUpgradeInfo) HasModelTypeInfos() bool`

HasModelTypeInfos returns a boolean if a field has been set.

### GetOccurrence

`func (o *AutoCheckUpgradeInfo) GetOccurrence() UpgradeBaseScheduleTimeOpenApiVO`

GetOccurrence returns the Occurrence field if non-nil, zero value otherwise.

### GetOccurrenceOk

`func (o *AutoCheckUpgradeInfo) GetOccurrenceOk() (*UpgradeBaseScheduleTimeOpenApiVO, bool)`

GetOccurrenceOk returns a tuple with the Occurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrence

`func (o *AutoCheckUpgradeInfo) SetOccurrence(v UpgradeBaseScheduleTimeOpenApiVO)`

SetOccurrence sets Occurrence field to given value.

### HasOccurrence

`func (o *AutoCheckUpgradeInfo) HasOccurrence() bool`

HasOccurrence returns a boolean if a field has been set.

### GetSiteNames

`func (o *AutoCheckUpgradeInfo) GetSiteNames() []string`

GetSiteNames returns the SiteNames field if non-nil, zero value otherwise.

### GetSiteNamesOk

`func (o *AutoCheckUpgradeInfo) GetSiteNamesOk() (*[]string, bool)`

GetSiteNamesOk returns a tuple with the SiteNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteNames

`func (o *AutoCheckUpgradeInfo) SetSiteNames(v []string)`

SetSiteNames sets SiteNames field to given value.

### HasSiteNames

`func (o *AutoCheckUpgradeInfo) HasSiteNames() bool`

HasSiteNames returns a boolean if a field has been set.

### GetSiteNum

`func (o *AutoCheckUpgradeInfo) GetSiteNum() int32`

GetSiteNum returns the SiteNum field if non-nil, zero value otherwise.

### GetSiteNumOk

`func (o *AutoCheckUpgradeInfo) GetSiteNumOk() (*int32, bool)`

GetSiteNumOk returns a tuple with the SiteNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteNum

`func (o *AutoCheckUpgradeInfo) SetSiteNum(v int32)`

SetSiteNum sets SiteNum field to given value.

### HasSiteNum

`func (o *AutoCheckUpgradeInfo) HasSiteNum() bool`

HasSiteNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


