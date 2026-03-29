# ClientConnectionTrend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientConnectionSummary** | Pointer to [**ClientConnectionSummary**](ClientConnectionSummary.md) |  | [optional] 
**ClientConnectionTrend** | Pointer to [**[]ClientCountStatisticsWithTime**](ClientCountStatisticsWithTime.md) | Client connection trend with time. | [optional] 

## Methods

### NewClientConnectionTrend

`func NewClientConnectionTrend() *ClientConnectionTrend`

NewClientConnectionTrend instantiates a new ClientConnectionTrend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientConnectionTrendWithDefaults

`func NewClientConnectionTrendWithDefaults() *ClientConnectionTrend`

NewClientConnectionTrendWithDefaults instantiates a new ClientConnectionTrend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientConnectionSummary

`func (o *ClientConnectionTrend) GetClientConnectionSummary() ClientConnectionSummary`

GetClientConnectionSummary returns the ClientConnectionSummary field if non-nil, zero value otherwise.

### GetClientConnectionSummaryOk

`func (o *ClientConnectionTrend) GetClientConnectionSummaryOk() (*ClientConnectionSummary, bool)`

GetClientConnectionSummaryOk returns a tuple with the ClientConnectionSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnectionSummary

`func (o *ClientConnectionTrend) SetClientConnectionSummary(v ClientConnectionSummary)`

SetClientConnectionSummary sets ClientConnectionSummary field to given value.

### HasClientConnectionSummary

`func (o *ClientConnectionTrend) HasClientConnectionSummary() bool`

HasClientConnectionSummary returns a boolean if a field has been set.

### GetClientConnectionTrend

`func (o *ClientConnectionTrend) GetClientConnectionTrend() []ClientCountStatisticsWithTime`

GetClientConnectionTrend returns the ClientConnectionTrend field if non-nil, zero value otherwise.

### GetClientConnectionTrendOk

`func (o *ClientConnectionTrend) GetClientConnectionTrendOk() (*[]ClientCountStatisticsWithTime, bool)`

GetClientConnectionTrendOk returns a tuple with the ClientConnectionTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnectionTrend

`func (o *ClientConnectionTrend) SetClientConnectionTrend(v []ClientCountStatisticsWithTime)`

SetClientConnectionTrend sets ClientConnectionTrend field to given value.

### HasClientConnectionTrend

`func (o *ClientConnectionTrend) HasClientConnectionTrend() bool`

HasClientConnectionTrend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


