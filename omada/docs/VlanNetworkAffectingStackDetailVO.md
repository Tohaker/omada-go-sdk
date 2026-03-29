# VlanNetworkAffectingStackDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedPorts** | Pointer to [**OswPortAndLagNetworkVO**](OswPortAndLagNetworkVO.md) |  | [optional] 
**Clients** | Pointer to [**[]OswClientVO**](OswClientVO.md) | Switch downlink clients. | [optional] 
**StackDetail** | Pointer to [**OswStackDetailVO**](OswStackDetailVO.md) |  | [optional] 
**StackLags** | Pointer to [**[]OswStackMemberLagVO**](OswStackMemberLagVO.md) | Stack lags | [optional] 

## Methods

### NewVlanNetworkAffectingStackDetailVO

`func NewVlanNetworkAffectingStackDetailVO() *VlanNetworkAffectingStackDetailVO`

NewVlanNetworkAffectingStackDetailVO instantiates a new VlanNetworkAffectingStackDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanNetworkAffectingStackDetailVOWithDefaults

`func NewVlanNetworkAffectingStackDetailVOWithDefaults() *VlanNetworkAffectingStackDetailVO`

NewVlanNetworkAffectingStackDetailVOWithDefaults instantiates a new VlanNetworkAffectingStackDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedPorts

`func (o *VlanNetworkAffectingStackDetailVO) GetAffectedPorts() OswPortAndLagNetworkVO`

GetAffectedPorts returns the AffectedPorts field if non-nil, zero value otherwise.

### GetAffectedPortsOk

`func (o *VlanNetworkAffectingStackDetailVO) GetAffectedPortsOk() (*OswPortAndLagNetworkVO, bool)`

GetAffectedPortsOk returns a tuple with the AffectedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedPorts

`func (o *VlanNetworkAffectingStackDetailVO) SetAffectedPorts(v OswPortAndLagNetworkVO)`

SetAffectedPorts sets AffectedPorts field to given value.

### HasAffectedPorts

`func (o *VlanNetworkAffectingStackDetailVO) HasAffectedPorts() bool`

HasAffectedPorts returns a boolean if a field has been set.

### GetClients

`func (o *VlanNetworkAffectingStackDetailVO) GetClients() []OswClientVO`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *VlanNetworkAffectingStackDetailVO) GetClientsOk() (*[]OswClientVO, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *VlanNetworkAffectingStackDetailVO) SetClients(v []OswClientVO)`

SetClients sets Clients field to given value.

### HasClients

`func (o *VlanNetworkAffectingStackDetailVO) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetStackDetail

`func (o *VlanNetworkAffectingStackDetailVO) GetStackDetail() OswStackDetailVO`

GetStackDetail returns the StackDetail field if non-nil, zero value otherwise.

### GetStackDetailOk

`func (o *VlanNetworkAffectingStackDetailVO) GetStackDetailOk() (*OswStackDetailVO, bool)`

GetStackDetailOk returns a tuple with the StackDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackDetail

`func (o *VlanNetworkAffectingStackDetailVO) SetStackDetail(v OswStackDetailVO)`

SetStackDetail sets StackDetail field to given value.

### HasStackDetail

`func (o *VlanNetworkAffectingStackDetailVO) HasStackDetail() bool`

HasStackDetail returns a boolean if a field has been set.

### GetStackLags

`func (o *VlanNetworkAffectingStackDetailVO) GetStackLags() []OswStackMemberLagVO`

GetStackLags returns the StackLags field if non-nil, zero value otherwise.

### GetStackLagsOk

`func (o *VlanNetworkAffectingStackDetailVO) GetStackLagsOk() (*[]OswStackMemberLagVO, bool)`

GetStackLagsOk returns a tuple with the StackLags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackLags

`func (o *VlanNetworkAffectingStackDetailVO) SetStackLags(v []OswStackMemberLagVO)`

SetStackLags sets StackLags field to given value.

### HasStackLags

`func (o *VlanNetworkAffectingStackDetailVO) HasStackLags() bool`

HasStackLags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


