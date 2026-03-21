# PolicyRoutingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupInterface** | **bool** | Use the other WAN port if the current one is down | 
**DestinationIds** | **[]string** | Destination IDs, which depends on destinationType, for example: if destinationType is network, destinationIds should be LAN network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | 
**DestinationType** | **int32** | DestinationType should be a value as follows: 0: Network; 1: IP Group; 2: IP port Group | 
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

### NewPolicyRoutingConfig

`func NewPolicyRoutingConfig(backupInterface bool, destinationIds []string, destinationType int32, name string, protocols []int32, sourceIds []string, sourceType int32, status bool, ) *PolicyRoutingConfig`

NewPolicyRoutingConfig instantiates a new PolicyRoutingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRoutingConfigWithDefaults

`func NewPolicyRoutingConfigWithDefaults() *PolicyRoutingConfig`

NewPolicyRoutingConfigWithDefaults instantiates a new PolicyRoutingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupInterface

`func (o *PolicyRoutingConfig) GetBackupInterface() bool`

GetBackupInterface returns the BackupInterface field if non-nil, zero value otherwise.

### GetBackupInterfaceOk

`func (o *PolicyRoutingConfig) GetBackupInterfaceOk() (*bool, bool)`

GetBackupInterfaceOk returns a tuple with the BackupInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupInterface

`func (o *PolicyRoutingConfig) SetBackupInterface(v bool)`

SetBackupInterface sets BackupInterface field to given value.


### GetDestinationIds

`func (o *PolicyRoutingConfig) GetDestinationIds() []string`

GetDestinationIds returns the DestinationIds field if non-nil, zero value otherwise.

### GetDestinationIdsOk

`func (o *PolicyRoutingConfig) GetDestinationIdsOk() (*[]string, bool)`

GetDestinationIdsOk returns a tuple with the DestinationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIds

`func (o *PolicyRoutingConfig) SetDestinationIds(v []string)`

SetDestinationIds sets DestinationIds field to given value.


### GetDestinationType

`func (o *PolicyRoutingConfig) GetDestinationType() int32`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *PolicyRoutingConfig) GetDestinationTypeOk() (*int32, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *PolicyRoutingConfig) SetDestinationType(v int32)`

SetDestinationType sets DestinationType field to given value.


### GetInterfaceId

`func (o *PolicyRoutingConfig) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *PolicyRoutingConfig) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *PolicyRoutingConfig) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *PolicyRoutingConfig) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetInterfaceType

`func (o *PolicyRoutingConfig) GetInterfaceType() int32`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *PolicyRoutingConfig) GetInterfaceTypeOk() (*int32, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *PolicyRoutingConfig) SetInterfaceType(v int32)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *PolicyRoutingConfig) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetName

`func (o *PolicyRoutingConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyRoutingConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyRoutingConfig) SetName(v string)`

SetName sets Name field to given value.


### GetProtocols

`func (o *PolicyRoutingConfig) GetProtocols() []int32`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *PolicyRoutingConfig) GetProtocolsOk() (*[]int32, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *PolicyRoutingConfig) SetProtocols(v []int32)`

SetProtocols sets Protocols field to given value.


### GetSourceIds

`func (o *PolicyRoutingConfig) GetSourceIds() []string`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *PolicyRoutingConfig) GetSourceIdsOk() (*[]string, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *PolicyRoutingConfig) SetSourceIds(v []string)`

SetSourceIds sets SourceIds field to given value.


### GetSourceType

`func (o *PolicyRoutingConfig) GetSourceType() int32`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *PolicyRoutingConfig) GetSourceTypeOk() (*int32, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *PolicyRoutingConfig) SetSourceType(v int32)`

SetSourceType sets SourceType field to given value.


### GetStatus

`func (o *PolicyRoutingConfig) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PolicyRoutingConfig) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PolicyRoutingConfig) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetVirtualWanIds

`func (o *PolicyRoutingConfig) GetVirtualWanIds() []string`

GetVirtualWanIds returns the VirtualWanIds field if non-nil, zero value otherwise.

### GetVirtualWanIdsOk

`func (o *PolicyRoutingConfig) GetVirtualWanIdsOk() (*[]string, bool)`

GetVirtualWanIdsOk returns a tuple with the VirtualWanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanIds

`func (o *PolicyRoutingConfig) SetVirtualWanIds(v []string)`

SetVirtualWanIds sets VirtualWanIds field to given value.

### HasVirtualWanIds

`func (o *PolicyRoutingConfig) HasVirtualWanIds() bool`

HasVirtualWanIds returns a boolean if a field has been set.

### GetVpnIds

`func (o *PolicyRoutingConfig) GetVpnIds() []string`

GetVpnIds returns the VpnIds field if non-nil, zero value otherwise.

### GetVpnIdsOk

`func (o *PolicyRoutingConfig) GetVpnIdsOk() (*[]string, bool)`

GetVpnIdsOk returns a tuple with the VpnIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnIds

`func (o *PolicyRoutingConfig) SetVpnIds(v []string)`

SetVpnIds sets VpnIds field to given value.

### HasVpnIds

`func (o *PolicyRoutingConfig) HasVpnIds() bool`

HasVpnIds returns a boolean if a field has been set.

### GetWanPortIds

`func (o *PolicyRoutingConfig) GetWanPortIds() []string`

GetWanPortIds returns the WanPortIds field if non-nil, zero value otherwise.

### GetWanPortIdsOk

`func (o *PolicyRoutingConfig) GetWanPortIdsOk() (*[]string, bool)`

GetWanPortIdsOk returns a tuple with the WanPortIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortIds

`func (o *PolicyRoutingConfig) SetWanPortIds(v []string)`

SetWanPortIds sets WanPortIds field to given value.

### HasWanPortIds

`func (o *PolicyRoutingConfig) HasWanPortIds() bool`

HasWanPortIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


