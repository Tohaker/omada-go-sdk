# CommonSubHealthInfoDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageNum** | Pointer to **int32** | Average value of common dimension, such as cpu、memory | [optional] 
**PastNums** | Pointer to [**[]TimeValueItemVO**](TimeValueItemVO.md) | List of common dimension value , such as cpu、memory | [optional] 
**SummaryScore** | Pointer to **int32** | Sub dimension health score | [optional] 
**Support** | Pointer to **bool** | Sub dimension support | [optional] 

## Methods

### NewCommonSubHealthInfoDetailVO

`func NewCommonSubHealthInfoDetailVO() *CommonSubHealthInfoDetailVO`

NewCommonSubHealthInfoDetailVO instantiates a new CommonSubHealthInfoDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonSubHealthInfoDetailVOWithDefaults

`func NewCommonSubHealthInfoDetailVOWithDefaults() *CommonSubHealthInfoDetailVO`

NewCommonSubHealthInfoDetailVOWithDefaults instantiates a new CommonSubHealthInfoDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageNum

`func (o *CommonSubHealthInfoDetailVO) GetAverageNum() int32`

GetAverageNum returns the AverageNum field if non-nil, zero value otherwise.

### GetAverageNumOk

`func (o *CommonSubHealthInfoDetailVO) GetAverageNumOk() (*int32, bool)`

GetAverageNumOk returns a tuple with the AverageNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageNum

`func (o *CommonSubHealthInfoDetailVO) SetAverageNum(v int32)`

SetAverageNum sets AverageNum field to given value.

### HasAverageNum

`func (o *CommonSubHealthInfoDetailVO) HasAverageNum() bool`

HasAverageNum returns a boolean if a field has been set.

### GetPastNums

`func (o *CommonSubHealthInfoDetailVO) GetPastNums() []TimeValueItemVO`

GetPastNums returns the PastNums field if non-nil, zero value otherwise.

### GetPastNumsOk

`func (o *CommonSubHealthInfoDetailVO) GetPastNumsOk() (*[]TimeValueItemVO, bool)`

GetPastNumsOk returns a tuple with the PastNums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastNums

`func (o *CommonSubHealthInfoDetailVO) SetPastNums(v []TimeValueItemVO)`

SetPastNums sets PastNums field to given value.

### HasPastNums

`func (o *CommonSubHealthInfoDetailVO) HasPastNums() bool`

HasPastNums returns a boolean if a field has been set.

### GetSummaryScore

`func (o *CommonSubHealthInfoDetailVO) GetSummaryScore() int32`

GetSummaryScore returns the SummaryScore field if non-nil, zero value otherwise.

### GetSummaryScoreOk

`func (o *CommonSubHealthInfoDetailVO) GetSummaryScoreOk() (*int32, bool)`

GetSummaryScoreOk returns a tuple with the SummaryScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryScore

`func (o *CommonSubHealthInfoDetailVO) SetSummaryScore(v int32)`

SetSummaryScore sets SummaryScore field to given value.

### HasSummaryScore

`func (o *CommonSubHealthInfoDetailVO) HasSummaryScore() bool`

HasSummaryScore returns a boolean if a field has been set.

### GetSupport

`func (o *CommonSubHealthInfoDetailVO) GetSupport() bool`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *CommonSubHealthInfoDetailVO) GetSupportOk() (*bool, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *CommonSubHealthInfoDetailVO) SetSupport(v bool)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *CommonSubHealthInfoDetailVO) HasSupport() bool`

HasSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


