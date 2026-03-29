# VlanNetworkAffectingEsDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedPorts** | Pointer to [**OswPortAndLagNetworkVO**](OswPortAndLagNetworkVO.md) |  | [optional] 
**Clients** | Pointer to [**[]OswClientVO**](OswClientVO.md) | Switch downlink clients. | [optional] 
**EsDetail** | Pointer to [**ESDetailVO**](ESDetailVO.md) |  | [optional] 

## Methods

### NewVlanNetworkAffectingEsDetailVO

`func NewVlanNetworkAffectingEsDetailVO() *VlanNetworkAffectingEsDetailVO`

NewVlanNetworkAffectingEsDetailVO instantiates a new VlanNetworkAffectingEsDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanNetworkAffectingEsDetailVOWithDefaults

`func NewVlanNetworkAffectingEsDetailVOWithDefaults() *VlanNetworkAffectingEsDetailVO`

NewVlanNetworkAffectingEsDetailVOWithDefaults instantiates a new VlanNetworkAffectingEsDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedPorts

`func (o *VlanNetworkAffectingEsDetailVO) GetAffectedPorts() OswPortAndLagNetworkVO`

GetAffectedPorts returns the AffectedPorts field if non-nil, zero value otherwise.

### GetAffectedPortsOk

`func (o *VlanNetworkAffectingEsDetailVO) GetAffectedPortsOk() (*OswPortAndLagNetworkVO, bool)`

GetAffectedPortsOk returns a tuple with the AffectedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedPorts

`func (o *VlanNetworkAffectingEsDetailVO) SetAffectedPorts(v OswPortAndLagNetworkVO)`

SetAffectedPorts sets AffectedPorts field to given value.

### HasAffectedPorts

`func (o *VlanNetworkAffectingEsDetailVO) HasAffectedPorts() bool`

HasAffectedPorts returns a boolean if a field has been set.

### GetClients

`func (o *VlanNetworkAffectingEsDetailVO) GetClients() []OswClientVO`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *VlanNetworkAffectingEsDetailVO) GetClientsOk() (*[]OswClientVO, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *VlanNetworkAffectingEsDetailVO) SetClients(v []OswClientVO)`

SetClients sets Clients field to given value.

### HasClients

`func (o *VlanNetworkAffectingEsDetailVO) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetEsDetail

`func (o *VlanNetworkAffectingEsDetailVO) GetEsDetail() ESDetailVO`

GetEsDetail returns the EsDetail field if non-nil, zero value otherwise.

### GetEsDetailOk

`func (o *VlanNetworkAffectingEsDetailVO) GetEsDetailOk() (*ESDetailVO, bool)`

GetEsDetailOk returns a tuple with the EsDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsDetail

`func (o *VlanNetworkAffectingEsDetailVO) SetEsDetail(v ESDetailVO)`

SetEsDetail sets EsDetail field to given value.

### HasEsDetail

`func (o *VlanNetworkAffectingEsDetailVO) HasEsDetail() bool`

HasEsDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


