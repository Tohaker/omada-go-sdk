# VpnSummaryVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveClients** | Pointer to **int64** | Active clients number of the VPN. | [optional] 
**CustomNetwork** | Pointer to [**[]IPSubnetsVO**](IPSubnetsVO.md) | Custom networks of the VPN. | [optional] 
**FailoverSetting** | Pointer to [**IPsecFailoverStatusOpenApiVO**](IPsecFailoverStatusOpenApiVO.md) |  | [optional] 
**FeatureDescription** | Pointer to [**[]FeatureInfoVO**](FeatureInfoVO.md) | Gateway Feature Description. | [optional] 
**GroupId** | Pointer to **string** | Default group ID of SSL VPN server. | [optional] 
**Id** | Pointer to **string** | ID of the VPN. | [optional] 
**Name** | Pointer to **string** | VPN name. | [optional] 
**NetworkList** | Pointer to **[]string** | Network list of the VPN. | [optional] 
**NetworkType** | Pointer to **int32** | Network type. 0: network list; 1: custom networks. | [optional] 
**SiteVpnType** | Pointer to **int32** | Site VPN type of the VPN. 0: Auto; 1: Manual. | [optional] 
**Status** | Pointer to **bool** | Status of the VPN. | [optional] 
**UserCount** | Pointer to **int64** | User number of the VPN. | [optional] 
**VpnType** | Pointer to **int32** | Server Vpn type. 0: L2TP; 1: PPTP; 2: IPSec; 3: OpenVPN; 4: WireGuard; 5: SSL VPN. | [optional] 
**Wans** | Pointer to **[]string** | WAN port ID. | [optional] 

## Methods

### NewVpnSummaryVO

`func NewVpnSummaryVO() *VpnSummaryVO`

NewVpnSummaryVO instantiates a new VpnSummaryVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnSummaryVOWithDefaults

`func NewVpnSummaryVOWithDefaults() *VpnSummaryVO`

NewVpnSummaryVOWithDefaults instantiates a new VpnSummaryVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveClients

`func (o *VpnSummaryVO) GetActiveClients() int64`

GetActiveClients returns the ActiveClients field if non-nil, zero value otherwise.

### GetActiveClientsOk

`func (o *VpnSummaryVO) GetActiveClientsOk() (*int64, bool)`

GetActiveClientsOk returns a tuple with the ActiveClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveClients

`func (o *VpnSummaryVO) SetActiveClients(v int64)`

SetActiveClients sets ActiveClients field to given value.

### HasActiveClients

`func (o *VpnSummaryVO) HasActiveClients() bool`

HasActiveClients returns a boolean if a field has been set.

### GetCustomNetwork

`func (o *VpnSummaryVO) GetCustomNetwork() []IPSubnetsVO`

GetCustomNetwork returns the CustomNetwork field if non-nil, zero value otherwise.

### GetCustomNetworkOk

`func (o *VpnSummaryVO) GetCustomNetworkOk() (*[]IPSubnetsVO, bool)`

GetCustomNetworkOk returns a tuple with the CustomNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetwork

`func (o *VpnSummaryVO) SetCustomNetwork(v []IPSubnetsVO)`

SetCustomNetwork sets CustomNetwork field to given value.

### HasCustomNetwork

`func (o *VpnSummaryVO) HasCustomNetwork() bool`

HasCustomNetwork returns a boolean if a field has been set.

### GetFailoverSetting

`func (o *VpnSummaryVO) GetFailoverSetting() IPsecFailoverStatusOpenApiVO`

GetFailoverSetting returns the FailoverSetting field if non-nil, zero value otherwise.

### GetFailoverSettingOk

`func (o *VpnSummaryVO) GetFailoverSettingOk() (*IPsecFailoverStatusOpenApiVO, bool)`

GetFailoverSettingOk returns a tuple with the FailoverSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverSetting

`func (o *VpnSummaryVO) SetFailoverSetting(v IPsecFailoverStatusOpenApiVO)`

