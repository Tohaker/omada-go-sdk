# OswIpSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmConflict** | Pointer to **bool** | Enable IP-MAC Conflict Detection or not. | [optional] 
**DhcpIp** | Pointer to **string** | The IP for Address reservation when useFixedAddr is enabled. | [optional] 
**Fallback** | Pointer to **bool** | The switch of fallback ip for mode 1. | [optional] 
**FallbackGate** | Pointer to **string** | Fallback gateway for dhcp mode when fallback is enabled. | [optional] 
**FallbackIp** | Pointer to **string** | Fallback ip for dhcp mode when fallback is enabled. | [optional] 
**FallbackMask** | Pointer to **string** | Fallback mask for dhcp mode when fallback is enabled. | [optional] 
**Gateway** | Pointer to **string** | Gateway, like 192.168.137.1 | [optional] 
**Ip** | Pointer to **string** | Static IP for mode 0, like 192.168.0.1 | [optional] 
**Mode** | **int32** | IP Setting mode. Static:0, DHCP:1 | 
**NetId** | Pointer to **string** | The LAN Network for Address reservation when useFixedAddr is enabled. Obtain the id from \&quot;Get LAN network list\&quot; | [optional] 
**Netmask** | Pointer to **string** | IP Mask, like 255.255.255.0 | [optional] 
**Option12** | Pointer to **string** | option12 | [optional] 
**PreDns** | Pointer to **string** | Primary DNS Server, like 127.0.0.1 | [optional] 
**SecDns** | Pointer to **string** | Second DNS Server, link 127.0.0.1 | [optional] 
**ServerMac** | Pointer to **string** | The Mac of DHCP Server for Address reservation when useFixedAddr is enabled. | [optional] 
**ServerStackId** | Pointer to **string** | The Stack ID of DHCP Server for Address reservation when useFixedAddr is enabled. | [optional] 
**ServerType** | Pointer to **string** | The Type of DHCP Server for Address reservation when useFixedAddr is enabled. | [optional] 
**UseFixedAddr** | Pointer to **bool** | The switch of Address reservation for dhcp mode. | [optional] 

## Methods

### NewOswIpSettingOpenApiVO

`func NewOswIpSettingOpenApiVO(mode int32, ) *OswIpSettingOpenApiVO`

NewOswIpSettingOpenApiVO instantiates a new OswIpSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswIpSettingOpenApiVOWithDefaults

`func NewOswIpSettingOpenApiVOWithDefaults() *OswIpSettingOpenApiVO`

NewOswIpSettingOpenApiVOWithDefaults instantiates a new OswIpSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmConflict

`func (o *OswIpSettingOpenApiVO) GetConfirmConflict() bool`

GetConfirmConflict returns the ConfirmConflict field if non-nil, zero value otherwise.

### GetConfirmConflictOk

`func (o *OswIpSettingOpenApiVO) GetConfirmConflictOk() (*bool, bool)`

GetConfirmConflictOk returns a tuple with the ConfirmConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmConflict

`func (o *OswIpSettingOpenApiVO) SetConfirmConflict(v bool)`

SetConfirmConflict sets ConfirmConflict field to given value.

### HasConfirmConflict

`func (o *OswIpSettingOpenApiVO) HasConfirmConflict() bool`

HasConfirmConflict returns a boolean if a field has been set.

### GetDhcpIp

`func (o *OswIpSettingOpenApiVO) GetDhcpIp() string`

GetDhcpIp returns the DhcpIp field if non-nil, zero value otherwise.

### GetDhcpIpOk

`func (o *OswIpSettingOpenApiVO) GetDhcpIpOk() (*string, bool)`

GetDhcpIpOk returns a tuple with the DhcpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpIp

`func (o *OswIpSettingOpenApiVO) SetDhcpIp(v string)`

SetDhcpIp sets DhcpIp field to given value.

### HasDhcpIp

`func (o *OswIpSettingOpenApiVO) HasDhcpIp() bool`

HasDhcpIp returns a boolean if a field has been set.

### GetFallback

`func (o *OswIpSettingOpenApiVO) GetFallback() bool`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *OswIpSettingOpenApiVO) GetFallbackOk() (*bool, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *OswIpSettingOpenApiVO) SetFallback(v bool)`

SetFallback sets Fallback field to given value.

### HasFallback

`func (o *OswIpSettingOpenApiVO) HasFallback() bool`

HasFallback returns a boolean if a field has been set.

### GetFallbackGate

`func (o *OswIpSettingOpenApiVO) GetFallbackGate() string`

GetFallbackGate returns the FallbackGate field if non-nil, zero value otherwise.

### GetFallbackGateOk

`func (o *OswIpSettingOpenApiVO) GetFallbackGateOk() (*string, bool)`

GetFallbackGateOk returns a tuple with the FallbackGate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackGate

`func (o *OswIpSettingOpenApiVO) SetFallbackGate(v string)`

SetFallbackGate sets FallbackGate field to given value.

### HasFallbackGate

`func (o *OswIpSettingOpenApiVO) HasFallbackGate() bool`

HasFallbackGate returns a boolean if a field has been set.

### GetFallbackIp

`func (o *OswIpSettingOpenApiVO) GetFallbackIp() string`

GetFallbackIp returns the FallbackIp field if non-nil, zero value otherwise.

### GetFallbackIpOk

`func (o *OswIpSettingOpenApiVO) GetFallbackIpOk() (*string, bool)`

GetFallbackIpOk returns a tuple with the FallbackIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackIp

`func (o *OswIpSettingOpenApiVO) SetFallbackIp(v string)`

SetFallbackIp sets FallbackIp field to given value.

### HasFallbackIp

`func (o *OswIpSettingOpenApiVO) HasFallbackIp() bool`

HasFallbackIp returns a boolean if a field has been set.

### GetFallbackMask

`func (o *OswIpSettingOpenApiVO) GetFallbackMask() string`

GetFallbackMask returns the FallbackMask field if non-nil, zero value otherwise.

### GetFallbackMaskOk

`func (o *OswIpSettingOpenApiVO) GetFallbackMaskOk() (*string, bool)`

GetFallbackMaskOk returns a tuple with the FallbackMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackMask

`func (o *OswIpSettingOpenApiVO) SetFallbackMask(v string)`

SetFallbackMask sets FallbackMask field to given value.

### HasFallbackMask

`func (o *OswIpSettingOpenApiVO) HasFallbackMask() bool`

HasFallbackMask returns a boolean if a field has been set.

### GetGateway

`func (o *OswIpSettingOpenApiVO) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *OswIpSettingOpenApiVO) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *OswIpSettingOpenApiVO) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *OswIpSettingOpenApiVO) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIp

