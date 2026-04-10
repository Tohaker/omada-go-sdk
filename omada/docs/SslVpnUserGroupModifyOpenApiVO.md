# SslVpnUserGroupModifyOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdapAttribute** | Pointer to **string** | Attribute value used for LDAP authentication. It should contain 1 to 20 characters. | [optional] 
**RadiusAttribute** | Pointer to **string** | Attribute value used for radius authentication. It should contain 1 to 20 characters. | [optional] 
**ResourceGroupList** | Pointer to **[]string** | Resource group ID list of the SSL VPN user group. Resource group can be created using &#39;Create SSL VPN resource group&#39; interface, and Resource group ID can be obtained from &#39;Get resource group list for SSL VPN server&#39; interface. | [optional] 

## Methods

### NewSslVpnUserGroupModifyOpenApiVO

`func NewSslVpnUserGroupModifyOpenApiVO() *SslVpnUserGroupModifyOpenApiVO`

NewSslVpnUserGroupModifyOpenApiVO instantiates a new SslVpnUserGroupModifyOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslVpnUserGroupModifyOpenApiVOWithDefaults

`func NewSslVpnUserGroupModifyOpenApiVOWithDefaults() *SslVpnUserGroupModifyOpenApiVO`

NewSslVpnUserGroupModifyOpenApiVOWithDefaults instantiates a new SslVpnUserGroupModifyOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdapAttribute

`func (o *SslVpnUserGroupModifyOpenApiVO) GetLdapAttribute() string`

GetLdapAttribute returns the LdapAttribute field if non-nil, zero value otherwise.

### GetLdapAttributeOk

`func (o *SslVpnUserGroupModifyOpenApiVO) GetLdapAttributeOk() (*string, bool)`

GetLdapAttributeOk returns a tuple with the LdapAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAttribute

`func (o *SslVpnUserGroupModifyOpenApiVO) SetLdapAttribute(v string)`

SetLdapAttribute sets LdapAttribute field to given value.

### HasLdapAttribute

`func (o *SslVpnUserGroupModifyOpenApiVO) HasLdapAttribute() bool`

HasLdapAttribute returns a boolean if a field has been set.

### GetRadiusAttribute

`func (o *SslVpnUserGroupModifyOpenApiVO) GetRadiusAttribute() string`

GetRadiusAttribute returns the RadiusAttribute field if non-nil, zero value otherwise.

### GetRadiusAttributeOk

`func (o *SslVpnUserGroupModifyOpenApiVO) GetRadiusAttributeOk() (*string, bool)`

GetRadiusAttributeOk returns a tuple with the RadiusAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAttribute

`func (o *SslVpnUserGroupModifyOpenApiVO) SetRadiusAttribute(v string)`

SetRadiusAttribute sets RadiusAttribute field to given value.

### HasRadiusAttribute

`func (o *SslVpnUserGroupModifyOpenApiVO) HasRadiusAttribute() bool`

HasRadiusAttribute returns a boolean if a field has been set.

### GetResourceGroupList

`func (o *SslVpnUserGroupModifyOpenApiVO) GetResourceGroupList() []string`

GetResourceGroupList returns the ResourceGroupList field if non-nil, zero value otherwise.

### GetResourceGroupListOk

`func (o *SslVpnUserGroupModifyOpenApiVO) GetResourceGroupListOk() (*[]string, bool)`

GetResourceGroupListOk returns a tuple with the ResourceGroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupList

`func (o *SslVpnUserGroupModifyOpenApiVO) SetResourceGroupList(v []string)`

SetResourceGroupList sets ResourceGroupList field to given value.

### HasResourceGroupList

`func (o *SslVpnUserGroupModifyOpenApiVO) HasResourceGroupList() bool`

HasResourceGroupList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


