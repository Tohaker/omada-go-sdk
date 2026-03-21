# SwitchPortStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** | Status should be a value as follows: 0: off; 1: on | 

## Methods

### NewSwitchPortStatus

`func NewSwitchPortStatus(status int32, ) *SwitchPortStatus`

NewSwitchPortStatus instantiates a new SwitchPortStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchPortStatusWithDefaults

`func NewSwitchPortStatusWithDefaults() *SwitchPortStatus`

NewSwitchPortStatusWithDefaults instantiates a new SwitchPortStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SwitchPortStatus) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SwitchPortStatus) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SwitchPortStatus) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


