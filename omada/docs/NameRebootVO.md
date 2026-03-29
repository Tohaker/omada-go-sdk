# NameRebootVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Switch name | [optional] 
**RebootTimes** | Pointer to **int32** | Reboot times | [optional] 

## Methods

### NewNameRebootVO

`func NewNameRebootVO() *NameRebootVO`

NewNameRebootVO instantiates a new NameRebootVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNameRebootVOWithDefaults

`func NewNameRebootVOWithDefaults() *NameRebootVO`

NewNameRebootVOWithDefaults instantiates a new NameRebootVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NameRebootVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NameRebootVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NameRebootVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NameRebootVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRebootTimes

`func (o *NameRebootVO) GetRebootTimes() int32`

GetRebootTimes returns the RebootTimes field if non-nil, zero value otherwise.

### GetRebootTimesOk

`func (o *NameRebootVO) GetRebootTimesOk() (*int32, bool)`

GetRebootTimesOk returns a tuple with the RebootTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootTimes

`func (o *NameRebootVO) SetRebootTimes(v int32)`

SetRebootTimes sets RebootTimes field to given value.

### HasRebootTimes

`func (o *NameRebootVO) HasRebootTimes() bool`

HasRebootTimes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


