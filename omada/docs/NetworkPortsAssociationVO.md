# NetworkPortsAssociationVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Network&#39;s ID | [optional] 
**Name** | Pointer to **string** | Network&#39;s name | [optional] 
**NativePorts** | Pointer to **[]int32** | Native ports list | [optional] 
**TaggedPorts** | Pointer to **[]int32** | Tagged ports list | [optional] 
**UntaggedPorts** | Pointer to **[]int32** | Untagged ports list | [optional] 
**Vlan** | Pointer to **int32** | VLAN ID | [optional] 

## Methods

### NewNetworkPortsAssociationVO

`func NewNetworkPortsAssociationVO() *NetworkPortsAssociationVO`

NewNetworkPortsAssociationVO instantiates a new NetworkPortsAssociationVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPortsAssociationVOWithDefaults

`func NewNetworkPortsAssociationVOWithDefaults() *NetworkPortsAssociationVO`

NewNetworkPortsAssociationVOWithDefaults instantiates a new NetworkPortsAssociationVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkPortsAssociationVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkPortsAssociationVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkPortsAssociationVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkPortsAssociationVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkPortsAssociationVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkPortsAssociationVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkPortsAssociationVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkPortsAssociationVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNativePorts

`func (o *NetworkPortsAssociationVO) GetNativePorts() []int32`

GetNativePorts returns the NativePorts field if non-nil, zero value otherwise.

### GetNativePortsOk

`func (o *NetworkPortsAssociationVO) GetNativePortsOk() (*[]int32, bool)`

GetNativePortsOk returns a tuple with the NativePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativePorts

`func (o *NetworkPortsAssociationVO) SetNativePorts(v []int32)`

SetNativePorts sets NativePorts field to given value.

### HasNativePorts

`func (o *NetworkPortsAssociationVO) HasNativePorts() bool`

HasNativePorts returns a boolean if a field has been set.

### GetTaggedPorts

`func (o *NetworkPortsAssociationVO) GetTaggedPorts() []int32`

GetTaggedPorts returns the TaggedPorts field if non-nil, zero value otherwise.

### GetTaggedPortsOk

`func (o *NetworkPortsAssociationVO) GetTaggedPortsOk() (*[]int32, bool)`

GetTaggedPortsOk returns a tuple with the TaggedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedPorts

`func (o *NetworkPortsAssociationVO) SetTaggedPorts(v []int32)`

SetTaggedPorts sets TaggedPorts field to given value.

### HasTaggedPorts

`func (o *NetworkPortsAssociationVO) HasTaggedPorts() bool`

HasTaggedPorts returns a boolean if a field has been set.

### GetUntaggedPorts

`func (o *NetworkPortsAssociationVO) GetUntaggedPorts() []int32`

GetUntaggedPorts returns the UntaggedPorts field if non-nil, zero value otherwise.

### GetUntaggedPortsOk

`func (o *NetworkPortsAssociationVO) GetUntaggedPortsOk() (*[]int32, bool)`

GetUntaggedPortsOk returns a tuple with the UntaggedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntaggedPorts

`func (o *NetworkPortsAssociationVO) SetUntaggedPorts(v []int32)`

SetUntaggedPorts sets UntaggedPorts field to given value.

### HasUntaggedPorts

`func (o *NetworkPortsAssociationVO) HasUntaggedPorts() bool`

HasUntaggedPorts returns a boolean if a field has been set.

### GetVlan

`func (o *NetworkPortsAssociationVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *NetworkPortsAssociationVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *NetworkPortsAssociationVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *NetworkPortsAssociationVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


