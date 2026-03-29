# LdapProfileOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseDn** | Pointer to **string** | LDAP server base distinguish name | [optional] 
**BindType** | Pointer to **int32** | Type of LDAP bind, 0: Simple Mode；1: Anonymous Mode; 2: Regular Mode | [optional] 
**Cn** | Pointer to **string** | LDAP server common name | [optional] 
**DstPort** | Pointer to **int32** | LDAP server listening port. When SSL is not enabled, it is generally 389, and when SSL is enabled, it is generally 636 | [optional] 
**Filter** | Pointer to **string** | Additional filter, optional when parameter [type] is 1 or 2 | [optional] 
**GroupDn** | Pointer to **string** | LDAP server group distinguish name, optional when parameter [type] is 1 or 2 | [optional] 
**LdapProfileId** | Pointer to **string** | LDAP profile ID | [optional] 
**Name** | Pointer to **string** | LDAP profile name | [optional] 
**RegularDn** | Pointer to **string** | Regular Dn, valid when parameter [type] is 2 | [optional] 
**RegularPassword** | Pointer to **string** | Regular Password, valid when parameter [type] is 2 | [optional] 
**Server** | Pointer to **string** | LDAP server address | [optional] 
**Status** | Pointer to **bool** | LDAP profile enable status | [optional] 
**UseSsl** | Pointer to **bool** | LDAP server enable ssl status | [optional] 

## Methods

### NewLdapProfileOpenApiVO

`func NewLdapProfileOpenApiVO() *LdapProfileOpenApiVO`

NewLdapProfileOpenApiVO instantiates a new LdapProfileOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapProfileOpenApiVOWithDefaults

`func NewLdapProfileOpenApiVOWithDefaults() *LdapProfileOpenApiVO`

NewLdapProfileOpenApiVOWithDefaults instantiates a new LdapProfileOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseDn

`func (o *LdapProfileOpenApiVO) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *LdapProfileOpenApiVO) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *LdapProfileOpenApiVO) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *LdapProfileOpenApiVO) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetBindType

`func (o *LdapProfileOpenApiVO) GetBindType() int32`

GetBindType returns the BindType field if non-nil, zero value otherwise.

### GetBindTypeOk

`func (o *LdapProfileOpenApiVO) GetBindTypeOk() (*int32, bool)`

GetBindTypeOk returns a tuple with the BindType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindType

`func (o *LdapProfileOpenApiVO) SetBindType(v int32)`

SetBindType sets BindType field to given value.

### HasBindType

`func (o *LdapProfileOpenApiVO) HasBindType() bool`

HasBindType returns a boolean if a field has been set.

### GetCn

`func (o *LdapProfileOpenApiVO) GetCn() string`

GetCn returns the Cn field if non-nil, zero value otherwise.

### GetCnOk

`func (o *LdapProfileOpenApiVO) GetCnOk() (*string, bool)`

GetCnOk returns a tuple with the Cn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCn

`func (o *LdapProfileOpenApiVO) SetCn(v string)`

SetCn sets Cn field to given value.

### HasCn

`func (o *LdapProfileOpenApiVO) HasCn() bool`

HasCn returns a boolean if a field has been set.

### GetDstPort

`func (o *LdapProfileOpenApiVO) GetDstPort() int32`

GetDstPort returns the DstPort field if non-nil, zero value otherwise.

### GetDstPortOk

`func (o *LdapProfileOpenApiVO) GetDstPortOk() (*int32, bool)`

GetDstPortOk returns a tuple with the DstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPort

`func (o *LdapProfileOpenApiVO) SetDstPort(v int32)`

SetDstPort sets DstPort field to given value.

### HasDstPort

`func (o *LdapProfileOpenApiVO) HasDstPort() bool`

HasDstPort returns a boolean if a field has been set.

### GetFilter

`func (o *LdapProfileOpenApiVO) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LdapProfileOpenApiVO) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LdapProfileOpenApiVO) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *LdapProfileOpenApiVO) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetGroupDn

`func (o *LdapProfileOpenApiVO) GetGroupDn() string`

GetGroupDn returns the GroupDn field if non-nil, zero value otherwise.

### GetGroupDnOk

`func (o *LdapProfileOpenApiVO) GetGroupDnOk() (*string, bool)`

GetGroupDnOk returns a tuple with the GroupDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDn

`func (o *LdapProfileOpenApiVO) SetGroupDn(v string)`

SetGroupDn sets GroupDn field to given value.

### HasGroupDn

`func (o *LdapProfileOpenApiVO) HasGroupDn() bool`

HasGroupDn returns a boolean if a field has been set.

### GetLdapProfileId

`func (o *LdapProfileOpenApiVO) GetLdapProfileId() string`

GetLdapProfileId returns the LdapProfileId field if non-nil, zero value otherwise.

### GetLdapProfileIdOk

`func (o *LdapProfileOpenApiVO) GetLdapProfileIdOk() (*string, bool)`

GetLdapProfileIdOk returns a tuple with the LdapProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProfileId

`func (o *LdapProfileOpenApiVO) SetLdapProfileId(v string)`

SetLdapProfileId sets LdapProfileId field to given value.

### HasLdapProfileId

`func (o *LdapProfileOpenApiVO) HasLdapProfileId() bool`

HasLdapProfileId returns a boolean if a field has been set.

### GetName

`func (o *LdapProfileOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LdapProfileOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LdapProfileOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LdapProfileOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegularDn

`func (o *LdapProfileOpenApiVO) GetRegularDn() string`

GetRegularDn returns the RegularDn field if non-nil, zero value otherwise.

### GetRegularDnOk

`func (o *LdapProfileOpenApiVO) GetRegularDnOk() (*string, bool)`

GetRegularDnOk returns a tuple with the RegularDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularDn

`func (o *LdapProfileOpenApiVO) SetRegularDn(v string)`

SetRegularDn sets RegularDn field to given value.

### HasRegularDn

`func (o *LdapProfileOpenApiVO) HasRegularDn() bool`

HasRegularDn returns a boolean if a field has been set.

### GetRegularPassword

`func (o *LdapProfileOpenApiVO) GetRegularPassword() string`

GetRegularPassword returns the RegularPassword field if non-nil, zero value otherwise.

### GetRegularPasswordOk

`func (o *LdapProfileOpenApiVO) GetRegularPasswordOk() (*string, bool)`

GetRegularPasswordOk returns a tuple with the RegularPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularPassword

`func (o *LdapProfileOpenApiVO) SetRegularPassword(v string)`

SetRegularPassword sets RegularPassword field to given value.

### HasRegularPassword

`func (o *LdapProfileOpenApiVO) HasRegularPassword() bool`

HasRegularPassword returns a boolean if a field has been set.

### GetServer

`func (o *LdapProfileOpenApiVO) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *LdapProfileOpenApiVO) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *LdapProfileOpenApiVO) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *LdapProfileOpenApiVO) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetStatus

`func (o *LdapProfileOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LdapProfileOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LdapProfileOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LdapProfileOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUseSsl

`func (o *LdapProfileOpenApiVO) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *LdapProfileOpenApiVO) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *LdapProfileOpenApiVO) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *LdapProfileOpenApiVO) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


