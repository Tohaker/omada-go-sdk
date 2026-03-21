# GatewayCustomACLAddEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Custom ACL description should contain 1 to 512 characters. | 
**DestinationList** | **[]string** | Custom ACL destination list. | 
**DestinationPort** | Pointer to **string** | Custom ACL destination port. when \&quot;protocol\&quot; is \&quot;TCP\&quot; or \&quot;UDP\&quot;, port is valid. | [optional] 
**Direction** | **int32** | Custom ACL direction should be a value as follows: 0: LAN-LAN; 1: LAN-WAN; 2: WAN-LAN; 3: LOCAL IN | 
**Index** | **int32** | Custom ACL index | 
**LogStatus** | **bool** | Custom ACL log status | 
**Policy** | **int32** | Custom ACL policy should be a value as follows: 0: Drop; 1:Allow | 
**Protocol** | **int32** | Custom ACL protocol. For the values of protocol, refer to section 5.5 of the Open API Access Guide. | 
**SourceList** | **[]string** | Custom ACL source list. | 
**SourcePort** | Pointer to **string** | Custom ACL source port. when \&quot;protocol\&quot; is \&quot;TCP\&quot; or \&quot;UDP\&quot;, port is valid. | [optional] 
**Status** | **bool** | Custom ACL status | 

## Methods

### NewGatewayCustomACLAddEntity

`func NewGatewayCustomACLAddEntity(description string, destinationList []string, direction int32, index int32, logStatus bool, policy int32, protocol int32, sourceList []string, status bool, ) *GatewayCustomACLAddEntity`

NewGatewayCustomACLAddEntity instantiates a new GatewayCustomACLAddEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCustomACLAddEntityWithDefaults

`func NewGatewayCustomACLAddEntityWithDefaults() *GatewayCustomACLAddEntity`

NewGatewayCustomACLAddEntityWithDefaults instantiates a new GatewayCustomACLAddEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GatewayCustomACLAddEntity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GatewayCustomACLAddEntity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GatewayCustomACLAddEntity) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDestinationList

`func (o *GatewayCustomACLAddEntity) GetDestinationList() []string`

GetDestinationList returns the DestinationList field if non-nil, zero value otherwise.

### GetDestinationListOk

`func (o *GatewayCustomACLAddEntity) GetDestinationListOk() (*[]string, bool)`

GetDestinationListOk returns a tuple with the DestinationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationList

`func (o *GatewayCustomACLAddEntity) SetDestinationList(v []string)`

SetDestinationList sets DestinationList field to given value.


### GetDestinationPort

`func (o *GatewayCustomACLAddEntity) GetDestinationPort() string`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *GatewayCustomACLAddEntity) GetDestinationPortOk() (*string, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *GatewayCustomACLAddEntity) SetDestinationPort(v string)`

SetDestinationPort sets DestinationPort field to given value.

### HasDestinationPort

`func (o *GatewayCustomACLAddEntity) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.

### GetDirection

`func (o *GatewayCustomACLAddEntity) GetDirection() int32`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GatewayCustomACLAddEntity) GetDirectionOk() (*int32, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GatewayCustomACLAddEntity) SetDirection(v int32)`

SetDirection sets Direction field to given value.


### GetIndex

`func (o *GatewayCustomACLAddEntity) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *GatewayCustomACLAddEntity) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *GatewayCustomACLAddEntity) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetLogStatus

`func (o *GatewayCustomACLAddEntity) GetLogStatus() bool`

GetLogStatus returns the LogStatus field if non-nil, zero value otherwise.

### GetLogStatusOk

`func (o *GatewayCustomACLAddEntity) GetLogStatusOk() (*bool, bool)`

GetLogStatusOk returns a tuple with the LogStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStatus

`func (o *GatewayCustomACLAddEntity) SetLogStatus(v bool)`

SetLogStatus sets LogStatus field to given value.


### GetPolicy

`func (o *GatewayCustomACLAddEntity) GetPolicy() int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GatewayCustomACLAddEntity) GetPolicyOk() (*int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GatewayCustomACLAddEntity) SetPolicy(v int32)`

SetPolicy sets Policy field to given value.


### GetProtocol

`func (o *GatewayCustomACLAddEntity) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GatewayCustomACLAddEntity) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GatewayCustomACLAddEntity) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.


### GetSourceList

`func (o *GatewayCustomACLAddEntity) GetSourceList() []string`

GetSourceList returns the SourceList field if non-nil, zero value otherwise.

### GetSourceListOk

`func (o *GatewayCustomACLAddEntity) GetSourceListOk() (*[]string, bool)`

GetSourceListOk returns a tuple with the SourceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceList

`func (o *GatewayCustomACLAddEntity) SetSourceList(v []string)`

SetSourceList sets SourceList field to given value.


### GetSourcePort

`func (o *GatewayCustomACLAddEntity) GetSourcePort() string`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *GatewayCustomACLAddEntity) GetSourcePortOk() (*string, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePort

`func (o *GatewayCustomACLAddEntity) SetSourcePort(v string)`

SetSourcePort sets SourcePort field to given value.

### HasSourcePort

`func (o *GatewayCustomACLAddEntity) HasSourcePort() bool`

HasSourcePort returns a boolean if a field has been set.

### GetStatus

`func (o *GatewayCustomACLAddEntity) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GatewayCustomACLAddEntity) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GatewayCustomACLAddEntity) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


