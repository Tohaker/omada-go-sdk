# PpskAutoCreateSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Length** | **int32** | PSK Password Length, should be within the range of 8-63. | 
**Number** | **int32** | Generate Number, should be within the range of 1-1024. | 
**Prefix** | **string** | PSK Name Prefix, should contain 1 to 60 visible ASCII characters. | 
**Vlan** | Pointer to **int32** | PSK Bound Vlan, should be within the range of 1-4094. | [optional] 

## Methods

### NewPpskAutoCreateSetting

`func NewPpskAutoCreateSetting(length int32, number int32, prefix string, ) *PpskAutoCreateSetting`

NewPpskAutoCreateSetting instantiates a new PpskAutoCreateSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPpskAutoCreateSettingWithDefaults

`func NewPpskAutoCreateSettingWithDefaults() *PpskAutoCreateSetting`

NewPpskAutoCreateSettingWithDefaults instantiates a new PpskAutoCreateSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLength

`func (o *PpskAutoCreateSetting) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *PpskAutoCreateSetting) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *PpskAutoCreateSetting) SetLength(v int32)`

SetLength sets Length field to given value.


### GetNumber

`func (o *PpskAutoCreateSetting) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PpskAutoCreateSetting) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PpskAutoCreateSetting) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetPrefix

`func (o *PpskAutoCreateSetting) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PpskAutoCreateSetting) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PpskAutoCreateSetting) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetVlan

`func (o *PpskAutoCreateSetting) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *PpskAutoCreateSetting) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *PpskAutoCreateSetting) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *PpskAutoCreateSetting) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


