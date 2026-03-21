# WifiDeviceAndClientQueryVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchKey** | Pointer to **string** | Fuzzy query parameters, support field name,mac | [optional] 
**Sorts** | Pointer to **map[string]string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. | [optional] 
**Time** | **int64** | Time(unit:ms) | 
**TopK** | Pointer to **int32** | The topK elements, which must be less than 100 | [optional] 

## Methods

### NewWifiDeviceAndClientQueryVO

`func NewWifiDeviceAndClientQueryVO(time int64, ) *WifiDeviceAndClientQueryVO`

NewWifiDeviceAndClientQueryVO instantiates a new WifiDeviceAndClientQueryVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWifiDeviceAndClientQueryVOWithDefaults

`func NewWifiDeviceAndClientQueryVOWithDefaults() *WifiDeviceAndClientQueryVO`

NewWifiDeviceAndClientQueryVOWithDefaults instantiates a new WifiDeviceAndClientQueryVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchKey

`func (o *WifiDeviceAndClientQueryVO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *WifiDeviceAndClientQueryVO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *WifiDeviceAndClientQueryVO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *WifiDeviceAndClientQueryVO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSorts

`func (o *WifiDeviceAndClientQueryVO) GetSorts() map[string]string`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *WifiDeviceAndClientQueryVO) GetSortsOk() (*map[string]string, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *WifiDeviceAndClientQueryVO) SetSorts(v map[string]string)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *WifiDeviceAndClientQueryVO) HasSorts() bool`

HasSorts returns a boolean if a field has been set.

### GetTime

`func (o *WifiDeviceAndClientQueryVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *WifiDeviceAndClientQueryVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *WifiDeviceAndClientQueryVO) SetTime(v int64)`

SetTime sets Time field to given value.


### GetTopK

`func (o *WifiDeviceAndClientQueryVO) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *WifiDeviceAndClientQueryVO) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *WifiDeviceAndClientQueryVO) SetTopK(v int32)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *WifiDeviceAndClientQueryVO) HasTopK() bool`

HasTopK returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


