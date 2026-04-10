# APLANPortList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandwidthControlEnable** | Pointer to **bool** | Whether to enable band width control. | [optional] 
**Custom** | Pointer to **bool** | Whether to enter a tagged vlan customically. Enter the VLAN ID manually if it is true | [optional] 
**Duplex** | Pointer to **int32** | Real-time duplex mode;It should be a value as follows: 1: Half; 2: Full. | [optional] 
**EgressRateLimit** | Pointer to [**ONUPortRateLimitVO**](ONUPortRateLimitVO.md) |  | [optional] 
**Id** | Pointer to **string** | Port Id | [optional] 
**IngressRateLimit** | Pointer to [**ONUPortRateLimitVO**](ONUPortRateLimitVO.md) |  | [optional] 
**LanPort** | Pointer to **string** | Lan Port | [optional] 
**LinkSpeed** | Pointer to **int32** | Real-time link speed;It should be a value as follows: 1: 10Mbps; 2: 100Mbps; 3: 1000Mbps; 4: 10Gbps. | [optional] 
**LinkStatus** | Pointer to **int32** | Link status; It should be a value as follows:  1:link up;  0: link down. | [optional] 
**LocalVlanEnable** | Pointer to **bool** | Whether the port enable VLANs. | [optional] 
**LocalVlanId** | Pointer to **int32** | Local vlan id. | [optional] 
**LocalVlanNetworkId** | Pointer to **string** | Local Vlan Network Id, used to enter the LAN field. | [optional] 
**Name** | Pointer to **string** | Port Name. It should contain 1 ~ 64 ASCII characters. If it is NULL, LanPort is displayed. | [optional] 
**OnuId** | Pointer to **string** | Onu Id | [optional] 
**PoeOutEnable** | Pointer to **bool** | Whether to enable poe out. | [optional] 
**PoeState** | Pointer to **int32** | This value is only available when supportPoe is true.It should be a value as follows: 0：In the power supply; 1：Not in the power supply. | [optional] 
**Port** | Pointer to **int32** | Port Number | [optional] 
**PortStatus** | Pointer to [**ApPortStatusVO**](ApPortStatusVO.md) |  | [optional] 
**PortType** | Pointer to **int32** | Port Type | [optional] 
**Status** | Pointer to **bool** | Whether to disable the port, defaults to true. | [optional] 
**SupportBandwidthControl** | Pointer to **bool** | Whether band width control is supported | [optional] 
**SupportPoe** | Pointer to **bool** | Whether poe is supported. | [optional] 
**SupportStatusEnable** | Pointer to **bool** | Whether the port disabling is supported. | [optional] 
**SupportStatusInform** | Pointer to **bool** | Whether report the status of port links through inform is supported. | [optional] 
**SupportVlan** | Pointer to **bool** | Whether the port supports VLANs. | [optional] 
**SupportVlanOption** | Pointer to **bool** | Whether configure a VLAN in profile is supported | [optional] 
**SupportVlanTagged** | Pointer to **bool** | Whether vlan tagged is supported | [optional] 
**Tagged** | Pointer to **string** | Tagged vlan list, it must be present when custom is true. | [optional] 
**TaggedNetworkId** | Pointer to **[]string** | The Id list of the tagged vlan network profile, it must be present when custom is true. | [optional] 
**Untagged** | Pointer to **string** | Untagged vlan list, it must be present when custom is true. Two data cannot be duplicated. | [optional] 
**UntaggedNetworkId** | Pointer to **[]string** | The Id list of the untagged vlan network profile, it must be present when custom is true. | [optional] 
**UplinkPort** | Pointer to **bool** | Whether it is an actual uplink port. | [optional] 
**VoipState** | Pointer to **int32** | This field has a value for the voice port.It should be a value as follows: 0：Off-hook; 1：On-hook. | [optional] 

## Methods

### NewAPLANPortList

`func NewAPLANPortList() *APLANPortList`

NewAPLANPortList instantiates a new APLANPortList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPLANPortListWithDefaults

`func NewAPLANPortListWithDefaults() *APLANPortList`

NewAPLANPortListWithDefaults instantiates a new APLANPortList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidthControlEnable

`func (o *APLANPortList) GetBandwidthControlEnable() bool`

GetBandwidthControlEnable returns the BandwidthControlEnable field if non-nil, zero value otherwise.

