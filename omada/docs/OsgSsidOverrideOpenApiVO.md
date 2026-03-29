# OsgSsidOverrideOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandInfo** | Pointer to **int32** | Band. Deprecated, use supportBands instead | [optional] 
**Enable** | **bool** | Enable/disable SSID override | 
**GlobalSsid** | **string** | SSID Name | 
**HidePwd** | **bool** | hide password | 
**Index** | **int32** | SSID Entry ID | 
**Psk** | Pointer to **string** | Override SSID Password | [optional] 
**Security** | **int32** | SSID Security Mode, must be the same as the current SSID Security | 
**Ssid** | **string** | Override SSID | 
**SsidEnable** | **bool** | Enable/disable SSID | 
**SupportBands** | Pointer to **[]int32** | support band 2.4GHz(0) 5GHz(1) 6GHz(2) | [optional] 
**VlanEnable** | **bool** | Enable/disable VLAN override | 
**VlanId** | Pointer to **int32** | VLAN ID | [optional] 

## Methods

### NewOsgSsidOverrideOpenApiVO

`func NewOsgSsidOverrideOpenApiVO(enable bool, globalSsid string, hidePwd bool, index int32, security int32, ssid string, ssidEnable bool, vlanEnable bool, ) *OsgSsidOverrideOpenApiVO`

NewOsgSsidOverrideOpenApiVO instantiates a new OsgSsidOverrideOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgSsidOverrideOpenApiVOWithDefaults

`func NewOsgSsidOverrideOpenApiVOWithDefaults() *OsgSsidOverrideOpenApiVO`

NewOsgSsidOverrideOpenApiVOWithDefaults instantiates a new OsgSsidOverrideOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandInfo

`func (o *OsgSsidOverrideOpenApiVO) GetBandInfo() int32`

GetBandInfo returns the BandInfo field if non-nil, zero value otherwise.

### GetBandInfoOk

`func (o *OsgSsidOverrideOpenApiVO) GetBandInfoOk() (*int32, bool)`

GetBandInfoOk returns a tuple with the BandInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandInfo

`func (o *OsgSsidOverrideOpenApiVO) SetBandInfo(v int32)`

SetBandInfo sets BandInfo field to given value.

### HasBandInfo

`func (o *OsgSsidOverrideOpenApiVO) HasBandInfo() bool`

HasBandInfo returns a boolean if a field has been set.

### GetEnable

`func (o *OsgSsidOverrideOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *OsgSsidOverrideOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *OsgSsidOverrideOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetGlobalSsid

`func (o *OsgSsidOverrideOpenApiVO) GetGlobalSsid() string`

GetGlobalSsid returns the GlobalSsid field if non-nil, zero value otherwise.

### GetGlobalSsidOk

`func (o *OsgSsidOverrideOpenApiVO) GetGlobalSsidOk() (*string, bool)`

GetGlobalSsidOk returns a tuple with the GlobalSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSsid

`func (o *OsgSsidOverrideOpenApiVO) SetGlobalSsid(v string)`

SetGlobalSsid sets GlobalSsid field to given value.


### GetHidePwd

`func (o *OsgSsidOverrideOpenApiVO) GetHidePwd() bool`

GetHidePwd returns the HidePwd field if non-nil, zero value otherwise.

### GetHidePwdOk

`func (o *OsgSsidOverrideOpenApiVO) GetHidePwdOk() (*bool, bool)`

GetHidePwdOk returns a tuple with the HidePwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePwd

`func (o *OsgSsidOverrideOpenApiVO) SetHidePwd(v bool)`

SetHidePwd sets HidePwd field to given value.


### GetIndex

`func (o *OsgSsidOverrideOpenApiVO) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *OsgSsidOverrideOpenApiVO) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *OsgSsidOverrideOpenApiVO) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetPsk

`func (o *OsgSsidOverrideOpenApiVO) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *OsgSsidOverrideOpenApiVO) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *OsgSsidOverrideOpenApiVO) SetPsk(v string)`

SetPsk sets Psk field to given value.

### HasPsk

`func (o *OsgSsidOverrideOpenApiVO) HasPsk() bool`

HasPsk returns a boolean if a field has been set.

### GetSecurity

`func (o *OsgSsidOverrideOpenApiVO) GetSecurity() int32`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *OsgSsidOverrideOpenApiVO) GetSecurityOk() (*int32, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *OsgSsidOverrideOpenApiVO) SetSecurity(v int32)`

SetSecurity sets Security field to given value.


### GetSsid

`func (o *OsgSsidOverrideOpenApiVO) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *OsgSsidOverrideOpenApiVO) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *OsgSsidOverrideOpenApiVO) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetSsidEnable

`func (o *OsgSsidOverrideOpenApiVO) GetSsidEnable() bool`

GetSsidEnable returns the SsidEnable field if non-nil, zero value otherwise.

### GetSsidEnableOk

`func (o *OsgSsidOverrideOpenApiVO) GetSsidEnableOk() (*bool, bool)`

GetSsidEnableOk returns a tuple with the SsidEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidEnable

`func (o *OsgSsidOverrideOpenApiVO) SetSsidEnable(v bool)`

SetSsidEnable sets SsidEnable field to given value.


### GetSupportBands

`func (o *OsgSsidOverrideOpenApiVO) GetSupportBands() []int32`

GetSupportBands returns the SupportBands field if non-nil, zero value otherwise.

### GetSupportBandsOk

`func (o *OsgSsidOverrideOpenApiVO) GetSupportBandsOk() (*[]int32, bool)`

GetSupportBandsOk returns a tuple with the SupportBands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportBands

`func (o *OsgSsidOverrideOpenApiVO) SetSupportBands(v []int32)`

SetSupportBands sets SupportBands field to given value.

### HasSupportBands

`func (o *OsgSsidOverrideOpenApiVO) HasSupportBands() bool`

HasSupportBands returns a boolean if a field has been set.

### GetVlanEnable

`func (o *OsgSsidOverrideOpenApiVO) GetVlanEnable() bool`

GetVlanEnable returns the VlanEnable field if non-nil, zero value otherwise.

### GetVlanEnableOk

`func (o *OsgSsidOverrideOpenApiVO) GetVlanEnableOk() (*bool, bool)`

GetVlanEnableOk returns a tuple with the VlanEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanEnable

`func (o *OsgSsidOverrideOpenApiVO) SetVlanEnable(v bool)`

SetVlanEnable sets VlanEnable field to given value.


### GetVlanId

`func (o *OsgSsidOverrideOpenApiVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *OsgSsidOverrideOpenApiVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *OsgSsidOverrideOpenApiVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *OsgSsidOverrideOpenApiVO) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


