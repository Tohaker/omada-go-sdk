# PortForwardingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DMZ** | **bool** | With DMZ enabled, all ports are open and the traffic from external network will be forwarded to the specific destination IP in the LAN. | 
**ExistVirtualWan** | Pointer to **bool** | Whether Virtual Wan has been configured in current Port Forwarding. | [optional] 
**ExistWanIp** | Pointer to **bool** | Whether WAN IP has been configured in current Port Forwarding.  | [optional] 
**ExternalPort** | Pointer to **string** | External port corresponds with source port in web. External port should be within the range of 1–65535. Must be filled in when DMS is false. | [optional] 
**ForwardIp** | **string** | Forward IP corresponds with destination IP in web. Forward IP | 
**ForwardPort** | Pointer to **string** | Forward port corresponds with destination port in web. Forward port should be 1 or 1-10, within the range of 1–65535. Must be filled in when DMS is false. | [optional] 
**From** | **int32** | From corresponds with source IP in web. From should be a value as follows: 0: where; 1: limited address | 
**InterfaceWanPortId** | Pointer to **[]string** | This field represents WAN port ID. WAN port ID can be obtained from can be obtained from &#39;Get internet basic info&#39; interface. | [optional] 
**LimitedAddresses** | Pointer to **[]string** | Only for limited address | [optional] 
**Name** | **string** | Name, name should contain 1 to 64 characters. | 
**Protocol** | Pointer to **int32** | Protocol should be a value as follows: 0: ALL; 1: TCP; 2: UDP. Must be filled in when DMS is false. | [optional] 
**Status** | **bool** | Port-forwarding enable status. | 
**VirtualWanId** | Pointer to **[]string** | This field represents Virtual Wan ID. Virtual Wan Id can be obtained from &#39;Query available virtual WAN list&#39; | [optional] 
**WanIps** | Pointer to [**[]PortIpOpenApiVO**](PortIpOpenApiVO.md) | WAN IPs. Wan IPs can be obtained from &#39;Get internet ports config&#39; interface | [optional] 

## Methods

### NewPortForwardingConfig

`func NewPortForwardingConfig(dMZ bool, forwardIp string, from int32, name string, status bool, ) *PortForwardingConfig`

NewPortForwardingConfig instantiates a new PortForwardingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortForwardingConfigWithDefaults

`func NewPortForwardingConfigWithDefaults() *PortForwardingConfig`

NewPortForwardingConfigWithDefaults instantiates a new PortForwardingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDMZ

`func (o *PortForwardingConfig) GetDMZ() bool`

GetDMZ returns the DMZ field if non-nil, zero value otherwise.

### GetDMZOk

`func (o *PortForwardingConfig) GetDMZOk() (*bool, bool)`

GetDMZOk returns a tuple with the DMZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDMZ

`func (o *PortForwardingConfig) SetDMZ(v bool)`

SetDMZ sets DMZ field to given value.


### GetExistVirtualWan

`func (o *PortForwardingConfig) GetExistVirtualWan() bool`

GetExistVirtualWan returns the ExistVirtualWan field if non-nil, zero value otherwise.

### GetExistVirtualWanOk

`func (o *PortForwardingConfig) GetExistVirtualWanOk() (*bool, bool)`

GetExistVirtualWanOk returns a tuple with the ExistVirtualWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistVirtualWan

`func (o *PortForwardingConfig) SetExistVirtualWan(v bool)`

SetExistVirtualWan sets ExistVirtualWan field to given value.

### HasExistVirtualWan

`func (o *PortForwardingConfig) HasExistVirtualWan() bool`

HasExistVirtualWan returns a boolean if a field has been set.

### GetExistWanIp

`func (o *PortForwardingConfig) GetExistWanIp() bool`

GetExistWanIp returns the ExistWanIp field if non-nil, zero value otherwise.

### GetExistWanIpOk

`func (o *PortForwardingConfig) GetExistWanIpOk() (*bool, bool)`

GetExistWanIpOk returns a tuple with the ExistWanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistWanIp

`func (o *PortForwardingConfig) SetExistWanIp(v bool)`

SetExistWanIp sets ExistWanIp field to given value.

### HasExistWanIp

`func (o *PortForwardingConfig) HasExistWanIp() bool`

HasExistWanIp returns a boolean if a field has been set.

### GetExternalPort

`func (o *PortForwardingConfig) GetExternalPort() string`

GetExternalPort returns the ExternalPort field if non-nil, zero value otherwise.

### GetExternalPortOk

`func (o *PortForwardingConfig) GetExternalPortOk() (*string, bool)`

GetExternalPortOk returns a tuple with the ExternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPort

