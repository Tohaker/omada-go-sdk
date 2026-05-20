# NatTraversalTunnelModifyOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppType** | **string** | App type of the remote access tunnel for the local target device. HTTP, HTTPS, SSH, TELNET. | 
**ClientDeviceType** | Pointer to **string** | Client&#39;s type of the remote access tunnel for the client. | [optional] 
**ClientManager** | Pointer to **bool** | Client&#39;s manager of the remote access tunnel for the client. | [optional] 
**ClientModel** | Pointer to **string** | Client&#39;s model of the remote access tunnel for the client. | [optional] 
**ClientName** | Pointer to **string** | Client&#39;s name of the remote access tunnel for the client. | [optional] 
**Duration** | Pointer to **int32** | Valid duration time of the remote access tunnel, 1-24 hours. | [optional] 
**LocalAddress** | Pointer to **string** | IP address of the local target device. | [optional] 
**LocalMac** | Pointer to **string** | Mac of the local target device. | [optional] 
**LocalPort** | Pointer to **int32** | Port of the local target device&#39;s service. | [optional] 
**Name** | Pointer to **string** | Name of the remote access tunnel. | [optional] 
**OpenStatus** | Pointer to **bool** | If open the remote access tunnel after create. | [optional] 
**TunnelEntryType** | Pointer to **int32** | Entry of the remote access tunnel.1: custom tunnel 2: device tunnel 3:device detail 4:client detail | [optional] 

## Methods

### NewNatTraversalTunnelModifyOpenApiVO

`func NewNatTraversalTunnelModifyOpenApiVO(appType string, ) *NatTraversalTunnelModifyOpenApiVO`

NewNatTraversalTunnelModifyOpenApiVO instantiates a new NatTraversalTunnelModifyOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatTraversalTunnelModifyOpenApiVOWithDefaults

`func NewNatTraversalTunnelModifyOpenApiVOWithDefaults() *NatTraversalTunnelModifyOpenApiVO`

NewNatTraversalTunnelModifyOpenApiVOWithDefaults instantiates a new NatTraversalTunnelModifyOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppType

`func (o *NatTraversalTunnelModifyOpenApiVO) GetAppType() string`

GetAppType returns the AppType field if non-nil, zero value otherwise.

### GetAppTypeOk

`func (o *NatTraversalTunnelModifyOpenApiVO) GetAppTypeOk() (*string, bool)`

GetAppTypeOk returns a tuple with the AppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppType

`func (o *NatTraversalTunnelModifyOpenApiVO) SetAppType(v string)`

SetAppType sets AppType field to given value.


### GetClientDeviceType

`func (o *NatTraversalTunnelModifyOpenApiVO) GetClientDeviceType() string`

GetClientDeviceType returns the ClientDeviceType field if non-nil, zero value otherwise.

### GetClientDeviceTypeOk

`func (o *NatTraversalTunnelModifyOpenApiVO) GetClientDeviceTypeOk() (*string, bool)`

GetClientDeviceTypeOk returns a tuple with the ClientDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDeviceType

`func (o *NatTraversalTunnelModifyOpenApiVO) SetClientDeviceType(v string)`

SetClientDeviceType sets ClientDeviceType field to given value.

### HasClientDeviceType

`func (o *NatTraversalTunnelModifyOpenApiVO) HasClientDeviceType() bool`

HasClientDeviceType returns a boolean if a field has been set.

### GetClientManager

`func (o *NatTraversalTunnelModifyOpenApiVO) GetClientManager() bool`

GetClientManager returns the ClientManager field if non-nil, zero value otherwise.

### GetClientManagerOk

`func (o *NatTraversalTunnelModifyOpenApiVO) GetClientManagerOk() (*bool, bool)`

GetClientManagerOk returns a tuple with the ClientManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientManager

`func (o *NatTraversalTunnelModifyOpenApiVO) SetClientManager(v bool)`

SetClientManager sets ClientManager field to given value.

### HasClientManager

`func (o *NatTraversalTunnelModifyOpenApiVO) HasClientManager() bool`

HasClientManager returns a boolean if a field has been set.

### GetClientModel

`func (o *NatTraversalTunnelModifyOpenApiVO) GetClientModel() string`

GetClientModel returns the ClientModel field if non-nil, zero value otherwise.

### GetClientModelOk

`func (o *NatTraversalTunnelModifyOpenApiVO) GetClientModelOk() (*string, bool)`

GetClientModelOk returns a tuple with the ClientModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientModel

