# SslVpnUserGroupEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LDAPAttribute** | Pointer to **string** | Attribute value used for LDAP authentication | [optional] 
**Id** | Pointer to **string** | ID of the SSL VPN user group | [optional] 
**Name** | **string** | Name of the SSL VPN user group | 
**RadiusAttribute** | Pointer to **string** | Attribute value used for radius authentication | [optional] 
**ResourceGroupList** | Pointer to [**[]SslVpnResourceGroupBriefInfo**](SslVpnResourceGroupBriefInfo.md) | Resource group list of the SSL VPN user group | [optional] 
**UserList** | Pointer to **[]string** | User list of the SSL VPN user group | [optional] 
**UserNumber** | Pointer to **int32** | User number of the SSL VPN user group | [optional] 

## Methods

### NewSslVpnUserGroupEntity

`func NewSslVpnUserGroupEntity(name string, ) *SslVpnUserGroupEntity`

NewSslVpnUserGroupEntity instantiates a new SslVpnUserGroupEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslVpnUserGroupEntityWithDefaults

`func NewSslVpnUserGroupEntityWithDefaults() *SslVpnUserGroupEntity`

NewSslVpnUserGroupEntityWithDefaults instantiates a new SslVpnUserGroupEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLDAPAttribute

`func (o *SslVpnUserGroupEntity) GetLDAPAttribute() string`

GetLDAPAttribute returns the LDAPAttribute field if non-nil, zero value otherwise.

### GetLDAPAttributeOk

`func (o *SslVpnUserGroupEntity) GetLDAPAttributeOk() (*string, bool)`

GetLDAPAttributeOk returns a tuple with the LDAPAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLDAPAttribute

`func (o *SslVpnUserGroupEntity) SetLDAPAttribute(v string)`

SetLDAPAttribute sets LDAPAttribute field to given value.

### HasLDAPAttribute

`func (o *SslVpnUserGroupEntity) HasLDAPAttribute() bool`

HasLDAPAttribute returns a boolean if a field has been set.

### GetId

`func (o *SslVpnUserGroupEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SslVpnUserGroupEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SslVpnUserGroupEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SslVpnUserGroupEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SslVpnUserGroupEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SslVpnUserGroupEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SslVpnUserGroupEntity) SetName(v string)`

SetName sets Name field to given value.


### GetRadiusAttribute

`func (o *SslVpnUserGroupEntity) GetRadiusAttribute() string`

GetRadiusAttribute returns the RadiusAttribute field if non-nil, zero value otherwise.

### GetRadiusAttributeOk

`func (o *SslVpnUserGroupEntity) GetRadiusAttributeOk() (*string, bool)`

GetRadiusAttributeOk returns a tuple with the RadiusAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAttribute

`func (o *SslVpnUserGroupEntity) SetRadiusAttribute(v string)`

SetRadiusAttribute sets RadiusAttribute field to given value.

### HasRadiusAttribute

`func (o *SslVpnUserGroupEntity) HasRadiusAttribute() bool`

HasRadiusAttribute returns a boolean if a field has been set.

### GetResourceGroupList

`func (o *SslVpnUserGroupEntity) GetResourceGroupList() []SslVpnResourceGroupBriefInfo`

GetResourceGroupList returns the ResourceGroupList field if non-nil, zero value otherwise.

### GetResourceGroupListOk

`func (o *SslVpnUserGroupEntity) GetResourceGroupListOk() (*[]SslVpnResourceGroupBriefInfo, bool)`

GetResourceGroupListOk returns a tuple with the ResourceGroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupList

`func (o *SslVpnUserGroupEntity) SetResourceGroupList(v []SslVpnResourceGroupBriefInfo)`

SetResourceGroupList sets ResourceGroupList field to given value.

### HasResourceGroupList

`func (o *SslVpnUserGroupEntity) HasResourceGroupList() bool`

HasResourceGroupList returns a boolean if a field has been set.

### GetUserList

`func (o *SslVpnUserGroupEntity) GetUserList() []string`

GetUserList returns the UserList field if non-nil, zero value otherwise.

### GetUserListOk

`func (o *SslVpnUserGroupEntity) GetUserListOk() (*[]string, bool)`

GetUserListOk returns a tuple with the UserList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserList

`func (o *SslVpnUserGroupEntity) SetUserList(v []string)`

SetUserList sets UserList field to given value.

### HasUserList

`func (o *SslVpnUserGroupEntity) HasUserList() bool`

HasUserList returns a boolean if a field has been set.

### GetUserNumber

`func (o *SslVpnUserGroupEntity) GetUserNumber() int32`

GetUserNumber returns the UserNumber field if non-nil, zero value otherwise.

### GetUserNumberOk

`func (o *SslVpnUserGroupEntity) GetUserNumberOk() (*int32, bool)`

GetUserNumberOk returns a tuple with the UserNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNumber

`func (o *SslVpnUserGroupEntity) SetUserNumber(v int32)`

SetUserNumber sets UserNumber field to given value.

### HasUserNumber

`func (o *SslVpnUserGroupEntity) HasUserNumber() bool`

HasUserNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


