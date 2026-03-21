# VoucherUnitPriceDistributionOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Currency Short Code of voucher. For the values of Currency Short Code, refer to section 5.4.2 of the Open API Access Guide. | [optional] 
**TotalAmount** | Pointer to **string** | Total amount of vouchers | [optional] 
**UnitPrice** | Pointer to **string** | Price of single voucher. It should be within the range of 1–999999999 | [optional] 
**UsedCount** | Pointer to **int32** | Used count of vouchers | [optional] 

## Methods

### NewVoucherUnitPriceDistributionOpenApiVO

`func NewVoucherUnitPriceDistributionOpenApiVO() *VoucherUnitPriceDistributionOpenApiVO`

NewVoucherUnitPriceDistributionOpenApiVO instantiates a new VoucherUnitPriceDistributionOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoucherUnitPriceDistributionOpenApiVOWithDefaults

`func NewVoucherUnitPriceDistributionOpenApiVOWithDefaults() *VoucherUnitPriceDistributionOpenApiVO`

NewVoucherUnitPriceDistributionOpenApiVOWithDefaults instantiates a new VoucherUnitPriceDistributionOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *VoucherUnitPriceDistributionOpenApiVO) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *VoucherUnitPriceDistributionOpenApiVO) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *VoucherUnitPriceDistributionOpenApiVO) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *VoucherUnitPriceDistributionOpenApiVO) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTotalAmount

`func (o *VoucherUnitPriceDistributionOpenApiVO) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *VoucherUnitPriceDistributionOpenApiVO) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *VoucherUnitPriceDistributionOpenApiVO) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *VoucherUnitPriceDistributionOpenApiVO) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetUnitPrice

`func (o *VoucherUnitPriceDistributionOpenApiVO) GetUnitPrice() string`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *VoucherUnitPriceDistributionOpenApiVO) GetUnitPriceOk() (*string, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *VoucherUnitPriceDistributionOpenApiVO) SetUnitPrice(v string)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *VoucherUnitPriceDistributionOpenApiVO) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetUsedCount

`func (o *VoucherUnitPriceDistributionOpenApiVO) GetUsedCount() int32`

GetUsedCount returns the UsedCount field if non-nil, zero value otherwise.

### GetUsedCountOk

`func (o *VoucherUnitPriceDistributionOpenApiVO) GetUsedCountOk() (*int32, bool)`

GetUsedCountOk returns a tuple with the UsedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCount

`func (o *VoucherUnitPriceDistributionOpenApiVO) SetUsedCount(v int32)`

SetUsedCount sets UsedCount field to given value.

### HasUsedCount

`func (o *VoucherUnitPriceDistributionOpenApiVO) HasUsedCount() bool`

HasUsedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


