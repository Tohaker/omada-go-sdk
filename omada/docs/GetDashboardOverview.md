# GetDashboardOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailablePorts** | Pointer to **int32** | Number of valid ports | [optional] 
**ConnectedApNum** | Pointer to **int32** | Number of APs in connected state | [optional] 
**ConnectedGatewayNum** | Pointer to **int32** | The number of connected gateways | [optional] 
**ConnectedSwitchNum** | Pointer to **int32** | Number of switches in connected state  | [optional] 
**DisconnectedApNum** | Pointer to **int32** | Number of APs in disconnected state | [optional] 
**DisconnectedGatewayNum** | Pointer to **int32** | The number of disconnected gateways | [optional] 
**DisconnectedSwitchNum** | Pointer to **int32** | Number of switches in disconnected state | [optional] 
**GuestNum** | Pointer to **int32** | Number of wireless guest clients | [optional] 
**IsolatedApNum** | Pointer to **int32** | Number of APs in isolated state | [optional] 
**NetCapacity** | Pointer to **int32** | Network Capacity measured in Mbps. Null means no value. | [optional] 
**NetUsage** | Pointer to **float64** | The current link usage rate under the Gateway node (percent conversion has been performed, such as 98.1) | [optional] 
**PowerConsumption** | Pointer to **float64** | Power consumption (unit: W) | [optional] 
**TotalApNum** | Pointer to **int32** | Total number of APs | [optional] 
**TotalClientNum** | Pointer to **int32** | Total number of clients | [optional] 
**TotalGatewayNum** | Pointer to **int32** | The total number of gateways | [optional] 
**TotalPorts** | Pointer to **int32** | Total number of ports | [optional] 
**TotalSwitchNum** | Pointer to **int32** | Total number of switches | [optional] 
**WiredClientNum** | Pointer to **int32** | Number of wired clients | [optional] 
**WirelessClientNum** | Pointer to **int32** | Number of wireless clients | [optional] 

## Methods

### NewGetDashboardOverview

`func NewGetDashboardOverview() *GetDashboardOverview`

NewGetDashboardOverview instantiates a new GetDashboardOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDashboardOverviewWithDefaults

`func NewGetDashboardOverviewWithDefaults() *GetDashboardOverview`

NewGetDashboardOverviewWithDefaults instantiates a new GetDashboardOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailablePorts

`func (o *GetDashboardOverview) GetAvailablePorts() int32`

GetAvailablePorts returns the AvailablePorts field if non-nil, zero value otherwise.

### GetAvailablePortsOk

`func (o *GetDashboardOverview) GetAvailablePortsOk() (*int32, bool)`

GetAvailablePortsOk returns a tuple with the AvailablePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePorts

`func (o *GetDashboardOverview) SetAvailablePorts(v int32)`

SetAvailablePorts sets AvailablePorts field to given value.

### HasAvailablePorts

`func (o *GetDashboardOverview) HasAvailablePorts() bool`

HasAvailablePorts returns a boolean if a field has been set.

### GetConnectedApNum

`func (o *GetDashboardOverview) GetConnectedApNum() int32`

GetConnectedApNum returns the ConnectedApNum field if non-nil, zero value otherwise.

### GetConnectedApNumOk

`func (o *GetDashboardOverview) GetConnectedApNumOk() (*int32, bool)`

GetConnectedApNumOk returns a tuple with the ConnectedApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedApNum

`func (o *GetDashboardOverview) SetConnectedApNum(v int32)`

SetConnectedApNum sets ConnectedApNum field to given value.

### HasConnectedApNum

`func (o *GetDashboardOverview) HasConnectedApNum() bool`

HasConnectedApNum returns a boolean if a field has been set.

### GetConnectedGatewayNum

`func (o *GetDashboardOverview) GetConnectedGatewayNum() int32`

GetConnectedGatewayNum returns the ConnectedGatewayNum field if non-nil, zero value otherwise.

### GetConnectedGatewayNumOk

`func (o *GetDashboardOverview) GetConnectedGatewayNumOk() (*int32, bool)`

