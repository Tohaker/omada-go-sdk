# FormAuthOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnswerNum** | Pointer to **int64** | The number of form&#39;s answers. | [optional] 
**AuthTimeout** | Pointer to [**AuthTimeOpenApiVO**](AuthTimeOpenApiVO.md) |  | [optional] 
**CardList** | Pointer to [**[]FormAuthCardOpenApiVO**](FormAuthCardOpenApiVO.md) | Form card list. | [optional] 
**CreateTime** | Pointer to **int64** | Created time of the form. | [optional] 
**Id** | Pointer to **string** | Auth form ID | [optional] 
**IsPublished** | Pointer to **bool** | Whether to publish. | [optional] 
**Name** | Pointer to **string** | Form name (display for the controller user). It should contain 1-2000 characters. | [optional] 
**Note** | Pointer to **string** | Note should contain 1-2000 characters. | [optional] 
**Portals** | Pointer to **[]string** | Portal names corresponding to the bound portal. | [optional] 
**Title** | Pointer to **string** | Form title (display for the authentication user). It should contain 1-2000 characters. | [optional] 

## Methods

### NewFormAuthOpenApiVO

`func NewFormAuthOpenApiVO() *FormAuthOpenApiVO`

NewFormAuthOpenApiVO instantiates a new FormAuthOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormAuthOpenApiVOWithDefaults

`func NewFormAuthOpenApiVOWithDefaults() *FormAuthOpenApiVO`

NewFormAuthOpenApiVOWithDefaults instantiates a new FormAuthOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswerNum

`func (o *FormAuthOpenApiVO) GetAnswerNum() int64`

GetAnswerNum returns the AnswerNum field if non-nil, zero value otherwise.

### GetAnswerNumOk

`func (o *FormAuthOpenApiVO) GetAnswerNumOk() (*int64, bool)`

GetAnswerNumOk returns a tuple with the AnswerNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerNum

`func (o *FormAuthOpenApiVO) SetAnswerNum(v int64)`

SetAnswerNum sets AnswerNum field to given value.

### HasAnswerNum

`func (o *FormAuthOpenApiVO) HasAnswerNum() bool`

HasAnswerNum returns a boolean if a field has been set.

### GetAuthTimeout

`func (o *FormAuthOpenApiVO) GetAuthTimeout() AuthTimeOpenApiVO`

GetAuthTimeout returns the AuthTimeout field if non-nil, zero value otherwise.

### GetAuthTimeoutOk

`func (o *FormAuthOpenApiVO) GetAuthTimeoutOk() (*AuthTimeOpenApiVO, bool)`

GetAuthTimeoutOk returns a tuple with the AuthTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTimeout

`func (o *FormAuthOpenApiVO) SetAuthTimeout(v AuthTimeOpenApiVO)`

SetAuthTimeout sets AuthTimeout field to given value.

### HasAuthTimeout

`func (o *FormAuthOpenApiVO) HasAuthTimeout() bool`

HasAuthTimeout returns a boolean if a field has been set.

### GetCardList

`func (o *FormAuthOpenApiVO) GetCardList() []FormAuthCardOpenApiVO`

GetCardList returns the CardList field if non-nil, zero value otherwise.

### GetCardListOk

`func (o *FormAuthOpenApiVO) GetCardListOk() (*[]FormAuthCardOpenApiVO, bool)`

GetCardListOk returns a tuple with the CardList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardList

`func (o *FormAuthOpenApiVO) SetCardList(v []FormAuthCardOpenApiVO)`

SetCardList sets CardList field to given value.

### HasCardList

`func (o *FormAuthOpenApiVO) HasCardList() bool`

HasCardList returns a boolean if a field has been set.

### GetCreateTime

`func (o *FormAuthOpenApiVO) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *FormAuthOpenApiVO) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *FormAuthOpenApiVO) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *FormAuthOpenApiVO) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetId

`func (o *FormAuthOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormAuthOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormAuthOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormAuthOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPublished

`func (o *FormAuthOpenApiVO) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *FormAuthOpenApiVO) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *FormAuthOpenApiVO) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *FormAuthOpenApiVO) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetName

`func (o *FormAuthOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormAuthOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormAuthOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FormAuthOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNote

`func (o *FormAuthOpenApiVO) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *FormAuthOpenApiVO) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *FormAuthOpenApiVO) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *FormAuthOpenApiVO) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPortals

`func (o *FormAuthOpenApiVO) GetPortals() []string`

GetPortals returns the Portals field if non-nil, zero value otherwise.

### GetPortalsOk

`func (o *FormAuthOpenApiVO) GetPortalsOk() (*[]string, bool)`

GetPortalsOk returns a tuple with the Portals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortals

`func (o *FormAuthOpenApiVO) SetPortals(v []string)`

SetPortals sets Portals field to given value.

### HasPortals

`func (o *FormAuthOpenApiVO) HasPortals() bool`

HasPortals returns a boolean if a field has been set.

### GetTitle

`func (o *FormAuthOpenApiVO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FormAuthOpenApiVO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FormAuthOpenApiVO) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FormAuthOpenApiVO) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


