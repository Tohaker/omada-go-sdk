# DuplicateSsidOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DuplicateSsidNames** | Pointer to **[]string** | The SSID Names that are duplicated by site | [optional] 
**SiteId** | Pointer to **string** | The site ID | [optional] 
**SsidNameMap** | Pointer to **map[string]int64** | The SSIDs that are duplicated by site | [optional] 

## Methods

### NewDuplicateSsidOpenApiVO

`func NewDuplicateSsidOpenApiVO() *DuplicateSsidOpenApiVO`

NewDuplicateSsidOpenApiVO instantiates a new DuplicateSsidOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDuplicateSsidOpenApiVOWithDefaults

`func NewDuplicateSsidOpenApiVOWithDefaults() *DuplicateSsidOpenApiVO`

NewDuplicateSsidOpenApiVOWithDefaults instantiates a new DuplicateSsidOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuplicateSsidNames

`func (o *DuplicateSsidOpenApiVO) GetDuplicateSsidNames() []string`

GetDuplicateSsidNames returns the DuplicateSsidNames field if non-nil, zero value otherwise.

### GetDuplicateSsidNamesOk

`func (o *DuplicateSsidOpenApiVO) GetDuplicateSsidNamesOk() (*[]string, bool)`

GetDuplicateSsidNamesOk returns a tuple with the DuplicateSsidNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateSsidNames

`func (o *DuplicateSsidOpenApiVO) SetDuplicateSsidNames(v []string)`

SetDuplicateSsidNames sets DuplicateSsidNames field to given value.

### HasDuplicateSsidNames

`func (o *DuplicateSsidOpenApiVO) HasDuplicateSsidNames() bool`

HasDuplicateSsidNames returns a boolean if a field has been set.

### GetSiteId

`func (o *DuplicateSsidOpenApiVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DuplicateSsidOpenApiVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DuplicateSsidOpenApiVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *DuplicateSsidOpenApiVO) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSsidNameMap

`func (o *DuplicateSsidOpenApiVO) GetSsidNameMap() map[string]int64`

GetSsidNameMap returns the SsidNameMap field if non-nil, zero value otherwise.

### GetSsidNameMapOk

`func (o *DuplicateSsidOpenApiVO) GetSsidNameMapOk() (*map[string]int64, bool)`

GetSsidNameMapOk returns a tuple with the SsidNameMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNameMap

`func (o *DuplicateSsidOpenApiVO) SetSsidNameMap(v map[string]int64)`

SetSsidNameMap sets SsidNameMap field to given value.

### HasSsidNameMap

`func (o *DuplicateSsidOpenApiVO) HasSsidNameMap() bool`

HasSsidNameMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


