# TrafficDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aps** | Pointer to [**[]DeviceTrafficDistribution**](DeviceTrafficDistribution.md) | Current Site Ap List | [optional] 
**Switches** | Pointer to [**[]DeviceTrafficDistribution**](DeviceTrafficDistribution.md) | Current Site Switch List | [optional] 

## Methods

### NewTrafficDistribution

`func NewTrafficDistribution() *TrafficDistribution`

NewTrafficDistribution instantiates a new TrafficDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficDistributionWithDefaults

`func NewTrafficDistributionWithDefaults() *TrafficDistribution`

NewTrafficDistributionWithDefaults instantiates a new TrafficDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAps

`func (o *TrafficDistribution) GetAps() []DeviceTrafficDistribution`

GetAps returns the Aps field if non-nil, zero value otherwise.

### GetApsOk

`func (o *TrafficDistribution) GetApsOk() (*[]DeviceTrafficDistribution, bool)`

GetApsOk returns a tuple with the Aps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAps

`func (o *TrafficDistribution) SetAps(v []DeviceTrafficDistribution)`

SetAps sets Aps field to given value.

### HasAps

`func (o *TrafficDistribution) HasAps() bool`

HasAps returns a boolean if a field has been set.

### GetSwitches

`func (o *TrafficDistribution) GetSwitches() []DeviceTrafficDistribution`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *TrafficDistribution) GetSwitchesOk() (*[]DeviceTrafficDistribution, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *TrafficDistribution) SetSwitches(v []DeviceTrafficDistribution)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *TrafficDistribution) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


