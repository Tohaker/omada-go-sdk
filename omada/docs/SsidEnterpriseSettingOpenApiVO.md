# SsidEnterpriseSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionEnt** | **int32** | WPA-Enterprise encryption. This is necessary when the value of security is 2(WPA-Enterprise); It should be a value as follows: 1: Auto; 3: AES; 4: AES-GCM 256; 5:AES-CNSA; 6:CCMP_128; When versionEnt is WPA3-Enterprise, Parameter [encryptionEnt] must be AES/AES-GCM 256/AES-CNSA/CCMP-128, and if ssid only enable 6G or contains 6G and enable mlo, Parameter [encryptionEnt] can not be AES. | 
**GikRekeyEntEnable** | **bool** | WPA-Enterprise SSID group key update period config status. True: enable, false: disable. | 
**IntervalEntType** | Pointer to **int32** | WPA-Enterprise SSID group key update period interval unit config. It should be a value as follows: 0: Seconds; 1: Minutes; 2: Hours. | [optional] 
**NasId** | Pointer to **string** | This field is necessary when the nasIdMode type is custom. | [optional] 
**NasIdMode** | Pointer to **int32** | Indicates the status of nasid under enterprise-level encryption. It should be a value as follows: 0: default (TP LINK: MAC Address), 1: follow device name, 2: custom. | [optional] 
**RadiusProfileId** | **string** | This field represents RADIUS Profile ID. RADIUS Profile(RADIUS Profile Template) can be created using Create a new Radius profile(Create a new Radius profile template) interface, and RADIUS Profile ID(RADIUS Profile Template ID) can be obtained from Get Radius profile list(Get Radius profile template list) interface. | 
**RekeyEntInterval** | Pointer to **int32** | WPA-Enterprise SSID group key update period interval config. When the value of Parameter [intervalEntType] is 0(Seconds), it should be within the range of 30-86400; when the value of Parameter [intervalEntType] is 1(Minutes), it should be within the range of 1-1440; when the value of Parameter [intervalEntType] is 2(Hours), it should be within the range of 1-24. | [optional] 
**VersionEnt** | **int32** | WPA-Enterprise version. This is necessary when the value of security is 2(WPA-Enterprise); It should be a value as follows: 1: WPA-Enterprise; 2: WPA2-Enterprise; 3: WPA/WPA2-Enterprise; 4.WPA3-Enterprise. | 

## Methods

### NewSsidEnterpriseSettingOpenApiVO

`func NewSsidEnterpriseSettingOpenApiVO(encryptionEnt int32, gikRekeyEntEnable bool, radiusProfileId string, versionEnt int32, ) *SsidEnterpriseSettingOpenApiVO`

NewSsidEnterpriseSettingOpenApiVO instantiates a new SsidEnterpriseSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsidEnterpriseSettingOpenApiVOWithDefaults

`func NewSsidEnterpriseSettingOpenApiVOWithDefaults() *SsidEnterpriseSettingOpenApiVO`

NewSsidEnterpriseSettingOpenApiVOWithDefaults instantiates a new SsidEnterpriseSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptionEnt

`func (o *SsidEnterpriseSettingOpenApiVO) GetEncryptionEnt() int32`

GetEncryptionEnt returns the EncryptionEnt field if non-nil, zero value otherwise.

### GetEncryptionEntOk

`func (o *SsidEnterpriseSettingOpenApiVO) GetEncryptionEntOk() (*int32, bool)`

GetEncryptionEntOk returns a tuple with the EncryptionEnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnt

`func (o *SsidEnterpriseSettingOpenApiVO) SetEncryptionEnt(v int32)`

SetEncryptionEnt sets EncryptionEnt field to given value.


### GetGikRekeyEntEnable

`func (o *SsidEnterpriseSettingOpenApiVO) GetGikRekeyEntEnable() bool`

GetGikRekeyEntEnable returns the GikRekeyEntEnable field if non-nil, zero value otherwise.

### GetGikRekeyEntEnableOk

`func (o *SsidEnterpriseSettingOpenApiVO) GetGikRekeyEntEnableOk() (*bool, bool)`

GetGikRekeyEntEnableOk returns a tuple with the GikRekeyEntEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGikRekeyEntEnable

`func (o *SsidEnterpriseSettingOpenApiVO) SetGikRekeyEntEnable(v bool)`

