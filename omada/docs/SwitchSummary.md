# SwitchSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients** | Pointer to **int32** | Number of wired clients | [optional] 
**ConnectedSwitchNum** | Pointer to **int32** | Number of online switches | [optional] 
**PortUtilization** | Pointer to **int32** | Port occupancy rate (integer) | [optional] 
**TotalTraffic** | Pointer to **int64** | Upstream and downstream traffic and the unit (Byte) of online switches | [optional] 

## Methods

### NewSwitchSummary

`func NewSwitchSummary() *SwitchSummary`

NewSwitchSummary instantiates a new SwitchSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchSummaryWithDefaults

`func NewSwitchSummaryWithDefaults() *SwitchSummary`

NewSwitchSummaryWithDefaults instantiates a new SwitchSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients

`func (o *SwitchSummary) GetClients() int32`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *SwitchSummary) GetClientsOk() (*int32, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *SwitchSummary) SetClients(v int32)`

SetClients sets Clients field to given value.

### HasClients

`func (o *SwitchSummary) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetConnectedSwitchNum

`func (o *SwitchSummary) GetConnectedSwitchNum() int32`

GetConnectedSwitchNum returns the ConnectedSwitchNum field if non-nil, zero value otherwise.

### GetConnectedSwitchNumOk

`func (o *SwitchSummary) GetConnectedSwitchNumOk() (*int32, bool)`

GetConnectedSwitchNumOk returns a tuple with the ConnectedSwitchNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedSwitchNum

`func (o *SwitchSummary) SetConnectedSwitchNum(v int32)`

SetConnectedSwitchNum sets ConnectedSwitchNum field to given value.

### HasConnectedSwitchNum

`func (o *SwitchSummary) HasConnectedSwitchNum() bool`

HasConnectedSwitchNum returns a boolean if a field has been set.

### GetPortUtilization

`func (o *SwitchSummary) GetPortUtilization() int32`

GetPortUtilization returns the PortUtilization field if non-nil, zero value otherwise.

### GetPortUtilizationOk

`func (o *SwitchSummary) GetPortUtilizationOk() (*int32, bool)`

GetPortUtilizationOk returns a tuple with the PortUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUtilization

`func (o *SwitchSummary) SetPortUtilization(v int32)`

SetPortUtilization sets PortUtilization field to given value.

### HasPortUtilization

`func (o *SwitchSummary) HasPortUtilization() bool`

HasPortUtilization returns a boolean if a field has been set.

### GetTotalTraffic

`func (o *SwitchSummary) GetTotalTraffic() int64`

GetTotalTraffic returns the TotalTraffic field if non-nil, zero value otherwise.

### GetTotalTrafficOk

`func (o *SwitchSummary) GetTotalTrafficOk() (*int64, bool)`

GetTotalTrafficOk returns a tuple with the TotalTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTraffic

`func (o *SwitchSummary) SetTotalTraffic(v int64)`

SetTotalTraffic sets TotalTraffic field to given value.

### HasTotalTraffic

`func (o *SwitchSummary) HasTotalTraffic() bool`

HasTotalTraffic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


