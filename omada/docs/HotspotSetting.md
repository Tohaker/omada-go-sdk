# HotspotSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnabledTypes** | **[]int32** | Hotspot enabled auth types, should be a value as follows: 3: Voucher, 5: Local User, 8: Hotspot RADIUS, 6: Sms, 12: Form Auth. | 

## Methods

### NewHotspotSetting

`func NewHotspotSetting(enabledTypes []int32, ) *HotspotSetting`

NewHotspotSetting instantiates a new HotspotSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHotspotSettingWithDefaults

`func NewHotspotSettingWithDefaults() *HotspotSetting`

NewHotspotSettingWithDefaults instantiates a new HotspotSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabledTypes

`func (o *HotspotSetting) GetEnabledTypes() []int32`

GetEnabledTypes returns the EnabledTypes field if non-nil, zero value otherwise.

### GetEnabledTypesOk

`func (o *HotspotSetting) GetEnabledTypesOk() (*[]int32, bool)`

GetEnabledTypesOk returns a tuple with the EnabledTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledTypes

`func (o *HotspotSetting) SetEnabledTypes(v []int32)`

SetEnabledTypes sets EnabledTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


