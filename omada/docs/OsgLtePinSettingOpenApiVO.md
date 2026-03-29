# OsgLtePinSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoUnLock** | Pointer to **bool** | Whether or not PIN is automatically unlocked. | [optional] 
**CardStatus** | Pointer to **int32** | cardStatus: 0:no_sim, 1:unknown, 2:pin_lock, 3:pin_verified, 4:ready, 5:puk_lock, 6:blocked | [optional] 
**PinLock** | Pointer to **bool** | Whether the PIN has a lock or not. | [optional] 
**SimCardUsed** | Pointer to **int32** | When the device supports Dual-SIM card, parameter [simCardUsed] should not be null.1: SIM1; 2: SIM2. | [optional] 

## Methods

### NewOsgLtePinSettingOpenApiVO

`func NewOsgLtePinSettingOpenApiVO() *OsgLtePinSettingOpenApiVO`

NewOsgLtePinSettingOpenApiVO instantiates a new OsgLtePinSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgLtePinSettingOpenApiVOWithDefaults

`func NewOsgLtePinSettingOpenApiVOWithDefaults() *OsgLtePinSettingOpenApiVO`

NewOsgLtePinSettingOpenApiVOWithDefaults instantiates a new OsgLtePinSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoUnLock

`func (o *OsgLtePinSettingOpenApiVO) GetAutoUnLock() bool`

GetAutoUnLock returns the AutoUnLock field if non-nil, zero value otherwise.

### GetAutoUnLockOk

`func (o *OsgLtePinSettingOpenApiVO) GetAutoUnLockOk() (*bool, bool)`

GetAutoUnLockOk returns a tuple with the AutoUnLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUnLock

`func (o *OsgLtePinSettingOpenApiVO) SetAutoUnLock(v bool)`

SetAutoUnLock sets AutoUnLock field to given value.

### HasAutoUnLock

`func (o *OsgLtePinSettingOpenApiVO) HasAutoUnLock() bool`

HasAutoUnLock returns a boolean if a field has been set.

### GetCardStatus

`func (o *OsgLtePinSettingOpenApiVO) GetCardStatus() int32`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *OsgLtePinSettingOpenApiVO) GetCardStatusOk() (*int32, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *OsgLtePinSettingOpenApiVO) SetCardStatus(v int32)`

SetCardStatus sets CardStatus field to given value.

### HasCardStatus

`func (o *OsgLtePinSettingOpenApiVO) HasCardStatus() bool`

HasCardStatus returns a boolean if a field has been set.

### GetPinLock

`func (o *OsgLtePinSettingOpenApiVO) GetPinLock() bool`

GetPinLock returns the PinLock field if non-nil, zero value otherwise.

### GetPinLockOk

`func (o *OsgLtePinSettingOpenApiVO) GetPinLockOk() (*bool, bool)`

GetPinLockOk returns a tuple with the PinLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinLock

`func (o *OsgLtePinSettingOpenApiVO) SetPinLock(v bool)`

SetPinLock sets PinLock field to given value.

### HasPinLock

`func (o *OsgLtePinSettingOpenApiVO) HasPinLock() bool`

HasPinLock returns a boolean if a field has been set.

### GetSimCardUsed

`func (o *OsgLtePinSettingOpenApiVO) GetSimCardUsed() int32`

GetSimCardUsed returns the SimCardUsed field if non-nil, zero value otherwise.

### GetSimCardUsedOk

`func (o *OsgLtePinSettingOpenApiVO) GetSimCardUsedOk() (*int32, bool)`

GetSimCardUsedOk returns a tuple with the SimCardUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardUsed

`func (o *OsgLtePinSettingOpenApiVO) SetSimCardUsed(v int32)`

SetSimCardUsed sets SimCardUsed field to given value.

### HasSimCardUsed

`func (o *OsgLtePinSettingOpenApiVO) HasSimCardUsed() bool`

HasSimCardUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


