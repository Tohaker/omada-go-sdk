# SmsSettingResOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthTimeout** | Pointer to [**AuthTimeOpenApiVO**](AuthTimeOpenApiVO.md) |  | [optional] 
**AuthToken** | Pointer to **string** | Twilio auth token | [optional] 
**CountryCode** | Pointer to **string** | Preset Contry code. String value such as \&quot;+86\&quot;. | [optional] 
**MaxVerificationCodeEnable** | Pointer to **bool** | Whether to control the limit of authentication for the same phone number. | [optional] 
**MaxVerificationCodeTimes** | Pointer to **int32** | User limit with the same phone number, should be within the range of 1–10. Required when parameter [maxVerificationCodeEnable] is true. | [optional] 
**PhoneNum** | Pointer to **string** | Twilio phone number. String value, should contain at least 6 digits such as \&quot;+123456\&quot;. | [optional] 
**Sid** | Pointer to **string** | Twilio SID | [optional] 

## Methods

### NewSmsSettingResOpenApiVO

`func NewSmsSettingResOpenApiVO() *SmsSettingResOpenApiVO`

NewSmsSettingResOpenApiVO instantiates a new SmsSettingResOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsSettingResOpenApiVOWithDefaults

`func NewSmsSettingResOpenApiVOWithDefaults() *SmsSettingResOpenApiVO`

NewSmsSettingResOpenApiVOWithDefaults instantiates a new SmsSettingResOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthTimeout

`func (o *SmsSettingResOpenApiVO) GetAuthTimeout() AuthTimeOpenApiVO`

GetAuthTimeout returns the AuthTimeout field if non-nil, zero value otherwise.

### GetAuthTimeoutOk

`func (o *SmsSettingResOpenApiVO) GetAuthTimeoutOk() (*AuthTimeOpenApiVO, bool)`

GetAuthTimeoutOk returns a tuple with the AuthTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTimeout

`func (o *SmsSettingResOpenApiVO) SetAuthTimeout(v AuthTimeOpenApiVO)`

SetAuthTimeout sets AuthTimeout field to given value.

### HasAuthTimeout

`func (o *SmsSettingResOpenApiVO) HasAuthTimeout() bool`

HasAuthTimeout returns a boolean if a field has been set.

### GetAuthToken

`func (o *SmsSettingResOpenApiVO) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *SmsSettingResOpenApiVO) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *SmsSettingResOpenApiVO) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *SmsSettingResOpenApiVO) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetCountryCode

`func (o *SmsSettingResOpenApiVO) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *SmsSettingResOpenApiVO) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *SmsSettingResOpenApiVO) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *SmsSettingResOpenApiVO) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetMaxVerificationCodeEnable

`func (o *SmsSettingResOpenApiVO) GetMaxVerificationCodeEnable() bool`

GetMaxVerificationCodeEnable returns the MaxVerificationCodeEnable field if non-nil, zero value otherwise.

### GetMaxVerificationCodeEnableOk

`func (o *SmsSettingResOpenApiVO) GetMaxVerificationCodeEnableOk() (*bool, bool)`

GetMaxVerificationCodeEnableOk returns a tuple with the MaxVerificationCodeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVerificationCodeEnable

`func (o *SmsSettingResOpenApiVO) SetMaxVerificationCodeEnable(v bool)`

SetMaxVerificationCodeEnable sets MaxVerificationCodeEnable field to given value.

### HasMaxVerificationCodeEnable

`func (o *SmsSettingResOpenApiVO) HasMaxVerificationCodeEnable() bool`

HasMaxVerificationCodeEnable returns a boolean if a field has been set.

### GetMaxVerificationCodeTimes

`func (o *SmsSettingResOpenApiVO) GetMaxVerificationCodeTimes() int32`

GetMaxVerificationCodeTimes returns the MaxVerificationCodeTimes field if non-nil, zero value otherwise.

### GetMaxVerificationCodeTimesOk

`func (o *SmsSettingResOpenApiVO) GetMaxVerificationCodeTimesOk() (*int32, bool)`

GetMaxVerificationCodeTimesOk returns a tuple with the MaxVerificationCodeTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVerificationCodeTimes

`func (o *SmsSettingResOpenApiVO) SetMaxVerificationCodeTimes(v int32)`

SetMaxVerificationCodeTimes sets MaxVerificationCodeTimes field to given value.

### HasMaxVerificationCodeTimes

`func (o *SmsSettingResOpenApiVO) HasMaxVerificationCodeTimes() bool`

HasMaxVerificationCodeTimes returns a boolean if a field has been set.

### GetPhoneNum

`func (o *SmsSettingResOpenApiVO) GetPhoneNum() string`

GetPhoneNum returns the PhoneNum field if non-nil, zero value otherwise.

### GetPhoneNumOk

`func (o *SmsSettingResOpenApiVO) GetPhoneNumOk() (*string, bool)`

GetPhoneNumOk returns a tuple with the PhoneNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNum

`func (o *SmsSettingResOpenApiVO) SetPhoneNum(v string)`

SetPhoneNum sets PhoneNum field to given value.

### HasPhoneNum

`func (o *SmsSettingResOpenApiVO) HasPhoneNum() bool`

HasPhoneNum returns a boolean if a field has been set.

### GetSid

`func (o *SmsSettingResOpenApiVO) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SmsSettingResOpenApiVO) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SmsSettingResOpenApiVO) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *SmsSettingResOpenApiVO) HasSid() bool`

HasSid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


