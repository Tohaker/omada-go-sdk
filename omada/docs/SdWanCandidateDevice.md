# SdWanCandidateDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityLevel** | Pointer to **int32** | The capacity level of a SD-WAN candidate device. | [optional] 
**DeviceMac** | Pointer to **string** | The MAC of a SD-WAN candidate device. | [optional] 
**DeviceName** | Pointer to **string** | The name of a SD-WAN candidate device. | [optional] 
**HasPublicIp** | Pointer to **bool** | Whether the a SD-WAN candidate has public IP. | [optional] 
**LanNetworks** | Pointer to [**[]LanNetworkBrief**](LanNetworkBrief.md) | A list of the lan network info for the candidate. | [optional] 
**Model** | Pointer to **string** | The model of a SD-WAN candidate device. | [optional] 
**ModelVersion** | Pointer to **string** | The model version of a SD-WAN candidate device. | [optional] 
**PortInfos** | Pointer to [**[]OsgPortStatBrief**](OsgPortStatBrief.md) | A list of the port info for the candidate. | [optional] 
**ShowModel** | Pointer to **string** | The showModel of a SD-WAN candidate device. | [optional] 
**SiteId** | Pointer to **string** | The ID of the site where the a SD-WAN candidate device located. | [optional] 
**SiteName** | Pointer to **string** | The name of the site where the a SD-WAN candidate device located. | [optional] 
**Status** | Pointer to **int32** | The online status of the candidate. | [optional] 
**SupportSdWanNat** | Pointer to **bool** | Whether the device support SD-WAN NAT. | [optional] 
**TunnelLimit** | Pointer to **int32** | The maximum number of VPN tunnels that can be created. | [optional] 
**Type** | Pointer to **string** | The device type of a SD-WAN candidate device. | [optional] 
**WanIp** | Pointer to **string** | The wan IP of the a SD-WAN candidate device. | [optional] 

## Methods

### NewSdWanCandidateDevice

`func NewSdWanCandidateDevice() *SdWanCandidateDevice`

NewSdWanCandidateDevice instantiates a new SdWanCandidateDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdWanCandidateDeviceWithDefaults

`func NewSdWanCandidateDeviceWithDefaults() *SdWanCandidateDevice`

NewSdWanCandidateDeviceWithDefaults instantiates a new SdWanCandidateDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityLevel

`func (o *SdWanCandidateDevice) GetCapacityLevel() int32`

GetCapacityLevel returns the CapacityLevel field if non-nil, zero value otherwise.

### GetCapacityLevelOk

`func (o *SdWanCandidateDevice) GetCapacityLevelOk() (*int32, bool)`

GetCapacityLevelOk returns a tuple with the CapacityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityLevel

`func (o *SdWanCandidateDevice) SetCapacityLevel(v int32)`

SetCapacityLevel sets CapacityLevel field to given value.

### HasCapacityLevel

`func (o *SdWanCandidateDevice) HasCapacityLevel() bool`

HasCapacityLevel returns a boolean if a field has been set.

### GetDeviceMac

`func (o *SdWanCandidateDevice) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *SdWanCandidateDevice) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *SdWanCandidateDevice) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *SdWanCandidateDevice) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDeviceName

`func (o *SdWanCandidateDevice) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *SdWanCandidateDevice) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *SdWanCandidateDevice) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *SdWanCandidateDevice) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetHasPublicIp

`func (o *SdWanCandidateDevice) GetHasPublicIp() bool`

GetHasPublicIp returns the HasPublicIp field if non-nil, zero value otherwise.

### GetHasPublicIpOk

`func (o *SdWanCandidateDevice) GetHasPublicIpOk() (*bool, bool)`

GetHasPublicIpOk returns a tuple with the HasPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPublicIp

`func (o *SdWanCandidateDevice) SetHasPublicIp(v bool)`

SetHasPublicIp sets HasPublicIp field to given value.

### HasHasPublicIp

`func (o *SdWanCandidateDevice) HasHasPublicIp() bool`

HasHasPublicIp returns a boolean if a field has been set.

### GetLanNetworks

`func (o *SdWanCandidateDevice) GetLanNetworks() []LanNetworkBrief`

GetLanNetworks returns the LanNetworks field if non-nil, zero value otherwise.

### GetLanNetworksOk

