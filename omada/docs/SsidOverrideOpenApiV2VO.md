# SsidOverrideOpenApiV2VO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverrideSsidEnable** | **bool** | Enable/disable SSID override | 
**OverrideSsidName** | Pointer to **string** | Override SSID name. It should contain 1 to 32 UTF-8 characters. | [optional] 
**OverrideSsidPassword** | Pointer to **string** | Override SSID password(when security is WPA-Personal need fill).It should contain 8-63 printable ASCII characters or 8-63 hexadecimal digits. | [optional] 
**OverrideVlanEnable** | **bool** | Enable/disable VLAN override | 
**SsidEnable** | Pointer to **bool** | Enable/disable SSID | [optional] 
**SsidEntryId** | **int32** | This field represents SSID Entry ID.SSID Entry ID can be obtained from &#39;Get ap WLANs override config&#39; interface. | 
**VlanId** | Pointer to **int32** | VLAN ID | [optional] 
**VlanPoolIds** | Pointer to **string** | VLAN POOL IDS | [optional] 

## Methods

### NewSsidOverrideOpenApiV2VO

`func NewSsidOverrideOpenApiV2VO(overrideSsidEnable bool, overrideVlanEnable bool, ssidEntryId int32, ) *SsidOverrideOpenApiV2VO`

NewSsidOverrideOpenApiV2VO instantiates a new SsidOverrideOpenApiV2VO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsidOverrideOpenApiV2VOWithDefaults

`func NewSsidOverrideOpenApiV2VOWithDefaults() *SsidOverrideOpenApiV2VO`

NewSsidOverrideOpenApiV2VOWithDefaults instantiates a new SsidOverrideOpenApiV2VO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverrideSsidEnable

`func (o *SsidOverrideOpenApiV2VO) GetOverrideSsidEnable() bool`

GetOverrideSsidEnable returns the OverrideSsidEnable field if non-nil, zero value otherwise.

### GetOverrideSsidEnableOk

`func (o *SsidOverrideOpenApiV2VO) GetOverrideSsidEnableOk() (*bool, bool)`

GetOverrideSsidEnableOk returns a tuple with the OverrideSsidEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSsidEnable

`func (o *SsidOverrideOpenApiV2VO) SetOverrideSsidEnable(v bool)`

SetOverrideSsidEnable sets OverrideSsidEnable field to given value.


### GetOverrideSsidName

`func (o *SsidOverrideOpenApiV2VO) GetOverrideSsidName() string`

GetOverrideSsidName returns the OverrideSsidName field if non-nil, zero value otherwise.

### GetOverrideSsidNameOk

`func (o *SsidOverrideOpenApiV2VO) GetOverrideSsidNameOk() (*string, bool)`

GetOverrideSsidNameOk returns a tuple with the OverrideSsidName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSsidName

`func (o *SsidOverrideOpenApiV2VO) SetOverrideSsidName(v string)`

SetOverrideSsidName sets OverrideSsidName field to given value.

### HasOverrideSsidName

`func (o *SsidOverrideOpenApiV2VO) HasOverrideSsidName() bool`

HasOverrideSsidName returns a boolean if a field has been set.

### GetOverrideSsidPassword

`func (o *SsidOverrideOpenApiV2VO) GetOverrideSsidPassword() string`

GetOverrideSsidPassword returns the OverrideSsidPassword field if non-nil, zero value otherwise.

### GetOverrideSsidPasswordOk

`func (o *SsidOverrideOpenApiV2VO) GetOverrideSsidPasswordOk() (*string, bool)`

GetOverrideSsidPasswordOk returns a tuple with the OverrideSsidPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSsidPassword

`func (o *SsidOverrideOpenApiV2VO) SetOverrideSsidPassword(v string)`

SetOverrideSsidPassword sets OverrideSsidPassword field to given value.

### HasOverrideSsidPassword

`func (o *SsidOverrideOpenApiV2VO) HasOverrideSsidPassword() bool`

HasOverrideSsidPassword returns a boolean if a field has been set.

### GetOverrideVlanEnable

`func (o *SsidOverrideOpenApiV2VO) GetOverrideVlanEnable() bool`

GetOverrideVlanEnable returns the OverrideVlanEnable field if non-nil, zero value otherwise.

### GetOverrideVlanEnableOk

`func (o *SsidOverrideOpenApiV2VO) GetOverrideVlanEnableOk() (*bool, bool)`

GetOverrideVlanEnableOk returns a tuple with the OverrideVlanEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideVlanEnable

`func (o *SsidOverrideOpenApiV2VO) SetOverrideVlanEnable(v bool)`

SetOverrideVlanEnable sets OverrideVlanEnable field to given value.


### GetSsidEnable

`func (o *SsidOverrideOpenApiV2VO) GetSsidEnable() bool`

GetSsidEnable returns the SsidEnable field if non-nil, zero value otherwise.

### GetSsidEnableOk

`func (o *SsidOverrideOpenApiV2VO) GetSsidEnableOk() (*bool, bool)`

GetSsidEnableOk returns a tuple with the SsidEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidEnable

`func (o *SsidOverrideOpenApiV2VO) SetSsidEnable(v bool)`

SetSsidEnable sets SsidEnable field to given value.

### HasSsidEnable

`func (o *SsidOverrideOpenApiV2VO) HasSsidEnable() bool`

HasSsidEnable returns a boolean if a field has been set.

### GetSsidEntryId

`func (o *SsidOverrideOpenApiV2VO) GetSsidEntryId() int32`

GetSsidEntryId returns the SsidEntryId field if non-nil, zero value otherwise.

### GetSsidEntryIdOk

`func (o *SsidOverrideOpenApiV2VO) GetSsidEntryIdOk() (*int32, bool)`

GetSsidEntryIdOk returns a tuple with the SsidEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidEntryId

`func (o *SsidOverrideOpenApiV2VO) SetSsidEntryId(v int32)`

SetSsidEntryId sets SsidEntryId field to given value.


### GetVlanId

`func (o *SsidOverrideOpenApiV2VO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *SsidOverrideOpenApiV2VO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *SsidOverrideOpenApiV2VO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *SsidOverrideOpenApiV2VO) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlanPoolIds

`func (o *SsidOverrideOpenApiV2VO) GetVlanPoolIds() string`

GetVlanPoolIds returns the VlanPoolIds field if non-nil, zero value otherwise.

### GetVlanPoolIdsOk

`func (o *SsidOverrideOpenApiV2VO) GetVlanPoolIdsOk() (*string, bool)`

GetVlanPoolIdsOk returns a tuple with the VlanPoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPoolIds

`func (o *SsidOverrideOpenApiV2VO) SetVlanPoolIds(v string)`

SetVlanPoolIds sets VlanPoolIds field to given value.

### HasVlanPoolIds

`func (o *SsidOverrideOpenApiV2VO) HasVlanPoolIds() bool`

HasVlanPoolIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


