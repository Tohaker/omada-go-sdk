# DhcpIpSettingEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpIp** | Pointer to **string** | Reserved IP address | [optional] 
**Fallback** | Pointer to **bool** | Enable/Disable fallback(if false, other params is not required) | [optional] 
**FallbackGate** | Pointer to **string** | Fallback gateway IP address | [optional] 
**FallbackIp** | Pointer to **string** | Fallback IP address | [optional] 
**FallbackMask** | Pointer to **string** | Fallback IP mask | [optional] 
**NetId** | Pointer to **string** | Network ID | [optional] 
**ServerMac** | Pointer to **string** | DHCP Server Mac | [optional] 
**ServerStackId** | Pointer to **string** | DHCP Server Stack ID | [optional] 
**ServerType** | Pointer to **string** | DHCP Server Type | [optional] 
**UseFixedAddr** | Pointer to **bool** | Enable reserved address(Gateway required) | [optional] 

## Methods

### NewDhcpIpSettingEntity

`func NewDhcpIpSettingEntity() *DhcpIpSettingEntity`

NewDhcpIpSettingEntity instantiates a new DhcpIpSettingEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpIpSettingEntityWithDefaults

`func NewDhcpIpSettingEntityWithDefaults() *DhcpIpSettingEntity`

NewDhcpIpSettingEntityWithDefaults instantiates a new DhcpIpSettingEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpIp

`func (o *DhcpIpSettingEntity) GetDhcpIp() string`

GetDhcpIp returns the DhcpIp field if non-nil, zero value otherwise.

### GetDhcpIpOk

`func (o *DhcpIpSettingEntity) GetDhcpIpOk() (*string, bool)`

GetDhcpIpOk returns a tuple with the DhcpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpIp

`func (o *DhcpIpSettingEntity) SetDhcpIp(v string)`

SetDhcpIp sets DhcpIp field to given value.

### HasDhcpIp

`func (o *DhcpIpSettingEntity) HasDhcpIp() bool`

HasDhcpIp returns a boolean if a field has been set.

### GetFallback

`func (o *DhcpIpSettingEntity) GetFallback() bool`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *DhcpIpSettingEntity) GetFallbackOk() (*bool, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *DhcpIpSettingEntity) SetFallback(v bool)`

SetFallback sets Fallback field to given value.

### HasFallback

`func (o *DhcpIpSettingEntity) HasFallback() bool`

HasFallback returns a boolean if a field has been set.

### GetFallbackGate

`func (o *DhcpIpSettingEntity) GetFallbackGate() string`

GetFallbackGate returns the FallbackGate field if non-nil, zero value otherwise.

### GetFallbackGateOk

`func (o *DhcpIpSettingEntity) GetFallbackGateOk() (*string, bool)`

GetFallbackGateOk returns a tuple with the FallbackGate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackGate

`func (o *DhcpIpSettingEntity) SetFallbackGate(v string)`

SetFallbackGate sets FallbackGate field to given value.

### HasFallbackGate

`func (o *DhcpIpSettingEntity) HasFallbackGate() bool`

HasFallbackGate returns a boolean if a field has been set.

### GetFallbackIp

`func (o *DhcpIpSettingEntity) GetFallbackIp() string`

GetFallbackIp returns the FallbackIp field if non-nil, zero value otherwise.

### GetFallbackIpOk

`func (o *DhcpIpSettingEntity) GetFallbackIpOk() (*string, bool)`

GetFallbackIpOk returns a tuple with the FallbackIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackIp

`func (o *DhcpIpSettingEntity) SetFallbackIp(v string)`

SetFallbackIp sets FallbackIp field to given value.

### HasFallbackIp

`func (o *DhcpIpSettingEntity) HasFallbackIp() bool`

HasFallbackIp returns a boolean if a field has been set.

### GetFallbackMask

`func (o *DhcpIpSettingEntity) GetFallbackMask() string`

GetFallbackMask returns the FallbackMask field if non-nil, zero value otherwise.

### GetFallbackMaskOk

`func (o *DhcpIpSettingEntity) GetFallbackMaskOk() (*string, bool)`

GetFallbackMaskOk returns a tuple with the FallbackMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackMask

`func (o *DhcpIpSettingEntity) SetFallbackMask(v string)`

SetFallbackMask sets FallbackMask field to given value.

### HasFallbackMask

`func (o *DhcpIpSettingEntity) HasFallbackMask() bool`

HasFallbackMask returns a boolean if a field has been set.

### GetNetId

`func (o *DhcpIpSettingEntity) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *DhcpIpSettingEntity) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *DhcpIpSettingEntity) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *DhcpIpSettingEntity) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetServerMac

`func (o *DhcpIpSettingEntity) GetServerMac() string`

GetServerMac returns the ServerMac field if non-nil, zero value otherwise.

### GetServerMacOk

`func (o *DhcpIpSettingEntity) GetServerMacOk() (*string, bool)`

GetServerMacOk returns a tuple with the ServerMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMac

`func (o *DhcpIpSettingEntity) SetServerMac(v string)`

SetServerMac sets ServerMac field to given value.

### HasServerMac

`func (o *DhcpIpSettingEntity) HasServerMac() bool`

HasServerMac returns a boolean if a field has been set.

### GetServerStackId

`func (o *DhcpIpSettingEntity) GetServerStackId() string`

GetServerStackId returns the ServerStackId field if non-nil, zero value otherwise.

### GetServerStackIdOk

`func (o *DhcpIpSettingEntity) GetServerStackIdOk() (*string, bool)`

GetServerStackIdOk returns a tuple with the ServerStackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStackId

`func (o *DhcpIpSettingEntity) SetServerStackId(v string)`

SetServerStackId sets ServerStackId field to given value.

### HasServerStackId

`func (o *DhcpIpSettingEntity) HasServerStackId() bool`

HasServerStackId returns a boolean if a field has been set.

### GetServerType

`func (o *DhcpIpSettingEntity) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DhcpIpSettingEntity) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DhcpIpSettingEntity) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DhcpIpSettingEntity) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetUseFixedAddr

`func (o *DhcpIpSettingEntity) GetUseFixedAddr() bool`

GetUseFixedAddr returns the UseFixedAddr field if non-nil, zero value otherwise.

### GetUseFixedAddrOk

`func (o *DhcpIpSettingEntity) GetUseFixedAddrOk() (*bool, bool)`

GetUseFixedAddrOk returns a tuple with the UseFixedAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFixedAddr

`func (o *DhcpIpSettingEntity) SetUseFixedAddr(v bool)`

SetUseFixedAddr sets UseFixedAddr field to given value.

### HasUseFixedAddr

`func (o *DhcpIpSettingEntity) HasUseFixedAddr() bool`

HasUseFixedAddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