`func (o *OswIpSettingOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OswIpSettingOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OswIpSettingOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OswIpSettingOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMode

`func (o *OswIpSettingOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OswIpSettingOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OswIpSettingOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetNetId

`func (o *OswIpSettingOpenApiVO) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *OswIpSettingOpenApiVO) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *OswIpSettingOpenApiVO) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *OswIpSettingOpenApiVO) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetNetmask

`func (o *OswIpSettingOpenApiVO) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *OswIpSettingOpenApiVO) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *OswIpSettingOpenApiVO) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *OswIpSettingOpenApiVO) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetOption12

`func (o *OswIpSettingOpenApiVO) GetOption12() string`

GetOption12 returns the Option12 field if non-nil, zero value otherwise.

### GetOption12Ok

`func (o *OswIpSettingOpenApiVO) GetOption12Ok() (*string, bool)`

GetOption12Ok returns a tuple with the Option12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption12

`func (o *OswIpSettingOpenApiVO) SetOption12(v string)`

SetOption12 sets Option12 field to given value.

### HasOption12

`func (o *OswIpSettingOpenApiVO) HasOption12() bool`

HasOption12 returns a boolean if a field has been set.

### GetPreDns

`func (o *OswIpSettingOpenApiVO) GetPreDns() string`

GetPreDns returns the PreDns field if non-nil, zero value otherwise.

### GetPreDnsOk

`func (o *OswIpSettingOpenApiVO) GetPreDnsOk() (*string, bool)`

GetPreDnsOk returns a tuple with the PreDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreDns

`func (o *OswIpSettingOpenApiVO) SetPreDns(v string)`

SetPreDns sets PreDns field to given value.

### HasPreDns

`func (o *OswIpSettingOpenApiVO) HasPreDns() bool`

HasPreDns returns a boolean if a field has been set.

### GetSecDns

`func (o *OswIpSettingOpenApiVO) GetSecDns() string`

GetSecDns returns the SecDns field if non-nil, zero value otherwise.

### GetSecDnsOk

`func (o *OswIpSettingOpenApiVO) GetSecDnsOk() (*string, bool)`

GetSecDnsOk returns a tuple with the SecDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecDns

`func (o *OswIpSettingOpenApiVO) SetSecDns(v string)`

SetSecDns sets SecDns field to given value.

### HasSecDns

`func (o *OswIpSettingOpenApiVO) HasSecDns() bool`

HasSecDns returns a boolean if a field has been set.

### GetServerMac

`func (o *OswIpSettingOpenApiVO) GetServerMac() string`

GetServerMac returns the ServerMac field if non-nil, zero value otherwise.

### GetServerMacOk

`func (o *OswIpSettingOpenApiVO) GetServerMacOk() (*string, bool)`

GetServerMacOk returns a tuple with the ServerMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMac

`func (o *OswIpSettingOpenApiVO) SetServerMac(v string)`

SetServerMac sets ServerMac field to given value.

### HasServerMac

`func (o *OswIpSettingOpenApiVO) HasServerMac() bool`

HasServerMac returns a boolean if a field has been set.

### GetServerStackId

`func (o *OswIpSettingOpenApiVO) GetServerStackId() string`

GetServerStackId returns the ServerStackId field if non-nil, zero value otherwise.

### GetServerStackIdOk

`func (o *OswIpSettingOpenApiVO) GetServerStackIdOk() (*string, bool)`

GetServerStackIdOk returns a tuple with the ServerStackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStackId

`func (o *OswIpSettingOpenApiVO) SetServerStackId(v string)`

SetServerStackId sets ServerStackId field to given value.

### HasServerStackId

`func (o *OswIpSettingOpenApiVO) HasServerStackId() bool`

HasServerStackId returns a boolean if a field has been set.

### GetServerType

`func (o *OswIpSettingOpenApiVO) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *OswIpSettingOpenApiVO) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *OswIpSettingOpenApiVO) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *OswIpSettingOpenApiVO) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetUseFixedAddr

`func (o *OswIpSettingOpenApiVO) GetUseFixedAddr() bool`

GetUseFixedAddr returns the UseFixedAddr field if non-nil, zero value otherwise.

### GetUseFixedAddrOk

`func (o *OswIpSettingOpenApiVO) GetUseFixedAddrOk() (*bool, bool)`

GetUseFixedAddrOk returns a tuple with the UseFixedAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFixedAddr

`func (o *OswIpSettingOpenApiVO) SetUseFixedAddr(v bool)`

SetUseFixedAddr sets UseFixedAddr field to given value.

### HasUseFixedAddr

`func (o *OswIpSettingOpenApiVO) HasUseFixedAddr() bool`

HasUseFixedAddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


