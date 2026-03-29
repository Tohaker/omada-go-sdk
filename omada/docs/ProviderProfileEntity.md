# ProviderProfileEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OmadacId** | Pointer to **string** | Omadac ID | [optional] 
**ProfileId** | Pointer to **string** | Provider profile ID | [optional] 
**ProfileName** | Pointer to **string** | Provider profile name | [optional] 
**ProviderSettings** | Pointer to [**ProviderSettingVO**](ProviderSettingVO.md) |  | [optional] 
**SiteId** | Pointer to **string** | Site ID | [optional] 

## Methods

### NewProviderProfileEntity

`func NewProviderProfileEntity() *ProviderProfileEntity`

NewProviderProfileEntity instantiates a new ProviderProfileEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderProfileEntityWithDefaults

`func NewProviderProfileEntityWithDefaults() *ProviderProfileEntity`

NewProviderProfileEntityWithDefaults instantiates a new ProviderProfileEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOmadacId

`func (o *ProviderProfileEntity) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *ProviderProfileEntity) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *ProviderProfileEntity) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *ProviderProfileEntity) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetProfileId

`func (o *ProviderProfileEntity) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *ProviderProfileEntity) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *ProviderProfileEntity) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *ProviderProfileEntity) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileName

`func (o *ProviderProfileEntity) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *ProviderProfileEntity) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *ProviderProfileEntity) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *ProviderProfileEntity) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetProviderSettings

`func (o *ProviderProfileEntity) GetProviderSettings() ProviderSettingVO`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *ProviderProfileEntity) GetProviderSettingsOk() (*ProviderSettingVO, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *ProviderProfileEntity) SetProviderSettings(v ProviderSettingVO)`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *ProviderProfileEntity) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.

### GetSiteId

`func (o *ProviderProfileEntity) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ProviderProfileEntity) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ProviderProfileEntity) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ProviderProfileEntity) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


