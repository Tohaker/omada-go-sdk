# VoucherGroupGridOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyToAllPortals** | Pointer to **bool** | Is the voucher effective for all portals, including all newly created portals | [optional] 
**CreatedTime** | Pointer to **int64** | Create timestamp for the voucher group, unit: millisecond | [optional] 
**CreatorName** | Pointer to **string** | Role of the creator of the voucher group | [optional] 
**Currency** | Pointer to **string** | Currency Short Code of voucher. For the values of Currency Short Code, refer to section 5.4.2 of the Open API Access Guide. | [optional] 
**CurrentPage** | Pointer to **int32** | Current page number | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page | [optional] 
**Data** | Pointer to [**[]SimpleVoucherOpenApiVO**](SimpleVoucherOpenApiVO.md) | Voucher pagination data of the voucher group | [optional] 
**Description** | Pointer to **string** | Description of the voucher group | [optional] 
**Duration** | Pointer to **int64** | Duration of one use, unit: minute. It should be within the range of 1–14400000. | [optional] 
**DurationType** | Pointer to **int32** | The duration type of the voucher. It should be a value as follows: 0: Client duration, each client expires after the duration is used. 1: Voucher duration, after reaching the voucher duration, clients using the voucher will expire | [optional] 
**EffectiveTime** | Pointer to **int64** | The timestamp when the voucher takes effect, unit: millisecond | [optional] 
**ExpirationTime** | Pointer to **int64** | The timestamp of the expiration of the voucher, unit: millisecond | [optional] 
**ExpiredCount** | Pointer to **int32** | Expired voucher counts of the voucher group, affected by search | [optional] 
**Id** | Pointer to **string** | Voucher group ID | [optional] 
**InUseCount** | Pointer to **int32** | In use voucher counts of the voucher group, affected by search | [optional] 
**LimitNum** | Pointer to **int32** | The number of limitations. It should be within the range of 1–999. If Parameter [limitType] is 0 or 1, [limitNum] should not be null.When Parameter [limitType] is 0, [limitNum] represents the maximum number of times this voucher can be used.When Parameter [limitType] is 1, [limitNum] represents the maximum number of users this voucher can be used at the same time. | [optional] 
**LimitType** | Pointer to **int32** | The limitations of the voucher. It should be a value as follows: 0: Limited Usage Counts, 1: Limited Online Users, 2: Unlimited | [optional] 
**Logout** | Pointer to **bool** | Whether the voucher support portal logout functionality | [optional] 
**Name** | Pointer to **string** | Voucher group ID | [optional] 
**PortalNames** | Pointer to **[]string** | Bound portal name list | [optional] 
**PrintComments** | Pointer to **string** | Print comments of the voucher group | [optional] 
**RateLimit** | Pointer to [**RateLimitOpenApiVO**](RateLimitOpenApiVO.md) |  | [optional] 
**Schedule** | Pointer to [**VoucherScheduleOpenApiVO**](VoucherScheduleOpenApiVO.md) |  | [optional] 
**StatisticsCount** | Pointer to [**StatisticsCountOpenApiVO**](StatisticsCountOpenApiVO.md) |  | [optional] 
**TimingType** | Pointer to **int32** | The timing type of the voucher. It should be a value as follows: 0: Timing by time, clients can use vouchers for specified time duration. 1: Timing by usage, clients can use vouchers for the duration of actual usage | [optional] 
**TotalAmount** | Pointer to **string** | Total voucher amount of the voucher group, unaffected by search | [optional] 
**TotalCount** | Pointer to **int32** | Total voucher counts of the voucher group, affected by search | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of vouchers in the voucher group, affected by search | [optional] 
**TrafficLimit** | Pointer to **int64** | Traffic limit in MB. It should be within the range of 1–10485760 | [optional] 
**TrafficLimitEnable** | Pointer to **bool** | Whether to enable traffic limit | [optional] 
**TrafficLimitFrequency** | Pointer to **int32** | Frequency of traffic limit should be a value as follows: 0: total; 1: daily; 2: weekly; 3: monthly. | [optional] 
**UnitPrice** | Pointer to **string** | Price of single voucher. It should be within the range of 1–999999999 | [optional] 
**UnusedAmount** | Pointer to **string** | Unused voucher amount of the voucher group, unaffected by search | [optional] 
**UnusedCount** | Pointer to **int32** | Unused voucher counts of the voucher group, affected by search | [optional] 
**UsedAmount** | Pointer to **string** | Used voucher amount of the voucher group, unaffected by search | [optional] 
**UsedCount** | Pointer to **int32** | Used voucher counts of the voucher group, affected by search | [optional] 
**ValidityType** | Pointer to **int32** | The validity type of the voucher. It should be a value as follows: 0: Voucher can be used at any time, parameter [effectiveTime], [expirationTime] and [schedule] should be null. 1: Voucher can be used between the effective time and expiration time, parameter [effectiveTime] and [expirationTime] should not be null, parameter [schedule] should be null. 2: Voucher can be used within a specified time period by schedule, parameter [effectiveTime] and [expirationTime] should be null, parameter [schedule] should not be null | [optional] 
**VoucherPattern** | Pointer to [**VoucherPatternOpenApiVO**](VoucherPatternOpenApiVO.md) |  | [optional] 

