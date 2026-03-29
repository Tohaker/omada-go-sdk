# SwitchPortPoe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoeMode** | **int32** | Poe mode should be a value as follows: 1: on(802.3at/af); 0: off. | 

## Methods

### NewSwitchPortPoe

`func NewSwitchPortPoe(poeMode int32, ) *SwitchPortPoe`

NewSwitchPortPoe instantiates a new SwitchPortPoe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchPortPoeWithDefaults

`func NewSwitchPortPoeWithDefaults() *SwitchPortPoe`

NewSwitchPortPoeWithDefaults instantiates a new SwitchPortPoe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoeMode

`func (o *SwitchPortPoe) GetPoeMode() int32`

GetPoeMode returns the PoeMode field if non-nil, zero value otherwise.

### GetPoeModeOk

`func (o *SwitchPortPoe) GetPoeModeOk() (*int32, bool)`

GetPoeModeOk returns a tuple with the PoeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeMode

`func (o *SwitchPortPoe) SetPoeMode(v int32)`

SetPoeMode sets PoeMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


