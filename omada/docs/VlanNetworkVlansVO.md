# VlanNetworkVlansVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VlanIds** | Pointer to **[]int32** | vlan list. The parameter is defined as an array of integers. However, the current implementation accepts non-array inputs (e.g., a single integer or a numeric string) and will treat them as a single-element array. For consistency and future compatibility, please provide &#x60;vlanIds&#x60; as an array | [optional] 

## Methods

### NewVlanNetworkVlansVO

`func NewVlanNetworkVlansVO() *VlanNetworkVlansVO`

NewVlanNetworkVlansVO instantiates a new VlanNetworkVlansVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanNetworkVlansVOWithDefaults

`func NewVlanNetworkVlansVOWithDefaults() *VlanNetworkVlansVO`

NewVlanNetworkVlansVOWithDefaults instantiates a new VlanNetworkVlansVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlanIds

`func (o *VlanNetworkVlansVO) GetVlanIds() []int32`

GetVlanIds returns the VlanIds field if non-nil, zero value otherwise.

### GetVlanIdsOk

`func (o *VlanNetworkVlansVO) GetVlanIdsOk() (*[]int32, bool)`

GetVlanIdsOk returns a tuple with the VlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIds

`func (o *VlanNetworkVlansVO) SetVlanIds(v []int32)`

SetVlanIds sets VlanIds field to given value.

### HasVlanIds

`func (o *VlanNetworkVlansVO) HasVlanIds() bool`

HasVlanIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


