# PpskProfileV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoCreatePsks** | **bool** | Whether to enable auto create psks | 
**Expiration** | Pointer to [**PPSKExpirationVO**](PPSKExpirationVO.md) |  | [optional] 
**Length** | Pointer to **int32** | This field is required when Parameter [autoCreatePsks] is true; PSK Password Length, should be within the range of 8-63. | [optional] 
**Number** | Pointer to **int32** | This field is required when Parameter [autoCreatePsks] is true; Generate Number, should be within the range of 1-1024. | [optional] 
**Ppsk** | Pointer to [**[]PpskSettingV2**](PpskSettingV2.md) | This field is required when Parameter [autoCreatePsks] is false; PPSK List In the PPSK Profile | [optional] 
**Prefix** | Pointer to **string** | This field is required when Parameter [autoCreatePsks] is true; PSK Name Prefix, should contain 1 to 60 visible ASCII characters. | [optional] 
**ProfileName** | **string** | PPSK Profile Name, should contain 1 to 64 characters. | 
**RateLimit** | Pointer to [**PPSKRateLimitSettingVO**](PPSKRateLimitSettingVO.md) |  | [optional] 
**Vlan** | Pointer to **int32** | This field is required when Parameter [autoCreatePsks] is true; PSK Bound Vlan, should be within the range of 1-4094. | [optional] 
**VlanInterval** | Pointer to **int32** | This field is required when Parameter [autoCreatePsks] is true; The interval of vlan when auto creating psk | [optional] 
**VlanPool** | Pointer to **string** | This field is required when Parameter [autoCreatePsks] is true; PSK Bound Vlan range, should be like: 10-1000. | [optional] 

## Methods

### NewPpskProfileV2

`func NewPpskProfileV2(autoCreatePsks bool, profileName string, ) *PpskProfileV2`

NewPpskProfileV2 instantiates a new PpskProfileV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPpskProfileV2WithDefaults

`func NewPpskProfileV2WithDefaults() *PpskProfileV2`

NewPpskProfileV2WithDefaults instantiates a new PpskProfileV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoCreatePsks

`func (o *PpskProfileV2) GetAutoCreatePsks() bool`

GetAutoCreatePsks returns the AutoCreatePsks field if non-nil, zero value otherwise.

### GetAutoCreatePsksOk

`func (o *PpskProfileV2) GetAutoCreatePsksOk() (*bool, bool)`

GetAutoCreatePsksOk returns a tuple with the AutoCreatePsks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreatePsks

`func (o *PpskProfileV2) SetAutoCreatePsks(v bool)`

SetAutoCreatePsks sets AutoCreatePsks field to given value.


### GetExpiration

`func (o *PpskProfileV2) GetExpiration() PPSKExpirationVO`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *PpskProfileV2) GetExpirationOk() (*PPSKExpirationVO, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *PpskProfileV2) SetExpiration(v PPSKExpirationVO)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *PpskProfileV2) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetLength

`func (o *PpskProfileV2) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *PpskProfileV2) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *PpskProfileV2) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *PpskProfileV2) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetNumber

`func (o *PpskProfileV2) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PpskProfileV2) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PpskProfileV2) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PpskProfileV2) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPpsk

`func (o *PpskProfileV2) GetPpsk() []PpskSettingV2`

GetPpsk returns the Ppsk field if non-nil, zero value otherwise.

### GetPpskOk

`func (o *PpskProfileV2) GetPpskOk() (*[]PpskSettingV2, bool)`

GetPpskOk returns a tuple with the Ppsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpsk

`func (o *PpskProfileV2) SetPpsk(v []PpskSettingV2)`

SetPpsk sets Ppsk field to given value.

### HasPpsk

`func (o *PpskProfileV2) HasPpsk() bool`

HasPpsk returns a boolean if a field has been set.

### GetPrefix

`func (o *PpskProfileV2) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PpskProfileV2) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PpskProfileV2) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PpskProfileV2) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetProfileName

`func (o *PpskProfileV2) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *PpskProfileV2) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *PpskProfileV2) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.


### GetRateLimit

`func (o *PpskProfileV2) GetRateLimit() PPSKRateLimitSettingVO`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *PpskProfileV2) GetRateLimitOk() (*PPSKRateLimitSettingVO, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *PpskProfileV2) SetRateLimit(v PPSKRateLimitSettingVO)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *PpskProfileV2) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetVlan

`func (o *PpskProfileV2) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *PpskProfileV2) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *PpskProfileV2) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *PpskProfileV2) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVlanInterval

`func (o *PpskProfileV2) GetVlanInterval() int32`

GetVlanInterval returns the VlanInterval field if non-nil, zero value otherwise.

### GetVlanIntervalOk

`func (o *PpskProfileV2) GetVlanIntervalOk() (*int32, bool)`

GetVlanIntervalOk returns a tuple with the VlanInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanInterval

`func (o *PpskProfileV2) SetVlanInterval(v int32)`

SetVlanInterval sets VlanInterval field to given value.

### HasVlanInterval

`func (o *PpskProfileV2) HasVlanInterval() bool`

HasVlanInterval returns a boolean if a field has been set.

### GetVlanPool

`func (o *PpskProfileV2) GetVlanPool() string`

GetVlanPool returns the VlanPool field if non-nil, zero value otherwise.

### GetVlanPoolOk

`func (o *PpskProfileV2) GetVlanPoolOk() (*string, bool)`

GetVlanPoolOk returns a tuple with the VlanPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPool

`func (o *PpskProfileV2) SetVlanPool(v string)`

SetVlanPool sets VlanPool field to given value.

### HasVlanPool

`func (o *PpskProfileV2) HasVlanPool() bool`

HasVlanPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


