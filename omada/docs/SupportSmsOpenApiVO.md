# SupportSmsOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportDualSim** | Pointer to **int32** | Whether the device supports Dual-SIM card. | [optional] 
**SupportDualSms** | Pointer to **int32** | Whether two SIM cards support SMS. 0: SIM1 and SIM2 not support; 1: SIM1 supports, SIM2 not supports; 2: SIM1 not supports, SIM2 supports; 3: SIM1 and SIM2 support. | [optional] 
**SupportLte** | Pointer to **bool** | Whether the device supports LTE. | [optional] 
**SupportSms** | Pointer to **bool** | Whether the using SIM card supports SMS. | [optional] 

## Methods

### NewSupportSmsOpenApiVO

`func NewSupportSmsOpenApiVO() *SupportSmsOpenApiVO`

NewSupportSmsOpenApiVO instantiates a new SupportSmsOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportSmsOpenApiVOWithDefaults

`func NewSupportSmsOpenApiVOWithDefaults() *SupportSmsOpenApiVO`

NewSupportSmsOpenApiVOWithDefaults instantiates a new SupportSmsOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportDualSim

`func (o *SupportSmsOpenApiVO) GetSupportDualSim() int32`

GetSupportDualSim returns the SupportDualSim field if non-nil, zero value otherwise.

### GetSupportDualSimOk

`func (o *SupportSmsOpenApiVO) GetSupportDualSimOk() (*int32, bool)`

GetSupportDualSimOk returns a tuple with the SupportDualSim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDualSim

`func (o *SupportSmsOpenApiVO) SetSupportDualSim(v int32)`

SetSupportDualSim sets SupportDualSim field to given value.

### HasSupportDualSim

`func (o *SupportSmsOpenApiVO) HasSupportDualSim() bool`

HasSupportDualSim returns a boolean if a field has been set.

### GetSupportDualSms

`func (o *SupportSmsOpenApiVO) GetSupportDualSms() int32`

GetSupportDualSms returns the SupportDualSms field if non-nil, zero value otherwise.

### GetSupportDualSmsOk

`func (o *SupportSmsOpenApiVO) GetSupportDualSmsOk() (*int32, bool)`

GetSupportDualSmsOk returns a tuple with the SupportDualSms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDualSms

`func (o *SupportSmsOpenApiVO) SetSupportDualSms(v int32)`

SetSupportDualSms sets SupportDualSms field to given value.

### HasSupportDualSms

`func (o *SupportSmsOpenApiVO) HasSupportDualSms() bool`

HasSupportDualSms returns a boolean if a field has been set.

### GetSupportLte

`func (o *SupportSmsOpenApiVO) GetSupportLte() bool`

GetSupportLte returns the SupportLte field if non-nil, zero value otherwise.

### GetSupportLteOk

`func (o *SupportSmsOpenApiVO) GetSupportLteOk() (*bool, bool)`

GetSupportLteOk returns a tuple with the SupportLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLte

`func (o *SupportSmsOpenApiVO) SetSupportLte(v bool)`

SetSupportLte sets SupportLte field to given value.

### HasSupportLte

`func (o *SupportSmsOpenApiVO) HasSupportLte() bool`

HasSupportLte returns a boolean if a field has been set.

### GetSupportSms

`func (o *SupportSmsOpenApiVO) GetSupportSms() bool`

GetSupportSms returns the SupportSms field if non-nil, zero value otherwise.

### GetSupportSmsOk

`func (o *SupportSmsOpenApiVO) GetSupportSmsOk() (*bool, bool)`

GetSupportSmsOk returns a tuple with the SupportSms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSms

`func (o *SupportSmsOpenApiVO) SetSupportSms(v bool)`

SetSupportSms sets SupportSms field to given value.

### HasSupportSms

`func (o *SupportSmsOpenApiVO) HasSupportSms() bool`

HasSupportSms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


