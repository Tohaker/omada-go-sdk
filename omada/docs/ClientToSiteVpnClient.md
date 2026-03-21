# ClientToSiteVpnClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientPassword** | Pointer to **string** | Client password should contain 1 to 64 characters. | [optional] 
**ClientUserName** | Pointer to **string** | Client username should contain 1 to 64 characters. | [optional] 
**ClientVpnType** | **int32** | Client Vpn type should be a value as follows: 0: L2TP; 1: PPTP; 2: IPSec; 3: OpenVPN. | 
**CustomNetwork** | Pointer to [**[]IPSubnetsVO**](IPSubnetsVO.md) | Custom networks of the VPN. | [optional] 
**Encryption** | Pointer to **int32** | Encryption should be a value as follows: 0: Encrypted; 1: Unencrypted. | [optional] 
**Id** | Pointer to **string** | ID of the VPN. | [optional] 
**Mode** | Pointer to **int32** | Mode(only for server OpenVPN) should be a value as follows: 0: certification; 1: certification+account. | [optional] 
**Name** | **string** | Name should contain 1 to 63 characters. | 
**NetworkList** | Pointer to **[]string** | Network list of the VPN. Network can be created using &#39;Create LAN network&#39; interface, and network ID can be obtained from &#39;Get LAN network list&#39; interface. | [optional] 
**NetworkType** | Pointer to **int32** | Network type should be a value as follows: 0: network list; 1: custom networks. | [optional] 
**OpenVpnMode** | Pointer to **int32** | OpenVPN mode should be a value as follows: 0: certification; 1: certification+account. | [optional] 
**OpenVpnTunnelMode** | Pointer to **int32** | OpenVPN tunnel mode should be a value as follows: 0: split; 1: full. | [optional] 
**PreSharedKey** | Pointer to **string** | Pre-shared key of the VPN. | [optional] 
**RemoteIp** | Pointer to **string** | Remote IP of the VPN. Get whether supports domain from interface &#39;Get client-to-site VPN client list&#39;. | [optional] 
**RemoteSite** | Pointer to **string** | Remote site of the VPN. | [optional] 
**RemoteSubnet** | Pointer to [**[]IPSubnetsVO**](IPSubnetsVO.md) | Remote subnet of the VPN. | [optional] 
**ServicePort** | Pointer to **int32** | Service port should be within the range of 1–65535. | [optional] 
**ServiceType** | Pointer to **int32** | Service type should be a value as follows: 0: UDP; 1: TCP. | [optional] 
**Status** | Pointer to **bool** | Status of the VPN. | [optional] 
**VpnConfiguration** | Pointer to [**VpnCertificateOpenApiVO**](VpnCertificateOpenApiVO.md) |  | [optional] 
**Wan** | **[]string** | WAN list of the VPN. WAN port ID can be obtained from &#39;Get internet basic info&#39; interface. | 
**WorkingMode** | Pointer to **int32** | Working mode should be a value as follows: 0: NAT; 1: Routing. | [optional] 

## Methods

### NewClientToSiteVpnClient

`func NewClientToSiteVpnClient(clientVpnType int32, name string, wan []string, ) *ClientToSiteVpnClient`

NewClientToSiteVpnClient instantiates a new ClientToSiteVpnClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientToSiteVpnClientWithDefaults

`func NewClientToSiteVpnClientWithDefaults() *ClientToSiteVpnClient`

NewClientToSiteVpnClientWithDefaults instantiates a new ClientToSiteVpnClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientPassword

`func (o *ClientToSiteVpnClient) GetClientPassword() string`

GetClientPassword returns the ClientPassword field if non-nil, zero value otherwise.

### GetClientPasswordOk

`func (o *ClientToSiteVpnClient) GetClientPasswordOk() (*string, bool)`

GetClientPasswordOk returns a tuple with the ClientPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPassword

`func (o *ClientToSiteVpnClient) SetClientPassword(v string)`

SetClientPassword sets ClientPassword field to given value.

### HasClientPassword

`func (o *ClientToSiteVpnClient) HasClientPassword() bool`

HasClientPassword returns a boolean if a field has been set.

### GetClientUserName

`func (o *ClientToSiteVpnClient) GetClientUserName() string`

GetClientUserName returns the ClientUserName field if non-nil, zero value otherwise.

### GetClientUserNameOk

`func (o *ClientToSiteVpnClient) GetClientUserNameOk() (*string, bool)`

