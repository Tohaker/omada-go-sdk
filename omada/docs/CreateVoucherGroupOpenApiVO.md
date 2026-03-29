# CreateVoucherGroupOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The amount of vouchers created. It should be within the range of 1-5000 | 
**ApplyToAllPortals** | **bool** | Is the voucher effective for all portals, including all newly created portals | 
**CodeForm** | **[]int32** | The character types contained in the voucher code. It should be a value as follows: 0: Number, 1: Letter. For example, [0] indicates that the code only contains numbers; [0, 1] indicates that the code contains numbers and letters | 
**CodeLength** | **int32** | The length of voucher code. It should be within the range of 6–10. | 
**Currency** | Pointer to **string** | Currency Short Code of voucher. For the values of Currency Short Code, refer to section 5.4.2 of the Open API Access Guide. | [optional] 
**Description** | Pointer to **string** | Description of the voucher group | [optional] 
**Duration** | **int64** | Duration of one use, unit: minute. It should be within the range of 1–14400000. | 
**DurationType** | **int32** | The duration type of the voucher. It should be a value as follows: 0: Client duration, each client expires after the duration is used. 1: Voucher duration, after reaching the voucher duration, clients using the voucher will expire | 
**EffectiveTime** | Pointer to **int64** | The timestamp when the voucher takes effect, unit: millisecond. When parameter [validityType] is 1, parameter [effectiveTime] is required | [optional] 
**ExpirationTime** | Pointer to **int64** | The timestamp of the expiration of the voucher, unit: millisecond. When parameter [validityType] is 1, parameter [expirationTime] is required | [optional] 
**LimitNum** | Pointer to **int32** | The number of limitations. It should be within the range of 1–999. If Parameter [limitType] is 0 or 1, [limitNum] should not be null.When Parameter [limitType] is 0, [limitNum] represents the maximum number of times this voucher can be used.When Parameter [limitType] is 1, [limitNum] represents the maximum number of users this voucher can be used at the same time. | [optional] 
**LimitType** | **int32** | The limitations of the voucher. It should be a value as follows: 0: Limited Usage Counts, 1: Limited Online Users, 2: Unlimited | 
**Logout** | Pointer to **bool** | Whether the voucher support portal logout functionality | [optional] 
**Name** | **string** | Voucher group name. It should contain 1-32 characters | 
**Portals** | Pointer to **[]string** | Bound portal ID list. Portal can be created using &#39;Add portal&#39; interface, and portal ID can be obtained from &#39;Get portal list in a site&#39; interface | [optional] 
**PrintComments** | Pointer to **string** | Print comments of the voucher group | [optional] 
**RateLimit** | [**RateLimitOpenApiVO**](RateLimitOpenApiVO.md) |  | 
**Schedule** | Pointer to [**VoucherScheduleOpenApiVO**](VoucherScheduleOpenApiVO.md) |  | [optional] 
**TimingType** | **int32** | The timing type of the voucher. It should be a value as follows: 0: Timing by time, clients can use vouchers for specified time duration. 1: Timing by usage, clients can use vouchers for the duration of actual usage | 
**TrafficLimit** | Pointer to **int64** | Traffic limit in MB. It should be within the range of 1–10485760 | [optional] 
**TrafficLimitEnable** | **bool** | Whether to enable traffic limit | 
**TrafficLimitFrequency** | Pointer to **int32** | Frequency of traffic limit should be a value as follows: 0: total; 1: daily; 2: weekly; 3: monthly. | [optional] 
**UnitPrice** | Pointer to **int64** | Price of single voucher. It should be within the range of 1–999999999 | [optional] 
**ValidityType** | Pointer to **int32** | The validity type of the voucher. It should be a value as follows: 0: Voucher can be used at any time, parameter [effectiveTime], [expirationTime] and [schedule] should be null. 1: Voucher can be used between the effective time and expiration time, parameter [effectiveTime] and [expirationTime] should not be null, parameter [schedule] should be null. 2: Voucher can be used within a specified time period by schedule, parameter [effectiveTime] and [expirationTime] should be null, parameter [schedule] should not be null | [optional] 

## Methods

### NewCreateVoucherGroupOpenApiVO

