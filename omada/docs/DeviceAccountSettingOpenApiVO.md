# DeviceAccountSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | Device account parameter [password] should contain 10 to 64 ASCII characters. And passwords must be a combination of uppercase letters, lowercase letters, numbers, and special symbols. Symbols such as ! # $ % &amp; * @ ^ are supported. The password should not contain consecutive identical characters. Username and Password should not be the same. | 
**Username** | **string** | Device account username should contain 1 to 64 ASCII characters. | 

## Methods

### NewDeviceAccountSettingOpenApiVO

`func NewDeviceAccountSettingOpenApiVO(password string, username string, ) *DeviceAccountSettingOpenApiVO`

NewDeviceAccountSettingOpenApiVO instantiates a new DeviceAccountSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAccountSettingOpenApiVOWithDefaults

`func NewDeviceAccountSettingOpenApiVOWithDefaults() *DeviceAccountSettingOpenApiVO`

NewDeviceAccountSettingOpenApiVOWithDefaults instantiates a new DeviceAccountSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *DeviceAccountSettingOpenApiVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DeviceAccountSettingOpenApiVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DeviceAccountSettingOpenApiVO) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *DeviceAccountSettingOpenApiVO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DeviceAccountSettingOpenApiVO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DeviceAccountSettingOpenApiVO) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


