# TopologyClientUplinkSwitchPortInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagId** | Pointer to **int32** | Lag Id. | [optional] 
**LagPort** | Pointer to **[]int32** | Lag Ports. | [optional] 
**Port** | Pointer to **int32** | Real Port Number or Port Number of Lag. | [optional] 
**StandardPort** | Pointer to **string** | Standard Port of Stack, for Uplink Device is Stack. | [optional] 

## Methods

### NewTopologyClientUplinkSwitchPortInfo

`func NewTopologyClientUplinkSwitchPortInfo() *TopologyClientUplinkSwitchPortInfo`

NewTopologyClientUplinkSwitchPortInfo instantiates a new TopologyClientUplinkSwitchPortInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyClientUplinkSwitchPortInfoWithDefaults

`func NewTopologyClientUplinkSwitchPortInfoWithDefaults() *TopologyClientUplinkSwitchPortInfo`

NewTopologyClientUplinkSwitchPortInfoWithDefaults instantiates a new TopologyClientUplinkSwitchPortInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagId

`func (o *TopologyClientUplinkSwitchPortInfo) GetLagId() int32`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *TopologyClientUplinkSwitchPortInfo) GetLagIdOk() (*int32, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *TopologyClientUplinkSwitchPortInfo) SetLagId(v int32)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *TopologyClientUplinkSwitchPortInfo) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetLagPort

`func (o *TopologyClientUplinkSwitchPortInfo) GetLagPort() []int32`

GetLagPort returns the LagPort field if non-nil, zero value otherwise.

### GetLagPortOk

`func (o *TopologyClientUplinkSwitchPortInfo) GetLagPortOk() (*[]int32, bool)`

GetLagPortOk returns a tuple with the LagPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagPort

`func (o *TopologyClientUplinkSwitchPortInfo) SetLagPort(v []int32)`

SetLagPort sets LagPort field to given value.

### HasLagPort

`func (o *TopologyClientUplinkSwitchPortInfo) HasLagPort() bool`

HasLagPort returns a boolean if a field has been set.

### GetPort

`func (o *TopologyClientUplinkSwitchPortInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TopologyClientUplinkSwitchPortInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TopologyClientUplinkSwitchPortInfo) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *TopologyClientUplinkSwitchPortInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStandardPort

`func (o *TopologyClientUplinkSwitchPortInfo) GetStandardPort() string`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *TopologyClientUplinkSwitchPortInfo) GetStandardPortOk() (*string, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *TopologyClientUplinkSwitchPortInfo) SetStandardPort(v string)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *TopologyClientUplinkSwitchPortInfo) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