SetGikRekeyEntEnable sets GikRekeyEntEnable field to given value.


### GetIntervalEntType

`func (o *SsidEnterpriseSettingOpenApiVO) GetIntervalEntType() int32`

GetIntervalEntType returns the IntervalEntType field if non-nil, zero value otherwise.

### GetIntervalEntTypeOk

`func (o *SsidEnterpriseSettingOpenApiVO) GetIntervalEntTypeOk() (*int32, bool)`

GetIntervalEntTypeOk returns a tuple with the IntervalEntType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalEntType

`func (o *SsidEnterpriseSettingOpenApiVO) SetIntervalEntType(v int32)`

SetIntervalEntType sets IntervalEntType field to given value.

### HasIntervalEntType

`func (o *SsidEnterpriseSettingOpenApiVO) HasIntervalEntType() bool`

HasIntervalEntType returns a boolean if a field has been set.

### GetNasId

`func (o *SsidEnterpriseSettingOpenApiVO) GetNasId() string`

GetNasId returns the NasId field if non-nil, zero value otherwise.

### GetNasIdOk

`func (o *SsidEnterpriseSettingOpenApiVO) GetNasIdOk() (*string, bool)`

GetNasIdOk returns a tuple with the NasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasId

`func (o *SsidEnterpriseSettingOpenApiVO) SetNasId(v string)`

SetNasId sets NasId field to given value.

### HasNasId

`func (o *SsidEnterpriseSettingOpenApiVO) HasNasId() bool`

HasNasId returns a boolean if a field has been set.

### GetNasIdMode

`func (o *SsidEnterpriseSettingOpenApiVO) GetNasIdMode() int32`

GetNasIdMode returns the NasIdMode field if non-nil, zero value otherwise.

### GetNasIdModeOk

`func (o *SsidEnterpriseSettingOpenApiVO) GetNasIdModeOk() (*int32, bool)`

GetNasIdModeOk returns a tuple with the NasIdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasIdMode

`func (o *SsidEnterpriseSettingOpenApiVO) SetNasIdMode(v int32)`

SetNasIdMode sets NasIdMode field to given value.

### HasNasIdMode

`func (o *SsidEnterpriseSettingOpenApiVO) HasNasIdMode() bool`

HasNasIdMode returns a boolean if a field has been set.

### GetRadiusProfileId

`func (o *SsidEnterpriseSettingOpenApiVO) GetRadiusProfileId() string`

GetRadiusProfileId returns the RadiusProfileId field if non-nil, zero value otherwise.

### GetRadiusProfileIdOk

`func (o *SsidEnterpriseSettingOpenApiVO) GetRadiusProfileIdOk() (*string, bool)`

GetRadiusProfileIdOk returns a tuple with the RadiusProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusProfileId

`func (o *SsidEnterpriseSettingOpenApiVO) SetRadiusProfileId(v string)`

SetRadiusProfileId sets RadiusProfileId field to given value.


### GetRekeyEntInterval

`func (o *SsidEnterpriseSettingOpenApiVO) GetRekeyEntInterval() int32`

GetRekeyEntInterval returns the RekeyEntInterval field if non-nil, zero value otherwise.

### GetRekeyEntIntervalOk

`func (o *SsidEnterpriseSettingOpenApiVO) GetRekeyEntIntervalOk() (*int32, bool)`

GetRekeyEntIntervalOk returns a tuple with the RekeyEntInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRekeyEntInterval

`func (o *SsidEnterpriseSettingOpenApiVO) SetRekeyEntInterval(v int32)`

SetRekeyEntInterval sets RekeyEntInterval field to given value.

### HasRekeyEntInterval

`func (o *SsidEnterpriseSettingOpenApiVO) HasRekeyEntInterval() bool`

HasRekeyEntInterval returns a boolean if a field has been set.

### GetVersionEnt

`func (o *SsidEnterpriseSettingOpenApiVO) GetVersionEnt() int32`

GetVersionEnt returns the VersionEnt field if non-nil, zero value otherwise.

### GetVersionEntOk

`func (o *SsidEnterpriseSettingOpenApiVO) GetVersionEntOk() (*int32, bool)`

GetVersionEntOk returns a tuple with the VersionEnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionEnt

`func (o *SsidEnterpriseSettingOpenApiVO) SetVersionEnt(v int32)`

SetVersionEnt sets VersionEnt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


