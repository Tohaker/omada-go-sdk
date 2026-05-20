# SdWanMemberConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMac** | **string** | The device MAC of the sdWan member. | 
**LanNetworkInfo** | Pointer to [**[]LanNetworkBrief**](LanNetworkBrief.md) | A list of lan network info for SdWan Member. | [optional] 
**PublicIp** | Pointer to **bool** | Whether the sdWan member has a public IP. | [optional] 
**Role** | **int32** | The role of sdWan member, hub or spoke. | 
**SiteId** | **string** | The ID of the site where the sdWan member is located. | 
**SiteName** | Pointer to **string** | The name of the site where the sdWan member is located. | [optional] 
**WanPortsInfo** | Pointer to [**[]OsgPortStatBrief**](OsgPortStatBrief.md) | A list of device port status info for SdWan Member. | [optional] 

## Methods

### NewSdWanMemberConfig

`func NewSdWanMemberConfig(deviceMac string, role int32, siteId string, ) *SdWanMemberConfig`

NewSdWanMemberConfig instantiates a new SdWanMemberConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdWanMemberConfigWithDefaults

`func NewSdWanMemberConfigWithDefaults() *SdWanMemberConfig`

NewSdWanMemberConfigWithDefaults instantiates a new SdWanMemberConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMac

`func (o *SdWanMemberConfig) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *SdWanMemberConfig) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *SdWanMemberConfig) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.


### GetLanNetworkInfo

`func (o *SdWanMemberConfig) GetLanNetworkInfo() []LanNetworkBrief`

GetLanNetworkInfo returns the LanNetworkInfo field if non-nil, zero value otherwise.

### GetLanNetworkInfoOk

`func (o *SdWanMemberConfig) GetLanNetworkInfoOk() (*[]LanNetworkBrief, bool)`

GetLanNetworkInfoOk returns a tuple with the LanNetworkInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetworkInfo

`func (o *SdWanMemberConfig) SetLanNetworkInfo(v []LanNetworkBrief)`

SetLanNetworkInfo sets LanNetworkInfo field to given value.

### HasLanNetworkInfo

`func (o *SdWanMemberConfig) HasLanNetworkInfo() bool`

HasLanNetworkInfo returns a boolean if a field has been set.

### GetPublicIp

`func (o *SdWanMemberConfig) GetPublicIp() bool`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *SdWanMemberConfig) GetPublicIpOk() (*bool, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *SdWanMemberConfig) SetPublicIp(v bool)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *SdWanMemberConfig) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRole

`func (o *SdWanMemberConfig) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SdWanMemberConfig) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SdWanMemberConfig) SetRole(v int32)`

SetRole sets Role field to given value.


### GetSiteId

`func (o *SdWanMemberConfig) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SdWanMemberConfig) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SdWanMemberConfig) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetSiteName

`func (o *SdWanMemberConfig) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *SdWanMemberConfig) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *SdWanMemberConfig) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *SdWanMemberConfig) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetWanPortsInfo

`func (o *SdWanMemberConfig) GetWanPortsInfo() []OsgPortStatBrief`

GetWanPortsInfo returns the WanPortsInfo field if non-nil, zero value otherwise.

### GetWanPortsInfoOk

`func (o *SdWanMemberConfig) GetWanPortsInfoOk() (*[]OsgPortStatBrief, bool)`

GetWanPortsInfoOk returns a tuple with the WanPortsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortsInfo

`func (o *SdWanMemberConfig) SetWanPortsInfo(v []OsgPortStatBrief)`

SetWanPortsInfo sets WanPortsInfo field to given value.

### HasWanPortsInfo

`func (o *SdWanMemberConfig) HasWanPortsInfo() bool`

HasWanPortsInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


