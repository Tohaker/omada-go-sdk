# VlanNetworkAffectingOsgDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedPorts** | Pointer to **[]int32** | Affected gateway port list. | [optional] 
**GatewayDetail** | Pointer to [**OsgDetailVO**](OsgDetailVO.md) |  | [optional] 

## Methods

### NewVlanNetworkAffectingOsgDetailVO

`func NewVlanNetworkAffectingOsgDetailVO() *VlanNetworkAffectingOsgDetailVO`

NewVlanNetworkAffectingOsgDetailVO instantiates a new VlanNetworkAffectingOsgDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanNetworkAffectingOsgDetailVOWithDefaults

`func NewVlanNetworkAffectingOsgDetailVOWithDefaults() *VlanNetworkAffectingOsgDetailVO`

NewVlanNetworkAffectingOsgDetailVOWithDefaults instantiates a new VlanNetworkAffectingOsgDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedPorts

`func (o *VlanNetworkAffectingOsgDetailVO) GetAffectedPorts() []int32`

GetAffectedPorts returns the AffectedPorts field if non-nil, zero value otherwise.

### GetAffectedPortsOk

`func (o *VlanNetworkAffectingOsgDetailVO) GetAffectedPortsOk() (*[]int32, bool)`

GetAffectedPortsOk returns a tuple with the AffectedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedPorts

`func (o *VlanNetworkAffectingOsgDetailVO) SetAffectedPorts(v []int32)`

SetAffectedPorts sets AffectedPorts field to given value.

### HasAffectedPorts

`func (o *VlanNetworkAffectingOsgDetailVO) HasAffectedPorts() bool`

HasAffectedPorts returns a boolean if a field has been set.

### GetGatewayDetail

`func (o *VlanNetworkAffectingOsgDetailVO) GetGatewayDetail() OsgDetailVO`

GetGatewayDetail returns the GatewayDetail field if non-nil, zero value otherwise.

### GetGatewayDetailOk

`func (o *VlanNetworkAffectingOsgDetailVO) GetGatewayDetailOk() (*OsgDetailVO, bool)`

GetGatewayDetailOk returns a tuple with the GatewayDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayDetail

`func (o *VlanNetworkAffectingOsgDetailVO) SetGatewayDetail(v OsgDetailVO)`

SetGatewayDetail sets GatewayDetail field to given value.

### HasGatewayDetail

`func (o *VlanNetworkAffectingOsgDetailVO) HasGatewayDetail() bool`

HasGatewayDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


