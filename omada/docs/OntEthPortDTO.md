# OntEthPortDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgmpForward** | Pointer to [**OntEthPortIGMPForwardDTO**](OntEthPortIGMPForwardDTO.md) |  | [optional] 
**MaxMacCount** | Pointer to **int32** | MaxMacCount should be within the range of 0 to 255 | [optional] 
**PortId** | **int32** | Port ID | 
**PriorityPolicy** | Pointer to **string** | Priority policy should be a value as follows:UNCONCERN,ASSIGNED,COPY_COS | [optional] 
**QinQ** | Pointer to **string** | QinQ mode,qinQ should be a value as follows:UNCONCERN;DISABLE;ENABLE | [optional] 
**TlsVlan** | Pointer to **string** | Tls vlan | [optional] 
**VlanConfig** | Pointer to [**OntEthPortVlanConfigDTO**](OntEthPortVlanConfigDTO.md) |  | [optional] 

## Methods

### NewOntEthPortDTO

`func NewOntEthPortDTO(portId int32, ) *OntEthPortDTO`

NewOntEthPortDTO instantiates a new OntEthPortDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOntEthPortDTOWithDefaults

`func NewOntEthPortDTOWithDefaults() *OntEthPortDTO`

NewOntEthPortDTOWithDefaults instantiates a new OntEthPortDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgmpForward

`func (o *OntEthPortDTO) GetIgmpForward() OntEthPortIGMPForwardDTO`

GetIgmpForward returns the IgmpForward field if non-nil, zero value otherwise.

### GetIgmpForwardOk

`func (o *OntEthPortDTO) GetIgmpForwardOk() (*OntEthPortIGMPForwardDTO, bool)`

GetIgmpForwardOk returns a tuple with the IgmpForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpForward

`func (o *OntEthPortDTO) SetIgmpForward(v OntEthPortIGMPForwardDTO)`

SetIgmpForward sets IgmpForward field to given value.

### HasIgmpForward

`func (o *OntEthPortDTO) HasIgmpForward() bool`

HasIgmpForward returns a boolean if a field has been set.

### GetMaxMacCount

`func (o *OntEthPortDTO) GetMaxMacCount() int32`

GetMaxMacCount returns the MaxMacCount field if non-nil, zero value otherwise.

### GetMaxMacCountOk

`func (o *OntEthPortDTO) GetMaxMacCountOk() (*int32, bool)`

GetMaxMacCountOk returns a tuple with the MaxMacCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMacCount

`func (o *OntEthPortDTO) SetMaxMacCount(v int32)`

SetMaxMacCount sets MaxMacCount field to given value.

### HasMaxMacCount

`func (o *OntEthPortDTO) HasMaxMacCount() bool`

HasMaxMacCount returns a boolean if a field has been set.

### GetPortId

`func (o *OntEthPortDTO) GetPortId() int32`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *OntEthPortDTO) GetPortIdOk() (*int32, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *OntEthPortDTO) SetPortId(v int32)`

SetPortId sets PortId field to given value.


### GetPriorityPolicy

`func (o *OntEthPortDTO) GetPriorityPolicy() string`

GetPriorityPolicy returns the PriorityPolicy field if non-nil, zero value otherwise.

### GetPriorityPolicyOk

`func (o *OntEthPortDTO) GetPriorityPolicyOk() (*string, bool)`

GetPriorityPolicyOk returns a tuple with the PriorityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityPolicy

`func (o *OntEthPortDTO) SetPriorityPolicy(v string)`

SetPriorityPolicy sets PriorityPolicy field to given value.

### HasPriorityPolicy

`func (o *OntEthPortDTO) HasPriorityPolicy() bool`

HasPriorityPolicy returns a boolean if a field has been set.

### GetQinQ

`func (o *OntEthPortDTO) GetQinQ() string`

GetQinQ returns the QinQ field if non-nil, zero value otherwise.

### GetQinQOk

`func (o *OntEthPortDTO) GetQinQOk() (*string, bool)`

GetQinQOk returns a tuple with the QinQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinQ

`func (o *OntEthPortDTO) SetQinQ(v string)`

SetQinQ sets QinQ field to given value.

### HasQinQ

`func (o *OntEthPortDTO) HasQinQ() bool`

HasQinQ returns a boolean if a field has been set.

### GetTlsVlan

`func (o *OntEthPortDTO) GetTlsVlan() string`

GetTlsVlan returns the TlsVlan field if non-nil, zero value otherwise.

### GetTlsVlanOk

`func (o *OntEthPortDTO) GetTlsVlanOk() (*string, bool)`

GetTlsVlanOk returns a tuple with the TlsVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVlan

`func (o *OntEthPortDTO) SetTlsVlan(v string)`

SetTlsVlan sets TlsVlan field to given value.

### HasTlsVlan

`func (o *OntEthPortDTO) HasTlsVlan() bool`

HasTlsVlan returns a boolean if a field has been set.

### GetVlanConfig

`func (o *OntEthPortDTO) GetVlanConfig() OntEthPortVlanConfigDTO`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *OntEthPortDTO) GetVlanConfigOk() (*OntEthPortVlanConfigDTO, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *OntEthPortDTO) SetVlanConfig(v OntEthPortVlanConfigDTO)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *OntEthPortDTO) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


