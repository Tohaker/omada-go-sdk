# FreeAuthClientPolicyOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIp** | Pointer to **string** | Free auth client IP Address | [optional] 
**ClientMac** | Pointer to **string** | Free auth client MAC Address, for example: AA-AA-AA-AA-AA-AA | [optional] 
**IdInt** | Pointer to **int32** | Entry ID of the policy. Except for newly added policies, this parameter should be retained | [optional] 
**Type** | **int32** | Type of the policy. It should be a value as follows: 3: Free auth client IP, and parameter [clientIp] is needed. 4: Free auth client MAC, and parameter [clientMac] is needed | 

## Methods

### NewFreeAuthClientPolicyOpenApiVO

`func NewFreeAuthClientPolicyOpenApiVO(type_ int32, ) *FreeAuthClientPolicyOpenApiVO`

NewFreeAuthClientPolicyOpenApiVO instantiates a new FreeAuthClientPolicyOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreeAuthClientPolicyOpenApiVOWithDefaults

`func NewFreeAuthClientPolicyOpenApiVOWithDefaults() *FreeAuthClientPolicyOpenApiVO`

NewFreeAuthClientPolicyOpenApiVOWithDefaults instantiates a new FreeAuthClientPolicyOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIp

`func (o *FreeAuthClientPolicyOpenApiVO) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *FreeAuthClientPolicyOpenApiVO) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *FreeAuthClientPolicyOpenApiVO) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *FreeAuthClientPolicyOpenApiVO) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetClientMac

`func (o *FreeAuthClientPolicyOpenApiVO) GetClientMac() string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *FreeAuthClientPolicyOpenApiVO) GetClientMacOk() (*string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *FreeAuthClientPolicyOpenApiVO) SetClientMac(v string)`

SetClientMac sets ClientMac field to given value.

### HasClientMac

`func (o *FreeAuthClientPolicyOpenApiVO) HasClientMac() bool`

HasClientMac returns a boolean if a field has been set.

### GetIdInt

`func (o *FreeAuthClientPolicyOpenApiVO) GetIdInt() int32`

GetIdInt returns the IdInt field if non-nil, zero value otherwise.

### GetIdIntOk

`func (o *FreeAuthClientPolicyOpenApiVO) GetIdIntOk() (*int32, bool)`

GetIdIntOk returns a tuple with the IdInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdInt

`func (o *FreeAuthClientPolicyOpenApiVO) SetIdInt(v int32)`

SetIdInt sets IdInt field to given value.

### HasIdInt

`func (o *FreeAuthClientPolicyOpenApiVO) HasIdInt() bool`

HasIdInt returns a boolean if a field has been set.

### GetType

`func (o *FreeAuthClientPolicyOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FreeAuthClientPolicyOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FreeAuthClientPolicyOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


