# GatewayDirectionEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LanToLan** | Pointer to **bool** | Whether select LAN-&gt;LAN direction, which conflicts with other directions | [optional] 
**LanToWan** | Pointer to **bool** | Whether select LAN-&gt;WAN direction | [optional] 
**VpnInIds** | Pointer to **[]string** | Selected VPN IDs | [optional] 
**WanInIds** | Pointer to **[]string** | Selected WAN port IDs | [optional] 

## Methods

### NewGatewayDirectionEntity

`func NewGatewayDirectionEntity() *GatewayDirectionEntity`

NewGatewayDirectionEntity instantiates a new GatewayDirectionEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayDirectionEntityWithDefaults

`func NewGatewayDirectionEntityWithDefaults() *GatewayDirectionEntity`

NewGatewayDirectionEntityWithDefaults instantiates a new GatewayDirectionEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanToLan

`func (o *GatewayDirectionEntity) GetLanToLan() bool`

GetLanToLan returns the LanToLan field if non-nil, zero value otherwise.

### GetLanToLanOk

`func (o *GatewayDirectionEntity) GetLanToLanOk() (*bool, bool)`

GetLanToLanOk returns a tuple with the LanToLan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanToLan

`func (o *GatewayDirectionEntity) SetLanToLan(v bool)`

SetLanToLan sets LanToLan field to given value.

### HasLanToLan

`func (o *GatewayDirectionEntity) HasLanToLan() bool`

HasLanToLan returns a boolean if a field has been set.

### GetLanToWan

`func (o *GatewayDirectionEntity) GetLanToWan() bool`

GetLanToWan returns the LanToWan field if non-nil, zero value otherwise.

### GetLanToWanOk

`func (o *GatewayDirectionEntity) GetLanToWanOk() (*bool, bool)`

GetLanToWanOk returns a tuple with the LanToWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanToWan

`func (o *GatewayDirectionEntity) SetLanToWan(v bool)`

SetLanToWan sets LanToWan field to given value.

### HasLanToWan

`func (o *GatewayDirectionEntity) HasLanToWan() bool`

HasLanToWan returns a boolean if a field has been set.

### GetVpnInIds

`func (o *GatewayDirectionEntity) GetVpnInIds() []string`

GetVpnInIds returns the VpnInIds field if non-nil, zero value otherwise.

### GetVpnInIdsOk

`func (o *GatewayDirectionEntity) GetVpnInIdsOk() (*[]string, bool)`

GetVpnInIdsOk returns a tuple with the VpnInIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnInIds

`func (o *GatewayDirectionEntity) SetVpnInIds(v []string)`

SetVpnInIds sets VpnInIds field to given value.

### HasVpnInIds

`func (o *GatewayDirectionEntity) HasVpnInIds() bool`

HasVpnInIds returns a boolean if a field has been set.

### GetWanInIds

`func (o *GatewayDirectionEntity) GetWanInIds() []string`

GetWanInIds returns the WanInIds field if non-nil, zero value otherwise.

### GetWanInIdsOk

`func (o *GatewayDirectionEntity) GetWanInIdsOk() (*[]string, bool)`

GetWanInIdsOk returns a tuple with the WanInIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanInIds

`func (o *GatewayDirectionEntity) SetWanInIds(v []string)`

SetWanInIds sets WanInIds field to given value.

### HasWanInIds

`func (o *GatewayDirectionEntity) HasWanInIds() bool`

HasWanInIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


