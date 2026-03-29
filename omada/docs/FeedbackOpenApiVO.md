# FeedbackOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | Feedback and evaluation of the controller service. | 
**Level** | **int32** | The rating of the controller service is in the range of 1~5. | 

## Methods

### NewFeedbackOpenApiVO

`func NewFeedbackOpenApiVO(content string, level int32, ) *FeedbackOpenApiVO`

NewFeedbackOpenApiVO instantiates a new FeedbackOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackOpenApiVOWithDefaults

`func NewFeedbackOpenApiVOWithDefaults() *FeedbackOpenApiVO`

NewFeedbackOpenApiVOWithDefaults instantiates a new FeedbackOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *FeedbackOpenApiVO) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *FeedbackOpenApiVO) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *FeedbackOpenApiVO) SetContent(v string)`

SetContent sets Content field to given value.


### GetLevel

`func (o *FeedbackOpenApiVO) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *FeedbackOpenApiVO) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *FeedbackOpenApiVO) SetLevel(v int32)`

SetLevel sets Level field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


