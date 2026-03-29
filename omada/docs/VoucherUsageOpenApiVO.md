# VoucherUsageOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | Amount of vouchers | [optional] 
**Count** | Pointer to **int32** | Count of vouchers | [optional] 
**Currency** | Pointer to **string** | Currency Short Code of voucher. For the values of Currency Short Code, refer to section 5.4.2 of the Open API Access Guide. | [optional] 
**Time** | Pointer to **int64** | Timestamp of the data point, unit: MS | [optional] 
**TimeInterval** | Pointer to **int32** | Time interval of each data point | [optional] 

## Methods

### NewVoucherUsageOpenApiVO

`func NewVoucherUsageOpenApiVO() *VoucherUsageOpenApiVO`

NewVoucherUsageOpenApiVO instantiates a new VoucherUsageOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoucherUsageOpenApiVOWithDefaults

`func NewVoucherUsageOpenApiVOWithDefaults() *VoucherUsageOpenApiVO`

NewVoucherUsageOpenApiVOWithDefaults instantiates a new VoucherUsageOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *VoucherUsageOpenApiVO) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *VoucherUsageOpenApiVO) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *VoucherUsageOpenApiVO) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *VoucherUsageOpenApiVO) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCount

`func (o *VoucherUsageOpenApiVO) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VoucherUsageOpenApiVO) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VoucherUsageOpenApiVO) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *VoucherUsageOpenApiVO) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCurrency

`func (o *VoucherUsageOpenApiVO) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *VoucherUsageOpenApiVO) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *VoucherUsageOpenApiVO) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *VoucherUsageOpenApiVO) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTime

`func (o *VoucherUsageOpenApiVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *VoucherUsageOpenApiVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *VoucherUsageOpenApiVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *VoucherUsageOpenApiVO) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTimeInterval

`func (o *VoucherUsageOpenApiVO) GetTimeInterval() int32`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *VoucherUsageOpenApiVO) GetTimeIntervalOk() (*int32, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *VoucherUsageOpenApiVO) SetTimeInterval(v int32)`

SetTimeInterval sets TimeInterval field to given value.

### HasTimeInterval

`func (o *VoucherUsageOpenApiVO) HasTimeInterval() bool`

HasTimeInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