`func NewCreateVoucherGroupOpenApiVO(amount int32, applyToAllPortals bool, codeForm []int32, codeLength int32, duration int64, durationType int32, limitType int32, name string, rateLimit RateLimitOpenApiVO, timingType int32, trafficLimitEnable bool, ) *CreateVoucherGroupOpenApiVO`

NewCreateVoucherGroupOpenApiVO instantiates a new CreateVoucherGroupOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVoucherGroupOpenApiVOWithDefaults

`func NewCreateVoucherGroupOpenApiVOWithDefaults() *CreateVoucherGroupOpenApiVO`

NewCreateVoucherGroupOpenApiVOWithDefaults instantiates a new CreateVoucherGroupOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CreateVoucherGroupOpenApiVO) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateVoucherGroupOpenApiVO) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateVoucherGroupOpenApiVO) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetApplyToAllPortals

`func (o *CreateVoucherGroupOpenApiVO) GetApplyToAllPortals() bool`

GetApplyToAllPortals returns the ApplyToAllPortals field if non-nil, zero value otherwise.

### GetApplyToAllPortalsOk

`func (o *CreateVoucherGroupOpenApiVO) GetApplyToAllPortalsOk() (*bool, bool)`

GetApplyToAllPortalsOk returns a tuple with the ApplyToAllPortals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToAllPortals

`func (o *CreateVoucherGroupOpenApiVO) SetApplyToAllPortals(v bool)`

SetApplyToAllPortals sets ApplyToAllPortals field to given value.


### GetCodeForm

`func (o *CreateVoucherGroupOpenApiVO) GetCodeForm() []int32`

GetCodeForm returns the CodeForm field if non-nil, zero value otherwise.

### GetCodeFormOk

`func (o *CreateVoucherGroupOpenApiVO) GetCodeFormOk() (*[]int32, bool)`

GetCodeFormOk returns a tuple with the CodeForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeForm

`func (o *CreateVoucherGroupOpenApiVO) SetCodeForm(v []int32)`

SetCodeForm sets CodeForm field to given value.


### GetCodeLength

`func (o *CreateVoucherGroupOpenApiVO) GetCodeLength() int32`

GetCodeLength returns the CodeLength field if non-nil, zero value otherwise.

### GetCodeLengthOk

`func (o *CreateVoucherGroupOpenApiVO) GetCodeLengthOk() (*int32, bool)`

GetCodeLengthOk returns a tuple with the CodeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeLength

`func (o *CreateVoucherGroupOpenApiVO) SetCodeLength(v int32)`

SetCodeLength sets CodeLength field to given value.


### GetCurrency

`func (o *CreateVoucherGroupOpenApiVO) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateVoucherGroupOpenApiVO) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateVoucherGroupOpenApiVO) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateVoucherGroupOpenApiVO) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *CreateVoucherGroupOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateVoucherGroupOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateVoucherGroupOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateVoucherGroupOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuration

`func (o *CreateVoucherGroupOpenApiVO) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CreateVoucherGroupOpenApiVO) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CreateVoucherGroupOpenApiVO) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetDurationType

`func (o *CreateVoucherGroupOpenApiVO) GetDurationType() int32`

GetDurationType returns the DurationType field if non-nil, zero value otherwise.

### GetDurationTypeOk

`func (o *CreateVoucherGroupOpenApiVO) GetDurationTypeOk() (*int32, bool)`

GetDurationTypeOk returns a tuple with the DurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationType

`func (o *CreateVoucherGroupOpenApiVO) SetDurationType(v int32)`

SetDurationType sets DurationType field to given value.


### GetEffectiveTime

`func (o *CreateVoucherGroupOpenApiVO) GetEffectiveTime() int64`

GetEffectiveTime returns the EffectiveTime field if non-nil, zero value otherwise.

### GetEffectiveTimeOk

`func (o *CreateVoucherGroupOpenApiVO) GetEffectiveTimeOk() (*int64, bool)`

GetEffectiveTimeOk returns a tuple with the EffectiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTime

`func (o *CreateVoucherGroupOpenApiVO) SetEffectiveTime(v int64)`

SetEffectiveTime sets EffectiveTime field to given value.

### HasEffectiveTime

