# IncomingCallsBlockingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberList** | Pointer to **[]string** | Field [numberList] is required when the value of field [types] contains 0. | [optional] 
**Types** | Pointer to **[]int32** | Incoming calls blocking types. 0 means specific number, 1 means anonymous number. | [optional] 

## Methods

### NewIncomingCallsBlockingVO

`func NewIncomingCallsBlockingVO() *IncomingCallsBlockingVO`

NewIncomingCallsBlockingVO instantiates a new IncomingCallsBlockingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomingCallsBlockingVOWithDefaults

`func NewIncomingCallsBlockingVOWithDefaults() *IncomingCallsBlockingVO`

NewIncomingCallsBlockingVOWithDefaults instantiates a new IncomingCallsBlockingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberList

`func (o *IncomingCallsBlockingVO) GetNumberList() []string`

GetNumberList returns the NumberList field if non-nil, zero value otherwise.

### GetNumberListOk

`func (o *IncomingCallsBlockingVO) GetNumberListOk() (*[]string, bool)`

GetNumberListOk returns a tuple with the NumberList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberList

`func (o *IncomingCallsBlockingVO) SetNumberList(v []string)`

SetNumberList sets NumberList field to given value.

### HasNumberList

`func (o *IncomingCallsBlockingVO) HasNumberList() bool`

HasNumberList returns a boolean if a field has been set.

### GetTypes

`func (o *IncomingCallsBlockingVO) GetTypes() []int32`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *IncomingCallsBlockingVO) GetTypesOk() (*[]int32, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *IncomingCallsBlockingVO) SetTypes(v []int32)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *IncomingCallsBlockingVO) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


