# VpnUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientMode** | Pointer to **int32** | Client mode should be a value as follows: 0: Client-To-Site, 1: Site-To-Site. | [optional] 
**LocalIp** | Pointer to **string** | Local IP of the VPN user. | [optional] 
**MaxConnections** | Pointer to **int32** | Max connections should be within the range of 1–100. | [optional] 
**Password** | **string** | Password of the VPN user. | 
**Protocol** | Pointer to **int32** | Protocol should be a value as follows: 0: L2TP or PPTP; 1: openVPN. | [optional] 
**Servers** | **[]string** | Servers of the VPN user. Server can be created using &#39;Create client-to-site VPN server&#39; interface, and server ID can be obtained from &#39;Get client-to-site VPN server list&#39; interface. | 
**UserRemoteSubnets** | Pointer to [**[]IPSubnetsVO**](IPSubnetsVO.md) | User remote subnets of the VPN user. | [optional] 
**Username** | **string** | Username of the VPN user. | 

## Methods

### NewVpnUserRequest

`func NewVpnUserRequest(password string, servers []string, username string, ) *VpnUserRequest`

NewVpnUserRequest instantiates a new VpnUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnUserRequestWithDefaults

`func NewVpnUserRequestWithDefaults() *VpnUserRequest`

NewVpnUserRequestWithDefaults instantiates a new VpnUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientMode

`func (o *VpnUserRequest) GetClientMode() int32`

GetClientMode returns the ClientMode field if non-nil, zero value otherwise.

### GetClientModeOk

`func (o *VpnUserRequest) GetClientModeOk() (*int32, bool)`

GetClientModeOk returns a tuple with the ClientMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMode

`func (o *VpnUserRequest) SetClientMode(v int32)`

SetClientMode sets ClientMode field to given value.

### HasClientMode

`func (o *VpnUserRequest) HasClientMode() bool`

HasClientMode returns a boolean if a field has been set.

### GetLocalIp

`func (o *VpnUserRequest) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *VpnUserRequest) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *VpnUserRequest) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.

### HasLocalIp

`func (o *VpnUserRequest) HasLocalIp() bool`

HasLocalIp returns a boolean if a field has been set.

### GetMaxConnections

`func (o *VpnUserRequest) GetMaxConnections() int32`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *VpnUserRequest) GetMaxConnectionsOk() (*int32, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *VpnUserRequest) SetMaxConnections(v int32)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *VpnUserRequest) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetPassword

`func (o *VpnUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VpnUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VpnUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetProtocol

`func (o *VpnUserRequest) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *VpnUserRequest) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *VpnUserRequest) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *VpnUserRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetServers

`func (o *VpnUserRequest) GetServers() []string`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *VpnUserRequest) GetServersOk() (*[]string, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *VpnUserRequest) SetServers(v []string)`

SetServers sets Servers field to given value.


### GetUserRemoteSubnets

`func (o *VpnUserRequest) GetUserRemoteSubnets() []IPSubnetsVO`

GetUserRemoteSubnets returns the UserRemoteSubnets field if non-nil, zero value otherwise.

### GetUserRemoteSubnetsOk

`func (o *VpnUserRequest) GetUserRemoteSubnetsOk() (*[]IPSubnetsVO, bool)`

GetUserRemoteSubnetsOk returns a tuple with the UserRemoteSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRemoteSubnets

`func (o *VpnUserRequest) SetUserRemoteSubnets(v []IPSubnetsVO)`

SetUserRemoteSubnets sets UserRemoteSubnets field to given value.

### HasUserRemoteSubnets

`func (o *VpnUserRequest) HasUserRemoteSubnets() bool`

HasUserRemoteSubnets returns a boolean if a field has been set.

### GetUsername

`func (o *VpnUserRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VpnUserRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VpnUserRequest) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


