# SsidPskSettingForIpptOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionPsk** | **int32** | WPA-Personal encryption. This is necessary when the value of security is 3(WPA-Personal); It should be a value as follows: 1: Auto; 3: AES; When versionPsk is WPA3-SAE, Parameter [encryptionPsk] must be AES. | 
**SecurityKey** | Pointer to **string** | WPA-Personal SSID password. This is necessary when the value of security is 3(WPA-Personal);It should contain 8-63 printable ASCII characters or 8-63 hexadecimal digits. | [optional] 
**VersionPsk** | **int32** | WPA-Personal version. This is necessary when the value of security is 3(WPA-Personal); It should be a value as follows: 1: WPA-PSK; 2: WPA2-PSK; 3: WPA/WPA2-PSK; 4: WPA3-SAE. | 

## Methods

### NewSsidPskSettingForIpptOpenApiVO

`func NewSsidPskSettingForIpptOpenApiVO(encryptionPsk int32, versionPsk int32, ) *SsidPskSettingForIpptOpenApiVO`

NewSsidPskSettingForIpptOpenApiVO instantiates a new SsidPskSettingForIpptOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsidPskSettingForIpptOpenApiVOWithDefaults

`func NewSsidPskSettingForIpptOpenApiVOWithDefaults() *SsidPskSettingForIpptOpenApiVO`

NewSsidPskSettingForIpptOpenApiVOWithDefaults instantiates a new SsidPskSettingForIpptOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptionPsk

`func (o *SsidPskSettingForIpptOpenApiVO) GetEncryptionPsk() int32`

GetEncryptionPsk returns the EncryptionPsk field if non-nil, zero value otherwise.

### GetEncryptionPskOk

`func (o *SsidPskSettingForIpptOpenApiVO) GetEncryptionPskOk() (*int32, bool)`

GetEncryptionPskOk returns a tuple with the EncryptionPsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionPsk

`func (o *SsidPskSettingForIpptOpenApiVO) SetEncryptionPsk(v int32)`

SetEncryptionPsk sets EncryptionPsk field to given value.


### GetSecurityKey

`func (o *SsidPskSettingForIpptOpenApiVO) GetSecurityKey() string`

GetSecurityKey returns the SecurityKey field if non-nil, zero value otherwise.

### GetSecurityKeyOk

`func (o *SsidPskSettingForIpptOpenApiVO) GetSecurityKeyOk() (*string, bool)`

GetSecurityKeyOk returns a tuple with the SecurityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityKey

`func (o *SsidPskSettingForIpptOpenApiVO) SetSecurityKey(v string)`

SetSecurityKey sets SecurityKey field to given value.

### HasSecurityKey

`func (o *SsidPskSettingForIpptOpenApiVO) HasSecurityKey() bool`

HasSecurityKey returns a boolean if a field has been set.

### GetVersionPsk

`func (o *SsidPskSettingForIpptOpenApiVO) GetVersionPsk() int32`

GetVersionPsk returns the VersionPsk field if non-nil, zero value otherwise.

### GetVersionPskOk

`func (o *SsidPskSettingForIpptOpenApiVO) GetVersionPskOk() (*int32, bool)`

GetVersionPskOk returns a tuple with the VersionPsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionPsk

`func (o *SsidPskSettingForIpptOpenApiVO) SetVersionPsk(v int32)`

SetVersionPsk sets VersionPsk field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


