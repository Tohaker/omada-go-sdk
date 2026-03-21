# OntPotsPortDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigItemList** | Pointer to [**[]OntPortVlanConfigItemDTO**](OntPortVlanConfigItemDTO.md) | Config item list | [optional] 
**VlanConfigMode** | **string** | Vlan config mode.VlanConfigMode should be a value as follows:TRANSPARENT,OTHERS | 

## Methods

### NewOntPotsPortDTO

`func NewOntPotsPortDTO(vlanConfigMode string, ) *OntPotsPortDTO`

NewOntPotsPortDTO instantiates a new OntPotsPortDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOntPotsPortDTOWithDefaults

`func NewOntPotsPortDTOWithDefaults() *OntPotsPortDTO`

NewOntPotsPortDTOWithDefaults instantiates a new OntPotsPortDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigItemList

`func (o *OntPotsPortDTO) GetConfigItemList() []OntPortVlanConfigItemDTO`

GetConfigItemList returns the ConfigItemList field if non-nil, zero value otherwise.

### GetConfigItemListOk

`func (o *OntPotsPortDTO) GetConfigItemListOk() (*[]OntPortVlanConfigItemDTO, bool)`

GetConfigItemListOk returns a tuple with the ConfigItemList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigItemList

`func (o *OntPotsPortDTO) SetConfigItemList(v []OntPortVlanConfigItemDTO)`

SetConfigItemList sets ConfigItemList field to given value.

### HasConfigItemList

`func (o *OntPotsPortDTO) HasConfigItemList() bool`

HasConfigItemList returns a boolean if a field has been set.

### GetVlanConfigMode

`func (o *OntPotsPortDTO) GetVlanConfigMode() string`

GetVlanConfigMode returns the VlanConfigMode field if non-nil, zero value otherwise.

### GetVlanConfigModeOk

`func (o *OntPotsPortDTO) GetVlanConfigModeOk() (*string, bool)`

GetVlanConfigModeOk returns a tuple with the VlanConfigMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfigMode

`func (o *OntPotsPortDTO) SetVlanConfigMode(v string)`

SetVlanConfigMode sets VlanConfigMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


