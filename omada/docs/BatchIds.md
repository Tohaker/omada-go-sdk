# BatchIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | When selectType is set to all, the ids do not need to be passed and all entries are processed, when selectType is set to include, the id entries contained in the ids are processed, when selectType is set to exclude, the id entries that are not contained in the ids are processed | 
**SearchKey** | Pointer to **string** | Fuzzy query parameters, support field: voucher code | [optional] 
**SelectType** | Pointer to **string** | SelectType all, include or exclude | [optional] 

## Methods

### NewBatchIds

`func NewBatchIds(ids []string, ) *BatchIds`

NewBatchIds instantiates a new BatchIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchIdsWithDefaults

`func NewBatchIdsWithDefaults() *BatchIds`

NewBatchIdsWithDefaults instantiates a new BatchIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *BatchIds) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BatchIds) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BatchIds) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetSearchKey

`func (o *BatchIds) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *BatchIds) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *BatchIds) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *BatchIds) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSelectType

`func (o *BatchIds) GetSelectType() string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *BatchIds) GetSelectTypeOk() (*string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *BatchIds) SetSelectType(v string)`

SetSelectType sets SelectType field to given value.

### HasSelectType

`func (o *BatchIds) HasSelectType() bool`

HasSelectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


