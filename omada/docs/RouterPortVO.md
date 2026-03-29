# RouterPortVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | Pointer to [**[]DevicePortVO**](DevicePortVO.md) | The collection of related devices. | [optional] 
**ForbiddenRouterPortsNum** | Pointer to **int32** | The total num of forbidden router ports related to this network. | [optional] 
**NetworkId** | Pointer to **string** | The unique identification of one network. | [optional] 
**NetworkName** | Pointer to **string** | The name of network. | [optional] 
**StaticRouterPortsNum** | Pointer to **int32** | The total num of static router ports related to this network. | [optional] 
**Vlan** | Pointer to **int32** | The vlan of the network. | [optional] 

## Methods

### NewRouterPortVO

`func NewRouterPortVO() *RouterPortVO`

NewRouterPortVO instantiates a new RouterPortVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouterPortVOWithDefaults

`func NewRouterPortVOWithDefaults() *RouterPortVO`

NewRouterPortVOWithDefaults instantiates a new RouterPortVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *RouterPortVO) GetDevices() []DevicePortVO`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *RouterPortVO) GetDevicesOk() (*[]DevicePortVO, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *RouterPortVO) SetDevices(v []DevicePortVO)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *RouterPortVO) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetForbiddenRouterPortsNum

`func (o *RouterPortVO) GetForbiddenRouterPortsNum() int32`

GetForbiddenRouterPortsNum returns the ForbiddenRouterPortsNum field if non-nil, zero value otherwise.

### GetForbiddenRouterPortsNumOk

`func (o *RouterPortVO) GetForbiddenRouterPortsNumOk() (*int32, bool)`

GetForbiddenRouterPortsNumOk returns a tuple with the ForbiddenRouterPortsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenRouterPortsNum

`func (o *RouterPortVO) SetForbiddenRouterPortsNum(v int32)`

SetForbiddenRouterPortsNum sets ForbiddenRouterPortsNum field to given value.

### HasForbiddenRouterPortsNum

`func (o *RouterPortVO) HasForbiddenRouterPortsNum() bool`

HasForbiddenRouterPortsNum returns a boolean if a field has been set.

### GetNetworkId

`func (o *RouterPortVO) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *RouterPortVO) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *RouterPortVO) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *RouterPortVO) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkName

`func (o *RouterPortVO) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *RouterPortVO) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *RouterPortVO) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *RouterPortVO) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetStaticRouterPortsNum

`func (o *RouterPortVO) GetStaticRouterPortsNum() int32`

GetStaticRouterPortsNum returns the StaticRouterPortsNum field if non-nil, zero value otherwise.

### GetStaticRouterPortsNumOk

`func (o *RouterPortVO) GetStaticRouterPortsNumOk() (*int32, bool)`

GetStaticRouterPortsNumOk returns a tuple with the StaticRouterPortsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRouterPortsNum

`func (o *RouterPortVO) SetStaticRouterPortsNum(v int32)`

SetStaticRouterPortsNum sets StaticRouterPortsNum field to given value.

### HasStaticRouterPortsNum

`func (o *RouterPortVO) HasStaticRouterPortsNum() bool`

HasStaticRouterPortsNum returns a boolean if a field has been set.

### GetVlan

`func (o *RouterPortVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *RouterPortVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *RouterPortVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *RouterPortVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


