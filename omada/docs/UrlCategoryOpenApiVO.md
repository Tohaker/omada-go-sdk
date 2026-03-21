# UrlCategoryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **map[string][]string** | categories of the URL filtering 5.15 | [optional] 
**ProtocolVer** | Pointer to **string** | protocolVer of the categories. | [optional] 

## Methods

### NewUrlCategoryOpenApiVO

`func NewUrlCategoryOpenApiVO() *UrlCategoryOpenApiVO`

NewUrlCategoryOpenApiVO instantiates a new UrlCategoryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrlCategoryOpenApiVOWithDefaults

`func NewUrlCategoryOpenApiVOWithDefaults() *UrlCategoryOpenApiVO`

NewUrlCategoryOpenApiVOWithDefaults instantiates a new UrlCategoryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *UrlCategoryOpenApiVO) GetCategories() map[string][]string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *UrlCategoryOpenApiVO) GetCategoriesOk() (*map[string][]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *UrlCategoryOpenApiVO) SetCategories(v map[string][]string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *UrlCategoryOpenApiVO) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetProtocolVer

`func (o *UrlCategoryOpenApiVO) GetProtocolVer() string`

GetProtocolVer returns the ProtocolVer field if non-nil, zero value otherwise.

### GetProtocolVerOk

`func (o *UrlCategoryOpenApiVO) GetProtocolVerOk() (*string, bool)`

GetProtocolVerOk returns a tuple with the ProtocolVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVer

`func (o *UrlCategoryOpenApiVO) SetProtocolVer(v string)`

SetProtocolVer sets ProtocolVer field to given value.

### HasProtocolVer

`func (o *UrlCategoryOpenApiVO) HasProtocolVer() bool`

HasProtocolVer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


