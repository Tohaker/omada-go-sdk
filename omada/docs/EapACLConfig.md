# EapACLConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | ACL rule description, description should contain 1 to 512 characters. | 
**DestinationIds** | Pointer to **[]string** | Source IDs, which depends on destinationType, for example: if destinationType is network, destinationIds should be LAN network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | [optional] 
**DestinationType** | **int32** | DestinationType should be a value as follows: 0: network; 1: IP Group; 2: IP-Port Group; 6: IPv6 Group; 7: IPv6-Port Group | 
**Policy** | **int32** | Policy should be a value as follows: 0: drop; 1: allow; | 
**Protocols** | **[]int32** | For the values of protocols, refer to section 5.5 of the Open API Access Guide. | 
**SourceIds** | **[]string** | Source IDs, which depends on sourceType, for example: if sourceType is network, sourceIds should be LAN network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | 
**SourceType** | **int32** | SourceType should be a value as follows: 0: network; 1: IP Group; 2: IP-Port Group; 4: SSID; 6: IPv6 Group; 7: IPv6-Port Group | 
**Status** | **bool** | Status should be a value as follows: 0: disable; 1: enable | 

## Methods

### NewEapACLConfig

`func NewEapACLConfig(description string, destinationType int32, policy int32, protocols []int32, sourceIds []string, sourceType int32, status bool, ) *EapACLConfig`

NewEapACLConfig instantiates a new EapACLConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEapACLConfigWithDefaults

`func NewEapACLConfigWithDefaults() *EapACLConfig`

NewEapACLConfigWithDefaults instantiates a new EapACLConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EapACLConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EapACLConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EapACLConfig) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDestinationIds

`func (o *EapACLConfig) GetDestinationIds() []string`

GetDestinationIds returns the DestinationIds field if non-nil, zero value otherwise.

### GetDestinationIdsOk

`func (o *EapACLConfig) GetDestinationIdsOk() (*[]string, bool)`

GetDestinationIdsOk returns a tuple with the DestinationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIds

`func (o *EapACLConfig) SetDestinationIds(v []string)`

SetDestinationIds sets DestinationIds field to given value.

### HasDestinationIds

`func (o *EapACLConfig) HasDestinationIds() bool`

HasDestinationIds returns a boolean if a field has been set.

### GetDestinationType

`func (o *EapACLConfig) GetDestinationType() int32`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *EapACLConfig) GetDestinationTypeOk() (*int32, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *EapACLConfig) SetDestinationType(v int32)`

SetDestinationType sets DestinationType field to given value.


### GetPolicy

`func (o *EapACLConfig) GetPolicy() int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *EapACLConfig) GetPolicyOk() (*int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *EapACLConfig) SetPolicy(v int32)`

SetPolicy sets Policy field to given value.


### GetProtocols

`func (o *EapACLConfig) GetProtocols() []int32`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *EapACLConfig) GetProtocolsOk() (*[]int32, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *EapACLConfig) SetProtocols(v []int32)`

SetProtocols sets Protocols field to given value.


### GetSourceIds

`func (o *EapACLConfig) GetSourceIds() []string`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *EapACLConfig) GetSourceIdsOk() (*[]string, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *EapACLConfig) SetSourceIds(v []string)`

SetSourceIds sets SourceIds field to given value.


### GetSourceType

`func (o *EapACLConfig) GetSourceType() int32`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *EapACLConfig) GetSourceTypeOk() (*int32, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *EapACLConfig) SetSourceType(v int32)`

SetSourceType sets SourceType field to given value.


### GetStatus

`func (o *EapACLConfig) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EapACLConfig) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EapACLConfig) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


