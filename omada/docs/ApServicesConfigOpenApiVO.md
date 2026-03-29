# ApServicesConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**L3AccessSetting** | Pointer to [**ApL3AccessVO**](ApL3AccessVO.md) |  | [optional] 
**LldpEnable** | Pointer to **int32** | Parameter [lldpEnable] should be a value as follows: 0:off; 1:on; 2:Use Site Settings. | [optional] 
**LoopbackDetectEnable** | Pointer to **bool** | Whether to enable loopback detection. | [optional] 
**MvlanSetting** | Pointer to [**ApMvlanSettingOpenApiVO**](ApMvlanSettingOpenApiVO.md) |  | [optional] 
**Snmp** | Pointer to [**ApSnmpVO**](ApSnmpVO.md) |  | [optional] 
**VoipVlanSetting** | Pointer to [**ApVoipVlanSettingOpenApiVO**](ApVoipVlanSettingOpenApiVO.md) |  | [optional] 

## Methods

### NewApServicesConfigOpenApiVO

`func NewApServicesConfigOpenApiVO() *ApServicesConfigOpenApiVO`

NewApServicesConfigOpenApiVO instantiates a new ApServicesConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApServicesConfigOpenApiVOWithDefaults

`func NewApServicesConfigOpenApiVOWithDefaults() *ApServicesConfigOpenApiVO`

NewApServicesConfigOpenApiVOWithDefaults instantiates a new ApServicesConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetL3AccessSetting

`func (o *ApServicesConfigOpenApiVO) GetL3AccessSetting() ApL3AccessVO`

GetL3AccessSetting returns the L3AccessSetting field if non-nil, zero value otherwise.

### GetL3AccessSettingOk

`func (o *ApServicesConfigOpenApiVO) GetL3AccessSettingOk() (*ApL3AccessVO, bool)`

GetL3AccessSettingOk returns a tuple with the L3AccessSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3AccessSetting

`func (o *ApServicesConfigOpenApiVO) SetL3AccessSetting(v ApL3AccessVO)`

SetL3AccessSetting sets L3AccessSetting field to given value.

### HasL3AccessSetting

`func (o *ApServicesConfigOpenApiVO) HasL3AccessSetting() bool`

HasL3AccessSetting returns a boolean if a field has been set.

### GetLldpEnable

`func (o *ApServicesConfigOpenApiVO) GetLldpEnable() int32`

GetLldpEnable returns the LldpEnable field if non-nil, zero value otherwise.

### GetLldpEnableOk

`func (o *ApServicesConfigOpenApiVO) GetLldpEnableOk() (*int32, bool)`

GetLldpEnableOk returns a tuple with the LldpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpEnable

`func (o *ApServicesConfigOpenApiVO) SetLldpEnable(v int32)`

SetLldpEnable sets LldpEnable field to given value.

### HasLldpEnable

`func (o *ApServicesConfigOpenApiVO) HasLldpEnable() bool`

HasLldpEnable returns a boolean if a field has been set.

### GetLoopbackDetectEnable

`func (o *ApServicesConfigOpenApiVO) GetLoopbackDetectEnable() bool`

GetLoopbackDetectEnable returns the LoopbackDetectEnable field if non-nil, zero value otherwise.

### GetLoopbackDetectEnableOk

`func (o *ApServicesConfigOpenApiVO) GetLoopbackDetectEnableOk() (*bool, bool)`

GetLoopbackDetectEnableOk returns a tuple with the LoopbackDetectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackDetectEnable

`func (o *ApServicesConfigOpenApiVO) SetLoopbackDetectEnable(v bool)`

SetLoopbackDetectEnable sets LoopbackDetectEnable field to given value.

### HasLoopbackDetectEnable

`func (o *ApServicesConfigOpenApiVO) HasLoopbackDetectEnable() bool`

HasLoopbackDetectEnable returns a boolean if a field has been set.

### GetMvlanSetting

`func (o *ApServicesConfigOpenApiVO) GetMvlanSetting() ApMvlanSettingOpenApiVO`

GetMvlanSetting returns the MvlanSetting field if non-nil, zero value otherwise.

### GetMvlanSettingOk

`func (o *ApServicesConfigOpenApiVO) GetMvlanSettingOk() (*ApMvlanSettingOpenApiVO, bool)`

GetMvlanSettingOk returns a tuple with the MvlanSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMvlanSetting

`func (o *ApServicesConfigOpenApiVO) SetMvlanSetting(v ApMvlanSettingOpenApiVO)`

SetMvlanSetting sets MvlanSetting field to given value.

### HasMvlanSetting

`func (o *ApServicesConfigOpenApiVO) HasMvlanSetting() bool`

HasMvlanSetting returns a boolean if a field has been set.

### GetSnmp

`func (o *ApServicesConfigOpenApiVO) GetSnmp() ApSnmpVO`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *ApServicesConfigOpenApiVO) GetSnmpOk() (*ApSnmpVO, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *ApServicesConfigOpenApiVO) SetSnmp(v ApSnmpVO)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *ApServicesConfigOpenApiVO) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetVoipVlanSetting

`func (o *ApServicesConfigOpenApiVO) GetVoipVlanSetting() ApVoipVlanSettingOpenApiVO`

GetVoipVlanSetting returns the VoipVlanSetting field if non-nil, zero value otherwise.

### GetVoipVlanSettingOk

`func (o *ApServicesConfigOpenApiVO) GetVoipVlanSettingOk() (*ApVoipVlanSettingOpenApiVO, bool)`

GetVoipVlanSettingOk returns a tuple with the VoipVlanSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoipVlanSetting

`func (o *ApServicesConfigOpenApiVO) SetVoipVlanSetting(v ApVoipVlanSettingOpenApiVO)`

SetVoipVlanSetting sets VoipVlanSetting field to given value.

### HasVoipVlanSetting

`func (o *ApServicesConfigOpenApiVO) HasVoipVlanSetting() bool`

HasVoipVlanSetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


