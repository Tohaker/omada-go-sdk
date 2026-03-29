# OswVrrpDeviceConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | **string** | Device Mac | 
**Name** | **string** | The name of the device | 
**Networks** | **[]string** | Store the IPv4 and IPv6 network segments corresponding to the VLAN interface, up to 2 entries are allowed for the networks. | 
**Priority** | **int32** | The device priority should be within the range of 1-254. | 
**ReducedPriority** | Pointer to **int32** | The ReducedPriority should be within the range of 1-254. | [optional] 
**StackId** | Pointer to **string** | ID of stack group | [optional] 
**TrackedInterface** | Pointer to **int32** | The TrackedInterface should be within the range of 1-4094. | [optional] 
**Type** | **string** | Device Type | 
**VlanInterface** | **int32** | The Vlan Interface should be within the range of 1-4094. | 

## Methods

### NewOswVrrpDeviceConfigOpenApiVO

`func NewOswVrrpDeviceConfigOpenApiVO(mac string, name string, networks []string, priority int32, type_ string, vlanInterface int32, ) *OswVrrpDeviceConfigOpenApiVO`

NewOswVrrpDeviceConfigOpenApiVO instantiates a new OswVrrpDeviceConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswVrrpDeviceConfigOpenApiVOWithDefaults

`func NewOswVrrpDeviceConfigOpenApiVOWithDefaults() *OswVrrpDeviceConfigOpenApiVO`

NewOswVrrpDeviceConfigOpenApiVOWithDefaults instantiates a new OswVrrpDeviceConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *OswVrrpDeviceConfigOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswVrrpDeviceConfigOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswVrrpDeviceConfigOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.


### GetName

`func (o *OswVrrpDeviceConfigOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswVrrpDeviceConfigOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswVrrpDeviceConfigOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetNetworks

`func (o *OswVrrpDeviceConfigOpenApiVO) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *OswVrrpDeviceConfigOpenApiVO) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *OswVrrpDeviceConfigOpenApiVO) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.


### GetPriority

`func (o *OswVrrpDeviceConfigOpenApiVO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *OswVrrpDeviceConfigOpenApiVO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *OswVrrpDeviceConfigOpenApiVO) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetReducedPriority

`func (o *OswVrrpDeviceConfigOpenApiVO) GetReducedPriority() int32`

GetReducedPriority returns the ReducedPriority field if non-nil, zero value otherwise.

### GetReducedPriorityOk

`func (o *OswVrrpDeviceConfigOpenApiVO) GetReducedPriorityOk() (*int32, bool)`

GetReducedPriorityOk returns a tuple with the ReducedPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReducedPriority

`func (o *OswVrrpDeviceConfigOpenApiVO) SetReducedPriority(v int32)`

SetReducedPriority sets ReducedPriority field to given value.

### HasReducedPriority

`func (o *OswVrrpDeviceConfigOpenApiVO) HasReducedPriority() bool`

HasReducedPriority returns a boolean if a field has been set.

### GetStackId

`func (o *OswVrrpDeviceConfigOpenApiVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OswVrrpDeviceConfigOpenApiVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OswVrrpDeviceConfigOpenApiVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OswVrrpDeviceConfigOpenApiVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetTrackedInterface

`func (o *OswVrrpDeviceConfigOpenApiVO) GetTrackedInterface() int32`

GetTrackedInterface returns the TrackedInterface field if non-nil, zero value otherwise.

### GetTrackedInterfaceOk

`func (o *OswVrrpDeviceConfigOpenApiVO) GetTrackedInterfaceOk() (*int32, bool)`

GetTrackedInterfaceOk returns a tuple with the TrackedInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackedInterface

`func (o *OswVrrpDeviceConfigOpenApiVO) SetTrackedInterface(v int32)`

SetTrackedInterface sets TrackedInterface field to given value.

### HasTrackedInterface

`func (o *OswVrrpDeviceConfigOpenApiVO) HasTrackedInterface() bool`

HasTrackedInterface returns a boolean if a field has been set.

### GetType

`func (o *OswVrrpDeviceConfigOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswVrrpDeviceConfigOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswVrrpDeviceConfigOpenApiVO) SetType(v string)`

SetType sets Type field to given value.


### GetVlanInterface

`func (o *OswVrrpDeviceConfigOpenApiVO) GetVlanInterface() int32`

GetVlanInterface returns the VlanInterface field if non-nil, zero value otherwise.

### GetVlanInterfaceOk

`func (o *OswVrrpDeviceConfigOpenApiVO) GetVlanInterfaceOk() (*int32, bool)`

GetVlanInterfaceOk returns a tuple with the VlanInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanInterface

`func (o *OswVrrpDeviceConfigOpenApiVO) SetVlanInterface(v int32)`

SetVlanInterface sets VlanInterface field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


