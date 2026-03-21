# PolicyRoutingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupInterface** | **bool** | Use another WAN port if the current one is down | 
**DestinationIds** | **[]string** | Destination IDs, which depends on destinationType, for example: if destinationType is network, destinationIds should be LAN network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | 
**DestinationType** | **int32** | DestinationType should be a value as follows: 0: Network; 1: IP Group; 2: IP port Group | 
**ExistDomainGroupDest** | Pointer to **bool** | Whether the destination type exists domain group | [optional] 
**ExistIpPortGroup** | Pointer to **bool** | Whether the destination type exists ip port group | [optional] 
**ExistLocationGroupDest** | Pointer to **bool** | Whether the destination type exists location group | [optional] 
**ExistMulti** | Pointer to **bool** | Whether the interface option exists multi | [optional] 
**ExistVirtualWan** | Pointer to **bool** | Whether the interface option exists virtual WAN | [optional] 
**ExistVpnClient** | Pointer to **bool** | Whether the interface option exists VPN client | [optional] 
**Id** | Pointer to **string** | ID | [optional] 
**Index** | Pointer to **int32** | Index | [optional] 
**InterfaceId** | Pointer to **string** | Interface ID | [optional] 
**InterfaceType** | Pointer to **int32** | InterfaceType should be a value as follows: 0: WAN; 2: L2TP; 3: PPTP; 4: multi-select WAN or VPN or virtual WAN. | [optional] 
**Name** | **string** | Name, name should contain 1 to 64 characters. | 
**Protocols** | **[]int32** | For the values of protocols, refer to section 5.5.2 of the Open API Access Guide. | 
**SourceIds** | **[]string** | Source IDs, which depends on sourceType, for example: if sourceType is network, sourceIds should be LAN network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | 
**SourceType** | **int32** | SourceType should be a value as follows: 0: Network; 1: IP Group; 2: IP port Group | 
**Status** | **bool** | Status | 
**VirtualWanIds** | Pointer to **[]string** | Virtual WAN list. When interfaceType is 4, at least one of the wanPortIds or vpnIds or virtualWanIds is not empty. | [optional] 
**VpnIds** | Pointer to **[]string** | VPN list. When interfaceType is 4, at least one of the wanPortIds or vpnIds or virtualWanIds is not empty. | [optional] 
**WanPortIds** | Pointer to **[]string** | WAN port list. When interfaceType is 4, at least one of the wanPortIds or vpnIds or virtualWanIds is not empty. | [optional] 

## Methods

### NewPolicyRoutingInfo

`func NewPolicyRoutingInfo(backupInterface bool, destinationIds []string, destinationType int32, name string, protocols []int32, sourceIds []string, sourceType int32, status bool, ) *PolicyRoutingInfo`

NewPolicyRoutingInfo instantiates a new PolicyRoutingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRoutingInfoWithDefaults

`func NewPolicyRoutingInfoWithDefaults() *PolicyRoutingInfo`

NewPolicyRoutingInfoWithDefaults instantiates a new PolicyRoutingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupInterface

`func (o *PolicyRoutingInfo) GetBackupInterface() bool`

GetBackupInterface returns the BackupInterface field if non-nil, zero value otherwise.

### GetBackupInterfaceOk

`func (o *PolicyRoutingInfo) GetBackupInterfaceOk() (*bool, bool)`

GetBackupInterfaceOk returns a tuple with the BackupInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupInterface

`func (o *PolicyRoutingInfo) SetBackupInterface(v bool)`

SetBackupInterface sets BackupInterface field to given value.


### GetDestinationIds

`func (o *PolicyRoutingInfo) GetDestinationIds() []string`

GetDestinationIds returns the DestinationIds field if non-nil, zero value otherwise.

### GetDestinationIdsOk

`func (o *PolicyRoutingInfo) GetDestinationIdsOk() (*[]string, bool)`

GetDestinationIdsOk returns a tuple with the DestinationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIds

`func (o *PolicyRoutingInfo) SetDestinationIds(v []string)`

SetDestinationIds sets DestinationIds field to given value.


### GetDestinationType

