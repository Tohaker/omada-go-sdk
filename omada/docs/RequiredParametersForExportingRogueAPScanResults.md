# RequiredParametersForExportingRogueAPScanResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **int32** | Parameter [format] indicates the type of the exported file. 0: CSV, 1: xlsx. | [optional] 
**SiteIds** | Pointer to **[]string** | Parameter [siteIds] indicates site IDs that need to export rogue AP scan results. | [optional] 

## Methods

### NewRequiredParametersForExportingRogueAPScanResults

`func NewRequiredParametersForExportingRogueAPScanResults() *RequiredParametersForExportingRogueAPScanResults`

NewRequiredParametersForExportingRogueAPScanResults instantiates a new RequiredParametersForExportingRogueAPScanResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequiredParametersForExportingRogueAPScanResultsWithDefaults

`func NewRequiredParametersForExportingRogueAPScanResultsWithDefaults() *RequiredParametersForExportingRogueAPScanResults`

NewRequiredParametersForExportingRogueAPScanResultsWithDefaults instantiates a new RequiredParametersForExportingRogueAPScanResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *RequiredParametersForExportingRogueAPScanResults) GetFormat() int32`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *RequiredParametersForExportingRogueAPScanResults) GetFormatOk() (*int32, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *RequiredParametersForExportingRogueAPScanResults) SetFormat(v int32)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *RequiredParametersForExportingRogueAPScanResults) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetSiteIds

`func (o *RequiredParametersForExportingRogueAPScanResults) GetSiteIds() []string`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *RequiredParametersForExportingRogueAPScanResults) GetSiteIdsOk() (*[]string, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *RequiredParametersForExportingRogueAPScanResults) SetSiteIds(v []string)`

SetSiteIds sets SiteIds field to given value.

### HasSiteIds

`func (o *RequiredParametersForExportingRogueAPScanResults) HasSiteIds() bool`

HasSiteIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