SetFailoverSetting sets FailoverSetting field to given value.

### HasFailoverSetting

`func (o *VpnSummaryVO) HasFailoverSetting() bool`

HasFailoverSetting returns a boolean if a field has been set.

### GetFeatureDescription

`func (o *VpnSummaryVO) GetFeatureDescription() []FeatureInfoVO`

GetFeatureDescription returns the FeatureDescription field if non-nil, zero value otherwise.

### GetFeatureDescriptionOk

`func (o *VpnSummaryVO) GetFeatureDescriptionOk() (*[]FeatureInfoVO, bool)`

GetFeatureDescriptionOk returns a tuple with the FeatureDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureDescription

`func (o *VpnSummaryVO) SetFeatureDescription(v []FeatureInfoVO)`

SetFeatureDescription sets FeatureDescription field to given value.

### HasFeatureDescription

`func (o *VpnSummaryVO) HasFeatureDescription() bool`

HasFeatureDescription returns a boolean if a field has been set.

### GetGroupId

`func (o *VpnSummaryVO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *VpnSummaryVO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *VpnSummaryVO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *VpnSummaryVO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *VpnSummaryVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpnSummaryVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpnSummaryVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VpnSummaryVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VpnSummaryVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnSummaryVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnSummaryVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpnSummaryVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkList

`func (o *VpnSummaryVO) GetNetworkList() []string`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *VpnSummaryVO) GetNetworkListOk() (*[]string, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *VpnSummaryVO) SetNetworkList(v []string)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *VpnSummaryVO) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.

### GetNetworkType

`func (o *VpnSummaryVO) GetNetworkType() int32`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *VpnSummaryVO) GetNetworkTypeOk() (*int32, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *VpnSummaryVO) SetNetworkType(v int32)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *VpnSummaryVO) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetSiteVpnType

`func (o *VpnSummaryVO) GetSiteVpnType() int32`

GetSiteVpnType returns the SiteVpnType field if non-nil, zero value otherwise.

### GetSiteVpnTypeOk

`func (o *VpnSummaryVO) GetSiteVpnTypeOk() (*int32, bool)`

GetSiteVpnTypeOk returns a tuple with the SiteVpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteVpnType

`func (o *VpnSummaryVO) SetSiteVpnType(v int32)`

SetSiteVpnType sets SiteVpnType field to given value.

### HasSiteVpnType

`func (o *VpnSummaryVO) HasSiteVpnType() bool`

HasSiteVpnType returns a boolean if a field has been set.

### GetStatus

`func (o *VpnSummaryVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnSummaryVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnSummaryVO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VpnSummaryVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserCount

`func (o *VpnSummaryVO) GetUserCount() int64`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *VpnSummaryVO) GetUserCountOk() (*int64, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *VpnSummaryVO) SetUserCount(v int64)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *VpnSummaryVO) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.

### GetVpnType

`func (o *VpnSummaryVO) GetVpnType() int32`

GetVpnType returns the VpnType field if non-nil, zero value otherwise.

### GetVpnTypeOk

`func (o *VpnSummaryVO) GetVpnTypeOk() (*int32, bool)`

GetVpnTypeOk returns a tuple with the VpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnType

`func (o *VpnSummaryVO) SetVpnType(v int32)`

SetVpnType sets VpnType field to given value.

### HasVpnType

`func (o *VpnSummaryVO) HasVpnType() bool`

HasVpnType returns a boolean if a field has been set.

### GetWans

`func (o *VpnSummaryVO) GetWans() []string`

GetWans returns the Wans field if non-nil, zero value otherwise.

### GetWansOk

`func (o *VpnSummaryVO) GetWansOk() (*[]string, bool)`

GetWansOk returns a tuple with the Wans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWans

`func (o *VpnSummaryVO) SetWans(v []string)`

SetWans sets Wans field to given value.

### HasWans

`func (o *VpnSummaryVO) HasWans() bool`

HasWans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


