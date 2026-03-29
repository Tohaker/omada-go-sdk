# ExportClientListOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientsDisplay** | Pointer to **[]string** | The information of client list for export. | [optional] 
**Format** | **int32** | Format of file exported.0: CSV 1: XLSX | 
**Mode** | **int32** | Export columns mode. 0：All Columns  1: CurrentDisplayColumns. | 
**QueryDataVO** | Pointer to [**OpenApiQueryDataVO**](OpenApiQueryDataVO.md) |  | [optional] 
**SiteIds** | **[]string** | IDs of the site of the data to be exported. | 

## Methods

### NewExportClientListOpenApiVO

`func NewExportClientListOpenApiVO(format int32, mode int32, siteIds []string, ) *ExportClientListOpenApiVO`

NewExportClientListOpenApiVO instantiates a new ExportClientListOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportClientListOpenApiVOWithDefaults

`func NewExportClientListOpenApiVOWithDefaults() *ExportClientListOpenApiVO`

NewExportClientListOpenApiVOWithDefaults instantiates a new ExportClientListOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientsDisplay

`func (o *ExportClientListOpenApiVO) GetClientsDisplay() []string`

GetClientsDisplay returns the ClientsDisplay field if non-nil, zero value otherwise.

### GetClientsDisplayOk

`func (o *ExportClientListOpenApiVO) GetClientsDisplayOk() (*[]string, bool)`

GetClientsDisplayOk returns a tuple with the ClientsDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsDisplay

`func (o *ExportClientListOpenApiVO) SetClientsDisplay(v []string)`

SetClientsDisplay sets ClientsDisplay field to given value.

### HasClientsDisplay

`func (o *ExportClientListOpenApiVO) HasClientsDisplay() bool`

HasClientsDisplay returns a boolean if a field has been set.

### GetFormat

`func (o *ExportClientListOpenApiVO) GetFormat() int32`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ExportClientListOpenApiVO) GetFormatOk() (*int32, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ExportClientListOpenApiVO) SetFormat(v int32)`

SetFormat sets Format field to given value.


### GetMode

`func (o *ExportClientListOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ExportClientListOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ExportClientListOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetQueryDataVO

`func (o *ExportClientListOpenApiVO) GetQueryDataVO() OpenApiQueryDataVO`

GetQueryDataVO returns the QueryDataVO field if non-nil, zero value otherwise.

### GetQueryDataVOOk

`func (o *ExportClientListOpenApiVO) GetQueryDataVOOk() (*OpenApiQueryDataVO, bool)`

GetQueryDataVOOk returns a tuple with the QueryDataVO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryDataVO

`func (o *ExportClientListOpenApiVO) SetQueryDataVO(v OpenApiQueryDataVO)`

SetQueryDataVO sets QueryDataVO field to given value.

### HasQueryDataVO

`func (o *ExportClientListOpenApiVO) HasQueryDataVO() bool`

HasQueryDataVO returns a boolean if a field has been set.

### GetSiteIds

`func (o *ExportClientListOpenApiVO) GetSiteIds() []string`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *ExportClientListOpenApiVO) GetSiteIdsOk() (*[]string, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *ExportClientListOpenApiVO) SetSiteIds(v []string)`

SetSiteIds sets SiteIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


