# OsgLtePinOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstTry** | Pointer to **bool** | Whether the first attempt was successful. | [optional] 
**NewPinCode** | Pointer to **string** | New PIN code | [optional] 
**PinCode** | Pointer to **string** | PIN code. | [optional] 
**PinSetting** | Pointer to [**OsgLtePinSettingOpenApiVO**](OsgLtePinSettingOpenApiVO.md) |  | [optional] 
**PukCode** | Pointer to **string** | PUK code. | [optional] 
**SimCard** | Pointer to **int32** | When the device supports Dual-SIM card, parameter [simCard] should not be null.1: SIM1; 2: SIM2. | [optional] 
**Type** | Pointer to **int32** | Type of change: 1: Changing the PIN setting, 2: Changing the PIN code. | [optional] 

## Methods

### NewOsgLtePinOpenApiVO

`func NewOsgLtePinOpenApiVO() *OsgLtePinOpenApiVO`

NewOsgLtePinOpenApiVO instantiates a new OsgLtePinOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgLtePinOpenApiVOWithDefaults

`func NewOsgLtePinOpenApiVOWithDefaults() *OsgLtePinOpenApiVO`

NewOsgLtePinOpenApiVOWithDefaults instantiates a new OsgLtePinOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstTry

`func (o *OsgLtePinOpenApiVO) GetFirstTry() bool`

GetFirstTry returns the FirstTry field if non-nil, zero value otherwise.

### GetFirstTryOk

`func (o *OsgLtePinOpenApiVO) GetFirstTryOk() (*bool, bool)`

GetFirstTryOk returns a tuple with the FirstTry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTry

`func (o *OsgLtePinOpenApiVO) SetFirstTry(v bool)`

SetFirstTry sets FirstTry field to given value.

### HasFirstTry

`func (o *OsgLtePinOpenApiVO) HasFirstTry() bool`

HasFirstTry returns a boolean if a field has been set.

### GetNewPinCode

`func (o *OsgLtePinOpenApiVO) GetNewPinCode() string`

GetNewPinCode returns the NewPinCode field if non-nil, zero value otherwise.

### GetNewPinCodeOk

`func (o *OsgLtePinOpenApiVO) GetNewPinCodeOk() (*string, bool)`

GetNewPinCodeOk returns a tuple with the NewPinCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPinCode

`func (o *OsgLtePinOpenApiVO) SetNewPinCode(v string)`

SetNewPinCode sets NewPinCode field to given value.

### HasNewPinCode

`func (o *OsgLtePinOpenApiVO) HasNewPinCode() bool`

HasNewPinCode returns a boolean if a field has been set.

### GetPinCode

`func (o *OsgLtePinOpenApiVO) GetPinCode() string`

GetPinCode returns the PinCode field if non-nil, zero value otherwise.

### GetPinCodeOk

`func (o *OsgLtePinOpenApiVO) GetPinCodeOk() (*string, bool)`

GetPinCodeOk returns a tuple with the PinCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinCode

`func (o *OsgLtePinOpenApiVO) SetPinCode(v string)`

SetPinCode sets PinCode field to given value.

### HasPinCode

`func (o *OsgLtePinOpenApiVO) HasPinCode() bool`

HasPinCode returns a boolean if a field has been set.

### GetPinSetting

`func (o *OsgLtePinOpenApiVO) GetPinSetting() OsgLtePinSettingOpenApiVO`

GetPinSetting returns the PinSetting field if non-nil, zero value otherwise.

### GetPinSettingOk

`func (o *OsgLtePinOpenApiVO) GetPinSettingOk() (*OsgLtePinSettingOpenApiVO, bool)`

GetPinSettingOk returns a tuple with the PinSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinSetting

`func (o *OsgLtePinOpenApiVO) SetPinSetting(v OsgLtePinSettingOpenApiVO)`

SetPinSetting sets PinSetting field to given value.

### HasPinSetting

`func (o *OsgLtePinOpenApiVO) HasPinSetting() bool`

HasPinSetting returns a boolean if a field has been set.

### GetPukCode

`func (o *OsgLtePinOpenApiVO) GetPukCode() string`

GetPukCode returns the PukCode field if non-nil, zero value otherwise.

### GetPukCodeOk

`func (o *OsgLtePinOpenApiVO) GetPukCodeOk() (*string, bool)`

GetPukCodeOk returns a tuple with the PukCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPukCode

`func (o *OsgLtePinOpenApiVO) SetPukCode(v string)`

SetPukCode sets PukCode field to given value.

### HasPukCode

`func (o *OsgLtePinOpenApiVO) HasPukCode() bool`

HasPukCode returns a boolean if a field has been set.

### GetSimCard

`func (o *OsgLtePinOpenApiVO) GetSimCard() int32`

GetSimCard returns the SimCard field if non-nil, zero value otherwise.

### GetSimCardOk

`func (o *OsgLtePinOpenApiVO) GetSimCardOk() (*int32, bool)`

GetSimCardOk returns a tuple with the SimCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCard

`func (o *OsgLtePinOpenApiVO) SetSimCard(v int32)`

SetSimCard sets SimCard field to given value.

### HasSimCard

`func (o *OsgLtePinOpenApiVO) HasSimCard() bool`

HasSimCard returns a boolean if a field has been set.

### GetType

`func (o *OsgLtePinOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OsgLtePinOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OsgLtePinOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *OsgLtePinOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