GetClientUserNameOk returns a tuple with the ClientUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUserName

`func (o *ClientToSiteVpnClient) SetClientUserName(v string)`

SetClientUserName sets ClientUserName field to given value.

### HasClientUserName

`func (o *ClientToSiteVpnClient) HasClientUserName() bool`

HasClientUserName returns a boolean if a field has been set.

### GetClientVpnType

`func (o *ClientToSiteVpnClient) GetClientVpnType() int32`

GetClientVpnType returns the ClientVpnType field if non-nil, zero value otherwise.

### GetClientVpnTypeOk

`func (o *ClientToSiteVpnClient) GetClientVpnTypeOk() (*int32, bool)`

GetClientVpnTypeOk returns a tuple with the ClientVpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVpnType

`func (o *ClientToSiteVpnClient) SetClientVpnType(v int32)`

SetClientVpnType sets ClientVpnType field to given value.


### GetCustomNetwork

`func (o *ClientToSiteVpnClient) GetCustomNetwork() []IPSubnetsVO`

GetCustomNetwork returns the CustomNetwork field if non-nil, zero value otherwise.

### GetCustomNetworkOk

`func (o *ClientToSiteVpnClient) GetCustomNetworkOk() (*[]IPSubnetsVO, bool)`

GetCustomNetworkOk returns a tuple with the CustomNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetwork

`func (o *ClientToSiteVpnClient) SetCustomNetwork(v []IPSubnetsVO)`

SetCustomNetwork sets CustomNetwork field to given value.

### HasCustomNetwork

`func (o *ClientToSiteVpnClient) HasCustomNetwork() bool`

HasCustomNetwork returns a boolean if a field has been set.

### GetEncryption

`func (o *ClientToSiteVpnClient) GetEncryption() int32`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *ClientToSiteVpnClient) GetEncryptionOk() (*int32, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *ClientToSiteVpnClient) SetEncryption(v int32)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *ClientToSiteVpnClient) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetId

`func (o *ClientToSiteVpnClient) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientToSiteVpnClient) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientToSiteVpnClient) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientToSiteVpnClient) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMode

`func (o *ClientToSiteVpnClient) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ClientToSiteVpnClient) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ClientToSiteVpnClient) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ClientToSiteVpnClient) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *ClientToSiteVpnClient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientToSiteVpnClient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientToSiteVpnClient) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkList

`func (o *ClientToSiteVpnClient) GetNetworkList() []string`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *ClientToSiteVpnClient) GetNetworkListOk() (*[]string, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *ClientToSiteVpnClient) SetNetworkList(v []string)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *ClientToSiteVpnClient) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.

### GetNetworkType

`func (o *ClientToSiteVpnClient) GetNetworkType() int32`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *ClientToSiteVpnClient) GetNetworkTypeOk() (*int32, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *ClientToSiteVpnClient) SetNetworkType(v int32)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *ClientToSiteVpnClient) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetOpenVpnMode

`func (o *ClientToSiteVpnClient) GetOpenVpnMode() int32`

GetOpenVpnMode returns the OpenVpnMode field if non-nil, zero value otherwise.

### GetOpenVpnModeOk

`func (o *ClientToSiteVpnClient) GetOpenVpnModeOk() (*int32, bool)`

GetOpenVpnModeOk returns a tuple with the OpenVpnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenVpnMode

`func (o *ClientToSiteVpnClient) SetOpenVpnMode(v int32)`

SetOpenVpnMode sets OpenVpnMode field to given value.

### HasOpenVpnMode

`func (o *ClientToSiteVpnClient) HasOpenVpnMode() bool`

HasOpenVpnMode returns a boolean if a field has been set.

### GetOpenVpnTunnelMode

`func (o *ClientToSiteVpnClient) GetOpenVpnTunnelMode() int32`

GetOpenVpnTunnelMode returns the OpenVpnTunnelMode field if non-nil, zero value otherwise.

### GetOpenVpnTunnelModeOk

`func (o *ClientToSiteVpnClient) GetOpenVpnTunnelModeOk() (*int32, bool)`

GetOpenVpnTunnelModeOk returns a tuple with the OpenVpnTunnelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenVpnTunnelMode

`func (o *ClientToSiteVpnClient) SetOpenVpnTunnelMode(v int32)`

SetOpenVpnTunnelMode sets OpenVpnTunnelMode field to given value.

### HasOpenVpnTunnelMode

