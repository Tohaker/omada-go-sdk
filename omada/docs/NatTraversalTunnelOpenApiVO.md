# NatTraversalTunnelOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppType** | **string** | App type of the remote access tunnel for the local target device. HTTP, HTTPS, SSH, TELNET. | 
**ClientDeviceType** | Pointer to **string** | Client&#39;s type of the remote access tunnel for the client. | [optional] 
**ClientManager** | Pointer to **bool** | Client&#39;s manager of the remote access tunnel for the client. | [optional] 
**ClientModel** | Pointer to **string** | Client&#39;s model of the remote access tunnel for the client. | [optional] 
**ClientName** | Pointer to **string** | Client&#39;s name of the remote access tunnel for the client. | [optional] 
**Duration** | Pointer to **int32** | Valid duration time of the remote access tunnel, 1-24 hours. | [optional] 
**LocalAddress** | **string** | IP address of the local target device. | 
**LocalMac** | Pointer to **string** | Mac of the local target device. | [optional] 
**LocalPort** | **int32** | Port of the local target device&#39;s service. | 
**Name** | **string** | Name of the remote access tunnel. | 
**OpenStatus** | Pointer to **bool** | If open the remote access tunnel after create. | [optional] 
**TunnelEntryType** | Pointer to **int32** | Entry of the remote access tunnel.1: custom tunnel 2: device tunnel 3:device detail 4:client detail | [optional] 

## Methods

### NewNatTraversalTunnelOpenApiVO

`func NewNatTraversalTunnelOpenApiVO(appType string, localAddress string, localPort int32, name string, ) *NatTraversalTunnelOpenApiVO`

NewNatTraversalTunnelOpenApiVO instantiates a new NatTraversalTunnelOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatTraversalTunnelOpenApiVOWithDefaults

`func NewNatTraversalTunnelOpenApiVOWithDefaults() *NatTraversalTunnelOpenApiVO`

NewNatTraversalTunnelOpenApiVOWithDefaults instantiates a new NatTraversalTunnelOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppType

`func (o *NatTraversalTunnelOpenApiVO) GetAppType() string`

GetAppType returns the AppType field if non-nil, zero value otherwise.

### GetAppTypeOk

`func (o *NatTraversalTunnelOpenApiVO) GetAppTypeOk() (*string, bool)`

GetAppTypeOk returns a tuple with the AppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppType

`func (o *NatTraversalTunnelOpenApiVO) SetAppType(v string)`

SetAppType sets AppType field to given value.


### GetClientDeviceType

`func (o *NatTraversalTunnelOpenApiVO) GetClientDeviceType() string`

GetClientDeviceType returns the ClientDeviceType field if non-nil, zero value otherwise.

### GetClientDeviceTypeOk

`func (o *NatTraversalTunnelOpenApiVO) GetClientDeviceTypeOk() (*string, bool)`

GetClientDeviceTypeOk returns a tuple with the ClientDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDeviceType

`func (o *NatTraversalTunnelOpenApiVO) SetClientDeviceType(v string)`

SetClientDeviceType sets ClientDeviceType field to given value.

### HasClientDeviceType

`func (o *NatTraversalTunnelOpenApiVO) HasClientDeviceType() bool`

HasClientDeviceType returns a boolean if a field has been set.

### GetClientManager

`func (o *NatTraversalTunnelOpenApiVO) GetClientManager() bool`

GetClientManager returns the ClientManager field if non-nil, zero value otherwise.

### GetClientManagerOk

`func (o *NatTraversalTunnelOpenApiVO) GetClientManagerOk() (*bool, bool)`

GetClientManagerOk returns a tuple with the ClientManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientManager

`func (o *NatTraversalTunnelOpenApiVO) SetClientManager(v bool)`

SetClientManager sets ClientManager field to given value.

### HasClientManager

`func (o *NatTraversalTunnelOpenApiVO) HasClientManager() bool`

HasClientManager returns a boolean if a field has been set.

### GetClientModel

`func (o *NatTraversalTunnelOpenApiVO) GetClientModel() string`

GetClientModel returns the ClientModel field if non-nil, zero value otherwise.

