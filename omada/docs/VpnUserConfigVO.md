# VpnUserConfigVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientMode** | Pointer to **int32** | Client mode should be a value as follows: 0: Client-To-Site, 1: Site-To-Site. | [optional] 
**GroupId** | Pointer to **string** | Group ID of the SSL VPN user. User group can be created using &#39;Create SSL VPN user group&#39; interface, and User Group ID can be obtained from &#39;Get user group list for SSL VPN server&#39; interface. | [optional] 
**LocalIp** | Pointer to **string** | Local IP of the VPN user. | [optional] 
**MaxConnections** | Pointer to **int32** | Max connections should be within the range of 1–100. | [optional] 
**Password** | Pointer to **string** | Password of the VPN user. | [optional] 
**Protocol** | Pointer to **int32** | Protocol should be a value as follows: 0: L2TP or PPTP; 1: openVPN; 2: SSL VPN. | [optional] 
**Servers** | Pointer to **[]string** | Servers of the VPN user. Server can be created using &#39;Create client-to-site VPN server&#39; interface, and server ID can be obtained from &#39;Get client-to-site VPN server list&#39; interface. | [optional] 
**Status** | Pointer to **bool** | Status of the SSL VPN user. | [optional] 
**UserRemoteSubnets** | Pointer to [**[]IPSubnetsVO**](IPSubnetsVO.md) | User remote subnets of the VPN user. | [optional] 
**Username** | Pointer to **string** | Username of the VPN user. | [optional] 
**Validity** | Pointer to **string** | Validity of the SSL VPN user. The format is Month/Day/Year, for example 08/20/2022. | [optional] 

## Methods

### NewVpnUserConfigVO

`func NewVpnUserConfigVO() *VpnUserConfigVO`

NewVpnUserConfigVO instantiates a new VpnUserConfigVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnUserConfigVOWithDefaults

`func NewVpnUserConfigVOWithDefaults() *VpnUserConfigVO`

NewVpnUserConfigVOWithDefaults instantiates a new VpnUserConfigVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientMode

`func (o *VpnUserConfigVO) GetClientMode() int32`

GetClientMode returns the ClientMode field if non-nil, zero value otherwise.

### GetClientModeOk

`func (o *VpnUserConfigVO) GetClientModeOk() (*int32, bool)`

GetClientModeOk returns a tuple with the ClientMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMode

`func (o *VpnUserConfigVO) SetClientMode(v int32)`

SetClientMode sets ClientMode field to given value.

### HasClientMode

`func (o *VpnUserConfigVO) HasClientMode() bool`

HasClientMode returns a boolean if a field has been set.

### GetGroupId

`func (o *VpnUserConfigVO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *VpnUserConfigVO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *VpnUserConfigVO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *VpnUserConfigVO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetLocalIp

`func (o *VpnUserConfigVO) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *VpnUserConfigVO) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *VpnUserConfigVO) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.

### HasLocalIp

`func (o *VpnUserConfigVO) HasLocalIp() bool`

HasLocalIp returns a boolean if a field has been set.

### GetMaxConnections

`func (o *VpnUserConfigVO) GetMaxConnections() int32`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *VpnUserConfigVO) GetMaxConnectionsOk() (*int32, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *VpnUserConfigVO) SetMaxConnections(v int32)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *VpnUserConfigVO) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetPassword

`func (o *VpnUserConfigVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VpnUserConfigVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VpnUserConfigVO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VpnUserConfigVO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProtocol

`func (o *VpnUserConfigVO) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *VpnUserConfigVO) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *VpnUserConfigVO) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *VpnUserConfigVO) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetServers

`func (o *VpnUserConfigVO) GetServers() []string`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *VpnUserConfigVO) GetServersOk() (*[]string, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *VpnUserConfigVO) SetServers(v []string)`

SetServers sets Servers field to given value.

### HasServers

`func (o *VpnUserConfigVO) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetStatus

`func (o *VpnUserConfigVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnUserConfigVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnUserConfigVO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VpnUserConfigVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserRemoteSubnets

`func (o *VpnUserConfigVO) GetUserRemoteSubnets() []IPSubnetsVO`

GetUserRemoteSubnets returns the UserRemoteSubnets field if non-nil, zero value otherwise.

### GetUserRemoteSubnetsOk

`func (o *VpnUserConfigVO) GetUserRemoteSubnetsOk() (*[]IPSubnetsVO, bool)`

GetUserRemoteSubnetsOk returns a tuple with the UserRemoteSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRemoteSubnets

`func (o *VpnUserConfigVO) SetUserRemoteSubnets(v []IPSubnetsVO)`

SetUserRemoteSubnets sets UserRemoteSubnets field to given value.

### HasUserRemoteSubnets

`func (o *VpnUserConfigVO) HasUserRemoteSubnets() bool`

HasUserRemoteSubnets returns a boolean if a field has been set.

### GetUsername

`func (o *VpnUserConfigVO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VpnUserConfigVO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VpnUserConfigVO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *VpnUserConfigVO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetValidity

`func (o *VpnUserConfigVO) GetValidity() string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *VpnUserConfigVO) GetValidityOk() (*string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *VpnUserConfigVO) SetValidity(v string)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *VpnUserConfigVO) HasValidity() bool`

HasValidity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


