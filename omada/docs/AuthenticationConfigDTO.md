# AuthenticationConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveStatus** | Pointer to **string** | Active status should be a value as follows:ACTIVE: ONU is in an active state, capable of data communication and supporting services.INACTIVE: ONU is in an inactive state, which may be due to reasons such as unactivated configuration, being offline, configuration failure, or configuration mismatch. | [optional] 
**AdminStatus** | Pointer to **string** | Admin status should be a value as follows:ACTIVATE,DEACTIVATE | [optional] 
**AuthenticationMethod** | Pointer to **string** | Authentication mode used by the automatically discovered ONU.AuthenticationMethod should be a value as follows:PASSWORD_AUTH,SN_AUTH,LOID_AUTH,SN_AND_PASSWORD_AUTH,LOID_AND_PASSWORD_AUTH | [optional] 
**ConfigStatus** | Pointer to **string** | Config status should be a value as follows:SUCCESS: Configuration successfully delivered and recognized by the ONU.FAILED: Configuration not successfully delivered or not recognized by the ONU. | [optional] 
**Description** | Pointer to **string** | Description of ONU.Only 1-32 bits numbers, Upper and lower letters, -@_:/. are allowed. | [optional] 
**DiscoveryMode** | Pointer to **string** | ONU auto discovery mode.DiscoveryMode should be a value as follows:ALWAYS_ON,ONCE_ON  | [optional] 
**Key** | Pointer to **string** | Identifier of entry | [optional] 
**LineProfile** | Pointer to **string** | Line profile bound to the ONU, lineProfile is displayed in the format id(name) | [optional] 
**Loid** | Pointer to **string** | Loid set by automatic discovery of the ONT, loid should be 1-32 characters, including letters, numbers, and symbols (-@_:/.). | [optional] 
**LoidPassword** | Pointer to **string** | Loid password set by automatic discovery of the ONT,loidPassword should be ASCII characters from \\x21 to \\x7e | [optional] 
**MacAddress** | Pointer to **string** | Mac address of ONU | [optional] 
**MatchStatus** | Pointer to **string** | Match status should be a value as follows:MATCH: ONU hardware capabilities are consistent with the bound service template.MISMATCH: ONU hardware capabilities are inconsistent with the bound service template. | [optional] 
**OnlineStatus** | Pointer to **string** | Online status should be a value as follows:ONLINE: ONU is online.OFFLINE: ONU is offline; it does not support version information queries or any service operations. | [optional] 
**OnuId** | Pointer to **int32** | ONU ID,This parameter is optional. If not specified, the system will automatically allocate the smallest available ONU number under the current port.OnuId should be within the range of 0 to 127 | [optional] 
**Password** | Pointer to **string** | Authentication password set by automatic discovery of the ONT, password should be ASCII characters from \\x21 to \\x7e and 1 to 20 hexadecimal digits | [optional] 
**PasswordType** | Pointer to **string** | Indicates the password type reported by the automatic discovery of the ONT.PasswordType should be a value as follows：ASCII,HEX | [optional] 
**PortId** | Pointer to **string** | Pon port ID,e.g.,PON 1/0/1 | [optional] 
**ReRegisterAuthentication** | Pointer to **string** | Re-Register-Authentication | [optional] 
**SerialNumber** | Pointer to **string** | SerialNumber of ONU should contain 12, 13, or 16 characters, in the format XXXXXXXXXXXX ,XXXX-XXXXXXXX,XXXXXXXXXXXXXXXX | [optional] 
**ServicePortProfile** | Pointer to **string** | Service port profile bound to the ONU is  in the format id(name) | [optional] 
**ServiceProfile** | Pointer to **string** | Service profile bound to the ONU, serviceProfile displayed in the format id(name) | [optional] 

## Methods

### NewAuthenticationConfigDTO

`func NewAuthenticationConfigDTO() *AuthenticationConfigDTO`

NewAuthenticationConfigDTO instantiates a new AuthenticationConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationConfigDTOWithDefaults

`func NewAuthenticationConfigDTOWithDefaults() *AuthenticationConfigDTO`

NewAuthenticationConfigDTOWithDefaults instantiates a new AuthenticationConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveStatus

`func (o *AuthenticationConfigDTO) GetActiveStatus() string`

GetActiveStatus returns the ActiveStatus field if non-nil, zero value otherwise.

### GetActiveStatusOk

`func (o *AuthenticationConfigDTO) GetActiveStatusOk() (*string, bool)`

GetActiveStatusOk returns a tuple with the ActiveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveStatus

`func (o *AuthenticationConfigDTO) SetActiveStatus(v string)`

SetActiveStatus sets ActiveStatus field to given value.

### HasActiveStatus

`func (o *AuthenticationConfigDTO) HasActiveStatus() bool`

