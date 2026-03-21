# CreateLdapProfileOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseDn** | **string** | LDAP server base distinguish name, for example: dc&#x3D;xxx,dc&#x3D;com. BaseDn should contain 1 to 512 characters. | 
**BindType** | **int32** | Type of LDAP bind. BindType should be a value as follows: 0: Simple Mode; 1: Anonymous Mode; 2: Regular Mode | 
**Cn** | **string** | LDAP server common name, for example: cn, uid. Cn should contain 1 to 64 characters | 
**DstPort** | **int32** | LDAP server listening port. dstPort should be within the range of 0-65535. When SSL is not enabled, it is generally 389, and when SSL is enabled, it is generally 636 | 
**Filter** | Pointer to **string** | Additional filter, optional when parameter [type] is 1 or 2. For example: ou&#x3D;xxx. Filter should contain 1 to 512 characters | [optional] 
**GroupDn** | Pointer to **string** | LDAP server group distinguish name, optional when parameter [type] is 1 or 2. For example: ou&#x3D;xxx,dc&#x3D;xxx,dc&#x3D;com. GroupDn should contain 1 to 512 characters | [optional] 
**Name** | **string** | LDAP profile name. Name should contain 1 to 64 characters | 
**RegularDn** | Pointer to **string** | Regular Dn, required when parameter [type] is 2. RegularDn should contain 1 to 256 characters | [optional] 
**RegularPassword** | Pointer to **string** | Regular Password, required when parameter [type] is 2. RegularPassword should contain 1 to 256 characters | [optional] 
**Server** | **string** | LDAP server address | 
**Status** | **bool** | LDAP profile enable status | 
**UseSsl** | **bool** | LDAP server enable ssl status | 

## Methods

### NewCreateLdapProfileOpenApiVO

`func NewCreateLdapProfileOpenApiVO(baseDn string, bindType int32, cn string, dstPort int32, name string, server string, status bool, useSsl bool, ) *CreateLdapProfileOpenApiVO`

NewCreateLdapProfileOpenApiVO instantiates a new CreateLdapProfileOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLdapProfileOpenApiVOWithDefaults

`func NewCreateLdapProfileOpenApiVOWithDefaults() *CreateLdapProfileOpenApiVO`

NewCreateLdapProfileOpenApiVOWithDefaults instantiates a new CreateLdapProfileOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseDn

`func (o *CreateLdapProfileOpenApiVO) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *CreateLdapProfileOpenApiVO) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *CreateLdapProfileOpenApiVO) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.


### GetBindType

`func (o *CreateLdapProfileOpenApiVO) GetBindType() int32`

GetBindType returns the BindType field if non-nil, zero value otherwise.

### GetBindTypeOk

`func (o *CreateLdapProfileOpenApiVO) GetBindTypeOk() (*int32, bool)`

GetBindTypeOk returns a tuple with the BindType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindType

`func (o *CreateLdapProfileOpenApiVO) SetBindType(v int32)`

SetBindType sets BindType field to given value.


### GetCn

`func (o *CreateLdapProfileOpenApiVO) GetCn() string`

GetCn returns the Cn field if non-nil, zero value otherwise.

### GetCnOk

`func (o *CreateLdapProfileOpenApiVO) GetCnOk() (*string, bool)`

GetCnOk returns a tuple with the Cn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCn

`func (o *CreateLdapProfileOpenApiVO) SetCn(v string)`

SetCn sets Cn field to given value.


### GetDstPort

`func (o *CreateLdapProfileOpenApiVO) GetDstPort() int32`

GetDstPort returns the DstPort field if non-nil, zero value otherwise.

### GetDstPortOk

`func (o *CreateLdapProfileOpenApiVO) GetDstPortOk() (*int32, bool)`

GetDstPortOk returns a tuple with the DstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPort

`func (o *CreateLdapProfileOpenApiVO) SetDstPort(v int32)`

SetDstPort sets DstPort field to given value.


### GetFilter

`func (o *CreateLdapProfileOpenApiVO) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CreateLdapProfileOpenApiVO) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CreateLdapProfileOpenApiVO) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CreateLdapProfileOpenApiVO) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetGroupDn

`func (o *CreateLdapProfileOpenApiVO) GetGroupDn() string`

GetGroupDn returns the GroupDn field if non-nil, zero value otherwise.

### GetGroupDnOk

`func (o *CreateLdapProfileOpenApiVO) GetGroupDnOk() (*string, bool)`

GetGroupDnOk returns a tuple with the GroupDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDn

`func (o *CreateLdapProfileOpenApiVO) SetGroupDn(v string)`

SetGroupDn sets GroupDn field to given value.

### HasGroupDn

`func (o *CreateLdapProfileOpenApiVO) HasGroupDn() bool`

HasGroupDn returns a boolean if a field has been set.

### GetName

`func (o *CreateLdapProfileOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLdapProfileOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLdapProfileOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetRegularDn

`func (o *CreateLdapProfileOpenApiVO) GetRegularDn() string`

GetRegularDn returns the RegularDn field if non-nil, zero value otherwise.

### GetRegularDnOk

`func (o *CreateLdapProfileOpenApiVO) GetRegularDnOk() (*string, bool)`

GetRegularDnOk returns a tuple with the RegularDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularDn

`func (o *CreateLdapProfileOpenApiVO) SetRegularDn(v string)`

SetRegularDn sets RegularDn field to given value.

### HasRegularDn

`func (o *CreateLdapProfileOpenApiVO) HasRegularDn() bool`

HasRegularDn returns a boolean if a field has been set.

### GetRegularPassword

`func (o *CreateLdapProfileOpenApiVO) GetRegularPassword() string`

GetRegularPassword returns the RegularPassword field if non-nil, zero value otherwise.

### GetRegularPasswordOk

`func (o *CreateLdapProfileOpenApiVO) GetRegularPasswordOk() (*string, bool)`

GetRegularPasswordOk returns a tuple with the RegularPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularPassword

`func (o *CreateLdapProfileOpenApiVO) SetRegularPassword(v string)`

SetRegularPassword sets RegularPassword field to given value.

### HasRegularPassword

`func (o *CreateLdapProfileOpenApiVO) HasRegularPassword() bool`

HasRegularPassword returns a boolean if a field has been set.

### GetServer

`func (o *CreateLdapProfileOpenApiVO) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *CreateLdapProfileOpenApiVO) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *CreateLdapProfileOpenApiVO) SetServer(v string)`

SetServer sets Server field to given value.


### GetStatus

`func (o *CreateLdapProfileOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateLdapProfileOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateLdapProfileOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetUseSsl

`func (o *CreateLdapProfileOpenApiVO) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *CreateLdapProfileOpenApiVO) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *CreateLdapProfileOpenApiVO) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


