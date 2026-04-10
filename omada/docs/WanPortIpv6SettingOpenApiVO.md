# WanPortIpv6SettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | IPv6 enable. | 
**Ipv6Dynamic** | Pointer to [**Ipv6DynamicOpenApiVO**](Ipv6DynamicOpenApiVO.md) |  | [optional] 
**Ipv6Pppoe** | Pointer to [**Ipv6PppoeOpenApiVO**](Ipv6PppoeOpenApiVO.md) |  | [optional] 
**Ipv6Static** | Pointer to [**Ipv6StaticOpenApiVO**](Ipv6StaticOpenApiVO.md) |  | [optional] 
**Ipv6Tunnel** | Pointer to [**Ipv6TunnelOpenApiVO**](Ipv6TunnelOpenApiVO.md) |  | [optional] 
**ProtoType** | Pointer to **int32** | IPv6 connection type should be a value as follows: 0: static; 1: dynamic; 2: PPPoE; 3: 6to4Tunnel; 4: bridge. | [optional] 

## Methods

### NewWanPortIpv6SettingOpenApiVO

`func NewWanPortIpv6SettingOpenApiVO(enable bool, ) *WanPortIpv6SettingOpenApiVO`

NewWanPortIpv6SettingOpenApiVO instantiates a new WanPortIpv6SettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWanPortIpv6SettingOpenApiVOWithDefaults

`func NewWanPortIpv6SettingOpenApiVOWithDefaults() *WanPortIpv6SettingOpenApiVO`

NewWanPortIpv6SettingOpenApiVOWithDefaults instantiates a new WanPortIpv6SettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *WanPortIpv6SettingOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *WanPortIpv6SettingOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *WanPortIpv6SettingOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetIpv6Dynamic

`func (o *WanPortIpv6SettingOpenApiVO) GetIpv6Dynamic() Ipv6DynamicOpenApiVO`

GetIpv6Dynamic returns the Ipv6Dynamic field if non-nil, zero value otherwise.

### GetIpv6DynamicOk

`func (o *WanPortIpv6SettingOpenApiVO) GetIpv6DynamicOk() (*Ipv6DynamicOpenApiVO, bool)`

GetIpv6DynamicOk returns a tuple with the Ipv6Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Dynamic

`func (o *WanPortIpv6SettingOpenApiVO) SetIpv6Dynamic(v Ipv6DynamicOpenApiVO)`

SetIpv6Dynamic sets Ipv6Dynamic field to given value.

### HasIpv6Dynamic

`func (o *WanPortIpv6SettingOpenApiVO) HasIpv6Dynamic() bool`

HasIpv6Dynamic returns a boolean if a field has been set.

### GetIpv6Pppoe

`func (o *WanPortIpv6SettingOpenApiVO) GetIpv6Pppoe() Ipv6PppoeOpenApiVO`

GetIpv6Pppoe returns the Ipv6Pppoe field if non-nil, zero value otherwise.

### GetIpv6PppoeOk

`func (o *WanPortIpv6SettingOpenApiVO) GetIpv6PppoeOk() (*Ipv6PppoeOpenApiVO, bool)`

GetIpv6PppoeOk returns a tuple with the Ipv6Pppoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Pppoe

`func (o *WanPortIpv6SettingOpenApiVO) SetIpv6Pppoe(v Ipv6PppoeOpenApiVO)`

SetIpv6Pppoe sets Ipv6Pppoe field to given value.

### HasIpv6Pppoe

`func (o *WanPortIpv6SettingOpenApiVO) HasIpv6Pppoe() bool`

HasIpv6Pppoe returns a boolean if a field has been set.

### GetIpv6Static

`func (o *WanPortIpv6SettingOpenApiVO) GetIpv6Static() Ipv6StaticOpenApiVO`

GetIpv6Static returns the Ipv6Static field if non-nil, zero value otherwise.

### GetIpv6StaticOk

`func (o *WanPortIpv6SettingOpenApiVO) GetIpv6StaticOk() (*Ipv6StaticOpenApiVO, bool)`

GetIpv6StaticOk returns a tuple with the Ipv6Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Static

`func (o *WanPortIpv6SettingOpenApiVO) SetIpv6Static(v Ipv6StaticOpenApiVO)`

SetIpv6Static sets Ipv6Static field to given value.

### HasIpv6Static

`func (o *WanPortIpv6SettingOpenApiVO) HasIpv6Static() bool`

HasIpv6Static returns a boolean if a field has been set.

### GetIpv6Tunnel

`func (o *WanPortIpv6SettingOpenApiVO) GetIpv6Tunnel() Ipv6TunnelOpenApiVO`

GetIpv6Tunnel returns the Ipv6Tunnel field if non-nil, zero value otherwise.

### GetIpv6TunnelOk

`func (o *WanPortIpv6SettingOpenApiVO) GetIpv6TunnelOk() (*Ipv6TunnelOpenApiVO, bool)`

GetIpv6TunnelOk returns a tuple with the Ipv6Tunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Tunnel

`func (o *WanPortIpv6SettingOpenApiVO) SetIpv6Tunnel(v Ipv6TunnelOpenApiVO)`

SetIpv6Tunnel sets Ipv6Tunnel field to given value.

### HasIpv6Tunnel

`func (o *WanPortIpv6SettingOpenApiVO) HasIpv6Tunnel() bool`

HasIpv6Tunnel returns a boolean if a field has been set.

### GetProtoType

`func (o *WanPortIpv6SettingOpenApiVO) GetProtoType() int32`

GetProtoType returns the ProtoType field if non-nil, zero value otherwise.

### GetProtoTypeOk

`func (o *WanPortIpv6SettingOpenApiVO) GetProtoTypeOk() (*int32, bool)`

GetProtoTypeOk returns a tuple with the ProtoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoType

`func (o *WanPortIpv6SettingOpenApiVO) SetProtoType(v int32)`

SetProtoType sets ProtoType field to given value.

### HasProtoType

`func (o *WanPortIpv6SettingOpenApiVO) HasProtoType() bool`

HasProtoType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


