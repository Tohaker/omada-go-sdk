# IotBtIbeaconOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvInterval** | Pointer to **int32** | Advertising interval. | [optional] 
**BoundDeviceNum** | Pointer to **int32** | The quantity of devices bound to this configuration. | [optional] 
**BuildIn** | Pointer to **int32** | Whether it is a built-in configuration: [0:true; 1:false]. | [optional] 
**Enable** | Pointer to **bool** | Whether to enable the Bluetooth Advertising setting. | [optional] 
**Id** | Pointer to **string** | The Bluetooth Advertising entry ID. | [optional] 
**MacList** | Pointer to **[]string** | List of device MAC addresses bound to this configuration. | [optional] 
**Major** | Pointer to **string** | The major value of adverting ibeacon packet, indicating a larger group. | [optional] 
**MeasurePower** | Pointer to **int32** | RSSI Calibration Value. The parameter [measurePower] is used to input the RSSI measured at a 1-meter distance from the device, enabling positioning functionality. | [optional] 
**Minor** | Pointer to **string** | The minor value of adverting ibeacon packet, indicating a smaller group. | [optional] 
**Name** | **string** | The Bluetooth Advertising seting name. | 
**SelectMacInfo** | Pointer to [**[]ApBtDetailOpenApiVO**](ApBtDetailOpenApiVO.md) | Detailed information about the devices bound to this configuration. | [optional] 
**TransmitPower** | Pointer to **int32** | Broadcast transmission power.&lt;br /&gt;The parameter [transmitPower] should be a value as follows:[-20, -18, -15, -12, -10, -9, -6, -5, -3, 0, 1, 2, 3, 4, 5, 14, 15, 16, 17, 18, 19, 20] | [optional] 
**Uuid** | Pointer to **string** | The UUID (Universally Unique Identifier) of the advertising ibeacon packet. | [optional] 

## Methods

### NewIotBtIbeaconOpenApiVO

`func NewIotBtIbeaconOpenApiVO(name string, ) *IotBtIbeaconOpenApiVO`

NewIotBtIbeaconOpenApiVO instantiates a new IotBtIbeaconOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIotBtIbeaconOpenApiVOWithDefaults

`func NewIotBtIbeaconOpenApiVOWithDefaults() *IotBtIbeaconOpenApiVO`

NewIotBtIbeaconOpenApiVOWithDefaults instantiates a new IotBtIbeaconOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvInterval

`func (o *IotBtIbeaconOpenApiVO) GetAdvInterval() int32`

GetAdvInterval returns the AdvInterval field if non-nil, zero value otherwise.

### GetAdvIntervalOk

`func (o *IotBtIbeaconOpenApiVO) GetAdvIntervalOk() (*int32, bool)`

GetAdvIntervalOk returns a tuple with the AdvInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvInterval

`func (o *IotBtIbeaconOpenApiVO) SetAdvInterval(v int32)`

SetAdvInterval sets AdvInterval field to given value.

### HasAdvInterval

`func (o *IotBtIbeaconOpenApiVO) HasAdvInterval() bool`

HasAdvInterval returns a boolean if a field has been set.

### GetBoundDeviceNum

`func (o *IotBtIbeaconOpenApiVO) GetBoundDeviceNum() int32`

GetBoundDeviceNum returns the BoundDeviceNum field if non-nil, zero value otherwise.

### GetBoundDeviceNumOk

`func (o *IotBtIbeaconOpenApiVO) GetBoundDeviceNumOk() (*int32, bool)`

GetBoundDeviceNumOk returns a tuple with the BoundDeviceNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundDeviceNum

`func (o *IotBtIbeaconOpenApiVO) SetBoundDeviceNum(v int32)`

SetBoundDeviceNum sets BoundDeviceNum field to given value.

### HasBoundDeviceNum

`func (o *IotBtIbeaconOpenApiVO) HasBoundDeviceNum() bool`

HasBoundDeviceNum returns a boolean if a field has been set.

### GetBuildIn

`func (o *IotBtIbeaconOpenApiVO) GetBuildIn() int32`

GetBuildIn returns the BuildIn field if non-nil, zero value otherwise.

### GetBuildInOk

`func (o *IotBtIbeaconOpenApiVO) GetBuildInOk() (*int32, bool)`

GetBuildInOk returns a tuple with the BuildIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildIn

`func (o *IotBtIbeaconOpenApiVO) SetBuildIn(v int32)`

SetBuildIn sets BuildIn field to given value.

### HasBuildIn

`func (o *IotBtIbeaconOpenApiVO) HasBuildIn() bool`

HasBuildIn returns a boolean if a field has been set.

### GetEnable

