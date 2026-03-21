# IotRadioSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgingTime** | Pointer to **int32** | The system automatically removes a device&#39;s registry entry if no data reports are received within a predefined aging period.&lt;br/&gt;When format &#x3D; 0, The parameter aging time should be within the range of 30-86400.&lt;br/&gt;When format &#x3D; 1, The parameter aging time should be within the range of 1-1440.&lt;br/&gt;When format &#x3D; 2, The parameter aging time should be within the range of 1-24.&lt;br/&gt; | [optional] 
**ConsoleMode** | Pointer to **int32** | Bluetooth consle mode should be a value as follow:[0:auto, 1:on, 2:off]. | [optional] 
**Enable** | **bool** | Bluetooth enable. | 
**Format** | Pointer to **int32** | The parameter [format] should be a value as follows: [0:second 1:minute; 2:hour] | [optional] 
**Passcode** | Pointer to **string** | Bluetooth consle passcode. | [optional] 
**Resource** | Pointer to **int32** |  | [optional] 
**TransmitPower** | Pointer to **int32** | Broadcast transmission power.&lt;br /&gt;The parameter [transmitPower] should be a value as follows:[-20, -18, -15, -12, -10, -9, -6, -5, -3, 0, 1, 2, 3, 4, 5, 14, 15, 16, 17, 18, 19, 20].(0 by default) | [optional] 

## Methods

### NewIotRadioSettingOpenApiVO

`func NewIotRadioSettingOpenApiVO(enable bool, ) *IotRadioSettingOpenApiVO`

NewIotRadioSettingOpenApiVO instantiates a new IotRadioSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIotRadioSettingOpenApiVOWithDefaults

`func NewIotRadioSettingOpenApiVOWithDefaults() *IotRadioSettingOpenApiVO`

NewIotRadioSettingOpenApiVOWithDefaults instantiates a new IotRadioSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgingTime

`func (o *IotRadioSettingOpenApiVO) GetAgingTime() int32`

GetAgingTime returns the AgingTime field if non-nil, zero value otherwise.

### GetAgingTimeOk

`func (o *IotRadioSettingOpenApiVO) GetAgingTimeOk() (*int32, bool)`

GetAgingTimeOk returns a tuple with the AgingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgingTime

`func (o *IotRadioSettingOpenApiVO) SetAgingTime(v int32)`

SetAgingTime sets AgingTime field to given value.

### HasAgingTime

`func (o *IotRadioSettingOpenApiVO) HasAgingTime() bool`

HasAgingTime returns a boolean if a field has been set.

### GetConsoleMode

`func (o *IotRadioSettingOpenApiVO) GetConsoleMode() int32`

GetConsoleMode returns the ConsoleMode field if non-nil, zero value otherwise.

### GetConsoleModeOk

`func (o *IotRadioSettingOpenApiVO) GetConsoleModeOk() (*int32, bool)`

GetConsoleModeOk returns a tuple with the ConsoleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleMode

`func (o *IotRadioSettingOpenApiVO) SetConsoleMode(v int32)`

SetConsoleMode sets ConsoleMode field to given value.

### HasConsoleMode

`func (o *IotRadioSettingOpenApiVO) HasConsoleMode() bool`

HasConsoleMode returns a boolean if a field has been set.

### GetEnable

`func (o *IotRadioSettingOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IotRadioSettingOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IotRadioSettingOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetFormat

`func (o *IotRadioSettingOpenApiVO) GetFormat() int32`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *IotRadioSettingOpenApiVO) GetFormatOk() (*int32, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *IotRadioSettingOpenApiVO) SetFormat(v int32)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *IotRadioSettingOpenApiVO) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetPasscode

`func (o *IotRadioSettingOpenApiVO) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *IotRadioSettingOpenApiVO) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *IotRadioSettingOpenApiVO) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.

### HasPasscode

`func (o *IotRadioSettingOpenApiVO) HasPasscode() bool`

HasPasscode returns a boolean if a field has been set.

### GetResource

`func (o *IotRadioSettingOpenApiVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *IotRadioSettingOpenApiVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *IotRadioSettingOpenApiVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *IotRadioSettingOpenApiVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetTransmitPower

`func (o *IotRadioSettingOpenApiVO) GetTransmitPower() int32`

GetTransmitPower returns the TransmitPower field if non-nil, zero value otherwise.

### GetTransmitPowerOk

`func (o *IotRadioSettingOpenApiVO) GetTransmitPowerOk() (*int32, bool)`

GetTransmitPowerOk returns a tuple with the TransmitPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitPower

`func (o *IotRadioSettingOpenApiVO) SetTransmitPower(v int32)`

SetTransmitPower sets TransmitPower field to given value.

### HasTransmitPower

`func (o *IotRadioSettingOpenApiVO) HasTransmitPower() bool`

HasTransmitPower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


