# OntPotsPortModifyDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigItemAddList** | Pointer to [**[]OntPortVlanConfigAddItemDTO**](OntPortVlanConfigAddItemDTO.md) | Config item add list | [optional] 
**ConfigItemDeleteList** | Pointer to **[]int32** | Config item delete list | [optional] 
**ConfigItemEditList** | Pointer to [**[]OntPortVlanConfigItemDTO**](OntPortVlanConfigItemDTO.md) | Config item edit list | [optional] 
**ServiceId** | Pointer to **int32** | The associated Service Profile ID | [optional] 
**VlanConfigMode** | **string** | Vlan config mode.VlanConfigMode should be a value as follows:TRANSPARENT,OTHERS | 

## Methods

### NewOntPotsPortModifyDTO

`func NewOntPotsPortModifyDTO(vlanConfigMode string, ) *OntPotsPortModifyDTO`

NewOntPotsPortModifyDTO instantiates a new OntPotsPortModifyDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOntPotsPortModifyDTOWithDefaults

`func NewOntPotsPortModifyDTOWithDefaults() *OntPotsPortModifyDTO`

NewOntPotsPortModifyDTOWithDefaults instantiates a new OntPotsPortModifyDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigItemAddList

`func (o *OntPotsPortModifyDTO) GetConfigItemAddList() []OntPortVlanConfigAddItemDTO`

GetConfigItemAddList returns the ConfigItemAddList field if non-nil, zero value otherwise.

### GetConfigItemAddListOk

`func (o *OntPotsPortModifyDTO) GetConfigItemAddListOk() (*[]OntPortVlanConfigAddItemDTO, bool)`

GetConfigItemAddListOk returns a tuple with the ConfigItemAddList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigItemAddList

`func (o *OntPotsPortModifyDTO) SetConfigItemAddList(v []OntPortVlanConfigAddItemDTO)`

SetConfigItemAddList sets ConfigItemAddList field to given value.

### HasConfigItemAddList

`func (o *OntPotsPortModifyDTO) HasConfigItemAddList() bool`

HasConfigItemAddList returns a boolean if a field has been set.

### GetConfigItemDeleteList

`func (o *OntPotsPortModifyDTO) GetConfigItemDeleteList() []int32`

GetConfigItemDeleteList returns the ConfigItemDeleteList field if non-nil, zero value otherwise.

### GetConfigItemDeleteListOk

`func (o *OntPotsPortModifyDTO) GetConfigItemDeleteListOk() (*[]int32, bool)`

GetConfigItemDeleteListOk returns a tuple with the ConfigItemDeleteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigItemDeleteList

`func (o *OntPotsPortModifyDTO) SetConfigItemDeleteList(v []int32)`

SetConfigItemDeleteList sets ConfigItemDeleteList field to given value.

### HasConfigItemDeleteList

`func (o *OntPotsPortModifyDTO) HasConfigItemDeleteList() bool`

HasConfigItemDeleteList returns a boolean if a field has been set.

### GetConfigItemEditList

`func (o *OntPotsPortModifyDTO) GetConfigItemEditList() []OntPortVlanConfigItemDTO`

GetConfigItemEditList returns the ConfigItemEditList field if non-nil, zero value otherwise.

### GetConfigItemEditListOk

`func (o *OntPotsPortModifyDTO) GetConfigItemEditListOk() (*[]OntPortVlanConfigItemDTO, bool)`

GetConfigItemEditListOk returns a tuple with the ConfigItemEditList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigItemEditList

`func (o *OntPotsPortModifyDTO) SetConfigItemEditList(v []OntPortVlanConfigItemDTO)`

SetConfigItemEditList sets ConfigItemEditList field to given value.

### HasConfigItemEditList

`func (o *OntPotsPortModifyDTO) HasConfigItemEditList() bool`

HasConfigItemEditList returns a boolean if a field has been set.

### GetServiceId

`func (o *OntPotsPortModifyDTO) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *OntPotsPortModifyDTO) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *OntPotsPortModifyDTO) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *OntPotsPortModifyDTO) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetVlanConfigMode

`func (o *OntPotsPortModifyDTO) GetVlanConfigMode() string`

GetVlanConfigMode returns the VlanConfigMode field if non-nil, zero value otherwise.

### GetVlanConfigModeOk

`func (o *OntPotsPortModifyDTO) GetVlanConfigModeOk() (*string, bool)`

GetVlanConfigModeOk returns a tuple with the VlanConfigMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfigMode

`func (o *OntPotsPortModifyDTO) SetVlanConfigMode(v string)`

SetVlanConfigMode sets VlanConfigMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


