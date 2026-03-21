# GetUserListResponseVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]UserBriefVO**](UserBriefVO.md) | User list | [optional] 

## Methods

### NewGetUserListResponseVO

`func NewGetUserListResponseVO() *GetUserListResponseVO`

NewGetUserListResponseVO instantiates a new GetUserListResponseVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserListResponseVOWithDefaults

`func NewGetUserListResponseVOWithDefaults() *GetUserListResponseVO`

NewGetUserListResponseVOWithDefaults instantiates a new GetUserListResponseVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *GetUserListResponseVO) GetUsers() []UserBriefVO`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetUserListResponseVO) GetUsersOk() (*[]UserBriefVO, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetUserListResponseVO) SetUsers(v []UserBriefVO)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GetUserListResponseVO) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


