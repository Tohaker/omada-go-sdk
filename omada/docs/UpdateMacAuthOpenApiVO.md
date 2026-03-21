# UpdateMacAuthOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmptyPwdEnable** | Pointer to **bool** | Whether to enable empty password | [optional] 
**Enable** | **bool** | MAC-Based Authentication enable status | 
**MabEnable** | Pointer to **bool** | MAB enable status | [optional] 
**MacFormat** | Pointer to **int32** | Format of the MAC address. MacFormat should be a value as follows: 0: aabbccddeeff, 1: aa-bb-cc-dd-ee-ff, 2: aa:bb:cc:dd:ee:ff, 3: AABBCCDDEEFF, 4: AA-BB-CC-DD-EE-FF, 5: AA:BB:CC:DD:EE:FF | [optional] 
**NasId** | Pointer to **string** | NAS ID issued to AP. NasId should contain 1 to 64 characters | [optional] 
**RadiusProfileId** | Pointer to **string** | This field represents radius profile ID. Radius profile can be created using &#39;Create a new Radius profile&#39; (&#39;Create a new Radius profile template&#39;) interface, and radius profile ID can be obtained from &#39;Get Radius profile list&#39; (&#39;Get Radius profile template list&#39;) interface | [optional] 
**Ssids** | **[]string** | SSID ID list with MAC-Based authentication configured. SSID can be created using &#39;Create new SSID&#39; (&#39;Create new SSID template&#39;) interface, and SSID ID can be obtained from &#39;Get SSID list&#39; (&#39;Get SSID template list&#39;) interface | 

## Methods

### NewUpdateMacAuthOpenApiVO

`func NewUpdateMacAuthOpenApiVO(enable bool, ssids []string, ) *UpdateMacAuthOpenApiVO`

NewUpdateMacAuthOpenApiVO instantiates a new UpdateMacAuthOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMacAuthOpenApiVOWithDefaults

`func NewUpdateMacAuthOpenApiVOWithDefaults() *UpdateMacAuthOpenApiVO`

NewUpdateMacAuthOpenApiVOWithDefaults instantiates a new UpdateMacAuthOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmptyPwdEnable

`func (o *UpdateMacAuthOpenApiVO) GetEmptyPwdEnable() bool`

GetEmptyPwdEnable returns the EmptyPwdEnable field if non-nil, zero value otherwise.

### GetEmptyPwdEnableOk

`func (o *UpdateMacAuthOpenApiVO) GetEmptyPwdEnableOk() (*bool, bool)`

GetEmptyPwdEnableOk returns a tuple with the EmptyPwdEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyPwdEnable

`func (o *UpdateMacAuthOpenApiVO) SetEmptyPwdEnable(v bool)`

SetEmptyPwdEnable sets EmptyPwdEnable field to given value.

### HasEmptyPwdEnable

`func (o *UpdateMacAuthOpenApiVO) HasEmptyPwdEnable() bool`

HasEmptyPwdEnable returns a boolean if a field has been set.

### GetEnable

`func (o *UpdateMacAuthOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *UpdateMacAuthOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *UpdateMacAuthOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetMabEnable

`func (o *UpdateMacAuthOpenApiVO) GetMabEnable() bool`

GetMabEnable returns the MabEnable field if non-nil, zero value otherwise.

### GetMabEnableOk

`func (o *UpdateMacAuthOpenApiVO) GetMabEnableOk() (*bool, bool)`

GetMabEnableOk returns a tuple with the MabEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMabEnable

`func (o *UpdateMacAuthOpenApiVO) SetMabEnable(v bool)`

SetMabEnable sets MabEnable field to given value.

### HasMabEnable

`func (o *UpdateMacAuthOpenApiVO) HasMabEnable() bool`

HasMabEnable returns a boolean if a field has been set.

### GetMacFormat

`func (o *UpdateMacAuthOpenApiVO) GetMacFormat() int32`

GetMacFormat returns the MacFormat field if non-nil, zero value otherwise.

### GetMacFormatOk

`func (o *UpdateMacAuthOpenApiVO) GetMacFormatOk() (*int32, bool)`

GetMacFormatOk returns a tuple with the MacFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacFormat

`func (o *UpdateMacAuthOpenApiVO) SetMacFormat(v int32)`

SetMacFormat sets MacFormat field to given value.

### HasMacFormat

`func (o *UpdateMacAuthOpenApiVO) HasMacFormat() bool`

HasMacFormat returns a boolean if a field has been set.

### GetNasId

`func (o *UpdateMacAuthOpenApiVO) GetNasId() string`

GetNasId returns the NasId field if non-nil, zero value otherwise.

### GetNasIdOk

`func (o *UpdateMacAuthOpenApiVO) GetNasIdOk() (*string, bool)`

GetNasIdOk returns a tuple with the NasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasId

`func (o *UpdateMacAuthOpenApiVO) SetNasId(v string)`

SetNasId sets NasId field to given value.

### HasNasId

`func (o *UpdateMacAuthOpenApiVO) HasNasId() bool`

HasNasId returns a boolean if a field has been set.

### GetRadiusProfileId

`func (o *UpdateMacAuthOpenApiVO) GetRadiusProfileId() string`

GetRadiusProfileId returns the RadiusProfileId field if non-nil, zero value otherwise.

### GetRadiusProfileIdOk

`func (o *UpdateMacAuthOpenApiVO) GetRadiusProfileIdOk() (*string, bool)`

GetRadiusProfileIdOk returns a tuple with the RadiusProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusProfileId

`func (o *UpdateMacAuthOpenApiVO) SetRadiusProfileId(v string)`

SetRadiusProfileId sets RadiusProfileId field to given value.

### HasRadiusProfileId

`func (o *UpdateMacAuthOpenApiVO) HasRadiusProfileId() bool`

HasRadiusProfileId returns a boolean if a field has been set.

### GetSsids

`func (o *UpdateMacAuthOpenApiVO) GetSsids() []string`

GetSsids returns the Ssids field if non-nil, zero value otherwise.

### GetSsidsOk

`func (o *UpdateMacAuthOpenApiVO) GetSsidsOk() (*[]string, bool)`

GetSsidsOk returns a tuple with the Ssids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsids

`func (o *UpdateMacAuthOpenApiVO) SetSsids(v []string)`

SetSsids sets Ssids field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


