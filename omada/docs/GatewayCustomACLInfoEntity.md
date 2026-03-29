# GatewayCustomACLInfoEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Custom ACL description should contain 1 to 512 characters. | 
**DestinationList** | **[]string** | Custom ACL destination list. | 
**DestinationPort** | Pointer to **string** | Custom ACL destination port. when \&quot;protocol\&quot; is \&quot;TCP\&quot; or \&quot;UDP\&quot;, port is valid. | [optional] 
**Direction** | **int32** | Custom ACL direction should be a value as follows: 0: LAN-LAN; 1: LAN-WAN; 2: WAN-LAN; 3: LOCAL IN | 
**HitCounts** | Pointer to **int32** | Show custom ACL hit counts. | [optional] 
**Id** | **string** | Custom ACL ID | 
**Index** | **int32** | Custom ACL index | 
**LogStatus** | **bool** | Custom ACL log status | 
**Policy** | **int32** | Custom ACL policy should be a value as follows: 0: Drop; 1:Allow | 
**Protocol** | **int32** | Custom ACL protocol. For the values of protocol, refer to section 5.5 of the Open API Access Guide. | 
**SourceList** | **[]string** | Custom ACL source list. | 
**SourcePort** | Pointer to **string** | Custom ACL source port. when \&quot;protocol\&quot; is \&quot;TCP\&quot; or \&quot;UDP\&quot;, port is valid. | [optional] 
**Status** | **bool** | Custom ACL status | 

## Methods

### NewGatewayCustomACLInfoEntity

`func NewGatewayCustomACLInfoEntity(description string, destinationList []string, direction int32, id string, index int32, logStatus bool, policy int32, protocol int32, sourceList []string, status bool, ) *GatewayCustomACLInfoEntity`

NewGatewayCustomACLInfoEntity instantiates a new GatewayCustomACLInfoEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCustomACLInfoEntityWithDefaults

`func NewGatewayCustomACLInfoEntityWithDefaults() *GatewayCustomACLInfoEntity`

NewGatewayCustomACLInfoEntityWithDefaults instantiates a new GatewayCustomACLInfoEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GatewayCustomACLInfoEntity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GatewayCustomACLInfoEntity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GatewayCustomACLInfoEntity) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDestinationList

`func (o *GatewayCustomACLInfoEntity) GetDestinationList() []string`

GetDestinationList returns the DestinationList field if non-nil, zero value otherwise.

### GetDestinationListOk

`func (o *GatewayCustomACLInfoEntity) GetDestinationListOk() (*[]string, bool)`

GetDestinationListOk returns a tuple with the DestinationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationList

`func (o *GatewayCustomACLInfoEntity) SetDestinationList(v []string)`

SetDestinationList sets DestinationList field to given value.


### GetDestinationPort

`func (o *GatewayCustomACLInfoEntity) GetDestinationPort() string`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *GatewayCustomACLInfoEntity) GetDestinationPortOk() (*string, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *GatewayCustomACLInfoEntity) SetDestinationPort(v string)`

SetDestinationPort sets DestinationPort field to given value.

### HasDestinationPort

`func (o *GatewayCustomACLInfoEntity) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.

### GetDirection

`func (o *GatewayCustomACLInfoEntity) GetDirection() int32`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GatewayCustomACLInfoEntity) GetDirectionOk() (*int32, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GatewayCustomACLInfoEntity) SetDirection(v int32)`

SetDirection sets Direction field to given value.


### GetHitCounts

`func (o *GatewayCustomACLInfoEntity) GetHitCounts() int32`

GetHitCounts returns the HitCounts field if non-nil, zero value otherwise.

### GetHitCountsOk

`func (o *GatewayCustomACLInfoEntity) GetHitCountsOk() (*int32, bool)`

GetHitCountsOk returns a tuple with the HitCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitCounts

`func (o *GatewayCustomACLInfoEntity) SetHitCounts(v int32)`

SetHitCounts sets HitCounts field to given value.

### HasHitCounts

`func (o *GatewayCustomACLInfoEntity) HasHitCounts() bool`

HasHitCounts returns a boolean if a field has been set.

### GetId

`func (o *GatewayCustomACLInfoEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayCustomACLInfoEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayCustomACLInfoEntity) SetId(v string)`

SetId sets Id field to given value.


### GetIndex

`func (o *GatewayCustomACLInfoEntity) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *GatewayCustomACLInfoEntity) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *GatewayCustomACLInfoEntity) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetLogStatus

`func (o *GatewayCustomACLInfoEntity) GetLogStatus() bool`

GetLogStatus returns the LogStatus field if non-nil, zero value otherwise.

### GetLogStatusOk

`func (o *GatewayCustomACLInfoEntity) GetLogStatusOk() (*bool, bool)`

GetLogStatusOk returns a tuple with the LogStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStatus

`func (o *GatewayCustomACLInfoEntity) SetLogStatus(v bool)`

SetLogStatus sets LogStatus field to given value.


### GetPolicy

`func (o *GatewayCustomACLInfoEntity) GetPolicy() int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GatewayCustomACLInfoEntity) GetPolicyOk() (*int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GatewayCustomACLInfoEntity) SetPolicy(v int32)`

SetPolicy sets Policy field to given value.


### GetProtocol

`func (o *GatewayCustomACLInfoEntity) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GatewayCustomACLInfoEntity) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GatewayCustomACLInfoEntity) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.


### GetSourceList

`func (o *GatewayCustomACLInfoEntity) GetSourceList() []string`

GetSourceList returns the SourceList field if non-nil, zero value otherwise.

### GetSourceListOk

`func (o *GatewayCustomACLInfoEntity) GetSourceListOk() (*[]string, bool)`

GetSourceListOk returns a tuple with the SourceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceList

`func (o *GatewayCustomACLInfoEntity) SetSourceList(v []string)`

SetSourceList sets SourceList field to given value.


### GetSourcePort

`func (o *GatewayCustomACLInfoEntity) GetSourcePort() string`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *GatewayCustomACLInfoEntity) GetSourcePortOk() (*string, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePort

`func (o *GatewayCustomACLInfoEntity) SetSourcePort(v string)`

SetSourcePort sets SourcePort field to given value.

### HasSourcePort

`func (o *GatewayCustomACLInfoEntity) HasSourcePort() bool`

HasSourcePort returns a boolean if a field has been set.

### GetStatus

`func (o *GatewayCustomACLInfoEntity) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GatewayCustomACLInfoEntity) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GatewayCustomACLInfoEntity) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


