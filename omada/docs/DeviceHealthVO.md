# DeviceHealthVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Good** | Pointer to **int32** | Number of devices with good health status | [optional] 
**NoData** | Pointer to **int32** | Number of devices with no health data | [optional] 
**Poor** | Pointer to **int32** | Number of devices with poor health status | [optional] 
**Time** | Pointer to **int64** | Timestamp | [optional] 
**Total** | Pointer to **int32** | Total number of devices | [optional] 

## Methods

### NewDeviceHealthVO

`func NewDeviceHealthVO() *DeviceHealthVO`

NewDeviceHealthVO instantiates a new DeviceHealthVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceHealthVOWithDefaults

`func NewDeviceHealthVOWithDefaults() *DeviceHealthVO`

NewDeviceHealthVOWithDefaults instantiates a new DeviceHealthVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGood

`func (o *DeviceHealthVO) GetGood() int32`

GetGood returns the Good field if non-nil, zero value otherwise.

### GetGoodOk

`func (o *DeviceHealthVO) GetGoodOk() (*int32, bool)`

GetGoodOk returns a tuple with the Good field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGood

`func (o *DeviceHealthVO) SetGood(v int32)`

SetGood sets Good field to given value.

### HasGood

`func (o *DeviceHealthVO) HasGood() bool`

HasGood returns a boolean if a field has been set.

### GetNoData

`func (o *DeviceHealthVO) GetNoData() int32`

GetNoData returns the NoData field if non-nil, zero value otherwise.

### GetNoDataOk

`func (o *DeviceHealthVO) GetNoDataOk() (*int32, bool)`

GetNoDataOk returns a tuple with the NoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoData

`func (o *DeviceHealthVO) SetNoData(v int32)`

SetNoData sets NoData field to given value.

### HasNoData

`func (o *DeviceHealthVO) HasNoData() bool`

HasNoData returns a boolean if a field has been set.

### GetPoor

`func (o *DeviceHealthVO) GetPoor() int32`

GetPoor returns the Poor field if non-nil, zero value otherwise.

### GetPoorOk

`func (o *DeviceHealthVO) GetPoorOk() (*int32, bool)`

GetPoorOk returns a tuple with the Poor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoor

`func (o *DeviceHealthVO) SetPoor(v int32)`

SetPoor sets Poor field to given value.

### HasPoor

`func (o *DeviceHealthVO) HasPoor() bool`

HasPoor returns a boolean if a field has been set.

### GetTime

`func (o *DeviceHealthVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DeviceHealthVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DeviceHealthVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *DeviceHealthVO) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTotal

`func (o *DeviceHealthVO) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DeviceHealthVO) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DeviceHealthVO) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DeviceHealthVO) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