## Methods

### NewVoucherGroupGridOpenApiVO

`func NewVoucherGroupGridOpenApiVO() *VoucherGroupGridOpenApiVO`

NewVoucherGroupGridOpenApiVO instantiates a new VoucherGroupGridOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoucherGroupGridOpenApiVOWithDefaults

`func NewVoucherGroupGridOpenApiVOWithDefaults() *VoucherGroupGridOpenApiVO`

NewVoucherGroupGridOpenApiVOWithDefaults instantiates a new VoucherGroupGridOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyToAllPortals

`func (o *VoucherGroupGridOpenApiVO) GetApplyToAllPortals() bool`

GetApplyToAllPortals returns the ApplyToAllPortals field if non-nil, zero value otherwise.

### GetApplyToAllPortalsOk

`func (o *VoucherGroupGridOpenApiVO) GetApplyToAllPortalsOk() (*bool, bool)`

GetApplyToAllPortalsOk returns a tuple with the ApplyToAllPortals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToAllPortals

`func (o *VoucherGroupGridOpenApiVO) SetApplyToAllPortals(v bool)`

SetApplyToAllPortals sets ApplyToAllPortals field to given value.

### HasApplyToAllPortals

`func (o *VoucherGroupGridOpenApiVO) HasApplyToAllPortals() bool`

HasApplyToAllPortals returns a boolean if a field has been set.

### GetCreatedTime

`func (o *VoucherGroupGridOpenApiVO) GetCreatedTime() int64`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *VoucherGroupGridOpenApiVO) GetCreatedTimeOk() (*int64, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *VoucherGroupGridOpenApiVO) SetCreatedTime(v int64)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *VoucherGroupGridOpenApiVO) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCreatorName

`func (o *VoucherGroupGridOpenApiVO) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *VoucherGroupGridOpenApiVO) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *VoucherGroupGridOpenApiVO) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *VoucherGroupGridOpenApiVO) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### GetCurrency

`func (o *VoucherGroupGridOpenApiVO) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *VoucherGroupGridOpenApiVO) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *VoucherGroupGridOpenApiVO) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *VoucherGroupGridOpenApiVO) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrentPage

`func (o *VoucherGroupGridOpenApiVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *VoucherGroupGridOpenApiVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *VoucherGroupGridOpenApiVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *VoucherGroupGridOpenApiVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *VoucherGroupGridOpenApiVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *VoucherGroupGridOpenApiVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *VoucherGroupGridOpenApiVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *VoucherGroupGridOpenApiVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *VoucherGroupGridOpenApiVO) GetData() []SimpleVoucherOpenApiVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VoucherGroupGridOpenApiVO) GetDataOk() (*[]SimpleVoucherOpenApiVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VoucherGroupGridOpenApiVO) SetData(v []SimpleVoucherOpenApiVO)`

SetData sets Data field to given value.

### HasData

`func (o *VoucherGroupGridOpenApiVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *VoucherGroupGridOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VoucherGroupGridOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VoucherGroupGridOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VoucherGroupGridOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuration

`func (o *VoucherGroupGridOpenApiVO) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VoucherGroupGridOpenApiVO) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VoucherGroupGridOpenApiVO) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VoucherGroupGridOpenApiVO) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDurationType

`func (o *VoucherGroupGridOpenApiVO) GetDurationType() int32`

GetDurationType returns the DurationType field if non-nil, zero value otherwise.

### GetDurationTypeOk

`func (o *VoucherGroupGridOpenApiVO) GetDurationTypeOk() (*int32, bool)`

GetDurationTypeOk returns a tuple with the DurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationType

`func (o *VoucherGroupGridOpenApiVO) SetDurationType(v int32)`

SetDurationType sets DurationType field to given value.

### HasDurationType

`func (o *VoucherGroupGridOpenApiVO) HasDurationType() bool`

HasDurationType returns a boolean if a field has been set.

### GetEffectiveTime

`func (o *VoucherGroupGridOpenApiVO) GetEffectiveTime() int64`

GetEffectiveTime returns the EffectiveTime field if non-nil, zero value otherwise.

### GetEffectiveTimeOk

`func (o *VoucherGroupGridOpenApiVO) GetEffectiveTimeOk() (*int64, bool)`

GetEffectiveTimeOk returns a tuple with the EffectiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTime

`func (o *VoucherGroupGridOpenApiVO) SetEffectiveTime(v int64)`

SetEffectiveTime sets EffectiveTime field to given value.

### HasEffectiveTime

`func (o *VoucherGroupGridOpenApiVO) HasEffectiveTime() bool`

HasEffectiveTime returns a boolean if a field has been set.

### GetExpirationTime

`func (o *VoucherGroupGridOpenApiVO) GetExpirationTime() int64`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *VoucherGroupGridOpenApiVO) GetExpirationTimeOk() (*int64, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *VoucherGroupGridOpenApiVO) SetExpirationTime(v int64)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *VoucherGroupGridOpenApiVO) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetExpiredCount

`func (o *VoucherGroupGridOpenApiVO) GetExpiredCount() int32`

GetExpiredCount returns the ExpiredCount field if non-nil, zero value otherwise.

### GetExpiredCountOk

`func (o *VoucherGroupGridOpenApiVO) GetExpiredCountOk() (*int32, bool)`

GetExpiredCountOk returns a tuple with the ExpiredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredCount

`func (o *VoucherGroupGridOpenApiVO) SetExpiredCount(v int32)`

SetExpiredCount sets ExpiredCount field to given value.

### HasExpiredCount

`func (o *VoucherGroupGridOpenApiVO) HasExpiredCount() bool`

HasExpiredCount returns a boolean if a field has been set.

### GetId

