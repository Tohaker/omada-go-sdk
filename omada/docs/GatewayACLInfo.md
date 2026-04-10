# GatewayACLInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | ACL rule description, description should contain 1 to 512 characters. | 
**DestinationIds** | **[]string** | Destination IDs, which depends on destinationType, for example: if destinationType is network, destinationIds should be LAN network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | 
**DestinationType** | **int32** | DestinationType should be a value as follows: 0: network; 1: IP Group; 2: IP-Port Group; 6: IPv6 Group; 7: IPv6-Port Group;10: Domain Group; 11: !Network; 12: !IP Group；13: !IP-Port Group；14: !IPv6 Group；15: !IPv6-port Group | 
**Direction** | Pointer to [**GatewayDirectionEntity**](GatewayDirectionEntity.md) |  | [optional] 
**Id** | **string** | ACL ID | 
**Index** | **int32** | Index | 
**Policy** | **int32** | Policy should be a value as follows: 0: drop; 1: allow; | 
**Protocols** | **[]int32** | For the values of protocols, refer to section 5.5 of the Open API Access Guide. | 
**SourceIds** | **[]string** | Source IDs, which depends on sourceType, for example: if sourceType is network, sourceIds should be LAN network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | 
**SourceType** | **int32** | SourceType should be a value as follows: 0: network; 1: IP Group; 2: IP-Port Group; 4: SSID; 6: IPv6 Group; 7: IPv6-Port Group; 8: Country; 9: Country Group; 11: !Network; 12: !IP Group; 13: !IP-Port Group; 14: !IPv6 Group; 15: !IPv6-port Group | 
**StateMode** | **int32** | StateMode should be a value as follows: 0: auto; 1: manual | 
**States** | Pointer to [**GatewayACLStatesEntity**](GatewayACLStatesEntity.md) |  | [optional] 
**Status** | **bool** | Status should be a value as follows: 0: disable; 1: enable | 
**Syslog** | **bool** | Enable remote log | 
**TimeRangeId** | Pointer to **string** | Gateway ACL time range ID. | [optional] 

## Methods

### NewGatewayACLInfo

`func NewGatewayACLInfo(description string, destinationIds []string, destinationType int32, id string, index int32, policy int32, protocols []int32, sourceIds []string, sourceType int32, stateMode int32, status bool, syslog bool, ) *GatewayACLInfo`

NewGatewayACLInfo instantiates a new GatewayACLInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayACLInfoWithDefaults

`func NewGatewayACLInfoWithDefaults() *GatewayACLInfo`

NewGatewayACLInfoWithDefaults instantiates a new GatewayACLInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GatewayACLInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GatewayACLInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GatewayACLInfo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDestinationIds

`func (o *GatewayACLInfo) GetDestinationIds() []string`

GetDestinationIds returns the DestinationIds field if non-nil, zero value otherwise.

### GetDestinationIdsOk

`func (o *GatewayACLInfo) GetDestinationIdsOk() (*[]string, bool)`

GetDestinationIdsOk returns a tuple with the DestinationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIds

`func (o *GatewayACLInfo) SetDestinationIds(v []string)`

SetDestinationIds sets DestinationIds field to given value.


### GetDestinationType

`func (o *GatewayACLInfo) GetDestinationType() int32`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *GatewayACLInfo) GetDestinationTypeOk() (*int32, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *GatewayACLInfo) SetDestinationType(v int32)`

SetDestinationType sets DestinationType field to given value.


### GetDirection

`func (o *GatewayACLInfo) GetDirection() GatewayDirectionEntity`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GatewayACLInfo) GetDirectionOk() (*GatewayDirectionEntity, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GatewayACLInfo) SetDirection(v GatewayDirectionEntity)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GatewayACLInfo) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetId

`func (o *GatewayACLInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayACLInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayACLInfo) SetId(v string)`

SetId sets Id field to given value.


### GetIndex

`func (o *GatewayACLInfo) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *GatewayACLInfo) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *GatewayACLInfo) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetPolicy

`func (o *GatewayACLInfo) GetPolicy() int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GatewayACLInfo) GetPolicyOk() (*int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GatewayACLInfo) SetPolicy(v int32)`

SetPolicy sets Policy field to given value.


### GetProtocols

`func (o *GatewayACLInfo) GetProtocols() []int32`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *GatewayACLInfo) GetProtocolsOk() (*[]int32, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *GatewayACLInfo) SetProtocols(v []int32)`

SetProtocols sets Protocols field to given value.


### GetSourceIds

`func (o *GatewayACLInfo) GetSourceIds() []string`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *GatewayACLInfo) GetSourceIdsOk() (*[]string, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *GatewayACLInfo) SetSourceIds(v []string)`

SetSourceIds sets SourceIds field to given value.


### GetSourceType

`func (o *GatewayACLInfo) GetSourceType() int32`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *GatewayACLInfo) GetSourceTypeOk() (*int32, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *GatewayACLInfo) SetSourceType(v int32)`

SetSourceType sets SourceType field to given value.


### GetStateMode

`func (o *GatewayACLInfo) GetStateMode() int32`

GetStateMode returns the StateMode field if non-nil, zero value otherwise.

### GetStateModeOk

`func (o *GatewayACLInfo) GetStateModeOk() (*int32, bool)`

GetStateModeOk returns a tuple with the StateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMode

`func (o *GatewayACLInfo) SetStateMode(v int32)`

SetStateMode sets StateMode field to given value.


### GetStates

`func (o *GatewayACLInfo) GetStates() GatewayACLStatesEntity`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *GatewayACLInfo) GetStatesOk() (*GatewayACLStatesEntity, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *GatewayACLInfo) SetStates(v GatewayACLStatesEntity)`

SetStates sets States field to given value.

### HasStates

`func (o *GatewayACLInfo) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetStatus

`func (o *GatewayACLInfo) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GatewayACLInfo) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GatewayACLInfo) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetSyslog

`func (o *GatewayACLInfo) GetSyslog() bool`

GetSyslog returns the Syslog field if non-nil, zero value otherwise.

### GetSyslogOk

`func (o *GatewayACLInfo) GetSyslogOk() (*bool, bool)`

GetSyslogOk returns a tuple with the Syslog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslog

`func (o *GatewayACLInfo) SetSyslog(v bool)`

SetSyslog sets Syslog field to given value.


### GetTimeRangeId

`func (o *GatewayACLInfo) GetTimeRangeId() string`

GetTimeRangeId returns the TimeRangeId field if non-nil, zero value otherwise.

### GetTimeRangeIdOk

`func (o *GatewayACLInfo) GetTimeRangeIdOk() (*string, bool)`

GetTimeRangeIdOk returns a tuple with the TimeRangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRangeId

`func (o *GatewayACLInfo) SetTimeRangeId(v string)`

SetTimeRangeId sets TimeRangeId field to given value.

### HasTimeRangeId

`func (o *GatewayACLInfo) HasTimeRangeId() bool`

HasTimeRangeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


