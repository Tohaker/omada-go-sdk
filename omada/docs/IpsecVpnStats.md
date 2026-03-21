# IpsecVpnStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** | SA direction should be a value as follows: in; out | [optional] 
**Id** | Pointer to **int32** | VPN ID | [optional] 
**LocalPeerIp** | Pointer to **string** | IP address of the local peer | [optional] 
**LocalSa** | Pointer to **string** | Local network segment of SA Cover | [optional] 
**Name** | Pointer to **string** | VPN name | [optional] 
**RemotePeerIp** | Pointer to **string** | IP address of the remote peer | [optional] 
**RemoteSa** | Pointer to **string** | Remote network segment of SA Cover | [optional] 
**Status** | Pointer to **bool** | VPN status, whether to enable VPN  | [optional] 

## Methods

### NewIpsecVpnStats

`func NewIpsecVpnStats() *IpsecVpnStats`

NewIpsecVpnStats instantiates a new IpsecVpnStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpsecVpnStatsWithDefaults

`func NewIpsecVpnStatsWithDefaults() *IpsecVpnStats`

NewIpsecVpnStatsWithDefaults instantiates a new IpsecVpnStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *IpsecVpnStats) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *IpsecVpnStats) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *IpsecVpnStats) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *IpsecVpnStats) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetDirection

`func (o *IpsecVpnStats) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *IpsecVpnStats) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *IpsecVpnStats) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *IpsecVpnStats) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetId

`func (o *IpsecVpnStats) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpsecVpnStats) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpsecVpnStats) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IpsecVpnStats) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocalPeerIp

`func (o *IpsecVpnStats) GetLocalPeerIp() string`

GetLocalPeerIp returns the LocalPeerIp field if non-nil, zero value otherwise.

### GetLocalPeerIpOk

`func (o *IpsecVpnStats) GetLocalPeerIpOk() (*string, bool)`

GetLocalPeerIpOk returns a tuple with the LocalPeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPeerIp

`func (o *IpsecVpnStats) SetLocalPeerIp(v string)`

SetLocalPeerIp sets LocalPeerIp field to given value.

### HasLocalPeerIp

`func (o *IpsecVpnStats) HasLocalPeerIp() bool`

HasLocalPeerIp returns a boolean if a field has been set.

### GetLocalSa

`func (o *IpsecVpnStats) GetLocalSa() string`

GetLocalSa returns the LocalSa field if non-nil, zero value otherwise.

### GetLocalSaOk

`func (o *IpsecVpnStats) GetLocalSaOk() (*string, bool)`

GetLocalSaOk returns a tuple with the LocalSa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSa

`func (o *IpsecVpnStats) SetLocalSa(v string)`

SetLocalSa sets LocalSa field to given value.

### HasLocalSa

`func (o *IpsecVpnStats) HasLocalSa() bool`

HasLocalSa returns a boolean if a field has been set.

### GetName

`func (o *IpsecVpnStats) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpsecVpnStats) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpsecVpnStats) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IpsecVpnStats) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemotePeerIp

`func (o *IpsecVpnStats) GetRemotePeerIp() string`

GetRemotePeerIp returns the RemotePeerIp field if non-nil, zero value otherwise.

### GetRemotePeerIpOk

`func (o *IpsecVpnStats) GetRemotePeerIpOk() (*string, bool)`

GetRemotePeerIpOk returns a tuple with the RemotePeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePeerIp

`func (o *IpsecVpnStats) SetRemotePeerIp(v string)`

SetRemotePeerIp sets RemotePeerIp field to given value.

### HasRemotePeerIp

`func (o *IpsecVpnStats) HasRemotePeerIp() bool`

HasRemotePeerIp returns a boolean if a field has been set.

### GetRemoteSa

`func (o *IpsecVpnStats) GetRemoteSa() string`

GetRemoteSa returns the RemoteSa field if non-nil, zero value otherwise.

### GetRemoteSaOk

`func (o *IpsecVpnStats) GetRemoteSaOk() (*string, bool)`

GetRemoteSaOk returns a tuple with the RemoteSa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSa

`func (o *IpsecVpnStats) SetRemoteSa(v string)`

SetRemoteSa sets RemoteSa field to given value.

### HasRemoteSa

`func (o *IpsecVpnStats) HasRemoteSa() bool`

HasRemoteSa returns a boolean if a field has been set.

### GetStatus

`func (o *IpsecVpnStats) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IpsecVpnStats) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IpsecVpnStats) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IpsecVpnStats) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


