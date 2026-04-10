# VpnTunnelRemoteStatusVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownBytes** | Pointer to **int64** |  | [optional] 
**DownPkts** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LocalIp** | Pointer to **string** |  | [optional] 
**LoginTime** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**RemoteIp** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**UpBytes** | Pointer to **int64** |  | [optional] 
**UpPkts** | Pointer to **int64** |  | [optional] 
**VpnId** | Pointer to **string** |  | [optional] 

## Methods

### NewVpnTunnelRemoteStatusVO

`func NewVpnTunnelRemoteStatusVO() *VpnTunnelRemoteStatusVO`

NewVpnTunnelRemoteStatusVO instantiates a new VpnTunnelRemoteStatusVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnTunnelRemoteStatusVOWithDefaults

`func NewVpnTunnelRemoteStatusVOWithDefaults() *VpnTunnelRemoteStatusVO`

NewVpnTunnelRemoteStatusVOWithDefaults instantiates a new VpnTunnelRemoteStatusVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownBytes

`func (o *VpnTunnelRemoteStatusVO) GetDownBytes() int64`

GetDownBytes returns the DownBytes field if non-nil, zero value otherwise.

### GetDownBytesOk

`func (o *VpnTunnelRemoteStatusVO) GetDownBytesOk() (*int64, bool)`

GetDownBytesOk returns a tuple with the DownBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownBytes

`func (o *VpnTunnelRemoteStatusVO) SetDownBytes(v int64)`

SetDownBytes sets DownBytes field to given value.

### HasDownBytes

`func (o *VpnTunnelRemoteStatusVO) HasDownBytes() bool`

HasDownBytes returns a boolean if a field has been set.

### GetDownPkts

`func (o *VpnTunnelRemoteStatusVO) GetDownPkts() int64`

GetDownPkts returns the DownPkts field if non-nil, zero value otherwise.

### GetDownPktsOk

`func (o *VpnTunnelRemoteStatusVO) GetDownPktsOk() (*int64, bool)`

GetDownPktsOk returns a tuple with the DownPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownPkts

`func (o *VpnTunnelRemoteStatusVO) SetDownPkts(v int64)`

SetDownPkts sets DownPkts field to given value.

### HasDownPkts

`func (o *VpnTunnelRemoteStatusVO) HasDownPkts() bool`

HasDownPkts returns a boolean if a field has been set.

### GetId

`func (o *VpnTunnelRemoteStatusVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpnTunnelRemoteStatusVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpnTunnelRemoteStatusVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VpnTunnelRemoteStatusVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocalIp

`func (o *VpnTunnelRemoteStatusVO) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *VpnTunnelRemoteStatusVO) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *VpnTunnelRemoteStatusVO) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.

### HasLocalIp

`func (o *VpnTunnelRemoteStatusVO) HasLocalIp() bool`

HasLocalIp returns a boolean if a field has been set.

### GetLoginTime

`func (o *VpnTunnelRemoteStatusVO) GetLoginTime() int64`

GetLoginTime returns the LoginTime field if non-nil, zero value otherwise.

### GetLoginTimeOk

`func (o *VpnTunnelRemoteStatusVO) GetLoginTimeOk() (*int64, bool)`

GetLoginTimeOk returns a tuple with the LoginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginTime

`func (o *VpnTunnelRemoteStatusVO) SetLoginTime(v int64)`

SetLoginTime sets LoginTime field to given value.

### HasLoginTime

`func (o *VpnTunnelRemoteStatusVO) HasLoginTime() bool`

HasLoginTime returns a boolean if a field has been set.

### GetName

`func (o *VpnTunnelRemoteStatusVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnTunnelRemoteStatusVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnTunnelRemoteStatusVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpnTunnelRemoteStatusVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *VpnTunnelRemoteStatusVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VpnTunnelRemoteStatusVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VpnTunnelRemoteStatusVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *VpnTunnelRemoteStatusVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRemoteIp

`func (o *VpnTunnelRemoteStatusVO) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *VpnTunnelRemoteStatusVO) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *VpnTunnelRemoteStatusVO) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *VpnTunnelRemoteStatusVO) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetStatus

`func (o *VpnTunnelRemoteStatusVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnTunnelRemoteStatusVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnTunnelRemoteStatusVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VpnTunnelRemoteStatusVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpBytes

`func (o *VpnTunnelRemoteStatusVO) GetUpBytes() int64`

GetUpBytes returns the UpBytes field if non-nil, zero value otherwise.

### GetUpBytesOk

`func (o *VpnTunnelRemoteStatusVO) GetUpBytesOk() (*int64, bool)`

GetUpBytesOk returns a tuple with the UpBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpBytes

`func (o *VpnTunnelRemoteStatusVO) SetUpBytes(v int64)`

SetUpBytes sets UpBytes field to given value.

### HasUpBytes

`func (o *VpnTunnelRemoteStatusVO) HasUpBytes() bool`

HasUpBytes returns a boolean if a field has been set.

### GetUpPkts

`func (o *VpnTunnelRemoteStatusVO) GetUpPkts() int64`

GetUpPkts returns the UpPkts field if non-nil, zero value otherwise.

### GetUpPktsOk

`func (o *VpnTunnelRemoteStatusVO) GetUpPktsOk() (*int64, bool)`

GetUpPktsOk returns a tuple with the UpPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPkts

`func (o *VpnTunnelRemoteStatusVO) SetUpPkts(v int64)`

SetUpPkts sets UpPkts field to given value.

### HasUpPkts

`func (o *VpnTunnelRemoteStatusVO) HasUpPkts() bool`

HasUpPkts returns a boolean if a field has been set.

### GetVpnId

`func (o *VpnTunnelRemoteStatusVO) GetVpnId() string`

GetVpnId returns the VpnId field if non-nil, zero value otherwise.

### GetVpnIdOk

`func (o *VpnTunnelRemoteStatusVO) GetVpnIdOk() (*string, bool)`

GetVpnIdOk returns a tuple with the VpnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnId

`func (o *VpnTunnelRemoteStatusVO) SetVpnId(v string)`

SetVpnId sets VpnId field to given value.

### HasVpnId

`func (o *VpnTunnelRemoteStatusVO) HasVpnId() bool`

HasVpnId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


