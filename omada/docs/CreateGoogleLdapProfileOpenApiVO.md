# CreateGoogleLdapProfileOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account, required when parameter [type] is 2. account should contain 1 to 256 characters. | [optional] 
**BaseDn** | **string** | LDAP server base distinguish name, for example: dc&#x3D;xxx,dc&#x3D;com. BaseDn should contain 1 to 512 characters. | 
**BindType** | **int32** | Type of LDAP bind. BindType should be a value as follows: 0: Simple Mode; 2: Regular Mode. | 
**Cn** | **string** | LDAP server common name, for example: cn, uid. Cn should contain 1 to 64 characters. | 
**Filter** | Pointer to **string** | Additional filter, optional when parameter [type] is 1 or 2. For example: ou&#x3D;xxx. Filter should contain 1 to 512 characters. | [optional] 
**Name** | **string** | Google LDAP profile name. Name should contain 1 to 64 characters. | 
**Password** | Pointer to **string** | Password, required when parameter [type] is 2. password should contain 6 to 256 characters. | [optional] 
**Port** | **int32** | LDAP server listening port. dstPort should be within the range of 0-65535. It is generally 636. | 
**Server** | **string** | LDAP server address. | 
**Status** | **bool** | LDAP profile enable status. | 

## Methods

### NewCreateGoogleLdapProfileOpenApiVO

`func NewCreateGoogleLdapProfileOpenApiVO(baseDn string, bindType int32, cn string, name string, port int32, server string, status bool, ) *CreateGoogleLdapProfileOpenApiVO`

NewCreateGoogleLdapProfileOpenApiVO instantiates a new CreateGoogleLdapProfileOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGoogleLdapProfileOpenApiVOWithDefaults

`func NewCreateGoogleLdapProfileOpenApiVOWithDefaults() *CreateGoogleLdapProfileOpenApiVO`

NewCreateGoogleLdapProfileOpenApiVOWithDefaults instantiates a new CreateGoogleLdapProfileOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CreateGoogleLdapProfileOpenApiVO) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreateGoogleLdapProfileOpenApiVO) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreateGoogleLdapProfileOpenApiVO) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CreateGoogleLdapProfileOpenApiVO) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetBaseDn

`func (o *CreateGoogleLdapProfileOpenApiVO) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *CreateGoogleLdapProfileOpenApiVO) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *CreateGoogleLdapProfileOpenApiVO) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.


### GetBindType

`func (o *CreateGoogleLdapProfileOpenApiVO) GetBindType() int32`

GetBindType returns the BindType field if non-nil, zero value otherwise.

### GetBindTypeOk

`func (o *CreateGoogleLdapProfileOpenApiVO) GetBindTypeOk() (*int32, bool)`

GetBindTypeOk returns a tuple with the BindType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindType

`func (o *CreateGoogleLdapProfileOpenApiVO) SetBindType(v int32)`

SetBindType sets BindType field to given value.


### GetCn

`func (o *CreateGoogleLdapProfileOpenApiVO) GetCn() string`

GetCn returns the Cn field if non-nil, zero value otherwise.

### GetCnOk

`func (o *CreateGoogleLdapProfileOpenApiVO) GetCnOk() (*string, bool)`

GetCnOk returns a tuple with the Cn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCn

`func (o *CreateGoogleLdapProfileOpenApiVO) SetCn(v string)`

SetCn sets Cn field to given value.


### GetFilter

`func (o *CreateGoogleLdapProfileOpenApiVO) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CreateGoogleLdapProfileOpenApiVO) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CreateGoogleLdapProfileOpenApiVO) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CreateGoogleLdapProfileOpenApiVO) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetName

`func (o *CreateGoogleLdapProfileOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGoogleLdapProfileOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGoogleLdapProfileOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CreateGoogleLdapProfileOpenApiVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateGoogleLdapProfileOpenApiVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateGoogleLdapProfileOpenApiVO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateGoogleLdapProfileOpenApiVO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *CreateGoogleLdapProfileOpenApiVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateGoogleLdapProfileOpenApiVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateGoogleLdapProfileOpenApiVO) SetPort(v int32)`

SetPort sets Port field to given value.


### GetServer

`func (o *CreateGoogleLdapProfileOpenApiVO) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *CreateGoogleLdapProfileOpenApiVO) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *CreateGoogleLdapProfileOpenApiVO) SetServer(v string)`

SetServer sets Server field to given value.


### GetStatus

`func (o *CreateGoogleLdapProfileOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateGoogleLdapProfileOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateGoogleLdapProfileOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


