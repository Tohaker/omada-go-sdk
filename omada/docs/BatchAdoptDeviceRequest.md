# BatchAdoptDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Macs** | Pointer to **[]string** | MAC list of devices. E.g. AA-BB-CC-DD-11-22 | [optional] 
**Password** | Pointer to **string** | Adopt device password. It should contain 1 to 64 characters. | [optional] 
**Username** | Pointer to **string** | Adopt device username. It should contain 1 to 64 characters. | [optional] 

## Methods

### NewBatchAdoptDeviceRequest

`func NewBatchAdoptDeviceRequest() *BatchAdoptDeviceRequest`

NewBatchAdoptDeviceRequest instantiates a new BatchAdoptDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchAdoptDeviceRequestWithDefaults

`func NewBatchAdoptDeviceRequestWithDefaults() *BatchAdoptDeviceRequest`

NewBatchAdoptDeviceRequestWithDefaults instantiates a new BatchAdoptDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacs

`func (o *BatchAdoptDeviceRequest) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *BatchAdoptDeviceRequest) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *BatchAdoptDeviceRequest) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *BatchAdoptDeviceRequest) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### GetPassword

`func (o *BatchAdoptDeviceRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BatchAdoptDeviceRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BatchAdoptDeviceRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BatchAdoptDeviceRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *BatchAdoptDeviceRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BatchAdoptDeviceRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BatchAdoptDeviceRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *BatchAdoptDeviceRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


