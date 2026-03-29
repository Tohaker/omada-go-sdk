# DeviceTrafficActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RxData** | Pointer to **float64** | Downlink traffic in MB | [optional] 
**Time** | Pointer to **int64** | The timestamp of this data sample, in seconds, such as 1682000000 | [optional] 
**TxData** | Pointer to **float64** | Uplink traffic in MB | [optional] 

## Methods

### NewDeviceTrafficActivity

`func NewDeviceTrafficActivity() *DeviceTrafficActivity`

NewDeviceTrafficActivity instantiates a new DeviceTrafficActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTrafficActivityWithDefaults

`func NewDeviceTrafficActivityWithDefaults() *DeviceTrafficActivity`

NewDeviceTrafficActivityWithDefaults instantiates a new DeviceTrafficActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRxData

`func (o *DeviceTrafficActivity) GetRxData() float64`

GetRxData returns the RxData field if non-nil, zero value otherwise.

### GetRxDataOk

`func (o *DeviceTrafficActivity) GetRxDataOk() (*float64, bool)`

GetRxDataOk returns a tuple with the RxData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxData

`func (o *DeviceTrafficActivity) SetRxData(v float64)`

SetRxData sets RxData field to given value.

### HasRxData

`func (o *DeviceTrafficActivity) HasRxData() bool`

HasRxData returns a boolean if a field has been set.

### GetTime

`func (o *DeviceTrafficActivity) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DeviceTrafficActivity) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DeviceTrafficActivity) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *DeviceTrafficActivity) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTxData

`func (o *DeviceTrafficActivity) GetTxData() float64`

GetTxData returns the TxData field if non-nil, zero value otherwise.

### GetTxDataOk

`func (o *DeviceTrafficActivity) GetTxDataOk() (*float64, bool)`

GetTxDataOk returns a tuple with the TxData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxData

`func (o *DeviceTrafficActivity) SetTxData(v float64)`

SetTxData sets TxData field to given value.

### HasTxData

`func (o *DeviceTrafficActivity) HasTxData() bool`

HasTxData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