GetConnectedGatewayNumOk returns a tuple with the ConnectedGatewayNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedGatewayNum

`func (o *GetDashboardOverview) SetConnectedGatewayNum(v int32)`

SetConnectedGatewayNum sets ConnectedGatewayNum field to given value.

### HasConnectedGatewayNum

`func (o *GetDashboardOverview) HasConnectedGatewayNum() bool`

HasConnectedGatewayNum returns a boolean if a field has been set.

### GetConnectedSwitchNum

`func (o *GetDashboardOverview) GetConnectedSwitchNum() int32`

GetConnectedSwitchNum returns the ConnectedSwitchNum field if non-nil, zero value otherwise.

### GetConnectedSwitchNumOk

`func (o *GetDashboardOverview) GetConnectedSwitchNumOk() (*int32, bool)`

GetConnectedSwitchNumOk returns a tuple with the ConnectedSwitchNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedSwitchNum

`func (o *GetDashboardOverview) SetConnectedSwitchNum(v int32)`

SetConnectedSwitchNum sets ConnectedSwitchNum field to given value.

### HasConnectedSwitchNum

`func (o *GetDashboardOverview) HasConnectedSwitchNum() bool`

HasConnectedSwitchNum returns a boolean if a field has been set.

### GetDisconnectedApNum

`func (o *GetDashboardOverview) GetDisconnectedApNum() int32`

GetDisconnectedApNum returns the DisconnectedApNum field if non-nil, zero value otherwise.

### GetDisconnectedApNumOk

`func (o *GetDashboardOverview) GetDisconnectedApNumOk() (*int32, bool)`

GetDisconnectedApNumOk returns a tuple with the DisconnectedApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedApNum

`func (o *GetDashboardOverview) SetDisconnectedApNum(v int32)`

SetDisconnectedApNum sets DisconnectedApNum field to given value.

### HasDisconnectedApNum

`func (o *GetDashboardOverview) HasDisconnectedApNum() bool`

HasDisconnectedApNum returns a boolean if a field has been set.

### GetDisconnectedGatewayNum

`func (o *GetDashboardOverview) GetDisconnectedGatewayNum() int32`

GetDisconnectedGatewayNum returns the DisconnectedGatewayNum field if non-nil, zero value otherwise.

### GetDisconnectedGatewayNumOk

`func (o *GetDashboardOverview) GetDisconnectedGatewayNumOk() (*int32, bool)`

GetDisconnectedGatewayNumOk returns a tuple with the DisconnectedGatewayNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedGatewayNum

`func (o *GetDashboardOverview) SetDisconnectedGatewayNum(v int32)`

SetDisconnectedGatewayNum sets DisconnectedGatewayNum field to given value.

### HasDisconnectedGatewayNum

`func (o *GetDashboardOverview) HasDisconnectedGatewayNum() bool`

HasDisconnectedGatewayNum returns a boolean if a field has been set.

### GetDisconnectedSwitchNum

`func (o *GetDashboardOverview) GetDisconnectedSwitchNum() int32`

GetDisconnectedSwitchNum returns the DisconnectedSwitchNum field if non-nil, zero value otherwise.

### GetDisconnectedSwitchNumOk

`func (o *GetDashboardOverview) GetDisconnectedSwitchNumOk() (*int32, bool)`

GetDisconnectedSwitchNumOk returns a tuple with the DisconnectedSwitchNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedSwitchNum

`func (o *GetDashboardOverview) SetDisconnectedSwitchNum(v int32)`

SetDisconnectedSwitchNum sets DisconnectedSwitchNum field to given value.

### HasDisconnectedSwitchNum

`func (o *GetDashboardOverview) HasDisconnectedSwitchNum() bool`

HasDisconnectedSwitchNum returns a boolean if a field has been set.

### GetGuestNum

`func (o *GetDashboardOverview) GetGuestNum() int32`

GetGuestNum returns the GuestNum field if non-nil, zero value otherwise.

### GetGuestNumOk

