# ClientExportOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **int32** | Format should be a value as follows: 0: csv; 1: xlsx. | 
**Tables** | **[]int32** | Tables should be a list of follows: 0: online; 1: offline; 2: blocked; 3: past-connection, 4: past-portal-auth. | 

## Methods

### NewClientExportOpenApiVO

`func NewClientExportOpenApiVO(format int32, tables []int32, ) *ClientExportOpenApiVO`

NewClientExportOpenApiVO instantiates a new ClientExportOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientExportOpenApiVOWithDefaults

`func NewClientExportOpenApiVOWithDefaults() *ClientExportOpenApiVO`

NewClientExportOpenApiVOWithDefaults instantiates a new ClientExportOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *ClientExportOpenApiVO) GetFormat() int32`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ClientExportOpenApiVO) GetFormatOk() (*int32, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ClientExportOpenApiVO) SetFormat(v int32)`

SetFormat sets Format field to given value.


### GetTables

`func (o *ClientExportOpenApiVO) GetTables() []int32`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *ClientExportOpenApiVO) GetTablesOk() (*[]int32, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *ClientExportOpenApiVO) SetTables(v []int32)`

SetTables sets Tables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


