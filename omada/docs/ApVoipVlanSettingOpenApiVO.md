# ApVoipVlanSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | VoIP VLAN IP. Parameter [ip] should not be null when [ipType] is 0. | [optional] 
**IpDns1** | Pointer to **string** | IP DNS1. | [optional] 
**IpDns2** | Pointer to **string** | IP DNS2. | [optional] 
**IpGateway** | Pointer to **string** | VoIP VLAN IP gateway. Parameter [ipGateway] should not be null when [ipType] is 0. | [optional] 
**IpMask** | Pointer to **string** | VoIP VLAN IP mask. Parameter [ipMask] should not be null when [ipType] is 0. | [optional] 
**IpType** | Pointer to **int32** | VoIP VLAN IP Type. 0: Static IP. 1: DHCP. | [optional] 
**LanNetworkId** | Pointer to **string** | LAN network ID. Parameter [lanNetworkId] should not be null, and should not be the same as the lanNetworkId of management VLAN when parameter [mode] is 1. | [optional] 
**Mode** | **int32** | VoIP VLAN mode. 0 : Follow Management VLAN. 1: Custom. | 
**VoipBridgeVlan** | Pointer to **int32** | VoIP bridge vlan. Parameter [voipBridgeVlan] should not be null and be between 1 and 4090 when parameter [lanNetworkId] corresponds to a multiple VLAN LAN network. | [optional] 

## Methods

### NewApVoipVlanSettingOpenApiVO

`func NewApVoipVlanSettingOpenApiVO(mode int32, ) *ApVoipVlanSettingOpenApiVO`

NewApVoipVlanSettingOpenApiVO instantiates a new ApVoipVlanSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApVoipVlanSettingOpenApiVOWithDefaults

`func NewApVoipVlanSettingOpenApiVOWithDefaults() *ApVoipVlanSettingOpenApiVO`

NewApVoipVlanSettingOpenApiVOWithDefaults instantiates a new ApVoipVlanSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *ApVoipVlanSettingOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ApVoipVlanSettingOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ApVoipVlanSettingOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ApVoipVlanSettingOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpDns1

`func (o *ApVoipVlanSettingOpenApiVO) GetIpDns1() string`

GetIpDns1 returns the IpDns1 field if non-nil, zero value otherwise.

### GetIpDns1Ok

`func (o *ApVoipVlanSettingOpenApiVO) GetIpDns1Ok() (*string, bool)`

GetIpDns1Ok returns a tuple with the IpDns1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDns1

`func (o *ApVoipVlanSettingOpenApiVO) SetIpDns1(v string)`

SetIpDns1 sets IpDns1 field to given value.

### HasIpDns1

`func (o *ApVoipVlanSettingOpenApiVO) HasIpDns1() bool`

HasIpDns1 returns a boolean if a field has been set.

### GetIpDns2

`func (o *ApVoipVlanSettingOpenApiVO) GetIpDns2() string`

GetIpDns2 returns the IpDns2 field if non-nil, zero value otherwise.

### GetIpDns2Ok

`func (o *ApVoipVlanSettingOpenApiVO) GetIpDns2Ok() (*string, bool)`

GetIpDns2Ok returns a tuple with the IpDns2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDns2

`func (o *ApVoipVlanSettingOpenApiVO) SetIpDns2(v string)`

SetIpDns2 sets IpDns2 field to given value.

### HasIpDns2

`func (o *ApVoipVlanSettingOpenApiVO) HasIpDns2() bool`

HasIpDns2 returns a boolean if a field has been set.

### GetIpGateway

`func (o *ApVoipVlanSettingOpenApiVO) GetIpGateway() string`

GetIpGateway returns the IpGateway field if non-nil, zero value otherwise.

### GetIpGatewayOk

`func (o *ApVoipVlanSettingOpenApiVO) GetIpGatewayOk() (*string, bool)`

GetIpGatewayOk returns a tuple with the IpGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpGateway

`func (o *ApVoipVlanSettingOpenApiVO) SetIpGateway(v string)`

SetIpGateway sets IpGateway field to given value.

### HasIpGateway

`func (o *ApVoipVlanSettingOpenApiVO) HasIpGateway() bool`

HasIpGateway returns a boolean if a field has been set.

### GetIpMask

`func (o *ApVoipVlanSettingOpenApiVO) GetIpMask() string`

GetIpMask returns the IpMask field if non-nil, zero value otherwise.

### GetIpMaskOk

`func (o *ApVoipVlanSettingOpenApiVO) GetIpMaskOk() (*string, bool)`

GetIpMaskOk returns a tuple with the IpMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMask

`func (o *ApVoipVlanSettingOpenApiVO) SetIpMask(v string)`

SetIpMask sets IpMask field to given value.

### HasIpMask

`func (o *ApVoipVlanSettingOpenApiVO) HasIpMask() bool`

HasIpMask returns a boolean if a field has been set.

### GetIpType

`func (o *ApVoipVlanSettingOpenApiVO) GetIpType() int32`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *ApVoipVlanSettingOpenApiVO) GetIpTypeOk() (*int32, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *ApVoipVlanSettingOpenApiVO) SetIpType(v int32)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *ApVoipVlanSettingOpenApiVO) HasIpType() bool`

HasIpType returns a boolean if a field has been set.

### GetLanNetworkId

`func (o *ApVoipVlanSettingOpenApiVO) GetLanNetworkId() string`

GetLanNetworkId returns the LanNetworkId field if non-nil, zero value otherwise.

### GetLanNetworkIdOk

`func (o *ApVoipVlanSettingOpenApiVO) GetLanNetworkIdOk() (*string, bool)`

GetLanNetworkIdOk returns a tuple with the LanNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetworkId

`func (o *ApVoipVlanSettingOpenApiVO) SetLanNetworkId(v string)`

SetLanNetworkId sets LanNetworkId field to given value.

### HasLanNetworkId

`func (o *ApVoipVlanSettingOpenApiVO) HasLanNetworkId() bool`

HasLanNetworkId returns a boolean if a field has been set.

### GetMode

`func (o *ApVoipVlanSettingOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ApVoipVlanSettingOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ApVoipVlanSettingOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetVoipBridgeVlan

`func (o *ApVoipVlanSettingOpenApiVO) GetVoipBridgeVlan() int32`

GetVoipBridgeVlan returns the VoipBridgeVlan field if non-nil, zero value otherwise.

### GetVoipBridgeVlanOk

`func (o *ApVoipVlanSettingOpenApiVO) GetVoipBridgeVlanOk() (*int32, bool)`

GetVoipBridgeVlanOk returns a tuple with the VoipBridgeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoipBridgeVlan

`func (o *ApVoipVlanSettingOpenApiVO) SetVoipBridgeVlan(v int32)`

SetVoipBridgeVlan sets VoipBridgeVlan field to given value.

### HasVoipBridgeVlan

`func (o *ApVoipVlanSettingOpenApiVO) HasVoipBridgeVlan() bool`

HasVoipBridgeVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


