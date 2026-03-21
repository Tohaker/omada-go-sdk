# LanNetworkSplitOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | LAN network ID | [optional] 
**Name** | **string** | LAN network name should contain 1 to 128 characters. | 
**Primary** | Pointer to **bool** | Primary | [optional] 
**Vlan** | Pointer to **int32** | Vlan should be within the range of 1-4090. | [optional] 
**VlanType** | Pointer to **int32** | When purpose is interface, VLANType should be a value as follows: 0: Single; 1: Multiple | [optional] 

## Methods

### NewLanNetworkSplitOpenApiVO

`func NewLanNetworkSplitOpenApiVO(name string, ) *LanNetworkSplitOpenApiVO`

NewLanNetworkSplitOpenApiVO instantiates a new LanNetworkSplitOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanNetworkSplitOpenApiVOWithDefaults

`func NewLanNetworkSplitOpenApiVOWithDefaults() *LanNetworkSplitOpenApiVO`

NewLanNetworkSplitOpenApiVOWithDefaults instantiates a new LanNetworkSplitOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LanNetworkSplitOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LanNetworkSplitOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LanNetworkSplitOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LanNetworkSplitOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LanNetworkSplitOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LanNetworkSplitOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LanNetworkSplitOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetPrimary

`func (o *LanNetworkSplitOpenApiVO) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *LanNetworkSplitOpenApiVO) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *LanNetworkSplitOpenApiVO) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *LanNetworkSplitOpenApiVO) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetVlan

`func (o *LanNetworkSplitOpenApiVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *LanNetworkSplitOpenApiVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *LanNetworkSplitOpenApiVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *LanNetworkSplitOpenApiVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVlanType

`func (o *LanNetworkSplitOpenApiVO) GetVlanType() int32`

GetVlanType returns the VlanType field if non-nil, zero value otherwise.

### GetVlanTypeOk

`func (o *LanNetworkSplitOpenApiVO) GetVlanTypeOk() (*int32, bool)`

GetVlanTypeOk returns a tuple with the VlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanType

`func (o *LanNetworkSplitOpenApiVO) SetVlanType(v int32)`

SetVlanType sets VlanType field to given value.

### HasVlanType

`func (o *LanNetworkSplitOpenApiVO) HasVlanType() bool`

HasVlanType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


