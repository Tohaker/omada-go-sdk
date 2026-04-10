# ApMgtSsidEnterpriseSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionEnt** | **int32** | WPA-Enterprise encryption. This is necessary when the value of security is 2(WPA-Enterprise); It should be a value as follows: 1: Auto; 3: AES. | 
**RadiusProfileId** | **string** | This field represents RADIUS Profile ID. RADIUS Profile(RADIUS Profile Template) can be created using Create a new Radius profile(Create a new Radius profile template) interface, and RADIUS Profile ID(RADIUS Profile Template ID) can be obtained from Get Radius profile list(Get Radius profile template list) interface. | 
**VersionEnt** | **int32** | WPA-Enterprise version. This is necessary when the value of security is 2(WPA-Enterprise); It should be a value as follows: 1: WPA-Enterprise; 2: WPA2-Enterprise; 3: WPA/WPA2-Enterprise. | 

## Methods

### NewApMgtSsidEnterpriseSettingOpenApiVO

`func NewApMgtSsidEnterpriseSettingOpenApiVO(encryptionEnt int32, radiusProfileId string, versionEnt int32, ) *ApMgtSsidEnterpriseSettingOpenApiVO`

NewApMgtSsidEnterpriseSettingOpenApiVO instantiates a new ApMgtSsidEnterpriseSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApMgtSsidEnterpriseSettingOpenApiVOWithDefaults

`func NewApMgtSsidEnterpriseSettingOpenApiVOWithDefaults() *ApMgtSsidEnterpriseSettingOpenApiVO`

NewApMgtSsidEnterpriseSettingOpenApiVOWithDefaults instantiates a new ApMgtSsidEnterpriseSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptionEnt

`func (o *ApMgtSsidEnterpriseSettingOpenApiVO) GetEncryptionEnt() int32`

GetEncryptionEnt returns the EncryptionEnt field if non-nil, zero value otherwise.

### GetEncryptionEntOk

`func (o *ApMgtSsidEnterpriseSettingOpenApiVO) GetEncryptionEntOk() (*int32, bool)`

GetEncryptionEntOk returns a tuple with the EncryptionEnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnt

`func (o *ApMgtSsidEnterpriseSettingOpenApiVO) SetEncryptionEnt(v int32)`

SetEncryptionEnt sets EncryptionEnt field to given value.


### GetRadiusProfileId

`func (o *ApMgtSsidEnterpriseSettingOpenApiVO) GetRadiusProfileId() string`

GetRadiusProfileId returns the RadiusProfileId field if non-nil, zero value otherwise.

### GetRadiusProfileIdOk

`func (o *ApMgtSsidEnterpriseSettingOpenApiVO) GetRadiusProfileIdOk() (*string, bool)`

GetRadiusProfileIdOk returns a tuple with the RadiusProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusProfileId

`func (o *ApMgtSsidEnterpriseSettingOpenApiVO) SetRadiusProfileId(v string)`

SetRadiusProfileId sets RadiusProfileId field to given value.


### GetVersionEnt

`func (o *ApMgtSsidEnterpriseSettingOpenApiVO) GetVersionEnt() int32`

GetVersionEnt returns the VersionEnt field if non-nil, zero value otherwise.

### GetVersionEntOk

`func (o *ApMgtSsidEnterpriseSettingOpenApiVO) GetVersionEntOk() (*int32, bool)`

GetVersionEntOk returns a tuple with the VersionEnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionEnt

`func (o *ApMgtSsidEnterpriseSettingOpenApiVO) SetVersionEnt(v int32)`

SetVersionEnt sets VersionEnt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


