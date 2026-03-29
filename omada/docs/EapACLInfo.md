# EapACLInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | ACL rule description, description should contain 1 to 512 characters. | 
**DestinationIds** | **[]string** | Destination IDs, which depends on destinationType, for example: if destinationType is network, destinationIds should be LAN network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | 
**DestinationType** | **int32** | DestinationType should be a value as follows: 0: network; 1: IP Group; 2: IP-Port Group; 6: IPv6 Group; 7: IPv6-Port Group | 
**Id** | **string** | ACL ID | 
**Index** | **int32** | Index | 
**Policy** | **int32** | Policy should be a value as follows: 0: drop; 1: allow; | 
**Protocols** | **[]int32** | For the values of protocols, refer to section 5.5 of the Open API Access Guide. | 
**SourceIds** | **[]string** | Source IDs, which depends on sourceType, for example: if sourceType is network, sourceIds should be LAN network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | 
**SourceType** | **int32** | SourceType should be a value as follows: 0: network; 1: IP Group; 2: IP-Port Group; 4: SSID; 6: IPv6 Group; 7: IPv6-Port Group. | 
**Status** | **bool** | Status should be a value as follows: 0: disable; 1: enable | 

## Methods

### NewEapACLInfo

`func NewEapACLInfo(description string, destinationIds []string, destinationType int32, id string, index int32, policy int32, protocols []int32, sourceIds []string, sourceType int32, status bool, ) *EapACLInfo`

NewEapACLInfo instantiates a new EapACLInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEapACLInfoWithDefaults

`func NewEapACLInfoWithDefaults() *EapACLInfo`

NewEapACLInfoWithDefaults instantiates a new EapACLInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EapACLInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EapACLInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EapACLInfo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDestinationIds

`func (o *EapACLInfo) GetDestinationIds() []string`

GetDestinationIds returns the DestinationIds field if non-nil, zero value otherwise.

### GetDestinationIdsOk

`func (o *EapACLInfo) GetDestinationIdsOk() (*[]string, bool)`

GetDestinationIdsOk returns a tuple with the DestinationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIds

`func (o *EapACLInfo) SetDestinationIds(v []string)`

SetDestinationIds sets DestinationIds field to given value.


### GetDestinationType

`func (o *EapACLInfo) GetDestinationType() int32`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *EapACLInfo) GetDestinationTypeOk() (*int32, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *EapACLInfo) SetDestinationType(v int32)`

SetDestinationType sets DestinationType field to given value.


### GetId

`func (o *EapACLInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EapACLInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EapACLInfo) SetId(v string)`

SetId sets Id field to given value.


### GetIndex

`func (o *EapACLInfo) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *EapACLInfo) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *EapACLInfo) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetPolicy

`func (o *EapACLInfo) GetPolicy() int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *EapACLInfo) GetPolicyOk() (*int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *EapACLInfo) SetPolicy(v int32)`

SetPolicy sets Policy field to given value.


### GetProtocols

`func (o *EapACLInfo) GetProtocols() []int32`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *EapACLInfo) GetProtocolsOk() (*[]int32, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *EapACLInfo) SetProtocols(v []int32)`

SetProtocols sets Protocols field to given value.


### GetSourceIds

`func (o *EapACLInfo) GetSourceIds() []string`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *EapACLInfo) GetSourceIdsOk() (*[]string, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *EapACLInfo) SetSourceIds(v []string)`

SetSourceIds sets SourceIds field to given value.


### GetSourceType

`func (o *EapACLInfo) GetSourceType() int32`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *EapACLInfo) GetSourceTypeOk() (*int32, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *EapACLInfo) SetSourceType(v int32)`

SetSourceType sets SourceType field to given value.


### GetStatus

`func (o *EapACLInfo) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EapACLInfo) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EapACLInfo) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


