# PortForwardingTemplateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DMZ** | **bool** | With DMZ enabled, all ports are open and the traffic from external network will be forwarded to the specific destination IP in the LAN. | 
**ExternalPort** | Pointer to **string** | External port corresponds with source port in web. External port should be within the range of 1–65535. Must be filled in when DMS is false. | [optional] 
**ForwardIp** | **string** | Forward IP corresponds with destination IP in web. Forward IP | 
**ForwardPort** | Pointer to **string** | Forward port corresponds with destination port in web. Forward port should be 1 or 1-10, within the range of 1–65535. Must be filled in when DMS is false. | [optional] 
**From** | **int32** | From corresponds with source IP in web. From should be a value as follows: 0: where; 1: limited address | 
**InterfaceWanPortId** | **[]string** | This field represents WAN port ID. WAN port ID can be obtained from can be obtained from &#39;Get internet basic info&#39; interface. | 
**LimitedAddresses** | Pointer to **[]string** | Only for limited address | [optional] 
**Name** | **string** | Name, name should contain 1 to 64 characters. | 
**Protocol** | Pointer to **int32** | Protocol should be a value as follows: 0: ALL; 1: TCP; 2: UDP. Must be filled in when DMS is false. | [optional] 
**Status** | **bool** | Port-forwarding enable status. | 

## Methods

### NewPortForwardingTemplateOpenApiVO

`func NewPortForwardingTemplateOpenApiVO(dMZ bool, forwardIp string, from int32, interfaceWanPortId []string, name string, status bool, ) *PortForwardingTemplateOpenApiVO`

NewPortForwardingTemplateOpenApiVO instantiates a new PortForwardingTemplateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortForwardingTemplateOpenApiVOWithDefaults

`func NewPortForwardingTemplateOpenApiVOWithDefaults() *PortForwardingTemplateOpenApiVO`

NewPortForwardingTemplateOpenApiVOWithDefaults instantiates a new PortForwardingTemplateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDMZ

`func (o *PortForwardingTemplateOpenApiVO) GetDMZ() bool`

GetDMZ returns the DMZ field if non-nil, zero value otherwise.

### GetDMZOk

`func (o *PortForwardingTemplateOpenApiVO) GetDMZOk() (*bool, bool)`

GetDMZOk returns a tuple with the DMZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDMZ

`func (o *PortForwardingTemplateOpenApiVO) SetDMZ(v bool)`

SetDMZ sets DMZ field to given value.


### GetExternalPort

`func (o *PortForwardingTemplateOpenApiVO) GetExternalPort() string`

GetExternalPort returns the ExternalPort field if non-nil, zero value otherwise.

### GetExternalPortOk

`func (o *PortForwardingTemplateOpenApiVO) GetExternalPortOk() (*string, bool)`

GetExternalPortOk returns a tuple with the ExternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPort

`func (o *PortForwardingTemplateOpenApiVO) SetExternalPort(v string)`

SetExternalPort sets ExternalPort field to given value.

### HasExternalPort

`func (o *PortForwardingTemplateOpenApiVO) HasExternalPort() bool`

HasExternalPort returns a boolean if a field has been set.

### GetForwardIp

`func (o *PortForwardingTemplateOpenApiVO) GetForwardIp() string`

GetForwardIp returns the ForwardIp field if non-nil, zero value otherwise.

### GetForwardIpOk

`func (o *PortForwardingTemplateOpenApiVO) GetForwardIpOk() (*string, bool)`

GetForwardIpOk returns a tuple with the ForwardIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardIp

`func (o *PortForwardingTemplateOpenApiVO) SetForwardIp(v string)`

SetForwardIp sets ForwardIp field to given value.


### GetForwardPort

`func (o *PortForwardingTemplateOpenApiVO) GetForwardPort() string`

GetForwardPort returns the ForwardPort field if non-nil, zero value otherwise.

### GetForwardPortOk

`func (o *PortForwardingTemplateOpenApiVO) GetForwardPortOk() (*string, bool)`

GetForwardPortOk returns a tuple with the ForwardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardPort

`func (o *PortForwardingTemplateOpenApiVO) SetForwardPort(v string)`

SetForwardPort sets ForwardPort field to given value.

### HasForwardPort

`func (o *PortForwardingTemplateOpenApiVO) HasForwardPort() bool`

HasForwardPort returns a boolean if a field has been set.

### GetFrom

`func (o *PortForwardingTemplateOpenApiVO) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PortForwardingTemplateOpenApiVO) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PortForwardingTemplateOpenApiVO) SetFrom(v int32)`

SetFrom sets From field to given value.


### GetInterfaceWanPortId

`func (o *PortForwardingTemplateOpenApiVO) GetInterfaceWanPortId() []string`

GetInterfaceWanPortId returns the InterfaceWanPortId field if non-nil, zero value otherwise.

### GetInterfaceWanPortIdOk

`func (o *PortForwardingTemplateOpenApiVO) GetInterfaceWanPortIdOk() (*[]string, bool)`

GetInterfaceWanPortIdOk returns a tuple with the InterfaceWanPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceWanPortId

`func (o *PortForwardingTemplateOpenApiVO) SetInterfaceWanPortId(v []string)`

SetInterfaceWanPortId sets InterfaceWanPortId field to given value.


### GetLimitedAddresses

`func (o *PortForwardingTemplateOpenApiVO) GetLimitedAddresses() []string`

GetLimitedAddresses returns the LimitedAddresses field if non-nil, zero value otherwise.

### GetLimitedAddressesOk

`func (o *PortForwardingTemplateOpenApiVO) GetLimitedAddressesOk() (*[]string, bool)`

GetLimitedAddressesOk returns a tuple with the LimitedAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitedAddresses

`func (o *PortForwardingTemplateOpenApiVO) SetLimitedAddresses(v []string)`

SetLimitedAddresses sets LimitedAddresses field to given value.

### HasLimitedAddresses

`func (o *PortForwardingTemplateOpenApiVO) HasLimitedAddresses() bool`

HasLimitedAddresses returns a boolean if a field has been set.

### GetName

`func (o *PortForwardingTemplateOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortForwardingTemplateOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortForwardingTemplateOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *PortForwardingTemplateOpenApiVO) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PortForwardingTemplateOpenApiVO) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PortForwardingTemplateOpenApiVO) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PortForwardingTemplateOpenApiVO) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetStatus

`func (o *PortForwardingTemplateOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortForwardingTemplateOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortForwardingTemplateOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


