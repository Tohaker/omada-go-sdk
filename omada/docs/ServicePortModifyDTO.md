# ServicePortModifyDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminStatus** | **string** | The enable status of the service flows matched by this Service Port. AdminStatus should be a value as follows: DISABLE,ENABLE. Default value: ENABLE. | 
**Description** | Pointer to **string** | Description of service port.Description should contain 1 to 32 characters including numbers, Upper and lower letters, -@_:/. . | [optional] 
**EtherType** | Pointer to **string** | EtherType should be a value as follows:NONE,IPV4OE,IPV6OE,PPPOE | [optional] 
**GemPortId** | Pointer to **int32** | GemPortId should be within the range of 1 to 1023 | [optional] 
**InboundTrafficProfileId** | Pointer to **int32** | InboundTrafficProfileId should be within the range of 0 to 512 | [optional] 
**Index** | **int32** | Service port ID should be within the range of 1 to 8100 | 
**InnerVlan** | Pointer to **int32** | InnerVlan should be within the range of 0 to 4095 | [optional] 
**InnerVlanPriority** | Pointer to **int32** | InnerVlanPriority should be within the range of -1 to 7 | [optional] 
**OnuId** | **int32** | Onu ID should be within the range of 0 to 127 | 
**OutboundTrafficProfileId** | Pointer to **int32** | OutboundTrafficProfileId should be within the range of 0 to 512 | [optional] 
**PonPortId** | Pointer to **int32** | PonPortId should be within the range of 1 to 16 | [optional] 
**PonPortStr** | Pointer to **string** | String form of pon port | [optional] 
**StatisticPerformance** | **string** | The traffic statistics switch status for the service flows matched by this Service Port. StatisticPerformance should be a value as follows: DISABLE,ENABLE. Default value: DISABLE. | 
**Svlan** | Pointer to **int32** | SVlan should be within the range of 1 to 4095 | [optional] 
**TagAction** | Pointer to **string** | TagAction should be a value as follows:DEFAULT,TRANSPARENT,TRANSLATE,TRANSLATE_AND_ADD,ADD_DOUBLE | [optional] 
**UserVlan** | Pointer to **int32** | UserVlan should be within the range of 0 to 4095 | [optional] 
**UserVlanPriority** | Pointer to **int32** | UserVlanPriority should be within the range of -1 to 7 | [optional] 

## Methods

### NewServicePortModifyDTO

`func NewServicePortModifyDTO(adminStatus string, index int32, onuId int32, statisticPerformance string, ) *ServicePortModifyDTO`

NewServicePortModifyDTO instantiates a new ServicePortModifyDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePortModifyDTOWithDefaults

`func NewServicePortModifyDTOWithDefaults() *ServicePortModifyDTO`

NewServicePortModifyDTOWithDefaults instantiates a new ServicePortModifyDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminStatus

