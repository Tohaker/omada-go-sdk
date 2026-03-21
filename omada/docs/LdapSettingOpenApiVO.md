# LdapSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LDAPProfile** | **string** | LDAP profile ID. LDAP profile can be created using &#39;Create a new LDAP profile&#39; interface, and LDAP profile ID can be obtained from &#39;Get LDAP profile list&#39; interface. | 
**DefaultGroup** | **string** | Default user group ID for the user on the LDAP server. User group can be created using &#39;Create SSL VPN user group&#39; interface, and user group ID can be obtained from &#39;Get user group list for SSL VPN server&#39; interface. | 

## Methods

### NewLdapSettingOpenApiVO

`func NewLdapSettingOpenApiVO(lDAPProfile string, defaultGroup string, ) *LdapSettingOpenApiVO`

NewLdapSettingOpenApiVO instantiates a new LdapSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapSettingOpenApiVOWithDefaults

`func NewLdapSettingOpenApiVOWithDefaults() *LdapSettingOpenApiVO`

NewLdapSettingOpenApiVOWithDefaults instantiates a new LdapSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLDAPProfile

`func (o *LdapSettingOpenApiVO) GetLDAPProfile() string`

GetLDAPProfile returns the LDAPProfile field if non-nil, zero value otherwise.

### GetLDAPProfileOk

`func (o *LdapSettingOpenApiVO) GetLDAPProfileOk() (*string, bool)`

GetLDAPProfileOk returns a tuple with the LDAPProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLDAPProfile

`func (o *LdapSettingOpenApiVO) SetLDAPProfile(v string)`

SetLDAPProfile sets LDAPProfile field to given value.


### GetDefaultGroup

`func (o *LdapSettingOpenApiVO) GetDefaultGroup() string`

GetDefaultGroup returns the DefaultGroup field if non-nil, zero value otherwise.

### GetDefaultGroupOk

`func (o *LdapSettingOpenApiVO) GetDefaultGroupOk() (*string, bool)`

GetDefaultGroupOk returns a tuple with the DefaultGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroup

`func (o *LdapSettingOpenApiVO) SetDefaultGroup(v string)`

SetDefaultGroup sets DefaultGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


