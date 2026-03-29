# PortForwardingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DMZ** | **bool** | With DMZ enabled, all ports are open and the traffic from external network will be forwarded to the specific destination IP in the LAN. | 
**ExistVirtualWan** | Pointer to **bool** | true: exist virtual WAN config | [optional] 
**ExistWanIp** | Pointer to **bool** | true: exist WAN IP config | [optional] 
**ExternalPort** | Pointer to **string** | External port corresponds with source port in web. ExternalPort should be within the range of 1–65535. | [optional] 
**ForwardIp** | **string** | Forward IP corresponds with destination IP in web. Forward IP | 
**ForwardPort** | Pointer to **string** | Forward port corresponds with destination port in web. ForwardPort should be 1 or 1-10, within the range of 1–65535. | [optional] 
**From** | **int32** | From corresponds with source IP in web. From should be a value as follows: 0: where; 1: limited address | 
**Id** | Pointer to **string** | Port-forwarding ID | [optional] 
**InterfaceWanPortId** | **[]string** | This field represents WAN port ID. WAN port ID can be obtained from &#39;Get internet basic info&#39; interface. | 
**LimitedAddresses** | Pointer to **[]string** | Only for limited address | [optional] 
**Name** | **string** | Name, name should contain 1 to 64 characters. | 
**Protocol** | Pointer to **int32** | Protocol should be a value as follows: 0: ALL; 1: TCP; 2: UDP. | [optional] 
**SiteId** | Pointer to **string** | Site ID | [optional] 
**Status** | **bool** | Port-forwarding enable status. | 
**VirtualWanId** | Pointer to **[]string** | This field represents VirtualWan ID. VirtualWan ID can be obtained from &#39;Query available virtual WAN list&#39; interface | [optional] 
**WanIps** | Pointer to [**[]PortIpOpenApiVO**](PortIpOpenApiVO.md) | WAN IPs | [optional] 

## Methods

### NewPortForwardingInfo

`func NewPortForwardingInfo(dMZ bool, forwardIp string, from int32, interfaceWanPortId []string, name string, status bool, ) *PortForwardingInfo`

NewPortForwardingInfo instantiates a new PortForwardingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortForwardingInfoWithDefaults

`func NewPortForwardingInfoWithDefaults() *PortForwardingInfo`

NewPortForwardingInfoWithDefaults instantiates a new PortForwardingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDMZ

`func (o *PortForwardingInfo) GetDMZ() bool`

GetDMZ returns the DMZ field if non-nil, zero value otherwise.

### GetDMZOk

`func (o *PortForwardingInfo) GetDMZOk() (*bool, bool)`

GetDMZOk returns a tuple with the DMZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDMZ

`func (o *PortForwardingInfo) SetDMZ(v bool)`

SetDMZ sets DMZ field to given value.


### GetExistVirtualWan

`func (o *PortForwardingInfo) GetExistVirtualWan() bool`

GetExistVirtualWan returns the ExistVirtualWan field if non-nil, zero value otherwise.

### GetExistVirtualWanOk

`func (o *PortForwardingInfo) GetExistVirtualWanOk() (*bool, bool)`

GetExistVirtualWanOk returns a tuple with the ExistVirtualWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistVirtualWan

`func (o *PortForwardingInfo) SetExistVirtualWan(v bool)`

SetExistVirtualWan sets ExistVirtualWan field to given value.

### HasExistVirtualWan

`func (o *PortForwardingInfo) HasExistVirtualWan() bool`

HasExistVirtualWan returns a boolean if a field has been set.

### GetExistWanIp

`func (o *PortForwardingInfo) GetExistWanIp() bool`

GetExistWanIp returns the ExistWanIp field if non-nil, zero value otherwise.

### GetExistWanIpOk

`func (o *PortForwardingInfo) GetExistWanIpOk() (*bool, bool)`

GetExistWanIpOk returns a tuple with the ExistWanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistWanIp

`func (o *PortForwardingInfo) SetExistWanIp(v bool)`

SetExistWanIp sets ExistWanIp field to given value.

### HasExistWanIp

`func (o *PortForwardingInfo) HasExistWanIp() bool`

HasExistWanIp returns a boolean if a field has been set.

### GetExternalPort

`func (o *PortForwardingInfo) GetExternalPort() string`

GetExternalPort returns the ExternalPort field if non-nil, zero value otherwise.

### GetExternalPortOk

`func (o *PortForwardingInfo) GetExternalPortOk() (*string, bool)`

GetExternalPortOk returns a tuple with the ExternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPort

`func (o *PortForwardingInfo) SetExternalPort(v string)`

SetExternalPort sets ExternalPort field to given value.

