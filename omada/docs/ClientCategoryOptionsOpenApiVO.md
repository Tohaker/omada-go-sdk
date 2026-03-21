# ClientCategoryOptionsOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Name of client category. | [optional] 
**Type** | Pointer to **[]string** | Type list of clients. | [optional] 

## Methods

### NewClientCategoryOptionsOpenApiVO

`func NewClientCategoryOptionsOpenApiVO() *ClientCategoryOptionsOpenApiVO`

NewClientCategoryOptionsOpenApiVO instantiates a new ClientCategoryOptionsOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCategoryOptionsOpenApiVOWithDefaults

`func NewClientCategoryOptionsOpenApiVOWithDefaults() *ClientCategoryOptionsOpenApiVO`

NewClientCategoryOptionsOpenApiVOWithDefaults instantiates a new ClientCategoryOptionsOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *ClientCategoryOptionsOpenApiVO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ClientCategoryOptionsOpenApiVO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ClientCategoryOptionsOpenApiVO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ClientCategoryOptionsOpenApiVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetType

`func (o *ClientCategoryOptionsOpenApiVO) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientCategoryOptionsOpenApiVO) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientCategoryOptionsOpenApiVO) SetType(v []string)`

SetType sets Type field to given value.

### HasType

`func (o *ClientCategoryOptionsOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