`func (o *IotBtIbeaconOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IotBtIbeaconOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IotBtIbeaconOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *IotBtIbeaconOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetId

`func (o *IotBtIbeaconOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IotBtIbeaconOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IotBtIbeaconOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IotBtIbeaconOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMacList

`func (o *IotBtIbeaconOpenApiVO) GetMacList() []string`

GetMacList returns the MacList field if non-nil, zero value otherwise.

### GetMacListOk

`func (o *IotBtIbeaconOpenApiVO) GetMacListOk() (*[]string, bool)`

GetMacListOk returns a tuple with the MacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacList

`func (o *IotBtIbeaconOpenApiVO) SetMacList(v []string)`

SetMacList sets MacList field to given value.

### HasMacList

`func (o *IotBtIbeaconOpenApiVO) HasMacList() bool`

HasMacList returns a boolean if a field has been set.

### GetMajor

`func (o *IotBtIbeaconOpenApiVO) GetMajor() string`

GetMajor returns the Major field if non-nil, zero value otherwise.

### GetMajorOk

`func (o *IotBtIbeaconOpenApiVO) GetMajorOk() (*string, bool)`

GetMajorOk returns a tuple with the Major field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajor

`func (o *IotBtIbeaconOpenApiVO) SetMajor(v string)`

SetMajor sets Major field to given value.

### HasMajor

`func (o *IotBtIbeaconOpenApiVO) HasMajor() bool`

HasMajor returns a boolean if a field has been set.

### GetMeasurePower

`func (o *IotBtIbeaconOpenApiVO) GetMeasurePower() int32`

GetMeasurePower returns the MeasurePower field if non-nil, zero value otherwise.

### GetMeasurePowerOk

`func (o *IotBtIbeaconOpenApiVO) GetMeasurePowerOk() (*int32, bool)`

GetMeasurePowerOk returns a tuple with the MeasurePower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurePower

`func (o *IotBtIbeaconOpenApiVO) SetMeasurePower(v int32)`

SetMeasurePower sets MeasurePower field to given value.

### HasMeasurePower

`func (o *IotBtIbeaconOpenApiVO) HasMeasurePower() bool`

HasMeasurePower returns a boolean if a field has been set.

### GetMinor

`func (o *IotBtIbeaconOpenApiVO) GetMinor() string`

GetMinor returns the Minor field if non-nil, zero value otherwise.

### GetMinorOk

`func (o *IotBtIbeaconOpenApiVO) GetMinorOk() (*string, bool)`

GetMinorOk returns a tuple with the Minor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinor

`func (o *IotBtIbeaconOpenApiVO) SetMinor(v string)`

SetMinor sets Minor field to given value.

### HasMinor

`func (o *IotBtIbeaconOpenApiVO) HasMinor() bool`

HasMinor returns a boolean if a field has been set.

### GetName

`func (o *IotBtIbeaconOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IotBtIbeaconOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IotBtIbeaconOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetSelectMacInfo

`func (o *IotBtIbeaconOpenApiVO) GetSelectMacInfo() []ApBtDetailOpenApiVO`

GetSelectMacInfo returns the SelectMacInfo field if non-nil, zero value otherwise.

### GetSelectMacInfoOk

`func (o *IotBtIbeaconOpenApiVO) GetSelectMacInfoOk() (*[]ApBtDetailOpenApiVO, bool)`

GetSelectMacInfoOk returns a tuple with the SelectMacInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectMacInfo

`func (o *IotBtIbeaconOpenApiVO) SetSelectMacInfo(v []ApBtDetailOpenApiVO)`

SetSelectMacInfo sets SelectMacInfo field to given value.

### HasSelectMacInfo

`func (o *IotBtIbeaconOpenApiVO) HasSelectMacInfo() bool`

HasSelectMacInfo returns a boolean if a field has been set.

### GetTransmitPower

`func (o *IotBtIbeaconOpenApiVO) GetTransmitPower() int32`

GetTransmitPower returns the TransmitPower field if non-nil, zero value otherwise.

### GetTransmitPowerOk

`func (o *IotBtIbeaconOpenApiVO) GetTransmitPowerOk() (*int32, bool)`

GetTransmitPowerOk returns a tuple with the TransmitPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitPower

`func (o *IotBtIbeaconOpenApiVO) SetTransmitPower(v int32)`

SetTransmitPower sets TransmitPower field to given value.

### HasTransmitPower

`func (o *IotBtIbeaconOpenApiVO) HasTransmitPower() bool`

HasTransmitPower returns a boolean if a field has been set.

### GetUuid

`func (o *IotBtIbeaconOpenApiVO) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *IotBtIbeaconOpenApiVO) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *IotBtIbeaconOpenApiVO) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *IotBtIbeaconOpenApiVO) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


