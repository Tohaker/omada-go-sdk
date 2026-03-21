# PowerThresholdVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **int32** | Parameter [mode] should be 0 or 1. 0: Auto. 1: Custom. | [optional] 
**Threshold2g** | Pointer to **int32** | Power threshold in 2.4 GHz. Parameter [threshold2g] should range from -75 and -60. | [optional] 
**Threshold5g** | Pointer to **int32** | Power threshold in 5 GHz. Parameter [threshold5g] should range from -75 and -60. | [optional] 
**Threshold6g** | Pointer to **int32** | Power threshold in 6 GHz. Parameter [threshold6g] should range from -75 and -60. | [optional] 

## Methods

### NewPowerThresholdVO

`func NewPowerThresholdVO() *PowerThresholdVO`

NewPowerThresholdVO instantiates a new PowerThresholdVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerThresholdVOWithDefaults

`func NewPowerThresholdVOWithDefaults() *PowerThresholdVO`

NewPowerThresholdVOWithDefaults instantiates a new PowerThresholdVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *PowerThresholdVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PowerThresholdVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PowerThresholdVO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PowerThresholdVO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetThreshold2g

`func (o *PowerThresholdVO) GetThreshold2g() int32`

GetThreshold2g returns the Threshold2g field if non-nil, zero value otherwise.

### GetThreshold2gOk

`func (o *PowerThresholdVO) GetThreshold2gOk() (*int32, bool)`

GetThreshold2gOk returns a tuple with the Threshold2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold2g

`func (o *PowerThresholdVO) SetThreshold2g(v int32)`

SetThreshold2g sets Threshold2g field to given value.

### HasThreshold2g

`func (o *PowerThresholdVO) HasThreshold2g() bool`

HasThreshold2g returns a boolean if a field has been set.

### GetThreshold5g

`func (o *PowerThresholdVO) GetThreshold5g() int32`

GetThreshold5g returns the Threshold5g field if non-nil, zero value otherwise.

### GetThreshold5gOk

`func (o *PowerThresholdVO) GetThreshold5gOk() (*int32, bool)`

GetThreshold5gOk returns a tuple with the Threshold5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold5g

`func (o *PowerThresholdVO) SetThreshold5g(v int32)`

SetThreshold5g sets Threshold5g field to given value.

### HasThreshold5g

`func (o *PowerThresholdVO) HasThreshold5g() bool`

HasThreshold5g returns a boolean if a field has been set.

### GetThreshold6g

`func (o *PowerThresholdVO) GetThreshold6g() int32`

GetThreshold6g returns the Threshold6g field if non-nil, zero value otherwise.

### GetThreshold6gOk

`func (o *PowerThresholdVO) GetThreshold6gOk() (*int32, bool)`

GetThreshold6gOk returns a tuple with the Threshold6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold6g

`func (o *PowerThresholdVO) SetThreshold6g(v int32)`

SetThreshold6g sets Threshold6g field to given value.

### HasThreshold6g

`func (o *PowerThresholdVO) HasThreshold6g() bool`

HasThreshold6g returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


