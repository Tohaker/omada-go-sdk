# OntEthPortIGMPForwardDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgmpForwardMode** | **string** | IgmpForwardMode should not be null,igmpForwardMode should be a value as follows:UNCONCERN,TRANSLATION,DEFAULT,TRANSPARENT | 
**Priority** | Pointer to **int32** | Priority should be within the range of -1 to 7 | [optional] 
**Vlan** | Pointer to **int32** | Vlan should be within the range of 1 to 4095 | [optional] 

## Methods

### NewOntEthPortIGMPForwardDTO

`func NewOntEthPortIGMPForwardDTO(igmpForwardMode string, ) *OntEthPortIGMPForwardDTO`

NewOntEthPortIGMPForwardDTO instantiates a new OntEthPortIGMPForwardDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOntEthPortIGMPForwardDTOWithDefaults

`func NewOntEthPortIGMPForwardDTOWithDefaults() *OntEthPortIGMPForwardDTO`

NewOntEthPortIGMPForwardDTOWithDefaults instantiates a new OntEthPortIGMPForwardDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgmpForwardMode

`func (o *OntEthPortIGMPForwardDTO) GetIgmpForwardMode() string`

GetIgmpForwardMode returns the IgmpForwardMode field if non-nil, zero value otherwise.

### GetIgmpForwardModeOk

`func (o *OntEthPortIGMPForwardDTO) GetIgmpForwardModeOk() (*string, bool)`

GetIgmpForwardModeOk returns a tuple with the IgmpForwardMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpForwardMode

`func (o *OntEthPortIGMPForwardDTO) SetIgmpForwardMode(v string)`

SetIgmpForwardMode sets IgmpForwardMode field to given value.


### GetPriority

`func (o *OntEthPortIGMPForwardDTO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *OntEthPortIGMPForwardDTO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *OntEthPortIGMPForwardDTO) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *OntEthPortIGMPForwardDTO) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetVlan

`func (o *OntEthPortIGMPForwardDTO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *OntEthPortIGMPForwardDTO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *OntEthPortIGMPForwardDTO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *OntEthPortIGMPForwardDTO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


