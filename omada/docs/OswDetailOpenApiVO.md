# OswDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Macs** | Pointer to **[]string** | The collection of macs of the general switches. | [optional] 
**NetworkId** | **string** | Network ID. | 
**StackIds** | Pointer to **[]string** | The collection of stack ids of the stack switches. | [optional] 
**Vlan** | **int32** | VLAN ID. | 

## Methods

### NewOswDetailOpenApiVO

`func NewOswDetailOpenApiVO(networkId string, vlan int32, ) *OswDetailOpenApiVO`

NewOswDetailOpenApiVO instantiates a new OswDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswDetailOpenApiVOWithDefaults

`func NewOswDetailOpenApiVOWithDefaults() *OswDetailOpenApiVO`

NewOswDetailOpenApiVOWithDefaults instantiates a new OswDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacs

`func (o *OswDetailOpenApiVO) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *OswDetailOpenApiVO) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *OswDetailOpenApiVO) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *OswDetailOpenApiVO) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### GetNetworkId

`func (o *OswDetailOpenApiVO) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *OswDetailOpenApiVO) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *OswDetailOpenApiVO) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetStackIds

`func (o *OswDetailOpenApiVO) GetStackIds() []string`

GetStackIds returns the StackIds field if non-nil, zero value otherwise.

### GetStackIdsOk

`func (o *OswDetailOpenApiVO) GetStackIdsOk() (*[]string, bool)`

GetStackIdsOk returns a tuple with the StackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackIds

`func (o *OswDetailOpenApiVO) SetStackIds(v []string)`

SetStackIds sets StackIds field to given value.

### HasStackIds

`func (o *OswDetailOpenApiVO) HasStackIds() bool`

HasStackIds returns a boolean if a field has been set.

### GetVlan

`func (o *OswDetailOpenApiVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *OswDetailOpenApiVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *OswDetailOpenApiVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


