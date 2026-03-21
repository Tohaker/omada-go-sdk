# ClientStatisticsOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageClientNum** | Pointer to **int32** | Average client number of each day. | [optional] 
**AverageClientTraffic** | Pointer to **int64** | Average traffic of each client. | [optional] 
**Clients2gNum** | Pointer to **int32** | Total number of wireless clients in the 2G band. | [optional] 
**Clients5gNum** | Pointer to **int32** | Total number of wireless clients in the 5G band. | [optional] 
**Clients6gNum** | Pointer to **int32** | Total number of wireless clients in the 6G band. | [optional] 
**ClientsSsidDistribution** | Pointer to [**SsidDistribution**](SsidDistribution.md) |  | [optional] 
**TotalClientsNum** | Pointer to **int32** | Total number of all clients. | [optional] 
**WiredClientsNum** | Pointer to **int32** | Total number of all wired clients. | [optional] 
**WirelessClientsNum** | Pointer to **int32** | Total number of all wireless clients. | [optional] 

## Methods

### NewClientStatisticsOverview

`func NewClientStatisticsOverview() *ClientStatisticsOverview`

NewClientStatisticsOverview instantiates a new ClientStatisticsOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientStatisticsOverviewWithDefaults

`func NewClientStatisticsOverviewWithDefaults() *ClientStatisticsOverview`

NewClientStatisticsOverviewWithDefaults instantiates a new ClientStatisticsOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageClientNum

`func (o *ClientStatisticsOverview) GetAverageClientNum() int32`

GetAverageClientNum returns the AverageClientNum field if non-nil, zero value otherwise.

### GetAverageClientNumOk

`func (o *ClientStatisticsOverview) GetAverageClientNumOk() (*int32, bool)`

GetAverageClientNumOk returns a tuple with the AverageClientNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageClientNum

`func (o *ClientStatisticsOverview) SetAverageClientNum(v int32)`

SetAverageClientNum sets AverageClientNum field to given value.

### HasAverageClientNum

`func (o *ClientStatisticsOverview) HasAverageClientNum() bool`

HasAverageClientNum returns a boolean if a field has been set.

### GetAverageClientTraffic

`func (o *ClientStatisticsOverview) GetAverageClientTraffic() int64`

GetAverageClientTraffic returns the AverageClientTraffic field if non-nil, zero value otherwise.

### GetAverageClientTrafficOk

`func (o *ClientStatisticsOverview) GetAverageClientTrafficOk() (*int64, bool)`

GetAverageClientTrafficOk returns a tuple with the AverageClientTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageClientTraffic

`func (o *ClientStatisticsOverview) SetAverageClientTraffic(v int64)`

SetAverageClientTraffic sets AverageClientTraffic field to given value.

### HasAverageClientTraffic

`func (o *ClientStatisticsOverview) HasAverageClientTraffic() bool`

HasAverageClientTraffic returns a boolean if a field has been set.

### GetClients2gNum

`func (o *ClientStatisticsOverview) GetClients2gNum() int32`

GetClients2gNum returns the Clients2gNum field if non-nil, zero value otherwise.

### GetClients2gNumOk

`func (o *ClientStatisticsOverview) GetClients2gNumOk() (*int32, bool)`

GetClients2gNumOk returns a tuple with the Clients2gNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients2gNum

`func (o *ClientStatisticsOverview) SetClients2gNum(v int32)`

SetClients2gNum sets Clients2gNum field to given value.

### HasClients2gNum

`func (o *ClientStatisticsOverview) HasClients2gNum() bool`

HasClients2gNum returns a boolean if a field has been set.

### GetClients5gNum

`func (o *ClientStatisticsOverview) GetClients5gNum() int32`

GetClients5gNum returns the Clients5gNum field if non-nil, zero value otherwise.

### GetClients5gNumOk

`func (o *ClientStatisticsOverview) GetClients5gNumOk() (*int32, bool)`

GetClients5gNumOk returns a tuple with the Clients5gNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients5gNum

`func (o *ClientStatisticsOverview) SetClients5gNum(v int32)`

