# VlanNetworkAffectingSwitchDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedPorts** | Pointer to [**OswPortAndLagNetworkVO**](OswPortAndLagNetworkVO.md) |  | [optional] 
**Clients** | Pointer to [**[]OswClientVO**](OswClientVO.md) | Switch downlink clients. | [optional] 
**SwitchDetail** | Pointer to [**OswDetailVO**](OswDetailVO.md) |  | [optional] 

## Methods

### NewVlanNetworkAffectingSwitchDetailVO

`func NewVlanNetworkAffectingSwitchDetailVO() *VlanNetworkAffectingSwitchDetailVO`

NewVlanNetworkAffectingSwitchDetailVO instantiates a new VlanNetworkAffectingSwitchDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanNetworkAffectingSwitchDetailVOWithDefaults

`func NewVlanNetworkAffectingSwitchDetailVOWithDefaults() *VlanNetworkAffectingSwitchDetailVO`

NewVlanNetworkAffectingSwitchDetailVOWithDefaults instantiates a new VlanNetworkAffectingSwitchDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedPorts

`func (o *VlanNetworkAffectingSwitchDetailVO) GetAffectedPorts() OswPortAndLagNetworkVO`

GetAffectedPorts returns the AffectedPorts field if non-nil, zero value otherwise.

### GetAffectedPortsOk

`func (o *VlanNetworkAffectingSwitchDetailVO) GetAffectedPortsOk() (*OswPortAndLagNetworkVO, bool)`

GetAffectedPortsOk returns a tuple with the AffectedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedPorts

`func (o *VlanNetworkAffectingSwitchDetailVO) SetAffectedPorts(v OswPortAndLagNetworkVO)`

SetAffectedPorts sets AffectedPorts field to given value.

### HasAffectedPorts

`func (o *VlanNetworkAffectingSwitchDetailVO) HasAffectedPorts() bool`

HasAffectedPorts returns a boolean if a field has been set.

### GetClients

`func (o *VlanNetworkAffectingSwitchDetailVO) GetClients() []OswClientVO`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *VlanNetworkAffectingSwitchDetailVO) GetClientsOk() (*[]OswClientVO, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *VlanNetworkAffectingSwitchDetailVO) SetClients(v []OswClientVO)`

SetClients sets Clients field to given value.

### HasClients

`func (o *VlanNetworkAffectingSwitchDetailVO) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetSwitchDetail

`func (o *VlanNetworkAffectingSwitchDetailVO) GetSwitchDetail() OswDetailVO`

GetSwitchDetail returns the SwitchDetail field if non-nil, zero value otherwise.

### GetSwitchDetailOk

`func (o *VlanNetworkAffectingSwitchDetailVO) GetSwitchDetailOk() (*OswDetailVO, bool)`

GetSwitchDetailOk returns a tuple with the SwitchDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchDetail

`func (o *VlanNetworkAffectingSwitchDetailVO) SetSwitchDetail(v OswDetailVO)`

SetSwitchDetail sets SwitchDetail field to given value.

### HasSwitchDetail

`func (o *VlanNetworkAffectingSwitchDetailVO) HasSwitchDetail() bool`

HasSwitchDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


