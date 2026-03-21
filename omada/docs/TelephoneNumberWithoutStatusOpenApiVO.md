# TelephoneNumberWithoutStatusOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMac** | **string** | Device MAC. | 
**Number** | **string** | The number of telephone number. | 
**Password** | Pointer to **string** | Password. | [optional] 
**PhoneNumberId** | Pointer to **string** | The phoneNumber ID of telephone number. | [optional] 
**Prefix** | Pointer to **string** | The prefix of telephone number. When the provider in the provider profile is 1&amp;1 Internet, Vodafone/Arcor, QSC/Q-DSL home, or Telekom, parameter [prefix] is required. When the provider is easybell, parameter [prefix] value is always 0049. In other cases, parameter [prefix] value is always null. | [optional] 
**ProfileId** | **string** | Provider profile ID used to bind phone numbers. | 
**Username** | Pointer to **string** | Username. | [optional] 

## Methods

### NewTelephoneNumberWithoutStatusOpenApiVO

`func NewTelephoneNumberWithoutStatusOpenApiVO(deviceMac string, number string, profileId string, ) *TelephoneNumberWithoutStatusOpenApiVO`

NewTelephoneNumberWithoutStatusOpenApiVO instantiates a new TelephoneNumberWithoutStatusOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelephoneNumberWithoutStatusOpenApiVOWithDefaults

`func NewTelephoneNumberWithoutStatusOpenApiVOWithDefaults() *TelephoneNumberWithoutStatusOpenApiVO`

NewTelephoneNumberWithoutStatusOpenApiVOWithDefaults instantiates a new TelephoneNumberWithoutStatusOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMac

`func (o *TelephoneNumberWithoutStatusOpenApiVO) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *TelephoneNumberWithoutStatusOpenApiVO) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *TelephoneNumberWithoutStatusOpenApiVO) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.


### GetNumber

`func (o *TelephoneNumberWithoutStatusOpenApiVO) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *TelephoneNumberWithoutStatusOpenApiVO) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *TelephoneNumberWithoutStatusOpenApiVO) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetPassword

`func (o *TelephoneNumberWithoutStatusOpenApiVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TelephoneNumberWithoutStatusOpenApiVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TelephoneNumberWithoutStatusOpenApiVO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TelephoneNumberWithoutStatusOpenApiVO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPhoneNumberId

`func (o *TelephoneNumberWithoutStatusOpenApiVO) GetPhoneNumberId() string`

GetPhoneNumberId returns the PhoneNumberId field if non-nil, zero value otherwise.

### GetPhoneNumberIdOk

`func (o *TelephoneNumberWithoutStatusOpenApiVO) GetPhoneNumberIdOk() (*string, bool)`

GetPhoneNumberIdOk returns a tuple with the PhoneNumberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberId

`func (o *TelephoneNumberWithoutStatusOpenApiVO) SetPhoneNumberId(v string)`

SetPhoneNumberId sets PhoneNumberId field to given value.

### HasPhoneNumberId

`func (o *TelephoneNumberWithoutStatusOpenApiVO) HasPhoneNumberId() bool`

HasPhoneNumberId returns a boolean if a field has been set.

### GetPrefix

`func (o *TelephoneNumberWithoutStatusOpenApiVO) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *TelephoneNumberWithoutStatusOpenApiVO) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *TelephoneNumberWithoutStatusOpenApiVO) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *TelephoneNumberWithoutStatusOpenApiVO) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetProfileId

`func (o *TelephoneNumberWithoutStatusOpenApiVO) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *TelephoneNumberWithoutStatusOpenApiVO) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *TelephoneNumberWithoutStatusOpenApiVO) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetUsername

`func (o *TelephoneNumberWithoutStatusOpenApiVO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TelephoneNumberWithoutStatusOpenApiVO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TelephoneNumberWithoutStatusOpenApiVO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TelephoneNumberWithoutStatusOpenApiVO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