HasActiveStatus returns a boolean if a field has been set.

### GetAdminStatus

`func (o *AuthenticationConfigDTO) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *AuthenticationConfigDTO) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *AuthenticationConfigDTO) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *AuthenticationConfigDTO) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *AuthenticationConfigDTO) GetAuthenticationMethod() string`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *AuthenticationConfigDTO) GetAuthenticationMethodOk() (*string, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *AuthenticationConfigDTO) SetAuthenticationMethod(v string)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *AuthenticationConfigDTO) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### GetConfigStatus

`func (o *AuthenticationConfigDTO) GetConfigStatus() string`

GetConfigStatus returns the ConfigStatus field if non-nil, zero value otherwise.

### GetConfigStatusOk

`func (o *AuthenticationConfigDTO) GetConfigStatusOk() (*string, bool)`

GetConfigStatusOk returns a tuple with the ConfigStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStatus

`func (o *AuthenticationConfigDTO) SetConfigStatus(v string)`

SetConfigStatus sets ConfigStatus field to given value.

### HasConfigStatus

`func (o *AuthenticationConfigDTO) HasConfigStatus() bool`

HasConfigStatus returns a boolean if a field has been set.

### GetDescription

`func (o *AuthenticationConfigDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthenticationConfigDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthenticationConfigDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthenticationConfigDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscoveryMode

`func (o *AuthenticationConfigDTO) GetDiscoveryMode() string`

GetDiscoveryMode returns the DiscoveryMode field if non-nil, zero value otherwise.

### GetDiscoveryModeOk

`func (o *AuthenticationConfigDTO) GetDiscoveryModeOk() (*string, bool)`

GetDiscoveryModeOk returns a tuple with the DiscoveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMode

`func (o *AuthenticationConfigDTO) SetDiscoveryMode(v string)`

SetDiscoveryMode sets DiscoveryMode field to given value.

### HasDiscoveryMode

`func (o *AuthenticationConfigDTO) HasDiscoveryMode() bool`

HasDiscoveryMode returns a boolean if a field has been set.

### GetKey

`func (o *AuthenticationConfigDTO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AuthenticationConfigDTO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AuthenticationConfigDTO) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AuthenticationConfigDTO) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLineProfile

`func (o *AuthenticationConfigDTO) GetLineProfile() string`

GetLineProfile returns the LineProfile field if non-nil, zero value otherwise.

### GetLineProfileOk

`func (o *AuthenticationConfigDTO) GetLineProfileOk() (*string, bool)`

GetLineProfileOk returns a tuple with the LineProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineProfile

`func (o *AuthenticationConfigDTO) SetLineProfile(v string)`

SetLineProfile sets LineProfile field to given value.

### HasLineProfile

`func (o *AuthenticationConfigDTO) HasLineProfile() bool`

HasLineProfile returns a boolean if a field has been set.

### GetLoid

`func (o *AuthenticationConfigDTO) GetLoid() string`

GetLoid returns the Loid field if non-nil, zero value otherwise.

### GetLoidOk

`func (o *AuthenticationConfigDTO) GetLoidOk() (*string, bool)`

GetLoidOk returns a tuple with the Loid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoid

`func (o *AuthenticationConfigDTO) SetLoid(v string)`

SetLoid sets Loid field to given value.

### HasLoid

`func (o *AuthenticationConfigDTO) HasLoid() bool`

HasLoid returns a boolean if a field has been set.

### GetLoidPassword

`func (o *AuthenticationConfigDTO) GetLoidPassword() string`

GetLoidPassword returns the LoidPassword field if non-nil, zero value otherwise.

### GetLoidPasswordOk

`func (o *AuthenticationConfigDTO) GetLoidPasswordOk() (*string, bool)`

GetLoidPasswordOk returns a tuple with the LoidPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoidPassword

`func (o *AuthenticationConfigDTO) SetLoidPassword(v string)`

SetLoidPassword sets LoidPassword field to given value.

### HasLoidPassword

`func (o *AuthenticationConfigDTO) HasLoidPassword() bool`

HasLoidPassword returns a boolean if a field has been set.

### GetMacAddress

