# PpskSettingV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Mac Bound With PSK.The MAC format requires the use of numbers and uppercase letters and connectors, such as AA-BB-CC-00-11-22. | [optional] 
**Name** | **string** | PPSK Name, should contain 1 to 64 characters. | 
**Psk** | **string** | Password, should contain 8 to 63 visible ASCII characters. | 
**Vlan** | Pointer to **int32** | Vlan Bound With PSK, should be within the range of 1-4094. | [optional] 

## Methods

### NewPpskSettingV2

`func NewPpskSettingV2(name string, psk string, ) *PpskSettingV2`

NewPpskSettingV2 instantiates a new PpskSettingV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPpskSettingV2WithDefaults

`func NewPpskSettingV2WithDefaults() *PpskSettingV2`

NewPpskSettingV2WithDefaults instantiates a new PpskSettingV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *PpskSettingV2) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *PpskSettingV2) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *PpskSettingV2) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *PpskSettingV2) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *PpskSettingV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PpskSettingV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PpskSettingV2) SetName(v string)`

SetName sets Name field to given value.


### GetPsk

`func (o *PpskSettingV2) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *PpskSettingV2) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *PpskSettingV2) SetPsk(v string)`

SetPsk sets Psk field to given value.


### GetVlan

`func (o *PpskSettingV2) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *PpskSettingV2) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *PpskSettingV2) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *PpskSettingV2) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


