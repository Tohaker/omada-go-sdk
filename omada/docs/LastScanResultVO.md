# LastScanResultVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfterIndex** | Pointer to **int32** | Index after WLAN Optimization, between 0 and 100. | [optional] 
**ApNum** | Pointer to **int32** | The number of EAPs. | [optional] 
**BeforeIndex** | Pointer to **int32** | Index before WLAN Optimization, between 0 and 100. | [optional] 
**IsAppliedSuccess** | Pointer to **bool** | Parameter [isAppliedSuccess] means whether WLAN Optimization executes successfully. | [optional] 
**Length** | Pointer to **int32** | Parameter [length] means the duration of the WLAN Optimization in seconds. | [optional] 
**Mode** | Pointer to **int32** | 0: by WLAN Optimization schedule. 1: by one-click WLAN Optimization. | [optional] 
**Time** | Pointer to **int64** | Timestamp corresponding to the start of optimization. | [optional] 

## Methods

### NewLastScanResultVO

`func NewLastScanResultVO() *LastScanResultVO`

NewLastScanResultVO instantiates a new LastScanResultVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastScanResultVOWithDefaults

`func NewLastScanResultVOWithDefaults() *LastScanResultVO`

NewLastScanResultVOWithDefaults instantiates a new LastScanResultVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfterIndex

`func (o *LastScanResultVO) GetAfterIndex() int32`

GetAfterIndex returns the AfterIndex field if non-nil, zero value otherwise.

### GetAfterIndexOk

`func (o *LastScanResultVO) GetAfterIndexOk() (*int32, bool)`

GetAfterIndexOk returns a tuple with the AfterIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterIndex

`func (o *LastScanResultVO) SetAfterIndex(v int32)`

SetAfterIndex sets AfterIndex field to given value.

### HasAfterIndex

`func (o *LastScanResultVO) HasAfterIndex() bool`

HasAfterIndex returns a boolean if a field has been set.

### GetApNum

`func (o *LastScanResultVO) GetApNum() int32`

GetApNum returns the ApNum field if non-nil, zero value otherwise.

### GetApNumOk

`func (o *LastScanResultVO) GetApNumOk() (*int32, bool)`

GetApNumOk returns a tuple with the ApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApNum

`func (o *LastScanResultVO) SetApNum(v int32)`

SetApNum sets ApNum field to given value.

### HasApNum

`func (o *LastScanResultVO) HasApNum() bool`

HasApNum returns a boolean if a field has been set.

### GetBeforeIndex

`func (o *LastScanResultVO) GetBeforeIndex() int32`

GetBeforeIndex returns the BeforeIndex field if non-nil, zero value otherwise.

### GetBeforeIndexOk

`func (o *LastScanResultVO) GetBeforeIndexOk() (*int32, bool)`

GetBeforeIndexOk returns a tuple with the BeforeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeIndex

`func (o *LastScanResultVO) SetBeforeIndex(v int32)`

SetBeforeIndex sets BeforeIndex field to given value.

### HasBeforeIndex

`func (o *LastScanResultVO) HasBeforeIndex() bool`

HasBeforeIndex returns a boolean if a field has been set.

### GetIsAppliedSuccess

`func (o *LastScanResultVO) GetIsAppliedSuccess() bool`

GetIsAppliedSuccess returns the IsAppliedSuccess field if non-nil, zero value otherwise.

### GetIsAppliedSuccessOk

`func (o *LastScanResultVO) GetIsAppliedSuccessOk() (*bool, bool)`

GetIsAppliedSuccessOk returns a tuple with the IsAppliedSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAppliedSuccess

`func (o *LastScanResultVO) SetIsAppliedSuccess(v bool)`

SetIsAppliedSuccess sets IsAppliedSuccess field to given value.

### HasIsAppliedSuccess

`func (o *LastScanResultVO) HasIsAppliedSuccess() bool`

HasIsAppliedSuccess returns a boolean if a field has been set.

### GetLength

`func (o *LastScanResultVO) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *LastScanResultVO) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *LastScanResultVO) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *LastScanResultVO) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetMode

`func (o *LastScanResultVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *LastScanResultVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *LastScanResultVO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *LastScanResultVO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetTime

`func (o *LastScanResultVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *LastScanResultVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *LastScanResultVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *LastScanResultVO) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


