# GatewayCustomACLModifyEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Custom ACL description should contain 1 to 512 characters. | 
**DestinationList** | **[]string** | Custom ACL destination list. | 
**DestinationPort** | Pointer to **string** | Custom ACL destination port. when \&quot;protocol\&quot; is \&quot;TCP\&quot; or \&quot;UDP\&quot;, port is valid. | [optional] 
**Direction** | **int32** | Custom ACL direction should be a value as follows: 0: LAN-LAN; 1: LAN-WAN; 2: WAN-LAN; 3: LOCAL IN | 
**Id** | **string** | Custom ACL ID | 
**Index** | **int32** | Custom ACL index | 
**LogStatus** | **bool** | Custom ACL log status | 
**Policy** | **int32** | Custom ACL policy should be a value as follows: 0: Drop; 1:Allow | 
**Protocol** | **int32** | Custom ACL protocol. For the values of protocol, refer to section 5.5 of the Open API Access Guide. | 
**SourceList** | **[]string** | Custom ACL source list. | 
**SourcePort** | Pointer to **string** | Custom ACL source port. when \&quot;protocol\&quot; is \&quot;TCP\&quot; or \&quot;UDP\&quot;, port is valid. | [optional] 
**Status** | **bool** | Custom ACL status | 

## Methods

### NewGatewayCustomACLModifyEntity

`func NewGatewayCustomACLModifyEntity(description string, destinationList []string, direction int32, id string, index int32, logStatus bool, policy int32, protocol int32, sourceList []string, status bool, ) *GatewayCustomACLModifyEntity`

NewGatewayCustomACLModifyEntity instantiates a new GatewayCustomACLModifyEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCustomACLModifyEntityWithDefaults

`func NewGatewayCustomACLModifyEntityWithDefaults() *GatewayCustomACLModifyEntity`

NewGatewayCustomACLModifyEntityWithDefaults instantiates a new GatewayCustomACLModifyEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GatewayCustomACLModifyEntity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GatewayCustomACLModifyEntity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GatewayCustomACLModifyEntity) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDestinationList

`func (o *GatewayCustomACLModifyEntity) GetDestinationList() []string`

GetDestinationList returns the DestinationList field if non-nil, zero value otherwise.

### GetDestinationListOk

`func (o *GatewayCustomACLModifyEntity) GetDestinationListOk() (*[]string, bool)`

GetDestinationListOk returns a tuple with the DestinationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationList

`func (o *GatewayCustomACLModifyEntity) SetDestinationList(v []string)`

SetDestinationList sets DestinationList field to given value.


### GetDestinationPort

`func (o *GatewayCustomACLModifyEntity) GetDestinationPort() string`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *GatewayCustomACLModifyEntity) GetDestinationPortOk() (*string, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *GatewayCustomACLModifyEntity) SetDestinationPort(v string)`

SetDestinationPort sets DestinationPort field to given value.

### HasDestinationPort

`func (o *GatewayCustomACLModifyEntity) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.

### GetDirection

`func (o *GatewayCustomACLModifyEntity) GetDirection() int32`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GatewayCustomACLModifyEntity) GetDirectionOk() (*int32, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GatewayCustomACLModifyEntity) SetDirection(v int32)`

SetDirection sets Direction field to given value.


### GetId

`func (o *GatewayCustomACLModifyEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayCustomACLModifyEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayCustomACLModifyEntity) SetId(v string)`

SetId sets Id field to given value.


### GetIndex

`func (o *GatewayCustomACLModifyEntity) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *GatewayCustomACLModifyEntity) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *GatewayCustomACLModifyEntity) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetLogStatus

`func (o *GatewayCustomACLModifyEntity) GetLogStatus() bool`

GetLogStatus returns the LogStatus field if non-nil, zero value otherwise.

### GetLogStatusOk

`func (o *GatewayCustomACLModifyEntity) GetLogStatusOk() (*bool, bool)`

GetLogStatusOk returns a tuple with the LogStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStatus

`func (o *GatewayCustomACLModifyEntity) SetLogStatus(v bool)`

SetLogStatus sets LogStatus field to given value.


### GetPolicy

`func (o *GatewayCustomACLModifyEntity) GetPolicy() int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GatewayCustomACLModifyEntity) GetPolicyOk() (*int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GatewayCustomACLModifyEntity) SetPolicy(v int32)`

SetPolicy sets Policy field to given value.


### GetProtocol

`func (o *GatewayCustomACLModifyEntity) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GatewayCustomACLModifyEntity) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GatewayCustomACLModifyEntity) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.


### GetSourceList

`func (o *GatewayCustomACLModifyEntity) GetSourceList() []string`

GetSourceList returns the SourceList field if non-nil, zero value otherwise.

### GetSourceListOk

`func (o *GatewayCustomACLModifyEntity) GetSourceListOk() (*[]string, bool)`

GetSourceListOk returns a tuple with the SourceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceList

`func (o *GatewayCustomACLModifyEntity) SetSourceList(v []string)`

SetSourceList sets SourceList field to given value.


### GetSourcePort

`func (o *GatewayCustomACLModifyEntity) GetSourcePort() string`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *GatewayCustomACLModifyEntity) GetSourcePortOk() (*string, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePort

`func (o *GatewayCustomACLModifyEntity) SetSourcePort(v string)`

SetSourcePort sets SourcePort field to given value.

### HasSourcePort

`func (o *GatewayCustomACLModifyEntity) HasSourcePort() bool`

HasSourcePort returns a boolean if a field has been set.

### GetStatus

`func (o *GatewayCustomACLModifyEntity) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GatewayCustomACLModifyEntity) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GatewayCustomACLModifyEntity) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


