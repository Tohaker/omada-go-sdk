# BatchAutoSelectWanPortReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberList** | Pointer to [**[]AutoSelectWanPortReq**](AutoSelectWanPortReq.md) | A list of the SD-WAN devices which use auto select. | [optional] 

## Methods

### NewBatchAutoSelectWanPortReq

`func NewBatchAutoSelectWanPortReq() *BatchAutoSelectWanPortReq`

NewBatchAutoSelectWanPortReq instantiates a new BatchAutoSelectWanPortReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchAutoSelectWanPortReqWithDefaults

`func NewBatchAutoSelectWanPortReqWithDefaults() *BatchAutoSelectWanPortReq`

NewBatchAutoSelectWanPortReqWithDefaults instantiates a new BatchAutoSelectWanPortReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberList

`func (o *BatchAutoSelectWanPortReq) GetMemberList() []AutoSelectWanPortReq`

GetMemberList returns the MemberList field if non-nil, zero value otherwise.

### GetMemberListOk

`func (o *BatchAutoSelectWanPortReq) GetMemberListOk() (*[]AutoSelectWanPortReq, bool)`

GetMemberListOk returns a tuple with the MemberList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberList

`func (o *BatchAutoSelectWanPortReq) SetMemberList(v []AutoSelectWanPortReq)`

SetMemberList sets MemberList field to given value.

### HasMemberList

`func (o *BatchAutoSelectWanPortReq) HasMemberList() bool`

HasMemberList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


