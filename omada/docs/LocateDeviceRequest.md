# LocateDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocateEnable** | Pointer to **bool** |  | [optional] 
**PortList** | Pointer to **[]int32** | Ports to be located | [optional] 

## Methods

### NewLocateDeviceRequest

`func NewLocateDeviceRequest() *LocateDeviceRequest`

NewLocateDeviceRequest instantiates a new LocateDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocateDeviceRequestWithDefaults

`func NewLocateDeviceRequestWithDefaults() *LocateDeviceRequest`

NewLocateDeviceRequestWithDefaults instantiates a new LocateDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocateEnable

`func (o *LocateDeviceRequest) GetLocateEnable() bool`

GetLocateEnable returns the LocateEnable field if non-nil, zero value otherwise.

### GetLocateEnableOk

`func (o *LocateDeviceRequest) GetLocateEnableOk() (*bool, bool)`

GetLocateEnableOk returns a tuple with the LocateEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocateEnable

`func (o *LocateDeviceRequest) SetLocateEnable(v bool)`

SetLocateEnable sets LocateEnable field to given value.

### HasLocateEnable

`func (o *LocateDeviceRequest) HasLocateEnable() bool`

HasLocateEnable returns a boolean if a field has been set.

### GetPortList

`func (o *LocateDeviceRequest) GetPortList() []int32`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *LocateDeviceRequest) GetPortListOk() (*[]int32, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *LocateDeviceRequest) SetPortList(v []int32)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *LocateDeviceRequest) HasPortList() bool`

HasPortList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


