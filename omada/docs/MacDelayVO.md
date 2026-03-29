# MacDelayVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelayTime** | Pointer to **int32** | DelayTime should be within the range of 0–60 | [optional] 
**Enable** | Pointer to **bool** | Enable | [optional] 

## Methods

### NewMacDelayVO

`func NewMacDelayVO() *MacDelayVO`

NewMacDelayVO instantiates a new MacDelayVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacDelayVOWithDefaults

`func NewMacDelayVOWithDefaults() *MacDelayVO`

NewMacDelayVOWithDefaults instantiates a new MacDelayVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelayTime

`func (o *MacDelayVO) GetDelayTime() int32`

GetDelayTime returns the DelayTime field if non-nil, zero value otherwise.

### GetDelayTimeOk

`func (o *MacDelayVO) GetDelayTimeOk() (*int32, bool)`

GetDelayTimeOk returns a tuple with the DelayTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayTime

`func (o *MacDelayVO) SetDelayTime(v int32)`

SetDelayTime sets DelayTime field to given value.

### HasDelayTime

`func (o *MacDelayVO) HasDelayTime() bool`

HasDelayTime returns a boolean if a field has been set.

### GetEnable

`func (o *MacDelayVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *MacDelayVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *MacDelayVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *MacDelayVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


