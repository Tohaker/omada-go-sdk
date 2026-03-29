# VoucherDurationDistributionOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int64** | Duration of single voucher, unit: minutes | [optional] 
**DurationType** | Pointer to **int32** | The duration type of the voucher. It should be a value as follows: 0: Client duration, each client expires after the duration is used. 1: Voucher duration, after reaching the voucher duration, clients using the voucher will expire | [optional] 
**TotalDuration** | Pointer to **int64** | Total duration of vouchers, unit: minutes | [optional] 
**UsedCount** | Pointer to **int32** | Used count of vouchers | [optional] 

## Methods

### NewVoucherDurationDistributionOpenApiVO

`func NewVoucherDurationDistributionOpenApiVO() *VoucherDurationDistributionOpenApiVO`

NewVoucherDurationDistributionOpenApiVO instantiates a new VoucherDurationDistributionOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoucherDurationDistributionOpenApiVOWithDefaults

`func NewVoucherDurationDistributionOpenApiVOWithDefaults() *VoucherDurationDistributionOpenApiVO`

NewVoucherDurationDistributionOpenApiVOWithDefaults instantiates a new VoucherDurationDistributionOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *VoucherDurationDistributionOpenApiVO) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VoucherDurationDistributionOpenApiVO) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VoucherDurationDistributionOpenApiVO) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VoucherDurationDistributionOpenApiVO) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDurationType

`func (o *VoucherDurationDistributionOpenApiVO) GetDurationType() int32`

GetDurationType returns the DurationType field if non-nil, zero value otherwise.

### GetDurationTypeOk

`func (o *VoucherDurationDistributionOpenApiVO) GetDurationTypeOk() (*int32, bool)`

GetDurationTypeOk returns a tuple with the DurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationType

`func (o *VoucherDurationDistributionOpenApiVO) SetDurationType(v int32)`

SetDurationType sets DurationType field to given value.

### HasDurationType

`func (o *VoucherDurationDistributionOpenApiVO) HasDurationType() bool`

HasDurationType returns a boolean if a field has been set.

### GetTotalDuration

`func (o *VoucherDurationDistributionOpenApiVO) GetTotalDuration() int64`

GetTotalDuration returns the TotalDuration field if non-nil, zero value otherwise.

### GetTotalDurationOk

`func (o *VoucherDurationDistributionOpenApiVO) GetTotalDurationOk() (*int64, bool)`

GetTotalDurationOk returns a tuple with the TotalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDuration

`func (o *VoucherDurationDistributionOpenApiVO) SetTotalDuration(v int64)`

SetTotalDuration sets TotalDuration field to given value.

### HasTotalDuration

`func (o *VoucherDurationDistributionOpenApiVO) HasTotalDuration() bool`

HasTotalDuration returns a boolean if a field has been set.

### GetUsedCount

`func (o *VoucherDurationDistributionOpenApiVO) GetUsedCount() int32`

GetUsedCount returns the UsedCount field if non-nil, zero value otherwise.

### GetUsedCountOk

`func (o *VoucherDurationDistributionOpenApiVO) GetUsedCountOk() (*int32, bool)`

GetUsedCountOk returns a tuple with the UsedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCount

`func (o *VoucherDurationDistributionOpenApiVO) SetUsedCount(v int32)`

SetUsedCount sets UsedCount field to given value.

### HasUsedCount

`func (o *VoucherDurationDistributionOpenApiVO) HasUsedCount() bool`

HasUsedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


