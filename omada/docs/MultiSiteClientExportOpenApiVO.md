# MultiSiteClientExportOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientsDisplay** | Pointer to **[]string** | The information of client list for export. | [optional] 
**ClientsDisplayOverride** | Pointer to **map[string][]string** | Override the export columns for given siteIds. | [optional] 
**Format** | **int32** | Format should be a value as follows: 0: csv; 1: xlsx. | 
**Mode** | **int32** | Export columns mode. 0：All Columns  1: CurrentDisplayColumns. | 
**SelectType** | **string** | Select type of the sites of clients to export. include: include selected sites, exclude: all but exclude selected sites, all: include all sites. | 
**SiteIds** | **[]string** | List of site id to export. SiteIds should contains at least 1 element when [selectType] is \&quot;include\&quot;. | 

## Methods

### NewMultiSiteClientExportOpenApiVO

`func NewMultiSiteClientExportOpenApiVO(format int32, mode int32, selectType string, siteIds []string, ) *MultiSiteClientExportOpenApiVO`

NewMultiSiteClientExportOpenApiVO instantiates a new MultiSiteClientExportOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiSiteClientExportOpenApiVOWithDefaults

`func NewMultiSiteClientExportOpenApiVOWithDefaults() *MultiSiteClientExportOpenApiVO`

NewMultiSiteClientExportOpenApiVOWithDefaults instantiates a new MultiSiteClientExportOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientsDisplay

`func (o *MultiSiteClientExportOpenApiVO) GetClientsDisplay() []string`

GetClientsDisplay returns the ClientsDisplay field if non-nil, zero value otherwise.

### GetClientsDisplayOk

`func (o *MultiSiteClientExportOpenApiVO) GetClientsDisplayOk() (*[]string, bool)`

GetClientsDisplayOk returns a tuple with the ClientsDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsDisplay

`func (o *MultiSiteClientExportOpenApiVO) SetClientsDisplay(v []string)`

SetClientsDisplay sets ClientsDisplay field to given value.

### HasClientsDisplay

`func (o *MultiSiteClientExportOpenApiVO) HasClientsDisplay() bool`

HasClientsDisplay returns a boolean if a field has been set.

### GetClientsDisplayOverride

`func (o *MultiSiteClientExportOpenApiVO) GetClientsDisplayOverride() map[string][]string`

GetClientsDisplayOverride returns the ClientsDisplayOverride field if non-nil, zero value otherwise.

### GetClientsDisplayOverrideOk

`func (o *MultiSiteClientExportOpenApiVO) GetClientsDisplayOverrideOk() (*map[string][]string, bool)`

GetClientsDisplayOverrideOk returns a tuple with the ClientsDisplayOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsDisplayOverride

`func (o *MultiSiteClientExportOpenApiVO) SetClientsDisplayOverride(v map[string][]string)`

SetClientsDisplayOverride sets ClientsDisplayOverride field to given value.

### HasClientsDisplayOverride

`func (o *MultiSiteClientExportOpenApiVO) HasClientsDisplayOverride() bool`

HasClientsDisplayOverride returns a boolean if a field has been set.

### GetFormat

`func (o *MultiSiteClientExportOpenApiVO) GetFormat() int32`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MultiSiteClientExportOpenApiVO) GetFormatOk() (*int32, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MultiSiteClientExportOpenApiVO) SetFormat(v int32)`

SetFormat sets Format field to given value.


### GetMode

`func (o *MultiSiteClientExportOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *MultiSiteClientExportOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *MultiSiteClientExportOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetSelectType

`func (o *MultiSiteClientExportOpenApiVO) GetSelectType() string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *MultiSiteClientExportOpenApiVO) GetSelectTypeOk() (*string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *MultiSiteClientExportOpenApiVO) SetSelectType(v string)`

SetSelectType sets SelectType field to given value.


### GetSiteIds

`func (o *MultiSiteClientExportOpenApiVO) GetSiteIds() []string`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *MultiSiteClientExportOpenApiVO) GetSiteIdsOk() (*[]string, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *MultiSiteClientExportOpenApiVO) SetSiteIds(v []string)`

SetSiteIds sets SiteIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


