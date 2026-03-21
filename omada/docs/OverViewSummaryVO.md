# OverViewSummaryVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientHealthScore** | Pointer to **int32** | client health score | [optional] 
**DeviceHealthScore** | Pointer to **int32** | device health score | [optional] 
**SiteHealthScore** | Pointer to **int32** | site health score | [optional] 
**TotalAps** | Pointer to **int64** | total number of ap | [optional] 
**TotalClients** | Pointer to **int64** | total number of client | [optional] 
**TotalDevices** | Pointer to **int64** | total number of device | [optional] 
**TotalGateways** | Pointer to **int64** | total number of gateway | [optional] 
**TotalIpcs** | Pointer to **int64** | total number of ipc | [optional] 
**TotalNvrs** | Pointer to **int64** | total number of nvr | [optional] 
**TotalOlts** | Pointer to **int64** | total number of olt | [optional] 
**TotalSwitches** | Pointer to **int64** | total number of switch | [optional] 
**TotalTraffic** | Pointer to **int64** | total traffic | [optional] 
**WanHealthScore** | Pointer to **int32** | wan health score | [optional] 
**WifiHealthScore** | Pointer to **int32** | Wi-Fi health score | [optional] 
**WiredClients** | Pointer to **int64** | total number of wired client | [optional] 
**WiredTraffic** | Pointer to **int64** | total traffic of wired device | [optional] 
**WirelessClients** | Pointer to **int64** | total number of wireless client | [optional] 
**WirelessTraffic** | Pointer to **int64** | total traffic of wireless device | [optional] 

## Methods

### NewOverViewSummaryVO

`func NewOverViewSummaryVO() *OverViewSummaryVO`

NewOverViewSummaryVO instantiates a new OverViewSummaryVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverViewSummaryVOWithDefaults

`func NewOverViewSummaryVOWithDefaults() *OverViewSummaryVO`

NewOverViewSummaryVOWithDefaults instantiates a new OverViewSummaryVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientHealthScore

`func (o *OverViewSummaryVO) GetClientHealthScore() int32`

GetClientHealthScore returns the ClientHealthScore field if non-nil, zero value otherwise.

### GetClientHealthScoreOk

`func (o *OverViewSummaryVO) GetClientHealthScoreOk() (*int32, bool)`

GetClientHealthScoreOk returns a tuple with the ClientHealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHealthScore

`func (o *OverViewSummaryVO) SetClientHealthScore(v int32)`

SetClientHealthScore sets ClientHealthScore field to given value.

### HasClientHealthScore

`func (o *OverViewSummaryVO) HasClientHealthScore() bool`

HasClientHealthScore returns a boolean if a field has been set.

### GetDeviceHealthScore

`func (o *OverViewSummaryVO) GetDeviceHealthScore() int32`

GetDeviceHealthScore returns the DeviceHealthScore field if non-nil, zero value otherwise.

### GetDeviceHealthScoreOk

`func (o *OverViewSummaryVO) GetDeviceHealthScoreOk() (*int32, bool)`

GetDeviceHealthScoreOk returns a tuple with the DeviceHealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceHealthScore

`func (o *OverViewSummaryVO) SetDeviceHealthScore(v int32)`

SetDeviceHealthScore sets DeviceHealthScore field to given value.

### HasDeviceHealthScore

`func (o *OverViewSummaryVO) HasDeviceHealthScore() bool`

HasDeviceHealthScore returns a boolean if a field has been set.

### GetSiteHealthScore

`func (o *OverViewSummaryVO) GetSiteHealthScore() int32`

GetSiteHealthScore returns the SiteHealthScore field if non-nil, zero value otherwise.

### GetSiteHealthScoreOk

`func (o *OverViewSummaryVO) GetSiteHealthScoreOk() (*int32, bool)`

GetSiteHealthScoreOk returns a tuple with the SiteHealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteHealthScore

`func (o *OverViewSummaryVO) SetSiteHealthScore(v int32)`

SetSiteHealthScore sets SiteHealthScore field to given value.

### HasSiteHealthScore

`func (o *OverViewSummaryVO) HasSiteHealthScore() bool`

HasSiteHealthScore returns a boolean if a field has been set.

### GetTotalAps

`func (o *OverViewSummaryVO) GetTotalAps() int64`

GetTotalAps returns the TotalAps field if non-nil, zero value otherwise.

### GetTotalApsOk

`func (o *OverViewSummaryVO) GetTotalApsOk() (*int64, bool)`

GetTotalApsOk returns a tuple with the TotalAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAps

`func (o *OverViewSummaryVO) SetTotalAps(v int64)`

SetTotalAps sets TotalAps field to given value.

### HasTotalAps

`func (o *OverViewSummaryVO) HasTotalAps() bool`

HasTotalAps returns a boolean if a field has been set.

