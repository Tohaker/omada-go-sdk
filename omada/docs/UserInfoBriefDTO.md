# UserInfoBriefDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**AllCustomer** | Pointer to **bool** |  | [optional] 
**AllSite** | Pointer to **bool** |  | [optional] 
**CustomerRoleId** | Pointer to **string** |  | [optional] 
**Customers** | Pointer to **[]string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EncryptFields** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OmadacId** | Pointer to **string** |  | [optional] 
**OpenApiRequest** | Pointer to **bool** |  | [optional] 
**RoleId** | Pointer to **string** |  | [optional] 
**RoleTypeName** | Pointer to **string** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**Sites** | Pointer to **[]string** |  | [optional] 
**TemporaryUser** | Pointer to **bool** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UserLevel** | Pointer to **int32** |  | [optional] 
**UserType** | Pointer to **int32** |  | [optional] 
**ViewerPermission** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserInfoBriefDTO

`func NewUserInfoBriefDTO() *UserInfoBriefDTO`

NewUserInfoBriefDTO instantiates a new UserInfoBriefDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInfoBriefDTOWithDefaults

`func NewUserInfoBriefDTOWithDefaults() *UserInfoBriefDTO`

NewUserInfoBriefDTOWithDefaults instantiates a new UserInfoBriefDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *UserInfoBriefDTO) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *UserInfoBriefDTO) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *UserInfoBriefDTO) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *UserInfoBriefDTO) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAccountId

`func (o *UserInfoBriefDTO) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserInfoBriefDTO) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserInfoBriefDTO) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserInfoBriefDTO) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAllCustomer

`func (o *UserInfoBriefDTO) GetAllCustomer() bool`

GetAllCustomer returns the AllCustomer field if non-nil, zero value otherwise.

### GetAllCustomerOk

`func (o *UserInfoBriefDTO) GetAllCustomerOk() (*bool, bool)`

GetAllCustomerOk returns a tuple with the AllCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllCustomer

`func (o *UserInfoBriefDTO) SetAllCustomer(v bool)`

SetAllCustomer sets AllCustomer field to given value.

### HasAllCustomer

`func (o *UserInfoBriefDTO) HasAllCustomer() bool`

HasAllCustomer returns a boolean if a field has been set.

### GetAllSite

`func (o *UserInfoBriefDTO) GetAllSite() bool`

GetAllSite returns the AllSite field if non-nil, zero value otherwise.

### GetAllSiteOk

`func (o *UserInfoBriefDTO) GetAllSiteOk() (*bool, bool)`

GetAllSiteOk returns a tuple with the AllSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSite

`func (o *UserInfoBriefDTO) SetAllSite(v bool)`

SetAllSite sets AllSite field to given value.

### HasAllSite

`func (o *UserInfoBriefDTO) HasAllSite() bool`

HasAllSite returns a boolean if a field has been set.

### GetCustomerRoleId

`func (o *UserInfoBriefDTO) GetCustomerRoleId() string`

GetCustomerRoleId returns the CustomerRoleId field if non-nil, zero value otherwise.

### GetCustomerRoleIdOk

`func (o *UserInfoBriefDTO) GetCustomerRoleIdOk() (*string, bool)`

GetCustomerRoleIdOk returns a tuple with the CustomerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoleId

`func (o *UserInfoBriefDTO) SetCustomerRoleId(v string)`

SetCustomerRoleId sets CustomerRoleId field to given value.

### HasCustomerRoleId

`func (o *UserInfoBriefDTO) HasCustomerRoleId() bool`

HasCustomerRoleId returns a boolean if a field has been set.

### GetCustomers

`func (o *UserInfoBriefDTO) GetCustomers() []string`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *UserInfoBriefDTO) GetCustomersOk() (*[]string, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *UserInfoBriefDTO) SetCustomers(v []string)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *UserInfoBriefDTO) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.

### GetEmail

`func (o *UserInfoBriefDTO) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserInfoBriefDTO) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserInfoBriefDTO) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserInfoBriefDTO) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEncryptFields

`func (o *UserInfoBriefDTO) GetEncryptFields() []string`

GetEncryptFields returns the EncryptFields field if non-nil, zero value otherwise.

### GetEncryptFieldsOk

`func (o *UserInfoBriefDTO) GetEncryptFieldsOk() (*[]string, bool)`

GetEncryptFieldsOk returns a tuple with the EncryptFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptFields

`func (o *UserInfoBriefDTO) SetEncryptFields(v []string)`

SetEncryptFields sets EncryptFields field to given value.

### HasEncryptFields

`func (o *UserInfoBriefDTO) HasEncryptFields() bool`

HasEncryptFields returns a boolean if a field has been set.

### GetId

