# CreateFormAuthOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthTimeout** | [**AuthTimeOpenApiVO**](AuthTimeOpenApiVO.md) |  | 
**CardList** | [**[]FormAuthCardOpenApiVO**](FormAuthCardOpenApiVO.md) | Form card list. | 
**Name** | **string** | Form name (display for the controller user). It should contain 1-2000 characters. | 
**Note** | **string** | Note should contain 1-2000 characters. | 
**Published** | **bool** | Whether to publish. | 
**Title** | **string** | Form title (display for the authentication user). It should contain 1-2000 characters. | 

## Methods

### NewCreateFormAuthOpenApiVO

`func NewCreateFormAuthOpenApiVO(authTimeout AuthTimeOpenApiVO, cardList []FormAuthCardOpenApiVO, name string, note string, published bool, title string, ) *CreateFormAuthOpenApiVO`

NewCreateFormAuthOpenApiVO instantiates a new CreateFormAuthOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFormAuthOpenApiVOWithDefaults

`func NewCreateFormAuthOpenApiVOWithDefaults() *CreateFormAuthOpenApiVO`

NewCreateFormAuthOpenApiVOWithDefaults instantiates a new CreateFormAuthOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthTimeout

`func (o *CreateFormAuthOpenApiVO) GetAuthTimeout() AuthTimeOpenApiVO`

GetAuthTimeout returns the AuthTimeout field if non-nil, zero value otherwise.

### GetAuthTimeoutOk

`func (o *CreateFormAuthOpenApiVO) GetAuthTimeoutOk() (*AuthTimeOpenApiVO, bool)`

GetAuthTimeoutOk returns a tuple with the AuthTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTimeout

`func (o *CreateFormAuthOpenApiVO) SetAuthTimeout(v AuthTimeOpenApiVO)`

SetAuthTimeout sets AuthTimeout field to given value.


### GetCardList

`func (o *CreateFormAuthOpenApiVO) GetCardList() []FormAuthCardOpenApiVO`

GetCardList returns the CardList field if non-nil, zero value otherwise.

### GetCardListOk

`func (o *CreateFormAuthOpenApiVO) GetCardListOk() (*[]FormAuthCardOpenApiVO, bool)`

GetCardListOk returns a tuple with the CardList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardList

`func (o *CreateFormAuthOpenApiVO) SetCardList(v []FormAuthCardOpenApiVO)`

SetCardList sets CardList field to given value.


### GetName

`func (o *CreateFormAuthOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFormAuthOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFormAuthOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetNote

`func (o *CreateFormAuthOpenApiVO) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CreateFormAuthOpenApiVO) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CreateFormAuthOpenApiVO) SetNote(v string)`

SetNote sets Note field to given value.


### GetPublished

`func (o *CreateFormAuthOpenApiVO) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *CreateFormAuthOpenApiVO) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *CreateFormAuthOpenApiVO) SetPublished(v bool)`

SetPublished sets Published field to given value.


### GetTitle

`func (o *CreateFormAuthOpenApiVO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateFormAuthOpenApiVO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateFormAuthOpenApiVO) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


