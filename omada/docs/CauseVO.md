# CauseVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advices** | Pointer to [**map[string]AnomalyAdviceVO**](AnomalyAdviceVO.md) | Advices for this root cause. Key is advice code, value is advice detail. | [optional] 
**CauseCode** | Pointer to **string** | Root cause code. | [optional] 
**CauseTitleParams** | Pointer to **map[string]string** | Title parameter map for rendering the root cause title template. | [optional] 

## Methods

### NewCauseVO

`func NewCauseVO() *CauseVO`

NewCauseVO instantiates a new CauseVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCauseVOWithDefaults

`func NewCauseVOWithDefaults() *CauseVO`

NewCauseVOWithDefaults instantiates a new CauseVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvices

`func (o *CauseVO) GetAdvices() map[string]AnomalyAdviceVO`

GetAdvices returns the Advices field if non-nil, zero value otherwise.

### GetAdvicesOk

`func (o *CauseVO) GetAdvicesOk() (*map[string]AnomalyAdviceVO, bool)`

GetAdvicesOk returns a tuple with the Advices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvices

`func (o *CauseVO) SetAdvices(v map[string]AnomalyAdviceVO)`

SetAdvices sets Advices field to given value.

### HasAdvices

`func (o *CauseVO) HasAdvices() bool`

HasAdvices returns a boolean if a field has been set.

### GetCauseCode

`func (o *CauseVO) GetCauseCode() string`

GetCauseCode returns the CauseCode field if non-nil, zero value otherwise.

### GetCauseCodeOk

`func (o *CauseVO) GetCauseCodeOk() (*string, bool)`

GetCauseCodeOk returns a tuple with the CauseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseCode

`func (o *CauseVO) SetCauseCode(v string)`

SetCauseCode sets CauseCode field to given value.

### HasCauseCode

`func (o *CauseVO) HasCauseCode() bool`

HasCauseCode returns a boolean if a field has been set.

### GetCauseTitleParams

`func (o *CauseVO) GetCauseTitleParams() map[string]string`

GetCauseTitleParams returns the CauseTitleParams field if non-nil, zero value otherwise.

### GetCauseTitleParamsOk

`func (o *CauseVO) GetCauseTitleParamsOk() (*map[string]string, bool)`

GetCauseTitleParamsOk returns a tuple with the CauseTitleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseTitleParams

`func (o *CauseVO) SetCauseTitleParams(v map[string]string)`

SetCauseTitleParams sets CauseTitleParams field to given value.

### HasCauseTitleParams

`func (o *CauseVO) HasCauseTitleParams() bool`

HasCauseTitleParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


