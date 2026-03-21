# SwitchPortsStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortList** | **[]int32** | Port list | 
**Status** | **int32** | Status should be a value as follows: 0: off; 1: on | 

## Methods

### NewSwitchPortsStatus

`func NewSwitchPortsStatus(portList []int32, status int32, ) *SwitchPortsStatus`

NewSwitchPortsStatus instantiates a new SwitchPortsStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchPortsStatusWithDefaults

`func NewSwitchPortsStatusWithDefaults() *SwitchPortsStatus`

NewSwitchPortsStatusWithDefaults instantiates a new SwitchPortsStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortList

`func (o *SwitchPortsStatus) GetPortList() []int32`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *SwitchPortsStatus) GetPortListOk() (*[]int32, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *SwitchPortsStatus) SetPortList(v []int32)`

SetPortList sets PortList field to given value.


### GetStatus

`func (o *SwitchPortsStatus) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SwitchPortsStatus) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SwitchPortsStatus) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


