# UpdateSsidBasicConfigForIpptOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoWanAccess** | Pointer to **bool** | Whether to enable auto wan access. True: enable, false: disable. | [optional] 
**Band** | **int32** | SSID band. The lowest bit indicates whether 2.4G is included; the second lowest bit indicates whether 5G is included; the third lowest bit indicates whether 6G is included; 1 means included while 0 means not included. For example, 7(111) means that 2G/5G/6G are enabled; 1(001) means that 2G is enabled. (When 5G is included，it means 5G/5G1/5G2 are enabled.) | 
**Broadcast** | **bool** | SSID broadcast config status. True: enable, false: disable. | 
**Name** | **string** | SSID name. It should contain 1 to 32 UTF-8 characters. | 
**OweEnable** | Pointer to **bool** | Opportunistic Wireless Encryption, also known as Enhanced Open, is a certification provided by the Wi-Fi Alliance as part of the WPA3 wireless security standard. OWE will enable two wireless VAPs per radio, one for access of OWE-supported stations, and one for access of other stations. An SSID with OWE enabled will be counted as two SSID entries for 2G and 5G. Only for security is None and band contains 2.4G or 5G. | [optional] 
**PskSetting** | Pointer to [**SsidPskSettingForIpptOpenApiVO**](SsidPskSettingForIpptOpenApiVO.md) |  | [optional] 
**Security** | **int32** | SSID security mode; Security should be a value as follows: 0: None; 2: WPA-Enterprise; 3: WPA-Personal; 4: PPSK without RADIUS; 5: PPSK with RADIUS. | 

## Methods

### NewUpdateSsidBasicConfigForIpptOpenApiVO

`func NewUpdateSsidBasicConfigForIpptOpenApiVO(band int32, broadcast bool, name string, security int32, ) *UpdateSsidBasicConfigForIpptOpenApiVO`

NewUpdateSsidBasicConfigForIpptOpenApiVO instantiates a new UpdateSsidBasicConfigForIpptOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSsidBasicConfigForIpptOpenApiVOWithDefaults

`func NewUpdateSsidBasicConfigForIpptOpenApiVOWithDefaults() *UpdateSsidBasicConfigForIpptOpenApiVO`

NewUpdateSsidBasicConfigForIpptOpenApiVOWithDefaults instantiates a new UpdateSsidBasicConfigForIpptOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoWanAccess

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) GetAutoWanAccess() bool`

GetAutoWanAccess returns the AutoWanAccess field if non-nil, zero value otherwise.

### GetAutoWanAccessOk

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) GetAutoWanAccessOk() (*bool, bool)`

GetAutoWanAccessOk returns a tuple with the AutoWanAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoWanAccess

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) SetAutoWanAccess(v bool)`

SetAutoWanAccess sets AutoWanAccess field to given value.

### HasAutoWanAccess

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) HasAutoWanAccess() bool`

HasAutoWanAccess returns a boolean if a field has been set.

### GetBand

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) GetBand() int32`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) GetBandOk() (*int32, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) SetBand(v int32)`

SetBand sets Band field to given value.


### GetBroadcast

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) GetBroadcast() bool`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) GetBroadcastOk() (*bool, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) SetBroadcast(v bool)`

SetBroadcast sets Broadcast field to given value.


### GetName

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetOweEnable

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) GetOweEnable() bool`

GetOweEnable returns the OweEnable field if non-nil, zero value otherwise.

### GetOweEnableOk

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) GetOweEnableOk() (*bool, bool)`

GetOweEnableOk returns a tuple with the OweEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOweEnable

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) SetOweEnable(v bool)`

SetOweEnable sets OweEnable field to given value.

### HasOweEnable

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) HasOweEnable() bool`

HasOweEnable returns a boolean if a field has been set.

### GetPskSetting

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) GetPskSetting() SsidPskSettingForIpptOpenApiVO`

GetPskSetting returns the PskSetting field if non-nil, zero value otherwise.

### GetPskSettingOk

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) GetPskSettingOk() (*SsidPskSettingForIpptOpenApiVO, bool)`

GetPskSettingOk returns a tuple with the PskSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPskSetting

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) SetPskSetting(v SsidPskSettingForIpptOpenApiVO)`

SetPskSetting sets PskSetting field to given value.

### HasPskSetting

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) HasPskSetting() bool`

HasPskSetting returns a boolean if a field has been set.

### GetSecurity

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) GetSecurity() int32`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) GetSecurityOk() (*int32, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *UpdateSsidBasicConfigForIpptOpenApiVO) SetSecurity(v int32)`

SetSecurity sets Security field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


