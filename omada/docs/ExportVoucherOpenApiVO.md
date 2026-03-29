# ExportVoucherOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **int32** | Export file format, should be a value as follows: 0: csv, 1: xlsx | 
**GroupIds** | Pointer to **[]string** | ID list of voucher groups. Voucher group can be created using &#39;Create Voucher Group&#39; interface, and Voucher Group ID can be obtained from &#39;Get Voucher Group list&#39; interface | [optional] 
**QueryData** | Pointer to [**OpenApiQueryDataVO**](OpenApiQueryDataVO.md) |  | [optional] 
**Type** | **int32** | Select type. It should be a value as follows: 0: Represents selecting all voucher groups, this selection does not pass parameter [groupIds]. 1: Parameter [groupIds] includes the IDs of the voucher groups to be selected. 2: Parameter [groupIds] includes the IDs of the voucher groups not to be selected | 

## Methods

### NewExportVoucherOpenApiVO

`func NewExportVoucherOpenApiVO(format int32, type_ int32, ) *ExportVoucherOpenApiVO`

NewExportVoucherOpenApiVO instantiates a new ExportVoucherOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportVoucherOpenApiVOWithDefaults

`func NewExportVoucherOpenApiVOWithDefaults() *ExportVoucherOpenApiVO`

NewExportVoucherOpenApiVOWithDefaults instantiates a new ExportVoucherOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *ExportVoucherOpenApiVO) GetFormat() int32`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ExportVoucherOpenApiVO) GetFormatOk() (*int32, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ExportVoucherOpenApiVO) SetFormat(v int32)`

SetFormat sets Format field to given value.


### GetGroupIds

`func (o *ExportVoucherOpenApiVO) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *ExportVoucherOpenApiVO) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *ExportVoucherOpenApiVO) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *ExportVoucherOpenApiVO) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetQueryData

`func (o *ExportVoucherOpenApiVO) GetQueryData() OpenApiQueryDataVO`

GetQueryData returns the QueryData field if non-nil, zero value otherwise.

### GetQueryDataOk

`func (o *ExportVoucherOpenApiVO) GetQueryDataOk() (*OpenApiQueryDataVO, bool)`

GetQueryDataOk returns a tuple with the QueryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryData

`func (o *ExportVoucherOpenApiVO) SetQueryData(v OpenApiQueryDataVO)`

SetQueryData sets QueryData field to given value.

### HasQueryData

`func (o *ExportVoucherOpenApiVO) HasQueryData() bool`

HasQueryData returns a boolean if a field has been set.

### GetType

`func (o *ExportVoucherOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExportVoucherOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExportVoucherOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


