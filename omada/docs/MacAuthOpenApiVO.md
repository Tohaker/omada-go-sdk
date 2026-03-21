# MacAuthOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmptyPwdEnable** | Pointer to **bool** | Whether to enable empty password | [optional] 
**Enable** | Pointer to **bool** | MAC-Based Authentication enable status | [optional] 
**MabEnable** | Pointer to **bool** | MAB enable status | [optional] 
**MacFormat** | Pointer to **int32** | Format of the MAC address. MacFormat should be a value as follows: 0: aabbccddeeff, 1: aa-bb-cc-dd-ee-ff, 2: aa:bb:cc:dd:ee:ff, 3: AABBCCDDEEFF, 4: AA-BB-CC-DD-EE-FF, 5: AA:BB:CC:DD:EE:FF | [optional] 
**NasId** | Pointer to **string** | NAS ID issued to AP. NasId should contain 1 to 64 characters | [optional] 
**RadiusProfileId** | Pointer to **string** | This field represents radius profile ID. Radius profile can be created using &#39;Create a new Radius profile&#39; (&#39;Create a new Radius profile template&#39;) interface, and radius profile ID can be obtained from &#39;Get Radius profile list&#39; (&#39;Get Radius profile template list&#39;) interface | [optional] 
**Ssids** | Pointer to [**[]SsidSimpleOpenApiVO**](SsidSimpleOpenApiVO.md) | SSID list with MAC-Based authentication configured | [optional] 

## Methods

### NewMacAuthOpenApiVO

`func NewMacAuthOpenApiVO() *MacAuthOpenApiVO`

NewMacAuthOpenApiVO instantiates a new MacAuthOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacAuthOpenApiVOWithDefaults

`func NewMacAuthOpenApiVOWithDefaults() *MacAuthOpenApiVO`

NewMacAuthOpenApiVOWithDefaults instantiates a new MacAuthOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmptyPwdEnable

`func (o *MacAuthOpenApiVO) GetEmptyPwdEnable() bool`

GetEmptyPwdEnable returns the EmptyPwdEnable field if non-nil, zero value otherwise.

### GetEmptyPwdEnableOk

`func (o *MacAuthOpenApiVO) GetEmptyPwdEnableOk() (*bool, bool)`

GetEmptyPwdEnableOk returns a tuple with the EmptyPwdEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyPwdEnable

`func (o *MacAuthOpenApiVO) SetEmptyPwdEnable(v bool)`

SetEmptyPwdEnable sets EmptyPwdEnable field to given value.

### HasEmptyPwdEnable

`func (o *MacAuthOpenApiVO) HasEmptyPwdEnable() bool`

HasEmptyPwdEnable returns a boolean if a field has been set.

### GetEnable

`func (o *MacAuthOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *MacAuthOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *MacAuthOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *MacAuthOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetMabEnable

`func (o *MacAuthOpenApiVO) GetMabEnable() bool`

GetMabEnable returns the MabEnable field if non-nil, zero value otherwise.

### GetMabEnableOk

`func (o *MacAuthOpenApiVO) GetMabEnableOk() (*bool, bool)`

GetMabEnableOk returns a tuple with the MabEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMabEnable

`func (o *MacAuthOpenApiVO) SetMabEnable(v bool)`

SetMabEnable sets MabEnable field to given value.

### HasMabEnable

`func (o *MacAuthOpenApiVO) HasMabEnable() bool`

HasMabEnable returns a boolean if a field has been set.

### GetMacFormat

`func (o *MacAuthOpenApiVO) GetMacFormat() int32`

GetMacFormat returns the MacFormat field if non-nil, zero value otherwise.

### GetMacFormatOk

`func (o *MacAuthOpenApiVO) GetMacFormatOk() (*int32, bool)`

GetMacFormatOk returns a tuple with the MacFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacFormat

`func (o *MacAuthOpenApiVO) SetMacFormat(v int32)`

SetMacFormat sets MacFormat field to given value.

### HasMacFormat

`func (o *MacAuthOpenApiVO) HasMacFormat() bool`

HasMacFormat returns a boolean if a field has been set.

### GetNasId

`func (o *MacAuthOpenApiVO) GetNasId() string`

GetNasId returns the NasId field if non-nil, zero value otherwise.

### GetNasIdOk

`func (o *MacAuthOpenApiVO) GetNasIdOk() (*string, bool)`

GetNasIdOk returns a tuple with the NasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasId

`func (o *MacAuthOpenApiVO) SetNasId(v string)`

SetNasId sets NasId field to given value.

### HasNasId

`func (o *MacAuthOpenApiVO) HasNasId() bool`

HasNasId returns a boolean if a field has been set.

### GetRadiusProfileId

`func (o *MacAuthOpenApiVO) GetRadiusProfileId() string`

GetRadiusProfileId returns the RadiusProfileId field if non-nil, zero value otherwise.

### GetRadiusProfileIdOk

`func (o *MacAuthOpenApiVO) GetRadiusProfileIdOk() (*string, bool)`

GetRadiusProfileIdOk returns a tuple with the RadiusProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusProfileId

`func (o *MacAuthOpenApiVO) SetRadiusProfileId(v string)`

SetRadiusProfileId sets RadiusProfileId field to given value.

### HasRadiusProfileId

`func (o *MacAuthOpenApiVO) HasRadiusProfileId() bool`

HasRadiusProfileId returns a boolean if a field has been set.

### GetSsids

`func (o *MacAuthOpenApiVO) GetSsids() []SsidSimpleOpenApiVO`

GetSsids returns the Ssids field if non-nil, zero value otherwise.

### GetSsidsOk

`func (o *MacAuthOpenApiVO) GetSsidsOk() (*[]SsidSimpleOpenApiVO, bool)`

GetSsidsOk returns a tuple with the Ssids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsids

`func (o *MacAuthOpenApiVO) SetSsids(v []SsidSimpleOpenApiVO)`

SetSsids sets Ssids field to given value.

### HasSsids

`func (o *MacAuthOpenApiVO) HasSsids() bool`

HasSsids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


