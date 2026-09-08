# DeviceReplaceSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationDeviceKey** | **string** | Device Key of the device. | 
**Password** | Pointer to **string** | password. | [optional] 
**Username** | Pointer to **string** | User Name. | [optional] 

## Methods

### NewDeviceReplaceSettingVO

`func NewDeviceReplaceSettingVO(destinationDeviceKey string, ) *DeviceReplaceSettingVO`

NewDeviceReplaceSettingVO instantiates a new DeviceReplaceSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceReplaceSettingVOWithDefaults

`func NewDeviceReplaceSettingVOWithDefaults() *DeviceReplaceSettingVO`

NewDeviceReplaceSettingVOWithDefaults instantiates a new DeviceReplaceSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationDeviceKey

`func (o *DeviceReplaceSettingVO) GetDestinationDeviceKey() string`

GetDestinationDeviceKey returns the DestinationDeviceKey field if non-nil, zero value otherwise.

### GetDestinationDeviceKeyOk

`func (o *DeviceReplaceSettingVO) GetDestinationDeviceKeyOk() (*string, bool)`

GetDestinationDeviceKeyOk returns a tuple with the DestinationDeviceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDeviceKey

`func (o *DeviceReplaceSettingVO) SetDestinationDeviceKey(v string)`

SetDestinationDeviceKey sets DestinationDeviceKey field to given value.


### GetPassword

`func (o *DeviceReplaceSettingVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DeviceReplaceSettingVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DeviceReplaceSettingVO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DeviceReplaceSettingVO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *DeviceReplaceSettingVO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DeviceReplaceSettingVO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DeviceReplaceSettingVO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DeviceReplaceSettingVO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


