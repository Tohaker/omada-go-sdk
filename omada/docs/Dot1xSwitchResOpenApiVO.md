# Dot1xSwitchResOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | Pointer to **int32** | Mode of authentication protocol. AuthMode should be a value as follows: 0: PAP, 1: EAP | [optional] 
**AuthType** | Pointer to **int32** | Type of the authentication. AuthType should be a value as follows: 0: MAC based, 1: Port based | [optional] 
**Enable** | **bool** | Switch 802.1x function enable status | 
**Mab** | Pointer to **bool** | MAB enable status | [optional] 
**MacFormat** | Pointer to **int32** | Format of the MAC address. MacFormat should be a value as follows: 0: aabbccddeeff, 1: aa-bb-cc-dd-ee-ff, 2: aa:bb:cc:dd:ee:ff, 3: AABBCCDDEEFF, 4: AA-BB-CC-DD-EE-FF, 5: AA:BB:CC:DD:EE:FF | [optional] 
**RadiusProfileId** | Pointer to **string** | This field represents radius profile ID. Radius profile can be created using &#39;Create a new Radius profile&#39; (&#39;Create a new Radius profile template&#39;) interface, and radius profile ID can be obtained from &#39;Get Radius profile list&#39; (&#39;Get Radius profile template list&#39;) interface | [optional] 
**VlanAssign** | Pointer to **bool** | VLAN Assignment enable status | [optional] 

## Methods

### NewDot1xSwitchResOpenApiVO

`func NewDot1xSwitchResOpenApiVO(enable bool, ) *Dot1xSwitchResOpenApiVO`

NewDot1xSwitchResOpenApiVO instantiates a new Dot1xSwitchResOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDot1xSwitchResOpenApiVOWithDefaults

`func NewDot1xSwitchResOpenApiVOWithDefaults() *Dot1xSwitchResOpenApiVO`

NewDot1xSwitchResOpenApiVOWithDefaults instantiates a new Dot1xSwitchResOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *Dot1xSwitchResOpenApiVO) GetAuthMode() int32`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *Dot1xSwitchResOpenApiVO) GetAuthModeOk() (*int32, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *Dot1xSwitchResOpenApiVO) SetAuthMode(v int32)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *Dot1xSwitchResOpenApiVO) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAuthType

`func (o *Dot1xSwitchResOpenApiVO) GetAuthType() int32`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *Dot1xSwitchResOpenApiVO) GetAuthTypeOk() (*int32, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *Dot1xSwitchResOpenApiVO) SetAuthType(v int32)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *Dot1xSwitchResOpenApiVO) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetEnable

`func (o *Dot1xSwitchResOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *Dot1xSwitchResOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *Dot1xSwitchResOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetMab

`func (o *Dot1xSwitchResOpenApiVO) GetMab() bool`

GetMab returns the Mab field if non-nil, zero value otherwise.

### GetMabOk

`func (o *Dot1xSwitchResOpenApiVO) GetMabOk() (*bool, bool)`

GetMabOk returns a tuple with the Mab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMab

`func (o *Dot1xSwitchResOpenApiVO) SetMab(v bool)`

SetMab sets Mab field to given value.

### HasMab

`func (o *Dot1xSwitchResOpenApiVO) HasMab() bool`

HasMab returns a boolean if a field has been set.

### GetMacFormat

`func (o *Dot1xSwitchResOpenApiVO) GetMacFormat() int32`

GetMacFormat returns the MacFormat field if non-nil, zero value otherwise.

### GetMacFormatOk

`func (o *Dot1xSwitchResOpenApiVO) GetMacFormatOk() (*int32, bool)`

GetMacFormatOk returns a tuple with the MacFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacFormat

`func (o *Dot1xSwitchResOpenApiVO) SetMacFormat(v int32)`

SetMacFormat sets MacFormat field to given value.

### HasMacFormat

`func (o *Dot1xSwitchResOpenApiVO) HasMacFormat() bool`

HasMacFormat returns a boolean if a field has been set.

### GetRadiusProfileId

`func (o *Dot1xSwitchResOpenApiVO) GetRadiusProfileId() string`

GetRadiusProfileId returns the RadiusProfileId field if non-nil, zero value otherwise.

### GetRadiusProfileIdOk

`func (o *Dot1xSwitchResOpenApiVO) GetRadiusProfileIdOk() (*string, bool)`

GetRadiusProfileIdOk returns a tuple with the RadiusProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusProfileId

`func (o *Dot1xSwitchResOpenApiVO) SetRadiusProfileId(v string)`

SetRadiusProfileId sets RadiusProfileId field to given value.

### HasRadiusProfileId

`func (o *Dot1xSwitchResOpenApiVO) HasRadiusProfileId() bool`

HasRadiusProfileId returns a boolean if a field has been set.

### GetVlanAssign

`func (o *Dot1xSwitchResOpenApiVO) GetVlanAssign() bool`

GetVlanAssign returns the VlanAssign field if non-nil, zero value otherwise.

### GetVlanAssignOk

`func (o *Dot1xSwitchResOpenApiVO) GetVlanAssignOk() (*bool, bool)`

GetVlanAssignOk returns a tuple with the VlanAssign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanAssign

`func (o *Dot1xSwitchResOpenApiVO) SetVlanAssign(v bool)`

SetVlanAssign sets VlanAssign field to given value.

### HasVlanAssign

`func (o *Dot1xSwitchResOpenApiVO) HasVlanAssign() bool`

HasVlanAssign returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


