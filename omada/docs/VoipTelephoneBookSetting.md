# VoipTelephoneBookSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactId** | Pointer to **string** | The contact ID of contact person. | [optional] 
**FirstName** | Pointer to **string** | The firstName of contact person. It should contain 0 to 64 characters. | [optional] 
**LastName** | Pointer to **string** | The lastName of contact person. It should contain 0 to 64 characters. | [optional] 
**MobilePhoneNumber** | Pointer to **string** | The mobilePhoneNumber of contact person. | [optional] 
**PrivatePhoneNumber** | Pointer to **string** | The privatePhoneNumber of contact person. | [optional] 
**SpeedDialEnable** | Pointer to **bool** | Whether to enable the speedDial. | [optional] 
**SpeedDialNumber** | Pointer to **string** | The speedDialNumber of contact person. | [optional] 
**SpeedDialNumberType** | Pointer to **int32** | speedDialNumberType should be a value as follows: 0: Private Phone Number; 1: Work Phone Number; 2: Mobile Phone Number. | [optional] 
**WorkPhoneNumber** | Pointer to **string** | The workPhoneNumber of contact person. | [optional] 

## Methods

### NewVoipTelephoneBookSetting

`func NewVoipTelephoneBookSetting() *VoipTelephoneBookSetting`

NewVoipTelephoneBookSetting instantiates a new VoipTelephoneBookSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoipTelephoneBookSettingWithDefaults

`func NewVoipTelephoneBookSettingWithDefaults() *VoipTelephoneBookSetting`

NewVoipTelephoneBookSettingWithDefaults instantiates a new VoipTelephoneBookSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactId

`func (o *VoipTelephoneBookSetting) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *VoipTelephoneBookSetting) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *VoipTelephoneBookSetting) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *VoipTelephoneBookSetting) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### GetFirstName

`func (o *VoipTelephoneBookSetting) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *VoipTelephoneBookSetting) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *VoipTelephoneBookSetting) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *VoipTelephoneBookSetting) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *VoipTelephoneBookSetting) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *VoipTelephoneBookSetting) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *VoipTelephoneBookSetting) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *VoipTelephoneBookSetting) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMobilePhoneNumber

`func (o *VoipTelephoneBookSetting) GetMobilePhoneNumber() string`

GetMobilePhoneNumber returns the MobilePhoneNumber field if non-nil, zero value otherwise.

### GetMobilePhoneNumberOk

`func (o *VoipTelephoneBookSetting) GetMobilePhoneNumberOk() (*string, bool)`

GetMobilePhoneNumberOk returns a tuple with the MobilePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneNumber

`func (o *VoipTelephoneBookSetting) SetMobilePhoneNumber(v string)`

SetMobilePhoneNumber sets MobilePhoneNumber field to given value.

### HasMobilePhoneNumber

`func (o *VoipTelephoneBookSetting) HasMobilePhoneNumber() bool`

HasMobilePhoneNumber returns a boolean if a field has been set.

### GetPrivatePhoneNumber

`func (o *VoipTelephoneBookSetting) GetPrivatePhoneNumber() string`

GetPrivatePhoneNumber returns the PrivatePhoneNumber field if non-nil, zero value otherwise.

### GetPrivatePhoneNumberOk

`func (o *VoipTelephoneBookSetting) GetPrivatePhoneNumberOk() (*string, bool)`

GetPrivatePhoneNumberOk returns a tuple with the PrivatePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivatePhoneNumber

`func (o *VoipTelephoneBookSetting) SetPrivatePhoneNumber(v string)`

SetPrivatePhoneNumber sets PrivatePhoneNumber field to given value.

### HasPrivatePhoneNumber

`func (o *VoipTelephoneBookSetting) HasPrivatePhoneNumber() bool`

HasPrivatePhoneNumber returns a boolean if a field has been set.

### GetSpeedDialEnable

`func (o *VoipTelephoneBookSetting) GetSpeedDialEnable() bool`

GetSpeedDialEnable returns the SpeedDialEnable field if non-nil, zero value otherwise.

### GetSpeedDialEnableOk

`func (o *VoipTelephoneBookSetting) GetSpeedDialEnableOk() (*bool, bool)`

GetSpeedDialEnableOk returns a tuple with the SpeedDialEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedDialEnable

`func (o *VoipTelephoneBookSetting) SetSpeedDialEnable(v bool)`

SetSpeedDialEnable sets SpeedDialEnable field to given value.

### HasSpeedDialEnable

`func (o *VoipTelephoneBookSetting) HasSpeedDialEnable() bool`

HasSpeedDialEnable returns a boolean if a field has been set.

### GetSpeedDialNumber

`func (o *VoipTelephoneBookSetting) GetSpeedDialNumber() string`

GetSpeedDialNumber returns the SpeedDialNumber field if non-nil, zero value otherwise.

### GetSpeedDialNumberOk

`func (o *VoipTelephoneBookSetting) GetSpeedDialNumberOk() (*string, bool)`

GetSpeedDialNumberOk returns a tuple with the SpeedDialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedDialNumber

`func (o *VoipTelephoneBookSetting) SetSpeedDialNumber(v string)`

SetSpeedDialNumber sets SpeedDialNumber field to given value.

### HasSpeedDialNumber

`func (o *VoipTelephoneBookSetting) HasSpeedDialNumber() bool`

HasSpeedDialNumber returns a boolean if a field has been set.

### GetSpeedDialNumberType

`func (o *VoipTelephoneBookSetting) GetSpeedDialNumberType() int32`

GetSpeedDialNumberType returns the SpeedDialNumberType field if non-nil, zero value otherwise.

### GetSpeedDialNumberTypeOk

`func (o *VoipTelephoneBookSetting) GetSpeedDialNumberTypeOk() (*int32, bool)`

GetSpeedDialNumberTypeOk returns a tuple with the SpeedDialNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedDialNumberType

`func (o *VoipTelephoneBookSetting) SetSpeedDialNumberType(v int32)`

SetSpeedDialNumberType sets SpeedDialNumberType field to given value.

### HasSpeedDialNumberType

`func (o *VoipTelephoneBookSetting) HasSpeedDialNumberType() bool`

HasSpeedDialNumberType returns a boolean if a field has been set.

### GetWorkPhoneNumber

`func (o *VoipTelephoneBookSetting) GetWorkPhoneNumber() string`

GetWorkPhoneNumber returns the WorkPhoneNumber field if non-nil, zero value otherwise.

### GetWorkPhoneNumberOk

`func (o *VoipTelephoneBookSetting) GetWorkPhoneNumberOk() (*string, bool)`

GetWorkPhoneNumberOk returns a tuple with the WorkPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkPhoneNumber

`func (o *VoipTelephoneBookSetting) SetWorkPhoneNumber(v string)`

SetWorkPhoneNumber sets WorkPhoneNumber field to given value.

### HasWorkPhoneNumber

`func (o *VoipTelephoneBookSetting) HasWorkPhoneNumber() bool`

HasWorkPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