`func (o *SdWanCandidateDevice) GetLanNetworksOk() (*[]LanNetworkBrief, bool)`

GetLanNetworksOk returns a tuple with the LanNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetworks

`func (o *SdWanCandidateDevice) SetLanNetworks(v []LanNetworkBrief)`

SetLanNetworks sets LanNetworks field to given value.

### HasLanNetworks

`func (o *SdWanCandidateDevice) HasLanNetworks() bool`

HasLanNetworks returns a boolean if a field has been set.

### GetModel

`func (o *SdWanCandidateDevice) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SdWanCandidateDevice) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SdWanCandidateDevice) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SdWanCandidateDevice) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *SdWanCandidateDevice) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *SdWanCandidateDevice) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *SdWanCandidateDevice) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *SdWanCandidateDevice) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetPortInfos

`func (o *SdWanCandidateDevice) GetPortInfos() []OsgPortStatBrief`

GetPortInfos returns the PortInfos field if non-nil, zero value otherwise.

### GetPortInfosOk

`func (o *SdWanCandidateDevice) GetPortInfosOk() (*[]OsgPortStatBrief, bool)`

GetPortInfosOk returns a tuple with the PortInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortInfos

`func (o *SdWanCandidateDevice) SetPortInfos(v []OsgPortStatBrief)`

SetPortInfos sets PortInfos field to given value.

### HasPortInfos

`func (o *SdWanCandidateDevice) HasPortInfos() bool`

HasPortInfos returns a boolean if a field has been set.

### GetShowModel

`func (o *SdWanCandidateDevice) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *SdWanCandidateDevice) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *SdWanCandidateDevice) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *SdWanCandidateDevice) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetSiteId

`func (o *SdWanCandidateDevice) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SdWanCandidateDevice) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SdWanCandidateDevice) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SdWanCandidateDevice) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteName

`func (o *SdWanCandidateDevice) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *SdWanCandidateDevice) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *SdWanCandidateDevice) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *SdWanCandidateDevice) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetStatus

`func (o *SdWanCandidateDevice) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SdWanCandidateDevice) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SdWanCandidateDevice) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SdWanCandidateDevice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportSdWanNat

`func (o *SdWanCandidateDevice) GetSupportSdWanNat() bool`

GetSupportSdWanNat returns the SupportSdWanNat field if non-nil, zero value otherwise.

### GetSupportSdWanNatOk

`func (o *SdWanCandidateDevice) GetSupportSdWanNatOk() (*bool, bool)`

GetSupportSdWanNatOk returns a tuple with the SupportSdWanNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSdWanNat

`func (o *SdWanCandidateDevice) SetSupportSdWanNat(v bool)`

SetSupportSdWanNat sets SupportSdWanNat field to given value.

### HasSupportSdWanNat

`func (o *SdWanCandidateDevice) HasSupportSdWanNat() bool`

HasSupportSdWanNat returns a boolean if a field has been set.

### GetTunnelLimit

`func (o *SdWanCandidateDevice) GetTunnelLimit() int32`

GetTunnelLimit returns the TunnelLimit field if non-nil, zero value otherwise.

### GetTunnelLimitOk

`func (o *SdWanCandidateDevice) GetTunnelLimitOk() (*int32, bool)`

GetTunnelLimitOk returns a tuple with the TunnelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelLimit

`func (o *SdWanCandidateDevice) SetTunnelLimit(v int32)`

SetTunnelLimit sets TunnelLimit field to given value.

### HasTunnelLimit

`func (o *SdWanCandidateDevice) HasTunnelLimit() bool`

HasTunnelLimit returns a boolean if a field has been set.

### GetType

`func (o *SdWanCandidateDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SdWanCandidateDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SdWanCandidateDevice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SdWanCandidateDevice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWanIp

`func (o *SdWanCandidateDevice) GetWanIp() string`

GetWanIp returns the WanIp field if non-nil, zero value otherwise.

### GetWanIpOk

`func (o *SdWanCandidateDevice) GetWanIpOk() (*string, bool)`

GetWanIpOk returns a tuple with the WanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanIp

`func (o *SdWanCandidateDevice) SetWanIp(v string)`

SetWanIp sets WanIp field to given value.

### HasWanIp

`func (o *SdWanCandidateDevice) HasWanIp() bool`

HasWanIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


