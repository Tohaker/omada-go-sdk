# OntPortVlanConfigAddItemDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpriority** | Pointer to **int32** | CVlan priority should be within the range of -1 to 7 | [optional] 
**Cvlan** | Pointer to **int32** | CVlan ID should be within the range of 1 to 4095 | [optional] 
**Spriority** | Pointer to **int32** | SVlan priority should be within the range of -2 to 7 | [optional] 
**Svlan** | **int32** | SVlan ID should be within the range of 1 to 4095 and should not be null | 
**VlanMode** | **string** | Vlan mode should not be null,vlanMode should be a value as follows:TRANSLATION,QINQ,TRUNK | 

## Methods

### NewOntPortVlanConfigAddItemDTO

`func NewOntPortVlanConfigAddItemDTO(svlan int32, vlanMode string, ) *OntPortVlanConfigAddItemDTO`

NewOntPortVlanConfigAddItemDTO instantiates a new OntPortVlanConfigAddItemDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOntPortVlanConfigAddItemDTOWithDefaults

`func NewOntPortVlanConfigAddItemDTOWithDefaults() *OntPortVlanConfigAddItemDTO`

NewOntPortVlanConfigAddItemDTOWithDefaults instantiates a new OntPortVlanConfigAddItemDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpriority

`func (o *OntPortVlanConfigAddItemDTO) GetCpriority() int32`

GetCpriority returns the Cpriority field if non-nil, zero value otherwise.

### GetCpriorityOk

`func (o *OntPortVlanConfigAddItemDTO) GetCpriorityOk() (*int32, bool)`

GetCpriorityOk returns a tuple with the Cpriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpriority

`func (o *OntPortVlanConfigAddItemDTO) SetCpriority(v int32)`

SetCpriority sets Cpriority field to given value.

### HasCpriority

`func (o *OntPortVlanConfigAddItemDTO) HasCpriority() bool`

HasCpriority returns a boolean if a field has been set.

### GetCvlan

`func (o *OntPortVlanConfigAddItemDTO) GetCvlan() int32`

GetCvlan returns the Cvlan field if non-nil, zero value otherwise.

### GetCvlanOk

`func (o *OntPortVlanConfigAddItemDTO) GetCvlanOk() (*int32, bool)`

GetCvlanOk returns a tuple with the Cvlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvlan

`func (o *OntPortVlanConfigAddItemDTO) SetCvlan(v int32)`

SetCvlan sets Cvlan field to given value.

### HasCvlan

`func (o *OntPortVlanConfigAddItemDTO) HasCvlan() bool`

HasCvlan returns a boolean if a field has been set.

### GetSpriority

`func (o *OntPortVlanConfigAddItemDTO) GetSpriority() int32`

GetSpriority returns the Spriority field if non-nil, zero value otherwise.

### GetSpriorityOk

`func (o *OntPortVlanConfigAddItemDTO) GetSpriorityOk() (*int32, bool)`

GetSpriorityOk returns a tuple with the Spriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpriority

`func (o *OntPortVlanConfigAddItemDTO) SetSpriority(v int32)`

SetSpriority sets Spriority field to given value.

### HasSpriority

`func (o *OntPortVlanConfigAddItemDTO) HasSpriority() bool`

HasSpriority returns a boolean if a field has been set.

### GetSvlan

`func (o *OntPortVlanConfigAddItemDTO) GetSvlan() int32`

GetSvlan returns the Svlan field if non-nil, zero value otherwise.

### GetSvlanOk

`func (o *OntPortVlanConfigAddItemDTO) GetSvlanOk() (*int32, bool)`

GetSvlanOk returns a tuple with the Svlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvlan

`func (o *OntPortVlanConfigAddItemDTO) SetSvlan(v int32)`

SetSvlan sets Svlan field to given value.


### GetVlanMode

`func (o *OntPortVlanConfigAddItemDTO) GetVlanMode() string`

GetVlanMode returns the VlanMode field if non-nil, zero value otherwise.

### GetVlanModeOk

`func (o *OntPortVlanConfigAddItemDTO) GetVlanModeOk() (*string, bool)`

GetVlanModeOk returns a tuple with the VlanMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanMode

`func (o *OntPortVlanConfigAddItemDTO) SetVlanMode(v string)`

SetVlanMode sets VlanMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


