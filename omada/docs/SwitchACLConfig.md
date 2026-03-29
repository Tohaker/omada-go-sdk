# SwitchACLConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BiDirectional** | Pointer to **bool** | Whether to enable the bidirectional, this field required only when creating a new entry | [optional] 
**BindingBridgeVlan** | Pointer to **int32** | Only for bindingType VLAN and network of bridge VLAN | [optional] 
**BindingType** | **int32** | BindingType should be a value as follows: 0: all ports; 1: custom ports; 2: all switch vlan; 3: custom switch vlan | 
**CustomAclDevices** | Pointer to [**[]CustomAclOswOpenApiVO**](CustomAclOswOpenApiVO.md) | Only for bindingType is custom switch vlan, list of selected device | [optional] 
**CustomAclOsws** | Pointer to **[]string** | Only for bindingType is custom switch vlan, list of selected switch mac | [optional] 
**CustomAclPorts** | Pointer to [**[]SwitchACLPortEntity**](SwitchACLPortEntity.md) | Only for bindingType is custom ports, select the custom port or LAG of the device | [optional] 
**CustomAclStacks** | Pointer to **[]string** | Only for bindingType is custom switch vlan, list of selected stack id | [optional] 
**Description** | **string** | ACL rule description, description should contain 1 to 512 characters. | 
**DestinationIds** | Pointer to **[]string** | Source IDs, which depends on destinationType, for example: if destinationType is network, destinationIds should be LAN network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | [optional] 
**DestinationType** | **int32** | DestinationType should be a value as follows: 0: network; 1: IP Group; 2: IP-Port Group; 6: IPv6 Group; 7: IPv6-Port Group | 
**EtherType** | [**SwitchACLEtherTypeEntity**](SwitchACLEtherTypeEntity.md) |  | 
**NetworkId** | Pointer to **string** | Only for bindingType VLAN | [optional] 
**Policy** | **int32** | Policy should be a value as follows: 0: drop; 1: allow; | 
**Protocols** | **[]int32** | For the values of protocols, refer to section 5.5 of the Open API Access Guide. | 
**SourceIds** | **[]string** | Source IDs, which depends on sourceType, for example: if sourceType is network, sourceIds should be LAN network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | 
**SourceType** | **int32** | SourceType should be a value as follows: 0: network; 1: IP Group; 2: IP-Port Group; 4: SSID; 6: IPv6 Group; 7: IPv6-Port Group | 
**Status** | **bool** | Status should be a value as follows: 0: disable; 1: enable | 
**TimeRangeId** | Pointer to **string** | Time range profile ID | [optional] 

## Methods

### NewSwitchACLConfig

`func NewSwitchACLConfig(bindingType int32, description string, destinationType int32, etherType SwitchACLEtherTypeEntity, policy int32, protocols []int32, sourceIds []string, sourceType int32, status bool, ) *SwitchACLConfig`

NewSwitchACLConfig instantiates a new SwitchACLConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchACLConfigWithDefaults

`func NewSwitchACLConfigWithDefaults() *SwitchACLConfig`

NewSwitchACLConfigWithDefaults instantiates a new SwitchACLConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBiDirectional

`func (o *SwitchACLConfig) GetBiDirectional() bool`

GetBiDirectional returns the BiDirectional field if non-nil, zero value otherwise.

### GetBiDirectionalOk

`func (o *SwitchACLConfig) GetBiDirectionalOk() (*bool, bool)`

GetBiDirectionalOk returns a tuple with the BiDirectional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiDirectional

`func (o *SwitchACLConfig) SetBiDirectional(v bool)`

SetBiDirectional sets BiDirectional field to given value.

### HasBiDirectional

`func (o *SwitchACLConfig) HasBiDirectional() bool`

HasBiDirectional returns a boolean if a field has been set.

### GetBindingBridgeVlan

`func (o *SwitchACLConfig) GetBindingBridgeVlan() int32`

