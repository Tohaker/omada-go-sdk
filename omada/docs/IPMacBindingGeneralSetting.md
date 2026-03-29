# IPMacBindingGeneralSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Enable of the IP MAC binding general setting, If applied, fill in at least one of LAN and WAN. | 
**Garp** | Pointer to **bool** | GARP of the IP MAC binding general setting. | [optional] 
**ImbPass** | Pointer to **bool** | ImbPass of the IP MAC binding general setting. | [optional] 
**Interval** | Pointer to **int32** | Interval should be within the range of 1–10000. | [optional] 
**LanIds** | Pointer to **[]string** | LANs of the IP MAC binding general setting. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | [optional] 
**WanIds** | Pointer to **[]string** | WANs of the IP MAC binding general setting. WAN port ID can be obtained from &#39;Get internet basic info&#39; interface. | [optional] 

## Methods

### NewIPMacBindingGeneralSetting

`func NewIPMacBindingGeneralSetting(enable bool, ) *IPMacBindingGeneralSetting`

NewIPMacBindingGeneralSetting instantiates a new IPMacBindingGeneralSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPMacBindingGeneralSettingWithDefaults

`func NewIPMacBindingGeneralSettingWithDefaults() *IPMacBindingGeneralSetting`

NewIPMacBindingGeneralSettingWithDefaults instantiates a new IPMacBindingGeneralSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *IPMacBindingGeneralSetting) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IPMacBindingGeneralSetting) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IPMacBindingGeneralSetting) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetGarp

`func (o *IPMacBindingGeneralSetting) GetGarp() bool`

GetGarp returns the Garp field if non-nil, zero value otherwise.

### GetGarpOk

`func (o *IPMacBindingGeneralSetting) GetGarpOk() (*bool, bool)`

GetGarpOk returns a tuple with the Garp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarp

`func (o *IPMacBindingGeneralSetting) SetGarp(v bool)`

SetGarp sets Garp field to given value.

### HasGarp

`func (o *IPMacBindingGeneralSetting) HasGarp() bool`

HasGarp returns a boolean if a field has been set.

### GetImbPass

`func (o *IPMacBindingGeneralSetting) GetImbPass() bool`

GetImbPass returns the ImbPass field if non-nil, zero value otherwise.

### GetImbPassOk

`func (o *IPMacBindingGeneralSetting) GetImbPassOk() (*bool, bool)`

GetImbPassOk returns a tuple with the ImbPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImbPass

`func (o *IPMacBindingGeneralSetting) SetImbPass(v bool)`

SetImbPass sets ImbPass field to given value.

### HasImbPass

`func (o *IPMacBindingGeneralSetting) HasImbPass() bool`

HasImbPass returns a boolean if a field has been set.

### GetInterval

`func (o *IPMacBindingGeneralSetting) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *IPMacBindingGeneralSetting) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *IPMacBindingGeneralSetting) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *IPMacBindingGeneralSetting) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetLanIds

`func (o *IPMacBindingGeneralSetting) GetLanIds() []string`

GetLanIds returns the LanIds field if non-nil, zero value otherwise.

### GetLanIdsOk

`func (o *IPMacBindingGeneralSetting) GetLanIdsOk() (*[]string, bool)`

GetLanIdsOk returns a tuple with the LanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIds

`func (o *IPMacBindingGeneralSetting) SetLanIds(v []string)`

SetLanIds sets LanIds field to given value.

### HasLanIds

`func (o *IPMacBindingGeneralSetting) HasLanIds() bool`

HasLanIds returns a boolean if a field has been set.

### GetWanIds

`func (o *IPMacBindingGeneralSetting) GetWanIds() []string`

GetWanIds returns the WanIds field if non-nil, zero value otherwise.

### GetWanIdsOk

`func (o *IPMacBindingGeneralSetting) GetWanIdsOk() (*[]string, bool)`

GetWanIdsOk returns a tuple with the WanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanIds

`func (o *IPMacBindingGeneralSetting) SetWanIds(v []string)`

SetWanIds sets WanIds field to given value.

### HasWanIds

`func (o *IPMacBindingGeneralSetting) HasWanIds() bool`

HasWanIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


