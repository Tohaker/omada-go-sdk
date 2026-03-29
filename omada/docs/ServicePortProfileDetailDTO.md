# ServicePortProfileDetailDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminStatus** | **string** | Whether to enable traffic flow  | 
**CreationMode** | **string** | The creation mode  | 
**Description** | Pointer to **string** | Service port description. Description should contain 1-32 characters, including uppercase and lowercase letters, numbers, and symbols(-@_:/.). | [optional] 
**EntriesNumber** | **int32** | Service port amount. The entriesNumber should be within the range of 0 to 127. | 
**EtherType** | **string** | Message type  | 
**GemId** | **int32** | Gem Port ID for automatic flow creation. The gemPort should be within the range of 1-1023. | 
**InUse** | **string** | Indicate whether the profile has been used  | 
**InboundTrafficProfileId** | **int32** | Inbound traffic profile ID. The inboundTrafficProfileId should be within the range of 1-1023, -1 indicating not configuring the item. | 
**InnerVlan** | **int32** | Inner VLAN ID for automatic flow creation. This parameter is associated with tagAction. When tagAction is TRANSLATE_AND_ADD/ADD_DOUBLE, this parameter is valid. The innerVlan should be within the range of 0-4094, -1 indicating not configuring the item.  | 
**InnerVlanPriority** | **int32** | Inner VLAN priority. This parameter is associated with tagAction. When tagAction is TRANSLATE_AND_ADD/ADD_DOUBLE, this parameter is valid. The innerVlanPriority should be within the range of -1 to 7, -1 indicating not configuring the item. | 
**OutboundTrafficProfileId** | **int32** | Outbound traffic profile ID. The outboundTrafficProfileId should be within the range of 1-1023, -1 indicating not configuring the item. | 
**ServicePortId** | **string** | Service port profile ID. The servicePortProfileId should be within the range of 1 to 127. | 
**ServicePortProfileId** | **string** | Service template ID. The servicePortProfileId should be within the range of 1 to 127. It can be obtained from \&quot;Get service port profile list\&quot; | 
**StatisticPerformance** | **string** | Whether to enable traffic flow statistics  | 
**Svlan** | **int32** | Bound network VLAN ID after automatic flow creation. The svlan should be within the range of 1-4094. | 
**TagAction** | **string** | VLAN TAG operation after automatic flow creation   | 
**UserVlan** | **int32** | User VLAN ID for automatic flow creation. The userVlan should be within the range of 0-4094. | 
**UserVlanPriority** | **int32** | User VLAN priority. The userVlanPriority should be within the range of -1 to 7, -1 indicating not configuring the item. | 

## Methods

### NewServicePortProfileDetailDTO

`func NewServicePortProfileDetailDTO(adminStatus string, creationMode string, entriesNumber int32, etherType string, gemId int32, inUse string, inboundTrafficProfileId int32, innerVlan int32, innerVlanPriority int32, outboundTrafficProfileId int32, servicePortId string, servicePortProfileId string, statisticPerformance string, svlan int32, tagAction string, userVlan int32, userVlanPriority int32, ) *ServicePortProfileDetailDTO`

NewServicePortProfileDetailDTO instantiates a new ServicePortProfileDetailDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePortProfileDetailDTOWithDefaults

`func NewServicePortProfileDetailDTOWithDefaults() *ServicePortProfileDetailDTO`

NewServicePortProfileDetailDTOWithDefaults instantiates a new ServicePortProfileDetailDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminStatus

