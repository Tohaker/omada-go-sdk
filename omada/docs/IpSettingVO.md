# IpSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternateDNS** | Pointer to **string** | Alternate DNS, only for \&quot;static\&quot; mode | [optional] 
**ConfigGate** | Pointer to **string** | Config Gate | [optional] 
**ConfigIp** | Pointer to **string** | Config IP | [optional] 
**ConfigMask** | Pointer to **string** | Config Mask | [optional] 
**ConfirmConflict** | Pointer to **bool** | Confirm Conflict | [optional] 
**DhcpIp** | Pointer to **string** | DHCP IP, only for \&quot;dhcp\&quot; mode | [optional] 
**Fallback** | Pointer to **bool** | Fallback, only for \&quot;dhcp\&quot; mode | [optional] 
**FallbackGate** | Pointer to **string** | Fallback Gate, only for \&quot;dhcp\&quot; mode | [optional] 
**FallbackIp** | Pointer to **string** | Fallback IP, only for \&quot;dhcp\&quot; mode | [optional] 
**FallbackMask** | Pointer to **string** | Fallback Mask, only for \&quot;dhcp\&quot; mode | [optional] 
**Mode** | **string** | IpSetting parameter [mode] should be \&quot;static\&quot; or \&quot;dhcp\&quot; | 
**NetId** | Pointer to **string** | Net ID, only for \&quot;dhcp\&quot; mode | [optional] 
**PreferredDNS** | Pointer to **string** | Preferred DNS, only for \&quot;static\&quot; mode | [optional] 
**ServerMac** | Pointer to **string** | DHCP Server Mac, only for \&quot;dhcp\&quot; mode | [optional] 
**ServerStackId** | Pointer to **string** | DHCP Server Stack ID, only for \&quot;dhcp\&quot; mode | [optional] 
**ServerType** | Pointer to **string** | DHCP Server Type, only for \&quot;dhcp\&quot; mode | [optional] 
**UseFoxedAddr** | Pointer to **bool** | Use Fixed Address, only for \&quot;dhcp\&quot; mode | [optional] 

## Methods

### NewIpSettingVO

`func NewIpSettingVO(mode string, ) *IpSettingVO`

NewIpSettingVO instantiates a new IpSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpSettingVOWithDefaults

`func NewIpSettingVOWithDefaults() *IpSettingVO`

NewIpSettingVOWithDefaults instantiates a new IpSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateDNS

`func (o *IpSettingVO) GetAlternateDNS() string`

GetAlternateDNS returns the AlternateDNS field if non-nil, zero value otherwise.

### GetAlternateDNSOk

`func (o *IpSettingVO) GetAlternateDNSOk() (*string, bool)`

GetAlternateDNSOk returns a tuple with the AlternateDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateDNS

`func (o *IpSettingVO) SetAlternateDNS(v string)`

SetAlternateDNS sets AlternateDNS field to given value.

### HasAlternateDNS

`func (o *IpSettingVO) HasAlternateDNS() bool`

HasAlternateDNS returns a boolean if a field has been set.

### GetConfigGate

`func (o *IpSettingVO) GetConfigGate() string`

GetConfigGate returns the ConfigGate field if non-nil, zero value otherwise.

### GetConfigGateOk

`func (o *IpSettingVO) GetConfigGateOk() (*string, bool)`

GetConfigGateOk returns a tuple with the ConfigGate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigGate

`func (o *IpSettingVO) SetConfigGate(v string)`

SetConfigGate sets ConfigGate field to given value.

### HasConfigGate

`func (o *IpSettingVO) HasConfigGate() bool`

HasConfigGate returns a boolean if a field has been set.

### GetConfigIp

`func (o *IpSettingVO) GetConfigIp() string`

GetConfigIp returns the ConfigIp field if non-nil, zero value otherwise.

### GetConfigIpOk

`func (o *IpSettingVO) GetConfigIpOk() (*string, bool)`

GetConfigIpOk returns a tuple with the ConfigIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigIp

`func (o *IpSettingVO) SetConfigIp(v string)`

SetConfigIp sets ConfigIp field to given value.

### HasConfigIp

`func (o *IpSettingVO) HasConfigIp() bool`

HasConfigIp returns a boolean if a field has been set.

### GetConfigMask

`func (o *IpSettingVO) GetConfigMask() string`

GetConfigMask returns the ConfigMask field if non-nil, zero value otherwise.

### GetConfigMaskOk

`func (o *IpSettingVO) GetConfigMaskOk() (*string, bool)`

GetConfigMaskOk returns a tuple with the ConfigMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMask

`func (o *IpSettingVO) SetConfigMask(v string)`

SetConfigMask sets ConfigMask field to given value.

### HasConfigMask

`func (o *IpSettingVO) HasConfigMask() bool`

HasConfigMask returns a boolean if a field has been set.

### GetConfirmConflict

`func (o *IpSettingVO) GetConfirmConflict() bool`

GetConfirmConflict returns the ConfirmConflict field if non-nil, zero value otherwise.

### GetConfirmConflictOk

`func (o *IpSettingVO) GetConfirmConflictOk() (*bool, bool)`

GetConfirmConflictOk returns a tuple with the ConfirmConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmConflict

`func (o *IpSettingVO) SetConfirmConflict(v bool)`

SetConfirmConflict sets ConfirmConflict field to given value.

### HasConfirmConflict

`func (o *IpSettingVO) HasConfirmConflict() bool`

HasConfirmConflict returns a boolean if a field has been set.

### GetDhcpIp

`func (o *IpSettingVO) GetDhcpIp() string`

GetDhcpIp returns the DhcpIp field if non-nil, zero value otherwise.

