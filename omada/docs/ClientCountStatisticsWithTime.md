# ClientCountStatisticsWithTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **int64** | Timestamp, unit is second. | [optional] 
**TotalClients** | Pointer to **int32** | The number of all clients. | [optional] 
**WiredClients** | Pointer to **int32** | The number of all wired clients. | [optional] 
**WirelessClients** | Pointer to **int32** | The number of all wireless clients. | [optional] 

## Methods

### NewClientCountStatisticsWithTime

`func NewClientCountStatisticsWithTime() *ClientCountStatisticsWithTime`

NewClientCountStatisticsWithTime instantiates a new ClientCountStatisticsWithTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCountStatisticsWithTimeWithDefaults

`func NewClientCountStatisticsWithTimeWithDefaults() *ClientCountStatisticsWithTime`

NewClientCountStatisticsWithTimeWithDefaults instantiates a new ClientCountStatisticsWithTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *ClientCountStatisticsWithTime) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ClientCountStatisticsWithTime) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ClientCountStatisticsWithTime) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *ClientCountStatisticsWithTime) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTotalClients

`func (o *ClientCountStatisticsWithTime) GetTotalClients() int32`

GetTotalClients returns the TotalClients field if non-nil, zero value otherwise.

### GetTotalClientsOk

`func (o *ClientCountStatisticsWithTime) GetTotalClientsOk() (*int32, bool)`

GetTotalClientsOk returns a tuple with the TotalClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalClients

`func (o *ClientCountStatisticsWithTime) SetTotalClients(v int32)`

SetTotalClients sets TotalClients field to given value.

### HasTotalClients

`func (o *ClientCountStatisticsWithTime) HasTotalClients() bool`

HasTotalClients returns a boolean if a field has been set.

### GetWiredClients

`func (o *ClientCountStatisticsWithTime) GetWiredClients() int32`

GetWiredClients returns the WiredClients field if non-nil, zero value otherwise.

### GetWiredClientsOk

`func (o *ClientCountStatisticsWithTime) GetWiredClientsOk() (*int32, bool)`

GetWiredClientsOk returns a tuple with the WiredClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredClients

`func (o *ClientCountStatisticsWithTime) SetWiredClients(v int32)`

SetWiredClients sets WiredClients field to given value.

### HasWiredClients

`func (o *ClientCountStatisticsWithTime) HasWiredClients() bool`

HasWiredClients returns a boolean if a field has been set.

### GetWirelessClients

`func (o *ClientCountStatisticsWithTime) GetWirelessClients() int32`

GetWirelessClients returns the WirelessClients field if non-nil, zero value otherwise.

### GetWirelessClientsOk

`func (o *ClientCountStatisticsWithTime) GetWirelessClientsOk() (*int32, bool)`

GetWirelessClientsOk returns a tuple with the WirelessClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessClients

`func (o *ClientCountStatisticsWithTime) SetWirelessClients(v int32)`

SetWirelessClients sets WirelessClients field to given value.

### HasWirelessClients

`func (o *ClientCountStatisticsWithTime) HasWirelessClients() bool`

HasWirelessClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


