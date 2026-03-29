# ServicePortProfileListDetailDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminStatus** | **string** | The admin status   | 
**CreationMode** | **string** | The creation mode  | 
**Description** | **string** | The description should be 1-32 characters, including uppercase and lowercase letters, numbers, and underscores. | 
**EntriesNumber** | **int32** | Number of service ports. The entriesNumber should be within the range of 0 to 127. | 
**EtherType** | **string** | The Ether Type value  | 
**GemId** | **int32** | The value of the GEM ID. The gemId should be within the range of 1-1023 | 
**InUse** | **string** | Indicate whether the profile has been used  | 
**InboundTrafficProfileId** | **int32** | The inboundTrafficProfileId should be within the range of 1 to 1023 (Traffic profiles that already exist on the system). Using -1 means that the item is not configured, and is displayed as --. | 
**InnerVlan** | **int32** | The innerVlan should be within the range of 1 to 4094.This configuration item appears if and only if Translate-And-Add/Add/Add-Double is selected for tag_action, using -1 means that the item is not configured, and the effect will be displayed as --. | 
**InnerVlanPriority** | **int32** | The innerVlanPriority should be within the range of 0 to 7. Only takes effect when Translate-And-Add/Add/Add-Double is selected for tag_action. Using -1 means that the item is not configured, and is displayed as --. | 
**OutboundTrafficProfileId** | **int32** | The outboundTrafficProfileId should be within the range of 1 to 1023 (Traffic profiles that already exist in the system). Using -1 means that the item is not configured, and is displayed as --. | 
**ServicePortId** | **string** | Display Service Port Entry ID. The servicePortId should be within the range of 1 to 127. | 
**ServicePortProfileId** | **string** | Service template ID. The servicePortProfileId should be within the range of 1 to 127. | 
**StatisticPerformance** | **string** | The statistic performance  | 
**Svlan** | **int32** | The value of the SVLAN. The svlan should be within the range of 1 to 4094. | 
**TagAction** | **string** | TAG Action  | 
**UserVlan** | **int32** | The value of the User VLAN. The userVlan should be within the range of 1 to 4094. | 
**UserVlanPriority** | **int32** | The value of the User VLAN Priority. The userVlanPriority should be within the range of 0 to 7.A value of -1 means that the item is not configured. | 

## Methods

### NewServicePortProfileListDetailDTO

`func NewServicePortProfileListDetailDTO(adminStatus string, creationMode string, description string, entriesNumber int32, etherType string, gemId int32, inUse string, inboundTrafficProfileId int32, innerVlan int32, innerVlanPriority int32, outboundTrafficProfileId int32, servicePortId string, servicePortProfileId string, statisticPerformance string, svlan int32, tagAction string, userVlan int32, userVlanPriority int32, ) *ServicePortProfileListDetailDTO`

NewServicePortProfileListDetailDTO instantiates a new ServicePortProfileListDetailDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePortProfileListDetailDTOWithDefaults

`func NewServicePortProfileListDetailDTOWithDefaults() *ServicePortProfileListDetailDTO`

NewServicePortProfileListDetailDTOWithDefaults instantiates a new ServicePortProfileListDetailDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminStatus

`func (o *ServicePortProfileListDetailDTO) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *ServicePortProfileListDetailDTO) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *ServicePortProfileListDetailDTO) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.


### GetCreationMode

`func (o *ServicePortProfileListDetailDTO) GetCreationMode() string`

GetCreationMode returns the CreationMode field if non-nil, zero value otherwise.

### GetCreationModeOk

`func (o *ServicePortProfileListDetailDTO) GetCreationModeOk() (*string, bool)`

GetCreationModeOk returns a tuple with the CreationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationMode

`func (o *ServicePortProfileListDetailDTO) SetCreationMode(v string)`

SetCreationMode sets CreationMode field to given value.


### GetDescription

`func (o *ServicePortProfileListDetailDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicePortProfileListDetailDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicePortProfileListDetailDTO) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEntriesNumber

`func (o *ServicePortProfileListDetailDTO) GetEntriesNumber() int32`

