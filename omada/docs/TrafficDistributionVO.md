# TrafficDistributionVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalTraffic** | Pointer to **int64** | total traffic | [optional] 
**TrafficDistribution** | Pointer to [**[]TrafficDistributionListVO**](TrafficDistributionListVO.md) | traffic data point | [optional] 
**WiredTraffic** | Pointer to **int64** | traffic of wired device | [optional] 
**WirelessTraffic** | Pointer to **int64** | traffic of wireless device | [optional] 

## Methods

### NewTrafficDistributionVO

`func NewTrafficDistributionVO() *TrafficDistributionVO`

NewTrafficDistributionVO instantiates a new TrafficDistributionVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficDistributionVOWithDefaults

`func NewTrafficDistributionVOWithDefaults() *TrafficDistributionVO`

NewTrafficDistributionVOWithDefaults instantiates a new TrafficDistributionVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalTraffic

`func (o *TrafficDistributionVO) GetTotalTraffic() int64`

GetTotalTraffic returns the TotalTraffic field if non-nil, zero value otherwise.

### GetTotalTrafficOk

`func (o *TrafficDistributionVO) GetTotalTrafficOk() (*int64, bool)`

GetTotalTrafficOk returns a tuple with the TotalTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTraffic

`func (o *TrafficDistributionVO) SetTotalTraffic(v int64)`

SetTotalTraffic sets TotalTraffic field to given value.

### HasTotalTraffic

`func (o *TrafficDistributionVO) HasTotalTraffic() bool`

HasTotalTraffic returns a boolean if a field has been set.

### GetTrafficDistribution

`func (o *TrafficDistributionVO) GetTrafficDistribution() []TrafficDistributionListVO`

GetTrafficDistribution returns the TrafficDistribution field if non-nil, zero value otherwise.

### GetTrafficDistributionOk

`func (o *TrafficDistributionVO) GetTrafficDistributionOk() (*[]TrafficDistributionListVO, bool)`

GetTrafficDistributionOk returns a tuple with the TrafficDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficDistribution

`func (o *TrafficDistributionVO) SetTrafficDistribution(v []TrafficDistributionListVO)`

SetTrafficDistribution sets TrafficDistribution field to given value.

### HasTrafficDistribution

`func (o *TrafficDistributionVO) HasTrafficDistribution() bool`

HasTrafficDistribution returns a boolean if a field has been set.

### GetWiredTraffic

`func (o *TrafficDistributionVO) GetWiredTraffic() int64`

GetWiredTraffic returns the WiredTraffic field if non-nil, zero value otherwise.

### GetWiredTrafficOk

`func (o *TrafficDistributionVO) GetWiredTrafficOk() (*int64, bool)`

GetWiredTrafficOk returns a tuple with the WiredTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredTraffic

`func (o *TrafficDistributionVO) SetWiredTraffic(v int64)`

SetWiredTraffic sets WiredTraffic field to given value.

### HasWiredTraffic

`func (o *TrafficDistributionVO) HasWiredTraffic() bool`

HasWiredTraffic returns a boolean if a field has been set.

### GetWirelessTraffic

`func (o *TrafficDistributionVO) GetWirelessTraffic() int64`

GetWirelessTraffic returns the WirelessTraffic field if non-nil, zero value otherwise.

### GetWirelessTrafficOk

`func (o *TrafficDistributionVO) GetWirelessTrafficOk() (*int64, bool)`

GetWirelessTrafficOk returns a tuple with the WirelessTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessTraffic

`func (o *TrafficDistributionVO) SetWirelessTraffic(v int64)`

SetWirelessTraffic sets WirelessTraffic field to given value.

### HasWirelessTraffic

`func (o *TrafficDistributionVO) HasWirelessTraffic() bool`

HasWirelessTraffic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


