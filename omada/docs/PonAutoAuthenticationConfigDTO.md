# PonAutoAuthenticationConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationMethod** | Pointer to **string** | Authentication mode used by the automatically discovered ONU.AuthenticationMethod should be a value as follows:PASSWORD_AUTH,SN_AUTH,LOID_AUTH,SN_AND_PASSWORD_AUTH,LOID_AND_PASSWORD_AUTH | [optional] 
**AutoAuthenticationStatus** | Pointer to **string** | Auto authentication status should be a value as follows:ENABLE,DISABLE | [optional] 
**DiscoveryMode** | Pointer to **string** | ONU auto discovery mode.DiscoveryMode should be a value as follows:ALWAYS_ON,ONCE_ON  | [optional] 
**OnuMatchMode** | Pointer to **string** | ONU match mode of automatic authentication.OnuMatchMode should be a value as follows:ALL_ONU,EQUID_AUTH,EQUID_SWVER_AUTH,VENDOR_AUTH | [optional] 
**PonPort** | **string** | Pon port.e.g.GPON 1/0/1 | 
**ReRegisterAuthentication** | Pointer to **string** | Re-Register-Authentication | [optional] 

## Methods

### NewPonAutoAuthenticationConfigDTO

`func NewPonAutoAuthenticationConfigDTO(ponPort string, ) *PonAutoAuthenticationConfigDTO`

NewPonAutoAuthenticationConfigDTO instantiates a new PonAutoAuthenticationConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPonAutoAuthenticationConfigDTOWithDefaults

`func NewPonAutoAuthenticationConfigDTOWithDefaults() *PonAutoAuthenticationConfigDTO`

NewPonAutoAuthenticationConfigDTOWithDefaults instantiates a new PonAutoAuthenticationConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationMethod

`func (o *PonAutoAuthenticationConfigDTO) GetAuthenticationMethod() string`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *PonAutoAuthenticationConfigDTO) GetAuthenticationMethodOk() (*string, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *PonAutoAuthenticationConfigDTO) SetAuthenticationMethod(v string)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *PonAutoAuthenticationConfigDTO) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### GetAutoAuthenticationStatus

`func (o *PonAutoAuthenticationConfigDTO) GetAutoAuthenticationStatus() string`

GetAutoAuthenticationStatus returns the AutoAuthenticationStatus field if non-nil, zero value otherwise.

### GetAutoAuthenticationStatusOk

`func (o *PonAutoAuthenticationConfigDTO) GetAutoAuthenticationStatusOk() (*string, bool)`

GetAutoAuthenticationStatusOk returns a tuple with the AutoAuthenticationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAuthenticationStatus

`func (o *PonAutoAuthenticationConfigDTO) SetAutoAuthenticationStatus(v string)`

SetAutoAuthenticationStatus sets AutoAuthenticationStatus field to given value.

### HasAutoAuthenticationStatus

`func (o *PonAutoAuthenticationConfigDTO) HasAutoAuthenticationStatus() bool`

HasAutoAuthenticationStatus returns a boolean if a field has been set.

### GetDiscoveryMode

`func (o *PonAutoAuthenticationConfigDTO) GetDiscoveryMode() string`

GetDiscoveryMode returns the DiscoveryMode field if non-nil, zero value otherwise.

### GetDiscoveryModeOk

`func (o *PonAutoAuthenticationConfigDTO) GetDiscoveryModeOk() (*string, bool)`

GetDiscoveryModeOk returns a tuple with the DiscoveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMode

`func (o *PonAutoAuthenticationConfigDTO) SetDiscoveryMode(v string)`

SetDiscoveryMode sets DiscoveryMode field to given value.

### HasDiscoveryMode

`func (o *PonAutoAuthenticationConfigDTO) HasDiscoveryMode() bool`

HasDiscoveryMode returns a boolean if a field has been set.

### GetOnuMatchMode

`func (o *PonAutoAuthenticationConfigDTO) GetOnuMatchMode() string`

GetOnuMatchMode returns the OnuMatchMode field if non-nil, zero value otherwise.

### GetOnuMatchModeOk

`func (o *PonAutoAuthenticationConfigDTO) GetOnuMatchModeOk() (*string, bool)`

GetOnuMatchModeOk returns a tuple with the OnuMatchMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuMatchMode

`func (o *PonAutoAuthenticationConfigDTO) SetOnuMatchMode(v string)`

SetOnuMatchMode sets OnuMatchMode field to given value.

### HasOnuMatchMode

`func (o *PonAutoAuthenticationConfigDTO) HasOnuMatchMode() bool`

HasOnuMatchMode returns a boolean if a field has been set.

### GetPonPort

`func (o *PonAutoAuthenticationConfigDTO) GetPonPort() string`

GetPonPort returns the PonPort field if non-nil, zero value otherwise.

### GetPonPortOk

`func (o *PonAutoAuthenticationConfigDTO) GetPonPortOk() (*string, bool)`

GetPonPortOk returns a tuple with the PonPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonPort

`func (o *PonAutoAuthenticationConfigDTO) SetPonPort(v string)`

SetPonPort sets PonPort field to given value.


### GetReRegisterAuthentication

`func (o *PonAutoAuthenticationConfigDTO) GetReRegisterAuthentication() string`

GetReRegisterAuthentication returns the ReRegisterAuthentication field if non-nil, zero value otherwise.

### GetReRegisterAuthenticationOk

`func (o *PonAutoAuthenticationConfigDTO) GetReRegisterAuthenticationOk() (*string, bool)`

GetReRegisterAuthenticationOk returns a tuple with the ReRegisterAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReRegisterAuthentication

`func (o *PonAutoAuthenticationConfigDTO) SetReRegisterAuthentication(v string)`

SetReRegisterAuthentication sets ReRegisterAuthentication field to given value.

### HasReRegisterAuthentication

`func (o *PonAutoAuthenticationConfigDTO) HasReRegisterAuthentication() bool`

HasReRegisterAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