GetEntriesNumber returns the EntriesNumber field if non-nil, zero value otherwise.

### GetEntriesNumberOk

`func (o *ServicePortProfileListDetailDTO) GetEntriesNumberOk() (*int32, bool)`

GetEntriesNumberOk returns a tuple with the EntriesNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntriesNumber

`func (o *ServicePortProfileListDetailDTO) SetEntriesNumber(v int32)`

SetEntriesNumber sets EntriesNumber field to given value.


### GetEtherType

`func (o *ServicePortProfileListDetailDTO) GetEtherType() string`

GetEtherType returns the EtherType field if non-nil, zero value otherwise.

### GetEtherTypeOk

`func (o *ServicePortProfileListDetailDTO) GetEtherTypeOk() (*string, bool)`

GetEtherTypeOk returns a tuple with the EtherType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherType

`func (o *ServicePortProfileListDetailDTO) SetEtherType(v string)`

SetEtherType sets EtherType field to given value.


### GetGemId

`func (o *ServicePortProfileListDetailDTO) GetGemId() int32`

GetGemId returns the GemId field if non-nil, zero value otherwise.

### GetGemIdOk

`func (o *ServicePortProfileListDetailDTO) GetGemIdOk() (*int32, bool)`

GetGemIdOk returns a tuple with the GemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemId

`func (o *ServicePortProfileListDetailDTO) SetGemId(v int32)`

SetGemId sets GemId field to given value.


### GetInUse

