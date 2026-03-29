# OntEthPortVlanConfigModifyDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigItemAddList** | Pointer to [**[]OntPortVlanConfigAddItemDTO**](OntPortVlanConfigAddItemDTO.md) |  | [optional] 
**ConfigItemDeleteList** | Pointer to **[]int32** |  | [optional] 
**ConfigItemEditList** | Pointer to [**[]OntPortVlanConfigItemDTO**](OntPortVlanConfigItemDTO.md) |  | [optional] 
**VlanConfigMode** | **string** |  | 

## Methods

### NewOntEthPortVlanConfigModifyDTO

`func NewOntEthPortVlanConfigModifyDTO(vlanConfigMode string, ) *OntEthPortVlanConfigModifyDTO`

NewOntEthPortVlanConfigModifyDTO instantiates a new OntEthPortVlanConfigModifyDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOntEthPortVlanConfigModifyDTOWithDefaults

`func NewOntEthPortVlanConfigModifyDTOWithDefaults() *OntEthPortVlanConfigModifyDTO`

NewOntEthPortVlanConfigModifyDTOWithDefaults instantiates a new OntEthPortVlanConfigModifyDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigItemAddList

`func (o *OntEthPortVlanConfigModifyDTO) GetConfigItemAddList() []OntPortVlanConfigAddItemDTO`

GetConfigItemAddList returns the ConfigItemAddList field if non-nil, zero value otherwise.

### GetConfigItemAddListOk

`func (o *OntEthPortVlanConfigModifyDTO) GetConfigItemAddListOk() (*[]OntPortVlanConfigAddItemDTO, bool)`

GetConfigItemAddListOk returns a tuple with the ConfigItemAddList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigItemAddList

`func (o *OntEthPortVlanConfigModifyDTO) SetConfigItemAddList(v []OntPortVlanConfigAddItemDTO)`

SetConfigItemAddList sets ConfigItemAddList field to given value.

### HasConfigItemAddList

`func (o *OntEthPortVlanConfigModifyDTO) HasConfigItemAddList() bool`

HasConfigItemAddList returns a boolean if a field has been set.

### GetConfigItemDeleteList

`func (o *OntEthPortVlanConfigModifyDTO) GetConfigItemDeleteList() []int32`

GetConfigItemDeleteList returns the ConfigItemDeleteList field if non-nil, zero value otherwise.

### GetConfigItemDeleteListOk

`func (o *OntEthPortVlanConfigModifyDTO) GetConfigItemDeleteListOk() (*[]int32, bool)`

GetConfigItemDeleteListOk returns a tuple with the ConfigItemDeleteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigItemDeleteList

`func (o *OntEthPortVlanConfigModifyDTO) SetConfigItemDeleteList(v []int32)`

SetConfigItemDeleteList sets ConfigItemDeleteList field to given value.

### HasConfigItemDeleteList

`func (o *OntEthPortVlanConfigModifyDTO) HasConfigItemDeleteList() bool`

HasConfigItemDeleteList returns a boolean if a field has been set.

### GetConfigItemEditList

`func (o *OntEthPortVlanConfigModifyDTO) GetConfigItemEditList() []OntPortVlanConfigItemDTO`

GetConfigItemEditList returns the ConfigItemEditList field if non-nil, zero value otherwise.

### GetConfigItemEditListOk

`func (o *OntEthPortVlanConfigModifyDTO) GetConfigItemEditListOk() (*[]OntPortVlanConfigItemDTO, bool)`

GetConfigItemEditListOk returns a tuple with the ConfigItemEditList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigItemEditList

`func (o *OntEthPortVlanConfigModifyDTO) SetConfigItemEditList(v []OntPortVlanConfigItemDTO)`

SetConfigItemEditList sets ConfigItemEditList field to given value.

### HasConfigItemEditList

`func (o *OntEthPortVlanConfigModifyDTO) HasConfigItemEditList() bool`

HasConfigItemEditList returns a boolean if a field has been set.

### GetVlanConfigMode

`func (o *OntEthPortVlanConfigModifyDTO) GetVlanConfigMode() string`

GetVlanConfigMode returns the VlanConfigMode field if non-nil, zero value otherwise.

### GetVlanConfigModeOk

`func (o *OntEthPortVlanConfigModifyDTO) GetVlanConfigModeOk() (*string, bool)`

GetVlanConfigModeOk returns a tuple with the VlanConfigMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfigMode

`func (o *OntEthPortVlanConfigModifyDTO) SetVlanConfigMode(v string)`

SetVlanConfigMode sets VlanConfigMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


