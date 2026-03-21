# CorrectSimQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **float32** | The amount of data usage(KB) in current billing cycle. | [optional] 
**SimCard** | Pointer to **int32** | When device supports Dual-SIM card, using parameter [simCard] to point which card to configure. 1: SIM1; 2:SIM2. | [optional] 
**Sms** | Pointer to **int32** | The amount of SMS in current billing cycle, valid date should be within the range of 0–100000. | [optional] 

## Methods

### NewCorrectSimQuota

`func NewCorrectSimQuota() *CorrectSimQuota`

NewCorrectSimQuota instantiates a new CorrectSimQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorrectSimQuotaWithDefaults

`func NewCorrectSimQuotaWithDefaults() *CorrectSimQuota`

NewCorrectSimQuotaWithDefaults instantiates a new CorrectSimQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CorrectSimQuota) GetData() float32`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CorrectSimQuota) GetDataOk() (*float32, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CorrectSimQuota) SetData(v float32)`

SetData sets Data field to given value.

### HasData

`func (o *CorrectSimQuota) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSimCard

`func (o *CorrectSimQuota) GetSimCard() int32`

GetSimCard returns the SimCard field if non-nil, zero value otherwise.

### GetSimCardOk

`func (o *CorrectSimQuota) GetSimCardOk() (*int32, bool)`

GetSimCardOk returns a tuple with the SimCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCard

`func (o *CorrectSimQuota) SetSimCard(v int32)`

SetSimCard sets SimCard field to given value.

### HasSimCard

`func (o *CorrectSimQuota) HasSimCard() bool`

HasSimCard returns a boolean if a field has been set.

### GetSms

`func (o *CorrectSimQuota) GetSms() int32`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *CorrectSimQuota) GetSmsOk() (*int32, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *CorrectSimQuota) SetSms(v int32)`

SetSms sets Sms field to given value.

### HasSms

`func (o *CorrectSimQuota) HasSms() bool`

HasSms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


