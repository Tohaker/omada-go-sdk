# ApBridgeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BridgeSsidName** | Pointer to **string** | Bridge SSID name. It should contain 1 to 32 UTF-8 characters. | [optional] 
**BridgeSsidPassword** | Pointer to **string** | Bridge SSID password. It should contain 8-63 printable ASCII characters. | [optional] 
**HwSwitch** | Pointer to **int32** | Bridge DIP Switch config status. 0: disable, 1: enable. | [optional] 
**ParingCode** | Pointer to **int32** | Bridge paring code. | [optional] 
**TdmaConfig** | Pointer to [**ApBridgeTdmaSettingOpenApiVO**](ApBridgeTdmaSettingOpenApiVO.md) |  | [optional] 

## Methods

### NewApBridgeInfo

`func NewApBridgeInfo() *ApBridgeInfo`

NewApBridgeInfo instantiates a new ApBridgeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApBridgeInfoWithDefaults

`func NewApBridgeInfoWithDefaults() *ApBridgeInfo`

NewApBridgeInfoWithDefaults instantiates a new ApBridgeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBridgeSsidName

`func (o *ApBridgeInfo) GetBridgeSsidName() string`

GetBridgeSsidName returns the BridgeSsidName field if non-nil, zero value otherwise.

### GetBridgeSsidNameOk

`func (o *ApBridgeInfo) GetBridgeSsidNameOk() (*string, bool)`

GetBridgeSsidNameOk returns a tuple with the BridgeSsidName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeSsidName

`func (o *ApBridgeInfo) SetBridgeSsidName(v string)`

SetBridgeSsidName sets BridgeSsidName field to given value.

### HasBridgeSsidName

`func (o *ApBridgeInfo) HasBridgeSsidName() bool`

HasBridgeSsidName returns a boolean if a field has been set.

### GetBridgeSsidPassword

`func (o *ApBridgeInfo) GetBridgeSsidPassword() string`

GetBridgeSsidPassword returns the BridgeSsidPassword field if non-nil, zero value otherwise.

### GetBridgeSsidPasswordOk

`func (o *ApBridgeInfo) GetBridgeSsidPasswordOk() (*string, bool)`

GetBridgeSsidPasswordOk returns a tuple with the BridgeSsidPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeSsidPassword

`func (o *ApBridgeInfo) SetBridgeSsidPassword(v string)`

SetBridgeSsidPassword sets BridgeSsidPassword field to given value.

### HasBridgeSsidPassword

`func (o *ApBridgeInfo) HasBridgeSsidPassword() bool`

HasBridgeSsidPassword returns a boolean if a field has been set.

### GetHwSwitch

`func (o *ApBridgeInfo) GetHwSwitch() int32`

GetHwSwitch returns the HwSwitch field if non-nil, zero value otherwise.

### GetHwSwitchOk

`func (o *ApBridgeInfo) GetHwSwitchOk() (*int32, bool)`

GetHwSwitchOk returns a tuple with the HwSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwSwitch

`func (o *ApBridgeInfo) SetHwSwitch(v int32)`

SetHwSwitch sets HwSwitch field to given value.

### HasHwSwitch

`func (o *ApBridgeInfo) HasHwSwitch() bool`

HasHwSwitch returns a boolean if a field has been set.

### GetParingCode

`func (o *ApBridgeInfo) GetParingCode() int32`

GetParingCode returns the ParingCode field if non-nil, zero value otherwise.

### GetParingCodeOk

`func (o *ApBridgeInfo) GetParingCodeOk() (*int32, bool)`

GetParingCodeOk returns a tuple with the ParingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParingCode

`func (o *ApBridgeInfo) SetParingCode(v int32)`

SetParingCode sets ParingCode field to given value.

### HasParingCode

`func (o *ApBridgeInfo) HasParingCode() bool`

HasParingCode returns a boolean if a field has been set.

### GetTdmaConfig

`func (o *ApBridgeInfo) GetTdmaConfig() ApBridgeTdmaSettingOpenApiVO`

GetTdmaConfig returns the TdmaConfig field if non-nil, zero value otherwise.

### GetTdmaConfigOk

`func (o *ApBridgeInfo) GetTdmaConfigOk() (*ApBridgeTdmaSettingOpenApiVO, bool)`

GetTdmaConfigOk returns a tuple with the TdmaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdmaConfig

`func (o *ApBridgeInfo) SetTdmaConfig(v ApBridgeTdmaSettingOpenApiVO)`

SetTdmaConfig sets TdmaConfig field to given value.

### HasTdmaConfig

`func (o *ApBridgeInfo) HasTdmaConfig() bool`

HasTdmaConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


