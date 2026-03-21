# ClassificationOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | Pointer to **int64** | Attempts. | [optional] 
**Classification** | Pointer to **string** | Classification. | [optional] 

## Methods

### NewClassificationOpenApiVO

`func NewClassificationOpenApiVO() *ClassificationOpenApiVO`

NewClassificationOpenApiVO instantiates a new ClassificationOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassificationOpenApiVOWithDefaults

`func NewClassificationOpenApiVOWithDefaults() *ClassificationOpenApiVO`

NewClassificationOpenApiVOWithDefaults instantiates a new ClassificationOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *ClassificationOpenApiVO) GetAttempts() int64`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *ClassificationOpenApiVO) GetAttemptsOk() (*int64, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *ClassificationOpenApiVO) SetAttempts(v int64)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *ClassificationOpenApiVO) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetClassification

`func (o *ClassificationOpenApiVO) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *ClassificationOpenApiVO) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *ClassificationOpenApiVO) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *ClassificationOpenApiVO) HasClassification() bool`

HasClassification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


