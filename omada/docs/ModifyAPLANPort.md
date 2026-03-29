# ModifyAPLANPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandwidthControlEnable** | Pointer to **bool** | bandwidth Control status | [optional] 
**Custom** | Pointer to **bool** | custom | [optional] 
**EgressRateLimit** | Pointer to [**ONUPortRateLimitVO**](ONUPortRateLimitVO.md) |  | [optional] 
**IngressRateLimit** | Pointer to [**ONUPortRateLimitVO**](ONUPortRateLimitVO.md) |  | [optional] 
**LanPort** | Pointer to **string** | lanPort | [optional] 
**LocalVlanEnable** | Pointer to **bool** | vlan status, true: enable; false: disable | [optional] 
**LocalVlanId** | Pointer to **int32** | port vlan id | [optional] 
**LocalVlanNetworkId** | Pointer to **string** | local vlan network Id | [optional] 
**Name** | Pointer to **string** | port name | [optional] 
**PoeOutEnable** | Pointer to **bool** | poe out status, true: enable; false: disable | [optional] 
**Status** | Pointer to **bool** | port status, true: enable; false: disable | [optional] 
**Tagged** | Pointer to **string** | tagged | [optional] 
**TaggedNetworkId** | Pointer to **[]string** | vlan tagged NetworkId | [optional] 
**Untagged** | Pointer to **string** | untagged | [optional] 
**UntaggedNetworkId** | Pointer to **[]string** | vlan untagged NetworkId | [optional] 

## Methods

### NewModifyAPLANPort

`func NewModifyAPLANPort() *ModifyAPLANPort`

NewModifyAPLANPort instantiates a new ModifyAPLANPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyAPLANPortWithDefaults

`func NewModifyAPLANPortWithDefaults() *ModifyAPLANPort`

NewModifyAPLANPortWithDefaults instantiates a new ModifyAPLANPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidthControlEnable

`func (o *ModifyAPLANPort) GetBandwidthControlEnable() bool`

GetBandwidthControlEnable returns the BandwidthControlEnable field if non-nil, zero value otherwise.

### GetBandwidthControlEnableOk

`func (o *ModifyAPLANPort) GetBandwidthControlEnableOk() (*bool, bool)`

GetBandwidthControlEnableOk returns a tuple with the BandwidthControlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthControlEnable

`func (o *ModifyAPLANPort) SetBandwidthControlEnable(v bool)`

SetBandwidthControlEnable sets BandwidthControlEnable field to given value.

### HasBandwidthControlEnable

`func (o *ModifyAPLANPort) HasBandwidthControlEnable() bool`

HasBandwidthControlEnable returns a boolean if a field has been set.

### GetCustom

