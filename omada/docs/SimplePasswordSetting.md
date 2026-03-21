# SimplePasswordSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | Auth password, should contain 1 to 128 characters. | 

## Methods

### NewSimplePasswordSetting

`func NewSimplePasswordSetting(password string, ) *SimplePasswordSetting`

NewSimplePasswordSetting instantiates a new SimplePasswordSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplePasswordSettingWithDefaults

`func NewSimplePasswordSettingWithDefaults() *SimplePasswordSetting`

NewSimplePasswordSettingWithDefaults instantiates a new SimplePasswordSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *SimplePasswordSetting) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SimplePasswordSetting) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SimplePasswordSetting) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


