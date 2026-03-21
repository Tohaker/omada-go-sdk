# PpskProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiration** | Pointer to [**PPSKExpirationVO**](PPSKExpirationVO.md) |  | [optional] 
**Ppsk** | [**[]PpskSetting**](PpskSetting.md) | PPSK List In the PPSK Profile | 
**ProfileName** | **string** | PPSK Profile Name, should contain 1 to 64 characters. | 
**RateLimit** | Pointer to [**PPSKRateLimitSettingVO**](PPSKRateLimitSettingVO.md) |  | [optional] 
**Type** | Pointer to **int32** | This field has been deprecated since version 6.1. PPSK Profile type: 0：PPSK Without RADIUS;Cloud Based Controller only support PPSK Without RADIUS. | [optional] 

## Methods

### NewPpskProfile

`func NewPpskProfile(ppsk []PpskSetting, profileName string, ) *PpskProfile`

NewPpskProfile instantiates a new PpskProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPpskProfileWithDefaults

`func NewPpskProfileWithDefaults() *PpskProfile`

NewPpskProfileWithDefaults instantiates a new PpskProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiration

`func (o *PpskProfile) GetExpiration() PPSKExpirationVO`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *PpskProfile) GetExpirationOk() (*PPSKExpirationVO, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *PpskProfile) SetExpiration(v PPSKExpirationVO)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *PpskProfile) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetPpsk

`func (o *PpskProfile) GetPpsk() []PpskSetting`

GetPpsk returns the Ppsk field if non-nil, zero value otherwise.

### GetPpskOk

`func (o *PpskProfile) GetPpskOk() (*[]PpskSetting, bool)`

GetPpskOk returns a tuple with the Ppsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpsk

`func (o *PpskProfile) SetPpsk(v []PpskSetting)`

SetPpsk sets Ppsk field to given value.


### GetProfileName

`func (o *PpskProfile) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *PpskProfile) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *PpskProfile) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.


### GetRateLimit

`func (o *PpskProfile) GetRateLimit() PPSKRateLimitSettingVO`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *PpskProfile) GetRateLimitOk() (*PPSKRateLimitSettingVO, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *PpskProfile) SetRateLimit(v PPSKRateLimitSettingVO)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *PpskProfile) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetType

`func (o *PpskProfile) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PpskProfile) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PpskProfile) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *PpskProfile) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


