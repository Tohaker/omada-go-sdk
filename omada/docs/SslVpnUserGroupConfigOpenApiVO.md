# SslVpnUserGroupConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LDAPAttribute** | Pointer to **string** | Attribute value used for LDAP authentication. It should contain 1 to 20 characters. | [optional] 
**Name** | **string** | Name of the SSL VPN user group should contain 1 to 20 characters. | 
**RadiusAttribute** | Pointer to **string** | Attribute value used for radius authentication. It should contain 1 to 20 characters. | [optional] 
**ResourceGroupList** | Pointer to **[]string** | Resource group ID list of the SSL VPN user group. Resource group can be created using &#39;Create SSL VPN resource group&#39; interface, and Resource group ID can be obtained from &#39;Get resource group list for SSL VPN server&#39; interface. | [optional] 

## Methods

### NewSslVpnUserGroupConfigOpenApiVO

`func NewSslVpnUserGroupConfigOpenApiVO(name string, ) *SslVpnUserGroupConfigOpenApiVO`

NewSslVpnUserGroupConfigOpenApiVO instantiates a new SslVpnUserGroupConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslVpnUserGroupConfigOpenApiVOWithDefaults

`func NewSslVpnUserGroupConfigOpenApiVOWithDefaults() *SslVpnUserGroupConfigOpenApiVO`

NewSslVpnUserGroupConfigOpenApiVOWithDefaults instantiates a new SslVpnUserGroupConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLDAPAttribute

`func (o *SslVpnUserGroupConfigOpenApiVO) GetLDAPAttribute() string`

GetLDAPAttribute returns the LDAPAttribute field if non-nil, zero value otherwise.

### GetLDAPAttributeOk

`func (o *SslVpnUserGroupConfigOpenApiVO) GetLDAPAttributeOk() (*string, bool)`

GetLDAPAttributeOk returns a tuple with the LDAPAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLDAPAttribute

`func (o *SslVpnUserGroupConfigOpenApiVO) SetLDAPAttribute(v string)`

SetLDAPAttribute sets LDAPAttribute field to given value.

### HasLDAPAttribute

`func (o *SslVpnUserGroupConfigOpenApiVO) HasLDAPAttribute() bool`

HasLDAPAttribute returns a boolean if a field has been set.

### GetName

`func (o *SslVpnUserGroupConfigOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SslVpnUserGroupConfigOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SslVpnUserGroupConfigOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetRadiusAttribute

`func (o *SslVpnUserGroupConfigOpenApiVO) GetRadiusAttribute() string`

GetRadiusAttribute returns the RadiusAttribute field if non-nil, zero value otherwise.

### GetRadiusAttributeOk

`func (o *SslVpnUserGroupConfigOpenApiVO) GetRadiusAttributeOk() (*string, bool)`

GetRadiusAttributeOk returns a tuple with the RadiusAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAttribute

`func (o *SslVpnUserGroupConfigOpenApiVO) SetRadiusAttribute(v string)`

SetRadiusAttribute sets RadiusAttribute field to given value.

### HasRadiusAttribute

`func (o *SslVpnUserGroupConfigOpenApiVO) HasRadiusAttribute() bool`

HasRadiusAttribute returns a boolean if a field has been set.

### GetResourceGroupList

`func (o *SslVpnUserGroupConfigOpenApiVO) GetResourceGroupList() []string`

GetResourceGroupList returns the ResourceGroupList field if non-nil, zero value otherwise.

### GetResourceGroupListOk

`func (o *SslVpnUserGroupConfigOpenApiVO) GetResourceGroupListOk() (*[]string, bool)`

GetResourceGroupListOk returns a tuple with the ResourceGroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupList

`func (o *SslVpnUserGroupConfigOpenApiVO) SetResourceGroupList(v []string)`

SetResourceGroupList sets ResourceGroupList field to given value.

### HasResourceGroupList

`func (o *SslVpnUserGroupConfigOpenApiVO) HasResourceGroupList() bool`

HasResourceGroupList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


