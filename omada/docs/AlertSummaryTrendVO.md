# AlertSummaryTrendVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertSummary** | Pointer to [**AlertSummaryVO**](AlertSummaryVO.md) |  | [optional] 
**AlertTrend** | Pointer to [**[]TrendBaseVO**](TrendBaseVO.md) | Alert trend | [optional] 

## Methods

### NewAlertSummaryTrendVO

`func NewAlertSummaryTrendVO() *AlertSummaryTrendVO`

NewAlertSummaryTrendVO instantiates a new AlertSummaryTrendVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertSummaryTrendVOWithDefaults

`func NewAlertSummaryTrendVOWithDefaults() *AlertSummaryTrendVO`

NewAlertSummaryTrendVOWithDefaults instantiates a new AlertSummaryTrendVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertSummary

`func (o *AlertSummaryTrendVO) GetAlertSummary() AlertSummaryVO`

GetAlertSummary returns the AlertSummary field if non-nil, zero value otherwise.

### GetAlertSummaryOk

`func (o *AlertSummaryTrendVO) GetAlertSummaryOk() (*AlertSummaryVO, bool)`

GetAlertSummaryOk returns a tuple with the AlertSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSummary

`func (o *AlertSummaryTrendVO) SetAlertSummary(v AlertSummaryVO)`

SetAlertSummary sets AlertSummary field to given value.

### HasAlertSummary

`func (o *AlertSummaryTrendVO) HasAlertSummary() bool`

HasAlertSummary returns a boolean if a field has been set.

### GetAlertTrend

`func (o *AlertSummaryTrendVO) GetAlertTrend() []TrendBaseVO`

GetAlertTrend returns the AlertTrend field if non-nil, zero value otherwise.

### GetAlertTrendOk

`func (o *AlertSummaryTrendVO) GetAlertTrendOk() (*[]TrendBaseVO, bool)`

GetAlertTrendOk returns a tuple with the AlertTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTrend

`func (o *AlertSummaryTrendVO) SetAlertTrend(v []TrendBaseVO)`

SetAlertTrend sets AlertTrend field to given value.

### HasAlertTrend

`func (o *AlertSummaryTrendVO) HasAlertTrend() bool`

HasAlertTrend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


