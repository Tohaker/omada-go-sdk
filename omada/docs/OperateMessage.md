# OperateMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]int32** | Message ID list. When parameter [operation] is 1 or 2, parameter [ids] should not be null. | [optional] 
**Operation** | **int32** | Operation Type. 1:delete; 2:read; 3:clear. | 
**SimCard** | Pointer to **int32** | When the device supports Dual-SIM card, parameter [simCard] should not be null.1: SIM1; 2: SIM2. | [optional] 
**Type** | **int32** | Message Type. 0: inbox; 1: outbox. | 

## Methods

### NewOperateMessage

`func NewOperateMessage(operation int32, type_ int32, ) *OperateMessage`

NewOperateMessage instantiates a new OperateMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperateMessageWithDefaults

`func NewOperateMessageWithDefaults() *OperateMessage`

NewOperateMessageWithDefaults instantiates a new OperateMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *OperateMessage) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *OperateMessage) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *OperateMessage) SetIds(v []int32)`

SetIds sets Ids field to given value.

### HasIds

`func (o *OperateMessage) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetOperation

`func (o *OperateMessage) GetOperation() int32`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *OperateMessage) GetOperationOk() (*int32, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *OperateMessage) SetOperation(v int32)`

SetOperation sets Operation field to given value.


### GetSimCard

`func (o *OperateMessage) GetSimCard() int32`

GetSimCard returns the SimCard field if non-nil, zero value otherwise.

### GetSimCardOk

`func (o *OperateMessage) GetSimCardOk() (*int32, bool)`

GetSimCardOk returns a tuple with the SimCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCard

`func (o *OperateMessage) SetSimCard(v int32)`

SetSimCard sets SimCard field to given value.

### HasSimCard

`func (o *OperateMessage) HasSimCard() bool`

HasSimCard returns a boolean if a field has been set.

### GetType

`func (o *OperateMessage) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OperateMessage) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OperateMessage) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