`func (o *ServicePortModifyDTO) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *ServicePortModifyDTO) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *ServicePortModifyDTO) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.


### GetDescription

`func (o *ServicePortModifyDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicePortModifyDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicePortModifyDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServicePortModifyDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEtherType

`func (o *ServicePortModifyDTO) GetEtherType() string`

GetEtherType returns the EtherType field if non-nil, zero value otherwise.

### GetEtherTypeOk

`func (o *ServicePortModifyDTO) GetEtherTypeOk() (*string, bool)`

GetEtherTypeOk returns a tuple with the EtherType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherType

`func (o *ServicePortModifyDTO) SetEtherType(v string)`

SetEtherType sets EtherType field to given value.

### HasEtherType

`func (o *ServicePortModifyDTO) HasEtherType() bool`

HasEtherType returns a boolean if a field has been set.

### GetGemPortId

`func (o *ServicePortModifyDTO) GetGemPortId() int32`

GetGemPortId returns the GemPortId field if non-nil, zero value otherwise.

### GetGemPortIdOk

`func (o *ServicePortModifyDTO) GetGemPortIdOk() (*int32, bool)`

GetGemPortIdOk returns a tuple with the GemPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemPortId

`func (o *ServicePortModifyDTO) SetGemPortId(v int32)`

SetGemPortId sets GemPortId field to given value.

### HasGemPortId

`func (o *ServicePortModifyDTO) HasGemPortId() bool`

HasGemPortId returns a boolean if a field has been set.

### GetInboundTrafficProfileId

`func (o *ServicePortModifyDTO) GetInboundTrafficProfileId() int32`

GetInboundTrafficProfileId returns the InboundTrafficProfileId field if non-nil, zero value otherwise.

### GetInboundTrafficProfileIdOk

`func (o *ServicePortModifyDTO) GetInboundTrafficProfileIdOk() (*int32, bool)`

GetInboundTrafficProfileIdOk returns a tuple with the InboundTrafficProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundTrafficProfileId

`func (o *ServicePortModifyDTO) SetInboundTrafficProfileId(v int32)`

SetInboundTrafficProfileId sets InboundTrafficProfileId field to given value.

### HasInboundTrafficProfileId

`func (o *ServicePortModifyDTO) HasInboundTrafficProfileId() bool`

HasInboundTrafficProfileId returns a boolean if a field has been set.

### GetIndex

`func (o *ServicePortModifyDTO) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ServicePortModifyDTO) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ServicePortModifyDTO) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetInnerVlan

`func (o *ServicePortModifyDTO) GetInnerVlan() int32`

GetInnerVlan returns the InnerVlan field if non-nil, zero value otherwise.

### GetInnerVlanOk

`func (o *ServicePortModifyDTO) GetInnerVlanOk() (*int32, bool)`

GetInnerVlanOk returns a tuple with the InnerVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerVlan

`func (o *ServicePortModifyDTO) SetInnerVlan(v int32)`

SetInnerVlan sets InnerVlan field to given value.

### HasInnerVlan

`func (o *ServicePortModifyDTO) HasInnerVlan() bool`

HasInnerVlan returns a boolean if a field has been set.

### GetInnerVlanPriority

`func (o *ServicePortModifyDTO) GetInnerVlanPriority() int32`

GetInnerVlanPriority returns the InnerVlanPriority field if non-nil, zero value otherwise.

### GetInnerVlanPriorityOk

`func (o *ServicePortModifyDTO) GetInnerVlanPriorityOk() (*int32, bool)`

GetInnerVlanPriorityOk returns a tuple with the InnerVlanPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerVlanPriority

`func (o *ServicePortModifyDTO) SetInnerVlanPriority(v int32)`

SetInnerVlanPriority sets InnerVlanPriority field to given value.

### HasInnerVlanPriority

`func (o *ServicePortModifyDTO) HasInnerVlanPriority() bool`

HasInnerVlanPriority returns a boolean if a field has been set.

### GetOnuId

`func (o *ServicePortModifyDTO) GetOnuId() int32`

GetOnuId returns the OnuId field if non-nil, zero value otherwise.

### GetOnuIdOk

`func (o *ServicePortModifyDTO) GetOnuIdOk() (*int32, bool)`

GetOnuIdOk returns a tuple with the OnuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuId

`func (o *ServicePortModifyDTO) SetOnuId(v int32)`

SetOnuId sets OnuId field to given value.


### GetOutboundTrafficProfileId

`func (o *ServicePortModifyDTO) GetOutboundTrafficProfileId() int32`

GetOutboundTrafficProfileId returns the OutboundTrafficProfileId field if non-nil, zero value otherwise.

### GetOutboundTrafficProfileIdOk

`func (o *ServicePortModifyDTO) GetOutboundTrafficProfileIdOk() (*int32, bool)`

GetOutboundTrafficProfileIdOk returns a tuple with the OutboundTrafficProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundTrafficProfileId

`func (o *ServicePortModifyDTO) SetOutboundTrafficProfileId(v int32)`

SetOutboundTrafficProfileId sets OutboundTrafficProfileId field to given value.

### HasOutboundTrafficProfileId

`func (o *ServicePortModifyDTO) HasOutboundTrafficProfileId() bool`

HasOutboundTrafficProfileId returns a boolean if a field has been set.

### GetPonPortId

`func (o *ServicePortModifyDTO) GetPonPortId() int32`

GetPonPortId returns the PonPortId field if non-nil, zero value otherwise.

### GetPonPortIdOk

`func (o *ServicePortModifyDTO) GetPonPortIdOk() (*int32, bool)`

GetPonPortIdOk returns a tuple with the PonPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonPortId

`func (o *ServicePortModifyDTO) SetPonPortId(v int32)`

SetPonPortId sets PonPortId field to given value.

### HasPonPortId

`func (o *ServicePortModifyDTO) HasPonPortId() bool`

HasPonPortId returns a boolean if a field has been set.

### GetPonPortStr

`func (o *ServicePortModifyDTO) GetPonPortStr() string`

GetPonPortStr returns the PonPortStr field if non-nil, zero value otherwise.

### GetPonPortStrOk

`func (o *ServicePortModifyDTO) GetPonPortStrOk() (*string, bool)`

GetPonPortStrOk returns a tuple with the PonPortStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonPortStr

`func (o *ServicePortModifyDTO) SetPonPortStr(v string)`

SetPonPortStr sets PonPortStr field to given value.

### HasPonPortStr

`func (o *ServicePortModifyDTO) HasPonPortStr() bool`

HasPonPortStr returns a boolean if a field has been set.

### GetStatisticPerformance

`func (o *ServicePortModifyDTO) GetStatisticPerformance() string`

GetStatisticPerformance returns the StatisticPerformance field if non-nil, zero value otherwise.

### GetStatisticPerformanceOk

`func (o *ServicePortModifyDTO) GetStatisticPerformanceOk() (*string, bool)`

GetStatisticPerformanceOk returns a tuple with the StatisticPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticPerformance

`func (o *ServicePortModifyDTO) SetStatisticPerformance(v string)`

SetStatisticPerformance sets StatisticPerformance field to given value.


### GetSvlan

`func (o *ServicePortModifyDTO) GetSvlan() int32`

GetSvlan returns the Svlan field if non-nil, zero value otherwise.

### GetSvlanOk

`func (o *ServicePortModifyDTO) GetSvlanOk() (*int32, bool)`

GetSvlanOk returns a tuple with the Svlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvlan

`func (o *ServicePortModifyDTO) SetSvlan(v int32)`

SetSvlan sets Svlan field to given value.

### HasSvlan

`func (o *ServicePortModifyDTO) HasSvlan() bool`

HasSvlan returns a boolean if a field has been set.

### GetTagAction

`func (o *ServicePortModifyDTO) GetTagAction() string`

GetTagAction returns the TagAction field if non-nil, zero value otherwise.

### GetTagActionOk

`func (o *ServicePortModifyDTO) GetTagActionOk() (*string, bool)`

GetTagActionOk returns a tuple with the TagAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagAction

`func (o *ServicePortModifyDTO) SetTagAction(v string)`

SetTagAction sets TagAction field to given value.

### HasTagAction

`func (o *ServicePortModifyDTO) HasTagAction() bool`

HasTagAction returns a boolean if a field has been set.

### GetUserVlan

`func (o *ServicePortModifyDTO) GetUserVlan() int32`

GetUserVlan returns the UserVlan field if non-nil, zero value otherwise.

### GetUserVlanOk

`func (o *ServicePortModifyDTO) GetUserVlanOk() (*int32, bool)`

GetUserVlanOk returns a tuple with the UserVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVlan

`func (o *ServicePortModifyDTO) SetUserVlan(v int32)`

SetUserVlan sets UserVlan field to given value.

### HasUserVlan

`func (o *ServicePortModifyDTO) HasUserVlan() bool`

HasUserVlan returns a boolean if a field has been set.

### GetUserVlanPriority

`func (o *ServicePortModifyDTO) GetUserVlanPriority() int32`

GetUserVlanPriority returns the UserVlanPriority field if non-nil, zero value otherwise.

### GetUserVlanPriorityOk

`func (o *ServicePortModifyDTO) GetUserVlanPriorityOk() (*int32, bool)`

GetUserVlanPriorityOk returns a tuple with the UserVlanPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVlanPriority

`func (o *ServicePortModifyDTO) SetUserVlanPriority(v int32)`

SetUserVlanPriority sets UserVlanPriority field to given value.

### HasUserVlanPriority

`func (o *ServicePortModifyDTO) HasUserVlanPriority() bool`

HasUserVlanPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


