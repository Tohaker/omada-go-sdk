# SiteToSiteManualWgPeerConfigVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | The comment of WireGuard peer should contain 0 to 128 characters. | [optional] 
**KeepAlive** | **int32** | The keepalive second of WireGuard peer should be within the range of 0-65535. | 
**Name** | **string** | The name of WireGuard peer should contain 1 to 64 characters. | 
**PreSharedKey** | Pointer to **string** | The preSharedKey of WireGuard peer must have 44 character of base64 and end with &#39;&#x3D;&#39;. | [optional] 
**RemoteIp** | Pointer to **string** | Remote IP of the VPN | [optional] 
**RemoteSubnet** | **[]string** | IP/MASK address list of client WireGuard peer allowed. | 
**ServerPublicKey** | **string** | The public key of WireGuard VPN must have 44 character of base64 and end with &#39;&#x3D;&#39;. | 
**ServicePort** | Pointer to **int32** | Service port should be within the range of 1–65535. | [optional] 
**Status** | **bool** | Status of the VPN. | 

## Methods

### NewSiteToSiteManualWgPeerConfigVO

`func NewSiteToSiteManualWgPeerConfigVO(keepAlive int32, name string, remoteSubnet []string, serverPublicKey string, status bool, ) *SiteToSiteManualWgPeerConfigVO`

NewSiteToSiteManualWgPeerConfigVO instantiates a new SiteToSiteManualWgPeerConfigVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteToSiteManualWgPeerConfigVOWithDefaults

`func NewSiteToSiteManualWgPeerConfigVOWithDefaults() *SiteToSiteManualWgPeerConfigVO`

NewSiteToSiteManualWgPeerConfigVOWithDefaults instantiates a new SiteToSiteManualWgPeerConfigVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *SiteToSiteManualWgPeerConfigVO) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SiteToSiteManualWgPeerConfigVO) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SiteToSiteManualWgPeerConfigVO) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SiteToSiteManualWgPeerConfigVO) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetKeepAlive

`func (o *SiteToSiteManualWgPeerConfigVO) GetKeepAlive() int32`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *SiteToSiteManualWgPeerConfigVO) GetKeepAliveOk() (*int32, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *SiteToSiteManualWgPeerConfigVO) SetKeepAlive(v int32)`

SetKeepAlive sets KeepAlive field to given value.


### GetName

`func (o *SiteToSiteManualWgPeerConfigVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteToSiteManualWgPeerConfigVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteToSiteManualWgPeerConfigVO) SetName(v string)`

SetName sets Name field to given value.


### GetPreSharedKey

`func (o *SiteToSiteManualWgPeerConfigVO) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *SiteToSiteManualWgPeerConfigVO) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *SiteToSiteManualWgPeerConfigVO) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.

### HasPreSharedKey

`func (o *SiteToSiteManualWgPeerConfigVO) HasPreSharedKey() bool`

HasPreSharedKey returns a boolean if a field has been set.

### GetRemoteIp

`func (o *SiteToSiteManualWgPeerConfigVO) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *SiteToSiteManualWgPeerConfigVO) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *SiteToSiteManualWgPeerConfigVO) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *SiteToSiteManualWgPeerConfigVO) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetRemoteSubnet

`func (o *SiteToSiteManualWgPeerConfigVO) GetRemoteSubnet() []string`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *SiteToSiteManualWgPeerConfigVO) GetRemoteSubnetOk() (*[]string, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *SiteToSiteManualWgPeerConfigVO) SetRemoteSubnet(v []string)`

SetRemoteSubnet sets RemoteSubnet field to given value.


### GetServerPublicKey

`func (o *SiteToSiteManualWgPeerConfigVO) GetServerPublicKey() string`

GetServerPublicKey returns the ServerPublicKey field if non-nil, zero value otherwise.

### GetServerPublicKeyOk

`func (o *SiteToSiteManualWgPeerConfigVO) GetServerPublicKeyOk() (*string, bool)`

GetServerPublicKeyOk returns a tuple with the ServerPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPublicKey

`func (o *SiteToSiteManualWgPeerConfigVO) SetServerPublicKey(v string)`

SetServerPublicKey sets ServerPublicKey field to given value.


### GetServicePort

`func (o *SiteToSiteManualWgPeerConfigVO) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *SiteToSiteManualWgPeerConfigVO) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *SiteToSiteManualWgPeerConfigVO) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *SiteToSiteManualWgPeerConfigVO) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetStatus

`func (o *SiteToSiteManualWgPeerConfigVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SiteToSiteManualWgPeerConfigVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SiteToSiteManualWgPeerConfigVO) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


