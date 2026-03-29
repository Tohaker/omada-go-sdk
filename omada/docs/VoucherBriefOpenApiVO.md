# VoucherBriefOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Voucher code | [optional] 
**Currency** | Pointer to **string** | Currency Short Code of voucher. For the values of Currency Short Code, refer to section 5.4.2 of the Open API Access Guide. | [optional] 
**Duration** | Pointer to **int64** | Duration of one use, unit: minute. It should be within the range of 1–14400000. | [optional] 
**DurationType** | Pointer to **int32** | The duration type of the voucher. It should be a value as follows: 0: Client duration, each client expires after the duration is used. 1: Voucher duration, after reaching the voucher duration, clients using the voucher will expire | [optional] 
**LimitNum** | Pointer to **int32** | The number of limitations. It should be within the range of 1–999. If Parameter [limitType] is 0 or 1, [limitNum] should not be null.When Parameter [limitType] is 0, [limitNum] represents the maximum number of times this voucher can be used.When Parameter [limitType] is 1, [limitNum] represents the maximum number of users this voucher can be used at the same time. | [optional] 
**LimitType** | Pointer to **int32** | The limitations of the voucher. It should be a value as follows: 0: Limited Usage Counts, 1: Limited Online Users, 2: Unlimited | [optional] 
**LogoSize** | Pointer to **int32** | Size of logo on the pattern of the voucher. It should be within the range of 12-18. | [optional] 
**NetworkNameList** | Pointer to **[]string** | Network name for voucher | [optional] 
**PatternType** | Pointer to **int32** | Voucher pattern type. 0: Logo, 1: Title, 3: Disabled | [optional] 
**PicId** | Pointer to **string** | Voucher logo picture ID | [optional] 
**Position** | Pointer to **int32** | Position for logo or title | [optional] 
**PrintComments** | Pointer to **string** | Print comments of the voucher | [optional] 
**SsidNameList** | Pointer to **[]string** | Ssid name for voucher | [optional] 
**Title** | Pointer to **string** | Title for voucher | [optional] 
**TitleSize** | Pointer to **int32** | Size of title on the pattern of the voucher. It should be within the range of 50-175 | [optional] 
**UnitPrice** | Pointer to **string** | Price of single voucher. It should be within the range of 1–999999999 | [optional] 
**Validity** | Pointer to **string** | The validity period information of the voucher | [optional] 

## Methods

### NewVoucherBriefOpenApiVO

`func NewVoucherBriefOpenApiVO() *VoucherBriefOpenApiVO`

NewVoucherBriefOpenApiVO instantiates a new VoucherBriefOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoucherBriefOpenApiVOWithDefaults

`func NewVoucherBriefOpenApiVOWithDefaults() *VoucherBriefOpenApiVO`

NewVoucherBriefOpenApiVOWithDefaults instantiates a new VoucherBriefOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *VoucherBriefOpenApiVO) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VoucherBriefOpenApiVO) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VoucherBriefOpenApiVO) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VoucherBriefOpenApiVO) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCurrency

`func (o *VoucherBriefOpenApiVO) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *VoucherBriefOpenApiVO) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *VoucherBriefOpenApiVO) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *VoucherBriefOpenApiVO) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDuration

`func (o *VoucherBriefOpenApiVO) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VoucherBriefOpenApiVO) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VoucherBriefOpenApiVO) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VoucherBriefOpenApiVO) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDurationType

`func (o *VoucherBriefOpenApiVO) GetDurationType() int32`

GetDurationType returns the DurationType field if non-nil, zero value otherwise.

### GetDurationTypeOk

`func (o *VoucherBriefOpenApiVO) GetDurationTypeOk() (*int32, bool)`

GetDurationTypeOk returns a tuple with the DurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationType

`func (o *VoucherBriefOpenApiVO) SetDurationType(v int32)`

SetDurationType sets DurationType field to given value.

### HasDurationType

`func (o *VoucherBriefOpenApiVO) HasDurationType() bool`

HasDurationType returns a boolean if a field has been set.

### GetLimitNum

`func (o *VoucherBriefOpenApiVO) GetLimitNum() int32`

GetLimitNum returns the LimitNum field if non-nil, zero value otherwise.

### GetLimitNumOk

`func (o *VoucherBriefOpenApiVO) GetLimitNumOk() (*int32, bool)`

GetLimitNumOk returns a tuple with the LimitNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitNum

`func (o *VoucherBriefOpenApiVO) SetLimitNum(v int32)`

SetLimitNum sets LimitNum field to given value.

### HasLimitNum

`func (o *VoucherBriefOpenApiVO) HasLimitNum() bool`

HasLimitNum returns a boolean if a field has been set.

### GetLimitType

`func (o *VoucherBriefOpenApiVO) GetLimitType() int32`

GetLimitType returns the LimitType field if non-nil, zero value otherwise.

### GetLimitTypeOk

`func (o *VoucherBriefOpenApiVO) GetLimitTypeOk() (*int32, bool)`

GetLimitTypeOk returns a tuple with the LimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitType

`func (o *VoucherBriefOpenApiVO) SetLimitType(v int32)`

SetLimitType sets LimitType field to given value.

### HasLimitType

`func (o *VoucherBriefOpenApiVO) HasLimitType() bool`

HasLimitType returns a boolean if a field has been set.

