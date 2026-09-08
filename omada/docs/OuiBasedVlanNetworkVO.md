# OuiBasedVlanNetworkVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | The id for the network. | [optional] 
**NetworkName** | Pointer to **string** | The name of the network. | [optional] 
**OuiBasedVlanId** | Pointer to **string** | The id of the oui based vlan rule. | [optional] 
**OuiBasedVlanName** | Pointer to **string** | The name of the oui based vlan rule. | [optional] 
**RuleMode** | Pointer to **int32** | The mode of the oui based vlan rule. 0 represents all device ports ,1 represents custom. | [optional] 
**Vlan** | Pointer to **int32** | The vlan of one network that multicast config used. | [optional] 
**VlanType** | Pointer to **int32** | 0:Single, 1:Multiple(Bridge Vlan) | [optional] 

## Methods

### NewOuiBasedVlanNetworkVO

`func NewOuiBasedVlanNetworkVO() *OuiBasedVlanNetworkVO`

NewOuiBasedVlanNetworkVO instantiates a new OuiBasedVlanNetworkVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOuiBasedVlanNetworkVOWithDefaults

`func NewOuiBasedVlanNetworkVOWithDefaults() *OuiBasedVlanNetworkVO`

NewOuiBasedVlanNetworkVOWithDefaults instantiates a new OuiBasedVlanNetworkVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *OuiBasedVlanNetworkVO) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *OuiBasedVlanNetworkVO) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *OuiBasedVlanNetworkVO) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *OuiBasedVlanNetworkVO) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkName

`func (o *OuiBasedVlanNetworkVO) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *OuiBasedVlanNetworkVO) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *OuiBasedVlanNetworkVO) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *OuiBasedVlanNetworkVO) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetOuiBasedVlanId

`func (o *OuiBasedVlanNetworkVO) GetOuiBasedVlanId() string`

GetOuiBasedVlanId returns the OuiBasedVlanId field if non-nil, zero value otherwise.

### GetOuiBasedVlanIdOk

`func (o *OuiBasedVlanNetworkVO) GetOuiBasedVlanIdOk() (*string, bool)`

GetOuiBasedVlanIdOk returns a tuple with the OuiBasedVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuiBasedVlanId

`func (o *OuiBasedVlanNetworkVO) SetOuiBasedVlanId(v string)`

SetOuiBasedVlanId sets OuiBasedVlanId field to given value.

### HasOuiBasedVlanId

`func (o *OuiBasedVlanNetworkVO) HasOuiBasedVlanId() bool`

HasOuiBasedVlanId returns a boolean if a field has been set.

### GetOuiBasedVlanName

`func (o *OuiBasedVlanNetworkVO) GetOuiBasedVlanName() string`

GetOuiBasedVlanName returns the OuiBasedVlanName field if non-nil, zero value otherwise.

### GetOuiBasedVlanNameOk

`func (o *OuiBasedVlanNetworkVO) GetOuiBasedVlanNameOk() (*string, bool)`

GetOuiBasedVlanNameOk returns a tuple with the OuiBasedVlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuiBasedVlanName

`func (o *OuiBasedVlanNetworkVO) SetOuiBasedVlanName(v string)`

SetOuiBasedVlanName sets OuiBasedVlanName field to given value.

### HasOuiBasedVlanName

`func (o *OuiBasedVlanNetworkVO) HasOuiBasedVlanName() bool`

HasOuiBasedVlanName returns a boolean if a field has been set.

### GetRuleMode

`func (o *OuiBasedVlanNetworkVO) GetRuleMode() int32`

GetRuleMode returns the RuleMode field if non-nil, zero value otherwise.

### GetRuleModeOk

`func (o *OuiBasedVlanNetworkVO) GetRuleModeOk() (*int32, bool)`

GetRuleModeOk returns a tuple with the RuleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleMode

`func (o *OuiBasedVlanNetworkVO) SetRuleMode(v int32)`

SetRuleMode sets RuleMode field to given value.

### HasRuleMode

`func (o *OuiBasedVlanNetworkVO) HasRuleMode() bool`

HasRuleMode returns a boolean if a field has been set.

### GetVlan

`func (o *OuiBasedVlanNetworkVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *OuiBasedVlanNetworkVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *OuiBasedVlanNetworkVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *OuiBasedVlanNetworkVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVlanType

`func (o *OuiBasedVlanNetworkVO) GetVlanType() int32`

GetVlanType returns the VlanType field if non-nil, zero value otherwise.

### GetVlanTypeOk

`func (o *OuiBasedVlanNetworkVO) GetVlanTypeOk() (*int32, bool)`

GetVlanTypeOk returns a tuple with the VlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanType

`func (o *OuiBasedVlanNetworkVO) SetVlanType(v int32)`

SetVlanType sets VlanType field to given value.

### HasVlanType

`func (o *OuiBasedVlanNetworkVO) HasVlanType() bool`

HasVlanType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


