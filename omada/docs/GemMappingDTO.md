# GemMappingDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GemMappingId** | **int32** | GemMappingId should be within the range of 1 to 8,and should not be null | 
**GemPortId** | **int32** | Gem port ID should be within the range of 1 to 1023 | 
**LineProfileId** | Pointer to **int32** | The ID of the associated Line Profile should be within the range of 1 to 512. | [optional] 
**PortId** | Pointer to **int32** | The mapped ONT port ID is valid when the portMappingType has a value. When portMappingType is ETH, portId should be within the range of 1 to 24. | [optional] 
**PortMappingType** | Pointer to **string** | The mapping relationship between Gem Port and ONT Port ID is valid when the Line Profile&#39;s mappingMode is set to: Port, Port-Vlan, Port-Priority, or Port-Vlan-Priority. PortMappingType should be a value as follows:ETH;POTS, with the default value being ETH. | [optional] 
**Priority** | Pointer to **int32** | The mapping relationship between Gem Port and 802.1p priority.It is valid when the Line Profile&#39;s mappingMode is set to: Priority, Vlan-Priority, Port-Priority, or Port-VLAN-Priority. Priority should be within the range of 0 to 7. | [optional] 
**VlanId** | Pointer to **int32** | VLAN ID is valid when the vlanType is set to Tagged. VlanId should be within the range of 1 to 4094. | [optional] 
**VlanType** | Pointer to **string** | The mapping relationship between Gem Port and VLAN is valid when the Line Profile&#39;s mappingMode is set to: Vlan, Vlan-Priority, Port-Vlan, or Port-VLAN-Priority. VlanType should be a value as follows: TAGGED;UNTAGGED, with the default value being TAGGED. | [optional] 

## Methods

### NewGemMappingDTO

`func NewGemMappingDTO(gemMappingId int32, gemPortId int32, ) *GemMappingDTO`

NewGemMappingDTO instantiates a new GemMappingDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGemMappingDTOWithDefaults

`func NewGemMappingDTOWithDefaults() *GemMappingDTO`

NewGemMappingDTOWithDefaults instantiates a new GemMappingDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGemMappingId

`func (o *GemMappingDTO) GetGemMappingId() int32`

GetGemMappingId returns the GemMappingId field if non-nil, zero value otherwise.

### GetGemMappingIdOk

`func (o *GemMappingDTO) GetGemMappingIdOk() (*int32, bool)`

GetGemMappingIdOk returns a tuple with the GemMappingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemMappingId

`func (o *GemMappingDTO) SetGemMappingId(v int32)`

SetGemMappingId sets GemMappingId field to given value.


### GetGemPortId

`func (o *GemMappingDTO) GetGemPortId() int32`

GetGemPortId returns the GemPortId field if non-nil, zero value otherwise.

### GetGemPortIdOk

`func (o *GemMappingDTO) GetGemPortIdOk() (*int32, bool)`

GetGemPortIdOk returns a tuple with the GemPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemPortId

`func (o *GemMappingDTO) SetGemPortId(v int32)`

SetGemPortId sets GemPortId field to given value.


### GetLineProfileId

`func (o *GemMappingDTO) GetLineProfileId() int32`

GetLineProfileId returns the LineProfileId field if non-nil, zero value otherwise.

### GetLineProfileIdOk

`func (o *GemMappingDTO) GetLineProfileIdOk() (*int32, bool)`

GetLineProfileIdOk returns a tuple with the LineProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineProfileId

`func (o *GemMappingDTO) SetLineProfileId(v int32)`

SetLineProfileId sets LineProfileId field to given value.

### HasLineProfileId

`func (o *GemMappingDTO) HasLineProfileId() bool`

HasLineProfileId returns a boolean if a field has been set.

### GetPortId

`func (o *GemMappingDTO) GetPortId() int32`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *GemMappingDTO) GetPortIdOk() (*int32, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *GemMappingDTO) SetPortId(v int32)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *GemMappingDTO) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPortMappingType

`func (o *GemMappingDTO) GetPortMappingType() string`

GetPortMappingType returns the PortMappingType field if non-nil, zero value otherwise.

### GetPortMappingTypeOk

`func (o *GemMappingDTO) GetPortMappingTypeOk() (*string, bool)`

GetPortMappingTypeOk returns a tuple with the PortMappingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMappingType

`func (o *GemMappingDTO) SetPortMappingType(v string)`

SetPortMappingType sets PortMappingType field to given value.

### HasPortMappingType

`func (o *GemMappingDTO) HasPortMappingType() bool`

HasPortMappingType returns a boolean if a field has been set.

### GetPriority

`func (o *GemMappingDTO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GemMappingDTO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GemMappingDTO) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GemMappingDTO) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetVlanId

`func (o *GemMappingDTO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *GemMappingDTO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *GemMappingDTO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *GemMappingDTO) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlanType

`func (o *GemMappingDTO) GetVlanType() string`

GetVlanType returns the VlanType field if non-nil, zero value otherwise.

### GetVlanTypeOk

`func (o *GemMappingDTO) GetVlanTypeOk() (*string, bool)`

GetVlanTypeOk returns a tuple with the VlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanType

`func (o *GemMappingDTO) SetVlanType(v string)`

SetVlanType sets VlanType field to given value.

### HasVlanType

`func (o *GemMappingDTO) HasVlanType() bool`

HasVlanType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


