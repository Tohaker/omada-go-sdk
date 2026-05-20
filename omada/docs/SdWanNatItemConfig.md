# SdWanNatItemConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewaySubnet** | **string** | The IP range of the original network before mapping | 
**LanNetworkId** | Pointer to **string** | The ID of the original network before mapping | [optional] 
**MappedNetwork** | Pointer to **string** | mapped network | [optional] 
**NetworkType** | Pointer to **int32** | Network type, 0/null: LAN network, 1: custom route | [optional] 
**SiteId** | **string** | The ID of the site | 

## Methods

### NewSdWanNatItemConfig

`func NewSdWanNatItemConfig(gatewaySubnet string, siteId string, ) *SdWanNatItemConfig`

NewSdWanNatItemConfig instantiates a new SdWanNatItemConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdWanNatItemConfigWithDefaults

`func NewSdWanNatItemConfigWithDefaults() *SdWanNatItemConfig`

NewSdWanNatItemConfigWithDefaults instantiates a new SdWanNatItemConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewaySubnet

`func (o *SdWanNatItemConfig) GetGatewaySubnet() string`

GetGatewaySubnet returns the GatewaySubnet field if non-nil, zero value otherwise.

### GetGatewaySubnetOk

`func (o *SdWanNatItemConfig) GetGatewaySubnetOk() (*string, bool)`

GetGatewaySubnetOk returns a tuple with the GatewaySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySubnet

`func (o *SdWanNatItemConfig) SetGatewaySubnet(v string)`

SetGatewaySubnet sets GatewaySubnet field to given value.


### GetLanNetworkId

`func (o *SdWanNatItemConfig) GetLanNetworkId() string`

GetLanNetworkId returns the LanNetworkId field if non-nil, zero value otherwise.

### GetLanNetworkIdOk

`func (o *SdWanNatItemConfig) GetLanNetworkIdOk() (*string, bool)`

GetLanNetworkIdOk returns a tuple with the LanNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetworkId

`func (o *SdWanNatItemConfig) SetLanNetworkId(v string)`

SetLanNetworkId sets LanNetworkId field to given value.

### HasLanNetworkId

`func (o *SdWanNatItemConfig) HasLanNetworkId() bool`

HasLanNetworkId returns a boolean if a field has been set.

### GetMappedNetwork

`func (o *SdWanNatItemConfig) GetMappedNetwork() string`

GetMappedNetwork returns the MappedNetwork field if non-nil, zero value otherwise.

### GetMappedNetworkOk

`func (o *SdWanNatItemConfig) GetMappedNetworkOk() (*string, bool)`

GetMappedNetworkOk returns a tuple with the MappedNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedNetwork

`func (o *SdWanNatItemConfig) SetMappedNetwork(v string)`

SetMappedNetwork sets MappedNetwork field to given value.

### HasMappedNetwork

`func (o *SdWanNatItemConfig) HasMappedNetwork() bool`

HasMappedNetwork returns a boolean if a field has been set.

### GetNetworkType

`func (o *SdWanNatItemConfig) GetNetworkType() int32`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *SdWanNatItemConfig) GetNetworkTypeOk() (*int32, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *SdWanNatItemConfig) SetNetworkType(v int32)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *SdWanNatItemConfig) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetSiteId

`func (o *SdWanNatItemConfig) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SdWanNatItemConfig) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SdWanNatItemConfig) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


