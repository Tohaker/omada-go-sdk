# OntEthPortModifyDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgmpForward** | Pointer to [**OntEthPortIGMPForwardDTO**](OntEthPortIGMPForwardDTO.md) |  | [optional] 
**MaxMacCount** | Pointer to **int32** | MaxMacCount should be within the range of 0 to 255 | [optional] 
**PortId** | **int32** | Port ID | 
**PriorityPolicy** | Pointer to **string** | Priority policy should be a value as follows:UNCONCERN,ASSIGNED,COPY_COS | [optional] 
**QinQ** | Pointer to **string** | QinQ mode,qinQ should be a value as follows:UNCONCERN,DISABLE,ENABLE | [optional] 
**ServiceId** | Pointer to **int32** | The associated Service Profile ID | [optional] 
**TlsVlan** | Pointer to **string** | Tls vlan | [optional] 
**VlanConfig** | Pointer to [**OntEthPortVlanConfigModifyDTO**](OntEthPortVlanConfigModifyDTO.md) |  | [optional] 

## Methods

### NewOntEthPortModifyDTO

`func NewOntEthPortModifyDTO(portId int32, ) *OntEthPortModifyDTO`

NewOntEthPortModifyDTO instantiates a new OntEthPortModifyDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOntEthPortModifyDTOWithDefaults

`func NewOntEthPortModifyDTOWithDefaults() *OntEthPortModifyDTO`

NewOntEthPortModifyDTOWithDefaults instantiates a new OntEthPortModifyDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgmpForward

`func (o *OntEthPortModifyDTO) GetIgmpForward() OntEthPortIGMPForwardDTO`

GetIgmpForward returns the IgmpForward field if non-nil, zero value otherwise.

### GetIgmpForwardOk

`func (o *OntEthPortModifyDTO) GetIgmpForwardOk() (*OntEthPortIGMPForwardDTO, bool)`

GetIgmpForwardOk returns a tuple with the IgmpForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpForward

`func (o *OntEthPortModifyDTO) SetIgmpForward(v OntEthPortIGMPForwardDTO)`

SetIgmpForward sets IgmpForward field to given value.

### HasIgmpForward

`func (o *OntEthPortModifyDTO) HasIgmpForward() bool`

HasIgmpForward returns a boolean if a field has been set.

### GetMaxMacCount

`func (o *OntEthPortModifyDTO) GetMaxMacCount() int32`

GetMaxMacCount returns the MaxMacCount field if non-nil, zero value otherwise.

### GetMaxMacCountOk

`func (o *OntEthPortModifyDTO) GetMaxMacCountOk() (*int32, bool)`

GetMaxMacCountOk returns a tuple with the MaxMacCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMacCount

`func (o *OntEthPortModifyDTO) SetMaxMacCount(v int32)`

SetMaxMacCount sets MaxMacCount field to given value.

### HasMaxMacCount

`func (o *OntEthPortModifyDTO) HasMaxMacCount() bool`

HasMaxMacCount returns a boolean if a field has been set.

### GetPortId

`func (o *OntEthPortModifyDTO) GetPortId() int32`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *OntEthPortModifyDTO) GetPortIdOk() (*int32, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *OntEthPortModifyDTO) SetPortId(v int32)`

SetPortId sets PortId field to given value.


### GetPriorityPolicy

`func (o *OntEthPortModifyDTO) GetPriorityPolicy() string`

GetPriorityPolicy returns the PriorityPolicy field if non-nil, zero value otherwise.

### GetPriorityPolicyOk

`func (o *OntEthPortModifyDTO) GetPriorityPolicyOk() (*string, bool)`

GetPriorityPolicyOk returns a tuple with the PriorityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityPolicy

`func (o *OntEthPortModifyDTO) SetPriorityPolicy(v string)`

SetPriorityPolicy sets PriorityPolicy field to given value.

### HasPriorityPolicy

`func (o *OntEthPortModifyDTO) HasPriorityPolicy() bool`

HasPriorityPolicy returns a boolean if a field has been set.

### GetQinQ

`func (o *OntEthPortModifyDTO) GetQinQ() string`

GetQinQ returns the QinQ field if non-nil, zero value otherwise.

### GetQinQOk

`func (o *OntEthPortModifyDTO) GetQinQOk() (*string, bool)`

GetQinQOk returns a tuple with the QinQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinQ

`func (o *OntEthPortModifyDTO) SetQinQ(v string)`

SetQinQ sets QinQ field to given value.

### HasQinQ

`func (o *OntEthPortModifyDTO) HasQinQ() bool`

HasQinQ returns a boolean if a field has been set.

### GetServiceId

`func (o *OntEthPortModifyDTO) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *OntEthPortModifyDTO) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *OntEthPortModifyDTO) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *OntEthPortModifyDTO) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetTlsVlan

`func (o *OntEthPortModifyDTO) GetTlsVlan() string`

GetTlsVlan returns the TlsVlan field if non-nil, zero value otherwise.

### GetTlsVlanOk

`func (o *OntEthPortModifyDTO) GetTlsVlanOk() (*string, bool)`

GetTlsVlanOk returns a tuple with the TlsVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVlan

`func (o *OntEthPortModifyDTO) SetTlsVlan(v string)`

SetTlsVlan sets TlsVlan field to given value.

### HasTlsVlan

`func (o *OntEthPortModifyDTO) HasTlsVlan() bool`

HasTlsVlan returns a boolean if a field has been set.

### GetVlanConfig

`func (o *OntEthPortModifyDTO) GetVlanConfig() OntEthPortVlanConfigModifyDTO`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *OntEthPortModifyDTO) GetVlanConfigOk() (*OntEthPortVlanConfigModifyDTO, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *OntEthPortModifyDTO) SetVlanConfig(v OntEthPortVlanConfigModifyDTO)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *OntEthPortModifyDTO) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


