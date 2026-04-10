# CheckWanLanStatusVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WanList** | Pointer to [**[]OsgWanStatusVO**](OsgWanStatusVO.md) |  | [optional] 
**AdoptedGateway** | Pointer to **bool** |  | [optional] 
**HasGateway** | Pointer to **bool** |  | [optional] 
**IptvPorts** | Pointer to **[]string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**ModelVersion** | Pointer to **string** |  | [optional] 
**NetworkComptent** | Pointer to **int32** |  | [optional] 
**OsgAddInAdvance** | Pointer to **bool** |  | [optional] 
**PortName** | Pointer to **map[string]string** |  | [optional] 
**PreOsgModel** | Pointer to **int32** |  | [optional] 
**SupportIpv6** | Pointer to **int32** |  | [optional] 
**WirelessRouter** | Pointer to **bool** |  | [optional] 

## Methods

### NewCheckWanLanStatusVO

`func NewCheckWanLanStatusVO() *CheckWanLanStatusVO`

NewCheckWanLanStatusVO instantiates a new CheckWanLanStatusVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckWanLanStatusVOWithDefaults

`func NewCheckWanLanStatusVOWithDefaults() *CheckWanLanStatusVO`

NewCheckWanLanStatusVOWithDefaults instantiates a new CheckWanLanStatusVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWanList

`func (o *CheckWanLanStatusVO) GetWanList() []OsgWanStatusVO`

GetWanList returns the WanList field if non-nil, zero value otherwise.

### GetWanListOk

`func (o *CheckWanLanStatusVO) GetWanListOk() (*[]OsgWanStatusVO, bool)`

GetWanListOk returns a tuple with the WanList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanList

`func (o *CheckWanLanStatusVO) SetWanList(v []OsgWanStatusVO)`

SetWanList sets WanList field to given value.

### HasWanList

`func (o *CheckWanLanStatusVO) HasWanList() bool`

HasWanList returns a boolean if a field has been set.

### GetAdoptedGateway

`func (o *CheckWanLanStatusVO) GetAdoptedGateway() bool`

GetAdoptedGateway returns the AdoptedGateway field if non-nil, zero value otherwise.

### GetAdoptedGatewayOk

`func (o *CheckWanLanStatusVO) GetAdoptedGatewayOk() (*bool, bool)`

GetAdoptedGatewayOk returns a tuple with the AdoptedGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdoptedGateway

`func (o *CheckWanLanStatusVO) SetAdoptedGateway(v bool)`

SetAdoptedGateway sets AdoptedGateway field to given value.

### HasAdoptedGateway

`func (o *CheckWanLanStatusVO) HasAdoptedGateway() bool`

HasAdoptedGateway returns a boolean if a field has been set.

### GetHasGateway

`func (o *CheckWanLanStatusVO) GetHasGateway() bool`

GetHasGateway returns the HasGateway field if non-nil, zero value otherwise.

### GetHasGatewayOk

`func (o *CheckWanLanStatusVO) GetHasGatewayOk() (*bool, bool)`

GetHasGatewayOk returns a tuple with the HasGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGateway

`func (o *CheckWanLanStatusVO) SetHasGateway(v bool)`

SetHasGateway sets HasGateway field to given value.

### HasHasGateway

`func (o *CheckWanLanStatusVO) HasHasGateway() bool`

HasHasGateway returns a boolean if a field has been set.

### GetIptvPorts

`func (o *CheckWanLanStatusVO) GetIptvPorts() []string`

GetIptvPorts returns the IptvPorts field if non-nil, zero value otherwise.

### GetIptvPortsOk

`func (o *CheckWanLanStatusVO) GetIptvPortsOk() (*[]string, bool)`

GetIptvPortsOk returns a tuple with the IptvPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIptvPorts

`func (o *CheckWanLanStatusVO) SetIptvPorts(v []string)`

SetIptvPorts sets IptvPorts field to given value.

### HasIptvPorts

`func (o *CheckWanLanStatusVO) HasIptvPorts() bool`

HasIptvPorts returns a boolean if a field has been set.

### GetModel

