# OswIpSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmConflict** | Pointer to **bool** | The switch to confirm conflict | [optional] 
**DhcpIp** | Pointer to **string** | DHCP IP | [optional] 
**Fallback** | Pointer to **bool** | The switch of fallback | [optional] 
**FallbackGate** | Pointer to **string** | Fallback Gateway | [optional] 
**FallbackIp** | Pointer to **string** | Fallback IP | [optional] 
**FallbackMask** | Pointer to **string** | Fallback Mask | [optional] 
**Gateway** | Pointer to **string** | Gateway, like 192.168.0.1 | [optional] 
**Ip** | Pointer to **string** | Static IP for mode 0, like 192.168.0.1 | [optional] 
**Mode** | **int32** | IP Setting mode. Static:0, DHCP:1 | 
**NetId** | Pointer to **string** | netId | [optional] 
**Netmask** | Pointer to **string** | IP Mask, like 255.255.255.0 | [optional] 
**Option12** | Pointer to **string** | option12 | [optional] 
**PreDns** | Pointer to **string** | primary DNS Server | [optional] 
**SecDns** | Pointer to **string** | second DNS Server | [optional] 
**ServerMac** | Pointer to **string** | Server MAC | [optional] 
**ServerStackId** | Pointer to **string** | serverStackId | [optional] 
**ServerType** | Pointer to **string** | Server Type | [optional] 
**UserFixedAddr** | Pointer to **bool** | The switch of Fixed Address | [optional] 

## Methods

### NewOswIpSettingVO

`func NewOswIpSettingVO(mode int32, ) *OswIpSettingVO`

NewOswIpSettingVO instantiates a new OswIpSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswIpSettingVOWithDefaults

`func NewOswIpSettingVOWithDefaults() *OswIpSettingVO`

NewOswIpSettingVOWithDefaults instantiates a new OswIpSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmConflict

`func (o *OswIpSettingVO) GetConfirmConflict() bool`

GetConfirmConflict returns the ConfirmConflict field if non-nil, zero value otherwise.

### GetConfirmConflictOk

`func (o *OswIpSettingVO) GetConfirmConflictOk() (*bool, bool)`

GetConfirmConflictOk returns a tuple with the ConfirmConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmConflict

`func (o *OswIpSettingVO) SetConfirmConflict(v bool)`

SetConfirmConflict sets ConfirmConflict field to given value.

### HasConfirmConflict

`func (o *OswIpSettingVO) HasConfirmConflict() bool`

HasConfirmConflict returns a boolean if a field has been set.

### GetDhcpIp

`func (o *OswIpSettingVO) GetDhcpIp() string`

GetDhcpIp returns the DhcpIp field if non-nil, zero value otherwise.

### GetDhcpIpOk

`func (o *OswIpSettingVO) GetDhcpIpOk() (*string, bool)`

GetDhcpIpOk returns a tuple with the DhcpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpIp

`func (o *OswIpSettingVO) SetDhcpIp(v string)`

SetDhcpIp sets DhcpIp field to given value.

### HasDhcpIp

`func (o *OswIpSettingVO) HasDhcpIp() bool`

HasDhcpIp returns a boolean if a field has been set.

### GetFallback

`func (o *OswIpSettingVO) GetFallback() bool`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *OswIpSettingVO) GetFallbackOk() (*bool, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *OswIpSettingVO) SetFallback(v bool)`

SetFallback sets Fallback field to given value.

### HasFallback

`func (o *OswIpSettingVO) HasFallback() bool`

HasFallback returns a boolean if a field has been set.

### GetFallbackGate

`func (o *OswIpSettingVO) GetFallbackGate() string`

GetFallbackGate returns the FallbackGate field if non-nil, zero value otherwise.

### GetFallbackGateOk

`func (o *OswIpSettingVO) GetFallbackGateOk() (*string, bool)`

GetFallbackGateOk returns a tuple with the FallbackGate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackGate

`func (o *OswIpSettingVO) SetFallbackGate(v string)`

SetFallbackGate sets FallbackGate field to given value.

### HasFallbackGate

`func (o *OswIpSettingVO) HasFallbackGate() bool`

HasFallbackGate returns a boolean if a field has been set.

### GetFallbackIp

`func (o *OswIpSettingVO) GetFallbackIp() string`

GetFallbackIp returns the FallbackIp field if non-nil, zero value otherwise.

### GetFallbackIpOk

`func (o *OswIpSettingVO) GetFallbackIpOk() (*string, bool)`

GetFallbackIpOk returns a tuple with the FallbackIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackIp

`func (o *OswIpSettingVO) SetFallbackIp(v string)`

SetFallbackIp sets FallbackIp field to given value.

### HasFallbackIp

`func (o *OswIpSettingVO) HasFallbackIp() bool`

HasFallbackIp returns a boolean if a field has been set.

### GetFallbackMask

`func (o *OswIpSettingVO) GetFallbackMask() string`

GetFallbackMask returns the FallbackMask field if non-nil, zero value otherwise.

### GetFallbackMaskOk

`func (o *OswIpSettingVO) GetFallbackMaskOk() (*string, bool)`

GetFallbackMaskOk returns a tuple with the FallbackMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackMask

`func (o *OswIpSettingVO) SetFallbackMask(v string)`

SetFallbackMask sets FallbackMask field to given value.

### HasFallbackMask

`func (o *OswIpSettingVO) HasFallbackMask() bool`

HasFallbackMask returns a boolean if a field has been set.

### GetGateway

`func (o *OswIpSettingVO) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *OswIpSettingVO) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *OswIpSettingVO) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *OswIpSettingVO) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIp

`func (o *OswIpSettingVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OswIpSettingVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OswIpSettingVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OswIpSettingVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMode

`func (o *OswIpSettingVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OswIpSettingVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OswIpSettingVO) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetNetId

`func (o *OswIpSettingVO) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *OswIpSettingVO) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *OswIpSettingVO) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *OswIpSettingVO) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetNetmask

`func (o *OswIpSettingVO) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *OswIpSettingVO) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *OswIpSettingVO) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *OswIpSettingVO) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetOption12

`func (o *OswIpSettingVO) GetOption12() string`

GetOption12 returns the Option12 field if non-nil, zero value otherwise.

### GetOption12Ok

`func (o *OswIpSettingVO) GetOption12Ok() (*string, bool)`

GetOption12Ok returns a tuple with the Option12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption12

`func (o *OswIpSettingVO) SetOption12(v string)`

SetOption12 sets Option12 field to given value.

### HasOption12

`func (o *OswIpSettingVO) HasOption12() bool`

HasOption12 returns a boolean if a field has been set.

### GetPreDns

`func (o *OswIpSettingVO) GetPreDns() string`

GetPreDns returns the PreDns field if non-nil, zero value otherwise.

### GetPreDnsOk

`func (o *OswIpSettingVO) GetPreDnsOk() (*string, bool)`

GetPreDnsOk returns a tuple with the PreDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreDns

`func (o *OswIpSettingVO) SetPreDns(v string)`

SetPreDns sets PreDns field to given value.

### HasPreDns

`func (o *OswIpSettingVO) HasPreDns() bool`

HasPreDns returns a boolean if a field has been set.

### GetSecDns

`func (o *OswIpSettingVO) GetSecDns() string`

GetSecDns returns the SecDns field if non-nil, zero value otherwise.

### GetSecDnsOk

`func (o *OswIpSettingVO) GetSecDnsOk() (*string, bool)`

GetSecDnsOk returns a tuple with the SecDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecDns

`func (o *OswIpSettingVO) SetSecDns(v string)`

SetSecDns sets SecDns field to given value.

### HasSecDns

`func (o *OswIpSettingVO) HasSecDns() bool`

HasSecDns returns a boolean if a field has been set.

### GetServerMac

`func (o *OswIpSettingVO) GetServerMac() string`

GetServerMac returns the ServerMac field if non-nil, zero value otherwise.

### GetServerMacOk

`func (o *OswIpSettingVO) GetServerMacOk() (*string, bool)`

GetServerMacOk returns a tuple with the ServerMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMac

`func (o *OswIpSettingVO) SetServerMac(v string)`

SetServerMac sets ServerMac field to given value.

### HasServerMac

`func (o *OswIpSettingVO) HasServerMac() bool`

HasServerMac returns a boolean if a field has been set.

### GetServerStackId

`func (o *OswIpSettingVO) GetServerStackId() string`

GetServerStackId returns the ServerStackId field if non-nil, zero value otherwise.

### GetServerStackIdOk

`func (o *OswIpSettingVO) GetServerStackIdOk() (*string, bool)`

GetServerStackIdOk returns a tuple with the ServerStackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStackId

`func (o *OswIpSettingVO) SetServerStackId(v string)`

SetServerStackId sets ServerStackId field to given value.

### HasServerStackId

`func (o *OswIpSettingVO) HasServerStackId() bool`

HasServerStackId returns a boolean if a field has been set.

### GetServerType

`func (o *OswIpSettingVO) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *OswIpSettingVO) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *OswIpSettingVO) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *OswIpSettingVO) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetUserFixedAddr

`func (o *OswIpSettingVO) GetUserFixedAddr() bool`

GetUserFixedAddr returns the UserFixedAddr field if non-nil, zero value otherwise.

### GetUserFixedAddrOk

`func (o *OswIpSettingVO) GetUserFixedAddrOk() (*bool, bool)`

GetUserFixedAddrOk returns a tuple with the UserFixedAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFixedAddr

`func (o *OswIpSettingVO) SetUserFixedAddr(v bool)`

SetUserFixedAddr sets UserFixedAddr field to given value.

### HasUserFixedAddr

`func (o *OswIpSettingVO) HasUserFixedAddr() bool`

HasUserFixedAddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