`func (o *ServicePortProfileListDetailDTO) GetInUse() string`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *ServicePortProfileListDetailDTO) GetInUseOk() (*string, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *ServicePortProfileListDetailDTO) SetInUse(v string)`

SetInUse sets InUse field to given value.


### GetInboundTrafficProfileId

`func (o *ServicePortProfileListDetailDTO) GetInboundTrafficProfileId() int32`

GetInboundTrafficProfileId returns the InboundTrafficProfileId field if non-nil, zero value otherwise.

### GetInboundTrafficProfileIdOk

`func (o *ServicePortProfileListDetailDTO) GetInboundTrafficProfileIdOk() (*int32, bool)`

GetInboundTrafficProfileIdOk returns a tuple with the InboundTrafficProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundTrafficProfileId

`func (o *ServicePortProfileListDetailDTO) SetInboundTrafficProfileId(v int32)`

SetInboundTrafficProfileId sets InboundTrafficProfileId field to given value.


### GetInnerVlan

`func (o *ServicePortProfileListDetailDTO) GetInnerVlan() int32`

GetInnerVlan returns the InnerVlan field if non-nil, zero value otherwise.

### GetInnerVlanOk

`func (o *ServicePortProfileListDetailDTO) GetInnerVlanOk() (*int32, bool)`

GetInnerVlanOk returns a tuple with the InnerVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerVlan

`func (o *ServicePortProfileListDetailDTO) SetInnerVlan(v int32)`

SetInnerVlan sets InnerVlan field to given value.


### GetInnerVlanPriority

`func (o *ServicePortProfileListDetailDTO) GetInnerVlanPriority() int32`

GetInnerVlanPriority returns the InnerVlanPriority field if non-nil, zero value otherwise.

### GetInnerVlanPriorityOk

`func (o *ServicePortProfileListDetailDTO) GetInnerVlanPriorityOk() (*int32, bool)`

GetInnerVlanPriorityOk returns a tuple with the InnerVlanPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerVlanPriority

`func (o *ServicePortProfileListDetailDTO) SetInnerVlanPriority(v int32)`

SetInnerVlanPriority sets InnerVlanPriority field to given value.


### GetOutboundTrafficProfileId

`func (o *ServicePortProfileListDetailDTO) GetOutboundTrafficProfileId() int32`

GetOutboundTrafficProfileId returns the OutboundTrafficProfileId field if non-nil, zero value otherwise.

### GetOutboundTrafficProfileIdOk

`func (o *ServicePortProfileListDetailDTO) GetOutboundTrafficProfileIdOk() (*int32, bool)`

GetOutboundTrafficProfileIdOk returns a tuple with the OutboundTrafficProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundTrafficProfileId

`func (o *ServicePortProfileListDetailDTO) SetOutboundTrafficProfileId(v int32)`

SetOutboundTrafficProfileId sets OutboundTrafficProfileId field to given value.


### GetServicePortId

`func (o *ServicePortProfileListDetailDTO) GetServicePortId() string`

GetServicePortId returns the ServicePortId field if non-nil, zero value otherwise.

### GetServicePortIdOk

`func (o *ServicePortProfileListDetailDTO) GetServicePortIdOk() (*string, bool)`

GetServicePortIdOk returns a tuple with the ServicePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePortId

`func (o *ServicePortProfileListDetailDTO) SetServicePortId(v string)`

SetServicePortId sets ServicePortId field to given value.


### GetServicePortProfileId

`func (o *ServicePortProfileListDetailDTO) GetServicePortProfileId() string`

GetServicePortProfileId returns the ServicePortProfileId field if non-nil, zero value otherwise.

### GetServicePortProfileIdOk

`func (o *ServicePortProfileListDetailDTO) GetServicePortProfileIdOk() (*string, bool)`

GetServicePortProfileIdOk returns a tuple with the ServicePortProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePortProfileId

`func (o *ServicePortProfileListDetailDTO) SetServicePortProfileId(v string)`

SetServicePortProfileId sets ServicePortProfileId field to given value.


### GetStatisticPerformance

`func (o *ServicePortProfileListDetailDTO) GetStatisticPerformance() string`

GetStatisticPerformance returns the StatisticPerformance field if non-nil, zero value otherwise.

### GetStatisticPerformanceOk

`func (o *ServicePortProfileListDetailDTO) GetStatisticPerformanceOk() (*string, bool)`

GetStatisticPerformanceOk returns a tuple with the StatisticPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticPerformance

`func (o *ServicePortProfileListDetailDTO) SetStatisticPerformance(v string)`

SetStatisticPerformance sets StatisticPerformance field to given value.


### GetSvlan

`func (o *ServicePortProfileListDetailDTO) GetSvlan() int32`

GetSvlan returns the Svlan field if non-nil, zero value otherwise.

### GetSvlanOk

`func (o *ServicePortProfileListDetailDTO) GetSvlanOk() (*int32, bool)`

GetSvlanOk returns a tuple with the Svlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvlan

`func (o *ServicePortProfileListDetailDTO) SetSvlan(v int32)`

SetSvlan sets Svlan field to given value.


### GetTagAction

`func (o *ServicePortProfileListDetailDTO) GetTagAction() string`

GetTagAction returns the TagAction field if non-nil, zero value otherwise.

### GetTagActionOk

`func (o *ServicePortProfileListDetailDTO) GetTagActionOk() (*string, bool)`

GetTagActionOk returns a tuple with the TagAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagAction

`func (o *ServicePortProfileListDetailDTO) SetTagAction(v string)`

SetTagAction sets TagAction field to given value.


### GetUserVlan

`func (o *ServicePortProfileListDetailDTO) GetUserVlan() int32`

GetUserVlan returns the UserVlan field if non-nil, zero value otherwise.

### GetUserVlanOk

`func (o *ServicePortProfileListDetailDTO) GetUserVlanOk() (*int32, bool)`

GetUserVlanOk returns a tuple with the UserVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVlan

`func (o *ServicePortProfileListDetailDTO) SetUserVlan(v int32)`

SetUserVlan sets UserVlan field to given value.


### GetUserVlanPriority

`func (o *ServicePortProfileListDetailDTO) GetUserVlanPriority() int32`

GetUserVlanPriority returns the UserVlanPriority field if non-nil, zero value otherwise.

### GetUserVlanPriorityOk

`func (o *ServicePortProfileListDetailDTO) GetUserVlanPriorityOk() (*int32, bool)`

GetUserVlanPriorityOk returns a tuple with the UserVlanPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVlanPriority

`func (o *ServicePortProfileListDetailDTO) SetUserVlanPriority(v int32)`

SetUserVlanPriority sets UserVlanPriority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