`func (o *VoucherGroupGridOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VoucherGroupGridOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VoucherGroupGridOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VoucherGroupGridOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInUseCount

`func (o *VoucherGroupGridOpenApiVO) GetInUseCount() int32`

GetInUseCount returns the InUseCount field if non-nil, zero value otherwise.

### GetInUseCountOk

`func (o *VoucherGroupGridOpenApiVO) GetInUseCountOk() (*int32, bool)`

GetInUseCountOk returns a tuple with the InUseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUseCount

`func (o *VoucherGroupGridOpenApiVO) SetInUseCount(v int32)`

SetInUseCount sets InUseCount field to given value.

### HasInUseCount

`func (o *VoucherGroupGridOpenApiVO) HasInUseCount() bool`

HasInUseCount returns a boolean if a field has been set.

### GetLimitNum

`func (o *VoucherGroupGridOpenApiVO) GetLimitNum() int32`

GetLimitNum returns the LimitNum field if non-nil, zero value otherwise.

### GetLimitNumOk

`func (o *VoucherGroupGridOpenApiVO) GetLimitNumOk() (*int32, bool)`

GetLimitNumOk returns a tuple with the LimitNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitNum

`func (o *VoucherGroupGridOpenApiVO) SetLimitNum(v int32)`

SetLimitNum sets LimitNum field to given value.

### HasLimitNum

`func (o *VoucherGroupGridOpenApiVO) HasLimitNum() bool`

HasLimitNum returns a boolean if a field has been set.

### GetLimitType

`func (o *VoucherGroupGridOpenApiVO) GetLimitType() int32`

GetLimitType returns the LimitType field if non-nil, zero value otherwise.

### GetLimitTypeOk

`func (o *VoucherGroupGridOpenApiVO) GetLimitTypeOk() (*int32, bool)`

GetLimitTypeOk returns a tuple with the LimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitType

`func (o *VoucherGroupGridOpenApiVO) SetLimitType(v int32)`

SetLimitType sets LimitType field to given value.

### HasLimitType

`func (o *VoucherGroupGridOpenApiVO) HasLimitType() bool`

HasLimitType returns a boolean if a field has been set.

### GetLogout

`func (o *VoucherGroupGridOpenApiVO) GetLogout() bool`

GetLogout returns the Logout field if non-nil, zero value otherwise.

### GetLogoutOk

`func (o *VoucherGroupGridOpenApiVO) GetLogoutOk() (*bool, bool)`

GetLogoutOk returns a tuple with the Logout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogout

`func (o *VoucherGroupGridOpenApiVO) SetLogout(v bool)`

SetLogout sets Logout field to given value.

### HasLogout

`func (o *VoucherGroupGridOpenApiVO) HasLogout() bool`

HasLogout returns a boolean if a field has been set.

### GetName

`func (o *VoucherGroupGridOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VoucherGroupGridOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VoucherGroupGridOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VoucherGroupGridOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortalNames

`func (o *VoucherGroupGridOpenApiVO) GetPortalNames() []string`

GetPortalNames returns the PortalNames field if non-nil, zero value otherwise.

### GetPortalNamesOk

`func (o *VoucherGroupGridOpenApiVO) GetPortalNamesOk() (*[]string, bool)`

GetPortalNamesOk returns a tuple with the PortalNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalNames

`func (o *VoucherGroupGridOpenApiVO) SetPortalNames(v []string)`

SetPortalNames sets PortalNames field to given value.

### HasPortalNames

`func (o *VoucherGroupGridOpenApiVO) HasPortalNames() bool`

HasPortalNames returns a boolean if a field has been set.

### GetPrintComments

`func (o *VoucherGroupGridOpenApiVO) GetPrintComments() string`

GetPrintComments returns the PrintComments field if non-nil, zero value otherwise.

### GetPrintCommentsOk

`func (o *VoucherGroupGridOpenApiVO) GetPrintCommentsOk() (*string, bool)`

GetPrintCommentsOk returns a tuple with the PrintComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintComments

`func (o *VoucherGroupGridOpenApiVO) SetPrintComments(v string)`

SetPrintComments sets PrintComments field to given value.

### HasPrintComments

`func (o *VoucherGroupGridOpenApiVO) HasPrintComments() bool`

HasPrintComments returns a boolean if a field has been set.

### GetRateLimit

`func (o *VoucherGroupGridOpenApiVO) GetRateLimit() RateLimitOpenApiVO`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *VoucherGroupGridOpenApiVO) GetRateLimitOk() (*RateLimitOpenApiVO, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *VoucherGroupGridOpenApiVO) SetRateLimit(v RateLimitOpenApiVO)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *VoucherGroupGridOpenApiVO) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetSchedule

`func (o *VoucherGroupGridOpenApiVO) GetSchedule() VoucherScheduleOpenApiVO`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *VoucherGroupGridOpenApiVO) GetScheduleOk() (*VoucherScheduleOpenApiVO, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *VoucherGroupGridOpenApiVO) SetSchedule(v VoucherScheduleOpenApiVO)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *VoucherGroupGridOpenApiVO) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetStatisticsCount

`func (o *VoucherGroupGridOpenApiVO) GetStatisticsCount() StatisticsCountOpenApiVO`

GetStatisticsCount returns the StatisticsCount field if non-nil, zero value otherwise.

### GetStatisticsCountOk

`func (o *VoucherGroupGridOpenApiVO) GetStatisticsCountOk() (*StatisticsCountOpenApiVO, bool)`

GetStatisticsCountOk returns a tuple with the StatisticsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticsCount

`func (o *VoucherGroupGridOpenApiVO) SetStatisticsCount(v StatisticsCountOpenApiVO)`

SetStatisticsCount sets StatisticsCount field to given value.

### HasStatisticsCount

`func (o *VoucherGroupGridOpenApiVO) HasStatisticsCount() bool`

HasStatisticsCount returns a boolean if a field has been set.

### GetTimingType

