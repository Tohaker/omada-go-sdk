# SiteTemplateSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dst** | Pointer to [**SiteSettingDst**](SiteSettingDst.md) |  | [optional] 
**Id** | Pointer to **string** | site template id | [optional] 
**Name** | Pointer to **string** | site template name | [optional] 
**NtpEnable** | Pointer to **bool** | Network Time Protocol enable | [optional] 
**NtpServers** | Pointer to [**[]NtpServer**](NtpServer.md) | ntp servers | [optional] 
**OmadacId** | Pointer to **string** | omadacId | [optional] 
**Settings** | Pointer to **[]string** | list of supported configurations | [optional] 
**TimeZone** | Pointer to **string** | time zone | [optional] 
**Type** | Pointer to **int32** | site type. 0：Basic；1：Pro | [optional] 

## Methods

### NewSiteTemplateSettings

`func NewSiteTemplateSettings() *SiteTemplateSettings`

NewSiteTemplateSettings instantiates a new SiteTemplateSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteTemplateSettingsWithDefaults

`func NewSiteTemplateSettingsWithDefaults() *SiteTemplateSettings`

NewSiteTemplateSettingsWithDefaults instantiates a new SiteTemplateSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDst

`func (o *SiteTemplateSettings) GetDst() SiteSettingDst`

GetDst returns the Dst field if non-nil, zero value otherwise.

### GetDstOk

`func (o *SiteTemplateSettings) GetDstOk() (*SiteSettingDst, bool)`

GetDstOk returns a tuple with the Dst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDst

`func (o *SiteTemplateSettings) SetDst(v SiteSettingDst)`

SetDst sets Dst field to given value.

### HasDst

`func (o *SiteTemplateSettings) HasDst() bool`

HasDst returns a boolean if a field has been set.

### GetId

`func (o *SiteTemplateSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteTemplateSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteTemplateSettings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SiteTemplateSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SiteTemplateSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteTemplateSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteTemplateSettings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SiteTemplateSettings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNtpEnable

`func (o *SiteTemplateSettings) GetNtpEnable() bool`

GetNtpEnable returns the NtpEnable field if non-nil, zero value otherwise.

### GetNtpEnableOk

`func (o *SiteTemplateSettings) GetNtpEnableOk() (*bool, bool)`

GetNtpEnableOk returns a tuple with the NtpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpEnable

`func (o *SiteTemplateSettings) SetNtpEnable(v bool)`

SetNtpEnable sets NtpEnable field to given value.

### HasNtpEnable

`func (o *SiteTemplateSettings) HasNtpEnable() bool`

HasNtpEnable returns a boolean if a field has been set.

### GetNtpServers

`func (o *SiteTemplateSettings) GetNtpServers() []NtpServer`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *SiteTemplateSettings) GetNtpServersOk() (*[]NtpServer, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *SiteTemplateSettings) SetNtpServers(v []NtpServer)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *SiteTemplateSettings) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetOmadacId

`func (o *SiteTemplateSettings) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *SiteTemplateSettings) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *SiteTemplateSettings) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *SiteTemplateSettings) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetSettings

`func (o *SiteTemplateSettings) GetSettings() []string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SiteTemplateSettings) GetSettingsOk() (*[]string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SiteTemplateSettings) SetSettings(v []string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *SiteTemplateSettings) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetTimeZone

`func (o *SiteTemplateSettings) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *SiteTemplateSettings) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *SiteTemplateSettings) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *SiteTemplateSettings) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetType

`func (o *SiteTemplateSettings) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SiteTemplateSettings) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SiteTemplateSettings) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SiteTemplateSettings) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