GetBindingBridgeVlan returns the BindingBridgeVlan field if non-nil, zero value otherwise.

### GetBindingBridgeVlanOk

`func (o *SwitchACLConfig) GetBindingBridgeVlanOk() (*int32, bool)`

GetBindingBridgeVlanOk returns a tuple with the BindingBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingBridgeVlan

`func (o *SwitchACLConfig) SetBindingBridgeVlan(v int32)`

SetBindingBridgeVlan sets BindingBridgeVlan field to given value.

### HasBindingBridgeVlan

`func (o *SwitchACLConfig) HasBindingBridgeVlan() bool`

HasBindingBridgeVlan returns a boolean if a field has been set.

### GetBindingType

`func (o *SwitchACLConfig) GetBindingType() int32`

GetBindingType returns the BindingType field if non-nil, zero value otherwise.

### GetBindingTypeOk

`func (o *SwitchACLConfig) GetBindingTypeOk() (*int32, bool)`

GetBindingTypeOk returns a tuple with the BindingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingType

`func (o *SwitchACLConfig) SetBindingType(v int32)`

SetBindingType sets BindingType field to given value.


### GetCustomAclDevices

`func (o *SwitchACLConfig) GetCustomAclDevices() []CustomAclOswOpenApiVO`

GetCustomAclDevices returns the CustomAclDevices field if non-nil, zero value otherwise.

### GetCustomAclDevicesOk

`func (o *SwitchACLConfig) GetCustomAclDevicesOk() (*[]CustomAclOswOpenApiVO, bool)`

GetCustomAclDevicesOk returns a tuple with the CustomAclDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAclDevices

`func (o *SwitchACLConfig) SetCustomAclDevices(v []CustomAclOswOpenApiVO)`

SetCustomAclDevices sets CustomAclDevices field to given value.

### HasCustomAclDevices

`func (o *SwitchACLConfig) HasCustomAclDevices() bool`

HasCustomAclDevices returns a boolean if a field has been set.

### GetCustomAclOsws

`func (o *SwitchACLConfig) GetCustomAclOsws() []string`

GetCustomAclOsws returns the CustomAclOsws field if non-nil, zero value otherwise.

### GetCustomAclOswsOk

`func (o *SwitchACLConfig) GetCustomAclOswsOk() (*[]string, bool)`

GetCustomAclOswsOk returns a tuple with the CustomAclOsws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAclOsws

`func (o *SwitchACLConfig) SetCustomAclOsws(v []string)`

SetCustomAclOsws sets CustomAclOsws field to given value.

### HasCustomAclOsws

`func (o *SwitchACLConfig) HasCustomAclOsws() bool`

HasCustomAclOsws returns a boolean if a field has been set.

### GetCustomAclPorts

`func (o *SwitchACLConfig) GetCustomAclPorts() []SwitchACLPortEntity`

GetCustomAclPorts returns the CustomAclPorts field if non-nil, zero value otherwise.

### GetCustomAclPortsOk

`func (o *SwitchACLConfig) GetCustomAclPortsOk() (*[]SwitchACLPortEntity, bool)`

GetCustomAclPortsOk returns a tuple with the CustomAclPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAclPorts

`func (o *SwitchACLConfig) SetCustomAclPorts(v []SwitchACLPortEntity)`

SetCustomAclPorts sets CustomAclPorts field to given value.

### HasCustomAclPorts

`func (o *SwitchACLConfig) HasCustomAclPorts() bool`

HasCustomAclPorts returns a boolean if a field has been set.

### GetCustomAclStacks

`func (o *SwitchACLConfig) GetCustomAclStacks() []string`

GetCustomAclStacks returns the CustomAclStacks field if non-nil, zero value otherwise.

### GetCustomAclStacksOk

`func (o *SwitchACLConfig) GetCustomAclStacksOk() (*[]string, bool)`

GetCustomAclStacksOk returns a tuple with the CustomAclStacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAclStacks

