# RestrictionResultEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **string** | Filter name | [optional] 
**FilterId** | Pointer to **int32** | Filter ID can be obtained from &#39;Get filter list&#39; interface. | [optional] 
**NetworkName** | Pointer to **string** | Network name. It should be the name of LAN network, can be obtained from &#39;Get LAN network list&#39; interface. | [optional] 
**RestrictionId** | Pointer to **int32** | Restriction ID | [optional] 
**Subnet** | Pointer to **string** | Subnet | [optional] 
**SubnetIpv6** | Pointer to **string** | Subnet Ipv6 | [optional] 
**Vlan** | Pointer to **int32** | Vlan ID | [optional] 
**Vlans** | Pointer to **string** | Vlan IDS | [optional] 

## Methods

### NewRestrictionResultEntity

`func NewRestrictionResultEntity() *RestrictionResultEntity`

NewRestrictionResultEntity instantiates a new RestrictionResultEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictionResultEntityWithDefaults

`func NewRestrictionResultEntityWithDefaults() *RestrictionResultEntity`

NewRestrictionResultEntityWithDefaults instantiates a new RestrictionResultEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *RestrictionResultEntity) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *RestrictionResultEntity) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *RestrictionResultEntity) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *RestrictionResultEntity) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFilterId

`func (o *RestrictionResultEntity) GetFilterId() int32`

GetFilterId returns the FilterId field if non-nil, zero value otherwise.

### GetFilterIdOk

`func (o *RestrictionResultEntity) GetFilterIdOk() (*int32, bool)`

GetFilterIdOk returns a tuple with the FilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterId

`func (o *RestrictionResultEntity) SetFilterId(v int32)`

SetFilterId sets FilterId field to given value.

### HasFilterId

`func (o *RestrictionResultEntity) HasFilterId() bool`

HasFilterId returns a boolean if a field has been set.

### GetNetworkName

`func (o *RestrictionResultEntity) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *RestrictionResultEntity) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *RestrictionResultEntity) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *RestrictionResultEntity) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetRestrictionId

`func (o *RestrictionResultEntity) GetRestrictionId() int32`

GetRestrictionId returns the RestrictionId field if non-nil, zero value otherwise.

### GetRestrictionIdOk

`func (o *RestrictionResultEntity) GetRestrictionIdOk() (*int32, bool)`

GetRestrictionIdOk returns a tuple with the RestrictionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionId

`func (o *RestrictionResultEntity) SetRestrictionId(v int32)`

SetRestrictionId sets RestrictionId field to given value.

### HasRestrictionId

`func (o *RestrictionResultEntity) HasRestrictionId() bool`

HasRestrictionId returns a boolean if a field has been set.

### GetSubnet

`func (o *RestrictionResultEntity) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *RestrictionResultEntity) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *RestrictionResultEntity) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *RestrictionResultEntity) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetSubnetIpv6

`func (o *RestrictionResultEntity) GetSubnetIpv6() string`

GetSubnetIpv6 returns the SubnetIpv6 field if non-nil, zero value otherwise.

### GetSubnetIpv6Ok

`func (o *RestrictionResultEntity) GetSubnetIpv6Ok() (*string, bool)`

GetSubnetIpv6Ok returns a tuple with the SubnetIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIpv6

`func (o *RestrictionResultEntity) SetSubnetIpv6(v string)`

SetSubnetIpv6 sets SubnetIpv6 field to given value.

### HasSubnetIpv6

`func (o *RestrictionResultEntity) HasSubnetIpv6() bool`

HasSubnetIpv6 returns a boolean if a field has been set.

### GetVlan

`func (o *RestrictionResultEntity) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *RestrictionResultEntity) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *RestrictionResultEntity) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *RestrictionResultEntity) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVlans

`func (o *RestrictionResultEntity) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *RestrictionResultEntity) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *RestrictionResultEntity) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *RestrictionResultEntity) HasVlans() bool`

HasVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


