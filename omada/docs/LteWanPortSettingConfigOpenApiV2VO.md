# LteWanPortSettingConfigOpenApiV2VO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DialupSetting** | [**DialupSettingOpenApiV2VO**](DialupSettingOpenApiV2VO.md) |  | 
**DialupSetting2** | Pointer to [**DialupSettingOpenApiV2VO**](DialupSettingOpenApiV2VO.md) |  | [optional] 
**MobileData** | **bool** | mobile data | 
**PortDesc** | Pointer to **string** | Port description should contain 1 to 32 characters. | [optional] 
**PortUuid** | **string** | Port Uuid | 
**SimPriority** | Pointer to **int32** | Set which SIM card is used first. SIM Priority takes effect only when the device is powered on and the priority is changed. If only one SIM card is inserted, this card is used by default. | [optional] 

## Methods

### NewLteWanPortSettingConfigOpenApiV2VO

`func NewLteWanPortSettingConfigOpenApiV2VO(dialupSetting DialupSettingOpenApiV2VO, mobileData bool, portUuid string, ) *LteWanPortSettingConfigOpenApiV2VO`

NewLteWanPortSettingConfigOpenApiV2VO instantiates a new LteWanPortSettingConfigOpenApiV2VO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLteWanPortSettingConfigOpenApiV2VOWithDefaults

`func NewLteWanPortSettingConfigOpenApiV2VOWithDefaults() *LteWanPortSettingConfigOpenApiV2VO`

NewLteWanPortSettingConfigOpenApiV2VOWithDefaults instantiates a new LteWanPortSettingConfigOpenApiV2VO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDialupSetting

`func (o *LteWanPortSettingConfigOpenApiV2VO) GetDialupSetting() DialupSettingOpenApiV2VO`

GetDialupSetting returns the DialupSetting field if non-nil, zero value otherwise.

### GetDialupSettingOk

`func (o *LteWanPortSettingConfigOpenApiV2VO) GetDialupSettingOk() (*DialupSettingOpenApiV2VO, bool)`

GetDialupSettingOk returns a tuple with the DialupSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialupSetting

`func (o *LteWanPortSettingConfigOpenApiV2VO) SetDialupSetting(v DialupSettingOpenApiV2VO)`

SetDialupSetting sets DialupSetting field to given value.


### GetDialupSetting2

`func (o *LteWanPortSettingConfigOpenApiV2VO) GetDialupSetting2() DialupSettingOpenApiV2VO`

GetDialupSetting2 returns the DialupSetting2 field if non-nil, zero value otherwise.

### GetDialupSetting2Ok

`func (o *LteWanPortSettingConfigOpenApiV2VO) GetDialupSetting2Ok() (*DialupSettingOpenApiV2VO, bool)`

GetDialupSetting2Ok returns a tuple with the DialupSetting2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialupSetting2

`func (o *LteWanPortSettingConfigOpenApiV2VO) SetDialupSetting2(v DialupSettingOpenApiV2VO)`

SetDialupSetting2 sets DialupSetting2 field to given value.

### HasDialupSetting2

`func (o *LteWanPortSettingConfigOpenApiV2VO) HasDialupSetting2() bool`

HasDialupSetting2 returns a boolean if a field has been set.

### GetMobileData

`func (o *LteWanPortSettingConfigOpenApiV2VO) GetMobileData() bool`

GetMobileData returns the MobileData field if non-nil, zero value otherwise.

### GetMobileDataOk

`func (o *LteWanPortSettingConfigOpenApiV2VO) GetMobileDataOk() (*bool, bool)`

GetMobileDataOk returns a tuple with the MobileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileData

`func (o *LteWanPortSettingConfigOpenApiV2VO) SetMobileData(v bool)`

SetMobileData sets MobileData field to given value.


### GetPortDesc

`func (o *LteWanPortSettingConfigOpenApiV2VO) GetPortDesc() string`

GetPortDesc returns the PortDesc field if non-nil, zero value otherwise.

### GetPortDescOk

`func (o *LteWanPortSettingConfigOpenApiV2VO) GetPortDescOk() (*string, bool)`

GetPortDescOk returns a tuple with the PortDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDesc

`func (o *LteWanPortSettingConfigOpenApiV2VO) SetPortDesc(v string)`

SetPortDesc sets PortDesc field to given value.

### HasPortDesc

`func (o *LteWanPortSettingConfigOpenApiV2VO) HasPortDesc() bool`

HasPortDesc returns a boolean if a field has been set.

### GetPortUuid

`func (o *LteWanPortSettingConfigOpenApiV2VO) GetPortUuid() string`

GetPortUuid returns the PortUuid field if non-nil, zero value otherwise.

### GetPortUuidOk

`func (o *LteWanPortSettingConfigOpenApiV2VO) GetPortUuidOk() (*string, bool)`

GetPortUuidOk returns a tuple with the PortUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUuid

`func (o *LteWanPortSettingConfigOpenApiV2VO) SetPortUuid(v string)`

SetPortUuid sets PortUuid field to given value.


### GetSimPriority

`func (o *LteWanPortSettingConfigOpenApiV2VO) GetSimPriority() int32`

GetSimPriority returns the SimPriority field if non-nil, zero value otherwise.

### GetSimPriorityOk

`func (o *LteWanPortSettingConfigOpenApiV2VO) GetSimPriorityOk() (*int32, bool)`

GetSimPriorityOk returns a tuple with the SimPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimPriority

`func (o *LteWanPortSettingConfigOpenApiV2VO) SetSimPriority(v int32)`

SetSimPriority sets SimPriority field to given value.

### HasSimPriority

`func (o *LteWanPortSettingConfigOpenApiV2VO) HasSimPriority() bool`

HasSimPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


