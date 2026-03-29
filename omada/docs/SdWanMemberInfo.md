# SdWanMemberInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMac** | Pointer to **string** | The device MAC of the sdWan member. | [optional] 
**DeviceName** | Pointer to **string** | The device name of the sdWan member. | [optional] 
**LanNetworkInfo** | Pointer to [**[]LanNetworkBrief**](LanNetworkBrief.md) | A list of lan network info for SdWan Member. | [optional] 
**LinkedToHub** | Pointer to **int32** | If a member is a spoke, the link connection to the hub is identified. | [optional] 
**Model** | Pointer to **string** | The device model of the sdWan member. | [optional] 
**ModelVersion** | Pointer to **string** | The device model version of the sdWan member. | [optional] 
**OnlineStatus** | Pointer to **int32** | The device online status of the sdWan member. | [optional] 
**PublicIp** | Pointer to **bool** | Whether the sdWan member has a public IP. | [optional] 
**Role** | **int32** | The role of sdWan member, hub or spoke. | 
**SdWanIp** | Pointer to **string** | The sdWan IP of the sdWan member. | [optional] 
**ShowModel** | Pointer to **string** | The device showmodel of the sdWan member. | [optional] 
**SiteId** | Pointer to **string** | The ID of the site where the sdWan member is located. | [optional] 
**SiteName** | Pointer to **string** | The name of the site where the sdWan member is located. | [optional] 
**Type** | Pointer to **string** | The device type of the sdWan member. | [optional] 
**WanPortsInfo** | Pointer to [**[]OsgPortStatBrief**](OsgPortStatBrief.md) | A list of device port status info for SdWan Member. | [optional] 

## Methods

### NewSdWanMemberInfo

`func NewSdWanMemberInfo(role int32, ) *SdWanMemberInfo`

NewSdWanMemberInfo instantiates a new SdWanMemberInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdWanMemberInfoWithDefaults

`func NewSdWanMemberInfoWithDefaults() *SdWanMemberInfo`

NewSdWanMemberInfoWithDefaults instantiates a new SdWanMemberInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMac

`func (o *SdWanMemberInfo) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *SdWanMemberInfo) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *SdWanMemberInfo) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *SdWanMemberInfo) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDeviceName

`func (o *SdWanMemberInfo) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *SdWanMemberInfo) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *SdWanMemberInfo) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *SdWanMemberInfo) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetLanNetworkInfo

`func (o *SdWanMemberInfo) GetLanNetworkInfo() []LanNetworkBrief`

GetLanNetworkInfo returns the LanNetworkInfo field if non-nil, zero value otherwise.

### GetLanNetworkInfoOk

`func (o *SdWanMemberInfo) GetLanNetworkInfoOk() (*[]LanNetworkBrief, bool)`

GetLanNetworkInfoOk returns a tuple with the LanNetworkInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetworkInfo

`func (o *SdWanMemberInfo) SetLanNetworkInfo(v []LanNetworkBrief)`

SetLanNetworkInfo sets LanNetworkInfo field to given value.

### HasLanNetworkInfo

`func (o *SdWanMemberInfo) HasLanNetworkInfo() bool`

HasLanNetworkInfo returns a boolean if a field has been set.

### GetLinkedToHub

`func (o *SdWanMemberInfo) GetLinkedToHub() int32`

GetLinkedToHub returns the LinkedToHub field if non-nil, zero value otherwise.

### GetLinkedToHubOk

`func (o *SdWanMemberInfo) GetLinkedToHubOk() (*int32, bool)`

GetLinkedToHubOk returns a tuple with the LinkedToHub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedToHub

`func (o *SdWanMemberInfo) SetLinkedToHub(v int32)`

SetLinkedToHub sets LinkedToHub field to given value.

### HasLinkedToHub

`func (o *SdWanMemberInfo) HasLinkedToHub() bool`

HasLinkedToHub returns a boolean if a field has been set.

### GetModel

`func (o *SdWanMemberInfo) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SdWanMemberInfo) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SdWanMemberInfo) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SdWanMemberInfo) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *SdWanMemberInfo) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *SdWanMemberInfo) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *SdWanMemberInfo) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *SdWanMemberInfo) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetOnlineStatus

`func (o *SdWanMemberInfo) GetOnlineStatus() int32`

GetOnlineStatus returns the OnlineStatus field if non-nil, zero value otherwise.

### GetOnlineStatusOk

`func (o *SdWanMemberInfo) GetOnlineStatusOk() (*int32, bool)`

GetOnlineStatusOk returns a tuple with the OnlineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineStatus

`func (o *SdWanMemberInfo) SetOnlineStatus(v int32)`

SetOnlineStatus sets OnlineStatus field to given value.

### HasOnlineStatus

`func (o *SdWanMemberInfo) HasOnlineStatus() bool`

HasOnlineStatus returns a boolean if a field has been set.

### GetPublicIp

`func (o *SdWanMemberInfo) GetPublicIp() bool`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *SdWanMemberInfo) GetPublicIpOk() (*bool, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *SdWanMemberInfo) SetPublicIp(v bool)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *SdWanMemberInfo) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRole

`func (o *SdWanMemberInfo) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SdWanMemberInfo) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SdWanMemberInfo) SetRole(v int32)`

SetRole sets Role field to given value.


### GetSdWanIp

`func (o *SdWanMemberInfo) GetSdWanIp() string`

GetSdWanIp returns the SdWanIp field if non-nil, zero value otherwise.

### GetSdWanIpOk

`func (o *SdWanMemberInfo) GetSdWanIpOk() (*string, bool)`

GetSdWanIpOk returns a tuple with the SdWanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdWanIp

`func (o *SdWanMemberInfo) SetSdWanIp(v string)`

SetSdWanIp sets SdWanIp field to given value.

### HasSdWanIp

`func (o *SdWanMemberInfo) HasSdWanIp() bool`

HasSdWanIp returns a boolean if a field has been set.

### GetShowModel

`func (o *SdWanMemberInfo) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *SdWanMemberInfo) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *SdWanMemberInfo) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *SdWanMemberInfo) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetSiteId

`func (o *SdWanMemberInfo) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SdWanMemberInfo) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SdWanMemberInfo) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SdWanMemberInfo) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteName

`func (o *SdWanMemberInfo) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *SdWanMemberInfo) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *SdWanMemberInfo) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *SdWanMemberInfo) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetType

`func (o *SdWanMemberInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SdWanMemberInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SdWanMemberInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SdWanMemberInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWanPortsInfo

`func (o *SdWanMemberInfo) GetWanPortsInfo() []OsgPortStatBrief`

GetWanPortsInfo returns the WanPortsInfo field if non-nil, zero value otherwise.

### GetWanPortsInfoOk

`func (o *SdWanMemberInfo) GetWanPortsInfoOk() (*[]OsgPortStatBrief, bool)`

GetWanPortsInfoOk returns a tuple with the WanPortsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortsInfo

`func (o *SdWanMemberInfo) SetWanPortsInfo(v []OsgPortStatBrief)`

SetWanPortsInfo sets WanPortsInfo field to given value.

### HasWanPortsInfo

`func (o *SdWanMemberInfo) HasWanPortsInfo() bool`

HasWanPortsInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


