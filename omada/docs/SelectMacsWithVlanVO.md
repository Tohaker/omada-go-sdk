# SelectMacsWithVlanVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Macs** | **[]string** | Device mac list for Site Setting or device template ID list for Site Template Setting | 
**Vlan** | Pointer to **int32** | vlan, only valid when vlanType is 0 (single vlan) | [optional] 
**VlanType** | **int32** | Network type, It should be a value as follows : 0:single vlan 1:multi vlan | 
**Vlans** | Pointer to **string** | vlans, only valid when vlanType is 1 (multi vlan). VLAN format: 200, 1-100. | [optional] 

## Methods

### NewSelectMacsWithVlanVO

`func NewSelectMacsWithVlanVO(macs []string, vlanType int32, ) *SelectMacsWithVlanVO`

NewSelectMacsWithVlanVO instantiates a new SelectMacsWithVlanVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectMacsWithVlanVOWithDefaults

`func NewSelectMacsWithVlanVOWithDefaults() *SelectMacsWithVlanVO`

NewSelectMacsWithVlanVOWithDefaults instantiates a new SelectMacsWithVlanVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacs

`func (o *SelectMacsWithVlanVO) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *SelectMacsWithVlanVO) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *SelectMacsWithVlanVO) SetMacs(v []string)`

SetMacs sets Macs field to given value.


### GetVlan

`func (o *SelectMacsWithVlanVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *SelectMacsWithVlanVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *SelectMacsWithVlanVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *SelectMacsWithVlanVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVlanType

`func (o *SelectMacsWithVlanVO) GetVlanType() int32`

GetVlanType returns the VlanType field if non-nil, zero value otherwise.

### GetVlanTypeOk

`func (o *SelectMacsWithVlanVO) GetVlanTypeOk() (*int32, bool)`

GetVlanTypeOk returns a tuple with the VlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanType

`func (o *SelectMacsWithVlanVO) SetVlanType(v int32)`

SetVlanType sets VlanType field to given value.


### GetVlans

`func (o *SelectMacsWithVlanVO) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *SelectMacsWithVlanVO) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *SelectMacsWithVlanVO) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *SelectMacsWithVlanVO) HasVlans() bool`

HasVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


