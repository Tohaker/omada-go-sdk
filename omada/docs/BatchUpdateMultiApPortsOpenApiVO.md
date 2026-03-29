# BatchUpdateMultiApPortsOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMacList** | **[]string** | AP mac list | 
**BandwidthControlEnable** | Pointer to **bool** | Whether to enable bandwidthControl. | [optional] 
**Custom** | Pointer to **bool** | Whether to enter a tagged vlan customically. Enter the VLAN ID manually if it is true. and select LAN network profile if it is false. | [optional] 
**EgressRateLimit** | Pointer to [**ONUPortRateLimitVO**](ONUPortRateLimitVO.md) |  | [optional] 
**IngressRateLimit** | Pointer to [**ONUPortRateLimitVO**](ONUPortRateLimitVO.md) |  | [optional] 
**LanPortList** | **[]string** | The list of ap ports which are configured in batches. | 
**LocalVlanEnable** | Pointer to **bool** | Specifies whether to enable the port. | [optional] 
**LocalVlanId** | Pointer to **int32** | Local vlan ID should be within the range of 1-4094, which has a value only if Status is true. | [optional] 
**LocalVlanNetworkId** | Pointer to **string** | The ID of the LAN network profile selected for the port. | [optional] 
**Name** | Pointer to **string** | Port name | [optional] 
**PoeOutEnable** | Pointer to **bool** | Whether to enable poe out. | [optional] 
**Status** | Pointer to **bool** | Port status | [optional] 
**Tagged** | Pointer to **string** | Tagged vlan list, it must be present when custom is true. | [optional] 
**TaggedNetworkId** | Pointer to **[]string** | The Id list of the tagged vlan network profile, it must be present when custom is true. | [optional] 
**Untagged** | Pointer to **string** | Untagged vlan list, it must be present when custom is true. Two data cannot be duplicated. | [optional] 
**UntaggedNetworkId** | Pointer to **[]string** | The Id list of the untagged vlan network profile, it must be present when custom is true. | [optional] 

## Methods

### NewBatchUpdateMultiApPortsOpenApiVO

`func NewBatchUpdateMultiApPortsOpenApiVO(apMacList []string, lanPortList []string, ) *BatchUpdateMultiApPortsOpenApiVO`

NewBatchUpdateMultiApPortsOpenApiVO instantiates a new BatchUpdateMultiApPortsOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchUpdateMultiApPortsOpenApiVOWithDefaults

`func NewBatchUpdateMultiApPortsOpenApiVOWithDefaults() *BatchUpdateMultiApPortsOpenApiVO`

NewBatchUpdateMultiApPortsOpenApiVOWithDefaults instantiates a new BatchUpdateMultiApPortsOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMacList

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetApMacList() []string`

GetApMacList returns the ApMacList field if non-nil, zero value otherwise.

### GetApMacListOk

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetApMacListOk() (*[]string, bool)`

GetApMacListOk returns a tuple with the ApMacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMacList

`func (o *BatchUpdateMultiApPortsOpenApiVO) SetApMacList(v []string)`

SetApMacList sets ApMacList field to given value.


### GetBandwidthControlEnable

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetBandwidthControlEnable() bool`

GetBandwidthControlEnable returns the BandwidthControlEnable field if non-nil, zero value otherwise.

### GetBandwidthControlEnableOk

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetBandwidthControlEnableOk() (*bool, bool)`

GetBandwidthControlEnableOk returns a tuple with the BandwidthControlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthControlEnable

`func (o *BatchUpdateMultiApPortsOpenApiVO) SetBandwidthControlEnable(v bool)`

SetBandwidthControlEnable sets BandwidthControlEnable field to given value.

### HasBandwidthControlEnable

`func (o *BatchUpdateMultiApPortsOpenApiVO) HasBandwidthControlEnable() bool`

HasBandwidthControlEnable returns a boolean if a field has been set.

