# ApRssiLedSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LedId** | Pointer to **string** | Led Name. | [optional] 
**Threshold** | Pointer to **int32** | Rssi threshold. | [optional] 

## Methods

### NewApRssiLedSettingVO

`func NewApRssiLedSettingVO() *ApRssiLedSettingVO`

NewApRssiLedSettingVO instantiates a new ApRssiLedSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApRssiLedSettingVOWithDefaults

`func NewApRssiLedSettingVOWithDefaults() *ApRssiLedSettingVO`

NewApRssiLedSettingVOWithDefaults instantiates a new ApRssiLedSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLedId

`func (o *ApRssiLedSettingVO) GetLedId() string`

GetLedId returns the LedId field if non-nil, zero value otherwise.

### GetLedIdOk

`func (o *ApRssiLedSettingVO) GetLedIdOk() (*string, bool)`

GetLedIdOk returns a tuple with the LedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedId

`func (o *ApRssiLedSettingVO) SetLedId(v string)`

SetLedId sets LedId field to given value.

### HasLedId

`func (o *ApRssiLedSettingVO) HasLedId() bool`

HasLedId returns a boolean if a field has been set.

### GetThreshold

`func (o *ApRssiLedSettingVO) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ApRssiLedSettingVO) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ApRssiLedSettingVO) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ApRssiLedSettingVO) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


