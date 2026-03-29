# AuthTimeoutSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomTimeout** | Pointer to **int32** | Custom timeout should be within the range of 1 - 1,000,000 min or 1 - 10,000 hour or 1 - 1,000 day. | [optional] 
**CustomTimeoutUnit** | Pointer to **int32** | Timeout unit, should be a value as follows: 1: min; 2: hour; 3: day. | [optional] 

## Methods

### NewAuthTimeoutSetting

`func NewAuthTimeoutSetting() *AuthTimeoutSetting`

NewAuthTimeoutSetting instantiates a new AuthTimeoutSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTimeoutSettingWithDefaults

`func NewAuthTimeoutSettingWithDefaults() *AuthTimeoutSetting`

NewAuthTimeoutSettingWithDefaults instantiates a new AuthTimeoutSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomTimeout

`func (o *AuthTimeoutSetting) GetCustomTimeout() int32`

GetCustomTimeout returns the CustomTimeout field if non-nil, zero value otherwise.

### GetCustomTimeoutOk

`func (o *AuthTimeoutSetting) GetCustomTimeoutOk() (*int32, bool)`

GetCustomTimeoutOk returns a tuple with the CustomTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTimeout

`func (o *AuthTimeoutSetting) SetCustomTimeout(v int32)`

SetCustomTimeout sets CustomTimeout field to given value.

### HasCustomTimeout

`func (o *AuthTimeoutSetting) HasCustomTimeout() bool`

HasCustomTimeout returns a boolean if a field has been set.

### GetCustomTimeoutUnit

`func (o *AuthTimeoutSetting) GetCustomTimeoutUnit() int32`

GetCustomTimeoutUnit returns the CustomTimeoutUnit field if non-nil, zero value otherwise.

### GetCustomTimeoutUnitOk

`func (o *AuthTimeoutSetting) GetCustomTimeoutUnitOk() (*int32, bool)`

GetCustomTimeoutUnitOk returns a tuple with the CustomTimeoutUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTimeoutUnit

`func (o *AuthTimeoutSetting) SetCustomTimeoutUnit(v int32)`

SetCustomTimeoutUnit sets CustomTimeoutUnit field to given value.

### HasCustomTimeoutUnit

`func (o *AuthTimeoutSetting) HasCustomTimeoutUnit() bool`

HasCustomTimeoutUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