`func (o *PolicyRoutingInfo) GetDestinationType() int32`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *PolicyRoutingInfo) GetDestinationTypeOk() (*int32, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *PolicyRoutingInfo) SetDestinationType(v int32)`

SetDestinationType sets DestinationType field to given value.


### GetExistDomainGroupDest

`func (o *PolicyRoutingInfo) GetExistDomainGroupDest() bool`

GetExistDomainGroupDest returns the ExistDomainGroupDest field if non-nil, zero value otherwise.

### GetExistDomainGroupDestOk

`func (o *PolicyRoutingInfo) GetExistDomainGroupDestOk() (*bool, bool)`

GetExistDomainGroupDestOk returns a tuple with the ExistDomainGroupDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistDomainGroupDest

`func (o *PolicyRoutingInfo) SetExistDomainGroupDest(v bool)`

SetExistDomainGroupDest sets ExistDomainGroupDest field to given value.

### HasExistDomainGroupDest

`func (o *PolicyRoutingInfo) HasExistDomainGroupDest() bool`

HasExistDomainGroupDest returns a boolean if a field has been set.

### GetExistIpPortGroup

`func (o *PolicyRoutingInfo) GetExistIpPortGroup() bool`

GetExistIpPortGroup returns the ExistIpPortGroup field if non-nil, zero value otherwise.

### GetExistIpPortGroupOk

`func (o *PolicyRoutingInfo) GetExistIpPortGroupOk() (*bool, bool)`

GetExistIpPortGroupOk returns a tuple with the ExistIpPortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistIpPortGroup

`func (o *PolicyRoutingInfo) SetExistIpPortGroup(v bool)`

SetExistIpPortGroup sets ExistIpPortGroup field to given value.

### HasExistIpPortGroup

`func (o *PolicyRoutingInfo) HasExistIpPortGroup() bool`

HasExistIpPortGroup returns a boolean if a field has been set.

### GetExistLocationGroupDest

`func (o *PolicyRoutingInfo) GetExistLocationGroupDest() bool`

GetExistLocationGroupDest returns the ExistLocationGroupDest field if non-nil, zero value otherwise.

### GetExistLocationGroupDestOk

`func (o *PolicyRoutingInfo) GetExistLocationGroupDestOk() (*bool, bool)`

GetExistLocationGroupDestOk returns a tuple with the ExistLocationGroupDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistLocationGroupDest

`func (o *PolicyRoutingInfo) SetExistLocationGroupDest(v bool)`

SetExistLocationGroupDest sets ExistLocationGroupDest field to given value.

### HasExistLocationGroupDest

`func (o *PolicyRoutingInfo) HasExistLocationGroupDest() bool`

HasExistLocationGroupDest returns a boolean if a field has been set.

### GetExistMulti

`func (o *PolicyRoutingInfo) GetExistMulti() bool`

GetExistMulti returns the ExistMulti field if non-nil, zero value otherwise.

### GetExistMultiOk

`func (o *PolicyRoutingInfo) GetExistMultiOk() (*bool, bool)`

GetExistMultiOk returns a tuple with the ExistMulti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistMulti

`func (o *PolicyRoutingInfo) SetExistMulti(v bool)`

SetExistMulti sets ExistMulti field to given value.

### HasExistMulti

`func (o *PolicyRoutingInfo) HasExistMulti() bool`

HasExistMulti returns a boolean if a field has been set.

### GetExistVirtualWan

`func (o *PolicyRoutingInfo) GetExistVirtualWan() bool`

GetExistVirtualWan returns the ExistVirtualWan field if non-nil, zero value otherwise.

### GetExistVirtualWanOk

`func (o *PolicyRoutingInfo) GetExistVirtualWanOk() (*bool, bool)`

GetExistVirtualWanOk returns a tuple with the ExistVirtualWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistVirtualWan

`func (o *PolicyRoutingInfo) SetExistVirtualWan(v bool)`

SetExistVirtualWan sets ExistVirtualWan field to given value.

### HasExistVirtualWan

`func (o *PolicyRoutingInfo) HasExistVirtualWan() bool`

HasExistVirtualWan returns a boolean if a field has been set.

### GetExistVpnClient

`func (o *PolicyRoutingInfo) GetExistVpnClient() bool`

GetExistVpnClient returns the ExistVpnClient field if non-nil, zero value otherwise.

### GetExistVpnClientOk

`func (o *PolicyRoutingInfo) GetExistVpnClientOk() (*bool, bool)`

GetExistVpnClientOk returns a tuple with the ExistVpnClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistVpnClient

`func (o *PolicyRoutingInfo) SetExistVpnClient(v bool)`

SetExistVpnClient sets ExistVpnClient field to given value.

### HasExistVpnClient

`func (o *PolicyRoutingInfo) HasExistVpnClient() bool`

HasExistVpnClient returns a boolean if a field has been set.

### GetId

`func (o *PolicyRoutingInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyRoutingInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyRoutingInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyRoutingInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndex

`func (o *PolicyRoutingInfo) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *PolicyRoutingInfo) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *PolicyRoutingInfo) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *PolicyRoutingInfo) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetInterfaceId

`func (o *PolicyRoutingInfo) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *PolicyRoutingInfo) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *PolicyRoutingInfo) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *PolicyRoutingInfo) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetInterfaceType

