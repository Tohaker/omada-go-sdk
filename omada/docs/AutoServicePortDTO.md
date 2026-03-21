# AutoServicePortDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoMode** | Pointer to **string** | AutoMOde should be a value as follows:Disable,Enable | [optional] 
**EtherType** | Pointer to **string** | EtherType should be a value as follows:NONE,IPV4OE,IPV6OE,PPPOE | [optional] 
**GemPortId** | Pointer to **int32** | GemPortId should be within the range of 1 to 1023 | [optional] 
**InboundTrafficProfileId** | Pointer to **int32** | InboundTrafficProfileId should be within the range of 0 to 512 | [optional] 
**InnerVlan** | Pointer to **int32** | InnerVlan should be within the range of 0 to 4095 | [optional] 
**InnerVlanPriority** | Pointer to **int32** | InnerVlanPriority should be within the range of -1 to 7 | [optional] 
**OutboundTrafficProfileId** | Pointer to **int32** | OutboundTrafficProfileId should be within the range of 0 to 512 | [optional] 
**PonPortId** | Pointer to **int32** | PonPortId should be within the range of 1 to 16 | [optional] 
**PonPortStr** | Pointer to **string** | String form of pon port | [optional] 
**Svlan** | Pointer to **int32** | SVlan should be within the range of 1 to 4095 | [optional] 
**TagAction** | Pointer to **string** | TagAction should be a value as follows:DEFAULT,TRANSPARENT,TRANSLATE,TRANSLATE_AND_ADD,ADD_DOUBLE | [optional] 
**UserVlan** | Pointer to **int32** | UserVlan should be within the range of 0 to 4095 | [optional] 
**UserVlanPriority** | Pointer to **int32** | UserVlanPriority should be within the range of -1 to 7 | [optional] 

## Methods

### NewAutoServicePortDTO

`func NewAutoServicePortDTO() *AutoServicePortDTO`

NewAutoServicePortDTO instantiates a new AutoServicePortDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoServicePortDTOWithDefaults

`func NewAutoServicePortDTOWithDefaults() *AutoServicePortDTO`

NewAutoServicePortDTOWithDefaults instantiates a new AutoServicePortDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoMode

`func (o *AutoServicePortDTO) GetAutoMode() string`

GetAutoMode returns the AutoMode field if non-nil, zero value otherwise.

### GetAutoModeOk

`func (o *AutoServicePortDTO) GetAutoModeOk() (*string, bool)`

GetAutoModeOk returns a tuple with the AutoMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoMode

`func (o *AutoServicePortDTO) SetAutoMode(v string)`

SetAutoMode sets AutoMode field to given value.

### HasAutoMode

`func (o *AutoServicePortDTO) HasAutoMode() bool`

HasAutoMode returns a boolean if a field has been set.

### GetEtherType

`func (o *AutoServicePortDTO) GetEtherType() string`

GetEtherType returns the EtherType field if non-nil, zero value otherwise.

### GetEtherTypeOk

`func (o *AutoServicePortDTO) GetEtherTypeOk() (*string, bool)`

GetEtherTypeOk returns a tuple with the EtherType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherType

`func (o *AutoServicePortDTO) SetEtherType(v string)`

SetEtherType sets EtherType field to given value.

### HasEtherType

`func (o *AutoServicePortDTO) HasEtherType() bool`

HasEtherType returns a boolean if a field has been set.

### GetGemPortId

`func (o *AutoServicePortDTO) GetGemPortId() int32`

GetGemPortId returns the GemPortId field if non-nil, zero value otherwise.

### GetGemPortIdOk

`func (o *AutoServicePortDTO) GetGemPortIdOk() (*int32, bool)`

GetGemPortIdOk returns a tuple with the GemPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemPortId

`func (o *AutoServicePortDTO) SetGemPortId(v int32)`

SetGemPortId sets GemPortId field to given value.

### HasGemPortId

`func (o *AutoServicePortDTO) HasGemPortId() bool`

HasGemPortId returns a boolean if a field has been set.

### GetInboundTrafficProfileId

`func (o *AutoServicePortDTO) GetInboundTrafficProfileId() int32`

GetInboundTrafficProfileId returns the InboundTrafficProfileId field if non-nil, zero value otherwise.

### GetInboundTrafficProfileIdOk

`func (o *AutoServicePortDTO) GetInboundTrafficProfileIdOk() (*int32, bool)`

GetInboundTrafficProfileIdOk returns a tuple with the InboundTrafficProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundTrafficProfileId

`func (o *AutoServicePortDTO) SetInboundTrafficProfileId(v int32)`

SetInboundTrafficProfileId sets InboundTrafficProfileId field to given value.

### HasInboundTrafficProfileId

`func (o *AutoServicePortDTO) HasInboundTrafficProfileId() bool`

HasInboundTrafficProfileId returns a boolean if a field has been set.

### GetInnerVlan

`func (o *AutoServicePortDTO) GetInnerVlan() int32`

GetInnerVlan returns the InnerVlan field if non-nil, zero value otherwise.

### GetInnerVlanOk

`func (o *AutoServicePortDTO) GetInnerVlanOk() (*int32, bool)`

GetInnerVlanOk returns a tuple with the InnerVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerVlan

`func (o *AutoServicePortDTO) SetInnerVlan(v int32)`

SetInnerVlan sets InnerVlan field to given value.

### HasInnerVlan

`func (o *AutoServicePortDTO) HasInnerVlan() bool`

HasInnerVlan returns a boolean if a field has been set.

