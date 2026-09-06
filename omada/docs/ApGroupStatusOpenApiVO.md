# ApGroupStatusOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApGroupNum** | Pointer to **int64** | the number of AP Groups | [optional] 
**Exceeded** | Pointer to **bool** | whether the number of AP Groups exceeds the limit | [optional] 

## Methods

### NewApGroupStatusOpenApiVO

`func NewApGroupStatusOpenApiVO() *ApGroupStatusOpenApiVO`

NewApGroupStatusOpenApiVO instantiates a new ApGroupStatusOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApGroupStatusOpenApiVOWithDefaults

`func NewApGroupStatusOpenApiVOWithDefaults() *ApGroupStatusOpenApiVO`

NewApGroupStatusOpenApiVOWithDefaults instantiates a new ApGroupStatusOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApGroupNum

`func (o *ApGroupStatusOpenApiVO) GetApGroupNum() int64`

GetApGroupNum returns the ApGroupNum field if non-nil, zero value otherwise.

### GetApGroupNumOk

`func (o *ApGroupStatusOpenApiVO) GetApGroupNumOk() (*int64, bool)`

GetApGroupNumOk returns a tuple with the ApGroupNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApGroupNum

`func (o *ApGroupStatusOpenApiVO) SetApGroupNum(v int64)`

SetApGroupNum sets ApGroupNum field to given value.

### HasApGroupNum

`func (o *ApGroupStatusOpenApiVO) HasApGroupNum() bool`

HasApGroupNum returns a boolean if a field has been set.

### GetExceeded

`func (o *ApGroupStatusOpenApiVO) GetExceeded() bool`

GetExceeded returns the Exceeded field if non-nil, zero value otherwise.

### GetExceededOk

`func (o *ApGroupStatusOpenApiVO) GetExceededOk() (*bool, bool)`

GetExceededOk returns a tuple with the Exceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceeded

`func (o *ApGroupStatusOpenApiVO) SetExceeded(v bool)`

SetExceeded sets Exceeded field to given value.

### HasExceeded

`func (o *ApGroupStatusOpenApiVO) HasExceeded() bool`

HasExceeded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


