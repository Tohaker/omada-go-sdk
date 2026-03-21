# UserBriefVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | User name. | [optional] 
**TemporaryEnable** | Pointer to **bool** | Temporary enable. | [optional] 
**UserId** | Pointer to **string** | User ID. | [optional] 

## Methods

### NewUserBriefVO

`func NewUserBriefVO() *UserBriefVO`

NewUserBriefVO instantiates a new UserBriefVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBriefVOWithDefaults

`func NewUserBriefVOWithDefaults() *UserBriefVO`

NewUserBriefVOWithDefaults instantiates a new UserBriefVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserBriefVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserBriefVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserBriefVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserBriefVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemporaryEnable

`func (o *UserBriefVO) GetTemporaryEnable() bool`

GetTemporaryEnable returns the TemporaryEnable field if non-nil, zero value otherwise.

### GetTemporaryEnableOk

`func (o *UserBriefVO) GetTemporaryEnableOk() (*bool, bool)`

GetTemporaryEnableOk returns a tuple with the TemporaryEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryEnable

`func (o *UserBriefVO) SetTemporaryEnable(v bool)`

SetTemporaryEnable sets TemporaryEnable field to given value.

### HasTemporaryEnable

`func (o *UserBriefVO) HasTemporaryEnable() bool`

HasTemporaryEnable returns a boolean if a field has been set.

### GetUserId

`func (o *UserBriefVO) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserBriefVO) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserBriefVO) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserBriefVO) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


