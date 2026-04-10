# DeviceLocationDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address of the device. | [optional] 
**Latitude** | Pointer to **float64** | Latitude of the device. | [optional] 
**Longitude** | Pointer to **float64** | Longitude of the device. | [optional] 
**Timestamp** | Pointer to **int64** | Last GPS fix timestamp | [optional] 

## Methods

### NewDeviceLocationDetailVO

`func NewDeviceLocationDetailVO() *DeviceLocationDetailVO`

NewDeviceLocationDetailVO instantiates a new DeviceLocationDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceLocationDetailVOWithDefaults

`func NewDeviceLocationDetailVOWithDefaults() *DeviceLocationDetailVO`

NewDeviceLocationDetailVOWithDefaults instantiates a new DeviceLocationDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DeviceLocationDetailVO) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DeviceLocationDetailVO) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DeviceLocationDetailVO) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DeviceLocationDetailVO) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLatitude

`func (o *DeviceLocationDetailVO) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *DeviceLocationDetailVO) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *DeviceLocationDetailVO) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *DeviceLocationDetailVO) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *DeviceLocationDetailVO) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *DeviceLocationDetailVO) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *DeviceLocationDetailVO) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *DeviceLocationDetailVO) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetTimestamp

`func (o *DeviceLocationDetailVO) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DeviceLocationDetailVO) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DeviceLocationDetailVO) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DeviceLocationDetailVO) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