`func (o *PolicyRoutingInfo) GetInterfaceType() int32`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *PolicyRoutingInfo) GetInterfaceTypeOk() (*int32, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *PolicyRoutingInfo) SetInterfaceType(v int32)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *PolicyRoutingInfo) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetName

`func (o *PolicyRoutingInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyRoutingInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyRoutingInfo) SetName(v string)`

SetName sets Name field to given value.


### GetProtocols

`func (o *PolicyRoutingInfo) GetProtocols() []int32`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *PolicyRoutingInfo) GetProtocolsOk() (*[]int32, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *PolicyRoutingInfo) SetProtocols(v []int32)`

SetProtocols sets Protocols field to given value.


### GetSourceIds

`func (o *PolicyRoutingInfo) GetSourceIds() []string`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *PolicyRoutingInfo) GetSourceIdsOk() (*[]string, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *PolicyRoutingInfo) SetSourceIds(v []string)`

SetSourceIds sets SourceIds field to given value.


### GetSourceType

`func (o *PolicyRoutingInfo) GetSourceType() int32`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *PolicyRoutingInfo) GetSourceTypeOk() (*int32, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *PolicyRoutingInfo) SetSourceType(v int32)`

SetSourceType sets SourceType field to given value.


### GetStatus

`func (o *PolicyRoutingInfo) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PolicyRoutingInfo) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PolicyRoutingInfo) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetVirtualWanIds

`func (o *PolicyRoutingInfo) GetVirtualWanIds() []string`

GetVirtualWanIds returns the VirtualWanIds field if non-nil, zero value otherwise.

### GetVirtualWanIdsOk

`func (o *PolicyRoutingInfo) GetVirtualWanIdsOk() (*[]string, bool)`

GetVirtualWanIdsOk returns a tuple with the VirtualWanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanIds

`func (o *PolicyRoutingInfo) SetVirtualWanIds(v []string)`

SetVirtualWanIds sets VirtualWanIds field to given value.

### HasVirtualWanIds

`func (o *PolicyRoutingInfo) HasVirtualWanIds() bool`

HasVirtualWanIds returns a boolean if a field has been set.

### GetVpnIds

`func (o *PolicyRoutingInfo) GetVpnIds() []string`

GetVpnIds returns the VpnIds field if non-nil, zero value otherwise.

### GetVpnIdsOk

`func (o *PolicyRoutingInfo) GetVpnIdsOk() (*[]string, bool)`

GetVpnIdsOk returns a tuple with the VpnIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnIds

`func (o *PolicyRoutingInfo) SetVpnIds(v []string)`

SetVpnIds sets VpnIds field to given value.

### HasVpnIds

`func (o *PolicyRoutingInfo) HasVpnIds() bool`

HasVpnIds returns a boolean if a field has been set.

### GetWanPortIds

`func (o *PolicyRoutingInfo) GetWanPortIds() []string`

GetWanPortIds returns the WanPortIds field if non-nil, zero value otherwise.

### GetWanPortIdsOk

`func (o *PolicyRoutingInfo) GetWanPortIdsOk() (*[]string, bool)`

GetWanPortIdsOk returns a tuple with the WanPortIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortIds

`func (o *PolicyRoutingInfo) SetWanPortIds(v []string)`

SetWanPortIds sets WanPortIds field to given value.

### HasWanPortIds

`func (o *PolicyRoutingInfo) HasWanPortIds() bool`

HasWanPortIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


