# DeleteDigitMapProfileEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** | Profile ID. Default digit map profile cannot be deleted. | 
**SkipConfirm** | **bool** | skipConfirm indicates whether to skip the query of the devices bound to digit map profile. false: Not to skip the query. true: Skip the query and delete the digit map profile corresponding to the profileId. | 

## Methods

### NewDeleteDigitMapProfileEntity

`func NewDeleteDigitMapProfileEntity(profileId string, skipConfirm bool, ) *DeleteDigitMapProfileEntity`

NewDeleteDigitMapProfileEntity instantiates a new DeleteDigitMapProfileEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDigitMapProfileEntityWithDefaults

`func NewDeleteDigitMapProfileEntityWithDefaults() *DeleteDigitMapProfileEntity`

NewDeleteDigitMapProfileEntityWithDefaults instantiates a new DeleteDigitMapProfileEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *DeleteDigitMapProfileEntity) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *DeleteDigitMapProfileEntity) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *DeleteDigitMapProfileEntity) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetSkipConfirm

`func (o *DeleteDigitMapProfileEntity) GetSkipConfirm() bool`

GetSkipConfirm returns the SkipConfirm field if non-nil, zero value otherwise.

### GetSkipConfirmOk

`func (o *DeleteDigitMapProfileEntity) GetSkipConfirmOk() (*bool, bool)`

GetSkipConfirmOk returns a tuple with the SkipConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipConfirm

`func (o *DeleteDigitMapProfileEntity) SetSkipConfirm(v bool)`

SetSkipConfirm sets SkipConfirm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