`func (o *UserInfoBriefDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserInfoBriefDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserInfoBriefDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserInfoBriefDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UserInfoBriefDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserInfoBriefDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserInfoBriefDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserInfoBriefDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOmadacId

`func (o *UserInfoBriefDTO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *UserInfoBriefDTO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *UserInfoBriefDTO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *UserInfoBriefDTO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetOpenApiRequest

`func (o *UserInfoBriefDTO) GetOpenApiRequest() bool`

GetOpenApiRequest returns the OpenApiRequest field if non-nil, zero value otherwise.

### GetOpenApiRequestOk

`func (o *UserInfoBriefDTO) GetOpenApiRequestOk() (*bool, bool)`

GetOpenApiRequestOk returns a tuple with the OpenApiRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenApiRequest

`func (o *UserInfoBriefDTO) SetOpenApiRequest(v bool)`

SetOpenApiRequest sets OpenApiRequest field to given value.

### HasOpenApiRequest

`func (o *UserInfoBriefDTO) HasOpenApiRequest() bool`

HasOpenApiRequest returns a boolean if a field has been set.

### GetRoleId

`func (o *UserInfoBriefDTO) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *UserInfoBriefDTO) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *UserInfoBriefDTO) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *UserInfoBriefDTO) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetRoleTypeName

`func (o *UserInfoBriefDTO) GetRoleTypeName() string`

GetRoleTypeName returns the RoleTypeName field if non-nil, zero value otherwise.

### GetRoleTypeNameOk

`func (o *UserInfoBriefDTO) GetRoleTypeNameOk() (*string, bool)`

GetRoleTypeNameOk returns a tuple with the RoleTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleTypeName

`func (o *UserInfoBriefDTO) SetRoleTypeName(v string)`

SetRoleTypeName sets RoleTypeName field to given value.

### HasRoleTypeName

`func (o *UserInfoBriefDTO) HasRoleTypeName() bool`

HasRoleTypeName returns a boolean if a field has been set.

### GetSessionId

`func (o *UserInfoBriefDTO) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *UserInfoBriefDTO) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *UserInfoBriefDTO) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *UserInfoBriefDTO) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSites

`func (o *UserInfoBriefDTO) GetSites() []string`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *UserInfoBriefDTO) GetSitesOk() (*[]string, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *UserInfoBriefDTO) SetSites(v []string)`

SetSites sets Sites field to given value.

### HasSites

`func (o *UserInfoBriefDTO) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetTemporaryUser

`func (o *UserInfoBriefDTO) GetTemporaryUser() bool`

GetTemporaryUser returns the TemporaryUser field if non-nil, zero value otherwise.

### GetTemporaryUserOk

`func (o *UserInfoBriefDTO) GetTemporaryUserOk() (*bool, bool)`

GetTemporaryUserOk returns a tuple with the TemporaryUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryUser

`func (o *UserInfoBriefDTO) SetTemporaryUser(v bool)`

SetTemporaryUser sets TemporaryUser field to given value.

### HasTemporaryUser

`func (o *UserInfoBriefDTO) HasTemporaryUser() bool`

HasTemporaryUser returns a boolean if a field has been set.

### GetTenantId

`func (o *UserInfoBriefDTO) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UserInfoBriefDTO) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UserInfoBriefDTO) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *UserInfoBriefDTO) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUserLevel

`func (o *UserInfoBriefDTO) GetUserLevel() int32`

GetUserLevel returns the UserLevel field if non-nil, zero value otherwise.

### GetUserLevelOk

`func (o *UserInfoBriefDTO) GetUserLevelOk() (*int32, bool)`

GetUserLevelOk returns a tuple with the UserLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLevel

`func (o *UserInfoBriefDTO) SetUserLevel(v int32)`

SetUserLevel sets UserLevel field to given value.

### HasUserLevel

`func (o *UserInfoBriefDTO) HasUserLevel() bool`

HasUserLevel returns a boolean if a field has been set.

### GetUserType

`func (o *UserInfoBriefDTO) GetUserType() int32`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UserInfoBriefDTO) GetUserTypeOk() (*int32, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UserInfoBriefDTO) SetUserType(v int32)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UserInfoBriefDTO) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetViewerPermission

`func (o *UserInfoBriefDTO) GetViewerPermission() bool`

GetViewerPermission returns the ViewerPermission field if non-nil, zero value otherwise.

### GetViewerPermissionOk

`func (o *UserInfoBriefDTO) GetViewerPermissionOk() (*bool, bool)`

GetViewerPermissionOk returns a tuple with the ViewerPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerPermission

`func (o *UserInfoBriefDTO) SetViewerPermission(v bool)`

SetViewerPermission sets ViewerPermission field to given value.

### HasViewerPermission

`func (o *UserInfoBriefDTO) HasViewerPermission() bool`

HasViewerPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


