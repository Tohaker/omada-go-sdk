# RrmSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **int32** | The mode of Auto WLAN Optimization, such as: 0: disable, 1: adaptive | 
**Resource** | Pointer to **int32** | The anomaly event setting creation resource, such as: 0: new created, 1: from template, 2: override | [optional] 
**TimeRangeAction** | Pointer to **int32** | The Time Range Action, such as: 0: out of range, 1: in range | [optional] 
**TimeRangeEnable** | **bool** | Whether the Time Range is enabled. True: enable, false: disable. | 
**TimeRangeId** | Pointer to **string** | This field represents Time Range Profile ID. Time Range Profile can be created using Create time range profile interface, and Time Range Profile ID can be obtained from Get time range profile list interface. | [optional] 

## Methods

### NewRrmSettingOpenApiVO

`func NewRrmSettingOpenApiVO(mode int32, timeRangeEnable bool, ) *RrmSettingOpenApiVO`

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

### GetTimeRangeAction

`func (o *RrmSettingOpenApiVO) GetTimeRangeAction() int32`

GetTimeRangeAction returns the TimeRangeAction field if non-nil, zero value otherwise.

### GetTimeRangeActionOk

`func (o *RrmSettingOpenApiVO) GetTimeRangeActionOk() (*int32, bool)`

GetTimeRangeActionOk returns a tuple with the TimeRangeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRangeAction

`func (o *RrmSettingOpenApiVO) SetTimeRangeAction(v int32)`

SetTimeRangeAction sets TimeRangeAction field to given value.

### HasTimeRangeAction

`func (o *RrmSettingOpenApiVO) HasTimeRangeAction() bool`

HasTimeRangeAction returns a boolean if a field has been set.

### GetTimeRangeEnable

`func (o *RrmSettingOpenApiVO) GetTimeRangeEnable() bool`

GetTimeRangeEnable returns the TimeRangeEnable field if non-nil, zero value otherwise.

### GetTimeRangeEnableOk

`func (o *RrmSettingOpenApiVO) GetTimeRangeEnableOk() (*bool, bool)`

GetTimeRangeEnableOk returns a tuple with the TimeRangeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRangeEnable

`func (o *RrmSettingOpenApiVO) SetTimeRangeEnable(v bool)`

SetTimeRangeEnable sets TimeRangeEnable field to given value.


### GetTimeRangeId

`func (o *RrmSettingOpenApiVO) GetTimeRangeId() string`

GetTimeRangeId returns the TimeRangeId field if non-nil, zero value otherwise.

### GetTimeRangeIdOk

`func (o *RrmSettingOpenApiVO) GetTimeRangeIdOk() (*string, bool)`

GetTimeRangeIdOk returns a tuple with the TimeRangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRangeId

`func (o *RrmSettingOpenApiVO) SetTimeRangeId(v string)`

SetTimeRangeId sets TimeRangeId field to given value.

### HasTimeRangeId

`func (o *RrmSettingOpenApiVO) HasTimeRangeId() bool`

HasTimeRangeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


