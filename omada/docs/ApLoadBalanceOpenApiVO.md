# ApLoadBalanceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LbEnable** | Pointer to **bool** | Load Balance enabled or not. | [optional] 
**MaxClients** | Pointer to **int32** | Maximum associated clients. | [optional] 

## Methods

### NewApLoadBalanceOpenApiVO

`func NewApLoadBalanceOpenApiVO() *ApLoadBalanceOpenApiVO`

NewApLoadBalanceOpenApiVO instantiates a new ApLoadBalanceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApLoadBalanceOpenApiVOWithDefaults

`func NewApLoadBalanceOpenApiVOWithDefaults() *ApLoadBalanceOpenApiVO`

NewApLoadBalanceOpenApiVOWithDefaults instantiates a new ApLoadBalanceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLbEnable

`func (o *ApLoadBalanceOpenApiVO) GetLbEnable() bool`

GetLbEnable returns the LbEnable field if non-nil, zero value otherwise.

### GetLbEnableOk

`func (o *ApLoadBalanceOpenApiVO) GetLbEnableOk() (*bool, bool)`

GetLbEnableOk returns a tuple with the LbEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbEnable

`func (o *ApLoadBalanceOpenApiVO) SetLbEnable(v bool)`

SetLbEnable sets LbEnable field to given value.

### HasLbEnable

`func (o *ApLoadBalanceOpenApiVO) HasLbEnable() bool`

HasLbEnable returns a boolean if a field has been set.

### GetMaxClients

`func (o *ApLoadBalanceOpenApiVO) GetMaxClients() int32`

GetMaxClients returns the MaxClients field if non-nil, zero value otherwise.

### GetMaxClientsOk

`func (o *ApLoadBalanceOpenApiVO) GetMaxClientsOk() (*int32, bool)`

GetMaxClientsOk returns a tuple with the MaxClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClients

`func (o *ApLoadBalanceOpenApiVO) SetMaxClients(v int32)`

SetMaxClients sets MaxClients field to given value.

### HasMaxClients

`func (o *ApLoadBalanceOpenApiVO) HasMaxClients() bool`

HasMaxClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


