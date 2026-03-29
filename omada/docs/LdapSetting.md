# LdapSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalUrl** | Pointer to **string** | External URL. Required when portalCustom is 2. | [optional] 
**ExternalUrlScheme** | Pointer to **string** | External URL scheme, should be a value as follows: http ; https. Required when portalCustom is 2. | [optional] 
**LdapProfileId** | **string** | LDAP profile ID. Ldap profile can be  created using &#39;Create a new LDAP profile&#39; (&#39;Create a new LDAP profile template&#39;) interface, and LDAP profile ID can be obtained from &#39;Get LDAP profile list&#39; (&#39;Get LDAP profile template list&#39;) interface | 
**PortalCustom** | **int32** | Portal customization, should be a value as follows: 1: use local; 2: use external | 

## Methods

### NewLdapSetting

`func NewLdapSetting(ldapProfileId string, portalCustom int32, ) *LdapSetting`

NewLdapSetting instantiates a new LdapSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapSettingWithDefaults

`func NewLdapSettingWithDefaults() *LdapSetting`

NewLdapSettingWithDefaults instantiates a new LdapSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalUrl

`func (o *LdapSetting) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *LdapSetting) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *LdapSetting) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.

### HasExternalUrl

`func (o *LdapSetting) HasExternalUrl() bool`

HasExternalUrl returns a boolean if a field has been set.

### GetExternalUrlScheme

`func (o *LdapSetting) GetExternalUrlScheme() string`

GetExternalUrlScheme returns the ExternalUrlScheme field if non-nil, zero value otherwise.

### GetExternalUrlSchemeOk

`func (o *LdapSetting) GetExternalUrlSchemeOk() (*string, bool)`

GetExternalUrlSchemeOk returns a tuple with the ExternalUrlScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrlScheme

`func (o *LdapSetting) SetExternalUrlScheme(v string)`

SetExternalUrlScheme sets ExternalUrlScheme field to given value.

### HasExternalUrlScheme

`func (o *LdapSetting) HasExternalUrlScheme() bool`

HasExternalUrlScheme returns a boolean if a field has been set.

### GetLdapProfileId

`func (o *LdapSetting) GetLdapProfileId() string`

GetLdapProfileId returns the LdapProfileId field if non-nil, zero value otherwise.

### GetLdapProfileIdOk

`func (o *LdapSetting) GetLdapProfileIdOk() (*string, bool)`

GetLdapProfileIdOk returns a tuple with the LdapProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProfileId

`func (o *LdapSetting) SetLdapProfileId(v string)`

SetLdapProfileId sets LdapProfileId field to given value.


### GetPortalCustom

`func (o *LdapSetting) GetPortalCustom() int32`

GetPortalCustom returns the PortalCustom field if non-nil, zero value otherwise.

### GetPortalCustomOk

`func (o *LdapSetting) GetPortalCustomOk() (*int32, bool)`

GetPortalCustomOk returns a tuple with the PortalCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalCustom

`func (o *LdapSetting) SetPortalCustom(v int32)`

SetPortalCustom sets PortalCustom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


