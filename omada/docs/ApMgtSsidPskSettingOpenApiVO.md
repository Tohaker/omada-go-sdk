# ApMgtSsidPskSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionPsk** | **int32** | WPA-Personal encryption. This is necessary when the value of security is 3(WPA-Personal); It should be a value as follows: 1: Auto; 3: AES. | 
**SecurityKey** | **string** | WPA-Personal SSID password. This is necessary when the value of security is 3(WPA-Personal);It should contain 8-63 printable ASCII characters or 8-63 hexadecimal digits. | 
**VersionPsk** | **int32** | WPA-Personal version. This is necessary when the value of security is 3(WPA-Personal); It should be a value as follows: 1: WPA-PSK; 2: WPA2-PSK; 3: WPA/WPA2-PSK. | 

## Methods

### NewApMgtSsidPskSettingOpenApiVO

`func NewApMgtSsidPskSettingOpenApiVO(encryptionPsk int32, securityKey string, versionPsk int32, ) *ApMgtSsidPskSettingOpenApiVO`

NewApMgtSsidPskSettingOpenApiVO instantiates a new ApMgtSsidPskSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApMgtSsidPskSettingOpenApiVOWithDefaults

`func NewApMgtSsidPskSettingOpenApiVOWithDefaults() *ApMgtSsidPskSettingOpenApiVO`

NewApMgtSsidPskSettingOpenApiVOWithDefaults instantiates a new ApMgtSsidPskSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptionPsk

`func (o *ApMgtSsidPskSettingOpenApiVO) GetEncryptionPsk() int32`

GetEncryptionPsk returns the EncryptionPsk field if non-nil, zero value otherwise.

### GetEncryptionPskOk

`func (o *ApMgtSsidPskSettingOpenApiVO) GetEncryptionPskOk() (*int32, bool)`

GetEncryptionPskOk returns a tuple with the EncryptionPsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionPsk

`func (o *ApMgtSsidPskSettingOpenApiVO) SetEncryptionPsk(v int32)`

SetEncryptionPsk sets EncryptionPsk field to given value.


### GetSecurityKey

`func (o *ApMgtSsidPskSettingOpenApiVO) GetSecurityKey() string`

GetSecurityKey returns the SecurityKey field if non-nil, zero value otherwise.

### GetSecurityKeyOk

`func (o *ApMgtSsidPskSettingOpenApiVO) GetSecurityKeyOk() (*string, bool)`

GetSecurityKeyOk returns a tuple with the SecurityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityKey

`func (o *ApMgtSsidPskSettingOpenApiVO) SetSecurityKey(v string)`

SetSecurityKey sets SecurityKey field to given value.


### GetVersionPsk

`func (o *ApMgtSsidPskSettingOpenApiVO) GetVersionPsk() int32`

GetVersionPsk returns the VersionPsk field if non-nil, zero value otherwise.

### GetVersionPskOk

`func (o *ApMgtSsidPskSettingOpenApiVO) GetVersionPskOk() (*int32, bool)`

GetVersionPskOk returns a tuple with the VersionPsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionPsk

`func (o *ApMgtSsidPskSettingOpenApiVO) SetVersionPsk(v int32)`

SetVersionPsk sets VersionPsk field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


