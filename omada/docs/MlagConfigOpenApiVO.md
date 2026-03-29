# MlagConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | MLAG group id should be between 1 and 31 characters. Only letters, numbers, and the following symbols are allowed: - . / : @ _ + # | [optional] 
**MembersConfig** | Pointer to [**[]MlagMemberConfigVO**](MlagMemberConfigVO.md) | M-LAG group members configuration | [optional] 
**Name** | **string** | MLAG group name should be between 1 and 31 characters. Only letters, numbers, and the following symbols are allowed: - . / : @ _ + # | 

## Methods

### NewMlagConfigOpenApiVO

`func NewMlagConfigOpenApiVO(name string, ) *MlagConfigOpenApiVO`

NewMlagConfigOpenApiVO instantiates a new MlagConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlagConfigOpenApiVOWithDefaults

`func NewMlagConfigOpenApiVOWithDefaults() *MlagConfigOpenApiVO`

NewMlagConfigOpenApiVOWithDefaults instantiates a new MlagConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *MlagConfigOpenApiVO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *MlagConfigOpenApiVO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *MlagConfigOpenApiVO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *MlagConfigOpenApiVO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetMembersConfig

`func (o *MlagConfigOpenApiVO) GetMembersConfig() []MlagMemberConfigVO`

GetMembersConfig returns the MembersConfig field if non-nil, zero value otherwise.

### GetMembersConfigOk

`func (o *MlagConfigOpenApiVO) GetMembersConfigOk() (*[]MlagMemberConfigVO, bool)`

GetMembersConfigOk returns a tuple with the MembersConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersConfig

`func (o *MlagConfigOpenApiVO) SetMembersConfig(v []MlagMemberConfigVO)`

SetMembersConfig sets MembersConfig field to given value.

### HasMembersConfig

`func (o *MlagConfigOpenApiVO) HasMembersConfig() bool`

HasMembersConfig returns a boolean if a field has been set.

### GetName

`func (o *MlagConfigOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MlagConfigOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MlagConfigOpenApiVO) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