`func (o *SwitchACLConfig) SetCustomAclStacks(v []string)`

SetCustomAclStacks sets CustomAclStacks field to given value.

### HasCustomAclStacks

`func (o *SwitchACLConfig) HasCustomAclStacks() bool`

HasCustomAclStacks returns a boolean if a field has been set.

### GetDescription

`func (o *SwitchACLConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SwitchACLConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SwitchACLConfig) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDestinationIds

`func (o *SwitchACLConfig) GetDestinationIds() []string`

GetDestinationIds returns the DestinationIds field if non-nil, zero value otherwise.

### GetDestinationIdsOk

`func (o *SwitchACLConfig) GetDestinationIdsOk() (*[]string, bool)`

GetDestinationIdsOk returns a tuple with the DestinationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIds

`func (o *SwitchACLConfig) SetDestinationIds(v []string)`

SetDestinationIds sets DestinationIds field to given value.

### HasDestinationIds

`func (o *SwitchACLConfig) HasDestinationIds() bool`

HasDestinationIds returns a boolean if a field has been set.

### GetDestinationType

`func (o *SwitchACLConfig) GetDestinationType() int32`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *SwitchACLConfig) GetDestinationTypeOk() (*int32, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *SwitchACLConfig) SetDestinationType(v int32)`

SetDestinationType sets DestinationType field to given value.


### GetEtherType

`func (o *SwitchACLConfig) GetEtherType() SwitchACLEtherTypeEntity`

GetEtherType returns the EtherType field if non-nil, zero value otherwise.

### GetEtherTypeOk

`func (o *SwitchACLConfig) GetEtherTypeOk() (*SwitchACLEtherTypeEntity, bool)`

GetEtherTypeOk returns a tuple with the EtherType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherType

`func (o *SwitchACLConfig) SetEtherType(v SwitchACLEtherTypeEntity)`

SetEtherType sets EtherType field to given value.


### GetNetworkId

`func (o *SwitchACLConfig) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *SwitchACLConfig) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *SwitchACLConfig) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *SwitchACLConfig) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetPolicy

`func (o *SwitchACLConfig) GetPolicy() int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *SwitchACLConfig) GetPolicyOk() (*int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *SwitchACLConfig) SetPolicy(v int32)`

SetPolicy sets Policy field to given value.


### GetProtocols

`func (o *SwitchACLConfig) GetProtocols() []int32`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *SwitchACLConfig) GetProtocolsOk() (*[]int32, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *SwitchACLConfig) SetProtocols(v []int32)`

SetProtocols sets Protocols field to given value.


### GetSourceIds

`func (o *SwitchACLConfig) GetSourceIds() []string`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *SwitchACLConfig) GetSourceIdsOk() (*[]string, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *SwitchACLConfig) SetSourceIds(v []string)`

SetSourceIds sets SourceIds field to given value.


### GetSourceType

`func (o *SwitchACLConfig) GetSourceType() int32`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *SwitchACLConfig) GetSourceTypeOk() (*int32, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *SwitchACLConfig) SetSourceType(v int32)`

SetSourceType sets SourceType field to given value.


### GetStatus

`func (o *SwitchACLConfig) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SwitchACLConfig) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SwitchACLConfig) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetTimeRangeId

`func (o *SwitchACLConfig) GetTimeRangeId() string`

GetTimeRangeId returns the TimeRangeId field if non-nil, zero value otherwise.

### GetTimeRangeIdOk

`func (o *SwitchACLConfig) GetTimeRangeIdOk() (*string, bool)`

GetTimeRangeIdOk returns a tuple with the TimeRangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRangeId

`func (o *SwitchACLConfig) SetTimeRangeId(v string)`

SetTimeRangeId sets TimeRangeId field to given value.

### HasTimeRangeId

`func (o *SwitchACLConfig) HasTimeRangeId() bool`

HasTimeRangeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


