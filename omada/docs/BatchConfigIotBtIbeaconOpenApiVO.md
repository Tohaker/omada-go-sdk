# BatchConfigIotBtIbeaconOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvInterval** | Pointer to **int32** | Advertising interval in milliseconds.(500 by default) | [optional] 
**Enable** | Pointer to **bool** | Whether to enable the Bluetooth Advertising setting.(Disable by default) | [optional] 
**Major** | Pointer to **string** | The major value of adverting ibeacon packet, indicating a larger group. | [optional] 
**MeasurePower** | Pointer to **int32** | RSSI Calibration Value. The parameter [measurePower] is used to input the RSSI measured at a 1-meter distance from the device, enabling positioning functionality.(-65 by default) | [optional] 
**Minor** | Pointer to **string** | The minor value of adverting ibeacon packet, indicating a smaller group. | [optional] 
**TransmitPower** | Pointer to **int32** | Broadcast transmission power.&lt;br /&gt;The parameter [transmitPower] should be a value as follows:[-20, -18, -15, -12, -10, -9, -6, -5, -3, 0, 1, 2, 3, 4, 5, 14, 15, 16, 17, 18, 19, 20].(0 by default) | [optional] 
**Uuid** | Pointer to **string** | The UUID (Universally Unique Identifier) of the advertising ibeacon packet. | [optional] 

## Methods

### NewBatchConfigIotBtIbeaconOpenApiVO

`func NewBatchConfigIotBtIbeaconOpenApiVO() *BatchConfigIotBtIbeaconOpenApiVO`

NewBatchConfigIotBtIbeaconOpenApiVO instantiates a new BatchConfigIotBtIbeaconOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchConfigIotBtIbeaconOpenApiVOWithDefaults

`func NewBatchConfigIotBtIbeaconOpenApiVOWithDefaults() *BatchConfigIotBtIbeaconOpenApiVO`

NewBatchConfigIotBtIbeaconOpenApiVOWithDefaults instantiates a new BatchConfigIotBtIbeaconOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvInterval

`func (o *BatchConfigIotBtIbeaconOpenApiVO) GetAdvInterval() int32`

GetAdvInterval returns the AdvInterval field if non-nil, zero value otherwise.

### GetAdvIntervalOk

`func (o *BatchConfigIotBtIbeaconOpenApiVO) GetAdvIntervalOk() (*int32, bool)`

GetAdvIntervalOk returns a tuple with the AdvInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvInterval

`func (o *BatchConfigIotBtIbeaconOpenApiVO) SetAdvInterval(v int32)`

SetAdvInterval sets AdvInterval field to given value.

### HasAdvInterval

`func (o *BatchConfigIotBtIbeaconOpenApiVO) HasAdvInterval() bool`

HasAdvInterval returns a boolean if a field has been set.

### GetEnable

`func (o *BatchConfigIotBtIbeaconOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *BatchConfigIotBtIbeaconOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *BatchConfigIotBtIbeaconOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *BatchConfigIotBtIbeaconOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetMajor

`func (o *BatchConfigIotBtIbeaconOpenApiVO) GetMajor() string`

GetMajor returns the Major field if non-nil, zero value otherwise.

### GetMajorOk

`func (o *BatchConfigIotBtIbeaconOpenApiVO) GetMajorOk() (*string, bool)`

GetMajorOk returns a tuple with the Major field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajor

`func (o *BatchConfigIotBtIbeaconOpenApiVO) SetMajor(v string)`

SetMajor sets Major field to given value.

### HasMajor

`func (o *BatchConfigIotBtIbeaconOpenApiVO) HasMajor() bool`

HasMajor returns a boolean if a field has been set.

### GetMeasurePower

`func (o *BatchConfigIotBtIbeaconOpenApiVO) GetMeasurePower() int32`

GetMeasurePower returns the MeasurePower field if non-nil, zero value otherwise.

### GetMeasurePowerOk

`func (o *BatchConfigIotBtIbeaconOpenApiVO) GetMeasurePowerOk() (*int32, bool)`

GetMeasurePowerOk returns a tuple with the MeasurePower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurePower

`func (o *BatchConfigIotBtIbeaconOpenApiVO) SetMeasurePower(v int32)`

SetMeasurePower sets MeasurePower field to given value.

### HasMeasurePower

`func (o *BatchConfigIotBtIbeaconOpenApiVO) HasMeasurePower() bool`

HasMeasurePower returns a boolean if a field has been set.

### GetMinor

`func (o *BatchConfigIotBtIbeaconOpenApiVO) GetMinor() string`

GetMinor returns the Minor field if non-nil, zero value otherwise.

### GetMinorOk

`func (o *BatchConfigIotBtIbeaconOpenApiVO) GetMinorOk() (*string, bool)`

GetMinorOk returns a tuple with the Minor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinor

`func (o *BatchConfigIotBtIbeaconOpenApiVO) SetMinor(v string)`

SetMinor sets Minor field to given value.

### HasMinor

`func (o *BatchConfigIotBtIbeaconOpenApiVO) HasMinor() bool`

HasMinor returns a boolean if a field has been set.

### GetTransmitPower

`func (o *BatchConfigIotBtIbeaconOpenApiVO) GetTransmitPower() int32`

GetTransmitPower returns the TransmitPower field if non-nil, zero value otherwise.

### GetTransmitPowerOk

`func (o *BatchConfigIotBtIbeaconOpenApiVO) GetTransmitPowerOk() (*int32, bool)`

GetTransmitPowerOk returns a tuple with the TransmitPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitPower

`func (o *BatchConfigIotBtIbeaconOpenApiVO) SetTransmitPower(v int32)`

SetTransmitPower sets TransmitPower field to given value.

### HasTransmitPower

`func (o *BatchConfigIotBtIbeaconOpenApiVO) HasTransmitPower() bool`

HasTransmitPower returns a boolean if a field has been set.

### GetUuid

`func (o *BatchConfigIotBtIbeaconOpenApiVO) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *BatchConfigIotBtIbeaconOpenApiVO) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *BatchConfigIotBtIbeaconOpenApiVO) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *BatchConfigIotBtIbeaconOpenApiVO) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


