# VoucherStatisticsHistoryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | Pointer to [**VoucherSummaryOpenApiVO**](VoucherSummaryOpenApiVO.md) |  | [optional] 
**Usage** | Pointer to [**[]VoucherUsageOpenApiVO**](VoucherUsageOpenApiVO.md) | Data points of vouchers | [optional] 

## Methods

### NewVoucherStatisticsHistoryOpenApiVO

`func NewVoucherStatisticsHistoryOpenApiVO() *VoucherStatisticsHistoryOpenApiVO`

NewVoucherStatisticsHistoryOpenApiVO instantiates a new VoucherStatisticsHistoryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoucherStatisticsHistoryOpenApiVOWithDefaults

`func NewVoucherStatisticsHistoryOpenApiVOWithDefaults() *VoucherStatisticsHistoryOpenApiVO`

NewVoucherStatisticsHistoryOpenApiVOWithDefaults instantiates a new VoucherStatisticsHistoryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *VoucherStatisticsHistoryOpenApiVO) GetSummary() VoucherSummaryOpenApiVO`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *VoucherStatisticsHistoryOpenApiVO) GetSummaryOk() (*VoucherSummaryOpenApiVO, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *VoucherStatisticsHistoryOpenApiVO) SetSummary(v VoucherSummaryOpenApiVO)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *VoucherStatisticsHistoryOpenApiVO) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUsage

`func (o *VoucherStatisticsHistoryOpenApiVO) GetUsage() []VoucherUsageOpenApiVO`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *VoucherStatisticsHistoryOpenApiVO) GetUsageOk() (*[]VoucherUsageOpenApiVO, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *VoucherStatisticsHistoryOpenApiVO) SetUsage(v []VoucherUsageOpenApiVO)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *VoucherStatisticsHistoryOpenApiVO) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


