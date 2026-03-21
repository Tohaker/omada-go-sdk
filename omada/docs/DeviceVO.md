# DeviceVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | The mac of the general device. | [optional] 
**Ports** | Pointer to [**[]PortVO**](PortVO.md) | The ports selected. | [optional] 
**StackId** | Pointer to **string** | The stack of the stack device. | [optional] 

## Methods

### NewDeviceVO

`func NewDeviceVO() *DeviceVO`

NewDeviceVO instantiates a new DeviceVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceVOWithDefaults

`func NewDeviceVOWithDefaults() *DeviceVO`

NewDeviceVOWithDefaults instantiates a new DeviceVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *DeviceVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DeviceVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DeviceVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *DeviceVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetPorts

`func (o *DeviceVO) GetPorts() []PortVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *DeviceVO) GetPortsOk() (*[]PortVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *DeviceVO) SetPorts(v []PortVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *DeviceVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetStackId

`func (o *DeviceVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *DeviceVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *DeviceVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *DeviceVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


