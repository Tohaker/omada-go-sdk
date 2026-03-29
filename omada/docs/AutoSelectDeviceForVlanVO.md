# AutoSelectDeviceForVlanVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmType** | Pointer to **int32** | It should be a value as follows: 1: need confirm when deleting gateway | [optional] 
**Deletable** | Pointer to **bool** | It indicates whether the device is deletable. | [optional] 
**Mac** | Pointer to **string** | Device Mac | [optional] 
**StackId** | Pointer to **string** | Stack Id, only valid when the device is stack. | [optional] 

## Methods

### NewAutoSelectDeviceForVlanVO

`func NewAutoSelectDeviceForVlanVO() *AutoSelectDeviceForVlanVO`

NewAutoSelectDeviceForVlanVO instantiates a new AutoSelectDeviceForVlanVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoSelectDeviceForVlanVOWithDefaults

`func NewAutoSelectDeviceForVlanVOWithDefaults() *AutoSelectDeviceForVlanVO`

NewAutoSelectDeviceForVlanVOWithDefaults instantiates a new AutoSelectDeviceForVlanVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmType

`func (o *AutoSelectDeviceForVlanVO) GetConfirmType() int32`

GetConfirmType returns the ConfirmType field if non-nil, zero value otherwise.

### GetConfirmTypeOk

`func (o *AutoSelectDeviceForVlanVO) GetConfirmTypeOk() (*int32, bool)`

GetConfirmTypeOk returns a tuple with the ConfirmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmType

`func (o *AutoSelectDeviceForVlanVO) SetConfirmType(v int32)`

SetConfirmType sets ConfirmType field to given value.

### HasConfirmType

`func (o *AutoSelectDeviceForVlanVO) HasConfirmType() bool`

HasConfirmType returns a boolean if a field has been set.

### GetDeletable

`func (o *AutoSelectDeviceForVlanVO) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *AutoSelectDeviceForVlanVO) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *AutoSelectDeviceForVlanVO) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *AutoSelectDeviceForVlanVO) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetMac

`func (o *AutoSelectDeviceForVlanVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *AutoSelectDeviceForVlanVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *AutoSelectDeviceForVlanVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *AutoSelectDeviceForVlanVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetStackId

`func (o *AutoSelectDeviceForVlanVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *AutoSelectDeviceForVlanVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *AutoSelectDeviceForVlanVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *AutoSelectDeviceForVlanVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


