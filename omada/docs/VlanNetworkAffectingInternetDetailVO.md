# VlanNetworkAffectingInternetDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedPorts** | Pointer to **[]int32** | Affected Internet gateway port list. | [optional] 
**InternetVO** | Pointer to [**CheckWanLanStatusVO**](CheckWanLanStatusVO.md) |  | [optional] 

## Methods

### NewVlanNetworkAffectingInternetDetailVO

`func NewVlanNetworkAffectingInternetDetailVO() *VlanNetworkAffectingInternetDetailVO`

NewVlanNetworkAffectingInternetDetailVO instantiates a new VlanNetworkAffectingInternetDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanNetworkAffectingInternetDetailVOWithDefaults

`func NewVlanNetworkAffectingInternetDetailVOWithDefaults() *VlanNetworkAffectingInternetDetailVO`

NewVlanNetworkAffectingInternetDetailVOWithDefaults instantiates a new VlanNetworkAffectingInternetDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedPorts

`func (o *VlanNetworkAffectingInternetDetailVO) GetAffectedPorts() []int32`

GetAffectedPorts returns the AffectedPorts field if non-nil, zero value otherwise.

### GetAffectedPortsOk

`func (o *VlanNetworkAffectingInternetDetailVO) GetAffectedPortsOk() (*[]int32, bool)`

GetAffectedPortsOk returns a tuple with the AffectedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedPorts

`func (o *VlanNetworkAffectingInternetDetailVO) SetAffectedPorts(v []int32)`

SetAffectedPorts sets AffectedPorts field to given value.

### HasAffectedPorts

`func (o *VlanNetworkAffectingInternetDetailVO) HasAffectedPorts() bool`

HasAffectedPorts returns a boolean if a field has been set.

### GetInternetVO

`func (o *VlanNetworkAffectingInternetDetailVO) GetInternetVO() CheckWanLanStatusVO`

GetInternetVO returns the InternetVO field if non-nil, zero value otherwise.

### GetInternetVOOk

`func (o *VlanNetworkAffectingInternetDetailVO) GetInternetVOOk() (*CheckWanLanStatusVO, bool)`

GetInternetVOOk returns a tuple with the InternetVO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetVO

`func (o *VlanNetworkAffectingInternetDetailVO) SetInternetVO(v CheckWanLanStatusVO)`

SetInternetVO sets InternetVO field to given value.

### HasInternetVO

`func (o *VlanNetworkAffectingInternetDetailVO) HasInternetVO() bool`

HasInternetVO returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


