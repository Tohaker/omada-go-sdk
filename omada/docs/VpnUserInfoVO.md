# VpnUserInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** |  | [optional] 
**ClientMode** | Pointer to **int32** | Client mode should be a value as follows: 0: Client-To-Site, 1: Site-To-Site. | [optional] 
**FeatureDescription** | Pointer to [**[]FeatureInfoVO**](FeatureInfoVO.md) | Gateway Feature Description. | [optional] 
**GroupId** | Pointer to **string** | Group ID of the SSL VPN user. User group can be created using &#39;Create SSL VPN user group&#39; interface, and User Group ID can be obtained from &#39;Get user group list for SSL VPN server&#39; interface. | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | ID of the VPN user. | [optional] 
**LocalIp** | Pointer to **string** | Local IP of the VPN user. | [optional] 
**MaxConnections** | Pointer to **int32** | Max connections should be within the range of 1–100. | [optional] 
**Password** | Pointer to **string** | Password of the VPN user. | [optional] 
**Protocol** | Pointer to **int32** | Protocol should be a value as follows: 0: L2TP or PPTP; 1: openVPN; 2: SSL VPN. | [optional] 
**ServerNames** | Pointer to **string** | Server names of the VPN user. | [optional] 
**Servers** | Pointer to **[]string** | Servers of the VPN user. Server can be created using &#39;Create client-to-site VPN server&#39; interface, and server ID can be obtained from &#39;Get client-to-site VPN server list&#39; interface. | [optional] 
**Status** | Pointer to **bool** | Status of the SSL VPN user. | [optional] 
**UserRemoteSubnets** | Pointer to [**[]IPSubnetsVO**](IPSubnetsVO.md) | User remote subnets of the VPN user. | [optional] 
**Username** | Pointer to **string** | Username of the VPN user. | [optional] 
**Validity** | Pointer to **string** | Validity of the SSL VPN user. The format is Month/Day/Year, for example 08/20/2022. | [optional] 

## Methods

### NewVpnUserInfoVO

`func NewVpnUserInfoVO() *VpnUserInfoVO`

NewVpnUserInfoVO instantiates a new VpnUserInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnUserInfoVOWithDefaults

`func NewVpnUserInfoVOWithDefaults() *VpnUserInfoVO`

NewVpnUserInfoVOWithDefaults instantiates a new VpnUserInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *VpnUserInfoVO) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *VpnUserInfoVO) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *VpnUserInfoVO) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *VpnUserInfoVO) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetClientMode

`func (o *VpnUserInfoVO) GetClientMode() int32`

GetClientMode returns the ClientMode field if non-nil, zero value otherwise.

### GetClientModeOk

`func (o *VpnUserInfoVO) GetClientModeOk() (*int32, bool)`

GetClientModeOk returns a tuple with the ClientMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMode

`func (o *VpnUserInfoVO) SetClientMode(v int32)`

SetClientMode sets ClientMode field to given value.

### HasClientMode

`func (o *VpnUserInfoVO) HasClientMode() bool`

HasClientMode returns a boolean if a field has been set.

### GetFeatureDescription

`func (o *VpnUserInfoVO) GetFeatureDescription() []FeatureInfoVO`

GetFeatureDescription returns the FeatureDescription field if non-nil, zero value otherwise.

### GetFeatureDescriptionOk

`func (o *VpnUserInfoVO) GetFeatureDescriptionOk() (*[]FeatureInfoVO, bool)`

GetFeatureDescriptionOk returns a tuple with the FeatureDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureDescription

`func (o *VpnUserInfoVO) SetFeatureDescription(v []FeatureInfoVO)`

SetFeatureDescription sets FeatureDescription field to given value.

### HasFeatureDescription

`func (o *VpnUserInfoVO) HasFeatureDescription() bool`

HasFeatureDescription returns a boolean if a field has been set.

### GetGroupId

`func (o *VpnUserInfoVO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *VpnUserInfoVO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *VpnUserInfoVO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *VpnUserInfoVO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *VpnUserInfoVO) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *VpnUserInfoVO) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *VpnUserInfoVO) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *VpnUserInfoVO) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetId

`func (o *VpnUserInfoVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpnUserInfoVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpnUserInfoVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VpnUserInfoVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocalIp

`func (o *VpnUserInfoVO) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *VpnUserInfoVO) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *VpnUserInfoVO) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.

### HasLocalIp

`func (o *VpnUserInfoVO) HasLocalIp() bool`

HasLocalIp returns a boolean if a field has been set.

### GetMaxConnections

`func (o *VpnUserInfoVO) GetMaxConnections() int32`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *VpnUserInfoVO) GetMaxConnectionsOk() (*int32, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *VpnUserInfoVO) SetMaxConnections(v int32)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *VpnUserInfoVO) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetPassword

`func (o *VpnUserInfoVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VpnUserInfoVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VpnUserInfoVO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VpnUserInfoVO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProtocol

`func (o *VpnUserInfoVO) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *VpnUserInfoVO) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *VpnUserInfoVO) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *VpnUserInfoVO) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetServerNames

`func (o *VpnUserInfoVO) GetServerNames() string`

GetServerNames returns the ServerNames field if non-nil, zero value otherwise.

### GetServerNamesOk

`func (o *VpnUserInfoVO) GetServerNamesOk() (*string, bool)`

GetServerNamesOk returns a tuple with the ServerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNames

`func (o *VpnUserInfoVO) SetServerNames(v string)`

SetServerNames sets ServerNames field to given value.

### HasServerNames

`func (o *VpnUserInfoVO) HasServerNames() bool`

HasServerNames returns a boolean if a field has been set.

### GetServers

`func (o *VpnUserInfoVO) GetServers() []string`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *VpnUserInfoVO) GetServersOk() (*[]string, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *VpnUserInfoVO) SetServers(v []string)`

SetServers sets Servers field to given value.

### HasServers

`func (o *VpnUserInfoVO) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetStatus

`func (o *VpnUserInfoVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnUserInfoVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnUserInfoVO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VpnUserInfoVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserRemoteSubnets

`func (o *VpnUserInfoVO) GetUserRemoteSubnets() []IPSubnetsVO`

GetUserRemoteSubnets returns the UserRemoteSubnets field if non-nil, zero value otherwise.

### GetUserRemoteSubnetsOk

`func (o *VpnUserInfoVO) GetUserRemoteSubnetsOk() (*[]IPSubnetsVO, bool)`

GetUserRemoteSubnetsOk returns a tuple with the UserRemoteSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRemoteSubnets

`func (o *VpnUserInfoVO) SetUserRemoteSubnets(v []IPSubnetsVO)`

SetUserRemoteSubnets sets UserRemoteSubnets field to given value.

### HasUserRemoteSubnets

`func (o *VpnUserInfoVO) HasUserRemoteSubnets() bool`

HasUserRemoteSubnets returns a boolean if a field has been set.

### GetUsername

`func (o *VpnUserInfoVO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VpnUserInfoVO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VpnUserInfoVO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *VpnUserInfoVO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetValidity

`func (o *VpnUserInfoVO) GetValidity() string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *VpnUserInfoVO) GetValidityOk() (*string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *VpnUserInfoVO) SetValidity(v string)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *VpnUserInfoVO) HasValidity() bool`

HasValidity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


