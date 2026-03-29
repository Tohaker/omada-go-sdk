# RaSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | The switch of Ra | [optional] 
**Preference** | Pointer to **int32** | Preference should be a value as follows: 0: \&quot;low\&quot;; 1: \&quot;medium\&quot;; 2: \&quot;high\&quot; | [optional] 
**PreferredLifetime** | Pointer to **int32** | PreferredLifetime | [optional] 
**ValidLifetime** | Pointer to **int32** | ValidLifetime should be larger than PreferredLifetime. | [optional] 

## Methods

### NewRaSetting

`func NewRaSetting() *RaSetting`

NewRaSetting instantiates a new RaSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRaSettingWithDefaults

`func NewRaSettingWithDefaults() *RaSetting`

NewRaSettingWithDefaults instantiates a new RaSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *RaSetting) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *RaSetting) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *RaSetting) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *RaSetting) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetPreference

`func (o *RaSetting) GetPreference() int32`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *RaSetting) GetPreferenceOk() (*int32, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *RaSetting) SetPreference(v int32)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *RaSetting) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### GetPreferredLifetime

`func (o *RaSetting) GetPreferredLifetime() int32`

GetPreferredLifetime returns the PreferredLifetime field if non-nil, zero value otherwise.

### GetPreferredLifetimeOk

`func (o *RaSetting) GetPreferredLifetimeOk() (*int32, bool)`

GetPreferredLifetimeOk returns a tuple with the PreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLifetime

`func (o *RaSetting) SetPreferredLifetime(v int32)`

SetPreferredLifetime sets PreferredLifetime field to given value.

### HasPreferredLifetime

`func (o *RaSetting) HasPreferredLifetime() bool`

HasPreferredLifetime returns a boolean if a field has been set.

### GetValidLifetime

`func (o *RaSetting) GetValidLifetime() int32`

GetValidLifetime returns the ValidLifetime field if non-nil, zero value otherwise.

### GetValidLifetimeOk

`func (o *RaSetting) GetValidLifetimeOk() (*int32, bool)`

GetValidLifetimeOk returns a tuple with the ValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidLifetime

`func (o *RaSetting) SetValidLifetime(v int32)`

SetValidLifetime sets ValidLifetime field to given value.

### HasValidLifetime

`func (o *RaSetting) HasValidLifetime() bool`

HasValidLifetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


