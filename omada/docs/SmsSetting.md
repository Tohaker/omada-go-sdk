# SmsSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthTimeout** | [**AuthTimeoutSetting**](AuthTimeoutSetting.md) |  | 
**AuthToken** | **string** | Twilio auth token | 
**CountryCode** | Pointer to **string** | Preset Contry code. String value such as \&quot;+86\&quot;. | [optional] 
**PhoneNum** | **string** | Twilio phone number. String value, should contain at least 6 digits such as \&quot;+123456\&quot;. | 
**Sid** | **string** | Twilio SID | 
**UserLimit** | Pointer to **int32** | User limit with the same phone number, should be within the range of 1–10. Required when parameter [userLimitEnable] is true. | [optional] 
**UserLimitEnable** | **bool** | Whether to control the limit of authentication for the same phone number. | 

## Methods

### NewSmsSetting

`func NewSmsSetting(authTimeout AuthTimeoutSetting, authToken string, phoneNum string, sid string, userLimitEnable bool, ) *SmsSetting`

NewSmsSetting instantiates a new SmsSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsSettingWithDefaults

`func NewSmsSettingWithDefaults() *SmsSetting`

NewSmsSettingWithDefaults instantiates a new SmsSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthTimeout

`func (o *SmsSetting) GetAuthTimeout() AuthTimeoutSetting`

GetAuthTimeout returns the AuthTimeout field if non-nil, zero value otherwise.

### GetAuthTimeoutOk

`func (o *SmsSetting) GetAuthTimeoutOk() (*AuthTimeoutSetting, bool)`

GetAuthTimeoutOk returns a tuple with the AuthTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTimeout

`func (o *SmsSetting) SetAuthTimeout(v AuthTimeoutSetting)`

SetAuthTimeout sets AuthTimeout field to given value.


### GetAuthToken

`func (o *SmsSetting) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *SmsSetting) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *SmsSetting) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.


### GetCountryCode

`func (o *SmsSetting) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *SmsSetting) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *SmsSetting) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *SmsSetting) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhoneNum

`func (o *SmsSetting) GetPhoneNum() string`

GetPhoneNum returns the PhoneNum field if non-nil, zero value otherwise.

### GetPhoneNumOk

`func (o *SmsSetting) GetPhoneNumOk() (*string, bool)`

GetPhoneNumOk returns a tuple with the PhoneNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNum

`func (o *SmsSetting) SetPhoneNum(v string)`

SetPhoneNum sets PhoneNum field to given value.


### GetSid

`func (o *SmsSetting) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SmsSetting) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SmsSetting) SetSid(v string)`

SetSid sets Sid field to given value.


### GetUserLimit

`func (o *SmsSetting) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *SmsSetting) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *SmsSetting) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *SmsSetting) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### GetUserLimitEnable

`func (o *SmsSetting) GetUserLimitEnable() bool`

GetUserLimitEnable returns the UserLimitEnable field if non-nil, zero value otherwise.

### GetUserLimitEnableOk

`func (o *SmsSetting) GetUserLimitEnableOk() (*bool, bool)`

GetUserLimitEnableOk returns a tuple with the UserLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimitEnable

`func (o *SmsSetting) SetUserLimitEnable(v bool)`

SetUserLimitEnable sets UserLimitEnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