### GetLogoSize

`func (o *VoucherBriefOpenApiVO) GetLogoSize() int32`

GetLogoSize returns the LogoSize field if non-nil, zero value otherwise.

### GetLogoSizeOk

`func (o *VoucherBriefOpenApiVO) GetLogoSizeOk() (*int32, bool)`

GetLogoSizeOk returns a tuple with the LogoSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoSize

`func (o *VoucherBriefOpenApiVO) SetLogoSize(v int32)`

SetLogoSize sets LogoSize field to given value.

### HasLogoSize

`func (o *VoucherBriefOpenApiVO) HasLogoSize() bool`

HasLogoSize returns a boolean if a field has been set.

### GetNetworkNameList

`func (o *VoucherBriefOpenApiVO) GetNetworkNameList() []string`

GetNetworkNameList returns the NetworkNameList field if non-nil, zero value otherwise.

### GetNetworkNameListOk

`func (o *VoucherBriefOpenApiVO) GetNetworkNameListOk() (*[]string, bool)`

GetNetworkNameListOk returns a tuple with the NetworkNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkNameList

`func (o *VoucherBriefOpenApiVO) SetNetworkNameList(v []string)`

SetNetworkNameList sets NetworkNameList field to given value.

### HasNetworkNameList

`func (o *VoucherBriefOpenApiVO) HasNetworkNameList() bool`

HasNetworkNameList returns a boolean if a field has been set.

### GetPatternType

`func (o *VoucherBriefOpenApiVO) GetPatternType() int32`

GetPatternType returns the PatternType field if non-nil, zero value otherwise.

### GetPatternTypeOk

`func (o *VoucherBriefOpenApiVO) GetPatternTypeOk() (*int32, bool)`

GetPatternTypeOk returns a tuple with the PatternType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternType

`func (o *VoucherBriefOpenApiVO) SetPatternType(v int32)`

SetPatternType sets PatternType field to given value.

### HasPatternType

`func (o *VoucherBriefOpenApiVO) HasPatternType() bool`

HasPatternType returns a boolean if a field has been set.

### GetPicId

`func (o *VoucherBriefOpenApiVO) GetPicId() string`

GetPicId returns the PicId field if non-nil, zero value otherwise.

### GetPicIdOk

`func (o *VoucherBriefOpenApiVO) GetPicIdOk() (*string, bool)`

GetPicIdOk returns a tuple with the PicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicId

`func (o *VoucherBriefOpenApiVO) SetPicId(v string)`

SetPicId sets PicId field to given value.

### HasPicId

`func (o *VoucherBriefOpenApiVO) HasPicId() bool`

HasPicId returns a boolean if a field has been set.

### GetPosition

`func (o *VoucherBriefOpenApiVO) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *VoucherBriefOpenApiVO) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *VoucherBriefOpenApiVO) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *VoucherBriefOpenApiVO) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPrintComments

`func (o *VoucherBriefOpenApiVO) GetPrintComments() string`

GetPrintComments returns the PrintComments field if non-nil, zero value otherwise.

### GetPrintCommentsOk

`func (o *VoucherBriefOpenApiVO) GetPrintCommentsOk() (*string, bool)`

GetPrintCommentsOk returns a tuple with the PrintComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintComments

`func (o *VoucherBriefOpenApiVO) SetPrintComments(v string)`

SetPrintComments sets PrintComments field to given value.

### HasPrintComments

`func (o *VoucherBriefOpenApiVO) HasPrintComments() bool`

HasPrintComments returns a boolean if a field has been set.

### GetSsidNameList

`func (o *VoucherBriefOpenApiVO) GetSsidNameList() []string`

GetSsidNameList returns the SsidNameList field if non-nil, zero value otherwise.

### GetSsidNameListOk

`func (o *VoucherBriefOpenApiVO) GetSsidNameListOk() (*[]string, bool)`

GetSsidNameListOk returns a tuple with the SsidNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNameList

`func (o *VoucherBriefOpenApiVO) SetSsidNameList(v []string)`

SetSsidNameList sets SsidNameList field to given value.

### HasSsidNameList

`func (o *VoucherBriefOpenApiVO) HasSsidNameList() bool`

HasSsidNameList returns a boolean if a field has been set.

### GetTitle

`func (o *VoucherBriefOpenApiVO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VoucherBriefOpenApiVO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VoucherBriefOpenApiVO) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *VoucherBriefOpenApiVO) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleSize

`func (o *VoucherBriefOpenApiVO) GetTitleSize() int32`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *VoucherBriefOpenApiVO) GetTitleSizeOk() (*int32, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *VoucherBriefOpenApiVO) SetTitleSize(v int32)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *VoucherBriefOpenApiVO) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetUnitPrice

`func (o *VoucherBriefOpenApiVO) GetUnitPrice() string`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *VoucherBriefOpenApiVO) GetUnitPriceOk() (*string, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *VoucherBriefOpenApiVO) SetUnitPrice(v string)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *VoucherBriefOpenApiVO) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetValidity

`func (o *VoucherBriefOpenApiVO) GetValidity() string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *VoucherBriefOpenApiVO) GetValidityOk() (*string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *VoucherBriefOpenApiVO) SetValidity(v string)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *VoucherBriefOpenApiVO) HasValidity() bool`

HasValidity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


