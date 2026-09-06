# ClientHealthTrendVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Average** | Pointer to **int32** | Fair health count | [optional] 
**Good** | Pointer to **int32** | Good health count | [optional] 
**NoData** | Pointer to **int32** | No data count | [optional] 
**Poor** | Pointer to **int32** | Poor health count | [optional] 
**Time** | Pointer to **int64** | Timestamp | [optional] 
**Total** | Pointer to **int32** | Total client count | [optional] 
**Wired** | Pointer to [**ClientHealthCategoryVO**](ClientHealthCategoryVO.md) |  | [optional] 
**Wireless** | Pointer to [**ClientHealthCategoryVO**](ClientHealthCategoryVO.md) |  | [optional] 

## Methods

### NewClientHealthTrendVO

`func NewClientHealthTrendVO() *ClientHealthTrendVO`

NewClientHealthTrendVO instantiates a new ClientHealthTrendVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientHealthTrendVOWithDefaults

`func NewClientHealthTrendVOWithDefaults() *ClientHealthTrendVO`

NewClientHealthTrendVOWithDefaults instantiates a new ClientHealthTrendVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverage

`func (o *ClientHealthTrendVO) GetAverage() int32`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *ClientHealthTrendVO) GetAverageOk() (*int32, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *ClientHealthTrendVO) SetAverage(v int32)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *ClientHealthTrendVO) HasAverage() bool`

HasAverage returns a boolean if a field has been set.

### GetGood

`func (o *ClientHealthTrendVO) GetGood() int32`

GetGood returns the Good field if non-nil, zero value otherwise.

### GetGoodOk

`func (o *ClientHealthTrendVO) GetGoodOk() (*int32, bool)`

GetGoodOk returns a tuple with the Good field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGood

`func (o *ClientHealthTrendVO) SetGood(v int32)`

SetGood sets Good field to given value.

### HasGood

`func (o *ClientHealthTrendVO) HasGood() bool`

HasGood returns a boolean if a field has been set.

### GetNoData

`func (o *ClientHealthTrendVO) GetNoData() int32`

GetNoData returns the NoData field if non-nil, zero value otherwise.

### GetNoDataOk

`func (o *ClientHealthTrendVO) GetNoDataOk() (*int32, bool)`

GetNoDataOk returns a tuple with the NoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoData

`func (o *ClientHealthTrendVO) SetNoData(v int32)`

SetNoData sets NoData field to given value.

### HasNoData

`func (o *ClientHealthTrendVO) HasNoData() bool`

HasNoData returns a boolean if a field has been set.

### GetPoor

`func (o *ClientHealthTrendVO) GetPoor() int32`

GetPoor returns the Poor field if non-nil, zero value otherwise.

### GetPoorOk

`func (o *ClientHealthTrendVO) GetPoorOk() (*int32, bool)`

GetPoorOk returns a tuple with the Poor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoor

`func (o *ClientHealthTrendVO) SetPoor(v int32)`

SetPoor sets Poor field to given value.

### HasPoor

`func (o *ClientHealthTrendVO) HasPoor() bool`

HasPoor returns a boolean if a field has been set.

### GetTime

`func (o *ClientHealthTrendVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ClientHealthTrendVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ClientHealthTrendVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *ClientHealthTrendVO) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTotal

`func (o *ClientHealthTrendVO) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ClientHealthTrendVO) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ClientHealthTrendVO) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ClientHealthTrendVO) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetWired

`func (o *ClientHealthTrendVO) GetWired() ClientHealthCategoryVO`

GetWired returns the Wired field if non-nil, zero value otherwise.

### GetWiredOk

`func (o *ClientHealthTrendVO) GetWiredOk() (*ClientHealthCategoryVO, bool)`

GetWiredOk returns a tuple with the Wired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWired

`func (o *ClientHealthTrendVO) SetWired(v ClientHealthCategoryVO)`

SetWired sets Wired field to given value.

### HasWired

`func (o *ClientHealthTrendVO) HasWired() bool`

HasWired returns a boolean if a field has been set.

### GetWireless

`func (o *ClientHealthTrendVO) GetWireless() ClientHealthCategoryVO`

GetWireless returns the Wireless field if non-nil, zero value otherwise.

### GetWirelessOk

`func (o *ClientHealthTrendVO) GetWirelessOk() (*ClientHealthCategoryVO, bool)`

GetWirelessOk returns a tuple with the Wireless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireless

`func (o *ClientHealthTrendVO) SetWireless(v ClientHealthCategoryVO)`

SetWireless sets Wireless field to given value.

### HasWireless

`func (o *ClientHealthTrendVO) HasWireless() bool`

HasWireless returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


