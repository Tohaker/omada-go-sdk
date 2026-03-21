# VoucherSummaryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | Amount of single voucher | [optional] 
**Count** | Pointer to **int64** | Count of vouchers | [optional] 
**Currency** | Pointer to **string** | Currency Short Code of voucher. For the values of Currency Short Code, refer to section 5.4.2 of the Open API Access Guide. | [optional] 
**Duration** | Pointer to **int64** | Duration of vouchers, unit: minutes | [optional] 

## Methods

### NewVoucherSummaryOpenApiVO

`func NewVoucherSummaryOpenApiVO() *VoucherSummaryOpenApiVO`

NewVoucherSummaryOpenApiVO instantiates a new VoucherSummaryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoucherSummaryOpenApiVOWithDefaults

`func NewVoucherSummaryOpenApiVOWithDefaults() *VoucherSummaryOpenApiVO`

NewVoucherSummaryOpenApiVOWithDefaults instantiates a new VoucherSummaryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *VoucherSummaryOpenApiVO) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *VoucherSummaryOpenApiVO) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *VoucherSummaryOpenApiVO) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *VoucherSummaryOpenApiVO) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCount

`func (o *VoucherSummaryOpenApiVO) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VoucherSummaryOpenApiVO) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VoucherSummaryOpenApiVO) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *VoucherSummaryOpenApiVO) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCurrency

`func (o *VoucherSummaryOpenApiVO) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *VoucherSummaryOpenApiVO) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *VoucherSummaryOpenApiVO) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *VoucherSummaryOpenApiVO) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDuration

`func (o *VoucherSummaryOpenApiVO) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VoucherSummaryOpenApiVO) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VoucherSummaryOpenApiVO) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VoucherSummaryOpenApiVO) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


