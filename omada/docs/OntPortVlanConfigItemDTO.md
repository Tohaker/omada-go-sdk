# OntPortVlanConfigItemDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpriority** | Pointer to **int32** | CVlan priority should be within the range of -1 to 7 | [optional] 
**Cvlan** | Pointer to **int32** | CVlan ID should be within the range of 1 to 4095 | [optional] 
**Id** | **int32** | Vlan ID should not be null | 
**Spriority** | Pointer to **int32** | SVlan priority should be within the range of -2 to 7 | [optional] 
**Svlan** | **int32** | SVlan ID should be within the range of 1 to 4095 and should not be null | 
**VlanMode** | **string** | Vlan mode should not be null,vlanMode should be a value as follows:TRANSLATION,QINQ,TRUNK | 

## Methods

### NewOntPortVlanConfigItemDTO

`func NewOntPortVlanConfigItemDTO(id int32, svlan int32, vlanMode string, ) *OntPortVlanConfigItemDTO`

NewOntPortVlanConfigItemDTO instantiates a new OntPortVlanConfigItemDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOntPortVlanConfigItemDTOWithDefaults

`func NewOntPortVlanConfigItemDTOWithDefaults() *OntPortVlanConfigItemDTO`

NewOntPortVlanConfigItemDTOWithDefaults instantiates a new OntPortVlanConfigItemDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpriority

`func (o *OntPortVlanConfigItemDTO) GetCpriority() int32`

GetCpriority returns the Cpriority field if non-nil, zero value otherwise.

### GetCpriorityOk

`func (o *OntPortVlanConfigItemDTO) GetCpriorityOk() (*int32, bool)`

GetCpriorityOk returns a tuple with the Cpriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpriority

`func (o *OntPortVlanConfigItemDTO) SetCpriority(v int32)`

SetCpriority sets Cpriority field to given value.

### HasCpriority

`func (o *OntPortVlanConfigItemDTO) HasCpriority() bool`

HasCpriority returns a boolean if a field has been set.

### GetCvlan

`func (o *OntPortVlanConfigItemDTO) GetCvlan() int32`

GetCvlan returns the Cvlan field if non-nil, zero value otherwise.

### GetCvlanOk

`func (o *OntPortVlanConfigItemDTO) GetCvlanOk() (*int32, bool)`

GetCvlanOk returns a tuple with the Cvlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvlan

`func (o *OntPortVlanConfigItemDTO) SetCvlan(v int32)`

SetCvlan sets Cvlan field to given value.

### HasCvlan

`func (o *OntPortVlanConfigItemDTO) HasCvlan() bool`

HasCvlan returns a boolean if a field has been set.

### GetId

`func (o *OntPortVlanConfigItemDTO) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OntPortVlanConfigItemDTO) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OntPortVlanConfigItemDTO) SetId(v int32)`

SetId sets Id field to given value.


### GetSpriority

`func (o *OntPortVlanConfigItemDTO) GetSpriority() int32`

GetSpriority returns the Spriority field if non-nil, zero value otherwise.

### GetSpriorityOk

`func (o *OntPortVlanConfigItemDTO) GetSpriorityOk() (*int32, bool)`

GetSpriorityOk returns a tuple with the Spriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpriority

`func (o *OntPortVlanConfigItemDTO) SetSpriority(v int32)`

SetSpriority sets Spriority field to given value.

### HasSpriority

`func (o *OntPortVlanConfigItemDTO) HasSpriority() bool`

HasSpriority returns a boolean if a field has been set.

### GetSvlan

`func (o *OntPortVlanConfigItemDTO) GetSvlan() int32`

GetSvlan returns the Svlan field if non-nil, zero value otherwise.

### GetSvlanOk

`func (o *OntPortVlanConfigItemDTO) GetSvlanOk() (*int32, bool)`

GetSvlanOk returns a tuple with the Svlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvlan

`func (o *OntPortVlanConfigItemDTO) SetSvlan(v int32)`

SetSvlan sets Svlan field to given value.


### GetVlanMode

`func (o *OntPortVlanConfigItemDTO) GetVlanMode() string`

GetVlanMode returns the VlanMode field if non-nil, zero value otherwise.

### GetVlanModeOk

`func (o *OntPortVlanConfigItemDTO) GetVlanModeOk() (*string, bool)`

GetVlanModeOk returns a tuple with the VlanMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanMode

`func (o *OntPortVlanConfigItemDTO) SetVlanMode(v string)`

SetVlanMode sets VlanMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


