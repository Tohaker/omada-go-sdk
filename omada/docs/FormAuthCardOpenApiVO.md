# FormAuthCardOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Choices** | **[]string** | Drop-down menu or list of options for multiple-choice questions (exist if and only if cardType is 0, 1, 2), excluding questions with other options. Choices should contain 1-2000 characters | 
**Others** | Pointer to **string** | The question of the other options. Exists if and only if type is 0, 1. Others should contain 1-2000 characters | [optional] 
**Required** | **bool** | Indicates whether the card is required (cardType is 5, meaningless in the prompt box). | 
**ScoreNotes** | Pointer to **[]string** | The prompt text corresponding to different scores of the scoring card. Exists when the cardType is 4 and needs to be transmitted in order. The subscript 0 corresponds to the score 1. | [optional] 
**Title** | **string** | Title should contain 1-2000 characters | 
**Type** | **int32** | Card type should be a value as follows: 0: single choice (single choice); 1: multiple choice (multiple choice); 2: ComboBox (drop-down menu); 3: input (input box); 4: score (score); 5: prompt frame | 

## Methods

### NewFormAuthCardOpenApiVO

`func NewFormAuthCardOpenApiVO(choices []string, required bool, title string, type_ int32, ) *FormAuthCardOpenApiVO`

NewFormAuthCardOpenApiVO instantiates a new FormAuthCardOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormAuthCardOpenApiVOWithDefaults

`func NewFormAuthCardOpenApiVOWithDefaults() *FormAuthCardOpenApiVO`

NewFormAuthCardOpenApiVOWithDefaults instantiates a new FormAuthCardOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoices

`func (o *FormAuthCardOpenApiVO) GetChoices() []string`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *FormAuthCardOpenApiVO) GetChoicesOk() (*[]string, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *FormAuthCardOpenApiVO) SetChoices(v []string)`

SetChoices sets Choices field to given value.


### GetOthers

`func (o *FormAuthCardOpenApiVO) GetOthers() string`

GetOthers returns the Others field if non-nil, zero value otherwise.

### GetOthersOk

`func (o *FormAuthCardOpenApiVO) GetOthersOk() (*string, bool)`

GetOthersOk returns a tuple with the Others field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOthers

`func (o *FormAuthCardOpenApiVO) SetOthers(v string)`

SetOthers sets Others field to given value.

### HasOthers

`func (o *FormAuthCardOpenApiVO) HasOthers() bool`

HasOthers returns a boolean if a field has been set.

### GetRequired

`func (o *FormAuthCardOpenApiVO) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormAuthCardOpenApiVO) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormAuthCardOpenApiVO) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetScoreNotes

`func (o *FormAuthCardOpenApiVO) GetScoreNotes() []string`

GetScoreNotes returns the ScoreNotes field if non-nil, zero value otherwise.

### GetScoreNotesOk

`func (o *FormAuthCardOpenApiVO) GetScoreNotesOk() (*[]string, bool)`

GetScoreNotesOk returns a tuple with the ScoreNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreNotes

`func (o *FormAuthCardOpenApiVO) SetScoreNotes(v []string)`

SetScoreNotes sets ScoreNotes field to given value.

### HasScoreNotes

`func (o *FormAuthCardOpenApiVO) HasScoreNotes() bool`

HasScoreNotes returns a boolean if a field has been set.

### GetTitle

`func (o *FormAuthCardOpenApiVO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FormAuthCardOpenApiVO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FormAuthCardOpenApiVO) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *FormAuthCardOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormAuthCardOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormAuthCardOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