### GetDhcpIpOk

`func (o *IpSettingVO) GetDhcpIpOk() (*string, bool)`

GetDhcpIpOk returns a tuple with the DhcpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpIp

`func (o *IpSettingVO) SetDhcpIp(v string)`

SetDhcpIp sets DhcpIp field to given value.

### HasDhcpIp

`func (o *IpSettingVO) HasDhcpIp() bool`

HasDhcpIp returns a boolean if a field has been set.

### GetFallback

`func (o *IpSettingVO) GetFallback() bool`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *IpSettingVO) GetFallbackOk() (*bool, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *IpSettingVO) SetFallback(v bool)`

SetFallback sets Fallback field to given value.

### HasFallback

`func (o *IpSettingVO) HasFallback() bool`

HasFallback returns a boolean if a field has been set.

### GetFallbackGate

`func (o *IpSettingVO) GetFallbackGate() string`

GetFallbackGate returns the FallbackGate field if non-nil, zero value otherwise.

### GetFallbackGateOk

`func (o *IpSettingVO) GetFallbackGateOk() (*string, bool)`

GetFallbackGateOk returns a tuple with the FallbackGate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackGate

`func (o *IpSettingVO) SetFallbackGate(v string)`

SetFallbackGate sets FallbackGate field to given value.

### HasFallbackGate

`func (o *IpSettingVO) HasFallbackGate() bool`

HasFallbackGate returns a boolean if a field has been set.

### GetFallbackIp

`func (o *IpSettingVO) GetFallbackIp() string`

GetFallbackIp returns the FallbackIp field if non-nil, zero value otherwise.

### GetFallbackIpOk

`func (o *IpSettingVO) GetFallbackIpOk() (*string, bool)`

GetFallbackIpOk returns a tuple with the FallbackIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackIp

`func (o *IpSettingVO) SetFallbackIp(v string)`

SetFallbackIp sets FallbackIp field to given value.

### HasFallbackIp

`func (o *IpSettingVO) HasFallbackIp() bool`

HasFallbackIp returns a boolean if a field has been set.

### GetFallbackMask

`func (o *IpSettingVO) GetFallbackMask() string`

GetFallbackMask returns the FallbackMask field if non-nil, zero value otherwise.

### GetFallbackMaskOk

`func (o *IpSettingVO) GetFallbackMaskOk() (*string, bool)`

GetFallbackMaskOk returns a tuple with the FallbackMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackMask

`func (o *IpSettingVO) SetFallbackMask(v string)`

SetFallbackMask sets FallbackMask field to given value.

### HasFallbackMask

`func (o *IpSettingVO) HasFallbackMask() bool`

HasFallbackMask returns a boolean if a field has been set.

### GetMode

`func (o *IpSettingVO) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *IpSettingVO) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *IpSettingVO) SetMode(v string)`

SetMode sets Mode field to given value.


### GetNetId

`func (o *IpSettingVO) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *IpSettingVO) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *IpSettingVO) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *IpSettingVO) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetPreferredDNS

`func (o *IpSettingVO) GetPreferredDNS() string`

GetPreferredDNS returns the PreferredDNS field if non-nil, zero value otherwise.

### GetPreferredDNSOk

`func (o *IpSettingVO) GetPreferredDNSOk() (*string, bool)`

GetPreferredDNSOk returns a tuple with the PreferredDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDNS

`func (o *IpSettingVO) SetPreferredDNS(v string)`

SetPreferredDNS sets PreferredDNS field to given value.

### HasPreferredDNS

`func (o *IpSettingVO) HasPreferredDNS() bool`

HasPreferredDNS returns a boolean if a field has been set.

### GetServerMac

`func (o *IpSettingVO) GetServerMac() string`

GetServerMac returns the ServerMac field if non-nil, zero value otherwise.

### GetServerMacOk

`func (o *IpSettingVO) GetServerMacOk() (*string, bool)`

GetServerMacOk returns a tuple with the ServerMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMac

`func (o *IpSettingVO) SetServerMac(v string)`

SetServerMac sets ServerMac field to given value.

### HasServerMac

`func (o *IpSettingVO) HasServerMac() bool`

HasServerMac returns a boolean if a field has been set.

### GetServerStackId

`func (o *IpSettingVO) GetServerStackId() string`

GetServerStackId returns the ServerStackId field if non-nil, zero value otherwise.

### GetServerStackIdOk

`func (o *IpSettingVO) GetServerStackIdOk() (*string, bool)`

GetServerStackIdOk returns a tuple with the ServerStackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStackId

`func (o *IpSettingVO) SetServerStackId(v string)`

SetServerStackId sets ServerStackId field to given value.

### HasServerStackId

`func (o *IpSettingVO) HasServerStackId() bool`

HasServerStackId returns a boolean if a field has been set.

### GetServerType

`func (o *IpSettingVO) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *IpSettingVO) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *IpSettingVO) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *IpSettingVO) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetUseFoxedAddr

`func (o *IpSettingVO) GetUseFoxedAddr() bool`

GetUseFoxedAddr returns the UseFoxedAddr field if non-nil, zero value otherwise.

### GetUseFoxedAddrOk

`func (o *IpSettingVO) GetUseFoxedAddrOk() (*bool, bool)`

GetUseFoxedAddrOk returns a tuple with the UseFoxedAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFoxedAddr

`func (o *IpSettingVO) SetUseFoxedAddr(v bool)`

SetUseFoxedAddr sets UseFoxedAddr field to given value.

### HasUseFoxedAddr

`func (o *IpSettingVO) HasUseFoxedAddr() bool`

HasUseFoxedAddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


