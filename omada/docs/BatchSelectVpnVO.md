# BatchSelectVpnVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | ID list of batch selected VPNs. | [optional] 
**SearchKey** | Pointer to **string** | The keywords of the searchIt is effected when [selectAll] is &#39;true&#39;. | [optional] 
**SelectAll** | Pointer to **bool** | Indicates whether to select all VPNs for deletion. The behavior depends on the combination with &#39;ids&#39;: 1: If selectAll is true and ids is not empty: perform a reverse selection (excluding the specified IDs). 2: If selectAll is true and ids is empty: select all VPNs. 3: If selectAll is false and ids is not empty: select only the specified VPNs in the ID list. | [optional] 
**Usage** | Pointer to **int32** | Parameter [usage] should be a value as follows: 0, Server; 1, Client; 2, Site-To-Site. | [optional] 

## Methods

### NewBatchSelectVpnVO

`func NewBatchSelectVpnVO() *BatchSelectVpnVO`

NewBatchSelectVpnVO instantiates a new BatchSelectVpnVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchSelectVpnVOWithDefaults

`func NewBatchSelectVpnVOWithDefaults() *BatchSelectVpnVO`

NewBatchSelectVpnVOWithDefaults instantiates a new BatchSelectVpnVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *BatchSelectVpnVO) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BatchSelectVpnVO) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BatchSelectVpnVO) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *BatchSelectVpnVO) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetSearchKey

`func (o *BatchSelectVpnVO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *BatchSelectVpnVO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *BatchSelectVpnVO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *BatchSelectVpnVO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSelectAll

`func (o *BatchSelectVpnVO) GetSelectAll() bool`

GetSelectAll returns the SelectAll field if non-nil, zero value otherwise.

### GetSelectAllOk

`func (o *BatchSelectVpnVO) GetSelectAllOk() (*bool, bool)`

GetSelectAllOk returns a tuple with the SelectAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectAll

`func (o *BatchSelectVpnVO) SetSelectAll(v bool)`

SetSelectAll sets SelectAll field to given value.

### HasSelectAll

`func (o *BatchSelectVpnVO) HasSelectAll() bool`

HasSelectAll returns a boolean if a field has been set.

### GetUsage

`func (o *BatchSelectVpnVO) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *BatchSelectVpnVO) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *BatchSelectVpnVO) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *BatchSelectVpnVO) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


