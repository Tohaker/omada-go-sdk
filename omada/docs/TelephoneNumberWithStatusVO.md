# TelephoneNumberWithStatusVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMac** | Pointer to **string** |  | [optional] 
**FeatureDescription** | Pointer to [**[]FeatureInfoVO**](FeatureInfoVO.md) |  | [optional] 
**Number** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PhoneNumberId** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**ProfileId** | Pointer to **string** |  | [optional] 
**ProfileName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewTelephoneNumberWithStatusVO

`func NewTelephoneNumberWithStatusVO() *TelephoneNumberWithStatusVO`

NewTelephoneNumberWithStatusVO instantiates a new TelephoneNumberWithStatusVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelephoneNumberWithStatusVOWithDefaults

`func NewTelephoneNumberWithStatusVOWithDefaults() *TelephoneNumberWithStatusVO`

NewTelephoneNumberWithStatusVOWithDefaults instantiates a new TelephoneNumberWithStatusVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMac

`func (o *TelephoneNumberWithStatusVO) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *TelephoneNumberWithStatusVO) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *TelephoneNumberWithStatusVO) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *TelephoneNumberWithStatusVO) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetFeatureDescription

`func (o *TelephoneNumberWithStatusVO) GetFeatureDescription() []FeatureInfoVO`

GetFeatureDescription returns the FeatureDescription field if non-nil, zero value otherwise.

### GetFeatureDescriptionOk

`func (o *TelephoneNumberWithStatusVO) GetFeatureDescriptionOk() (*[]FeatureInfoVO, bool)`

GetFeatureDescriptionOk returns a tuple with the FeatureDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureDescription

`func (o *TelephoneNumberWithStatusVO) SetFeatureDescription(v []FeatureInfoVO)`

SetFeatureDescription sets FeatureDescription field to given value.

### HasFeatureDescription

`func (o *TelephoneNumberWithStatusVO) HasFeatureDescription() bool`

HasFeatureDescription returns a boolean if a field has been set.

### GetNumber

`func (o *TelephoneNumberWithStatusVO) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *TelephoneNumberWithStatusVO) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *TelephoneNumberWithStatusVO) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *TelephoneNumberWithStatusVO) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPassword

`func (o *TelephoneNumberWithStatusVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TelephoneNumberWithStatusVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TelephoneNumberWithStatusVO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TelephoneNumberWithStatusVO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPhoneNumberId

`func (o *TelephoneNumberWithStatusVO) GetPhoneNumberId() string`

GetPhoneNumberId returns the PhoneNumberId field if non-nil, zero value otherwise.

### GetPhoneNumberIdOk

`func (o *TelephoneNumberWithStatusVO) GetPhoneNumberIdOk() (*string, bool)`

GetPhoneNumberIdOk returns a tuple with the PhoneNumberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberId

`func (o *TelephoneNumberWithStatusVO) SetPhoneNumberId(v string)`

SetPhoneNumberId sets PhoneNumberId field to given value.

### HasPhoneNumberId

`func (o *TelephoneNumberWithStatusVO) HasPhoneNumberId() bool`

HasPhoneNumberId returns a boolean if a field has been set.

### GetPort

`func (o *TelephoneNumberWithStatusVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TelephoneNumberWithStatusVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TelephoneNumberWithStatusVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *TelephoneNumberWithStatusVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrefix

`func (o *TelephoneNumberWithStatusVO) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *TelephoneNumberWithStatusVO) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *TelephoneNumberWithStatusVO) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *TelephoneNumberWithStatusVO) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetProfileId

`func (o *TelephoneNumberWithStatusVO) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *TelephoneNumberWithStatusVO) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *TelephoneNumberWithStatusVO) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *TelephoneNumberWithStatusVO) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileName

`func (o *TelephoneNumberWithStatusVO) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *TelephoneNumberWithStatusVO) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *TelephoneNumberWithStatusVO) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *TelephoneNumberWithStatusVO) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetStatus

`func (o *TelephoneNumberWithStatusVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TelephoneNumberWithStatusVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TelephoneNumberWithStatusVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TelephoneNumberWithStatusVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsername

`func (o *TelephoneNumberWithStatusVO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TelephoneNumberWithStatusVO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TelephoneNumberWithStatusVO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TelephoneNumberWithStatusVO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


