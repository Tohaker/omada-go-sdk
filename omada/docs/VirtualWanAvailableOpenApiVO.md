# VirtualWanAvailableOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Proto** | Pointer to **string** | Virtual WAN IPv4 proto. | [optional] 
**Name** | Pointer to **string** | Virtual WAN name. | [optional] 
**PhysicalWanId** | Pointer to **string** | Physical WAN ID. | [optional] 
**Type** | Pointer to **int32** | Physical WAN port type, 0: WAN; 1:WAN/LAN; 2:LAN; 3:SFP WAN; 4:USB LTE WAN; 5: LTE WAN; 6:DSL WAN; | [optional] 
**VirtualWanId** | Pointer to **string** | Virtual WAN ID. | [optional] 

## Methods

### NewVirtualWanAvailableOpenApiVO

`func NewVirtualWanAvailableOpenApiVO() *VirtualWanAvailableOpenApiVO`

NewVirtualWanAvailableOpenApiVO instantiates a new VirtualWanAvailableOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualWanAvailableOpenApiVOWithDefaults

`func NewVirtualWanAvailableOpenApiVOWithDefaults() *VirtualWanAvailableOpenApiVO`

NewVirtualWanAvailableOpenApiVOWithDefaults instantiates a new VirtualWanAvailableOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Proto

`func (o *VirtualWanAvailableOpenApiVO) GetIpv4Proto() string`

GetIpv4Proto returns the Ipv4Proto field if non-nil, zero value otherwise.

### GetIpv4ProtoOk

`func (o *VirtualWanAvailableOpenApiVO) GetIpv4ProtoOk() (*string, bool)`

GetIpv4ProtoOk returns a tuple with the Ipv4Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Proto

`func (o *VirtualWanAvailableOpenApiVO) SetIpv4Proto(v string)`

SetIpv4Proto sets Ipv4Proto field to given value.

### HasIpv4Proto

`func (o *VirtualWanAvailableOpenApiVO) HasIpv4Proto() bool`

HasIpv4Proto returns a boolean if a field has been set.

### GetName

`func (o *VirtualWanAvailableOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualWanAvailableOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualWanAvailableOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualWanAvailableOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhysicalWanId

`func (o *VirtualWanAvailableOpenApiVO) GetPhysicalWanId() string`

GetPhysicalWanId returns the PhysicalWanId field if non-nil, zero value otherwise.

### GetPhysicalWanIdOk

`func (o *VirtualWanAvailableOpenApiVO) GetPhysicalWanIdOk() (*string, bool)`

GetPhysicalWanIdOk returns a tuple with the PhysicalWanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalWanId

`func (o *VirtualWanAvailableOpenApiVO) SetPhysicalWanId(v string)`

SetPhysicalWanId sets PhysicalWanId field to given value.

### HasPhysicalWanId

`func (o *VirtualWanAvailableOpenApiVO) HasPhysicalWanId() bool`

HasPhysicalWanId returns a boolean if a field has been set.

### GetType

`func (o *VirtualWanAvailableOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualWanAvailableOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualWanAvailableOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *VirtualWanAvailableOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVirtualWanId

`func (o *VirtualWanAvailableOpenApiVO) GetVirtualWanId() string`

GetVirtualWanId returns the VirtualWanId field if non-nil, zero value otherwise.

### GetVirtualWanIdOk

`func (o *VirtualWanAvailableOpenApiVO) GetVirtualWanIdOk() (*string, bool)`

GetVirtualWanIdOk returns a tuple with the VirtualWanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanId

`func (o *VirtualWanAvailableOpenApiVO) SetVirtualWanId(v string)`

SetVirtualWanId sets VirtualWanId field to given value.

### HasVirtualWanId

`func (o *VirtualWanAvailableOpenApiVO) HasVirtualWanId() bool`

HasVirtualWanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


