# SsidPskSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionPsk** | **int32** | WPA-Personal encryption. This is necessary when the value of security is 3(WPA-Personal); It should be a value as follows: 1: Auto; 3: AES; When versionPsk is WPA3-SAE, Parameter [encryptionPsk] must be AES. | 
**GikRekeyPskEnable** | **bool** | WPA-Personal SSID group key update period config status. True: enable, false: disable. | 
**IntervalPskType** | Pointer to **int32** | WPA-Personal SSID group key update period interval unit config. It should be a value as follows: 0: Seconds; 1: Minutes; 2: Hours. | [optional] 
**RekeyPskInterval** | Pointer to **int32** | WPA-Personal SSID group key update period interval config. When the value of Parameter [intervalPskType] is 0 (Seconds), it should be within the range of 30-86400; when the value of Parameter [intervalPskType] is 1 (Minutes), it should be within the range of 1-1440; when the value of Parameter [intervalPskType] is 2 (Hours), it should be within the range of 1-24. | [optional] 
**SecurityKey** | Pointer to **string** | WPA-Personal SSID password. This is necessary when the value of security is 3(WPA-Personal);It should contain 8-63 printable ASCII characters or 8-63 hexadecimal digits. | [optional] 
**VersionPsk** | **int32** | WPA-Personal version. This is necessary when the value of security is 3(WPA-Personal); It should be a value as follows: 1: WPA-PSK; 2: WPA2-PSK; 3: WPA/WPA2-PSK; 4: WPA3-SAE. | 

## Methods

### NewSsidPskSettingOpenApiVO

`func NewSsidPskSettingOpenApiVO(encryptionPsk int32, gikRekeyPskEnable bool, versionPsk int32, ) *SsidPskSettingOpenApiVO`

NewSsidPskSettingOpenApiVO instantiates a new SsidPskSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsidPskSettingOpenApiVOWithDefaults

`func NewSsidPskSettingOpenApiVOWithDefaults() *SsidPskSettingOpenApiVO`

NewSsidPskSettingOpenApiVOWithDefaults instantiates a new SsidPskSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptionPsk

`func (o *SsidPskSettingOpenApiVO) GetEncryptionPsk() int32`

GetEncryptionPsk returns the EncryptionPsk field if non-nil, zero value otherwise.

### GetEncryptionPskOk

`func (o *SsidPskSettingOpenApiVO) GetEncryptionPskOk() (*int32, bool)`

GetEncryptionPskOk returns a tuple with the EncryptionPsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionPsk

`func (o *SsidPskSettingOpenApiVO) SetEncryptionPsk(v int32)`

SetEncryptionPsk sets EncryptionPsk field to given value.


### GetGikRekeyPskEnable

`func (o *SsidPskSettingOpenApiVO) GetGikRekeyPskEnable() bool`

GetGikRekeyPskEnable returns the GikRekeyPskEnable field if non-nil, zero value otherwise.

### GetGikRekeyPskEnableOk

`func (o *SsidPskSettingOpenApiVO) GetGikRekeyPskEnableOk() (*bool, bool)`

GetGikRekeyPskEnableOk returns a tuple with the GikRekeyPskEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGikRekeyPskEnable

`func (o *SsidPskSettingOpenApiVO) SetGikRekeyPskEnable(v bool)`

SetGikRekeyPskEnable sets GikRekeyPskEnable field to given value.


### GetIntervalPskType

`func (o *SsidPskSettingOpenApiVO) GetIntervalPskType() int32`

GetIntervalPskType returns the IntervalPskType field if non-nil, zero value otherwise.

### GetIntervalPskTypeOk

`func (o *SsidPskSettingOpenApiVO) GetIntervalPskTypeOk() (*int32, bool)`

GetIntervalPskTypeOk returns a tuple with the IntervalPskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalPskType

`func (o *SsidPskSettingOpenApiVO) SetIntervalPskType(v int32)`

SetIntervalPskType sets IntervalPskType field to given value.

### HasIntervalPskType

`func (o *SsidPskSettingOpenApiVO) HasIntervalPskType() bool`

HasIntervalPskType returns a boolean if a field has been set.

### GetRekeyPskInterval

`func (o *SsidPskSettingOpenApiVO) GetRekeyPskInterval() int32`

GetRekeyPskInterval returns the RekeyPskInterval field if non-nil, zero value otherwise.

### GetRekeyPskIntervalOk

`func (o *SsidPskSettingOpenApiVO) GetRekeyPskIntervalOk() (*int32, bool)`

GetRekeyPskIntervalOk returns a tuple with the RekeyPskInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRekeyPskInterval

`func (o *SsidPskSettingOpenApiVO) SetRekeyPskInterval(v int32)`

SetRekeyPskInterval sets RekeyPskInterval field to given value.

### HasRekeyPskInterval

`func (o *SsidPskSettingOpenApiVO) HasRekeyPskInterval() bool`

HasRekeyPskInterval returns a boolean if a field has been set.

### GetSecurityKey

`func (o *SsidPskSettingOpenApiVO) GetSecurityKey() string`

GetSecurityKey returns the SecurityKey field if non-nil, zero value otherwise.

### GetSecurityKeyOk

`func (o *SsidPskSettingOpenApiVO) GetSecurityKeyOk() (*string, bool)`

GetSecurityKeyOk returns a tuple with the SecurityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityKey

`func (o *SsidPskSettingOpenApiVO) SetSecurityKey(v string)`

SetSecurityKey sets SecurityKey field to given value.

### HasSecurityKey

`func (o *SsidPskSettingOpenApiVO) HasSecurityKey() bool`

HasSecurityKey returns a boolean if a field has been set.

### GetVersionPsk

`func (o *SsidPskSettingOpenApiVO) GetVersionPsk() int32`

GetVersionPsk returns the VersionPsk field if non-nil, zero value otherwise.

### GetVersionPskOk

`func (o *SsidPskSettingOpenApiVO) GetVersionPskOk() (*int32, bool)`

GetVersionPskOk returns a tuple with the VersionPsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionPsk

`func (o *SsidPskSettingOpenApiVO) SetVersionPsk(v int32)`

SetVersionPsk sets VersionPsk field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


