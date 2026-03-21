# BatchRequestEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]RequestActionEntity**](RequestActionEntity.md) | List of OpenAPIs that require batch execution. Up to 20 entries are allowed for the action list | [optional] 
**Interrupt** | Pointer to **bool** | Indicates whether to interrupt execution when encountering an error while executing openAPI, defaults to true | [optional] 

## Methods

### NewBatchRequestEntity

`func NewBatchRequestEntity() *BatchRequestEntity`

NewBatchRequestEntity instantiates a new BatchRequestEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRequestEntityWithDefaults

`func NewBatchRequestEntityWithDefaults() *BatchRequestEntity`

NewBatchRequestEntityWithDefaults instantiates a new BatchRequestEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *BatchRequestEntity) GetActions() []RequestActionEntity`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *BatchRequestEntity) GetActionsOk() (*[]RequestActionEntity, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *BatchRequestEntity) SetActions(v []RequestActionEntity)`

SetActions sets Actions field to given value.

### HasActions

`func (o *BatchRequestEntity) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetInterrupt

`func (o *BatchRequestEntity) GetInterrupt() bool`

GetInterrupt returns the Interrupt field if non-nil, zero value otherwise.

### GetInterruptOk

`func (o *BatchRequestEntity) GetInterruptOk() (*bool, bool)`

GetInterruptOk returns a tuple with the Interrupt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterrupt

`func (o *BatchRequestEntity) SetInterrupt(v bool)`

SetInterrupt sets Interrupt field to given value.

### HasInterrupt

`func (o *BatchRequestEntity) HasInterrupt() bool`

HasInterrupt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