`func (o *VoucherGroupGridOpenApiVO) GetTimingType() int32`

GetTimingType returns the TimingType field if non-nil, zero value otherwise.

### GetTimingTypeOk

`func (o *VoucherGroupGridOpenApiVO) GetTimingTypeOk() (*int32, bool)`

GetTimingTypeOk returns a tuple with the TimingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimingType

`func (o *VoucherGroupGridOpenApiVO) SetTimingType(v int32)`

SetTimingType sets TimingType field to given value.

### HasTimingType

`func (o *VoucherGroupGridOpenApiVO) HasTimingType() bool`

HasTimingType returns a boolean if a field has been set.

### GetTotalAmount

`func (o *VoucherGroupGridOpenApiVO) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *VoucherGroupGridOpenApiVO) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *VoucherGroupGridOpenApiVO) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *VoucherGroupGridOpenApiVO) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetTotalCount

`func (o *VoucherGroupGridOpenApiVO) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *VoucherGroupGridOpenApiVO) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *VoucherGroupGridOpenApiVO) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *VoucherGroupGridOpenApiVO) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTotalRows

`func (o *VoucherGroupGridOpenApiVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *VoucherGroupGridOpenApiVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *VoucherGroupGridOpenApiVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *VoucherGroupGridOpenApiVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.

### GetTrafficLimit

`func (o *VoucherGroupGridOpenApiVO) GetTrafficLimit() int64`

GetTrafficLimit returns the TrafficLimit field if non-nil, zero value otherwise.

### GetTrafficLimitOk

`func (o *VoucherGroupGridOpenApiVO) GetTrafficLimitOk() (*int64, bool)`

GetTrafficLimitOk returns a tuple with the TrafficLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimit

`func (o *VoucherGroupGridOpenApiVO) SetTrafficLimit(v int64)`

SetTrafficLimit sets TrafficLimit field to given value.

### HasTrafficLimit

`func (o *VoucherGroupGridOpenApiVO) HasTrafficLimit() bool`

HasTrafficLimit returns a boolean if a field has been set.

### GetTrafficLimitEnable

`func (o *VoucherGroupGridOpenApiVO) GetTrafficLimitEnable() bool`

GetTrafficLimitEnable returns the TrafficLimitEnable field if non-nil, zero value otherwise.

### GetTrafficLimitEnableOk

`func (o *VoucherGroupGridOpenApiVO) GetTrafficLimitEnableOk() (*bool, bool)`

GetTrafficLimitEnableOk returns a tuple with the TrafficLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimitEnable

`func (o *VoucherGroupGridOpenApiVO) SetTrafficLimitEnable(v bool)`

SetTrafficLimitEnable sets TrafficLimitEnable field to given value.

### HasTrafficLimitEnable

`func (o *VoucherGroupGridOpenApiVO) HasTrafficLimitEnable() bool`

HasTrafficLimitEnable returns a boolean if a field has been set.

### GetTrafficLimitFrequency

`func (o *VoucherGroupGridOpenApiVO) GetTrafficLimitFrequency() int32`

GetTrafficLimitFrequency returns the TrafficLimitFrequency field if non-nil, zero value otherwise.

### GetTrafficLimitFrequencyOk

`func (o *VoucherGroupGridOpenApiVO) GetTrafficLimitFrequencyOk() (*int32, bool)`

GetTrafficLimitFrequencyOk returns a tuple with the TrafficLimitFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimitFrequency

`func (o *VoucherGroupGridOpenApiVO) SetTrafficLimitFrequency(v int32)`

SetTrafficLimitFrequency sets TrafficLimitFrequency field to given value.

### HasTrafficLimitFrequency

`func (o *VoucherGroupGridOpenApiVO) HasTrafficLimitFrequency() bool`

HasTrafficLimitFrequency returns a boolean if a field has been set.

### GetUnitPrice

`func (o *VoucherGroupGridOpenApiVO) GetUnitPrice() string`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *VoucherGroupGridOpenApiVO) GetUnitPriceOk() (*string, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *VoucherGroupGridOpenApiVO) SetUnitPrice(v string)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *VoucherGroupGridOpenApiVO) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetUnusedAmount

`func (o *VoucherGroupGridOpenApiVO) GetUnusedAmount() string`

GetUnusedAmount returns the UnusedAmount field if non-nil, zero value otherwise.

### GetUnusedAmountOk

`func (o *VoucherGroupGridOpenApiVO) GetUnusedAmountOk() (*string, bool)`

GetUnusedAmountOk returns a tuple with the UnusedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnusedAmount

`func (o *VoucherGroupGridOpenApiVO) SetUnusedAmount(v string)`

SetUnusedAmount sets UnusedAmount field to given value.

### HasUnusedAmount

`func (o *VoucherGroupGridOpenApiVO) HasUnusedAmount() bool`

HasUnusedAmount returns a boolean if a field has been set.

### GetUnusedCount

`func (o *VoucherGroupGridOpenApiVO) GetUnusedCount() int32`

GetUnusedCount returns the UnusedCount field if non-nil, zero value otherwise.

### GetUnusedCountOk

`func (o *VoucherGroupGridOpenApiVO) GetUnusedCountOk() (*int32, bool)`

GetUnusedCountOk returns a tuple with the UnusedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnusedCount

`func (o *VoucherGroupGridOpenApiVO) SetUnusedCount(v int32)`

SetUnusedCount sets UnusedCount field to given value.

### HasUnusedCount

`func (o *VoucherGroupGridOpenApiVO) HasUnusedCount() bool`

HasUnusedCount returns a boolean if a field has been set.

### GetUsedAmount

`func (o *VoucherGroupGridOpenApiVO) GetUsedAmount() string`

GetUsedAmount returns the UsedAmount field if non-nil, zero value otherwise.

### GetUsedAmountOk

`func (o *VoucherGroupGridOpenApiVO) GetUsedAmountOk() (*string, bool)`

GetUsedAmountOk returns a tuple with the UsedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedAmount

`func (o *VoucherGroupGridOpenApiVO) SetUsedAmount(v string)`

SetUsedAmount sets UsedAmount field to given value.

### HasUsedAmount

`func (o *VoucherGroupGridOpenApiVO) HasUsedAmount() bool`

HasUsedAmount returns a boolean if a field has been set.

### GetUsedCount

`func (o *VoucherGroupGridOpenApiVO) GetUsedCount() int32`

GetUsedCount returns the UsedCount field if non-nil, zero value otherwise.

### GetUsedCountOk

`func (o *VoucherGroupGridOpenApiVO) GetUsedCountOk() (*int32, bool)`

GetUsedCountOk returns a tuple with the UsedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCount

`func (o *VoucherGroupGridOpenApiVO) SetUsedCount(v int32)`

SetUsedCount sets UsedCount field to given value.

### HasUsedCount

`func (o *VoucherGroupGridOpenApiVO) HasUsedCount() bool`

HasUsedCount returns a boolean if a field has been set.

### GetValidityType

`func (o *VoucherGroupGridOpenApiVO) GetValidityType() int32`

GetValidityType returns the ValidityType field if non-nil, zero value otherwise.

### GetValidityTypeOk

`func (o *VoucherGroupGridOpenApiVO) GetValidityTypeOk() (*int32, bool)`

GetValidityTypeOk returns a tuple with the ValidityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityType

`func (o *VoucherGroupGridOpenApiVO) SetValidityType(v int32)`

SetValidityType sets ValidityType field to given value.

### HasValidityType

`func (o *VoucherGroupGridOpenApiVO) HasValidityType() bool`

HasValidityType returns a boolean if a field has been set.

### GetVoucherPattern

`func (o *VoucherGroupGridOpenApiVO) GetVoucherPattern() VoucherPatternOpenApiVO`

GetVoucherPattern returns the VoucherPattern field if non-nil, zero value otherwise.

### GetVoucherPatternOk

`func (o *VoucherGroupGridOpenApiVO) GetVoucherPatternOk() (*VoucherPatternOpenApiVO, bool)`

GetVoucherPatternOk returns a tuple with the VoucherPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherPattern

`func (o *VoucherGroupGridOpenApiVO) SetVoucherPattern(v VoucherPatternOpenApiVO)`

SetVoucherPattern sets VoucherPattern field to given value.

### HasVoucherPattern

`func (o *VoucherGroupGridOpenApiVO) HasVoucherPattern() bool`

HasVoucherPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


