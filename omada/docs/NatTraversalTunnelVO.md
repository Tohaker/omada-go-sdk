# NatTraversalTunnelVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppType** | Pointer to **int32** | App type of the remote access tunnel for the local target device. 21:HTTP, 22:HTTPS, 31:SSH, 32:TELNET. | [optional] 
**AutoLogin** | Pointer to **bool** | If the remote access tunnel enabled auto login. | [optional] 
**ClientDeviceType** | Pointer to **string** | Client&#39;s type of the remote access tunnel for the client. | [optional] 
**ClientManager** | Pointer to **bool** | Client&#39;s manager of the remote access tunnel for the client. | [optional] 
**ClientModel** | Pointer to **string** | Client&#39;s model of the remote access tunnel for the client. | [optional] 
**ClientName** | Pointer to **string** | Client&#39;s name of the remote access tunnel for the client. | [optional] 
**ClosureTime** | Pointer to **int64** | Closure time of the remote access tunnel. | [optional] 
**CustomTunnel** | Pointer to **bool** | If the remote access tunnel is a custom tunnel. | [optional] 
**Duration** | Pointer to **int32** | Valid duration time of the remote access tunnel, 1-24 hours. | [optional] 
**Id** | Pointer to **string** | ID of the remote access tunnel. | [optional] 
**JumpMac** | Pointer to **string** | Mac of the jumper device. | [optional] 
**LocalAddress** | Pointer to **string** | IP address of the local target device. | [optional] 
**LocalMac** | Pointer to **string** | Mac of the local target device. | [optional] 
**LocalPort** | **int32** | Port of the local target device&#39;s service. | 
**Name** | Pointer to **string** | Name of the remote access tunnel. | [optional] 
**OpenStatus** | Pointer to **bool** | Open Status of the remote access tunnel. | [optional] 
**Status** | Pointer to **int32** | Status of the remote access tunnel. 0: Disconnected, 1: Connected, 2: Opening, -1: Heartbeat Missed, -2: Expired. | [optional] 
**SupportSshTelnetTunnel** | Pointer to **bool** | If the remote access tunnel supports SSH and telnet connections. | [optional] 
**SupportWebAutoLogin** | Pointer to **bool** | If the remote access tunnel supports web auto login. | [optional] 
**TunnelEntryType** | Pointer to **int32** | Entry of the remote access tunnel.1: custom tunnel 2: device tunnel 3:device detail 4:client detail | [optional] 
**TunnelId** | Pointer to **int32** | Tunnel ID of the remote access tunnel. | [optional] 
**Type** | Pointer to **string** | Type of the remote access tunnel. | [optional] 

## Methods

### NewNatTraversalTunnelVO

`func NewNatTraversalTunnelVO(localPort int32, ) *NatTraversalTunnelVO`

NewNatTraversalTunnelVO instantiates a new NatTraversalTunnelVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatTraversalTunnelVOWithDefaults

`func NewNatTraversalTunnelVOWithDefaults() *NatTraversalTunnelVO`

NewNatTraversalTunnelVOWithDefaults instantiates a new NatTraversalTunnelVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppType

`func (o *NatTraversalTunnelVO) GetAppType() int32`

GetAppType returns the AppType field if non-nil, zero value otherwise.

### GetAppTypeOk

`func (o *NatTraversalTunnelVO) GetAppTypeOk() (*int32, bool)`

GetAppTypeOk returns a tuple with the AppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppType

`func (o *NatTraversalTunnelVO) SetAppType(v int32)`

SetAppType sets AppType field to given value.

### HasAppType

`func (o *NatTraversalTunnelVO) HasAppType() bool`

HasAppType returns a boolean if a field has been set.

### GetAutoLogin

`func (o *NatTraversalTunnelVO) GetAutoLogin() bool`

GetAutoLogin returns the AutoLogin field if non-nil, zero value otherwise.

### GetAutoLoginOk

`func (o *NatTraversalTunnelVO) GetAutoLoginOk() (*bool, bool)`

GetAutoLoginOk returns a tuple with the AutoLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLogin

`func (o *NatTraversalTunnelVO) SetAutoLogin(v bool)`

SetAutoLogin sets AutoLogin field to given value.

### HasAutoLogin

`func (o *NatTraversalTunnelVO) HasAutoLogin() bool`

HasAutoLogin returns a boolean if a field has been set.

### GetClientDeviceType

`func (o *NatTraversalTunnelVO) GetClientDeviceType() string`

GetClientDeviceType returns the ClientDeviceType field if non-nil, zero value otherwise.

### GetClientDeviceTypeOk

`func (o *NatTraversalTunnelVO) GetClientDeviceTypeOk() (*string, bool)`

GetClientDeviceTypeOk returns a tuple with the ClientDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDeviceType

`func (o *NatTraversalTunnelVO) SetClientDeviceType(v string)`

SetClientDeviceType sets ClientDeviceType field to given value.