### GetCustom

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *BatchUpdateMultiApPortsOpenApiVO) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *BatchUpdateMultiApPortsOpenApiVO) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetEgressRateLimit

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetEgressRateLimit() ONUPortRateLimitVO`

GetEgressRateLimit returns the EgressRateLimit field if non-nil, zero value otherwise.

### GetEgressRateLimitOk

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetEgressRateLimitOk() (*ONUPortRateLimitVO, bool)`

GetEgressRateLimitOk returns a tuple with the EgressRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressRateLimit

`func (o *BatchUpdateMultiApPortsOpenApiVO) SetEgressRateLimit(v ONUPortRateLimitVO)`

SetEgressRateLimit sets EgressRateLimit field to given value.

### HasEgressRateLimit

`func (o *BatchUpdateMultiApPortsOpenApiVO) HasEgressRateLimit() bool`

HasEgressRateLimit returns a boolean if a field has been set.

### GetIngressRateLimit

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetIngressRateLimit() ONUPortRateLimitVO`

GetIngressRateLimit returns the IngressRateLimit field if non-nil, zero value otherwise.

### GetIngressRateLimitOk

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetIngressRateLimitOk() (*ONUPortRateLimitVO, bool)`

GetIngressRateLimitOk returns a tuple with the IngressRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressRateLimit

`func (o *BatchUpdateMultiApPortsOpenApiVO) SetIngressRateLimit(v ONUPortRateLimitVO)`

SetIngressRateLimit sets IngressRateLimit field to given value.

### HasIngressRateLimit

`func (o *BatchUpdateMultiApPortsOpenApiVO) HasIngressRateLimit() bool`

HasIngressRateLimit returns a boolean if a field has been set.

### GetLanPortList

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetLanPortList() []string`

GetLanPortList returns the LanPortList field if non-nil, zero value otherwise.

### GetLanPortListOk

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetLanPortListOk() (*[]string, bool)`

GetLanPortListOk returns a tuple with the LanPortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanPortList

`func (o *BatchUpdateMultiApPortsOpenApiVO) SetLanPortList(v []string)`

SetLanPortList sets LanPortList field to given value.


### GetLocalVlanEnable

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetLocalVlanEnable() bool`

GetLocalVlanEnable returns the LocalVlanEnable field if non-nil, zero value otherwise.

### GetLocalVlanEnableOk

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetLocalVlanEnableOk() (*bool, bool)`

GetLocalVlanEnableOk returns a tuple with the LocalVlanEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVlanEnable

`func (o *BatchUpdateMultiApPortsOpenApiVO) SetLocalVlanEnable(v bool)`

SetLocalVlanEnable sets LocalVlanEnable field to given value.

### HasLocalVlanEnable

`func (o *BatchUpdateMultiApPortsOpenApiVO) HasLocalVlanEnable() bool`

HasLocalVlanEnable returns a boolean if a field has been set.

### GetLocalVlanId

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetLocalVlanId() int32`

GetLocalVlanId returns the LocalVlanId field if non-nil, zero value otherwise.

### GetLocalVlanIdOk

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetLocalVlanIdOk() (*int32, bool)`

GetLocalVlanIdOk returns a tuple with the LocalVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVlanId

`func (o *BatchUpdateMultiApPortsOpenApiVO) SetLocalVlanId(v int32)`

SetLocalVlanId sets LocalVlanId field to given value.

### HasLocalVlanId

`func (o *BatchUpdateMultiApPortsOpenApiVO) HasLocalVlanId() bool`

HasLocalVlanId returns a boolean if a field has been set.

### GetLocalVlanNetworkId

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetLocalVlanNetworkId() string`

GetLocalVlanNetworkId returns the LocalVlanNetworkId field if non-nil, zero value otherwise.

### GetLocalVlanNetworkIdOk

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetLocalVlanNetworkIdOk() (*string, bool)`

GetLocalVlanNetworkIdOk returns a tuple with the LocalVlanNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVlanNetworkId

`func (o *BatchUpdateMultiApPortsOpenApiVO) SetLocalVlanNetworkId(v string)`

SetLocalVlanNetworkId sets LocalVlanNetworkId field to given value.

### HasLocalVlanNetworkId

`func (o *BatchUpdateMultiApPortsOpenApiVO) HasLocalVlanNetworkId() bool`

HasLocalVlanNetworkId returns a boolean if a field has been set.

### GetName

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BatchUpdateMultiApPortsOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BatchUpdateMultiApPortsOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPoeOutEnable

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetPoeOutEnable() bool`

GetPoeOutEnable returns the PoeOutEnable field if non-nil, zero value otherwise.

### GetPoeOutEnableOk

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetPoeOutEnableOk() (*bool, bool)`

GetPoeOutEnableOk returns a tuple with the PoeOutEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeOutEnable

`func (o *BatchUpdateMultiApPortsOpenApiVO) SetPoeOutEnable(v bool)`

SetPoeOutEnable sets PoeOutEnable field to given value.

### HasPoeOutEnable

`func (o *BatchUpdateMultiApPortsOpenApiVO) HasPoeOutEnable() bool`

HasPoeOutEnable returns a boolean if a field has been set.

### GetStatus

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchUpdateMultiApPortsOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BatchUpdateMultiApPortsOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTagged

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetTagged() string`

GetTagged returns the Tagged field if non-nil, zero value otherwise.

### GetTaggedOk

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetTaggedOk() (*string, bool)`

GetTaggedOk returns a tuple with the Tagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagged

`func (o *BatchUpdateMultiApPortsOpenApiVO) SetTagged(v string)`

SetTagged sets Tagged field to given value.

### HasTagged

`func (o *BatchUpdateMultiApPortsOpenApiVO) HasTagged() bool`

HasTagged returns a boolean if a field has been set.

### GetTaggedNetworkId

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetTaggedNetworkId() []string`

GetTaggedNetworkId returns the TaggedNetworkId field if non-nil, zero value otherwise.

### GetTaggedNetworkIdOk

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetTaggedNetworkIdOk() (*[]string, bool)`

GetTaggedNetworkIdOk returns a tuple with the TaggedNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedNetworkId

`func (o *BatchUpdateMultiApPortsOpenApiVO) SetTaggedNetworkId(v []string)`

SetTaggedNetworkId sets TaggedNetworkId field to given value.

### HasTaggedNetworkId

`func (o *BatchUpdateMultiApPortsOpenApiVO) HasTaggedNetworkId() bool`

HasTaggedNetworkId returns a boolean if a field has been set.

### GetUntagged

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetUntagged() string`

GetUntagged returns the Untagged field if non-nil, zero value otherwise.

### GetUntaggedOk

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetUntaggedOk() (*string, bool)`

GetUntaggedOk returns a tuple with the Untagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagged

`func (o *BatchUpdateMultiApPortsOpenApiVO) SetUntagged(v string)`

SetUntagged sets Untagged field to given value.

### HasUntagged

`func (o *BatchUpdateMultiApPortsOpenApiVO) HasUntagged() bool`

HasUntagged returns a boolean if a field has been set.

### GetUntaggedNetworkId

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetUntaggedNetworkId() []string`

GetUntaggedNetworkId returns the UntaggedNetworkId field if non-nil, zero value otherwise.

### GetUntaggedNetworkIdOk

`func (o *BatchUpdateMultiApPortsOpenApiVO) GetUntaggedNetworkIdOk() (*[]string, bool)`

GetUntaggedNetworkIdOk returns a tuple with the UntaggedNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntaggedNetworkId

`func (o *BatchUpdateMultiApPortsOpenApiVO) SetUntaggedNetworkId(v []string)`

SetUntaggedNetworkId sets UntaggedNetworkId field to given value.

### HasUntaggedNetworkId

`func (o *BatchUpdateMultiApPortsOpenApiVO) HasUntaggedNetworkId() bool`

HasUntaggedNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


