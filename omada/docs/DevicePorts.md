# DevicePorts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortIds** | Pointer to **[]int32** | List of device ports supporting port schedule. | [optional] 
**PortNum** | Pointer to **int32** | Device total ports number. | [optional] 

## Methods

### NewDevicePorts

`func NewDevicePorts() *DevicePorts`

NewDevicePorts instantiates a new DevicePorts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicePortsWithDefaults

`func NewDevicePortsWithDefaults() *DevicePorts`

NewDevicePortsWithDefaults instantiates a new DevicePorts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortIds

`func (o *DevicePorts) GetPortIds() []int32`

GetPortIds returns the PortIds field if non-nil, zero value otherwise.

### GetPortIdsOk

`func (o *DevicePorts) GetPortIdsOk() (*[]int32, bool)`

GetPortIdsOk returns a tuple with the PortIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIds

`func (o *DevicePorts) SetPortIds(v []int32)`

SetPortIds sets PortIds field to given value.

### HasPortIds

`func (o *DevicePorts) HasPortIds() bool`

HasPortIds returns a boolean if a field has been set.

### GetPortNum

`func (o *DevicePorts) GetPortNum() int32`

GetPortNum returns the PortNum field if non-nil, zero value otherwise.

### GetPortNumOk

`func (o *DevicePorts) GetPortNumOk() (*int32, bool)`

GetPortNumOk returns a tuple with the PortNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNum

`func (o *DevicePorts) SetPortNum(v int32)`

SetPortNum sets PortNum field to given value.

### HasPortNum

`func (o *DevicePorts) HasPortNum() bool`

HasPortNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


