# ServicePortVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveStatus** | Pointer to **string** | Whether this service port is active.ActiveStatus should be a value as follows:ACTIVE,INACTIVE | [optional] 
**AdminStatus** | **string** | The enable status of the service flows matched by this Service Port. AdminStatus should be a value as follows: DISABLE,ENABLE. Default value: ENABLE. | 
**Description** | **string** | Description of service port.Description should contain 1 to 32 characters including numbers, Upper and lower letters, -@_:/. . | 
**EtherType** | Pointer to **string** | EtherType should be a value as follows:NONE,IPV4OE,IPV6OE,PPPOE | [optional] 
**GemPortId** | Pointer to **int32** | GemPortId should be within the range of 1 to 1023 | [optional] 
**InboundTrafficProfileId** | Pointer to **string** | InboundTrafficProfileId should be within the range of 0 to 512 | [optional] 
**Index** | **int32** | ID of service port.Index should be within the range of 1 to 8100 | 
**InnerVlan** | Pointer to **int32** | InnerVlan should be within the range of 0 to 4095 | [optional] 
**InnerVlanPriority** | Pointer to **string** | InnerVlanPriority should be within the range of -1 to 7 | [optional] 
**OnuId** | **int32** | OnuId should be within the range of 0 to 127 | 
**OutboundTrafficProfileId** | Pointer to **string** | OutboundTrafficProfileId should be within the range of 0 to 512 | [optional] 
**PonPortId** | Pointer to **int32** | PonPortId should be within the range of 1 to 16 | [optional] 
**StatisticPerformance** | **string** | The traffic statistics switch status for the service flows matched by this Service Port. StatisticPerformance should be a value as follows: DISABLE,ENABLE. Default value: DISABLE. | 
**Svlan** | Pointer to **int32** | SVlan should be within the range of 1 to 4095 | [optional] 
**TagAction** | Pointer to **string** | TagAction should be a value as follows:DEFAULT,TRANSPARENT,TRANSLATE,TRANSLATE_AND_ADD,ADD_DOUBLE | [optional] 
**UserVlan** | Pointer to **int32** | UserVlan should be within the range of 0 to 4095 | [optional] 
**UserVlanPriority** | Pointer to **string** | UserVlanPriority should be within the range of -1 to 7 | [optional] 

## Methods

### NewServicePortVO

`func NewServicePortVO(adminStatus string, description string, index int32, onuId int32, statisticPerformance string, ) *ServicePortVO`

NewServicePortVO instantiates a new ServicePortVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePortVOWithDefaults

`func NewServicePortVOWithDefaults() *ServicePortVO`

NewServicePortVOWithDefaults instantiates a new ServicePortVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveStatus

`func (o *ServicePortVO) GetActiveStatus() string`

GetActiveStatus returns the ActiveStatus field if non-nil, zero value otherwise.

### GetActiveStatusOk

`func (o *ServicePortVO) GetActiveStatusOk() (*string, bool)`

GetActiveStatusOk returns a tuple with the ActiveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveStatus

`func (o *ServicePortVO) SetActiveStatus(v string)`

SetActiveStatus sets ActiveStatus field to given value.

### HasActiveStatus

`func (o *ServicePortVO) HasActiveStatus() bool`

HasActiveStatus returns a boolean if a field has been set.

### GetAdminStatus

`func (o *ServicePortVO) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *ServicePortVO) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *ServicePortVO) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.


### GetDescription

`func (o *ServicePortVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicePortVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicePortVO) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEtherType

`func (o *ServicePortVO) GetEtherType() string`

GetEtherType returns the EtherType field if non-nil, zero value otherwise.

### GetEtherTypeOk

`func (o *ServicePortVO) GetEtherTypeOk() (*string, bool)`

GetEtherTypeOk returns a tuple with the EtherType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherType

`func (o *ServicePortVO) SetEtherType(v string)`

SetEtherType sets EtherType field to given value.

### HasEtherType

