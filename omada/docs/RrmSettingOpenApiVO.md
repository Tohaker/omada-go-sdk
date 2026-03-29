# RrmSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **int32** | The mode of Auto WLAN Optimization. | 
**Resource** | Pointer to **int32** | The anomaly event setting creation resource, such as: 0: new created, 1: from template, 2: override | [optional] 

## Methods

### NewRrmSettingOpenApiVO

`func NewRrmSettingOpenApiVO(mode int32, ) *RrmSettingOpenApiVO`

NewRrmSettingOpenApiVO instantiates a new RrmSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRrmSettingOpenApiVOWithDefaults

`func NewRrmSettingOpenApiVOWithDefaults() *RrmSettingOpenApiVO`

NewRrmSettingOpenApiVOWithDefaults instantiates a new RrmSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *RrmSettingOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RrmSettingOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RrmSettingOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetResource

`func (o *RrmSettingOpenApiVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *RrmSettingOpenApiVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *RrmSettingOpenApiVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *RrmSettingOpenApiVO) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


