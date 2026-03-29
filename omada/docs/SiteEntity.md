# SiteEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address of the site | [optional] 
**Dst** | Pointer to [**DstDTO**](DstDTO.md) |  | [optional] 
**Latitude** | Pointer to **float64** | Latitude of the site should be within the range of -90 - 90. | [optional] 
**Longitude** | Pointer to **float64** | Longitude of the site should be within the range of -180 - 180. | [optional] 
**Name** | Pointer to **string** | Name of the site should contain 1 to 64 characters. | [optional] 
**NtpEnable** | Pointer to **bool** | NTP server status of the site | [optional] 
**NtpServers** | Pointer to **[]string** | NTP server address; Up to 5 entries are allowed for the NTP server address list. | [optional] 
**Region** | Pointer to **string** | Country/Region of the site; For the values of region, refer to the abbreviation of the ISO country code; For example, you need to input \&quot;United States\&quot; for the United States of America. | [optional] 
**Scenario** | Pointer to **string** | For the values of the scenario of the site, refer to result of the interface for Get scenario list. | [optional] 
**SiteId** | Pointer to **string** | Site ID | [optional] 
**SupportES** | Pointer to **bool** | Whether the site supports adopting Agile Series Switches | [optional] 
**SupportL2** | Pointer to **bool** | Whether the site supports adopting Non-Agile Series Switches | [optional] 
**TagIds** | Pointer to **[]string** | Site tag ID | [optional] 
**TimeZone** | Pointer to **string** | For the values of the timezone of the site, refer to section 5.1 of the Open API Access Guide. | [optional] 
**Type** | Pointer to **int32** | Type of the site should be 0 or 1, and 0 means basic site, 1 means pro site. | [optional] 

## Methods

### NewSiteEntity

`func NewSiteEntity() *SiteEntity`

NewSiteEntity instantiates a new SiteEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteEntityWithDefaults

`func NewSiteEntityWithDefaults() *SiteEntity`

NewSiteEntityWithDefaults instantiates a new SiteEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SiteEntity) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SiteEntity) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SiteEntity) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SiteEntity) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDst

`func (o *SiteEntity) GetDst() DstDTO`

GetDst returns the Dst field if non-nil, zero value otherwise.

### GetDstOk

`func (o *SiteEntity) GetDstOk() (*DstDTO, bool)`

GetDstOk returns a tuple with the Dst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDst

`func (o *SiteEntity) SetDst(v DstDTO)`

SetDst sets Dst field to given value.

### HasDst

`func (o *SiteEntity) HasDst() bool`

HasDst returns a boolean if a field has been set.

### GetLatitude

`func (o *SiteEntity) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *SiteEntity) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *SiteEntity) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *SiteEntity) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *SiteEntity) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *SiteEntity) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *SiteEntity) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *SiteEntity) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetName

`func (o *SiteEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteEntity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SiteEntity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNtpEnable

`func (o *SiteEntity) GetNtpEnable() bool`

GetNtpEnable returns the NtpEnable field if non-nil, zero value otherwise.

### GetNtpEnableOk

`func (o *SiteEntity) GetNtpEnableOk() (*bool, bool)`

GetNtpEnableOk returns a tuple with the NtpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpEnable

`func (o *SiteEntity) SetNtpEnable(v bool)`

SetNtpEnable sets NtpEnable field to given value.

### HasNtpEnable

`func (o *SiteEntity) HasNtpEnable() bool`

HasNtpEnable returns a boolean if a field has been set.

### GetNtpServers

`func (o *SiteEntity) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *SiteEntity) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *SiteEntity) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *SiteEntity) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetRegion

`func (o *SiteEntity) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SiteEntity) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SiteEntity) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *SiteEntity) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetScenario

`func (o *SiteEntity) GetScenario() string`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *SiteEntity) GetScenarioOk() (*string, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *SiteEntity) SetScenario(v string)`

SetScenario sets Scenario field to given value.

### HasScenario

`func (o *SiteEntity) HasScenario() bool`

HasScenario returns a boolean if a field has been set.

### GetSiteId

`func (o *SiteEntity) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SiteEntity) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SiteEntity) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SiteEntity) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSupportES

`func (o *SiteEntity) GetSupportES() bool`

GetSupportES returns the SupportES field if non-nil, zero value otherwise.

### GetSupportESOk

`func (o *SiteEntity) GetSupportESOk() (*bool, bool)`

GetSupportESOk returns a tuple with the SupportES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportES

`func (o *SiteEntity) SetSupportES(v bool)`

SetSupportES sets SupportES field to given value.

### HasSupportES

`func (o *SiteEntity) HasSupportES() bool`

HasSupportES returns a boolean if a field has been set.

### GetSupportL2

`func (o *SiteEntity) GetSupportL2() bool`

GetSupportL2 returns the SupportL2 field if non-nil, zero value otherwise.

### GetSupportL2Ok

`func (o *SiteEntity) GetSupportL2Ok() (*bool, bool)`

GetSupportL2Ok returns a tuple with the SupportL2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportL2

`func (o *SiteEntity) SetSupportL2(v bool)`

SetSupportL2 sets SupportL2 field to given value.

### HasSupportL2

`func (o *SiteEntity) HasSupportL2() bool`

HasSupportL2 returns a boolean if a field has been set.

### GetTagIds

`func (o *SiteEntity) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *SiteEntity) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *SiteEntity) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *SiteEntity) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTimeZone

`func (o *SiteEntity) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *SiteEntity) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *SiteEntity) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *SiteEntity) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetType

`func (o *SiteEntity) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SiteEntity) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SiteEntity) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SiteEntity) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


