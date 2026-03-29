# PortForwardingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | Pointer to **int64** | Total flow in byte. | [optional] 
**EntryId** | Pointer to **int32** | Port Forwarding entry id. | [optional] 
**ExternalPort** | Pointer to **string** | External Port. | [optional] 
**From** | Pointer to **int32** | 0:Anywhere 1:Limited Address. | [optional] 
**InterfaceWanPortId** | Pointer to **[]string** | Port Forwarding item id. | [optional] 
**InternalIp** | Pointer to **string** | Internal IP. | [optional] 
**InternalPort** | Pointer to **string** | Internal Port. | [optional] 
**LeaseDuration** | Pointer to **int64** | Displays the duration of the UPnP port forwarding in seconds. | [optional] 
**LimitedAddresses** | Pointer to **[]string** | If the field from is 1,limitedAddresses includes limited addresses. | [optional] 
**Name** | Pointer to **string** | Port Forwarding rule name. | [optional] 
**OmadacId** | Pointer to **string** | Omada ID | [optional] 
**Packets** | Pointer to **int64** | Packet quantity. | [optional] 
**Protocol** | Pointer to **int32** | 0:ALL, 1:TCP, 2:UDP. | [optional] 
**SiteId** | Pointer to **string** | Site ID | [optional] 

## Methods

### NewPortForwardingOpenApiVO

`func NewPortForwardingOpenApiVO() *PortForwardingOpenApiVO`

NewPortForwardingOpenApiVO instantiates a new PortForwardingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortForwardingOpenApiVOWithDefaults

`func NewPortForwardingOpenApiVOWithDefaults() *PortForwardingOpenApiVO`

NewPortForwardingOpenApiVOWithDefaults instantiates a new PortForwardingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *PortForwardingOpenApiVO) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *PortForwardingOpenApiVO) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *PortForwardingOpenApiVO) SetBytes(v int64)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *PortForwardingOpenApiVO) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetEntryId

`func (o *PortForwardingOpenApiVO) GetEntryId() int32`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *PortForwardingOpenApiVO) GetEntryIdOk() (*int32, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *PortForwardingOpenApiVO) SetEntryId(v int32)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *PortForwardingOpenApiVO) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetExternalPort

`func (o *PortForwardingOpenApiVO) GetExternalPort() string`

GetExternalPort returns the ExternalPort field if non-nil, zero value otherwise.

### GetExternalPortOk

`func (o *PortForwardingOpenApiVO) GetExternalPortOk() (*string, bool)`

GetExternalPortOk returns a tuple with the ExternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPort

`func (o *PortForwardingOpenApiVO) SetExternalPort(v string)`

SetExternalPort sets ExternalPort field to given value.

### HasExternalPort

`func (o *PortForwardingOpenApiVO) HasExternalPort() bool`

HasExternalPort returns a boolean if a field has been set.

### GetFrom

`func (o *PortForwardingOpenApiVO) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PortForwardingOpenApiVO) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PortForwardingOpenApiVO) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PortForwardingOpenApiVO) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetInterfaceWanPortId

`func (o *PortForwardingOpenApiVO) GetInterfaceWanPortId() []string`

GetInterfaceWanPortId returns the InterfaceWanPortId field if non-nil, zero value otherwise.

### GetInterfaceWanPortIdOk

`func (o *PortForwardingOpenApiVO) GetInterfaceWanPortIdOk() (*[]string, bool)`

GetInterfaceWanPortIdOk returns a tuple with the InterfaceWanPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceWanPortId

`func (o *PortForwardingOpenApiVO) SetInterfaceWanPortId(v []string)`

SetInterfaceWanPortId sets InterfaceWanPortId field to given value.

### HasInterfaceWanPortId

`func (o *PortForwardingOpenApiVO) HasInterfaceWanPortId() bool`

HasInterfaceWanPortId returns a boolean if a field has been set.

### GetInternalIp

`func (o *PortForwardingOpenApiVO) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *PortForwardingOpenApiVO) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *PortForwardingOpenApiVO) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *PortForwardingOpenApiVO) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetInternalPort

`func (o *PortForwardingOpenApiVO) GetInternalPort() string`

GetInternalPort returns the InternalPort field if non-nil, zero value otherwise.

### GetInternalPortOk

`func (o *PortForwardingOpenApiVO) GetInternalPortOk() (*string, bool)`

GetInternalPortOk returns a tuple with the InternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalPort

`func (o *PortForwardingOpenApiVO) SetInternalPort(v string)`

SetInternalPort sets InternalPort field to given value.

### HasInternalPort

`func (o *PortForwardingOpenApiVO) HasInternalPort() bool`

HasInternalPort returns a boolean if a field has been set.

### GetLeaseDuration

`func (o *PortForwardingOpenApiVO) GetLeaseDuration() int64`

GetLeaseDuration returns the LeaseDuration field if non-nil, zero value otherwise.

### GetLeaseDurationOk

`func (o *PortForwardingOpenApiVO) GetLeaseDurationOk() (*int64, bool)`

GetLeaseDurationOk returns a tuple with the LeaseDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseDuration

`func (o *PortForwardingOpenApiVO) SetLeaseDuration(v int64)`

SetLeaseDuration sets LeaseDuration field to given value.

### HasLeaseDuration

`func (o *PortForwardingOpenApiVO) HasLeaseDuration() bool`

HasLeaseDuration returns a boolean if a field has been set.

### GetLimitedAddresses

`func (o *PortForwardingOpenApiVO) GetLimitedAddresses() []string`

GetLimitedAddresses returns the LimitedAddresses field if non-nil, zero value otherwise.

### GetLimitedAddressesOk

`func (o *PortForwardingOpenApiVO) GetLimitedAddressesOk() (*[]string, bool)`

GetLimitedAddressesOk returns a tuple with the LimitedAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitedAddresses

`func (o *PortForwardingOpenApiVO) SetLimitedAddresses(v []string)`

SetLimitedAddresses sets LimitedAddresses field to given value.

### HasLimitedAddresses

`func (o *PortForwardingOpenApiVO) HasLimitedAddresses() bool`

HasLimitedAddresses returns a boolean if a field has been set.

### GetName

`func (o *PortForwardingOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortForwardingOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortForwardingOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PortForwardingOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOmadacId

`func (o *PortForwardingOpenApiVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *PortForwardingOpenApiVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *PortForwardingOpenApiVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *PortForwardingOpenApiVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetPackets

`func (o *PortForwardingOpenApiVO) GetPackets() int64`

GetPackets returns the Packets field if non-nil, zero value otherwise.

### GetPacketsOk

`func (o *PortForwardingOpenApiVO) GetPacketsOk() (*int64, bool)`

GetPacketsOk returns a tuple with the Packets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackets

`func (o *PortForwardingOpenApiVO) SetPackets(v int64)`

SetPackets sets Packets field to given value.

### HasPackets

`func (o *PortForwardingOpenApiVO) HasPackets() bool`

HasPackets returns a boolean if a field has been set.

### GetProtocol

`func (o *PortForwardingOpenApiVO) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PortForwardingOpenApiVO) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PortForwardingOpenApiVO) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PortForwardingOpenApiVO) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSiteId

`func (o *PortForwardingOpenApiVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *PortForwardingOpenApiVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *PortForwardingOpenApiVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *PortForwardingOpenApiVO) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


