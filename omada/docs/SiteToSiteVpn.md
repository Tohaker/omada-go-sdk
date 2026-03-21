# SiteToSiteVpn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedSetting** | Pointer to [**VpnAdvancedSettingOpenApiVO**](VpnAdvancedSettingOpenApiVO.md) |  | [optional] 
**CustomNetwork** | Pointer to [**[]IPSubnetsVO**](IPSubnetsVO.md) | Custom networks of the VPN, only for Manual IPSec type. | [optional] 
**Id** | Pointer to **string** | ID of the VPN. | [optional] 
**Name** | **string** | Name should contain 1 to 63 characters. | 
**NetworkList** | Pointer to **[]string** | Network list of the VPN, only for Manual IPSec type. Network can be created using &#39;Create LAN network&#39; interface, and network ID can be obtained from &#39;Get LAN network list&#39; interface. | [optional] 
**NetworkType** | Pointer to **int32** | Network type should be a value as follows: 0: network list; 1: custom networks. | [optional] 
**PreSharedKey** | Pointer to **string** | Pre-shared key of the VPN, only for Manual IPSec type. | [optional] 
**RemoteIp** | Pointer to **string** | Remote IP of the VPN, only for Manual IPSec type. | [optional] 
**RemoteSite** | Pointer to **string** | Remote site of the VPN, only for Auto IPSec type. | [optional] 
**RemoteSubnet** | Pointer to [**[]IPSubnetsVO**](IPSubnetsVO.md) | Remote subnet of the VPN, only for Manual IPSec type. | [optional] 
**SiteVpnType** | **int32** | Site VPN type should be a value as follows: 0: Auto IPSec; 1: Manual IPSec. | 
**Status** | **bool** | Status of the VPN. | 
**Wan** | Pointer to **[]string** | WAN list of the VPN, only for Manual IPSec type. WAN port ID can be obtained from &#39;Get internet basic info&#39; interface. | [optional] 

## Methods

### NewSiteToSiteVpn

`func NewSiteToSiteVpn(name string, siteVpnType int32, status bool, ) *SiteToSiteVpn`

NewSiteToSiteVpn instantiates a new SiteToSiteVpn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteToSiteVpnWithDefaults

`func NewSiteToSiteVpnWithDefaults() *SiteToSiteVpn`

NewSiteToSiteVpnWithDefaults instantiates a new SiteToSiteVpn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedSetting

`func (o *SiteToSiteVpn) GetAdvancedSetting() VpnAdvancedSettingOpenApiVO`

GetAdvancedSetting returns the AdvancedSetting field if non-nil, zero value otherwise.

### GetAdvancedSettingOk

`func (o *SiteToSiteVpn) GetAdvancedSettingOk() (*VpnAdvancedSettingOpenApiVO, bool)`

GetAdvancedSettingOk returns a tuple with the AdvancedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSetting

`func (o *SiteToSiteVpn) SetAdvancedSetting(v VpnAdvancedSettingOpenApiVO)`

SetAdvancedSetting sets AdvancedSetting field to given value.

### HasAdvancedSetting

`func (o *SiteToSiteVpn) HasAdvancedSetting() bool`

HasAdvancedSetting returns a boolean if a field has been set.

### GetCustomNetwork

`func (o *SiteToSiteVpn) GetCustomNetwork() []IPSubnetsVO`

GetCustomNetwork returns the CustomNetwork field if non-nil, zero value otherwise.

### GetCustomNetworkOk

`func (o *SiteToSiteVpn) GetCustomNetworkOk() (*[]IPSubnetsVO, bool)`

GetCustomNetworkOk returns a tuple with the CustomNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetwork

`func (o *SiteToSiteVpn) SetCustomNetwork(v []IPSubnetsVO)`

SetCustomNetwork sets CustomNetwork field to given value.

### HasCustomNetwork

`func (o *SiteToSiteVpn) HasCustomNetwork() bool`

HasCustomNetwork returns a boolean if a field has been set.

### GetId

`func (o *SiteToSiteVpn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteToSiteVpn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteToSiteVpn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SiteToSiteVpn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SiteToSiteVpn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteToSiteVpn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteToSiteVpn) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkList

`func (o *SiteToSiteVpn) GetNetworkList() []string`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *SiteToSiteVpn) GetNetworkListOk() (*[]string, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *SiteToSiteVpn) SetNetworkList(v []string)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *SiteToSiteVpn) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.

### GetNetworkType

`func (o *SiteToSiteVpn) GetNetworkType() int32`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *SiteToSiteVpn) GetNetworkTypeOk() (*int32, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *SiteToSiteVpn) SetNetworkType(v int32)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *SiteToSiteVpn) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetPreSharedKey

`func (o *SiteToSiteVpn) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *SiteToSiteVpn) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *SiteToSiteVpn) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.

### HasPreSharedKey

`func (o *SiteToSiteVpn) HasPreSharedKey() bool`

HasPreSharedKey returns a boolean if a field has been set.

### GetRemoteIp

`func (o *SiteToSiteVpn) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *SiteToSiteVpn) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *SiteToSiteVpn) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *SiteToSiteVpn) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetRemoteSite

`func (o *SiteToSiteVpn) GetRemoteSite() string`

GetRemoteSite returns the RemoteSite field if non-nil, zero value otherwise.

### GetRemoteSiteOk

`func (o *SiteToSiteVpn) GetRemoteSiteOk() (*string, bool)`

GetRemoteSiteOk returns a tuple with the RemoteSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSite

`func (o *SiteToSiteVpn) SetRemoteSite(v string)`

SetRemoteSite sets RemoteSite field to given value.

### HasRemoteSite

`func (o *SiteToSiteVpn) HasRemoteSite() bool`

HasRemoteSite returns a boolean if a field has been set.

### GetRemoteSubnet

`func (o *SiteToSiteVpn) GetRemoteSubnet() []IPSubnetsVO`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *SiteToSiteVpn) GetRemoteSubnetOk() (*[]IPSubnetsVO, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *SiteToSiteVpn) SetRemoteSubnet(v []IPSubnetsVO)`

SetRemoteSubnet sets RemoteSubnet field to given value.

### HasRemoteSubnet

`func (o *SiteToSiteVpn) HasRemoteSubnet() bool`

HasRemoteSubnet returns a boolean if a field has been set.

### GetSiteVpnType

`func (o *SiteToSiteVpn) GetSiteVpnType() int32`

GetSiteVpnType returns the SiteVpnType field if non-nil, zero value otherwise.

### GetSiteVpnTypeOk

`func (o *SiteToSiteVpn) GetSiteVpnTypeOk() (*int32, bool)`

GetSiteVpnTypeOk returns a tuple with the SiteVpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteVpnType

`func (o *SiteToSiteVpn) SetSiteVpnType(v int32)`

SetSiteVpnType sets SiteVpnType field to given value.


### GetStatus

`func (o *SiteToSiteVpn) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SiteToSiteVpn) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SiteToSiteVpn) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetWan

`func (o *SiteToSiteVpn) GetWan() []string`

GetWan returns the Wan field if non-nil, zero value otherwise.

### GetWanOk

`func (o *SiteToSiteVpn) GetWanOk() (*[]string, bool)`

GetWanOk returns a tuple with the Wan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWan

`func (o *SiteToSiteVpn) SetWan(v []string)`

SetWan sets Wan field to given value.

### HasWan

`func (o *SiteToSiteVpn) HasWan() bool`

HasWan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


