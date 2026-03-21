# OutboxMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Message content. | [optional] 
**Id** | Pointer to **int32** | Item ID. | [optional] 
**Receiver** | Pointer to **string** | Receiver. | [optional] 
**Time** | Pointer to **int64** | Time. | [optional] 

## Methods

### NewOutboxMessage

`func NewOutboxMessage() *OutboxMessage`

NewOutboxMessage instantiates a new OutboxMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboxMessageWithDefaults

`func NewOutboxMessageWithDefaults() *OutboxMessage`

NewOutboxMessageWithDefaults instantiates a new OutboxMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *OutboxMessage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OutboxMessage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OutboxMessage) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *OutboxMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetId

`func (o *OutboxMessage) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutboxMessage) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutboxMessage) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OutboxMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReceiver

`func (o *OutboxMessage) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *OutboxMessage) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *OutboxMessage) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *OutboxMessage) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetTime

`func (o *OutboxMessage) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *OutboxMessage) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *OutboxMessage) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *OutboxMessage) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