`func (o *AuthenticationConfigDTO) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *AuthenticationConfigDTO) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *AuthenticationConfigDTO) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *AuthenticationConfigDTO) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMatchStatus

`func (o *AuthenticationConfigDTO) GetMatchStatus() string`

GetMatchStatus returns the MatchStatus field if non-nil, zero value otherwise.

### GetMatchStatusOk

`func (o *AuthenticationConfigDTO) GetMatchStatusOk() (*string, bool)`

GetMatchStatusOk returns a tuple with the MatchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchStatus

`func (o *AuthenticationConfigDTO) SetMatchStatus(v string)`

SetMatchStatus sets MatchStatus field to given value.

### HasMatchStatus

`func (o *AuthenticationConfigDTO) HasMatchStatus() bool`

HasMatchStatus returns a boolean if a field has been set.

### GetOnlineStatus

`func (o *AuthenticationConfigDTO) GetOnlineStatus() string`

GetOnlineStatus returns the OnlineStatus field if non-nil, zero value otherwise.

### GetOnlineStatusOk

`func (o *AuthenticationConfigDTO) GetOnlineStatusOk() (*string, bool)`

GetOnlineStatusOk returns a tuple with the OnlineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineStatus

`func (o *AuthenticationConfigDTO) SetOnlineStatus(v string)`

SetOnlineStatus sets OnlineStatus field to given value.

### HasOnlineStatus

`func (o *AuthenticationConfigDTO) HasOnlineStatus() bool`

HasOnlineStatus returns a boolean if a field has been set.

### GetOnuId

`func (o *AuthenticationConfigDTO) GetOnuId() int32`

GetOnuId returns the OnuId field if non-nil, zero value otherwise.

### GetOnuIdOk

`func (o *AuthenticationConfigDTO) GetOnuIdOk() (*int32, bool)`

GetOnuIdOk returns a tuple with the OnuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuId

`func (o *AuthenticationConfigDTO) SetOnuId(v int32)`

SetOnuId sets OnuId field to given value.

### HasOnuId

`func (o *AuthenticationConfigDTO) HasOnuId() bool`

HasOnuId returns a boolean if a field has been set.

### GetPassword

`func (o *AuthenticationConfigDTO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthenticationConfigDTO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthenticationConfigDTO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AuthenticationConfigDTO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordType

`func (o *AuthenticationConfigDTO) GetPasswordType() string`

GetPasswordType returns the PasswordType field if non-nil, zero value otherwise.

### GetPasswordTypeOk

`func (o *AuthenticationConfigDTO) GetPasswordTypeOk() (*string, bool)`

GetPasswordTypeOk returns a tuple with the PasswordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordType

`func (o *AuthenticationConfigDTO) SetPasswordType(v string)`

SetPasswordType sets PasswordType field to given value.

### HasPasswordType

`func (o *AuthenticationConfigDTO) HasPasswordType() bool`

HasPasswordType returns a boolean if a field has been set.

### GetPortId

`func (o *AuthenticationConfigDTO) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *AuthenticationConfigDTO) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *AuthenticationConfigDTO) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *AuthenticationConfigDTO) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetReRegisterAuthentication

`func (o *AuthenticationConfigDTO) GetReRegisterAuthentication() string`

GetReRegisterAuthentication returns the ReRegisterAuthentication field if non-nil, zero value otherwise.

### GetReRegisterAuthenticationOk

`func (o *AuthenticationConfigDTO) GetReRegisterAuthenticationOk() (*string, bool)`

GetReRegisterAuthenticationOk returns a tuple with the ReRegisterAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReRegisterAuthentication

`func (o *AuthenticationConfigDTO) SetReRegisterAuthentication(v string)`

SetReRegisterAuthentication sets ReRegisterAuthentication field to given value.

### HasReRegisterAuthentication

`func (o *AuthenticationConfigDTO) HasReRegisterAuthentication() bool`

HasReRegisterAuthentication returns a boolean if a field has been set.

### GetSerialNumber

`func (o *AuthenticationConfigDTO) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *AuthenticationConfigDTO) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *AuthenticationConfigDTO) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *AuthenticationConfigDTO) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetServicePortProfile

`func (o *AuthenticationConfigDTO) GetServicePortProfile() string`

GetServicePortProfile returns the ServicePortProfile field if non-nil, zero value otherwise.

### GetServicePortProfileOk

`func (o *AuthenticationConfigDTO) GetServicePortProfileOk() (*string, bool)`

GetServicePortProfileOk returns a tuple with the ServicePortProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePortProfile

`func (o *AuthenticationConfigDTO) SetServicePortProfile(v string)`

SetServicePortProfile sets ServicePortProfile field to given value.

### HasServicePortProfile

`func (o *AuthenticationConfigDTO) HasServicePortProfile() bool`

HasServicePortProfile returns a boolean if a field has been set.

### GetServiceProfile

`func (o *AuthenticationConfigDTO) GetServiceProfile() string`

GetServiceProfile returns the ServiceProfile field if non-nil, zero value otherwise.

### GetServiceProfileOk

`func (o *AuthenticationConfigDTO) GetServiceProfileOk() (*string, bool)`

GetServiceProfileOk returns a tuple with the ServiceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProfile

`func (o *AuthenticationConfigDTO) SetServiceProfile(v string)`

SetServiceProfile sets ServiceProfile field to given value.

### HasServiceProfile

`func (o *AuthenticationConfigDTO) HasServiceProfile() bool`

HasServiceProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