### GetClientModelOk

`func (o *NatTraversalTunnelOpenApiVO) GetClientModelOk() (*string, bool)`

GetClientModelOk returns a tuple with the ClientModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientModel

`func (o *NatTraversalTunnelOpenApiVO) SetClientModel(v string)`

SetClientModel sets ClientModel field to given value.

### HasClientModel

`func (o *NatTraversalTunnelOpenApiVO) HasClientModel() bool`

HasClientModel returns a boolean if a field has been set.

### GetClientName

`func (o *NatTraversalTunnelOpenApiVO) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *NatTraversalTunnelOpenApiVO) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *NatTraversalTunnelOpenApiVO) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *NatTraversalTunnelOpenApiVO) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetDuration

`func (o *NatTraversalTunnelOpenApiVO) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *NatTraversalTunnelOpenApiVO) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *NatTraversalTunnelOpenApiVO) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *NatTraversalTunnelOpenApiVO) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetLocalAddress

`func (o *NatTraversalTunnelOpenApiVO) GetLocalAddress() string`

GetLocalAddress returns the LocalAddress field if non-nil, zero value otherwise.

### GetLocalAddressOk

`func (o *NatTraversalTunnelOpenApiVO) GetLocalAddressOk() (*string, bool)`

GetLocalAddressOk returns a tuple with the LocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddress

`func (o *NatTraversalTunnelOpenApiVO) SetLocalAddress(v string)`

SetLocalAddress sets LocalAddress field to given value.


### GetLocalMac

`func (o *NatTraversalTunnelOpenApiVO) GetLocalMac() string`

GetLocalMac returns the LocalMac field if non-nil, zero value otherwise.

### GetLocalMacOk

`func (o *NatTraversalTunnelOpenApiVO) GetLocalMacOk() (*string, bool)`

GetLocalMacOk returns a tuple with the LocalMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalMac

`func (o *NatTraversalTunnelOpenApiVO) SetLocalMac(v string)`

SetLocalMac sets LocalMac field to given value.

### HasLocalMac

`func (o *NatTraversalTunnelOpenApiVO) HasLocalMac() bool`

HasLocalMac returns a boolean if a field has been set.

### GetLocalPort

`func (o *NatTraversalTunnelOpenApiVO) GetLocalPort() int32`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *NatTraversalTunnelOpenApiVO) GetLocalPortOk() (*int32, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *NatTraversalTunnelOpenApiVO) SetLocalPort(v int32)`

SetLocalPort sets LocalPort field to given value.


### GetName

`func (o *NatTraversalTunnelOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NatTraversalTunnelOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NatTraversalTunnelOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetOpenStatus

`func (o *NatTraversalTunnelOpenApiVO) GetOpenStatus() bool`

GetOpenStatus returns the OpenStatus field if non-nil, zero value otherwise.

### GetOpenStatusOk

`func (o *NatTraversalTunnelOpenApiVO) GetOpenStatusOk() (*bool, bool)`

GetOpenStatusOk returns a tuple with the OpenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStatus

`func (o *NatTraversalTunnelOpenApiVO) SetOpenStatus(v bool)`

SetOpenStatus sets OpenStatus field to given value.

### HasOpenStatus

`func (o *NatTraversalTunnelOpenApiVO) HasOpenStatus() bool`

HasOpenStatus returns a boolean if a field has been set.

### GetTunnelEntryType

`func (o *NatTraversalTunnelOpenApiVO) GetTunnelEntryType() int32`

GetTunnelEntryType returns the TunnelEntryType field if non-nil, zero value otherwise.

### GetTunnelEntryTypeOk

`func (o *NatTraversalTunnelOpenApiVO) GetTunnelEntryTypeOk() (*int32, bool)`

GetTunnelEntryTypeOk returns a tuple with the TunnelEntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelEntryType

`func (o *NatTraversalTunnelOpenApiVO) SetTunnelEntryType(v int32)`

SetTunnelEntryType sets TunnelEntryType field to given value.

### HasTunnelEntryType

`func (o *NatTraversalTunnelOpenApiVO) HasTunnelEntryType() bool`

HasTunnelEntryType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


