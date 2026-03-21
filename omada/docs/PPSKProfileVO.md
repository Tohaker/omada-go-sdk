# PPSKProfileVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiration** | Pointer to [**PPSKExpirationVO**](PPSKExpirationVO.md) |  | [optional] 
**Id** | Pointer to **string** | PPSK Profile ID | [optional] 
**Ppsk** | [**[]PSKVO**](PSKVO.md) | PSK List In the PPSK Profile | 
**ProfileName** | **string** | PPSK Profile Name | 
**RateLimit** | Pointer to [**PPSKRateLimitSettingVO**](PPSKRateLimitSettingVO.md) |  | [optional] 
**Resource** | Pointer to **int32** |  | [optional] 
**Ssid** | Pointer to **[]string** | Ssid List Bound With PPSK Profile | [optional] 
**Type** | Pointer to **int32** | PPSK Profile type: 0：PPSK Without RADIUS; 1: PPSK With Built-In RADIUS. | [optional] 

## Methods

### NewPPSKProfileVO

`func NewPPSKProfileVO(ppsk []PSKVO, profileName string, ) *PPSKProfileVO`

NewPPSKProfileVO instantiates a new PPSKProfileVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPPSKProfileVOWithDefaults

`func NewPPSKProfileVOWithDefaults() *PPSKProfileVO`

NewPPSKProfileVOWithDefaults instantiates a new PPSKProfileVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiration

`func (o *PPSKProfileVO) GetExpiration() PPSKExpirationVO`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *PPSKProfileVO) GetExpirationOk() (*PPSKExpirationVO, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *PPSKProfileVO) SetExpiration(v PPSKExpirationVO)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *PPSKProfileVO) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetId

`func (o *PPSKProfileVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PPSKProfileVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PPSKProfileVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PPSKProfileVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPpsk

`func (o *PPSKProfileVO) GetPpsk() []PSKVO`

GetPpsk returns the Ppsk field if non-nil, zero value otherwise.

### GetPpskOk

`func (o *PPSKProfileVO) GetPpskOk() (*[]PSKVO, bool)`

GetPpskOk returns a tuple with the Ppsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpsk

`func (o *PPSKProfileVO) SetPpsk(v []PSKVO)`

SetPpsk sets Ppsk field to given value.


### GetProfileName

`func (o *PPSKProfileVO) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *PPSKProfileVO) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *PPSKProfileVO) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.


### GetRateLimit

`func (o *PPSKProfileVO) GetRateLimit() PPSKRateLimitSettingVO`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *PPSKProfileVO) GetRateLimitOk() (*PPSKRateLimitSettingVO, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *PPSKProfileVO) SetRateLimit(v PPSKRateLimitSettingVO)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *PPSKProfileVO) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetResource

`func (o *PPSKProfileVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *PPSKProfileVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *PPSKProfileVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *PPSKProfileVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSsid

`func (o *PPSKProfileVO) GetSsid() []string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *PPSKProfileVO) GetSsidOk() (*[]string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *PPSKProfileVO) SetSsid(v []string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *PPSKProfileVO) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetType

`func (o *PPSKProfileVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PPSKProfileVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PPSKProfileVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *PPSKProfileVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


