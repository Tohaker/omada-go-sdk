# SwitchMultiPortName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Port name should contain 1 to 128 characters. | 
**Port** | **int32** | Port ID | 

## Methods

### NewSwitchMultiPortName

`func NewSwitchMultiPortName(name string, port int32, ) *SwitchMultiPortName`

NewSwitchMultiPortName instantiates a new SwitchMultiPortName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchMultiPortNameWithDefaults

`func NewSwitchMultiPortNameWithDefaults() *SwitchMultiPortName`

NewSwitchMultiPortNameWithDefaults instantiates a new SwitchMultiPortName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SwitchMultiPortName) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SwitchMultiPortName) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SwitchMultiPortName) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *SwitchMultiPortName) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SwitchMultiPortName) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SwitchMultiPortName) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


