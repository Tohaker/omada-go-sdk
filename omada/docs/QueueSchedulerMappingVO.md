# QueueSchedulerMappingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildIn** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Mapping** | Pointer to [**[]QueueScheduleConfigVO**](QueueScheduleConfigVO.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Types** | Pointer to **int32** |  | [optional] 

## Methods

### NewQueueSchedulerMappingVO

`func NewQueueSchedulerMappingVO() *QueueSchedulerMappingVO`

NewQueueSchedulerMappingVO instantiates a new QueueSchedulerMappingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueSchedulerMappingVOWithDefaults

`func NewQueueSchedulerMappingVOWithDefaults() *QueueSchedulerMappingVO`

NewQueueSchedulerMappingVOWithDefaults instantiates a new QueueSchedulerMappingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildIn

`func (o *QueueSchedulerMappingVO) GetBuildIn() bool`

GetBuildIn returns the BuildIn field if non-nil, zero value otherwise.

### GetBuildInOk

`func (o *QueueSchedulerMappingVO) GetBuildInOk() (*bool, bool)`

GetBuildInOk returns a tuple with the BuildIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildIn

`func (o *QueueSchedulerMappingVO) SetBuildIn(v bool)`

SetBuildIn sets BuildIn field to given value.

### HasBuildIn

`func (o *QueueSchedulerMappingVO) HasBuildIn() bool`

HasBuildIn returns a boolean if a field has been set.

### GetId

`func (o *QueueSchedulerMappingVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueueSchedulerMappingVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueueSchedulerMappingVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QueueSchedulerMappingVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMapping

`func (o *QueueSchedulerMappingVO) GetMapping() []QueueScheduleConfigVO`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *QueueSchedulerMappingVO) GetMappingOk() (*[]QueueScheduleConfigVO, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *QueueSchedulerMappingVO) SetMapping(v []QueueScheduleConfigVO)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *QueueSchedulerMappingVO) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### GetName

`func (o *QueueSchedulerMappingVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueueSchedulerMappingVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueueSchedulerMappingVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QueueSchedulerMappingVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTypes

`func (o *QueueSchedulerMappingVO) GetTypes() int32`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *QueueSchedulerMappingVO) GetTypesOk() (*int32, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *QueueSchedulerMappingVO) SetTypes(v int32)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *QueueSchedulerMappingVO) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