### GetTotalClients

`func (o *OverViewSummaryVO) GetTotalClients() int64`

GetTotalClients returns the TotalClients field if non-nil, zero value otherwise.

### GetTotalClientsOk

`func (o *OverViewSummaryVO) GetTotalClientsOk() (*int64, bool)`

GetTotalClientsOk returns a tuple with the TotalClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalClients

`func (o *OverViewSummaryVO) SetTotalClients(v int64)`

SetTotalClients sets TotalClients field to given value.

### HasTotalClients

`func (o *OverViewSummaryVO) HasTotalClients() bool`

HasTotalClients returns a boolean if a field has been set.

### GetTotalDevices

`func (o *OverViewSummaryVO) GetTotalDevices() int64`

GetTotalDevices returns the TotalDevices field if non-nil, zero value otherwise.

### GetTotalDevicesOk

`func (o *OverViewSummaryVO) GetTotalDevicesOk() (*int64, bool)`

GetTotalDevicesOk returns a tuple with the TotalDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDevices

`func (o *OverViewSummaryVO) SetTotalDevices(v int64)`

SetTotalDevices sets TotalDevices field to given value.

### HasTotalDevices

`func (o *OverViewSummaryVO) HasTotalDevices() bool`

HasTotalDevices returns a boolean if a field has been set.

### GetTotalGateways

`func (o *OverViewSummaryVO) GetTotalGateways() int64`

GetTotalGateways returns the TotalGateways field if non-nil, zero value otherwise.

### GetTotalGatewaysOk

`func (o *OverViewSummaryVO) GetTotalGatewaysOk() (*int64, bool)`

GetTotalGatewaysOk returns a tuple with the TotalGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGateways

`func (o *OverViewSummaryVO) SetTotalGateways(v int64)`

SetTotalGateways sets TotalGateways field to given value.

### HasTotalGateways

`func (o *OverViewSummaryVO) HasTotalGateways() bool`

HasTotalGateways returns a boolean if a field has been set.

### GetTotalIpcs

`func (o *OverViewSummaryVO) GetTotalIpcs() int64`

GetTotalIpcs returns the TotalIpcs field if non-nil, zero value otherwise.

### GetTotalIpcsOk

`func (o *OverViewSummaryVO) GetTotalIpcsOk() (*int64, bool)`

GetTotalIpcsOk returns a tuple with the TotalIpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIpcs

`func (o *OverViewSummaryVO) SetTotalIpcs(v int64)`

SetTotalIpcs sets TotalIpcs field to given value.

### HasTotalIpcs

`func (o *OverViewSummaryVO) HasTotalIpcs() bool`

HasTotalIpcs returns a boolean if a field has been set.

### GetTotalNvrs

`func (o *OverViewSummaryVO) GetTotalNvrs() int64`

GetTotalNvrs returns the TotalNvrs field if non-nil, zero value otherwise.

### GetTotalNvrsOk

`func (o *OverViewSummaryVO) GetTotalNvrsOk() (*int64, bool)`

GetTotalNvrsOk returns a tuple with the TotalNvrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNvrs

`func (o *OverViewSummaryVO) SetTotalNvrs(v int64)`

SetTotalNvrs sets TotalNvrs field to given value.

### HasTotalNvrs

`func (o *OverViewSummaryVO) HasTotalNvrs() bool`

HasTotalNvrs returns a boolean if a field has been set.

### GetTotalOlts

`func (o *OverViewSummaryVO) GetTotalOlts() int64`

GetTotalOlts returns the TotalOlts field if non-nil, zero value otherwise.

### GetTotalOltsOk

`func (o *OverViewSummaryVO) GetTotalOltsOk() (*int64, bool)`

GetTotalOltsOk returns a tuple with the TotalOlts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOlts

`func (o *OverViewSummaryVO) SetTotalOlts(v int64)`

SetTotalOlts sets TotalOlts field to given value.

### HasTotalOlts

`func (o *OverViewSummaryVO) HasTotalOlts() bool`

HasTotalOlts returns a boolean if a field has been set.

### GetTotalSwitches

`func (o *OverViewSummaryVO) GetTotalSwitches() int64`

GetTotalSwitches returns the TotalSwitches field if non-nil, zero value otherwise.

### GetTotalSwitchesOk

`func (o *OverViewSummaryVO) GetTotalSwitchesOk() (*int64, bool)`

GetTotalSwitchesOk returns a tuple with the TotalSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSwitches

`func (o *OverViewSummaryVO) SetTotalSwitches(v int64)`

SetTotalSwitches sets TotalSwitches field to given value.

### HasTotalSwitches

`func (o *OverViewSummaryVO) HasTotalSwitches() bool`

HasTotalSwitches returns a boolean if a field has been set.

### GetTotalTraffic