`func (o *ModifyAPLANPort) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *ModifyAPLANPort) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *ModifyAPLANPort) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *ModifyAPLANPort) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetEgressRateLimit

`func (o *ModifyAPLANPort) GetEgressRateLimit() ONUPortRateLimitVO`

GetEgressRateLimit returns the EgressRateLimit field if non-nil, zero value otherwise.

### GetEgressRateLimitOk

`func (o *ModifyAPLANPort) GetEgressRateLimitOk() (*ONUPortRateLimitVO, bool)`

GetEgressRateLimitOk returns a tuple with the EgressRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressRateLimit

`func (o *ModifyAPLANPort) SetEgressRateLimit(v ONUPortRateLimitVO)`

SetEgressRateLimit sets EgressRateLimit field to given value.

### HasEgressRateLimit

`func (o *ModifyAPLANPort) HasEgressRateLimit() bool`

HasEgressRateLimit returns a boolean if a field has been set.

### GetIngressRateLimit

`func (o *ModifyAPLANPort) GetIngressRateLimit() ONUPortRateLimitVO`

GetIngressRateLimit returns the IngressRateLimit field if non-nil, zero value otherwise.

### GetIngressRateLimitOk

`func (o *ModifyAPLANPort) GetIngressRateLimitOk() (*ONUPortRateLimitVO, bool)`

GetIngressRateLimitOk returns a tuple with the IngressRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressRateLimit

`func (o *ModifyAPLANPort) SetIngressRateLimit(v ONUPortRateLimitVO)`

SetIngressRateLimit sets IngressRateLimit field to given value.

### HasIngressRateLimit

`func (o *ModifyAPLANPort) HasIngressRateLimit() bool`

HasIngressRateLimit returns a boolean if a field has been set.

### GetLanPort

`func (o *ModifyAPLANPort) GetLanPort() string`

GetLanPort returns the LanPort field if non-nil, zero value otherwise.

### GetLanPortOk

`func (o *ModifyAPLANPort) GetLanPortOk() (*string, bool)`

GetLanPortOk returns a tuple with the LanPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanPort

`func (o *ModifyAPLANPort) SetLanPort(v string)`

SetLanPort sets LanPort field to given value.

### HasLanPort

`func (o *ModifyAPLANPort) HasLanPort() bool`

HasLanPort returns a boolean if a field has been set.

### GetLocalVlanEnable

`func (o *ModifyAPLANPort) GetLocalVlanEnable() bool`

GetLocalVlanEnable returns the LocalVlanEnable field if non-nil, zero value otherwise.

### GetLocalVlanEnableOk

`func (o *ModifyAPLANPort) GetLocalVlanEnableOk() (*bool, bool)`

GetLocalVlanEnableOk returns a tuple with the LocalVlanEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVlanEnable

`func (o *ModifyAPLANPort) SetLocalVlanEnable(v bool)`

SetLocalVlanEnable sets LocalVlanEnable field to given value.

### HasLocalVlanEnable

`func (o *ModifyAPLANPort) HasLocalVlanEnable() bool`

HasLocalVlanEnable returns a boolean if a field has been set.

### GetLocalVlanId

`func (o *ModifyAPLANPort) GetLocalVlanId() int32`

GetLocalVlanId returns the LocalVlanId field if non-nil, zero value otherwise.

### GetLocalVlanIdOk

`func (o *ModifyAPLANPort) GetLocalVlanIdOk() (*int32, bool)`

GetLocalVlanIdOk returns a tuple with the LocalVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVlanId

`func (o *ModifyAPLANPort) SetLocalVlanId(v int32)`

SetLocalVlanId sets LocalVlanId field to given value.

### HasLocalVlanId

`func (o *ModifyAPLANPort) HasLocalVlanId() bool`

HasLocalVlanId returns a boolean if a field has been set.

### GetLocalVlanNetworkId

`func (o *ModifyAPLANPort) GetLocalVlanNetworkId() string`

GetLocalVlanNetworkId returns the LocalVlanNetworkId field if non-nil, zero value otherwise.

### GetLocalVlanNetworkIdOk

`func (o *ModifyAPLANPort) GetLocalVlanNetworkIdOk() (*string, bool)`

GetLocalVlanNetworkIdOk returns a tuple with the LocalVlanNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVlanNetworkId

`func (o *ModifyAPLANPort) SetLocalVlanNetworkId(v string)`

SetLocalVlanNetworkId sets LocalVlanNetworkId field to given value.

### HasLocalVlanNetworkId

`func (o *ModifyAPLANPort) HasLocalVlanNetworkId() bool`

HasLocalVlanNetworkId returns a boolean if a field has been set.

### GetName

`func (o *ModifyAPLANPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyAPLANPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyAPLANPort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModifyAPLANPort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPoeOutEnable

`func (o *ModifyAPLANPort) GetPoeOutEnable() bool`

GetPoeOutEnable returns the PoeOutEnable field if non-nil, zero value otherwise.

### GetPoeOutEnableOk

`func (o *ModifyAPLANPort) GetPoeOutEnableOk() (*bool, bool)`

GetPoeOutEnableOk returns a tuple with the PoeOutEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeOutEnable

`func (o *ModifyAPLANPort) SetPoeOutEnable(v bool)`

SetPoeOutEnable sets PoeOutEnable field to given value.

### HasPoeOutEnable

`func (o *ModifyAPLANPort) HasPoeOutEnable() bool`

HasPoeOutEnable returns a boolean if a field has been set.

### GetStatus

`func (o *ModifyAPLANPort) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModifyAPLANPort) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModifyAPLANPort) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModifyAPLANPort) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTagged

`func (o *ModifyAPLANPort) GetTagged() string`

GetTagged returns the Tagged field if non-nil, zero value otherwise.

### GetTaggedOk

`func (o *ModifyAPLANPort) GetTaggedOk() (*string, bool)`

GetTaggedOk returns a tuple with the Tagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagged

`func (o *ModifyAPLANPort) SetTagged(v string)`

SetTagged sets Tagged field to given value.

### HasTagged

`func (o *ModifyAPLANPort) HasTagged() bool`

HasTagged returns a boolean if a field has been set.

### GetTaggedNetworkId

`func (o *ModifyAPLANPort) GetTaggedNetworkId() []string`

GetTaggedNetworkId returns the TaggedNetworkId field if non-nil, zero value otherwise.

### GetTaggedNetworkIdOk

`func (o *ModifyAPLANPort) GetTaggedNetworkIdOk() (*[]string, bool)`

GetTaggedNetworkIdOk returns a tuple with the TaggedNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedNetworkId

`func (o *ModifyAPLANPort) SetTaggedNetworkId(v []string)`

SetTaggedNetworkId sets TaggedNetworkId field to given value.

### HasTaggedNetworkId

`func (o *ModifyAPLANPort) HasTaggedNetworkId() bool`

HasTaggedNetworkId returns a boolean if a field has been set.

### GetUntagged

`func (o *ModifyAPLANPort) GetUntagged() string`

GetUntagged returns the Untagged field if non-nil, zero value otherwise.

### GetUntaggedOk

`func (o *ModifyAPLANPort) GetUntaggedOk() (*string, bool)`

GetUntaggedOk returns a tuple with the Untagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagged

`func (o *ModifyAPLANPort) SetUntagged(v string)`

SetUntagged sets Untagged field to given value.

### HasUntagged

`func (o *ModifyAPLANPort) HasUntagged() bool`

HasUntagged returns a boolean if a field has been set.

### GetUntaggedNetworkId

`func (o *ModifyAPLANPort) GetUntaggedNetworkId() []string`

GetUntaggedNetworkId returns the UntaggedNetworkId field if non-nil, zero value otherwise.

### GetUntaggedNetworkIdOk

`func (o *ModifyAPLANPort) GetUntaggedNetworkIdOk() (*[]string, bool)`

GetUntaggedNetworkIdOk returns a tuple with the UntaggedNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntaggedNetworkId

`func (o *ModifyAPLANPort) SetUntaggedNetworkId(v []string)`

SetUntaggedNetworkId sets UntaggedNetworkId field to given value.

### HasUntaggedNetworkId

`func (o *ModifyAPLANPort) HasUntaggedNetworkId() bool`

HasUntaggedNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


