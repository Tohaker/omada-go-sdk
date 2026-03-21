# WifiSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelUtilization** | Pointer to **int32** | The average channel utilization of the 2.4GHz and 5GHz band for each AP in percentage. Null means no value. | [optional] 
**Clients** | Pointer to **int32** | The number of connected wireless clients | [optional] 
**ConnectedApNum** | Pointer to **int32** | The number of connected APs | [optional] 
**TotalTraffic** | Pointer to **float64** | Total traffic measured in KB | [optional] 

## Methods

### NewWifiSummary

`func NewWifiSummary() *WifiSummary`

NewWifiSummary instantiates a new WifiSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWifiSummaryWithDefaults

`func NewWifiSummaryWithDefaults() *WifiSummary`

NewWifiSummaryWithDefaults instantiates a new WifiSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelUtilization

`func (o *WifiSummary) GetChannelUtilization() int32`

GetChannelUtilization returns the ChannelUtilization field if non-nil, zero value otherwise.

### GetChannelUtilizationOk

`func (o *WifiSummary) GetChannelUtilizationOk() (*int32, bool)`

GetChannelUtilizationOk returns a tuple with the ChannelUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelUtilization

`func (o *WifiSummary) SetChannelUtilization(v int32)`

SetChannelUtilization sets ChannelUtilization field to given value.

### HasChannelUtilization

`func (o *WifiSummary) HasChannelUtilization() bool`

HasChannelUtilization returns a boolean if a field has been set.

### GetClients

`func (o *WifiSummary) GetClients() int32`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *WifiSummary) GetClientsOk() (*int32, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *WifiSummary) SetClients(v int32)`

SetClients sets Clients field to given value.

### HasClients

`func (o *WifiSummary) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetConnectedApNum

`func (o *WifiSummary) GetConnectedApNum() int32`

GetConnectedApNum returns the ConnectedApNum field if non-nil, zero value otherwise.

### GetConnectedApNumOk

`func (o *WifiSummary) GetConnectedApNumOk() (*int32, bool)`

GetConnectedApNumOk returns a tuple with the ConnectedApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedApNum

`func (o *WifiSummary) SetConnectedApNum(v int32)`

SetConnectedApNum sets ConnectedApNum field to given value.

### HasConnectedApNum

`func (o *WifiSummary) HasConnectedApNum() bool`

HasConnectedApNum returns a boolean if a field has been set.

### GetTotalTraffic

`func (o *WifiSummary) GetTotalTraffic() float64`

GetTotalTraffic returns the TotalTraffic field if non-nil, zero value otherwise.

### GetTotalTrafficOk

`func (o *WifiSummary) GetTotalTrafficOk() (*float64, bool)`

GetTotalTrafficOk returns a tuple with the TotalTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTraffic

`func (o *WifiSummary) SetTotalTraffic(v float64)`

SetTotalTraffic sets TotalTraffic field to given value.

### HasTotalTraffic

`func (o *WifiSummary) HasTotalTraffic() bool`

HasTotalTraffic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


