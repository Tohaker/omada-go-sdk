# CheckWanLanStatusOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WanList** | Pointer to [**[]WanList**](WanList.md) | The WAN list of gateway | [optional] 
**AdoptedGateway** | Pointer to **bool** | Indicates whether the gateway is connected | [optional] 
**HasGateway** | Pointer to **bool** | Whether adopt a gateway or enable WAN Settings Overrides. | [optional] 
**IptvPorts** | Pointer to **[]string** | IPTVPorts | [optional] 
**Model** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**NetworkComptent** | Pointer to **int32** | NetworkComptent | [optional] 
**PortName** | Pointer to **map[string]string** | The map of port name, key: \&quot;port ID\&quot;, value: \&quot;port name\&quot; | [optional] 
**PreOsgModel** | Pointer to **int32** | PreOSGModel | [optional] 
**SupportIpv6** | Pointer to **int32** | IIndicates whether the gateway support IPv6 | [optional] 

## Methods

### NewCheckWanLanStatusOpenApiVO

`func NewCheckWanLanStatusOpenApiVO() *CheckWanLanStatusOpenApiVO`

NewCheckWanLanStatusOpenApiVO instantiates a new CheckWanLanStatusOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckWanLanStatusOpenApiVOWithDefaults

`func NewCheckWanLanStatusOpenApiVOWithDefaults() *CheckWanLanStatusOpenApiVO`

NewCheckWanLanStatusOpenApiVOWithDefaults instantiates a new CheckWanLanStatusOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWanList

`func (o *CheckWanLanStatusOpenApiVO) GetWanList() []WanList`

GetWanList returns the WanList field if non-nil, zero value otherwise.

### GetWanListOk

`func (o *CheckWanLanStatusOpenApiVO) GetWanListOk() (*[]WanList, bool)`

GetWanListOk returns a tuple with the WanList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanList

`func (o *CheckWanLanStatusOpenApiVO) SetWanList(v []WanList)`

SetWanList sets WanList field to given value.

### HasWanList

`func (o *CheckWanLanStatusOpenApiVO) HasWanList() bool`

HasWanList returns a boolean if a field has been set.

### GetAdoptedGateway

`func (o *CheckWanLanStatusOpenApiVO) GetAdoptedGateway() bool`

GetAdoptedGateway returns the AdoptedGateway field if non-nil, zero value otherwise.

### GetAdoptedGatewayOk

`func (o *CheckWanLanStatusOpenApiVO) GetAdoptedGatewayOk() (*bool, bool)`

GetAdoptedGatewayOk returns a tuple with the AdoptedGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdoptedGateway

`func (o *CheckWanLanStatusOpenApiVO) SetAdoptedGateway(v bool)`

SetAdoptedGateway sets AdoptedGateway field to given value.

### HasAdoptedGateway

`func (o *CheckWanLanStatusOpenApiVO) HasAdoptedGateway() bool`

HasAdoptedGateway returns a boolean if a field has been set.

### GetHasGateway

`func (o *CheckWanLanStatusOpenApiVO) GetHasGateway() bool`

GetHasGateway returns the HasGateway field if non-nil, zero value otherwise.

### GetHasGatewayOk

`func (o *CheckWanLanStatusOpenApiVO) GetHasGatewayOk() (*bool, bool)`

GetHasGatewayOk returns a tuple with the HasGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGateway

`func (o *CheckWanLanStatusOpenApiVO) SetHasGateway(v bool)`

SetHasGateway sets HasGateway field to given value.

### HasHasGateway

`func (o *CheckWanLanStatusOpenApiVO) HasHasGateway() bool`

HasHasGateway returns a boolean if a field has been set.

### GetIptvPorts

`func (o *CheckWanLanStatusOpenApiVO) GetIptvPorts() []string`

GetIptvPorts returns the IptvPorts field if non-nil, zero value otherwise.

### GetIptvPortsOk

`func (o *CheckWanLanStatusOpenApiVO) GetIptvPortsOk() (*[]string, bool)`

GetIptvPortsOk returns a tuple with the IptvPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIptvPorts

`func (o *CheckWanLanStatusOpenApiVO) SetIptvPorts(v []string)`

SetIptvPorts sets IptvPorts field to given value.

### HasIptvPorts

`func (o *CheckWanLanStatusOpenApiVO) HasIptvPorts() bool`

HasIptvPorts returns a boolean if a field has been set.

### GetModel

`func (o *CheckWanLanStatusOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CheckWanLanStatusOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CheckWanLanStatusOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CheckWanLanStatusOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *CheckWanLanStatusOpenApiVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *CheckWanLanStatusOpenApiVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *CheckWanLanStatusOpenApiVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *CheckWanLanStatusOpenApiVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetNetworkComptent

`func (o *CheckWanLanStatusOpenApiVO) GetNetworkComptent() int32`

GetNetworkComptent returns the NetworkComptent field if non-nil, zero value otherwise.

### GetNetworkComptentOk

`func (o *CheckWanLanStatusOpenApiVO) GetNetworkComptentOk() (*int32, bool)`

GetNetworkComptentOk returns a tuple with the NetworkComptent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComptent

`func (o *CheckWanLanStatusOpenApiVO) SetNetworkComptent(v int32)`

SetNetworkComptent sets NetworkComptent field to given value.

### HasNetworkComptent

`func (o *CheckWanLanStatusOpenApiVO) HasNetworkComptent() bool`

HasNetworkComptent returns a boolean if a field has been set.

### GetPortName

`func (o *CheckWanLanStatusOpenApiVO) GetPortName() map[string]string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *CheckWanLanStatusOpenApiVO) GetPortNameOk() (*map[string]string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *CheckWanLanStatusOpenApiVO) SetPortName(v map[string]string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *CheckWanLanStatusOpenApiVO) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetPreOsgModel

`func (o *CheckWanLanStatusOpenApiVO) GetPreOsgModel() int32`

GetPreOsgModel returns the PreOsgModel field if non-nil, zero value otherwise.

### GetPreOsgModelOk

`func (o *CheckWanLanStatusOpenApiVO) GetPreOsgModelOk() (*int32, bool)`

GetPreOsgModelOk returns a tuple with the PreOsgModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreOsgModel

`func (o *CheckWanLanStatusOpenApiVO) SetPreOsgModel(v int32)`

SetPreOsgModel sets PreOsgModel field to given value.

### HasPreOsgModel

`func (o *CheckWanLanStatusOpenApiVO) HasPreOsgModel() bool`

HasPreOsgModel returns a boolean if a field has been set.

### GetSupportIpv6

`func (o *CheckWanLanStatusOpenApiVO) GetSupportIpv6() int32`

GetSupportIpv6 returns the SupportIpv6 field if non-nil, zero value otherwise.

### GetSupportIpv6Ok

`func (o *CheckWanLanStatusOpenApiVO) GetSupportIpv6Ok() (*int32, bool)`

GetSupportIpv6Ok returns a tuple with the SupportIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIpv6

`func (o *CheckWanLanStatusOpenApiVO) SetSupportIpv6(v int32)`

SetSupportIpv6 sets SupportIpv6 field to given value.

### HasSupportIpv6

`func (o *CheckWanLanStatusOpenApiVO) HasSupportIpv6() bool`

HasSupportIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


