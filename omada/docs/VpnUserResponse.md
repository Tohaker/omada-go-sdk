# VpnUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | ID of current Site. | [optional] 
**ClientMode** | Pointer to **int32** | Client mode should be a value as follows: 0: Client-To-Site, 1: Site-To-Site. | [optional] 
**ExistLocalIp** | Pointer to **bool** | Whether the Local IP is configured. | [optional] 
**ExistProtocol** | Pointer to **bool** | Whether an OpenVPN user exists. | [optional] 
**GroupId** | Pointer to **string** | Group ID of the SSL VPN user. User group can be created using &#39;Create SSL VPN user group&#39; interface, and User Group ID can be obtained from &#39;Get user group list for SSL VPN server&#39; interface. | [optional] 
**Id** | Pointer to **string** | ID of the VPN user. | [optional] 
**LocalIp** | Pointer to **string** | Local IP of the VPN user. | [optional] 
**MaxConnections** | Pointer to **int32** | Max connections should be within the range of 1–100. | [optional] 
**OmadacId** | Pointer to **string** | ID of current Controller. | [optional] 
**Password** | Pointer to **string** | Password of the VPN user. | [optional] 
**Protocol** | Pointer to **int32** | Protocol should be a value as follows: 0: L2TP or PPTP; 1: openVPN. | [optional] 
**ServerNames** | Pointer to **string** | Server names of the VPN user. | [optional] 
**Servers** | Pointer to **[]string** | Servers of the VPN user. Server can be created using &#39;Create client-to-site VPN server&#39; interface, and server ID can be obtained from &#39;Get client-to-site VPN server list&#39; interface. | [optional] 
**Status** | Pointer to **bool** | Status of the SSL VPN user. | [optional] 
**UserRemoteSubnets** | Pointer to [**[]IPSubnetsVO**](IPSubnetsVO.md) | User remote subnets of the VPN user. | [optional] 
**Username** | Pointer to **string** | Username of the VPN user. | [optional] 
**Validity** | Pointer to **string** | Validity of the SSL VPN user. The format is Month/Day/Year, for example 08/20/2022. | [optional] 

## Methods

### NewVpnUserResponse

`func NewVpnUserResponse() *VpnUserResponse`

NewVpnUserResponse instantiates a new VpnUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnUserResponseWithDefaults

`func NewVpnUserResponseWithDefaults() *VpnUserResponse`

NewVpnUserResponseWithDefaults instantiates a new VpnUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *VpnUserResponse) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *VpnUserResponse) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *VpnUserResponse) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *VpnUserResponse) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetClientMode

`func (o *VpnUserResponse) GetClientMode() int32`

GetClientMode returns the ClientMode field if non-nil, zero value otherwise.

### GetClientModeOk

`func (o *VpnUserResponse) GetClientModeOk() (*int32, bool)`

GetClientModeOk returns a tuple with the ClientMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMode

`func (o *VpnUserResponse) SetClientMode(v int32)`

SetClientMode sets ClientMode field to given value.

### HasClientMode

`func (o *VpnUserResponse) HasClientMode() bool`

HasClientMode returns a boolean if a field has been set.

### GetExistLocalIp

`func (o *VpnUserResponse) GetExistLocalIp() bool`

GetExistLocalIp returns the ExistLocalIp field if non-nil, zero value otherwise.

### GetExistLocalIpOk

`func (o *VpnUserResponse) GetExistLocalIpOk() (*bool, bool)`

GetExistLocalIpOk returns a tuple with the ExistLocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistLocalIp

`func (o *VpnUserResponse) SetExistLocalIp(v bool)`

SetExistLocalIp sets ExistLocalIp field to given value.

### HasExistLocalIp

`func (o *VpnUserResponse) HasExistLocalIp() bool`

HasExistLocalIp returns a boolean if a field has been set.

### GetExistProtocol

`func (o *VpnUserResponse) GetExistProtocol() bool`

GetExistProtocol returns the ExistProtocol field if non-nil, zero value otherwise.

### GetExistProtocolOk

`func (o *VpnUserResponse) GetExistProtocolOk() (*bool, bool)`

GetExistProtocolOk returns a tuple with the ExistProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistProtocol

`func (o *VpnUserResponse) SetExistProtocol(v bool)`

SetExistProtocol sets ExistProtocol field to given value.

### HasExistProtocol

`func (o *VpnUserResponse) HasExistProtocol() bool`

HasExistProtocol returns a boolean if a field has been set.

### GetGroupId

`func (o *VpnUserResponse) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *VpnUserResponse) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *VpnUserResponse) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *VpnUserResponse) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *VpnUserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpnUserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpnUserResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VpnUserResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocalIp

`func (o *VpnUserResponse) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *VpnUserResponse) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *VpnUserResponse) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.

### HasLocalIp

`func (o *VpnUserResponse) HasLocalIp() bool`

HasLocalIp returns a boolean if a field has been set.

### GetMaxConnections

`func (o *VpnUserResponse) GetMaxConnections() int32`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *VpnUserResponse) GetMaxConnectionsOk() (*int32, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *VpnUserResponse) SetMaxConnections(v int32)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *VpnUserResponse) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetOmadacId

`func (o *VpnUserResponse) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *VpnUserResponse) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *VpnUserResponse) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *VpnUserResponse) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetPassword

`func (o *VpnUserResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VpnUserResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VpnUserResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VpnUserResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProtocol

`func (o *VpnUserResponse) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *VpnUserResponse) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *VpnUserResponse) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *VpnUserResponse) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetServerNames

`func (o *VpnUserResponse) GetServerNames() string`

GetServerNames returns the ServerNames field if non-nil, zero value otherwise.

### GetServerNamesOk

`func (o *VpnUserResponse) GetServerNamesOk() (*string, bool)`

GetServerNamesOk returns a tuple with the ServerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNames

`func (o *VpnUserResponse) SetServerNames(v string)`

SetServerNames sets ServerNames field to given value.

### HasServerNames

`func (o *VpnUserResponse) HasServerNames() bool`

HasServerNames returns a boolean if a field has been set.

### GetServers

`func (o *VpnUserResponse) GetServers() []string`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *VpnUserResponse) GetServersOk() (*[]string, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *VpnUserResponse) SetServers(v []string)`

SetServers sets Servers field to given value.

### HasServers

`func (o *VpnUserResponse) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetStatus

`func (o *VpnUserResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnUserResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnUserResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VpnUserResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserRemoteSubnets

`func (o *VpnUserResponse) GetUserRemoteSubnets() []IPSubnetsVO`

GetUserRemoteSubnets returns the UserRemoteSubnets field if non-nil, zero value otherwise.

### GetUserRemoteSubnetsOk

`func (o *VpnUserResponse) GetUserRemoteSubnetsOk() (*[]IPSubnetsVO, bool)`

GetUserRemoteSubnetsOk returns a tuple with the UserRemoteSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRemoteSubnets

`func (o *VpnUserResponse) SetUserRemoteSubnets(v []IPSubnetsVO)`

SetUserRemoteSubnets sets UserRemoteSubnets field to given value.

### HasUserRemoteSubnets

`func (o *VpnUserResponse) HasUserRemoteSubnets() bool`

HasUserRemoteSubnets returns a boolean if a field has been set.

### GetUsername

`func (o *VpnUserResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VpnUserResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VpnUserResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *VpnUserResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetValidity

`func (o *VpnUserResponse) GetValidity() string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *VpnUserResponse) GetValidityOk() (*string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *VpnUserResponse) SetValidity(v string)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *VpnUserResponse) HasValidity() bool`

HasValidity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