SetClients5gNum sets Clients5gNum field to given value.

### HasClients5gNum

`func (o *ClientStatisticsOverview) HasClients5gNum() bool`

HasClients5gNum returns a boolean if a field has been set.

### GetClients6gNum

`func (o *ClientStatisticsOverview) GetClients6gNum() int32`

GetClients6gNum returns the Clients6gNum field if non-nil, zero value otherwise.

### GetClients6gNumOk

`func (o *ClientStatisticsOverview) GetClients6gNumOk() (*int32, bool)`

GetClients6gNumOk returns a tuple with the Clients6gNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients6gNum

`func (o *ClientStatisticsOverview) SetClients6gNum(v int32)`

SetClients6gNum sets Clients6gNum field to given value.

### HasClients6gNum

`func (o *ClientStatisticsOverview) HasClients6gNum() bool`

HasClients6gNum returns a boolean if a field has been set.

### GetClientsSsidDistribution

`func (o *ClientStatisticsOverview) GetClientsSsidDistribution() SsidDistribution`

GetClientsSsidDistribution returns the ClientsSsidDistribution field if non-nil, zero value otherwise.

### GetClientsSsidDistributionOk

`func (o *ClientStatisticsOverview) GetClientsSsidDistributionOk() (*SsidDistribution, bool)`

GetClientsSsidDistributionOk returns a tuple with the ClientsSsidDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsSsidDistribution

`func (o *ClientStatisticsOverview) SetClientsSsidDistribution(v SsidDistribution)`

SetClientsSsidDistribution sets ClientsSsidDistribution field to given value.

### HasClientsSsidDistribution

`func (o *ClientStatisticsOverview) HasClientsSsidDistribution() bool`

HasClientsSsidDistribution returns a boolean if a field has been set.

### GetTotalClientsNum

`func (o *ClientStatisticsOverview) GetTotalClientsNum() int32`

GetTotalClientsNum returns the TotalClientsNum field if non-nil, zero value otherwise.

### GetTotalClientsNumOk

`func (o *ClientStatisticsOverview) GetTotalClientsNumOk() (*int32, bool)`

GetTotalClientsNumOk returns a tuple with the TotalClientsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalClientsNum

`func (o *ClientStatisticsOverview) SetTotalClientsNum(v int32)`

SetTotalClientsNum sets TotalClientsNum field to given value.

### HasTotalClientsNum

`func (o *ClientStatisticsOverview) HasTotalClientsNum() bool`

HasTotalClientsNum returns a boolean if a field has been set.

### GetWiredClientsNum

`func (o *ClientStatisticsOverview) GetWiredClientsNum() int32`

GetWiredClientsNum returns the WiredClientsNum field if non-nil, zero value otherwise.

### GetWiredClientsNumOk

`func (o *ClientStatisticsOverview) GetWiredClientsNumOk() (*int32, bool)`

GetWiredClientsNumOk returns a tuple with the WiredClientsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredClientsNum

`func (o *ClientStatisticsOverview) SetWiredClientsNum(v int32)`

SetWiredClientsNum sets WiredClientsNum field to given value.

### HasWiredClientsNum

`func (o *ClientStatisticsOverview) HasWiredClientsNum() bool`

HasWiredClientsNum returns a boolean if a field has been set.

### GetWirelessClientsNum

`func (o *ClientStatisticsOverview) GetWirelessClientsNum() int32`

GetWirelessClientsNum returns the WirelessClientsNum field if non-nil, zero value otherwise.

### GetWirelessClientsNumOk

`func (o *ClientStatisticsOverview) GetWirelessClientsNumOk() (*int32, bool)`

GetWirelessClientsNumOk returns a tuple with the WirelessClientsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessClientsNum

`func (o *ClientStatisticsOverview) SetWirelessClientsNum(v int32)`

SetWirelessClientsNum sets WirelessClientsNum field to given value.

### HasWirelessClientsNum

`func (o *ClientStatisticsOverview) HasWirelessClientsNum() bool`

HasWirelessClientsNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


