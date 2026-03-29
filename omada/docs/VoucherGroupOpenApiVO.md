# VoucherGroupOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyToAllPortals** | Pointer to **bool** | Is the voucher effective for all portals, including all newly created portals | [optional] 
**CreatedTime** | Pointer to **int64** | Create timestamp for the voucher group, unit: millisecond | [optional] 
**CreatorName** | Pointer to **string** | Role of the creator of the voucher group | [optional] 
**Currency** | Pointer to **string** | Currency Short Code of voucher. For the values of Currency Short Code, refer to section 5.4.2 of the Open API Access Guide. | [optional] 
**Description** | Pointer to **string** | Description of the voucher group | [optional] 
**Duration** | Pointer to **int64** | Duration of one use, unit: minute. It should be within the range of 1–14400000. | [optional] 
**DurationType** | Pointer to **int32** | The duration type of the voucher. It should be a value as follows: 0: Client duration, each client expires after the duration is used. 1: Voucher duration, after reaching the voucher duration, clients using the voucher will expire | [optional] 
**EffectiveTime** | Pointer to **int64** | The timestamp when the voucher takes effect, unit: millisecond | [optional] 
**ExpirationTime** | Pointer to **int64** | The timestamp of the expiration of the voucher, unit: millisecond | [optional] 
**ExpiredCount** | Pointer to **int32** | Expired voucher counts of the voucher group | [optional] 
**Id** | Pointer to **string** | Voucher group ID | [optional] 
**InUseCount** | Pointer to **int32** | In use voucher counts of the voucher group | [optional] 
**LimitNum** | Pointer to **int32** | The number of limitations. It should be within the range of 1–999. If Parameter [limitType] is 0 or 1, [limitNum] should not be null.When Parameter [limitType] is 0, [limitNum] represents the maximum number of times this voucher can be used.When Parameter [limitType] is 1, [limitNum] represents the maximum number of users this voucher can be used at the same time. | [optional] 
**LimitType** | Pointer to **int32** | The limitations of the voucher. It should be a value as follows: 0: Limited Usage Counts, 1: Limited Online Users, 2: Unlimited | [optional] 
**Logout** | Pointer to **bool** | Whether the voucher support portal logout functionality | [optional] 
**Name** | Pointer to **string** | Voucher group ID | [optional] 
**PortalNames** | Pointer to **[]string** | Bound portal name list | [optional] 
**PrintComments** | Pointer to **string** | Print comments of the voucher group | [optional] 
**RateLimit** | Pointer to [**RateLimitOpenApiVO**](RateLimitOpenApiVO.md) |  | [optional] 
**Schedule** | Pointer to [**VoucherScheduleOpenApiVO**](VoucherScheduleOpenApiVO.md) |  | [optional] 
**TimingType** | Pointer to **int32** | The timing type of the voucher. It should be a value as follows: 0: Timing by time, clients can use vouchers for specified time duration. 1: Timing by usage, clients can use vouchers for the duration of actual usage | [optional] 
**TotalAmount** | Pointer to **string** | Total voucher amount of the voucher group | [optional] 
**TotalCount** | Pointer to **int32** | Total voucher counts of the voucher group | [optional] 
**TrafficLimit** | Pointer to **int64** | Traffic limit in MB. It should be within the range of 1–10485760 | [optional] 
**TrafficLimitEnable** | Pointer to **bool** | Whether to enable traffic limit | [optional] 
**TrafficLimitFrequency** | Pointer to **int32** | Frequency of traffic limit should be a value as follows: 0: total; 1: daily; 2: weekly; 3: monthly. | [optional] 
**UnitPrice** | Pointer to **string** | Price of single voucher. It should be within the range of 1–999999999 | [optional] 
**UnusedCount** | Pointer to **int32** | Unused voucher counts of the voucher group | [optional] 
**UsedCount** | Pointer to **int32** | Used voucher counts of the voucher group | [optional] 
**ValidityType** | Pointer to **int32** | The validity type of the voucher. It should be a value as follows: 0: Voucher can be used at any time, parameter [effectiveTime], [expirationTime] and [schedule] should be null. 1: Voucher can be used between the effective time and expiration time, parameter [effectiveTime] and [expirationTime] should not be null, parameter [schedule] should be null. 2: Voucher can be used within a specified time period by schedule, parameter [effectiveTime] and [expirationTime] should be null, parameter [schedule] should not be null | [optional] 

