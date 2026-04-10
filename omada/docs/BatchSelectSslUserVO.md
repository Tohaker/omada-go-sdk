# BatchSelectSslUserVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | ID list of batch selected SSL VPN users. | [optional] 
**SearchKey** | Pointer to **string** | The keywords of the searchIt is effected when [selectAll] is &#39;true&#39;. | [optional] 
**SelectAll** | Pointer to **bool** | Indicates whether to select all SSL VPN users for deletion.The behavior depends on the combination with &#39;ids&#39;: - If selectAll is true and ids is not empty: perform a reverse selection (excluding the specified IDs). - If selectAll is true and ids is empty: select all SSL VPN users. - If selectAll is false and ids is not empty: select only the specified SSL VPN users in the ID list. | [optional] 

## Methods

### NewBatchSelectSslUserVO

`func NewBatchSelectSslUserVO() *BatchSelectSslUserVO`

NewBatchSelectSslUserVO instantiates a new BatchSelectSslUserVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchSelectSslUserVOWithDefaults

`func NewBatchSelectSslUserVOWithDefaults() *BatchSelectSslUserVO`

NewBatchSelectSslUserVOWithDefaults instantiates a new BatchSelectSslUserVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *BatchSelectSslUserVO) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BatchSelectSslUserVO) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BatchSelectSslUserVO) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *BatchSelectSslUserVO) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetSearchKey

`func (o *BatchSelectSslUserVO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *BatchSelectSslUserVO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *BatchSelectSslUserVO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *BatchSelectSslUserVO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSelectAll

`func (o *BatchSelectSslUserVO) GetSelectAll() bool`

GetSelectAll returns the SelectAll field if non-nil, zero value otherwise.

### GetSelectAllOk

`func (o *BatchSelectSslUserVO) GetSelectAllOk() (*bool, bool)`

GetSelectAllOk returns a tuple with the SelectAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectAll

`func (o *BatchSelectSslUserVO) SetSelectAll(v bool)`

SetSelectAll sets SelectAll field to given value.

### HasSelectAll

`func (o *BatchSelectSslUserVO) HasSelectAll() bool`

HasSelectAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


