# VoipContactPersonSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | The firstName of contact person. It should contain 0 to 64 characters. | [optional] 
**LastName** | Pointer to **string** | The lastName of contact person. It should contain 0 to 64 characters. | [optional] 
**MobilePhoneNumber** | Pointer to **string** | The mobilePhoneNumber of contact person. | [optional] 
**PrivatePhoneNumber** | Pointer to **string** | The privatePhoneNumber of contact person. | [optional] 
**SpeedDialEnable** | Pointer to **bool** | Whether to enable the speedDial. | [optional] 
**SpeedDialNumber** | Pointer to **string** | The speedDialNumber of contact person. | [optional] 
**SpeedDialNumberType** | Pointer to **int32** | speedDialNumberType should be a value as follows: 0: Private Phone Number; 1: Work Phone Number; 2: Mobile Phone Number. | [optional] 
**WorkPhoneNumber** | Pointer to **string** | The workPhoneNumber of contact person. | [optional] 

## Methods

### NewVoipContactPersonSettings

`func NewVoipContactPersonSettings() *VoipContactPersonSettings`

NewVoipContactPersonSettings instantiates a new VoipContactPersonSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoipContactPersonSettingsWithDefaults

`func NewVoipContactPersonSettingsWithDefaults() *VoipContactPersonSettings`

NewVoipContactPersonSettingsWithDefaults instantiates a new VoipContactPersonSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *VoipContactPersonSettings) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *VoipContactPersonSettings) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *VoipContactPersonSettings) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *VoipContactPersonSettings) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *VoipContactPersonSettings) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *VoipContactPersonSettings) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *VoipContactPersonSettings) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *VoipContactPersonSettings) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMobilePhoneNumber

`func (o *VoipContactPersonSettings) GetMobilePhoneNumber() string`

GetMobilePhoneNumber returns the MobilePhoneNumber field if non-nil, zero value otherwise.

### GetMobilePhoneNumberOk

`func (o *VoipContactPersonSettings) GetMobilePhoneNumberOk() (*string, bool)`

GetMobilePhoneNumberOk returns a tuple with the MobilePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneNumber

`func (o *VoipContactPersonSettings) SetMobilePhoneNumber(v string)`

SetMobilePhoneNumber sets MobilePhoneNumber field to given value.

### HasMobilePhoneNumber

`func (o *VoipContactPersonSettings) HasMobilePhoneNumber() bool`

HasMobilePhoneNumber returns a boolean if a field has been set.

### GetPrivatePhoneNumber

`func (o *VoipContactPersonSettings) GetPrivatePhoneNumber() string`

GetPrivatePhoneNumber returns the PrivatePhoneNumber field if non-nil, zero value otherwise.

### GetPrivatePhoneNumberOk

`func (o *VoipContactPersonSettings) GetPrivatePhoneNumberOk() (*string, bool)`

GetPrivatePhoneNumberOk returns a tuple with the PrivatePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivatePhoneNumber

`func (o *VoipContactPersonSettings) SetPrivatePhoneNumber(v string)`

SetPrivatePhoneNumber sets PrivatePhoneNumber field to given value.

### HasPrivatePhoneNumber

`func (o *VoipContactPersonSettings) HasPrivatePhoneNumber() bool`

HasPrivatePhoneNumber returns a boolean if a field has been set.

### GetSpeedDialEnable

`func (o *VoipContactPersonSettings) GetSpeedDialEnable() bool`

GetSpeedDialEnable returns the SpeedDialEnable field if non-nil, zero value otherwise.

### GetSpeedDialEnableOk

`func (o *VoipContactPersonSettings) GetSpeedDialEnableOk() (*bool, bool)`

GetSpeedDialEnableOk returns a tuple with the SpeedDialEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedDialEnable

`func (o *VoipContactPersonSettings) SetSpeedDialEnable(v bool)`

SetSpeedDialEnable sets SpeedDialEnable field to given value.

### HasSpeedDialEnable

`func (o *VoipContactPersonSettings) HasSpeedDialEnable() bool`

HasSpeedDialEnable returns a boolean if a field has been set.

### GetSpeedDialNumber

`func (o *VoipContactPersonSettings) GetSpeedDialNumber() string`

GetSpeedDialNumber returns the SpeedDialNumber field if non-nil, zero value otherwise.

### GetSpeedDialNumberOk

`func (o *VoipContactPersonSettings) GetSpeedDialNumberOk() (*string, bool)`

GetSpeedDialNumberOk returns a tuple with the SpeedDialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedDialNumber

`func (o *VoipContactPersonSettings) SetSpeedDialNumber(v string)`

SetSpeedDialNumber sets SpeedDialNumber field to given value.

### HasSpeedDialNumber

`func (o *VoipContactPersonSettings) HasSpeedDialNumber() bool`

HasSpeedDialNumber returns a boolean if a field has been set.

### GetSpeedDialNumberType

`func (o *VoipContactPersonSettings) GetSpeedDialNumberType() int32`

GetSpeedDialNumberType returns the SpeedDialNumberType field if non-nil, zero value otherwise.

### GetSpeedDialNumberTypeOk

`func (o *VoipContactPersonSettings) GetSpeedDialNumberTypeOk() (*int32, bool)`

GetSpeedDialNumberTypeOk returns a tuple with the SpeedDialNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedDialNumberType

`func (o *VoipContactPersonSettings) SetSpeedDialNumberType(v int32)`

SetSpeedDialNumberType sets SpeedDialNumberType field to given value.

### HasSpeedDialNumberType

`func (o *VoipContactPersonSettings) HasSpeedDialNumberType() bool`

HasSpeedDialNumberType returns a boolean if a field has been set.

### GetWorkPhoneNumber

`func (o *VoipContactPersonSettings) GetWorkPhoneNumber() string`

GetWorkPhoneNumber returns the WorkPhoneNumber field if non-nil, zero value otherwise.

### GetWorkPhoneNumberOk

`func (o *VoipContactPersonSettings) GetWorkPhoneNumberOk() (*string, bool)`

GetWorkPhoneNumberOk returns a tuple with the WorkPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkPhoneNumber

`func (o *VoipContactPersonSettings) SetWorkPhoneNumber(v string)`

SetWorkPhoneNumber sets WorkPhoneNumber field to given value.

### HasWorkPhoneNumber

`func (o *VoipContactPersonSettings) HasWorkPhoneNumber() bool`

HasWorkPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


