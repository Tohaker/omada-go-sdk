# RetryDropRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DroppedEaps** | Pointer to [**[]DropEap**](DropEap.md) | AP list of packet loss rate tab, in descending order of average value | [optional] 
**RetryEaps** | Pointer to [**[]RetryEap**](RetryEap.md) | AP list of retransmission packet rates tab, in descending order of average | [optional] 

## Methods

### NewRetryDropRate

`func NewRetryDropRate() *RetryDropRate`

NewRetryDropRate instantiates a new RetryDropRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryDropRateWithDefaults

`func NewRetryDropRateWithDefaults() *RetryDropRate`

NewRetryDropRateWithDefaults instantiates a new RetryDropRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDroppedEaps

`func (o *RetryDropRate) GetDroppedEaps() []DropEap`

GetDroppedEaps returns the DroppedEaps field if non-nil, zero value otherwise.

### GetDroppedEapsOk

`func (o *RetryDropRate) GetDroppedEapsOk() (*[]DropEap, bool)`

GetDroppedEapsOk returns a tuple with the DroppedEaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroppedEaps

`func (o *RetryDropRate) SetDroppedEaps(v []DropEap)`

SetDroppedEaps sets DroppedEaps field to given value.

### HasDroppedEaps

`func (o *RetryDropRate) HasDroppedEaps() bool`

HasDroppedEaps returns a boolean if a field has been set.

### GetRetryEaps

`func (o *RetryDropRate) GetRetryEaps() []RetryEap`

GetRetryEaps returns the RetryEaps field if non-nil, zero value otherwise.

### GetRetryEapsOk

`func (o *RetryDropRate) GetRetryEapsOk() (*[]RetryEap, bool)`

GetRetryEapsOk returns a tuple with the RetryEaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryEaps

`func (o *RetryDropRate) SetRetryEaps(v []RetryEap)`

SetRetryEaps sets RetryEaps field to given value.

### HasRetryEaps

`func (o *RetryDropRate) HasRetryEaps() bool`

HasRetryEaps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