`func (o *GetDashboardOverview) GetGuestNumOk() (*int32, bool)`

GetGuestNumOk returns a tuple with the GuestNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestNum

`func (o *GetDashboardOverview) SetGuestNum(v int32)`

SetGuestNum sets GuestNum field to given value.

### HasGuestNum

`func (o *GetDashboardOverview) HasGuestNum() bool`

HasGuestNum returns a boolean if a field has been set.

### GetIsolatedApNum

`func (o *GetDashboardOverview) GetIsolatedApNum() int32`

GetIsolatedApNum returns the IsolatedApNum field if non-nil, zero value otherwise.

### GetIsolatedApNumOk

`func (o *GetDashboardOverview) GetIsolatedApNumOk() (*int32, bool)`

GetIsolatedApNumOk returns a tuple with the IsolatedApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolatedApNum

`func (o *GetDashboardOverview) SetIsolatedApNum(v int32)`

SetIsolatedApNum sets IsolatedApNum field to given value.

### HasIsolatedApNum

`func (o *GetDashboardOverview) HasIsolatedApNum() bool`

HasIsolatedApNum returns a boolean if a field has been set.

### GetNetCapacity

`func (o *GetDashboardOverview) GetNetCapacity() int32`

GetNetCapacity returns the NetCapacity field if non-nil, zero value otherwise.

### GetNetCapacityOk

`func (o *GetDashboardOverview) GetNetCapacityOk() (*int32, bool)`

GetNetCapacityOk returns a tuple with the NetCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetCapacity

`func (o *GetDashboardOverview) SetNetCapacity(v int32)`

SetNetCapacity sets NetCapacity field to given value.

### HasNetCapacity

`func (o *GetDashboardOverview) HasNetCapacity() bool`

HasNetCapacity returns a boolean if a field has been set.

### GetNetUsage

`func (o *GetDashboardOverview) GetNetUsage() float64`

GetNetUsage returns the NetUsage field if non-nil, zero value otherwise.

### GetNetUsageOk

`func (o *GetDashboardOverview) GetNetUsageOk() (*float64, bool)`

GetNetUsageOk returns a tuple with the NetUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetUsage

`func (o *GetDashboardOverview) SetNetUsage(v float64)`

SetNetUsage sets NetUsage field to given value.

### HasNetUsage

`func (o *GetDashboardOverview) HasNetUsage() bool`

HasNetUsage returns a boolean if a field has been set.

### GetPowerConsumption

`func (o *GetDashboardOverview) GetPowerConsumption() float64`

GetPowerConsumption returns the PowerConsumption field if non-nil, zero value otherwise.

### GetPowerConsumptionOk

`func (o *GetDashboardOverview) GetPowerConsumptionOk() (*float64, bool)`

GetPowerConsumptionOk returns a tuple with the PowerConsumption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerConsumption

`func (o *GetDashboardOverview) SetPowerConsumption(v float64)`

SetPowerConsumption sets PowerConsumption field to given value.

### HasPowerConsumption

`func (o *GetDashboardOverview) HasPowerConsumption() bool`

HasPowerConsumption returns a boolean if a field has been set.

### GetTotalApNum

`func (o *GetDashboardOverview) GetTotalApNum() int32`

GetTotalApNum returns the TotalApNum field if non-nil, zero value otherwise.

### GetTotalApNumOk

`func (o *GetDashboardOverview) GetTotalApNumOk() (*int32, bool)`

GetTotalApNumOk returns a tuple with the TotalApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApNum

`func (o *GetDashboardOverview) SetTotalApNum(v int32)`

SetTotalApNum sets TotalApNum field to given value.

### HasTotalApNum

`func (o *GetDashboardOverview) HasTotalApNum() bool`

HasTotalApNum returns a boolean if a field has been set.

### GetTotalClientNum

`func (o *GetDashboardOverview) GetTotalClientNum() int32`

GetTotalClientNum returns the TotalClientNum field if non-nil, zero value otherwise.

### GetTotalClientNumOk

