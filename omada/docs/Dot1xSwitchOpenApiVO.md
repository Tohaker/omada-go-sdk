# Dot1xSwitchOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | Pointer to **int32** | Mode of authentication protocol. AuthMode should be a value as follows: 0: PAP, 1: EAP | [optional] 
**AuthType** | Pointer to **int32** | Type of the authentication. AuthType should be a value as follows: 0: MAC based, 1: Port based | [optional] 
**Enable** | **bool** | Switch 802.1x function enable status | 
**Mab** | Pointer to **bool** | MAB enable status | [optional] 
**MacFormat** | Pointer to **int32** | Format of the MAC address. MacFormat should be a value as follows: 0: aabbccddeeff, 1: aa-bb-cc-dd-ee-ff, 2: aa:bb:cc:dd:ee:ff, 3: AABBCCDDEEFF, 4: AA-BB-CC-DD-EE-FF, 5: AA:BB:CC:DD:EE:FF | [optional] 
**RadiusProfileId** | Pointer to **string** | This field represents radius profile ID. Radius profile can be created using &#39;Create a new Radius profile&#39; (&#39;Create a new Radius profile template&#39;) interface, and radius profile ID can be obtained from &#39;Get Radius profile list&#39; (&#39;Get Radius profile template list&#39;) interface | [optional] 
**Switches** | Pointer to [**[]Dot1xSwitchSettingOpenApiVO**](Dot1xSwitchSettingOpenApiVO.md) | Enabled switch ports, optional when update switch 802.1x setting | [optional] 
**VlanAssign** | Pointer to **bool** | VLAN Assignment enable status | [optional] 

## Methods

### NewDot1xSwitchOpenApiVO

`func NewDot1xSwitchOpenApiVO(enable bool, ) *Dot1xSwitchOpenApiVO`

NewDot1xSwitchOpenApiVO instantiates a new Dot1xSwitchOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDot1xSwitchOpenApiVOWithDefaults

`func NewDot1xSwitchOpenApiVOWithDefaults() *Dot1xSwitchOpenApiVO`

NewDot1xSwitchOpenApiVOWithDefaults instantiates a new Dot1xSwitchOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *Dot1xSwitchOpenApiVO) GetAuthMode() int32`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *Dot1xSwitchOpenApiVO) GetAuthModeOk() (*int32, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *Dot1xSwitchOpenApiVO) SetAuthMode(v int32)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *Dot1xSwitchOpenApiVO) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAuthType

`func (o *Dot1xSwitchOpenApiVO) GetAuthType() int32`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *Dot1xSwitchOpenApiVO) GetAuthTypeOk() (*int32, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *Dot1xSwitchOpenApiVO) SetAuthType(v int32)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *Dot1xSwitchOpenApiVO) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetEnable

`func (o *Dot1xSwitchOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *Dot1xSwitchOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *Dot1xSwitchOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetMab

`func (o *Dot1xSwitchOpenApiVO) GetMab() bool`

GetMab returns the Mab field if non-nil, zero value otherwise.

### GetMabOk

`func (o *Dot1xSwitchOpenApiVO) GetMabOk() (*bool, bool)`

GetMabOk returns a tuple with the Mab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMab

`func (o *Dot1xSwitchOpenApiVO) SetMab(v bool)`

SetMab sets Mab field to given value.

### HasMab

`func (o *Dot1xSwitchOpenApiVO) HasMab() bool`

HasMab returns a boolean if a field has been set.

### GetMacFormat

`func (o *Dot1xSwitchOpenApiVO) GetMacFormat() int32`

GetMacFormat returns the MacFormat field if non-nil, zero value otherwise.

### GetMacFormatOk

`func (o *Dot1xSwitchOpenApiVO) GetMacFormatOk() (*int32, bool)`

GetMacFormatOk returns a tuple with the MacFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacFormat

`func (o *Dot1xSwitchOpenApiVO) SetMacFormat(v int32)`

SetMacFormat sets MacFormat field to given value.

### HasMacFormat

`func (o *Dot1xSwitchOpenApiVO) HasMacFormat() bool`

HasMacFormat returns a boolean if a field has been set.

### GetRadiusProfileId

`func (o *Dot1xSwitchOpenApiVO) GetRadiusProfileId() string`

GetRadiusProfileId returns the RadiusProfileId field if non-nil, zero value otherwise.

### GetRadiusProfileIdOk

`func (o *Dot1xSwitchOpenApiVO) GetRadiusProfileIdOk() (*string, bool)`

GetRadiusProfileIdOk returns a tuple with the RadiusProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusProfileId

`func (o *Dot1xSwitchOpenApiVO) SetRadiusProfileId(v string)`

SetRadiusProfileId sets RadiusProfileId field to given value.

### HasRadiusProfileId

`func (o *Dot1xSwitchOpenApiVO) HasRadiusProfileId() bool`

HasRadiusProfileId returns a boolean if a field has been set.

### GetSwitches

`func (o *Dot1xSwitchOpenApiVO) GetSwitches() []Dot1xSwitchSettingOpenApiVO`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *Dot1xSwitchOpenApiVO) GetSwitchesOk() (*[]Dot1xSwitchSettingOpenApiVO, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *Dot1xSwitchOpenApiVO) SetSwitches(v []Dot1xSwitchSettingOpenApiVO)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *Dot1xSwitchOpenApiVO) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.

### GetVlanAssign

`func (o *Dot1xSwitchOpenApiVO) GetVlanAssign() bool`

GetVlanAssign returns the VlanAssign field if non-nil, zero value otherwise.

### GetVlanAssignOk

`func (o *Dot1xSwitchOpenApiVO) GetVlanAssignOk() (*bool, bool)`

GetVlanAssignOk returns a tuple with the VlanAssign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanAssign

`func (o *Dot1xSwitchOpenApiVO) SetVlanAssign(v bool)`

SetVlanAssign sets VlanAssign field to given value.

### HasVlanAssign

`func (o *Dot1xSwitchOpenApiVO) HasVlanAssign() bool`

HasVlanAssign returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


