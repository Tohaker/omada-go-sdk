# AuthInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to **int32** | Auth type should be a value as follows: 0: No Auth; 1: Simple Password; 2: External Radius; 3: Voucher; 4: External Portal Server; 5: Local User; 6: SMS; 7: Facebook; 8: Hotspot Radius; 9: Mac Auth (with fail over); 10: Admin auth; 12: Form auth; 15: LDAP | [optional] 
**Info** | Pointer to **string** | Auth brief information | [optional] 

## Methods

### NewAuthInfoOpenApiVO

`func NewAuthInfoOpenApiVO() *AuthInfoOpenApiVO`

NewAuthInfoOpenApiVO instantiates a new AuthInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthInfoOpenApiVOWithDefaults

`func NewAuthInfoOpenApiVOWithDefaults() *AuthInfoOpenApiVO`

NewAuthInfoOpenApiVOWithDefaults instantiates a new AuthInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *AuthInfoOpenApiVO) GetAuthType() int32`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *AuthInfoOpenApiVO) GetAuthTypeOk() (*int32, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *AuthInfoOpenApiVO) SetAuthType(v int32)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *AuthInfoOpenApiVO) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetInfo

`func (o *AuthInfoOpenApiVO) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AuthInfoOpenApiVO) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AuthInfoOpenApiVO) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AuthInfoOpenApiVO) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