`func (o *CreateVoucherGroupOpenApiVO) HasEffectiveTime() bool`

HasEffectiveTime returns a boolean if a field has been set.

### GetExpirationTime

`func (o *CreateVoucherGroupOpenApiVO) GetExpirationTime() int64`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *CreateVoucherGroupOpenApiVO) GetExpirationTimeOk() (*int64, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *CreateVoucherGroupOpenApiVO) SetExpirationTime(v int64)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *CreateVoucherGroupOpenApiVO) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetLimitNum

`func (o *CreateVoucherGroupOpenApiVO) GetLimitNum() int32`

GetLimitNum returns the LimitNum field if non-nil, zero value otherwise.

### GetLimitNumOk

`func (o *CreateVoucherGroupOpenApiVO) GetLimitNumOk() (*int32, bool)`

GetLimitNumOk returns a tuple with the LimitNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitNum

`func (o *CreateVoucherGroupOpenApiVO) SetLimitNum(v int32)`

SetLimitNum sets LimitNum field to given value.

### HasLimitNum

`func (o *CreateVoucherGroupOpenApiVO) HasLimitNum() bool`

HasLimitNum returns a boolean if a field has been set.

### GetLimitType

`func (o *CreateVoucherGroupOpenApiVO) GetLimitType() int32`

GetLimitType returns the LimitType field if non-nil, zero value otherwise.

### GetLimitTypeOk

`func (o *CreateVoucherGroupOpenApiVO) GetLimitTypeOk() (*int32, bool)`

GetLimitTypeOk returns a tuple with the LimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitType

`func (o *CreateVoucherGroupOpenApiVO) SetLimitType(v int32)`

SetLimitType sets LimitType field to given value.


### GetLogout

`func (o *CreateVoucherGroupOpenApiVO) GetLogout() bool`

GetLogout returns the Logout field if non-nil, zero value otherwise.

### GetLogoutOk

`func (o *CreateVoucherGroupOpenApiVO) GetLogoutOk() (*bool, bool)`

GetLogoutOk returns a tuple with the Logout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogout

`func (o *CreateVoucherGroupOpenApiVO) SetLogout(v bool)`

SetLogout sets Logout field to given value.

### HasLogout

`func (o *CreateVoucherGroupOpenApiVO) HasLogout() bool`

HasLogout returns a boolean if a field has been set.

### GetName

`func (o *CreateVoucherGroupOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVoucherGroupOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVoucherGroupOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetPortals

`func (o *CreateVoucherGroupOpenApiVO) GetPortals() []string`

GetPortals returns the Portals field if non-nil, zero value otherwise.

### GetPortalsOk

`func (o *CreateVoucherGroupOpenApiVO) GetPortalsOk() (*[]string, bool)`

GetPortalsOk returns a tuple with the Portals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortals

`func (o *CreateVoucherGroupOpenApiVO) SetPortals(v []string)`

SetPortals sets Portals field to given value.

### HasPortals

`func (o *CreateVoucherGroupOpenApiVO) HasPortals() bool`

HasPortals returns a boolean if a field has been set.

### GetPrintComments

`func (o *CreateVoucherGroupOpenApiVO) GetPrintComments() string`

GetPrintComments returns the PrintComments field if non-nil, zero value otherwise.

### GetPrintCommentsOk

`func (o *CreateVoucherGroupOpenApiVO) GetPrintCommentsOk() (*string, bool)`

GetPrintCommentsOk returns a tuple with the PrintComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintComments

`func (o *CreateVoucherGroupOpenApiVO) SetPrintComments(v string)`

SetPrintComments sets PrintComments field to given value.

### HasPrintComments

`func (o *CreateVoucherGroupOpenApiVO) HasPrintComments() bool`

HasPrintComments returns a boolean if a field has been set.

### GetRateLimit

`func (o *CreateVoucherGroupOpenApiVO) GetRateLimit() RateLimitOpenApiVO`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *CreateVoucherGroupOpenApiVO) GetRateLimitOk() (*RateLimitOpenApiVO, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *CreateVoucherGroupOpenApiVO) SetRateLimit(v RateLimitOpenApiVO)`

SetRateLimit sets RateLimit field to given value.


### GetSchedule

