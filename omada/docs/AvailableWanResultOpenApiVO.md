# AvailableWanResultOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | enable | [optional] 
**WanPorts** | Pointer to [**[]AvailableWanPortOpenApiVO**](AvailableWanPortOpenApiVO.md) | available wan ports | [optional] 

## Methods

### NewAvailableWanResultOpenApiVO

`func NewAvailableWanResultOpenApiVO() *AvailableWanResultOpenApiVO`

NewAvailableWanResultOpenApiVO instantiates a new AvailableWanResultOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableWanResultOpenApiVOWithDefaults

`func NewAvailableWanResultOpenApiVOWithDefaults() *AvailableWanResultOpenApiVO`

NewAvailableWanResultOpenApiVOWithDefaults instantiates a new AvailableWanResultOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *AvailableWanResultOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *AvailableWanResultOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *AvailableWanResultOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *AvailableWanResultOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetWanPorts

`func (o *AvailableWanResultOpenApiVO) GetWanPorts() []AvailableWanPortOpenApiVO`

GetWanPorts returns the WanPorts field if non-nil, zero value otherwise.

### GetWanPortsOk

`func (o *AvailableWanResultOpenApiVO) GetWanPortsOk() (*[]AvailableWanPortOpenApiVO, bool)`

GetWanPortsOk returns a tuple with the WanPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPorts

`func (o *AvailableWanResultOpenApiVO) SetWanPorts(v []AvailableWanPortOpenApiVO)`

SetWanPorts sets WanPorts field to given value.

### HasWanPorts

`func (o *AvailableWanResultOpenApiVO) HasWanPorts() bool`

HasWanPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


