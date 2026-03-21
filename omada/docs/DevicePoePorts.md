# DevicePoePorts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoePortIds** | Pointer to **[]int32** | List of ports supporting Poe. | [optional] 
**PortNum** | Pointer to **int32** | Total port num. | [optional] 

## Methods

### NewDevicePoePorts

`func NewDevicePoePorts() *DevicePoePorts`

NewDevicePoePorts instantiates a new DevicePoePorts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicePoePortsWithDefaults

`func NewDevicePoePortsWithDefaults() *DevicePoePorts`

NewDevicePoePortsWithDefaults instantiates a new DevicePoePorts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoePortIds

`func (o *DevicePoePorts) GetPoePortIds() []int32`

GetPoePortIds returns the PoePortIds field if non-nil, zero value otherwise.

### GetPoePortIdsOk

`func (o *DevicePoePorts) GetPoePortIdsOk() (*[]int32, bool)`

GetPoePortIdsOk returns a tuple with the PoePortIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePortIds

`func (o *DevicePoePorts) SetPoePortIds(v []int32)`

SetPoePortIds sets PoePortIds field to given value.

### HasPoePortIds

`func (o *DevicePoePorts) HasPoePortIds() bool`

HasPoePortIds returns a boolean if a field has been set.

### GetPortNum

`func (o *DevicePoePorts) GetPortNum() int32`

GetPortNum returns the PortNum field if non-nil, zero value otherwise.

### GetPortNumOk

`func (o *DevicePoePorts) GetPortNumOk() (*int32, bool)`

GetPortNumOk returns a tuple with the PortNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNum

`func (o *DevicePoePorts) SetPortNum(v int32)`

SetPortNum sets PortNum field to given value.

### HasPortNum

`func (o *DevicePoePorts) HasPortNum() bool`

HasPortNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


