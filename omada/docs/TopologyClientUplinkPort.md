# TopologyClientUplinkPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagId** | Pointer to **int32** | Lag Id. | [optional] 
**LagPorts** | Pointer to **[]int32** | Lag Ports. | [optional] 
**Name** | Pointer to **string** | Name of the Port. | [optional] 
**PoePower** | Pointer to **float64** | Poe power, unit: W. | [optional] 
**PoePowerDecimal** | Pointer to **float64** | Poe Power Decimal, unit: W. | [optional] 
**Port** | Pointer to **int32** | Real Port Number or Port Number of Lag. | [optional] 
**StandardLagPorts** | Pointer to **[]string** | Standard Lag Ports. | [optional] 
**StandardPort** | Pointer to **string** | Standard Port of Stack, for Uplink Device is Stack. | [optional] 

## Methods

### NewTopologyClientUplinkPort

`func NewTopologyClientUplinkPort() *TopologyClientUplinkPort`

NewTopologyClientUplinkPort instantiates a new TopologyClientUplinkPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyClientUplinkPortWithDefaults

`func NewTopologyClientUplinkPortWithDefaults() *TopologyClientUplinkPort`

NewTopologyClientUplinkPortWithDefaults instantiates a new TopologyClientUplinkPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagId

`func (o *TopologyClientUplinkPort) GetLagId() int32`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *TopologyClientUplinkPort) GetLagIdOk() (*int32, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *TopologyClientUplinkPort) SetLagId(v int32)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *TopologyClientUplinkPort) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetLagPorts

`func (o *TopologyClientUplinkPort) GetLagPorts() []int32`

GetLagPorts returns the LagPorts field if non-nil, zero value otherwise.

### GetLagPortsOk

`func (o *TopologyClientUplinkPort) GetLagPortsOk() (*[]int32, bool)`

GetLagPortsOk returns a tuple with the LagPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagPorts

`func (o *TopologyClientUplinkPort) SetLagPorts(v []int32)`

SetLagPorts sets LagPorts field to given value.

### HasLagPorts

`func (o *TopologyClientUplinkPort) HasLagPorts() bool`

HasLagPorts returns a boolean if a field has been set.

### GetName

`func (o *TopologyClientUplinkPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopologyClientUplinkPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopologyClientUplinkPort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopologyClientUplinkPort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPoePower

`func (o *TopologyClientUplinkPort) GetPoePower() float64`

GetPoePower returns the PoePower field if non-nil, zero value otherwise.

### GetPoePowerOk

`func (o *TopologyClientUplinkPort) GetPoePowerOk() (*float64, bool)`

GetPoePowerOk returns a tuple with the PoePower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePower

`func (o *TopologyClientUplinkPort) SetPoePower(v float64)`

SetPoePower sets PoePower field to given value.

### HasPoePower

`func (o *TopologyClientUplinkPort) HasPoePower() bool`

HasPoePower returns a boolean if a field has been set.

### GetPoePowerDecimal

`func (o *TopologyClientUplinkPort) GetPoePowerDecimal() float64`

GetPoePowerDecimal returns the PoePowerDecimal field if non-nil, zero value otherwise.

### GetPoePowerDecimalOk

`func (o *TopologyClientUplinkPort) GetPoePowerDecimalOk() (*float64, bool)`

GetPoePowerDecimalOk returns a tuple with the PoePowerDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePowerDecimal

`func (o *TopologyClientUplinkPort) SetPoePowerDecimal(v float64)`

SetPoePowerDecimal sets PoePowerDecimal field to given value.

### HasPoePowerDecimal

`func (o *TopologyClientUplinkPort) HasPoePowerDecimal() bool`

HasPoePowerDecimal returns a boolean if a field has been set.

### GetPort

`func (o *TopologyClientUplinkPort) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TopologyClientUplinkPort) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TopologyClientUplinkPort) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *TopologyClientUplinkPort) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStandardLagPorts

`func (o *TopologyClientUplinkPort) GetStandardLagPorts() []string`

GetStandardLagPorts returns the StandardLagPorts field if non-nil, zero value otherwise.

### GetStandardLagPortsOk

`func (o *TopologyClientUplinkPort) GetStandardLagPortsOk() (*[]string, bool)`

GetStandardLagPortsOk returns a tuple with the StandardLagPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardLagPorts

`func (o *TopologyClientUplinkPort) SetStandardLagPorts(v []string)`

SetStandardLagPorts sets StandardLagPorts field to given value.

### HasStandardLagPorts

`func (o *TopologyClientUplinkPort) HasStandardLagPorts() bool`

HasStandardLagPorts returns a boolean if a field has been set.

### GetStandardPort

`func (o *TopologyClientUplinkPort) GetStandardPort() string`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *TopologyClientUplinkPort) GetStandardPortOk() (*string, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *TopologyClientUplinkPort) SetStandardPort(v string)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *TopologyClientUplinkPort) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


