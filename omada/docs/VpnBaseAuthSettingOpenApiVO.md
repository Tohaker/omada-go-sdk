# VpnBaseAuthSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultGroup** | **string** | Only for SSL VPN. Default user group ID for the user on the LDAP/Radius server. User group can be created using &#39;Create SSL VPN user group&#39; interface, and user group ID can be obtained from &#39;Get user group list for SSL VPN server&#39; interface. | 
**ProfileId** | **string** | Profile ID. LDAP profile can be created using &#39;Create a new LDAP profile&#39; interface, and LDAP profile ID can be obtained from &#39;Get LDAP profile list&#39; interface. Radius profile can be created using &#39;Create a new Radius profile&#39; interface, and Radius profile ID can be obtained from &#39;Get Radius profile list&#39; interface. | 

## Methods

### NewVpnBaseAuthSettingOpenApiVO

`func NewVpnBaseAuthSettingOpenApiVO(defaultGroup string, profileId string, ) *VpnBaseAuthSettingOpenApiVO`

NewVpnBaseAuthSettingOpenApiVO instantiates a new VpnBaseAuthSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnBaseAuthSettingOpenApiVOWithDefaults

`func NewVpnBaseAuthSettingOpenApiVOWithDefaults() *VpnBaseAuthSettingOpenApiVO`

NewVpnBaseAuthSettingOpenApiVOWithDefaults instantiates a new VpnBaseAuthSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultGroup

`func (o *VpnBaseAuthSettingOpenApiVO) GetDefaultGroup() string`

GetDefaultGroup returns the DefaultGroup field if non-nil, zero value otherwise.

### GetDefaultGroupOk

`func (o *VpnBaseAuthSettingOpenApiVO) GetDefaultGroupOk() (*string, bool)`

GetDefaultGroupOk returns a tuple with the DefaultGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroup

`func (o *VpnBaseAuthSettingOpenApiVO) SetDefaultGroup(v string)`

SetDefaultGroup sets DefaultGroup field to given value.


### GetProfileId

`func (o *VpnBaseAuthSettingOpenApiVO) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *VpnBaseAuthSettingOpenApiVO) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *VpnBaseAuthSettingOpenApiVO) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


