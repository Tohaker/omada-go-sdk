# BatchIsolateInterfaceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | A list of network Id which is selected to isolate or de-isolate. | [optional] 
**Isolation** | Pointer to **bool** | The isolation status after batch isolate or de-isolate. | [optional] 
**SelectAll** | Pointer to **bool** | Whether select all networks. | [optional] 

## Methods

### NewBatchIsolateInterfaceOpenApiVO

`func NewBatchIsolateInterfaceOpenApiVO() *BatchIsolateInterfaceOpenApiVO`

NewBatchIsolateInterfaceOpenApiVO instantiates a new BatchIsolateInterfaceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchIsolateInterfaceOpenApiVOWithDefaults

`func NewBatchIsolateInterfaceOpenApiVOWithDefaults() *BatchIsolateInterfaceOpenApiVO`

NewBatchIsolateInterfaceOpenApiVOWithDefaults instantiates a new BatchIsolateInterfaceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *BatchIsolateInterfaceOpenApiVO) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BatchIsolateInterfaceOpenApiVO) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BatchIsolateInterfaceOpenApiVO) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *BatchIsolateInterfaceOpenApiVO) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetIsolation

`func (o *BatchIsolateInterfaceOpenApiVO) GetIsolation() bool`

GetIsolation returns the Isolation field if non-nil, zero value otherwise.

### GetIsolationOk

`func (o *BatchIsolateInterfaceOpenApiVO) GetIsolationOk() (*bool, bool)`

GetIsolationOk returns a tuple with the Isolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolation

`func (o *BatchIsolateInterfaceOpenApiVO) SetIsolation(v bool)`

SetIsolation sets Isolation field to given value.

### HasIsolation

`func (o *BatchIsolateInterfaceOpenApiVO) HasIsolation() bool`

HasIsolation returns a boolean if a field has been set.

### GetSelectAll

`func (o *BatchIsolateInterfaceOpenApiVO) GetSelectAll() bool`

GetSelectAll returns the SelectAll field if non-nil, zero value otherwise.

### GetSelectAllOk

`func (o *BatchIsolateInterfaceOpenApiVO) GetSelectAllOk() (*bool, bool)`

GetSelectAllOk returns a tuple with the SelectAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectAll

`func (o *BatchIsolateInterfaceOpenApiVO) SetSelectAll(v bool)`

SetSelectAll sets SelectAll field to given value.

### HasSelectAll

`func (o *BatchIsolateInterfaceOpenApiVO) HasSelectAll() bool`

HasSelectAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


