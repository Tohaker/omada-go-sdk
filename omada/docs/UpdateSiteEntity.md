# UpdateSiteEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address of the site | [optional] 
**Dst** | Pointer to [**ModifyDstDTO**](ModifyDstDTO.md) |  | [optional] 
**Latitude** | Pointer to **float64** | Latitude of the site should be within the range of -90 - 90. | [optional] 
**Longitude** | Pointer to **float64** | Longitude of the site should be within the range of -180 - 180. | [optional] 
**Name** | Pointer to **string** | Site name should contain 1 to 64 characters. | [optional] 
**NtpEnable** | Pointer to **bool** | NTP server status of the site | [optional] 
**NtpServers** | Pointer to [**[]NtpServer**](NtpServer.md) | NTP server address; Up to 5 entries are allowed for the NTP server address list. | [optional] 
**Region** | **string** | Country/Region of the site; For the values of region, refer to the abbreviation of the ISO country code; For example, you need to input \&quot;United States\&quot; for the United States of America. | 
**Scenario** | **string** | For the values of the scenario of the site, refer to result of the interface for Get scenario list. | 
**SupportES** | Pointer to **bool** | Whether the site supports adopting Agile Series Switches | [optional] 
**SupportL2** | Pointer to **bool** | Whether the site supports adopting Non-Agile Series Switches | [optional] 
**TagIds** | Pointer to **[]string** | Site tag ID, Site tag ID can be created using \&quot;Create new site tag\&quot; interface, and site tag ID can be obtained from \&quot;Get site tag list\&quot; interface | [optional] 
**TimeZone** | **string** | For the values of the timezone of the site, refer to section 5.1 of the Open API Access Guide. | 

## Methods

### NewUpdateSiteEntity

`func NewUpdateSiteEntity(region string, scenario string, timeZone string, ) *UpdateSiteEntity`

NewUpdateSiteEntity instantiates a new UpdateSiteEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSiteEntityWithDefaults

`func NewUpdateSiteEntityWithDefaults() *UpdateSiteEntity`

NewUpdateSiteEntityWithDefaults instantiates a new UpdateSiteEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UpdateSiteEntity) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateSiteEntity) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateSiteEntity) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdateSiteEntity) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDst

`func (o *UpdateSiteEntity) GetDst() ModifyDstDTO`

GetDst returns the Dst field if non-nil, zero value otherwise.

### GetDstOk

`func (o *UpdateSiteEntity) GetDstOk() (*ModifyDstDTO, bool)`

GetDstOk returns a tuple with the Dst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDst

`func (o *UpdateSiteEntity) SetDst(v ModifyDstDTO)`

SetDst sets Dst field to given value.

### HasDst

`func (o *UpdateSiteEntity) HasDst() bool`

HasDst returns a boolean if a field has been set.

### GetLatitude

`func (o *UpdateSiteEntity) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *UpdateSiteEntity) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *UpdateSiteEntity) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *UpdateSiteEntity) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *UpdateSiteEntity) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *UpdateSiteEntity) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *UpdateSiteEntity) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *UpdateSiteEntity) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetName

`func (o *UpdateSiteEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSiteEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSiteEntity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSiteEntity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNtpEnable

`func (o *UpdateSiteEntity) GetNtpEnable() bool`

GetNtpEnable returns the NtpEnable field if non-nil, zero value otherwise.

### GetNtpEnableOk

`func (o *UpdateSiteEntity) GetNtpEnableOk() (*bool, bool)`

GetNtpEnableOk returns a tuple with the NtpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpEnable

`func (o *UpdateSiteEntity) SetNtpEnable(v bool)`

SetNtpEnable sets NtpEnable field to given value.

### HasNtpEnable

`func (o *UpdateSiteEntity) HasNtpEnable() bool`

HasNtpEnable returns a boolean if a field has been set.

### GetNtpServers

`func (o *UpdateSiteEntity) GetNtpServers() []NtpServer`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *UpdateSiteEntity) GetNtpServersOk() (*[]NtpServer, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *UpdateSiteEntity) SetNtpServers(v []NtpServer)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *UpdateSiteEntity) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetRegion

`func (o *UpdateSiteEntity) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *UpdateSiteEntity) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *UpdateSiteEntity) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetScenario

`func (o *UpdateSiteEntity) GetScenario() string`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *UpdateSiteEntity) GetScenarioOk() (*string, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *UpdateSiteEntity) SetScenario(v string)`

SetScenario sets Scenario field to given value.


### GetSupportES

`func (o *UpdateSiteEntity) GetSupportES() bool`

GetSupportES returns the SupportES field if non-nil, zero value otherwise.

### GetSupportESOk

`func (o *UpdateSiteEntity) GetSupportESOk() (*bool, bool)`

GetSupportESOk returns a tuple with the SupportES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportES

`func (o *UpdateSiteEntity) SetSupportES(v bool)`

SetSupportES sets SupportES field to given value.

### HasSupportES

`func (o *UpdateSiteEntity) HasSupportES() bool`

HasSupportES returns a boolean if a field has been set.

### GetSupportL2

`func (o *UpdateSiteEntity) GetSupportL2() bool`

GetSupportL2 returns the SupportL2 field if non-nil, zero value otherwise.

### GetSupportL2Ok

`func (o *UpdateSiteEntity) GetSupportL2Ok() (*bool, bool)`

GetSupportL2Ok returns a tuple with the SupportL2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportL2

`func (o *UpdateSiteEntity) SetSupportL2(v bool)`

SetSupportL2 sets SupportL2 field to given value.

### HasSupportL2

`func (o *UpdateSiteEntity) HasSupportL2() bool`

HasSupportL2 returns a boolean if a field has been set.

### GetTagIds

`func (o *UpdateSiteEntity) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *UpdateSiteEntity) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *UpdateSiteEntity) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *UpdateSiteEntity) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTimeZone

`func (o *UpdateSiteEntity) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *UpdateSiteEntity) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *UpdateSiteEntity) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


