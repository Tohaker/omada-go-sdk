# AuthenticationConfigInfoDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminStatus** | Pointer to **string** | Admin status of ONU.AdminStatus should be a value as follows:ACTIVE,DEACTIVE | [optional] 
**AuthenticationMethod** | **string** | Authentication mode used by the automatically discovered ONU.AuthenticationMethod should be a value as follows:PASSWORD_AUTH,SN_AUTH,LOID_AUTH,SN_AND_PASSWORD_AUTH,LOID_AND_PASSWORD_AUTH | 
**ConfigStatus** | Pointer to **string** | Whether OLT has successfully delivered configurations to ONU.ConfigStatus should be a value as follows:FAILED,SUCCESS | [optional] 
**Description** | Pointer to **string** | Display the configured port description.Only 1-32 bits numbers, Upper and lower letters, -@_:/. are allowed. | [optional] 
**DiscoveryMode** | Pointer to **string** | ONU auto discovery mode.DiscoveryMode should be a value as follows:ALWAYS_ON,ONCE_ON  | [optional] 
**LineProfile** | **string** | Line profile id bound to the ONU | 
**Loid** | Pointer to **string** | Loid set by automatic discovery of the ONT, loid should be 1-32 characters, including letters, numbers, and symbols (-@_:/.). | [optional] 
**LoidPassword** | Pointer to **string** | Loid password set by automatic discovery of the ONT,loidPassword should be ASCII characters from \\x21 to \\x7e | [optional] 
**OnlineStatus** | Pointer to **string** | Online status should be a value as follows:DISABLE,ENABLE | [optional] 
**Password** | Pointer to **string** | Automatic discovery ONT Authentication Password.Password should contain 1 to 10 characters in ASCII (from \\x21 to \\x7e) or 1 to 20 characters in hexadecimal. | [optional] 
**PasswordType** | Pointer to **string** | Indicates the password type reported by the automatic discovery of the ONT.PasswordType should be a value as follows：ASCII,HEX | [optional] 
**ReRegisterAuthentication** | Pointer to **string** | Re-Register-Authentication | [optional] 
**SerialNumber** | Pointer to **string** | SerialNumber of ONU, serialNumber should contain 12, 13, or 16 characters, in the format XXXXXXXXXXXX,XXXX-XXXXXXXX,XXXXXXXXXXXXXXXX.When the serialNumber contains 16 characters, the first 8 characters will be grouped into pairs of two, resulting in 4 hexadecimal characters, and each hexadecimal character must fall within the range of 0x21 to 0x7E. | [optional] 
**ServicePortProfile** | Pointer to **string** | Service port profile id bound to the ONU | [optional] 
**ServiceProfile** | **string** | Service profile id bound to the ONU | 

## Methods

### NewAuthenticationConfigInfoDTO

`func NewAuthenticationConfigInfoDTO(authenticationMethod string, lineProfile string, serviceProfile string, ) *AuthenticationConfigInfoDTO`

NewAuthenticationConfigInfoDTO instantiates a new AuthenticationConfigInfoDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationConfigInfoDTOWithDefaults

`func NewAuthenticationConfigInfoDTOWithDefaults() *AuthenticationConfigInfoDTO`

NewAuthenticationConfigInfoDTOWithDefaults instantiates a new AuthenticationConfigInfoDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminStatus

`func (o *AuthenticationConfigInfoDTO) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *AuthenticationConfigInfoDTO) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *AuthenticationConfigInfoDTO) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *AuthenticationConfigInfoDTO) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *AuthenticationConfigInfoDTO) GetAuthenticationMethod() string`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *AuthenticationConfigInfoDTO) GetAuthenticationMethodOk() (*string, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *AuthenticationConfigInfoDTO) SetAuthenticationMethod(v string)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.


### GetConfigStatus

`func (o *AuthenticationConfigInfoDTO) GetConfigStatus() string`

GetConfigStatus returns the ConfigStatus field if non-nil, zero value otherwise.

### GetConfigStatusOk

`func (o *AuthenticationConfigInfoDTO) GetConfigStatusOk() (*string, bool)`

GetConfigStatusOk returns a tuple with the ConfigStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStatus

`func (o *AuthenticationConfigInfoDTO) SetConfigStatus(v string)`

SetConfigStatus sets ConfigStatus field to given value.

### HasConfigStatus

`func (o *AuthenticationConfigInfoDTO) HasConfigStatus() bool`

HasConfigStatus returns a boolean if a field has been set.

### GetDescription

`func (o *AuthenticationConfigInfoDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthenticationConfigInfoDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthenticationConfigInfoDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthenticationConfigInfoDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscoveryMode

`func (o *AuthenticationConfigInfoDTO) GetDiscoveryMode() string`

GetDiscoveryMode returns the DiscoveryMode field if non-nil, zero value otherwise.

### GetDiscoveryModeOk

`func (o *AuthenticationConfigInfoDTO) GetDiscoveryModeOk() (*string, bool)`

GetDiscoveryModeOk returns a tuple with the DiscoveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMode

`func (o *AuthenticationConfigInfoDTO) SetDiscoveryMode(v string)`

SetDiscoveryMode sets DiscoveryMode field to given value.

### HasDiscoveryMode

`func (o *AuthenticationConfigInfoDTO) HasDiscoveryMode() bool`

HasDiscoveryMode returns a boolean if a field has been set.

### GetLineProfile

`func (o *AuthenticationConfigInfoDTO) GetLineProfile() string`

GetLineProfile returns the LineProfile field if non-nil, zero value otherwise.

### GetLineProfileOk

`func (o *AuthenticationConfigInfoDTO) GetLineProfileOk() (*string, bool)`

GetLineProfileOk returns a tuple with the LineProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineProfile

`func (o *AuthenticationConfigInfoDTO) SetLineProfile(v string)`

SetLineProfile sets LineProfile field to given value.


### GetLoid

`func (o *AuthenticationConfigInfoDTO) GetLoid() string`

GetLoid returns the Loid field if non-nil, zero value otherwise.

### GetLoidOk

`func (o *AuthenticationConfigInfoDTO) GetLoidOk() (*string, bool)`

GetLoidOk returns a tuple with the Loid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoid

`func (o *AuthenticationConfigInfoDTO) SetLoid(v string)`

SetLoid sets Loid field to given value.

### HasLoid

`func (o *AuthenticationConfigInfoDTO) HasLoid() bool`

HasLoid returns a boolean if a field has been set.

### GetLoidPassword

`func (o *AuthenticationConfigInfoDTO) GetLoidPassword() string`

GetLoidPassword returns the LoidPassword field if non-nil, zero value otherwise.

### GetLoidPasswordOk

`func (o *AuthenticationConfigInfoDTO) GetLoidPasswordOk() (*string, bool)`

GetLoidPasswordOk returns a tuple with the LoidPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoidPassword

`func (o *AuthenticationConfigInfoDTO) SetLoidPassword(v string)`

SetLoidPassword sets LoidPassword field to given value.

### HasLoidPassword

`func (o *AuthenticationConfigInfoDTO) HasLoidPassword() bool`

HasLoidPassword returns a boolean if a field has been set.

### GetOnlineStatus

`func (o *AuthenticationConfigInfoDTO) GetOnlineStatus() string`

GetOnlineStatus returns the OnlineStatus field if non-nil, zero value otherwise.

### GetOnlineStatusOk

`func (o *AuthenticationConfigInfoDTO) GetOnlineStatusOk() (*string, bool)`

GetOnlineStatusOk returns a tuple with the OnlineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineStatus

`func (o *AuthenticationConfigInfoDTO) SetOnlineStatus(v string)`

SetOnlineStatus sets OnlineStatus field to given value.

### HasOnlineStatus

`func (o *AuthenticationConfigInfoDTO) HasOnlineStatus() bool`

HasOnlineStatus returns a boolean if a field has been set.

### GetPassword

`func (o *AuthenticationConfigInfoDTO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthenticationConfigInfoDTO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthenticationConfigInfoDTO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AuthenticationConfigInfoDTO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordType

`func (o *AuthenticationConfigInfoDTO) GetPasswordType() string`

GetPasswordType returns the PasswordType field if non-nil, zero value otherwise.

### GetPasswordTypeOk

`func (o *AuthenticationConfigInfoDTO) GetPasswordTypeOk() (*string, bool)`

GetPasswordTypeOk returns a tuple with the PasswordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordType

`func (o *AuthenticationConfigInfoDTO) SetPasswordType(v string)`

SetPasswordType sets PasswordType field to given value.

### HasPasswordType

`func (o *AuthenticationConfigInfoDTO) HasPasswordType() bool`

HasPasswordType returns a boolean if a field has been set.

### GetReRegisterAuthentication

`func (o *AuthenticationConfigInfoDTO) GetReRegisterAuthentication() string`

GetReRegisterAuthentication returns the ReRegisterAuthentication field if non-nil, zero value otherwise.

### GetReRegisterAuthenticationOk

`func (o *AuthenticationConfigInfoDTO) GetReRegisterAuthenticationOk() (*string, bool)`

GetReRegisterAuthenticationOk returns a tuple with the ReRegisterAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReRegisterAuthentication

`func (o *AuthenticationConfigInfoDTO) SetReRegisterAuthentication(v string)`

SetReRegisterAuthentication sets ReRegisterAuthentication field to given value.

### HasReRegisterAuthentication

`func (o *AuthenticationConfigInfoDTO) HasReRegisterAuthentication() bool`

HasReRegisterAuthentication returns a boolean if a field has been set.

### GetSerialNumber

`func (o *AuthenticationConfigInfoDTO) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *AuthenticationConfigInfoDTO) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *AuthenticationConfigInfoDTO) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *AuthenticationConfigInfoDTO) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetServicePortProfile

`func (o *AuthenticationConfigInfoDTO) GetServicePortProfile() string`

GetServicePortProfile returns the ServicePortProfile field if non-nil, zero value otherwise.

### GetServicePortProfileOk

`func (o *AuthenticationConfigInfoDTO) GetServicePortProfileOk() (*string, bool)`

GetServicePortProfileOk returns a tuple with the ServicePortProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePortProfile

`func (o *AuthenticationConfigInfoDTO) SetServicePortProfile(v string)`

SetServicePortProfile sets ServicePortProfile field to given value.

### HasServicePortProfile

`func (o *AuthenticationConfigInfoDTO) HasServicePortProfile() bool`

HasServicePortProfile returns a boolean if a field has been set.

### GetServiceProfile

`func (o *AuthenticationConfigInfoDTO) GetServiceProfile() string`

GetServiceProfile returns the ServiceProfile field if non-nil, zero value otherwise.

### GetServiceProfileOk

`func (o *AuthenticationConfigInfoDTO) GetServiceProfileOk() (*string, bool)`

GetServiceProfileOk returns a tuple with the ServiceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProfile

`func (o *AuthenticationConfigInfoDTO) SetServiceProfile(v string)`

SetServiceProfile sets ServiceProfile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


