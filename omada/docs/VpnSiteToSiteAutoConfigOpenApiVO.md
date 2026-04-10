# VpnSiteToSiteAutoConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name should contain 1 to 63 characters. | 
**RemoteSite** | **string** | Remote site of the VPN. | 
**Status** | **bool** | Status of the VPN. | 
**VpnType** | **int32** | Vpn type should be a value as follows: 2: IPSec; 4: WireGuard. | 

## Methods

### NewVpnSiteToSiteAutoConfigOpenApiVO

`func NewVpnSiteToSiteAutoConfigOpenApiVO(name string, remoteSite string, status bool, vpnType int32, ) *VpnSiteToSiteAutoConfigOpenApiVO`

NewVpnSiteToSiteAutoConfigOpenApiVO instantiates a new VpnSiteToSiteAutoConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnSiteToSiteAutoConfigOpenApiVOWithDefaults

`func NewVpnSiteToSiteAutoConfigOpenApiVOWithDefaults() *VpnSiteToSiteAutoConfigOpenApiVO`

NewVpnSiteToSiteAutoConfigOpenApiVOWithDefaults instantiates a new VpnSiteToSiteAutoConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VpnSiteToSiteAutoConfigOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnSiteToSiteAutoConfigOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnSiteToSiteAutoConfigOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetRemoteSite

`func (o *VpnSiteToSiteAutoConfigOpenApiVO) GetRemoteSite() string`

GetRemoteSite returns the RemoteSite field if non-nil, zero value otherwise.

### GetRemoteSiteOk

`func (o *VpnSiteToSiteAutoConfigOpenApiVO) GetRemoteSiteOk() (*string, bool)`

GetRemoteSiteOk returns a tuple with the RemoteSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSite

`func (o *VpnSiteToSiteAutoConfigOpenApiVO) SetRemoteSite(v string)`

SetRemoteSite sets RemoteSite field to given value.


### GetStatus

`func (o *VpnSiteToSiteAutoConfigOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnSiteToSiteAutoConfigOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnSiteToSiteAutoConfigOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetVpnType

`func (o *VpnSiteToSiteAutoConfigOpenApiVO) GetVpnType() int32`

GetVpnType returns the VpnType field if non-nil, zero value otherwise.

### GetVpnTypeOk

`func (o *VpnSiteToSiteAutoConfigOpenApiVO) GetVpnTypeOk() (*int32, bool)`

GetVpnTypeOk returns a tuple with the VpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnType

`func (o *VpnSiteToSiteAutoConfigOpenApiVO) SetVpnType(v int32)`

SetVpnType sets VpnType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


