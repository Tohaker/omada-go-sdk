# PSKVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Mac Bound With PSK | [optional] 
**Name** | **string** | PSK Name | 
**Psk** | **string** | Password | 
**Vlan** | Pointer to **int32** | Vlan Bound With PSK | [optional] 

## Methods

### NewPSKVO

`func NewPSKVO(name string, psk string, ) *PSKVO`

NewPSKVO instantiates a new PSKVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPSKVOWithDefaults

`func NewPSKVOWithDefaults() *PSKVO`

NewPSKVOWithDefaults instantiates a new PSKVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *PSKVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *PSKVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *PSKVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *PSKVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *PSKVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PSKVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PSKVO) SetName(v string)`

SetName sets Name field to given value.


### GetPsk

`func (o *PSKVO) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *PSKVO) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *PSKVO) SetPsk(v string)`

SetPsk sets Psk field to given value.


### GetVlan

`func (o *PSKVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *PSKVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *PSKVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *PSKVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


