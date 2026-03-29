# SwitchPortsPoe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoeMode** | **int32** | Poe mode should be a value as follows: 1: on(802.3at/af); 0: off. | 
**PortList** | **[]int32** | Port ID List. | 

## Methods

### NewSwitchPortsPoe

`func NewSwitchPortsPoe(poeMode int32, portList []int32, ) *SwitchPortsPoe`

NewSwitchPortsPoe instantiates a new SwitchPortsPoe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchPortsPoeWithDefaults

`func NewSwitchPortsPoeWithDefaults() *SwitchPortsPoe`

NewSwitchPortsPoeWithDefaults instantiates a new SwitchPortsPoe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoeMode

`func (o *SwitchPortsPoe) GetPoeMode() int32`

GetPoeMode returns the PoeMode field if non-nil, zero value otherwise.

### GetPoeModeOk

`func (o *SwitchPortsPoe) GetPoeModeOk() (*int32, bool)`

GetPoeModeOk returns a tuple with the PoeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeMode

`func (o *SwitchPortsPoe) SetPoeMode(v int32)`

SetPoeMode sets PoeMode field to given value.


### GetPortList

`func (o *SwitchPortsPoe) GetPortList() []int32`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *SwitchPortsPoe) GetPortListOk() (*[]int32, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *SwitchPortsPoe) SetPortList(v []int32)`

SetPortList sets PortList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


