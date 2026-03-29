# FormAuthResultOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answers** | Pointer to [**[]FormAuthCardAnswerOpenApiVO**](FormAuthCardAnswerOpenApiVO.md) | Answers submitted by users. | [optional] 
**FormId** | Pointer to **string** | This field represents authentication survey ID. Authentication survey can be created using &#39;Create a new authentication survey&#39; interface, and authentication survey ID can be obtained from &#39;Get authentication survey list&#39; interface. | [optional] 
**Id** | Pointer to **string** | Form result ID | [optional] 
**Network** | Pointer to **string** | Network (exists when wired connection). | [optional] 
**Ssid** | Pointer to **string** | Client connected SSID. | [optional] 
**Time** | Pointer to **int64** | Authenticated timestamp in ms. | [optional] 

## Methods

### NewFormAuthResultOpenApiVO

`func NewFormAuthResultOpenApiVO() *FormAuthResultOpenApiVO`

NewFormAuthResultOpenApiVO instantiates a new FormAuthResultOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormAuthResultOpenApiVOWithDefaults

`func NewFormAuthResultOpenApiVOWithDefaults() *FormAuthResultOpenApiVO`

NewFormAuthResultOpenApiVOWithDefaults instantiates a new FormAuthResultOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswers

`func (o *FormAuthResultOpenApiVO) GetAnswers() []FormAuthCardAnswerOpenApiVO`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *FormAuthResultOpenApiVO) GetAnswersOk() (*[]FormAuthCardAnswerOpenApiVO, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *FormAuthResultOpenApiVO) SetAnswers(v []FormAuthCardAnswerOpenApiVO)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *FormAuthResultOpenApiVO) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetFormId

`func (o *FormAuthResultOpenApiVO) GetFormId() string`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *FormAuthResultOpenApiVO) GetFormIdOk() (*string, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *FormAuthResultOpenApiVO) SetFormId(v string)`

SetFormId sets FormId field to given value.

### HasFormId

`func (o *FormAuthResultOpenApiVO) HasFormId() bool`

HasFormId returns a boolean if a field has been set.

### GetId

`func (o *FormAuthResultOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormAuthResultOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormAuthResultOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormAuthResultOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetwork

`func (o *FormAuthResultOpenApiVO) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *FormAuthResultOpenApiVO) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *FormAuthResultOpenApiVO) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *FormAuthResultOpenApiVO) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSsid

`func (o *FormAuthResultOpenApiVO) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *FormAuthResultOpenApiVO) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *FormAuthResultOpenApiVO) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *FormAuthResultOpenApiVO) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetTime

`func (o *FormAuthResultOpenApiVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *FormAuthResultOpenApiVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *FormAuthResultOpenApiVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *FormAuthResultOpenApiVO) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


