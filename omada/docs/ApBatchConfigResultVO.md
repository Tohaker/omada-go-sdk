# ApBatchConfigResultVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigResultList** | Pointer to [**[]ApConfigResultVO**](ApConfigResultVO.md) | Configuration results for devices with partial or complete configuration failures. | [optional] 

## Methods

### NewApBatchConfigResultVO

`func NewApBatchConfigResultVO() *ApBatchConfigResultVO`

NewApBatchConfigResultVO instantiates a new ApBatchConfigResultVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApBatchConfigResultVOWithDefaults

`func NewApBatchConfigResultVOWithDefaults() *ApBatchConfigResultVO`

NewApBatchConfigResultVOWithDefaults instantiates a new ApBatchConfigResultVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigResultList

`func (o *ApBatchConfigResultVO) GetConfigResultList() []ApConfigResultVO`

GetConfigResultList returns the ConfigResultList field if non-nil, zero value otherwise.

### GetConfigResultListOk

`func (o *ApBatchConfigResultVO) GetConfigResultListOk() (*[]ApConfigResultVO, bool)`

GetConfigResultListOk returns a tuple with the ConfigResultList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResultList

`func (o *ApBatchConfigResultVO) SetConfigResultList(v []ApConfigResultVO)`

SetConfigResultList sets ConfigResultList field to given value.

### HasConfigResultList

`func (o *ApBatchConfigResultVO) HasConfigResultList() bool`

HasConfigResultList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


