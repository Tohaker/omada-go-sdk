# ConfigIotBtIbeaconV2OpenApiVO

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
**Uuid** | **string** | The UUID (Universally Unique Identifier) of the advertising ibeacon packet. | 

## Methods

### NewConfigIotBtIbeaconV2OpenApiVO

`func NewConfigIotBtIbeaconV2OpenApiVO(macList []string, major string, minor string, name string, uuid string, ) *ConfigIotBtIbeaconV2OpenApiVO`

NewConfigIotBtIbeaconV2OpenApiVO instantiates a new ConfigIotBtIbeaconV2OpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigIotBtIbeaconV2OpenApiVOWithDefaults

`func NewConfigIotBtIbeaconV2OpenApiVOWithDefaults() *ConfigIotBtIbeaconV2OpenApiVO`

NewConfigIotBtIbeaconV2OpenApiVOWithDefaults instantiates a new ConfigIotBtIbeaconV2OpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvInterval

`func (o *ConfigIotBtIbeaconV2OpenApiVO) GetAdvInterval() int32`

GetAdvInterval returns the AdvInterval field if non-nil, zero value otherwise.

### GetAdvIntervalOk

`func (o *ConfigIotBtIbeaconV2OpenApiVO) GetAdvIntervalOk() (*int32, bool)`

GetAdvIntervalOk returns a tuple with the AdvInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvInterval

`func (o *ConfigIotBtIbeaconV2OpenApiVO) SetAdvInterval(v int32)`

SetAdvInterval sets AdvInterval field to given value.

### HasAdvInterval

`func (o *ConfigIotBtIbeaconV2OpenApiVO) HasAdvInterval() bool`

HasAdvInterval returns a boolean if a field has been set.

### GetEnable

`func (o *ConfigIotBtIbeaconV2OpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ConfigIotBtIbeaconV2OpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ConfigIotBtIbeaconV2OpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ConfigIotBtIbeaconV2OpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetMacList

`func (o *ConfigIotBtIbeaconV2OpenApiVO) GetMacList() []string`

GetMacList returns the MacList field if non-nil, zero value otherwise.

### GetMacListOk

`func (o *ConfigIotBtIbeaconV2OpenApiVO) GetMacListOk() (*[]string, bool)`

GetMacListOk returns a tuple with the MacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacList

`func (o *ConfigIotBtIbeaconV2OpenApiVO) SetMacList(v []string)`

SetMacList sets MacList field to given value.


### GetMajor

`func (o *ConfigIotBtIbeaconV2OpenApiVO) GetMajor() string`

GetMajor returns the Major field if non-nil, zero value otherwise.

### GetMajorOk

`func (o *ConfigIotBtIbeaconV2OpenApiVO) GetMajorOk() (*string, bool)`

GetMajorOk returns a tuple with the Major field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajor

`func (o *ConfigIotBtIbeaconV2OpenApiVO) SetMajor(v string)`

SetMajor sets Major field to given value.


### GetMeasurePower

`func (o *ConfigIotBtIbeaconV2OpenApiVO) GetMeasurePower() int32`

GetMeasurePower returns the MeasurePower field if non-nil, zero value otherwise.

### GetMeasurePowerOk

`func (o *ConfigIotBtIbeaconV2OpenApiVO) GetMeasurePowerOk() (*int32, bool)`

GetMeasurePowerOk returns a tuple with the MeasurePower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurePower

`func (o *ConfigIotBtIbeaconV2OpenApiVO) SetMeasurePower(v int32)`

SetMeasurePower sets MeasurePower field to given value.

### HasMeasurePower

`func (o *ConfigIotBtIbeaconV2OpenApiVO) HasMeasurePower() bool`

HasMeasurePower returns a boolean if a field has been set.

### GetMinor

`func (o *ConfigIotBtIbeaconV2OpenApiVO) GetMinor() string`

GetMinor returns the Minor field if non-nil, zero value otherwise.

### GetMinorOk

`func (o *ConfigIotBtIbeaconV2OpenApiVO) GetMinorOk() (*string, bool)`

GetMinorOk returns a tuple with the Minor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinor

`func (o *ConfigIotBtIbeaconV2OpenApiVO) SetMinor(v string)`

SetMinor sets Minor field to given value.


### GetName

`func (o *ConfigIotBtIbeaconV2OpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigIotBtIbeaconV2OpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigIotBtIbeaconV2OpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetUuid

`func (o *ConfigIotBtIbeaconV2OpenApiVO) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ConfigIotBtIbeaconV2OpenApiVO) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ConfigIotBtIbeaconV2OpenApiVO) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


