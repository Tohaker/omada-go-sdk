# ExportLogOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **int32** | Parameter [format] indicates the type of the exported file. 0: CSV, 1: xlsx. | 
**SiteIds** | **[]string** | Parameter [siteIds] indicates site IDs that need to export log results. | 

## Methods

### NewExportLogOpenApiVO

`func NewExportLogOpenApiVO(format int32, siteIds []string, ) *ExportLogOpenApiVO`

NewExportLogOpenApiVO instantiates a new ExportLogOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportLogOpenApiVOWithDefaults

`func NewExportLogOpenApiVOWithDefaults() *ExportLogOpenApiVO`

NewExportLogOpenApiVOWithDefaults instantiates a new ExportLogOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *ExportLogOpenApiVO) GetFormat() int32`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ExportLogOpenApiVO) GetFormatOk() (*int32, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ExportLogOpenApiVO) SetFormat(v int32)`

SetFormat sets Format field to given value.


### GetSiteIds

`func (o *ExportLogOpenApiVO) GetSiteIds() []string`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *ExportLogOpenApiVO) GetSiteIdsOk() (*[]string, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *ExportLogOpenApiVO) SetSiteIds(v []string)`

SetSiteIds sets SiteIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