## Methods

### NewVoucherGroupOpenApiVO

`func NewVoucherGroupOpenApiVO() *VoucherGroupOpenApiVO`

NewVoucherGroupOpenApiVO instantiates a new VoucherGroupOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoucherGroupOpenApiVOWithDefaults

`func NewVoucherGroupOpenApiVOWithDefaults() *VoucherGroupOpenApiVO`

NewVoucherGroupOpenApiVOWithDefaults instantiates a new VoucherGroupOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyToAllPortals

`func (o *VoucherGroupOpenApiVO) GetApplyToAllPortals() bool`

GetApplyToAllPortals returns the ApplyToAllPortals field if non-nil, zero value otherwise.

### GetApplyToAllPortalsOk

`func (o *VoucherGroupOpenApiVO) GetApplyToAllPortalsOk() (*bool, bool)`

GetApplyToAllPortalsOk returns a tuple with the ApplyToAllPortals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToAllPortals

`func (o *VoucherGroupOpenApiVO) SetApplyToAllPortals(v bool)`

SetApplyToAllPortals sets ApplyToAllPortals field to given value.

### HasApplyToAllPortals

`func (o *VoucherGroupOpenApiVO) HasApplyToAllPortals() bool`

HasApplyToAllPortals returns a boolean if a field has been set.

### GetCreatedTime

`func (o *VoucherGroupOpenApiVO) GetCreatedTime() int64`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *VoucherGroupOpenApiVO) GetCreatedTimeOk() (*int64, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *VoucherGroupOpenApiVO) SetCreatedTime(v int64)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *VoucherGroupOpenApiVO) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCreatorName

`func (o *VoucherGroupOpenApiVO) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *VoucherGroupOpenApiVO) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *VoucherGroupOpenApiVO) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *VoucherGroupOpenApiVO) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### GetCurrency

`func (o *VoucherGroupOpenApiVO) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *VoucherGroupOpenApiVO) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *VoucherGroupOpenApiVO) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *VoucherGroupOpenApiVO) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *VoucherGroupOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VoucherGroupOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VoucherGroupOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VoucherGroupOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuration

`func (o *VoucherGroupOpenApiVO) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VoucherGroupOpenApiVO) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VoucherGroupOpenApiVO) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VoucherGroupOpenApiVO) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDurationType

`func (o *VoucherGroupOpenApiVO) GetDurationType() int32`

GetDurationType returns the DurationType field if non-nil, zero value otherwise.

### GetDurationTypeOk

`func (o *VoucherGroupOpenApiVO) GetDurationTypeOk() (*int32, bool)`

GetDurationTypeOk returns a tuple with the DurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationType

`func (o *VoucherGroupOpenApiVO) SetDurationType(v int32)`

SetDurationType sets DurationType field to given value.

### HasDurationType

`func (o *VoucherGroupOpenApiVO) HasDurationType() bool`

HasDurationType returns a boolean if a field has been set.

### GetEffectiveTime

`func (o *VoucherGroupOpenApiVO) GetEffectiveTime() int64`

GetEffectiveTime returns the EffectiveTime field if non-nil, zero value otherwise.

### GetEffectiveTimeOk

`func (o *VoucherGroupOpenApiVO) GetEffectiveTimeOk() (*int64, bool)`

GetEffectiveTimeOk returns a tuple with the EffectiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTime

`func (o *VoucherGroupOpenApiVO) SetEffectiveTime(v int64)`

SetEffectiveTime sets EffectiveTime field to given value.

### HasEffectiveTime

`func (o *VoucherGroupOpenApiVO) HasEffectiveTime() bool`

HasEffectiveTime returns a boolean if a field has been set.

### GetExpirationTime

`func (o *VoucherGroupOpenApiVO) GetExpirationTime() int64`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *VoucherGroupOpenApiVO) GetExpirationTimeOk() (*int64, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *VoucherGroupOpenApiVO) SetExpirationTime(v int64)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *VoucherGroupOpenApiVO) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetExpiredCount

`func (o *VoucherGroupOpenApiVO) GetExpiredCount() int32`

GetExpiredCount returns the ExpiredCount field if non-nil, zero value otherwise.

### GetExpiredCountOk

`func (o *VoucherGroupOpenApiVO) GetExpiredCountOk() (*int32, bool)`

GetExpiredCountOk returns a tuple with the ExpiredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredCount

`func (o *VoucherGroupOpenApiVO) SetExpiredCount(v int32)`

SetExpiredCount sets ExpiredCount field to given value.

### HasExpiredCount

`func (o *VoucherGroupOpenApiVO) HasExpiredCount() bool`

HasExpiredCount returns a boolean if a field has been set.

### GetId

`func (o *VoucherGroupOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VoucherGroupOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VoucherGroupOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VoucherGroupOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInUseCount

`func (o *VoucherGroupOpenApiVO) GetInUseCount() int32`

GetInUseCount returns the InUseCount field if non-nil, zero value otherwise.

### GetInUseCountOk

`func (o *VoucherGroupOpenApiVO) GetInUseCountOk() (*int32, bool)`

GetInUseCountOk returns a tuple with the InUseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUseCount

`func (o *VoucherGroupOpenApiVO) SetInUseCount(v int32)`

SetInUseCount sets InUseCount field to given value.

### HasInUseCount

`func (o *VoucherGroupOpenApiVO) HasInUseCount() bool`

HasInUseCount returns a boolean if a field has been set.

### GetLimitNum

`func (o *VoucherGroupOpenApiVO) GetLimitNum() int32`

GetLimitNum returns the LimitNum field if non-nil, zero value otherwise.

### GetLimitNumOk

`func (o *VoucherGroupOpenApiVO) GetLimitNumOk() (*int32, bool)`

GetLimitNumOk returns a tuple with the LimitNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitNum

`func (o *VoucherGroupOpenApiVO) SetLimitNum(v int32)`

SetLimitNum sets LimitNum field to given value.

### HasLimitNum

`func (o *VoucherGroupOpenApiVO) HasLimitNum() bool`

HasLimitNum returns a boolean if a field has been set.

### GetLimitType

`func (o *VoucherGroupOpenApiVO) GetLimitType() int32`

GetLimitType returns the LimitType field if non-nil, zero value otherwise.

### GetLimitTypeOk

`func (o *VoucherGroupOpenApiVO) GetLimitTypeOk() (*int32, bool)`

GetLimitTypeOk returns a tuple with the LimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitType

`func (o *VoucherGroupOpenApiVO) SetLimitType(v int32)`

SetLimitType sets LimitType field to given value.

### HasLimitType

`func (o *VoucherGroupOpenApiVO) HasLimitType() bool`

HasLimitType returns a boolean if a field has been set.

### GetLogout

`func (o *VoucherGroupOpenApiVO) GetLogout() bool`

GetLogout returns the Logout field if non-nil, zero value otherwise.

### GetLogoutOk

`func (o *VoucherGroupOpenApiVO) GetLogoutOk() (*bool, bool)`

GetLogoutOk returns a tuple with the Logout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogout

`func (o *VoucherGroupOpenApiVO) SetLogout(v bool)`

SetLogout sets Logout field to given value.

### HasLogout

`func (o *VoucherGroupOpenApiVO) HasLogout() bool`

HasLogout returns a boolean if a field has been set.

### GetName

`func (o *VoucherGroupOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VoucherGroupOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VoucherGroupOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VoucherGroupOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortalNames

`func (o *VoucherGroupOpenApiVO) GetPortalNames() []string`

GetPortalNames returns the PortalNames field if non-nil, zero value otherwise.

### GetPortalNamesOk

`func (o *VoucherGroupOpenApiVO) GetPortalNamesOk() (*[]string, bool)`

GetPortalNamesOk returns a tuple with the PortalNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalNames

`func (o *VoucherGroupOpenApiVO) SetPortalNames(v []string)`

SetPortalNames sets PortalNames field to given value.

### HasPortalNames

`func (o *VoucherGroupOpenApiVO) HasPortalNames() bool`

HasPortalNames returns a boolean if a field has been set.

### GetPrintComments

`func (o *VoucherGroupOpenApiVO) GetPrintComments() string`

GetPrintComments returns the PrintComments field if non-nil, zero value otherwise.

### GetPrintCommentsOk

`func (o *VoucherGroupOpenApiVO) GetPrintCommentsOk() (*string, bool)`

GetPrintCommentsOk returns a tuple with the PrintComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintComments

`func (o *VoucherGroupOpenApiVO) SetPrintComments(v string)`

SetPrintComments sets PrintComments field to given value.

### HasPrintComments

`func (o *VoucherGroupOpenApiVO) HasPrintComments() bool`

HasPrintComments returns a boolean if a field has been set.

### GetRateLimit

`func (o *VoucherGroupOpenApiVO) GetRateLimit() RateLimitOpenApiVO`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *VoucherGroupOpenApiVO) GetRateLimitOk() (*RateLimitOpenApiVO, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *VoucherGroupOpenApiVO) SetRateLimit(v RateLimitOpenApiVO)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *VoucherGroupOpenApiVO) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetSchedule

`func (o *VoucherGroupOpenApiVO) GetSchedule() VoucherScheduleOpenApiVO`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *VoucherGroupOpenApiVO) GetScheduleOk() (*VoucherScheduleOpenApiVO, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *VoucherGroupOpenApiVO) SetSchedule(v VoucherScheduleOpenApiVO)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *VoucherGroupOpenApiVO) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetTimingType

`func (o *VoucherGroupOpenApiVO) GetTimingType() int32`

GetTimingType returns the TimingType field if non-nil, zero value otherwise.

### GetTimingTypeOk

`func (o *VoucherGroupOpenApiVO) GetTimingTypeOk() (*int32, bool)`

GetTimingTypeOk returns a tuple with the TimingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimingType

`func (o *VoucherGroupOpenApiVO) SetTimingType(v int32)`

SetTimingType sets TimingType field to given value.

### HasTimingType

`func (o *VoucherGroupOpenApiVO) HasTimingType() bool`

HasTimingType returns a boolean if a field has been set.

### GetTotalAmount

`func (o *VoucherGroupOpenApiVO) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *VoucherGroupOpenApiVO) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *VoucherGroupOpenApiVO) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *VoucherGroupOpenApiVO) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetTotalCount

