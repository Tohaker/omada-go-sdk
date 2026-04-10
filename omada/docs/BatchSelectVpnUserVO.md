# BatchSelectVpnUserVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | ID list of batch selected VPN users. | [optional] 
**Protocol** | Pointer to **int32** | Protocol should be a value as follows: 0: L2TP or PPTP; 1: OpenVPN, 2: SSL VPN. | [optional] 
**SearchKey** | Pointer to **string** | The keywords of the searchIt is effected when [selectAll] is &#39;true&#39;. | [optional] 
**SelectAll** | Pointer to **bool** | Indicates whether to select all VPN users for deletion. The behavior depends on the combination with &#39;ids&#39;: 1: If selectAll is true and ids is not empty: perform a reverse selection (excluding the specified IDs). 2: If selectAll is true and ids is empty: select all VPN users. 3: If selectAll is false and ids is not empty: select only the specified VPN users in the ID list. | [optional] 

## Methods

### NewBatchSelectVpnUserVO

`func NewBatchSelectVpnUserVO() *BatchSelectVpnUserVO`

NewBatchSelectVpnUserVO instantiates a new BatchSelectVpnUserVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchSelectVpnUserVOWithDefaults

`func NewBatchSelectVpnUserVOWithDefaults() *BatchSelectVpnUserVO`

NewBatchSelectVpnUserVOWithDefaults instantiates a new BatchSelectVpnUserVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *BatchSelectVpnUserVO) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BatchSelectVpnUserVO) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BatchSelectVpnUserVO) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *BatchSelectVpnUserVO) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetProtocol

`func (o *BatchSelectVpnUserVO) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BatchSelectVpnUserVO) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BatchSelectVpnUserVO) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *BatchSelectVpnUserVO) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSearchKey

`func (o *BatchSelectVpnUserVO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *BatchSelectVpnUserVO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *BatchSelectVpnUserVO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *BatchSelectVpnUserVO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSelectAll

`func (o *BatchSelectVpnUserVO) GetSelectAll() bool`

GetSelectAll returns the SelectAll field if non-nil, zero value otherwise.

### GetSelectAllOk

`func (o *BatchSelectVpnUserVO) GetSelectAllOk() (*bool, bool)`

GetSelectAllOk returns a tuple with the SelectAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectAll

`func (o *BatchSelectVpnUserVO) SetSelectAll(v bool)`

SetSelectAll sets SelectAll field to given value.

### HasSelectAll

`func (o *BatchSelectVpnUserVO) HasSelectAll() bool`

HasSelectAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


