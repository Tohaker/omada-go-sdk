# QueueScheduleConfigVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queue** | **int32** |  | 
**Type** | **int32** |  | 
**Weight** | Pointer to **int32** |  | [optional] 

## Methods

### NewQueueScheduleConfigVO

`func NewQueueScheduleConfigVO(queue int32, type_ int32, ) *QueueScheduleConfigVO`

NewQueueScheduleConfigVO instantiates a new QueueScheduleConfigVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueScheduleConfigVOWithDefaults

`func NewQueueScheduleConfigVOWithDefaults() *QueueScheduleConfigVO`

NewQueueScheduleConfigVOWithDefaults instantiates a new QueueScheduleConfigVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueue

`func (o *QueueScheduleConfigVO) GetQueue() int32`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *QueueScheduleConfigVO) GetQueueOk() (*int32, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *QueueScheduleConfigVO) SetQueue(v int32)`

SetQueue sets Queue field to given value.


### GetType

`func (o *QueueScheduleConfigVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueueScheduleConfigVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueueScheduleConfigVO) SetType(v int32)`

SetType sets Type field to given value.


### GetWeight

`func (o *QueueScheduleConfigVO) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *QueueScheduleConfigVO) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *QueueScheduleConfigVO) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *QueueScheduleConfigVO) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


