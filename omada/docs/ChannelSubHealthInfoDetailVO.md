# ChannelSubHealthInfoDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageNum2g** | Pointer to **int32** | Average value of 2g channel | [optional] 
**AverageNum5g** | Pointer to **int32** | Average value of 5g channel | [optional] 
**AverageNum6g** | Pointer to **int32** | Average value of 6g channel | [optional] 
**PastNums2g** | Pointer to [**[]TimeValueItemVO**](TimeValueItemVO.md) | List of 2g channel value | [optional] 
**PastNums5g** | Pointer to [**[]TimeValueItemVO**](TimeValueItemVO.md) | List of 5g channel value | [optional] 
**PastNums6g** | Pointer to [**[]TimeValueItemVO**](TimeValueItemVO.md) | List of 6g channel value | [optional] 
**SummaryScore** | Pointer to **int32** | Sub dimension health score | [optional] 
**Support** | Pointer to **bool** | Sub dimension support | [optional] 

## Methods

### NewChannelSubHealthInfoDetailVO

`func NewChannelSubHealthInfoDetailVO() *ChannelSubHealthInfoDetailVO`

NewChannelSubHealthInfoDetailVO instantiates a new ChannelSubHealthInfoDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelSubHealthInfoDetailVOWithDefaults

`func NewChannelSubHealthInfoDetailVOWithDefaults() *ChannelSubHealthInfoDetailVO`

NewChannelSubHealthInfoDetailVOWithDefaults instantiates a new ChannelSubHealthInfoDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageNum2g

`func (o *ChannelSubHealthInfoDetailVO) GetAverageNum2g() int32`

GetAverageNum2g returns the AverageNum2g field if non-nil, zero value otherwise.

### GetAverageNum2gOk

`func (o *ChannelSubHealthInfoDetailVO) GetAverageNum2gOk() (*int32, bool)`

GetAverageNum2gOk returns a tuple with the AverageNum2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageNum2g

`func (o *ChannelSubHealthInfoDetailVO) SetAverageNum2g(v int32)`

SetAverageNum2g sets AverageNum2g field to given value.

### HasAverageNum2g

`func (o *ChannelSubHealthInfoDetailVO) HasAverageNum2g() bool`

HasAverageNum2g returns a boolean if a field has been set.

### GetAverageNum5g

`func (o *ChannelSubHealthInfoDetailVO) GetAverageNum5g() int32`

GetAverageNum5g returns the AverageNum5g field if non-nil, zero value otherwise.

### GetAverageNum5gOk

`func (o *ChannelSubHealthInfoDetailVO) GetAverageNum5gOk() (*int32, bool)`

GetAverageNum5gOk returns a tuple with the AverageNum5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageNum5g

`func (o *ChannelSubHealthInfoDetailVO) SetAverageNum5g(v int32)`

SetAverageNum5g sets AverageNum5g field to given value.

### HasAverageNum5g

`func (o *ChannelSubHealthInfoDetailVO) HasAverageNum5g() bool`

HasAverageNum5g returns a boolean if a field has been set.

### GetAverageNum6g

`func (o *ChannelSubHealthInfoDetailVO) GetAverageNum6g() int32`

GetAverageNum6g returns the AverageNum6g field if non-nil, zero value otherwise.

### GetAverageNum6gOk

`func (o *ChannelSubHealthInfoDetailVO) GetAverageNum6gOk() (*int32, bool)`

GetAverageNum6gOk returns a tuple with the AverageNum6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageNum6g

`func (o *ChannelSubHealthInfoDetailVO) SetAverageNum6g(v int32)`

SetAverageNum6g sets AverageNum6g field to given value.

### HasAverageNum6g

`func (o *ChannelSubHealthInfoDetailVO) HasAverageNum6g() bool`

HasAverageNum6g returns a boolean if a field has been set.

### GetPastNums2g

`func (o *ChannelSubHealthInfoDetailVO) GetPastNums2g() []TimeValueItemVO`

GetPastNums2g returns the PastNums2g field if non-nil, zero value otherwise.

### GetPastNums2gOk

`func (o *ChannelSubHealthInfoDetailVO) GetPastNums2gOk() (*[]TimeValueItemVO, bool)`

GetPastNums2gOk returns a tuple with the PastNums2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastNums2g

`func (o *ChannelSubHealthInfoDetailVO) SetPastNums2g(v []TimeValueItemVO)`

SetPastNums2g sets PastNums2g field to given value.

### HasPastNums2g

`func (o *ChannelSubHealthInfoDetailVO) HasPastNums2g() bool`

HasPastNums2g returns a boolean if a field has been set.

### GetPastNums5g

`func (o *ChannelSubHealthInfoDetailVO) GetPastNums5g() []TimeValueItemVO`

GetPastNums5g returns the PastNums5g field if non-nil, zero value otherwise.

### GetPastNums5gOk

`func (o *ChannelSubHealthInfoDetailVO) GetPastNums5gOk() (*[]TimeValueItemVO, bool)`

GetPastNums5gOk returns a tuple with the PastNums5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastNums5g

`func (o *ChannelSubHealthInfoDetailVO) SetPastNums5g(v []TimeValueItemVO)`

SetPastNums5g sets PastNums5g field to given value.

### HasPastNums5g

`func (o *ChannelSubHealthInfoDetailVO) HasPastNums5g() bool`

HasPastNums5g returns a boolean if a field has been set.

### GetPastNums6g

`func (o *ChannelSubHealthInfoDetailVO) GetPastNums6g() []TimeValueItemVO`

GetPastNums6g returns the PastNums6g field if non-nil, zero value otherwise.

### GetPastNums6gOk

`func (o *ChannelSubHealthInfoDetailVO) GetPastNums6gOk() (*[]TimeValueItemVO, bool)`

GetPastNums6gOk returns a tuple with the PastNums6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastNums6g

`func (o *ChannelSubHealthInfoDetailVO) SetPastNums6g(v []TimeValueItemVO)`

SetPastNums6g sets PastNums6g field to given value.

### HasPastNums6g

`func (o *ChannelSubHealthInfoDetailVO) HasPastNums6g() bool`

HasPastNums6g returns a boolean if a field has been set.

### GetSummaryScore

`func (o *ChannelSubHealthInfoDetailVO) GetSummaryScore() int32`

GetSummaryScore returns the SummaryScore field if non-nil, zero value otherwise.

### GetSummaryScoreOk

`func (o *ChannelSubHealthInfoDetailVO) GetSummaryScoreOk() (*int32, bool)`

GetSummaryScoreOk returns a tuple with the SummaryScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryScore

`func (o *ChannelSubHealthInfoDetailVO) SetSummaryScore(v int32)`

SetSummaryScore sets SummaryScore field to given value.

### HasSummaryScore

`func (o *ChannelSubHealthInfoDetailVO) HasSummaryScore() bool`

HasSummaryScore returns a boolean if a field has been set.

### GetSupport

`func (o *ChannelSubHealthInfoDetailVO) GetSupport() bool`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *ChannelSubHealthInfoDetailVO) GetSupportOk() (*bool, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *ChannelSubHealthInfoDetailVO) SetSupport(v bool)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *ChannelSubHealthInfoDetailVO) HasSupport() bool`

HasSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