### HasClientDeviceType

`func (o *NatTraversalTunnelVO) HasClientDeviceType() bool`

HasClientDeviceType returns a boolean if a field has been set.

### GetClientManager

`func (o *NatTraversalTunnelVO) GetClientManager() bool`

GetClientManager returns the ClientManager field if non-nil, zero value otherwise.

### GetClientManagerOk

`func (o *NatTraversalTunnelVO) GetClientManagerOk() (*bool, bool)`

GetClientManagerOk returns a tuple with the ClientManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientManager

`func (o *NatTraversalTunnelVO) SetClientManager(v bool)`

SetClientManager sets ClientManager field to given value.

### HasClientManager

`func (o *NatTraversalTunnelVO) HasClientManager() bool`

HasClientManager returns a boolean if a field has been set.

### GetClientModel

`func (o *NatTraversalTunnelVO) GetClientModel() string`

GetClientModel returns the ClientModel field if non-nil, zero value otherwise.

### GetClientModelOk

`func (o *NatTraversalTunnelVO) GetClientModelOk() (*string, bool)`

GetClientModelOk returns a tuple with the ClientModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientModel

`func (o *NatTraversalTunnelVO) SetClientModel(v string)`

SetClientModel sets ClientModel field to given value.

### HasClientModel

`func (o *NatTraversalTunnelVO) HasClientModel() bool`

HasClientModel returns a boolean if a field has been set.

### GetClientName

`func (o *NatTraversalTunnelVO) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *NatTraversalTunnelVO) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *NatTraversalTunnelVO) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *NatTraversalTunnelVO) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetClosureTime

`func (o *NatTraversalTunnelVO) GetClosureTime() int64`

GetClosureTime returns the ClosureTime field if non-nil, zero value otherwise.

### GetClosureTimeOk

`func (o *NatTraversalTunnelVO) GetClosureTimeOk() (*int64, bool)`

GetClosureTimeOk returns a tuple with the ClosureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosureTime

`func (o *NatTraversalTunnelVO) SetClosureTime(v int64)`

SetClosureTime sets ClosureTime field to given value.

### HasClosureTime

`func (o *NatTraversalTunnelVO) HasClosureTime() bool`

HasClosureTime returns a boolean if a field has been set.

### GetCustomTunnel

`func (o *NatTraversalTunnelVO) GetCustomTunnel() bool`

GetCustomTunnel returns the CustomTunnel field if non-nil, zero value otherwise.

### GetCustomTunnelOk

`func (o *NatTraversalTunnelVO) GetCustomTunnelOk() (*bool, bool)`

GetCustomTunnelOk returns a tuple with the CustomTunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTunnel

`func (o *NatTraversalTunnelVO) SetCustomTunnel(v bool)`

SetCustomTunnel sets CustomTunnel field to given value.

### HasCustomTunnel

`func (o *NatTraversalTunnelVO) HasCustomTunnel() bool`

HasCustomTunnel returns a boolean if a field has been set.

### GetDuration

`func (o *NatTraversalTunnelVO) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *NatTraversalTunnelVO) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *NatTraversalTunnelVO) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *NatTraversalTunnelVO) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetId

`func (o *NatTraversalTunnelVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NatTraversalTunnelVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NatTraversalTunnelVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NatTraversalTunnelVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJumpMac

`func (o *NatTraversalTunnelVO) GetJumpMac() string`

GetJumpMac returns the JumpMac field if non-nil, zero value otherwise.

### GetJumpMacOk

`func (o *NatTraversalTunnelVO) GetJumpMacOk() (*string, bool)`

GetJumpMacOk returns a tuple with the JumpMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumpMac

`func (o *NatTraversalTunnelVO) SetJumpMac(v string)`

SetJumpMac sets JumpMac field to given value.

### HasJumpMac

`func (o *NatTraversalTunnelVO) HasJumpMac() bool`

HasJumpMac returns a boolean if a field has been set.

### GetLocalAddress

`func (o *NatTraversalTunnelVO) GetLocalAddress() string`

GetLocalAddress returns the LocalAddress field if non-nil, zero value otherwise.

### GetLocalAddressOk

`func (o *NatTraversalTunnelVO) GetLocalAddressOk() (*string, bool)`

GetLocalAddressOk returns a tuple with the LocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddress

`func (o *NatTraversalTunnelVO) SetLocalAddress(v string)`

SetLocalAddress sets LocalAddress field to given value.

### HasLocalAddress

`func (o *NatTraversalTunnelVO) HasLocalAddress() bool`

HasLocalAddress returns a boolean if a field has been set.

### GetLocalMac

`func (o *NatTraversalTunnelVO) GetLocalMac() string`

GetLocalMac returns the LocalMac field if non-nil, zero value otherwise.

### GetLocalMacOk

`func (o *NatTraversalTunnelVO) GetLocalMacOk() (*string, bool)`

GetLocalMacOk returns a tuple with the LocalMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalMac

`func (o *NatTraversalTunnelVO) SetLocalMac(v string)`

SetLocalMac sets LocalMac field to given value.

### HasLocalMac

`func (o *NatTraversalTunnelVO) HasLocalMac() bool`

HasLocalMac returns a boolean if a field has been set.

### GetLocalPort

`func (o *NatTraversalTunnelVO) GetLocalPort() int32`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *NatTraversalTunnelVO) GetLocalPortOk() (*int32, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *NatTraversalTunnelVO) SetLocalPort(v int32)`

