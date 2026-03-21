# DeviceTrafficDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Device MAC | [optional] 
**Name** | Pointer to **string** | Device Name | [optional] 
**Traffic** | Pointer to **float64** | Device traffic measured in MB | [optional] 
**TrafficProportion** | Pointer to **float64** | The proportion of AP traffic in percentage | [optional] 

## Methods

### NewDeviceTrafficDistribution

`func NewDeviceTrafficDistribution() *DeviceTrafficDistribution`

NewDeviceTrafficDistribution instantiates a new DeviceTrafficDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTrafficDistributionWithDefaults

`func NewDeviceTrafficDistributionWithDefaults() *DeviceTrafficDistribution`

NewDeviceTrafficDistributionWithDefaults instantiates a new DeviceTrafficDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *DeviceTrafficDistribution) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DeviceTrafficDistribution) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DeviceTrafficDistribution) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *DeviceTrafficDistribution) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *DeviceTrafficDistribution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceTrafficDistribution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceTrafficDistribution) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceTrafficDistribution) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTraffic

`func (o *DeviceTrafficDistribution) GetTraffic() float64`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *DeviceTrafficDistribution) GetTrafficOk() (*float64, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *DeviceTrafficDistribution) SetTraffic(v float64)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *DeviceTrafficDistribution) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.

### GetTrafficProportion

`func (o *DeviceTrafficDistribution) GetTrafficProportion() float64`

GetTrafficProportion returns the TrafficProportion field if non-nil, zero value otherwise.

### GetTrafficProportionOk

`func (o *DeviceTrafficDistribution) GetTrafficProportionOk() (*float64, bool)`

GetTrafficProportionOk returns a tuple with the TrafficProportion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficProportion

`func (o *DeviceTrafficDistribution) SetTrafficProportion(v float64)`

SetTrafficProportion sets TrafficProportion field to given value.

### HasTrafficProportion

`func (o *DeviceTrafficDistribution) HasTrafficProportion() bool`

HasTrafficProportion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


