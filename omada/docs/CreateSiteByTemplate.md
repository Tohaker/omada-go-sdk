# CreateSiteByTemplate

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
**SiteTemplateId** | **string** | Site template ID. | 
**SupportES** | Pointer to **bool** | Whether the site supports adopting Agile Series switches | [optional] 
**SupportL2** | Pointer to **bool** | Whether the site supports adopting Non-Agile Series switches | [optional] 
**TagIds** | Pointer to **[]string** | Site tag ID, Site tag ID can be created using \&quot;Create new site tag\&quot; interface, and site tag ID can be obtained from \&quot;Get site tag list\&quot; interface | [optional] 

## Methods

### NewCreateSiteByTemplate

`func NewCreateSiteByTemplate(deviceAccountSetting DeviceAccountSettingOpenApiVO, name string, region string, scenario string, siteTemplateId string, ) *CreateSiteByTemplate`

NewCreateSiteByTemplate instantiates a new CreateSiteByTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSiteByTemplateWithDefaults

`func NewCreateSiteByTemplateWithDefaults() *CreateSiteByTemplate`

NewCreateSiteByTemplateWithDefaults instantiates a new CreateSiteByTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CreateSiteByTemplate) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateSiteByTemplate) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateSiteByTemplate) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CreateSiteByTemplate) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDeviceAccountSetting

`func (o *CreateSiteByTemplate) GetDeviceAccountSetting() DeviceAccountSettingOpenApiVO`

GetDeviceAccountSetting returns the DeviceAccountSetting field if non-nil, zero value otherwise.

### GetDeviceAccountSettingOk

`func (o *CreateSiteByTemplate) GetDeviceAccountSettingOk() (*DeviceAccountSettingOpenApiVO, bool)`

GetDeviceAccountSettingOk returns a tuple with the DeviceAccountSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAccountSetting

`func (o *CreateSiteByTemplate) SetDeviceAccountSetting(v DeviceAccountSettingOpenApiVO)`

SetDeviceAccountSetting sets DeviceAccountSetting field to given value.


### GetLatitude

`func (o *CreateSiteByTemplate) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *CreateSiteByTemplate) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *CreateSiteByTemplate) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *CreateSiteByTemplate) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *CreateSiteByTemplate) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *CreateSiteByTemplate) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *CreateSiteByTemplate) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *CreateSiteByTemplate) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetName

`func (o *CreateSiteByTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSiteByTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSiteByTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *CreateSiteByTemplate) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateSiteByTemplate) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateSiteByTemplate) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetScenario

`func (o *CreateSiteByTemplate) GetScenario() string`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *CreateSiteByTemplate) GetScenarioOk() (*string, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *CreateSiteByTemplate) SetScenario(v string)`

SetScenario sets Scenario field to given value.


### GetSiteTemplateId

`func (o *CreateSiteByTemplate) GetSiteTemplateId() string`

GetSiteTemplateId returns the SiteTemplateId field if non-nil, zero value otherwise.

### GetSiteTemplateIdOk

`func (o *CreateSiteByTemplate) GetSiteTemplateIdOk() (*string, bool)`

GetSiteTemplateIdOk returns a tuple with the SiteTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTemplateId

`func (o *CreateSiteByTemplate) SetSiteTemplateId(v string)`

SetSiteTemplateId sets SiteTemplateId field to given value.


### GetSupportES

`func (o *CreateSiteByTemplate) GetSupportES() bool`

GetSupportES returns the SupportES field if non-nil, zero value otherwise.

### GetSupportESOk

`func (o *CreateSiteByTemplate) GetSupportESOk() (*bool, bool)`

GetSupportESOk returns a tuple with the SupportES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportES

`func (o *CreateSiteByTemplate) SetSupportES(v bool)`

SetSupportES sets SupportES field to given value.

### HasSupportES

`func (o *CreateSiteByTemplate) HasSupportES() bool`

HasSupportES returns a boolean if a field has been set.

### GetSupportL2

`func (o *CreateSiteByTemplate) GetSupportL2() bool`

GetSupportL2 returns the SupportL2 field if non-nil, zero value otherwise.

### GetSupportL2Ok

`func (o *CreateSiteByTemplate) GetSupportL2Ok() (*bool, bool)`

GetSupportL2Ok returns a tuple with the SupportL2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportL2

`func (o *CreateSiteByTemplate) SetSupportL2(v bool)`

SetSupportL2 sets SupportL2 field to given value.

### HasSupportL2

`func (o *CreateSiteByTemplate) HasSupportL2() bool`

HasSupportL2 returns a boolean if a field has been set.

### GetTagIds

`func (o *CreateSiteByTemplate) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *CreateSiteByTemplate) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *CreateSiteByTemplate) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *CreateSiteByTemplate) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


