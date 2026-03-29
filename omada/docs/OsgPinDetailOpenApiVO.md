# OsgPinDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoUnLock** | Pointer to **bool** | Whether or not PIN is automatically unlocked. | [optional] 
**CardStatus** | Pointer to **int32** | cardStatus: 0:no_sim, 1:unknown, 2:pin_lock, 3:pin_verified, 4:ready, 5:puk_lock, 6:blocked | [optional] 
**PinLock** | Pointer to **bool** | Whether the PIN has a lock or not. | [optional] 
**SimCardUsed** | Pointer to **int32** | When the device supports Dual-SIM card, parameter [simCardUsed] should not be null.1: SIM1; 2: SIM2. | [optional] 
**TryPinLimit** | Pointer to **int32** | The number of times the PIN can be entered incorrectly. | [optional] 
**TryPukLimit** | Pointer to **int32** | The number of times the PUK can be entered incorrectly. | [optional] 

## Methods

### NewOsgPinDetailOpenApiVO

`func NewOsgPinDetailOpenApiVO() *OsgPinDetailOpenApiVO`

NewOsgPinDetailOpenApiVO instantiates a new OsgPinDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgPinDetailOpenApiVOWithDefaults

`func NewOsgPinDetailOpenApiVOWithDefaults() *OsgPinDetailOpenApiVO`

NewOsgPinDetailOpenApiVOWithDefaults instantiates a new OsgPinDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoUnLock

`func (o *OsgPinDetailOpenApiVO) GetAutoUnLock() bool`

GetAutoUnLock returns the AutoUnLock field if non-nil, zero value otherwise.

### GetAutoUnLockOk

`func (o *OsgPinDetailOpenApiVO) GetAutoUnLockOk() (*bool, bool)`

GetAutoUnLockOk returns a tuple with the AutoUnLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUnLock

`func (o *OsgPinDetailOpenApiVO) SetAutoUnLock(v bool)`

SetAutoUnLock sets AutoUnLock field to given value.

### HasAutoUnLock

`func (o *OsgPinDetailOpenApiVO) HasAutoUnLock() bool`

HasAutoUnLock returns a boolean if a field has been set.

### GetCardStatus

`func (o *OsgPinDetailOpenApiVO) GetCardStatus() int32`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *OsgPinDetailOpenApiVO) GetCardStatusOk() (*int32, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *OsgPinDetailOpenApiVO) SetCardStatus(v int32)`

SetCardStatus sets CardStatus field to given value.

### HasCardStatus

`func (o *OsgPinDetailOpenApiVO) HasCardStatus() bool`

HasCardStatus returns a boolean if a field has been set.

### GetPinLock

`func (o *OsgPinDetailOpenApiVO) GetPinLock() bool`

GetPinLock returns the PinLock field if non-nil, zero value otherwise.

### GetPinLockOk

`func (o *OsgPinDetailOpenApiVO) GetPinLockOk() (*bool, bool)`

GetPinLockOk returns a tuple with the PinLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinLock

`func (o *OsgPinDetailOpenApiVO) SetPinLock(v bool)`

SetPinLock sets PinLock field to given value.

### HasPinLock

`func (o *OsgPinDetailOpenApiVO) HasPinLock() bool`

HasPinLock returns a boolean if a field has been set.

### GetSimCardUsed

`func (o *OsgPinDetailOpenApiVO) GetSimCardUsed() int32`

GetSimCardUsed returns the SimCardUsed field if non-nil, zero value otherwise.

### GetSimCardUsedOk

`func (o *OsgPinDetailOpenApiVO) GetSimCardUsedOk() (*int32, bool)`

GetSimCardUsedOk returns a tuple with the SimCardUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardUsed

`func (o *OsgPinDetailOpenApiVO) SetSimCardUsed(v int32)`

SetSimCardUsed sets SimCardUsed field to given value.

### HasSimCardUsed

`func (o *OsgPinDetailOpenApiVO) HasSimCardUsed() bool`

HasSimCardUsed returns a boolean if a field has been set.

### GetTryPinLimit

`func (o *OsgPinDetailOpenApiVO) GetTryPinLimit() int32`

GetTryPinLimit returns the TryPinLimit field if non-nil, zero value otherwise.

### GetTryPinLimitOk

`func (o *OsgPinDetailOpenApiVO) GetTryPinLimitOk() (*int32, bool)`

GetTryPinLimitOk returns a tuple with the TryPinLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryPinLimit

`func (o *OsgPinDetailOpenApiVO) SetTryPinLimit(v int32)`

SetTryPinLimit sets TryPinLimit field to given value.

### HasTryPinLimit

`func (o *OsgPinDetailOpenApiVO) HasTryPinLimit() bool`

HasTryPinLimit returns a boolean if a field has been set.

### GetTryPukLimit

`func (o *OsgPinDetailOpenApiVO) GetTryPukLimit() int32`

GetTryPukLimit returns the TryPukLimit field if non-nil, zero value otherwise.

### GetTryPukLimitOk

`func (o *OsgPinDetailOpenApiVO) GetTryPukLimitOk() (*int32, bool)`

GetTryPukLimitOk returns a tuple with the TryPukLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryPukLimit

`func (o *OsgPinDetailOpenApiVO) SetTryPukLimit(v int32)`

SetTryPukLimit sets TryPukLimit field to given value.

### HasTryPukLimit

`func (o *OsgPinDetailOpenApiVO) HasTryPukLimit() bool`

HasTryPukLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