`func (o *NatTraversalTunnelModifyOpenApiVO) SetClientModel(v string)`

SetClientModel sets ClientModel field to given value.

### HasClientModel

`func (o *NatTraversalTunnelModifyOpenApiVO) HasClientModel() bool`

HasClientModel returns a boolean if a field has been set.

### GetClientName

`func (o *NatTraversalTunnelModifyOpenApiVO) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *NatTraversalTunnelModifyOpenApiVO) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *NatTraversalTunnelModifyOpenApiVO) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *NatTraversalTunnelModifyOpenApiVO) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetDuration

`func (o *NatTraversalTunnelModifyOpenApiVO) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *NatTraversalTunnelModifyOpenApiVO) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *NatTraversalTunnelModifyOpenApiVO) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *NatTraversalTunnelModifyOpenApiVO) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetLocalAddress

`func (o *NatTraversalTunnelModifyOpenApiVO) GetLocalAddress() string`

GetLocalAddress returns the LocalAddress field if non-nil, zero value otherwise.

### GetLocalAddressOk

`func (o *NatTraversalTunnelModifyOpenApiVO) GetLocalAddressOk() (*string, bool)`

GetLocalAddressOk returns a tuple with the LocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddress

`func (o *NatTraversalTunnelModifyOpenApiVO) SetLocalAddress(v string)`

SetLocalAddress sets LocalAddress field to given value.

### HasLocalAddress

`func (o *NatTraversalTunnelModifyOpenApiVO) HasLocalAddress() bool`

HasLocalAddress returns a boolean if a field has been set.

### GetLocalMac

`func (o *NatTraversalTunnelModifyOpenApiVO) GetLocalMac() string`

GetLocalMac returns the LocalMac field if non-nil, zero value otherwise.

### GetLocalMacOk

`func (o *NatTraversalTunnelModifyOpenApiVO) GetLocalMacOk() (*string, bool)`

GetLocalMacOk returns a tuple with the LocalMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalMac

`func (o *NatTraversalTunnelModifyOpenApiVO) SetLocalMac(v string)`

SetLocalMac sets LocalMac field to given value.

### HasLocalMac

`func (o *NatTraversalTunnelModifyOpenApiVO) HasLocalMac() bool`

HasLocalMac returns a boolean if a field has been set.

### GetLocalPort

`func (o *NatTraversalTunnelModifyOpenApiVO) GetLocalPort() int32`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *NatTraversalTunnelModifyOpenApiVO) GetLocalPortOk() (*int32, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *NatTraversalTunnelModifyOpenApiVO) SetLocalPort(v int32)`

SetLocalPort sets LocalPort field to given value.

### HasLocalPort

`func (o *NatTraversalTunnelModifyOpenApiVO) HasLocalPort() bool`

HasLocalPort returns a boolean if a field has been set.

### GetName

`func (o *NatTraversalTunnelModifyOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NatTraversalTunnelModifyOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NatTraversalTunnelModifyOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NatTraversalTunnelModifyOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpenStatus

`func (o *NatTraversalTunnelModifyOpenApiVO) GetOpenStatus() bool`

GetOpenStatus returns the OpenStatus field if non-nil, zero value otherwise.

### GetOpenStatusOk

`func (o *NatTraversalTunnelModifyOpenApiVO) GetOpenStatusOk() (*bool, bool)`

GetOpenStatusOk returns a tuple with the OpenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStatus

`func (o *NatTraversalTunnelModifyOpenApiVO) SetOpenStatus(v bool)`

SetOpenStatus sets OpenStatus field to given value.

### HasOpenStatus

`func (o *NatTraversalTunnelModifyOpenApiVO) HasOpenStatus() bool`

HasOpenStatus returns a boolean if a field has been set.

### GetTunnelEntryType

`func (o *NatTraversalTunnelModifyOpenApiVO) GetTunnelEntryType() int32`

GetTunnelEntryType returns the TunnelEntryType field if non-nil, zero value otherwise.

### GetTunnelEntryTypeOk

`func (o *NatTraversalTunnelModifyOpenApiVO) GetTunnelEntryTypeOk() (*int32, bool)`

GetTunnelEntryTypeOk returns a tuple with the TunnelEntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelEntryType

`func (o *NatTraversalTunnelModifyOpenApiVO) SetTunnelEntryType(v int32)`

SetTunnelEntryType sets TunnelEntryType field to given value.

### HasTunnelEntryType

`func (o *NatTraversalTunnelModifyOpenApiVO) HasTunnelEntryType() bool`

HasTunnelEntryType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


