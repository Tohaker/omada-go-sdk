# AuthInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to **int32** |  | [optional] 
**Info** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthInfoVO

`func NewAuthInfoVO() *AuthInfoVO`

NewAuthInfoVO instantiates a new AuthInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthInfoVOWithDefaults

`func NewAuthInfoVOWithDefaults() *AuthInfoVO`

NewAuthInfoVOWithDefaults instantiates a new AuthInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *AuthInfoVO) GetAuthType() int32`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *AuthInfoVO) GetAuthTypeOk() (*int32, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *AuthInfoVO) SetAuthType(v int32)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *AuthInfoVO) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetInfo

`func (o *AuthInfoVO) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AuthInfoVO) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AuthInfoVO) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AuthInfoVO) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


