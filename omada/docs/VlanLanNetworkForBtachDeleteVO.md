# VlanLanNetworkForBtachDeleteVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | network id | [optional] 
**Name** | Pointer to **string** | network name | [optional] 
**Purpose** | **string** | LAN network purpose, 0: VLAN, 1: interface | 
**Vlan** | Pointer to **int32** | vlan id | [optional] 

## Methods

### NewVlanLanNetworkForBtachDeleteVO

`func NewVlanLanNetworkForBtachDeleteVO(purpose string, ) *VlanLanNetworkForBtachDeleteVO`

NewVlanLanNetworkForBtachDeleteVO instantiates a new VlanLanNetworkForBtachDeleteVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanLanNetworkForBtachDeleteVOWithDefaults

`func NewVlanLanNetworkForBtachDeleteVOWithDefaults() *VlanLanNetworkForBtachDeleteVO`

NewVlanLanNetworkForBtachDeleteVOWithDefaults instantiates a new VlanLanNetworkForBtachDeleteVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VlanLanNetworkForBtachDeleteVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VlanLanNetworkForBtachDeleteVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VlanLanNetworkForBtachDeleteVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VlanLanNetworkForBtachDeleteVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VlanLanNetworkForBtachDeleteVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VlanLanNetworkForBtachDeleteVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VlanLanNetworkForBtachDeleteVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VlanLanNetworkForBtachDeleteVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPurpose

`func (o *VlanLanNetworkForBtachDeleteVO) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *VlanLanNetworkForBtachDeleteVO) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *VlanLanNetworkForBtachDeleteVO) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.


### GetVlan

`func (o *VlanLanNetworkForBtachDeleteVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *VlanLanNetworkForBtachDeleteVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *VlanLanNetworkForBtachDeleteVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *VlanLanNetworkForBtachDeleteVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


