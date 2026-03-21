# PpskSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Mac Bound With PSK.The MAC format requires the use of numbers and uppercase letters and connectors, such as AA-BB-CC-00-11-22. | [optional] 
**Name** | **string** | PPSK Name, should contain 1 to 64 characters. | 
**Psk** | **string** | Password, should contain 8 to 63 visible ASCII characters. | 
**Vlan** | Pointer to **int32** | Vlan Bound With PSK, should be within the range of 1-4094. | [optional] 

## Methods

### NewPpskSetting

`func NewPpskSetting(name string, psk string, ) *PpskSetting`

NewPpskSetting instantiates a new PpskSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPpskSettingWithDefaults

`func NewPpskSettingWithDefaults() *PpskSetting`

NewPpskSettingWithDefaults instantiates a new PpskSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *PpskSetting) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *PpskSetting) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *PpskSetting) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *PpskSetting) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *PpskSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PpskSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PpskSetting) SetName(v string)`

SetName sets Name field to given value.


### GetPsk

`func (o *PpskSetting) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *PpskSetting) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *PpskSetting) SetPsk(v string)`

SetPsk sets Psk field to given value.


### GetVlan

`func (o *PpskSetting) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *PpskSetting) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *PpskSetting) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *PpskSetting) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