`func (o *CheckWanLanStatusVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CheckWanLanStatusVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CheckWanLanStatusVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CheckWanLanStatusVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *CheckWanLanStatusVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *CheckWanLanStatusVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *CheckWanLanStatusVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *CheckWanLanStatusVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetNetworkComptent

`func (o *CheckWanLanStatusVO) GetNetworkComptent() int32`

GetNetworkComptent returns the NetworkComptent field if non-nil, zero value otherwise.

### GetNetworkComptentOk

`func (o *CheckWanLanStatusVO) GetNetworkComptentOk() (*int32, bool)`

GetNetworkComptentOk returns a tuple with the NetworkComptent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComptent

`func (o *CheckWanLanStatusVO) SetNetworkComptent(v int32)`

SetNetworkComptent sets NetworkComptent field to given value.

### HasNetworkComptent

`func (o *CheckWanLanStatusVO) HasNetworkComptent() bool`

HasNetworkComptent returns a boolean if a field has been set.

### GetOsgAddInAdvance

`func (o *CheckWanLanStatusVO) GetOsgAddInAdvance() bool`

GetOsgAddInAdvance returns the OsgAddInAdvance field if non-nil, zero value otherwise.

### GetOsgAddInAdvanceOk

`func (o *CheckWanLanStatusVO) GetOsgAddInAdvanceOk() (*bool, bool)`

GetOsgAddInAdvanceOk returns a tuple with the OsgAddInAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsgAddInAdvance

`func (o *CheckWanLanStatusVO) SetOsgAddInAdvance(v bool)`

SetOsgAddInAdvance sets OsgAddInAdvance field to given value.

### HasOsgAddInAdvance

`func (o *CheckWanLanStatusVO) HasOsgAddInAdvance() bool`

HasOsgAddInAdvance returns a boolean if a field has been set.

### GetPortName

`func (o *CheckWanLanStatusVO) GetPortName() map[string]string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *CheckWanLanStatusVO) GetPortNameOk() (*map[string]string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *CheckWanLanStatusVO) SetPortName(v map[string]string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *CheckWanLanStatusVO) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetPreOsgModel

`func (o *CheckWanLanStatusVO) GetPreOsgModel() int32`

GetPreOsgModel returns the PreOsgModel field if non-nil, zero value otherwise.

### GetPreOsgModelOk

`func (o *CheckWanLanStatusVO) GetPreOsgModelOk() (*int32, bool)`

GetPreOsgModelOk returns a tuple with the PreOsgModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreOsgModel

`func (o *CheckWanLanStatusVO) SetPreOsgModel(v int32)`

SetPreOsgModel sets PreOsgModel field to given value.

### HasPreOsgModel

`func (o *CheckWanLanStatusVO) HasPreOsgModel() bool`

HasPreOsgModel returns a boolean if a field has been set.

### GetSupportIpv6

`func (o *CheckWanLanStatusVO) GetSupportIpv6() int32`

GetSupportIpv6 returns the SupportIpv6 field if non-nil, zero value otherwise.

### GetSupportIpv6Ok

`func (o *CheckWanLanStatusVO) GetSupportIpv6Ok() (*int32, bool)`

GetSupportIpv6Ok returns a tuple with the SupportIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIpv6

`func (o *CheckWanLanStatusVO) SetSupportIpv6(v int32)`

SetSupportIpv6 sets SupportIpv6 field to given value.

### HasSupportIpv6

`func (o *CheckWanLanStatusVO) HasSupportIpv6() bool`

HasSupportIpv6 returns a boolean if a field has been set.

### GetWirelessRouter

`func (o *CheckWanLanStatusVO) GetWirelessRouter() bool`

GetWirelessRouter returns the WirelessRouter field if non-nil, zero value otherwise.

### GetWirelessRouterOk

`func (o *CheckWanLanStatusVO) GetWirelessRouterOk() (*bool, bool)`

GetWirelessRouterOk returns a tuple with the WirelessRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessRouter

`func (o *CheckWanLanStatusVO) SetWirelessRouter(v bool)`

SetWirelessRouter sets WirelessRouter field to given value.

### HasWirelessRouter

`func (o *CheckWanLanStatusVO) HasWirelessRouter() bool`

HasWirelessRouter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