### HasExternalPort

`func (o *PortForwardingInfo) HasExternalPort() bool`

HasExternalPort returns a boolean if a field has been set.

### GetForwardIp

`func (o *PortForwardingInfo) GetForwardIp() string`

GetForwardIp returns the ForwardIp field if non-nil, zero value otherwise.

### GetForwardIpOk

`func (o *PortForwardingInfo) GetForwardIpOk() (*string, bool)`

GetForwardIpOk returns a tuple with the ForwardIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardIp

`func (o *PortForwardingInfo) SetForwardIp(v string)`

SetForwardIp sets ForwardIp field to given value.


### GetForwardPort

`func (o *PortForwardingInfo) GetForwardPort() string`

GetForwardPort returns the ForwardPort field if non-nil, zero value otherwise.

### GetForwardPortOk

`func (o *PortForwardingInfo) GetForwardPortOk() (*string, bool)`

GetForwardPortOk returns a tuple with the ForwardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardPort

`func (o *PortForwardingInfo) SetForwardPort(v string)`

SetForwardPort sets ForwardPort field to given value.

### HasForwardPort

`func (o *PortForwardingInfo) HasForwardPort() bool`

HasForwardPort returns a boolean if a field has been set.

### GetFrom

`func (o *PortForwardingInfo) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PortForwardingInfo) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PortForwardingInfo) SetFrom(v int32)`

SetFrom sets From field to given value.


### GetId

`func (o *PortForwardingInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortForwardingInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortForwardingInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortForwardingInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceWanPortId

`func (o *PortForwardingInfo) GetInterfaceWanPortId() []string`

GetInterfaceWanPortId returns the InterfaceWanPortId field if non-nil, zero value otherwise.

### GetInterfaceWanPortIdOk

`func (o *PortForwardingInfo) GetInterfaceWanPortIdOk() (*[]string, bool)`

GetInterfaceWanPortIdOk returns a tuple with the InterfaceWanPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceWanPortId

`func (o *PortForwardingInfo) SetInterfaceWanPortId(v []string)`

SetInterfaceWanPortId sets InterfaceWanPortId field to given value.


### GetLimitedAddresses

`func (o *PortForwardingInfo) GetLimitedAddresses() []string`

GetLimitedAddresses returns the LimitedAddresses field if non-nil, zero value otherwise.

### GetLimitedAddressesOk

`func (o *PortForwardingInfo) GetLimitedAddressesOk() (*[]string, bool)`

GetLimitedAddressesOk returns a tuple with the LimitedAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitedAddresses

`func (o *PortForwardingInfo) SetLimitedAddresses(v []string)`

SetLimitedAddresses sets LimitedAddresses field to given value.

### HasLimitedAddresses

`func (o *PortForwardingInfo) HasLimitedAddresses() bool`

HasLimitedAddresses returns a boolean if a field has been set.

### GetName

`func (o *PortForwardingInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortForwardingInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortForwardingInfo) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *PortForwardingInfo) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PortForwardingInfo) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PortForwardingInfo) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PortForwardingInfo) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSiteId

`func (o *PortForwardingInfo) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *PortForwardingInfo) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *PortForwardingInfo) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *PortForwardingInfo) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetStatus

`func (o *PortForwardingInfo) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortForwardingInfo) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortForwardingInfo) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetVirtualWanId

`func (o *PortForwardingInfo) GetVirtualWanId() []string`

GetVirtualWanId returns the VirtualWanId field if non-nil, zero value otherwise.

### GetVirtualWanIdOk

`func (o *PortForwardingInfo) GetVirtualWanIdOk() (*[]string, bool)`

GetVirtualWanIdOk returns a tuple with the VirtualWanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanId

`func (o *PortForwardingInfo) SetVirtualWanId(v []string)`

SetVirtualWanId sets VirtualWanId field to given value.

### HasVirtualWanId

`func (o *PortForwardingInfo) HasVirtualWanId() bool`

HasVirtualWanId returns a boolean if a field has been set.

### GetWanIps

`func (o *PortForwardingInfo) GetWanIps() []PortIpOpenApiVO`

GetWanIps returns the WanIps field if non-nil, zero value otherwise.

### GetWanIpsOk

`func (o *PortForwardingInfo) GetWanIpsOk() (*[]PortIpOpenApiVO, bool)`

GetWanIpsOk returns a tuple with the WanIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanIps

`func (o *PortForwardingInfo) SetWanIps(v []PortIpOpenApiVO)`

SetWanIps sets WanIps field to given value.

### HasWanIps

`func (o *PortForwardingInfo) HasWanIps() bool`

HasWanIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


