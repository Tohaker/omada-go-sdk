# OswVrrpDeviceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceSeriesType** | Pointer to **int32** | Device type. | [optional] 
**Mac** | **string** | Device Mac | 
**Model** | Pointer to **string** | Device model. | [optional] 
**ModelVersion** | Pointer to **string** | Device Model Version. | [optional] 
**Name** | **string** | The name of the device | 
**Networks** | **[]string** | Store the IPv4 and IPv6 network segments corresponding to the VLAN interface, up to 2 entries are allowed for the networks. | 
**Priority** | **int32** | The device priority should be within the range of 1-254. | 
**ReducedPriority** | Pointer to **int32** | The ReducedPriority should be within the range of 1-254. | [optional] 
**StackId** | Pointer to **string** | ID of stack group | [optional] 
**TrackedInterface** | Pointer to **int32** | The TrackedInterface should be within the range of 1-4094. | [optional] 
**Type** | **string** | Device Type | 
**VlanInterface** | **int32** | The Vlan Interface should be within the range of 1-4094. | 

## Methods

### NewOswVrrpDeviceOpenApiVO

`func NewOswVrrpDeviceOpenApiVO(mac string, name string, networks []string, priority int32, type_ string, vlanInterface int32, ) *OswVrrpDeviceOpenApiVO`

NewOswVrrpDeviceOpenApiVO instantiates a new OswVrrpDeviceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswVrrpDeviceOpenApiVOWithDefaults

`func NewOswVrrpDeviceOpenApiVOWithDefaults() *OswVrrpDeviceOpenApiVO`

NewOswVrrpDeviceOpenApiVOWithDefaults instantiates a new OswVrrpDeviceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceSeriesType

`func (o *OswVrrpDeviceOpenApiVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *OswVrrpDeviceOpenApiVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *OswVrrpDeviceOpenApiVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *OswVrrpDeviceOpenApiVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetMac

`func (o *OswVrrpDeviceOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswVrrpDeviceOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswVrrpDeviceOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.


### GetModel

`func (o *OswVrrpDeviceOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OswVrrpDeviceOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OswVrrpDeviceOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OswVrrpDeviceOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OswVrrpDeviceOpenApiVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OswVrrpDeviceOpenApiVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OswVrrpDeviceOpenApiVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OswVrrpDeviceOpenApiVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *OswVrrpDeviceOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswVrrpDeviceOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswVrrpDeviceOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetNetworks

`func (o *OswVrrpDeviceOpenApiVO) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *OswVrrpDeviceOpenApiVO) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *OswVrrpDeviceOpenApiVO) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.


### GetPriority

`func (o *OswVrrpDeviceOpenApiVO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *OswVrrpDeviceOpenApiVO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *OswVrrpDeviceOpenApiVO) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetReducedPriority

`func (o *OswVrrpDeviceOpenApiVO) GetReducedPriority() int32`

GetReducedPriority returns the ReducedPriority field if non-nil, zero value otherwise.

### GetReducedPriorityOk

`func (o *OswVrrpDeviceOpenApiVO) GetReducedPriorityOk() (*int32, bool)`

GetReducedPriorityOk returns a tuple with the ReducedPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReducedPriority

`func (o *OswVrrpDeviceOpenApiVO) SetReducedPriority(v int32)`

SetReducedPriority sets ReducedPriority field to given value.

### HasReducedPriority

`func (o *OswVrrpDeviceOpenApiVO) HasReducedPriority() bool`

HasReducedPriority returns a boolean if a field has been set.

### GetStackId

`func (o *OswVrrpDeviceOpenApiVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OswVrrpDeviceOpenApiVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OswVrrpDeviceOpenApiVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OswVrrpDeviceOpenApiVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetTrackedInterface

`func (o *OswVrrpDeviceOpenApiVO) GetTrackedInterface() int32`

GetTrackedInterface returns the TrackedInterface field if non-nil, zero value otherwise.

### GetTrackedInterfaceOk

`func (o *OswVrrpDeviceOpenApiVO) GetTrackedInterfaceOk() (*int32, bool)`

GetTrackedInterfaceOk returns a tuple with the TrackedInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackedInterface

`func (o *OswVrrpDeviceOpenApiVO) SetTrackedInterface(v int32)`

SetTrackedInterface sets TrackedInterface field to given value.

### HasTrackedInterface

`func (o *OswVrrpDeviceOpenApiVO) HasTrackedInterface() bool`

HasTrackedInterface returns a boolean if a field has been set.

### GetType

`func (o *OswVrrpDeviceOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswVrrpDeviceOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswVrrpDeviceOpenApiVO) SetType(v string)`

SetType sets Type field to given value.


### GetVlanInterface

`func (o *OswVrrpDeviceOpenApiVO) GetVlanInterface() int32`

GetVlanInterface returns the VlanInterface field if non-nil, zero value otherwise.

### GetVlanInterfaceOk

`func (o *OswVrrpDeviceOpenApiVO) GetVlanInterfaceOk() (*int32, bool)`

GetVlanInterfaceOk returns a tuple with the VlanInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanInterface

`func (o *OswVrrpDeviceOpenApiVO) SetVlanInterface(v int32)`

SetVlanInterface sets VlanInterface field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


