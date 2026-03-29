# NewRoamingSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiRoamingEnable** | **bool** | Whether to enable AI roaming, this configuration will take effect only when fast roaming is enabled | 
**FastRoamingEnable** | **bool** | Whether to enable fast roaming | 
**NonPingPongRoamingEnable** | Pointer to **bool** | Whether to enable non-ping-pong roaming | [optional] 
**NonStickRoamingEnable** | Pointer to **bool** | Whether to enable non-stick roaming | [optional] 
**PingPongAssocThreshold** | Pointer to **int32** | Association frequency threshold of ping-pong roaming | [optional] 

## Methods

### NewNewRoamingSettingOpenApiVO

`func NewNewRoamingSettingOpenApiVO(aiRoamingEnable bool, fastRoamingEnable bool, ) *NewRoamingSettingOpenApiVO`

NewNewRoamingSettingOpenApiVO instantiates a new NewRoamingSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewRoamingSettingOpenApiVOWithDefaults

`func NewNewRoamingSettingOpenApiVOWithDefaults() *NewRoamingSettingOpenApiVO`

NewNewRoamingSettingOpenApiVOWithDefaults instantiates a new NewRoamingSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiRoamingEnable

`func (o *NewRoamingSettingOpenApiVO) GetAiRoamingEnable() bool`

GetAiRoamingEnable returns the AiRoamingEnable field if non-nil, zero value otherwise.

### GetAiRoamingEnableOk

`func (o *NewRoamingSettingOpenApiVO) GetAiRoamingEnableOk() (*bool, bool)`

GetAiRoamingEnableOk returns a tuple with the AiRoamingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiRoamingEnable

`func (o *NewRoamingSettingOpenApiVO) SetAiRoamingEnable(v bool)`

SetAiRoamingEnable sets AiRoamingEnable field to given value.


### GetFastRoamingEnable

`func (o *NewRoamingSettingOpenApiVO) GetFastRoamingEnable() bool`

GetFastRoamingEnable returns the FastRoamingEnable field if non-nil, zero value otherwise.

### GetFastRoamingEnableOk

`func (o *NewRoamingSettingOpenApiVO) GetFastRoamingEnableOk() (*bool, bool)`

GetFastRoamingEnableOk returns a tuple with the FastRoamingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastRoamingEnable

`func (o *NewRoamingSettingOpenApiVO) SetFastRoamingEnable(v bool)`

SetFastRoamingEnable sets FastRoamingEnable field to given value.


### GetNonPingPongRoamingEnable

`func (o *NewRoamingSettingOpenApiVO) GetNonPingPongRoamingEnable() bool`

GetNonPingPongRoamingEnable returns the NonPingPongRoamingEnable field if non-nil, zero value otherwise.

### GetNonPingPongRoamingEnableOk

`func (o *NewRoamingSettingOpenApiVO) GetNonPingPongRoamingEnableOk() (*bool, bool)`

GetNonPingPongRoamingEnableOk returns a tuple with the NonPingPongRoamingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonPingPongRoamingEnable

`func (o *NewRoamingSettingOpenApiVO) SetNonPingPongRoamingEnable(v bool)`

SetNonPingPongRoamingEnable sets NonPingPongRoamingEnable field to given value.

### HasNonPingPongRoamingEnable

`func (o *NewRoamingSettingOpenApiVO) HasNonPingPongRoamingEnable() bool`

HasNonPingPongRoamingEnable returns a boolean if a field has been set.

### GetNonStickRoamingEnable

`func (o *NewRoamingSettingOpenApiVO) GetNonStickRoamingEnable() bool`

GetNonStickRoamingEnable returns the NonStickRoamingEnable field if non-nil, zero value otherwise.

### GetNonStickRoamingEnableOk

`func (o *NewRoamingSettingOpenApiVO) GetNonStickRoamingEnableOk() (*bool, bool)`

GetNonStickRoamingEnableOk returns a tuple with the NonStickRoamingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonStickRoamingEnable

`func (o *NewRoamingSettingOpenApiVO) SetNonStickRoamingEnable(v bool)`

SetNonStickRoamingEnable sets NonStickRoamingEnable field to given value.

### HasNonStickRoamingEnable

`func (o *NewRoamingSettingOpenApiVO) HasNonStickRoamingEnable() bool`

HasNonStickRoamingEnable returns a boolean if a field has been set.

### GetPingPongAssocThreshold

`func (o *NewRoamingSettingOpenApiVO) GetPingPongAssocThreshold() int32`

GetPingPongAssocThreshold returns the PingPongAssocThreshold field if non-nil, zero value otherwise.

### GetPingPongAssocThresholdOk

`func (o *NewRoamingSettingOpenApiVO) GetPingPongAssocThresholdOk() (*int32, bool)`

GetPingPongAssocThresholdOk returns a tuple with the PingPongAssocThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingPongAssocThreshold

`func (o *NewRoamingSettingOpenApiVO) SetPingPongAssocThreshold(v int32)`

SetPingPongAssocThreshold sets PingPongAssocThreshold field to given value.

### HasPingPongAssocThreshold

`func (o *NewRoamingSettingOpenApiVO) HasPingPongAssocThreshold() bool`

HasPingPongAssocThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