`func (o *ServicePortProfileDetailDTO) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *ServicePortProfileDetailDTO) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *ServicePortProfileDetailDTO) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.


### GetCreationMode

`func (o *ServicePortProfileDetailDTO) GetCreationMode() string`

GetCreationMode returns the CreationMode field if non-nil, zero value otherwise.

### GetCreationModeOk

`func (o *ServicePortProfileDetailDTO) GetCreationModeOk() (*string, bool)`

GetCreationModeOk returns a tuple with the CreationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationMode

`func (o *ServicePortProfileDetailDTO) SetCreationMode(v string)`

SetCreationMode sets CreationMode field to given value.


### GetDescription

`func (o *ServicePortProfileDetailDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicePortProfileDetailDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicePortProfileDetailDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServicePortProfileDetailDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntriesNumber

`func (o *ServicePortProfileDetailDTO) GetEntriesNumber() int32`

GetEntriesNumber returns the EntriesNumber field if non-nil, zero value otherwise.

### GetEntriesNumberOk

`func (o *ServicePortProfileDetailDTO) GetEntriesNumberOk() (*int32, bool)`

GetEntriesNumberOk returns a tuple with the EntriesNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntriesNumber

`func (o *ServicePortProfileDetailDTO) SetEntriesNumber(v int32)`

SetEntriesNumber sets EntriesNumber field to given value.


### GetEtherType

`func (o *ServicePortProfileDetailDTO) GetEtherType() string`

GetEtherType returns the EtherType field if non-nil, zero value otherwise.

### GetEtherTypeOk

`func (o *ServicePortProfileDetailDTO) GetEtherTypeOk() (*string, bool)`

GetEtherTypeOk returns a tuple with the EtherType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherType

`func (o *ServicePortProfileDetailDTO) SetEtherType(v string)`

SetEtherType sets EtherType field to given value.


### GetGemId

`func (o *ServicePortProfileDetailDTO) GetGemId() int32`

GetGemId returns the GemId field if non-nil, zero value otherwise.

### GetGemIdOk

`func (o *ServicePortProfileDetailDTO) GetGemIdOk() (*int32, bool)`

GetGemIdOk returns a tuple with the GemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemId

`func (o *ServicePortProfileDetailDTO) SetGemId(v int32)`

SetGemId sets GemId field to given value.


### GetInUse

`func (o *ServicePortProfileDetailDTO) GetInUse() string`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *ServicePortProfileDetailDTO) GetInUseOk() (*string, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *ServicePortProfileDetailDTO) SetInUse(v string)`

SetInUse sets InUse field to given value.


### GetInboundTrafficProfileId

`func (o *ServicePortProfileDetailDTO) GetInboundTrafficProfileId() int32`

GetInboundTrafficProfileId returns the InboundTrafficProfileId field if non-nil, zero value otherwise.

### GetInboundTrafficProfileIdOk

`func (o *ServicePortProfileDetailDTO) GetInboundTrafficProfileIdOk() (*int32, bool)`

GetInboundTrafficProfileIdOk returns a tuple with the InboundTrafficProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundTrafficProfileId

`func (o *ServicePortProfileDetailDTO) SetInboundTrafficProfileId(v int32)`

SetInboundTrafficProfileId sets InboundTrafficProfileId field to given value.


### GetInnerVlan

`func (o *ServicePortProfileDetailDTO) GetInnerVlan() int32`

GetInnerVlan returns the InnerVlan field if non-nil, zero value otherwise.

### GetInnerVlanOk

`func (o *ServicePortProfileDetailDTO) GetInnerVlanOk() (*int32, bool)`

GetInnerVlanOk returns a tuple with the InnerVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerVlan

`func (o *ServicePortProfileDetailDTO) SetInnerVlan(v int32)`

SetInnerVlan sets InnerVlan field to given value.


### GetInnerVlanPriority

`func (o *ServicePortProfileDetailDTO) GetInnerVlanPriority() int32`

GetInnerVlanPriority returns the InnerVlanPriority field if non-nil, zero value otherwise.

### GetInnerVlanPriorityOk

`func (o *ServicePortProfileDetailDTO) GetInnerVlanPriorityOk() (*int32, bool)`

GetInnerVlanPriorityOk returns a tuple with the InnerVlanPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerVlanPriority

`func (o *ServicePortProfileDetailDTO) SetInnerVlanPriority(v int32)`

SetInnerVlanPriority sets InnerVlanPriority field to given value.


### GetOutboundTrafficProfileId

`func (o *ServicePortProfileDetailDTO) GetOutboundTrafficProfileId() int32`

GetOutboundTrafficProfileId returns the OutboundTrafficProfileId field if non-nil, zero value otherwise.

### GetOutboundTrafficProfileIdOk

`func (o *ServicePortProfileDetailDTO) GetOutboundTrafficProfileIdOk() (*int32, bool)`

GetOutboundTrafficProfileIdOk returns a tuple with the OutboundTrafficProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundTrafficProfileId

`func (o *ServicePortProfileDetailDTO) SetOutboundTrafficProfileId(v int32)`

SetOutboundTrafficProfileId sets OutboundTrafficProfileId field to given value.


### GetServicePortId

`func (o *ServicePortProfileDetailDTO) GetServicePortId() string`

GetServicePortId returns the ServicePortId field if non-nil, zero value otherwise.

### GetServicePortIdOk

`func (o *ServicePortProfileDetailDTO) GetServicePortIdOk() (*string, bool)`

GetServicePortIdOk returns a tuple with the ServicePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePortId

`func (o *ServicePortProfileDetailDTO) SetServicePortId(v string)`

SetServicePortId sets ServicePortId field to given value.


### GetServicePortProfileId

`func (o *ServicePortProfileDetailDTO) GetServicePortProfileId() string`

GetServicePortProfileId returns the ServicePortProfileId field if non-nil, zero value otherwise.

### GetServicePortProfileIdOk

`func (o *ServicePortProfileDetailDTO) GetServicePortProfileIdOk() (*string, bool)`

GetServicePortProfileIdOk returns a tuple with the ServicePortProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePortProfileId

`func (o *ServicePortProfileDetailDTO) SetServicePortProfileId(v string)`

SetServicePortProfileId sets ServicePortProfileId field to given value.


### GetStatisticPerformance

`func (o *ServicePortProfileDetailDTO) GetStatisticPerformance() string`

GetStatisticPerformance returns the StatisticPerformance field if non-nil, zero value otherwise.

### GetStatisticPerformanceOk

`func (o *ServicePortProfileDetailDTO) GetStatisticPerformanceOk() (*string, bool)`

GetStatisticPerformanceOk returns a tuple with the StatisticPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticPerformance

`func (o *ServicePortProfileDetailDTO) SetStatisticPerformance(v string)`

SetStatisticPerformance sets StatisticPerformance field to given value.


### GetSvlan

`func (o *ServicePortProfileDetailDTO) GetSvlan() int32`

GetSvlan returns the Svlan field if non-nil, zero value otherwise.

### GetSvlanOk

`func (o *ServicePortProfileDetailDTO) GetSvlanOk() (*int32, bool)`

GetSvlanOk returns a tuple with the Svlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvlan

`func (o *ServicePortProfileDetailDTO) SetSvlan(v int32)`

SetSvlan sets Svlan field to given value.


### GetTagAction

`func (o *ServicePortProfileDetailDTO) GetTagAction() string`

GetTagAction returns the TagAction field if non-nil, zero value otherwise.

### GetTagActionOk

`func (o *ServicePortProfileDetailDTO) GetTagActionOk() (*string, bool)`

GetTagActionOk returns a tuple with the TagAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagAction

`func (o *ServicePortProfileDetailDTO) SetTagAction(v string)`

SetTagAction sets TagAction field to given value.


### GetUserVlan

`func (o *ServicePortProfileDetailDTO) GetUserVlan() int32`

GetUserVlan returns the UserVlan field if non-nil, zero value otherwise.

### GetUserVlanOk

`func (o *ServicePortProfileDetailDTO) GetUserVlanOk() (*int32, bool)`

GetUserVlanOk returns a tuple with the UserVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVlan

`func (o *ServicePortProfileDetailDTO) SetUserVlan(v int32)`

SetUserVlan sets UserVlan field to given value.


### GetUserVlanPriority

`func (o *ServicePortProfileDetailDTO) GetUserVlanPriority() int32`

GetUserVlanPriority returns the UserVlanPriority field if non-nil, zero value otherwise.

### GetUserVlanPriorityOk

`func (o *ServicePortProfileDetailDTO) GetUserVlanPriorityOk() (*int32, bool)`

GetUserVlanPriorityOk returns a tuple with the UserVlanPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVlanPriority

`func (o *ServicePortProfileDetailDTO) SetUserVlanPriority(v int32)`

SetUserVlanPriority sets UserVlanPriority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


