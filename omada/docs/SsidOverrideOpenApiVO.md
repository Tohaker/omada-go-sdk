# SsidOverrideOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverrideSsidEnable** | **bool** | Enable/disable SSID override | 
**OverrideSsidName** | **string** | Override SSID name | 
**OverrideSsidPassword** | Pointer to **string** | Override SSID password(when security is WPA-Personal need fill) | [optional] 
**OverrideVlanEnable** | **bool** | Enable/disable VLAN override | 
**SsidEnable** | **bool** | Enable/disable SSID | 
**SsidId** | **int32** | SSID Entry ID | 
**VlanId** | Pointer to **int32** | VLAN ID | [optional] 

## Methods

### NewSsidOverrideOpenApiVO

`func NewSsidOverrideOpenApiVO(overrideSsidEnable bool, overrideSsidName string, overrideVlanEnable bool, ssidEnable bool, ssidId int32, ) *SsidOverrideOpenApiVO`

NewSsidOverrideOpenApiVO instantiates a new SsidOverrideOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsidOverrideOpenApiVOWithDefaults

`func NewSsidOverrideOpenApiVOWithDefaults() *SsidOverrideOpenApiVO`

NewSsidOverrideOpenApiVOWithDefaults instantiates a new SsidOverrideOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverrideSsidEnable

`func (o *SsidOverrideOpenApiVO) GetOverrideSsidEnable() bool`

GetOverrideSsidEnable returns the OverrideSsidEnable field if non-nil, zero value otherwise.

### GetOverrideSsidEnableOk

`func (o *SsidOverrideOpenApiVO) GetOverrideSsidEnableOk() (*bool, bool)`

GetOverrideSsidEnableOk returns a tuple with the OverrideSsidEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSsidEnable

`func (o *SsidOverrideOpenApiVO) SetOverrideSsidEnable(v bool)`

SetOverrideSsidEnable sets OverrideSsidEnable field to given value.


### GetOverrideSsidName

`func (o *SsidOverrideOpenApiVO) GetOverrideSsidName() string`

GetOverrideSsidName returns the OverrideSsidName field if non-nil, zero value otherwise.

### GetOverrideSsidNameOk

`func (o *SsidOverrideOpenApiVO) GetOverrideSsidNameOk() (*string, bool)`

GetOverrideSsidNameOk returns a tuple with the OverrideSsidName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSsidName

`func (o *SsidOverrideOpenApiVO) SetOverrideSsidName(v string)`

SetOverrideSsidName sets OverrideSsidName field to given value.


### GetOverrideSsidPassword

`func (o *SsidOverrideOpenApiVO) GetOverrideSsidPassword() string`

GetOverrideSsidPassword returns the OverrideSsidPassword field if non-nil, zero value otherwise.

### GetOverrideSsidPasswordOk

`func (o *SsidOverrideOpenApiVO) GetOverrideSsidPasswordOk() (*string, bool)`

GetOverrideSsidPasswordOk returns a tuple with the OverrideSsidPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSsidPassword

`func (o *SsidOverrideOpenApiVO) SetOverrideSsidPassword(v string)`

SetOverrideSsidPassword sets OverrideSsidPassword field to given value.

### HasOverrideSsidPassword

`func (o *SsidOverrideOpenApiVO) HasOverrideSsidPassword() bool`

HasOverrideSsidPassword returns a boolean if a field has been set.

### GetOverrideVlanEnable

`func (o *SsidOverrideOpenApiVO) GetOverrideVlanEnable() bool`

GetOverrideVlanEnable returns the OverrideVlanEnable field if non-nil, zero value otherwise.

### GetOverrideVlanEnableOk

`func (o *SsidOverrideOpenApiVO) GetOverrideVlanEnableOk() (*bool, bool)`

GetOverrideVlanEnableOk returns a tuple with the OverrideVlanEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideVlanEnable

`func (o *SsidOverrideOpenApiVO) SetOverrideVlanEnable(v bool)`

SetOverrideVlanEnable sets OverrideVlanEnable field to given value.


### GetSsidEnable

`func (o *SsidOverrideOpenApiVO) GetSsidEnable() bool`

GetSsidEnable returns the SsidEnable field if non-nil, zero value otherwise.

### GetSsidEnableOk

`func (o *SsidOverrideOpenApiVO) GetSsidEnableOk() (*bool, bool)`

GetSsidEnableOk returns a tuple with the SsidEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidEnable

`func (o *SsidOverrideOpenApiVO) SetSsidEnable(v bool)`

SetSsidEnable sets SsidEnable field to given value.


### GetSsidId

`func (o *SsidOverrideOpenApiVO) GetSsidId() int32`

GetSsidId returns the SsidId field if non-nil, zero value otherwise.

### GetSsidIdOk

`func (o *SsidOverrideOpenApiVO) GetSsidIdOk() (*int32, bool)`

GetSsidIdOk returns a tuple with the SsidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidId

`func (o *SsidOverrideOpenApiVO) SetSsidId(v int32)`

SetSsidId sets SsidId field to given value.


### GetVlanId

`func (o *SsidOverrideOpenApiVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *SsidOverrideOpenApiVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *SsidOverrideOpenApiVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *SsidOverrideOpenApiVO) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


