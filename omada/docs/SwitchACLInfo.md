# SwitchACLInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindingBridgeVlan** | Pointer to **int32** | Only for bindingType VLAN and network of bridge VLAN | [optional] 
**BindingType** | **int32** | BindingType should be a value as follows: 0: all ports; 1: custom ports; 2: all switch vlan; 3: custom switch vlan | 
**CustomAclDevices** | Pointer to [**[]CustomAclOswOpenApiVO**](CustomAclOswOpenApiVO.md) | Only for bindingType is custom switch vlan, list of selected device | [optional] 
**CustomAclOsws** | Pointer to **[]string** | Only for bindingType is custom switch vlan, list of selected switch mac | [optional] 
**CustomAclPorts** | Pointer to [**[]SwitchACLPortEntity**](SwitchACLPortEntity.md) | Only for bindingType is custom ports, select the custom port or LAG of the device | [optional] 
**CustomAclStacks** | Pointer to **[]string** | Only for bindingType is custom switch vlan, list of selected stack id | [optional] 
**Description** | **string** | ACL rule description, description should contain 1 to 512 characters. | 
**DestinationIds** | **[]string** | Destination IDs, which depends on destinationType, for example: if destinationType is network, destinationIds should be LAN network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | 
**DestinationType** | **int32** | DestinationType should be a value as follows: 0: network; 1: IP Group; 2: IP-Port Group; 6: IPv6 Group; 7: IPv6-Port Group | 
**EtherType** | [**SwitchACLEtherTypeEntity**](SwitchACLEtherTypeEntity.md) |  | 
**Id** | **string** | ACL ID | 
**Index** | **int32** | Index | 
**NetworkId** | Pointer to **string** | Only for bindingType VLAN | [optional] 
**Policy** | **int32** | Policy should be a value as follows: 0: drop; 1: allow; | 
**Protocols** | **[]int32** | For the values of protocols, refer to section 5.5 of the Open API Access Guide. | 
**SourceIds** | **[]string** | Source IDs, which depends on sourceType, for example: if sourceType is network, sourceIds should be LAN network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | 
**SourceType** | **int32** | SourceType should be a value as follows: 0: network; 1: IP Group; 2: IP-Port Group; 4: SSID; 6: IPv6 Group; 7: IPv6-Port Group | 
**Status** | **bool** | Status should be a value as follows: 0: disable; 1: enable | 
**TimeRangeId** | Pointer to **string** | Time range profile ID | [optional] 

## Methods

### NewSwitchACLInfo

`func NewSwitchACLInfo(bindingType int32, description string, destinationIds []string, destinationType int32, etherType SwitchACLEtherTypeEntity, id string, index int32, policy int32, protocols []int32, sourceIds []string, sourceType int32, status bool, ) *SwitchACLInfo`

NewSwitchACLInfo instantiates a new SwitchACLInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchACLInfoWithDefaults

`func NewSwitchACLInfoWithDefaults() *SwitchACLInfo`

NewSwitchACLInfoWithDefaults instantiates a new SwitchACLInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindingBridgeVlan

`func (o *SwitchACLInfo) GetBindingBridgeVlan() int32`

GetBindingBridgeVlan returns the BindingBridgeVlan field if non-nil, zero value otherwise.

### GetBindingBridgeVlanOk

`func (o *SwitchACLInfo) GetBindingBridgeVlanOk() (*int32, bool)`

GetBindingBridgeVlanOk returns a tuple with the BindingBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingBridgeVlan

`func (o *SwitchACLInfo) SetBindingBridgeVlan(v int32)`

SetBindingBridgeVlan sets BindingBridgeVlan field to given value.

### HasBindingBridgeVlan

`func (o *SwitchACLInfo) HasBindingBridgeVlan() bool`

HasBindingBridgeVlan returns a boolean if a field has been set.

### GetBindingType

`func (o *SwitchACLInfo) GetBindingType() int32`

GetBindingType returns the BindingType field if non-nil, zero value otherwise.

### GetBindingTypeOk

`func (o *SwitchACLInfo) GetBindingTypeOk() (*int32, bool)`

GetBindingTypeOk returns a tuple with the BindingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingType

`func (o *SwitchACLInfo) SetBindingType(v int32)`

SetBindingType sets BindingType field to given value.


### GetCustomAclDevices

`func (o *SwitchACLInfo) GetCustomAclDevices() []CustomAclOswOpenApiVO`

GetCustomAclDevices returns the CustomAclDevices field if non-nil, zero value otherwise.

### GetCustomAclDevicesOk

`func (o *SwitchACLInfo) GetCustomAclDevicesOk() (*[]CustomAclOswOpenApiVO, bool)`

GetCustomAclDevicesOk returns a tuple with the CustomAclDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAclDevices

`func (o *SwitchACLInfo) SetCustomAclDevices(v []CustomAclOswOpenApiVO)`

SetCustomAclDevices sets CustomAclDevices field to given value.

### HasCustomAclDevices

`func (o *SwitchACLInfo) HasCustomAclDevices() bool`

HasCustomAclDevices returns a boolean if a field has been set.

### GetCustomAclOsws

`func (o *SwitchACLInfo) GetCustomAclOsws() []string`

GetCustomAclOsws returns the CustomAclOsws field if non-nil, zero value otherwise.

### GetCustomAclOswsOk

`func (o *SwitchACLInfo) GetCustomAclOswsOk() (*[]string, bool)`

GetCustomAclOswsOk returns a tuple with the CustomAclOsws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAclOsws

`func (o *SwitchACLInfo) SetCustomAclOsws(v []string)`

SetCustomAclOsws sets CustomAclOsws field to given value.

### HasCustomAclOsws

`func (o *SwitchACLInfo) HasCustomAclOsws() bool`

HasCustomAclOsws returns a boolean if a field has been set.

### GetCustomAclPorts

`func (o *SwitchACLInfo) GetCustomAclPorts() []SwitchACLPortEntity`

GetCustomAclPorts returns the CustomAclPorts field if non-nil, zero value otherwise.

### GetCustomAclPortsOk

`func (o *SwitchACLInfo) GetCustomAclPortsOk() (*[]SwitchACLPortEntity, bool)`

GetCustomAclPortsOk returns a tuple with the CustomAclPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAclPorts

`func (o *SwitchACLInfo) SetCustomAclPorts(v []SwitchACLPortEntity)`

SetCustomAclPorts sets CustomAclPorts field to given value.

### HasCustomAclPorts

`func (o *SwitchACLInfo) HasCustomAclPorts() bool`

HasCustomAclPorts returns a boolean if a field has been set.

### GetCustomAclStacks

`func (o *SwitchACLInfo) GetCustomAclStacks() []string`

GetCustomAclStacks returns the CustomAclStacks field if non-nil, zero value otherwise.

### GetCustomAclStacksOk

`func (o *SwitchACLInfo) GetCustomAclStacksOk() (*[]string, bool)`

GetCustomAclStacksOk returns a tuple with the CustomAclStacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAclStacks

`func (o *SwitchACLInfo) SetCustomAclStacks(v []string)`

SetCustomAclStacks sets CustomAclStacks field to given value.

### HasCustomAclStacks

`func (o *SwitchACLInfo) HasCustomAclStacks() bool`

HasCustomAclStacks returns a boolean if a field has been set.

### GetDescription

`func (o *SwitchACLInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SwitchACLInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SwitchACLInfo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDestinationIds

`func (o *SwitchACLInfo) GetDestinationIds() []string`

GetDestinationIds returns the DestinationIds field if non-nil, zero value otherwise.

### GetDestinationIdsOk

`func (o *SwitchACLInfo) GetDestinationIdsOk() (*[]string, bool)`

GetDestinationIdsOk returns a tuple with the DestinationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIds

`func (o *SwitchACLInfo) SetDestinationIds(v []string)`

SetDestinationIds sets DestinationIds field to given value.


### GetDestinationType

`func (o *SwitchACLInfo) GetDestinationType() int32`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *SwitchACLInfo) GetDestinationTypeOk() (*int32, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *SwitchACLInfo) SetDestinationType(v int32)`

SetDestinationType sets DestinationType field to given value.


### GetEtherType

`func (o *SwitchACLInfo) GetEtherType() SwitchACLEtherTypeEntity`

GetEtherType returns the EtherType field if non-nil, zero value otherwise.

### GetEtherTypeOk

`func (o *SwitchACLInfo) GetEtherTypeOk() (*SwitchACLEtherTypeEntity, bool)`

GetEtherTypeOk returns a tuple with the EtherType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherType

`func (o *SwitchACLInfo) SetEtherType(v SwitchACLEtherTypeEntity)`

SetEtherType sets EtherType field to given value.


### GetId

`func (o *SwitchACLInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SwitchACLInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SwitchACLInfo) SetId(v string)`

SetId sets Id field to given value.


### GetIndex

`func (o *SwitchACLInfo) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *SwitchACLInfo) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *SwitchACLInfo) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetNetworkId

`func (o *SwitchACLInfo) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *SwitchACLInfo) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *SwitchACLInfo) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *SwitchACLInfo) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetPolicy

`func (o *SwitchACLInfo) GetPolicy() int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *SwitchACLInfo) GetPolicyOk() (*int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *SwitchACLInfo) SetPolicy(v int32)`

SetPolicy sets Policy field to given value.


### GetProtocols

`func (o *SwitchACLInfo) GetProtocols() []int32`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *SwitchACLInfo) GetProtocolsOk() (*[]int32, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *SwitchACLInfo) SetProtocols(v []int32)`

SetProtocols sets Protocols field to given value.


### GetSourceIds

`func (o *SwitchACLInfo) GetSourceIds() []string`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *SwitchACLInfo) GetSourceIdsOk() (*[]string, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *SwitchACLInfo) SetSourceIds(v []string)`

SetSourceIds sets SourceIds field to given value.


### GetSourceType

`func (o *SwitchACLInfo) GetSourceType() int32`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *SwitchACLInfo) GetSourceTypeOk() (*int32, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *SwitchACLInfo) SetSourceType(v int32)`

SetSourceType sets SourceType field to given value.


### GetStatus

`func (o *SwitchACLInfo) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SwitchACLInfo) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SwitchACLInfo) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetTimeRangeId

`func (o *SwitchACLInfo) GetTimeRangeId() string`

GetTimeRangeId returns the TimeRangeId field if non-nil, zero value otherwise.

### GetTimeRangeIdOk

`func (o *SwitchACLInfo) GetTimeRangeIdOk() (*string, bool)`

GetTimeRangeIdOk returns a tuple with the TimeRangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRangeId

`func (o *SwitchACLInfo) SetTimeRangeId(v string)`

SetTimeRangeId sets TimeRangeId field to given value.

### HasTimeRangeId

`func (o *SwitchACLInfo) HasTimeRangeId() bool`

HasTimeRangeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


