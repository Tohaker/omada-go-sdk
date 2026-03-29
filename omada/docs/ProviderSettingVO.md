# ProviderSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** |  | [optional] 
**PortSettings** | Pointer to [**PortSettingVO**](PortSettingVO.md) |  | [optional] 
**Provider** | **int32** | The telephony providers supported. The range is between 0 and 13. 0: Other Provider; 1: 1 &amp; 1 Internet; 2: Congstar; 3: Vodafone/Arcor; 4: DUS.net; 5: Easybell; 6: Kabel Deutschland; 7: QSC/Q-DSL home; 8: Sipgate.de; 9: Sipgate Team; 10: Sipload; 11: Ventengo; 12: Telekom; 13: Bellsip. The default value is 0. | 
**RegistrarAddress** | Pointer to **string** | When parameter [provider] is a value in[0, 6], parameter [registrarAddress] should not be null. | [optional] 
**Username** | Pointer to **string** | When parameter [provider] is a value in [0, 2, 4, 5, 8, 9, 10, 11, 12, 13], parameter [username] should not be null. | [optional] 

## Methods

### NewProviderSettingVO

`func NewProviderSettingVO(provider int32, ) *ProviderSettingVO`

NewProviderSettingVO instantiates a new ProviderSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSettingVOWithDefaults

`func NewProviderSettingVOWithDefaults() *ProviderSettingVO`

NewProviderSettingVOWithDefaults instantiates a new ProviderSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *ProviderSettingVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ProviderSettingVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ProviderSettingVO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ProviderSettingVO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPortSettings

`func (o *ProviderSettingVO) GetPortSettings() PortSettingVO`

GetPortSettings returns the PortSettings field if non-nil, zero value otherwise.

### GetPortSettingsOk

`func (o *ProviderSettingVO) GetPortSettingsOk() (*PortSettingVO, bool)`

GetPortSettingsOk returns a tuple with the PortSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSettings

`func (o *ProviderSettingVO) SetPortSettings(v PortSettingVO)`

SetPortSettings sets PortSettings field to given value.

### HasPortSettings

`func (o *ProviderSettingVO) HasPortSettings() bool`

HasPortSettings returns a boolean if a field has been set.

### GetProvider

`func (o *ProviderSettingVO) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ProviderSettingVO) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ProviderSettingVO) SetProvider(v int32)`

SetProvider sets Provider field to given value.


### GetRegistrarAddress

`func (o *ProviderSettingVO) GetRegistrarAddress() string`

GetRegistrarAddress returns the RegistrarAddress field if non-nil, zero value otherwise.

### GetRegistrarAddressOk

`func (o *ProviderSettingVO) GetRegistrarAddressOk() (*string, bool)`

GetRegistrarAddressOk returns a tuple with the RegistrarAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrarAddress

`func (o *ProviderSettingVO) SetRegistrarAddress(v string)`

SetRegistrarAddress sets RegistrarAddress field to given value.

### HasRegistrarAddress

`func (o *ProviderSettingVO) HasRegistrarAddress() bool`

HasRegistrarAddress returns a boolean if a field has been set.

### GetUsername

`func (o *ProviderSettingVO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ProviderSettingVO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ProviderSettingVO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ProviderSettingVO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


