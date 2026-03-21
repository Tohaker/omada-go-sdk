# SubSingleHealthInfoDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scores** | Pointer to [**[]TimeScoreItemVO**](TimeScoreItemVO.md) |  | [optional] 
**SummaryScore** | Pointer to **int32** | Sub dimension health score | [optional] 
**Support** | Pointer to **bool** | Sub dimension support | [optional] 

## Methods

### NewSubSingleHealthInfoDetailVO

`func NewSubSingleHealthInfoDetailVO() *SubSingleHealthInfoDetailVO`

NewSubSingleHealthInfoDetailVO instantiates a new SubSingleHealthInfoDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubSingleHealthInfoDetailVOWithDefaults

`func NewSubSingleHealthInfoDetailVOWithDefaults() *SubSingleHealthInfoDetailVO`

NewSubSingleHealthInfoDetailVOWithDefaults instantiates a new SubSingleHealthInfoDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScores

`func (o *SubSingleHealthInfoDetailVO) GetScores() []TimeScoreItemVO`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *SubSingleHealthInfoDetailVO) GetScoresOk() (*[]TimeScoreItemVO, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *SubSingleHealthInfoDetailVO) SetScores(v []TimeScoreItemVO)`

SetScores sets Scores field to given value.

### HasScores

`func (o *SubSingleHealthInfoDetailVO) HasScores() bool`

HasScores returns a boolean if a field has been set.

### GetSummaryScore

`func (o *SubSingleHealthInfoDetailVO) GetSummaryScore() int32`

GetSummaryScore returns the SummaryScore field if non-nil, zero value otherwise.

### GetSummaryScoreOk

`func (o *SubSingleHealthInfoDetailVO) GetSummaryScoreOk() (*int32, bool)`

GetSummaryScoreOk returns a tuple with the SummaryScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryScore

`func (o *SubSingleHealthInfoDetailVO) SetSummaryScore(v int32)`

SetSummaryScore sets SummaryScore field to given value.

### HasSummaryScore

`func (o *SubSingleHealthInfoDetailVO) HasSummaryScore() bool`

HasSummaryScore returns a boolean if a field has been set.

### GetSupport

`func (o *SubSingleHealthInfoDetailVO) GetSupport() bool`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *SubSingleHealthInfoDetailVO) GetSupportOk() (*bool, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *SubSingleHealthInfoDetailVO) SetSupport(v bool)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *SubSingleHealthInfoDetailVO) HasSupport() bool`

HasSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


