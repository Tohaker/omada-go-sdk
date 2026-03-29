# GatewayACLConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | ACL rule description, description should contain 1 to 512 characters. | 
**DestinationIds** | Pointer to **[]string** | Source IDs, which depends on destinationType, for example: if destinationType is network, destinationIds should be LAN network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | [optional] 
**DestinationType** | **int32** | DestinationType should be a value as follows: 0: network; 1: IP Group; 2: IP-Port Group; 6: IPv6 Group; 7: IPv6-Port Group;10: Domain Group. | 
**Direction** | [**GatewayDirectionEntity**](GatewayDirectionEntity.md) |  | 
**Policy** | **int32** | Policy should be a value as follows: 0: drop; 1: allow; | 
**Protocols** | **[]int32** | For the values of protocols, refer to section 5.5 of the Open API Access Guide. | 
**SourceIds** | **[]string** | Source IDs, which depends on sourceType, for example: if sourceType is network, sourceIds should be LAN network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | 
**SourceType** | **int32** | SourceType should be a value as follows: 0: network; 1: IP Group; 2: IP-Port Group; 4: SSID; 6: IPv6 Group; 7: IPv6-Port Group. | 
**StateMode** | **int32** | StateMode should be a value as follows: 0: auto; 1: manual | 
**States** | Pointer to [**GatewayACLStatesEntity**](GatewayACLStatesEntity.md) |  | [optional] 
**Status** | **bool** | Status should be a value as follows: 0: disable; 1: enable | 
**Syslog** | **bool** | Enable remote log | 
**TimeRangeId** | Pointer to **string** | Gateway ACL time range ID. | [optional] 

## Methods

### NewGatewayACLConfig

`func NewGatewayACLConfig(description string, destinationType int32, direction GatewayDirectionEntity, policy int32, protocols []int32, sourceIds []string, sourceType int32, stateMode int32, status bool, syslog bool, ) *GatewayACLConfig`

NewGatewayACLConfig instantiates a new GatewayACLConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayACLConfigWithDefaults

`func NewGatewayACLConfigWithDefaults() *GatewayACLConfig`

NewGatewayACLConfigWithDefaults instantiates a new GatewayACLConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GatewayACLConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GatewayACLConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GatewayACLConfig) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDestinationIds

`func (o *GatewayACLConfig) GetDestinationIds() []string`

GetDestinationIds returns the DestinationIds field if non-nil, zero value otherwise.

### GetDestinationIdsOk

`func (o *GatewayACLConfig) GetDestinationIdsOk() (*[]string, bool)`

GetDestinationIdsOk returns a tuple with the DestinationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIds

`func (o *GatewayACLConfig) SetDestinationIds(v []string)`

SetDestinationIds sets DestinationIds field to given value.

### HasDestinationIds

`func (o *GatewayACLConfig) HasDestinationIds() bool`

HasDestinationIds returns a boolean if a field has been set.

### GetDestinationType

`func (o *GatewayACLConfig) GetDestinationType() int32`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *GatewayACLConfig) GetDestinationTypeOk() (*int32, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *GatewayACLConfig) SetDestinationType(v int32)`

SetDestinationType sets DestinationType field to given value.


### GetDirection

`func (o *GatewayACLConfig) GetDirection() GatewayDirectionEntity`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GatewayACLConfig) GetDirectionOk() (*GatewayDirectionEntity, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GatewayACLConfig) SetDirection(v GatewayDirectionEntity)`

SetDirection sets Direction field to given value.


### GetPolicy

`func (o *GatewayACLConfig) GetPolicy() int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GatewayACLConfig) GetPolicyOk() (*int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GatewayACLConfig) SetPolicy(v int32)`

SetPolicy sets Policy field to given value.


### GetProtocols

`func (o *GatewayACLConfig) GetProtocols() []int32`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *GatewayACLConfig) GetProtocolsOk() (*[]int32, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *GatewayACLConfig) SetProtocols(v []int32)`

SetProtocols sets Protocols field to given value.


### GetSourceIds

`func (o *GatewayACLConfig) GetSourceIds() []string`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *GatewayACLConfig) GetSourceIdsOk() (*[]string, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *GatewayACLConfig) SetSourceIds(v []string)`

SetSourceIds sets SourceIds field to given value.


### GetSourceType

`func (o *GatewayACLConfig) GetSourceType() int32`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *GatewayACLConfig) GetSourceTypeOk() (*int32, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *GatewayACLConfig) SetSourceType(v int32)`

SetSourceType sets SourceType field to given value.


### GetStateMode

`func (o *GatewayACLConfig) GetStateMode() int32`

GetStateMode returns the StateMode field if non-nil, zero value otherwise.

### GetStateModeOk

`func (o *GatewayACLConfig) GetStateModeOk() (*int32, bool)`

GetStateModeOk returns a tuple with the StateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMode

`func (o *GatewayACLConfig) SetStateMode(v int32)`

SetStateMode sets StateMode field to given value.


### GetStates

`func (o *GatewayACLConfig) GetStates() GatewayACLStatesEntity`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *GatewayACLConfig) GetStatesOk() (*GatewayACLStatesEntity, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *GatewayACLConfig) SetStates(v GatewayACLStatesEntity)`

SetStates sets States field to given value.

### HasStates

`func (o *GatewayACLConfig) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetStatus

`func (o *GatewayACLConfig) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GatewayACLConfig) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GatewayACLConfig) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetSyslog

`func (o *GatewayACLConfig) GetSyslog() bool`

GetSyslog returns the Syslog field if non-nil, zero value otherwise.

### GetSyslogOk

`func (o *GatewayACLConfig) GetSyslogOk() (*bool, bool)`

GetSyslogOk returns a tuple with the Syslog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslog

`func (o *GatewayACLConfig) SetSyslog(v bool)`

SetSyslog sets Syslog field to given value.


### GetTimeRangeId

`func (o *GatewayACLConfig) GetTimeRangeId() string`

GetTimeRangeId returns the TimeRangeId field if non-nil, zero value otherwise.

### GetTimeRangeIdOk

`func (o *GatewayACLConfig) GetTimeRangeIdOk() (*string, bool)`

GetTimeRangeIdOk returns a tuple with the TimeRangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRangeId

`func (o *GatewayACLConfig) SetTimeRangeId(v string)`

SetTimeRangeId sets TimeRangeId field to given value.

### HasTimeRangeId

`func (o *GatewayACLConfig) HasTimeRangeId() bool`

HasTimeRangeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


