# LanNetworkVlansOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LanNetworkId** | Pointer to **string** | Lan network ID. | [optional] 
**Vlan** | Pointer to **int32** | Created vlan in site. Valid range is from 2 to 4090 in \&quot;Oui based vlan\&quot;. | [optional] 

## Methods

### NewLanNetworkVlansOpenApiVO

`func NewLanNetworkVlansOpenApiVO() *LanNetworkVlansOpenApiVO`

NewLanNetworkVlansOpenApiVO instantiates a new LanNetworkVlansOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanNetworkVlansOpenApiVOWithDefaults

`func NewLanNetworkVlansOpenApiVOWithDefaults() *LanNetworkVlansOpenApiVO`

NewLanNetworkVlansOpenApiVOWithDefaults instantiates a new LanNetworkVlansOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanNetworkId

`func (o *LanNetworkVlansOpenApiVO) GetLanNetworkId() string`

GetLanNetworkId returns the LanNetworkId field if non-nil, zero value otherwise.

### GetLanNetworkIdOk

`func (o *LanNetworkVlansOpenApiVO) GetLanNetworkIdOk() (*string, bool)`

GetLanNetworkIdOk returns a tuple with the LanNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetworkId

`func (o *LanNetworkVlansOpenApiVO) SetLanNetworkId(v string)`

SetLanNetworkId sets LanNetworkId field to given value.

### HasLanNetworkId

`func (o *LanNetworkVlansOpenApiVO) HasLanNetworkId() bool`

HasLanNetworkId returns a boolean if a field has been set.

### GetVlan

`func (o *LanNetworkVlansOpenApiVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *LanNetworkVlansOpenApiVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *LanNetworkVlansOpenApiVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *LanNetworkVlansOpenApiVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


