# RestrictionEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterId** | **int32** | Filter ID can be obtained from &#39;Get filter list&#39; interface. | 
**NetworkId** | Pointer to **string** | Network ID. It should be the ID of lan network. It can be queried by request: Get LAN network list. | [optional] 
**NetworkName** | **string** | Network name. It should be the name of LAN network, can be obtained from &#39;Get LAN network list&#39; interface. | 

## Methods

### NewRestrictionEntity

`func NewRestrictionEntity(filterId int32, networkName string, ) *RestrictionEntity`

NewRestrictionEntity instantiates a new RestrictionEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictionEntityWithDefaults

`func NewRestrictionEntityWithDefaults() *RestrictionEntity`

NewRestrictionEntityWithDefaults instantiates a new RestrictionEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterId

`func (o *RestrictionEntity) GetFilterId() int32`

GetFilterId returns the FilterId field if non-nil, zero value otherwise.

### GetFilterIdOk

`func (o *RestrictionEntity) GetFilterIdOk() (*int32, bool)`

GetFilterIdOk returns a tuple with the FilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterId

`func (o *RestrictionEntity) SetFilterId(v int32)`

SetFilterId sets FilterId field to given value.


### GetNetworkId

`func (o *RestrictionEntity) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *RestrictionEntity) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *RestrictionEntity) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *RestrictionEntity) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkName

`func (o *RestrictionEntity) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *RestrictionEntity) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *RestrictionEntity) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


