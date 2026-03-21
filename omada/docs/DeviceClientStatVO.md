# DeviceClientStatVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewDeviceClientStatVO

`func NewDeviceClientStatVO() *DeviceClientStatVO`

NewDeviceClientStatVO instantiates a new DeviceClientStatVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceClientStatVOWithDefaults

`func NewDeviceClientStatVOWithDefaults() *DeviceClientStatVO`

NewDeviceClientStatVOWithDefaults instantiates a new DeviceClientStatVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients

`func (o *DeviceClientStatVO) GetClients() int32`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *DeviceClientStatVO) GetClientsOk() (*int32, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *DeviceClientStatVO) SetClients(v int32)`

SetClients sets Clients field to given value.

### HasClients

`func (o *DeviceClientStatVO) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetTime

`func (o *DeviceClientStatVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DeviceClientStatVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DeviceClientStatVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *DeviceClientStatVO) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


