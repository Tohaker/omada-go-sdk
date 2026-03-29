# AlertEmailSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertEmailEnable** | **bool** | Enable alert emails | 
**Delay** | Pointer to **int32** | Send similar alerts within x seconds in one email, Note that when the number of incidents reaches 100, the log will be sent immediately | [optional] 
**DelayEnable** | Pointer to **bool** | Enable alert emails delay | [optional] 

## Methods

### NewAlertEmailSettingVO

`func NewAlertEmailSettingVO(alertEmailEnable bool, ) *AlertEmailSettingVO`

NewAlertEmailSettingVO instantiates a new AlertEmailSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertEmailSettingVOWithDefaults

`func NewAlertEmailSettingVOWithDefaults() *AlertEmailSettingVO`

NewAlertEmailSettingVOWithDefaults instantiates a new AlertEmailSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertEmailEnable

`func (o *AlertEmailSettingVO) GetAlertEmailEnable() bool`

GetAlertEmailEnable returns the AlertEmailEnable field if non-nil, zero value otherwise.

### GetAlertEmailEnableOk

`func (o *AlertEmailSettingVO) GetAlertEmailEnableOk() (*bool, bool)`

GetAlertEmailEnableOk returns a tuple with the AlertEmailEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmailEnable

`func (o *AlertEmailSettingVO) SetAlertEmailEnable(v bool)`

SetAlertEmailEnable sets AlertEmailEnable field to given value.


### GetDelay

`func (o *AlertEmailSettingVO) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *AlertEmailSettingVO) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *AlertEmailSettingVO) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *AlertEmailSettingVO) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetDelayEnable

`func (o *AlertEmailSettingVO) GetDelayEnable() bool`

GetDelayEnable returns the DelayEnable field if non-nil, zero value otherwise.

### GetDelayEnableOk

`func (o *AlertEmailSettingVO) GetDelayEnableOk() (*bool, bool)`

GetDelayEnableOk returns a tuple with the DelayEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayEnable

`func (o *AlertEmailSettingVO) SetDelayEnable(v bool)`

SetDelayEnable sets DelayEnable field to given value.

### HasDelayEnable

`func (o *AlertEmailSettingVO) HasDelayEnable() bool`

HasDelayEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


