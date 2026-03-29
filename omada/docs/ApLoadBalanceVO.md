# ApLoadBalanceVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LbEnable** | Pointer to **bool** |  | [optional] 
**MaxClients** | Pointer to **int32** |  | [optional] 

## Methods

### NewApLoadBalanceVO

`func NewApLoadBalanceVO() *ApLoadBalanceVO`

NewApLoadBalanceVO instantiates a new ApLoadBalanceVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApLoadBalanceVOWithDefaults

`func NewApLoadBalanceVOWithDefaults() *ApLoadBalanceVO`

NewApLoadBalanceVOWithDefaults instantiates a new ApLoadBalanceVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLbEnable

`func (o *ApLoadBalanceVO) GetLbEnable() bool`

GetLbEnable returns the LbEnable field if non-nil, zero value otherwise.

### GetLbEnableOk

`func (o *ApLoadBalanceVO) GetLbEnableOk() (*bool, bool)`

GetLbEnableOk returns a tuple with the LbEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbEnable

`func (o *ApLoadBalanceVO) SetLbEnable(v bool)`

SetLbEnable sets LbEnable field to given value.

### HasLbEnable

`func (o *ApLoadBalanceVO) HasLbEnable() bool`

HasLbEnable returns a boolean if a field has been set.

### GetMaxClients

`func (o *ApLoadBalanceVO) GetMaxClients() int32`

GetMaxClients returns the MaxClients field if non-nil, zero value otherwise.

### GetMaxClientsOk

`func (o *ApLoadBalanceVO) GetMaxClientsOk() (*int32, bool)`

GetMaxClientsOk returns a tuple with the MaxClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClients

`func (o *ApLoadBalanceVO) SetMaxClients(v int32)`

SetMaxClients sets MaxClients field to given value.

### HasMaxClients

`func (o *ApLoadBalanceVO) HasMaxClients() bool`

HasMaxClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


