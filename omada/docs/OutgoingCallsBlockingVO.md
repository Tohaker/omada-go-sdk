# OutgoingCallsBlockingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrefixList** | Pointer to **[]string** | Field [prefixList] is required when the value of field [types] contains 4. | [optional] 
**Types** | Pointer to **[]int32** | Outgoing calls blocking types. 0 means mobile, 1 means landline, 2 means long distance, 3 means international, 4 means calls with specific number prefix. | [optional] 

## Methods

### NewOutgoingCallsBlockingVO

`func NewOutgoingCallsBlockingVO() *OutgoingCallsBlockingVO`

NewOutgoingCallsBlockingVO instantiates a new OutgoingCallsBlockingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutgoingCallsBlockingVOWithDefaults

`func NewOutgoingCallsBlockingVOWithDefaults() *OutgoingCallsBlockingVO`

NewOutgoingCallsBlockingVOWithDefaults instantiates a new OutgoingCallsBlockingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefixList

`func (o *OutgoingCallsBlockingVO) GetPrefixList() []string`

GetPrefixList returns the PrefixList field if non-nil, zero value otherwise.

### GetPrefixListOk

`func (o *OutgoingCallsBlockingVO) GetPrefixListOk() (*[]string, bool)`

GetPrefixListOk returns a tuple with the PrefixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixList

`func (o *OutgoingCallsBlockingVO) SetPrefixList(v []string)`

SetPrefixList sets PrefixList field to given value.

### HasPrefixList

`func (o *OutgoingCallsBlockingVO) HasPrefixList() bool`

HasPrefixList returns a boolean if a field has been set.

### GetTypes

`func (o *OutgoingCallsBlockingVO) GetTypes() []int32`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *OutgoingCallsBlockingVO) GetTypesOk() (*[]int32, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *OutgoingCallsBlockingVO) SetTypes(v []int32)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *OutgoingCallsBlockingVO) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


