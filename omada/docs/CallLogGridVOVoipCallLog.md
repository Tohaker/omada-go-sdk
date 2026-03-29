# CallLogGridVOVoipCallLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]VoipCallLog**](VoipCallLog.md) |  | [optional] 
**Statistic** | Pointer to [**CallLogStatisticVO**](CallLogStatisticVO.md) |  | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewCallLogGridVOVoipCallLog

`func NewCallLogGridVOVoipCallLog() *CallLogGridVOVoipCallLog`

NewCallLogGridVOVoipCallLog instantiates a new CallLogGridVOVoipCallLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallLogGridVOVoipCallLogWithDefaults

`func NewCallLogGridVOVoipCallLogWithDefaults() *CallLogGridVOVoipCallLog`

NewCallLogGridVOVoipCallLogWithDefaults instantiates a new CallLogGridVOVoipCallLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *CallLogGridVOVoipCallLog) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *CallLogGridVOVoipCallLog) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *CallLogGridVOVoipCallLog) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *CallLogGridVOVoipCallLog) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *CallLogGridVOVoipCallLog) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *CallLogGridVOVoipCallLog) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *CallLogGridVOVoipCallLog) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *CallLogGridVOVoipCallLog) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *CallLogGridVOVoipCallLog) GetData() []VoipCallLog`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CallLogGridVOVoipCallLog) GetDataOk() (*[]VoipCallLog, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CallLogGridVOVoipCallLog) SetData(v []VoipCallLog)`

SetData sets Data field to given value.

### HasData

`func (o *CallLogGridVOVoipCallLog) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatistic

`func (o *CallLogGridVOVoipCallLog) GetStatistic() CallLogStatisticVO`

GetStatistic returns the Statistic field if non-nil, zero value otherwise.

### GetStatisticOk

`func (o *CallLogGridVOVoipCallLog) GetStatisticOk() (*CallLogStatisticVO, bool)`

GetStatisticOk returns a tuple with the Statistic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistic

`func (o *CallLogGridVOVoipCallLog) SetStatistic(v CallLogStatisticVO)`

SetStatistic sets Statistic field to given value.

### HasStatistic

`func (o *CallLogGridVOVoipCallLog) HasStatistic() bool`

HasStatistic returns a boolean if a field has been set.

### GetTotalRows

`func (o *CallLogGridVOVoipCallLog) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *CallLogGridVOVoipCallLog) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *CallLogGridVOVoipCallLog) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *CallLogGridVOVoipCallLog) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


