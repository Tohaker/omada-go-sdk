# TelephoneNumberAdvancedSettingApOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DscpForRtp** | Pointer to **int32** | The dscp For Rtp of telephone number. | [optional] 
**DscpForSip** | Pointer to **int32** | The dscp For Sip of telephone number. | [optional] 
**DtmfRelaySetting** | Pointer to **int32** | The dtmf relay setting of telephone number. | [optional] 
**EndWithNumberSign** | Pointer to **bool** | Whether the telephone number setting end with number sign. | [optional] 
**ExpirationTime** | Pointer to **int32** | The expiration time of telephone number. | [optional] 
**Locale** | Pointer to **int32** | The country code of telephone number. | [optional] 
**RetryInterval** | Pointer to **int32** | The retry interval of telephone number. | [optional] 
**T38Support** | Pointer to **bool** | Whether to enable t38 Support. | [optional] 

## Methods

### NewTelephoneNumberAdvancedSettingApOpenApiVO

`func NewTelephoneNumberAdvancedSettingApOpenApiVO() *TelephoneNumberAdvancedSettingApOpenApiVO`

NewTelephoneNumberAdvancedSettingApOpenApiVO instantiates a new TelephoneNumberAdvancedSettingApOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelephoneNumberAdvancedSettingApOpenApiVOWithDefaults

`func NewTelephoneNumberAdvancedSettingApOpenApiVOWithDefaults() *TelephoneNumberAdvancedSettingApOpenApiVO`

NewTelephoneNumberAdvancedSettingApOpenApiVOWithDefaults instantiates a new TelephoneNumberAdvancedSettingApOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDscpForRtp

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) GetDscpForRtp() int32`

GetDscpForRtp returns the DscpForRtp field if non-nil, zero value otherwise.

### GetDscpForRtpOk

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) GetDscpForRtpOk() (*int32, bool)`

GetDscpForRtpOk returns a tuple with the DscpForRtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscpForRtp

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) SetDscpForRtp(v int32)`

SetDscpForRtp sets DscpForRtp field to given value.

### HasDscpForRtp

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) HasDscpForRtp() bool`

HasDscpForRtp returns a boolean if a field has been set.

### GetDscpForSip

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) GetDscpForSip() int32`

GetDscpForSip returns the DscpForSip field if non-nil, zero value otherwise.

### GetDscpForSipOk

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) GetDscpForSipOk() (*int32, bool)`

GetDscpForSipOk returns a tuple with the DscpForSip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscpForSip

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) SetDscpForSip(v int32)`

SetDscpForSip sets DscpForSip field to given value.

### HasDscpForSip

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) HasDscpForSip() bool`

HasDscpForSip returns a boolean if a field has been set.

### GetDtmfRelaySetting

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) GetDtmfRelaySetting() int32`

GetDtmfRelaySetting returns the DtmfRelaySetting field if non-nil, zero value otherwise.

### GetDtmfRelaySettingOk

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) GetDtmfRelaySettingOk() (*int32, bool)`

GetDtmfRelaySettingOk returns a tuple with the DtmfRelaySetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfRelaySetting

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) SetDtmfRelaySetting(v int32)`

SetDtmfRelaySetting sets DtmfRelaySetting field to given value.

### HasDtmfRelaySetting

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) HasDtmfRelaySetting() bool`

HasDtmfRelaySetting returns a boolean if a field has been set.

### GetEndWithNumberSign

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) GetEndWithNumberSign() bool`

GetEndWithNumberSign returns the EndWithNumberSign field if non-nil, zero value otherwise.

### GetEndWithNumberSignOk

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) GetEndWithNumberSignOk() (*bool, bool)`

GetEndWithNumberSignOk returns a tuple with the EndWithNumberSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndWithNumberSign

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) SetEndWithNumberSign(v bool)`

SetEndWithNumberSign sets EndWithNumberSign field to given value.

### HasEndWithNumberSign

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) HasEndWithNumberSign() bool`

HasEndWithNumberSign returns a boolean if a field has been set.

### GetExpirationTime

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) GetExpirationTime() int32`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) GetExpirationTimeOk() (*int32, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) SetExpirationTime(v int32)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetLocale

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) GetLocale() int32`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) GetLocaleOk() (*int32, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) SetLocale(v int32)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetRetryInterval

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) GetRetryInterval() int32`

GetRetryInterval returns the RetryInterval field if non-nil, zero value otherwise.

### GetRetryIntervalOk

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) GetRetryIntervalOk() (*int32, bool)`

GetRetryIntervalOk returns a tuple with the RetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryInterval

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) SetRetryInterval(v int32)`

SetRetryInterval sets RetryInterval field to given value.

### HasRetryInterval

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) HasRetryInterval() bool`

HasRetryInterval returns a boolean if a field has been set.

### GetT38Support

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) GetT38Support() bool`

GetT38Support returns the T38Support field if non-nil, zero value otherwise.

### GetT38SupportOk

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) GetT38SupportOk() (*bool, bool)`

GetT38SupportOk returns a tuple with the T38Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT38Support

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) SetT38Support(v bool)`

SetT38Support sets T38Support field to given value.

### HasT38Support

`func (o *TelephoneNumberAdvancedSettingApOpenApiVO) HasT38Support() bool`

HasT38Support returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


