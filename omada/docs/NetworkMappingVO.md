# NetworkMappingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestModel** | Pointer to **int32** |  | [optional] 
**LanMapping** | Pointer to [**LanMappingVO**](LanMappingVO.md) |  | [optional] 
**WanMapping** | Pointer to [**WanMappingVO**](WanMappingVO.md) |  | [optional] 
**WanNum** | Pointer to **int32** |  | [optional] 

## Methods

### NewNetworkMappingVO

`func NewNetworkMappingVO() *NetworkMappingVO`

NewNetworkMappingVO instantiates a new NetworkMappingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkMappingVOWithDefaults

`func NewNetworkMappingVOWithDefaults() *NetworkMappingVO`

NewNetworkMappingVOWithDefaults instantiates a new NetworkMappingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestModel

`func (o *NetworkMappingVO) GetDestModel() int32`

GetDestModel returns the DestModel field if non-nil, zero value otherwise.

### GetDestModelOk

`func (o *NetworkMappingVO) GetDestModelOk() (*int32, bool)`

GetDestModelOk returns a tuple with the DestModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestModel

`func (o *NetworkMappingVO) SetDestModel(v int32)`

SetDestModel sets DestModel field to given value.

### HasDestModel

`func (o *NetworkMappingVO) HasDestModel() bool`

HasDestModel returns a boolean if a field has been set.

### GetLanMapping

`func (o *NetworkMappingVO) GetLanMapping() LanMappingVO`

GetLanMapping returns the LanMapping field if non-nil, zero value otherwise.

### GetLanMappingOk

`func (o *NetworkMappingVO) GetLanMappingOk() (*LanMappingVO, bool)`

GetLanMappingOk returns a tuple with the LanMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanMapping

`func (o *NetworkMappingVO) SetLanMapping(v LanMappingVO)`

SetLanMapping sets LanMapping field to given value.

### HasLanMapping

`func (o *NetworkMappingVO) HasLanMapping() bool`

HasLanMapping returns a boolean if a field has been set.

### GetWanMapping

`func (o *NetworkMappingVO) GetWanMapping() WanMappingVO`

GetWanMapping returns the WanMapping field if non-nil, zero value otherwise.

### GetWanMappingOk

`func (o *NetworkMappingVO) GetWanMappingOk() (*WanMappingVO, bool)`

GetWanMappingOk returns a tuple with the WanMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanMapping

`func (o *NetworkMappingVO) SetWanMapping(v WanMappingVO)`

SetWanMapping sets WanMapping field to given value.

### HasWanMapping

`func (o *NetworkMappingVO) HasWanMapping() bool`

HasWanMapping returns a boolean if a field has been set.

### GetWanNum

`func (o *NetworkMappingVO) GetWanNum() int32`

GetWanNum returns the WanNum field if non-nil, zero value otherwise.

### GetWanNumOk

`func (o *NetworkMappingVO) GetWanNumOk() (*int32, bool)`

GetWanNumOk returns a tuple with the WanNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanNum

`func (o *NetworkMappingVO) SetWanNum(v int32)`

SetWanNum sets WanNum field to given value.

### HasWanNum

`func (o *NetworkMappingVO) HasWanNum() bool`

HasWanNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


