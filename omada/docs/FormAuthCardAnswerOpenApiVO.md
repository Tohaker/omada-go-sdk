# FormAuthCardAnswerOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChoiceAnswer** | Pointer to **[]int32** | Drop-down menu or answer list for multiple-choice questions (exists if and only if cardType is 0, 1, 2) | [optional] 
**InputAnswer** | Pointer to **string** | Answers entered by the user, which corresponds to the question and answer question or the scored evaluation, or the input of other options in the multiple-choice question. Determined according to the type field. 0, 1, and 2 indicate other options answers, 3 indicates the answer in the input box, and 4 indicates the scored evaluation. InputAnswer should contain 1-2000 characters | [optional] 
**Others** | Pointer to **string** | Exists when cardType is 0, 1 and the user adds other options. Used to save the option answer corresponding to the prompt of other options. Assuming that the user adds two other options, D and E, and choiceAnswer contains 3 ( D), 4 (E), then others[0] represents the answer of option D. Others should contain 1-2000 characters | [optional] 
**Score** | Pointer to **int32** | User rating (should be within the range of 1-5), exists when type is 4 | [optional] 
**Type** | Pointer to **int32** | Type should be a value as follows: 0: single choice; 1: multiple choice; 2: ComboBox; 3: input; 4: score; 5: prompt box | [optional] 

## Methods

### NewFormAuthCardAnswerOpenApiVO

`func NewFormAuthCardAnswerOpenApiVO() *FormAuthCardAnswerOpenApiVO`

NewFormAuthCardAnswerOpenApiVO instantiates a new FormAuthCardAnswerOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormAuthCardAnswerOpenApiVOWithDefaults

`func NewFormAuthCardAnswerOpenApiVOWithDefaults() *FormAuthCardAnswerOpenApiVO`

NewFormAuthCardAnswerOpenApiVOWithDefaults instantiates a new FormAuthCardAnswerOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoiceAnswer

`func (o *FormAuthCardAnswerOpenApiVO) GetChoiceAnswer() []int32`

GetChoiceAnswer returns the ChoiceAnswer field if non-nil, zero value otherwise.

### GetChoiceAnswerOk

`func (o *FormAuthCardAnswerOpenApiVO) GetChoiceAnswerOk() (*[]int32, bool)`

GetChoiceAnswerOk returns a tuple with the ChoiceAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoiceAnswer

`func (o *FormAuthCardAnswerOpenApiVO) SetChoiceAnswer(v []int32)`

SetChoiceAnswer sets ChoiceAnswer field to given value.

### HasChoiceAnswer

`func (o *FormAuthCardAnswerOpenApiVO) HasChoiceAnswer() bool`

HasChoiceAnswer returns a boolean if a field has been set.

### GetInputAnswer

`func (o *FormAuthCardAnswerOpenApiVO) GetInputAnswer() string`

GetInputAnswer returns the InputAnswer field if non-nil, zero value otherwise.

### GetInputAnswerOk

`func (o *FormAuthCardAnswerOpenApiVO) GetInputAnswerOk() (*string, bool)`

GetInputAnswerOk returns a tuple with the InputAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputAnswer

`func (o *FormAuthCardAnswerOpenApiVO) SetInputAnswer(v string)`

SetInputAnswer sets InputAnswer field to given value.

### HasInputAnswer

`func (o *FormAuthCardAnswerOpenApiVO) HasInputAnswer() bool`

HasInputAnswer returns a boolean if a field has been set.

### GetOthers

`func (o *FormAuthCardAnswerOpenApiVO) GetOthers() string`

GetOthers returns the Others field if non-nil, zero value otherwise.

### GetOthersOk

`func (o *FormAuthCardAnswerOpenApiVO) GetOthersOk() (*string, bool)`

GetOthersOk returns a tuple with the Others field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOthers

`func (o *FormAuthCardAnswerOpenApiVO) SetOthers(v string)`

SetOthers sets Others field to given value.

### HasOthers

`func (o *FormAuthCardAnswerOpenApiVO) HasOthers() bool`

HasOthers returns a boolean if a field has been set.

### GetScore

`func (o *FormAuthCardAnswerOpenApiVO) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *FormAuthCardAnswerOpenApiVO) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *FormAuthCardAnswerOpenApiVO) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *FormAuthCardAnswerOpenApiVO) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetType

`func (o *FormAuthCardAnswerOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormAuthCardAnswerOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormAuthCardAnswerOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *FormAuthCardAnswerOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