`func (o *ServicePortVO) HasEtherType() bool`

HasEtherType returns a boolean if a field has been set.

### GetGemPortId

`func (o *ServicePortVO) GetGemPortId() int32`

GetGemPortId returns the GemPortId field if non-nil, zero value otherwise.

### GetGemPortIdOk

`func (o *ServicePortVO) GetGemPortIdOk() (*int32, bool)`

GetGemPortIdOk returns a tuple with the GemPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemPortId

`func (o *ServicePortVO) SetGemPortId(v int32)`

SetGemPortId sets GemPortId field to given value.

### HasGemPortId

`func (o *ServicePortVO) HasGemPortId() bool`

HasGemPortId returns a boolean if a field has been set.

### GetInboundTrafficProfileId

`func (o *ServicePortVO) GetInboundTrafficProfileId() string`

GetInboundTrafficProfileId returns the InboundTrafficProfileId field if non-nil, zero value otherwise.

### GetInboundTrafficProfileIdOk

`func (o *ServicePortVO) GetInboundTrafficProfileIdOk() (*string, bool)`

GetInboundTrafficProfileIdOk returns a tuple with the InboundTrafficProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundTrafficProfileId

`func (o *ServicePortVO) SetInboundTrafficProfileId(v string)`

SetInboundTrafficProfileId sets InboundTrafficProfileId field to given value.

### HasInboundTrafficProfileId

`func (o *ServicePortVO) HasInboundTrafficProfileId() bool`

HasInboundTrafficProfileId returns a boolean if a field has been set.

### GetIndex

