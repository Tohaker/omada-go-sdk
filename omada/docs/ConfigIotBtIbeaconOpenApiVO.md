# ConfigIotBtIbeaconOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvInterval** | Pointer to **int32** | Advertising interval in milliseconds.(500 by default) | [optional] 
**Enable** | Pointer to **bool** | Whether to enable the Bluetooth Advertising setting.(Disable by default) | [optional] 
**MacList** | **[]string** | List of device MAC addresses bound to this configuration. | 
**Major** | **string** | The major value of adverting ibeacon packet, indicating a larger group. | 
**MeasurePower** | Pointer to **int32** | RSSI Calibration Value. The parameter [measurePower] is used to input the RSSI measured at a 1-meter distance from the device, enabling positioning functionality.(-65 by default) | [optional] 
**Minor** | **string** | The minor value of adverting ibeacon packet, indicating a smaller group. | 
**Name** | **string** | The Bluetooth Advertising seting name. The name of built-in entry are not allowed to be modified. | 
**TransmitPower** | Pointer to **int32** | Broadcast transmission power.&lt;br /&gt;The parameter [transmitPower] should be a value as follows:[-20, -18, -15, -12, -10, -9, -6, -5, -3, 0, 1, 2, 3, 4, 5, 14, 15, 16, 17, 18, 19, 20].(0 by default) | [optional] 
**Uuid** | **string** | The UUID (Universally Unique Identifier) of the advertising ibeacon packet. | 

## Methods

### NewConfigIotBtIbeaconOpenApiVO

`func NewConfigIotBtIbeaconOpenApiVO(macList []string, major string, minor string, name string, uuid string, ) *ConfigIotBtIbeaconOpenApiVO`

NewConfigIotBtIbeaconOpenApiVO instantiates a new ConfigIotBtIbeaconOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigIotBtIbeaconOpenApiVOWithDefaults

`func NewConfigIotBtIbeaconOpenApiVOWithDefaults() *ConfigIotBtIbeaconOpenApiVO`

NewConfigIotBtIbeaconOpenApiVOWithDefaults instantiates a new ConfigIotBtIbeaconOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvInterval

`func (o *ConfigIotBtIbeaconOpenApiVO) GetAdvInterval() int32`

GetAdvInterval returns the AdvInterval field if non-nil, zero value otherwise.

### GetAdvIntervalOk

`func (o *ConfigIotBtIbeaconOpenApiVO) GetAdvIntervalOk() (*int32, bool)`

GetAdvIntervalOk returns a tuple with the AdvInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvInterval

`func (o *ConfigIotBtIbeaconOpenApiVO) SetAdvInterval(v int32)`

SetAdvInterval sets AdvInterval field to given value.

### HasAdvInterval

`func (o *ConfigIotBtIbeaconOpenApiVO) HasAdvInterval() bool`

HasAdvInterval returns a boolean if a field has been set.

### GetEnable

`func (o *ConfigIotBtIbeaconOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ConfigIotBtIbeaconOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ConfigIotBtIbeaconOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ConfigIotBtIbeaconOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetMacList

`func (o *ConfigIotBtIbeaconOpenApiVO) GetMacList() []string`

GetMacList returns the MacList field if non-nil, zero value otherwise.

### GetMacListOk

`func (o *ConfigIotBtIbeaconOpenApiVO) GetMacListOk() (*[]string, bool)`

GetMacListOk returns a tuple with the MacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacList

`func (o *ConfigIotBtIbeaconOpenApiVO) SetMacList(v []string)`

SetMacList sets MacList field to given value.


### GetMajor

`func (o *ConfigIotBtIbeaconOpenApiVO) GetMajor() string`

GetMajor returns the Major field if non-nil, zero value otherwise.

### GetMajorOk

`func (o *ConfigIotBtIbeaconOpenApiVO) GetMajorOk() (*string, bool)`

GetMajorOk returns a tuple with the Major field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajor

`func (o *ConfigIotBtIbeaconOpenApiVO) SetMajor(v string)`

SetMajor sets Major field to given value.


### GetMeasurePower

`func (o *ConfigIotBtIbeaconOpenApiVO) GetMeasurePower() int32`

GetMeasurePower returns the MeasurePower field if non-nil, zero value otherwise.

### GetMeasurePowerOk

`func (o *ConfigIotBtIbeaconOpenApiVO) GetMeasurePowerOk() (*int32, bool)`

GetMeasurePowerOk returns a tuple with the MeasurePower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurePower

`func (o *ConfigIotBtIbeaconOpenApiVO) SetMeasurePower(v int32)`

SetMeasurePower sets MeasurePower field to given value.

### HasMeasurePower

`func (o *ConfigIotBtIbeaconOpenApiVO) HasMeasurePower() bool`

HasMeasurePower returns a boolean if a field has been set.

### GetMinor

`func (o *ConfigIotBtIbeaconOpenApiVO) GetMinor() string`

GetMinor returns the Minor field if non-nil, zero value otherwise.

### GetMinorOk

`func (o *ConfigIotBtIbeaconOpenApiVO) GetMinorOk() (*string, bool)`

GetMinorOk returns a tuple with the Minor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinor

`func (o *ConfigIotBtIbeaconOpenApiVO) SetMinor(v string)`

SetMinor sets Minor field to given value.


### GetName

`func (o *ConfigIotBtIbeaconOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigIotBtIbeaconOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigIotBtIbeaconOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetTransmitPower

`func (o *ConfigIotBtIbeaconOpenApiVO) GetTransmitPower() int32`

GetTransmitPower returns the TransmitPower field if non-nil, zero value otherwise.

### GetTransmitPowerOk

`func (o *ConfigIotBtIbeaconOpenApiVO) GetTransmitPowerOk() (*int32, bool)`

GetTransmitPowerOk returns a tuple with the TransmitPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitPower

`func (o *ConfigIotBtIbeaconOpenApiVO) SetTransmitPower(v int32)`

SetTransmitPower sets TransmitPower field to given value.

### HasTransmitPower

`func (o *ConfigIotBtIbeaconOpenApiVO) HasTransmitPower() bool`

HasTransmitPower returns a boolean if a field has been set.

### GetUuid

`func (o *ConfigIotBtIbeaconOpenApiVO) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ConfigIotBtIbeaconOpenApiVO) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ConfigIotBtIbeaconOpenApiVO) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