`func (o *OverViewSummaryVO) GetTotalTraffic() int64`

GetTotalTraffic returns the TotalTraffic field if non-nil, zero value otherwise.

### GetTotalTrafficOk

`func (o *OverViewSummaryVO) GetTotalTrafficOk() (*int64, bool)`

GetTotalTrafficOk returns a tuple with the TotalTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTraffic

`func (o *OverViewSummaryVO) SetTotalTraffic(v int64)`

SetTotalTraffic sets TotalTraffic field to given value.

### HasTotalTraffic

`func (o *OverViewSummaryVO) HasTotalTraffic() bool`

HasTotalTraffic returns a boolean if a field has been set.

### GetWanHealthScore

`func (o *OverViewSummaryVO) GetWanHealthScore() int32`

GetWanHealthScore returns the WanHealthScore field if non-nil, zero value otherwise.

### GetWanHealthScoreOk

`func (o *OverViewSummaryVO) GetWanHealthScoreOk() (*int32, bool)`

GetWanHealthScoreOk returns a tuple with the WanHealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanHealthScore

`func (o *OverViewSummaryVO) SetWanHealthScore(v int32)`

SetWanHealthScore sets WanHealthScore field to given value.

### HasWanHealthScore

`func (o *OverViewSummaryVO) HasWanHealthScore() bool`

HasWanHealthScore returns a boolean if a field has been set.

### GetWifiHealthScore

`func (o *OverViewSummaryVO) GetWifiHealthScore() int32`

GetWifiHealthScore returns the WifiHealthScore field if non-nil, zero value otherwise.

### GetWifiHealthScoreOk

`func (o *OverViewSummaryVO) GetWifiHealthScoreOk() (*int32, bool)`

GetWifiHealthScoreOk returns a tuple with the WifiHealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiHealthScore

`func (o *OverViewSummaryVO) SetWifiHealthScore(v int32)`

SetWifiHealthScore sets WifiHealthScore field to given value.

### HasWifiHealthScore

`func (o *OverViewSummaryVO) HasWifiHealthScore() bool`

HasWifiHealthScore returns a boolean if a field has been set.

### GetWiredClients

`func (o *OverViewSummaryVO) GetWiredClients() int64`

GetWiredClients returns the WiredClients field if non-nil, zero value otherwise.

### GetWiredClientsOk

`func (o *OverViewSummaryVO) GetWiredClientsOk() (*int64, bool)`

GetWiredClientsOk returns a tuple with the WiredClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredClients

`func (o *OverViewSummaryVO) SetWiredClients(v int64)`

SetWiredClients sets WiredClients field to given value.

### HasWiredClients

`func (o *OverViewSummaryVO) HasWiredClients() bool`

HasWiredClients returns a boolean if a field has been set.

### GetWiredTraffic

`func (o *OverViewSummaryVO) GetWiredTraffic() int64`

GetWiredTraffic returns the WiredTraffic field if non-nil, zero value otherwise.

### GetWiredTrafficOk

`func (o *OverViewSummaryVO) GetWiredTrafficOk() (*int64, bool)`

GetWiredTrafficOk returns a tuple with the WiredTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredTraffic

`func (o *OverViewSummaryVO) SetWiredTraffic(v int64)`

SetWiredTraffic sets WiredTraffic field to given value.

### HasWiredTraffic

`func (o *OverViewSummaryVO) HasWiredTraffic() bool`

HasWiredTraffic returns a boolean if a field has been set.

### GetWirelessClients

`func (o *OverViewSummaryVO) GetWirelessClients() int64`

GetWirelessClients returns the WirelessClients field if non-nil, zero value otherwise.

### GetWirelessClientsOk

`func (o *OverViewSummaryVO) GetWirelessClientsOk() (*int64, bool)`

GetWirelessClientsOk returns a tuple with the WirelessClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessClients

`func (o *OverViewSummaryVO) SetWirelessClients(v int64)`

SetWirelessClients sets WirelessClients field to given value.

### HasWirelessClients

`func (o *OverViewSummaryVO) HasWirelessClients() bool`

HasWirelessClients returns a boolean if a field has been set.

### GetWirelessTraffic

`func (o *OverViewSummaryVO) GetWirelessTraffic() int64`

GetWirelessTraffic returns the WirelessTraffic field if non-nil, zero value otherwise.

### GetWirelessTrafficOk

`func (o *OverViewSummaryVO) GetWirelessTrafficOk() (*int64, bool)`

GetWirelessTrafficOk returns a tuple with the WirelessTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessTraffic

`func (o *OverViewSummaryVO) SetWirelessTraffic(v int64)`

SetWirelessTraffic sets WirelessTraffic field to given value.

### HasWirelessTraffic

`func (o *OverViewSummaryVO) HasWirelessTraffic() bool`

HasWirelessTraffic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


