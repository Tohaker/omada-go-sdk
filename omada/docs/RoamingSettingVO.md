# RoamingSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiRoamingEnable** | **bool** | Whether to enable AI roaming, this configuration will take effect only when fast roaming is enabled | 
**DualBand11kReportEnable** | Pointer to **bool** | Whether to enable 802.11k report | [optional] 
**FastRoamingEnable** | **bool** | Whether to enable fast roaming | 
**ForceDisassociationEnable** | Pointer to **bool** | Whether to enable forced disassociation | [optional] 
**NonPingPongRoamingEnable** | Pointer to **bool** | Whether to enable non-ping-pong roaming | [optional] 
**NonStickRoamingEnable** | Pointer to **bool** | Whether to enable non-stick roaming | [optional] 
**PingPongAssocThreshold** | Pointer to **int32** | Association frequency threshold of ping-pong roaming | [optional] 

## Methods

### NewRoamingSettingVO

`func NewRoamingSettingVO(aiRoamingEnable bool, fastRoamingEnable bool, ) *RoamingSettingVO`

NewRoamingSettingVO instantiates a new RoamingSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoamingSettingVOWithDefaults

`func NewRoamingSettingVOWithDefaults() *RoamingSettingVO`

NewRoamingSettingVOWithDefaults instantiates a new RoamingSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiRoamingEnable

`func (o *RoamingSettingVO) GetAiRoamingEnable() bool`

GetAiRoamingEnable returns the AiRoamingEnable field if non-nil, zero value otherwise.

### GetAiRoamingEnableOk

`func (o *RoamingSettingVO) GetAiRoamingEnableOk() (*bool, bool)`

GetAiRoamingEnableOk returns a tuple with the AiRoamingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiRoamingEnable

`func (o *RoamingSettingVO) SetAiRoamingEnable(v bool)`

SetAiRoamingEnable sets AiRoamingEnable field to given value.


### GetDualBand11kReportEnable

`func (o *RoamingSettingVO) GetDualBand11kReportEnable() bool`

GetDualBand11kReportEnable returns the DualBand11kReportEnable field if non-nil, zero value otherwise.

### GetDualBand11kReportEnableOk

`func (o *RoamingSettingVO) GetDualBand11kReportEnableOk() (*bool, bool)`

GetDualBand11kReportEnableOk returns a tuple with the DualBand11kReportEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDualBand11kReportEnable

`func (o *RoamingSettingVO) SetDualBand11kReportEnable(v bool)`

SetDualBand11kReportEnable sets DualBand11kReportEnable field to given value.

### HasDualBand11kReportEnable

`func (o *RoamingSettingVO) HasDualBand11kReportEnable() bool`

HasDualBand11kReportEnable returns a boolean if a field has been set.

### GetFastRoamingEnable

`func (o *RoamingSettingVO) GetFastRoamingEnable() bool`

GetFastRoamingEnable returns the FastRoamingEnable field if non-nil, zero value otherwise.

### GetFastRoamingEnableOk

`func (o *RoamingSettingVO) GetFastRoamingEnableOk() (*bool, bool)`

GetFastRoamingEnableOk returns a tuple with the FastRoamingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastRoamingEnable

`func (o *RoamingSettingVO) SetFastRoamingEnable(v bool)`

SetFastRoamingEnable sets FastRoamingEnable field to given value.


### GetForceDisassociationEnable

`func (o *RoamingSettingVO) GetForceDisassociationEnable() bool`

GetForceDisassociationEnable returns the ForceDisassociationEnable field if non-nil, zero value otherwise.

### GetForceDisassociationEnableOk

`func (o *RoamingSettingVO) GetForceDisassociationEnableOk() (*bool, bool)`

GetForceDisassociationEnableOk returns a tuple with the ForceDisassociationEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDisassociationEnable

`func (o *RoamingSettingVO) SetForceDisassociationEnable(v bool)`

SetForceDisassociationEnable sets ForceDisassociationEnable field to given value.

### HasForceDisassociationEnable

`func (o *RoamingSettingVO) HasForceDisassociationEnable() bool`

HasForceDisassociationEnable returns a boolean if a field has been set.

### GetNonPingPongRoamingEnable

`func (o *RoamingSettingVO) GetNonPingPongRoamingEnable() bool`

GetNonPingPongRoamingEnable returns the NonPingPongRoamingEnable field if non-nil, zero value otherwise.

### GetNonPingPongRoamingEnableOk

`func (o *RoamingSettingVO) GetNonPingPongRoamingEnableOk() (*bool, bool)`

GetNonPingPongRoamingEnableOk returns a tuple with the NonPingPongRoamingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonPingPongRoamingEnable

`func (o *RoamingSettingVO) SetNonPingPongRoamingEnable(v bool)`

SetNonPingPongRoamingEnable sets NonPingPongRoamingEnable field to given value.

### HasNonPingPongRoamingEnable

`func (o *RoamingSettingVO) HasNonPingPongRoamingEnable() bool`

HasNonPingPongRoamingEnable returns a boolean if a field has been set.

### GetNonStickRoamingEnable

`func (o *RoamingSettingVO) GetNonStickRoamingEnable() bool`

GetNonStickRoamingEnable returns the NonStickRoamingEnable field if non-nil, zero value otherwise.

### GetNonStickRoamingEnableOk

`func (o *RoamingSettingVO) GetNonStickRoamingEnableOk() (*bool, bool)`

GetNonStickRoamingEnableOk returns a tuple with the NonStickRoamingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonStickRoamingEnable

`func (o *RoamingSettingVO) SetNonStickRoamingEnable(v bool)`

SetNonStickRoamingEnable sets NonStickRoamingEnable field to given value.

### HasNonStickRoamingEnable

`func (o *RoamingSettingVO) HasNonStickRoamingEnable() bool`

HasNonStickRoamingEnable returns a boolean if a field has been set.

### GetPingPongAssocThreshold

`func (o *RoamingSettingVO) GetPingPongAssocThreshold() int32`

GetPingPongAssocThreshold returns the PingPongAssocThreshold field if non-nil, zero value otherwise.

### GetPingPongAssocThresholdOk

`func (o *RoamingSettingVO) GetPingPongAssocThresholdOk() (*int32, bool)`

GetPingPongAssocThresholdOk returns a tuple with the PingPongAssocThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingPongAssocThreshold

`func (o *RoamingSettingVO) SetPingPongAssocThreshold(v int32)`

SetPingPongAssocThreshold sets PingPongAssocThreshold field to given value.

### HasPingPongAssocThreshold

`func (o *RoamingSettingVO) HasPingPongAssocThreshold() bool`

HasPingPongAssocThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


