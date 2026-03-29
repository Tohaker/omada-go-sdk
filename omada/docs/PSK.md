# PSK

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | PSK Name | 
**Psk** | **string** | Password | 
**Vlan** | Pointer to **int32** | Vlan Bound With PSK | [optional] 

## Methods

### NewPSK

`func NewPSK(name string, psk string, ) *PSK`

NewPSK instantiates a new PSK object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPSKWithDefaults

`func NewPSKWithDefaults() *PSK`

NewPSKWithDefaults instantiates a new PSK object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PSK) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PSK) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PSK) SetName(v string)`

SetName sets Name field to given value.


### GetPsk

`func (o *PSK) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *PSK) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *PSK) SetPsk(v string)`

SetPsk sets Psk field to given value.


### GetVlan

`func (o *PSK) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *PSK) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *PSK) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *PSK) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