`func (o *ClientToSiteVpnClient) HasOpenVpnTunnelMode() bool`

HasOpenVpnTunnelMode returns a boolean if a field has been set.

### GetPreSharedKey

`func (o *ClientToSiteVpnClient) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *ClientToSiteVpnClient) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *ClientToSiteVpnClient) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.

### HasPreSharedKey

`func (o *ClientToSiteVpnClient) HasPreSharedKey() bool`

HasPreSharedKey returns a boolean if a field has been set.

### GetRemoteIp

`func (o *ClientToSiteVpnClient) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *ClientToSiteVpnClient) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *ClientToSiteVpnClient) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *ClientToSiteVpnClient) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetRemoteSite

`func (o *ClientToSiteVpnClient) GetRemoteSite() string`

GetRemoteSite returns the RemoteSite field if non-nil, zero value otherwise.

### GetRemoteSiteOk

`func (o *ClientToSiteVpnClient) GetRemoteSiteOk() (*string, bool)`

GetRemoteSiteOk returns a tuple with the RemoteSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSite

`func (o *ClientToSiteVpnClient) SetRemoteSite(v string)`

SetRemoteSite sets RemoteSite field to given value.

### HasRemoteSite

`func (o *ClientToSiteVpnClient) HasRemoteSite() bool`

HasRemoteSite returns a boolean if a field has been set.

### GetRemoteSubnet

`func (o *ClientToSiteVpnClient) GetRemoteSubnet() []IPSubnetsVO`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *ClientToSiteVpnClient) GetRemoteSubnetOk() (*[]IPSubnetsVO, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *ClientToSiteVpnClient) SetRemoteSubnet(v []IPSubnetsVO)`

SetRemoteSubnet sets RemoteSubnet field to given value.

### HasRemoteSubnet

`func (o *ClientToSiteVpnClient) HasRemoteSubnet() bool`

HasRemoteSubnet returns a boolean if a field has been set.

### GetServicePort

`func (o *ClientToSiteVpnClient) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *ClientToSiteVpnClient) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *ClientToSiteVpnClient) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *ClientToSiteVpnClient) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetServiceType

`func (o *ClientToSiteVpnClient) GetServiceType() int32`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ClientToSiteVpnClient) GetServiceTypeOk() (*int32, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ClientToSiteVpnClient) SetServiceType(v int32)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *ClientToSiteVpnClient) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetStatus

`func (o *ClientToSiteVpnClient) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClientToSiteVpnClient) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClientToSiteVpnClient) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClientToSiteVpnClient) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVpnConfiguration

`func (o *ClientToSiteVpnClient) GetVpnConfiguration() VpnCertificateOpenApiVO`

GetVpnConfiguration returns the VpnConfiguration field if non-nil, zero value otherwise.

### GetVpnConfigurationOk

`func (o *ClientToSiteVpnClient) GetVpnConfigurationOk() (*VpnCertificateOpenApiVO, bool)`

GetVpnConfigurationOk returns a tuple with the VpnConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConfiguration

`func (o *ClientToSiteVpnClient) SetVpnConfiguration(v VpnCertificateOpenApiVO)`

SetVpnConfiguration sets VpnConfiguration field to given value.

### HasVpnConfiguration

`func (o *ClientToSiteVpnClient) HasVpnConfiguration() bool`

HasVpnConfiguration returns a boolean if a field has been set.

### GetWan

`func (o *ClientToSiteVpnClient) GetWan() []string`

GetWan returns the Wan field if non-nil, zero value otherwise.

### GetWanOk

`func (o *ClientToSiteVpnClient) GetWanOk() (*[]string, bool)`

GetWanOk returns a tuple with the Wan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWan

`func (o *ClientToSiteVpnClient) SetWan(v []string)`

SetWan sets Wan field to given value.


### GetWorkingMode

`func (o *ClientToSiteVpnClient) GetWorkingMode() int32`

GetWorkingMode returns the WorkingMode field if non-nil, zero value otherwise.

### GetWorkingModeOk

`func (o *ClientToSiteVpnClient) GetWorkingModeOk() (*int32, bool)`

GetWorkingModeOk returns a tuple with the WorkingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingMode

`func (o *ClientToSiteVpnClient) SetWorkingMode(v int32)`

SetWorkingMode sets WorkingMode field to given value.

### HasWorkingMode

`func (o *ClientToSiteVpnClient) HasWorkingMode() bool`

HasWorkingMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


