# SiteSummaryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address of the site | [optional] 
**Latitude** | Pointer to **float64** | Latitude of the site should be within the range of -90 - 90. | [optional] 
**Longitude** | Pointer to **float64** | Longitude of the site should be within the range of -180 - 180. | [optional] 
**Name** | Pointer to **string** | Name of the site should contain 1 to 64 characters. | [optional] 
**Primary** | Pointer to **bool** | Default Site mark | [optional] 
**Region** | Pointer to **string** | Country/Region of the site; For the values of region, refer to the abbreviation of the ISO country code; For example, you need to input \&quot;United States\&quot; for the United States of America. | [optional] 
**Scenario** | Pointer to **string** | For the values of the scenario of the site, refer to result of the interface for Get scenario list. | [optional] 
**SiteId** | Pointer to **string** | Site ID | [optional] 
**SitePublicIp** | Pointer to **string** | Adopted gateway public ip of the site, only useful for cloud based controller and remote management local Controller | [optional] 
**SupportES** | Pointer to **bool** | Whether the site supports adopting Agile Series Switches | [optional] 
**SupportL2** | Pointer to **bool** | Whether the site supports adopting Non-Agile Series Switches | [optional] 
**TagIds** | Pointer to **[]string** | Site tag ID | [optional] 
**TimeZone** | Pointer to **string** | For the values of the timezone of the site, refer to section 5.1 of the Open API Access Guide. | [optional] 
**Type** | Pointer to **int32** | Site type(only for pro controller). It should be a value as follows: 0: Basic Site; 1: Pro Site | [optional] 

## Methods

### NewSiteSummaryInfo

`func NewSiteSummaryInfo() *SiteSummaryInfo`

NewSiteSummaryInfo instantiates a new SiteSummaryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteSummaryInfoWithDefaults

`func NewSiteSummaryInfoWithDefaults() *SiteSummaryInfo`

NewSiteSummaryInfoWithDefaults instantiates a new SiteSummaryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SiteSummaryInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SiteSummaryInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SiteSummaryInfo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SiteSummaryInfo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLatitude

`func (o *SiteSummaryInfo) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *SiteSummaryInfo) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *SiteSummaryInfo) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *SiteSummaryInfo) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *SiteSummaryInfo) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *SiteSummaryInfo) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *SiteSummaryInfo) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *SiteSummaryInfo) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetName

`func (o *SiteSummaryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteSummaryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteSummaryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SiteSummaryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimary

`func (o *SiteSummaryInfo) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *SiteSummaryInfo) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *SiteSummaryInfo) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *SiteSummaryInfo) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetRegion

`func (o *SiteSummaryInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SiteSummaryInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SiteSummaryInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *SiteSummaryInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetScenario

`func (o *SiteSummaryInfo) GetScenario() string`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *SiteSummaryInfo) GetScenarioOk() (*string, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *SiteSummaryInfo) SetScenario(v string)`

SetScenario sets Scenario field to given value.

### HasScenario

`func (o *SiteSummaryInfo) HasScenario() bool`

HasScenario returns a boolean if a field has been set.

### GetSiteId

`func (o *SiteSummaryInfo) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SiteSummaryInfo) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SiteSummaryInfo) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SiteSummaryInfo) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSitePublicIp

`func (o *SiteSummaryInfo) GetSitePublicIp() string`

GetSitePublicIp returns the SitePublicIp field if non-nil, zero value otherwise.

### GetSitePublicIpOk

`func (o *SiteSummaryInfo) GetSitePublicIpOk() (*string, bool)`

GetSitePublicIpOk returns a tuple with the SitePublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitePublicIp

`func (o *SiteSummaryInfo) SetSitePublicIp(v string)`

SetSitePublicIp sets SitePublicIp field to given value.

### HasSitePublicIp

`func (o *SiteSummaryInfo) HasSitePublicIp() bool`

HasSitePublicIp returns a boolean if a field has been set.

### GetSupportES

`func (o *SiteSummaryInfo) GetSupportES() bool`

GetSupportES returns the SupportES field if non-nil, zero value otherwise.

### GetSupportESOk

`func (o *SiteSummaryInfo) GetSupportESOk() (*bool, bool)`

GetSupportESOk returns a tuple with the SupportES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportES

`func (o *SiteSummaryInfo) SetSupportES(v bool)`

SetSupportES sets SupportES field to given value.

### HasSupportES

`func (o *SiteSummaryInfo) HasSupportES() bool`

HasSupportES returns a boolean if a field has been set.

### GetSupportL2

`func (o *SiteSummaryInfo) GetSupportL2() bool`

GetSupportL2 returns the SupportL2 field if non-nil, zero value otherwise.

### GetSupportL2Ok

`func (o *SiteSummaryInfo) GetSupportL2Ok() (*bool, bool)`

GetSupportL2Ok returns a tuple with the SupportL2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportL2

`func (o *SiteSummaryInfo) SetSupportL2(v bool)`

SetSupportL2 sets SupportL2 field to given value.

### HasSupportL2

`func (o *SiteSummaryInfo) HasSupportL2() bool`

HasSupportL2 returns a boolean if a field has been set.

### GetTagIds

`func (o *SiteSummaryInfo) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *SiteSummaryInfo) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *SiteSummaryInfo) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *SiteSummaryInfo) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTimeZone

`func (o *SiteSummaryInfo) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *SiteSummaryInfo) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *SiteSummaryInfo) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *SiteSummaryInfo) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetType

`func (o *SiteSummaryInfo) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SiteSummaryInfo) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SiteSummaryInfo) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SiteSummaryInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


