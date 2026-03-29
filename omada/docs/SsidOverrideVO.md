# SsidOverrideVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandInfo** | Pointer to **int32** |  | [optional] 
**Enable** | **bool** |  | 
**GlobalSsid** | **string** |  | 
**HidePwd** | Pointer to **bool** |  | [optional] 
**Index** | **int32** |  | 
**Psk** | Pointer to **string** |  | [optional] 
**Security** | Pointer to **int32** |  | [optional] 
**Ssid** | Pointer to **string** |  | [optional] 
**SsidEnable** | Pointer to **bool** |  | [optional] 
**SupportBands** | Pointer to **[]int32** |  | [optional] 
**VlanEnable** | Pointer to **bool** |  | [optional] 
**VlanId** | Pointer to **int32** |  | [optional] 
**VlanPoolIds** | Pointer to **string** |  | [optional] 

## Methods

### NewSsidOverrideVO

`func NewSsidOverrideVO(enable bool, globalSsid string, index int32, ) *SsidOverrideVO`

NewSsidOverrideVO instantiates a new SsidOverrideVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsidOverrideVOWithDefaults

`func NewSsidOverrideVOWithDefaults() *SsidOverrideVO`

NewSsidOverrideVOWithDefaults instantiates a new SsidOverrideVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandInfo

`func (o *SsidOverrideVO) GetBandInfo() int32`

GetBandInfo returns the BandInfo field if non-nil, zero value otherwise.

### GetBandInfoOk

`func (o *SsidOverrideVO) GetBandInfoOk() (*int32, bool)`

GetBandInfoOk returns a tuple with the BandInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandInfo

`func (o *SsidOverrideVO) SetBandInfo(v int32)`

SetBandInfo sets BandInfo field to given value.

### HasBandInfo

`func (o *SsidOverrideVO) HasBandInfo() bool`

HasBandInfo returns a boolean if a field has been set.

### GetEnable

`func (o *SsidOverrideVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *SsidOverrideVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *SsidOverrideVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetGlobalSsid

`func (o *SsidOverrideVO) GetGlobalSsid() string`

GetGlobalSsid returns the GlobalSsid field if non-nil, zero value otherwise.

### GetGlobalSsidOk

`func (o *SsidOverrideVO) GetGlobalSsidOk() (*string, bool)`

GetGlobalSsidOk returns a tuple with the GlobalSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSsid

`func (o *SsidOverrideVO) SetGlobalSsid(v string)`

SetGlobalSsid sets GlobalSsid field to given value.


### GetHidePwd

`func (o *SsidOverrideVO) GetHidePwd() bool`

GetHidePwd returns the HidePwd field if non-nil, zero value otherwise.

### GetHidePwdOk

`func (o *SsidOverrideVO) GetHidePwdOk() (*bool, bool)`

GetHidePwdOk returns a tuple with the HidePwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePwd

`func (o *SsidOverrideVO) SetHidePwd(v bool)`

SetHidePwd sets HidePwd field to given value.

### HasHidePwd

`func (o *SsidOverrideVO) HasHidePwd() bool`

HasHidePwd returns a boolean if a field has been set.

### GetIndex

`func (o *SsidOverrideVO) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *SsidOverrideVO) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *SsidOverrideVO) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetPsk

`func (o *SsidOverrideVO) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *SsidOverrideVO) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *SsidOverrideVO) SetPsk(v string)`

SetPsk sets Psk field to given value.

### HasPsk

`func (o *SsidOverrideVO) HasPsk() bool`

HasPsk returns a boolean if a field has been set.

### GetSecurity

`func (o *SsidOverrideVO) GetSecurity() int32`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *SsidOverrideVO) GetSecurityOk() (*int32, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *SsidOverrideVO) SetSecurity(v int32)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *SsidOverrideVO) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetSsid

`func (o *SsidOverrideVO) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *SsidOverrideVO) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *SsidOverrideVO) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *SsidOverrideVO) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetSsidEnable

`func (o *SsidOverrideVO) GetSsidEnable() bool`

GetSsidEnable returns the SsidEnable field if non-nil, zero value otherwise.

### GetSsidEnableOk

`func (o *SsidOverrideVO) GetSsidEnableOk() (*bool, bool)`

GetSsidEnableOk returns a tuple with the SsidEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidEnable

`func (o *SsidOverrideVO) SetSsidEnable(v bool)`

SetSsidEnable sets SsidEnable field to given value.

### HasSsidEnable

`func (o *SsidOverrideVO) HasSsidEnable() bool`

HasSsidEnable returns a boolean if a field has been set.

### GetSupportBands

`func (o *SsidOverrideVO) GetSupportBands() []int32`

GetSupportBands returns the SupportBands field if non-nil, zero value otherwise.

### GetSupportBandsOk

`func (o *SsidOverrideVO) GetSupportBandsOk() (*[]int32, bool)`

GetSupportBandsOk returns a tuple with the SupportBands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportBands

`func (o *SsidOverrideVO) SetSupportBands(v []int32)`

SetSupportBands sets SupportBands field to given value.

### HasSupportBands

`func (o *SsidOverrideVO) HasSupportBands() bool`

HasSupportBands returns a boolean if a field has been set.

### GetVlanEnable

`func (o *SsidOverrideVO) GetVlanEnable() bool`

GetVlanEnable returns the VlanEnable field if non-nil, zero value otherwise.

### GetVlanEnableOk

`func (o *SsidOverrideVO) GetVlanEnableOk() (*bool, bool)`

GetVlanEnableOk returns a tuple with the VlanEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanEnable

`func (o *SsidOverrideVO) SetVlanEnable(v bool)`

SetVlanEnable sets VlanEnable field to given value.

### HasVlanEnable

`func (o *SsidOverrideVO) HasVlanEnable() bool`

HasVlanEnable returns a boolean if a field has been set.

### GetVlanId

`func (o *SsidOverrideVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *SsidOverrideVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *SsidOverrideVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *SsidOverrideVO) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlanPoolIds

`func (o *SsidOverrideVO) GetVlanPoolIds() string`

GetVlanPoolIds returns the VlanPoolIds field if non-nil, zero value otherwise.

### GetVlanPoolIdsOk

`func (o *SsidOverrideVO) GetVlanPoolIdsOk() (*string, bool)`

GetVlanPoolIdsOk returns a tuple with the VlanPoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPoolIds

`func (o *SsidOverrideVO) SetVlanPoolIds(v string)`

SetVlanPoolIds sets VlanPoolIds field to given value.

### HasVlanPoolIds

`func (o *SsidOverrideVO) HasVlanPoolIds() bool`

HasVlanPoolIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


