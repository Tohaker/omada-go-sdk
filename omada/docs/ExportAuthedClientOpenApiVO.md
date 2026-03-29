# ExportAuthedClientOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **int32** | Export authed client format. | [optional] 
**SiteIds** | Pointer to **[]string** | Site IDs of the authed clients to export. | [optional] 

## Methods

### NewExportAuthedClientOpenApiVO

`func NewExportAuthedClientOpenApiVO() *ExportAuthedClientOpenApiVO`

NewExportAuthedClientOpenApiVO instantiates a new ExportAuthedClientOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportAuthedClientOpenApiVOWithDefaults

`func NewExportAuthedClientOpenApiVOWithDefaults() *ExportAuthedClientOpenApiVO`

NewExportAuthedClientOpenApiVOWithDefaults instantiates a new ExportAuthedClientOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *ExportAuthedClientOpenApiVO) GetFormat() int32`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ExportAuthedClientOpenApiVO) GetFormatOk() (*int32, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ExportAuthedClientOpenApiVO) SetFormat(v int32)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ExportAuthedClientOpenApiVO) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetSiteIds

`func (o *ExportAuthedClientOpenApiVO) GetSiteIds() []string`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *ExportAuthedClientOpenApiVO) GetSiteIdsOk() (*[]string, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *ExportAuthedClientOpenApiVO) SetSiteIds(v []string)`

SetSiteIds sets SiteIds field to given value.

### HasSiteIds

`func (o *ExportAuthedClientOpenApiVO) HasSiteIds() bool`

HasSiteIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


