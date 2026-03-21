# DeleteUserVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceDelete** | Pointer to **bool** | Force delete target user. If false, target user can not be deleted, when target user has child users. If true, target user will be deleted anyway. Target user&#39;s child users will be root&#39;s child user. | [optional] 

## Methods

### NewDeleteUserVO

`func NewDeleteUserVO() *DeleteUserVO`

NewDeleteUserVO instantiates a new DeleteUserVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteUserVOWithDefaults

`func NewDeleteUserVOWithDefaults() *DeleteUserVO`

NewDeleteUserVOWithDefaults instantiates a new DeleteUserVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceDelete

`func (o *DeleteUserVO) GetForceDelete() bool`

GetForceDelete returns the ForceDelete field if non-nil, zero value otherwise.

### GetForceDeleteOk

`func (o *DeleteUserVO) GetForceDeleteOk() (*bool, bool)`

GetForceDeleteOk returns a tuple with the ForceDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDelete

`func (o *DeleteUserVO) SetForceDelete(v bool)`

SetForceDelete sets ForceDelete field to given value.

### HasForceDelete

`func (o *DeleteUserVO) HasForceDelete() bool`

HasForceDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


