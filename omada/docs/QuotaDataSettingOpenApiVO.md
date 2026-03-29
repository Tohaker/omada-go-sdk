# QuotaDataSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | Pointer to **bool** | SMS alert for usage is enabled/disabled. | [optional] 
**CallingCode** | Pointer to **string** | Calling code should contain 2 to 5 characters. Calling code must be entered when entering the country code. For the values of Calling code, refer to section 5.4.1 of the Open API Access Guide. | [optional] 
**CountryCode** | Pointer to **string** | Country code should contain 2 characters. Country code must be entered when entering the calling code. For the values of Country code, refer to section 5.4.1 of the Open API Access Guide. | [optional] 
**Credit** | Pointer to **float32** | The amount of data allowance in current billing cycle. | [optional] 
**CreditUnit** | Pointer to **int32** | Unit for data allowance should be a value as follows: 2:MB; 3:GB. | [optional] 
**Limit** | **bool** | Quota limit is enabled/disabled. | 
**Phone** | Pointer to **string** | The phone number to receive SMS alerts. | [optional] 
**StartDate** | Pointer to **int32** | Start date of monthly billing cycle type, valid date should be within the range of 1–31. | [optional] 
**Type** | **int32** | Billing cycle type should be a value as follows: 0:total; 1:monthly. | 
**Usage** | Pointer to **int32** | SMS alert for usage when reach percentage of allowance, valid date should be within the range of 0–100. | [optional] 
**Used** | Pointer to **float32** | The amount of data usage (KB) in current billing cycle. | [optional] 

## Methods

### NewQuotaDataSettingOpenApiVO

`func NewQuotaDataSettingOpenApiVO(limit bool, type_ int32, ) *QuotaDataSettingOpenApiVO`

NewQuotaDataSettingOpenApiVO instantiates a new QuotaDataSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaDataSettingOpenApiVOWithDefaults

`func NewQuotaDataSettingOpenApiVOWithDefaults() *QuotaDataSettingOpenApiVO`

NewQuotaDataSettingOpenApiVOWithDefaults instantiates a new QuotaDataSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *QuotaDataSettingOpenApiVO) GetAlert() bool`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *QuotaDataSettingOpenApiVO) GetAlertOk() (*bool, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *QuotaDataSettingOpenApiVO) SetAlert(v bool)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *QuotaDataSettingOpenApiVO) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetCallingCode

`func (o *QuotaDataSettingOpenApiVO) GetCallingCode() string`

GetCallingCode returns the CallingCode field if non-nil, zero value otherwise.

### GetCallingCodeOk

`func (o *QuotaDataSettingOpenApiVO) GetCallingCodeOk() (*string, bool)`

GetCallingCodeOk returns a tuple with the CallingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallingCode

`func (o *QuotaDataSettingOpenApiVO) SetCallingCode(v string)`

SetCallingCode sets CallingCode field to given value.

### HasCallingCode

`func (o *QuotaDataSettingOpenApiVO) HasCallingCode() bool`

HasCallingCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *QuotaDataSettingOpenApiVO) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *QuotaDataSettingOpenApiVO) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *QuotaDataSettingOpenApiVO) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *QuotaDataSettingOpenApiVO) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCredit

`func (o *QuotaDataSettingOpenApiVO) GetCredit() float32`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *QuotaDataSettingOpenApiVO) GetCreditOk() (*float32, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *QuotaDataSettingOpenApiVO) SetCredit(v float32)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *QuotaDataSettingOpenApiVO) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetCreditUnit

`func (o *QuotaDataSettingOpenApiVO) GetCreditUnit() int32`

GetCreditUnit returns the CreditUnit field if non-nil, zero value otherwise.

### GetCreditUnitOk

`func (o *QuotaDataSettingOpenApiVO) GetCreditUnitOk() (*int32, bool)`

GetCreditUnitOk returns a tuple with the CreditUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditUnit

`func (o *QuotaDataSettingOpenApiVO) SetCreditUnit(v int32)`

SetCreditUnit sets CreditUnit field to given value.

### HasCreditUnit

`func (o *QuotaDataSettingOpenApiVO) HasCreditUnit() bool`

HasCreditUnit returns a boolean if a field has been set.

### GetLimit

`func (o *QuotaDataSettingOpenApiVO) GetLimit() bool`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QuotaDataSettingOpenApiVO) GetLimitOk() (*bool, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QuotaDataSettingOpenApiVO) SetLimit(v bool)`

SetLimit sets Limit field to given value.


### GetPhone

`func (o *QuotaDataSettingOpenApiVO) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *QuotaDataSettingOpenApiVO) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *QuotaDataSettingOpenApiVO) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *QuotaDataSettingOpenApiVO) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetStartDate

`func (o *QuotaDataSettingOpenApiVO) GetStartDate() int32`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *QuotaDataSettingOpenApiVO) GetStartDateOk() (*int32, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *QuotaDataSettingOpenApiVO) SetStartDate(v int32)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *QuotaDataSettingOpenApiVO) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetType

`func (o *QuotaDataSettingOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QuotaDataSettingOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QuotaDataSettingOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.


### GetUsage

`func (o *QuotaDataSettingOpenApiVO) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *QuotaDataSettingOpenApiVO) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *QuotaDataSettingOpenApiVO) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *QuotaDataSettingOpenApiVO) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUsed

`func (o *QuotaDataSettingOpenApiVO) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *QuotaDataSettingOpenApiVO) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *QuotaDataSettingOpenApiVO) SetUsed(v float32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *QuotaDataSettingOpenApiVO) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