`func (o *VoucherGroupOpenApiVO) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *VoucherGroupOpenApiVO) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *VoucherGroupOpenApiVO) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *VoucherGroupOpenApiVO) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTrafficLimit

`func (o *VoucherGroupOpenApiVO) GetTrafficLimit() int64`

GetTrafficLimit returns the TrafficLimit field if non-nil, zero value otherwise.

### GetTrafficLimitOk

`func (o *VoucherGroupOpenApiVO) GetTrafficLimitOk() (*int64, bool)`

GetTrafficLimitOk returns a tuple with the TrafficLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimit

`func (o *VoucherGroupOpenApiVO) SetTrafficLimit(v int64)`

SetTrafficLimit sets TrafficLimit field to given value.

### HasTrafficLimit

`func (o *VoucherGroupOpenApiVO) HasTrafficLimit() bool`

HasTrafficLimit returns a boolean if a field has been set.

### GetTrafficLimitEnable

`func (o *VoucherGroupOpenApiVO) GetTrafficLimitEnable() bool`

GetTrafficLimitEnable returns the TrafficLimitEnable field if non-nil, zero value otherwise.

### GetTrafficLimitEnableOk

`func (o *VoucherGroupOpenApiVO) GetTrafficLimitEnableOk() (*bool, bool)`

GetTrafficLimitEnableOk returns a tuple with the TrafficLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimitEnable

`func (o *VoucherGroupOpenApiVO) SetTrafficLimitEnable(v bool)`

SetTrafficLimitEnable sets TrafficLimitEnable field to given value.

### HasTrafficLimitEnable

`func (o *VoucherGroupOpenApiVO) HasTrafficLimitEnable() bool`

HasTrafficLimitEnable returns a boolean if a field has been set.

### GetTrafficLimitFrequency

`func (o *VoucherGroupOpenApiVO) GetTrafficLimitFrequency() int32`

GetTrafficLimitFrequency returns the TrafficLimitFrequency field if non-nil, zero value otherwise.

### GetTrafficLimitFrequencyOk

`func (o *VoucherGroupOpenApiVO) GetTrafficLimitFrequencyOk() (*int32, bool)`

GetTrafficLimitFrequencyOk returns a tuple with the TrafficLimitFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimitFrequency

`func (o *VoucherGroupOpenApiVO) SetTrafficLimitFrequency(v int32)`

SetTrafficLimitFrequency sets TrafficLimitFrequency field to given value.

### HasTrafficLimitFrequency

`func (o *VoucherGroupOpenApiVO) HasTrafficLimitFrequency() bool`

HasTrafficLimitFrequency returns a boolean if a field has been set.

### GetUnitPrice

`func (o *VoucherGroupOpenApiVO) GetUnitPrice() string`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *VoucherGroupOpenApiVO) GetUnitPriceOk() (*string, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *VoucherGroupOpenApiVO) SetUnitPrice(v string)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *VoucherGroupOpenApiVO) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetUnusedCount

`func (o *VoucherGroupOpenApiVO) GetUnusedCount() int32`

GetUnusedCount returns the UnusedCount field if non-nil, zero value otherwise.

### GetUnusedCountOk

`func (o *VoucherGroupOpenApiVO) GetUnusedCountOk() (*int32, bool)`

GetUnusedCountOk returns a tuple with the UnusedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnusedCount

`func (o *VoucherGroupOpenApiVO) SetUnusedCount(v int32)`

SetUnusedCount sets UnusedCount field to given value.

### HasUnusedCount

`func (o *VoucherGroupOpenApiVO) HasUnusedCount() bool`

HasUnusedCount returns a boolean if a field has been set.

### GetUsedCount

`func (o *VoucherGroupOpenApiVO) GetUsedCount() int32`

GetUsedCount returns the UsedCount field if non-nil, zero value otherwise.

### GetUsedCountOk

`func (o *VoucherGroupOpenApiVO) GetUsedCountOk() (*int32, bool)`

GetUsedCountOk returns a tuple with the UsedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCount

`func (o *VoucherGroupOpenApiVO) SetUsedCount(v int32)`

SetUsedCount sets UsedCount field to given value.

### HasUsedCount

`func (o *VoucherGroupOpenApiVO) HasUsedCount() bool`

HasUsedCount returns a boolean if a field has been set.

### GetValidityType

`func (o *VoucherGroupOpenApiVO) GetValidityType() int32`

GetValidityType returns the ValidityType field if non-nil, zero value otherwise.

### GetValidityTypeOk

`func (o *VoucherGroupOpenApiVO) GetValidityTypeOk() (*int32, bool)`

GetValidityTypeOk returns a tuple with the ValidityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityType

`func (o *VoucherGroupOpenApiVO) SetValidityType(v int32)`

SetValidityType sets ValidityType field to given value.

### HasValidityType

`func (o *VoucherGroupOpenApiVO) HasValidityType() bool`

HasValidityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


