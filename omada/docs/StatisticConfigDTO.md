# StatisticConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRefresh** | Pointer to **string** | AutoRefresh should be a value as follows:DISABLE,ENABLE | [optional] 
**RefreshInterval** | Pointer to **int32** | RefreshInterval should be within the range of 3 to 300 | [optional] 

## Methods

### NewStatisticConfigDTO

`func NewStatisticConfigDTO() *StatisticConfigDTO`

NewStatisticConfigDTO instantiates a new StatisticConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticConfigDTOWithDefaults

`func NewStatisticConfigDTOWithDefaults() *StatisticConfigDTO`

NewStatisticConfigDTOWithDefaults instantiates a new StatisticConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRefresh

`func (o *StatisticConfigDTO) GetAutoRefresh() string`

GetAutoRefresh returns the AutoRefresh field if non-nil, zero value otherwise.

### GetAutoRefreshOk

`func (o *StatisticConfigDTO) GetAutoRefreshOk() (*string, bool)`

GetAutoRefreshOk returns a tuple with the AutoRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRefresh

`func (o *StatisticConfigDTO) SetAutoRefresh(v string)`

SetAutoRefresh sets AutoRefresh field to given value.

### HasAutoRefresh

`func (o *StatisticConfigDTO) HasAutoRefresh() bool`

HasAutoRefresh returns a boolean if a field has been set.

### GetRefreshInterval

`func (o *StatisticConfigDTO) GetRefreshInterval() int32`

GetRefreshInterval returns the RefreshInterval field if non-nil, zero value otherwise.

### GetRefreshIntervalOk

`func (o *StatisticConfigDTO) GetRefreshIntervalOk() (*int32, bool)`

GetRefreshIntervalOk returns a tuple with the RefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshInterval

`func (o *StatisticConfigDTO) SetRefreshInterval(v int32)`

SetRefreshInterval sets RefreshInterval field to given value.

### HasRefreshInterval

`func (o *StatisticConfigDTO) HasRefreshInterval() bool`

HasRefreshInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


