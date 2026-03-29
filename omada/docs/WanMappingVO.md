# WanMappingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableWanPorts** | Pointer to **[]string** |  | [optional] 
**WanRelation** | Pointer to [**[]WanPortMappingVO**](WanPortMappingVO.md) |  | [optional] 

## Methods

### NewWanMappingVO

`func NewWanMappingVO() *WanMappingVO`

NewWanMappingVO instantiates a new WanMappingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWanMappingVOWithDefaults

`func NewWanMappingVOWithDefaults() *WanMappingVO`

NewWanMappingVOWithDefaults instantiates a new WanMappingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableWanPorts

`func (o *WanMappingVO) GetEnableWanPorts() []string`

GetEnableWanPorts returns the EnableWanPorts field if non-nil, zero value otherwise.

### GetEnableWanPortsOk

`func (o *WanMappingVO) GetEnableWanPortsOk() (*[]string, bool)`

GetEnableWanPortsOk returns a tuple with the EnableWanPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWanPorts

`func (o *WanMappingVO) SetEnableWanPorts(v []string)`

SetEnableWanPorts sets EnableWanPorts field to given value.

### HasEnableWanPorts

`func (o *WanMappingVO) HasEnableWanPorts() bool`

HasEnableWanPorts returns a boolean if a field has been set.

### GetWanRelation

`func (o *WanMappingVO) GetWanRelation() []WanPortMappingVO`

GetWanRelation returns the WanRelation field if non-nil, zero value otherwise.

### GetWanRelationOk

`func (o *WanMappingVO) GetWanRelationOk() (*[]WanPortMappingVO, bool)`

GetWanRelationOk returns a tuple with the WanRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanRelation

`func (o *WanMappingVO) SetWanRelation(v []WanPortMappingVO)`

SetWanRelation sets WanRelation field to given value.

### HasWanRelation

`func (o *WanMappingVO) HasWanRelation() bool`

HasWanRelation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


