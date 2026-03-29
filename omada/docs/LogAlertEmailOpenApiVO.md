# LogAlertEmailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertEmailEnable** | Pointer to **bool** | Log Enable or Disable email | [optional] 
**Delay** | Pointer to **int32** | Time of Log delay email (unit:s). The value should be within the range of 0–99999. | [optional] 
**DelayEnable** | Pointer to **bool** | Log Enable or Disable delay email | [optional] 

## Methods

### NewLogAlertEmailOpenApiVO

`func NewLogAlertEmailOpenApiVO() *LogAlertEmailOpenApiVO`

NewLogAlertEmailOpenApiVO instantiates a new LogAlertEmailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogAlertEmailOpenApiVOWithDefaults

`func NewLogAlertEmailOpenApiVOWithDefaults() *LogAlertEmailOpenApiVO`

NewLogAlertEmailOpenApiVOWithDefaults instantiates a new LogAlertEmailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertEmailEnable

`func (o *LogAlertEmailOpenApiVO) GetAlertEmailEnable() bool`

GetAlertEmailEnable returns the AlertEmailEnable field if non-nil, zero value otherwise.

### GetAlertEmailEnableOk

`func (o *LogAlertEmailOpenApiVO) GetAlertEmailEnableOk() (*bool, bool)`

GetAlertEmailEnableOk returns a tuple with the AlertEmailEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmailEnable

`func (o *LogAlertEmailOpenApiVO) SetAlertEmailEnable(v bool)`

SetAlertEmailEnable sets AlertEmailEnable field to given value.

### HasAlertEmailEnable

`func (o *LogAlertEmailOpenApiVO) HasAlertEmailEnable() bool`

HasAlertEmailEnable returns a boolean if a field has been set.

### GetDelay

`func (o *LogAlertEmailOpenApiVO) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *LogAlertEmailOpenApiVO) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *LogAlertEmailOpenApiVO) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *LogAlertEmailOpenApiVO) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetDelayEnable

`func (o *LogAlertEmailOpenApiVO) GetDelayEnable() bool`

GetDelayEnable returns the DelayEnable field if non-nil, zero value otherwise.

### GetDelayEnableOk

`func (o *LogAlertEmailOpenApiVO) GetDelayEnableOk() (*bool, bool)`

GetDelayEnableOk returns a tuple with the DelayEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayEnable

`func (o *LogAlertEmailOpenApiVO) SetDelayEnable(v bool)`

SetDelayEnable sets DelayEnable field to given value.

### HasDelayEnable

`func (o *LogAlertEmailOpenApiVO) HasDelayEnable() bool`

HasDelayEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