`func (o *PortForwardingConfig) SetExternalPort(v string)`

SetExternalPort sets ExternalPort field to given value.

### HasExternalPort

`func (o *PortForwardingConfig) HasExternalPort() bool`

HasExternalPort returns a boolean if a field has been set.

### GetForwardIp

`func (o *PortForwardingConfig) GetForwardIp() string`

GetForwardIp returns the ForwardIp field if non-nil, zero value otherwise.

### GetForwardIpOk

`func (o *PortForwardingConfig) GetForwardIpOk() (*string, bool)`

GetForwardIpOk returns a tuple with the ForwardIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardIp

`func (o *PortForwardingConfig) SetForwardIp(v string)`

SetForwardIp sets ForwardIp field to given value.


### GetForwardPort

`func (o *PortForwardingConfig) GetForwardPort() string`

GetForwardPort returns the ForwardPort field if non-nil, zero value otherwise.

### GetForwardPortOk

`func (o *PortForwardingConfig) GetForwardPortOk() (*string, bool)`

GetForwardPortOk returns a tuple with the ForwardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardPort

`func (o *PortForwardingConfig) SetForwardPort(v string)`

SetForwardPort sets ForwardPort field to given value.

### HasForwardPort

`func (o *PortForwardingConfig) HasForwardPort() bool`

HasForwardPort returns a boolean if a field has been set.

### GetFrom

`func (o *PortForwardingConfig) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PortForwardingConfig) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PortForwardingConfig) SetFrom(v int32)`

SetFrom sets From field to given value.


### GetInterfaceWanPortId

`func (o *PortForwardingConfig) GetInterfaceWanPortId() []string`

GetInterfaceWanPortId returns the InterfaceWanPortId field if non-nil, zero value otherwise.

### GetInterfaceWanPortIdOk

`func (o *PortForwardingConfig) GetInterfaceWanPortIdOk() (*[]string, bool)`

GetInterfaceWanPortIdOk returns a tuple with the InterfaceWanPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceWanPortId

`func (o *PortForwardingConfig) SetInterfaceWanPortId(v []string)`

SetInterfaceWanPortId sets InterfaceWanPortId field to given value.

### HasInterfaceWanPortId

`func (o *PortForwardingConfig) HasInterfaceWanPortId() bool`

HasInterfaceWanPortId returns a boolean if a field has been set.

### GetLimitedAddresses

`func (o *PortForwardingConfig) GetLimitedAddresses() []string`

GetLimitedAddresses returns the LimitedAddresses field if non-nil, zero value otherwise.

### GetLimitedAddressesOk

`func (o *PortForwardingConfig) GetLimitedAddressesOk() (*[]string, bool)`

GetLimitedAddressesOk returns a tuple with the LimitedAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitedAddresses

`func (o *PortForwardingConfig) SetLimitedAddresses(v []string)`

SetLimitedAddresses sets LimitedAddresses field to given value.

### HasLimitedAddresses

`func (o *PortForwardingConfig) HasLimitedAddresses() bool`

HasLimitedAddresses returns a boolean if a field has been set.

### GetName

`func (o *PortForwardingConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortForwardingConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortForwardingConfig) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *PortForwardingConfig) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PortForwardingConfig) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PortForwardingConfig) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PortForwardingConfig) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetStatus

`func (o *PortForwardingConfig) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortForwardingConfig) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortForwardingConfig) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetVirtualWanId

`func (o *PortForwardingConfig) GetVirtualWanId() []string`

GetVirtualWanId returns the VirtualWanId field if non-nil, zero value otherwise.

### GetVirtualWanIdOk

`func (o *PortForwardingConfig) GetVirtualWanIdOk() (*[]string, bool)`

GetVirtualWanIdOk returns a tuple with the VirtualWanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanId

`func (o *PortForwardingConfig) SetVirtualWanId(v []string)`

SetVirtualWanId sets VirtualWanId field to given value.

### HasVirtualWanId

`func (o *PortForwardingConfig) HasVirtualWanId() bool`

HasVirtualWanId returns a boolean if a field has been set.

### GetWanIps

`func (o *PortForwardingConfig) GetWanIps() []PortIpOpenApiVO`

GetWanIps returns the WanIps field if non-nil, zero value otherwise.

### GetWanIpsOk

`func (o *PortForwardingConfig) GetWanIpsOk() (*[]PortIpOpenApiVO, bool)`

GetWanIpsOk returns a tuple with the WanIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanIps

`func (o *PortForwardingConfig) SetWanIps(v []PortIpOpenApiVO)`

SetWanIps sets WanIps field to given value.

### HasWanIps

`func (o *PortForwardingConfig) HasWanIps() bool`

HasWanIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


