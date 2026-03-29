# ServicePortAddDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminStatus** | **string** | The enable status of the service flows matched by this Service Port. AdminStatus should be a value as follows: DISABLE,ENABLE. Default value: ENABLE. | 
**BatchConfig** | **bool** | Whether it is batch addition.Default value:false | 
**Description** | Pointer to **string** | Description of service port.Description should be 1-32 characters, including letters, numbers, and symbols (-@_:/.). | [optional] 
**EtherType** | Pointer to **string** | EtherType should be a value as follows:NONE,IPV4OE,IPV6OE,PPPOE | [optional] 
**GemPortId** | Pointer to **int32** | GemPortId should be within the range of 1 to 1023 | [optional] 
**InboundTrafficProfileId** | Pointer to **int32** | InboundTrafficProfileId should be within the range of 0 to 512 | [optional] 
**Index** | Pointer to **int32** | ID of service port.Index should be within the range of 1 to 8100 | [optional] 
**InnerVlan** | Pointer to **int32** | InnerVlan should be within the range of 0 to 4095 | [optional] 
**InnerVlanPriority** | Pointer to **int32** | InnerVlanPriority should be within the range of -1 to 7 | [optional] 
**OnuId** | **string** | OnuId should be within the range of 0 to 127 | 
**OutboundTrafficProfileId** | Pointer to **int32** | OutboundTrafficProfileId should be within the range of 0 to 512 | [optional] 
**PonPortId** | Pointer to **int32** | PonPortId should be within the range of 1 to 16 | [optional] 
**PonPortStr** | Pointer to **string** | String form of pon port | [optional] 
**StatisticPerformance** | **string** | The traffic statistics switch status for the service flows matched by this Service Port. StatisticPerformance should be a value as follows: DISABLE,ENABLE. Default value: DISABLE. | 
**Svlan** | Pointer to **int32** | SVlan should be within the range of 1 to 4095 | [optional] 
**TagAction** | Pointer to **string** | TagAction should be a value as follows:DEFAULT,TRANSPARENT,TRANSLATE,TRANSLATE_AND_ADD,ADD_DOUBLE | [optional] 
**UserVlan** | Pointer to **int32** | UserVlan should be within the range of 0 to 4095 | [optional] 
**UserVlanPriority** | Pointer to **int32** | UserVlanPriority should be within the range of -1 to 7 | [optional] 

## Methods

### NewServicePortAddDTO

`func NewServicePortAddDTO(adminStatus string, batchConfig bool, onuId string, statisticPerformance string, ) *ServicePortAddDTO`

NewServicePortAddDTO instantiates a new ServicePortAddDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePortAddDTOWithDefaults

`func NewServicePortAddDTOWithDefaults() *ServicePortAddDTO`

NewServicePortAddDTOWithDefaults instantiates a new ServicePortAddDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminStatus

`func (o *ServicePortAddDTO) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *ServicePortAddDTO) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *ServicePortAddDTO) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.


### GetBatchConfig

`func (o *ServicePortAddDTO) GetBatchConfig() bool`

GetBatchConfig returns the BatchConfig field if non-nil, zero value otherwise.

### GetBatchConfigOk

`func (o *ServicePortAddDTO) GetBatchConfigOk() (*bool, bool)`

GetBatchConfigOk returns a tuple with the BatchConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchConfig

`func (o *ServicePortAddDTO) SetBatchConfig(v bool)`

SetBatchConfig sets BatchConfig field to given value.


### GetDescription

`func (o *ServicePortAddDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicePortAddDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicePortAddDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServicePortAddDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEtherType

`func (o *ServicePortAddDTO) GetEtherType() string`

GetEtherType returns the EtherType field if non-nil, zero value otherwise.

### GetEtherTypeOk

`func (o *ServicePortAddDTO) GetEtherTypeOk() (*string, bool)`

GetEtherTypeOk returns a tuple with the EtherType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherType

`func (o *ServicePortAddDTO) SetEtherType(v string)`

SetEtherType sets EtherType field to given value.

### HasEtherType

`func (o *ServicePortAddDTO) HasEtherType() bool`

HasEtherType returns a boolean if a field has been set.

### GetGemPortId

`func (o *ServicePortAddDTO) GetGemPortId() int32`

GetGemPortId returns the GemPortId field if non-nil, zero value otherwise.

### GetGemPortIdOk

`func (o *ServicePortAddDTO) GetGemPortIdOk() (*int32, bool)`

GetGemPortIdOk returns a tuple with the GemPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemPortId

`func (o *ServicePortAddDTO) SetGemPortId(v int32)`

SetGemPortId sets GemPortId field to given value.

### HasGemPortId

`func (o *ServicePortAddDTO) HasGemPortId() bool`

HasGemPortId returns a boolean if a field has been set.

### GetInboundTrafficProfileId

`func (o *ServicePortAddDTO) GetInboundTrafficProfileId() int32`

GetInboundTrafficProfileId returns the InboundTrafficProfileId field if non-nil, zero value otherwise.

### GetInboundTrafficProfileIdOk

`func (o *ServicePortAddDTO) GetInboundTrafficProfileIdOk() (*int32, bool)`

GetInboundTrafficProfileIdOk returns a tuple with the InboundTrafficProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundTrafficProfileId

`func (o *ServicePortAddDTO) SetInboundTrafficProfileId(v int32)`

SetInboundTrafficProfileId sets InboundTrafficProfileId field to given value.

### HasInboundTrafficProfileId

`func (o *ServicePortAddDTO) HasInboundTrafficProfileId() bool`

HasInboundTrafficProfileId returns a boolean if a field has been set.

### GetIndex

`func (o *ServicePortAddDTO) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ServicePortAddDTO) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ServicePortAddDTO) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ServicePortAddDTO) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetInnerVlan

`func (o *ServicePortAddDTO) GetInnerVlan() int32`

GetInnerVlan returns the InnerVlan field if non-nil, zero value otherwise.

### GetInnerVlanOk

`func (o *ServicePortAddDTO) GetInnerVlanOk() (*int32, bool)`

GetInnerVlanOk returns a tuple with the InnerVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerVlan

`func (o *ServicePortAddDTO) SetInnerVlan(v int32)`

SetInnerVlan sets InnerVlan field to given value.

### HasInnerVlan

`func (o *ServicePortAddDTO) HasInnerVlan() bool`

HasInnerVlan returns a boolean if a field has been set.

### GetInnerVlanPriority

`func (o *ServicePortAddDTO) GetInnerVlanPriority() int32`

GetInnerVlanPriority returns the InnerVlanPriority field if non-nil, zero value otherwise.

### GetInnerVlanPriorityOk

`func (o *ServicePortAddDTO) GetInnerVlanPriorityOk() (*int32, bool)`

GetInnerVlanPriorityOk returns a tuple with the InnerVlanPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerVlanPriority

`func (o *ServicePortAddDTO) SetInnerVlanPriority(v int32)`

SetInnerVlanPriority sets InnerVlanPriority field to given value.

### HasInnerVlanPriority

`func (o *ServicePortAddDTO) HasInnerVlanPriority() bool`

HasInnerVlanPriority returns a boolean if a field has been set.

### GetOnuId

`func (o *ServicePortAddDTO) GetOnuId() string`

GetOnuId returns the OnuId field if non-nil, zero value otherwise.

### GetOnuIdOk

`func (o *ServicePortAddDTO) GetOnuIdOk() (*string, bool)`

GetOnuIdOk returns a tuple with the OnuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuId

`func (o *ServicePortAddDTO) SetOnuId(v string)`

SetOnuId sets OnuId field to given value.


### GetOutboundTrafficProfileId

`func (o *ServicePortAddDTO) GetOutboundTrafficProfileId() int32`

GetOutboundTrafficProfileId returns the OutboundTrafficProfileId field if non-nil, zero value otherwise.

### GetOutboundTrafficProfileIdOk

`func (o *ServicePortAddDTO) GetOutboundTrafficProfileIdOk() (*int32, bool)`

GetOutboundTrafficProfileIdOk returns a tuple with the OutboundTrafficProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundTrafficProfileId

`func (o *ServicePortAddDTO) SetOutboundTrafficProfileId(v int32)`

SetOutboundTrafficProfileId sets OutboundTrafficProfileId field to given value.

### HasOutboundTrafficProfileId

`func (o *ServicePortAddDTO) HasOutboundTrafficProfileId() bool`

HasOutboundTrafficProfileId returns a boolean if a field has been set.

### GetPonPortId

`func (o *ServicePortAddDTO) GetPonPortId() int32`

GetPonPortId returns the PonPortId field if non-nil, zero value otherwise.

### GetPonPortIdOk

`func (o *ServicePortAddDTO) GetPonPortIdOk() (*int32, bool)`

GetPonPortIdOk returns a tuple with the PonPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonPortId

`func (o *ServicePortAddDTO) SetPonPortId(v int32)`

SetPonPortId sets PonPortId field to given value.

### HasPonPortId

`func (o *ServicePortAddDTO) HasPonPortId() bool`

HasPonPortId returns a boolean if a field has been set.

### GetPonPortStr

`func (o *ServicePortAddDTO) GetPonPortStr() string`

GetPonPortStr returns the PonPortStr field if non-nil, zero value otherwise.

### GetPonPortStrOk

`func (o *ServicePortAddDTO) GetPonPortStrOk() (*string, bool)`

GetPonPortStrOk returns a tuple with the PonPortStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonPortStr

`func (o *ServicePortAddDTO) SetPonPortStr(v string)`

SetPonPortStr sets PonPortStr field to given value.

### HasPonPortStr

`func (o *ServicePortAddDTO) HasPonPortStr() bool`

HasPonPortStr returns a boolean if a field has been set.

### GetStatisticPerformance

`func (o *ServicePortAddDTO) GetStatisticPerformance() string`

GetStatisticPerformance returns the StatisticPerformance field if non-nil, zero value otherwise.

### GetStatisticPerformanceOk

`func (o *ServicePortAddDTO) GetStatisticPerformanceOk() (*string, bool)`

GetStatisticPerformanceOk returns a tuple with the StatisticPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticPerformance

`func (o *ServicePortAddDTO) SetStatisticPerformance(v string)`

SetStatisticPerformance sets StatisticPerformance field to given value.


### GetSvlan

`func (o *ServicePortAddDTO) GetSvlan() int32`

GetSvlan returns the Svlan field if non-nil, zero value otherwise.

### GetSvlanOk

`func (o *ServicePortAddDTO) GetSvlanOk() (*int32, bool)`

GetSvlanOk returns a tuple with the Svlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvlan

`func (o *ServicePortAddDTO) SetSvlan(v int32)`

SetSvlan sets Svlan field to given value.

### HasSvlan

`func (o *ServicePortAddDTO) HasSvlan() bool`

HasSvlan returns a boolean if a field has been set.

### GetTagAction

`func (o *ServicePortAddDTO) GetTagAction() string`

GetTagAction returns the TagAction field if non-nil, zero value otherwise.

### GetTagActionOk

`func (o *ServicePortAddDTO) GetTagActionOk() (*string, bool)`

GetTagActionOk returns a tuple with the TagAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagAction

`func (o *ServicePortAddDTO) SetTagAction(v string)`

SetTagAction sets TagAction field to given value.

### HasTagAction

`func (o *ServicePortAddDTO) HasTagAction() bool`

HasTagAction returns a boolean if a field has been set.

### GetUserVlan

`func (o *ServicePortAddDTO) GetUserVlan() int32`

GetUserVlan returns the UserVlan field if non-nil, zero value otherwise.

### GetUserVlanOk

`func (o *ServicePortAddDTO) GetUserVlanOk() (*int32, bool)`

GetUserVlanOk returns a tuple with the UserVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVlan

`func (o *ServicePortAddDTO) SetUserVlan(v int32)`

SetUserVlan sets UserVlan field to given value.

### HasUserVlan

`func (o *ServicePortAddDTO) HasUserVlan() bool`

HasUserVlan returns a boolean if a field has been set.

### GetUserVlanPriority

`func (o *ServicePortAddDTO) GetUserVlanPriority() int32`

GetUserVlanPriority returns the UserVlanPriority field if non-nil, zero value otherwise.

### GetUserVlanPriorityOk

`func (o *ServicePortAddDTO) GetUserVlanPriorityOk() (*int32, bool)`

GetUserVlanPriorityOk returns a tuple with the UserVlanPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVlanPriority

`func (o *ServicePortAddDTO) SetUserVlanPriority(v int32)`

SetUserVlanPriority sets UserVlanPriority field to given value.

### HasUserVlanPriority

`func (o *ServicePortAddDTO) HasUserVlanPriority() bool`

HasUserVlanPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