### GetInnerVlanPriority

`func (o *AutoServicePortDTO) GetInnerVlanPriority() int32`

GetInnerVlanPriority returns the InnerVlanPriority field if non-nil, zero value otherwise.

### GetInnerVlanPriorityOk

`func (o *AutoServicePortDTO) GetInnerVlanPriorityOk() (*int32, bool)`

GetInnerVlanPriorityOk returns a tuple with the InnerVlanPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerVlanPriority

`func (o *AutoServicePortDTO) SetInnerVlanPriority(v int32)`

SetInnerVlanPriority sets InnerVlanPriority field to given value.

### HasInnerVlanPriority

`func (o *AutoServicePortDTO) HasInnerVlanPriority() bool`

HasInnerVlanPriority returns a boolean if a field has been set.

### GetOutboundTrafficProfileId

`func (o *AutoServicePortDTO) GetOutboundTrafficProfileId() int32`

GetOutboundTrafficProfileId returns the OutboundTrafficProfileId field if non-nil, zero value otherwise.

### GetOutboundTrafficProfileIdOk

`func (o *AutoServicePortDTO) GetOutboundTrafficProfileIdOk() (*int32, bool)`

GetOutboundTrafficProfileIdOk returns a tuple with the OutboundTrafficProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundTrafficProfileId

`func (o *AutoServicePortDTO) SetOutboundTrafficProfileId(v int32)`

SetOutboundTrafficProfileId sets OutboundTrafficProfileId field to given value.

### HasOutboundTrafficProfileId

`func (o *AutoServicePortDTO) HasOutboundTrafficProfileId() bool`

HasOutboundTrafficProfileId returns a boolean if a field has been set.

### GetPonPortId

`func (o *AutoServicePortDTO) GetPonPortId() int32`

GetPonPortId returns the PonPortId field if non-nil, zero value otherwise.

### GetPonPortIdOk

`func (o *AutoServicePortDTO) GetPonPortIdOk() (*int32, bool)`

GetPonPortIdOk returns a tuple with the PonPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonPortId

`func (o *AutoServicePortDTO) SetPonPortId(v int32)`

SetPonPortId sets PonPortId field to given value.

### HasPonPortId

`func (o *AutoServicePortDTO) HasPonPortId() bool`

HasPonPortId returns a boolean if a field has been set.

### GetPonPortStr

`func (o *AutoServicePortDTO) GetPonPortStr() string`

GetPonPortStr returns the PonPortStr field if non-nil, zero value otherwise.

### GetPonPortStrOk

`func (o *AutoServicePortDTO) GetPonPortStrOk() (*string, bool)`

GetPonPortStrOk returns a tuple with the PonPortStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonPortStr

`func (o *AutoServicePortDTO) SetPonPortStr(v string)`

SetPonPortStr sets PonPortStr field to given value.

### HasPonPortStr

`func (o *AutoServicePortDTO) HasPonPortStr() bool`

HasPonPortStr returns a boolean if a field has been set.

### GetSvlan

`func (o *AutoServicePortDTO) GetSvlan() int32`

GetSvlan returns the Svlan field if non-nil, zero value otherwise.

### GetSvlanOk

`func (o *AutoServicePortDTO) GetSvlanOk() (*int32, bool)`

GetSvlanOk returns a tuple with the Svlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvlan

`func (o *AutoServicePortDTO) SetSvlan(v int32)`

SetSvlan sets Svlan field to given value.

### HasSvlan

`func (o *AutoServicePortDTO) HasSvlan() bool`

HasSvlan returns a boolean if a field has been set.

### GetTagAction

`func (o *AutoServicePortDTO) GetTagAction() string`

GetTagAction returns the TagAction field if non-nil, zero value otherwise.

### GetTagActionOk

`func (o *AutoServicePortDTO) GetTagActionOk() (*string, bool)`

GetTagActionOk returns a tuple with the TagAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagAction

`func (o *AutoServicePortDTO) SetTagAction(v string)`

SetTagAction sets TagAction field to given value.

### HasTagAction

`func (o *AutoServicePortDTO) HasTagAction() bool`

HasTagAction returns a boolean if a field has been set.

### GetUserVlan

`func (o *AutoServicePortDTO) GetUserVlan() int32`

GetUserVlan returns the UserVlan field if non-nil, zero value otherwise.

### GetUserVlanOk

`func (o *AutoServicePortDTO) GetUserVlanOk() (*int32, bool)`

GetUserVlanOk returns a tuple with the UserVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVlan

`func (o *AutoServicePortDTO) SetUserVlan(v int32)`

SetUserVlan sets UserVlan field to given value.

### HasUserVlan

`func (o *AutoServicePortDTO) HasUserVlan() bool`

HasUserVlan returns a boolean if a field has been set.

### GetUserVlanPriority

`func (o *AutoServicePortDTO) GetUserVlanPriority() int32`

GetUserVlanPriority returns the UserVlanPriority field if non-nil, zero value otherwise.

### GetUserVlanPriorityOk

`func (o *AutoServicePortDTO) GetUserVlanPriorityOk() (*int32, bool)`

GetUserVlanPriorityOk returns a tuple with the UserVlanPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVlanPriority

`func (o *AutoServicePortDTO) SetUserVlanPriority(v int32)`

SetUserVlanPriority sets UserVlanPriority field to given value.

### HasUserVlanPriority

`func (o *AutoServicePortDTO) HasUserVlanPriority() bool`

HasUserVlanPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


