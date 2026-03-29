# StaticIpSettingEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternateDNS** | Pointer to **string** | Secondary DNS server | [optional] 
**ConfigGate** | Pointer to **string** | Gateway IP address | [optional] 
**ConfigIp** | Pointer to **string** | Static IP address | [optional] 
**ConfigMask** | Pointer to **string** | Subnet mask | [optional] 
**PreferredDNS** | Pointer to **string** | Primary DNS server | [optional] 

## Methods

### NewStaticIpSettingEntity

`func NewStaticIpSettingEntity() *StaticIpSettingEntity`

NewStaticIpSettingEntity instantiates a new StaticIpSettingEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticIpSettingEntityWithDefaults

`func NewStaticIpSettingEntityWithDefaults() *StaticIpSettingEntity`

NewStaticIpSettingEntityWithDefaults instantiates a new StaticIpSettingEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateDNS

`func (o *StaticIpSettingEntity) GetAlternateDNS() string`

GetAlternateDNS returns the AlternateDNS field if non-nil, zero value otherwise.

### GetAlternateDNSOk

`func (o *StaticIpSettingEntity) GetAlternateDNSOk() (*string, bool)`

GetAlternateDNSOk returns a tuple with the AlternateDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateDNS

`func (o *StaticIpSettingEntity) SetAlternateDNS(v string)`

SetAlternateDNS sets AlternateDNS field to given value.

### HasAlternateDNS

`func (o *StaticIpSettingEntity) HasAlternateDNS() bool`

HasAlternateDNS returns a boolean if a field has been set.

### GetConfigGate

`func (o *StaticIpSettingEntity) GetConfigGate() string`

GetConfigGate returns the ConfigGate field if non-nil, zero value otherwise.

### GetConfigGateOk

`func (o *StaticIpSettingEntity) GetConfigGateOk() (*string, bool)`

GetConfigGateOk returns a tuple with the ConfigGate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigGate

`func (o *StaticIpSettingEntity) SetConfigGate(v string)`

SetConfigGate sets ConfigGate field to given value.

### HasConfigGate

`func (o *StaticIpSettingEntity) HasConfigGate() bool`

HasConfigGate returns a boolean if a field has been set.

### GetConfigIp

`func (o *StaticIpSettingEntity) GetConfigIp() string`

GetConfigIp returns the ConfigIp field if non-nil, zero value otherwise.

### GetConfigIpOk

`func (o *StaticIpSettingEntity) GetConfigIpOk() (*string, bool)`

GetConfigIpOk returns a tuple with the ConfigIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigIp

`func (o *StaticIpSettingEntity) SetConfigIp(v string)`

SetConfigIp sets ConfigIp field to given value.

### HasConfigIp

`func (o *StaticIpSettingEntity) HasConfigIp() bool`

HasConfigIp returns a boolean if a field has been set.

### GetConfigMask

`func (o *StaticIpSettingEntity) GetConfigMask() string`

GetConfigMask returns the ConfigMask field if non-nil, zero value otherwise.

### GetConfigMaskOk

`func (o *StaticIpSettingEntity) GetConfigMaskOk() (*string, bool)`

GetConfigMaskOk returns a tuple with the ConfigMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMask

`func (o *StaticIpSettingEntity) SetConfigMask(v string)`

SetConfigMask sets ConfigMask field to given value.

### HasConfigMask

`func (o *StaticIpSettingEntity) HasConfigMask() bool`

HasConfigMask returns a boolean if a field has been set.

### GetPreferredDNS

`func (o *StaticIpSettingEntity) GetPreferredDNS() string`

GetPreferredDNS returns the PreferredDNS field if non-nil, zero value otherwise.

### GetPreferredDNSOk

`func (o *StaticIpSettingEntity) GetPreferredDNSOk() (*string, bool)`

GetPreferredDNSOk returns a tuple with the PreferredDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDNS

`func (o *StaticIpSettingEntity) SetPreferredDNS(v string)`

SetPreferredDNS sets PreferredDNS field to given value.

### HasPreferredDNS

`func (o *StaticIpSettingEntity) HasPreferredDNS() bool`

HasPreferredDNS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