### GetBandwidthControlEnableOk

`func (o *APLANPortList) GetBandwidthControlEnableOk() (*bool, bool)`

GetBandwidthControlEnableOk returns a tuple with the BandwidthControlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthControlEnable

`func (o *APLANPortList) SetBandwidthControlEnable(v bool)`

SetBandwidthControlEnable sets BandwidthControlEnable field to given value.

### HasBandwidthControlEnable

`func (o *APLANPortList) HasBandwidthControlEnable() bool`

HasBandwidthControlEnable returns a boolean if a field has been set.

### GetCustom

`func (o *APLANPortList) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *APLANPortList) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *APLANPortList) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *APLANPortList) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetDuplex

`func (o *APLANPortList) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *APLANPortList) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *APLANPortList) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *APLANPortList) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetEgressRateLimit

`func (o *APLANPortList) GetEgressRateLimit() ONUPortRateLimitVO`

GetEgressRateLimit returns the EgressRateLimit field if non-nil, zero value otherwise.

### GetEgressRateLimitOk

`func (o *APLANPortList) GetEgressRateLimitOk() (*ONUPortRateLimitVO, bool)`

GetEgressRateLimitOk returns a tuple with the EgressRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressRateLimit

`func (o *APLANPortList) SetEgressRateLimit(v ONUPortRateLimitVO)`

SetEgressRateLimit sets EgressRateLimit field to given value.

### HasEgressRateLimit

`func (o *APLANPortList) HasEgressRateLimit() bool`

HasEgressRateLimit returns a boolean if a field has been set.

### GetId

`func (o *APLANPortList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APLANPortList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APLANPortList) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *APLANPortList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIngressRateLimit

`func (o *APLANPortList) GetIngressRateLimit() ONUPortRateLimitVO`

GetIngressRateLimit returns the IngressRateLimit field if non-nil, zero value otherwise.

### GetIngressRateLimitOk

`func (o *APLANPortList) GetIngressRateLimitOk() (*ONUPortRateLimitVO, bool)`

GetIngressRateLimitOk returns a tuple with the IngressRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressRateLimit

`func (o *APLANPortList) SetIngressRateLimit(v ONUPortRateLimitVO)`

SetIngressRateLimit sets IngressRateLimit field to given value.

### HasIngressRateLimit

`func (o *APLANPortList) HasIngressRateLimit() bool`

HasIngressRateLimit returns a boolean if a field has been set.

### GetLanPort

`func (o *APLANPortList) GetLanPort() string`

GetLanPort returns the LanPort field if non-nil, zero value otherwise.

### GetLanPortOk

`func (o *APLANPortList) GetLanPortOk() (*string, bool)`

GetLanPortOk returns a tuple with the LanPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanPort

`func (o *APLANPortList) SetLanPort(v string)`

SetLanPort sets LanPort field to given value.

### HasLanPort

`func (o *APLANPortList) HasLanPort() bool`

HasLanPort returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *APLANPortList) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *APLANPortList) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *APLANPortList) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *APLANPortList) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLinkStatus

`func (o *APLANPortList) GetLinkStatus() int32`

GetLinkStatus returns the LinkStatus field if non-nil, zero value otherwise.

### GetLinkStatusOk

`func (o *APLANPortList) GetLinkStatusOk() (*int32, bool)`

GetLinkStatusOk returns a tuple with the LinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStatus

`func (o *APLANPortList) SetLinkStatus(v int32)`

SetLinkStatus sets LinkStatus field to given value.

### HasLinkStatus

`func (o *APLANPortList) HasLinkStatus() bool`

HasLinkStatus returns a boolean if a field has been set.

### GetLocalVlanEnable

`func (o *APLANPortList) GetLocalVlanEnable() bool`

GetLocalVlanEnable returns the LocalVlanEnable field if non-nil, zero value otherwise.

### GetLocalVlanEnableOk

`func (o *APLANPortList) GetLocalVlanEnableOk() (*bool, bool)`

GetLocalVlanEnableOk returns a tuple with the LocalVlanEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVlanEnable

`func (o *APLANPortList) SetLocalVlanEnable(v bool)`

SetLocalVlanEnable sets LocalVlanEnable field to given value.

### HasLocalVlanEnable

`func (o *APLANPortList) HasLocalVlanEnable() bool`

HasLocalVlanEnable returns a boolean if a field has been set.

### GetLocalVlanId

`func (o *APLANPortList) GetLocalVlanId() int32`

GetLocalVlanId returns the LocalVlanId field if non-nil, zero value otherwise.

### GetLocalVlanIdOk

`func (o *APLANPortList) GetLocalVlanIdOk() (*int32, bool)`

GetLocalVlanIdOk returns a tuple with the LocalVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVlanId

`func (o *APLANPortList) SetLocalVlanId(v int32)`

SetLocalVlanId sets LocalVlanId field to given value.

### HasLocalVlanId

`func (o *APLANPortList) HasLocalVlanId() bool`

HasLocalVlanId returns a boolean if a field has been set.

### GetLocalVlanNetworkId

`func (o *APLANPortList) GetLocalVlanNetworkId() string`

GetLocalVlanNetworkId returns the LocalVlanNetworkId field if non-nil, zero value otherwise.

### GetLocalVlanNetworkIdOk

`func (o *APLANPortList) GetLocalVlanNetworkIdOk() (*string, bool)`

GetLocalVlanNetworkIdOk returns a tuple with the LocalVlanNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVlanNetworkId

`func (o *APLANPortList) SetLocalVlanNetworkId(v string)`

SetLocalVlanNetworkId sets LocalVlanNetworkId field to given value.

### HasLocalVlanNetworkId

`func (o *APLANPortList) HasLocalVlanNetworkId() bool`

HasLocalVlanNetworkId returns a boolean if a field has been set.

### GetName

`func (o *APLANPortList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *APLANPortList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *APLANPortList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *APLANPortList) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnuId

`func (o *APLANPortList) GetOnuId() string`

GetOnuId returns the OnuId field if non-nil, zero value otherwise.

### GetOnuIdOk

`func (o *APLANPortList) GetOnuIdOk() (*string, bool)`

GetOnuIdOk returns a tuple with the OnuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuId

`func (o *APLANPortList) SetOnuId(v string)`

SetOnuId sets OnuId field to given value.

### HasOnuId

`func (o *APLANPortList) HasOnuId() bool`

HasOnuId returns a boolean if a field has been set.

### GetPoeOutEnable

`func (o *APLANPortList) GetPoeOutEnable() bool`

GetPoeOutEnable returns the PoeOutEnable field if non-nil, zero value otherwise.

### GetPoeOutEnableOk

`func (o *APLANPortList) GetPoeOutEnableOk() (*bool, bool)`

GetPoeOutEnableOk returns a tuple with the PoeOutEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeOutEnable

`func (o *APLANPortList) SetPoeOutEnable(v bool)`

SetPoeOutEnable sets PoeOutEnable field to given value.

### HasPoeOutEnable

`func (o *APLANPortList) HasPoeOutEnable() bool`

HasPoeOutEnable returns a boolean if a field has been set.

### GetPoeState

`func (o *APLANPortList) GetPoeState() int32`

GetPoeState returns the PoeState field if non-nil, zero value otherwise.

### GetPoeStateOk

`func (o *APLANPortList) GetPoeStateOk() (*int32, bool)`

GetPoeStateOk returns a tuple with the PoeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeState

`func (o *APLANPortList) SetPoeState(v int32)`

SetPoeState sets PoeState field to given value.

### HasPoeState

`func (o *APLANPortList) HasPoeState() bool`

HasPoeState returns a boolean if a field has been set.

### GetPort

`func (o *APLANPortList) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *APLANPortList) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *APLANPortList) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *APLANPortList) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortStatus

`func (o *APLANPortList) GetPortStatus() ApPortStatusVO`

GetPortStatus returns the PortStatus field if non-nil, zero value otherwise.

### GetPortStatusOk

`func (o *APLANPortList) GetPortStatusOk() (*ApPortStatusVO, bool)`

GetPortStatusOk returns a tuple with the PortStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStatus

`func (o *APLANPortList) SetPortStatus(v ApPortStatusVO)`

SetPortStatus sets PortStatus field to given value.

### HasPortStatus

`func (o *APLANPortList) HasPortStatus() bool`

HasPortStatus returns a boolean if a field has been set.

### GetPortType

`func (o *APLANPortList) GetPortType() int32`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *APLANPortList) GetPortTypeOk() (*int32, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *APLANPortList) SetPortType(v int32)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *APLANPortList) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetStatus

`func (o *APLANPortList) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *APLANPortList) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *APLANPortList) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *APLANPortList) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportBandwidthControl

`func (o *APLANPortList) GetSupportBandwidthControl() bool`

GetSupportBandwidthControl returns the SupportBandwidthControl field if non-nil, zero value otherwise.

### GetSupportBandwidthControlOk

`func (o *APLANPortList) GetSupportBandwidthControlOk() (*bool, bool)`

GetSupportBandwidthControlOk returns a tuple with the SupportBandwidthControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportBandwidthControl

`func (o *APLANPortList) SetSupportBandwidthControl(v bool)`

SetSupportBandwidthControl sets SupportBandwidthControl field to given value.

### HasSupportBandwidthControl

`func (o *APLANPortList) HasSupportBandwidthControl() bool`

HasSupportBandwidthControl returns a boolean if a field has been set.

### GetSupportPoe

`func (o *APLANPortList) GetSupportPoe() bool`

GetSupportPoe returns the SupportPoe field if non-nil, zero value otherwise.

### GetSupportPoeOk

`func (o *APLANPortList) GetSupportPoeOk() (*bool, bool)`

GetSupportPoeOk returns a tuple with the SupportPoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPoe

`func (o *APLANPortList) SetSupportPoe(v bool)`

SetSupportPoe sets SupportPoe field to given value.

### HasSupportPoe

`func (o *APLANPortList) HasSupportPoe() bool`

HasSupportPoe returns a boolean if a field has been set.

### GetSupportStatusEnable

`func (o *APLANPortList) GetSupportStatusEnable() bool`

GetSupportStatusEnable returns the SupportStatusEnable field if non-nil, zero value otherwise.

### GetSupportStatusEnableOk

`func (o *APLANPortList) GetSupportStatusEnableOk() (*bool, bool)`

GetSupportStatusEnableOk returns a tuple with the SupportStatusEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportStatusEnable

`func (o *APLANPortList) SetSupportStatusEnable(v bool)`

SetSupportStatusEnable sets SupportStatusEnable field to given value.

### HasSupportStatusEnable

`func (o *APLANPortList) HasSupportStatusEnable() bool`

HasSupportStatusEnable returns a boolean if a field has been set.

### GetSupportStatusInform

`func (o *APLANPortList) GetSupportStatusInform() bool`

GetSupportStatusInform returns the SupportStatusInform field if non-nil, zero value otherwise.

### GetSupportStatusInformOk

`func (o *APLANPortList) GetSupportStatusInformOk() (*bool, bool)`

GetSupportStatusInformOk returns a tuple with the SupportStatusInform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportStatusInform

`func (o *APLANPortList) SetSupportStatusInform(v bool)`

SetSupportStatusInform sets SupportStatusInform field to given value.

### HasSupportStatusInform

`func (o *APLANPortList) HasSupportStatusInform() bool`

HasSupportStatusInform returns a boolean if a field has been set.

### GetSupportVlan

`func (o *APLANPortList) GetSupportVlan() bool`

GetSupportVlan returns the SupportVlan field if non-nil, zero value otherwise.

### GetSupportVlanOk

`func (o *APLANPortList) GetSupportVlanOk() (*bool, bool)`

GetSupportVlanOk returns a tuple with the SupportVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVlan

`func (o *APLANPortList) SetSupportVlan(v bool)`

SetSupportVlan sets SupportVlan field to given value.

### HasSupportVlan

`func (o *APLANPortList) HasSupportVlan() bool`

HasSupportVlan returns a boolean if a field has been set.

### GetSupportVlanOption

`func (o *APLANPortList) GetSupportVlanOption() bool`

GetSupportVlanOption returns the SupportVlanOption field if non-nil, zero value otherwise.

### GetSupportVlanOptionOk

`func (o *APLANPortList) GetSupportVlanOptionOk() (*bool, bool)`

GetSupportVlanOptionOk returns a tuple with the SupportVlanOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVlanOption

`func (o *APLANPortList) SetSupportVlanOption(v bool)`

SetSupportVlanOption sets SupportVlanOption field to given value.

### HasSupportVlanOption

`func (o *APLANPortList) HasSupportVlanOption() bool`

HasSupportVlanOption returns a boolean if a field has been set.

### GetSupportVlanTagged

`func (o *APLANPortList) GetSupportVlanTagged() bool`

