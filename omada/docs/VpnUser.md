# VpnUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientMode** | Pointer to **int32** | Client mode should be a value as follows: 0: Client-To-Site, 1: Site-To-Site. | [optional] 
**Id** | Pointer to **string** | ID of the VPN user. | [optional] 
**LocalIp** | Pointer to **string** | Local IP of the VPN user. | [optional] 
**MaxConnections** | Pointer to **int32** | Max connections should be within the range of 1–100. | [optional] 
**Password** | Pointer to **string** | Password of the VPN user. | [optional] 
**Protocol** | Pointer to **int32** | Protocol should be a value as follows: 0: L2TP or PPTP; 1: openVPN. | [optional] 
**ServerNames** | Pointer to **string** | Server names of the VPN user. | [optional] 
**Servers** | Pointer to **[]string** | Servers of the VPN user. Server can be created using &#39;Create client-to-site VPN server&#39; interface, and server ID can be obtained from &#39;Get client-to-site VPN server list&#39; interface. | [optional] 
**UserRemoteSubnets** | Pointer to [**[]IPSubnetsVO**](IPSubnetsVO.md) | User remote subnets of the VPN user. | [optional] 
**Username** | Pointer to **string** | Username of the VPN user. | [optional] 

## Methods

### NewVpnUser

`func NewVpnUser() *VpnUser`

NewVpnUser instantiates a new VpnUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnUserWithDefaults

`func NewVpnUserWithDefaults() *VpnUser`

NewVpnUserWithDefaults instantiates a new VpnUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientMode

`func (o *VpnUser) GetClientMode() int32`

GetClientMode returns the ClientMode field if non-nil, zero value otherwise.

### GetClientModeOk

`func (o *VpnUser) GetClientModeOk() (*int32, bool)`

GetClientModeOk returns a tuple with the ClientMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMode

`func (o *VpnUser) SetClientMode(v int32)`

SetClientMode sets ClientMode field to given value.

### HasClientMode

`func (o *VpnUser) HasClientMode() bool`

HasClientMode returns a boolean if a field has been set.

### GetId

`func (o *VpnUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpnUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpnUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VpnUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocalIp

`func (o *VpnUser) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *VpnUser) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *VpnUser) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.

### HasLocalIp

`func (o *VpnUser) HasLocalIp() bool`

HasLocalIp returns a boolean if a field has been set.

### GetMaxConnections

`func (o *VpnUser) GetMaxConnections() int32`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *VpnUser) GetMaxConnectionsOk() (*int32, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *VpnUser) SetMaxConnections(v int32)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *VpnUser) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetPassword

`func (o *VpnUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VpnUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VpnUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VpnUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProtocol

`func (o *VpnUser) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *VpnUser) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *VpnUser) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *VpnUser) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetServerNames

`func (o *VpnUser) GetServerNames() string`

GetServerNames returns the ServerNames field if non-nil, zero value otherwise.

### GetServerNamesOk

`func (o *VpnUser) GetServerNamesOk() (*string, bool)`

GetServerNamesOk returns a tuple with the ServerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNames

`func (o *VpnUser) SetServerNames(v string)`

SetServerNames sets ServerNames field to given value.

### HasServerNames

`func (o *VpnUser) HasServerNames() bool`

HasServerNames returns a boolean if a field has been set.

### GetServers

`func (o *VpnUser) GetServers() []string`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *VpnUser) GetServersOk() (*[]string, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *VpnUser) SetServers(v []string)`

SetServers sets Servers field to given value.

### HasServers

`func (o *VpnUser) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetUserRemoteSubnets

`func (o *VpnUser) GetUserRemoteSubnets() []IPSubnetsVO`

GetUserRemoteSubnets returns the UserRemoteSubnets field if non-nil, zero value otherwise.

### GetUserRemoteSubnetsOk

`func (o *VpnUser) GetUserRemoteSubnetsOk() (*[]IPSubnetsVO, bool)`

GetUserRemoteSubnetsOk returns a tuple with the UserRemoteSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRemoteSubnets

`func (o *VpnUser) SetUserRemoteSubnets(v []IPSubnetsVO)`

SetUserRemoteSubnets sets UserRemoteSubnets field to given value.

### HasUserRemoteSubnets

`func (o *VpnUser) HasUserRemoteSubnets() bool`

HasUserRemoteSubnets returns a boolean if a field has been set.

### GetUsername

`func (o *VpnUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VpnUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VpnUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *VpnUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


