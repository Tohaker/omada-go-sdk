# CreateSiteEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address of the site | [optional] 
**DeviceAccountSetting** | [**DeviceAccountSettingOpenApiVO**](DeviceAccountSettingOpenApiVO.md) |  | 
**Latitude** | Pointer to **float64** | Latitude of the site should be within the range of -90 - 90. | [optional] 
**Longitude** | Pointer to **float64** | Longitude of the site should be within the range of -180 - 180. | [optional] 
**Name** | **string** | Name of the site should contain 1 to 64 characters. | 
**Region** | **string** | Country/Region of the site; For the values of region, refer to the abbreviation of the ISO country code; For example, you need to input \&quot;United States\&quot; for the United States of America. | 
**Scenario** | **string** | For the values of the scenario of the site, refer to result of the interface for Get scenario list. | 
**SupportES** | Pointer to **bool** | Whether the site supports adopting Agile Series Switches | [optional] 
**SupportL2** | Pointer to **bool** | Whether the site supports adopting Non-Agile Series Switches | [optional] 
**TagIds** | Pointer to **[]string** | Site tag ID, Site tag ID can be created using \&quot;Create new site tag\&quot; interface, and site tag ID can be obtained from \&quot;Get site tag list\&quot; interface | [optional] 
**TimeZone** | **string** | For the values of the timezone of the site, refer to section 5.1 of the Open API Access Guide. | 
**Type** | Pointer to **int32** | Type of the site should be 0 or 1, and 0 means basic site, 1 means pro site. | [optional] 

## Methods

### NewCreateSiteEntity

`func NewCreateSiteEntity(deviceAccountSetting DeviceAccountSettingOpenApiVO, name string, region string, scenario string, timeZone string, ) *CreateSiteEntity`

NewCreateSiteEntity instantiates a new CreateSiteEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSiteEntityWithDefaults

`func NewCreateSiteEntityWithDefaults() *CreateSiteEntity`

NewCreateSiteEntityWithDefaults instantiates a new CreateSiteEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CreateSiteEntity) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateSiteEntity) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateSiteEntity) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CreateSiteEntity) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDeviceAccountSetting

`func (o *CreateSiteEntity) GetDeviceAccountSetting() DeviceAccountSettingOpenApiVO`

GetDeviceAccountSetting returns the DeviceAccountSetting field if non-nil, zero value otherwise.

### GetDeviceAccountSettingOk

`func (o *CreateSiteEntity) GetDeviceAccountSettingOk() (*DeviceAccountSettingOpenApiVO, bool)`

GetDeviceAccountSettingOk returns a tuple with the DeviceAccountSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAccountSetting

`func (o *CreateSiteEntity) SetDeviceAccountSetting(v DeviceAccountSettingOpenApiVO)`

SetDeviceAccountSetting sets DeviceAccountSetting field to given value.


### GetLatitude

`func (o *CreateSiteEntity) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *CreateSiteEntity) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *CreateSiteEntity) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *CreateSiteEntity) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *CreateSiteEntity) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *CreateSiteEntity) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *CreateSiteEntity) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *CreateSiteEntity) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetName

`func (o *CreateSiteEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSiteEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSiteEntity) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *CreateSiteEntity) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateSiteEntity) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateSiteEntity) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetScenario

`func (o *CreateSiteEntity) GetScenario() string`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *CreateSiteEntity) GetScenarioOk() (*string, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *CreateSiteEntity) SetScenario(v string)`

SetScenario sets Scenario field to given value.


### GetSupportES

`func (o *CreateSiteEntity) GetSupportES() bool`

GetSupportES returns the SupportES field if non-nil, zero value otherwise.

### GetSupportESOk

`func (o *CreateSiteEntity) GetSupportESOk() (*bool, bool)`

GetSupportESOk returns a tuple with the SupportES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportES

`func (o *CreateSiteEntity) SetSupportES(v bool)`

SetSupportES sets SupportES field to given value.

### HasSupportES

`func (o *CreateSiteEntity) HasSupportES() bool`

HasSupportES returns a boolean if a field has been set.

### GetSupportL2

`func (o *CreateSiteEntity) GetSupportL2() bool`

GetSupportL2 returns the SupportL2 field if non-nil, zero value otherwise.

### GetSupportL2Ok

`func (o *CreateSiteEntity) GetSupportL2Ok() (*bool, bool)`

GetSupportL2Ok returns a tuple with the SupportL2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportL2

`func (o *CreateSiteEntity) SetSupportL2(v bool)`

SetSupportL2 sets SupportL2 field to given value.

### HasSupportL2

`func (o *CreateSiteEntity) HasSupportL2() bool`

HasSupportL2 returns a boolean if a field has been set.

### GetTagIds

`func (o *CreateSiteEntity) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *CreateSiteEntity) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *CreateSiteEntity) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *CreateSiteEntity) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTimeZone

`func (o *CreateSiteEntity) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *CreateSiteEntity) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *CreateSiteEntity) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetType

`func (o *CreateSiteEntity) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateSiteEntity) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateSiteEntity) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *CreateSiteEntity) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