GetSupportVlanTagged returns the SupportVlanTagged field if non-nil, zero value otherwise.

### GetSupportVlanTaggedOk

`func (o *APLANPortList) GetSupportVlanTaggedOk() (*bool, bool)`

GetSupportVlanTaggedOk returns a tuple with the SupportVlanTagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVlanTagged

`func (o *APLANPortList) SetSupportVlanTagged(v bool)`

SetSupportVlanTagged sets SupportVlanTagged field to given value.

### HasSupportVlanTagged

`func (o *APLANPortList) HasSupportVlanTagged() bool`

HasSupportVlanTagged returns a boolean if a field has been set.

### GetTagged

`func (o *APLANPortList) GetTagged() string`

GetTagged returns the Tagged field if non-nil, zero value otherwise.

### GetTaggedOk

`func (o *APLANPortList) GetTaggedOk() (*string, bool)`

GetTaggedOk returns a tuple with the Tagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagged

`func (o *APLANPortList) SetTagged(v string)`

SetTagged sets Tagged field to given value.

### HasTagged

`func (o *APLANPortList) HasTagged() bool`

HasTagged returns a boolean if a field has been set.

### GetTaggedNetworkId

`func (o *APLANPortList) GetTaggedNetworkId() []string`

GetTaggedNetworkId returns the TaggedNetworkId field if non-nil, zero value otherwise.

### GetTaggedNetworkIdOk

`func (o *APLANPortList) GetTaggedNetworkIdOk() (*[]string, bool)`

GetTaggedNetworkIdOk returns a tuple with the TaggedNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedNetworkId

`func (o *APLANPortList) SetTaggedNetworkId(v []string)`

SetTaggedNetworkId sets TaggedNetworkId field to given value.

### HasTaggedNetworkId

`func (o *APLANPortList) HasTaggedNetworkId() bool`

HasTaggedNetworkId returns a boolean if a field has been set.

### GetUntagged

`func (o *APLANPortList) GetUntagged() string`

GetUntagged returns the Untagged field if non-nil, zero value otherwise.

### GetUntaggedOk

`func (o *APLANPortList) GetUntaggedOk() (*string, bool)`

GetUntaggedOk returns a tuple with the Untagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagged

`func (o *APLANPortList) SetUntagged(v string)`

SetUntagged sets Untagged field to given value.

### HasUntagged

`func (o *APLANPortList) HasUntagged() bool`

HasUntagged returns a boolean if a field has been set.

### GetUntaggedNetworkId

`func (o *APLANPortList) GetUntaggedNetworkId() []string`

GetUntaggedNetworkId returns the UntaggedNetworkId field if non-nil, zero value otherwise.

### GetUntaggedNetworkIdOk

`func (o *APLANPortList) GetUntaggedNetworkIdOk() (*[]string, bool)`

GetUntaggedNetworkIdOk returns a tuple with the UntaggedNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntaggedNetworkId

`func (o *APLANPortList) SetUntaggedNetworkId(v []string)`

SetUntaggedNetworkId sets UntaggedNetworkId field to given value.

### HasUntaggedNetworkId

`func (o *APLANPortList) HasUntaggedNetworkId() bool`

HasUntaggedNetworkId returns a boolean if a field has been set.

### GetUplinkPort

`func (o *APLANPortList) GetUplinkPort() bool`

GetUplinkPort returns the UplinkPort field if non-nil, zero value otherwise.

### GetUplinkPortOk

`func (o *APLANPortList) GetUplinkPortOk() (*bool, bool)`

GetUplinkPortOk returns a tuple with the UplinkPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkPort

`func (o *APLANPortList) SetUplinkPort(v bool)`

SetUplinkPort sets UplinkPort field to given value.

### HasUplinkPort

`func (o *APLANPortList) HasUplinkPort() bool`

HasUplinkPort returns a boolean if a field has been set.

### GetVoipState

`func (o *APLANPortList) GetVoipState() int32`

GetVoipState returns the VoipState field if non-nil, zero value otherwise.

### GetVoipStateOk

`func (o *APLANPortList) GetVoipStateOk() (*int32, bool)`

GetVoipStateOk returns a tuple with the VoipState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoipState

`func (o *APLANPortList) SetVoipState(v int32)`

SetVoipState sets VoipState field to given value.

### HasVoipState

`func (o *APLANPortList) HasVoipState() bool`

HasVoipState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