`func (o *ServicePortVO) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ServicePortVO) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ServicePortVO) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetInnerVlan

`func (o *ServicePortVO) GetInnerVlan() int32`

GetInnerVlan returns the InnerVlan field if non-nil, zero value otherwise.

### GetInnerVlanOk

`func (o *ServicePortVO) GetInnerVlanOk() (*int32, bool)`

GetInnerVlanOk returns a tuple with the InnerVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerVlan

`func (o *ServicePortVO) SetInnerVlan(v int32)`

SetInnerVlan sets InnerVlan field to given value.

### HasInnerVlan

`func (o *ServicePortVO) HasInnerVlan() bool`

HasInnerVlan returns a boolean if a field has been set.

### GetInnerVlanPriority

`func (o *ServicePortVO) GetInnerVlanPriority() string`

GetInnerVlanPriority returns the InnerVlanPriority field if non-nil, zero value otherwise.

### GetInnerVlanPriorityOk

`func (o *ServicePortVO) GetInnerVlanPriorityOk() (*string, bool)`

GetInnerVlanPriorityOk returns a tuple with the InnerVlanPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerVlanPriority

`func (o *ServicePortVO) SetInnerVlanPriority(v string)`

SetInnerVlanPriority sets InnerVlanPriority field to given value.

### HasInnerVlanPriority

`func (o *ServicePortVO) HasInnerVlanPriority() bool`

HasInnerVlanPriority returns a boolean if a field has been set.

### GetOnuId

`func (o *ServicePortVO) GetOnuId() int32`

GetOnuId returns the OnuId field if non-nil, zero value otherwise.

### GetOnuIdOk

`func (o *ServicePortVO) GetOnuIdOk() (*int32, bool)`

GetOnuIdOk returns a tuple with the OnuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuId

`func (o *ServicePortVO) SetOnuId(v int32)`

SetOnuId sets OnuId field to given value.


### GetOutboundTrafficProfileId

`func (o *ServicePortVO) GetOutboundTrafficProfileId() string`

GetOutboundTrafficProfileId returns the OutboundTrafficProfileId field if non-nil, zero value otherwise.

### GetOutboundTrafficProfileIdOk

`func (o *ServicePortVO) GetOutboundTrafficProfileIdOk() (*string, bool)`

GetOutboundTrafficProfileIdOk returns a tuple with the OutboundTrafficProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundTrafficProfileId

`func (o *ServicePortVO) SetOutboundTrafficProfileId(v string)`

SetOutboundTrafficProfileId sets OutboundTrafficProfileId field to given value.

### HasOutboundTrafficProfileId

`func (o *ServicePortVO) HasOutboundTrafficProfileId() bool`

HasOutboundTrafficProfileId returns a boolean if a field has been set.

### GetPonPortId

`func (o *ServicePortVO) GetPonPortId() int32`

GetPonPortId returns the PonPortId field if non-nil, zero value otherwise.

### GetPonPortIdOk

`func (o *ServicePortVO) GetPonPortIdOk() (*int32, bool)`

GetPonPortIdOk returns a tuple with the PonPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonPortId

`func (o *ServicePortVO) SetPonPortId(v int32)`

SetPonPortId sets PonPortId field to given value.

### HasPonPortId

`func (o *ServicePortVO) HasPonPortId() bool`

HasPonPortId returns a boolean if a field has been set.

### GetStatisticPerformance

`func (o *ServicePortVO) GetStatisticPerformance() string`

GetStatisticPerformance returns the StatisticPerformance field if non-nil, zero value otherwise.

### GetStatisticPerformanceOk

`func (o *ServicePortVO) GetStatisticPerformanceOk() (*string, bool)`

GetStatisticPerformanceOk returns a tuple with the StatisticPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticPerformance

`func (o *ServicePortVO) SetStatisticPerformance(v string)`

SetStatisticPerformance sets StatisticPerformance field to given value.


### GetSvlan

`func (o *ServicePortVO) GetSvlan() int32`

GetSvlan returns the Svlan field if non-nil, zero value otherwise.

### GetSvlanOk

`func (o *ServicePortVO) GetSvlanOk() (*int32, bool)`

GetSvlanOk returns a tuple with the Svlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvlan

`func (o *ServicePortVO) SetSvlan(v int32)`

SetSvlan sets Svlan field to given value.

### HasSvlan

`func (o *ServicePortVO) HasSvlan() bool`

HasSvlan returns a boolean if a field has been set.

### GetTagAction

`func (o *ServicePortVO) GetTagAction() string`

GetTagAction returns the TagAction field if non-nil, zero value otherwise.

### GetTagActionOk

`func (o *ServicePortVO) GetTagActionOk() (*string, bool)`

GetTagActionOk returns a tuple with the TagAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagAction

`func (o *ServicePortVO) SetTagAction(v string)`

SetTagAction sets TagAction field to given value.

### HasTagAction

`func (o *ServicePortVO) HasTagAction() bool`

HasTagAction returns a boolean if a field has been set.

### GetUserVlan

`func (o *ServicePortVO) GetUserVlan() int32`

GetUserVlan returns the UserVlan field if non-nil, zero value otherwise.

### GetUserVlanOk

`func (o *ServicePortVO) GetUserVlanOk() (*int32, bool)`

GetUserVlanOk returns a tuple with the UserVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVlan

`func (o *ServicePortVO) SetUserVlan(v int32)`

SetUserVlan sets UserVlan field to given value.

### HasUserVlan

`func (o *ServicePortVO) HasUserVlan() bool`

HasUserVlan returns a boolean if a field has been set.

### GetUserVlanPriority

`func (o *ServicePortVO) GetUserVlanPriority() string`

GetUserVlanPriority returns the UserVlanPriority field if non-nil, zero value otherwise.

### GetUserVlanPriorityOk

`func (o *ServicePortVO) GetUserVlanPriorityOk() (*string, bool)`

GetUserVlanPriorityOk returns a tuple with the UserVlanPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVlanPriority

`func (o *ServicePortVO) SetUserVlanPriority(v string)`

SetUserVlanPriority sets UserVlanPriority field to given value.

### HasUserVlanPriority

`func (o *ServicePortVO) HasUserVlanPriority() bool`

HasUserVlanPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