`func (o *GetDashboardOverview) GetTotalClientNumOk() (*int32, bool)`

GetTotalClientNumOk returns a tuple with the TotalClientNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalClientNum

`func (o *GetDashboardOverview) SetTotalClientNum(v int32)`

SetTotalClientNum sets TotalClientNum field to given value.

### HasTotalClientNum

`func (o *GetDashboardOverview) HasTotalClientNum() bool`

HasTotalClientNum returns a boolean if a field has been set.

### GetTotalGatewayNum

`func (o *GetDashboardOverview) GetTotalGatewayNum() int32`

GetTotalGatewayNum returns the TotalGatewayNum field if non-nil, zero value otherwise.

### GetTotalGatewayNumOk

`func (o *GetDashboardOverview) GetTotalGatewayNumOk() (*int32, bool)`

GetTotalGatewayNumOk returns a tuple with the TotalGatewayNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGatewayNum

`func (o *GetDashboardOverview) SetTotalGatewayNum(v int32)`

SetTotalGatewayNum sets TotalGatewayNum field to given value.

### HasTotalGatewayNum

`func (o *GetDashboardOverview) HasTotalGatewayNum() bool`

HasTotalGatewayNum returns a boolean if a field has been set.

### GetTotalPorts

`func (o *GetDashboardOverview) GetTotalPorts() int32`

GetTotalPorts returns the TotalPorts field if non-nil, zero value otherwise.

### GetTotalPortsOk

`func (o *GetDashboardOverview) GetTotalPortsOk() (*int32, bool)`

GetTotalPortsOk returns a tuple with the TotalPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPorts

`func (o *GetDashboardOverview) SetTotalPorts(v int32)`

SetTotalPorts sets TotalPorts field to given value.

### HasTotalPorts

`func (o *GetDashboardOverview) HasTotalPorts() bool`

HasTotalPorts returns a boolean if a field has been set.

### GetTotalSwitchNum

`func (o *GetDashboardOverview) GetTotalSwitchNum() int32`

GetTotalSwitchNum returns the TotalSwitchNum field if non-nil, zero value otherwise.

### GetTotalSwitchNumOk

`func (o *GetDashboardOverview) GetTotalSwitchNumOk() (*int32, bool)`

GetTotalSwitchNumOk returns a tuple with the TotalSwitchNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSwitchNum

`func (o *GetDashboardOverview) SetTotalSwitchNum(v int32)`

SetTotalSwitchNum sets TotalSwitchNum field to given value.

### HasTotalSwitchNum

`func (o *GetDashboardOverview) HasTotalSwitchNum() bool`

HasTotalSwitchNum returns a boolean if a field has been set.

### GetWiredClientNum

`func (o *GetDashboardOverview) GetWiredClientNum() int32`

GetWiredClientNum returns the WiredClientNum field if non-nil, zero value otherwise.

### GetWiredClientNumOk

`func (o *GetDashboardOverview) GetWiredClientNumOk() (*int32, bool)`

GetWiredClientNumOk returns a tuple with the WiredClientNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredClientNum

`func (o *GetDashboardOverview) SetWiredClientNum(v int32)`

SetWiredClientNum sets WiredClientNum field to given value.

### HasWiredClientNum

`func (o *GetDashboardOverview) HasWiredClientNum() bool`

HasWiredClientNum returns a boolean if a field has been set.

### GetWirelessClientNum

`func (o *GetDashboardOverview) GetWirelessClientNum() int32`

GetWirelessClientNum returns the WirelessClientNum field if non-nil, zero value otherwise.

### GetWirelessClientNumOk

`func (o *GetDashboardOverview) GetWirelessClientNumOk() (*int32, bool)`

GetWirelessClientNumOk returns a tuple with the WirelessClientNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessClientNum

`func (o *GetDashboardOverview) SetWirelessClientNum(v int32)`

SetWirelessClientNum sets WirelessClientNum field to given value.

### HasWirelessClientNum

`func (o *GetDashboardOverview) HasWirelessClientNum() bool`

HasWirelessClientNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


