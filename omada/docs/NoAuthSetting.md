# NoAuthSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DailyLimitEnable** | **bool** | If enabled, authentication can only be performed once a day. | 

## Methods

### NewNoAuthSetting

`func NewNoAuthSetting(dailyLimitEnable bool, ) *NoAuthSetting`

NewNoAuthSetting instantiates a new NoAuthSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoAuthSettingWithDefaults

`func NewNoAuthSettingWithDefaults() *NoAuthSetting`

NewNoAuthSettingWithDefaults instantiates a new NoAuthSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDailyLimitEnable

`func (o *NoAuthSetting) GetDailyLimitEnable() bool`

GetDailyLimitEnable returns the DailyLimitEnable field if non-nil, zero value otherwise.

### GetDailyLimitEnableOk

`func (o *NoAuthSetting) GetDailyLimitEnableOk() (*bool, bool)`

GetDailyLimitEnableOk returns a tuple with the DailyLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLimitEnable

`func (o *NoAuthSetting) SetDailyLimitEnable(v bool)`

SetDailyLimitEnable sets DailyLimitEnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