SetLocalPort sets LocalPort field to given value.


### GetName

`func (o *NatTraversalTunnelVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NatTraversalTunnelVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NatTraversalTunnelVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NatTraversalTunnelVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpenStatus

`func (o *NatTraversalTunnelVO) GetOpenStatus() bool`

GetOpenStatus returns the OpenStatus field if non-nil, zero value otherwise.

### GetOpenStatusOk

`func (o *NatTraversalTunnelVO) GetOpenStatusOk() (*bool, bool)`

GetOpenStatusOk returns a tuple with the OpenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStatus

`func (o *NatTraversalTunnelVO) SetOpenStatus(v bool)`

SetOpenStatus sets OpenStatus field to given value.

### HasOpenStatus

`func (o *NatTraversalTunnelVO) HasOpenStatus() bool`

HasOpenStatus returns a boolean if a field has been set.

### GetStatus

`func (o *NatTraversalTunnelVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NatTraversalTunnelVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NatTraversalTunnelVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NatTraversalTunnelVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportSshTelnetTunnel

`func (o *NatTraversalTunnelVO) GetSupportSshTelnetTunnel() bool`

GetSupportSshTelnetTunnel returns the SupportSshTelnetTunnel field if non-nil, zero value otherwise.

### GetSupportSshTelnetTunnelOk

`func (o *NatTraversalTunnelVO) GetSupportSshTelnetTunnelOk() (*bool, bool)`

GetSupportSshTelnetTunnelOk returns a tuple with the SupportSshTelnetTunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSshTelnetTunnel

`func (o *NatTraversalTunnelVO) SetSupportSshTelnetTunnel(v bool)`

SetSupportSshTelnetTunnel sets SupportSshTelnetTunnel field to given value.

### HasSupportSshTelnetTunnel

`func (o *NatTraversalTunnelVO) HasSupportSshTelnetTunnel() bool`

HasSupportSshTelnetTunnel returns a boolean if a field has been set.

### GetSupportWebAutoLogin

`func (o *NatTraversalTunnelVO) GetSupportWebAutoLogin() bool`

GetSupportWebAutoLogin returns the SupportWebAutoLogin field if non-nil, zero value otherwise.

### GetSupportWebAutoLoginOk

`func (o *NatTraversalTunnelVO) GetSupportWebAutoLoginOk() (*bool, bool)`

GetSupportWebAutoLoginOk returns a tuple with the SupportWebAutoLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportWebAutoLogin

`func (o *NatTraversalTunnelVO) SetSupportWebAutoLogin(v bool)`

SetSupportWebAutoLogin sets SupportWebAutoLogin field to given value.

### HasSupportWebAutoLogin

`func (o *NatTraversalTunnelVO) HasSupportWebAutoLogin() bool`

HasSupportWebAutoLogin returns a boolean if a field has been set.

### GetTunnelEntryType

`func (o *NatTraversalTunnelVO) GetTunnelEntryType() int32`

GetTunnelEntryType returns the TunnelEntryType field if non-nil, zero value otherwise.

### GetTunnelEntryTypeOk

`func (o *NatTraversalTunnelVO) GetTunnelEntryTypeOk() (*int32, bool)`

GetTunnelEntryTypeOk returns a tuple with the TunnelEntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelEntryType

`func (o *NatTraversalTunnelVO) SetTunnelEntryType(v int32)`

SetTunnelEntryType sets TunnelEntryType field to given value.

### HasTunnelEntryType

`func (o *NatTraversalTunnelVO) HasTunnelEntryType() bool`

HasTunnelEntryType returns a boolean if a field has been set.

### GetTunnelId

`func (o *NatTraversalTunnelVO) GetTunnelId() int32`

GetTunnelId returns the TunnelId field if non-nil, zero value otherwise.

### GetTunnelIdOk

`func (o *NatTraversalTunnelVO) GetTunnelIdOk() (*int32, bool)`

GetTunnelIdOk returns a tuple with the TunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelId

`func (o *NatTraversalTunnelVO) SetTunnelId(v int32)`

SetTunnelId sets TunnelId field to given value.

### HasTunnelId

`func (o *NatTraversalTunnelVO) HasTunnelId() bool`

HasTunnelId returns a boolean if a field has been set.

### GetType

`func (o *NatTraversalTunnelVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NatTraversalTunnelVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NatTraversalTunnelVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NatTraversalTunnelVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


