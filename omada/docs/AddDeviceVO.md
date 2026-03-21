# AddDeviceVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceKey** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Device name should contain 1 to 128 characters. | [optional] 
**Password** | Pointer to **string** | Device password. It should contain 1 to 64 characters. | [optional] 
**Sn** | Pointer to **string** | Device serial number. It should contains 13 characters. | [optional] 
**Username** | Pointer to **string** | Device username. It should contain 1 to 64 characters. | [optional] 

## Methods

### NewAddDeviceVO

`func NewAddDeviceVO() *AddDeviceVO`

NewAddDeviceVO instantiates a new AddDeviceVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDeviceVOWithDefaults

`func NewAddDeviceVOWithDefaults() *AddDeviceVO`

NewAddDeviceVOWithDefaults instantiates a new AddDeviceVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceKey

`func (o *AddDeviceVO) GetDeviceKey() string`

GetDeviceKey returns the DeviceKey field if non-nil, zero value otherwise.

### GetDeviceKeyOk

`func (o *AddDeviceVO) GetDeviceKeyOk() (*string, bool)`

GetDeviceKeyOk returns a tuple with the DeviceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceKey

`func (o *AddDeviceVO) SetDeviceKey(v string)`

SetDeviceKey sets DeviceKey field to given value.

### HasDeviceKey

`func (o *AddDeviceVO) HasDeviceKey() bool`

HasDeviceKey returns a boolean if a field has been set.

### GetName

`func (o *AddDeviceVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddDeviceVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddDeviceVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddDeviceVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *AddDeviceVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddDeviceVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddDeviceVO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddDeviceVO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSn

`func (o *AddDeviceVO) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *AddDeviceVO) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *AddDeviceVO) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *AddDeviceVO) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetUsername

`func (o *AddDeviceVO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddDeviceVO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddDeviceVO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AddDeviceVO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


