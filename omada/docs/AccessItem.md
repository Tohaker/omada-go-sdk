# AccessItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallingCode** | Pointer to **string** | Calling code should contain 2 to 5 characters. Calling code must be entered when entering the country code. For the values of Calling code, refer to section 5.4.1 of the Open API Access Guide. | [optional] 
**CountryCode** | Pointer to **string** | Country code should contain 2 characters. Country code must be entered when entering the calling code. For the values of Country code, refer to section 5.4.1 of the Open API Access Guide. | [optional] 
**Phone** | **string** | Allowed phone numbers. | 

## Methods

### NewAccessItem

`func NewAccessItem(phone string, ) *AccessItem`

NewAccessItem instantiates a new AccessItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessItemWithDefaults

`func NewAccessItemWithDefaults() *AccessItem`

NewAccessItemWithDefaults instantiates a new AccessItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallingCode

`func (o *AccessItem) GetCallingCode() string`

GetCallingCode returns the CallingCode field if non-nil, zero value otherwise.

### GetCallingCodeOk

`func (o *AccessItem) GetCallingCodeOk() (*string, bool)`

GetCallingCodeOk returns a tuple with the CallingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallingCode

`func (o *AccessItem) SetCallingCode(v string)`

SetCallingCode sets CallingCode field to given value.

### HasCallingCode

`func (o *AccessItem) HasCallingCode() bool`

HasCallingCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *AccessItem) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AccessItem) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AccessItem) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *AccessItem) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhone

`func (o *AccessItem) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AccessItem) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AccessItem) SetPhone(v string)`

SetPhone sets Phone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


