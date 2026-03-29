# CommonSubHealthInfoDetailVOInteger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageNum** | Pointer to [**[]TimeValueItemVOInteger**](TimeValueItemVOInteger.md) | List of common dimension value, such as rssi | [optional] 
**SummaryScore** | Pointer to **int32** | Sub dimension health score | [optional] 

## Methods

### NewCommonSubHealthInfoDetailVOInteger

`func NewCommonSubHealthInfoDetailVOInteger() *CommonSubHealthInfoDetailVOInteger`

NewCommonSubHealthInfoDetailVOInteger instantiates a new CommonSubHealthInfoDetailVOInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonSubHealthInfoDetailVOIntegerWithDefaults

`func NewCommonSubHealthInfoDetailVOIntegerWithDefaults() *CommonSubHealthInfoDetailVOInteger`

NewCommonSubHealthInfoDetailVOIntegerWithDefaults instantiates a new CommonSubHealthInfoDetailVOInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageNum

`func (o *CommonSubHealthInfoDetailVOInteger) GetAverageNum() []TimeValueItemVOInteger`

GetAverageNum returns the AverageNum field if non-nil, zero value otherwise.

### GetAverageNumOk

`func (o *CommonSubHealthInfoDetailVOInteger) GetAverageNumOk() (*[]TimeValueItemVOInteger, bool)`

GetAverageNumOk returns a tuple with the AverageNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageNum

`func (o *CommonSubHealthInfoDetailVOInteger) SetAverageNum(v []TimeValueItemVOInteger)`

SetAverageNum sets AverageNum field to given value.

### HasAverageNum

`func (o *CommonSubHealthInfoDetailVOInteger) HasAverageNum() bool`

HasAverageNum returns a boolean if a field has been set.

### GetSummaryScore

`func (o *CommonSubHealthInfoDetailVOInteger) GetSummaryScore() int32`

GetSummaryScore returns the SummaryScore field if non-nil, zero value otherwise.

### GetSummaryScoreOk

`func (o *CommonSubHealthInfoDetailVOInteger) GetSummaryScoreOk() (*int32, bool)`

GetSummaryScoreOk returns a tuple with the SummaryScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryScore

`func (o *CommonSubHealthInfoDetailVOInteger) SetSummaryScore(v int32)`

SetSummaryScore sets SummaryScore field to given value.

### HasSummaryScore

`func (o *CommonSubHealthInfoDetailVOInteger) HasSummaryScore() bool`

HasSummaryScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


