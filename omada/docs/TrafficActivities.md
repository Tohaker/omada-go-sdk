# TrafficActivities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApTrafficActivities** | Pointer to [**[]DeviceTrafficActivity**](DeviceTrafficActivity.md) | Wireless network total traffic timing list | [optional] 
**SwitchTrafficActivities** | Pointer to [**[]DeviceTrafficActivity**](DeviceTrafficActivity.md) | Wired network total traffic timing list | [optional] 

## Methods

### NewTrafficActivities

`func NewTrafficActivities() *TrafficActivities`

NewTrafficActivities instantiates a new TrafficActivities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficActivitiesWithDefaults

`func NewTrafficActivitiesWithDefaults() *TrafficActivities`

NewTrafficActivitiesWithDefaults instantiates a new TrafficActivities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApTrafficActivities

`func (o *TrafficActivities) GetApTrafficActivities() []DeviceTrafficActivity`

GetApTrafficActivities returns the ApTrafficActivities field if non-nil, zero value otherwise.

### GetApTrafficActivitiesOk

`func (o *TrafficActivities) GetApTrafficActivitiesOk() (*[]DeviceTrafficActivity, bool)`

GetApTrafficActivitiesOk returns a tuple with the ApTrafficActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApTrafficActivities

`func (o *TrafficActivities) SetApTrafficActivities(v []DeviceTrafficActivity)`

SetApTrafficActivities sets ApTrafficActivities field to given value.

### HasApTrafficActivities

`func (o *TrafficActivities) HasApTrafficActivities() bool`

HasApTrafficActivities returns a boolean if a field has been set.

### GetSwitchTrafficActivities

`func (o *TrafficActivities) GetSwitchTrafficActivities() []DeviceTrafficActivity`

GetSwitchTrafficActivities returns the SwitchTrafficActivities field if non-nil, zero value otherwise.

### GetSwitchTrafficActivitiesOk

`func (o *TrafficActivities) GetSwitchTrafficActivitiesOk() (*[]DeviceTrafficActivity, bool)`

GetSwitchTrafficActivitiesOk returns a tuple with the SwitchTrafficActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchTrafficActivities

`func (o *TrafficActivities) SetSwitchTrafficActivities(v []DeviceTrafficActivity)`

SetSwitchTrafficActivities sets SwitchTrafficActivities field to given value.

### HasSwitchTrafficActivities

`func (o *TrafficActivities) HasSwitchTrafficActivities() bool`

HasSwitchTrafficActivities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


