# AddDeviceBySnOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Device name should contain 1 to 128 characters. | [optional] 
**Password** | Pointer to **string** | Device password. It should contain 1 to 128 characters. | [optional] 
**Sn** | Pointer to **string** | Device serial number. It should contains 13 characters. | [optional] 
**Username** | Pointer to **string** | Device username. It should contain 1 to 128 characters. | [optional] 

## Methods

### NewAddDeviceBySnOpenApiVO

`func NewAddDeviceBySnOpenApiVO() *AddDeviceBySnOpenApiVO`

NewAddDeviceBySnOpenApiVO instantiates a new AddDeviceBySnOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDeviceBySnOpenApiVOWithDefaults

`func NewAddDeviceBySnOpenApiVOWithDefaults() *AddDeviceBySnOpenApiVO`

NewAddDeviceBySnOpenApiVOWithDefaults instantiates a new AddDeviceBySnOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddDeviceBySnOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddDeviceBySnOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddDeviceBySnOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddDeviceBySnOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *AddDeviceBySnOpenApiVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddDeviceBySnOpenApiVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddDeviceBySnOpenApiVO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddDeviceBySnOpenApiVO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSn

`func (o *AddDeviceBySnOpenApiVO) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *AddDeviceBySnOpenApiVO) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *AddDeviceBySnOpenApiVO) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *AddDeviceBySnOpenApiVO) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetUsername

`func (o *AddDeviceBySnOpenApiVO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddDeviceBySnOpenApiVO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddDeviceBySnOpenApiVO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AddDeviceBySnOpenApiVO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


