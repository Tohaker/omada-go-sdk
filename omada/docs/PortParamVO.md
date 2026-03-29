# PortParamVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **int32** |  | 
**SwitchMac** | **string** |  | 

## Methods

### NewPortParamVO

`func NewPortParamVO(port int32, switchMac string, ) *PortParamVO`

NewPortParamVO instantiates a new PortParamVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortParamVOWithDefaults

`func NewPortParamVOWithDefaults() *PortParamVO`

NewPortParamVOWithDefaults instantiates a new PortParamVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *PortParamVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PortParamVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PortParamVO) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSwitchMac

`func (o *PortParamVO) GetSwitchMac() string`

GetSwitchMac returns the SwitchMac field if non-nil, zero value otherwise.

### GetSwitchMacOk

`func (o *PortParamVO) GetSwitchMacOk() (*string, bool)`

GetSwitchMacOk returns a tuple with the SwitchMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMac

`func (o *PortParamVO) SetSwitchMac(v string)`

SetSwitchMac sets SwitchMac field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


