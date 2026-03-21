# VoucherOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Voucher code | [optional] 
**CreatedTime** | Pointer to **int64** | Create timestamp for the voucher, unit: millisecond | [optional] 
**Currency** | Pointer to **string** | Currency Short Code of voucher. For the values of Currency Short Code, refer to section 5.4.2 of the Open API Access Guide. | [optional] 
**Description** | Pointer to **string** | Description of the voucher | [optional] 
**Duration** | Pointer to **int64** | Duration of one use, unit: minute. It should be within the range of 1–14400000. | [optional] 
**DurationType** | Pointer to **int32** | The duration type of the voucher. It should be a value as follows: 0: Client duration, each client expires after the duration is used. 1: Voucher duration, after reaching the voucher duration, clients using the voucher will expire | [optional] 
**EffectiveTime** | Pointer to **int64** | The timestamp when the voucher takes effect, unit: millisecond | [optional] 
**ExpirationTime** | Pointer to **int64** | The timestamp of the expiration of the voucher, unit: millisecond | [optional] 
**Id** | Pointer to **string** | Voucher ID | [optional] 
**LimitNum** | Pointer to **int32** | The number of limitations. It should be within the range of 1–999. If Parameter [limitType] is 0 or 1, [limitNum] should not be null.When Parameter [limitType] is 0, [limitNum] represents the maximum number of times this voucher can be used.When Parameter [limitType] is 1, [limitNum] represents the maximum number of users this voucher can be used at the same time. | [optional] 
**LimitType** | Pointer to **int32** | The limitations of the voucher. It should be a value as follows: 0: Limited Usage Counts, 1: Limited Online Users, 2: Unlimited | [optional] 
**LogoSize** | Pointer to **int32** | Size of logo on the pattern of the voucher. It should be within the range of 12-18. | [optional] 
**Logout** | Pointer to **bool** | Whether the voucher support portal logout functionality | [optional] 
**NetworkNameList** | Pointer to **[]string** | Networks for voucher | [optional] 
**PicId** | Pointer to **string** | Voucher logo picture ID | [optional] 
**PortalNames** | Pointer to **[]string** | Bound portal name list | [optional] 
**Position** | Pointer to **int32** | Position for logo or title | [optional] 
**PrintComments** | Pointer to **string** | Customized print information for voters | [optional] 
**RateLimit** | Pointer to [**RateLimitOpenApiVO**](RateLimitOpenApiVO.md) |  | [optional] 
**SsidNameList** | Pointer to **[]string** | SSIDs for voucher | [optional] 
**StartTime** | Pointer to **int64** | The expiration date of the voucher | [optional] 
**TimingType** | Pointer to **int32** | The timing type of the voucher. It should be a value as follows: 0: Timing by time, clients can use vouchers for specified time duration. 1: Timing by usage, clients can use vouchers for the duration of actual usage | [optional] 
**Title** | Pointer to **string** | Title for voucher | [optional] 
**TitleSize** | Pointer to **int32** | Size of title on the pattern of the voucher. It should be within the range of 50-175 | [optional] 
**TrafficLeft** | Pointer to **bool** | Is there remaining traffic of traffic for the voucher | [optional] 
**TrafficLimit** | Pointer to **int64** | Traffic limit in MB. It should be within the range of 1–10485760 | [optional] 
**TrafficLimitEnable** | Pointer to **bool** | Whether to enable traffic limit | [optional] 
**TrafficLimitFrequency** | Pointer to **int32** | Frequency of traffic limit should be a value as follows: 0: total; 1: daily; 2: weekly; 3: monthly. | [optional] 
**TrafficUsed** | Pointer to **int64** | Used traffic of the voucher, unit: Byte | [optional] 
**UnitPrice** | Pointer to **string** | Price of single voucher. It should be within the range of 1–999999999 | [optional] 
**Used** | Pointer to **int32** | The number of times the voucher is used | [optional] 
**Valid** | Pointer to **bool** | Can the voucher still be used | [optional] 
**Validity** | Pointer to **string** | Information on the validity period of the voucher | [optional] 

## Methods

### NewVoucherOpenApiVO

`func NewVoucherOpenApiVO() *VoucherOpenApiVO`

NewVoucherOpenApiVO instantiates a new VoucherOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoucherOpenApiVOWithDefaults

`func NewVoucherOpenApiVOWithDefaults() *VoucherOpenApiVO`

NewVoucherOpenApiVOWithDefaults instantiates a new VoucherOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *VoucherOpenApiVO) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VoucherOpenApiVO) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VoucherOpenApiVO) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VoucherOpenApiVO) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreatedTime

`func (o *VoucherOpenApiVO) GetCreatedTime() int64`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *VoucherOpenApiVO) GetCreatedTimeOk() (*int64, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *VoucherOpenApiVO) SetCreatedTime(v int64)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *VoucherOpenApiVO) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrency

`func (o *VoucherOpenApiVO) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *VoucherOpenApiVO) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *VoucherOpenApiVO) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *VoucherOpenApiVO) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *VoucherOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VoucherOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VoucherOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VoucherOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuration

`func (o *VoucherOpenApiVO) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VoucherOpenApiVO) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VoucherOpenApiVO) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VoucherOpenApiVO) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDurationType

`func (o *VoucherOpenApiVO) GetDurationType() int32`

GetDurationType returns the DurationType field if non-nil, zero value otherwise.

### GetDurationTypeOk

`func (o *VoucherOpenApiVO) GetDurationTypeOk() (*int32, bool)`

GetDurationTypeOk returns a tuple with the DurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationType

`func (o *VoucherOpenApiVO) SetDurationType(v int32)`

SetDurationType sets DurationType field to given value.

### HasDurationType

`func (o *VoucherOpenApiVO) HasDurationType() bool`

HasDurationType returns a boolean if a field has been set.

### GetEffectiveTime

`func (o *VoucherOpenApiVO) GetEffectiveTime() int64`

GetEffectiveTime returns the EffectiveTime field if non-nil, zero value otherwise.

### GetEffectiveTimeOk

`func (o *VoucherOpenApiVO) GetEffectiveTimeOk() (*int64, bool)`

GetEffectiveTimeOk returns a tuple with the EffectiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTime

`func (o *VoucherOpenApiVO) SetEffectiveTime(v int64)`

SetEffectiveTime sets EffectiveTime field to given value.

### HasEffectiveTime

`func (o *VoucherOpenApiVO) HasEffectiveTime() bool`

HasEffectiveTime returns a boolean if a field has been set.

### GetExpirationTime

`func (o *VoucherOpenApiVO) GetExpirationTime() int64`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *VoucherOpenApiVO) GetExpirationTimeOk() (*int64, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *VoucherOpenApiVO) SetExpirationTime(v int64)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *VoucherOpenApiVO) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetId

`func (o *VoucherOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VoucherOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VoucherOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VoucherOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLimitNum

`func (o *VoucherOpenApiVO) GetLimitNum() int32`

GetLimitNum returns the LimitNum field if non-nil, zero value otherwise.

### GetLimitNumOk

`func (o *VoucherOpenApiVO) GetLimitNumOk() (*int32, bool)`

GetLimitNumOk returns a tuple with the LimitNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitNum

`func (o *VoucherOpenApiVO) SetLimitNum(v int32)`

SetLimitNum sets LimitNum field to given value.

### HasLimitNum

`func (o *VoucherOpenApiVO) HasLimitNum() bool`

HasLimitNum returns a boolean if a field has been set.

### GetLimitType

`func (o *VoucherOpenApiVO) GetLimitType() int32`

GetLimitType returns the LimitType field if non-nil, zero value otherwise.

### GetLimitTypeOk

`func (o *VoucherOpenApiVO) GetLimitTypeOk() (*int32, bool)`

GetLimitTypeOk returns a tuple with the LimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitType

`func (o *VoucherOpenApiVO) SetLimitType(v int32)`

SetLimitType sets LimitType field to given value.

### HasLimitType

`func (o *VoucherOpenApiVO) HasLimitType() bool`

HasLimitType returns a boolean if a field has been set.

### GetLogoSize

`func (o *VoucherOpenApiVO) GetLogoSize() int32`

GetLogoSize returns the LogoSize field if non-nil, zero value otherwise.

### GetLogoSizeOk

`func (o *VoucherOpenApiVO) GetLogoSizeOk() (*int32, bool)`

GetLogoSizeOk returns a tuple with the LogoSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoSize

`func (o *VoucherOpenApiVO) SetLogoSize(v int32)`

SetLogoSize sets LogoSize field to given value.

### HasLogoSize

`func (o *VoucherOpenApiVO) HasLogoSize() bool`

HasLogoSize returns a boolean if a field has been set.

### GetLogout

`func (o *VoucherOpenApiVO) GetLogout() bool`

GetLogout returns the Logout field if non-nil, zero value otherwise.

### GetLogoutOk

`func (o *VoucherOpenApiVO) GetLogoutOk() (*bool, bool)`

GetLogoutOk returns a tuple with the Logout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogout

`func (o *VoucherOpenApiVO) SetLogout(v bool)`

SetLogout sets Logout field to given value.

### HasLogout

`func (o *VoucherOpenApiVO) HasLogout() bool`

HasLogout returns a boolean if a field has been set.

### GetNetworkNameList

`func (o *VoucherOpenApiVO) GetNetworkNameList() []string`

GetNetworkNameList returns the NetworkNameList field if non-nil, zero value otherwise.

### GetNetworkNameListOk

`func (o *VoucherOpenApiVO) GetNetworkNameListOk() (*[]string, bool)`

GetNetworkNameListOk returns a tuple with the NetworkNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkNameList

`func (o *VoucherOpenApiVO) SetNetworkNameList(v []string)`

SetNetworkNameList sets NetworkNameList field to given value.

### HasNetworkNameList

`func (o *VoucherOpenApiVO) HasNetworkNameList() bool`

HasNetworkNameList returns a boolean if a field has been set.

### GetPicId

`func (o *VoucherOpenApiVO) GetPicId() string`

GetPicId returns the PicId field if non-nil, zero value otherwise.

### GetPicIdOk

`func (o *VoucherOpenApiVO) GetPicIdOk() (*string, bool)`

GetPicIdOk returns a tuple with the PicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicId

`func (o *VoucherOpenApiVO) SetPicId(v string)`

SetPicId sets PicId field to given value.

### HasPicId

`func (o *VoucherOpenApiVO) HasPicId() bool`

HasPicId returns a boolean if a field has been set.

### GetPortalNames

`func (o *VoucherOpenApiVO) GetPortalNames() []string`

GetPortalNames returns the PortalNames field if non-nil, zero value otherwise.

### GetPortalNamesOk

`func (o *VoucherOpenApiVO) GetPortalNamesOk() (*[]string, bool)`

GetPortalNamesOk returns a tuple with the PortalNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalNames

`func (o *VoucherOpenApiVO) SetPortalNames(v []string)`

SetPortalNames sets PortalNames field to given value.

### HasPortalNames

`func (o *VoucherOpenApiVO) HasPortalNames() bool`

HasPortalNames returns a boolean if a field has been set.

### GetPosition

`func (o *VoucherOpenApiVO) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *VoucherOpenApiVO) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *VoucherOpenApiVO) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *VoucherOpenApiVO) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPrintComments

`func (o *VoucherOpenApiVO) GetPrintComments() string`

GetPrintComments returns the PrintComments field if non-nil, zero value otherwise.

### GetPrintCommentsOk

`func (o *VoucherOpenApiVO) GetPrintCommentsOk() (*string, bool)`

GetPrintCommentsOk returns a tuple with the PrintComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintComments

`func (o *VoucherOpenApiVO) SetPrintComments(v string)`

SetPrintComments sets PrintComments field to given value.

### HasPrintComments

`func (o *VoucherOpenApiVO) HasPrintComments() bool`

HasPrintComments returns a boolean if a field has been set.

### GetRateLimit

`func (o *VoucherOpenApiVO) GetRateLimit() RateLimitOpenApiVO`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *VoucherOpenApiVO) GetRateLimitOk() (*RateLimitOpenApiVO, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *VoucherOpenApiVO) SetRateLimit(v RateLimitOpenApiVO)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *VoucherOpenApiVO) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetSsidNameList

`func (o *VoucherOpenApiVO) GetSsidNameList() []string`

GetSsidNameList returns the SsidNameList field if non-nil, zero value otherwise.

### GetSsidNameListOk

`func (o *VoucherOpenApiVO) GetSsidNameListOk() (*[]string, bool)`

GetSsidNameListOk returns a tuple with the SsidNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNameList

`func (o *VoucherOpenApiVO) SetSsidNameList(v []string)`

SetSsidNameList sets SsidNameList field to given value.

### HasSsidNameList

`func (o *VoucherOpenApiVO) HasSsidNameList() bool`

HasSsidNameList returns a boolean if a field has been set.

### GetStartTime

`func (o *VoucherOpenApiVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *VoucherOpenApiVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *VoucherOpenApiVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *VoucherOpenApiVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTimingType

`func (o *VoucherOpenApiVO) GetTimingType() int32`

GetTimingType returns the TimingType field if non-nil, zero value otherwise.

### GetTimingTypeOk

`func (o *VoucherOpenApiVO) GetTimingTypeOk() (*int32, bool)`

GetTimingTypeOk returns a tuple with the TimingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimingType

`func (o *VoucherOpenApiVO) SetTimingType(v int32)`

SetTimingType sets TimingType field to given value.

### HasTimingType

`func (o *VoucherOpenApiVO) HasTimingType() bool`

HasTimingType returns a boolean if a field has been set.

### GetTitle

`func (o *VoucherOpenApiVO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VoucherOpenApiVO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VoucherOpenApiVO) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *VoucherOpenApiVO) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleSize

`func (o *VoucherOpenApiVO) GetTitleSize() int32`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *VoucherOpenApiVO) GetTitleSizeOk() (*int32, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *VoucherOpenApiVO) SetTitleSize(v int32)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *VoucherOpenApiVO) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetTrafficLeft

`func (o *VoucherOpenApiVO) GetTrafficLeft() bool`

GetTrafficLeft returns the TrafficLeft field if non-nil, zero value otherwise.

### GetTrafficLeftOk

`func (o *VoucherOpenApiVO) GetTrafficLeftOk() (*bool, bool)`

GetTrafficLeftOk returns a tuple with the TrafficLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLeft

`func (o *VoucherOpenApiVO) SetTrafficLeft(v bool)`

SetTrafficLeft sets TrafficLeft field to given value.

### HasTrafficLeft

`func (o *VoucherOpenApiVO) HasTrafficLeft() bool`

HasTrafficLeft returns a boolean if a field has been set.

### GetTrafficLimit

`func (o *VoucherOpenApiVO) GetTrafficLimit() int64`

GetTrafficLimit returns the TrafficLimit field if non-nil, zero value otherwise.

### GetTrafficLimitOk

`func (o *VoucherOpenApiVO) GetTrafficLimitOk() (*int64, bool)`

GetTrafficLimitOk returns a tuple with the TrafficLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimit

`func (o *VoucherOpenApiVO) SetTrafficLimit(v int64)`

SetTrafficLimit sets TrafficLimit field to given value.

### HasTrafficLimit

`func (o *VoucherOpenApiVO) HasTrafficLimit() bool`

HasTrafficLimit returns a boolean if a field has been set.

### GetTrafficLimitEnable

`func (o *VoucherOpenApiVO) GetTrafficLimitEnable() bool`

GetTrafficLimitEnable returns the TrafficLimitEnable field if non-nil, zero value otherwise.

### GetTrafficLimitEnableOk

`func (o *VoucherOpenApiVO) GetTrafficLimitEnableOk() (*bool, bool)`

GetTrafficLimitEnableOk returns a tuple with the TrafficLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimitEnable

`func (o *VoucherOpenApiVO) SetTrafficLimitEnable(v bool)`

SetTrafficLimitEnable sets TrafficLimitEnable field to given value.

### HasTrafficLimitEnable

`func (o *VoucherOpenApiVO) HasTrafficLimitEnable() bool`

HasTrafficLimitEnable returns a boolean if a field has been set.

### GetTrafficLimitFrequency

`func (o *VoucherOpenApiVO) GetTrafficLimitFrequency() int32`

GetTrafficLimitFrequency returns the TrafficLimitFrequency field if non-nil, zero value otherwise.

### GetTrafficLimitFrequencyOk

`func (o *VoucherOpenApiVO) GetTrafficLimitFrequencyOk() (*int32, bool)`

GetTrafficLimitFrequencyOk returns a tuple with the TrafficLimitFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimitFrequency

`func (o *VoucherOpenApiVO) SetTrafficLimitFrequency(v int32)`

SetTrafficLimitFrequency sets TrafficLimitFrequency field to given value.

### HasTrafficLimitFrequency

`func (o *VoucherOpenApiVO) HasTrafficLimitFrequency() bool`

HasTrafficLimitFrequency returns a boolean if a field has been set.

### GetTrafficUsed

`func (o *VoucherOpenApiVO) GetTrafficUsed() int64`

GetTrafficUsed returns the TrafficUsed field if non-nil, zero value otherwise.

### GetTrafficUsedOk

`func (o *VoucherOpenApiVO) GetTrafficUsedOk() (*int64, bool)`

GetTrafficUsedOk returns a tuple with the TrafficUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficUsed

`func (o *VoucherOpenApiVO) SetTrafficUsed(v int64)`

SetTrafficUsed sets TrafficUsed field to given value.

### HasTrafficUsed

`func (o *VoucherOpenApiVO) HasTrafficUsed() bool`

HasTrafficUsed returns a boolean if a field has been set.

### GetUnitPrice

`func (o *VoucherOpenApiVO) GetUnitPrice() string`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *VoucherOpenApiVO) GetUnitPriceOk() (*string, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *VoucherOpenApiVO) SetUnitPrice(v string)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *VoucherOpenApiVO) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetUsed

`func (o *VoucherOpenApiVO) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *VoucherOpenApiVO) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *VoucherOpenApiVO) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *VoucherOpenApiVO) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetValid

`func (o *VoucherOpenApiVO) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *VoucherOpenApiVO) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *VoucherOpenApiVO) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *VoucherOpenApiVO) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetValidity

`func (o *VoucherOpenApiVO) GetValidity() string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *VoucherOpenApiVO) GetValidityOk() (*string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *VoucherOpenApiVO) SetValidity(v string)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *VoucherOpenApiVO) HasValidity() bool`

HasValidity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


