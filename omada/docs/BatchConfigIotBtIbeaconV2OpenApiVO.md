# BatchConfigIotBtIbeaconV2OpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvInterval** | Pointer to **int32** | Advertising interval in milliseconds.(500 by default) | [optional] 
**Enable** | Pointer to **bool** | Whether to enable the Bluetooth Advertising setting.(Disable by default) | [optional] 
**Major** | Pointer to **string** | The major value of adverting ibeacon packet, indicating a larger group. | [optional] 
**MeasurePower** | Pointer to **int32** | RSSI Calibration Value. The parameter [measurePower] is used to input the RSSI measured at a 1-meter distance from the device, enabling positioning functionality.(-65 by default) | [optional] 
**Minor** | Pointer to **string** | The minor value of adverting ibeacon packet, indicating a smaller group. | [optional] 
**Uuid** | Pointer to **string** | The UUID (Universally Unique Identifier) of the advertising ibeacon packet. | [optional] 

## Methods

### NewBatchConfigIotBtIbeaconV2OpenApiVO

`func NewBatchConfigIotBtIbeaconV2OpenApiVO() *BatchConfigIotBtIbeaconV2OpenApiVO`

NewBatchConfigIotBtIbeaconV2OpenApiVO instantiates a new BatchConfigIotBtIbeaconV2OpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchConfigIotBtIbeaconV2OpenApiVOWithDefaults

`func NewBatchConfigIotBtIbeaconV2OpenApiVOWithDefaults() *BatchConfigIotBtIbeaconV2OpenApiVO`

NewBatchConfigIotBtIbeaconV2OpenApiVOWithDefaults instantiates a new BatchConfigIotBtIbeaconV2OpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvInterval

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) GetAdvInterval() int32`

GetAdvInterval returns the AdvInterval field if non-nil, zero value otherwise.

### GetAdvIntervalOk

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) GetAdvIntervalOk() (*int32, bool)`

GetAdvIntervalOk returns a tuple with the AdvInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvInterval

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) SetAdvInterval(v int32)`

SetAdvInterval sets AdvInterval field to given value.

### HasAdvInterval

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) HasAdvInterval() bool`

HasAdvInterval returns a boolean if a field has been set.

### GetEnable

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetMajor

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) GetMajor() string`

GetMajor returns the Major field if non-nil, zero value otherwise.

### GetMajorOk

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) GetMajorOk() (*string, bool)`

GetMajorOk returns a tuple with the Major field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajor

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) SetMajor(v string)`

SetMajor sets Major field to given value.

### HasMajor

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) HasMajor() bool`

HasMajor returns a boolean if a field has been set.

### GetMeasurePower

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) GetMeasurePower() int32`

GetMeasurePower returns the MeasurePower field if non-nil, zero value otherwise.

### GetMeasurePowerOk

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) GetMeasurePowerOk() (*int32, bool)`

GetMeasurePowerOk returns a tuple with the MeasurePower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurePower

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) SetMeasurePower(v int32)`

SetMeasurePower sets MeasurePower field to given value.

### HasMeasurePower

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) HasMeasurePower() bool`

HasMeasurePower returns a boolean if a field has been set.

### GetMinor

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) GetMinor() string`

GetMinor returns the Minor field if non-nil, zero value otherwise.

### GetMinorOk

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) GetMinorOk() (*string, bool)`

GetMinorOk returns a tuple with the Minor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinor

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) SetMinor(v string)`

SetMinor sets Minor field to given value.

### HasMinor

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) HasMinor() bool`

HasMinor returns a boolean if a field has been set.

### GetUuid

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *BatchConfigIotBtIbeaconV2OpenApiVO) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


