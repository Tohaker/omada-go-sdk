# LanMulticastVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FloodKnownEnable** | Pointer to **bool** | Whether open flood known protocols. | [optional] 
**Id** | Pointer to **string** | The primary id of the multicast snooping. | [optional] 
**IgmpQueriers** | Pointer to [**[]IgmpConfigVO**](IgmpConfigVO.md) | The querier configs for IGMP snooping(ipv4). | [optional] 
**MldQueriers** | Pointer to [**[]MldConfigVO**](MldConfigVO.md) | The querier configs for MLD snooping(ipv6). | [optional] 
**Name** | Pointer to **string** | The name of the multicast snooping. | [optional] 
**Networks** | Pointer to [**[]NetworkVO**](NetworkVO.md) | The collection of snooping network ids related to this multicast snooping config. | [optional] 
**Protocol** | Pointer to **int32** | When multicast snooping value 0,then represents the IGMP snooping type,else if value 1,then represents MLD snooping type. | [optional] 
**QuerierEnable** | Pointer to **bool** | Whether open querier. | [optional] 
**Resource** | Pointer to **int32** | resource | [optional] 
**RouterPortEnable** | Pointer to **bool** | Whether open manual router port config. | [optional] 
**RouterPorts** | Pointer to [**[]RouterPortVO**](RouterPortVO.md) | The specific list of router ports info,including network, devices and the ports on them. | [optional] 
**SnoopConfig** | Pointer to [**SnoopConfigVO**](SnoopConfigVO.md) |  | [optional] 
**UnknownMulticastExceptDevice** | Pointer to [**UnknownMulticastConfigVO**](UnknownMulticastConfigVO.md) |  | [optional] 
**UnknownMulticastRule** | Pointer to **int32** | When it selects 0,then send forward, it selects 1,then discard info, it selects 2,then route port first. | [optional] 

## Methods

### NewLanMulticastVO

`func NewLanMulticastVO() *LanMulticastVO`

NewLanMulticastVO instantiates a new LanMulticastVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanMulticastVOWithDefaults

`func NewLanMulticastVOWithDefaults() *LanMulticastVO`

NewLanMulticastVOWithDefaults instantiates a new LanMulticastVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFloodKnownEnable

`func (o *LanMulticastVO) GetFloodKnownEnable() bool`

GetFloodKnownEnable returns the FloodKnownEnable field if non-nil, zero value otherwise.

### GetFloodKnownEnableOk

`func (o *LanMulticastVO) GetFloodKnownEnableOk() (*bool, bool)`

GetFloodKnownEnableOk returns a tuple with the FloodKnownEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloodKnownEnable

`func (o *LanMulticastVO) SetFloodKnownEnable(v bool)`

SetFloodKnownEnable sets FloodKnownEnable field to given value.

### HasFloodKnownEnable

`func (o *LanMulticastVO) HasFloodKnownEnable() bool`

HasFloodKnownEnable returns a boolean if a field has been set.

### GetId

`func (o *LanMulticastVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LanMulticastVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LanMulticastVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LanMulticastVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgmpQueriers

`func (o *LanMulticastVO) GetIgmpQueriers() []IgmpConfigVO`

GetIgmpQueriers returns the IgmpQueriers field if non-nil, zero value otherwise.

### GetIgmpQueriersOk

`func (o *LanMulticastVO) GetIgmpQueriersOk() (*[]IgmpConfigVO, bool)`

GetIgmpQueriersOk returns a tuple with the IgmpQueriers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpQueriers

`func (o *LanMulticastVO) SetIgmpQueriers(v []IgmpConfigVO)`

SetIgmpQueriers sets IgmpQueriers field to given value.

### HasIgmpQueriers

`func (o *LanMulticastVO) HasIgmpQueriers() bool`

HasIgmpQueriers returns a boolean if a field has been set.

### GetMldQueriers

`func (o *LanMulticastVO) GetMldQueriers() []MldConfigVO`

GetMldQueriers returns the MldQueriers field if non-nil, zero value otherwise.

### GetMldQueriersOk

`func (o *LanMulticastVO) GetMldQueriersOk() (*[]MldConfigVO, bool)`

GetMldQueriersOk returns a tuple with the MldQueriers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMldQueriers

`func (o *LanMulticastVO) SetMldQueriers(v []MldConfigVO)`

SetMldQueriers sets MldQueriers field to given value.

### HasMldQueriers

`func (o *LanMulticastVO) HasMldQueriers() bool`

HasMldQueriers returns a boolean if a field has been set.

### GetName

`func (o *LanMulticastVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LanMulticastVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LanMulticastVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LanMulticastVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworks

`func (o *LanMulticastVO) GetNetworks() []NetworkVO`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *LanMulticastVO) GetNetworksOk() (*[]NetworkVO, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *LanMulticastVO) SetNetworks(v []NetworkVO)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *LanMulticastVO) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetProtocol

`func (o *LanMulticastVO) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LanMulticastVO) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LanMulticastVO) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *LanMulticastVO) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetQuerierEnable

`func (o *LanMulticastVO) GetQuerierEnable() bool`

GetQuerierEnable returns the QuerierEnable field if non-nil, zero value otherwise.

### GetQuerierEnableOk

`func (o *LanMulticastVO) GetQuerierEnableOk() (*bool, bool)`

GetQuerierEnableOk returns a tuple with the QuerierEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerierEnable

`func (o *LanMulticastVO) SetQuerierEnable(v bool)`

SetQuerierEnable sets QuerierEnable field to given value.

### HasQuerierEnable

`func (o *LanMulticastVO) HasQuerierEnable() bool`

HasQuerierEnable returns a boolean if a field has been set.

### GetResource

`func (o *LanMulticastVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *LanMulticastVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *LanMulticastVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *LanMulticastVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetRouterPortEnable

`func (o *LanMulticastVO) GetRouterPortEnable() bool`

GetRouterPortEnable returns the RouterPortEnable field if non-nil, zero value otherwise.

### GetRouterPortEnableOk

`func (o *LanMulticastVO) GetRouterPortEnableOk() (*bool, bool)`

GetRouterPortEnableOk returns a tuple with the RouterPortEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterPortEnable

`func (o *LanMulticastVO) SetRouterPortEnable(v bool)`

SetRouterPortEnable sets RouterPortEnable field to given value.

### HasRouterPortEnable

`func (o *LanMulticastVO) HasRouterPortEnable() bool`

HasRouterPortEnable returns a boolean if a field has been set.

### GetRouterPorts

`func (o *LanMulticastVO) GetRouterPorts() []RouterPortVO`

GetRouterPorts returns the RouterPorts field if non-nil, zero value otherwise.

### GetRouterPortsOk

`func (o *LanMulticastVO) GetRouterPortsOk() (*[]RouterPortVO, bool)`

GetRouterPortsOk returns a tuple with the RouterPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterPorts

`func (o *LanMulticastVO) SetRouterPorts(v []RouterPortVO)`

SetRouterPorts sets RouterPorts field to given value.

### HasRouterPorts

`func (o *LanMulticastVO) HasRouterPorts() bool`

HasRouterPorts returns a boolean if a field has been set.

### GetSnoopConfig

`func (o *LanMulticastVO) GetSnoopConfig() SnoopConfigVO`

GetSnoopConfig returns the SnoopConfig field if non-nil, zero value otherwise.

### GetSnoopConfigOk

`func (o *LanMulticastVO) GetSnoopConfigOk() (*SnoopConfigVO, bool)`

GetSnoopConfigOk returns a tuple with the SnoopConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnoopConfig

`func (o *LanMulticastVO) SetSnoopConfig(v SnoopConfigVO)`

SetSnoopConfig sets SnoopConfig field to given value.

### HasSnoopConfig

`func (o *LanMulticastVO) HasSnoopConfig() bool`

HasSnoopConfig returns a boolean if a field has been set.

### GetUnknownMulticastExceptDevice

`func (o *LanMulticastVO) GetUnknownMulticastExceptDevice() UnknownMulticastConfigVO`

GetUnknownMulticastExceptDevice returns the UnknownMulticastExceptDevice field if non-nil, zero value otherwise.

### GetUnknownMulticastExceptDeviceOk

`func (o *LanMulticastVO) GetUnknownMulticastExceptDeviceOk() (*UnknownMulticastConfigVO, bool)`

GetUnknownMulticastExceptDeviceOk returns a tuple with the UnknownMulticastExceptDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownMulticastExceptDevice

`func (o *LanMulticastVO) SetUnknownMulticastExceptDevice(v UnknownMulticastConfigVO)`

SetUnknownMulticastExceptDevice sets UnknownMulticastExceptDevice field to given value.

### HasUnknownMulticastExceptDevice

`func (o *LanMulticastVO) HasUnknownMulticastExceptDevice() bool`

HasUnknownMulticastExceptDevice returns a boolean if a field has been set.

### GetUnknownMulticastRule

`func (o *LanMulticastVO) GetUnknownMulticastRule() int32`

GetUnknownMulticastRule returns the UnknownMulticastRule field if non-nil, zero value otherwise.

### GetUnknownMulticastRuleOk

`func (o *LanMulticastVO) GetUnknownMulticastRuleOk() (*int32, bool)`

GetUnknownMulticastRuleOk returns a tuple with the UnknownMulticastRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownMulticastRule

`func (o *LanMulticastVO) SetUnknownMulticastRule(v int32)`

SetUnknownMulticastRule sets UnknownMulticastRule field to given value.

### HasUnknownMulticastRule

`func (o *LanMulticastVO) HasUnknownMulticastRule() bool`

HasUnknownMulticastRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


