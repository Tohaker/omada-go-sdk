# QuotaSmsSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | Pointer to **bool** | SMS alert for usage is enabled/disabled. | [optional] 
**CallingCode** | Pointer to **string** | Calling code should contain 2 to 5 characters. Calling code must be entered when entering the country code. For the values of Calling code, refer to section 5.4.1 of the Open API Access Guide. | [optional] 
**CountryCode** | Pointer to **string** | Country code should contain 2 characters. Country code must be entered when entering the calling code. For the values of Country code, refer to section 5.4.1 of the Open API Access Guide. | [optional] 
**Credit** | Pointer to **int32** | The amount of data allowance in current billing cycle, valid date should be within the range of 0–100000. | [optional] 
**Limit** | **bool** | Quota limit is enabled/disabled. | 
**Phone** | Pointer to **string** | The phone number to receive SMS alerts. | [optional] 
**StartDate** | Pointer to **int32** | Start date of monthly billing cycle type, valid date should be within the range of 1–31. | [optional] 
**Type** | **int32** | Billing cycle type should be a value as follows: 0:total; 1:monthly. | 
**Usage** | Pointer to **int32** | SMS alert for usage when reach percentage of allowance, valid date should be within the range of 0–100. | [optional] 
**Used** | Pointer to **int32** | The amount of SMS in current billing cycle. | [optional] 

## Methods

### NewQuotaSmsSettingOpenApiVO

`func NewQuotaSmsSettingOpenApiVO(limit bool, type_ int32, ) *QuotaSmsSettingOpenApiVO`

NewQuotaSmsSettingOpenApiVO instantiates a new QuotaSmsSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaSmsSettingOpenApiVOWithDefaults

`func NewQuotaSmsSettingOpenApiVOWithDefaults() *QuotaSmsSettingOpenApiVO`

NewQuotaSmsSettingOpenApiVOWithDefaults instantiates a new QuotaSmsSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *QuotaSmsSettingOpenApiVO) GetAlert() bool`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *QuotaSmsSettingOpenApiVO) GetAlertOk() (*bool, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *QuotaSmsSettingOpenApiVO) SetAlert(v bool)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *QuotaSmsSettingOpenApiVO) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetCallingCode

`func (o *QuotaSmsSettingOpenApiVO) GetCallingCode() string`

GetCallingCode returns the CallingCode field if non-nil, zero value otherwise.

### GetCallingCodeOk

`func (o *QuotaSmsSettingOpenApiVO) GetCallingCodeOk() (*string, bool)`

GetCallingCodeOk returns a tuple with the CallingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallingCode

`func (o *QuotaSmsSettingOpenApiVO) SetCallingCode(v string)`

SetCallingCode sets CallingCode field to given value.

### HasCallingCode

`func (o *QuotaSmsSettingOpenApiVO) HasCallingCode() bool`

HasCallingCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *QuotaSmsSettingOpenApiVO) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *QuotaSmsSettingOpenApiVO) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *QuotaSmsSettingOpenApiVO) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *QuotaSmsSettingOpenApiVO) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCredit

`func (o *QuotaSmsSettingOpenApiVO) GetCredit() int32`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *QuotaSmsSettingOpenApiVO) GetCreditOk() (*int32, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *QuotaSmsSettingOpenApiVO) SetCredit(v int32)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *QuotaSmsSettingOpenApiVO) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetLimit

`func (o *QuotaSmsSettingOpenApiVO) GetLimit() bool`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QuotaSmsSettingOpenApiVO) GetLimitOk() (*bool, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QuotaSmsSettingOpenApiVO) SetLimit(v bool)`

SetLimit sets Limit field to given value.


### GetPhone

`func (o *QuotaSmsSettingOpenApiVO) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *QuotaSmsSettingOpenApiVO) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *QuotaSmsSettingOpenApiVO) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *QuotaSmsSettingOpenApiVO) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetStartDate

`func (o *QuotaSmsSettingOpenApiVO) GetStartDate() int32`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *QuotaSmsSettingOpenApiVO) GetStartDateOk() (*int32, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *QuotaSmsSettingOpenApiVO) SetStartDate(v int32)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *QuotaSmsSettingOpenApiVO) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetType

`func (o *QuotaSmsSettingOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QuotaSmsSettingOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QuotaSmsSettingOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.


### GetUsage

`func (o *QuotaSmsSettingOpenApiVO) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *QuotaSmsSettingOpenApiVO) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *QuotaSmsSettingOpenApiVO) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *QuotaSmsSettingOpenApiVO) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUsed

`func (o *QuotaSmsSettingOpenApiVO) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *QuotaSmsSettingOpenApiVO) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *QuotaSmsSettingOpenApiVO) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *QuotaSmsSettingOpenApiVO) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