`func (o *CreateVoucherGroupOpenApiVO) GetSchedule() VoucherScheduleOpenApiVO`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CreateVoucherGroupOpenApiVO) GetScheduleOk() (*VoucherScheduleOpenApiVO, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CreateVoucherGroupOpenApiVO) SetSchedule(v VoucherScheduleOpenApiVO)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CreateVoucherGroupOpenApiVO) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetTimingType

`func (o *CreateVoucherGroupOpenApiVO) GetTimingType() int32`

GetTimingType returns the TimingType field if non-nil, zero value otherwise.

### GetTimingTypeOk

`func (o *CreateVoucherGroupOpenApiVO) GetTimingTypeOk() (*int32, bool)`

GetTimingTypeOk returns a tuple with the TimingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimingType

`func (o *CreateVoucherGroupOpenApiVO) SetTimingType(v int32)`

SetTimingType sets TimingType field to given value.


### GetTrafficLimit

`func (o *CreateVoucherGroupOpenApiVO) GetTrafficLimit() int64`

GetTrafficLimit returns the TrafficLimit field if non-nil, zero value otherwise.

### GetTrafficLimitOk

`func (o *CreateVoucherGroupOpenApiVO) GetTrafficLimitOk() (*int64, bool)`

GetTrafficLimitOk returns a tuple with the TrafficLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimit

`func (o *CreateVoucherGroupOpenApiVO) SetTrafficLimit(v int64)`

SetTrafficLimit sets TrafficLimit field to given value.

### HasTrafficLimit

`func (o *CreateVoucherGroupOpenApiVO) HasTrafficLimit() bool`

HasTrafficLimit returns a boolean if a field has been set.

### GetTrafficLimitEnable

`func (o *CreateVoucherGroupOpenApiVO) GetTrafficLimitEnable() bool`

GetTrafficLimitEnable returns the TrafficLimitEnable field if non-nil, zero value otherwise.

### GetTrafficLimitEnableOk

`func (o *CreateVoucherGroupOpenApiVO) GetTrafficLimitEnableOk() (*bool, bool)`

GetTrafficLimitEnableOk returns a tuple with the TrafficLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimitEnable

`func (o *CreateVoucherGroupOpenApiVO) SetTrafficLimitEnable(v bool)`

SetTrafficLimitEnable sets TrafficLimitEnable field to given value.


### GetTrafficLimitFrequency

`func (o *CreateVoucherGroupOpenApiVO) GetTrafficLimitFrequency() int32`

GetTrafficLimitFrequency returns the TrafficLimitFrequency field if non-nil, zero value otherwise.

### GetTrafficLimitFrequencyOk

`func (o *CreateVoucherGroupOpenApiVO) GetTrafficLimitFrequencyOk() (*int32, bool)`

GetTrafficLimitFrequencyOk returns a tuple with the TrafficLimitFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimitFrequency

`func (o *CreateVoucherGroupOpenApiVO) SetTrafficLimitFrequency(v int32)`

SetTrafficLimitFrequency sets TrafficLimitFrequency field to given value.

### HasTrafficLimitFrequency

`func (o *CreateVoucherGroupOpenApiVO) HasTrafficLimitFrequency() bool`

HasTrafficLimitFrequency returns a boolean if a field has been set.

### GetUnitPrice

`func (o *CreateVoucherGroupOpenApiVO) GetUnitPrice() int64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *CreateVoucherGroupOpenApiVO) GetUnitPriceOk() (*int64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *CreateVoucherGroupOpenApiVO) SetUnitPrice(v int64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *CreateVoucherGroupOpenApiVO) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetValidityType

`func (o *CreateVoucherGroupOpenApiVO) GetValidityType() int32`

GetValidityType returns the ValidityType field if non-nil, zero value otherwise.

### GetValidityTypeOk

`func (o *CreateVoucherGroupOpenApiVO) GetValidityTypeOk() (*int32, bool)`

GetValidityTypeOk returns a tuple with the ValidityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityType

`func (o *CreateVoucherGroupOpenApiVO) SetValidityType(v int32)`

SetValidityType sets ValidityType field to given value.

### HasValidityType

`func (o *CreateVoucherGroupOpenApiVO) HasValidityType() bool`

HasValidityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


