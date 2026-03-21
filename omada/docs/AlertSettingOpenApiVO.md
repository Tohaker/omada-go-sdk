# AlertSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delay** | Pointer to **int32** | Time of Log delay email (unit:s). The value should be within the range of 0–99999. | [optional] 
**DelayEnable** | Pointer to **bool** | Alert delay enable. | [optional] 
**Enable** | **bool** | Alert enable. | 

## Methods

### NewAlertSettingOpenApiVO

`func NewAlertSettingOpenApiVO(enable bool, ) *AlertSettingOpenApiVO`

NewAlertSettingOpenApiVO instantiates a new AlertSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertSettingOpenApiVOWithDefaults

`func NewAlertSettingOpenApiVOWithDefaults() *AlertSettingOpenApiVO`

NewAlertSettingOpenApiVOWithDefaults instantiates a new AlertSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelay

`func (o *AlertSettingOpenApiVO) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *AlertSettingOpenApiVO) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *AlertSettingOpenApiVO) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *AlertSettingOpenApiVO) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetDelayEnable

`func (o *AlertSettingOpenApiVO) GetDelayEnable() bool`

GetDelayEnable returns the DelayEnable field if non-nil, zero value otherwise.

### GetDelayEnableOk

`func (o *AlertSettingOpenApiVO) GetDelayEnableOk() (*bool, bool)`

GetDelayEnableOk returns a tuple with the DelayEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayEnable

`func (o *AlertSettingOpenApiVO) SetDelayEnable(v bool)`

SetDelayEnable sets DelayEnable field to given value.

### HasDelayEnable

`func (o *AlertSettingOpenApiVO) HasDelayEnable() bool`

HasDelayEnable returns a boolean if a field has been set.

### GetEnable

`func (o *AlertSettingOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *AlertSettingOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *AlertSettingOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


